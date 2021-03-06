// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dproto/results.proto

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

// Platform is GetPlatform result
type Platform struct {
	// @inject_tag: json:"model,omitempty"
	Model string `protobuf:"bytes,1,opt,name=Model,proto3" json:"model,omitempty"`
	// @inject_tag: json:"revision,omitempty"
	Revision string `protobuf:"bytes,2,opt,name=Revision,proto3" json:"revision,omitempty"`
	// @inject_tag: json:"version,omitempty"
	Version string `protobuf:"bytes,3,opt,name=Version,proto3" json:"version,omitempty"`
	// @inject_tag: json:"serial,omitempty"
	Serial string `protobuf:"bytes,4,opt,name=Serial,proto3" json:"serial,omitempty"`
	// @inject_tag: json:"macs"
	Macs                 []string `protobuf:"bytes,5,rep,name=Macs,proto3" json:"macs"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Platform) Reset()         { *m = Platform{} }
func (m *Platform) String() string { return proto.CompactTextString(m) }
func (*Platform) ProtoMessage()    {}
func (*Platform) Descriptor() ([]byte, []int) {
	return fileDescriptor_results_f15d742ddb5a6112, []int{0}
}
func (m *Platform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform.Unmarshal(m, b)
}
func (m *Platform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform.Marshal(b, m, deterministic)
}
func (dst *Platform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform.Merge(dst, src)
}
func (m *Platform) XXX_Size() int {
	return xxx_messageInfo_Platform.Size(m)
}
func (m *Platform) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform.DiscardUnknown(m)
}

var xxx_messageInfo_Platform proto.InternalMessageInfo

func (m *Platform) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Platform) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *Platform) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Platform) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func (m *Platform) GetMacs() []string {
	if m != nil {
		return m.Macs
	}
	return nil
}

// Interface is for GetInterfaces() result
type Interface struct {
	// @inject_tag: json:"type,omitempty"
	Type InterfaceType `protobuf:"varint,1,opt,name=Type,proto3,enum=dproto.InterfaceType" json:"type,omitempty"`
	// @inject_tag: json:"name,omitempty"
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"name,omitempty"`
	// @inject_tag: json:"shortname,omitempty"
	Shortname string `protobuf:"bytes,3,opt,name=Shortname,proto3" json:"shortname,omitempty"`
	// @inject_tag: json:"description,omitempty"
	Description string `protobuf:"bytes,4,opt,name=Description,proto3" json:"description,omitempty"`
	// @inject_tag: json:"lldp_id,omitempty"
	LldpID string `protobuf:"bytes,5,opt,name=LldpID,proto3" json:"lldp_id,omitempty"`
	// @inject_tag: json:"po_members,omitempty"
	PoMembers            []string `protobuf:"bytes,6,rep,name=PoMembers,proto3" json:"po_members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Interface) Reset()         { *m = Interface{} }
func (m *Interface) String() string { return proto.CompactTextString(m) }
func (*Interface) ProtoMessage()    {}
func (*Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_results_f15d742ddb5a6112, []int{1}
}
func (m *Interface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Interface.Unmarshal(m, b)
}
func (m *Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Interface.Marshal(b, m, deterministic)
}
func (dst *Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Interface.Merge(dst, src)
}
func (m *Interface) XXX_Size() int {
	return xxx_messageInfo_Interface.Size(m)
}
func (m *Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_Interface.DiscardUnknown(m)
}

var xxx_messageInfo_Interface proto.InternalMessageInfo

func (m *Interface) GetType() InterfaceType {
	if m != nil {
		return m.Type
	}
	return InterfaceType_UNKNOWN
}

func (m *Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Interface) GetShortname() string {
	if m != nil {
		return m.Shortname
	}
	return ""
}

func (m *Interface) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Interface) GetLldpID() string {
	if m != nil {
		return m.LldpID
	}
	return ""
}

func (m *Interface) GetPoMembers() []string {
	if m != nil {
		return m.PoMembers
	}
	return nil
}

// LldpNeighbor for GetLldp()
type LldpNeighbor struct {
	// @inject_tag: json:"local_port,omitempty"
	LocalPort string `protobuf:"bytes,1,opt,name=LocalPort,proto3" json:"local_port,omitempty"`
	// @inject_tag: json:"chassis_id,omitempty"
	ChassisID string `protobuf:"bytes,2,opt,name=ChassisID,proto3" json:"chassis_id,omitempty"`
	// @inject_tag: json:"port_id,omitempty"
	PortID               string   `protobuf:"bytes,3,opt,name=PortID,proto3" json:"port_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LldpNeighbor) Reset()         { *m = LldpNeighbor{} }
func (m *LldpNeighbor) String() string { return proto.CompactTextString(m) }
func (*LldpNeighbor) ProtoMessage()    {}
func (*LldpNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_results_f15d742ddb5a6112, []int{2}
}
func (m *LldpNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpNeighbor.Unmarshal(m, b)
}
func (m *LldpNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpNeighbor.Marshal(b, m, deterministic)
}
func (dst *LldpNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpNeighbor.Merge(dst, src)
}
func (m *LldpNeighbor) XXX_Size() int {
	return xxx_messageInfo_LldpNeighbor.Size(m)
}
func (m *LldpNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_LldpNeighbor proto.InternalMessageInfo

func (m *LldpNeighbor) GetLocalPort() string {
	if m != nil {
		return m.LocalPort
	}
	return ""
}

func (m *LldpNeighbor) GetChassisID() string {
	if m != nil {
		return m.ChassisID
	}
	return ""
}

func (m *LldpNeighbor) GetPortID() string {
	if m != nil {
		return m.PortID
	}
	return ""
}

// Vlan for GetVlans()
type Vlan struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	ID                   int64    `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	TrunkPorts           []string `protobuf:"bytes,3,rep,name=TrunkPorts,proto3" json:"TrunkPorts,omitempty"`
	AccessPorts          []string `protobuf:"bytes,4,rep,name=AccessPorts,proto3" json:"AccessPorts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vlan) Reset()         { *m = Vlan{} }
func (m *Vlan) String() string { return proto.CompactTextString(m) }
func (*Vlan) ProtoMessage()    {}
func (*Vlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_results_f15d742ddb5a6112, []int{3}
}
func (m *Vlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vlan.Unmarshal(m, b)
}
func (m *Vlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vlan.Marshal(b, m, deterministic)
}
func (dst *Vlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vlan.Merge(dst, src)
}
func (m *Vlan) XXX_Size() int {
	return xxx_messageInfo_Vlan.Size(m)
}
func (m *Vlan) XXX_DiscardUnknown() {
	xxx_messageInfo_Vlan.DiscardUnknown(m)
}

var xxx_messageInfo_Vlan proto.InternalMessageInfo

func (m *Vlan) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vlan) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Vlan) GetTrunkPorts() []string {
	if m != nil {
		return m.TrunkPorts
	}
	return nil
}

func (m *Vlan) GetAccessPorts() []string {
	if m != nil {
		return m.AccessPorts
	}
	return nil
}

// Ipif is ip interface for GetIps()
type Ipif struct {
	Interface            string   `protobuf:"bytes,1,opt,name=Interface,proto3" json:"Interface,omitempty"`
	IP                   string   `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	Mask                 string   `protobuf:"bytes,3,opt,name=Mask,proto3" json:"Mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipif) Reset()         { *m = Ipif{} }
func (m *Ipif) String() string { return proto.CompactTextString(m) }
func (*Ipif) ProtoMessage()    {}
func (*Ipif) Descriptor() ([]byte, []int) {
	return fileDescriptor_results_f15d742ddb5a6112, []int{4}
}
func (m *Ipif) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipif.Unmarshal(m, b)
}
func (m *Ipif) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipif.Marshal(b, m, deterministic)
}
func (dst *Ipif) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipif.Merge(dst, src)
}
func (m *Ipif) XXX_Size() int {
	return xxx_messageInfo_Ipif.Size(m)
}
func (m *Ipif) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipif.DiscardUnknown(m)
}

var xxx_messageInfo_Ipif proto.InternalMessageInfo

func (m *Ipif) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *Ipif) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *Ipif) GetMask() string {
	if m != nil {
		return m.Mask
	}
	return ""
}

func init() {
	proto.RegisterType((*Platform)(nil), "dproto.Platform")
	proto.RegisterType((*Interface)(nil), "dproto.Interface")
	proto.RegisterType((*LldpNeighbor)(nil), "dproto.LldpNeighbor")
	proto.RegisterType((*Vlan)(nil), "dproto.Vlan")
	proto.RegisterType((*Ipif)(nil), "dproto.Ipif")
}

func init() { proto.RegisterFile("dproto/results.proto", fileDescriptor_results_f15d742ddb5a6112) }

var fileDescriptor_results_f15d742ddb5a6112 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xdb, 0x8a, 0xe2, 0x30,
	0x18, 0xc7, 0xe9, 0xc1, 0xae, 0xfd, 0x5c, 0x64, 0x09, 0xee, 0x12, 0x64, 0x59, 0xa4, 0x57, 0xee,
	0x8d, 0x0b, 0xbb, 0x4f, 0xb0, 0x4c, 0x2f, 0xa6, 0x83, 0x4a, 0xa9, 0xe2, 0x7d, 0x5a, 0xe3, 0x58,
	0x8c, 0x4d, 0x49, 0xe2, 0x80, 0x6f, 0x30, 0x8f, 0x35, 0x8f, 0x36, 0xe4, 0x60, 0xdb, 0xbb, 0xfc,
	0x0f, 0x4d, 0x7f, 0x5f, 0x3e, 0x98, 0x1d, 0x5b, 0xc1, 0x15, 0xff, 0x23, 0xa8, 0xbc, 0x31, 0x25,
	0x57, 0x46, 0xa1, 0xc8, 0xba, 0x73, 0xe4, 0x52, 0x75, 0x6f, 0xa9, 0xcb, 0x92, 0x77, 0x0f, 0xc6,
	0x39, 0x23, 0xea, 0xc4, 0xc5, 0x15, 0xcd, 0x60, 0xb4, 0xe1, 0x47, 0xca, 0xb0, 0xb7, 0xf0, 0x96,
	0x71, 0x61, 0x05, 0x9a, 0xc3, 0xb8, 0xa0, 0x6f, 0xb5, 0xac, 0x79, 0x83, 0x7d, 0x13, 0x74, 0x1a,
	0x61, 0xf8, 0x72, 0xa0, 0xc2, 0x44, 0x81, 0x89, 0x1e, 0x12, 0xfd, 0x80, 0x68, 0x47, 0x45, 0x4d,
	0x18, 0x0e, 0x4d, 0xe0, 0x14, 0x42, 0x10, 0x6e, 0x48, 0x25, 0xf1, 0x68, 0x11, 0x2c, 0xe3, 0xc2,
	0x9c, 0x5f, 0xc2, 0x71, 0xf4, 0x6d, 0x92, 0x7c, 0x78, 0x10, 0x67, 0x8d, 0xa2, 0xe2, 0x44, 0x2a,
	0x8a, 0x7e, 0x43, 0xb8, 0xbf, 0xb7, 0xd4, 0xa0, 0x4c, 0xff, 0x7e, 0x5f, 0x59, 0xf6, 0x55, 0x57,
	0xd0, 0x61, 0x61, 0x2a, 0xfa, 0xca, 0x2d, 0xb9, 0x52, 0x07, 0x67, 0xce, 0xe8, 0x27, 0xc4, 0xbb,
	0x33, 0x17, 0xaa, 0xd1, 0x81, 0x45, 0xeb, 0x0d, 0xb4, 0x80, 0x49, 0x4a, 0x65, 0x25, 0xea, 0x56,
	0x69, 0x74, 0x4b, 0x38, 0xb4, 0x34, 0xfe, 0x9a, 0x1d, 0xdb, 0x2c, 0xc5, 0x23, 0x8b, 0x6f, 0x95,
	0xbe, 0x37, 0xe7, 0x1b, 0x7a, 0x2d, 0xa9, 0x90, 0x38, 0x32, 0x33, 0xf4, 0x46, 0x52, 0xc2, 0x57,
	0xdd, 0xdb, 0xd2, 0xfa, 0xf5, 0x5c, 0x72, 0xa1, 0xdb, 0x6b, 0x5e, 0x11, 0x96, 0x73, 0xa1, 0xdc,
	0xa3, 0xf6, 0x86, 0x4e, 0x9f, 0xce, 0x44, 0xca, 0x5a, 0x66, 0xa9, 0x83, 0xef, 0x0d, 0x4d, 0xa0,
	0x5b, 0x59, 0xea, 0xf0, 0x9d, 0x4a, 0x18, 0x84, 0x07, 0x46, 0x9a, 0x6e, 0x6a, 0x6f, 0x30, 0xf5,
	0x14, 0x7c, 0x77, 0x55, 0x50, 0xf8, 0x59, 0x8a, 0x7e, 0x01, 0xec, 0xc5, 0xad, 0xb9, 0xe8, 0x4f,
	0x25, 0x0e, 0x0c, 0xee, 0xc0, 0xd1, 0xef, 0xf0, 0xbf, 0xaa, 0xa8, 0x94, 0xb6, 0x10, 0x9a, 0xc2,
	0xd0, 0x4a, 0x9e, 0x21, 0xcc, 0xda, 0xfa, 0xa4, 0x59, 0xbb, 0xa7, 0x7f, 0x4c, 0xd2, 0x2f, 0x4b,
	0xff, 0x37, 0x77, 0x23, 0xf8, 0x59, 0x6e, 0x97, 0x2c, 0x2f, 0x8e, 0xdc, 0x9c, 0xcb, 0xc8, 0x2c,
	0xf0, 0xdf, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0xb0, 0x49, 0x0e, 0xa4, 0x02, 0x00, 0x00,
}
