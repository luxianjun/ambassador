// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/cluster/v3/circuit_breaker.proto

package envoy_config_cluster_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/config/core/v3"
	v31 "github.com/datawire/ambassador/pkg/api/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// :ref:`Circuit breaking<arch_overview_circuit_break>` settings can be
// specified individually for each defined priority.
type CircuitBreakers struct {
	// If multiple :ref:`Thresholds<envoy_api_msg_config.cluster.v3.CircuitBreakers.Thresholds>`
	// are defined with the same :ref:`RoutingPriority<envoy_api_enum_config.core.v3.RoutingPriority>`,
	// the first one in the list is used. If no Thresholds is defined for a given
	// :ref:`RoutingPriority<envoy_api_enum_config.core.v3.RoutingPriority>`, the default values
	// are used.
	Thresholds           []*CircuitBreakers_Thresholds `protobuf:"bytes,1,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CircuitBreakers) Reset()         { *m = CircuitBreakers{} }
func (m *CircuitBreakers) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers) ProtoMessage()    {}
func (*CircuitBreakers) Descriptor() ([]byte, []int) {
	return fileDescriptor_131e6721a0838844, []int{0}
}
func (m *CircuitBreakers) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CircuitBreakers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CircuitBreakers.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CircuitBreakers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers.Merge(m, src)
}
func (m *CircuitBreakers) XXX_Size() int {
	return m.Size()
}
func (m *CircuitBreakers) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers proto.InternalMessageInfo

func (m *CircuitBreakers) GetThresholds() []*CircuitBreakers_Thresholds {
	if m != nil {
		return m.Thresholds
	}
	return nil
}

// A Thresholds defines CircuitBreaker settings for a
// :ref:`RoutingPriority<envoy_api_enum_config.core.v3.RoutingPriority>`.
// [#next-free-field: 9]
type CircuitBreakers_Thresholds struct {
	// The :ref:`RoutingPriority<envoy_api_enum_config.core.v3.RoutingPriority>`
	// the specified CircuitBreaker settings apply to.
	Priority v3.RoutingPriority `protobuf:"varint,1,opt,name=priority,proto3,enum=envoy.config.core.v3.RoutingPriority" json:"priority,omitempty"`
	// The maximum number of connections that Envoy will make to the upstream
	// cluster. If not specified, the default is 1024.
	MaxConnections *types.UInt32Value `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	// The maximum number of pending requests that Envoy will allow to the
	// upstream cluster. If not specified, the default is 1024.
	MaxPendingRequests *types.UInt32Value `protobuf:"bytes,3,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	// The maximum number of parallel requests that Envoy will make to the
	// upstream cluster. If not specified, the default is 1024.
	MaxRequests *types.UInt32Value `protobuf:"bytes,4,opt,name=max_requests,json=maxRequests,proto3" json:"max_requests,omitempty"`
	// The maximum number of parallel retries that Envoy will allow to the
	// upstream cluster. If not specified, the default is 3.
	MaxRetries *types.UInt32Value `protobuf:"bytes,5,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// Specifies a limit on concurrent retries in relation to the number of active requests. This
	// parameter is optional.
	//
	// .. note::
	//
	//    If this field is set, the retry budget will override any configured retry circuit
	//    breaker.
	RetryBudget *CircuitBreakers_Thresholds_RetryBudget `protobuf:"bytes,8,opt,name=retry_budget,json=retryBudget,proto3" json:"retry_budget,omitempty"`
	// If track_remaining is true, then stats will be published that expose
	// the number of resources remaining until the circuit breakers open. If
	// not specified, the default is false.
	//
	// .. note::
	//
	//    If a retry budget is used in lieu of the max_retries circuit breaker,
	//    the remaining retry resources remaining will not be tracked.
	TrackRemaining bool `protobuf:"varint,6,opt,name=track_remaining,json=trackRemaining,proto3" json:"track_remaining,omitempty"`
	// The maximum number of connection pools per cluster that Envoy will concurrently support at
	// once. If not specified, the default is unlimited. Set this for clusters which create a
	// large number of connection pools. See
	// :ref:`Circuit Breaking <arch_overview_circuit_break_cluster_maximum_connection_pools>` for
	// more details.
	MaxConnectionPools   *types.UInt32Value `protobuf:"bytes,7,opt,name=max_connection_pools,json=maxConnectionPools,proto3" json:"max_connection_pools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CircuitBreakers_Thresholds) Reset()         { *m = CircuitBreakers_Thresholds{} }
func (m *CircuitBreakers_Thresholds) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds) Descriptor() ([]byte, []int) {
	return fileDescriptor_131e6721a0838844, []int{0, 0}
}
func (m *CircuitBreakers_Thresholds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CircuitBreakers_Thresholds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CircuitBreakers_Thresholds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CircuitBreakers_Thresholds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers_Thresholds.Merge(m, src)
}
func (m *CircuitBreakers_Thresholds) XXX_Size() int {
	return m.Size()
}
func (m *CircuitBreakers_Thresholds) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers_Thresholds.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers_Thresholds proto.InternalMessageInfo

func (m *CircuitBreakers_Thresholds) GetPriority() v3.RoutingPriority {
	if m != nil {
		return m.Priority
	}
	return v3.RoutingPriority_DEFAULT
}

func (m *CircuitBreakers_Thresholds) GetMaxConnections() *types.UInt32Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxPendingRequests() *types.UInt32Value {
	if m != nil {
		return m.MaxPendingRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRequests() *types.UInt32Value {
	if m != nil {
		return m.MaxRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRetries() *types.UInt32Value {
	if m != nil {
		return m.MaxRetries
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetRetryBudget() *CircuitBreakers_Thresholds_RetryBudget {
	if m != nil {
		return m.RetryBudget
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetTrackRemaining() bool {
	if m != nil {
		return m.TrackRemaining
	}
	return false
}

func (m *CircuitBreakers_Thresholds) GetMaxConnectionPools() *types.UInt32Value {
	if m != nil {
		return m.MaxConnectionPools
	}
	return nil
}

type CircuitBreakers_Thresholds_RetryBudget struct {
	// Specifies the limit on concurrent retries as a percentage of the sum of active requests and
	// active pending requests. For example, if there are 100 active requests and the
	// budget_percent is set to 25, there may be 25 active retries.
	//
	// This parameter is optional. Defaults to 20%.
	BudgetPercent *v31.Percent `protobuf:"bytes,1,opt,name=budget_percent,json=budgetPercent,proto3" json:"budget_percent,omitempty"`
	// Specifies the minimum retry concurrency allowed for the retry budget. The limit on the
	// number of active retries may never go below this number.
	//
	// This parameter is optional. Defaults to 3.
	MinRetryConcurrency  *types.UInt32Value `protobuf:"bytes,2,opt,name=min_retry_concurrency,json=minRetryConcurrency,proto3" json:"min_retry_concurrency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CircuitBreakers_Thresholds_RetryBudget) Reset() {
	*m = CircuitBreakers_Thresholds_RetryBudget{}
}
func (m *CircuitBreakers_Thresholds_RetryBudget) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds_RetryBudget) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds_RetryBudget) Descriptor() ([]byte, []int) {
	return fileDescriptor_131e6721a0838844, []int{0, 0, 0}
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Merge(m, src)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Size() int {
	return m.Size()
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget proto.InternalMessageInfo

func (m *CircuitBreakers_Thresholds_RetryBudget) GetBudgetPercent() *v31.Percent {
	if m != nil {
		return m.BudgetPercent
	}
	return nil
}

func (m *CircuitBreakers_Thresholds_RetryBudget) GetMinRetryConcurrency() *types.UInt32Value {
	if m != nil {
		return m.MinRetryConcurrency
	}
	return nil
}

func init() {
	proto.RegisterType((*CircuitBreakers)(nil), "envoy.config.cluster.v3.CircuitBreakers")
	proto.RegisterType((*CircuitBreakers_Thresholds)(nil), "envoy.config.cluster.v3.CircuitBreakers.Thresholds")
	proto.RegisterType((*CircuitBreakers_Thresholds_RetryBudget)(nil), "envoy.config.cluster.v3.CircuitBreakers.Thresholds.RetryBudget")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/v3/circuit_breaker.proto", fileDescriptor_131e6721a0838844)
}

var fileDescriptor_131e6721a0838844 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xe5, 0x8d, 0x8d, 0xca, 0x19, 0x1d, 0xca, 0x80, 0x45, 0x05, 0x95, 0x82, 0x18, 0x54,
	0x42, 0x38, 0xa2, 0x95, 0x38, 0x0c, 0x4d, 0x93, 0x32, 0x71, 0x00, 0x24, 0x14, 0x85, 0x3f, 0xd7,
	0xc8, 0x4d, 0xbd, 0xcc, 0x5a, 0x63, 0x07, 0xdb, 0x09, 0xcd, 0x0d, 0x71, 0xe2, 0x33, 0xf0, 0x7d,
	0x90, 0x38, 0xf2, 0x11, 0x50, 0x8f, 0x7c, 0x04, 0x4e, 0xc8, 0x76, 0xda, 0x6c, 0x93, 0x26, 0x05,
	0x6e, 0x71, 0xde, 0xe7, 0xf9, 0xd9, 0xef, 0xe3, 0xd7, 0xf0, 0x09, 0x61, 0x25, 0xaf, 0xfc, 0x84,
	0xb3, 0x63, 0x9a, 0xfa, 0xc9, 0xac, 0x90, 0x8a, 0x08, 0xbf, 0x1c, 0xfb, 0x09, 0x15, 0x49, 0x41,
	0x55, 0x3c, 0x11, 0x04, 0x9f, 0x12, 0x81, 0x72, 0xc1, 0x15, 0x77, 0x77, 0x8d, 0x1c, 0x59, 0x39,
	0xaa, 0xe5, 0xa8, 0x1c, 0xf7, 0xee, 0x9e, 0xe7, 0x70, 0x41, 0x34, 0x64, 0x82, 0x25, 0xb1, 0xce,
	0xde, 0x6d, 0x2b, 0x50, 0x55, 0x6e, 0x2a, 0x39, 0x11, 0x09, 0x61, 0xaa, 0x2e, 0xf6, 0x53, 0xce,
	0xd3, 0x19, 0xf1, 0xcd, 0x6a, 0x52, 0x1c, 0xfb, 0x9f, 0x04, 0xce, 0x73, 0x22, 0x64, 0x5d, 0xbf,
	0x57, 0x4c, 0x73, 0xec, 0x63, 0xc6, 0xb8, 0xc2, 0x8a, 0x72, 0x26, 0xfd, 0x92, 0x08, 0x49, 0x39,
	0xa3, 0x2c, 0xad, 0x25, 0xbb, 0x25, 0x9e, 0xd1, 0x29, 0x56, 0xc4, 0x5f, 0x7e, 0xd8, 0xc2, 0xfd,
	0xcf, 0x1d, 0xb8, 0x7d, 0x64, 0x9b, 0x09, 0x6c, 0x2f, 0xd2, 0x7d, 0x0b, 0xa1, 0x3a, 0x11, 0x44,
	0x9e, 0xf0, 0xd9, 0x54, 0x7a, 0x60, 0xb0, 0x3e, 0x74, 0x46, 0x63, 0x74, 0x49, 0x6f, 0xe8, 0x82,
	0x1b, 0xbd, 0x5b, 0x59, 0xa3, 0x33, 0x98, 0xde, 0xef, 0x4d, 0x08, 0x9b, 0x92, 0xfb, 0x1a, 0x76,
	0x72, 0x41, 0xb9, 0xa0, 0xaa, 0xf2, 0xc0, 0x00, 0x0c, 0xbb, 0xa3, 0xbd, 0x0b, 0x3b, 0x70, 0x41,
	0x34, 0x3e, 0xe2, 0x85, 0xa2, 0x2c, 0x0d, 0x6b, 0x71, 0xd0, 0xf9, 0x13, 0x6c, 0x7c, 0x01, 0x6b,
	0xd7, 0x41, 0xb4, 0x02, 0xb8, 0x2f, 0xe0, 0x76, 0x86, 0xe7, 0x71, 0xc2, 0x19, 0x23, 0x89, 0x49,
	0xc0, 0x5b, 0x1b, 0x80, 0xa1, 0x33, 0xba, 0x83, 0x6c, 0x74, 0x68, 0x19, 0x1d, 0x7a, 0xff, 0x92,
	0xa9, 0xf1, 0xe8, 0x03, 0x9e, 0x15, 0x24, 0xea, 0x66, 0x78, 0x7e, 0xd4, 0x78, 0xdc, 0x37, 0xf0,
	0x86, 0xc6, 0xe4, 0x84, 0x4d, 0x29, 0x4b, 0x63, 0x41, 0x3e, 0x16, 0x44, 0x2a, 0xe9, 0xad, 0xb7,
	0x60, 0xb9, 0x19, 0x9e, 0x87, 0xd6, 0x18, 0xd5, 0x3e, 0xf7, 0x10, 0x6e, 0x69, 0xde, 0x8a, 0x73,
	0xa5, 0x05, 0xc7, 0xc9, 0xf0, 0x7c, 0x05, 0x38, 0x80, 0x8e, 0x05, 0x28, 0x41, 0x89, 0xf4, 0x36,
	0x5a, 0xf8, 0xa1, 0xf1, 0x1b, 0xbd, 0x3b, 0x81, 0x5b, 0xda, 0x5a, 0xc5, 0x93, 0x62, 0x9a, 0x12,
	0xe5, 0x75, 0x8c, 0xff, 0xf0, 0x3f, 0x6e, 0x12, 0x69, 0x64, 0x15, 0x18, 0x4c, 0xe4, 0x88, 0x66,
	0xe1, 0x3e, 0x82, 0xdb, 0x4a, 0xe0, 0xe4, 0x34, 0x16, 0x24, 0xc3, 0x54, 0x4f, 0x9c, 0xb7, 0x39,
	0x00, 0xc3, 0x4e, 0xd4, 0x35, 0xbf, 0xa3, 0xe5, 0xdf, 0x65, 0xb8, 0xcd, 0x1d, 0xc5, 0x39, 0xe7,
	0x33, 0xe9, 0x5d, 0x6d, 0x19, 0x6e, 0x73, 0x51, 0xa1, 0xf6, 0xf5, 0x16, 0x00, 0x3a, 0x67, 0x4e,
	0xe5, 0x1e, 0xc0, 0xae, 0x6d, 0x33, 0xae, 0x1f, 0x8f, 0x19, 0x2b, 0x67, 0x74, 0xab, 0x6e, 0x57,
	0x3f, 0x2d, 0xdd, 0x64, 0x68, 0xab, 0xd1, 0x35, 0xab, 0xae, 0x97, 0x6e, 0x08, 0x6f, 0x66, 0x94,
	0xc5, 0x36, 0xaf, 0x84, 0xb3, 0xa4, 0x10, 0x82, 0xb0, 0xa4, 0x6a, 0x35, 0x48, 0x3b, 0x19, 0x65,
	0xe6, 0x2c, 0x47, 0x8d, 0x71, 0x3f, 0xf8, 0xf6, 0xfd, 0x6b, 0xff, 0x00, 0x3e, 0xb7, 0xdb, 0xe3,
	0x9c, 0xa2, 0x72, 0xb4, 0x4a, 0xbb, 0x5d, 0xd4, 0xfb, 0xcf, 0x34, 0xe3, 0x29, 0xf4, 0xff, 0x91,
	0xb1, 0xff, 0x58, 0xfb, 0x1e, 0xc2, 0x07, 0x6d, 0x7c, 0xc1, 0xab, 0x1f, 0x8b, 0x3e, 0xf8, 0xb9,
	0xe8, 0x83, 0x5f, 0x8b, 0x3e, 0x80, 0x7b, 0x94, 0xdb, 0xc4, 0x72, 0xc1, 0xe7, 0xd5, 0x65, 0xb3,
	0x12, 0xec, 0x9c, 0xa7, 0x84, 0x3a, 0x96, 0x10, 0x4c, 0x36, 0x4d, 0x3e, 0xe3, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xea, 0x32, 0x46, 0xb5, 0x39, 0x05, 0x00, 0x00,
}

func (m *CircuitBreakers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CircuitBreakers) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CircuitBreakers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Thresholds) > 0 {
		for iNdEx := len(m.Thresholds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Thresholds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CircuitBreakers_Thresholds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CircuitBreakers_Thresholds) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CircuitBreakers_Thresholds) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.RetryBudget != nil {
		{
			size, err := m.RetryBudget.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.MaxConnectionPools != nil {
		{
			size, err := m.MaxConnectionPools.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.TrackRemaining {
		i--
		if m.TrackRemaining {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.MaxRetries != nil {
		{
			size, err := m.MaxRetries.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.MaxRequests != nil {
		{
			size, err := m.MaxRequests.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.MaxPendingRequests != nil {
		{
			size, err := m.MaxPendingRequests.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.MaxConnections != nil {
		{
			size, err := m.MaxConnections.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Priority != 0 {
		i = encodeVarintCircuitBreaker(dAtA, i, uint64(m.Priority))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CircuitBreakers_Thresholds_RetryBudget) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CircuitBreakers_Thresholds_RetryBudget) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CircuitBreakers_Thresholds_RetryBudget) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MinRetryConcurrency != nil {
		{
			size, err := m.MinRetryConcurrency.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BudgetPercent != nil {
		{
			size, err := m.BudgetPercent.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCircuitBreaker(dAtA []byte, offset int, v uint64) int {
	offset -= sovCircuitBreaker(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CircuitBreakers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Thresholds) > 0 {
		for _, e := range m.Thresholds {
			l = e.Size()
			n += 1 + l + sovCircuitBreaker(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CircuitBreakers_Thresholds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Priority != 0 {
		n += 1 + sovCircuitBreaker(uint64(m.Priority))
	}
	if m.MaxConnections != nil {
		l = m.MaxConnections.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.MaxPendingRequests != nil {
		l = m.MaxPendingRequests.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.MaxRequests != nil {
		l = m.MaxRequests.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.MaxRetries != nil {
		l = m.MaxRetries.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.TrackRemaining {
		n += 2
	}
	if m.MaxConnectionPools != nil {
		l = m.MaxConnectionPools.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.RetryBudget != nil {
		l = m.RetryBudget.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CircuitBreakers_Thresholds_RetryBudget) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BudgetPercent != nil {
		l = m.BudgetPercent.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.MinRetryConcurrency != nil {
		l = m.MinRetryConcurrency.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCircuitBreaker(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCircuitBreaker(x uint64) (n int) {
	return sovCircuitBreaker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CircuitBreakers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCircuitBreaker
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
			return fmt.Errorf("proto: CircuitBreakers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CircuitBreakers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Thresholds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Thresholds = append(m.Thresholds, &CircuitBreakers_Thresholds{})
			if err := m.Thresholds[len(m.Thresholds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCircuitBreaker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCircuitBreaker
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
func (m *CircuitBreakers_Thresholds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCircuitBreaker
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
			return fmt.Errorf("proto: Thresholds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Thresholds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= v3.RoutingPriority(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxConnections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxConnections == nil {
				m.MaxConnections = &types.UInt32Value{}
			}
			if err := m.MaxConnections.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPendingRequests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxPendingRequests == nil {
				m.MaxPendingRequests = &types.UInt32Value{}
			}
			if err := m.MaxPendingRequests.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRequests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxRequests == nil {
				m.MaxRequests = &types.UInt32Value{}
			}
			if err := m.MaxRequests.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRetries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxRetries == nil {
				m.MaxRetries = &types.UInt32Value{}
			}
			if err := m.MaxRetries.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrackRemaining", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TrackRemaining = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxConnectionPools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxConnectionPools == nil {
				m.MaxConnectionPools = &types.UInt32Value{}
			}
			if err := m.MaxConnectionPools.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetryBudget", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RetryBudget == nil {
				m.RetryBudget = &CircuitBreakers_Thresholds_RetryBudget{}
			}
			if err := m.RetryBudget.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCircuitBreaker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCircuitBreaker
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
func (m *CircuitBreakers_Thresholds_RetryBudget) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCircuitBreaker
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
			return fmt.Errorf("proto: RetryBudget: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RetryBudget: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BudgetPercent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BudgetPercent == nil {
				m.BudgetPercent = &v31.Percent{}
			}
			if err := m.BudgetPercent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinRetryConcurrency", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MinRetryConcurrency == nil {
				m.MinRetryConcurrency = &types.UInt32Value{}
			}
			if err := m.MinRetryConcurrency.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCircuitBreaker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCircuitBreaker
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
func skipCircuitBreaker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCircuitBreaker
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
					return 0, ErrIntOverflowCircuitBreaker
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
					return 0, ErrIntOverflowCircuitBreaker
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
				return 0, ErrInvalidLengthCircuitBreaker
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCircuitBreaker
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCircuitBreaker
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCircuitBreaker        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCircuitBreaker          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCircuitBreaker = fmt.Errorf("proto: unexpected end of group")
)
