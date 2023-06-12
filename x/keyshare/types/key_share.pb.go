// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/keyshare/key_share.proto

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

type KeyShare struct {
	Validator           string `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	BlockHeight         uint64 `protobuf:"varint,2,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	Commitment          string `protobuf:"bytes,3,opt,name=commitment,proto3" json:"commitment,omitempty"`
	KeyShare            string `protobuf:"bytes,4,opt,name=keyShare,proto3" json:"keyShare,omitempty"`
	KeyShareIndex       uint64 `protobuf:"varint,5,opt,name=keyShareIndex,proto3" json:"keyShareIndex,omitempty"`
	ReceivedTimestamp   uint64 `protobuf:"varint,6,opt,name=receivedTimestamp,proto3" json:"receivedTimestamp,omitempty"`
	ReceivedBlockHeight uint64 `protobuf:"varint,7,opt,name=receivedBlockHeight,proto3" json:"receivedBlockHeight,omitempty"`
}

func (m *KeyShare) Reset()         { *m = KeyShare{} }
func (m *KeyShare) String() string { return proto.CompactTextString(m) }
func (*KeyShare) ProtoMessage()    {}
func (*KeyShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb45212b5123dd29, []int{0}
}
func (m *KeyShare) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyShare.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyShare.Merge(m, src)
}
func (m *KeyShare) XXX_Size() int {
	return m.Size()
}
func (m *KeyShare) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyShare.DiscardUnknown(m)
}

var xxx_messageInfo_KeyShare proto.InternalMessageInfo

func (m *KeyShare) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *KeyShare) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *KeyShare) GetCommitment() string {
	if m != nil {
		return m.Commitment
	}
	return ""
}

func (m *KeyShare) GetKeyShare() string {
	if m != nil {
		return m.KeyShare
	}
	return ""
}

func (m *KeyShare) GetKeyShareIndex() uint64 {
	if m != nil {
		return m.KeyShareIndex
	}
	return 0
}

func (m *KeyShare) GetReceivedTimestamp() uint64 {
	if m != nil {
		return m.ReceivedTimestamp
	}
	return 0
}

func (m *KeyShare) GetReceivedBlockHeight() uint64 {
	if m != nil {
		return m.ReceivedBlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*KeyShare)(nil), "fairyring.keyshare.KeyShare")
}

func init() {
	proto.RegisterFile("fairyring/keyshare/key_share.proto", fileDescriptor_cb45212b5123dd29)
}

var fileDescriptor_cb45212b5123dd29 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0x4b, 0xcc, 0x2c,
	0xaa, 0x2c, 0xca, 0xcc, 0x4b, 0xd7, 0xcf, 0x4e, 0xad, 0x2c, 0xce, 0x48, 0x2c, 0x4a, 0x05, 0x31,
	0xe2, 0xc1, 0x2c, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x21, 0xb8, 0x1a, 0x3d, 0x98, 0x1a,
	0xa5, 0x1e, 0x26, 0x2e, 0x0e, 0xef, 0xd4, 0xca, 0x60, 0x10, 0x47, 0x48, 0x86, 0x8b, 0xb3, 0x2c,
	0x31, 0x27, 0x33, 0x25, 0xb1, 0x24, 0xbf, 0x48, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0x21,
	0x20, 0xa4, 0xc0, 0xc5, 0x9d, 0x94, 0x93, 0x9f, 0x9c, 0xed, 0x91, 0x9a, 0x99, 0x9e, 0x51, 0x22,
	0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x12, 0x84, 0x2c, 0x24, 0x24, 0xc7, 0xc5, 0x95, 0x9c, 0x9f, 0x9b,
	0x9b, 0x59, 0x92, 0x9b, 0x9a, 0x57, 0x22, 0xc1, 0x0c, 0x36, 0x00, 0x49, 0x44, 0x48, 0x8a, 0x8b,
	0x23, 0x1b, 0x6a, 0x97, 0x04, 0x0b, 0x58, 0x16, 0xce, 0x17, 0x52, 0xe1, 0xe2, 0x85, 0xb1, 0x3d,
	0xf3, 0x52, 0x52, 0x2b, 0x24, 0x58, 0xc1, 0xe6, 0xa3, 0x0a, 0x0a, 0xe9, 0x70, 0x09, 0x16, 0xa5,
	0x26, 0xa7, 0x66, 0x96, 0xa5, 0xa6, 0x84, 0x64, 0xe6, 0xa6, 0x16, 0x97, 0x24, 0xe6, 0x16, 0x48,
	0xb0, 0x81, 0x55, 0x62, 0x4a, 0x08, 0x19, 0x70, 0x09, 0xc3, 0x04, 0x9d, 0x90, 0x5c, 0xce, 0x0e,
	0x56, 0x8f, 0x4d, 0xca, 0xc9, 0xe4, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2,
	0xa4, 0x10, 0x01, 0x5c, 0x81, 0x08, 0xe2, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0xf8,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x9a, 0x2b, 0xcb, 0x85, 0x01, 0x00, 0x00,
}

func (m *KeyShare) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyShare) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyShare) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReceivedBlockHeight != 0 {
		i = encodeVarintKeyShare(dAtA, i, uint64(m.ReceivedBlockHeight))
		i--
		dAtA[i] = 0x38
	}
	if m.ReceivedTimestamp != 0 {
		i = encodeVarintKeyShare(dAtA, i, uint64(m.ReceivedTimestamp))
		i--
		dAtA[i] = 0x30
	}
	if m.KeyShareIndex != 0 {
		i = encodeVarintKeyShare(dAtA, i, uint64(m.KeyShareIndex))
		i--
		dAtA[i] = 0x28
	}
	if len(m.KeyShare) > 0 {
		i -= len(m.KeyShare)
		copy(dAtA[i:], m.KeyShare)
		i = encodeVarintKeyShare(dAtA, i, uint64(len(m.KeyShare)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Commitment) > 0 {
		i -= len(m.Commitment)
		copy(dAtA[i:], m.Commitment)
		i = encodeVarintKeyShare(dAtA, i, uint64(len(m.Commitment)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintKeyShare(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintKeyShare(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintKeyShare(dAtA []byte, offset int, v uint64) int {
	offset -= sovKeyShare(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KeyShare) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovKeyShare(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovKeyShare(uint64(m.BlockHeight))
	}
	l = len(m.Commitment)
	if l > 0 {
		n += 1 + l + sovKeyShare(uint64(l))
	}
	l = len(m.KeyShare)
	if l > 0 {
		n += 1 + l + sovKeyShare(uint64(l))
	}
	if m.KeyShareIndex != 0 {
		n += 1 + sovKeyShare(uint64(m.KeyShareIndex))
	}
	if m.ReceivedTimestamp != 0 {
		n += 1 + sovKeyShare(uint64(m.ReceivedTimestamp))
	}
	if m.ReceivedBlockHeight != 0 {
		n += 1 + sovKeyShare(uint64(m.ReceivedBlockHeight))
	}
	return n
}

func sovKeyShare(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKeyShare(x uint64) (n int) {
	return sovKeyShare(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyShare) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeyShare
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
			return fmt.Errorf("proto: KeyShare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyShare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyShare
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
				return ErrInvalidLengthKeyShare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeyShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commitment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyShare
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
				return ErrInvalidLengthKeyShare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeyShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commitment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyShare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyShare
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
				return ErrInvalidLengthKeyShare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeyShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyShare = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyShareIndex", wireType)
			}
			m.KeyShareIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyShareIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceivedTimestamp", wireType)
			}
			m.ReceivedTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReceivedTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceivedBlockHeight", wireType)
			}
			m.ReceivedBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReceivedBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKeyShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeyShare
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
func skipKeyShare(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeyShare
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
					return 0, ErrIntOverflowKeyShare
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
					return 0, ErrIntOverflowKeyShare
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
				return 0, ErrInvalidLengthKeyShare
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKeyShare
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKeyShare
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKeyShare        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeyShare          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKeyShare = fmt.Errorf("proto: unexpected end of group")
)