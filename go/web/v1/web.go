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

package web

import (
	"net/url"
	"path/filepath"
	"strings"
	"time"
)

func UrlKey(u string) string {

	uo, err := url.Parse(u)
	if err != nil {
		return u
	}

	p := filepath.Clean(uo.Path)
	if p == "" || p == "." || p == ".." {
		p = "/"
	}

	if n := strings.Index(uo.Host, ":"); n > 0 {
		uo.Host = uo.Host[:n]
	}

	rs := "page/" + uo.Host + p

	if uo.RawQuery != "" {
		rs += "?" + uo.RawQuery
	}

	return strings.ToLower(rs)
}

func NewPage(u string) *Page {
	// idhash.HashToHexString([]byte(arg.(HashSeed)), 16)
	item := &Page{
		Url: u,
	}
	item.Updated = time.Now().Unix()
	item.Created = item.Updated
	return item
}
