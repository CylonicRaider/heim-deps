// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/firestore/v1beta1/common.proto

package firestore

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A set of field paths on a document.
// Used to restrict a get or update operation on a document to a subset of its
// fields.
// This is different from standard field masks, as this is always scoped to a
// [Document][google.firestore.v1beta1.Document], and takes in account the dynamic nature of [Value][google.firestore.v1beta1.Value].
type DocumentMask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of field paths in the mask. See [Document.fields][google.firestore.v1beta1.Document.fields] for a field
	// path syntax reference.
	FieldPaths []string `protobuf:"bytes,1,rep,name=field_paths,json=fieldPaths,proto3" json:"field_paths,omitempty"`
}

func (x *DocumentMask) Reset() {
	*x = DocumentMask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_v1beta1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentMask) ProtoMessage() {}

func (x *DocumentMask) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_v1beta1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentMask.ProtoReflect.Descriptor instead.
func (*DocumentMask) Descriptor() ([]byte, []int) {
	return file_google_firestore_v1beta1_common_proto_rawDescGZIP(), []int{0}
}

func (x *DocumentMask) GetFieldPaths() []string {
	if x != nil {
		return x.FieldPaths
	}
	return nil
}

// A precondition on a document, used for conditional operations.
type Precondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of precondition.
	//
	// Types that are assignable to ConditionType:
	//
	//	*Precondition_Exists
	//	*Precondition_UpdateTime
	ConditionType isPrecondition_ConditionType `protobuf_oneof:"condition_type"`
}

func (x *Precondition) Reset() {
	*x = Precondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_v1beta1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Precondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Precondition) ProtoMessage() {}

func (x *Precondition) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_v1beta1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Precondition.ProtoReflect.Descriptor instead.
func (*Precondition) Descriptor() ([]byte, []int) {
	return file_google_firestore_v1beta1_common_proto_rawDescGZIP(), []int{1}
}

func (m *Precondition) GetConditionType() isPrecondition_ConditionType {
	if m != nil {
		return m.ConditionType
	}
	return nil
}

func (x *Precondition) GetExists() bool {
	if x, ok := x.GetConditionType().(*Precondition_Exists); ok {
		return x.Exists
	}
	return false
}

func (x *Precondition) GetUpdateTime() *timestamppb.Timestamp {
	if x, ok := x.GetConditionType().(*Precondition_UpdateTime); ok {
		return x.UpdateTime
	}
	return nil
}

type isPrecondition_ConditionType interface {
	isPrecondition_ConditionType()
}

type Precondition_Exists struct {
	// When set to `true`, the target document must exist.
	// When set to `false`, the target document must not exist.
	Exists bool `protobuf:"varint,1,opt,name=exists,proto3,oneof"`
}

type Precondition_UpdateTime struct {
	// When set, the target document must exist and have been last updated at
	// that time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=update_time,json=updateTime,proto3,oneof"`
}

func (*Precondition_Exists) isPrecondition_ConditionType() {}

func (*Precondition_UpdateTime) isPrecondition_ConditionType() {}

// Options for creating a new transaction.
type TransactionOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mode of the transaction.
	//
	// Types that are assignable to Mode:
	//
	//	*TransactionOptions_ReadOnly_
	//	*TransactionOptions_ReadWrite_
	Mode isTransactionOptions_Mode `protobuf_oneof:"mode"`
}

func (x *TransactionOptions) Reset() {
	*x = TransactionOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_v1beta1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionOptions) ProtoMessage() {}

func (x *TransactionOptions) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_v1beta1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionOptions.ProtoReflect.Descriptor instead.
func (*TransactionOptions) Descriptor() ([]byte, []int) {
	return file_google_firestore_v1beta1_common_proto_rawDescGZIP(), []int{2}
}

func (m *TransactionOptions) GetMode() isTransactionOptions_Mode {
	if m != nil {
		return m.Mode
	}
	return nil
}

func (x *TransactionOptions) GetReadOnly() *TransactionOptions_ReadOnly {
	if x, ok := x.GetMode().(*TransactionOptions_ReadOnly_); ok {
		return x.ReadOnly
	}
	return nil
}

func (x *TransactionOptions) GetReadWrite() *TransactionOptions_ReadWrite {
	if x, ok := x.GetMode().(*TransactionOptions_ReadWrite_); ok {
		return x.ReadWrite
	}
	return nil
}

type isTransactionOptions_Mode interface {
	isTransactionOptions_Mode()
}

type TransactionOptions_ReadOnly_ struct {
	// The transaction can only be used for read operations.
	ReadOnly *TransactionOptions_ReadOnly `protobuf:"bytes,2,opt,name=read_only,json=readOnly,proto3,oneof"`
}

type TransactionOptions_ReadWrite_ struct {
	// The transaction can be used for both read and write operations.
	ReadWrite *TransactionOptions_ReadWrite `protobuf:"bytes,3,opt,name=read_write,json=readWrite,proto3,oneof"`
}

func (*TransactionOptions_ReadOnly_) isTransactionOptions_Mode() {}

func (*TransactionOptions_ReadWrite_) isTransactionOptions_Mode() {}

// Options for a transaction that can be used to read and write documents.
type TransactionOptions_ReadWrite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An optional transaction to retry.
	RetryTransaction []byte `protobuf:"bytes,1,opt,name=retry_transaction,json=retryTransaction,proto3" json:"retry_transaction,omitempty"`
}

func (x *TransactionOptions_ReadWrite) Reset() {
	*x = TransactionOptions_ReadWrite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_v1beta1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionOptions_ReadWrite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionOptions_ReadWrite) ProtoMessage() {}

func (x *TransactionOptions_ReadWrite) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_v1beta1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionOptions_ReadWrite.ProtoReflect.Descriptor instead.
func (*TransactionOptions_ReadWrite) Descriptor() ([]byte, []int) {
	return file_google_firestore_v1beta1_common_proto_rawDescGZIP(), []int{2, 0}
}

func (x *TransactionOptions_ReadWrite) GetRetryTransaction() []byte {
	if x != nil {
		return x.RetryTransaction
	}
	return nil
}

// Options for a transaction that can only be used to read documents.
type TransactionOptions_ReadOnly struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The consistency mode for this transaction. If not set, defaults to strong
	// consistency.
	//
	// Types that are assignable to ConsistencySelector:
	//
	//	*TransactionOptions_ReadOnly_ReadTime
	ConsistencySelector isTransactionOptions_ReadOnly_ConsistencySelector `protobuf_oneof:"consistency_selector"`
}

func (x *TransactionOptions_ReadOnly) Reset() {
	*x = TransactionOptions_ReadOnly{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_v1beta1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionOptions_ReadOnly) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionOptions_ReadOnly) ProtoMessage() {}

func (x *TransactionOptions_ReadOnly) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_v1beta1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionOptions_ReadOnly.ProtoReflect.Descriptor instead.
func (*TransactionOptions_ReadOnly) Descriptor() ([]byte, []int) {
	return file_google_firestore_v1beta1_common_proto_rawDescGZIP(), []int{2, 1}
}

func (m *TransactionOptions_ReadOnly) GetConsistencySelector() isTransactionOptions_ReadOnly_ConsistencySelector {
	if m != nil {
		return m.ConsistencySelector
	}
	return nil
}

func (x *TransactionOptions_ReadOnly) GetReadTime() *timestamppb.Timestamp {
	if x, ok := x.GetConsistencySelector().(*TransactionOptions_ReadOnly_ReadTime); ok {
		return x.ReadTime
	}
	return nil
}

type isTransactionOptions_ReadOnly_ConsistencySelector interface {
	isTransactionOptions_ReadOnly_ConsistencySelector()
}

type TransactionOptions_ReadOnly_ReadTime struct {
	// Reads documents at the given time.
	// This may not be older than 60 seconds.
	ReadTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=read_time,json=readTime,proto3,oneof"`
}

func (*TransactionOptions_ReadOnly_ReadTime) isTransactionOptions_ReadOnly_ConsistencySelector() {}

var File_google_firestore_v1beta1_common_proto protoreflect.FileDescriptor

var file_google_firestore_v1beta1_common_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61,
	0x73, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61,
	0x74, 0x68, 0x73, 0x22, 0x79, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x3d, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x10, 0x0a, 0x0e,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xe4,
	0x02, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x54, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e,
	0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x48,
	0x00, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x57, 0x0a, 0x0a, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65, 0x61, 0x64, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x1a, 0x38, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x72, 0x65,
	0x74, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x5d,
	0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x39, 0x0a, 0x09, 0x72, 0x65,
	0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x61,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x16, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x06, 0x0a,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x42, 0xdd, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x66, 0x69, 0x72,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x66,
	0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0xa2, 0x02, 0x04, 0x47, 0x43, 0x46, 0x53, 0xaa,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x46,
	0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31,
	0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_firestore_v1beta1_common_proto_rawDescOnce sync.Once
	file_google_firestore_v1beta1_common_proto_rawDescData = file_google_firestore_v1beta1_common_proto_rawDesc
)

func file_google_firestore_v1beta1_common_proto_rawDescGZIP() []byte {
	file_google_firestore_v1beta1_common_proto_rawDescOnce.Do(func() {
		file_google_firestore_v1beta1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_firestore_v1beta1_common_proto_rawDescData)
	})
	return file_google_firestore_v1beta1_common_proto_rawDescData
}

var file_google_firestore_v1beta1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_firestore_v1beta1_common_proto_goTypes = []interface{}{
	(*DocumentMask)(nil),                 // 0: google.firestore.v1beta1.DocumentMask
	(*Precondition)(nil),                 // 1: google.firestore.v1beta1.Precondition
	(*TransactionOptions)(nil),           // 2: google.firestore.v1beta1.TransactionOptions
	(*TransactionOptions_ReadWrite)(nil), // 3: google.firestore.v1beta1.TransactionOptions.ReadWrite
	(*TransactionOptions_ReadOnly)(nil),  // 4: google.firestore.v1beta1.TransactionOptions.ReadOnly
	(*timestamppb.Timestamp)(nil),        // 5: google.protobuf.Timestamp
}
var file_google_firestore_v1beta1_common_proto_depIdxs = []int32{
	5, // 0: google.firestore.v1beta1.Precondition.update_time:type_name -> google.protobuf.Timestamp
	4, // 1: google.firestore.v1beta1.TransactionOptions.read_only:type_name -> google.firestore.v1beta1.TransactionOptions.ReadOnly
	3, // 2: google.firestore.v1beta1.TransactionOptions.read_write:type_name -> google.firestore.v1beta1.TransactionOptions.ReadWrite
	5, // 3: google.firestore.v1beta1.TransactionOptions.ReadOnly.read_time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_firestore_v1beta1_common_proto_init() }
func file_google_firestore_v1beta1_common_proto_init() {
	if File_google_firestore_v1beta1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_firestore_v1beta1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentMask); i {
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
		file_google_firestore_v1beta1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Precondition); i {
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
		file_google_firestore_v1beta1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionOptions); i {
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
		file_google_firestore_v1beta1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionOptions_ReadWrite); i {
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
		file_google_firestore_v1beta1_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionOptions_ReadOnly); i {
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
	file_google_firestore_v1beta1_common_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Precondition_Exists)(nil),
		(*Precondition_UpdateTime)(nil),
	}
	file_google_firestore_v1beta1_common_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*TransactionOptions_ReadOnly_)(nil),
		(*TransactionOptions_ReadWrite_)(nil),
	}
	file_google_firestore_v1beta1_common_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*TransactionOptions_ReadOnly_ReadTime)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_firestore_v1beta1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_firestore_v1beta1_common_proto_goTypes,
		DependencyIndexes: file_google_firestore_v1beta1_common_proto_depIdxs,
		MessageInfos:      file_google_firestore_v1beta1_common_proto_msgTypes,
	}.Build()
	File_google_firestore_v1beta1_common_proto = out.File
	file_google_firestore_v1beta1_common_proto_rawDesc = nil
	file_google_firestore_v1beta1_common_proto_goTypes = nil
	file_google_firestore_v1beta1_common_proto_depIdxs = nil
}
