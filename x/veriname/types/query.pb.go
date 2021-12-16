// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: veriname/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
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

func init() { proto.RegisterFile("veriname/query.proto", fileDescriptor_0472b6a522181f5b) }

var fileDescriptor_0472b6a522181f5b = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x4b, 0x2d, 0xca,
	0xcc, 0x4b, 0xcc, 0x4d, 0xd5, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x48, 0xc9, 0xcd, 0xcf, 0xc9, 0xcc, 0xd6, 0x83, 0x49, 0xc2, 0x19, 0x52, 0x32, 0xe9,
	0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89,
	0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x7d, 0x52, 0x5a, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xfa,
	0x49, 0x89, 0xc5, 0x50, 0x03, 0xf5, 0xcb, 0x0c, 0x93, 0x52, 0x4b, 0x12, 0x0d, 0xf5, 0x0b, 0x12,
	0xd3, 0x33, 0xf3, 0xc0, 0x8a, 0x21, 0x6a, 0x8d, 0xd8, 0xb9, 0x58, 0x03, 0x41, 0x2a, 0x9c, 0x5c,
	0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5,
	0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x3b, 0x3d, 0xb3, 0x24, 0xa3,
	0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xe2, 0x22, 0x7d, 0xb8, 0x73, 0x2b, 0x10, 0xcc, 0x92,
	0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xb1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c,
	0x52, 0x73, 0x67, 0xd2, 0x00, 0x00, 0x00,
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
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dmolik.veriname.veriname.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "veriname/query.proto",
}
