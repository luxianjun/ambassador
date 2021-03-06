// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/retry/previous_hosts/v2/previous_hosts.proto

package envoy_config_retry_previous_hosts_v2

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

type PreviousHostsPredicate struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreviousHostsPredicate) Reset()         { *m = PreviousHostsPredicate{} }
func (m *PreviousHostsPredicate) String() string { return proto.CompactTextString(m) }
func (*PreviousHostsPredicate) ProtoMessage()    {}
func (*PreviousHostsPredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4d889d3482108fe, []int{0}
}
func (m *PreviousHostsPredicate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PreviousHostsPredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PreviousHostsPredicate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PreviousHostsPredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreviousHostsPredicate.Merge(m, src)
}
func (m *PreviousHostsPredicate) XXX_Size() int {
	return m.Size()
}
func (m *PreviousHostsPredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_PreviousHostsPredicate.DiscardUnknown(m)
}

var xxx_messageInfo_PreviousHostsPredicate proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PreviousHostsPredicate)(nil), "envoy.config.retry.previous_hosts.v2.PreviousHostsPredicate")
}

func init() {
	proto.RegisterFile("envoy/config/retry/previous_hosts/v2/previous_hosts.proto", fileDescriptor_e4d889d3482108fe)
}

var fileDescriptor_e4d889d3482108fe = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x4a, 0x2d, 0x29, 0xaa, 0xd4, 0x2f,
	0x28, 0x4a, 0x2d, 0xcb, 0xcc, 0x2f, 0x2d, 0x8e, 0xcf, 0xc8, 0x2f, 0x2e, 0x29, 0xd6, 0x2f, 0x33,
	0x42, 0x13, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x01, 0x6b, 0xd5, 0x83, 0x68, 0xd5,
	0x03, 0x6b, 0xd5, 0x43, 0x53, 0x58, 0x66, 0xa4, 0x24, 0xc1, 0x25, 0x16, 0x00, 0x15, 0xf4, 0x00,
	0x89, 0x05, 0x14, 0xa5, 0xa6, 0x64, 0x26, 0x27, 0x96, 0xa4, 0x3a, 0x85, 0x9d, 0x78, 0x24, 0xc7,
	0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x5c, 0x46, 0x99, 0xf9, 0x7a, 0x60, 0x03,
	0x0b, 0x8a, 0xf2, 0x2b, 0x2a, 0xf5, 0x88, 0x31, 0xdb, 0x49, 0x08, 0xcd, 0xe4, 0xfc, 0x92, 0xfc,
	0x00, 0xc6, 0x24, 0x36, 0xb0, 0xf3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x4f, 0xb1,
	0xba, 0xdb, 0x00, 0x00, 0x00,
}

func (m *PreviousHostsPredicate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PreviousHostsPredicate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PreviousHostsPredicate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintPreviousHosts(dAtA []byte, offset int, v uint64) int {
	offset -= sovPreviousHosts(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PreviousHostsPredicate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPreviousHosts(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPreviousHosts(x uint64) (n int) {
	return sovPreviousHosts(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PreviousHostsPredicate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPreviousHosts
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
			return fmt.Errorf("proto: PreviousHostsPredicate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PreviousHostsPredicate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPreviousHosts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPreviousHosts
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPreviousHosts
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
func skipPreviousHosts(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPreviousHosts
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
					return 0, ErrIntOverflowPreviousHosts
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
					return 0, ErrIntOverflowPreviousHosts
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
				return 0, ErrInvalidLengthPreviousHosts
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPreviousHosts
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPreviousHosts
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPreviousHosts        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPreviousHosts          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPreviousHosts = fmt.Errorf("proto: unexpected end of group")
)
