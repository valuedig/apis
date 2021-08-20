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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.17.3
// source: tsd.proto

package tsd

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CycleFeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unit   int64         `protobuf:"varint,2,opt,name=unit,proto3" json:"unit,omitempty" toml:"unit,omitempty"`
	Items  []*CycleEntry `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty" toml:"items,omitempty"`
	Labels []string      `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" toml:"labels,omitempty"`
	Keys   []int64       `protobuf:"varint,9,rep,packed,name=keys,proto3" json:"keys,omitempty" toml:"keys,omitempty"`
}

func (x *CycleFeed) Reset() {
	*x = CycleFeed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tsd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CycleFeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CycleFeed) ProtoMessage() {}

func (x *CycleFeed) ProtoReflect() protoreflect.Message {
	mi := &file_tsd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CycleFeed.ProtoReflect.Descriptor instead.
func (*CycleFeed) Descriptor() ([]byte, []int) {
	return file_tsd_proto_rawDescGZIP(), []int{0}
}

func (x *CycleFeed) GetUnit() int64 {
	if x != nil {
		return x.Unit
	}
	return 0
}

func (x *CycleFeed) GetItems() []*CycleEntry {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CycleFeed) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CycleFeed) GetKeys() []int64 {
	if x != nil {
		return x.Keys
	}
	return nil
}

type CycleEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty"`
	Unit   int64   `protobuf:"varint,2,opt,name=unit,proto3" json:"unit,omitempty" toml:"unit,omitempty"`
	Keys   []int64 `protobuf:"varint,9,rep,packed,name=keys,proto3" json:"keys,omitempty" toml:"keys,omitempty"`
	Values []int64 `protobuf:"varint,10,rep,packed,name=values,proto3" json:"values,omitempty" toml:"values,omitempty"`
	Attrs  uint64  `protobuf:"varint,11,opt,name=attrs,proto3" json:"attrs,omitempty" toml:"attrs,omitempty"`
}

func (x *CycleEntry) Reset() {
	*x = CycleEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tsd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CycleEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CycleEntry) ProtoMessage() {}

func (x *CycleEntry) ProtoReflect() protoreflect.Message {
	mi := &file_tsd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CycleEntry.ProtoReflect.Descriptor instead.
func (*CycleEntry) Descriptor() ([]byte, []int) {
	return file_tsd_proto_rawDescGZIP(), []int{1}
}

func (x *CycleEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CycleEntry) GetUnit() int64 {
	if x != nil {
		return x.Unit
	}
	return 0
}

func (x *CycleEntry) GetKeys() []int64 {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *CycleEntry) GetValues() []int64 {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *CycleEntry) GetAttrs() uint64 {
	if x != nil {
		return x.Attrs
	}
	return 0
}

type CycleExportOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names     []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty" toml:"names,omitempty"`
	TimeUnit  int64    `protobuf:"varint,2,opt,name=time_unit,json=timeUnit,proto3" json:"time_unit,omitempty" toml:"time_unit,omitempty"`
	TimeStart int64    `protobuf:"varint,3,opt,name=time_start,json=timeStart,proto3" json:"time_start,omitempty" toml:"time_start,omitempty"`
	TimeEnd   int64    `protobuf:"varint,4,opt,name=time_end,json=timeEnd,proto3" json:"time_end,omitempty" toml:"time_end,omitempty"`
	TimeZone  int64    `protobuf:"varint,5,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty" toml:"time_zone,omitempty"`
}

func (x *CycleExportOptions) Reset() {
	*x = CycleExportOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tsd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CycleExportOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CycleExportOptions) ProtoMessage() {}

func (x *CycleExportOptions) ProtoReflect() protoreflect.Message {
	mi := &file_tsd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CycleExportOptions.ProtoReflect.Descriptor instead.
func (*CycleExportOptions) Descriptor() ([]byte, []int) {
	return file_tsd_proto_rawDescGZIP(), []int{2}
}

func (x *CycleExportOptions) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *CycleExportOptions) GetTimeUnit() int64 {
	if x != nil {
		return x.TimeUnit
	}
	return 0
}

func (x *CycleExportOptions) GetTimeStart() int64 {
	if x != nil {
		return x.TimeStart
	}
	return 0
}

func (x *CycleExportOptions) GetTimeEnd() int64 {
	if x != nil {
		return x.TimeEnd
	}
	return 0
}

func (x *CycleExportOptions) GetTimeZone() int64 {
	if x != nil {
		return x.TimeZone
	}
	return 0
}

var File_tsd_proto protoreflect.FileDescriptor

var file_tsd_proto_rawDesc = []byte{
	0x0a, 0x09, 0x74, 0x73, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x64, 0x69, 0x67, 0x2e, 0x74, 0x73, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x7e, 0x0a, 0x09,
	0x43, 0x79, 0x63, 0x6c, 0x65, 0x46, 0x65, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x31, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x64, 0x69, 0x67, 0x2e, 0x74, 0x73, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x79, 0x63, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x03, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x76, 0x0a, 0x0a,
	0x43, 0x79, 0x63, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x61,
	0x74, 0x74, 0x72, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x12, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x42, 0x09, 0x48, 0x03, 0x5a, 0x05, 0x2e, 0x3b, 0x74, 0x73, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tsd_proto_rawDescOnce sync.Once
	file_tsd_proto_rawDescData = file_tsd_proto_rawDesc
)

func file_tsd_proto_rawDescGZIP() []byte {
	file_tsd_proto_rawDescOnce.Do(func() {
		file_tsd_proto_rawDescData = protoimpl.X.CompressGZIP(file_tsd_proto_rawDescData)
	})
	return file_tsd_proto_rawDescData
}

var file_tsd_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tsd_proto_goTypes = []interface{}{
	(*CycleFeed)(nil),          // 0: valuedig.tsd.v1.CycleFeed
	(*CycleEntry)(nil),         // 1: valuedig.tsd.v1.CycleEntry
	(*CycleExportOptions)(nil), // 2: valuedig.tsd.v1.CycleExportOptions
}
var file_tsd_proto_depIdxs = []int32{
	1, // 0: valuedig.tsd.v1.CycleFeed.items:type_name -> valuedig.tsd.v1.CycleEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tsd_proto_init() }
func file_tsd_proto_init() {
	if File_tsd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tsd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CycleFeed); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tsd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CycleEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tsd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CycleExportOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tsd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tsd_proto_goTypes,
		DependencyIndexes: file_tsd_proto_depIdxs,
		MessageInfos:      file_tsd_proto_msgTypes,
	}.Build()
	File_tsd_proto = out.File
	file_tsd_proto_rawDesc = nil
	file_tsd_proto_goTypes = nil
	file_tsd_proto_depIdxs = nil
}
