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
	"testing"
)

func Test_RX(t *testing.T) {

	for _, str := range []string{
		"AAA_aaa-bbb",
	} {
		for _, rx := range []*regexp.Regexp{
			MetricNameRX, LabelNameRX, LabelValueRX,
		} {
			if !rx.MatchString(str) {
				t.Fatalf("fail with %s", str)
			}
		}
	}

	for _, str := range []string{
		"AAA_aaa-bbb*",
	} {
		if !LabelValueRX.MatchString(str) {
			t.Fatalf("fail with %s", str)
		}
	}
}
