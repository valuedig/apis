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
	"testing"
	"time"
)

func Test_Common(t *testing.T) {

	ms := NewSampler()

	for i := 0; i < 10; i++ {
		time.Sleep(500e6)
		for j := 0; j <= i; j++ {
			ms.Metric("hit").Inc(1)
			ms.Metric("label").With(map[string]string{
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
