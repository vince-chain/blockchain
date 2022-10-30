// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vince/vrf/v1/randomval.proto

package types

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Randomval struct {
	Index      string  `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Creator    string  `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Vrv        string  `protobuf:"bytes,3,opt,name=vrv,proto3" json:"vrv,omitempty"`
	Multiplier uint64  `protobuf:"varint,4,opt,name=multiplier,proto3" json:"multiplier,omitempty"`
	Proof      string  `protobuf:"bytes,5,opt,name=proof,proto3" json:"proof,omitempty"`
	Pubk       string  `protobuf:"bytes,6,opt,name=pubk,proto3" json:"pubk,omitempty"`
	Message    string  `protobuf:"bytes,7,opt,name=message,proto3" json:"message,omitempty"`
	Parsedvrv  uint64  `protobuf:"varint,8,opt,name=parsedvrv,proto3" json:"parsedvrv,omitempty"`
	Floatvrv   float64 `protobuf:"fixed64,9,opt,name=floatvrv,proto3" json:"floatvrv,omitempty"`
	// between or equal to 0 and 1
	Finalvrv   uint64  `protobuf:"varint,10,opt,name=finalvrv,proto3" json:"finalvrv,omitempty"`
	Finalvrvfl float64 `protobuf:"fixed64,11,opt,name=finalvrvfl,proto3" json:"finalvrvfl,omitempty"`
}

func (m *Randomval) Reset()         { *m = Randomval{} }
func (m *Randomval) String() string { return proto.CompactTextString(m) }
func (*Randomval) ProtoMessage()    {}
func (*Randomval) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f0232beec08fea7, []int{0}
}
func (m *Randomval) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Randomval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Randomval.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Randomval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Randomval.Merge(m, src)
}
func (m *Randomval) XXX_Size() int {
	return m.Size()
}
func (m *Randomval) XXX_DiscardUnknown() {
	xxx_messageInfo_Randomval.DiscardUnknown(m)
}

var xxx_messageInfo_Randomval proto.InternalMessageInfo

func (m *Randomval) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Randomval) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Randomval) GetVrv() string {
	if m != nil {
		return m.Vrv
	}
	return ""
}

func (m *Randomval) GetMultiplier() uint64 {
	if m != nil {
		return m.Multiplier
	}
	return 0
}

func (m *Randomval) GetProof() string {
	if m != nil {
		return m.Proof
	}
	return ""
}

func (m *Randomval) GetPubk() string {
	if m != nil {
		return m.Pubk
	}
	return ""
}

func (m *Randomval) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Randomval) GetParsedvrv() uint64 {
	if m != nil {
		return m.Parsedvrv
	}
	return 0
}

func (m *Randomval) GetFloatvrv() float64 {
	if m != nil {
		return m.Floatvrv
	}
	return 0
}

func (m *Randomval) GetFinalvrv() uint64 {
	if m != nil {
		return m.Finalvrv
	}
	return 0
}

func (m *Randomval) GetFinalvrvfl() float64 {
	if m != nil {
		return m.Finalvrvfl
	}
	return 0
}

func init() {
	proto.RegisterType((*Randomval)(nil), "vince.vrf.v1.Randomval")
}

func init() { proto.RegisterFile("vince/vrf/v1/randomval.proto", fileDescriptor_6f0232beec08fea7) }

var fileDescriptor_6f0232beec08fea7 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xbd, 0x6e, 0xf3, 0x20,
	0x14, 0x86, 0x43, 0xfe, 0xcd, 0x27, 0x7d, 0xaa, 0x50, 0x07, 0x54, 0x55, 0x28, 0xea, 0x94, 0xc9,
	0x28, 0xca, 0x1d, 0x74, 0xef, 0xe2, 0xb1, 0x1b, 0x89, 0x21, 0x41, 0xc5, 0x80, 0x30, 0x46, 0xe9,
	0x5d, 0xf4, 0x2e, 0x7a, 0x2b, 0x1d, 0x33, 0x76, 0xac, 0x92, 0x1b, 0xa9, 0xc0, 0x3f, 0xcd, 0x76,
	0xde, 0xe7, 0xf5, 0xd1, 0x63, 0x71, 0x20, 0xe1, 0xfb, 0x23, 0x57, 0x46, 0xd3, 0xe0, 0x04, 0x0d,
	0x1b, 0xea, 0x98, 0x2e, 0x4d, 0x15, 0x98, 0xca, 0xad, 0x33, 0xde, 0xa0, 0xff, 0x5d, 0x9f, 0x07,
	0x27, 0xf2, 0xb0, 0x79, 0xfa, 0x1c, 0xc3, 0xac, 0xe8, 0xbf, 0x41, 0xf7, 0x70, 0x26, 0x75, 0xc9,
	0x4f, 0x18, 0xac, 0xc0, 0x3a, 0x2b, 0xda, 0x80, 0x30, 0x5c, 0xec, 0x1d, 0x67, 0xde, 0x38, 0x3c,
	0x4e, 0xbc, 0x8f, 0xe8, 0x0e, 0x4e, 0x82, 0x0b, 0x78, 0x92, 0x68, 0x1c, 0x11, 0x81, 0xb0, 0x6a,
	0x94, 0x97, 0x56, 0x49, 0xee, 0xf0, 0x74, 0x05, 0xd6, 0xd3, 0xe2, 0x86, 0x44, 0x83, 0x75, 0xc6,
	0x08, 0x3c, 0x6b, 0x0d, 0x29, 0x20, 0x04, 0xa7, 0xb6, 0xd9, 0xbd, 0xe1, 0x79, 0x82, 0x69, 0x8e,
	0xd6, 0x8a, 0xd7, 0x35, 0x3b, 0x70, 0xbc, 0x68, 0xad, 0x5d, 0x44, 0x8f, 0x30, 0xb3, 0xcc, 0xd5,
	0xbc, 0x8c, 0xee, 0x65, 0x52, 0xfc, 0x01, 0xf4, 0x00, 0x97, 0x42, 0x19, 0xe6, 0x63, 0x99, 0xad,
	0xc0, 0x1a, 0x14, 0x43, 0x4e, 0x9d, 0xd4, 0x4c, 0xc5, 0x0e, 0xa6, 0xc5, 0x21, 0xc7, 0x3f, 0xef,
	0x67, 0xa1, 0xf0, 0xbf, 0xb4, 0x79, 0x43, 0x9e, 0x5f, 0xbe, 0x2e, 0x04, 0x9c, 0x2f, 0x04, 0xfc,
	0x5c, 0x08, 0xf8, 0xb8, 0x92, 0xd1, 0xf9, 0x4a, 0x46, 0xdf, 0x57, 0x32, 0x7a, 0xdd, 0x1e, 0xa4,
	0x3f, 0x36, 0xbb, 0x7c, 0x6f, 0x2a, 0xda, 0x3d, 0xaf, 0x30, 0x8d, 0x2e, 0x99, 0x97, 0x46, 0xd3,
	0xe1, 0x20, 0x5b, 0x7a, 0x4a, 0x57, 0xf1, 0xef, 0x96, 0xd7, 0xbb, 0x79, 0xba, 0xc7, 0xf6, 0x37,
	0x00, 0x00, 0xff, 0xff, 0xf0, 0x53, 0x81, 0x3d, 0xb1, 0x01, 0x00, 0x00,
}

func (m *Randomval) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Randomval) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Randomval) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Finalvrvfl != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Finalvrvfl))))
		i--
		dAtA[i] = 0x59
	}
	if m.Finalvrv != 0 {
		i = encodeVarintRandomval(dAtA, i, uint64(m.Finalvrv))
		i--
		dAtA[i] = 0x50
	}
	if m.Floatvrv != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Floatvrv))))
		i--
		dAtA[i] = 0x49
	}
	if m.Parsedvrv != 0 {
		i = encodeVarintRandomval(dAtA, i, uint64(m.Parsedvrv))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintRandomval(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Pubk) > 0 {
		i -= len(m.Pubk)
		copy(dAtA[i:], m.Pubk)
		i = encodeVarintRandomval(dAtA, i, uint64(len(m.Pubk)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Proof) > 0 {
		i -= len(m.Proof)
		copy(dAtA[i:], m.Proof)
		i = encodeVarintRandomval(dAtA, i, uint64(len(m.Proof)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Multiplier != 0 {
		i = encodeVarintRandomval(dAtA, i, uint64(m.Multiplier))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Vrv) > 0 {
		i -= len(m.Vrv)
		copy(dAtA[i:], m.Vrv)
		i = encodeVarintRandomval(dAtA, i, uint64(len(m.Vrv)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRandomval(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintRandomval(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRandomval(dAtA []byte, offset int, v uint64) int {
	offset -= sovRandomval(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Randomval) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovRandomval(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRandomval(uint64(l))
	}
	l = len(m.Vrv)
	if l > 0 {
		n += 1 + l + sovRandomval(uint64(l))
	}
	if m.Multiplier != 0 {
		n += 1 + sovRandomval(uint64(m.Multiplier))
	}
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovRandomval(uint64(l))
	}
	l = len(m.Pubk)
	if l > 0 {
		n += 1 + l + sovRandomval(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovRandomval(uint64(l))
	}
	if m.Parsedvrv != 0 {
		n += 1 + sovRandomval(uint64(m.Parsedvrv))
	}
	if m.Floatvrv != 0 {
		n += 9
	}
	if m.Finalvrv != 0 {
		n += 1 + sovRandomval(uint64(m.Finalvrv))
	}
	if m.Finalvrvfl != 0 {
		n += 9
	}
	return n
}

func sovRandomval(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRandomval(x uint64) (n int) {
	return sovRandomval(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Randomval) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRandomval
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Randomval: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Randomval: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandomval
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRandomval
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRandomval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandomval
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRandomval
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRandomval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vrv", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandomval
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRandomval
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRandomval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vrv = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
			}
			m.Multiplier = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandomval
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Multiplier |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandomval
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRandomval
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRandomval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubk", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandomval
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRandomval
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRandomval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubk = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandomval
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRandomval
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRandomval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parsedvrv", wireType)
			}
			m.Parsedvrv = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandomval
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Parsedvrv |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Floatvrv", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Floatvrv = float64(math.Float64frombits(v))
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Finalvrv", wireType)
			}
			m.Finalvrv = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandomval
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Finalvrv |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Finalvrvfl", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Finalvrvfl = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipRandomval(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRandomval
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
func skipRandomval(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRandomval
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
					return 0, ErrIntOverflowRandomval
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRandomval
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
			if length < 0 {
				return 0, ErrInvalidLengthRandomval
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRandomval
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRandomval
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRandomval        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRandomval          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRandomval = fmt.Errorf("proto: unexpected end of group")
)
