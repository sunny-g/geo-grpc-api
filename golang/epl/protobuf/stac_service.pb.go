// Code generated by protoc-gen-go. DO NOT EDIT.
// source: epl/protobuf/stac_service.proto

package protobuf // import "github.com/geo-grpc/api/golang/epl/protobuf"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// StacServiceClient is the client API for StacService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StacServiceClient interface {
	Search(ctx context.Context, in *StacRequest, opts ...grpc.CallOption) (StacService_SearchClient, error)
	Insert(ctx context.Context, opts ...grpc.CallOption) (StacService_InsertClient, error)
	Update(ctx context.Context, opts ...grpc.CallOption) (StacService_UpdateClient, error)
	SearchOne(ctx context.Context, in *StacRequest, opts ...grpc.CallOption) (*StacItem, error)
	InsertOne(ctx context.Context, in *StacItem, opts ...grpc.CallOption) (*StacUpsertResponse, error)
	UpdateOne(ctx context.Context, in *StacItem, opts ...grpc.CallOption) (*StacUpsertResponse, error)
}

type stacServiceClient struct {
	cc *grpc.ClientConn
}

func NewStacServiceClient(cc *grpc.ClientConn) StacServiceClient {
	return &stacServiceClient{cc}
}

func (c *stacServiceClient) Search(ctx context.Context, in *StacRequest, opts ...grpc.CallOption) (StacService_SearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StacService_serviceDesc.Streams[0], "/epl.protobuf.StacService/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &stacServiceSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StacService_SearchClient interface {
	Recv() (*StacItem, error)
	grpc.ClientStream
}

type stacServiceSearchClient struct {
	grpc.ClientStream
}

func (x *stacServiceSearchClient) Recv() (*StacItem, error) {
	m := new(StacItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stacServiceClient) Insert(ctx context.Context, opts ...grpc.CallOption) (StacService_InsertClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StacService_serviceDesc.Streams[1], "/epl.protobuf.StacService/Insert", opts...)
	if err != nil {
		return nil, err
	}
	x := &stacServiceInsertClient{stream}
	return x, nil
}

type StacService_InsertClient interface {
	Send(*StacItem) error
	Recv() (*StacUpsertResponse, error)
	grpc.ClientStream
}

type stacServiceInsertClient struct {
	grpc.ClientStream
}

func (x *stacServiceInsertClient) Send(m *StacItem) error {
	return x.ClientStream.SendMsg(m)
}

func (x *stacServiceInsertClient) Recv() (*StacUpsertResponse, error) {
	m := new(StacUpsertResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stacServiceClient) Update(ctx context.Context, opts ...grpc.CallOption) (StacService_UpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StacService_serviceDesc.Streams[2], "/epl.protobuf.StacService/Update", opts...)
	if err != nil {
		return nil, err
	}
	x := &stacServiceUpdateClient{stream}
	return x, nil
}

type StacService_UpdateClient interface {
	Send(*StacItem) error
	Recv() (*StacUpsertResponse, error)
	grpc.ClientStream
}

type stacServiceUpdateClient struct {
	grpc.ClientStream
}

func (x *stacServiceUpdateClient) Send(m *StacItem) error {
	return x.ClientStream.SendMsg(m)
}

func (x *stacServiceUpdateClient) Recv() (*StacUpsertResponse, error) {
	m := new(StacUpsertResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stacServiceClient) SearchOne(ctx context.Context, in *StacRequest, opts ...grpc.CallOption) (*StacItem, error) {
	out := new(StacItem)
	err := c.cc.Invoke(ctx, "/epl.protobuf.StacService/SearchOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stacServiceClient) InsertOne(ctx context.Context, in *StacItem, opts ...grpc.CallOption) (*StacUpsertResponse, error) {
	out := new(StacUpsertResponse)
	err := c.cc.Invoke(ctx, "/epl.protobuf.StacService/InsertOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stacServiceClient) UpdateOne(ctx context.Context, in *StacItem, opts ...grpc.CallOption) (*StacUpsertResponse, error) {
	out := new(StacUpsertResponse)
	err := c.cc.Invoke(ctx, "/epl.protobuf.StacService/UpdateOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StacServiceServer is the server API for StacService service.
type StacServiceServer interface {
	Search(*StacRequest, StacService_SearchServer) error
	Insert(StacService_InsertServer) error
	Update(StacService_UpdateServer) error
	SearchOne(context.Context, *StacRequest) (*StacItem, error)
	InsertOne(context.Context, *StacItem) (*StacUpsertResponse, error)
	UpdateOne(context.Context, *StacItem) (*StacUpsertResponse, error)
}

func RegisterStacServiceServer(s *grpc.Server, srv StacServiceServer) {
	s.RegisterService(&_StacService_serviceDesc, srv)
}

func _StacService_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StacRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StacServiceServer).Search(m, &stacServiceSearchServer{stream})
}

type StacService_SearchServer interface {
	Send(*StacItem) error
	grpc.ServerStream
}

type stacServiceSearchServer struct {
	grpc.ServerStream
}

func (x *stacServiceSearchServer) Send(m *StacItem) error {
	return x.ServerStream.SendMsg(m)
}

func _StacService_Insert_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StacServiceServer).Insert(&stacServiceInsertServer{stream})
}

type StacService_InsertServer interface {
	Send(*StacUpsertResponse) error
	Recv() (*StacItem, error)
	grpc.ServerStream
}

type stacServiceInsertServer struct {
	grpc.ServerStream
}

func (x *stacServiceInsertServer) Send(m *StacUpsertResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *stacServiceInsertServer) Recv() (*StacItem, error) {
	m := new(StacItem)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StacService_Update_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StacServiceServer).Update(&stacServiceUpdateServer{stream})
}

type StacService_UpdateServer interface {
	Send(*StacUpsertResponse) error
	Recv() (*StacItem, error)
	grpc.ServerStream
}

type stacServiceUpdateServer struct {
	grpc.ServerStream
}

func (x *stacServiceUpdateServer) Send(m *StacUpsertResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *stacServiceUpdateServer) Recv() (*StacItem, error) {
	m := new(StacItem)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StacService_SearchOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StacRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacServiceServer).SearchOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/epl.protobuf.StacService/SearchOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacServiceServer).SearchOne(ctx, req.(*StacRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StacService_InsertOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StacItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacServiceServer).InsertOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/epl.protobuf.StacService/InsertOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacServiceServer).InsertOne(ctx, req.(*StacItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _StacService_UpdateOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StacItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacServiceServer).UpdateOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/epl.protobuf.StacService/UpdateOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacServiceServer).UpdateOne(ctx, req.(*StacItem))
	}
	return interceptor(ctx, in, info, handler)
}

var _StacService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "epl.protobuf.StacService",
	HandlerType: (*StacServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchOne",
			Handler:    _StacService_SearchOne_Handler,
		},
		{
			MethodName: "InsertOne",
			Handler:    _StacService_InsertOne_Handler,
		},
		{
			MethodName: "UpdateOne",
			Handler:    _StacService_UpdateOne_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _StacService_Search_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Insert",
			Handler:       _StacService_Insert_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Update",
			Handler:       _StacService_Update_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "epl/protobuf/stac_service.proto",
}

func init() {
	proto.RegisterFile("epl/protobuf/stac_service.proto", fileDescriptor_stac_service_42ac76d8fc0e3bd4)
}

var fileDescriptor_stac_service_42ac76d8fc0e3bd4 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2d, 0xc8, 0xd1,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x2e, 0x49, 0x4c, 0x8e, 0x2f, 0x4e,
	0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0x8b, 0x0a, 0xf1, 0xa4, 0x16, 0xe4, 0xe8, 0xc1, 0x14,
	0x48, 0x89, 0x63, 0x28, 0x87, 0xc8, 0x19, 0x6d, 0x63, 0xe6, 0xe2, 0x0e, 0x2e, 0x49, 0x4c, 0x0e,
	0x86, 0x68, 0x16, 0xb2, 0xe7, 0x62, 0x0b, 0x4e, 0x4d, 0x2c, 0x4a, 0xce, 0x10, 0x92, 0xd4, 0x43,
	0x36, 0x41, 0x0f, 0xa4, 0x28, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x4a, 0x0c, 0x53, 0xca,
	0xb3, 0x24, 0x35, 0x57, 0x89, 0xc1, 0x80, 0x51, 0xc8, 0x83, 0x8b, 0xcd, 0x33, 0xaf, 0x38, 0xb5,
	0xa8, 0x44, 0x08, 0x87, 0x2a, 0x29, 0x05, 0x4c, 0xf1, 0xd0, 0x02, 0x90, 0x8e, 0xa0, 0xd4, 0xe2,
	0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x25, 0x06, 0x0d, 0x46, 0x88, 0x49, 0xa1, 0x05, 0x29, 0x89, 0x25,
	0xa9, 0x14, 0x9b, 0xe4, 0xc0, 0xc5, 0x09, 0xf1, 0x94, 0x7f, 0x5e, 0x2a, 0x59, 0xfe, 0x12, 0x72,
	0xe7, 0xe2, 0x84, 0xf8, 0x0a, 0x64, 0x02, 0x05, 0xce, 0x01, 0x19, 0x04, 0xf1, 0x14, 0x85, 0x06,
	0x39, 0x45, 0x72, 0x09, 0x24, 0xe7, 0xe7, 0xa2, 0x28, 0x74, 0x12, 0x40, 0x8a, 0xc9, 0x00, 0x90,
	0x60, 0x00, 0x63, 0x94, 0x76, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e,
	0x7a, 0x6a, 0xbe, 0x6e, 0x7a, 0x51, 0x41, 0xb2, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x7a, 0x7e, 0x4e,
	0x62, 0x5e, 0xba, 0x3e, 0x72, 0xc2, 0x58, 0xc4, 0xc4, 0x1c, 0x1c, 0x12, 0x9c, 0xc4, 0x06, 0xe6,
	0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xe5, 0x24, 0xb9, 0x64, 0x02, 0x00, 0x00,
}
