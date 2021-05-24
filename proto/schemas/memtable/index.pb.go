// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: index.proto

package proto_memtable

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IndexSchemas struct {
	P                    []*PositionSchemas `protobuf:"bytes,1,rep,name=p,proto3" json:"p,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IndexSchemas) Reset()         { *m = IndexSchemas{} }
func (m *IndexSchemas) String() string { return proto.CompactTextString(m) }
func (*IndexSchemas) ProtoMessage()    {}
func (*IndexSchemas) Descriptor() ([]byte, []int) {
	return fileDescriptor_f750e0f7889345b5, []int{0}
}
func (m *IndexSchemas) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexSchemas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexSchemas.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexSchemas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexSchemas.Merge(m, src)
}
func (m *IndexSchemas) XXX_Size() int {
	return m.Size()
}
func (m *IndexSchemas) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexSchemas.DiscardUnknown(m)
}

var xxx_messageInfo_IndexSchemas proto.InternalMessageInfo

func (m *IndexSchemas) GetP() []*PositionSchemas {
	if m != nil {
		return m.P
	}
	return nil
}

type PositionSchemas struct {
	H64                  uint64   `protobuf:"fixed64,1,opt,name=H64,proto3" json:"H64,omitempty"`
	Offset               uint32   `protobuf:"fixed32,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PositionSchemas) Reset()         { *m = PositionSchemas{} }
func (m *PositionSchemas) String() string { return proto.CompactTextString(m) }
func (*PositionSchemas) ProtoMessage()    {}
func (*PositionSchemas) Descriptor() ([]byte, []int) {
	return fileDescriptor_f750e0f7889345b5, []int{1}
}
func (m *PositionSchemas) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PositionSchemas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PositionSchemas.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PositionSchemas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositionSchemas.Merge(m, src)
}
func (m *PositionSchemas) XXX_Size() int {
	return m.Size()
}
func (m *PositionSchemas) XXX_DiscardUnknown() {
	xxx_messageInfo_PositionSchemas.DiscardUnknown(m)
}

var xxx_messageInfo_PositionSchemas proto.InternalMessageInfo

func (m *PositionSchemas) GetH64() uint64 {
	if m != nil {
		return m.H64
	}
	return 0
}

func (m *PositionSchemas) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func init() {
	proto.RegisterType((*IndexSchemas)(nil), "proto.memtable.IndexSchemas")
	proto.RegisterType((*PositionSchemas)(nil), "proto.memtable.PositionSchemas")
}

func init() { proto.RegisterFile("index.proto", fileDescriptor_f750e0f7889345b5) }

var fileDescriptor_f750e0f7889345b5 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcc, 0x4b, 0x49,
	0xad, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03, 0x53, 0x7a, 0xb9, 0xa9, 0xb9, 0x25,
	0x89, 0x49, 0x39, 0xa9, 0x4a, 0xb6, 0x5c, 0x3c, 0x9e, 0x20, 0xe9, 0xe0, 0xe4, 0x8c, 0xd4, 0xdc,
	0xc4, 0x62, 0x21, 0x5d, 0x2e, 0xc6, 0x02, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x79, 0x3d,
	0x54, 0xb5, 0x7a, 0x01, 0xf9, 0xc5, 0x99, 0x25, 0x99, 0xf9, 0x79, 0x50, 0xb5, 0x41, 0x8c, 0x05,
	0x4a, 0xd6, 0x5c, 0xfc, 0x68, 0xa2, 0x42, 0x02, 0x5c, 0xcc, 0x1e, 0x66, 0x26, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x6c, 0x41, 0x20, 0xa6, 0x90, 0x18, 0x17, 0x9b, 0x7f, 0x5a, 0x5a, 0x71, 0x6a, 0x89,
	0x04, 0x93, 0x02, 0xa3, 0x06, 0x7b, 0x10, 0x94, 0xe7, 0x24, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0xb6, 0xd1,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x45, 0x02, 0xd7, 0xd1, 0xb3, 0x00, 0x00, 0x00,
}

func (m *IndexSchemas) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexSchemas) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexSchemas) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.P) > 0 {
		for iNdEx := len(m.P) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.P[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIndex(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PositionSchemas) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PositionSchemas) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PositionSchemas) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Offset != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.Offset))
		i--
		dAtA[i] = 0x15
	}
	if m.H64 != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.H64))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func encodeVarintIndex(dAtA []byte, offset int, v uint64) int {
	offset -= sovIndex(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IndexSchemas) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.P) > 0 {
		for _, e := range m.P {
			l = e.Size()
			n += 1 + l + sovIndex(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PositionSchemas) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.H64 != 0 {
		n += 9
	}
	if m.Offset != 0 {
		n += 5
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIndex(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIndex(x uint64) (n int) {
	return sovIndex(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IndexSchemas) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndex
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
			return fmt.Errorf("proto: IndexSchemas: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexSchemas: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIndex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIndex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.P = append(m.P, &PositionSchemas{})
			if err := m.P[len(m.P)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIndex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIndex
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PositionSchemas) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndex
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
			return fmt.Errorf("proto: PositionSchemas: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PositionSchemas: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field H64", wireType)
			}
			m.H64 = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.H64 = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.Offset = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		default:
			iNdEx = preIndex
			skippy, err := skipIndex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIndex
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIndex(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIndex
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
					return 0, ErrIntOverflowIndex
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
					return 0, ErrIntOverflowIndex
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
				return 0, ErrInvalidLengthIndex
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIndex
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIndex
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIndex        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIndex          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIndex = fmt.Errorf("proto: unexpected end of group")
)
