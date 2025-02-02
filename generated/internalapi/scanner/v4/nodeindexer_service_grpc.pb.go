// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: internalapi/scanner/v4/nodeindexer_service.proto

package v4

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NodeIndexer_CreateNodeIndexReport_FullMethodName = "/scanner.v4.NodeIndexer/CreateNodeIndexReport"
)

// NodeIndexerClient is the client API for NodeIndexer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeIndexerClient interface {
	// CreateNodeIndexReport creates an index report for the node the container runs on and returns the report.
	CreateNodeIndexReport(ctx context.Context, in *CreateNodeIndexReportRequest, opts ...grpc.CallOption) (*IndexReport, error)
}

type nodeIndexerClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeIndexerClient(cc grpc.ClientConnInterface) NodeIndexerClient {
	return &nodeIndexerClient{cc}
}

func (c *nodeIndexerClient) CreateNodeIndexReport(ctx context.Context, in *CreateNodeIndexReportRequest, opts ...grpc.CallOption) (*IndexReport, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IndexReport)
	err := c.cc.Invoke(ctx, NodeIndexer_CreateNodeIndexReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeIndexerServer is the server API for NodeIndexer service.
// All implementations should embed UnimplementedNodeIndexerServer
// for forward compatibility.
type NodeIndexerServer interface {
	// CreateNodeIndexReport creates an index report for the node the container runs on and returns the report.
	CreateNodeIndexReport(context.Context, *CreateNodeIndexReportRequest) (*IndexReport, error)
}

// UnimplementedNodeIndexerServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNodeIndexerServer struct{}

func (UnimplementedNodeIndexerServer) CreateNodeIndexReport(context.Context, *CreateNodeIndexReportRequest) (*IndexReport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodeIndexReport not implemented")
}
func (UnimplementedNodeIndexerServer) testEmbeddedByValue() {}

// UnsafeNodeIndexerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeIndexerServer will
// result in compilation errors.
type UnsafeNodeIndexerServer interface {
	mustEmbedUnimplementedNodeIndexerServer()
}

func RegisterNodeIndexerServer(s grpc.ServiceRegistrar, srv NodeIndexerServer) {
	// If the following call pancis, it indicates UnimplementedNodeIndexerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NodeIndexer_ServiceDesc, srv)
}

func _NodeIndexer_CreateNodeIndexReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeIndexReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeIndexerServer).CreateNodeIndexReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeIndexer_CreateNodeIndexReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeIndexerServer).CreateNodeIndexReport(ctx, req.(*CreateNodeIndexReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeIndexer_ServiceDesc is the grpc.ServiceDesc for NodeIndexer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeIndexer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scanner.v4.NodeIndexer",
	HandlerType: (*NodeIndexerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNodeIndexReport",
			Handler:    _NodeIndexer_CreateNodeIndexReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internalapi/scanner/v4/nodeindexer_service.proto",
}
