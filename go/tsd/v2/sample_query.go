package tsd

import (
	"sort"
	"strings"
	"sync"
	"time"
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

	ms := &MetricSet{}

	if len(req.Metrics) == 0 {
		return ms, nil
	}

	var (
		tn        = timesec()
		reqMetric *SampleQueryRequest_Metric
		ok        bool
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

	it.mu.RLock()
	defer it.mu.RUnlock()

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

				// p.Count = 0
				// p.Sum = 0
			}

			ms.Metrics = append(ms.Metrics, m)
		}
	}

	sort.Slice(ms.Metrics, func(i, j int) bool {
		cmp := strings.Compare(ms.Metrics[i].Name, ms.Metrics[j].Name)
		if cmp != 0 {
			return cmp < 0
		}
		m1 := ms.Metrics[i]
		m2 := ms.Metrics[j]
		if len(m1.Labels) == 0 || len(m1.Labels) < len(m2.Labels) {
			return true
		}
		cmp = strings.Compare(m1.Labels[0].Name, m2.Labels[0].Name)
		if cmp != 0 {
			return cmp < 0
		}
		cmp = strings.Compare(m1.Labels[0].Value, m2.Labels[0].Value)
		return cmp < 0
	})

	return ms, nil
}
