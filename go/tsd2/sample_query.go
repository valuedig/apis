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
	"strings"
	"sync"
	"time"

	"github.com/hooto/hlog4g/hlog"
)

var (
	mu sync.RWMutex
)

func (it *SampleQueryRequest) Select(metricName string, labels ...string) *SampleQueryRequest_Metric {
	m := it.MetricSelect(metricName)
	if len(labels) > 0 {
		m.LabelSelect(labels...)
	}
	return m
}

func (it *SampleQueryRequest) MetricSelect(name string) *SampleQueryRequest_Metric {
	if !LabelNameRX.MatchString(name) {
		return &SampleQueryRequest_Metric{}
	}
	if it.Metrics == nil {
		it.Metrics = map[string]*SampleQueryRequest_Metric{}
	}
	req, ok := it.Metrics[name]
	if !ok {
		req = &SampleQueryRequest_Metric{
			Type: MetricType_UNTYPED,
		}
		it.Metrics[name] = req
	}
	return req
}

func (it *SampleQueryRequest_Metric) LabelSelect(labels ...string) *SampleQueryRequest_Metric {
	if it.Labels == nil {
		it.Labels = map[string]string{}
	}
	for _, name := range labels {
		if name != "*" && !LabelNameRX.MatchString(name) {
			continue
		}
		it.Labels[name] = "*"
	}
	return it
}

func (it *SampleQueryRequest_Metric) labelMatch(name, value string) bool {

	if it.Labels != nil {
		for k, v := range it.Labels {
			if k == "*" {
				return true
			}
			if name != k {
				continue
			}
			if v == "" || v == "*" || v == value {
				return true
			}
			if v[len(v)-1] == '*' &&
				strings.HasPrefix(value, v[:len(v)-1]) {
				return true
			}
		}
	}

	return false
}

func (it *SampleQueryRequest_Metric) Gauge() *SampleQueryRequest_Metric {
	it.Type = MetricType_GAUGE
	return it
}

func (it *SampleQueryRequest_Metric) Delta() *SampleQueryRequest_Metric {
	it.Type = MetricType_DELTA
	return it
}

func NewSampleQueryRequest() *SampleQueryRequest {
	return &SampleQueryRequest{
		Metrics: map[string]*SampleQueryRequest_Metric{},
	}
}

func (req *SampleQueryRequest) reset() {

	apFilter := func(v int64) int64 {
		var v2 int64
		for _, v2 = range []int64{
			1, 2, 3, 5, 10, 15, 30, 60, 120, 300, 600, 900, 1200, 1800, 3600,
		} {
			if v2 >= v {
				break
			}
		}
		return v2
	}
	req.AlignmentPeriod = apFilter(req.AlignmentPeriod)

	if req.LastTimeRange > 0 {
		req.EndTime = timesec()
		req.StartTime = req.EndTime - req.LastTimeRange
	}

	if req.EndTime < req.StartTime {
		req.EndTime = req.StartTime
	}

	{
		ust := time.Unix(req.StartTime, 0)
		if f := int64(ust.Second()) % req.AlignmentPeriod; f > 0 {
			req.StartTime -= f
		}
		if f := int64(ust.Minute()*60) % req.AlignmentPeriod; f > 0 {
			req.StartTime -= f
		}
	}

	{
		uet := time.Unix(req.EndTime, 0)
		fix := req.EndTime
		if f := int64(uet.Second()) % req.AlignmentPeriod; f > 0 {
			fix -= f
		}
		if f := int64(uet.Minute()*60) % req.AlignmentPeriod; f > 0 {
			fix -= f
		}
		if fix < req.EndTime {
			req.EndTime = fix + req.AlignmentPeriod
		}
	}

	if req.Metrics == nil {
		req.Metrics = map[string]*SampleQueryRequest_Metric{}
	}
}

func (it *SampleSet) Query(req *SampleQueryRequest) (*MetricSet, error) {

	req.reset()

	ms := &MetricSet{
		Status: &MetricSet_Status{
			AlignmentPeriod: req.AlignmentPeriod,
		},
	}

	if len(req.Metrics) == 0 {
		return ms, nil
	}

	var (
		tn = timesec()
	)

	for st := req.StartTime; st <= req.EndTime; st += req.AlignmentPeriod {
		if (st + 1) >= tn {
			break
		}
		ms.TimeBuckets = append(ms.TimeBuckets, st+req.AlignmentPeriod)
	}

	if req.LastTimeRange > req.AlignmentPeriod {
		n := int(req.LastTimeRange / req.AlignmentPeriod)
		if n > 1 && n < len(ms.TimeBuckets) {
			ms.TimeBuckets = ms.TimeBuckets[1:]
		}
	}

	if len(ms.TimeBuckets) < 1 {
		return ms, nil
	}

	if ms.TimeBuckets[0] < (tn-it.cleanInterval) && it.storageClient != nil {
		return it.dbQuery(ms, req)
	}

	return it.memQuery(ms, req)
}

func (it *SampleSet) memQuery(ms *MetricSet, req *SampleQueryRequest) (*MetricSet, error) {

	it.mu.RLock()
	defer it.mu.RUnlock()

	var (
		reqMetric *SampleQueryRequest_Metric
		ok        bool
	)

	for name, family := range it.families {

		reqMetric, ok = req.Metrics[name]
		if !ok {
			continue
		}

		for _, v := range family.metrics {

			if !reqMetric.labelMatch(v.labelName, v.labelValue) {
				continue
			}

			m := &Metric{
				Name: name,
			}

			switch reqMetric.Type {

			case MetricType_UNTYPED, MetricType_GAUGE, MetricType_DELTA:
				m.Type = reqMetric.Type

			default:
				continue
			}

			if v.labelName != "" && v.labelValue != "" {
				m.Labels = append(m.Labels, &MetricLabel{
					Name:  v.labelName,
					Value: v.labelValue,
				})
			}

			for _, _ = range ms.TimeBuckets {
				m.Points = append(m.Points, &MetricPoint{
					// Time: t,
				})
			}

			offset := 0
			for _, p := range v.points {

				if offset >= len(m.Points) {
					break
				}

				if (p.time + req.AlignmentPeriod) < ms.TimeBuckets[offset] {
					continue
				}

				for offset < len(m.Points) && p.time > ms.TimeBuckets[offset] {
					offset++
				}

				if offset >= len(m.Points) {
					break
				}

				m.Points[offset].Count += p.count
				m.Points[offset].Sum += p.sum
			}

			for _, p := range m.Points {

				switch reqMetric.Type {

				case MetricType_UNTYPED, MetricType_GAUGE:
					if p.Count > 1 {
						p.Value = p.Sum / p.Count
					} else {
						p.Value = p.Sum
					}

				case MetricType_DELTA:
					p.Value = p.Count
				}
			}

			ms.Metrics = append(ms.Metrics, m)
		}
	}

	ms.Metrics = metricsSort(ms.Metrics)

	return ms, nil
}

func (it *SampleSet) dbQuery(ms *MetricSet, req *SampleQueryRequest) (*MetricSet, error) {

	var (
		tl          = ms.TimeBuckets[0]
		tr          = ms.TimeBuckets[len(ms.TimeBuckets)-1]
		stl, str    = sampleStorageTimeBucketAlign(tl-req.AlignmentPeriod, tr)
		keyOffset   = storageKeyEncode(it.instanceId, stl)
		keyCutset   = storageKeyEncode(it.instanceId, str)
		reqMetric   *SampleQueryRequest_Metric
		ok          bool
		statsBytes  = 0
		statsPoint  = 0
		statsMetric = 0
	)

	if rs, err := it.storageClient.Scan(keyOffset, keyCutset); err == nil {

		for _, bs := range rs {

			var set MetricStorageSet

			if err := StdProto.Decode(bs, &set); err != nil {
				continue
			}

			for _, v := range set.Metrics {

				reqMetric, ok = req.Metrics[v.Name]
				if !ok {
					continue
				}

				labelName, labelValue := "", ""
				if len(v.Labels) > 0 {
					labelName, labelValue = v.Labels[0].Name, v.Labels[0].Value
				}

				if !reqMetric.labelMatch(labelName, labelValue) {
					continue
				}

				switch reqMetric.Type {

				case MetricType_UNTYPED, MetricType_GAUGE, MetricType_DELTA:

				default:
					continue
				}

				m := ms.getMetric(v.Name, labelName, labelValue)

				if len(m.Points) == 0 {
					for _, _ = range ms.TimeBuckets {
						m.Points = append(m.Points, &MetricPoint{
							// Time: t,
						})
					}
				}

				offset := 0
				for _, p := range v.Points {

					if offset >= len(m.Points) {
						break
					}

					if (p.Time + req.AlignmentPeriod) < ms.TimeBuckets[offset] {
						continue
					}

					for offset < len(m.Points) && p.Time > ms.TimeBuckets[offset] {
						offset++
					}

					if offset >= len(m.Points) {
						break
					}

					m.Points[offset].Count += p.Count
					m.Points[offset].Sum += p.Sum

					statsPoint++
				}

				for _, p := range m.Points {

					switch reqMetric.Type {

					case MetricType_UNTYPED, MetricType_GAUGE:
						if p.Count > 1 {
							p.Value = p.Sum / p.Count
						} else {
							p.Value = p.Sum
						}

					case MetricType_DELTA:
						p.Value = p.Count
					}
				}

				statsMetric++
			}
			statsBytes += len(bs)
		}
	}

	hlog.Printf("info", "metric/storage query metrics %d, points %d, bytes %d",
		statsMetric, statsPoint, statsBytes)

	ms.Metrics = metricsSort(ms.Metrics)

	return ms, nil
}
