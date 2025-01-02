// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: firmachain/nft/nft_item.proto

package types

import (
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

// An NFT Item
type NftItem struct {
	// Owner
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Id
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// TokenURI
	TokenURI string `protobuf:"bytes,3,opt,name=tokenURI,proto3" json:"tokenURI,omitempty"`
}

func (m *NftItem) Reset()         { *m = NftItem{} }
func (m *NftItem) String() string { return proto.CompactTextString(m) }
func (*NftItem) ProtoMessage()    {}
func (*NftItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfa18a50b429599c, []int{0}
}
func (m *NftItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NftItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NftItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NftItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NftItem.Merge(m, src)
}
func (m *NftItem) XXX_Size() int {
	return m.Size()
}
func (m *NftItem) XXX_DiscardUnknown() {
	xxx_messageInfo_NftItem.DiscardUnknown(m)
}

var xxx_messageInfo_NftItem proto.InternalMessageInfo

func (m *NftItem) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *NftItem) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NftItem) GetTokenURI() string {
	if m != nil {
		return m.TokenURI
	}
	return ""
}

func init() {
	proto.RegisterType((*NftItem)(nil), "firmachain.firmachain.nft.NftItem")
}

func init() { proto.RegisterFile("firmachain/nft/nft_item.proto", fileDescriptor_dfa18a50b429599c) }

var fileDescriptor_dfa18a50b429599c = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xcb, 0x2c, 0xca,
	0x4d, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xcf, 0x4b, 0x2b, 0x01, 0xe1, 0xf8, 0xcc, 0x92, 0xd4,
	0x5c, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x49, 0x84, 0xb4, 0x1e, 0x12, 0x33, 0x2f, 0xad,
	0x44, 0xc9, 0x9b, 0x8b, 0xdd, 0x2f, 0xad, 0xc4, 0xb3, 0x24, 0x35, 0x57, 0x48, 0x84, 0x8b, 0x35,
	0xbf, 0x3c, 0x2f, 0xb5, 0x48, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x11, 0xe2, 0xe3,
	0x62, 0xca, 0x4c, 0x91, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x09, 0x62, 0xca, 0x4c, 0x11, 0x92, 0xe2,
	0xe2, 0x28, 0xc9, 0xcf, 0x4e, 0xcd, 0x0b, 0x0d, 0xf2, 0x94, 0x60, 0x06, 0x2b, 0x84, 0xf3, 0x9d,
	0xbc, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f,
	0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x20, 0x3d, 0xb3, 0x24,
	0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xc9, 0xad, 0x48, 0xcc, 0x32, 0x03, 0x53, 0xfd,
	0x0a, 0xb0, 0xe3, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x4e, 0x37, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x96, 0xe2, 0x69, 0x0d, 0xdb, 0x00, 0x00, 0x00,
}

func (m *NftItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NftItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenURI) > 0 {
		i -= len(m.TokenURI)
		copy(dAtA[i:], m.TokenURI)
		i = encodeVarintNftItem(dAtA, i, uint64(len(m.TokenURI)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintNftItem(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintNftItem(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftItem(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftItem(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NftItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovNftItem(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovNftItem(uint64(m.Id))
	}
	l = len(m.TokenURI)
	if l > 0 {
		n += 1 + l + sovNftItem(uint64(l))
	}
	return n
}

func sovNftItem(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftItem(x uint64) (n int) {
	return sovNftItem(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NftItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftItem
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
			return fmt.Errorf("proto: NftItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NftItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftItem
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
				return ErrInvalidLengthNftItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenURI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftItem
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
				return ErrInvalidLengthNftItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenURI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftItem
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
func skipNftItem(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftItem
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
					return 0, ErrIntOverflowNftItem
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
					return 0, ErrIntOverflowNftItem
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
				return 0, ErrInvalidLengthNftItem
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftItem
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftItem
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftItem        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftItem          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftItem = fmt.Errorf("proto: unexpected end of group")
)
