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
