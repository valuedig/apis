package tsd

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/golang/protobuf/proto"
)

var (
	MetricNameRX = regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9\-\_]{0,48}([a-zA-Z0-9])`)
	LabelNameRX  = regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9\-\_]{0,48}([a-zA-Z0-9])`)
	LabelValueRX = regexp.MustCompile(`[a-zA-Z0-9\.\/\-\_]{0,50}([*]?)`)
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
