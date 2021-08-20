package tsd

import (
	"testing"
	"time"
)

func Test_Common(t *testing.T) {

	ms := NewSampler()

	for i := 0; i < 10; i++ {
		time.Sleep(500e6)
		for j := 0; j <= i; j++ {
			ms.Delta("hit").Inc(1)
			ms.Delta("label").With(map[string]string{
				"type": "01",
			}).Add(int64(i))
		}
	}

	q := &SampleQueryRequest{
		AlignmentPeriod: 2,
		LastTimeRange:   12,
	}

	q.MetricSelect("hit").Delta()
	q.MetricSelect("label").Gauge().LabelSelect("type", "*")

	rs, _ := ms.Query(q)

	objPrint("test_common", rs)
}
