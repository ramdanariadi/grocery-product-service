// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: proto/wishlist.proto

package wishlist

import (
	context "context"
	"github.com/ramdanariadi/grocery-product-service/main/response"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WishlistServiceClient is the client API for WishlistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WishlistServiceClient interface {
	Save(ctx context.Context, in *Wishlist, opts ...grpc.CallOption) (*response.Response, error)
	Delete(ctx context.Context, in *UserAndWishlistId, opts ...grpc.CallOption) (*response.Response, error)
	FindByUserId(ctx context.Context, in *WishlistUserId, opts ...grpc.CallOption) (*MultipleWishlistResponse, error)
	FindWishlistByProductId(ctx context.Context, in *UserAndProductId, opts ...grpc.CallOption) (*WishlistResponse, error)
}

type wishlistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWishlistServiceClient(cc grpc.ClientConnInterface) WishlistServiceClient {
	return &wishlistServiceClient{cc}
}

func (c *wishlistServiceClient) Save(ctx context.Context, in *Wishlist, opts ...grpc.CallOption) (*response.Response, error) {
	out := new(response.Response)
	err := c.cc.Invoke(ctx, "/proto.WishlistService/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wishlistServiceClient) Delete(ctx context.Context, in *UserAndWishlistId, opts ...grpc.CallOption) (*response.Response, error) {
	out := new(response.Response)
	err := c.cc.Invoke(ctx, "/proto.WishlistService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wishlistServiceClient) FindByUserId(ctx context.Context, in *WishlistUserId, opts ...grpc.CallOption) (*MultipleWishlistResponse, error) {
	out := new(MultipleWishlistResponse)
	err := c.cc.Invoke(ctx, "/proto.WishlistService/FindByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wishlistServiceClient) FindWishlistByProductId(ctx context.Context, in *UserAndProductId, opts ...grpc.CallOption) (*WishlistResponse, error) {
	out := new(WishlistResponse)
	err := c.cc.Invoke(ctx, "/proto.WishlistService/FindWishlistByProductId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WishlistServiceServer is the server API for WishlistService service.
// All implementations must embed UnimplementedWishlistServiceServer
// for forward compatibility
type WishlistServiceServer interface {
	Save(context.Context, *Wishlist) (*response.Response, error)
	Delete(context.Context, *UserAndWishlistId) (*response.Response, error)
	FindByUserId(context.Context, *WishlistUserId) (*MultipleWishlistResponse, error)
	FindWishlistByProductId(context.Context, *UserAndProductId) (*WishlistResponse, error)
	mustEmbedUnimplementedWishlistServiceServer()
}

// UnimplementedWishlistServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWishlistServiceServer struct {
}

func (UnimplementedWishlistServiceServer) Save(context.Context, *Wishlist) (*response.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedWishlistServiceServer) Delete(context.Context, *UserAndWishlistId) (*response.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedWishlistServiceServer) FindByUserId(context.Context, *WishlistUserId) (*MultipleWishlistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByUserId not implemented")
}
func (UnimplementedWishlistServiceServer) FindWishlistByProductId(context.Context, *UserAndProductId) (*WishlistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindWishlistByProductId not implemented")
}
func (UnimplementedWishlistServiceServer) mustEmbedUnimplementedWishlistServiceServer() {}

// UnsafeWishlistServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WishlistServiceServer will
// result in compilation errors.
type UnsafeWishlistServiceServer interface {
	mustEmbedUnimplementedWishlistServiceServer()
}

func RegisterWishlistServiceServer(s grpc.ServiceRegistrar, srv WishlistServiceServer) {
	s.RegisterService(&WishlistService_ServiceDesc, srv)
}

func _WishlistService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Wishlist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishlistServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WishlistService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishlistServiceServer).Save(ctx, req.(*Wishlist))
	}
	return interceptor(ctx, in, info, handler)
}

func _WishlistService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAndWishlistId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishlistServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WishlistService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishlistServiceServer).Delete(ctx, req.(*UserAndWishlistId))
	}
	return interceptor(ctx, in, info, handler)
}

func _WishlistService_FindByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WishlistUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishlistServiceServer).FindByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WishlistService/FindByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishlistServiceServer).FindByUserId(ctx, req.(*WishlistUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _WishlistService_FindWishlistByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAndProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishlistServiceServer).FindWishlistByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WishlistService/FindWishlistByProductId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishlistServiceServer).FindWishlistByProductId(ctx, req.(*UserAndProductId))
	}
	return interceptor(ctx, in, info, handler)
}

// WishlistService_ServiceDesc is the grpc.ServiceDesc for WishlistService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WishlistService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WishlistService",
	HandlerType: (*WishlistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _WishlistService_Save_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _WishlistService_Delete_Handler,
		},
		{
			MethodName: "FindByUserId",
			Handler:    _WishlistService_FindByUserId_Handler,
		},
		{
			MethodName: "FindWishlistByProductId",
			Handler:    _WishlistService_FindWishlistByProductId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/wishlist.proto",
}