// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource_handle.proto

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

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type ResourceHandleProto struct {
	// Unique name for the device containing the resource.
	Device string `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	// Container in which this resource is placed.
	Container string `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
	// Unique name of this resource.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Hash code for the type of the resource. Is only valid in the same device
	// and in the same execution.
	HashCode uint64 `protobuf:"varint,4,opt,name=hash_code,json=hashCode,proto3" json:"hash_code,omitempty"`
	// For debug-only, the name of the type pointed to by this handle, if
	// available.
	MaybeTypeName        string   `protobuf:"bytes,5,opt,name=maybe_type_name,json=maybeTypeName,proto3" json:"maybe_type_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceHandleProto) Reset()         { *m = ResourceHandleProto{} }
func (m *ResourceHandleProto) String() string { return proto.CompactTextString(m) }
func (*ResourceHandleProto) ProtoMessage()    {}
func (*ResourceHandleProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ca6d7983246ca36, []int{0}
}

func (m *ResourceHandleProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceHandleProto.Unmarshal(m, b)
}
func (m *ResourceHandleProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceHandleProto.Marshal(b, m, deterministic)
}
func (m *ResourceHandleProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceHandleProto.Merge(m, src)
}
func (m *ResourceHandleProto) XXX_Size() int {
	return xxx_messageInfo_ResourceHandleProto.Size(m)
}
func (m *ResourceHandleProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceHandleProto.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceHandleProto proto.InternalMessageInfo

func (m *ResourceHandleProto) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *ResourceHandleProto) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *ResourceHandleProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceHandleProto) GetHashCode() uint64 {
	if m != nil {
		return m.HashCode
	}
	return 0
}

func (m *ResourceHandleProto) GetMaybeTypeName() string {
	if m != nil {
		return m.MaybeTypeName
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceHandleProto)(nil), "tensorflow.ResourceHandleProto")
}

func init() { proto.RegisterFile("resource_handle.proto", fileDescriptor_2ca6d7983246ca36) }

var fileDescriptor_2ca6d7983246ca36 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0xd6, 0xc5, 0x0e, 0xa8, 0x10, 0x51, 0x02, 0x7a, 0x58, 0x3c, 0xc8, 0x9e, 0xea,
	0xc1, 0x37, 0x58, 0x2f, 0x9e, 0x44, 0x8a, 0xf7, 0x90, 0x4d, 0xa6, 0xb6, 0xd8, 0x64, 0xca, 0x34,
	0x5a, 0xfa, 0x3e, 0x3e, 0xa4, 0x47, 0x71, 0x28, 0x14, 0xbc, 0x25, 0xdf, 0xff, 0xfd, 0xf0, 0x0f,
	0x5c, 0x31, 0x8e, 0xf4, 0xc9, 0x1e, 0x6d, 0xeb, 0x52, 0xe8, 0xb1, 0x1a, 0x98, 0x32, 0x69, 0xc8,
	0x98, 0x46, 0xe2, 0xa6, 0xa7, 0xe9, 0xee, 0x5b, 0xc1, 0x65, 0xbd, 0x58, 0xcf, 0x22, 0xbd, 0x8a,
	0x73, 0x0d, 0x9b, 0x80, 0x5f, 0x9d, 0x47, 0xa3, 0xb6, 0x6a, 0x57, 0xd6, 0xcb, 0x4f, 0xdf, 0x42,
	0xe9, 0x29, 0x65, 0xd7, 0x25, 0x64, 0x73, 0x24, 0xd1, 0x0a, 0xb4, 0x86, 0x22, 0xb9, 0x88, 0xe6,
	0x58, 0x02, 0x79, 0xeb, 0x1b, 0x28, 0x5b, 0x37, 0xb6, 0xd6, 0x53, 0x40, 0x53, 0x6c, 0xd5, 0xae,
	0xa8, 0x4f, 0xff, 0xc0, 0x13, 0x05, 0xd4, 0xf7, 0x70, 0x11, 0xdd, 0x7c, 0x40, 0x9b, 0xe7, 0x01,
	0xad, 0x74, 0x4f, 0xa4, 0x7b, 0x26, 0xf8, 0x6d, 0x1e, 0xf0, 0xc5, 0x45, 0xdc, 0x3f, 0x80, 0x21,
	0x7e, 0xaf, 0xd6, 0xe1, 0x55, 0xc3, 0x2e, 0xe2, 0x44, 0xfc, 0xb1, 0x3f, 0xff, 0xb7, 0x5f, 0xfd,
	0x28, 0x75, 0xd8, 0xc8, 0xa9, 0x8f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x3a, 0xb0, 0x8d,
	0x03, 0x01, 0x00, 0x00,
}
