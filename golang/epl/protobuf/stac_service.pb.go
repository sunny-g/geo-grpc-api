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
	//
	// using a search request, stream all the results that match the search filter
	Search(ctx context.Context, in *StacRequest, opts ...grpc.CallOption) (StacService_SearchClient, error)
	//
	// insert a stream of items into the STAC service
	Insert(ctx context.Context, opts ...grpc.CallOption) (StacService_InsertClient, error)
	//
	// update a stream of items in the STAC service
	Update(ctx context.Context, opts ...grpc.CallOption) (StacService_UpdateClient, error)
	//
	// count all the items in the Stac service according to the StacRequest filter
	Count(ctx context.Context, in *StacRequest, opts ...grpc.CallOption) (*StacDbResponse, error)
	//
	// delete an item from the STAC service
	DeleteOne(ctx context.Context, in *StacItem, opts ...grpc.CallOption) (*StacDbResponse, error)
	//
	// using a search request get the first item that matches the request
	SearchOne(ctx context.Context, in *StacRequest, opts ...grpc.CallOption) (*StacItem, error)
	//
	// Insert one item into the STAC service
	InsertOne(ctx context.Context, in *StacItem, opts ...grpc.CallOption) (*StacDbResponse, error)
	//
	// Update one item in the STAC service
	UpdateOne(ctx context.Context, in *StacItem, opts ...grpc.CallOption) (*StacDbResponse, error)
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
	Recv() (*StacDbResponse, error)
	grpc.ClientStream
}

type stacServiceInsertClient struct {
	grpc.ClientStream
}

func (x *stacServiceInsertClient) Send(m *StacItem) error {
	return x.ClientStream.SendMsg(m)
}

func (x *stacServiceInsertClient) Recv() (*StacDbResponse, error) {
	m := new(StacDbResponse)
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
	Recv() (*StacDbResponse, error)
	grpc.ClientStream
}

type stacServiceUpdateClient struct {
	grpc.ClientStream
}

func (x *stacServiceUpdateClient) Send(m *StacItem) error {
	return x.ClientStream.SendMsg(m)
}

func (x *stacServiceUpdateClient) Recv() (*StacDbResponse, error) {
	m := new(StacDbResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stacServiceClient) Count(ctx context.Context, in *StacRequest, opts ...grpc.CallOption) (*StacDbResponse, error) {
	out := new(StacDbResponse)
	err := c.cc.Invoke(ctx, "/epl.protobuf.StacService/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stacServiceClient) DeleteOne(ctx context.Context, in *StacItem, opts ...grpc.CallOption) (*StacDbResponse, error) {
	out := new(StacDbResponse)
	err := c.cc.Invoke(ctx, "/epl.protobuf.StacService/DeleteOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stacServiceClient) SearchOne(ctx context.Context, in *StacRequest, opts ...grpc.CallOption) (*StacItem, error) {
	out := new(StacItem)
	err := c.cc.Invoke(ctx, "/epl.protobuf.StacService/SearchOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stacServiceClient) InsertOne(ctx context.Context, in *StacItem, opts ...grpc.CallOption) (*StacDbResponse, error) {
	out := new(StacDbResponse)
	err := c.cc.Invoke(ctx, "/epl.protobuf.StacService/InsertOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stacServiceClient) UpdateOne(ctx context.Context, in *StacItem, opts ...grpc.CallOption) (*StacDbResponse, error) {
	out := new(StacDbResponse)
	err := c.cc.Invoke(ctx, "/epl.protobuf.StacService/UpdateOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StacServiceServer is the server API for StacService service.
type StacServiceServer interface {
	//
	// using a search request, stream all the results that match the search filter
	Search(*StacRequest, StacService_SearchServer) error
	//
	// insert a stream of items into the STAC service
	Insert(StacService_InsertServer) error
	//
	// update a stream of items in the STAC service
	Update(StacService_UpdateServer) error
	//
	// count all the items in the Stac service according to the StacRequest filter
	Count(context.Context, *StacRequest) (*StacDbResponse, error)
	//
	// delete an item from the STAC service
	DeleteOne(context.Context, *StacItem) (*StacDbResponse, error)
	//
	// using a search request get the first item that matches the request
	SearchOne(context.Context, *StacRequest) (*StacItem, error)
	//
	// Insert one item into the STAC service
	InsertOne(context.Context, *StacItem) (*StacDbResponse, error)
	//
	// Update one item in the STAC service
	UpdateOne(context.Context, *StacItem) (*StacDbResponse, error)
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
	Send(*StacDbResponse) error
	Recv() (*StacItem, error)
	grpc.ServerStream
}

type stacServiceInsertServer struct {
	grpc.ServerStream
}

func (x *stacServiceInsertServer) Send(m *StacDbResponse) error {
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
	Send(*StacDbResponse) error
	Recv() (*StacItem, error)
	grpc.ServerStream
}

type stacServiceUpdateServer struct {
	grpc.ServerStream
}

func (x *stacServiceUpdateServer) Send(m *StacDbResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *stacServiceUpdateServer) Recv() (*StacItem, error) {
	m := new(StacItem)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StacService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StacRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/epl.protobuf.StacService/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacServiceServer).Count(ctx, req.(*StacRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StacService_DeleteOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StacItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StacServiceServer).DeleteOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/epl.protobuf.StacService/DeleteOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StacServiceServer).DeleteOne(ctx, req.(*StacItem))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "Count",
			Handler:    _StacService_Count_Handler,
		},
		{
			MethodName: "DeleteOne",
			Handler:    _StacService_DeleteOne_Handler,
		},
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
	proto.RegisterFile("epl/protobuf/stac_service.proto", fileDescriptor_stac_service_d8c84c09f1bcf066)
}

var fileDescriptor_stac_service_d8c84c09f1bcf066 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0xa7, 0xfe, 0x14, 0x1a, 0x5d, 0x0c, 0x59, 0x28, 0x16, 0x41, 0x98, 0x95, 0x20, 0x36,
	0xa2, 0x0f, 0xa0, 0xb4, 0xdd, 0xcc, 0xca, 0x61, 0xaa, 0x0b, 0xdd, 0x48, 0x1a, 0xaf, 0x99, 0x42,
	0x9b, 0xc4, 0xe4, 0xd6, 0x07, 0xf2, 0x05, 0x7c, 0x45, 0xe9, 0x04, 0xa1, 0x52, 0x74, 0x60, 0xba,
	0xcc, 0xb9, 0x27, 0x1f, 0xe7, 0x1e, 0x2e, 0x39, 0x03, 0x53, 0x33, 0x63, 0x35, 0xea, 0xb2, 0x7d,
	0x63, 0x0e, 0xb9, 0x78, 0x71, 0x60, 0x3f, 0x2a, 0x01, 0xc9, 0x5a, 0xa5, 0x87, 0x60, 0xea, 0xe4,
	0xc7, 0x10, 0x1f, 0x0f, 0xec, 0x7e, 0x76, 0xfd, 0xb5, 0x47, 0x0e, 0x0a, 0xe4, 0xa2, 0xf0, 0x9f,
	0xe9, 0x2d, 0x09, 0x0b, 0xe0, 0x56, 0xac, 0xe8, 0x49, 0xd2, 0x27, 0x24, 0x9d, 0x69, 0x09, 0xef,
	0x2d, 0x38, 0x8c, 0x8f, 0x86, 0xa3, 0x39, 0x42, 0x33, 0x9b, 0x5c, 0x05, 0x34, 0x27, 0xe1, 0x5c,
	0x39, 0xb0, 0x48, 0xff, 0x70, 0xc5, 0xa7, 0x43, 0x3d, 0x2f, 0x97, 0xe0, 0x8c, 0x56, 0x0e, 0x66,
	0x93, 0xf3, 0xc0, 0x53, 0x1e, 0xcd, 0x2b, 0x47, 0x18, 0x45, 0x49, 0xc9, 0x7e, 0xa6, 0x5b, 0x85,
	0xff, 0xed, 0xb2, 0x81, 0x43, 0x33, 0x12, 0xe5, 0x50, 0x03, 0xc2, 0xbd, 0xda, 0x3a, 0x0c, 0xbd,
	0x23, 0x91, 0x6f, 0xb5, 0x83, 0x6c, 0x53, 0x6c, 0x17, 0xc3, 0xd7, 0x3a, 0x26, 0x46, 0x46, 0x22,
	0xdf, 0xea, 0x08, 0x48, 0xfa, 0x44, 0xa6, 0x42, 0x37, 0xbf, 0x4c, 0xe9, 0xb4, 0x77, 0x42, 0x8b,
	0x4e, 0x5c, 0x04, 0xcf, 0x17, 0xb2, 0xc2, 0x55, 0x5b, 0x26, 0x42, 0x37, 0x4c, 0x82, 0xbe, 0x94,
	0xd6, 0x08, 0xc6, 0x4d, 0xc5, 0xa4, 0xae, 0xb9, 0x92, 0xac, 0x7f, 0x91, 0x9f, 0x3b, 0xbb, 0xc5,
	0x43, 0x51, 0x86, 0xeb, 0xf7, 0xcd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xa2, 0xf9, 0x77,
	0xdd, 0x02, 0x00, 0x00,
}
