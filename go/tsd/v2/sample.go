package tsd

import (
	"sync"
	"sync/atomic"
	"time"
)

const (
	defaultLabelName  = ""
	defaultLabelValue = ""
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
	mu            sync.RWMutex
	families      map[string]*SampleFamily
	cleanInterval int64
	cleanUpdated  int64
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

type SamplePoint struct {
	count int64
	sum   int64
	time  int64
}

type multiPointsSetter struct {
	points []*SamplePoint
}

func init() {
	StdSampler = NewSampler()
}

func NewSampler() *SampleSet {
	return &SampleSet{
		families:      map[string]*SampleFamily{},
		cleanInterval: 3600,
		cleanUpdated:  timesec(),
	}
}

func (it *SampleSet) Metric(name string) SampleSetter {
	return it.getFamily(name, MetricType_UNTYPED)
}

func (it *SampleSet) Gauge(name string) SampleSetter {
	return it.getFamily(name, MetricType_GAUGE)
}

func (it *SampleSet) Delta(name string) SampleSetter {
	return it.getFamily(name, MetricType_DELTA)
}

func (it *SampleSet) getFamily(name string, typ MetricType) *SampleFamily {

	it.mu.Lock()
	defer it.mu.Unlock()

	if it.families == nil {
		it.families = map[string]*SampleFamily{}
	}

	it.clean(false)

	mf, ok := it.families[name]
	if !ok {
		mf = &SampleFamily{
			name: name,
		}
		it.families[name] = mf
	}
	return mf
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

	it.mu.Lock()
	defer it.mu.Unlock()

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

func (it *SampleFamily) Inc(i int64) {
	it.getEntry(defaultLabelName, defaultLabelValue).point(timesec()).Inc(i)
}

func (it *SampleFamily) Add(v int64) {
	it.getEntry(defaultLabelName, defaultLabelValue).point(timesec()).Add(v)
}

func (it *SampleFamily) Set(i, v int64) {
	it.getEntry(defaultLabelName, defaultLabelValue).point(timesec()).Set(i, v)
}

func (it *SampleFamily) With(m map[string]string) SampleSetter {

	e := &multiPointsSetter{}
	if m == nil {
		return e
	}
	tn := timesec()

	for k, v := range m {

		if !LabelNameRX.MatchString(k) ||
			!LabelValueRX.MatchString(v) {
			continue
		}

		e.points = append(e.points, it.getEntry(k, v).point(tn))
	}

	return e
}

func (it *SampleMetric) point(tn int64) *SamplePoint {
	it.mu.Lock()
	defer it.mu.Unlock()
	if len(it.points) == 0 || it.points[len(it.points)-1].time < tn {
		it.points = append(it.points, &SamplePoint{
			time: tn,
		})
	}
	return it.points[len(it.points)-1]
}

//
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

//
func (it *multiPointsSetter) Inc(v int64) {
	for _, p := range it.points {
		p.Inc(v)
	}
}

func (it *multiPointsSetter) Add(v int64) {
	for _, p := range it.points {
		p.Add(v)
	}
}

func (it *multiPointsSetter) Set(count int64, sum int64) {
	for _, p := range it.points {
		p.Set(count, sum)
	}
}

func (it *multiPointsSetter) With(m map[string]string) SampleSetter {
	return it
}

func (it *SampleSet) Load(b []byte) error {

	it.mu.Lock()
	defer it.mu.Unlock()

	var ms MetricSet
	if err := StdProto.Decode(b, &ms); err != nil {
		return err
	}

	var m2 *SampleMetric
	it.families = map[string]*SampleFamily{}

	for _, m := range ms.Metrics {
		mf, ok := it.families[m.Name]
		if !ok {
			mf = &SampleFamily{
				name: m.Name,
			}
			it.families[m.Name] = mf
		}

		if len(m.Labels) > 0 {
			for _, label := range m.Labels {
				m2 = mf.getEntry(label.Name, label.Value)
				if m2 != nil {
					m2.labelName = label.Name
					m2.labelValue = label.Value
					break
				}
			}
		} else {
			m2 = mf.getEntry("", "")
		}

		if m2 == nil {
			continue
		}
		for _, v := range m.Points {
			m2.points = append(m2.points, &SamplePoint{
				count: v.Count,
				sum:   v.Sum,
				time:  v.Time,
			})
		}
	}

	return nil
}

func (it *SampleSet) Dump() ([]byte, error) {

	it.mu.Lock()
	defer it.mu.Unlock()

	it.clean(true)

	ms := &MetricSet{}

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
				m2.Points = append(m2.Points, &MetricPoint{
					Count: p.count,
					Sum:   p.sum,
					Time:  p.time,
				})
			}
			ms.Metrics = append(ms.Metrics, m2)
		}
	}
	return StdProto.Encode(ms)
}

func timesec() int64 {
	return time.Now().Unix()
}
