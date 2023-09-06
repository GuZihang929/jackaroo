// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: position.proto

package position

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

const (
	Position_Create_FullMethodName  = "/position.Position/Create"
	Position_Update_FullMethodName  = "/position.Position/Update"
	Position_Remove_FullMethodName  = "/position.Position/Remove"
	Position_Refresh_FullMethodName = "/position.Position/Refresh"
)

// PositionClient is the client API for Position service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PositionClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
}

type positionClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionClient(cc grpc.ClientConnInterface) PositionClient {
	return &positionClient{cc}
}

func (c *positionClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, Position_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, Position_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, Position_Remove_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, Position_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PositionServer is the server API for Position service.
// All implementations must embed UnimplementedPositionServer
// for forward compatibility
type PositionServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
	mustEmbedUnimplementedPositionServer()
}

// UnimplementedPositionServer must be embedded to have forward compatible implementations.
type UnimplementedPositionServer struct {
}

func (UnimplementedPositionServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPositionServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPositionServer) Remove(context.Context, *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedPositionServer) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedPositionServer) mustEmbedUnimplementedPositionServer() {}

// UnsafePositionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PositionServer will
// result in compilation errors.
type UnsafePositionServer interface {
	mustEmbedUnimplementedPositionServer()
}

func RegisterPositionServer(s grpc.ServiceRegistrar, srv PositionServer) {
	s.RegisterService(&Position_ServiceDesc, srv)
}

func _Position_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Position_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Position_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Position_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Position_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Position_Remove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Position_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Position_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Position_ServiceDesc is the grpc.ServiceDesc for Position service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Position_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "position.Position",
	HandlerType: (*PositionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Position_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Position_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Position_Remove_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _Position_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "position.proto",
}
