// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/retry/priority/previous_priorities/v3/previous_priorities_config.proto

package envoy_extensions_retry_priority_previous_priorities_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// A retry host selector that attempts to spread retries between priorities, even if certain
// priorities would not normally be attempted due to higher priorities being available.
//
// As priorities get excluded, load will be distributed amongst the remaining healthy priorities
// based on the relative health of the priorities, matching how load is distributed during regular
// host selection. For example, given priority healths of {100, 50, 50}, the original load will be
// {100, 0, 0} (since P0 has capacity to handle 100% of the traffic). If P0 is excluded, the load
// changes to {0, 50, 50}, because P1 is only able to handle 50% of the traffic, causing the
// remaining to spill over to P2.
//
// Each priority attempted will be excluded until there are no healthy priorities left, at which
// point the list of attempted priorities will be reset, essentially starting from the beginning.
// For example, given three priorities P0, P1, P2 with healthy % of 100, 0 and 50 respectively, the
// following sequence of priorities would be selected (assuming update_frequency = 1):
// Attempt 1: P0 (P0 is 100% healthy)
// Attempt 2: P2 (P0 already attempted, P2 only healthy priority)
// Attempt 3: P0 (no healthy priorities, reset)
// Attempt 4: P2
//
// In the case of all upstream hosts being unhealthy, no adjustments will be made to the original
// priority load, so behavior should be identical to not using this plugin.
//
// Using this PriorityFilter requires rebuilding the priority load, which runs in O(# of
// priorities), which might incur significant overhead for clusters with many priorities.
// [#extension: envoy.retry_priorities.previous_priorities]
type PreviousPrioritiesConfig struct {
	// How often the priority load should be updated based on previously attempted priorities. Useful
	// to allow each priorities to receive more than one request before being excluded or to reduce
	// the number of times that the priority load has to be recomputed.
	//
	// For example, by setting this to 2, then the first two attempts (initial attempt and first
	// retry) will use the unmodified priority load. The third and fourth attempt will use priority
	// load which excludes the priorities routed to with the first two attempts, and the fifth and
	// sixth attempt will use the priority load excluding the priorities used for the first four
	// attempts.
	//
	// Must be greater than 0.
	UpdateFrequency      int32    `protobuf:"varint,1,opt,name=update_frequency,json=updateFrequency,proto3" json:"update_frequency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreviousPrioritiesConfig) Reset()         { *m = PreviousPrioritiesConfig{} }
func (m *PreviousPrioritiesConfig) String() string { return proto.CompactTextString(m) }
func (*PreviousPrioritiesConfig) ProtoMessage()    {}
func (*PreviousPrioritiesConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b022e36ffd9806c, []int{0}
}
func (m *PreviousPrioritiesConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PreviousPrioritiesConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PreviousPrioritiesConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PreviousPrioritiesConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreviousPrioritiesConfig.Merge(m, src)
}
func (m *PreviousPrioritiesConfig) XXX_Size() int {
	return m.Size()
}
func (m *PreviousPrioritiesConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PreviousPrioritiesConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PreviousPrioritiesConfig proto.InternalMessageInfo

func (m *PreviousPrioritiesConfig) GetUpdateFrequency() int32 {
	if m != nil {
		return m.UpdateFrequency
	}
	return 0
}

func init() {
	proto.RegisterType((*PreviousPrioritiesConfig)(nil), "envoy.extensions.retry.priority.previous_priorities.v3.PreviousPrioritiesConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/retry/priority/previous_priorities/v3/previous_priorities_config.proto", fileDescriptor_1b022e36ffd9806c)
}

var fileDescriptor_1b022e36ffd9806c = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x0a, 0x4f, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0x4a, 0x2d,
	0x29, 0xaa, 0xd4, 0x2f, 0x28, 0xca, 0xcc, 0x2f, 0xca, 0x2c, 0x01, 0x31, 0x52, 0xcb, 0x32, 0xf3,
	0x4b, 0x8b, 0xe3, 0xa1, 0x22, 0x99, 0xa9, 0xc5, 0xfa, 0x65, 0xc6, 0xd8, 0x84, 0xe3, 0x93, 0xf3,
	0xf3, 0xd2, 0x32, 0xd3, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xcc, 0xc0, 0x06, 0xeb, 0x21,
	0x0c, 0xd6, 0x03, 0x1b, 0xac, 0x07, 0x33, 0x58, 0x0f, 0x8b, 0x09, 0x7a, 0x65, 0xc6, 0x52, 0x8a,
	0xa5, 0x29, 0x05, 0x89, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x60, 0x07, 0x95, 0xa5,
	0x16, 0x81, 0x0c, 0xc8, 0xcc, 0x83, 0x1a, 0x2d, 0x25, 0x5e, 0x96, 0x98, 0x93, 0x99, 0x92, 0x58,
	0x92, 0xaa, 0x0f, 0x63, 0x40, 0x24, 0x94, 0xa6, 0x31, 0x72, 0x49, 0x04, 0x40, 0x8d, 0x0d, 0x80,
	0x9b, 0xea, 0x0c, 0x76, 0x96, 0x90, 0x11, 0x97, 0x40, 0x69, 0x01, 0x48, 0x71, 0x7c, 0x5a, 0x51,
	0x6a, 0x61, 0x69, 0x6a, 0x5e, 0x72, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xab, 0x13, 0xfb, 0x2f,
	0x27, 0x16, 0x29, 0x26, 0x05, 0x86, 0x20, 0x7e, 0x88, 0x02, 0x37, 0x98, 0xbc, 0x95, 0xdb, 0xac,
	0xa3, 0x1d, 0x72, 0x8e, 0x5c, 0xf6, 0x10, 0xbf, 0x40, 0xfd, 0x07, 0xf3, 0x07, 0xa6, 0xf3, 0x71,
	0xd9, 0xed, 0x54, 0x7c, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31,
	0x72, 0xb9, 0x64, 0xe6, 0xeb, 0x81, 0x4d, 0x2c, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x23, 0x2f, 0xa0,
	0x9c, 0x64, 0x71, 0xd9, 0x16, 0x00, 0x0a, 0x8b, 0x00, 0xc6, 0x24, 0x36, 0x70, 0xa0, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x11, 0x93, 0x98, 0xe3, 0x01, 0x00, 0x00,
}

func (m *PreviousPrioritiesConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PreviousPrioritiesConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PreviousPrioritiesConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.UpdateFrequency != 0 {
		i = encodeVarintPreviousPrioritiesConfig(dAtA, i, uint64(m.UpdateFrequency))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPreviousPrioritiesConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovPreviousPrioritiesConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PreviousPrioritiesConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpdateFrequency != 0 {
		n += 1 + sovPreviousPrioritiesConfig(uint64(m.UpdateFrequency))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPreviousPrioritiesConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPreviousPrioritiesConfig(x uint64) (n int) {
	return sovPreviousPrioritiesConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PreviousPrioritiesConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPreviousPrioritiesConfig
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
			return fmt.Errorf("proto: PreviousPrioritiesConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PreviousPrioritiesConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateFrequency", wireType)
			}
			m.UpdateFrequency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPreviousPrioritiesConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateFrequency |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPreviousPrioritiesConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPreviousPrioritiesConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPreviousPrioritiesConfig
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
func skipPreviousPrioritiesConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPreviousPrioritiesConfig
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
					return 0, ErrIntOverflowPreviousPrioritiesConfig
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
					return 0, ErrIntOverflowPreviousPrioritiesConfig
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
				return 0, ErrInvalidLengthPreviousPrioritiesConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPreviousPrioritiesConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPreviousPrioritiesConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPreviousPrioritiesConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPreviousPrioritiesConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPreviousPrioritiesConfig = fmt.Errorf("proto: unexpected end of group")
)
