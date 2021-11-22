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
	"regexp"
	"sort"
	"strings"
)

var (
	MetricNameRX = regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9\.\-\_]{0,48}([a-zA-Z0-9])`)
	LabelNameRX  = regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9\.\-\_]{0,48}([a-zA-Z0-9])`)
	LabelValueRX = regexp.MustCompile(`[a-zA-Z0-9\.\/\-\_]{0,50}([*]?)`)
)

func (it *MetricSet) getMetric(metricName, labelName, labelValue string) *Metric {

	var m *Metric

	for _, v := range it.Metrics {

		if v.Name != metricName {
			continue
		}

		if labelName == "" && len(v.Labels) == 0 {
			m = v
		} else if labelName != "" {
			for _, l := range v.Labels {
				if l.Name == labelName &&
					l.Value == labelValue {
					m = v
					break
				}
			}
		}

		if m != nil {
			break
		}
	}

	if m == nil {
		m = &Metric{
			Name: metricName,
		}
		if labelName != "" {
			m.Labels = append(m.Labels, &MetricLabel{
				Name:  labelName,
				Value: labelValue,
			})
		}

		it.Metrics = append(it.Metrics, m)
	}

	return m
}

func metricsSort(metrics []*Metric) []*Metric {

	sort.Slice(metrics, func(i, j int) bool {
		cmp := strings.Compare(metrics[i].Name, metrics[j].Name)
		if cmp != 0 {
			return cmp < 0
		}
		m1 := metrics[i]
		m2 := metrics[j]
		if len(m1.Labels) == 0 {
			return true
		}
		if len(m2.Labels) == 0 {
			return false
		}
		cmp = strings.Compare(m1.Labels[0].Name, m2.Labels[0].Name)
		if cmp != 0 {
			return cmp < 0
		}
		cmp = strings.Compare(m1.Labels[0].Value, m2.Labels[0].Value)
		return cmp < 0
	})

	return metrics
}
