// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reader_base.proto

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

// For serializing and restoring the state of ReaderBase, see
// reader_base.h for details.
type ReaderBaseState struct {
	WorkStarted          int64    `protobuf:"varint,1,opt,name=work_started,json=workStarted,proto3" json:"work_started,omitempty"`
	WorkFinished         int64    `protobuf:"varint,2,opt,name=work_finished,json=workFinished,proto3" json:"work_finished,omitempty"`
	NumRecordsProduced   int64    `protobuf:"varint,3,opt,name=num_records_produced,json=numRecordsProduced,proto3" json:"num_records_produced,omitempty"`
	CurrentWork          []byte   `protobuf:"bytes,4,opt,name=current_work,json=currentWork,proto3" json:"current_work,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReaderBaseState) Reset()         { *m = ReaderBaseState{} }
func (m *ReaderBaseState) String() string { return proto.CompactTextString(m) }
func (*ReaderBaseState) ProtoMessage()    {}
func (*ReaderBaseState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e786c556d0df4cd2, []int{0}
}

func (m *ReaderBaseState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReaderBaseState.Unmarshal(m, b)
}
func (m *ReaderBaseState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReaderBaseState.Marshal(b, m, deterministic)
}
func (m *ReaderBaseState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReaderBaseState.Merge(m, src)
}
func (m *ReaderBaseState) XXX_Size() int {
	return xxx_messageInfo_ReaderBaseState.Size(m)
}
func (m *ReaderBaseState) XXX_DiscardUnknown() {
	xxx_messageInfo_ReaderBaseState.DiscardUnknown(m)
}

var xxx_messageInfo_ReaderBaseState proto.InternalMessageInfo

func (m *ReaderBaseState) GetWorkStarted() int64 {
	if m != nil {
		return m.WorkStarted
	}
	return 0
}

func (m *ReaderBaseState) GetWorkFinished() int64 {
	if m != nil {
		return m.WorkFinished
	}
	return 0
}

func (m *ReaderBaseState) GetNumRecordsProduced() int64 {
	if m != nil {
		return m.NumRecordsProduced
	}
	return 0
}

func (m *ReaderBaseState) GetCurrentWork() []byte {
	if m != nil {
		return m.CurrentWork
	}
	return nil
}

func init() {
	proto.RegisterType((*ReaderBaseState)(nil), "tensorflow.ReaderBaseState")
}

func init() { proto.RegisterFile("reader_base.proto", fileDescriptor_e786c556d0df4cd2) }

var fileDescriptor_e786c556d0df4cd2 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xb1, 0x4e, 0x85, 0x30,
	0x14, 0x86, 0x53, 0xaf, 0x71, 0xe8, 0xbd, 0x46, 0x6d, 0x1c, 0x3a, 0xa2, 0x2e, 0x4c, 0x44, 0xe3,
	0x1b, 0x30, 0x38, 0x93, 0x32, 0x38, 0x36, 0x85, 0x1e, 0x94, 0x20, 0x3d, 0xe4, 0xb4, 0x0d, 0x6f,
	0xe5, 0xf3, 0x39, 0x9a, 0x16, 0x12, 0xd6, 0xef, 0xff, 0xda, 0x93, 0x8f, 0x3f, 0x10, 0x18, 0x0b,
	0xa4, 0x3b, 0xe3, 0xa1, 0x5a, 0x08, 0x03, 0x0a, 0x1e, 0xc0, 0x79, 0xa4, 0xe1, 0x07, 0xd7, 0xe7,
	0x5f, 0xc6, 0xef, 0x54, 0x36, 0x6a, 0xe3, 0xa1, 0x0d, 0x26, 0x80, 0x78, 0xe2, 0x97, 0x15, 0x69,
	0xd2, 0x3e, 0x18, 0x0a, 0x60, 0x25, 0x2b, 0x58, 0x79, 0x52, 0xe7, 0xc4, 0xda, 0x0d, 0x89, 0x17,
	0x7e, 0x9b, 0x95, 0x61, 0x74, 0xa3, 0xff, 0x06, 0x2b, 0xaf, 0xb2, 0x93, 0xdf, 0x7d, 0xec, 0x4c,
	0xbc, 0xf2, 0x47, 0x17, 0x67, 0x4d, 0xd0, 0x23, 0x59, 0xaf, 0x17, 0x42, 0x1b, 0x7b, 0xb0, 0xf2,
	0x94, 0x5d, 0xe1, 0xe2, 0xac, 0xb6, 0xa9, 0xd9, 0x97, 0x74, 0xb9, 0x8f, 0x44, 0xe0, 0x82, 0x4e,
	0x3f, 0xc9, 0xeb, 0x82, 0x95, 0x17, 0x75, 0xde, 0xd9, 0x27, 0xd2, 0x54, 0xbf, 0x71, 0x89, 0xf4,
	0x55, 0x1d, 0x09, 0xd5, 0x40, 0x66, 0x86, 0xa4, 0xd7, 0xf7, 0x47, 0x49, 0x93, 0x4a, 0x7d, 0xc3,
	0xfe, 0x18, 0xeb, 0x6e, 0x72, 0xf6, 0xfb, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x71, 0xb8,
	0xe1, 0x0b, 0x01, 0x00, 0x00,
}