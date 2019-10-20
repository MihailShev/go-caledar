// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calendar.proto

package calendarpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Event struct {
	UUID                 int64                `protobuf:"varint,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Start                *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	NotifyTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=notifyTime,proto3" json:"notifyTime,omitempty"`
	Description          string               `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	UserId               uint64               `protobuf:"varint,7,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetUUID() int64 {
	if m != nil {
		return m.UUID
	}
	return 0
}

func (m *Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Event) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Event) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *Event) GetNotifyTime() *timestamp.Timestamp {
	if m != nil {
		return m.NotifyTime
	}
	return nil
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateEventReq struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEventReq) Reset()         { *m = CreateEventReq{} }
func (m *CreateEventReq) String() string { return proto.CompactTextString(m) }
func (*CreateEventReq) ProtoMessage()    {}
func (*CreateEventReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{1}
}

func (m *CreateEventReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventReq.Unmarshal(m, b)
}
func (m *CreateEventReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventReq.Marshal(b, m, deterministic)
}
func (m *CreateEventReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventReq.Merge(m, src)
}
func (m *CreateEventReq) XXX_Size() int {
	return xxx_messageInfo_CreateEventReq.Size(m)
}
func (m *CreateEventReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventReq proto.InternalMessageInfo

func (m *CreateEventReq) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type CreateEventRes struct {
	UUID                 int64    `protobuf:"varint,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEventRes) Reset()         { *m = CreateEventRes{} }
func (m *CreateEventRes) String() string { return proto.CompactTextString(m) }
func (*CreateEventRes) ProtoMessage()    {}
func (*CreateEventRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{2}
}

func (m *CreateEventRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventRes.Unmarshal(m, b)
}
func (m *CreateEventRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventRes.Marshal(b, m, deterministic)
}
func (m *CreateEventRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventRes.Merge(m, src)
}
func (m *CreateEventRes) XXX_Size() int {
	return xxx_messageInfo_CreateEventRes.Size(m)
}
func (m *CreateEventRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventRes proto.InternalMessageInfo

func (m *CreateEventRes) GetUUID() int64 {
	if m != nil {
		return m.UUID
	}
	return 0
}

func (m *CreateEventRes) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GetEventReq struct {
	UUID                 int64    `protobuf:"varint,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventReq) Reset()         { *m = GetEventReq{} }
func (m *GetEventReq) String() string { return proto.CompactTextString(m) }
func (*GetEventReq) ProtoMessage()    {}
func (*GetEventReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{3}
}

func (m *GetEventReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventReq.Unmarshal(m, b)
}
func (m *GetEventReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventReq.Marshal(b, m, deterministic)
}
func (m *GetEventReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventReq.Merge(m, src)
}
func (m *GetEventReq) XXX_Size() int {
	return xxx_messageInfo_GetEventReq.Size(m)
}
func (m *GetEventReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventReq proto.InternalMessageInfo

func (m *GetEventReq) GetUUID() int64 {
	if m != nil {
		return m.UUID
	}
	return 0
}

type GetEventRes struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventRes) Reset()         { *m = GetEventRes{} }
func (m *GetEventRes) String() string { return proto.CompactTextString(m) }
func (*GetEventRes) ProtoMessage()    {}
func (*GetEventRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{4}
}

func (m *GetEventRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventRes.Unmarshal(m, b)
}
func (m *GetEventRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventRes.Marshal(b, m, deterministic)
}
func (m *GetEventRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventRes.Merge(m, src)
}
func (m *GetEventRes) XXX_Size() int {
	return xxx_messageInfo_GetEventRes.Size(m)
}
func (m *GetEventRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventRes proto.InternalMessageInfo

func (m *GetEventRes) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *GetEventRes) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type UpdateEventReq struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEventReq) Reset()         { *m = UpdateEventReq{} }
func (m *UpdateEventReq) String() string { return proto.CompactTextString(m) }
func (*UpdateEventReq) ProtoMessage()    {}
func (*UpdateEventReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{5}
}

func (m *UpdateEventReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEventReq.Unmarshal(m, b)
}
func (m *UpdateEventReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEventReq.Marshal(b, m, deterministic)
}
func (m *UpdateEventReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEventReq.Merge(m, src)
}
func (m *UpdateEventReq) XXX_Size() int {
	return xxx_messageInfo_UpdateEventReq.Size(m)
}
func (m *UpdateEventReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEventReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEventReq proto.InternalMessageInfo

func (m *UpdateEventReq) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type UpdateEventRes struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEventRes) Reset()         { *m = UpdateEventRes{} }
func (m *UpdateEventRes) String() string { return proto.CompactTextString(m) }
func (*UpdateEventRes) ProtoMessage()    {}
func (*UpdateEventRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{6}
}

func (m *UpdateEventRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEventRes.Unmarshal(m, b)
}
func (m *UpdateEventRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEventRes.Marshal(b, m, deterministic)
}
func (m *UpdateEventRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEventRes.Merge(m, src)
}
func (m *UpdateEventRes) XXX_Size() int {
	return xxx_messageInfo_UpdateEventRes.Size(m)
}
func (m *UpdateEventRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEventRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEventRes proto.InternalMessageInfo

func (m *UpdateEventRes) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *UpdateEventRes) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "calendarpb.Event")
	proto.RegisterType((*CreateEventReq)(nil), "calendarpb.CreateEventReq")
	proto.RegisterType((*CreateEventRes)(nil), "calendarpb.CreateEventRes")
	proto.RegisterType((*GetEventReq)(nil), "calendarpb.GetEventReq")
	proto.RegisterType((*GetEventRes)(nil), "calendarpb.GetEventRes")
	proto.RegisterType((*UpdateEventReq)(nil), "calendarpb.UpdateEventReq")
	proto.RegisterType((*UpdateEventRes)(nil), "calendarpb.UpdateEventRes")
}

func init() { proto.RegisterFile("calendar.proto", fileDescriptor_e3d25d49f056cdb2) }

var fileDescriptor_e3d25d49f056cdb2 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0x24, 0x6d, 0x53, 0xeb, 0x0b, 0x14, 0x7c, 0x88, 0x2e, 0xb9, 0x18, 0x73, 0x31, 0x07, 0x49,
	0xa5, 0x9e, 0x5a, 0xbc, 0xd5, 0x22, 0x05, 0x41, 0x08, 0xf6, 0x03, 0xd2, 0xe6, 0xb5, 0x04, 0xda,
	0x6c, 0xdc, 0xdd, 0x0a, 0x7e, 0x81, 0x7f, 0xe8, 0xf7, 0x48, 0x36, 0x8d, 0x6e, 0x8b, 0xa1, 0xa2,
	0xb7, 0x7d, 0x3b, 0x33, 0x99, 0x79, 0x93, 0x85, 0xee, 0x3c, 0x5e, 0x51, 0x96, 0xc4, 0x22, 0xcc,
	0x05, 0x57, 0x1c, 0xa1, 0x9a, 0xf3, 0x99, 0x7b, 0xb1, 0xe4, 0x7c, 0xb9, 0xa2, 0x9e, 0x46, 0x66,
	0x9b, 0x45, 0x4f, 0xa5, 0x6b, 0x92, 0x2a, 0x5e, 0xe7, 0x25, 0xd9, 0x7f, 0x6f, 0x80, 0x3d, 0x7e,
	0xa5, 0x4c, 0x21, 0x42, 0x6b, 0x3a, 0x9d, 0xdc, 0x33, 0xcb, 0xb3, 0x82, 0x66, 0xa4, 0xcf, 0x78,
	0x0a, 0xb6, 0x4a, 0xd5, 0x8a, 0x58, 0xc3, 0xb3, 0x82, 0xe3, 0xa8, 0x1c, 0xf0, 0x06, 0x6c, 0xa9,
	0x62, 0xa1, 0x58, 0xd3, 0xb3, 0x02, 0xa7, 0xef, 0x86, 0xa5, 0x49, 0x58, 0x99, 0x84, 0xcf, 0x95,
	0x49, 0x54, 0x12, 0xf1, 0x1a, 0x9a, 0x94, 0x25, 0xac, 0x75, 0x90, 0x5f, 0xd0, 0x70, 0x08, 0x90,
	0x71, 0x95, 0x2e, 0xde, 0x8a, 0x7b, 0x66, 0x1f, 0x14, 0x19, 0x6c, 0xf4, 0xc0, 0x49, 0x48, 0xce,
	0x45, 0x9a, 0xab, 0x94, 0x67, 0xac, 0xad, 0x73, 0x9b, 0x57, 0x78, 0x06, 0xed, 0x8d, 0x24, 0x31,
	0x49, 0xd8, 0x91, 0x67, 0x05, 0xad, 0x68, 0x3b, 0xf9, 0x03, 0xe8, 0x8e, 0x04, 0xc5, 0x8a, 0x74,
	0x1d, 0x11, 0xbd, 0xe0, 0x15, 0xd8, 0x54, 0x9c, 0x75, 0x25, 0x4e, 0xff, 0x24, 0xfc, 0x2e, 0x36,
	0x2c, 0x49, 0x25, 0xee, 0x0f, 0xf7, 0xa4, 0xb2, 0xae, 0x4c, 0x12, 0x82, 0x8b, 0xaa, 0x4c, 0x3d,
	0xf8, 0x97, 0xe0, 0x3c, 0x90, 0xfa, 0xf2, 0xfc, 0x41, 0xe8, 0x3f, 0x9a, 0x14, 0xf9, 0xeb, 0x58,
	0x35, 0x86, 0x03, 0xe8, 0x4e, 0xf3, 0xe4, 0x4f, 0x7b, 0x3e, 0xed, 0x49, 0xff, 0x9b, 0xa5, 0xff,
	0x61, 0x41, 0x67, 0xb4, 0x55, 0xe0, 0x18, 0x1c, 0xa3, 0x45, 0x74, 0xcd, 0x6f, 0xed, 0xfe, 0x19,
	0xb7, 0x1e, 0x93, 0x78, 0x07, 0x9d, 0xaa, 0x2d, 0x3c, 0x37, 0x79, 0x46, 0xcd, 0x6e, 0x0d, 0x20,
	0x8b, 0x10, 0xc6, 0x8a, 0xbb, 0x21, 0x76, 0x6b, 0x73, 0xeb, 0x31, 0x39, 0x6b, 0xeb, 0x67, 0x7a,
	0xfb, 0x19, 0x00, 0x00, 0xff, 0xff, 0xea, 0xc9, 0xd3, 0xa9, 0x9c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarClient interface {
	CreateEvent(ctx context.Context, in *CreateEventReq, opts ...grpc.CallOption) (*CreateEventRes, error)
	GetEvent(ctx context.Context, in *GetEventReq, opts ...grpc.CallOption) (*GetEventRes, error)
	UpdateEvent(ctx context.Context, in *UpdateEventReq, opts ...grpc.CallOption) (*UpdateEventRes, error)
}

type calendarClient struct {
	cc *grpc.ClientConn
}

func NewCalendarClient(cc *grpc.ClientConn) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) CreateEvent(ctx context.Context, in *CreateEventReq, opts ...grpc.CallOption) (*CreateEventRes, error) {
	out := new(CreateEventRes)
	err := c.cc.Invoke(ctx, "/calendarpb.Calendar/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetEvent(ctx context.Context, in *GetEventReq, opts ...grpc.CallOption) (*GetEventRes, error) {
	out := new(GetEventRes)
	err := c.cc.Invoke(ctx, "/calendarpb.Calendar/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) UpdateEvent(ctx context.Context, in *UpdateEventReq, opts ...grpc.CallOption) (*UpdateEventRes, error) {
	out := new(UpdateEventRes)
	err := c.cc.Invoke(ctx, "/calendarpb.Calendar/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
type CalendarServer interface {
	CreateEvent(context.Context, *CreateEventReq) (*CreateEventRes, error)
	GetEvent(context.Context, *GetEventReq) (*GetEventRes, error)
	UpdateEvent(context.Context, *UpdateEventReq) (*UpdateEventRes, error)
}

// UnimplementedCalendarServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (*UnimplementedCalendarServer) CreateEvent(ctx context.Context, req *CreateEventReq) (*CreateEventRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (*UnimplementedCalendarServer) GetEvent(ctx context.Context, req *GetEventReq) (*GetEventRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (*UnimplementedCalendarServer) UpdateEvent(ctx context.Context, req *UpdateEventReq) (*UpdateEventRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}

func RegisterCalendarServer(s *grpc.Server, srv CalendarServer) {
	s.RegisterService(&_Calendar_serviceDesc, srv)
}

func _Calendar_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendarpb.Calendar/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).CreateEvent(ctx, req.(*CreateEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendarpb.Calendar/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetEvent(ctx, req.(*GetEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendarpb.Calendar/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).UpdateEvent(ctx, req.(*UpdateEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calendar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calendarpb.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _Calendar_CreateEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _Calendar_GetEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _Calendar_UpdateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calendar.proto",
}