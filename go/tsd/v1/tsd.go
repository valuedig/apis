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
	"sync"
	"time"
)

const (
	CycleKeysMax       = 100000
	CycleUnitMax int64 = 3600
	CycleUnitMin int64 = 1
)

var (
	cycleFeedMU  sync.RWMutex
	cycleEntryMU sync.RWMutex
)

func NewCycleFeed(unit int64) *CycleFeed {
	if fix := unit % CycleUnitMin; fix > 0 {
		unit -= unit
	}
	if unit < CycleUnitMin {
		unit = CycleUnitMin
	} else if unit > CycleUnitMax {
		unit = CycleUnitMax
	}
	return &CycleFeed{
		Unit: unit,
	}
}

func (it *CycleFeed) Entry(name string) *CycleEntry {

	cycleFeedMU.Lock()
	defer cycleFeedMU.Unlock()

	for _, v := range it.Items {
		if name == v.Name {
			return v
		}
	}
	entry := &CycleEntry{
		Name: name,
		Unit: it.Unit,
	}
	it.Items = append(it.Items, entry)
	return entry
}

func (it *CycleFeed) Sync(name string, key, value int64, add bool) {
	it.Entry(name).Sync(key, value, add)
}

func (it *CycleFeed) Trim(sec int64) {

	if sec < 1 {
		sec = 1
	}
	tn := time.Now().Unix()
	tl := tn - sec

	cycleFeedMU.Lock()
	defer cycleFeedMU.Unlock()

	for _, v := range it.Items {
		for i, v2 := range v.Keys {
			if v2 > tl {
				continue
			}
			v.Keys = v.Keys[i:]
			v.Values = v.Values[i:]
			break
		}
	}
}

func (it *CycleEntry) add(key, value int64) {
	for i := 0; i < len(it.Keys); i++ {

		if key > it.Keys[i] {
			continue
		}

		if key == it.Keys[i] {
			it.Values[i] += value
		} else {
			it.Keys = append(append(it.Keys[:i], key), it.Keys[i:]...)
			it.Values = append(append(it.Values[:i], value), it.Values[i:]...)
		}

		return
	}

	it.Keys = append(it.Keys, key)
	it.Values = append(it.Values, value)
}

func (it *CycleEntry) Sync(key, value int64, add bool) {

	var kt time.Time

	if key <= 0 {
		kt = time.Now()
		key = kt.Unix()
	} else {
		kt = time.Unix(key, 0)
	}

	if it.Unit < 1 {
		it.Unit = 1
	}

	if it.Unit >= 3600 {
		key -= int64(kt.Minute()) * 60
		key -= int64(kt.Second())
	} else if it.Unit >= 60 {
		key -= (int64(kt.Minute()) % (it.Unit / 60)) * 60
		key -= int64(kt.Second())
	} else {
		key -= int64(kt.Second()) % it.Unit
	}

	cycleEntryMU.Lock()
	defer cycleEntryMU.Unlock()

	for i := 0; i < len(it.Keys); i++ {

		if key > it.Keys[i] {
			continue
		}

		if key == it.Keys[i] {
			if add {
				it.Values[i] += value
			} else {
				it.Values[i] = value
			}
		} else {
			it.Keys = append(append(it.Keys[:i], key), it.Keys[i:]...)
			it.Values = append(append(it.Values[:i], value), it.Values[i:]...)
		}

		return
	}

	it.Keys = append(it.Keys, key)
	it.Values = append(it.Values, value)

	if len(it.Keys) > CycleKeysMax {
		it.Keys = it.Keys[CycleKeysMax-len(it.Keys):]
		it.Values = it.Keys[CycleKeysMax-len(it.Values):]
	}
}

type CycleTimeExportZone int

type CycleTimeExportUnit int64

func (it *CycleFeed) Export(args ...interface{}) *CycleFeed {

	cycleFeedMU.Lock()
	defer cycleFeedMU.Unlock()

	var (
		unit int64 = 0
		tzz        = time.UTC
	)

	for _, arg := range args {

		switch arg.(type) {

		case CycleTimeExportZone:
			tz := int(arg.(CycleTimeExportZone))
			if tz < -11 {
				tz = -11
			} else if tz > 11 {
				tz = 11
			}
			if tz != 0 {
				tzz = time.FixedZone("CST", (tz * 3600))
			}

		case CycleTimeExportUnit:
			unit = int64(arg.(CycleTimeExportUnit))
			if unit < 1 {
				unit = 1
			} else if unit > 86400 {
				unit = 86400
			}
		}
	}

	var (
		unitDay    int64 = 0
		unitHour   int64 = 0
		unitMinute int64 = 0
	)

	if unit >= 86400 {
		unitDay = unit / 86400
	} else if unit >= 3600 {
		unitHour = unit / 3600
	} else {
		unitMinute = unit / 60
		if unitMinute < 1 {
			unitMinute = 1
		}
	}

	feed := &CycleFeed{
		Unit: unit,
	}

	for _, v := range it.Items {

		if len(v.Keys) != len(v.Values) {
			continue
		}

		entry := &CycleEntry{
			Name: v.Name,
			Unit: unit,
		}

		for j, k1 := range v.Keys {

			kt := time.Unix(k1, 0).In(tzz)

			k1 = int64(kt.Year()) * 10000
			k1 += int64(kt.Month()) * 100
			k1 += int64(kt.Day())

			if unitDay > 0 {
				if fix := int64(kt.Day()) % unitDay; fix > 0 {
					k1 -= fix
				}
			} else if unitHour > 0 {
				k1 *= 100
				k1 += int64(kt.Hour())
				if fix := int64(kt.Hour()) % unitHour; fix > 0 {
					k1 -= fix
				}
			} else {
				k1 *= 10000
				k1 += int64(kt.Hour()) * 100
				k1 += int64(kt.Minute())
				if fix := int64(kt.Minute()) % unitMinute; fix > 0 {
					k1 -= fix
				}
			}

			entry.add(k1, v.Values[j])
		}

		if len(entry.Keys) > 0 {
			feed.Items = append(feed.Items, entry)
		}
	}

	return feed
}
