// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dproto/packet_types.proto

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

type PacketType int32

const (
	PacketType_ERROR        PacketType = 0
	PacketType_BOX_REQUEST  PacketType = 1
	PacketType_BOX_REPLY    PacketType = 2
	PacketType_BOX_ERROR    PacketType = 3
	PacketType_SYNC_REQUEST PacketType = 4
	PacketType_SYNC_REPLY   PacketType = 5
)

var PacketType_name = map[int32]string{
	0: "ERROR",
	1: "BOX_REQUEST",
	2: "BOX_REPLY",
	3: "BOX_ERROR",
	4: "SYNC_REQUEST",
	5: "SYNC_REPLY",
}
var PacketType_value = map[string]int32{
	"ERROR":        0,
	"BOX_REQUEST":  1,
	"BOX_REPLY":    2,
	"BOX_ERROR":    3,
	"SYNC_REQUEST": 4,
	"SYNC_REPLY":   5,
}

func (x PacketType) String() string {
	return proto.EnumName(PacketType_name, int32(x))
}
func (PacketType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_packet_types_48008c22c0c3bb12, []int{0}
}

func init() {
	proto.RegisterEnum("dproto.PacketType", PacketType_name, PacketType_value)
}

func init() {
	proto.RegisterFile("dproto/packet_types.proto", fileDescriptor_packet_types_48008c22c0c3bb12)
}

var fileDescriptor_packet_types_48008c22c0c3bb12 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x29, 0x28, 0xca,
	0x2f, 0xc9, 0xd7, 0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0x89, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6,
	0x03, 0x0b, 0x09, 0xb1, 0x41, 0xa4, 0xb4, 0x32, 0xb8, 0xb8, 0x02, 0xc0, 0xb2, 0x21, 0x95, 0x05,
	0xa9, 0x42, 0x9c, 0x5c, 0xac, 0xae, 0x41, 0x41, 0xfe, 0x41, 0x02, 0x0c, 0x42, 0xfc, 0x5c, 0xdc,
	0x4e, 0xfe, 0x11, 0xf1, 0x41, 0xae, 0x81, 0xa1, 0xae, 0xc1, 0x21, 0x02, 0x8c, 0x42, 0xbc, 0x5c,
	0x9c, 0x10, 0x81, 0x00, 0x9f, 0x48, 0x01, 0x26, 0x18, 0x17, 0xa2, 0x9c, 0x59, 0x48, 0x80, 0x8b,
	0x27, 0x38, 0xd2, 0xcf, 0x19, 0xae, 0x9e, 0x45, 0x88, 0x8f, 0x8b, 0x0b, 0x2a, 0x02, 0xd2, 0xc0,
	0x9a, 0xc4, 0x06, 0xb6, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x78, 0xd5, 0xae, 0x74, 0x95,
	0x00, 0x00, 0x00,
}
