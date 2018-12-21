// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reader_base.proto

package framework

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// For serializing and restoring the state of ReaderBase, see
// reader_base.h for details.
type ReaderBaseState struct {
	WorkStarted          int64    `protobuf:"varint,1,opt,name=work_started,json=workStarted,proto3" json:"work_started,omitempty"`
	WorkFinished         int64    `protobuf:"varint,2,opt,name=work_finished,json=workFinished,proto3" json:"work_finished,omitempty"`
	NumRecordsProduced   int64    `protobuf:"varint,3,opt,name=num_records_produced,json=numRecordsProduced,proto3" json:"num_records_produced,omitempty"`
	CurrentWork          []byte   `protobuf:"bytes,4,opt,name=current_work,json=currentWork,proto3" json:"current_work,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReaderBaseState) Reset()      { *m = ReaderBaseState{} }
func (*ReaderBaseState) ProtoMessage() {}
func (*ReaderBaseState) Descriptor() ([]byte, []int) {
	return fileDescriptor_reader_base_55051ebfe3eeca26, []int{0}
}
func (m *ReaderBaseState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReaderBaseState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReaderBaseState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ReaderBaseState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReaderBaseState.Merge(dst, src)
}
func (m *ReaderBaseState) XXX_Size() int {
	return m.Size()
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
	proto.RegisterType((*ReaderBaseState)(nil), "framework.ReaderBaseState")
}
func (this *ReaderBaseState) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ReaderBaseState)
	if !ok {
		that2, ok := that.(ReaderBaseState)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ReaderBaseState")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ReaderBaseState but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ReaderBaseState but is not nil && this == nil")
	}
	if this.WorkStarted != that1.WorkStarted {
		return fmt.Errorf("WorkStarted this(%v) Not Equal that(%v)", this.WorkStarted, that1.WorkStarted)
	}
	if this.WorkFinished != that1.WorkFinished {
		return fmt.Errorf("WorkFinished this(%v) Not Equal that(%v)", this.WorkFinished, that1.WorkFinished)
	}
	if this.NumRecordsProduced != that1.NumRecordsProduced {
		return fmt.Errorf("NumRecordsProduced this(%v) Not Equal that(%v)", this.NumRecordsProduced, that1.NumRecordsProduced)
	}
	if !bytes.Equal(this.CurrentWork, that1.CurrentWork) {
		return fmt.Errorf("CurrentWork this(%v) Not Equal that(%v)", this.CurrentWork, that1.CurrentWork)
	}
	return nil
}
func (this *ReaderBaseState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReaderBaseState)
	if !ok {
		that2, ok := that.(ReaderBaseState)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.WorkStarted != that1.WorkStarted {
		return false
	}
	if this.WorkFinished != that1.WorkFinished {
		return false
	}
	if this.NumRecordsProduced != that1.NumRecordsProduced {
		return false
	}
	if !bytes.Equal(this.CurrentWork, that1.CurrentWork) {
		return false
	}
	return true
}
func (this *ReaderBaseState) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&framework.ReaderBaseState{")
	s = append(s, "WorkStarted: "+fmt.Sprintf("%#v", this.WorkStarted)+",\n")
	s = append(s, "WorkFinished: "+fmt.Sprintf("%#v", this.WorkFinished)+",\n")
	s = append(s, "NumRecordsProduced: "+fmt.Sprintf("%#v", this.NumRecordsProduced)+",\n")
	s = append(s, "CurrentWork: "+fmt.Sprintf("%#v", this.CurrentWork)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringReaderBase(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ReaderBaseState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReaderBaseState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.WorkStarted != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintReaderBase(dAtA, i, uint64(m.WorkStarted))
	}
	if m.WorkFinished != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintReaderBase(dAtA, i, uint64(m.WorkFinished))
	}
	if m.NumRecordsProduced != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintReaderBase(dAtA, i, uint64(m.NumRecordsProduced))
	}
	if len(m.CurrentWork) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintReaderBase(dAtA, i, uint64(len(m.CurrentWork)))
		i += copy(dAtA[i:], m.CurrentWork)
	}
	return i, nil
}

func encodeVarintReaderBase(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedReaderBaseState(r randyReaderBase, easy bool) *ReaderBaseState {
	this := &ReaderBaseState{}
	this.WorkStarted = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.WorkStarted *= -1
	}
	this.WorkFinished = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.WorkFinished *= -1
	}
	this.NumRecordsProduced = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.NumRecordsProduced *= -1
	}
	v1 := r.Intn(100)
	this.CurrentWork = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.CurrentWork[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyReaderBase interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneReaderBase(r randyReaderBase) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringReaderBase(r randyReaderBase) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneReaderBase(r)
	}
	return string(tmps)
}
func randUnrecognizedReaderBase(r randyReaderBase, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldReaderBase(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldReaderBase(dAtA []byte, r randyReaderBase, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateReaderBase(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateReaderBase(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateReaderBase(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateReaderBase(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateReaderBase(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateReaderBase(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateReaderBase(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ReaderBaseState) Size() (n int) {
	var l int
	_ = l
	if m.WorkStarted != 0 {
		n += 1 + sovReaderBase(uint64(m.WorkStarted))
	}
	if m.WorkFinished != 0 {
		n += 1 + sovReaderBase(uint64(m.WorkFinished))
	}
	if m.NumRecordsProduced != 0 {
		n += 1 + sovReaderBase(uint64(m.NumRecordsProduced))
	}
	l = len(m.CurrentWork)
	if l > 0 {
		n += 1 + l + sovReaderBase(uint64(l))
	}
	return n
}

func sovReaderBase(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozReaderBase(x uint64) (n int) {
	return sovReaderBase(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ReaderBaseState) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReaderBaseState{`,
		`WorkStarted:` + fmt.Sprintf("%v", this.WorkStarted) + `,`,
		`WorkFinished:` + fmt.Sprintf("%v", this.WorkFinished) + `,`,
		`NumRecordsProduced:` + fmt.Sprintf("%v", this.NumRecordsProduced) + `,`,
		`CurrentWork:` + fmt.Sprintf("%v", this.CurrentWork) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringReaderBase(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ReaderBaseState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReaderBase
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReaderBaseState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReaderBaseState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkStarted", wireType)
			}
			m.WorkStarted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReaderBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WorkStarted |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkFinished", wireType)
			}
			m.WorkFinished = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReaderBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WorkFinished |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumRecordsProduced", wireType)
			}
			m.NumRecordsProduced = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReaderBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumRecordsProduced |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentWork", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReaderBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthReaderBase
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentWork = append(m.CurrentWork[:0], dAtA[iNdEx:postIndex]...)
			if m.CurrentWork == nil {
				m.CurrentWork = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReaderBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReaderBase
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipReaderBase(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReaderBase
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowReaderBase
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowReaderBase
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthReaderBase
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowReaderBase
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipReaderBase(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthReaderBase = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReaderBase   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("reader_base.proto", fileDescriptor_reader_base_55051ebfe3eeca26) }

var fileDescriptor_reader_base_55051ebfe3eeca26 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x7b, 0x14, 0x21, 0x91, 0x16, 0x21, 0x2c, 0x86, 0x88, 0xe1, 0x54, 0xca, 0xd2, 0x01,
	0x51, 0x24, 0xde, 0xa0, 0x03, 0x33, 0x4a, 0x07, 0x46, 0xcb, 0x8d, 0x2f, 0xa5, 0xaa, 0x12, 0x57,
	0x67, 0x5b, 0xac, 0x3c, 0x02, 0x8f, 0xc1, 0xc4, 0xcc, 0x23, 0x30, 0x32, 0x32, 0x36, 0xe6, 0x05,
	0x18, 0x19, 0x91, 0x9d, 0x6c, 0xbe, 0xef, 0xbe, 0xff, 0xb7, 0x2e, 0x3b, 0x63, 0x52, 0x9a, 0x58,
	0xae, 0x94, 0xa5, 0x9b, 0x1d, 0x1b, 0x67, 0xc4, 0x71, 0xc5, 0xaa, 0xa6, 0x67, 0xc3, 0xdb, 0x8b,
	0xe9, 0xda, 0xac, 0xcd, 0x3c, 0xe1, 0x95, 0xaf, 0xe6, 0x71, 0x4a, 0x43, 0x7a, 0x75, 0xfa, 0xf4,
	0x1d, 0xb2, 0xd3, 0x22, 0x95, 0x2c, 0x94, 0xa5, 0xa5, 0x53, 0x8e, 0xc4, 0x65, 0x36, 0x8e, 0x79,
	0x69, 0x9d, 0x62, 0x47, 0x3a, 0x87, 0x09, 0xcc, 0x86, 0xc5, 0x28, 0xb2, 0x65, 0x87, 0xc4, 0x55,
	0x76, 0x92, 0x94, 0x6a, 0xd3, 0x6c, 0xec, 0x13, 0xe9, 0xfc, 0x20, 0x39, 0x29, 0x77, 0xdf, 0x33,
	0x71, 0x9b, 0x9d, 0x37, 0xbe, 0x96, 0x4c, 0xa5, 0x61, 0x6d, 0xe5, 0x8e, 0x8d, 0xf6, 0x25, 0xe9,
	0x7c, 0x98, 0x5c, 0xd1, 0xf8, 0xba, 0xe8, 0x56, 0x0f, 0xfd, 0x26, 0xfe, 0x5c, 0x7a, 0x66, 0x6a,
	0x9c, 0x8c, 0x4d, 0xf9, 0xe1, 0x04, 0x66, 0xe3, 0x62, 0xd4, 0xb3, 0x47, 0xc3, 0xdb, 0xc5, 0xf5,
	0x77, 0x8b, 0x83, 0x7d, 0x8b, 0xf0, 0xdb, 0x22, 0xfc, 0xb5, 0x08, 0x2f, 0x01, 0xe1, 0x2d, 0x20,
	0x7c, 0x04, 0x84, 0xcf, 0x80, 0xf0, 0x15, 0x10, 0xf6, 0x01, 0xe1, 0xf5, 0x07, 0x07, 0xab, 0xa3,
	0x74, 0xe5, 0xdd, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x8c, 0xe2, 0xa6, 0x29, 0x01, 0x00,
	0x00,
}
