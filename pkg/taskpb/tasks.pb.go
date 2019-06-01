// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tasks.proto

package taskpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Task describe todo task
type Task struct {
	// Unique task identifier
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// title of the task
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// details description of the task
	Desc string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	// remind time for todo task
	Reminder             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=reminder,proto3" json:"reminder,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3834c8ef8464a3f, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Task) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Task) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Task) GetReminder() *timestamp.Timestamp {
	if m != nil {
		return m.Reminder
	}
	return nil
}

// Status describe status of performed action
type Status struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3834c8ef8464a3f, []int{1}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CreateRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3834c8ef8464a3f, []int{2}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateRequest) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type CreateResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3834c8ef8464a3f, []int{3}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3834c8ef8464a3f, []int{4}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadResponse struct {
	Task                 *Task    `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3834c8ef8464a3f, []int{5}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type UpdateRequest struct {
	Task                 *Task    `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3834c8ef8464a3f, []int{6}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type UpdateResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3834c8ef8464a3f, []int{7}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type DeleteRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3834c8ef8464a3f, []int{8}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3834c8ef8464a3f, []int{9}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReadAllResponse struct {
	Task                 []*Task  `protobuf:"bytes,1,rep,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllResponse) Reset()         { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()    {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3834c8ef8464a3f, []int{10}
}

func (m *ReadAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllResponse.Unmarshal(m, b)
}
func (m *ReadAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllResponse.Marshal(b, m, deterministic)
}
func (m *ReadAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllResponse.Merge(m, src)
}
func (m *ReadAllResponse) XXX_Size() int {
	return xxx_messageInfo_ReadAllResponse.Size(m)
}
func (m *ReadAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllResponse proto.InternalMessageInfo

func (m *ReadAllResponse) GetTask() []*Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func init() {
	proto.RegisterType((*Task)(nil), "taskprotobuf.Task")
	proto.RegisterType((*Status)(nil), "taskprotobuf.Status")
	proto.RegisterType((*CreateRequest)(nil), "taskprotobuf.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "taskprotobuf.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "taskprotobuf.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "taskprotobuf.ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "taskprotobuf.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "taskprotobuf.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "taskprotobuf.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "taskprotobuf.DeleteResponse")
	proto.RegisterType((*ReadAllResponse)(nil), "taskprotobuf.ReadAllResponse")
}

func init() { proto.RegisterFile("tasks.proto", fileDescriptor_b3834c8ef8464a3f) }

var fileDescriptor_b3834c8ef8464a3f = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x5d, 0x8f, 0x93, 0x50,
	0x10, 0x0d, 0x14, 0xd9, 0x75, 0xd8, 0x62, 0x72, 0xb3, 0x31, 0xc8, 0xba, 0x59, 0xc2, 0x83, 0xe9,
	0x83, 0x61, 0x93, 0x9a, 0xd4, 0xec, 0x83, 0x8d, 0x8a, 0xfe, 0x01, 0x5a, 0x5f, 0x7c, 0xa3, 0x65,
	0x6c, 0x48, 0xa1, 0x20, 0xf7, 0xd6, 0xe8, 0x8f, 0xf5, 0xbf, 0x98, 0xfb, 0x65, 0xe1, 0x96, 0x18,
	0xb3, 0x6f, 0x77, 0x98, 0x73, 0x66, 0xce, 0x9c, 0x13, 0xc0, 0x63, 0x39, 0xdd, 0xd3, 0xa4, 0xed,
	0x1a, 0xd6, 0x90, 0x2b, 0x5e, 0x88, 0xe7, 0xe6, 0xf8, 0x2d, 0xbc, 0xdb, 0x35, 0xcd, 0xae, 0xc2,
	0x7b, 0xfd, 0xe1, 0x9e, 0x95, 0x35, 0x52, 0x96, 0xd7, 0xad, 0x84, 0x87, 0x37, 0x26, 0x00, 0xeb,
	0x96, 0xfd, 0x92, 0xcd, 0xf8, 0x27, 0x38, 0xeb, 0x9c, 0xee, 0x89, 0x0f, 0x76, 0x59, 0x04, 0x56,
	0x64, 0xcd, 0x26, 0x99, 0x5d, 0x16, 0xe4, 0x1a, 0x9e, 0xb0, 0x92, 0x55, 0x18, 0xd8, 0x91, 0x35,
	0x7b, 0x9a, 0xc9, 0x82, 0x10, 0x70, 0x0a, 0xa4, 0xdb, 0x60, 0x22, 0x3e, 0x8a, 0x37, 0x59, 0xc0,
	0x65, 0x87, 0x75, 0x79, 0x28, 0xb0, 0x0b, 0x9c, 0xc8, 0x9a, 0x79, 0xf3, 0x30, 0x91, 0x1b, 0x13,
	0xbd, 0x31, 0x59, 0x6b, 0x49, 0xd9, 0x5f, 0x6c, 0xbc, 0x00, 0x77, 0xc5, 0x72, 0x76, 0xa4, 0x7c,
	0xea, 0xb6, 0x29, 0x50, 0x6d, 0x17, 0x6f, 0x12, 0xc0, 0x45, 0x8d, 0x94, 0xe6, 0x3b, 0xad, 0x40,
	0x97, 0xf1, 0x03, 0x4c, 0xd3, 0x0e, 0x73, 0x86, 0x19, 0x7e, 0x3f, 0x22, 0x65, 0x27, 0xa9, 0xd6,
	0x98, 0x54, 0xfb, 0x24, 0x35, 0x8e, 0xc0, 0xd7, 0x54, 0xda, 0x36, 0x07, 0x8a, 0xe6, 0xd9, 0xf1,
	0x2d, 0x78, 0x19, 0xe6, 0x85, 0x1e, 0x6d, 0xb6, 0x17, 0x70, 0x25, 0xdb, 0x8a, 0xfe, 0x0a, 0x1c,
	0x9e, 0x85, 0x40, 0x78, 0x73, 0x92, 0xf4, 0x83, 0x49, 0xb8, 0xaf, 0x99, 0xe8, 0xc7, 0x6f, 0x61,
	0xfa, 0xa5, 0x2d, 0x7a, 0x9a, 0xff, 0x97, 0xb8, 0x04, 0x5f, 0x13, 0xd5, 0xca, 0xd7, 0xe0, 0x52,
	0x61, 0x9b, 0xe2, 0x5e, 0x0f, 0xb9, 0xd2, 0xd2, 0x4c, 0x61, 0xe2, 0x3b, 0x98, 0x7e, 0xc2, 0x0a,
	0x4f, 0x8b, 0xcd, 0x8b, 0x96, 0xe0, 0x6b, 0xc0, 0xa3, 0x16, 0x3c, 0xc0, 0x33, 0xee, 0xc8, 0x87,
	0xaa, 0x1a, 0x31, 0x65, 0xf2, 0xaf, 0xdb, 0xe6, 0xbf, 0x6d, 0xf0, 0x78, 0xb9, 0xc2, 0xee, 0x47,
	0xb9, 0x45, 0x92, 0x82, 0x2b, 0xd3, 0x21, 0x37, 0x43, 0xce, 0x20, 0xee, 0xf0, 0xe5, 0x78, 0x53,
	0x2d, 0x7f, 0x07, 0x0e, 0xd7, 0x43, 0x5e, 0x0c, 0x51, 0xbd, 0x50, 0xc3, 0x70, 0xac, 0xa5, 0xe8,
	0x29, 0xb8, 0xd2, 0x6f, 0x53, 0xc3, 0x20, 0x3e, 0x53, 0x83, 0x11, 0x51, 0x0a, 0xae, 0xf4, 0xd4,
	0x1c, 0x32, 0x88, 0xc2, 0x1c, 0x62, 0xc4, 0xf0, 0x1e, 0x2e, 0x94, 0xb1, 0xe4, 0xf9, 0xd9, 0xff,
	0xf4, 0x99, 0xff, 0xc1, 0xe1, 0xed, 0xf9, 0x21, 0xbd, 0x1c, 0x3e, 0x5e, 0x7e, 0x75, 0x45, 0x7f,
	0xb3, 0x71, 0x05, 0xe6, 0xcd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x21, 0x79, 0x65, 0x46,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskServiceClient interface {
	// Create creates new todo task
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Read todo task
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	// Update todo task
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete todo task
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Read all todo tasks
	ReadAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadAllResponse, error)
}

type taskServiceClient struct {
	cc *grpc.ClientConn
}

func NewTaskServiceClient(cc *grpc.ClientConn) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/taskprotobuf.TaskService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/taskprotobuf.TaskService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/taskprotobuf.TaskService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/taskprotobuf.TaskService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ReadAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/taskprotobuf.TaskService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
type TaskServiceServer interface {
	// Create creates new todo task
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Read todo task
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	// Update todo task
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete todo task
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Read all todo tasks
	ReadAll(context.Context, *empty.Empty) (*ReadAllResponse, error)
}

// UnimplementedTaskServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (*UnimplementedTaskServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedTaskServiceServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedTaskServiceServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedTaskServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedTaskServiceServer) ReadAll(ctx context.Context, req *empty.Empty) (*ReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

func _TaskService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taskprotobuf.TaskService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taskprotobuf.TaskService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taskprotobuf.TaskService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taskprotobuf.TaskService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/taskprotobuf.TaskService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ReadAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "taskprotobuf.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TaskService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _TaskService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TaskService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TaskService_Delete_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _TaskService_ReadAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tasks.proto",
}