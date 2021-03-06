// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: raza_shell.proto

package raza

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

type StartSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host            string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	User            string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	StartTimeEpochS int64  `protobuf:"varint,3,opt,name=start_time_epoch_s,json=startTimeEpochS,proto3" json:"start_time_epoch_s,omitempty"`
}

func (x *StartSessionRequest) Reset() {
	*x = StartSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raza_shell_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartSessionRequest) ProtoMessage() {}

func (x *StartSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raza_shell_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartSessionRequest.ProtoReflect.Descriptor instead.
func (*StartSessionRequest) Descriptor() ([]byte, []int) {
	return file_raza_shell_proto_rawDescGZIP(), []int{0}
}

func (x *StartSessionRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *StartSessionRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *StartSessionRequest) GetStartTimeEpochS() int64 {
	if x != nil {
		return x.StartTimeEpochS
	}
	return 0
}

type StartSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId int64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *StartSessionResponse) Reset() {
	*x = StartSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raza_shell_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartSessionResponse) ProtoMessage() {}

func (x *StartSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raza_shell_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartSessionResponse.ProtoReflect.Descriptor instead.
func (*StartSessionResponse) Descriptor() ([]byte, []int) {
	return file_raza_shell_proto_rawDescGZIP(), []int{1}
}

func (x *StartSessionResponse) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

type EndSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId     int64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	EndTimeEpochS int64 `protobuf:"varint,2,opt,name=end_time_epoch_s,json=endTimeEpochS,proto3" json:"end_time_epoch_s,omitempty"`
}

func (x *EndSessionRequest) Reset() {
	*x = EndSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raza_shell_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndSessionRequest) ProtoMessage() {}

func (x *EndSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raza_shell_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndSessionRequest.ProtoReflect.Descriptor instead.
func (*EndSessionRequest) Descriptor() ([]byte, []int) {
	return file_raza_shell_proto_rawDescGZIP(), []int{2}
}

func (x *EndSessionRequest) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *EndSessionRequest) GetEndTimeEpochS() int64 {
	if x != nil {
		return x.EndTimeEpochS
	}
	return 0
}

type EndSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *EndSessionResponse) Reset() {
	*x = EndSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raza_shell_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndSessionResponse) ProtoMessage() {}

func (x *EndSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raza_shell_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndSessionResponse.ProtoReflect.Descriptor instead.
func (*EndSessionResponse) Descriptor() ([]byte, []int) {
	return file_raza_shell_proto_rawDescGZIP(), []int{3}
}

func (x *EndSessionResponse) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type StartCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId       int64            `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Cmd             string           `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Pwd             string           `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	StartTimeEpochS int64            `protobuf:"varint,4,opt,name=start_time_epoch_s,json=startTimeEpochS,proto3" json:"start_time_epoch_s,omitempty"`
	Metadata        []*MetadataEntry `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *StartCommandRequest) Reset() {
	*x = StartCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raza_shell_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartCommandRequest) ProtoMessage() {}

func (x *StartCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raza_shell_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartCommandRequest.ProtoReflect.Descriptor instead.
func (*StartCommandRequest) Descriptor() ([]byte, []int) {
	return file_raza_shell_proto_rawDescGZIP(), []int{4}
}

func (x *StartCommandRequest) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *StartCommandRequest) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *StartCommandRequest) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

func (x *StartCommandRequest) GetStartTimeEpochS() int64 {
	if x != nil {
		return x.StartTimeEpochS
	}
	return 0
}

func (x *StartCommandRequest) GetMetadata() []*MetadataEntry {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type StartCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandId int64 `protobuf:"varint,1,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
}

func (x *StartCommandResponse) Reset() {
	*x = StartCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raza_shell_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartCommandResponse) ProtoMessage() {}

func (x *StartCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raza_shell_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartCommandResponse.ProtoReflect.Descriptor instead.
func (*StartCommandResponse) Descriptor() ([]byte, []int) {
	return file_raza_shell_proto_rawDescGZIP(), []int{5}
}

func (x *StartCommandResponse) GetCommandId() int64 {
	if x != nil {
		return x.CommandId
	}
	return 0
}

type EndSessionsLastCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId     int64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Retval        int32 `protobuf:"varint,2,opt,name=retval,proto3" json:"retval,omitempty"`
	EndTimeEpochS int64 `protobuf:"varint,3,opt,name=end_time_epoch_s,json=endTimeEpochS,proto3" json:"end_time_epoch_s,omitempty"`
}

func (x *EndSessionsLastCommandRequest) Reset() {
	*x = EndSessionsLastCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raza_shell_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndSessionsLastCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndSessionsLastCommandRequest) ProtoMessage() {}

func (x *EndSessionsLastCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raza_shell_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndSessionsLastCommandRequest.ProtoReflect.Descriptor instead.
func (*EndSessionsLastCommandRequest) Descriptor() ([]byte, []int) {
	return file_raza_shell_proto_rawDescGZIP(), []int{6}
}

func (x *EndSessionsLastCommandRequest) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *EndSessionsLastCommandRequest) GetRetval() int32 {
	if x != nil {
		return x.Retval
	}
	return 0
}

func (x *EndSessionsLastCommandRequest) GetEndTimeEpochS() int64 {
	if x != nil {
		return x.EndTimeEpochS
	}
	return 0
}

type EndCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command *Command `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *EndCommandResponse) Reset() {
	*x = EndCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raza_shell_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndCommandResponse) ProtoMessage() {}

func (x *EndCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raza_shell_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndCommandResponse.ProtoReflect.Descriptor instead.
func (*EndCommandResponse) Descriptor() ([]byte, []int) {
	return file_raza_shell_proto_rawDescGZIP(), []int{7}
}

func (x *EndCommandResponse) GetCommand() *Command {
	if x != nil {
		return x.Command
	}
	return nil
}

var File_raza_shell_proto protoreflect.FileDescriptor

var file_raza_shell_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x61, 0x7a, 0x61, 0x5f, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x72, 0x61, 0x7a, 0x61, 0x1a, 0x13, 0x72, 0x61, 0x7a, 0x61, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a,
	0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x12,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x5f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x53, 0x22, 0x35, 0x0a, 0x14, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0x5b, 0x0a, 0x11, 0x45, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x10, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x53, 0x22, 0x3d, 0x0a,
	0x12, 0x45, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x61, 0x7a, 0x61, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xb6, 0x01, 0x0a,
	0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x77, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x70, 0x77, 0x64, 0x12, 0x2b, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x53, 0x12, 0x2f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x61, 0x7a, 0x61, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x35, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x7f, 0x0a, 0x1d,
	0x45, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x74, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65,
	0x74, 0x76, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x10, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x53, 0x22, 0x3d, 0x0a,
	0x12, 0x45, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x61, 0x7a, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x32, 0xc2, 0x02, 0x0a,
	0x10, 0x52, 0x61, 0x7a, 0x61, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x12, 0x47, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x2e, 0x72, 0x61, 0x7a, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72,
	0x61, 0x7a, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x45, 0x6e,
	0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x72, 0x61, 0x7a, 0x61, 0x2e,
	0x45, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x72, 0x61, 0x7a, 0x61, 0x2e, 0x45, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a,
	0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x19, 0x2e,
	0x72, 0x61, 0x7a, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x61, 0x7a, 0x61, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x16, 0x45, 0x6e, 0x64, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x23, 0x2e, 0x72, 0x61, 0x7a, 0x61, 0x2e, 0x45, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x61, 0x7a, 0x61, 0x2e, 0x45, 0x6e, 0x64,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0a, 0x5a, 0x08, 0x6c, 0x69, 0x62, 0x2f, 0x72, 0x61, 0x7a, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_raza_shell_proto_rawDescOnce sync.Once
	file_raza_shell_proto_rawDescData = file_raza_shell_proto_rawDesc
)

func file_raza_shell_proto_rawDescGZIP() []byte {
	file_raza_shell_proto_rawDescOnce.Do(func() {
		file_raza_shell_proto_rawDescData = protoimpl.X.CompressGZIP(file_raza_shell_proto_rawDescData)
	})
	return file_raza_shell_proto_rawDescData
}

var file_raza_shell_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_raza_shell_proto_goTypes = []interface{}{
	(*StartSessionRequest)(nil),           // 0: raza.StartSessionRequest
	(*StartSessionResponse)(nil),          // 1: raza.StartSessionResponse
	(*EndSessionRequest)(nil),             // 2: raza.EndSessionRequest
	(*EndSessionResponse)(nil),            // 3: raza.EndSessionResponse
	(*StartCommandRequest)(nil),           // 4: raza.StartCommandRequest
	(*StartCommandResponse)(nil),          // 5: raza.StartCommandResponse
	(*EndSessionsLastCommandRequest)(nil), // 6: raza.EndSessionsLastCommandRequest
	(*EndCommandResponse)(nil),            // 7: raza.EndCommandResponse
	(*Session)(nil),                       // 8: raza.Session
	(*MetadataEntry)(nil),                 // 9: raza.MetadataEntry
	(*Command)(nil),                       // 10: raza.Command
}
var file_raza_shell_proto_depIdxs = []int32{
	8,  // 0: raza.EndSessionResponse.session:type_name -> raza.Session
	9,  // 1: raza.StartCommandRequest.metadata:type_name -> raza.MetadataEntry
	10, // 2: raza.EndCommandResponse.command:type_name -> raza.Command
	0,  // 3: raza.RazaShellWrapper.StartSession:input_type -> raza.StartSessionRequest
	2,  // 4: raza.RazaShellWrapper.EndSession:input_type -> raza.EndSessionRequest
	4,  // 5: raza.RazaShellWrapper.StartCommand:input_type -> raza.StartCommandRequest
	6,  // 6: raza.RazaShellWrapper.EndSessionsLastCommand:input_type -> raza.EndSessionsLastCommandRequest
	1,  // 7: raza.RazaShellWrapper.StartSession:output_type -> raza.StartSessionResponse
	3,  // 8: raza.RazaShellWrapper.EndSession:output_type -> raza.EndSessionResponse
	5,  // 9: raza.RazaShellWrapper.StartCommand:output_type -> raza.StartCommandResponse
	7,  // 10: raza.RazaShellWrapper.EndSessionsLastCommand:output_type -> raza.EndCommandResponse
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_raza_shell_proto_init() }
func file_raza_shell_proto_init() {
	if File_raza_shell_proto != nil {
		return
	}
	file_raza_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_raza_shell_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartSessionRequest); i {
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
		file_raza_shell_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartSessionResponse); i {
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
		file_raza_shell_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndSessionRequest); i {
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
		file_raza_shell_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndSessionResponse); i {
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
		file_raza_shell_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartCommandRequest); i {
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
		file_raza_shell_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartCommandResponse); i {
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
		file_raza_shell_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndSessionsLastCommandRequest); i {
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
		file_raza_shell_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndCommandResponse); i {
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
			RawDescriptor: file_raza_shell_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_raza_shell_proto_goTypes,
		DependencyIndexes: file_raza_shell_proto_depIdxs,
		MessageInfos:      file_raza_shell_proto_msgTypes,
	}.Build()
	File_raza_shell_proto = out.File
	file_raza_shell_proto_rawDesc = nil
	file_raza_shell_proto_goTypes = nil
	file_raza_shell_proto_depIdxs = nil
}
