// Based on https://github.com/openstreetmap/OSM-binary/blob/65e7e976f5c8e47f057a0d921639ea8e6309ef06/osmpbf/osmformat.proto

//* Copyright (c) 2010 Scott A. Crosby. <scott@sacrosby.com>
//Permission is hereby granted, free of charge, to any person obtaining a copy of
//this software and associated documentation files (the "Software"), to deal in
//the Software without restriction, including without limitation the rights to
//use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
//of the Software, and to permit persons to whom the Software is furnished to do
//so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.2
// source: osmformat.proto

package pbfproto

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

type Relation_MemberType int32

const (
	Relation_NODE     Relation_MemberType = 0
	Relation_WAY      Relation_MemberType = 1
	Relation_RELATION Relation_MemberType = 2
)

// Enum value maps for Relation_MemberType.
var (
	Relation_MemberType_name = map[int32]string{
		0: "NODE",
		1: "WAY",
		2: "RELATION",
	}
	Relation_MemberType_value = map[string]int32{
		"NODE":     0,
		"WAY":      1,
		"RELATION": 2,
	}
)

func (x Relation_MemberType) Enum() *Relation_MemberType {
	p := new(Relation_MemberType)
	*p = x
	return p
}

func (x Relation_MemberType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Relation_MemberType) Descriptor() protoreflect.EnumDescriptor {
	return file_osmformat_proto_enumTypes[0].Descriptor()
}

func (Relation_MemberType) Type() protoreflect.EnumType {
	return &file_osmformat_proto_enumTypes[0]
}

func (x Relation_MemberType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Relation_MemberType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Relation_MemberType(num)
	return nil
}

// Deprecated: Use Relation_MemberType.Descriptor instead.
func (Relation_MemberType) EnumDescriptor() ([]byte, []int) {
	return file_osmformat_proto_rawDescGZIP(), []int{7, 0}
}

type HeaderBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Additional tags to aid in parsing this dataset
	RequiredFeatures []string `protobuf:"bytes,4,rep,name=required_features,json=requiredFeatures" json:"required_features,omitempty"`
	OptionalFeatures []string `protobuf:"bytes,5,rep,name=optional_features,json=optionalFeatures" json:"optional_features,omitempty"`
}

func (x *HeaderBlock) Reset() {
	*x = HeaderBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_osmformat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderBlock) ProtoMessage() {}

func (x *HeaderBlock) ProtoReflect() protoreflect.Message {
	mi := &file_osmformat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderBlock.ProtoReflect.Descriptor instead.
func (*HeaderBlock) Descriptor() ([]byte, []int) {
	return file_osmformat_proto_rawDescGZIP(), []int{0}
}

func (x *HeaderBlock) GetRequiredFeatures() []string {
	if x != nil {
		return x.RequiredFeatures
	}
	return nil
}

func (x *HeaderBlock) GetOptionalFeatures() []string {
	if x != nil {
		return x.OptionalFeatures
	}
	return nil
}

type PrimitiveBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stringtable    *StringTable      `protobuf:"bytes,1,req,name=stringtable" json:"stringtable,omitempty"`
	Primitivegroup []*PrimitiveGroup `protobuf:"bytes,2,rep,name=primitivegroup" json:"primitivegroup,omitempty"`
	// Granularity, units of nanodegrees, used to store coordinates in this block.
	Granularity *int32 `protobuf:"varint,17,opt,name=granularity,def=100" json:"granularity,omitempty"`
	// Offset value between the output coordinates and the granularity grid in units of nanodegrees.
	LatOffset *int64 `protobuf:"varint,19,opt,name=lat_offset,json=latOffset,def=0" json:"lat_offset,omitempty"`
	LonOffset *int64 `protobuf:"varint,20,opt,name=lon_offset,json=lonOffset,def=0" json:"lon_offset,omitempty"`
}

// Default values for PrimitiveBlock fields.
const (
	Default_PrimitiveBlock_Granularity = int32(100)
	Default_PrimitiveBlock_LatOffset   = int64(0)
	Default_PrimitiveBlock_LonOffset   = int64(0)
)

func (x *PrimitiveBlock) Reset() {
	*x = PrimitiveBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_osmformat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimitiveBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimitiveBlock) ProtoMessage() {}

func (x *PrimitiveBlock) ProtoReflect() protoreflect.Message {
	mi := &file_osmformat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimitiveBlock.ProtoReflect.Descriptor instead.
func (*PrimitiveBlock) Descriptor() ([]byte, []int) {
	return file_osmformat_proto_rawDescGZIP(), []int{1}
}

func (x *PrimitiveBlock) GetStringtable() *StringTable {
	if x != nil {
		return x.Stringtable
	}
	return nil
}

func (x *PrimitiveBlock) GetPrimitivegroup() []*PrimitiveGroup {
	if x != nil {
		return x.Primitivegroup
	}
	return nil
}

func (x *PrimitiveBlock) GetGranularity() int32 {
	if x != nil && x.Granularity != nil {
		return *x.Granularity
	}
	return Default_PrimitiveBlock_Granularity
}

func (x *PrimitiveBlock) GetLatOffset() int64 {
	if x != nil && x.LatOffset != nil {
		return *x.LatOffset
	}
	return Default_PrimitiveBlock_LatOffset
}

func (x *PrimitiveBlock) GetLonOffset() int64 {
	if x != nil && x.LonOffset != nil {
		return *x.LonOffset
	}
	return Default_PrimitiveBlock_LonOffset
}

// Group of OSMPrimitives. All primitives in a group must be the same type.
type PrimitiveGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes     []*Node     `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	Dense     *DenseNodes `protobuf:"bytes,2,opt,name=dense" json:"dense,omitempty"`
	Ways      []*Way      `protobuf:"bytes,3,rep,name=ways" json:"ways,omitempty"`
	Relations []*Relation `protobuf:"bytes,4,rep,name=relations" json:"relations,omitempty"`
}

func (x *PrimitiveGroup) Reset() {
	*x = PrimitiveGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_osmformat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimitiveGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimitiveGroup) ProtoMessage() {}

func (x *PrimitiveGroup) ProtoReflect() protoreflect.Message {
	mi := &file_osmformat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimitiveGroup.ProtoReflect.Descriptor instead.
func (*PrimitiveGroup) Descriptor() ([]byte, []int) {
	return file_osmformat_proto_rawDescGZIP(), []int{2}
}

func (x *PrimitiveGroup) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *PrimitiveGroup) GetDense() *DenseNodes {
	if x != nil {
		return x.Dense
	}
	return nil
}

func (x *PrimitiveGroup) GetWays() []*Way {
	if x != nil {
		return x.Ways
	}
	return nil
}

func (x *PrimitiveGroup) GetRelations() []*Relation {
	if x != nil {
		return x.Relations
	}
	return nil
}

type StringTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S [][]byte `protobuf:"bytes,1,rep,name=s" json:"s,omitempty"`
}

func (x *StringTable) Reset() {
	*x = StringTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_osmformat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringTable) ProtoMessage() {}

func (x *StringTable) ProtoReflect() protoreflect.Message {
	mi := &file_osmformat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringTable.ProtoReflect.Descriptor instead.
func (*StringTable) Descriptor() ([]byte, []int) {
	return file_osmformat_proto_rawDescGZIP(), []int{3}
}

func (x *StringTable) GetS() [][]byte {
	if x != nil {
		return x.S
	}
	return nil
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *int64 `protobuf:"zigzag64,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"` // String IDs.
	Vals []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"` // String IDs.
	Lat  *int64   `protobuf:"zigzag64,8,req,name=lat" json:"lat,omitempty"`
	Lon  *int64   `protobuf:"zigzag64,9,req,name=lon" json:"lon,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_osmformat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_osmformat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_osmformat_proto_rawDescGZIP(), []int{4}
}

func (x *Node) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Node) GetKeys() []uint32 {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *Node) GetVals() []uint32 {
	if x != nil {
		return x.Vals
	}
	return nil
}

func (x *Node) GetLat() int64 {
	if x != nil && x.Lat != nil {
		return *x.Lat
	}
	return 0
}

func (x *Node) GetLon() int64 {
	if x != nil && x.Lon != nil {
		return *x.Lon
	}
	return 0
}

type DenseNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  []int64 `protobuf:"zigzag64,1,rep,packed,name=id" json:"id,omitempty"`   // DELTA coded
	Lat []int64 `protobuf:"zigzag64,8,rep,packed,name=lat" json:"lat,omitempty"` // DELTA coded
	Lon []int64 `protobuf:"zigzag64,9,rep,packed,name=lon" json:"lon,omitempty"` // DELTA coded
	// Special packing of keys and vals into one array. May be empty if all nodes in this block are tagless.
	KeysVals []int32 `protobuf:"varint,10,rep,packed,name=keys_vals,json=keysVals" json:"keys_vals,omitempty"`
}

func (x *DenseNodes) Reset() {
	*x = DenseNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_osmformat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DenseNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenseNodes) ProtoMessage() {}

func (x *DenseNodes) ProtoReflect() protoreflect.Message {
	mi := &file_osmformat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenseNodes.ProtoReflect.Descriptor instead.
func (*DenseNodes) Descriptor() ([]byte, []int) {
	return file_osmformat_proto_rawDescGZIP(), []int{5}
}

func (x *DenseNodes) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *DenseNodes) GetLat() []int64 {
	if x != nil {
		return x.Lat
	}
	return nil
}

func (x *DenseNodes) GetLon() []int64 {
	if x != nil {
		return x.Lon
	}
	return nil
}

func (x *DenseNodes) GetKeysVals() []int32 {
	if x != nil {
		return x.KeysVals
	}
	return nil
}

type Way struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	Refs []int64  `protobuf:"zigzag64,8,rep,packed,name=refs" json:"refs,omitempty"` // DELTA coded
}

func (x *Way) Reset() {
	*x = Way{}
	if protoimpl.UnsafeEnabled {
		mi := &file_osmformat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Way) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Way) ProtoMessage() {}

func (x *Way) ProtoReflect() protoreflect.Message {
	mi := &file_osmformat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Way.ProtoReflect.Descriptor instead.
func (*Way) Descriptor() ([]byte, []int) {
	return file_osmformat_proto_rawDescGZIP(), []int{6}
}

func (x *Way) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Way) GetKeys() []uint32 {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *Way) GetVals() []uint32 {
	if x != nil {
		return x.Vals
	}
	return nil
}

func (x *Way) GetRefs() []int64 {
	if x != nil {
		return x.Refs
	}
	return nil
}

type Relation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	// Parallel arrays
	RolesSid []int32               `protobuf:"varint,8,rep,packed,name=roles_sid,json=rolesSid" json:"roles_sid,omitempty"` // This should have been defined as uint32 for consistency, but it is now too late to change it
	Memids   []int64               `protobuf:"zigzag64,9,rep,packed,name=memids" json:"memids,omitempty"`                   // DELTA encoded
	Types    []Relation_MemberType `protobuf:"varint,10,rep,packed,name=types,enum=OSMPBF.Relation_MemberType" json:"types,omitempty"`
}

func (x *Relation) Reset() {
	*x = Relation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_osmformat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relation) ProtoMessage() {}

func (x *Relation) ProtoReflect() protoreflect.Message {
	mi := &file_osmformat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relation.ProtoReflect.Descriptor instead.
func (*Relation) Descriptor() ([]byte, []int) {
	return file_osmformat_proto_rawDescGZIP(), []int{7}
}

func (x *Relation) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Relation) GetKeys() []uint32 {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *Relation) GetVals() []uint32 {
	if x != nil {
		return x.Vals
	}
	return nil
}

func (x *Relation) GetRolesSid() []int32 {
	if x != nil {
		return x.RolesSid
	}
	return nil
}

func (x *Relation) GetMemids() []int64 {
	if x != nil {
		return x.Memids
	}
	return nil
}

func (x *Relation) GetTypes() []Relation_MemberType {
	if x != nil {
		return x.Types
	}
	return nil
}

var File_osmformat_proto protoreflect.FileDescriptor

var file_osmformat_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x73, 0x6d, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x4f, 0x53, 0x4d, 0x50, 0x42, 0x46, 0x22, 0x67, 0x0a, 0x0b, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x22, 0xf2, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x35, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4f, 0x53, 0x4d,
	0x50, 0x42, 0x46, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3e, 0x0a, 0x0e,
	0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x4f, 0x53, 0x4d, 0x50, 0x42, 0x46, 0x2e, 0x50, 0x72,
	0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0e, 0x70, 0x72,
	0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x25, 0x0a, 0x0b,
	0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x05, 0x3a, 0x03, 0x31, 0x30, 0x30, 0x52, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0a, 0x6c, 0x61, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x3a, 0x01, 0x30, 0x52, 0x09, 0x6c, 0x61, 0x74, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0a, 0x6c, 0x6f, 0x6e, 0x5f, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x3a, 0x01, 0x30, 0x52, 0x09, 0x6c, 0x6f,
	0x6e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xaf, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x69, 0x6d,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x22, 0x0a, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x53, 0x4d, 0x50,
	0x42, 0x46, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x28,
	0x0a, 0x05, 0x64, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x4f, 0x53, 0x4d, 0x50, 0x42, 0x46, 0x2e, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x77, 0x61, 0x79, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4f, 0x53, 0x4d, 0x50, 0x42, 0x46, 0x2e,
	0x57, 0x61, 0x79, 0x52, 0x04, 0x77, 0x61, 0x79, 0x73, 0x12, 0x2e, 0x0a, 0x09, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4f,
	0x53, 0x4d, 0x50, 0x42, 0x46, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1b, 0x0a, 0x0b, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x01, 0x73, 0x22, 0x6a, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x12, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x16, 0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x01, 0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x08, 0x20, 0x02, 0x28, 0x12, 0x52, 0x03, 0x6c, 0x61, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x02, 0x28, 0x12, 0x52, 0x03, 0x6c,
	0x6f, 0x6e, 0x22, 0x6d, 0x0a, 0x0a, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x12, 0x42, 0x02, 0x10, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x12, 0x42, 0x02, 0x10, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x03, 0x6c, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x03, 0x28, 0x12, 0x42, 0x02, 0x10, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x05, 0x42, 0x02, 0x10, 0x01, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x73, 0x56, 0x61, 0x6c,
	0x73, 0x22, 0x5d, 0x0a, 0x03, 0x57, 0x61, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x12, 0x16, 0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x04, 0x72, 0x65, 0x66, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x12, 0x42, 0x02, 0x10, 0x01, 0x52, 0x04, 0x72, 0x65, 0x66, 0x73,
	0x22, 0xed, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x01, 0x52,
	0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x16, 0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x01, 0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x12, 0x1f, 0x0a,
	0x09, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x08, 0x20, 0x03, 0x28, 0x05,
	0x42, 0x02, 0x10, 0x01, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x53, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x12, 0x42, 0x02,
	0x10, 0x01, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x69, 0x64, 0x73, 0x12, 0x35, 0x0a, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x4f, 0x53, 0x4d, 0x50,
	0x42, 0x46, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x02, 0x10, 0x01, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x22, 0x2d, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x41, 0x59,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x62, 0x66, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_osmformat_proto_rawDescOnce sync.Once
	file_osmformat_proto_rawDescData = file_osmformat_proto_rawDesc
)

func file_osmformat_proto_rawDescGZIP() []byte {
	file_osmformat_proto_rawDescOnce.Do(func() {
		file_osmformat_proto_rawDescData = protoimpl.X.CompressGZIP(file_osmformat_proto_rawDescData)
	})
	return file_osmformat_proto_rawDescData
}

var file_osmformat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_osmformat_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_osmformat_proto_goTypes = []any{
	(Relation_MemberType)(0), // 0: OSMPBF.Relation.MemberType
	(*HeaderBlock)(nil),      // 1: OSMPBF.HeaderBlock
	(*PrimitiveBlock)(nil),   // 2: OSMPBF.PrimitiveBlock
	(*PrimitiveGroup)(nil),   // 3: OSMPBF.PrimitiveGroup
	(*StringTable)(nil),      // 4: OSMPBF.StringTable
	(*Node)(nil),             // 5: OSMPBF.Node
	(*DenseNodes)(nil),       // 6: OSMPBF.DenseNodes
	(*Way)(nil),              // 7: OSMPBF.Way
	(*Relation)(nil),         // 8: OSMPBF.Relation
}
var file_osmformat_proto_depIdxs = []int32{
	4, // 0: OSMPBF.PrimitiveBlock.stringtable:type_name -> OSMPBF.StringTable
	3, // 1: OSMPBF.PrimitiveBlock.primitivegroup:type_name -> OSMPBF.PrimitiveGroup
	5, // 2: OSMPBF.PrimitiveGroup.nodes:type_name -> OSMPBF.Node
	6, // 3: OSMPBF.PrimitiveGroup.dense:type_name -> OSMPBF.DenseNodes
	7, // 4: OSMPBF.PrimitiveGroup.ways:type_name -> OSMPBF.Way
	8, // 5: OSMPBF.PrimitiveGroup.relations:type_name -> OSMPBF.Relation
	0, // 6: OSMPBF.Relation.types:type_name -> OSMPBF.Relation.MemberType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_osmformat_proto_init() }
func file_osmformat_proto_init() {
	if File_osmformat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_osmformat_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HeaderBlock); i {
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
		file_osmformat_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PrimitiveBlock); i {
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
		file_osmformat_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PrimitiveGroup); i {
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
		file_osmformat_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*StringTable); i {
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
		file_osmformat_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Node); i {
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
		file_osmformat_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DenseNodes); i {
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
		file_osmformat_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Way); i {
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
		file_osmformat_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Relation); i {
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
			RawDescriptor: file_osmformat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_osmformat_proto_goTypes,
		DependencyIndexes: file_osmformat_proto_depIdxs,
		EnumInfos:         file_osmformat_proto_enumTypes,
		MessageInfos:      file_osmformat_proto_msgTypes,
	}.Build()
	File_osmformat_proto = out.File
	file_osmformat_proto_rawDesc = nil
	file_osmformat_proto_goTypes = nil
	file_osmformat_proto_depIdxs = nil
}