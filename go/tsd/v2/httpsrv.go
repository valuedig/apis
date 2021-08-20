package tsd

import (
	"github.com/hooto/httpsrv"
)

type HttpService struct {
	*httpsrv.Controller
}

func (c HttpService) QueryAction() {

	var (
		req SampleQueryRequest
		rsp MetricSet
	)
	defer c.RenderJson(&rsp)

	if err := c.Request.JsonDecode(&req); err != nil {
		return
	}

	if rs, err := StdSampler.Query(&req); err != nil {
		return
	} else {
		rsp = *rs
	}
}

func NewHttpServiceModule() httpsrv.Module {

	module := httpsrv.NewModule("valuedig_apis_tsd_v2")

	module.ControllerRegister(new(Metric))

	return module
}
