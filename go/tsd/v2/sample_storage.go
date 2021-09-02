// Copyright 2020 Eryx <evorui аt gmail dοt com>, All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tsd

import (
	"hash/crc32"
	"time"

	"github.com/hooto/hlog4g/hlog"
)

var (
	storageFlushPeriod  = int64(600)
	storageFlushRefresh = storageFlushPeriod / 10
)

type MetricStorage interface {
	Put(key, value []byte) error
	Scan(keyOffset, keyCutset []byte) ([][]byte, error)
}

func storageKeyEncode(instanceId string, t int64) []byte {
	return []byte("/valuedig/apis/tsd/v2/" + instanceId + "/" + uint32ToHexString(uint32(t)))
}

func (it *SampleSet) StorageSetup(instanceId string, drv MetricStorage) error {

	it.mu.Lock()
	defer it.mu.Unlock()

	if it.families == nil {
		it.families = map[string]*SampleFamily{}
	}

	it.instanceId = instanceId

	if len(it.flushes) == 0 {

		var (
			tn          = timesec()
			keyOffset   = storageKeyEncode(instanceId, tn-4200)
			keyCutset   = storageKeyEncode(instanceId, tn+600)
			sm          *SampleMetric
			statsBytes  = 0
			statsPoint  = 0
			statsMetric = 0
		)

		if rs, err := drv.Scan(keyOffset, keyCutset); err == nil {

			for _, bs := range rs {

				var set MetricStorageSet

				if err := StdProto.Decode(bs, &set); err != nil {
					continue
				}

				it.flushes[set.TimeBucket] = crc32.ChecksumIEEE(bs)

				for _, m := range set.Metrics {

					mf, ok := it.families[m.Name]
					if !ok {
						mf = &SampleFamily{
							name: m.Name,
						}
						it.families[m.Name] = mf
					}

					if len(m.Labels) > 0 {
						sm = mf.getEntry(m.Labels[0].Name, m.Labels[0].Value)
					} else {
						sm = mf.getEntry("", "")
					}

					for _, p := range m.Points {
						if p2 := sm.point(p.Time); p2 != nil {
							p2.count = p.Count
							p2.sum = p.Sum
							statsPoint++
						}
					}

					statsMetric++
				}
				statsBytes += len(bs)
			}
		}

		hlog.Printf("info", "metric/storage reload metrics %d, points %d, bytes %d",
			statsMetric, statsPoint, statsBytes)
	}

	it.storageClient = drv

	return nil
}

func (it *SampleSet) Close() error {
	it.close = true
	it.events <- nil
	return nil
}

func (it *SampleSet) Flush(force bool) error {
	it.mu.Lock()
	defer it.mu.Unlock()
	return it.flush(force)
}

func sampleStorageTimeBucketAlign(l, r int64) (int64, int64) {

	lt := time.Unix(l, 0)
	rt := time.Unix(r, 0)

	if f := (int64(lt.Second()) + int64(lt.Minute()*60)) % storageFlushPeriod; f > 0 {
		l -= f
	}

	if f := (int64(rt.Second()) + int64(rt.Minute()*60)) % storageFlushPeriod; f > 0 {
		r -= f
		r += storageFlushPeriod
	}

	return l, r
}

func (it *SampleSet) flush(force bool) error {

	if it.storageClient == nil {
		return nil
	}

	var (
		tn      = time.Now()
		endTime = tn.Unix()
		fixTime = tn.Unix()
	)

	if !force && (it.storageFlushed+storageFlushRefresh) >= endTime {
		return nil
	}

	if f := int64(tn.Second()) % storageFlushPeriod; f > 0 {
		fixTime -= f
	}

	if f := int64(tn.Minute()*60) % storageFlushPeriod; f > 0 {
		fixTime -= f
	}

	if fixTime != endTime {
		endTime = fixTime + storageFlushPeriod
	}

	for t := endTime - (storageFlushPeriod * 2); t <= endTime; t += storageFlushPeriod {

		ms := &MetricStorageSet{
			InstanceId:      it.instanceId,
			TimeBucket:      t,
			AlignmentPeriod: storageFlushPeriod,
		}

		for name, family := range it.families {
			for _, m := range family.metrics {

				m2 := &Metric{
					Name: name,
				}
				if m.labelName != "" {
					m2.Labels = append(m2.Labels, &MetricLabel{
						Name:  m.labelName,
						Value: m.labelValue,
					})
				}
				for _, p := range m.points {

					if p.time <= (t - storageFlushPeriod) {
						continue
					}
					if p.time > t {
						break
					}
					m2.Points = append(m2.Points, &MetricPoint{
						Count: p.count,
						Sum:   p.sum,
						Time:  p.time,
					})
				}
				if len(m2.Points) > 0 {
					ms.Metrics = append(ms.Metrics, m2)
				}
			}
		}

		if len(ms.Metrics) == 0 {
			continue
		}

		ms.Metrics = metricsSort(ms.Metrics)

		bs, err := StdProto.Encode(ms)
		if err != nil {
			continue
		}

		psum, ok := it.flushes[t]
		csum := crc32.ChecksumIEEE(bs)
		if ok && psum == csum {
			continue
		}

		if err := it.storageClient.Put(storageKeyEncode(it.instanceId, t), bs); err == nil {
			hlog.Printf("debug", "instance %s, time bucket %d, flush %d bytes, sumcheck %d",
				it.instanceId, t, len(bs), csum)
			it.flushes[t] = csum
		} else {
			hlog.Printf("warn", "instance %s, time bucket %d, flush %d bytes, err %s",
				it.instanceId, t, len(bs), err.Error())
		}
	}

	dels := []int64{}
	for t, _ := range it.flushes {
		if (t + 7200) < tn.Unix() {
			dels = append(dels, t)
		}
	}

	for _, t := range dels {
		delete(it.flushes, t)
	}

	it.storageFlushed = tn.Unix()

	return nil
}
