// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/network/sni_cluster/v2/sni_cluster.proto

package envoy_config_filter_network_sni_cluster_v2

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

type SniCluster struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SniCluster) Reset()         { *m = SniCluster{} }
func (m *SniCluster) String() string { return proto.CompactTextString(m) }
func (*SniCluster) ProtoMessage()    {}
func (*SniCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_c62c0a0da3b665be, []int{0}
}
func (m *SniCluster) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SniCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SniCluster.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SniCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SniCluster.Merge(m, src)
}
func (m *SniCluster) XXX_Size() int {
	return m.Size()
}
func (m *SniCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_SniCluster.DiscardUnknown(m)
}

var xxx_messageInfo_SniCluster proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SniCluster)(nil), "envoy.config.filter.network.sni_cluster.v2.SniCluster")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/sni_cluster/v2/sni_cluster.proto", fileDescriptor_c62c0a0da3b665be)
}

var fileDescriptor_c62c0a0da3b665be = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2f, 0xce, 0xcb, 0x8c, 0x4f, 0xce, 0x29, 0x2d,
	0x06, 0x89, 0x95, 0x19, 0x21, 0x73, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb4, 0xc0, 0xba,
	0xf5, 0x20, 0xba, 0xf5, 0x20, 0xba, 0xf5, 0xa0, 0xba, 0xf5, 0x90, 0x95, 0x97, 0x19, 0x49, 0xc9,
	0x95, 0xa6, 0x14, 0x24, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15,
	0xeb, 0xe7, 0x66, 0xa6, 0x17, 0x25, 0x96, 0xa4, 0x42, 0xcc, 0x52, 0xe2, 0xe1, 0xe2, 0x0a, 0xce,
	0xcb, 0x74, 0x86, 0x68, 0x70, 0x9a, 0xc0, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0x7e, 0x9a, 0xf1, 0xaf, 0x9f, 0xd5, 0x50, 0x48, 0x1f, 0x62, 0x5d, 0x6a,
	0x45, 0x49, 0x6a, 0x5e, 0x31, 0xc8, 0x08, 0xa8, 0x95, 0xc5, 0xd8, 0xed, 0x34, 0xe6, 0xb2, 0xc8,
	0xcc, 0xd7, 0x03, 0xeb, 0x29, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x23, 0xde, 0xb5, 0x4e, 0xfc, 0x08,
	0xb7, 0x04, 0x80, 0x9c, 0x17, 0xc0, 0x98, 0xc4, 0x06, 0x76, 0xa7, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xea, 0xd4, 0x8f, 0x19, 0x33, 0x01, 0x00, 0x00,
}

func (m *SniCluster) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SniCluster) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SniCluster) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func encodeVarintSniCluster(dAtA []byte, offset int, v uint64) int {
	offset -= sovSniCluster(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SniCluster) Size() (n int) {
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

func sovSniCluster(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSniCluster(x uint64) (n int) {
	return sovSniCluster(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SniCluster) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSniCluster
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
			return fmt.Errorf("proto: SniCluster: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SniCluster: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipSniCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSniCluster
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSniCluster
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
func skipSniCluster(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSniCluster
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
					return 0, ErrIntOverflowSniCluster
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
					return 0, ErrIntOverflowSniCluster
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
				return 0, ErrInvalidLengthSniCluster
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSniCluster
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSniCluster
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
				next, err := skipSniCluster(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSniCluster
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
	ErrInvalidLengthSniCluster = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSniCluster   = fmt.Errorf("proto: integer overflow")
)