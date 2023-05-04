// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mint/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0718dda172d2cb4, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0718dda172d2cb4, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryMinterRequest is the request type for the
// Query/Minter RPC method.
type QueryMinterRequest struct {
}

func (m *QueryMinterRequest) Reset()         { *m = QueryMinterRequest{} }
func (m *QueryMinterRequest) String() string { return proto.CompactTextString(m) }
func (*QueryMinterRequest) ProtoMessage()    {}
func (*QueryMinterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0718dda172d2cb4, []int{2}
}
func (m *QueryMinterRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMinterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMinterRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMinterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMinterRequest.Merge(m, src)
}
func (m *QueryMinterRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryMinterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMinterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMinterRequest proto.InternalMessageInfo

// QueryMinterResponse is the response type for the
// Query/Minter RPC method.
type QueryMinterResponse struct {
	StartDate           string `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty" yaml:"start_date"`
	EndDate             string `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty" yaml:"end_date"`
	Denom               string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
	TotalMintAmount     uint64 `protobuf:"varint,4,opt,name=total_mint_amount,json=totalMintAmount,proto3" json:"total_mint_amount,omitempty" yaml:"total_mint_amount"`
	RemainingMintAmount uint64 `protobuf:"varint,5,opt,name=remaining_mint_amount,json=remainingMintAmount,proto3" json:"remaining_mint_amount,omitempty" yaml:"remaining_mint_amount"`
	LastMintAmount      uint64 `protobuf:"varint,6,opt,name=last_mint_amount,json=lastMintAmount,proto3" json:"last_mint_amount,omitempty" yaml:"last_mint_amount"`
	LastMintDate        string `protobuf:"bytes,7,opt,name=last_mint_date,json=lastMintDate,proto3" json:"last_mint_date,omitempty" yaml:"last_mint_date"`
	LastMintHeight      uint64 `protobuf:"varint,8,opt,name=last_mint_height,json=lastMintHeight,proto3" json:"last_mint_height,omitempty" yaml:"last_mint_height"`
}

func (m *QueryMinterResponse) Reset()         { *m = QueryMinterResponse{} }
func (m *QueryMinterResponse) String() string { return proto.CompactTextString(m) }
func (*QueryMinterResponse) ProtoMessage()    {}
func (*QueryMinterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0718dda172d2cb4, []int{3}
}
func (m *QueryMinterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMinterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMinterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMinterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMinterResponse.Merge(m, src)
}
func (m *QueryMinterResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryMinterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMinterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMinterResponse proto.InternalMessageInfo

func (m *QueryMinterResponse) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *QueryMinterResponse) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *QueryMinterResponse) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *QueryMinterResponse) GetTotalMintAmount() uint64 {
	if m != nil {
		return m.TotalMintAmount
	}
	return 0
}

func (m *QueryMinterResponse) GetRemainingMintAmount() uint64 {
	if m != nil {
		return m.RemainingMintAmount
	}
	return 0
}

func (m *QueryMinterResponse) GetLastMintAmount() uint64 {
	if m != nil {
		return m.LastMintAmount
	}
	return 0
}

func (m *QueryMinterResponse) GetLastMintDate() string {
	if m != nil {
		return m.LastMintDate
	}
	return ""
}

func (m *QueryMinterResponse) GetLastMintHeight() uint64 {
	if m != nil {
		return m.LastMintHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "seiprotocol.seichain.mint.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "seiprotocol.seichain.mint.QueryParamsResponse")
	proto.RegisterType((*QueryMinterRequest)(nil), "seiprotocol.seichain.mint.QueryMinterRequest")
	proto.RegisterType((*QueryMinterResponse)(nil), "seiprotocol.seichain.mint.QueryMinterResponse")
}

func init() { proto.RegisterFile("mint/v1beta1/query.proto", fileDescriptor_b0718dda172d2cb4) }

var fileDescriptor_b0718dda172d2cb4 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x18, 0x8c, 0x43, 0x92, 0xb6, 0x4b, 0xd5, 0x1f, 0xa7, 0x51, 0xdd, 0x10, 0xec, 0xb0, 0x12, 0xa8,
	0x97, 0xda, 0x6a, 0xe1, 0xc4, 0xa5, 0x22, 0x02, 0x29, 0x17, 0x24, 0xb0, 0x10, 0x07, 0x2e, 0xd1,
	0x26, 0x59, 0x39, 0x2b, 0xc5, 0xbb, 0xae, 0x77, 0x83, 0xc8, 0x95, 0x07, 0x40, 0x08, 0x9e, 0x82,
	0x37, 0xe9, 0xb1, 0x12, 0x17, 0x4e, 0x16, 0x4a, 0x78, 0x02, 0x3f, 0x01, 0xf2, 0xb7, 0x71, 0x1b,
	0xb7, 0x94, 0xf4, 0xe6, 0x9d, 0x6f, 0xbe, 0x99, 0xf1, 0x6a, 0x16, 0x59, 0x21, 0xe3, 0xca, 0xfb,
	0x78, 0xdc, 0xa7, 0x8a, 0x1c, 0x7b, 0x67, 0x13, 0x1a, 0x4f, 0xdd, 0x28, 0x16, 0x4a, 0x98, 0x07,
	0x92, 0x32, 0xf8, 0x1a, 0x88, 0xb1, 0x2b, 0x29, 0x1b, 0x8c, 0x08, 0xe3, 0x6e, 0x46, 0x6f, 0xee,
	0x05, 0x22, 0x10, 0x30, 0xf3, 0xb2, 0x2f, 0xbd, 0xd0, 0x6c, 0x05, 0x42, 0x04, 0x63, 0xea, 0x91,
	0x88, 0x79, 0x84, 0x73, 0xa1, 0x88, 0x62, 0x82, 0xcb, 0xc5, 0x74, 0xbf, 0x60, 0x94, 0x1d, 0xf4,
	0x00, 0xef, 0x21, 0xf3, 0x6d, 0x66, 0xfb, 0x86, 0xc4, 0x24, 0x94, 0x3e, 0x3d, 0x9b, 0x50, 0xa9,
	0xf0, 0x7b, 0x54, 0x2f, 0xa0, 0x32, 0x12, 0x5c, 0x52, 0xf3, 0x14, 0xd5, 0x22, 0x40, 0x2c, 0xa3,
	0x6d, 0x1c, 0xde, 0x3f, 0x79, 0xe4, 0xde, 0x9a, 0xd2, 0xd5, 0xab, 0x9d, 0xca, 0x79, 0xe2, 0x94,
	0xfc, 0xc5, 0xda, 0xa5, 0xdb, 0x6b, 0xc6, 0x15, 0x8d, 0x73, 0xb7, 0x6f, 0x95, 0x85, 0x5d, 0x0e,
	0x2f, 0xec, 0x9e, 0x21, 0x24, 0x15, 0x89, 0x55, 0x6f, 0x48, 0x14, 0x05, 0xcb, 0x8d, 0x4e, 0x23,
	0x4d, 0x9c, 0xdd, 0x29, 0x09, 0xc7, 0xcf, 0xf1, 0xd5, 0x0c, 0xfb, 0x1b, 0x70, 0x78, 0x49, 0x14,
	0x35, 0x5d, 0xb4, 0x4e, 0xf9, 0x50, 0xef, 0x94, 0x61, 0xa7, 0x9e, 0x26, 0xce, 0xb6, 0xde, 0xc9,
	0x27, 0xd8, 0x5f, 0xa3, 0x7c, 0x08, 0xfc, 0x27, 0xa8, 0x3a, 0xa4, 0x5c, 0x84, 0xd6, 0x3d, 0x20,
	0xef, 0xa4, 0x89, 0xb3, 0xa9, 0xc9, 0x00, 0x63, 0x5f, 0x8f, 0xcd, 0x2e, 0xda, 0x55, 0x42, 0x91,
	0x71, 0x2f, 0xfb, 0xbd, 0x1e, 0x09, 0xc5, 0x84, 0x2b, 0xab, 0xd2, 0x36, 0x0e, 0x2b, 0x9d, 0x56,
	0x9a, 0x38, 0x96, 0xde, 0xb9, 0x41, 0xc1, 0xfe, 0x36, 0x60, 0xd9, 0xbf, 0xbd, 0x00, 0xc4, 0x7c,
	0x87, 0x1a, 0x31, 0x0d, 0x09, 0xe3, 0x8c, 0x07, 0x05, 0xb5, 0x2a, 0xa8, 0xb5, 0xd3, 0xc4, 0x69,
	0x69, 0xb5, 0x7f, 0xd2, 0xb0, 0x5f, 0xbf, 0xc4, 0x97, 0x54, 0x5f, 0xa1, 0x9d, 0x31, 0x91, 0xaa,
	0x20, 0x58, 0x03, 0xc1, 0x07, 0x69, 0xe2, 0xec, 0x6b, 0xc1, 0xeb, 0x0c, 0xec, 0x6f, 0x65, 0xd0,
	0x92, 0xcc, 0x29, 0xda, 0xba, 0x22, 0xc1, 0x25, 0xae, 0xc1, 0xbd, 0x1c, 0xa4, 0x89, 0xd3, 0xb8,
	0x2e, 0xa2, 0xaf, 0x72, 0x33, 0x97, 0x80, 0xfb, 0x2c, 0xe4, 0x18, 0x51, 0x16, 0x8c, 0x94, 0xb5,
	0x7e, 0x7b, 0x0e, 0xcd, 0x58, 0xca, 0xd1, 0x05, 0xe0, 0xe4, 0x47, 0x19, 0x55, 0xa1, 0x14, 0xe6,
	0x17, 0x03, 0xd5, 0x74, 0x9b, 0xcc, 0xa3, 0xff, 0x14, 0xee, 0x66, 0x8d, 0x9b, 0xee, 0x5d, 0xe9,
	0xba, 0x70, 0xf8, 0xf1, 0xe7, 0x9f, 0x7f, 0xbe, 0x97, 0x1d, 0xf3, 0xa1, 0x97, 0x73, 0xbd, 0xc2,
	0xbb, 0xd1, 0x2d, 0x86, 0x40, 0xba, 0xaa, 0xab, 0x03, 0x15, 0x9a, 0xbe, 0x3a, 0x50, 0xf1, 0x05,
	0xac, 0x0c, 0x14, 0x02, 0xbd, 0xd3, 0x3d, 0x9f, 0xd9, 0xc6, 0xc5, 0xcc, 0x36, 0x7e, 0xcf, 0x6c,
	0xe3, 0xeb, 0xdc, 0x2e, 0x5d, 0xcc, 0xed, 0xd2, 0xaf, 0xb9, 0x5d, 0xfa, 0xe0, 0x06, 0x4c, 0x8d,
	0x26, 0x7d, 0x77, 0x20, 0xc2, 0x4c, 0xe2, 0x28, 0xf7, 0x86, 0x83, 0x16, 0xfc, 0xa4, 0x25, 0xd5,
	0x34, 0xa2, 0xb2, 0x5f, 0x03, 0xc2, 0xd3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x6e, 0x23,
	0xe2, 0x99, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params returns the total set of minting parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// EpochProvisions current minting epoch provisions value.
	Minter(ctx context.Context, in *QueryMinterRequest, opts ...grpc.CallOption) (*QueryMinterResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/seiprotocol.seichain.mint.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Minter(ctx context.Context, in *QueryMinterRequest, opts ...grpc.CallOption) (*QueryMinterResponse, error) {
	out := new(QueryMinterResponse)
	err := c.cc.Invoke(ctx, "/seiprotocol.seichain.mint.Query/Minter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params returns the total set of minting parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// EpochProvisions current minting epoch provisions value.
	Minter(context.Context, *QueryMinterRequest) (*QueryMinterResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Minter(ctx context.Context, req *QueryMinterRequest) (*QueryMinterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Minter not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seiprotocol.seichain.mint.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Minter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMinterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Minter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seiprotocol.seichain.mint.Query/Minter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Minter(ctx, req.(*QueryMinterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "seiprotocol.seichain.mint.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Minter",
			Handler:    _Query_Minter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mint/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryMinterRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMinterRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMinterRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryMinterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMinterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMinterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastMintHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.LastMintHeight))
		i--
		dAtA[i] = 0x40
	}
	if len(m.LastMintDate) > 0 {
		i -= len(m.LastMintDate)
		copy(dAtA[i:], m.LastMintDate)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.LastMintDate)))
		i--
		dAtA[i] = 0x3a
	}
	if m.LastMintAmount != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.LastMintAmount))
		i--
		dAtA[i] = 0x30
	}
	if m.RemainingMintAmount != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.RemainingMintAmount))
		i--
		dAtA[i] = 0x28
	}
	if m.TotalMintAmount != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.TotalMintAmount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EndDate) > 0 {
		i -= len(m.EndDate)
		copy(dAtA[i:], m.EndDate)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.EndDate)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.StartDate) > 0 {
		i -= len(m.StartDate)
		copy(dAtA[i:], m.StartDate)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.StartDate)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryMinterRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryMinterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StartDate)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.EndDate)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.TotalMintAmount != 0 {
		n += 1 + sovQuery(uint64(m.TotalMintAmount))
	}
	if m.RemainingMintAmount != 0 {
		n += 1 + sovQuery(uint64(m.RemainingMintAmount))
	}
	if m.LastMintAmount != 0 {
		n += 1 + sovQuery(uint64(m.LastMintAmount))
	}
	l = len(m.LastMintDate)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.LastMintHeight != 0 {
		n += 1 + sovQuery(uint64(m.LastMintHeight))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryMinterRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryMinterRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMinterRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryMinterResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryMinterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMinterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalMintAmount", wireType)
			}
			m.TotalMintAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalMintAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingMintAmount", wireType)
			}
			m.RemainingMintAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RemainingMintAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastMintAmount", wireType)
			}
			m.LastMintAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastMintAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastMintDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastMintDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastMintHeight", wireType)
			}
			m.LastMintHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastMintHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
