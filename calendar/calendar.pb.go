// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calendar.proto

package calendar

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Start                int64    `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	End                  int64    `protobuf:"varint,4,opt,name=end,proto3" json:"end,omitempty"`
	Id                   uint32   `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Event) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Event) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type EventList struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventList) Reset()         { *m = EventList{} }
func (m *EventList) String() string { return proto.CompactTextString(m) }
func (*EventList) ProtoMessage()    {}
func (*EventList) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{1}
}

func (m *EventList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventList.Unmarshal(m, b)
}
func (m *EventList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventList.Marshal(b, m, deterministic)
}
func (m *EventList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventList.Merge(m, src)
}
func (m *EventList) XXX_Size() int {
	return xxx_messageInfo_EventList.Size(m)
}
func (m *EventList) XXX_DiscardUnknown() {
	xxx_messageInfo_EventList.DiscardUnknown(m)
}

var xxx_messageInfo_EventList proto.InternalMessageInfo

func (m *EventList) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "calendar.Event")
	proto.RegisterType((*EventList)(nil), "calendar.EventList")
}

func init() { proto.RegisterFile("calendar.proto", fileDescriptor_e3d25d49f056cdb2) }

var fileDescriptor_e3d25d49f056cdb2 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xbd, 0x0e, 0xc2, 0x20,
	0x14, 0x85, 0x43, 0x69, 0x1b, 0x7b, 0x1b, 0xab, 0xb9, 0x71, 0x60, 0x24, 0x5d, 0x64, 0xea, 0xa0,
	0xbe, 0x82, 0x9b, 0x13, 0x6f, 0x80, 0x85, 0x81, 0x44, 0x69, 0x03, 0xe8, 0xf3, 0x9b, 0x5e, 0x7f,
	0xe2, 0x76, 0xbe, 0x73, 0x6e, 0x72, 0x3f, 0xe8, 0x46, 0x73, 0x73, 0xc1, 0x9a, 0x38, 0xcc, 0x71,
	0xca, 0x13, 0xae, 0xbe, 0xdc, 0x3f, 0xa0, 0x3a, 0x3f, 0x5d, 0xc8, 0x88, 0x50, 0x06, 0x73, 0x77,
	0x82, 0x49, 0xa6, 0x1a, 0x4d, 0x19, 0x25, 0xb4, 0xd6, 0xa5, 0x31, 0xfa, 0x39, 0xfb, 0x29, 0x88,
	0x82, 0xa6, 0xff, 0x0a, 0x77, 0x50, 0xa5, 0x6c, 0x62, 0x16, 0x5c, 0x32, 0xc5, 0xf5, 0x1b, 0x70,
	0x0b, 0xdc, 0x05, 0x2b, 0x4a, 0xea, 0x96, 0x88, 0x1d, 0x14, 0xde, 0x8a, 0x4a, 0x32, 0xb5, 0xd6,
	0x85, 0xb7, 0xfd, 0x09, 0x1a, 0x7a, 0x7b, 0xf1, 0x29, 0xe3, 0x1e, 0x6a, 0xb7, 0x40, 0x12, 0x4c,
	0x72, 0xd5, 0x1e, 0x36, 0xc3, 0x4f, 0x97, 0x8e, 0xf4, 0x67, 0xbe, 0xd6, 0x64, 0x7f, 0x7c, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x58, 0xd8, 0xa4, 0xe3, 0xcf, 0x00, 0x00, 0x00,
}
