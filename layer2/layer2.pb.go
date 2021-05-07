//
// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: layer2/layer2.proto

package layer2

import (
	types "github.com/openconfig/gnoi/types"
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

type PerformBERTResponse_BERTState int32

const (
	PerformBERTResponse_UNKNOWN  PerformBERTResponse_BERTState = 0
	PerformBERTResponse_DISABLED PerformBERTResponse_BERTState = 1
	PerformBERTResponse_RUNNING  PerformBERTResponse_BERTState = 2
	PerformBERTResponse_COMPLETE PerformBERTResponse_BERTState = 3
	PerformBERTResponse_ERROR    PerformBERTResponse_BERTState = 4
)

// Enum value maps for PerformBERTResponse_BERTState.
var (
	PerformBERTResponse_BERTState_name = map[int32]string{
		0: "UNKNOWN",
		1: "DISABLED",
		2: "RUNNING",
		3: "COMPLETE",
		4: "ERROR",
	}
	PerformBERTResponse_BERTState_value = map[string]int32{
		"UNKNOWN":  0,
		"DISABLED": 1,
		"RUNNING":  2,
		"COMPLETE": 3,
		"ERROR":    4,
	}
)

func (x PerformBERTResponse_BERTState) Enum() *PerformBERTResponse_BERTState {
	p := new(PerformBERTResponse_BERTState)
	*p = x
	return p
}

func (x PerformBERTResponse_BERTState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PerformBERTResponse_BERTState) Descriptor() protoreflect.EnumDescriptor {
	return file_layer2_layer2_proto_enumTypes[0].Descriptor()
}

func (PerformBERTResponse_BERTState) Type() protoreflect.EnumType {
	return &file_layer2_layer2_proto_enumTypes[0]
}

func (x PerformBERTResponse_BERTState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PerformBERTResponse_BERTState.Descriptor instead.
func (PerformBERTResponse_BERTState) EnumDescriptor() ([]byte, []int) {
	return file_layer2_layer2_proto_rawDescGZIP(), []int{5, 0}
}

type ClearNeighborDiscoveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol types.L3Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=gnoi.types.L3Protocol" json:"protocol,omitempty"`
	Address  string           `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ClearNeighborDiscoveryRequest) Reset() {
	*x = ClearNeighborDiscoveryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer2_layer2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearNeighborDiscoveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearNeighborDiscoveryRequest) ProtoMessage() {}

func (x *ClearNeighborDiscoveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_layer2_layer2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearNeighborDiscoveryRequest.ProtoReflect.Descriptor instead.
func (*ClearNeighborDiscoveryRequest) Descriptor() ([]byte, []int) {
	return file_layer2_layer2_proto_rawDescGZIP(), []int{0}
}

func (x *ClearNeighborDiscoveryRequest) GetProtocol() types.L3Protocol {
	if x != nil {
		return x.Protocol
	}
	return types.L3Protocol_UNSPECIFIED
}

func (x *ClearNeighborDiscoveryRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ClearNeighborDiscoveryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearNeighborDiscoveryResponse) Reset() {
	*x = ClearNeighborDiscoveryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer2_layer2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearNeighborDiscoveryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearNeighborDiscoveryResponse) ProtoMessage() {}

func (x *ClearNeighborDiscoveryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_layer2_layer2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearNeighborDiscoveryResponse.ProtoReflect.Descriptor instead.
func (*ClearNeighborDiscoveryResponse) Descriptor() ([]byte, []int) {
	return file_layer2_layer2_proto_rawDescGZIP(), []int{1}
}

type ClearSpanningTreeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interface *types.Path `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
}

func (x *ClearSpanningTreeRequest) Reset() {
	*x = ClearSpanningTreeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer2_layer2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearSpanningTreeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearSpanningTreeRequest) ProtoMessage() {}

func (x *ClearSpanningTreeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_layer2_layer2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearSpanningTreeRequest.ProtoReflect.Descriptor instead.
func (*ClearSpanningTreeRequest) Descriptor() ([]byte, []int) {
	return file_layer2_layer2_proto_rawDescGZIP(), []int{2}
}

func (x *ClearSpanningTreeRequest) GetInterface() *types.Path {
	if x != nil {
		return x.Interface
	}
	return nil
}

type ClearSpanningTreeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearSpanningTreeResponse) Reset() {
	*x = ClearSpanningTreeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer2_layer2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearSpanningTreeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearSpanningTreeResponse) ProtoMessage() {}

func (x *ClearSpanningTreeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_layer2_layer2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearSpanningTreeResponse.ProtoReflect.Descriptor instead.
func (*ClearSpanningTreeResponse) Descriptor() ([]byte, []int) {
	return file_layer2_layer2_proto_rawDescGZIP(), []int{3}
}

type PerformBERTRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID for retrieving a previous BERT run data - optional.
	Id        string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Interface *types.Path `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
}

func (x *PerformBERTRequest) Reset() {
	*x = PerformBERTRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer2_layer2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerformBERTRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerformBERTRequest) ProtoMessage() {}

func (x *PerformBERTRequest) ProtoReflect() protoreflect.Message {
	mi := &file_layer2_layer2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerformBERTRequest.ProtoReflect.Descriptor instead.
func (*PerformBERTRequest) Descriptor() ([]byte, []int) {
	return file_layer2_layer2_proto_rawDescGZIP(), []int{4}
}

func (x *PerformBERTRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PerformBERTRequest) GetInterface() *types.Path {
	if x != nil {
		return x.Interface
	}
	return nil
}

type PerformBERTResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State         PerformBERTResponse_BERTState `protobuf:"varint,2,opt,name=state,proto3,enum=gnoi.layer2.PerformBERTResponse_BERTState" json:"state,omitempty"`
	ElapsedPeriod int64                         `protobuf:"varint,3,opt,name=elapsed_period,json=elapsedPeriod,proto3" json:"elapsed_period,omitempty"` // BERT test length in nanoseconds.
	Pattern       []byte                        `protobuf:"bytes,4,opt,name=pattern,proto3" json:"pattern,omitempty"`                                   // Pattern used for the BERT test.
	// Number of errors experienced since the start of the BERT test.
	Errors int64 `protobuf:"varint,5,opt,name=errors,proto3" json:"errors,omitempty"`
	// Number of bits received since the start of the BERT test.
	ReceivedBits int64 `protobuf:"varint,6,opt,name=received_bits,json=receivedBits,proto3" json:"received_bits,omitempty"`
}

func (x *PerformBERTResponse) Reset() {
	*x = PerformBERTResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer2_layer2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerformBERTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerformBERTResponse) ProtoMessage() {}

func (x *PerformBERTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_layer2_layer2_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerformBERTResponse.ProtoReflect.Descriptor instead.
func (*PerformBERTResponse) Descriptor() ([]byte, []int) {
	return file_layer2_layer2_proto_rawDescGZIP(), []int{5}
}

func (x *PerformBERTResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PerformBERTResponse) GetState() PerformBERTResponse_BERTState {
	if x != nil {
		return x.State
	}
	return PerformBERTResponse_UNKNOWN
}

func (x *PerformBERTResponse) GetElapsedPeriod() int64 {
	if x != nil {
		return x.ElapsedPeriod
	}
	return 0
}

func (x *PerformBERTResponse) GetPattern() []byte {
	if x != nil {
		return x.Pattern
	}
	return nil
}

func (x *PerformBERTResponse) GetErrors() int64 {
	if x != nil {
		return x.Errors
	}
	return 0
}

func (x *PerformBERTResponse) GetReceivedBits() int64 {
	if x != nil {
		return x.ReceivedBits
	}
	return 0
}

type ClearLLDPInterfaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interface *types.Path `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
}

func (x *ClearLLDPInterfaceRequest) Reset() {
	*x = ClearLLDPInterfaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer2_layer2_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearLLDPInterfaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearLLDPInterfaceRequest) ProtoMessage() {}

func (x *ClearLLDPInterfaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_layer2_layer2_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearLLDPInterfaceRequest.ProtoReflect.Descriptor instead.
func (*ClearLLDPInterfaceRequest) Descriptor() ([]byte, []int) {
	return file_layer2_layer2_proto_rawDescGZIP(), []int{6}
}

func (x *ClearLLDPInterfaceRequest) GetInterface() *types.Path {
	if x != nil {
		return x.Interface
	}
	return nil
}

type ClearLLDPInterfaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearLLDPInterfaceResponse) Reset() {
	*x = ClearLLDPInterfaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer2_layer2_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearLLDPInterfaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearLLDPInterfaceResponse) ProtoMessage() {}

func (x *ClearLLDPInterfaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_layer2_layer2_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearLLDPInterfaceResponse.ProtoReflect.Descriptor instead.
func (*ClearLLDPInterfaceResponse) Descriptor() ([]byte, []int) {
	return file_layer2_layer2_proto_rawDescGZIP(), []int{7}
}

type SendWakeOnLANRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interface  *types.Path `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	Address    string      `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`                         // IP address of the WOL target.
	MacAddress []byte      `protobuf:"bytes,3,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"` // MAC address of the target.
}

func (x *SendWakeOnLANRequest) Reset() {
	*x = SendWakeOnLANRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer2_layer2_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendWakeOnLANRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendWakeOnLANRequest) ProtoMessage() {}

func (x *SendWakeOnLANRequest) ProtoReflect() protoreflect.Message {
	mi := &file_layer2_layer2_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendWakeOnLANRequest.ProtoReflect.Descriptor instead.
func (*SendWakeOnLANRequest) Descriptor() ([]byte, []int) {
	return file_layer2_layer2_proto_rawDescGZIP(), []int{8}
}

func (x *SendWakeOnLANRequest) GetInterface() *types.Path {
	if x != nil {
		return x.Interface
	}
	return nil
}

func (x *SendWakeOnLANRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SendWakeOnLANRequest) GetMacAddress() []byte {
	if x != nil {
		return x.MacAddress
	}
	return nil
}

type SendWakeOnLANResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendWakeOnLANResponse) Reset() {
	*x = SendWakeOnLANResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_layer2_layer2_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendWakeOnLANResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendWakeOnLANResponse) ProtoMessage() {}

func (x *SendWakeOnLANResponse) ProtoReflect() protoreflect.Message {
	mi := &file_layer2_layer2_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendWakeOnLANResponse.ProtoReflect.Descriptor instead.
func (*SendWakeOnLANResponse) Descriptor() ([]byte, []int) {
	return file_layer2_layer2_proto_rawDescGZIP(), []int{9}
}

var File_layer2_layer2_proto protoreflect.FileDescriptor

var file_layer2_layer2_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x2f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x32, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x6f, 0x69, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6d, 0x0a, 0x1d, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f,
	0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4c, 0x33, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x20, 0x0a, 0x1e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x4a, 0x0a, 0x18, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a,
	0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x61,
	0x74, 0x68, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x22, 0x1b, 0x0a,
	0x19, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x72,
	0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54, 0x0a, 0x12, 0x50, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x45, 0x52, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2e, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x22, 0xb3, 0x02, 0x0a, 0x13, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x45, 0x52, 0x54,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x32, 0x2e, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x45, 0x52,
	0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x45, 0x52, 0x54, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6c,
	0x61, 0x70, 0x73, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f,
	0x62, 0x69, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x42, 0x69, 0x74, 0x73, 0x22, 0x4c, 0x0a, 0x09, 0x42, 0x45, 0x52, 0x54,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x22, 0x4b, 0x0a, 0x19, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x4c,
	0x4c, 0x44, 0x50, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x22, 0x1c, 0x0a, 0x1a, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x4c, 0x4c, 0x44, 0x50,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x81, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x61, 0x6b, 0x65, 0x4f, 0x6e,
	0x4c, 0x41, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52,
	0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x61, 0x6b,
	0x65, 0x4f, 0x6e, 0x4c, 0x41, 0x4e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xfc,
	0x03, 0x0a, 0x06, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x12, 0x73, 0x0a, 0x16, 0x43, 0x6c, 0x65,
	0x61, 0x72, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x12, 0x2a, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x32, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x2e, 0x43, 0x6c,
	0x65, 0x61, 0x72, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64,
	0x0a, 0x11, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x54,
	0x72, 0x65, 0x65, 0x12, 0x25, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x32, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x54,
	0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x6e, 0x6f,
	0x69, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x53, 0x70,
	0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0b, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x42,
	0x45, 0x52, 0x54, 0x12, 0x1f, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x32, 0x2e, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x45, 0x52, 0x54, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x32, 0x2e, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x45, 0x52, 0x54, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x67, 0x0a, 0x12, 0x43, 0x6c,
	0x65, 0x61, 0x72, 0x4c, 0x4c, 0x44, 0x50, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x12, 0x26, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x2e, 0x43,
	0x6c, 0x65, 0x61, 0x72, 0x4c, 0x4c, 0x44, 0x50, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x4c, 0x4c, 0x44, 0x50,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x61, 0x6b, 0x65, 0x4f,
	0x6e, 0x4c, 0x41, 0x4e, 0x12, 0x21, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x32, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x61, 0x6b, 0x65, 0x4f, 0x6e, 0x4c, 0x41, 0x4e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x32, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x61, 0x6b, 0x65, 0x4f, 0x6e,
	0x4c, 0x41, 0x4e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a,
	0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x6f, 0x69, 0x2f, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x32, 0xd2, 0x3e, 0x05, 0x30, 0x2e, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_layer2_layer2_proto_rawDescOnce sync.Once
	file_layer2_layer2_proto_rawDescData = file_layer2_layer2_proto_rawDesc
)

func file_layer2_layer2_proto_rawDescGZIP() []byte {
	file_layer2_layer2_proto_rawDescOnce.Do(func() {
		file_layer2_layer2_proto_rawDescData = protoimpl.X.CompressGZIP(file_layer2_layer2_proto_rawDescData)
	})
	return file_layer2_layer2_proto_rawDescData
}

var file_layer2_layer2_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_layer2_layer2_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_layer2_layer2_proto_goTypes = []interface{}{
	(PerformBERTResponse_BERTState)(0),     // 0: gnoi.layer2.PerformBERTResponse.BERTState
	(*ClearNeighborDiscoveryRequest)(nil),  // 1: gnoi.layer2.ClearNeighborDiscoveryRequest
	(*ClearNeighborDiscoveryResponse)(nil), // 2: gnoi.layer2.ClearNeighborDiscoveryResponse
	(*ClearSpanningTreeRequest)(nil),       // 3: gnoi.layer2.ClearSpanningTreeRequest
	(*ClearSpanningTreeResponse)(nil),      // 4: gnoi.layer2.ClearSpanningTreeResponse
	(*PerformBERTRequest)(nil),             // 5: gnoi.layer2.PerformBERTRequest
	(*PerformBERTResponse)(nil),            // 6: gnoi.layer2.PerformBERTResponse
	(*ClearLLDPInterfaceRequest)(nil),      // 7: gnoi.layer2.ClearLLDPInterfaceRequest
	(*ClearLLDPInterfaceResponse)(nil),     // 8: gnoi.layer2.ClearLLDPInterfaceResponse
	(*SendWakeOnLANRequest)(nil),           // 9: gnoi.layer2.SendWakeOnLANRequest
	(*SendWakeOnLANResponse)(nil),          // 10: gnoi.layer2.SendWakeOnLANResponse
	(types.L3Protocol)(0),                  // 11: gnoi.types.L3Protocol
	(*types.Path)(nil),                     // 12: gnoi.types.Path
}
var file_layer2_layer2_proto_depIdxs = []int32{
	11, // 0: gnoi.layer2.ClearNeighborDiscoveryRequest.protocol:type_name -> gnoi.types.L3Protocol
	12, // 1: gnoi.layer2.ClearSpanningTreeRequest.interface:type_name -> gnoi.types.Path
	12, // 2: gnoi.layer2.PerformBERTRequest.interface:type_name -> gnoi.types.Path
	0,  // 3: gnoi.layer2.PerformBERTResponse.state:type_name -> gnoi.layer2.PerformBERTResponse.BERTState
	12, // 4: gnoi.layer2.ClearLLDPInterfaceRequest.interface:type_name -> gnoi.types.Path
	12, // 5: gnoi.layer2.SendWakeOnLANRequest.interface:type_name -> gnoi.types.Path
	1,  // 6: gnoi.layer2.Layer2.ClearNeighborDiscovery:input_type -> gnoi.layer2.ClearNeighborDiscoveryRequest
	3,  // 7: gnoi.layer2.Layer2.ClearSpanningTree:input_type -> gnoi.layer2.ClearSpanningTreeRequest
	5,  // 8: gnoi.layer2.Layer2.PerformBERT:input_type -> gnoi.layer2.PerformBERTRequest
	7,  // 9: gnoi.layer2.Layer2.ClearLLDPInterface:input_type -> gnoi.layer2.ClearLLDPInterfaceRequest
	9,  // 10: gnoi.layer2.Layer2.SendWakeOnLAN:input_type -> gnoi.layer2.SendWakeOnLANRequest
	2,  // 11: gnoi.layer2.Layer2.ClearNeighborDiscovery:output_type -> gnoi.layer2.ClearNeighborDiscoveryResponse
	4,  // 12: gnoi.layer2.Layer2.ClearSpanningTree:output_type -> gnoi.layer2.ClearSpanningTreeResponse
	6,  // 13: gnoi.layer2.Layer2.PerformBERT:output_type -> gnoi.layer2.PerformBERTResponse
	8,  // 14: gnoi.layer2.Layer2.ClearLLDPInterface:output_type -> gnoi.layer2.ClearLLDPInterfaceResponse
	10, // 15: gnoi.layer2.Layer2.SendWakeOnLAN:output_type -> gnoi.layer2.SendWakeOnLANResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_layer2_layer2_proto_init() }
func file_layer2_layer2_proto_init() {
	if File_layer2_layer2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_layer2_layer2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearNeighborDiscoveryRequest); i {
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
		file_layer2_layer2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearNeighborDiscoveryResponse); i {
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
		file_layer2_layer2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearSpanningTreeRequest); i {
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
		file_layer2_layer2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearSpanningTreeResponse); i {
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
		file_layer2_layer2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerformBERTRequest); i {
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
		file_layer2_layer2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerformBERTResponse); i {
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
		file_layer2_layer2_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearLLDPInterfaceRequest); i {
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
		file_layer2_layer2_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearLLDPInterfaceResponse); i {
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
		file_layer2_layer2_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendWakeOnLANRequest); i {
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
		file_layer2_layer2_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendWakeOnLANResponse); i {
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
			RawDescriptor: file_layer2_layer2_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_layer2_layer2_proto_goTypes,
		DependencyIndexes: file_layer2_layer2_proto_depIdxs,
		EnumInfos:         file_layer2_layer2_proto_enumTypes,
		MessageInfos:      file_layer2_layer2_proto_msgTypes,
	}.Build()
	File_layer2_layer2_proto = out.File
	file_layer2_layer2_proto_rawDesc = nil
	file_layer2_layer2_proto_goTypes = nil
	file_layer2_layer2_proto_depIdxs = nil
}
