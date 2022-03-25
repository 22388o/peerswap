// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.1
// source: peerswaprpc/peerswaprpc.proto

package peerswaprpc

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

// PeerSwapClient is the client API for PeerSwap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeerSwapClient interface {
	SwapOut(ctx context.Context, in *SwapOutRequest, opts ...grpc.CallOption) (*SwapResponse, error)
	SwapIn(ctx context.Context, in *SwapInRequest, opts ...grpc.CallOption) (*SwapResponse, error)
	GetSwap(ctx context.Context, in *GetSwapRequest, opts ...grpc.CallOption) (*SwapResponse, error)
	ListSwaps(ctx context.Context, in *ListSwapsRequest, opts ...grpc.CallOption) (*ListSwapsResponse, error)
	ListPeers(ctx context.Context, in *ListPeersRequest, opts ...grpc.CallOption) (*ListPeersResponse, error)
	ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error)
	ListRequestedSwaps(ctx context.Context, in *ListRequestedSwapsRequest, opts ...grpc.CallOption) (*ListRequestedSwapsResponse, error)
	ListActiveSwaps(ctx context.Context, in *ListSwapsRequest, opts ...grpc.CallOption) (*ListSwapsResponse, error)
	RejectSwaps(ctx context.Context, in *RejectSwapsRequest, opts ...grpc.CallOption) (*RejectSwapsResponse, error)
	// policy
	ReloadPolicyFile(ctx context.Context, in *ReloadPolicyFileRequest, opts ...grpc.CallOption) (*ReloadPolicyFileResponse, error)
	AddPeer(ctx context.Context, in *AddPeerRequest, opts ...grpc.CallOption) (*AddPeerResponse, error)
	RemovePeer(ctx context.Context, in *RemovePeerRequest, opts ...grpc.CallOption) (*RemovePeerResponse, error)
	// Liquid Stuff
	LiquidGetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error)
	LiquidGetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	LiquidSendToAddress(ctx context.Context, in *SendToAddressRequest, opts ...grpc.CallOption) (*SendToAddressResponse, error)
	Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type peerSwapClient struct {
	cc grpc.ClientConnInterface
}

func NewPeerSwapClient(cc grpc.ClientConnInterface) PeerSwapClient {
	return &peerSwapClient{cc}
}

func (c *peerSwapClient) SwapOut(ctx context.Context, in *SwapOutRequest, opts ...grpc.CallOption) (*SwapResponse, error) {
	out := new(SwapResponse)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/SwapOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) SwapIn(ctx context.Context, in *SwapInRequest, opts ...grpc.CallOption) (*SwapResponse, error) {
	out := new(SwapResponse)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/SwapIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) GetSwap(ctx context.Context, in *GetSwapRequest, opts ...grpc.CallOption) (*SwapResponse, error) {
	out := new(SwapResponse)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/GetSwap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) ListSwaps(ctx context.Context, in *ListSwapsRequest, opts ...grpc.CallOption) (*ListSwapsResponse, error) {
	out := new(ListSwapsResponse)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/ListSwaps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) ListPeers(ctx context.Context, in *ListPeersRequest, opts ...grpc.CallOption) (*ListPeersResponse, error) {
	out := new(ListPeersResponse)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/ListPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error) {
	out := new(ListNodesResponse)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/ListNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) ListRequestedSwaps(ctx context.Context, in *ListRequestedSwapsRequest, opts ...grpc.CallOption) (*ListRequestedSwapsResponse, error) {
	out := new(ListRequestedSwapsResponse)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/ListRequestedSwaps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) ListActiveSwaps(ctx context.Context, in *ListSwapsRequest, opts ...grpc.CallOption) (*ListSwapsResponse, error) {
	out := new(ListSwapsResponse)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/ListActiveSwaps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) RejectSwaps(ctx context.Context, in *RejectSwapsRequest, opts ...grpc.CallOption) (*RejectSwapsResponse, error) {
	out := new(RejectSwapsResponse)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/RejectSwaps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) ReloadPolicyFile(ctx context.Context, in *ReloadPolicyFileRequest, opts ...grpc.CallOption) (*ReloadPolicyFileResponse, error) {
	out := new(ReloadPolicyFileResponse)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/ReloadPolicyFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) AddPeer(ctx context.Context, in *AddPeerRequest, opts ...grpc.CallOption) (*AddPeerResponse, error) {
	out := new(AddPeerResponse)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/AddPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) RemovePeer(ctx context.Context, in *RemovePeerRequest, opts ...grpc.CallOption) (*RemovePeerResponse, error) {
	out := new(RemovePeerResponse)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/RemovePeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) LiquidGetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error) {
	out := new(GetAddressResponse)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/LiquidGetAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) LiquidGetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/LiquidGetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) LiquidSendToAddress(ctx context.Context, in *SendToAddressRequest, opts ...grpc.CallOption) (*SendToAddressResponse, error) {
	out := new(SendToAddressResponse)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/LiquidSendToAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerSwapClient) Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/peerswap.PeerSwap/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerSwapServer is the server API for PeerSwap service.
// All implementations must embed UnimplementedPeerSwapServer
// for forward compatibility
type PeerSwapServer interface {
	SwapOut(context.Context, *SwapOutRequest) (*SwapResponse, error)
	SwapIn(context.Context, *SwapInRequest) (*SwapResponse, error)
	GetSwap(context.Context, *GetSwapRequest) (*SwapResponse, error)
	ListSwaps(context.Context, *ListSwapsRequest) (*ListSwapsResponse, error)
	ListPeers(context.Context, *ListPeersRequest) (*ListPeersResponse, error)
	ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error)
	ListRequestedSwaps(context.Context, *ListRequestedSwapsRequest) (*ListRequestedSwapsResponse, error)
	ListActiveSwaps(context.Context, *ListSwapsRequest) (*ListSwapsResponse, error)
	RejectSwaps(context.Context, *RejectSwapsRequest) (*RejectSwapsResponse, error)
	// policy
	ReloadPolicyFile(context.Context, *ReloadPolicyFileRequest) (*ReloadPolicyFileResponse, error)
	AddPeer(context.Context, *AddPeerRequest) (*AddPeerResponse, error)
	RemovePeer(context.Context, *RemovePeerRequest) (*RemovePeerResponse, error)
	// Liquid Stuff
	LiquidGetAddress(context.Context, *GetAddressRequest) (*GetAddressResponse, error)
	LiquidGetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	LiquidSendToAddress(context.Context, *SendToAddressRequest) (*SendToAddressResponse, error)
	Stop(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedPeerSwapServer()
}

// UnimplementedPeerSwapServer must be embedded to have forward compatible implementations.
type UnimplementedPeerSwapServer struct {
}

func (UnimplementedPeerSwapServer) SwapOut(context.Context, *SwapOutRequest) (*SwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapOut not implemented")
}
func (UnimplementedPeerSwapServer) SwapIn(context.Context, *SwapInRequest) (*SwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapIn not implemented")
}
func (UnimplementedPeerSwapServer) GetSwap(context.Context, *GetSwapRequest) (*SwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSwap not implemented")
}
func (UnimplementedPeerSwapServer) ListSwaps(context.Context, *ListSwapsRequest) (*ListSwapsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSwaps not implemented")
}
func (UnimplementedPeerSwapServer) ListPeers(context.Context, *ListPeersRequest) (*ListPeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPeers not implemented")
}
func (UnimplementedPeerSwapServer) ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodes not implemented")
}
func (UnimplementedPeerSwapServer) ListRequestedSwaps(context.Context, *ListRequestedSwapsRequest) (*ListRequestedSwapsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRequestedSwaps not implemented")
}
func (UnimplementedPeerSwapServer) ListActiveSwaps(context.Context, *ListSwapsRequest) (*ListSwapsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActiveSwaps not implemented")
}
func (UnimplementedPeerSwapServer) RejectSwaps(context.Context, *RejectSwapsRequest) (*RejectSwapsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectSwaps not implemented")
}
func (UnimplementedPeerSwapServer) ReloadPolicyFile(context.Context, *ReloadPolicyFileRequest) (*ReloadPolicyFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReloadPolicyFile not implemented")
}
func (UnimplementedPeerSwapServer) AddPeer(context.Context, *AddPeerRequest) (*AddPeerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPeer not implemented")
}
func (UnimplementedPeerSwapServer) RemovePeer(context.Context, *RemovePeerRequest) (*RemovePeerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePeer not implemented")
}
func (UnimplementedPeerSwapServer) LiquidGetAddress(context.Context, *GetAddressRequest) (*GetAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidGetAddress not implemented")
}
func (UnimplementedPeerSwapServer) LiquidGetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidGetBalance not implemented")
}
func (UnimplementedPeerSwapServer) LiquidSendToAddress(context.Context, *SendToAddressRequest) (*SendToAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidSendToAddress not implemented")
}
func (UnimplementedPeerSwapServer) Stop(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedPeerSwapServer) mustEmbedUnimplementedPeerSwapServer() {}

// UnsafePeerSwapServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeerSwapServer will
// result in compilation errors.
type UnsafePeerSwapServer interface {
	mustEmbedUnimplementedPeerSwapServer()
}

func RegisterPeerSwapServer(s grpc.ServiceRegistrar, srv PeerSwapServer) {
	s.RegisterService(&PeerSwap_ServiceDesc, srv)
}

func _PeerSwap_SwapOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwapOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).SwapOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/SwapOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).SwapOut(ctx, req.(*SwapOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_SwapIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwapInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).SwapIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/SwapIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).SwapIn(ctx, req.(*SwapInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_GetSwap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).GetSwap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/GetSwap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).GetSwap(ctx, req.(*GetSwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_ListSwaps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSwapsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).ListSwaps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/ListSwaps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).ListSwaps(ctx, req.(*ListSwapsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_ListPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).ListPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/ListPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).ListPeers(ctx, req.(*ListPeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/ListNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).ListNodes(ctx, req.(*ListNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_ListRequestedSwaps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequestedSwapsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).ListRequestedSwaps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/ListRequestedSwaps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).ListRequestedSwaps(ctx, req.(*ListRequestedSwapsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_ListActiveSwaps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSwapsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).ListActiveSwaps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/ListActiveSwaps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).ListActiveSwaps(ctx, req.(*ListSwapsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_RejectSwaps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectSwapsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).RejectSwaps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/RejectSwaps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).RejectSwaps(ctx, req.(*RejectSwapsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_ReloadPolicyFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadPolicyFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).ReloadPolicyFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/ReloadPolicyFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).ReloadPolicyFile(ctx, req.(*ReloadPolicyFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_AddPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).AddPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/AddPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).AddPeer(ctx, req.(*AddPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_RemovePeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).RemovePeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/RemovePeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).RemovePeer(ctx, req.(*RemovePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_LiquidGetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).LiquidGetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/LiquidGetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).LiquidGetAddress(ctx, req.(*GetAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_LiquidGetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).LiquidGetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/LiquidGetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).LiquidGetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_LiquidSendToAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendToAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).LiquidSendToAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/LiquidSendToAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).LiquidSendToAddress(ctx, req.(*SendToAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerSwap_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerSwapServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peerswap.PeerSwap/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerSwapServer).Stop(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PeerSwap_ServiceDesc is the grpc.ServiceDesc for PeerSwap service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PeerSwap_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "peerswap.PeerSwap",
	HandlerType: (*PeerSwapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SwapOut",
			Handler:    _PeerSwap_SwapOut_Handler,
		},
		{
			MethodName: "SwapIn",
			Handler:    _PeerSwap_SwapIn_Handler,
		},
		{
			MethodName: "GetSwap",
			Handler:    _PeerSwap_GetSwap_Handler,
		},
		{
			MethodName: "ListSwaps",
			Handler:    _PeerSwap_ListSwaps_Handler,
		},
		{
			MethodName: "ListPeers",
			Handler:    _PeerSwap_ListPeers_Handler,
		},
		{
			MethodName: "ListNodes",
			Handler:    _PeerSwap_ListNodes_Handler,
		},
		{
			MethodName: "ListRequestedSwaps",
			Handler:    _PeerSwap_ListRequestedSwaps_Handler,
		},
		{
			MethodName: "ListActiveSwaps",
			Handler:    _PeerSwap_ListActiveSwaps_Handler,
		},
		{
			MethodName: "RejectSwaps",
			Handler:    _PeerSwap_RejectSwaps_Handler,
		},
		{
			MethodName: "ReloadPolicyFile",
			Handler:    _PeerSwap_ReloadPolicyFile_Handler,
		},
		{
			MethodName: "AddPeer",
			Handler:    _PeerSwap_AddPeer_Handler,
		},
		{
			MethodName: "RemovePeer",
			Handler:    _PeerSwap_RemovePeer_Handler,
		},
		{
			MethodName: "LiquidGetAddress",
			Handler:    _PeerSwap_LiquidGetAddress_Handler,
		},
		{
			MethodName: "LiquidGetBalance",
			Handler:    _PeerSwap_LiquidGetBalance_Handler,
		},
		{
			MethodName: "LiquidSendToAddress",
			Handler:    _PeerSwap_LiquidSendToAddress_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _PeerSwap_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peerswaprpc/peerswaprpc.proto",
}
