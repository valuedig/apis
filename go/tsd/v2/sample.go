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
	"sync"
	"sync/atomic"
	"time"
)

const (
	defaultLabelName  = ""
	defaultLabelValue = ""
	maxEventQueue     = 1000
)

var (
	StdSampler *SampleSet
)

type SampleSetter interface {
	Inc(int64)
	Add(int64)
	Set(int64, int64)
	With(map[string]string) SampleSetter
}

type SampleSet struct {
	mu             sync.RWMutex
	running        bool
	events         chan *sampleMetricEvent
	families       map[string]*SampleFamily
	cleanInterval  int64
	cleanUpdated   int64
	instanceId     string
	flushes        map[int64]uint32
	storageFlushed int64
	storageClient  MetricStorage
	close          bool
}

type SampleFamily struct {
	mu      sync.RWMutex
	name    string
	metrics map[string]*SampleMetric
}

type SampleMetric struct {
	mu         sync.RWMutex
	labelName  string
	labelValue string
	points     []*SamplePoint
}

type sampleMetricEvent struct {
	metricName string
	labelName  string
	labelValue string
	delta      bool
	count      int64
	sum        int64
}

type sampleSetter struct {
	set        *SampleSet
	metricName string
	labels     []string
}

type SamplePoint struct {
	count int64
	sum   int64
	time  int64
}

func init() {
	StdSampler = NewSampler()
}

func NewSampler() *SampleSet {
	set := &SampleSet{
		events:        make(chan *sampleMetricEvent, 2*maxEventQueue),
		families:      map[string]*SampleFamily{},
		cleanInterval: 3600,
		cleanUpdated:  timesec(),
		flushes:       map[int64]uint32{},
		close:         false,
	}
	go set.run()
	return set
}

func (it *SampleSet) Metric(name string) SampleSetter {
	return &sampleSetter{
		set:        it,
		metricName: name,
	}
}

func (it *SampleSet) run() {

	it.mu.Lock()

	if it.running {
		it.mu.Unlock()
		return
	}

	it.running = true
	if it.families == nil {
		it.families = map[string]*SampleFamily{}
	}

	it.mu.Unlock()

	for !it.close {

		ev := <-it.events
		if ev == nil {
			break
		}

		tn := timesec()

		mf, ok := it.families[ev.metricName]
		if !ok {
			mf = &SampleFamily{
				name: ev.metricName,
			}
			it.families[ev.metricName] = mf
		}

		p := mf.getEntry(ev.labelName, ev.labelValue).point(tn)
		if ev.delta {
			if ev.count != 0 {
				atomic.AddInt64(&p.count, ev.count)
			}
			if ev.sum != 0 {
				atomic.AddInt64(&p.sum, ev.sum)
			}
		} else {
			atomic.StoreInt64(&p.count, ev.count)
			atomic.StoreInt64(&p.sum, ev.sum)
		}

		it.clean(false)
		it.flush(false)
	}
}

func (it *SampleSet) clean(force bool) {

	tn := time.Now()

	if it.cleanInterval < 3600 {
		it.cleanInterval = 3600
	}

	if !force && (it.cleanUpdated+(2*it.cleanInterval)) > tn.Unix() {
		return
	}

	var (
		num = 0
		ttl = tn.Unix() - it.cleanInterval
		idx = 0
		p   *SamplePoint
	)

	for _, mf := range it.families {
		dels := []string{}
		for mname, m := range mf.metrics {
			idx = -1
			for idx, p = range m.points {
				if p.time >= ttl {
					break
				}
			}
			if idx > 0 {
				num += idx
				m.points = m.points[idx:]
			}
			if len(m.points) <= 1 {
				dels = append(dels, mname)
			}
		}
		for _, mname := range dels {
			delete(mf.metrics, mname)
		}
	}

	it.cleanUpdated = tn.Unix()
}

func (it *SampleFamily) getEntry(labelName, labelValue string) *SampleMetric {

	if labelName == "" {
		labelValue = ""
	}

	labelKey := labelName + "/" + labelValue

	// it.mu.Lock()
	// defer it.mu.Unlock()

	if it.metrics == nil {
		it.metrics = map[string]*SampleMetric{}
	}

	s, ok := it.metrics[labelKey]
	if !ok {
		s = &SampleMetric{
			labelName:  labelName,
			labelValue: labelValue,
		}
		it.metrics[labelKey] = s
	}

	return s
}

func (it *SampleMetric) point(tn int64) *SamplePoint {
	// it.mu.Lock()
	// defer it.mu.Unlock()
	if len(it.points) == 0 || it.points[len(it.points)-1].time < tn {
		it.points = append(it.points, &SamplePoint{
			time: tn,
		})
	}
	return it.points[len(it.points)-1]
}

func (it *SamplePoint) Inc(v int64) {
	atomic.AddInt64(&it.count, v)
}

func (it *SamplePoint) Add(v int64) {
	atomic.AddInt64(&it.count, 1)
	atomic.AddInt64(&it.sum, v)
}

func (it *SamplePoint) Set(c int64, v int64) {
	atomic.StoreInt64(&it.count, c)
	atomic.StoreInt64(&it.sum, v)
}

func (it *sampleSetter) Inc(v int64) {
	it.add(v, 0, true)
}

func (it *sampleSetter) Add(v int64) {
	it.add(1, v, true)
}

func (it *sampleSetter) Set(count, sum int64) {
	it.add(count, sum, false)
}

func (it *sampleSetter) add(count, sum int64, delta bool) {
	if len(it.set.events) > maxEventQueue {
		return
	}
	if len(it.labels) > 1 {
		for i := 1; i < len(it.labels); i += 2 {
			it.set.events <- &sampleMetricEvent{
				metricName: it.metricName,
				labelName:  it.labels[i-1],
				labelValue: it.labels[i],
				delta:      delta,
				count:      count,
				sum:        sum,
			}
		}
	} else {
		it.set.events <- &sampleMetricEvent{
			metricName: it.metricName,
			delta:      delta,
			count:      count,
			sum:        sum,
		}
	}
}

func (it *sampleSetter) With(m map[string]string) SampleSetter {
	setter := &sampleSetter{
		metricName: it.metricName,
		set:        it.set,
	}
	for k, v := range m {
		if !LabelNameRX.MatchString(k) ||
			!LabelValueRX.MatchString(v) {
			continue
		}
		setter.labels = append(setter.labels, k)
		setter.labels = append(setter.labels, v)
	}
	return setter
}
