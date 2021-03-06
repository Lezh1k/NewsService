// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type RegisterItemErrorCode int32

const (
	RegisterItemErrorCode_SUCCESS         RegisterItemErrorCode = 0
	RegisterItemErrorCode_INCORRECT_INPUT RegisterItemErrorCode = 1
	RegisterItemErrorCode_SERVER_ERROR    RegisterItemErrorCode = 2
)

var RegisterItemErrorCode_name = map[int32]string{
	0: "SUCCESS",
	1: "INCORRECT_INPUT",
	2: "SERVER_ERROR",
}

var RegisterItemErrorCode_value = map[string]int32{
	"SUCCESS":         0,
	"INCORRECT_INPUT": 1,
	"SERVER_ERROR":    2,
}

func (x RegisterItemErrorCode) String() string {
	return proto.EnumName(RegisterItemErrorCode_name, int32(x))
}

func (RegisterItemErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

type NewsItem struct {
	Header               string               `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NewsItem) Reset()         { *m = NewsItem{} }
func (m *NewsItem) String() string { return proto.CompactTextString(m) }
func (*NewsItem) ProtoMessage()    {}
func (*NewsItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *NewsItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewsItem.Unmarshal(m, b)
}
func (m *NewsItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewsItem.Marshal(b, m, deterministic)
}
func (m *NewsItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewsItem.Merge(m, src)
}
func (m *NewsItem) XXX_Size() int {
	return xxx_messageInfo_NewsItem.Size(m)
}
func (m *NewsItem) XXX_DiscardUnknown() {
	xxx_messageInfo_NewsItem.DiscardUnknown(m)
}

var xxx_messageInfo_NewsItem proto.InternalMessageInfo

func (m *NewsItem) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *NewsItem) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

// оборачиваем запросы/ответы в контейнеры, чтоб можно было игнорировать не относящиеся к конкретному запросу ответы в NATS
// запрос Register
type NewsRegisterRequestContainer struct {
	Guid                 string    `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Item                 *NewsItem `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *NewsRegisterRequestContainer) Reset()         { *m = NewsRegisterRequestContainer{} }
func (m *NewsRegisterRequestContainer) String() string { return proto.CompactTextString(m) }
func (*NewsRegisterRequestContainer) ProtoMessage()    {}
func (*NewsRegisterRequestContainer) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *NewsRegisterRequestContainer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewsRegisterRequestContainer.Unmarshal(m, b)
}
func (m *NewsRegisterRequestContainer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewsRegisterRequestContainer.Marshal(b, m, deterministic)
}
func (m *NewsRegisterRequestContainer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewsRegisterRequestContainer.Merge(m, src)
}
func (m *NewsRegisterRequestContainer) XXX_Size() int {
	return xxx_messageInfo_NewsRegisterRequestContainer.Size(m)
}
func (m *NewsRegisterRequestContainer) XXX_DiscardUnknown() {
	xxx_messageInfo_NewsRegisterRequestContainer.DiscardUnknown(m)
}

var xxx_messageInfo_NewsRegisterRequestContainer proto.InternalMessageInfo

func (m *NewsRegisterRequestContainer) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *NewsRegisterRequestContainer) GetItem() *NewsItem {
	if m != nil {
		return m.Item
	}
	return nil
}

// ответ на запрос Register
type NewsRegisterResponseContainer struct {
	Guid                 string   `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewsRegisterResponseContainer) Reset()         { *m = NewsRegisterResponseContainer{} }
func (m *NewsRegisterResponseContainer) String() string { return proto.CompactTextString(m) }
func (*NewsRegisterResponseContainer) ProtoMessage()    {}
func (*NewsRegisterResponseContainer) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *NewsRegisterResponseContainer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewsRegisterResponseContainer.Unmarshal(m, b)
}
func (m *NewsRegisterResponseContainer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewsRegisterResponseContainer.Marshal(b, m, deterministic)
}
func (m *NewsRegisterResponseContainer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewsRegisterResponseContainer.Merge(m, src)
}
func (m *NewsRegisterResponseContainer) XXX_Size() int {
	return xxx_messageInfo_NewsRegisterResponseContainer.Size(m)
}
func (m *NewsRegisterResponseContainer) XXX_DiscardUnknown() {
	xxx_messageInfo_NewsRegisterResponseContainer.DiscardUnknown(m)
}

var xxx_messageInfo_NewsRegisterResponseContainer proto.InternalMessageInfo

func (m *NewsRegisterResponseContainer) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *NewsRegisterResponseContainer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// запрос на инфо по новости
type NewsInfoRequestContainer struct {
	Guid                 string   `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewsInfoRequestContainer) Reset()         { *m = NewsInfoRequestContainer{} }
func (m *NewsInfoRequestContainer) String() string { return proto.CompactTextString(m) }
func (*NewsInfoRequestContainer) ProtoMessage()    {}
func (*NewsInfoRequestContainer) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *NewsInfoRequestContainer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewsInfoRequestContainer.Unmarshal(m, b)
}
func (m *NewsInfoRequestContainer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewsInfoRequestContainer.Marshal(b, m, deterministic)
}
func (m *NewsInfoRequestContainer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewsInfoRequestContainer.Merge(m, src)
}
func (m *NewsInfoRequestContainer) XXX_Size() int {
	return xxx_messageInfo_NewsInfoRequestContainer.Size(m)
}
func (m *NewsInfoRequestContainer) XXX_DiscardUnknown() {
	xxx_messageInfo_NewsInfoRequestContainer.DiscardUnknown(m)
}

var xxx_messageInfo_NewsInfoRequestContainer proto.InternalMessageInfo

func (m *NewsInfoRequestContainer) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *NewsInfoRequestContainer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// ответ на инфо о новости
type NewsInfoResponseContainer struct {
	Guid                 string    `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Item                 *NewsItem `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *NewsInfoResponseContainer) Reset()         { *m = NewsInfoResponseContainer{} }
func (m *NewsInfoResponseContainer) String() string { return proto.CompactTextString(m) }
func (*NewsInfoResponseContainer) ProtoMessage()    {}
func (*NewsInfoResponseContainer) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *NewsInfoResponseContainer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewsInfoResponseContainer.Unmarshal(m, b)
}
func (m *NewsInfoResponseContainer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewsInfoResponseContainer.Marshal(b, m, deterministic)
}
func (m *NewsInfoResponseContainer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewsInfoResponseContainer.Merge(m, src)
}
func (m *NewsInfoResponseContainer) XXX_Size() int {
	return xxx_messageInfo_NewsInfoResponseContainer.Size(m)
}
func (m *NewsInfoResponseContainer) XXX_DiscardUnknown() {
	xxx_messageInfo_NewsInfoResponseContainer.DiscardUnknown(m)
}

var xxx_messageInfo_NewsInfoResponseContainer proto.InternalMessageInfo

func (m *NewsInfoResponseContainer) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *NewsInfoResponseContainer) GetItem() *NewsItem {
	if m != nil {
		return m.Item
	}
	return nil
}

func init() {
	proto.RegisterEnum("message.RegisterItemErrorCode", RegisterItemErrorCode_name, RegisterItemErrorCode_value)
	proto.RegisterType((*NewsItem)(nil), "message.NewsItem")
	proto.RegisterType((*NewsRegisterRequestContainer)(nil), "message.NewsRegisterRequestContainer")
	proto.RegisterType((*NewsRegisterResponseContainer)(nil), "message.NewsRegisterResponseContainer")
	proto.RegisterType((*NewsInfoRequestContainer)(nil), "message.NewsInfoRequestContainer")
	proto.RegisterType((*NewsInfoResponseContainer)(nil), "message.NewsInfoResponseContainer")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xbf, 0x09, 0xa5, 0xfd, 0x76, 0xea, 0x8f, 0xb8, 0xa2, 0xd4, 0xa2, 0x58, 0x02, 0x42,
	0xf1, 0x90, 0x82, 0xde, 0xbd, 0x2c, 0x7b, 0x28, 0x42, 0x2b, 0x93, 0xb6, 0xe0, 0xa9, 0xa4, 0xec,
	0x34, 0x2e, 0x98, 0x6c, 0xdd, 0xdd, 0xe2, 0xbf, 0x2f, 0xdd, 0x26, 0x88, 0x1e, 0x24, 0xb7, 0x9d,
	0xd9, 0x99, 0xcf, 0x7b, 0x8f, 0x81, 0xe3, 0x82, 0xac, 0xcd, 0x72, 0x4a, 0xb6, 0x46, 0x3b, 0xcd,
	0x3a, 0x55, 0x39, 0xb8, 0xcd, 0xb5, 0xce, 0xdf, 0x69, 0xec, 0xdb, 0xeb, 0xdd, 0x66, 0xec, 0x54,
	0x41, 0xd6, 0x65, 0xc5, 0xf6, 0x30, 0x19, 0x23, 0xfc, 0x9f, 0xd2, 0xa7, 0x9d, 0x38, 0x2a, 0xd8,
	0x25, 0xb4, 0xdf, 0x28, 0x93, 0x64, 0xfa, 0xc1, 0x30, 0x18, 0x75, 0xb1, 0xaa, 0x58, 0x02, 0x2d,
	0x99, 0x39, 0xea, 0x87, 0xc3, 0x60, 0xd4, 0x7b, 0x18, 0x24, 0x07, 0x66, 0x52, 0x33, 0x93, 0x79,
	0xcd, 0x44, 0x3f, 0x17, 0xbf, 0xc2, 0xf5, 0x9e, 0x89, 0x94, 0x2b, 0xeb, 0xc8, 0x20, 0x7d, 0xec,
	0xc8, 0x3a, 0xae, 0x4b, 0x97, 0xa9, 0x92, 0x0c, 0x63, 0xd0, 0xca, 0x77, 0x4a, 0x56, 0x2a, 0xfe,
	0xcd, 0xee, 0xa0, 0xa5, 0x1c, 0x15, 0x95, 0xc6, 0x59, 0x52, 0xe7, 0xa9, 0xcd, 0xa1, 0xff, 0x8e,
	0x39, 0xdc, 0xfc, 0x44, 0xdb, 0xad, 0x2e, 0x2d, 0xfd, 0xcd, 0x3e, 0x81, 0x50, 0x49, 0x4f, 0xee,
	0x62, 0xa8, 0x64, 0xfc, 0x04, 0x7d, 0x8f, 0x2d, 0x37, 0xba, 0x91, 0xb7, 0xdf, 0xfb, 0x4b, 0xb8,
	0xfa, 0xde, 0x6f, 0x62, 0xa0, 0x59, 0xb8, 0xfb, 0x67, 0xb8, 0xa8, 0x83, 0xed, 0xbb, 0xc2, 0x18,
	0x6d, 0xb8, 0x96, 0xc4, 0x7a, 0xd0, 0x49, 0x17, 0x9c, 0x8b, 0x34, 0x8d, 0xfe, 0xb1, 0x73, 0x38,
	0x9d, 0x4c, 0xf9, 0x0c, 0x51, 0xf0, 0xf9, 0x6a, 0x32, 0x7d, 0x59, 0xcc, 0xa3, 0x80, 0x45, 0x70,
	0x94, 0x0a, 0x5c, 0x0a, 0x5c, 0x09, 0xc4, 0x19, 0x46, 0xe1, 0xba, 0xed, 0xcf, 0xf3, 0xf8, 0x15,
	0x00, 0x00, 0xff, 0xff, 0x9f, 0x70, 0x2b, 0x7d, 0x1a, 0x02, 0x00, 0x00,
}
