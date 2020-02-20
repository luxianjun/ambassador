// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/v3/semantic_version.proto

package envoy_type_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

// Envoy uses SemVer (https://semver.org/). Major/minor versions indicate
// expected behaviors and APIs, the patch version field is used only
// for security fixes and can be generally ignored.
type SemanticVersion struct {
	MajorNumber          uint32   `protobuf:"varint,1,opt,name=major_number,json=majorNumber,proto3" json:"major_number,omitempty"`
	MinorNumber          uint32   `protobuf:"varint,2,opt,name=minor_number,json=minorNumber,proto3" json:"minor_number,omitempty"`
	Patch                uint32   `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SemanticVersion) Reset()         { *m = SemanticVersion{} }
func (m *SemanticVersion) String() string { return proto.CompactTextString(m) }
func (*SemanticVersion) ProtoMessage()    {}
func (*SemanticVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d8ec35e89ef71eb, []int{0}
}
func (m *SemanticVersion) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SemanticVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SemanticVersion.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SemanticVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SemanticVersion.Merge(m, src)
}
func (m *SemanticVersion) XXX_Size() int {
	return m.Size()
}
func (m *SemanticVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_SemanticVersion.DiscardUnknown(m)
}

var xxx_messageInfo_SemanticVersion proto.InternalMessageInfo

func (m *SemanticVersion) GetMajorNumber() uint32 {
	if m != nil {
		return m.MajorNumber
	}
	return 0
}

func (m *SemanticVersion) GetMinorNumber() uint32 {
	if m != nil {
		return m.MinorNumber
	}
	return 0
}

func (m *SemanticVersion) GetPatch() uint32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

func init() {
	proto.RegisterType((*SemanticVersion)(nil), "envoy.type.v3.SemanticVersion")
}

func init() {
	proto.RegisterFile("envoy/type/v3/semantic_version.proto", fileDescriptor_5d8ec35e89ef71eb)
}

var fileDescriptor_5d8ec35e89ef71eb = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0x33, 0xd6, 0x2f, 0x4e, 0xcd, 0x4d, 0xcc, 0x2b,
	0xc9, 0x4c, 0x8e, 0x2f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x05, 0xab, 0xd2, 0x03, 0xa9, 0xd2, 0x2b, 0x33, 0x96, 0x52, 0x2c, 0x4d, 0x29, 0x48,
	0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x87, 0x2a, 0xcf,
	0xcc, 0x4b, 0x87, 0xe8, 0x50, 0x9a, 0xc0, 0xc8, 0xc5, 0x1f, 0x0c, 0x35, 0x2c, 0x0c, 0x22, 0x29,
	0xa4, 0xc8, 0xc5, 0x93, 0x9b, 0x98, 0x95, 0x5f, 0x14, 0x9f, 0x57, 0x9a, 0x9b, 0x94, 0x5a, 0x24,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1b, 0xc4, 0x0d, 0x16, 0xf3, 0x03, 0x0b, 0x81, 0x95, 0x64, 0xe6,
	0x21, 0x94, 0x30, 0x41, 0x95, 0x80, 0xc4, 0xa0, 0x4a, 0x44, 0xb8, 0x58, 0x0b, 0x12, 0x4b, 0x92,
	0x33, 0x24, 0x98, 0xc1, 0x72, 0x10, 0x8e, 0x95, 0xe2, 0xac, 0xa3, 0x1d, 0x72, 0x32, 0x5c, 0x52,
	0x48, 0x0e, 0x45, 0xb3, 0xde, 0xc9, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0xe4, 0x92, 0xce, 0xcc, 0xd7, 0x03, 0x2b, 0x2e, 0x28, 0xca, 0xaf, 0xa8, 0xd4,
	0x43, 0xf1, 0xa0, 0x93, 0x08, 0x9a, 0xde, 0x00, 0x90, 0x9f, 0x02, 0x18, 0x93, 0xd8, 0xc0, 0x9e,
	0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x0e, 0x51, 0xad, 0x36, 0x01, 0x00, 0x00,
}

func (m *SemanticVersion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SemanticVersion) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SemanticVersion) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Patch != 0 {
		i = encodeVarintSemanticVersion(dAtA, i, uint64(m.Patch))
		i--
		dAtA[i] = 0x18
	}
	if m.MinorNumber != 0 {
		i = encodeVarintSemanticVersion(dAtA, i, uint64(m.MinorNumber))
		i--
		dAtA[i] = 0x10
	}
	if m.MajorNumber != 0 {
		i = encodeVarintSemanticVersion(dAtA, i, uint64(m.MajorNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSemanticVersion(dAtA []byte, offset int, v uint64) int {
	offset -= sovSemanticVersion(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SemanticVersion) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MajorNumber != 0 {
		n += 1 + sovSemanticVersion(uint64(m.MajorNumber))
	}
	if m.MinorNumber != 0 {
		n += 1 + sovSemanticVersion(uint64(m.MinorNumber))
	}
	if m.Patch != 0 {
		n += 1 + sovSemanticVersion(uint64(m.Patch))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSemanticVersion(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSemanticVersion(x uint64) (n int) {
	return sovSemanticVersion(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SemanticVersion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSemanticVersion
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
			return fmt.Errorf("proto: SemanticVersion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SemanticVersion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MajorNumber", wireType)
			}
			m.MajorNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSemanticVersion
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MajorNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinorNumber", wireType)
			}
			m.MinorNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSemanticVersion
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinorNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Patch", wireType)
			}
			m.Patch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSemanticVersion
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Patch |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSemanticVersion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSemanticVersion
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSemanticVersion
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
func skipSemanticVersion(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSemanticVersion
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
					return 0, ErrIntOverflowSemanticVersion
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
					return 0, ErrIntOverflowSemanticVersion
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
				return 0, ErrInvalidLengthSemanticVersion
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSemanticVersion
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSemanticVersion
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
				next, err := skipSemanticVersion(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSemanticVersion
				}
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
	ErrInvalidLengthSemanticVersion = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSemanticVersion   = fmt.Errorf("proto: integer overflow")
)