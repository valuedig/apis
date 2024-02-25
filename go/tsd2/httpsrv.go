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

func NewHttpServiceModule() *httpsrv.Module {

	module := httpsrv.NewModule()

	module.RegisterController(new(Metric))

	return module
}
