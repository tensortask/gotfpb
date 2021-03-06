// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device_attributes.proto

package tensorflow

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

type InterconnectLink struct {
	DeviceId             int32    `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Strength             int32    `protobuf:"varint,3,opt,name=strength,proto3" json:"strength,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterconnectLink) Reset()         { *m = InterconnectLink{} }
func (m *InterconnectLink) String() string { return proto.CompactTextString(m) }
func (*InterconnectLink) ProtoMessage()    {}
func (*InterconnectLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d48eb250312832, []int{0}
}

func (m *InterconnectLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterconnectLink.Unmarshal(m, b)
}
func (m *InterconnectLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterconnectLink.Marshal(b, m, deterministic)
}
func (m *InterconnectLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterconnectLink.Merge(m, src)
}
func (m *InterconnectLink) XXX_Size() int {
	return xxx_messageInfo_InterconnectLink.Size(m)
}
func (m *InterconnectLink) XXX_DiscardUnknown() {
	xxx_messageInfo_InterconnectLink.DiscardUnknown(m)
}

var xxx_messageInfo_InterconnectLink proto.InternalMessageInfo

func (m *InterconnectLink) GetDeviceId() int32 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *InterconnectLink) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *InterconnectLink) GetStrength() int32 {
	if m != nil {
		return m.Strength
	}
	return 0
}

type LocalLinks struct {
	Link                 []*InterconnectLink `protobuf:"bytes,1,rep,name=link,proto3" json:"link,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LocalLinks) Reset()         { *m = LocalLinks{} }
func (m *LocalLinks) String() string { return proto.CompactTextString(m) }
func (*LocalLinks) ProtoMessage()    {}
func (*LocalLinks) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d48eb250312832, []int{1}
}

func (m *LocalLinks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalLinks.Unmarshal(m, b)
}
func (m *LocalLinks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalLinks.Marshal(b, m, deterministic)
}
func (m *LocalLinks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalLinks.Merge(m, src)
}
func (m *LocalLinks) XXX_Size() int {
	return xxx_messageInfo_LocalLinks.Size(m)
}
func (m *LocalLinks) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalLinks.DiscardUnknown(m)
}

var xxx_messageInfo_LocalLinks proto.InternalMessageInfo

func (m *LocalLinks) GetLink() []*InterconnectLink {
	if m != nil {
		return m.Link
	}
	return nil
}

type DeviceLocality struct {
	// Optional bus locality of device.  Default value of 0 means
	// no specific locality.  Specific localities are indexed from 1.
	BusId int32 `protobuf:"varint,1,opt,name=bus_id,json=busId,proto3" json:"bus_id,omitempty"`
	// Optional NUMA locality of device.
	NumaNode int32 `protobuf:"varint,2,opt,name=numa_node,json=numaNode,proto3" json:"numa_node,omitempty"`
	// Optional local interconnect links to other devices.
	Links                *LocalLinks `protobuf:"bytes,3,opt,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeviceLocality) Reset()         { *m = DeviceLocality{} }
func (m *DeviceLocality) String() string { return proto.CompactTextString(m) }
func (*DeviceLocality) ProtoMessage()    {}
func (*DeviceLocality) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d48eb250312832, []int{2}
}

func (m *DeviceLocality) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceLocality.Unmarshal(m, b)
}
func (m *DeviceLocality) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceLocality.Marshal(b, m, deterministic)
}
func (m *DeviceLocality) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceLocality.Merge(m, src)
}
func (m *DeviceLocality) XXX_Size() int {
	return xxx_messageInfo_DeviceLocality.Size(m)
}
func (m *DeviceLocality) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceLocality.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceLocality proto.InternalMessageInfo

func (m *DeviceLocality) GetBusId() int32 {
	if m != nil {
		return m.BusId
	}
	return 0
}

func (m *DeviceLocality) GetNumaNode() int32 {
	if m != nil {
		return m.NumaNode
	}
	return 0
}

func (m *DeviceLocality) GetLinks() *LocalLinks {
	if m != nil {
		return m.Links
	}
	return nil
}

type DeviceAttributes struct {
	// Fully specified name of the device within a cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// String representation of device_type.
	DeviceType string `protobuf:"bytes,2,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	// Memory capacity of device in bytes.
	MemoryLimit int64 `protobuf:"varint,4,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	// Platform-specific data about device that may be useful
	// for supporting efficient data transfers.
	Locality *DeviceLocality `protobuf:"bytes,5,opt,name=locality,proto3" json:"locality,omitempty"`
	// A device is assigned a global unique number each time it is
	// initialized. "incarnation" should never be 0.
	Incarnation uint64 `protobuf:"fixed64,6,opt,name=incarnation,proto3" json:"incarnation,omitempty"`
	// String representation of the physical device that this device maps to.
	PhysicalDeviceDesc   string   `protobuf:"bytes,7,opt,name=physical_device_desc,json=physicalDeviceDesc,proto3" json:"physical_device_desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceAttributes) Reset()         { *m = DeviceAttributes{} }
func (m *DeviceAttributes) String() string { return proto.CompactTextString(m) }
func (*DeviceAttributes) ProtoMessage()    {}
func (*DeviceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d48eb250312832, []int{3}
}

func (m *DeviceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceAttributes.Unmarshal(m, b)
}
func (m *DeviceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceAttributes.Marshal(b, m, deterministic)
}
func (m *DeviceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceAttributes.Merge(m, src)
}
func (m *DeviceAttributes) XXX_Size() int {
	return xxx_messageInfo_DeviceAttributes.Size(m)
}
func (m *DeviceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceAttributes proto.InternalMessageInfo

func (m *DeviceAttributes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceAttributes) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *DeviceAttributes) GetMemoryLimit() int64 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *DeviceAttributes) GetLocality() *DeviceLocality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *DeviceAttributes) GetIncarnation() uint64 {
	if m != nil {
		return m.Incarnation
	}
	return 0
}

func (m *DeviceAttributes) GetPhysicalDeviceDesc() string {
	if m != nil {
		return m.PhysicalDeviceDesc
	}
	return ""
}

func init() {
	proto.RegisterType((*InterconnectLink)(nil), "tensorflow.InterconnectLink")
	proto.RegisterType((*LocalLinks)(nil), "tensorflow.LocalLinks")
	proto.RegisterType((*DeviceLocality)(nil), "tensorflow.DeviceLocality")
	proto.RegisterType((*DeviceAttributes)(nil), "tensorflow.DeviceAttributes")
}

func init() { proto.RegisterFile("device_attributes.proto", fileDescriptor_04d48eb250312832) }

var fileDescriptor_04d48eb250312832 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xc1, 0x8b, 0xd4, 0x30,
	0x18, 0xc5, 0x89, 0x33, 0xad, 0x9d, 0xaf, 0x22, 0x4b, 0xd0, 0x35, 0xac, 0x82, 0xb5, 0xa7, 0x1e,
	0x64, 0x58, 0x56, 0xd0, 0x9b, 0xe0, 0xb2, 0x97, 0x81, 0x41, 0x96, 0xe0, 0xbd, 0xa4, 0x69, 0x76,
	0x37, 0x4c, 0x9b, 0x0c, 0x49, 0xea, 0xd2, 0x7f, 0x5c, 0x3c, 0x4a, 0x92, 0xce, 0x4c, 0x9d, 0x5b,
	0xfa, 0xfa, 0xf2, 0xbd, 0xdf, 0xf7, 0x08, 0xbc, 0x6b, 0xc5, 0x6f, 0xc9, 0x45, 0xcd, 0x9c, 0x33,
	0xb2, 0x19, 0x9c, 0xb0, 0xeb, 0xbd, 0xd1, 0x4e, 0x63, 0x70, 0x42, 0x59, 0x6d, 0x1e, 0x3a, 0xfd,
	0x5c, 0xd6, 0x70, 0xb1, 0x51, 0x4e, 0x18, 0xae, 0x95, 0x12, 0xdc, 0x6d, 0xa5, 0xda, 0xe1, 0xf7,
	0xb0, 0x9a, 0xae, 0xca, 0x96, 0xa0, 0x02, 0x55, 0x09, 0xcd, 0xa2, 0xb0, 0x69, 0x31, 0x86, 0xa5,
	0x1b, 0xf7, 0x82, 0xbc, 0x28, 0x50, 0xb5, 0xa2, 0xe1, 0x8c, 0xaf, 0x20, 0xb3, 0xce, 0x08, 0xf5,
	0xe8, 0x9e, 0xc8, 0x22, 0xfa, 0x0f, 0xdf, 0xe5, 0x77, 0x80, 0xad, 0xe6, 0xac, 0xf3, 0x93, 0x2d,
	0xbe, 0x86, 0x65, 0x27, 0xd5, 0x8e, 0xa0, 0x62, 0x51, 0xe5, 0x37, 0x1f, 0xd6, 0x27, 0x92, 0xf5,
	0x39, 0x06, 0x0d, 0xce, 0xd2, 0xc0, 0xeb, 0xbb, 0x90, 0x1d, 0xa6, 0x48, 0x37, 0xe2, 0xb7, 0x90,
	0x36, 0x83, 0x3d, 0xb1, 0x25, 0xcd, 0x60, 0x37, 0xad, 0xa7, 0x56, 0x43, 0xcf, 0x6a, 0xa5, 0xdb,
	0x48, 0x97, 0xd0, 0xcc, 0x0b, 0x3f, 0x75, 0x2b, 0xf0, 0x67, 0x48, 0xfc, 0x34, 0x1b, 0xf0, 0xf2,
	0x9b, 0xcb, 0x79, 0xf0, 0x09, 0x8f, 0x46, 0x53, 0xf9, 0x07, 0xc1, 0x45, 0x0c, 0xfd, 0x71, 0xec,
	0xce, 0x2f, 0xae, 0x58, 0x2f, 0x42, 0xe8, 0x8a, 0x86, 0x33, 0xfe, 0x08, 0xf9, 0xd4, 0xd4, 0xac,
	0x13, 0x88, 0xd2, 0x2f, 0xdf, 0xcc, 0x27, 0x78, 0xd5, 0x8b, 0x5e, 0x9b, 0xb1, 0xee, 0x64, 0x2f,
	0x1d, 0x59, 0x16, 0xa8, 0x5a, 0xd0, 0x3c, 0x6a, 0x5b, 0x2f, 0xe1, 0xaf, 0x90, 0x75, 0xd3, 0x6a,
	0x24, 0x09, 0x74, 0x57, 0x73, 0xba, 0xff, 0x97, 0xa7, 0x47, 0x2f, 0x2e, 0x20, 0x97, 0x8a, 0x33,
	0xa3, 0x98, 0x93, 0x5a, 0x91, 0xb4, 0x40, 0x55, 0x4a, 0xe7, 0x12, 0xbe, 0x86, 0x37, 0xfb, 0xa7,
	0xd1, 0x4a, 0xce, 0xba, 0x7a, 0xc2, 0x6c, 0x85, 0xe5, 0xe4, 0x65, 0xc0, 0xc4, 0x87, 0x7f, 0x31,
	0xe1, 0x4e, 0x58, 0x7e, 0xfb, 0x0d, 0x88, 0x36, 0x8f, 0xf3, 0xf8, 0x07, 0xc3, 0x7a, 0xf1, 0xac,
	0xcd, 0xee, 0xf6, 0xf2, 0xbc, 0x91, 0x7b, 0xff, 0x98, 0xec, 0x3d, 0xfa, 0x8b, 0x50, 0x93, 0x86,
	0x97, 0xf5, 0xe5, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0xa2, 0xbd, 0x47, 0x74, 0x02, 0x00,
	0x00,
}
