// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/dataflow/v1beta3/snapshots.proto

package dataflowpb

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

// SnapshotsV1Beta3Client is the client API for SnapshotsV1Beta3 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SnapshotsV1Beta3Client interface {
	// Gets information about a snapshot.
	GetSnapshot(ctx context.Context, in *GetSnapshotRequest, opts ...grpc.CallOption) (*Snapshot, error)
	// Deletes a snapshot.
	DeleteSnapshot(ctx context.Context, in *DeleteSnapshotRequest, opts ...grpc.CallOption) (*DeleteSnapshotResponse, error)
	// Lists snapshots.
	ListSnapshots(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotsResponse, error)
}

type snapshotsV1Beta3Client struct {
	cc grpc.ClientConnInterface
}

func NewSnapshotsV1Beta3Client(cc grpc.ClientConnInterface) SnapshotsV1Beta3Client {
	return &snapshotsV1Beta3Client{cc}
}

func (c *snapshotsV1Beta3Client) GetSnapshot(ctx context.Context, in *GetSnapshotRequest, opts ...grpc.CallOption) (*Snapshot, error) {
	out := new(Snapshot)
	err := c.cc.Invoke(ctx, "/mockgcp.dataflow.v1beta3.SnapshotsV1Beta3/GetSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsV1Beta3Client) DeleteSnapshot(ctx context.Context, in *DeleteSnapshotRequest, opts ...grpc.CallOption) (*DeleteSnapshotResponse, error) {
	out := new(DeleteSnapshotResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.dataflow.v1beta3.SnapshotsV1Beta3/DeleteSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsV1Beta3Client) ListSnapshots(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotsResponse, error) {
	out := new(ListSnapshotsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.dataflow.v1beta3.SnapshotsV1Beta3/ListSnapshots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnapshotsV1Beta3Server is the server API for SnapshotsV1Beta3 service.
// All implementations must embed UnimplementedSnapshotsV1Beta3Server
// for forward compatibility
type SnapshotsV1Beta3Server interface {
	// Gets information about a snapshot.
	GetSnapshot(context.Context, *GetSnapshotRequest) (*Snapshot, error)
	// Deletes a snapshot.
	DeleteSnapshot(context.Context, *DeleteSnapshotRequest) (*DeleteSnapshotResponse, error)
	// Lists snapshots.
	ListSnapshots(context.Context, *ListSnapshotsRequest) (*ListSnapshotsResponse, error)
	mustEmbedUnimplementedSnapshotsV1Beta3Server()
}

// UnimplementedSnapshotsV1Beta3Server must be embedded to have forward compatible implementations.
type UnimplementedSnapshotsV1Beta3Server struct {
}

func (UnimplementedSnapshotsV1Beta3Server) GetSnapshot(context.Context, *GetSnapshotRequest) (*Snapshot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnapshot not implemented")
}
func (UnimplementedSnapshotsV1Beta3Server) DeleteSnapshot(context.Context, *DeleteSnapshotRequest) (*DeleteSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSnapshot not implemented")
}
func (UnimplementedSnapshotsV1Beta3Server) ListSnapshots(context.Context, *ListSnapshotsRequest) (*ListSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSnapshots not implemented")
}
func (UnimplementedSnapshotsV1Beta3Server) mustEmbedUnimplementedSnapshotsV1Beta3Server() {}

// UnsafeSnapshotsV1Beta3Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SnapshotsV1Beta3Server will
// result in compilation errors.
type UnsafeSnapshotsV1Beta3Server interface {
	mustEmbedUnimplementedSnapshotsV1Beta3Server()
}

func RegisterSnapshotsV1Beta3Server(s grpc.ServiceRegistrar, srv SnapshotsV1Beta3Server) {
	s.RegisterService(&SnapshotsV1Beta3_ServiceDesc, srv)
}

func _SnapshotsV1Beta3_GetSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsV1Beta3Server).GetSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.dataflow.v1beta3.SnapshotsV1Beta3/GetSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsV1Beta3Server).GetSnapshot(ctx, req.(*GetSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotsV1Beta3_DeleteSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsV1Beta3Server).DeleteSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.dataflow.v1beta3.SnapshotsV1Beta3/DeleteSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsV1Beta3Server).DeleteSnapshot(ctx, req.(*DeleteSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotsV1Beta3_ListSnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsV1Beta3Server).ListSnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.dataflow.v1beta3.SnapshotsV1Beta3/ListSnapshots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsV1Beta3Server).ListSnapshots(ctx, req.(*ListSnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SnapshotsV1Beta3_ServiceDesc is the grpc.ServiceDesc for SnapshotsV1Beta3 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SnapshotsV1Beta3_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.dataflow.v1beta3.SnapshotsV1Beta3",
	HandlerType: (*SnapshotsV1Beta3Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSnapshot",
			Handler:    _SnapshotsV1Beta3_GetSnapshot_Handler,
		},
		{
			MethodName: "DeleteSnapshot",
			Handler:    _SnapshotsV1Beta3_DeleteSnapshot_Handler,
		},
		{
			MethodName: "ListSnapshots",
			Handler:    _SnapshotsV1Beta3_ListSnapshots_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/dataflow/v1beta3/snapshots.proto",
}
