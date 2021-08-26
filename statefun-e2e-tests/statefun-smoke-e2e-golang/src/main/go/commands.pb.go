//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: commands.proto

package main

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

type SourceCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target   int32     `protobuf:"varint,1,opt,name=target,proto3" json:"target,omitempty"`
	Commands *Commands `protobuf:"bytes,2,opt,name=commands,proto3" json:"commands,omitempty"`
}

func (x *SourceCommand) Reset() {
	*x = SourceCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceCommand) ProtoMessage() {}

func (x *SourceCommand) ProtoReflect() protoreflect.Message {
	mi := &file_commands_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceCommand.ProtoReflect.Descriptor instead.
func (*SourceCommand) Descriptor() ([]byte, []int) {
	return file_commands_proto_rawDescGZIP(), []int{0}
}

func (x *SourceCommand) GetTarget() int32 {
	if x != nil {
		return x.Target
	}
	return 0
}

func (x *SourceCommand) GetCommands() *Commands {
	if x != nil {
		return x.Commands
	}
	return nil
}

type Commands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command []*Command `protobuf:"bytes,1,rep,name=command,proto3" json:"command,omitempty"`
}

func (x *Commands) Reset() {
	*x = Commands{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commands) ProtoMessage() {}

func (x *Commands) ProtoReflect() protoreflect.Message {
	mi := &file_commands_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commands.ProtoReflect.Descriptor instead.
func (*Commands) Descriptor() ([]byte, []int) {
	return file_commands_proto_rawDescGZIP(), []int{1}
}

func (x *Commands) GetCommand() []*Command {
	if x != nil {
		return x.Command
	}
	return nil
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Command:
	//	*Command_Increment
	//	*Command_Send_
	//	*Command_SendAfter_
	//	*Command_SendEgress_
	//	*Command_AsyncOperation_
	//	*Command_Verify_
	Command isCommand_Command `protobuf_oneof:"command"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_commands_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_commands_proto_rawDescGZIP(), []int{2}
}

func (m *Command) GetCommand() isCommand_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (x *Command) GetIncrement() *Command_IncrementState {
	if x, ok := x.GetCommand().(*Command_Increment); ok {
		return x.Increment
	}
	return nil
}

func (x *Command) GetSend() *Command_Send {
	if x, ok := x.GetCommand().(*Command_Send_); ok {
		return x.Send
	}
	return nil
}

func (x *Command) GetSendAfter() *Command_SendAfter {
	if x, ok := x.GetCommand().(*Command_SendAfter_); ok {
		return x.SendAfter
	}
	return nil
}

func (x *Command) GetSendEgress() *Command_SendEgress {
	if x, ok := x.GetCommand().(*Command_SendEgress_); ok {
		return x.SendEgress
	}
	return nil
}

func (x *Command) GetAsyncOperation() *Command_AsyncOperation {
	if x, ok := x.GetCommand().(*Command_AsyncOperation_); ok {
		return x.AsyncOperation
	}
	return nil
}

func (x *Command) GetVerify() *Command_Verify {
	if x, ok := x.GetCommand().(*Command_Verify_); ok {
		return x.Verify
	}
	return nil
}

type isCommand_Command interface {
	isCommand_Command()
}

type Command_Increment struct {
	Increment *Command_IncrementState `protobuf:"bytes,1,opt,name=increment,proto3,oneof"`
}

type Command_Send_ struct {
	Send *Command_Send `protobuf:"bytes,2,opt,name=send,proto3,oneof"`
}

type Command_SendAfter_ struct {
	SendAfter *Command_SendAfter `protobuf:"bytes,3,opt,name=send_after,json=sendAfter,proto3,oneof"`
}

type Command_SendEgress_ struct {
	SendEgress *Command_SendEgress `protobuf:"bytes,4,opt,name=send_egress,json=sendEgress,proto3,oneof"`
}

type Command_AsyncOperation_ struct {
	AsyncOperation *Command_AsyncOperation `protobuf:"bytes,5,opt,name=async_operation,json=asyncOperation,proto3,oneof"`
}

type Command_Verify_ struct {
	Verify *Command_Verify `protobuf:"bytes,6,opt,name=verify,proto3,oneof"`
}

func (*Command_Increment) isCommand_Command() {}

func (*Command_Send_) isCommand_Command() {}

func (*Command_SendAfter_) isCommand_Command() {}

func (*Command_SendEgress_) isCommand_Command() {}

func (*Command_AsyncOperation_) isCommand_Command() {}

func (*Command_Verify_) isCommand_Command() {}

type VerificationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Expected int64 `protobuf:"varint,2,opt,name=expected,proto3" json:"expected,omitempty"`
	Actual   int64 `protobuf:"varint,3,opt,name=actual,proto3" json:"actual,omitempty"`
}

func (x *VerificationResult) Reset() {
	*x = VerificationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerificationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationResult) ProtoMessage() {}

func (x *VerificationResult) ProtoReflect() protoreflect.Message {
	mi := &file_commands_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationResult.ProtoReflect.Descriptor instead.
func (*VerificationResult) Descriptor() ([]byte, []int) {
	return file_commands_proto_rawDescGZIP(), []int{3}
}

func (x *VerificationResult) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VerificationResult) GetExpected() int64 {
	if x != nil {
		return x.Expected
	}
	return 0
}

func (x *VerificationResult) GetActual() int64 {
	if x != nil {
		return x.Actual
	}
	return 0
}

type Command_IncrementState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Command_IncrementState) Reset() {
	*x = Command_IncrementState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command_IncrementState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command_IncrementState) ProtoMessage() {}

func (x *Command_IncrementState) ProtoReflect() protoreflect.Message {
	mi := &file_commands_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command_IncrementState.ProtoReflect.Descriptor instead.
func (*Command_IncrementState) Descriptor() ([]byte, []int) {
	return file_commands_proto_rawDescGZIP(), []int{2, 0}
}

type Command_Send struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target   int32     `protobuf:"varint,1,opt,name=target,proto3" json:"target,omitempty"`
	Commands *Commands `protobuf:"bytes,2,opt,name=commands,proto3" json:"commands,omitempty"`
}

func (x *Command_Send) Reset() {
	*x = Command_Send{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command_Send) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command_Send) ProtoMessage() {}

func (x *Command_Send) ProtoReflect() protoreflect.Message {
	mi := &file_commands_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command_Send.ProtoReflect.Descriptor instead.
func (*Command_Send) Descriptor() ([]byte, []int) {
	return file_commands_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Command_Send) GetTarget() int32 {
	if x != nil {
		return x.Target
	}
	return 0
}

func (x *Command_Send) GetCommands() *Commands {
	if x != nil {
		return x.Commands
	}
	return nil
}

type Command_SendAfter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target   int32     `protobuf:"varint,1,opt,name=target,proto3" json:"target,omitempty"`
	Commands *Commands `protobuf:"bytes,2,opt,name=commands,proto3" json:"commands,omitempty"`
}

func (x *Command_SendAfter) Reset() {
	*x = Command_SendAfter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command_SendAfter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command_SendAfter) ProtoMessage() {}

func (x *Command_SendAfter) ProtoReflect() protoreflect.Message {
	mi := &file_commands_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command_SendAfter.ProtoReflect.Descriptor instead.
func (*Command_SendAfter) Descriptor() ([]byte, []int) {
	return file_commands_proto_rawDescGZIP(), []int{2, 2}
}

func (x *Command_SendAfter) GetTarget() int32 {
	if x != nil {
		return x.Target
	}
	return 0
}

func (x *Command_SendAfter) GetCommands() *Commands {
	if x != nil {
		return x.Commands
	}
	return nil
}

type Command_SendEgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Command_SendEgress) Reset() {
	*x = Command_SendEgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command_SendEgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command_SendEgress) ProtoMessage() {}

func (x *Command_SendEgress) ProtoReflect() protoreflect.Message {
	mi := &file_commands_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command_SendEgress.ProtoReflect.Descriptor instead.
func (*Command_SendEgress) Descriptor() ([]byte, []int) {
	return file_commands_proto_rawDescGZIP(), []int{2, 3}
}

type Command_AsyncOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failure          bool      `protobuf:"varint,1,opt,name=failure,proto3" json:"failure,omitempty"`
	ResolvedCommands *Commands `protobuf:"bytes,2,opt,name=resolved_commands,json=resolvedCommands,proto3" json:"resolved_commands,omitempty"`
}

func (x *Command_AsyncOperation) Reset() {
	*x = Command_AsyncOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command_AsyncOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command_AsyncOperation) ProtoMessage() {}

func (x *Command_AsyncOperation) ProtoReflect() protoreflect.Message {
	mi := &file_commands_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command_AsyncOperation.ProtoReflect.Descriptor instead.
func (*Command_AsyncOperation) Descriptor() ([]byte, []int) {
	return file_commands_proto_rawDescGZIP(), []int{2, 4}
}

func (x *Command_AsyncOperation) GetFailure() bool {
	if x != nil {
		return x.Failure
	}
	return false
}

func (x *Command_AsyncOperation) GetResolvedCommands() *Commands {
	if x != nil {
		return x.ResolvedCommands
	}
	return nil
}

type Command_Verify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expected int64 `protobuf:"varint,1,opt,name=expected,proto3" json:"expected,omitempty"`
}

func (x *Command_Verify) Reset() {
	*x = Command_Verify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commands_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command_Verify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command_Verify) ProtoMessage() {}

func (x *Command_Verify) ProtoReflect() protoreflect.Message {
	mi := &file_commands_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command_Verify.ProtoReflect.Descriptor instead.
func (*Command_Verify) Descriptor() ([]byte, []int) {
	return file_commands_proto_rawDescGZIP(), []int{2, 5}
}

func (x *Command_Verify) GetExpected() int64 {
	if x != nil {
		return x.Expected
	}
	return 0
}

var File_commands_proto protoreflect.FileDescriptor

var file_commands_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x23, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x66, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6e, 0x2e, 0x65, 0x32, 0x65, 0x2e,
	0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x22, 0x72, 0x0a, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x49,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x66, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6e, 0x2e, 0x65, 0x32, 0x65,
	0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x52, 0x0a, 0x08, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x46, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61,
	0x63, 0x68, 0x65, 0x2e, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66,
	0x75, 0x6e, 0x2e, 0x65, 0x32, 0x65, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0xd0, 0x07,
	0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x5b, 0x0a, 0x09, 0x69, 0x6e, 0x63,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6e, 0x2e, 0x65, 0x32, 0x65, 0x2e, 0x73, 0x6d, 0x6f,
	0x6b, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x09, 0x69, 0x6e, 0x63,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6e,
	0x2e, 0x65, 0x32, 0x65, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x12,
	0x57, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6e, 0x2e,
	0x65, 0x32, 0x65, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x09, 0x73,
	0x65, 0x6e, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x5a, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64,
	0x5f, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x66, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6e, 0x2e, 0x65, 0x32, 0x65, 0x2e, 0x73, 0x6d,
	0x6f, 0x6b, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x45, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x66, 0x0a, 0x0f, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x66, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6e, 0x2e, 0x65, 0x32, 0x65, 0x2e, 0x73, 0x6d,
	0x6f, 0x6b, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x41, 0x73, 0x79, 0x6e,
	0x63, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x61, 0x73,
	0x79, 0x6e, 0x63, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x06,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6e, 0x2e, 0x65, 0x32, 0x65, 0x2e, 0x73, 0x6d, 0x6f,
	0x6b, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x48, 0x00, 0x52, 0x06, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x1a, 0x10, 0x0a, 0x0e, 0x49,
	0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x69, 0x0a,
	0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x49, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x66, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6e, 0x2e, 0x65, 0x32, 0x65, 0x2e,
	0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x1a, 0x6e, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64,
	0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x49, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x66, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6e, 0x2e, 0x65, 0x32, 0x65, 0x2e,
	0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x1a, 0x0c, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64,
	0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x86, 0x01, 0x0a, 0x0e, 0x41, 0x73, 0x79, 0x6e, 0x63,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x12, 0x5a, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x66, 0x6c, 0x69, 0x6e,
	0x6b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6e, 0x2e, 0x65, 0x32, 0x65, 0x2e, 0x73,
	0x6d, 0x6f, 0x6b, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x10, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x1a,
	0x24, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x70,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x22, 0x58, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b,
	0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commands_proto_rawDescOnce sync.Once
	file_commands_proto_rawDescData = file_commands_proto_rawDesc
)

func file_commands_proto_rawDescGZIP() []byte {
	file_commands_proto_rawDescOnce.Do(func() {
		file_commands_proto_rawDescData = protoimpl.X.CompressGZIP(file_commands_proto_rawDescData)
	})
	return file_commands_proto_rawDescData
}

var file_commands_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_commands_proto_goTypes = []interface{}{
	(*SourceCommand)(nil),          // 0: org.apache.flink.statefun.e2e.smoke.SourceCommand
	(*Commands)(nil),               // 1: org.apache.flink.statefun.e2e.smoke.Commands
	(*Command)(nil),                // 2: org.apache.flink.statefun.e2e.smoke.Command
	(*VerificationResult)(nil),     // 3: org.apache.flink.statefun.e2e.smoke.VerificationResult
	(*Command_IncrementState)(nil), // 4: org.apache.flink.statefun.e2e.smoke.Command.IncrementState
	(*Command_Send)(nil),           // 5: org.apache.flink.statefun.e2e.smoke.Command.Send
	(*Command_SendAfter)(nil),      // 6: org.apache.flink.statefun.e2e.smoke.Command.SendAfter
	(*Command_SendEgress)(nil),     // 7: org.apache.flink.statefun.e2e.smoke.Command.SendEgress
	(*Command_AsyncOperation)(nil), // 8: org.apache.flink.statefun.e2e.smoke.Command.AsyncOperation
	(*Command_Verify)(nil),         // 9: org.apache.flink.statefun.e2e.smoke.Command.Verify
}
var file_commands_proto_depIdxs = []int32{
	1,  // 0: org.apache.flink.statefun.e2e.smoke.SourceCommand.commands:type_name -> org.apache.flink.statefun.e2e.smoke.Commands
	2,  // 1: org.apache.flink.statefun.e2e.smoke.Commands.command:type_name -> org.apache.flink.statefun.e2e.smoke.Command
	4,  // 2: org.apache.flink.statefun.e2e.smoke.Command.increment:type_name -> org.apache.flink.statefun.e2e.smoke.Command.IncrementState
	5,  // 3: org.apache.flink.statefun.e2e.smoke.Command.send:type_name -> org.apache.flink.statefun.e2e.smoke.Command.Send
	6,  // 4: org.apache.flink.statefun.e2e.smoke.Command.send_after:type_name -> org.apache.flink.statefun.e2e.smoke.Command.SendAfter
	7,  // 5: org.apache.flink.statefun.e2e.smoke.Command.send_egress:type_name -> org.apache.flink.statefun.e2e.smoke.Command.SendEgress
	8,  // 6: org.apache.flink.statefun.e2e.smoke.Command.async_operation:type_name -> org.apache.flink.statefun.e2e.smoke.Command.AsyncOperation
	9,  // 7: org.apache.flink.statefun.e2e.smoke.Command.verify:type_name -> org.apache.flink.statefun.e2e.smoke.Command.Verify
	1,  // 8: org.apache.flink.statefun.e2e.smoke.Command.Send.commands:type_name -> org.apache.flink.statefun.e2e.smoke.Commands
	1,  // 9: org.apache.flink.statefun.e2e.smoke.Command.SendAfter.commands:type_name -> org.apache.flink.statefun.e2e.smoke.Commands
	1,  // 10: org.apache.flink.statefun.e2e.smoke.Command.AsyncOperation.resolved_commands:type_name -> org.apache.flink.statefun.e2e.smoke.Commands
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_commands_proto_init() }
func file_commands_proto_init() {
	if File_commands_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commands_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceCommand); i {
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
		file_commands_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commands); i {
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
		file_commands_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_commands_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerificationResult); i {
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
		file_commands_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command_IncrementState); i {
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
		file_commands_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command_Send); i {
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
		file_commands_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command_SendAfter); i {
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
		file_commands_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command_SendEgress); i {
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
		file_commands_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command_AsyncOperation); i {
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
		file_commands_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command_Verify); i {
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
	file_commands_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Command_Increment)(nil),
		(*Command_Send_)(nil),
		(*Command_SendAfter_)(nil),
		(*Command_SendEgress_)(nil),
		(*Command_AsyncOperation_)(nil),
		(*Command_Verify_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commands_proto_goTypes,
		DependencyIndexes: file_commands_proto_depIdxs,
		MessageInfos:      file_commands_proto_msgTypes,
	}.Build()
	File_commands_proto = out.File
	file_commands_proto_rawDesc = nil
	file_commands_proto_goTypes = nil
	file_commands_proto_depIdxs = nil
}