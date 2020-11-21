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
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: web.proto

package web

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty"`
	Status  uint64     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty" toml:"status,omitempty"`
	Url     string     `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty" toml:"url,omitempty"`
	Body    []byte     `protobuf:"bytes,9,opt,name=body,proto3" json:"body,omitempty" toml:"body,omitempty"`
	Type    string     `protobuf:"bytes,11,opt,name=type,proto3" json:"type,omitempty" toml:"type,omitempty"`
	Created int64      `protobuf:"varint,14,opt,name=created,proto3" json:"created,omitempty" toml:"created,omitempty"`
	Updated int64      `protobuf:"varint,15,opt,name=updated,proto3" json:"updated,omitempty" toml:"updated,omitempty"`
	Logs    []*PageLog `protobuf:"bytes,16,rep,name=logs,proto3" json:"logs,omitempty" toml:"logs,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{0}
}

func (x *Page) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Page) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Page) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Page) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Page) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Page) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Page) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Page) GetLogs() []*PageLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

type PageLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  uint64 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty" toml:"status,omitempty"`
	Desc    string `protobuf:"bytes,9,opt,name=desc,proto3" json:"desc,omitempty" toml:"desc,omitempty"`
	Created int64  `protobuf:"varint,14,opt,name=created,proto3" json:"created,omitempty" toml:"created,omitempty"`
}

func (x *PageLog) Reset() {
	*x = PageLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageLog) ProtoMessage() {}

func (x *PageLog) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageLog.ProtoReflect.Descriptor instead.
func (*PageLog) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{1}
}

func (x *PageLog) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PageLog) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *PageLog) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

type ScrapeMatchEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// url, body
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty" toml:"type,omitempty"`
	// Regular Expression Matching
	Regexp string `protobuf:"bytes,3,opt,name=regexp,proto3" json:"regexp,omitempty" toml:"regexp,omitempty"`
	Desc   string `protobuf:"bytes,9,opt,name=desc,proto3" json:"desc,omitempty" toml:"desc,omitempty"`
}

func (x *ScrapeMatchEntry) Reset() {
	*x = ScrapeMatchEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScrapeMatchEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScrapeMatchEntry) ProtoMessage() {}

func (x *ScrapeMatchEntry) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScrapeMatchEntry.ProtoReflect.Descriptor instead.
func (*ScrapeMatchEntry) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{2}
}

func (x *ScrapeMatchEntry) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ScrapeMatchEntry) GetRegexp() string {
	if x != nil {
		return x.Regexp
	}
	return ""
}

func (x *ScrapeMatchEntry) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type Scraper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty"`
	Visits  []string            `protobuf:"bytes,5,rep,name=visits,proto3" json:"visits,omitempty" toml:"visits,omitempty"`
	Allow   []*ScrapeMatchEntry `protobuf:"bytes,6,rep,name=allow,proto3" json:"allow,omitempty" toml:"allow,omitempty"`
	Deny    []*ScrapeMatchEntry `protobuf:"bytes,7,rep,name=deny,proto3" json:"deny,omitempty" toml:"deny,omitempty"`
	ItemHit []*ScrapeMatchEntry `protobuf:"bytes,8,rep,name=item_hit,json=itemHit,proto3" json:"item_hit,omitempty" toml:"item_hit,omitempty"`
	ListHit []*ScrapeMatchEntry `protobuf:"bytes,9,rep,name=list_hit,json=listHit,proto3" json:"list_hit,omitempty" toml:"list_hit,omitempty"`
}

func (x *Scraper) Reset() {
	*x = Scraper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scraper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scraper) ProtoMessage() {}

func (x *Scraper) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scraper.ProtoReflect.Descriptor instead.
func (*Scraper) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{3}
}

func (x *Scraper) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Scraper) GetVisits() []string {
	if x != nil {
		return x.Visits
	}
	return nil
}

func (x *Scraper) GetAllow() []*ScrapeMatchEntry {
	if x != nil {
		return x.Allow
	}
	return nil
}

func (x *Scraper) GetDeny() []*ScrapeMatchEntry {
	if x != nil {
		return x.Deny
	}
	return nil
}

func (x *Scraper) GetItemHit() []*ScrapeMatchEntry {
	if x != nil {
		return x.ItemHit
	}
	return nil
}

func (x *Scraper) GetListHit() []*ScrapeMatchEntry {
	if x != nil {
		return x.ListHit
	}
	return nil
}

var File_web_proto protoreflect.FileDescriptor

var file_web_proto_rawDesc = []byte{
	0x0a, 0x09, 0x77, 0x65, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x64, 0x69, 0x67, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x22, 0xca, 0x01, 0x0a,
	0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x6c,
	0x6f, 0x67, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x64, 0x69, 0x67, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x4f, 0x0a, 0x07, 0x50, 0x61, 0x67,
	0x65, 0x4c, 0x6f, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x52, 0x0a, 0x10, 0x53, 0x63,
	0x72, 0x61, 0x70, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0xa1,
	0x02, 0x0a, 0x07, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x69, 0x73, 0x69, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x76, 0x69, 0x73, 0x69, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x64, 0x69, 0x67,
	0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12,
	0x35, 0x0a, 0x04, 0x64, 0x65, 0x6e, 0x79, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x64, 0x69, 0x67, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x64, 0x65, 0x6e, 0x79, 0x12, 0x3c, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x68,
	0x69, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x64, 0x69, 0x67, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x72, 0x61, 0x70,
	0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x69, 0x74, 0x65,
	0x6d, 0x48, 0x69, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x68, 0x69, 0x74,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x64, 0x69,
	0x67, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x48,
	0x69, 0x74, 0x42, 0x09, 0x48, 0x03, 0x5a, 0x05, 0x2e, 0x3b, 0x77, 0x65, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_proto_rawDescOnce sync.Once
	file_web_proto_rawDescData = file_web_proto_rawDesc
)

func file_web_proto_rawDescGZIP() []byte {
	file_web_proto_rawDescOnce.Do(func() {
		file_web_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_proto_rawDescData)
	})
	return file_web_proto_rawDescData
}

var file_web_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_web_proto_goTypes = []interface{}{
	(*Page)(nil),             // 0: valuedig.web.v1.Page
	(*PageLog)(nil),          // 1: valuedig.web.v1.PageLog
	(*ScrapeMatchEntry)(nil), // 2: valuedig.web.v1.ScrapeMatchEntry
	(*Scraper)(nil),          // 3: valuedig.web.v1.Scraper
}
var file_web_proto_depIdxs = []int32{
	1, // 0: valuedig.web.v1.Page.logs:type_name -> valuedig.web.v1.PageLog
	2, // 1: valuedig.web.v1.Scraper.allow:type_name -> valuedig.web.v1.ScrapeMatchEntry
	2, // 2: valuedig.web.v1.Scraper.deny:type_name -> valuedig.web.v1.ScrapeMatchEntry
	2, // 3: valuedig.web.v1.Scraper.item_hit:type_name -> valuedig.web.v1.ScrapeMatchEntry
	2, // 4: valuedig.web.v1.Scraper.list_hit:type_name -> valuedig.web.v1.ScrapeMatchEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_web_proto_init() }
func file_web_proto_init() {
	if File_web_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_web_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageLog); i {
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
		file_web_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScrapeMatchEntry); i {
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
		file_web_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scraper); i {
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
			RawDescriptor: file_web_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_proto_goTypes,
		DependencyIndexes: file_web_proto_depIdxs,
		MessageInfos:      file_web_proto_msgTypes,
	}.Build()
	File_web_proto = out.File
	file_web_proto_rawDesc = nil
	file_web_proto_goTypes = nil
	file_web_proto_depIdxs = nil
}
