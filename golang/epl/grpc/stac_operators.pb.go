// Code generated by protoc-gen-go. DO NOT EDIT.
// source: epl/grpc/stac_operators.proto

package grpc // import "github.com/geo-grpc/api/golang/epl/grpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import protobuf "github.com/geo-grpc/api/golang/epl/protobuf"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetadataOperatorsClient is the client API for MetadataOperators service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetadataOperatorsClient interface {
	Search(ctx context.Context, in *protobuf.MetadataRequest, opts ...grpc.CallOption) (MetadataOperators_SearchClient, error)
	Insert(ctx context.Context, in *protobuf.MetadataResult, opts ...grpc.CallOption) (*protobuf.DBResult, error)
}

type metadataOperatorsClient struct {
	cc *grpc.ClientConn
}

func NewMetadataOperatorsClient(cc *grpc.ClientConn) MetadataOperatorsClient {
	return &metadataOperatorsClient{cc}
}

func (c *metadataOperatorsClient) Search(ctx context.Context, in *protobuf.MetadataRequest, opts ...grpc.CallOption) (MetadataOperators_SearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MetadataOperators_serviceDesc.Streams[0], "/epl.grpc.MetadataOperators/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &metadataOperatorsSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MetadataOperators_SearchClient interface {
	Recv() (*protobuf.MetadataResult, error)
	grpc.ClientStream
}

type metadataOperatorsSearchClient struct {
	grpc.ClientStream
}

func (x *metadataOperatorsSearchClient) Recv() (*protobuf.MetadataResult, error) {
	m := new(protobuf.MetadataResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metadataOperatorsClient) Insert(ctx context.Context, in *protobuf.MetadataResult, opts ...grpc.CallOption) (*protobuf.DBResult, error) {
	out := new(protobuf.DBResult)
	err := c.cc.Invoke(ctx, "/epl.grpc.MetadataOperators/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataOperatorsServer is the server API for MetadataOperators service.
type MetadataOperatorsServer interface {
	Search(*protobuf.MetadataRequest, MetadataOperators_SearchServer) error
	Insert(context.Context, *protobuf.MetadataResult) (*protobuf.DBResult, error)
}

func RegisterMetadataOperatorsServer(s *grpc.Server, srv MetadataOperatorsServer) {
	s.RegisterService(&_MetadataOperators_serviceDesc, srv)
}

func _MetadataOperators_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(protobuf.MetadataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetadataOperatorsServer).Search(m, &metadataOperatorsSearchServer{stream})
}

type MetadataOperators_SearchServer interface {
	Send(*protobuf.MetadataResult) error
	grpc.ServerStream
}

type metadataOperatorsSearchServer struct {
	grpc.ServerStream
}

func (x *metadataOperatorsSearchServer) Send(m *protobuf.MetadataResult) error {
	return x.ServerStream.SendMsg(m)
}

func _MetadataOperators_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.MetadataResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataOperatorsServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/epl.grpc.MetadataOperators/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataOperatorsServer).Insert(ctx, req.(*protobuf.MetadataResult))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetadataOperators_serviceDesc = grpc.ServiceDesc{
	ServiceName: "epl.grpc.MetadataOperators",
	HandlerType: (*MetadataOperatorsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _MetadataOperators_Insert_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _MetadataOperators_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "epl/grpc/stac_operators.proto",
}

func init() {
	proto.RegisterFile("epl/grpc/stac_operators.proto", fileDescriptor_stac_operators_0f29c59608698f04)
}

var fileDescriptor_stac_operators_0f29c59608698f04 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2d, 0xc8, 0xd1,
	0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x2e, 0x49, 0x4c, 0x8e, 0xcf, 0x2f, 0x48, 0x2d, 0x4a, 0x2c,
	0xc9, 0x2f, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x2d, 0xc8, 0xd1, 0x03,
	0x49, 0x4b, 0x89, 0x83, 0x14, 0x82, 0x05, 0x93, 0x4a, 0xd3, 0xc0, 0x8a, 0x21, 0x4a, 0x8c, 0x16,
	0x30, 0x72, 0x09, 0xfa, 0xa6, 0x96, 0x24, 0xa6, 0x24, 0x96, 0x24, 0xfa, 0xc3, 0xb4, 0x0b, 0x79,
	0x72, 0xb1, 0x05, 0xa7, 0x26, 0x16, 0x25, 0x67, 0x08, 0xc9, 0xea, 0x81, 0xcc, 0x80, 0xe9, 0xd4,
	0x83, 0x29, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x91, 0x92, 0xc1, 0x25, 0x5d, 0x5c, 0x9a,
	0x53, 0xa2, 0xc4, 0x60, 0xc0, 0x28, 0xe4, 0xc0, 0xc5, 0xe6, 0x99, 0x57, 0x9c, 0x5a, 0x54, 0x22,
	0x84, 0x57, 0xad, 0x94, 0x18, 0xaa, 0xac, 0x8b, 0x13, 0xcc, 0x0c, 0x27, 0xcd, 0x28, 0xf5, 0xf4,
	0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xf4, 0xd4, 0x7c, 0x5d, 0xb0, 0x8f,
	0x13, 0x0b, 0x32, 0xf5, 0xd3, 0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0x61, 0xa1, 0x90, 0xc4, 0x06,
	0xd6, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x32, 0x70, 0x17, 0xf1, 0x18, 0x01, 0x00, 0x00,
}
