// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/resourcemanager/v3/tag_holds.proto

package resourcemanagerpb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TagHoldsClient is the client API for TagHolds service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagHoldsClient interface {
	// Creates a TagHold. Returns ALREADY_EXISTS if a TagHold with the same
	// resource and origin exists under the same TagValue.
	CreateTagHold(ctx context.Context, in *CreateTagHoldRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes a TagHold.
	DeleteTagHold(ctx context.Context, in *DeleteTagHoldRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Lists TagHolds under a TagValue.
	ListTagHolds(ctx context.Context, in *ListTagHoldsRequest, opts ...grpc.CallOption) (*ListTagHoldsResponse, error)
}

type tagHoldsClient struct {
	cc grpc.ClientConnInterface
}

func NewTagHoldsClient(cc grpc.ClientConnInterface) TagHoldsClient {
	return &tagHoldsClient{cc}
}

func (c *tagHoldsClient) CreateTagHold(ctx context.Context, in *CreateTagHoldRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.resourcemanager.v3.TagHolds/CreateTagHold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagHoldsClient) DeleteTagHold(ctx context.Context, in *DeleteTagHoldRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.resourcemanager.v3.TagHolds/DeleteTagHold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagHoldsClient) ListTagHolds(ctx context.Context, in *ListTagHoldsRequest, opts ...grpc.CallOption) (*ListTagHoldsResponse, error) {
	out := new(ListTagHoldsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.resourcemanager.v3.TagHolds/ListTagHolds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagHoldsServer is the server API for TagHolds service.
// All implementations must embed UnimplementedTagHoldsServer
// for forward compatibility
type TagHoldsServer interface {
	// Creates a TagHold. Returns ALREADY_EXISTS if a TagHold with the same
	// resource and origin exists under the same TagValue.
	CreateTagHold(context.Context, *CreateTagHoldRequest) (*longrunningpb.Operation, error)
	// Deletes a TagHold.
	DeleteTagHold(context.Context, *DeleteTagHoldRequest) (*longrunningpb.Operation, error)
	// Lists TagHolds under a TagValue.
	ListTagHolds(context.Context, *ListTagHoldsRequest) (*ListTagHoldsResponse, error)
	mustEmbedUnimplementedTagHoldsServer()
}

// UnimplementedTagHoldsServer must be embedded to have forward compatible implementations.
type UnimplementedTagHoldsServer struct {
}

func (UnimplementedTagHoldsServer) CreateTagHold(context.Context, *CreateTagHoldRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTagHold not implemented")
}
func (UnimplementedTagHoldsServer) DeleteTagHold(context.Context, *DeleteTagHoldRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTagHold not implemented")
}
func (UnimplementedTagHoldsServer) ListTagHolds(context.Context, *ListTagHoldsRequest) (*ListTagHoldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTagHolds not implemented")
}
func (UnimplementedTagHoldsServer) mustEmbedUnimplementedTagHoldsServer() {}

// UnsafeTagHoldsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagHoldsServer will
// result in compilation errors.
type UnsafeTagHoldsServer interface {
	mustEmbedUnimplementedTagHoldsServer()
}

func RegisterTagHoldsServer(s grpc.ServiceRegistrar, srv TagHoldsServer) {
	s.RegisterService(&TagHolds_ServiceDesc, srv)
}

func _TagHolds_CreateTagHold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagHoldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagHoldsServer).CreateTagHold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.resourcemanager.v3.TagHolds/CreateTagHold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagHoldsServer).CreateTagHold(ctx, req.(*CreateTagHoldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagHolds_DeleteTagHold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTagHoldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagHoldsServer).DeleteTagHold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.resourcemanager.v3.TagHolds/DeleteTagHold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagHoldsServer).DeleteTagHold(ctx, req.(*DeleteTagHoldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagHolds_ListTagHolds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagHoldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagHoldsServer).ListTagHolds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.resourcemanager.v3.TagHolds/ListTagHolds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagHoldsServer).ListTagHolds(ctx, req.(*ListTagHoldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TagHolds_ServiceDesc is the grpc.ServiceDesc for TagHolds service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TagHolds_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.resourcemanager.v3.TagHolds",
	HandlerType: (*TagHoldsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTagHold",
			Handler:    _TagHolds_CreateTagHold_Handler,
		},
		{
			MethodName: "DeleteTagHold",
			Handler:    _TagHolds_DeleteTagHold_Handler,
		},
		{
			MethodName: "ListTagHolds",
			Handler:    _TagHolds_ListTagHolds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/resourcemanager/v3/tag_holds.proto",
}
