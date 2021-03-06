// Code generated by protoc-gen-go. DO NOT EDIT.
// source: allocation_description.proto

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

type AllocationDescription struct {
	// Total number of bytes requested
	RequestedBytes int64 `protobuf:"varint,1,opt,name=requested_bytes,json=requestedBytes,proto3" json:"requested_bytes,omitempty"`
	// Total number of bytes allocated if known
	AllocatedBytes int64 `protobuf:"varint,2,opt,name=allocated_bytes,json=allocatedBytes,proto3" json:"allocated_bytes,omitempty"`
	// Name of the allocator used
	AllocatorName string `protobuf:"bytes,3,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
	// Identifier of the allocated buffer if known
	AllocationId int64 `protobuf:"varint,4,opt,name=allocation_id,json=allocationId,proto3" json:"allocation_id,omitempty"`
	// Set if this tensor only has one remaining reference
	HasSingleReference bool `protobuf:"varint,5,opt,name=has_single_reference,json=hasSingleReference,proto3" json:"has_single_reference,omitempty"`
	// Address of the allocation.
	Ptr                  uint64   `protobuf:"varint,6,opt,name=ptr,proto3" json:"ptr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocationDescription) Reset()         { *m = AllocationDescription{} }
func (m *AllocationDescription) String() string { return proto.CompactTextString(m) }
func (*AllocationDescription) ProtoMessage()    {}
func (*AllocationDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebde754afd94ec1, []int{0}
}

func (m *AllocationDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocationDescription.Unmarshal(m, b)
}
func (m *AllocationDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocationDescription.Marshal(b, m, deterministic)
}
func (m *AllocationDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationDescription.Merge(m, src)
}
func (m *AllocationDescription) XXX_Size() int {
	return xxx_messageInfo_AllocationDescription.Size(m)
}
func (m *AllocationDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationDescription.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationDescription proto.InternalMessageInfo

func (m *AllocationDescription) GetRequestedBytes() int64 {
	if m != nil {
		return m.RequestedBytes
	}
	return 0
}

func (m *AllocationDescription) GetAllocatedBytes() int64 {
	if m != nil {
		return m.AllocatedBytes
	}
	return 0
}

func (m *AllocationDescription) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

func (m *AllocationDescription) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *AllocationDescription) GetHasSingleReference() bool {
	if m != nil {
		return m.HasSingleReference
	}
	return false
}

func (m *AllocationDescription) GetPtr() uint64 {
	if m != nil {
		return m.Ptr
	}
	return 0
}

func init() {
	proto.RegisterType((*AllocationDescription)(nil), "tensorflow.AllocationDescription")
}

func init() { proto.RegisterFile("allocation_description.proto", fileDescriptor_eebde754afd94ec1) }

var fileDescriptor_eebde754afd94ec1 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x40, 0x89, 0x5d, 0x17, 0x0d, 0xae, 0x4a, 0x50, 0x08, 0xe8, 0xa1, 0x28, 0x62, 0x4f, 0x45,
	0xf0, 0xea, 0xc5, 0xe2, 0xc5, 0x8b, 0x2c, 0xf1, 0x03, 0x42, 0xb6, 0x9d, 0xee, 0x16, 0xdb, 0x4c,
	0x9d, 0x44, 0x16, 0xff, 0x5c, 0x6f, 0x92, 0xae, 0x4d, 0x3d, 0x78, 0x1b, 0xde, 0x3c, 0x92, 0xe1,
	0xf1, 0x4b, 0xd3, 0xb6, 0x58, 0x1a, 0xdf, 0xa0, 0xd5, 0x15, 0xb8, 0x92, 0x9a, 0x3e, 0xcc, 0x79,
	0x4f, 0xe8, 0x51, 0x70, 0x0f, 0xd6, 0x21, 0xd5, 0x2d, 0x6e, 0xaf, 0xbe, 0x19, 0x3f, 0x7f, 0x8c,
	0xf2, 0xd3, 0xe4, 0x8a, 0x5b, 0x7e, 0x42, 0xf0, 0xfe, 0x01, 0xce, 0x43, 0xa5, 0x57, 0x9f, 0x1e,
	0x9c, 0x64, 0x29, 0xcb, 0x12, 0x75, 0x1c, 0x71, 0x11, 0x68, 0x10, 0x7f, 0xbf, 0x8b, 0xe2, 0xde,
	0x4e, 0x8c, 0x78, 0x27, 0xde, 0xf0, 0x91, 0x20, 0x69, 0x6b, 0x3a, 0x90, 0x49, 0xca, 0xb2, 0x43,
	0xb5, 0x88, 0xf4, 0xc5, 0x74, 0x20, 0xae, 0xf9, 0xe2, 0xcf, 0xf9, 0x4d, 0x25, 0x67, 0xc3, 0x6b,
	0x47, 0x13, 0x7c, 0xae, 0xc4, 0x1d, 0x3f, 0xdb, 0x18, 0xa7, 0x5d, 0x63, 0xd7, 0x2d, 0x68, 0x82,
	0x1a, 0x08, 0x6c, 0x09, 0x72, 0x3f, 0x65, 0xd9, 0x81, 0x12, 0x1b, 0xe3, 0x5e, 0x87, 0x95, 0x1a,
	0x37, 0xe2, 0x94, 0x27, 0xbd, 0x27, 0x39, 0x4f, 0x59, 0x36, 0x53, 0x61, 0x2c, 0x1e, 0xb8, 0x44,
	0x5a, 0xe7, 0x53, 0x8d, 0xbc, 0x26, 0xd3, 0xc1, 0x16, 0xe9, 0xad, 0xb8, 0xf8, 0x37, 0xca, 0x32,
	0xf4, 0x73, 0x4b, 0xf6, 0xc5, 0xd8, 0x6a, 0x3e, 0xc4, 0xbc, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xa2, 0x9f, 0xdd, 0xe5, 0x6c, 0x01, 0x00, 0x00,
}
