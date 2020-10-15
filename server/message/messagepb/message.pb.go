// GO: protoc protofiles/message.proto --go_out=plugins=grpc:./server/
// TS: protoc --ts_out="service=grpc-web:owly-client/src" protofiles/message.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: protofiles/message.proto

package messagepb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Attachments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Bucket   string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
}

func (x *Attachments) Reset() {
	*x = Attachments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attachments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachments) ProtoMessage() {}

func (x *Attachments) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachments.ProtoReflect.Descriptor instead.
func (*Attachments) Descriptor() ([]byte, []int) {
	return file_protofiles_message_proto_rawDescGZIP(), []int{0}
}

func (x *Attachments) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Attachments) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorUUID      string         `protobuf:"bytes,1,opt,name=authorUUID,proto3" json:"authorUUID,omitempty"`
	ChatroomID      string         `protobuf:"bytes,2,opt,name=chatroomID,proto3" json:"chatroomID,omitempty"`
	AuthorNAME      string         `protobuf:"bytes,3,opt,name=authorNAME,proto3" json:"authorNAME,omitempty"`
	Content         string         `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp       int64          `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	HasFileAttached bool           `protobuf:"varint,6,opt,name=hasFileAttached,proto3" json:"hasFileAttached,omitempty"`
	Attach          []*Attachments `protobuf:"bytes,7,rep,name=attach,proto3" json:"attach,omitempty"`
	IsAnswer        bool           `protobuf:"varint,8,opt,name=isAnswer,proto3" json:"isAnswer,omitempty"`
	AnswersTo       string         `protobuf:"bytes,9,opt,name=answersTo,proto3" json:"answersTo,omitempty"`
	Id              string         `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_protofiles_message_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetAuthorUUID() string {
	if x != nil {
		return x.AuthorUUID
	}
	return ""
}

func (x *Message) GetChatroomID() string {
	if x != nil {
		return x.ChatroomID
	}
	return ""
}

func (x *Message) GetAuthorNAME() string {
	if x != nil {
		return x.AuthorNAME
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Message) GetHasFileAttached() bool {
	if x != nil {
		return x.HasFileAttached
	}
	return false
}

func (x *Message) GetAttach() []*Attachments {
	if x != nil {
		return x.Attach
	}
	return nil
}

func (x *Message) GetIsAnswer() bool {
	if x != nil {
		return x.IsAnswer
	}
	return false
}

func (x *Message) GetAnswersTo() string {
	if x != nil {
		return x.AnswersTo
	}
	return ""
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StreamMessagesByChatroomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatroomID string `protobuf:"bytes,1,opt,name=chatroomID,proto3" json:"chatroomID,omitempty"`
}

func (x *StreamMessagesByChatroomRequest) Reset() {
	*x = StreamMessagesByChatroomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamMessagesByChatroomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMessagesByChatroomRequest) ProtoMessage() {}

func (x *StreamMessagesByChatroomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMessagesByChatroomRequest.ProtoReflect.Descriptor instead.
func (*StreamMessagesByChatroomRequest) Descriptor() ([]byte, []int) {
	return file_protofiles_message_proto_rawDescGZIP(), []int{2}
}

func (x *StreamMessagesByChatroomRequest) GetChatroomID() string {
	if x != nil {
		return x.ChatroomID
	}
	return ""
}

type StreamMessagesByChatroomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StreamMessagesByChatroomResponse) Reset() {
	*x = StreamMessagesByChatroomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamMessagesByChatroomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMessagesByChatroomResponse) ProtoMessage() {}

func (x *StreamMessagesByChatroomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMessagesByChatroomResponse.ProtoReflect.Descriptor instead.
func (*StreamMessagesByChatroomResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_message_proto_rawDescGZIP(), []int{3}
}

func (x *StreamMessagesByChatroomResponse) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_protofiles_message_proto_rawDescGZIP(), []int{4}
}

func (x *SendMessageRequest) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_message_proto_rawDescGZIP(), []int{5}
}

func (x *SendMessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetMessagesByChatroomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatroomID string `protobuf:"bytes,1,opt,name=chatroomID,proto3" json:"chatroomID,omitempty"`
}

func (x *GetMessagesByChatroomRequest) Reset() {
	*x = GetMessagesByChatroomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesByChatroomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesByChatroomRequest) ProtoMessage() {}

func (x *GetMessagesByChatroomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesByChatroomRequest.ProtoReflect.Descriptor instead.
func (*GetMessagesByChatroomRequest) Descriptor() ([]byte, []int) {
	return file_protofiles_message_proto_rawDescGZIP(), []int{6}
}

func (x *GetMessagesByChatroomRequest) GetChatroomID() string {
	if x != nil {
		return x.ChatroomID
	}
	return ""
}

type GetMessagesByChatroomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *GetMessagesByChatroomResponse) Reset() {
	*x = GetMessagesByChatroomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesByChatroomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesByChatroomResponse) ProtoMessage() {}

func (x *GetMessagesByChatroomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesByChatroomResponse.ProtoReflect.Descriptor instead.
func (*GetMessagesByChatroomResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_message_proto_rawDescGZIP(), []int{7}
}

func (x *GetMessagesByChatroomResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_protofiles_message_proto protoreflect.FileDescriptor

var file_protofiles_message_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x41, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x22, 0xc3, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x55, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x55,
	0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x41, 0x4d, 0x45,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x41,
	0x4d, 0x45, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x28, 0x0a, 0x0f, 0x68, 0x61,
	0x73, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x68, 0x61, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x06, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x1f,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79,
	0x43, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x22,
	0x4e, 0x0a, 0x20, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x40, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x2f, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x3e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x44, 0x22, 0x4d, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x32, 0xbb, 0x02, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x18, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d,
	0x12, 0x28, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x74, 0x72,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4a, 0x0a, 0x0b, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x25,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x43, 0x68, 0x61,
	0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x13, 0x5a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protofiles_message_proto_rawDescOnce sync.Once
	file_protofiles_message_proto_rawDescData = file_protofiles_message_proto_rawDesc
)

func file_protofiles_message_proto_rawDescGZIP() []byte {
	file_protofiles_message_proto_rawDescOnce.Do(func() {
		file_protofiles_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_protofiles_message_proto_rawDescData)
	})
	return file_protofiles_message_proto_rawDescData
}

var file_protofiles_message_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protofiles_message_proto_goTypes = []interface{}{
	(*Attachments)(nil),                      // 0: message.Attachments
	(*Message)(nil),                          // 1: message.Message
	(*StreamMessagesByChatroomRequest)(nil),  // 2: message.StreamMessagesByChatroomRequest
	(*StreamMessagesByChatroomResponse)(nil), // 3: message.StreamMessagesByChatroomResponse
	(*SendMessageRequest)(nil),               // 4: message.SendMessageRequest
	(*SendMessageResponse)(nil),              // 5: message.SendMessageResponse
	(*GetMessagesByChatroomRequest)(nil),     // 6: message.GetMessagesByChatroomRequest
	(*GetMessagesByChatroomResponse)(nil),    // 7: message.GetMessagesByChatroomResponse
}
var file_protofiles_message_proto_depIdxs = []int32{
	0, // 0: message.Message.attach:type_name -> message.Attachments
	1, // 1: message.StreamMessagesByChatroomResponse.message:type_name -> message.Message
	1, // 2: message.SendMessageRequest.message:type_name -> message.Message
	1, // 3: message.GetMessagesByChatroomResponse.messages:type_name -> message.Message
	2, // 4: message.MessageService.StreamMessagesByChatroom:input_type -> message.StreamMessagesByChatroomRequest
	4, // 5: message.MessageService.SendMessage:input_type -> message.SendMessageRequest
	6, // 6: message.MessageService.GetMessagesByChatroom:input_type -> message.GetMessagesByChatroomRequest
	3, // 7: message.MessageService.StreamMessagesByChatroom:output_type -> message.StreamMessagesByChatroomResponse
	5, // 8: message.MessageService.SendMessage:output_type -> message.SendMessageResponse
	7, // 9: message.MessageService.GetMessagesByChatroom:output_type -> message.GetMessagesByChatroomResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protofiles_message_proto_init() }
func file_protofiles_message_proto_init() {
	if File_protofiles_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protofiles_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attachments); i {
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
		file_protofiles_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_protofiles_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamMessagesByChatroomRequest); i {
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
		file_protofiles_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamMessagesByChatroomResponse); i {
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
		file_protofiles_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_protofiles_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
		file_protofiles_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessagesByChatroomRequest); i {
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
		file_protofiles_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessagesByChatroomResponse); i {
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
			RawDescriptor: file_protofiles_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protofiles_message_proto_goTypes,
		DependencyIndexes: file_protofiles_message_proto_depIdxs,
		MessageInfos:      file_protofiles_message_proto_msgTypes,
	}.Build()
	File_protofiles_message_proto = out.File
	file_protofiles_message_proto_rawDesc = nil
	file_protofiles_message_proto_goTypes = nil
	file_protofiles_message_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageServiceClient interface {
	StreamMessagesByChatroom(ctx context.Context, in *StreamMessagesByChatroomRequest, opts ...grpc.CallOption) (MessageService_StreamMessagesByChatroomClient, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	GetMessagesByChatroom(ctx context.Context, in *GetMessagesByChatroomRequest, opts ...grpc.CallOption) (*GetMessagesByChatroomResponse, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) StreamMessagesByChatroom(ctx context.Context, in *StreamMessagesByChatroomRequest, opts ...grpc.CallOption) (MessageService_StreamMessagesByChatroomClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MessageService_serviceDesc.Streams[0], "/message.MessageService/StreamMessagesByChatroom", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceStreamMessagesByChatroomClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageService_StreamMessagesByChatroomClient interface {
	Recv() (*StreamMessagesByChatroomResponse, error)
	grpc.ClientStream
}

type messageServiceStreamMessagesByChatroomClient struct {
	grpc.ClientStream
}

func (x *messageServiceStreamMessagesByChatroomClient) Recv() (*StreamMessagesByChatroomResponse, error) {
	m := new(StreamMessagesByChatroomResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetMessagesByChatroom(ctx context.Context, in *GetMessagesByChatroomRequest, opts ...grpc.CallOption) (*GetMessagesByChatroomResponse, error) {
	out := new(GetMessagesByChatroomResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/GetMessagesByChatroom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
type MessageServiceServer interface {
	StreamMessagesByChatroom(*StreamMessagesByChatroomRequest, MessageService_StreamMessagesByChatroomServer) error
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	GetMessagesByChatroom(context.Context, *GetMessagesByChatroomRequest) (*GetMessagesByChatroomResponse, error)
}

// UnimplementedMessageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (*UnimplementedMessageServiceServer) StreamMessagesByChatroom(*StreamMessagesByChatroomRequest, MessageService_StreamMessagesByChatroomServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMessagesByChatroom not implemented")
}
func (*UnimplementedMessageServiceServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (*UnimplementedMessageServiceServer) GetMessagesByChatroom(context.Context, *GetMessagesByChatroomRequest) (*GetMessagesByChatroomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessagesByChatroom not implemented")
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_StreamMessagesByChatroom_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamMessagesByChatroomRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageServiceServer).StreamMessagesByChatroom(m, &messageServiceStreamMessagesByChatroomServer{stream})
}

type MessageService_StreamMessagesByChatroomServer interface {
	Send(*StreamMessagesByChatroomResponse) error
	grpc.ServerStream
}

type messageServiceStreamMessagesByChatroomServer struct {
	grpc.ServerStream
}

func (x *messageServiceStreamMessagesByChatroomServer) Send(m *StreamMessagesByChatroomResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MessageService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetMessagesByChatroom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesByChatroomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetMessagesByChatroom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/GetMessagesByChatroom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetMessagesByChatroom(ctx, req.(*GetMessagesByChatroomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _MessageService_SendMessage_Handler,
		},
		{
			MethodName: "GetMessagesByChatroom",
			Handler:    _MessageService_GetMessagesByChatroom_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMessagesByChatroom",
			Handler:       _MessageService_StreamMessagesByChatroom_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protofiles/message.proto",
}
