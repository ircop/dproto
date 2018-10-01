// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dproto/request.proto

package dproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BoxRequest struct {
	RequestID            string      `protobuf:"bytes,1,opt,name=RequestID,proto3" json:"RequestID,omitempty"`
	Profile              ProfileType `protobuf:"varint,2,opt,name=Profile,proto3,enum=dproto.ProfileType" json:"Profile,omitempty"`
	Proto                Protocol    `protobuf:"varint,3,opt,name=Proto,proto3,enum=dproto.Protocol" json:"Proto,omitempty"`
	Host                 string      `protobuf:"bytes,4,opt,name=Host,proto3" json:"Host,omitempty"`
	Port                 int64       `protobuf:"varint,5,opt,name=Port,proto3" json:"Port,omitempty"`
	Login                string      `protobuf:"bytes,6,opt,name=Login,proto3" json:"Login,omitempty"`
	Password             string      `protobuf:"bytes,7,opt,name=Password,proto3" json:"Password,omitempty"`
	Enable               string      `protobuf:"bytes,8,opt,name=Enable,proto3" json:"Enable,omitempty"`
	Timeout              int32       `protobuf:"varint,9,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BoxRequest) Reset()         { *m = BoxRequest{} }
func (m *BoxRequest) String() string { return proto.CompactTextString(m) }
func (*BoxRequest) ProtoMessage()    {}
func (*BoxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_976614bd856be8da, []int{0}
}
func (m *BoxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoxRequest.Unmarshal(m, b)
}
func (m *BoxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoxRequest.Marshal(b, m, deterministic)
}
func (dst *BoxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoxRequest.Merge(dst, src)
}
func (m *BoxRequest) XXX_Size() int {
	return xxx_messageInfo_BoxRequest.Size(m)
}
func (m *BoxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BoxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BoxRequest proto.InternalMessageInfo

func (m *BoxRequest) GetRequestID() string {
	if m != nil {
		return m.RequestID
	}
	return ""
}

func (m *BoxRequest) GetProfile() ProfileType {
	if m != nil {
		return m.Profile
	}
	return ProfileType_DXS
}

func (m *BoxRequest) GetProto() Protocol {
	if m != nil {
		return m.Proto
	}
	return Protocol_NONE
}

func (m *BoxRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *BoxRequest) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *BoxRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *BoxRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *BoxRequest) GetEnable() string {
	if m != nil {
		return m.Enable
	}
	return ""
}

func (m *BoxRequest) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func init() {
	proto.RegisterType((*BoxRequest)(nil), "dproto.BoxRequest")
}

func init() { proto.RegisterFile("dproto/request.proto", fileDescriptor_request_976614bd856be8da) }

var fileDescriptor_request_976614bd856be8da = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x49, 0x3b, 0xbf, 0x57, 0x90, 0x72, 0xad, 0x72, 0x29, 0x2e, 0x06, 0x17, 0x32, 0x1b,
	0x47, 0xd0, 0x37, 0x10, 0x05, 0x15, 0x17, 0x43, 0xe8, 0x0b, 0xb4, 0x36, 0xca, 0xc0, 0xd8, 0x3b,
	0x26, 0x29, 0xda, 0xe7, 0xf0, 0x85, 0xa5, 0x37, 0xa9, 0x76, 0x77, 0xbe, 0xf3, 0x9d, 0xfc, 0xc0,
	0x74, 0x35, 0x58, 0xf6, 0x7c, 0x6d, 0xcd, 0xe7, 0xc6, 0x38, 0xdf, 0x08, 0x61, 0x16, 0xda, 0x19,
	0x46, 0xeb, 0xb7, 0x83, 0x71, 0xc1, 0x5d, 0xfc, 0x8c, 0x00, 0xee, 0xf8, 0x5b, 0x87, 0x03, 0x78,
	0x0e, 0x65, 0x8c, 0x4f, 0xf7, 0xa4, 0x2a, 0x55, 0x97, 0xfa, 0xbf, 0xc0, 0x2b, 0xc8, 0x5b, 0xcb,
	0x6f, 0x5d, 0x6f, 0x68, 0x54, 0xa9, 0xfa, 0xf8, 0xe6, 0xa4, 0x09, 0x57, 0x36, 0xb1, 0x9e, 0x6f,
	0x07, 0xa3, 0xf7, 0x1b, 0xbc, 0x84, 0xb4, 0xdd, 0x59, 0x1a, 0xcb, 0x78, 0x72, 0x30, 0xf6, 0xfc,
	0xca, 0xbd, 0x0e, 0x1a, 0x11, 0x92, 0x47, 0x76, 0x9e, 0x12, 0x79, 0x4f, 0xf2, 0xae, 0x6b, 0xd9,
	0x7a, 0x4a, 0x2b, 0x55, 0x8f, 0xb5, 0x64, 0x9c, 0x42, 0xfa, 0xc2, 0xef, 0xdd, 0x9a, 0x32, 0x19,
	0x06, 0xc0, 0x19, 0x14, 0xed, 0xc2, 0xb9, 0x2f, 0xb6, 0x2b, 0xca, 0x45, 0xfc, 0x31, 0x9e, 0x41,
	0xf6, 0xb0, 0x5e, 0x2c, 0x7b, 0x43, 0x85, 0x98, 0x48, 0x48, 0x90, 0xcf, 0xbb, 0x0f, 0xc3, 0x1b,
	0x4f, 0x65, 0xa5, 0xea, 0x54, 0xef, 0xf1, 0x39, 0x29, 0x8e, 0x26, 0xa7, 0xcb, 0x4c, 0x3e, 0x7a,
	0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xb3, 0xb2, 0x49, 0x50, 0x01, 0x00, 0x00,
}
