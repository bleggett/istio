// Copyright Istio Authors
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
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: zdsapi/zds.proto

// GRPC package - part of the URL. Service is added.
// URL: /PACKAGE.SERVICE/METHOD

package zdsapi

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

type Version int32

const (
	Version_NOT_USED Version = 0
	Version_V1       Version = 1
)

// Enum value maps for Version.
var (
	Version_name = map[int32]string{
		0: "NOT_USED",
		1: "V1",
	}
	Version_value = map[string]int32{
		"NOT_USED": 0,
		"V1":       1,
	}
)

func (x Version) Enum() *Version {
	p := new(Version)
	*p = x
	return p
}

func (x Version) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Version) Descriptor() protoreflect.EnumDescriptor {
	return file_zdsapi_zds_proto_enumTypes[0].Descriptor()
}

func (Version) Type() protoreflect.EnumType {
	return &file_zdsapi_zds_proto_enumTypes[0]
}

func (x Version) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Version.Descriptor instead.
func (Version) EnumDescriptor() ([]byte, []int) {
	return file_zdsapi_zds_proto_rawDescGZIP(), []int{0}
}

type ZdsHello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version Version `protobuf:"varint,1,opt,name=version,proto3,enum=istio.workload.zds.Version" json:"version,omitempty"`
}

func (x *ZdsHello) Reset() {
	*x = ZdsHello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zdsapi_zds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZdsHello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZdsHello) ProtoMessage() {}

func (x *ZdsHello) ProtoReflect() protoreflect.Message {
	mi := &file_zdsapi_zds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZdsHello.ProtoReflect.Descriptor instead.
func (*ZdsHello) Descriptor() ([]byte, []int) {
	return file_zdsapi_zds_proto_rawDescGZIP(), []int{0}
}

func (x *ZdsHello) GetVersion() Version {
	if x != nil {
		return x.Version
	}
	return Version_NOT_USED
}

type WorkloadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace      string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ServiceAccount string `protobuf:"bytes,3,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
}

func (x *WorkloadInfo) Reset() {
	*x = WorkloadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zdsapi_zds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadInfo) ProtoMessage() {}

func (x *WorkloadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_zdsapi_zds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadInfo.ProtoReflect.Descriptor instead.
func (*WorkloadInfo) Descriptor() ([]byte, []int) {
	return file_zdsapi_zds_proto_rawDescGZIP(), []int{1}
}

func (x *WorkloadInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkloadInfo) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *WorkloadInfo) GetServiceAccount() string {
	if x != nil {
		return x.ServiceAccount
	}
	return ""
}

// Add a workload to the ztunnel. this will be accompanied by ancillary data contianing
// the workload's netns file descriptor.
type AddWorkload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid          string        `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	WorkloadInfo *WorkloadInfo `protobuf:"bytes,2,opt,name=workload_info,json=workloadInfo,proto3" json:"workload_info,omitempty"`
}

func (x *AddWorkload) Reset() {
	*x = AddWorkload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zdsapi_zds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddWorkload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddWorkload) ProtoMessage() {}

func (x *AddWorkload) ProtoReflect() protoreflect.Message {
	mi := &file_zdsapi_zds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddWorkload.ProtoReflect.Descriptor instead.
func (*AddWorkload) Descriptor() ([]byte, []int) {
	return file_zdsapi_zds_proto_rawDescGZIP(), []int{2}
}

func (x *AddWorkload) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *AddWorkload) GetWorkloadInfo() *WorkloadInfo {
	if x != nil {
		return x.WorkloadInfo
	}
	return nil
}

// Keep workload that we can't find in the fd cache. This can only be sent before SnapshotSent is sent
// to signal ztunnel to not delete the workload if it has it.
type KeepWorkload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *KeepWorkload) Reset() {
	*x = KeepWorkload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zdsapi_zds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeepWorkload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeepWorkload) ProtoMessage() {}

func (x *KeepWorkload) ProtoReflect() protoreflect.Message {
	mi := &file_zdsapi_zds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeepWorkload.ProtoReflect.Descriptor instead.
func (*KeepWorkload) Descriptor() ([]byte, []int) {
	return file_zdsapi_zds_proto_rawDescGZIP(), []int{3}
}

func (x *KeepWorkload) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

// Delete a workload from the ztunnel. Ztunnel should shutdown the workload's proxy.
type DelWorkload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DelWorkload) Reset() {
	*x = DelWorkload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zdsapi_zds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelWorkload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelWorkload) ProtoMessage() {}

func (x *DelWorkload) ProtoReflect() protoreflect.Message {
	mi := &file_zdsapi_zds_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelWorkload.ProtoReflect.Descriptor instead.
func (*DelWorkload) Descriptor() ([]byte, []int) {
	return file_zdsapi_zds_proto_rawDescGZIP(), []int{4}
}

func (x *DelWorkload) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

// Let ztunnel know that a full snapshot was sent. Ztunnel should reconcile its internal state
// and remove internal entries that were not sent.
type SnapshotSent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SnapshotSent) Reset() {
	*x = SnapshotSent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zdsapi_zds_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotSent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotSent) ProtoMessage() {}

func (x *SnapshotSent) ProtoReflect() protoreflect.Message {
	mi := &file_zdsapi_zds_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotSent.ProtoReflect.Descriptor instead.
func (*SnapshotSent) Descriptor() ([]byte, []int) {
	return file_zdsapi_zds_proto_rawDescGZIP(), []int{5}
}

// Ztunnel ack message. If error is not empty, this is an error message.
type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zdsapi_zds_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_zdsapi_zds_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_zdsapi_zds_proto_rawDescGZIP(), []int{6}
}

func (x *Ack) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// Sent from CNI to ztunnel
type WorkloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*WorkloadRequest_Add
	//	*WorkloadRequest_Keep
	//	*WorkloadRequest_Del
	//	*WorkloadRequest_SnapshotSent
	Payload isWorkloadRequest_Payload `protobuf_oneof:"payload"`
}

func (x *WorkloadRequest) Reset() {
	*x = WorkloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zdsapi_zds_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadRequest) ProtoMessage() {}

func (x *WorkloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zdsapi_zds_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadRequest.ProtoReflect.Descriptor instead.
func (*WorkloadRequest) Descriptor() ([]byte, []int) {
	return file_zdsapi_zds_proto_rawDescGZIP(), []int{7}
}

func (m *WorkloadRequest) GetPayload() isWorkloadRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *WorkloadRequest) GetAdd() *AddWorkload {
	if x, ok := x.GetPayload().(*WorkloadRequest_Add); ok {
		return x.Add
	}
	return nil
}

func (x *WorkloadRequest) GetKeep() *KeepWorkload {
	if x, ok := x.GetPayload().(*WorkloadRequest_Keep); ok {
		return x.Keep
	}
	return nil
}

func (x *WorkloadRequest) GetDel() *DelWorkload {
	if x, ok := x.GetPayload().(*WorkloadRequest_Del); ok {
		return x.Del
	}
	return nil
}

func (x *WorkloadRequest) GetSnapshotSent() *SnapshotSent {
	if x, ok := x.GetPayload().(*WorkloadRequest_SnapshotSent); ok {
		return x.SnapshotSent
	}
	return nil
}

type isWorkloadRequest_Payload interface {
	isWorkloadRequest_Payload()
}

type WorkloadRequest_Add struct {
	Add *AddWorkload `protobuf:"bytes,1,opt,name=add,proto3,oneof"`
}

type WorkloadRequest_Keep struct {
	Keep *KeepWorkload `protobuf:"bytes,5,opt,name=keep,proto3,oneof"`
}

type WorkloadRequest_Del struct {
	Del *DelWorkload `protobuf:"bytes,2,opt,name=del,proto3,oneof"`
}

type WorkloadRequest_SnapshotSent struct {
	SnapshotSent *SnapshotSent `protobuf:"bytes,3,opt,name=snapshot_sent,json=snapshotSent,proto3,oneof"`
}

func (*WorkloadRequest_Add) isWorkloadRequest_Payload() {}

func (*WorkloadRequest_Keep) isWorkloadRequest_Payload() {}

func (*WorkloadRequest_Del) isWorkloadRequest_Payload() {}

func (*WorkloadRequest_SnapshotSent) isWorkloadRequest_Payload() {}

// Sent from ztunnel to CNI
type WorkloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*WorkloadResponse_Ack
	Payload isWorkloadResponse_Payload `protobuf_oneof:"payload"`
}

func (x *WorkloadResponse) Reset() {
	*x = WorkloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zdsapi_zds_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadResponse) ProtoMessage() {}

func (x *WorkloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_zdsapi_zds_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadResponse.ProtoReflect.Descriptor instead.
func (*WorkloadResponse) Descriptor() ([]byte, []int) {
	return file_zdsapi_zds_proto_rawDescGZIP(), []int{8}
}

func (m *WorkloadResponse) GetPayload() isWorkloadResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *WorkloadResponse) GetAck() *Ack {
	if x, ok := x.GetPayload().(*WorkloadResponse_Ack); ok {
		return x.Ack
	}
	return nil
}

type isWorkloadResponse_Payload interface {
	isWorkloadResponse_Payload()
}

type WorkloadResponse_Ack struct {
	Ack *Ack `protobuf:"bytes,1,opt,name=ack,proto3,oneof"`
}

func (*WorkloadResponse_Ack) isWorkloadResponse_Payload() {}

var File_zdsapi_zds_proto protoreflect.FileDescriptor

var file_zdsapi_zds_proto_rawDesc = []byte{
	0x0a, 0x10, 0x7a, 0x64, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x7a, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x7a, 0x64, 0x73, 0x22, 0x41, 0x0a, 0x08, 0x5a, 0x64, 0x73, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x12, 0x35, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x7a, 0x64, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x69, 0x0a, 0x0c, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x66, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x7a, 0x64,
	0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c,
	0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x20, 0x0a, 0x0c,
	0x4b, 0x65, 0x65, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x1f,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x0e, 0x0a, 0x0c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x22,
	0x1b, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x87, 0x02, 0x0a,
	0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x33, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x7a,
	0x64, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00,
	0x52, 0x03, 0x61, 0x64, 0x64, 0x12, 0x36, 0x0a, 0x04, 0x6b, 0x65, 0x65, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x7a, 0x64, 0x73, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x04, 0x6b, 0x65, 0x65, 0x70, 0x12, 0x33, 0x0a,
	0x03, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x73, 0x74,
	0x69, 0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x7a, 0x64, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x03, 0x64,
	0x65, 0x6c, 0x12, 0x47, 0x0a, 0x0d, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x73,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x73, 0x74, 0x69,
	0x6f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x7a, 0x64, 0x73, 0x2e, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x73,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x4a, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x03, 0x61, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x7a, 0x64, 0x73, 0x2e, 0x41, 0x63, 0x6b,
	0x48, 0x00, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x2a, 0x1f, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a,
	0x08, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x56,
	0x31, 0x10, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x6b, 0x67, 0x2f, 0x7a, 0x64, 0x73, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zdsapi_zds_proto_rawDescOnce sync.Once
	file_zdsapi_zds_proto_rawDescData = file_zdsapi_zds_proto_rawDesc
)

func file_zdsapi_zds_proto_rawDescGZIP() []byte {
	file_zdsapi_zds_proto_rawDescOnce.Do(func() {
		file_zdsapi_zds_proto_rawDescData = protoimpl.X.CompressGZIP(file_zdsapi_zds_proto_rawDescData)
	})
	return file_zdsapi_zds_proto_rawDescData
}

var file_zdsapi_zds_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_zdsapi_zds_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_zdsapi_zds_proto_goTypes = []interface{}{
	(Version)(0),             // 0: istio.workload.zds.Version
	(*ZdsHello)(nil),         // 1: istio.workload.zds.ZdsHello
	(*WorkloadInfo)(nil),     // 2: istio.workload.zds.WorkloadInfo
	(*AddWorkload)(nil),      // 3: istio.workload.zds.AddWorkload
	(*KeepWorkload)(nil),     // 4: istio.workload.zds.KeepWorkload
	(*DelWorkload)(nil),      // 5: istio.workload.zds.DelWorkload
	(*SnapshotSent)(nil),     // 6: istio.workload.zds.SnapshotSent
	(*Ack)(nil),              // 7: istio.workload.zds.Ack
	(*WorkloadRequest)(nil),  // 8: istio.workload.zds.WorkloadRequest
	(*WorkloadResponse)(nil), // 9: istio.workload.zds.WorkloadResponse
}
var file_zdsapi_zds_proto_depIdxs = []int32{
	0, // 0: istio.workload.zds.ZdsHello.version:type_name -> istio.workload.zds.Version
	2, // 1: istio.workload.zds.AddWorkload.workload_info:type_name -> istio.workload.zds.WorkloadInfo
	3, // 2: istio.workload.zds.WorkloadRequest.add:type_name -> istio.workload.zds.AddWorkload
	4, // 3: istio.workload.zds.WorkloadRequest.keep:type_name -> istio.workload.zds.KeepWorkload
	5, // 4: istio.workload.zds.WorkloadRequest.del:type_name -> istio.workload.zds.DelWorkload
	6, // 5: istio.workload.zds.WorkloadRequest.snapshot_sent:type_name -> istio.workload.zds.SnapshotSent
	7, // 6: istio.workload.zds.WorkloadResponse.ack:type_name -> istio.workload.zds.Ack
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_zdsapi_zds_proto_init() }
func file_zdsapi_zds_proto_init() {
	if File_zdsapi_zds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zdsapi_zds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZdsHello); i {
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
		file_zdsapi_zds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadInfo); i {
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
		file_zdsapi_zds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddWorkload); i {
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
		file_zdsapi_zds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeepWorkload); i {
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
		file_zdsapi_zds_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelWorkload); i {
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
		file_zdsapi_zds_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotSent); i {
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
		file_zdsapi_zds_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
		file_zdsapi_zds_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadRequest); i {
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
		file_zdsapi_zds_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadResponse); i {
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
	file_zdsapi_zds_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*WorkloadRequest_Add)(nil),
		(*WorkloadRequest_Keep)(nil),
		(*WorkloadRequest_Del)(nil),
		(*WorkloadRequest_SnapshotSent)(nil),
	}
	file_zdsapi_zds_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*WorkloadResponse_Ack)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zdsapi_zds_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zdsapi_zds_proto_goTypes,
		DependencyIndexes: file_zdsapi_zds_proto_depIdxs,
		EnumInfos:         file_zdsapi_zds_proto_enumTypes,
		MessageInfos:      file_zdsapi_zds_proto_msgTypes,
	}.Build()
	File_zdsapi_zds_proto = out.File
	file_zdsapi_zds_proto_rawDesc = nil
	file_zdsapi_zds_proto_goTypes = nil
	file_zdsapi_zds_proto_depIdxs = nil
}
