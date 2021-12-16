// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: veriname/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
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

type QueryGetIdentRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryGetIdentRequest) Reset()         { *m = QueryGetIdentRequest{} }
func (m *QueryGetIdentRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetIdentRequest) ProtoMessage()    {}
func (*QueryGetIdentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0472b6a522181f5b, []int{0}
}
func (m *QueryGetIdentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetIdentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetIdentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetIdentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetIdentRequest.Merge(m, src)
}
func (m *QueryGetIdentRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetIdentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetIdentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetIdentRequest proto.InternalMessageInfo

func (m *QueryGetIdentRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type QueryGetIdentResponse struct {
	Ident Ident `protobuf:"bytes,1,opt,name=ident,proto3" json:"ident"`
}

func (m *QueryGetIdentResponse) Reset()         { *m = QueryGetIdentResponse{} }
func (m *QueryGetIdentResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetIdentResponse) ProtoMessage()    {}
func (*QueryGetIdentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0472b6a522181f5b, []int{1}
}
func (m *QueryGetIdentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetIdentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetIdentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetIdentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetIdentResponse.Merge(m, src)
}
func (m *QueryGetIdentResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetIdentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetIdentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetIdentResponse proto.InternalMessageInfo

func (m *QueryGetIdentResponse) GetIdent() Ident {
	if m != nil {
		return m.Ident
	}
	return Ident{}
}

type QueryAllIdentRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllIdentRequest) Reset()         { *m = QueryAllIdentRequest{} }
func (m *QueryAllIdentRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllIdentRequest) ProtoMessage()    {}
func (*QueryAllIdentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0472b6a522181f5b, []int{2}
}
func (m *QueryAllIdentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllIdentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllIdentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllIdentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllIdentRequest.Merge(m, src)
}
func (m *QueryAllIdentRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllIdentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllIdentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllIdentRequest proto.InternalMessageInfo

func (m *QueryAllIdentRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllIdentResponse struct {
	Ident      []Ident             `protobuf:"bytes,1,rep,name=ident,proto3" json:"ident"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllIdentResponse) Reset()         { *m = QueryAllIdentResponse{} }
func (m *QueryAllIdentResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllIdentResponse) ProtoMessage()    {}
func (*QueryAllIdentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0472b6a522181f5b, []int{3}
}
func (m *QueryAllIdentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllIdentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllIdentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllIdentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllIdentResponse.Merge(m, src)
}
func (m *QueryAllIdentResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllIdentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllIdentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllIdentResponse proto.InternalMessageInfo

func (m *QueryAllIdentResponse) GetIdent() []Ident {
	if m != nil {
		return m.Ident
	}
	return nil
}

func (m *QueryAllIdentResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetIdentRequest)(nil), "dmolik.veriname.veriname.QueryGetIdentRequest")
	proto.RegisterType((*QueryGetIdentResponse)(nil), "dmolik.veriname.veriname.QueryGetIdentResponse")
	proto.RegisterType((*QueryAllIdentRequest)(nil), "dmolik.veriname.veriname.QueryAllIdentRequest")
	proto.RegisterType((*QueryAllIdentResponse)(nil), "dmolik.veriname.veriname.QueryAllIdentResponse")
}

func init() { proto.RegisterFile("veriname/query.proto", fileDescriptor_0472b6a522181f5b) }

var fileDescriptor_0472b6a522181f5b = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x8f, 0x12, 0x31,
	0x18, 0xc6, 0xa7, 0xe8, 0x18, 0xad, 0xb7, 0x06, 0x13, 0x42, 0xcc, 0xa0, 0x73, 0x10, 0xff, 0xa5,
	0x0d, 0x78, 0xf4, 0x04, 0x89, 0x12, 0x6f, 0x3a, 0xf1, 0xe4, 0xc1, 0xa4, 0x03, 0xcd, 0xd8, 0x38,
	0xd3, 0x0e, 0xb4, 0x10, 0x88, 0xf1, 0xe2, 0x27, 0x30, 0xd1, 0x8b, 0x07, 0xbf, 0xc8, 0x7e, 0x02,
	0x8e, 0x24, 0x7b, 0xd9, 0xd3, 0x66, 0x03, 0xfb, 0x41, 0x36, 0xd3, 0x0e, 0x0c, 0x7f, 0x96, 0xcc,
	0xee, 0xed, 0xa5, 0x3c, 0xcf, 0xfb, 0xfc, 0xfa, 0xbe, 0x1d, 0x58, 0x9d, 0xb0, 0x11, 0x17, 0x34,
	0x61, 0x64, 0x38, 0x66, 0xa3, 0x19, 0x4e, 0x47, 0x52, 0x4b, 0x54, 0x1b, 0x24, 0x32, 0xe6, 0xdf,
	0xf1, 0xfa, 0xcf, 0x4d, 0x51, 0x7f, 0x1c, 0x49, 0x19, 0xc5, 0x8c, 0xd0, 0x94, 0x13, 0x2a, 0x84,
	0xd4, 0x54, 0x73, 0x29, 0x94, 0xf5, 0xd5, 0x5f, 0xf6, 0xa5, 0x4a, 0xa4, 0x22, 0x21, 0x55, 0x79,
	0x43, 0x32, 0x69, 0x85, 0x4c, 0xd3, 0x16, 0x49, 0x69, 0xc4, 0x85, 0x11, 0xe7, 0xda, 0x22, 0x99,
	0x0f, 0x98, 0xd0, 0xeb, 0xd3, 0x48, 0x46, 0xd2, 0x94, 0x24, 0xab, 0xec, 0xa9, 0xff, 0x1a, 0x56,
	0x3f, 0x65, 0xdd, 0x7a, 0x4c, 0x7f, 0xc8, 0xc4, 0x01, 0x1b, 0x8e, 0x99, 0xd2, 0xa8, 0x0a, 0x5d,
	0x2e, 0x06, 0x6c, 0x5a, 0x03, 0x4f, 0xc0, 0xf3, 0x07, 0x81, 0xfd, 0xe1, 0x7f, 0x86, 0x8f, 0xf6,
	0xd4, 0x2a, 0x95, 0x42, 0x31, 0xf4, 0x16, 0xba, 0x26, 0xcb, 0xc8, 0x1f, 0xb6, 0x1b, 0xf8, 0xd8,
	0x35, 0xb1, 0xf1, 0x75, 0xef, 0xce, 0xcf, 0x1b, 0x4e, 0x60, 0x3d, 0xfe, 0xd7, 0x9c, 0xa1, 0x13,
	0xc7, 0x3b, 0x0c, 0xef, 0x21, 0x2c, 0xee, 0x96, 0x77, 0x7e, 0x86, 0xed, 0x20, 0x70, 0x36, 0x08,
	0x6c, 0x27, 0x9b, 0x0f, 0x02, 0x7f, 0xa4, 0x11, 0xcb, 0xbd, 0xc1, 0x96, 0xd3, 0xff, 0x0f, 0x72,
	0xec, 0x22, 0xe0, 0x10, 0xfb, 0xce, 0x6d, 0xb1, 0x51, 0x6f, 0x07, 0xaf, 0x62, 0xf0, 0x9a, 0xa5,
	0x78, 0x36, 0x79, 0x9b, 0xaf, 0x7d, 0x52, 0x81, 0xae, 0xe1, 0x43, 0xff, 0x00, 0x74, 0x4d, 0x12,
	0xc2, 0xc7, 0x51, 0xae, 0xdb, 0x57, 0x9d, 0xdc, 0x58, 0x6f, 0x01, 0x7c, 0xf2, 0xeb, 0xf4, 0xf2,
	0x4f, 0xe5, 0x05, 0x6a, 0x12, 0x6b, 0x24, 0x9b, 0x47, 0xb3, 0xfb, 0x7a, 0xc8, 0x0f, 0xb3, 0xfa,
	0x9f, 0xe8, 0x2f, 0x80, 0xf7, 0x4d, 0x8b, 0x4e, 0x1c, 0x97, 0xe2, 0xed, 0xad, 0xb2, 0x14, 0x6f,
	0x7f, 0x33, 0x7e, 0xd3, 0xe0, 0x3d, 0x45, 0x8d, 0x12, 0xbc, 0xee, 0xbb, 0xf9, 0xd2, 0x03, 0x8b,
	0xa5, 0x07, 0x2e, 0x96, 0x1e, 0xf8, 0xbd, 0xf2, 0x9c, 0xc5, 0xca, 0x73, 0xce, 0x56, 0x9e, 0xf3,
	0xe5, 0x55, 0xc4, 0xf5, 0xb7, 0x71, 0x88, 0xfb, 0x32, 0x39, 0x68, 0x32, 0x2d, 0x4a, 0x3d, 0x4b,
	0x99, 0x0a, 0xef, 0x99, 0xcf, 0xe1, 0xcd, 0x55, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x78, 0x15,
	0x17, 0xb6, 0x03, 0x00, 0x00,
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
	// Queries a ident by index.
	Ident(ctx context.Context, in *QueryGetIdentRequest, opts ...grpc.CallOption) (*QueryGetIdentResponse, error)
	// Queries a list of ident items.
	IdentAll(ctx context.Context, in *QueryAllIdentRequest, opts ...grpc.CallOption) (*QueryAllIdentResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Ident(ctx context.Context, in *QueryGetIdentRequest, opts ...grpc.CallOption) (*QueryGetIdentResponse, error) {
	out := new(QueryGetIdentResponse)
	err := c.cc.Invoke(ctx, "/dmolik.veriname.veriname.Query/Ident", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IdentAll(ctx context.Context, in *QueryAllIdentRequest, opts ...grpc.CallOption) (*QueryAllIdentResponse, error) {
	out := new(QueryAllIdentResponse)
	err := c.cc.Invoke(ctx, "/dmolik.veriname.veriname.Query/IdentAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a ident by index.
	Ident(context.Context, *QueryGetIdentRequest) (*QueryGetIdentResponse, error)
	// Queries a list of ident items.
	IdentAll(context.Context, *QueryAllIdentRequest) (*QueryAllIdentResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Ident(ctx context.Context, req *QueryGetIdentRequest) (*QueryGetIdentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ident not implemented")
}
func (*UnimplementedQueryServer) IdentAll(ctx context.Context, req *QueryAllIdentRequest) (*QueryAllIdentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IdentAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Ident_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetIdentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Ident(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dmolik.veriname.veriname.Query/Ident",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Ident(ctx, req.(*QueryGetIdentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IdentAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllIdentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IdentAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dmolik.veriname.veriname.Query/IdentAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IdentAll(ctx, req.(*QueryAllIdentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dmolik.veriname.veriname.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ident",
			Handler:    _Query_Ident_Handler,
		},
		{
			MethodName: "IdentAll",
			Handler:    _Query_IdentAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "veriname/query.proto",
}

func (m *QueryGetIdentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetIdentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetIdentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetIdentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetIdentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetIdentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Ident.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllIdentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllIdentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllIdentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllIdentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllIdentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllIdentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Ident) > 0 {
		for iNdEx := len(m.Ident) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Ident[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *QueryGetIdentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetIdentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Ident.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllIdentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllIdentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ident) > 0 {
		for _, e := range m.Ident {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetIdentRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetIdentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetIdentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
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
			m.Index = string(dAtA[iNdEx:postIndex])
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
func (m *QueryGetIdentResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetIdentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetIdentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ident", wireType)
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
			if err := m.Ident.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllIdentRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllIdentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllIdentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllIdentResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllIdentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllIdentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ident", wireType)
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
			m.Ident = append(m.Ident, Ident{})
			if err := m.Ident[len(m.Ident)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
