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
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
)

func objPrint(name string, obj interface{}) {
	js, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Println(name, string(js))
}

type ProtoCodec struct{}

func (ProtoCodec) Encode(obj proto.Message) ([]byte, error) {
	return proto.Marshal(obj)
}

func (ProtoCodec) Decode(bs []byte, obj proto.Message) error {
	return proto.Unmarshal(bs, obj)
}

var StdProto = &ProtoCodec{}

func uint32ToBytes(v uint32) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, v)
	return bs
}

func uint32ToHexString(v uint32) string {
	return hex.EncodeToString(uint32ToBytes(v))
}

func timesec() int64 {
	return time.Now().Unix()
}
