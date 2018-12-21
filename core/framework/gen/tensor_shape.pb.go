// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensor_shape.proto

package framework

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

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

// Dimensions of a tensor.
type TensorShapeProto struct {
	// Dimensions of the tensor, such as {"input", 30}, {"output", 40}
	// for a 30 x 40 2D tensor.  If an entry has size -1, this
	// corresponds to a dimension of unknown size. The names are
	// optional.
	//
	// The order of entries in "dim" matters: It indicates the layout of the
	// values in the tensor in-memory representation.
	//
	// The first entry in "dim" is the outermost dimension used to layout the
	// values, the last entry is the innermost dimension.  This matches the
	// in-memory layout of RowMajor Eigen tensors.
	//
	// If "dim.size()" > 0, "unknown_rank" must be false.
	Dim []*TensorShapeProto_Dim `protobuf:"bytes,2,rep,name=dim" json:"dim,omitempty"`
	// If true, the number of dimensions in the shape is unknown.
	//
	// If true, "dim.size()" must be 0.
	UnknownRank          bool     `protobuf:"varint,3,opt,name=unknown_rank,json=unknownRank,proto3" json:"unknown_rank,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TensorShapeProto) Reset()      { *m = TensorShapeProto{} }
func (*TensorShapeProto) ProtoMessage() {}
func (*TensorShapeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_tensor_shape_7d1bcec5c29cc248, []int{0}
}
func (m *TensorShapeProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TensorShapeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TensorShapeProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TensorShapeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TensorShapeProto.Merge(dst, src)
}
func (m *TensorShapeProto) XXX_Size() int {
	return m.Size()
}
func (m *TensorShapeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TensorShapeProto.DiscardUnknown(m)
}

var xxx_messageInfo_TensorShapeProto proto.InternalMessageInfo

func (m *TensorShapeProto) GetDim() []*TensorShapeProto_Dim {
	if m != nil {
		return m.Dim
	}
	return nil
}

func (m *TensorShapeProto) GetUnknownRank() bool {
	if m != nil {
		return m.UnknownRank
	}
	return false
}

// One dimension of the tensor.
type TensorShapeProto_Dim struct {
	// Size of the tensor in that dimension.
	// This value must be >= -1, but values of -1 are reserved for "unknown"
	// shapes (values of -1 mean "unknown" dimension).  Certain wrappers
	// that work with TensorShapeProto may fail at runtime when deserializing
	// a TensorShapeProto containing a dim value of -1.
	Size_ int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// Optional name of the tensor dimension.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TensorShapeProto_Dim) Reset()      { *m = TensorShapeProto_Dim{} }
func (*TensorShapeProto_Dim) ProtoMessage() {}
func (*TensorShapeProto_Dim) Descriptor() ([]byte, []int) {
	return fileDescriptor_tensor_shape_7d1bcec5c29cc248, []int{0, 0}
}
func (m *TensorShapeProto_Dim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TensorShapeProto_Dim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TensorShapeProto_Dim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TensorShapeProto_Dim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TensorShapeProto_Dim.Merge(dst, src)
}
func (m *TensorShapeProto_Dim) XXX_Size() int {
	return m.Size()
}
func (m *TensorShapeProto_Dim) XXX_DiscardUnknown() {
	xxx_messageInfo_TensorShapeProto_Dim.DiscardUnknown(m)
}

var xxx_messageInfo_TensorShapeProto_Dim proto.InternalMessageInfo

func (m *TensorShapeProto_Dim) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *TensorShapeProto_Dim) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*TensorShapeProto)(nil), "framework.TensorShapeProto")
	proto.RegisterType((*TensorShapeProto_Dim)(nil), "framework.TensorShapeProto.Dim")
}
func (this *TensorShapeProto) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*TensorShapeProto)
	if !ok {
		that2, ok := that.(TensorShapeProto)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *TensorShapeProto")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *TensorShapeProto but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *TensorShapeProto but is not nil && this == nil")
	}
	if len(this.Dim) != len(that1.Dim) {
		return fmt.Errorf("Dim this(%v) Not Equal that(%v)", len(this.Dim), len(that1.Dim))
	}
	for i := range this.Dim {
		if !this.Dim[i].Equal(that1.Dim[i]) {
			return fmt.Errorf("Dim this[%v](%v) Not Equal that[%v](%v)", i, this.Dim[i], i, that1.Dim[i])
		}
	}
	if this.UnknownRank != that1.UnknownRank {
		return fmt.Errorf("UnknownRank this(%v) Not Equal that(%v)", this.UnknownRank, that1.UnknownRank)
	}
	return nil
}
func (this *TensorShapeProto) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TensorShapeProto)
	if !ok {
		that2, ok := that.(TensorShapeProto)
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
	if len(this.Dim) != len(that1.Dim) {
		return false
	}
	for i := range this.Dim {
		if !this.Dim[i].Equal(that1.Dim[i]) {
			return false
		}
	}
	if this.UnknownRank != that1.UnknownRank {
		return false
	}
	return true
}
func (this *TensorShapeProto_Dim) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*TensorShapeProto_Dim)
	if !ok {
		that2, ok := that.(TensorShapeProto_Dim)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *TensorShapeProto_Dim")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *TensorShapeProto_Dim but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *TensorShapeProto_Dim but is not nil && this == nil")
	}
	if this.Size_ != that1.Size_ {
		return fmt.Errorf("Size_ this(%v) Not Equal that(%v)", this.Size_, that1.Size_)
	}
	if this.Name != that1.Name {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	return nil
}
func (this *TensorShapeProto_Dim) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TensorShapeProto_Dim)
	if !ok {
		that2, ok := that.(TensorShapeProto_Dim)
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
	if this.Size_ != that1.Size_ {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *TensorShapeProto) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&framework.TensorShapeProto{")
	if this.Dim != nil {
		s = append(s, "Dim: "+fmt.Sprintf("%#v", this.Dim)+",\n")
	}
	s = append(s, "UnknownRank: "+fmt.Sprintf("%#v", this.UnknownRank)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TensorShapeProto_Dim) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&framework.TensorShapeProto_Dim{")
	s = append(s, "Size_: "+fmt.Sprintf("%#v", this.Size_)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTensorShape(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *TensorShapeProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TensorShapeProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Dim) > 0 {
		for _, msg := range m.Dim {
			dAtA[i] = 0x12
			i++
			i = encodeVarintTensorShape(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.UnknownRank {
		dAtA[i] = 0x18
		i++
		if m.UnknownRank {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *TensorShapeProto_Dim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TensorShapeProto_Dim) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Size_ != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTensorShape(dAtA, i, uint64(m.Size_))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTensorShape(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func encodeVarintTensorShape(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedTensorShapeProto(r randyTensorShape, easy bool) *TensorShapeProto {
	this := &TensorShapeProto{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.Dim = make([]*TensorShapeProto_Dim, v1)
		for i := 0; i < v1; i++ {
			this.Dim[i] = NewPopulatedTensorShapeProto_Dim(r, easy)
		}
	}
	this.UnknownRank = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedTensorShapeProto_Dim(r randyTensorShape, easy bool) *TensorShapeProto_Dim {
	this := &TensorShapeProto_Dim{}
	this.Size_ = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Size_ *= -1
	}
	this.Name = string(randStringTensorShape(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTensorShape interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTensorShape(r randyTensorShape) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTensorShape(r randyTensorShape) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneTensorShape(r)
	}
	return string(tmps)
}
func randUnrecognizedTensorShape(r randyTensorShape, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTensorShape(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTensorShape(dAtA []byte, r randyTensorShape, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTensorShape(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateTensorShape(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateTensorShape(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTensorShape(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTensorShape(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTensorShape(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTensorShape(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *TensorShapeProto) Size() (n int) {
	var l int
	_ = l
	if len(m.Dim) > 0 {
		for _, e := range m.Dim {
			l = e.Size()
			n += 1 + l + sovTensorShape(uint64(l))
		}
	}
	if m.UnknownRank {
		n += 2
	}
	return n
}

func (m *TensorShapeProto_Dim) Size() (n int) {
	var l int
	_ = l
	if m.Size_ != 0 {
		n += 1 + sovTensorShape(uint64(m.Size_))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTensorShape(uint64(l))
	}
	return n
}

func sovTensorShape(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTensorShape(x uint64) (n int) {
	return sovTensorShape(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TensorShapeProto) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TensorShapeProto{`,
		`Dim:` + strings.Replace(fmt.Sprintf("%v", this.Dim), "TensorShapeProto_Dim", "TensorShapeProto_Dim", 1) + `,`,
		`UnknownRank:` + fmt.Sprintf("%v", this.UnknownRank) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TensorShapeProto_Dim) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TensorShapeProto_Dim{`,
		`Size_:` + fmt.Sprintf("%v", this.Size_) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTensorShape(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TensorShapeProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTensorShape
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
			return fmt.Errorf("proto: TensorShapeProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TensorShapeProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorShape
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTensorShape
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dim = append(m.Dim, &TensorShapeProto_Dim{})
			if err := m.Dim[len(m.Dim)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnknownRank", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorShape
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UnknownRank = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTensorShape(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTensorShape
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
func (m *TensorShapeProto_Dim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTensorShape
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
			return fmt.Errorf("proto: Dim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Dim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorShape
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorShape
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTensorShape
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTensorShape(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTensorShape
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
func skipTensorShape(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTensorShape
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
					return 0, ErrIntOverflowTensorShape
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
					return 0, ErrIntOverflowTensorShape
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
				return 0, ErrInvalidLengthTensorShape
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTensorShape
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
				next, err := skipTensorShape(dAtA[start:])
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
	ErrInvalidLengthTensorShape = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTensorShape   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("tensor_shape.proto", fileDescriptor_tensor_shape_7d1bcec5c29cc248) }

var fileDescriptor_tensor_shape_7d1bcec5c29cc248 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x8a, 0x2f, 0xce, 0x48, 0x2c, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x4c, 0x2b, 0x4a, 0xcc, 0x4d, 0x2d, 0xcf, 0x2f, 0xca, 0x96, 0x52, 0x4a, 0xcf, 0x4f, 0xcf, 0xd7,
	0x07, 0x0b, 0x27, 0x95, 0xa6, 0xe9, 0x83, 0x78, 0x60, 0x0e, 0x98, 0x05, 0x51, 0xae, 0x34, 0x9d,
	0x91, 0x4b, 0x20, 0x04, 0x6c, 0x4a, 0x30, 0xc8, 0x90, 0x00, 0xb0, 0x19, 0x86, 0x5c, 0xcc, 0x29,
	0x99, 0xb9, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xf2, 0x7a, 0x70, 0x13, 0xf5, 0xd0, 0x55,
	0xea, 0xb9, 0x64, 0xe6, 0x06, 0x81, 0xd4, 0x0a, 0x29, 0x72, 0xf1, 0x94, 0xe6, 0x65, 0xe7, 0xe5,
	0x97, 0xe7, 0xc5, 0x17, 0x25, 0xe6, 0x65, 0x4b, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x04, 0x71, 0x43,
	0xc5, 0x82, 0x12, 0xf3, 0xb2, 0xa5, 0x74, 0xb9, 0x98, 0x5d, 0x32, 0x73, 0x85, 0x84, 0xb8, 0x58,
	0x8a, 0x33, 0xab, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x5e,
	0x62, 0x6e, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xed, 0xa4, 0x73, 0xe3, 0xa1,
	0x1c, 0xc3, 0x83, 0x87, 0x72, 0x8c, 0x1f, 0x1e, 0xca, 0x31, 0xfe, 0x78, 0x28, 0xc7, 0xd8, 0xf0,
	0x48, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x1d, 0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0xde,
	0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x56, 0x54, 0xcc, 0x5a, 0x13, 0x01, 0x00, 0x00,
}
