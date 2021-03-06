// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VideoUIClient is the client API for VideoUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoUIClient interface {
	// ListEngineSpecs returns a list of Video Engine(s) that can be started through the UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (VideoUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type videoUIClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoUIClient(cc grpc.ClientConnInterface) VideoUIClient {
	return &videoUIClient{cc}
}

func (c *videoUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (VideoUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &VideoUI_ServiceDesc.Streams[0], "/v1.VideoUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &videoUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VideoUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type videoUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *videoUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *videoUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.VideoUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoUIServer is the server API for VideoUI service.
// All implementations must embed UnimplementedVideoUIServer
// for forward compatibility
type VideoUIServer interface {
	// ListEngineSpecs returns a list of Video Engine(s) that can be started through the UI.
	ListEngineSpecs(*ListEngineSpecsRequest, VideoUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedVideoUIServer()
}

// UnimplementedVideoUIServer must be embedded to have forward compatible implementations.
type UnimplementedVideoUIServer struct {
}

func (UnimplementedVideoUIServer) ListEngineSpecs(*ListEngineSpecsRequest, VideoUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedVideoUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedVideoUIServer) mustEmbedUnimplementedVideoUIServer() {}

// UnsafeVideoUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoUIServer will
// result in compilation errors.
type UnsafeVideoUIServer interface {
	mustEmbedUnimplementedVideoUIServer()
}

func RegisterVideoUIServer(s grpc.ServiceRegistrar, srv VideoUIServer) {
	s.RegisterService(&VideoUI_ServiceDesc, srv)
}

func _VideoUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VideoUIServer).ListEngineSpecs(m, &videoUIListEngineSpecsServer{stream})
}

type VideoUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type videoUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *videoUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _VideoUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.VideoUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoUI_ServiceDesc is the grpc.ServiceDesc for VideoUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.VideoUI",
	HandlerType: (*VideoUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _VideoUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _VideoUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "video-ui.proto",
}
