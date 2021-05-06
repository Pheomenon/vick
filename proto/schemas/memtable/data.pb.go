// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: data.proto

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

type DataSchemas struct {
	Es                   []*EntrySchemas `protobuf:"bytes,1,rep,name=es,proto3" json:"es,omitempty"`
	Info                 *InfoSchemas    `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DataSchemas) Reset()         { *m = DataSchemas{} }
func (m *DataSchemas) String() string { return proto.CompactTextString(m) }
func (*DataSchemas) ProtoMessage()    {}
func (*DataSchemas) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}
func (m *DataSchemas) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataSchemas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataSchemas.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataSchemas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSchemas.Merge(m, src)
}
func (m *DataSchemas) XXX_Size() int {
	return m.Size()
}
func (m *DataSchemas) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSchemas.DiscardUnknown(m)
}

var xxx_messageInfo_DataSchemas proto.InternalMessageInfo

func (m *DataSchemas) GetEs() []*EntrySchemas {
	if m != nil {
		return m.Es
	}
	return nil
}

func (m *DataSchemas) GetInfo() *InfoSchemas {
	if m != nil {
		return m.Info
	}
	return nil
}

type EntrySchemas struct {
	KeyLen               uint32   `protobuf:"fixed32,1,opt,name=keyLen,proto3" json:"keyLen,omitempty"`
	ValLen               uint32   `protobuf:"fixed32,2,opt,name=valLen,proto3" json:"valLen,omitempty"`
	Key                  []byte   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Val                  []byte   `protobuf:"bytes,4,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntrySchemas) Reset()         { *m = EntrySchemas{} }
func (m *EntrySchemas) String() string { return proto.CompactTextString(m) }
func (*EntrySchemas) ProtoMessage()    {}
func (*EntrySchemas) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{1}
}
func (m *EntrySchemas) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntrySchemas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntrySchemas.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntrySchemas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntrySchemas.Merge(m, src)
}
func (m *EntrySchemas) XXX_Size() int {
	return m.Size()
}
func (m *EntrySchemas) XXX_DiscardUnknown() {
	xxx_messageInfo_EntrySchemas.DiscardUnknown(m)
}

var xxx_messageInfo_EntrySchemas proto.InternalMessageInfo

func (m *EntrySchemas) GetKeyLen() uint32 {
	if m != nil {
		return m.KeyLen
	}
	return 0
}

func (m *EntrySchemas) GetValLen() uint32 {
	if m != nil {
		return m.ValLen
	}
	return 0
}

func (m *EntrySchemas) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *EntrySchemas) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

type InfoSchemas struct {
	MetaOffset           uint32   `protobuf:"varint,1,opt,name=metaOffset,proto3" json:"metaOffset,omitempty"`
	Entries              int32    `protobuf:"varint,2,opt,name=entries,proto3" json:"entries,omitempty"`
	MinRange             uint32   `protobuf:"fixed32,3,opt,name=minRange,proto3" json:"minRange,omitempty"`
	MaxRange             uint32   `protobuf:"fixed32,4,opt,name=maxRange,proto3" json:"maxRange,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoSchemas) Reset()         { *m = InfoSchemas{} }
func (m *InfoSchemas) String() string { return proto.CompactTextString(m) }
func (*InfoSchemas) ProtoMessage()    {}
func (*InfoSchemas) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{2}
}
func (m *InfoSchemas) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InfoSchemas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InfoSchemas.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InfoSchemas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoSchemas.Merge(m, src)
}
func (m *InfoSchemas) XXX_Size() int {
	return m.Size()
}
func (m *InfoSchemas) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoSchemas.DiscardUnknown(m)
}

var xxx_messageInfo_InfoSchemas proto.InternalMessageInfo

func (m *InfoSchemas) GetMetaOffset() uint32 {
	if m != nil {
		return m.MetaOffset
	}
	return 0
}

func (m *InfoSchemas) GetEntries() int32 {
	if m != nil {
		return m.Entries
	}
	return 0
}

func (m *InfoSchemas) GetMinRange() uint32 {
	if m != nil {
		return m.MinRange
	}
	return 0
}

func (m *InfoSchemas) GetMaxRange() uint32 {
	if m != nil {
		return m.MaxRange
	}
	return 0
}

func init() {
	proto.RegisterType((*DataSchemas)(nil), "proto.memtable.DataSchemas")
	proto.RegisterType((*EntrySchemas)(nil), "proto.memtable.EntrySchemas")
	proto.RegisterType((*InfoSchemas)(nil), "proto.memtable.InfoSchemas")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_871986018790d2fd) }

var fileDescriptor_871986018790d2fd = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x4d, 0x5b, 0xa7, 0xf2, 0x3a, 0xca, 0x90, 0x85, 0x04, 0x95, 0x52, 0xba, 0xea, 0x42,
	0x2a, 0x8c, 0x37, 0x10, 0x5d, 0x08, 0x82, 0x10, 0x4f, 0xf0, 0xaa, 0xaf, 0x5a, 0xa6, 0x4d, 0xa5,
	0x09, 0xc5, 0xae, 0xbc, 0x86, 0x47, 0x72, 0xe9, 0x11, 0xa4, 0x5e, 0x44, 0x9a, 0x76, 0xa4, 0xce,
	0x2a, 0xf9, 0xbf, 0xff, 0x23, 0x3f, 0x04, 0xe0, 0x09, 0x0d, 0xa6, 0xaf, 0x4d, 0x6d, 0x6a, 0x7e,
	0x64, 0x8f, 0xb4, 0xa2, 0xca, 0x60, 0x56, 0x52, 0x5c, 0x42, 0x70, 0x8d, 0x06, 0x1f, 0x1e, 0x5f,
	0xa8, 0x42, 0xcd, 0xcf, 0xc1, 0x21, 0x2d, 0x58, 0xe4, 0x26, 0xc1, 0xfa, 0x2c, 0xfd, 0xef, 0xa6,
	0x37, 0xca, 0x34, 0xdd, 0x64, 0x4a, 0x87, 0x34, 0xbf, 0x00, 0xaf, 0x50, 0x79, 0x2d, 0x9c, 0x88,
	0x25, 0xc1, 0xfa, 0x74, 0xd7, 0xbf, 0x55, 0x79, 0xbd, 0xd5, 0xad, 0x18, 0x67, 0xb0, 0x9c, 0x3f,
	0xc2, 0x8f, 0x61, 0xb1, 0xa1, 0xee, 0x8e, 0x94, 0x60, 0x11, 0x4b, 0x7c, 0x39, 0xa5, 0x81, 0xb7,
	0x58, 0x0e, 0xdc, 0x19, 0xf9, 0x98, 0xf8, 0x0a, 0xdc, 0x0d, 0x75, 0xc2, 0x8d, 0x58, 0xb2, 0x94,
	0xc3, 0x75, 0x20, 0x2d, 0x96, 0xc2, 0x1b, 0x49, 0x8b, 0x65, 0xfc, 0x0e, 0xc1, 0x6c, 0x98, 0x87,
	0x00, 0x15, 0x19, 0xbc, 0xcf, 0x73, 0x4d, 0xc6, 0xce, 0x1c, 0xca, 0x19, 0xe1, 0x02, 0x7c, 0x52,
	0xa6, 0x29, 0x48, 0xdb, 0xad, 0x7d, 0xb9, 0x8d, 0xfc, 0x04, 0x0e, 0xaa, 0x42, 0x49, 0x54, 0xcf,
	0x64, 0x17, 0x7d, 0xf9, 0x97, 0x6d, 0x87, 0x6f, 0x63, 0xe7, 0x4d, 0xdd, 0x94, 0xaf, 0x56, 0x9f,
	0x7d, 0xc8, 0xbe, 0xfa, 0x90, 0x7d, 0xf7, 0x21, 0xfb, 0xf8, 0x09, 0xf7, 0xb2, 0x85, 0xfd, 0x98,
	0xcb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xfd, 0x91, 0x48, 0x89, 0x01, 0x00, 0x00,
}

func (m *DataSchemas) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataSchemas) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataSchemas) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Es) > 0 {
		for iNdEx := len(m.Es) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Es[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *EntrySchemas) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntrySchemas) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntrySchemas) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Val) > 0 {
		i -= len(m.Val)
		copy(dAtA[i:], m.Val)
		i = encodeVarintData(dAtA, i, uint64(len(m.Val)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintData(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ValLen != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.ValLen))
		i--
		dAtA[i] = 0x15
	}
	if m.KeyLen != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.KeyLen))
		i--
		dAtA[i] = 0xd
	}
	return len(dAtA) - i, nil
}

func (m *InfoSchemas) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfoSchemas) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InfoSchemas) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MaxRange != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.MaxRange))
		i--
		dAtA[i] = 0x25
	}
	if m.MinRange != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.MinRange))
		i--
		dAtA[i] = 0x1d
	}
	if m.Entries != 0 {
		i = encodeVarintData(dAtA, i, uint64(m.Entries))
		i--
		dAtA[i] = 0x10
	}
	if m.MetaOffset != 0 {
		i = encodeVarintData(dAtA, i, uint64(m.MetaOffset))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintData(dAtA []byte, offset int, v uint64) int {
	offset -= sovData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataSchemas) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Es) > 0 {
		for _, e := range m.Es {
			l = e.Size()
			n += 1 + l + sovData(uint64(l))
		}
	}
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovData(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EntrySchemas) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.KeyLen != 0 {
		n += 5
	}
	if m.ValLen != 0 {
		n += 5
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	l = len(m.Val)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *InfoSchemas) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MetaOffset != 0 {
		n += 1 + sovData(uint64(m.MetaOffset))
	}
	if m.Entries != 0 {
		n += 1 + sovData(uint64(m.Entries))
	}
	if m.MinRange != 0 {
		n += 5
	}
	if m.MaxRange != 0 {
		n += 5
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozData(x uint64) (n int) {
	return sovData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataSchemas) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
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
			return fmt.Errorf("proto: DataSchemas: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataSchemas: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Es", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Es = append(m.Es, &EntrySchemas{})
			if err := m.Es[len(m.Es)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &InfoSchemas{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthData
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
func (m *EntrySchemas) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
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
			return fmt.Errorf("proto: EntrySchemas: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntrySchemas: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyLen", wireType)
			}
			m.KeyLen = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyLen = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValLen", wireType)
			}
			m.ValLen = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.ValLen = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Val = append(m.Val[:0], dAtA[iNdEx:postIndex]...)
			if m.Val == nil {
				m.Val = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthData
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
func (m *InfoSchemas) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
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
			return fmt.Errorf("proto: InfoSchemas: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfoSchemas: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaOffset", wireType)
			}
			m.MetaOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MetaOffset |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			m.Entries = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Entries |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinRange", wireType)
			}
			m.MinRange = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.MinRange = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRange", wireType)
			}
			m.MaxRange = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxRange = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		default:
			iNdEx = preIndex
			skippy, err := skipData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthData
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
func skipData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowData
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
					return 0, ErrIntOverflowData
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
					return 0, ErrIntOverflowData
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
				return 0, ErrInvalidLengthData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupData = fmt.Errorf("proto: unexpected end of group")
)