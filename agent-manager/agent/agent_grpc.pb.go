// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.2
// source: agent.proto

package agent

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

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	RegisterAgent(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (*AgentResponse, error)
	AgentStream(ctx context.Context, opts ...grpc.CallOption) (AgentService_AgentStreamClient, error)
	DeleteAgent(ctx context.Context, in *AgentDelete, opts ...grpc.CallOption) (*AgentResponse, error)
	Ping(ctx context.Context, opts ...grpc.CallOption) (AgentService_PingClient, error)
	ListAgents(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAgentsResponse, error)
	UpdateAgentType(ctx context.Context, in *AgentTypeUpdate, opts ...grpc.CallOption) (*Agent, error)
	UpdateAgentGroup(ctx context.Context, in *AgentGroupUpdate, opts ...grpc.CallOption) (*Agent, error)
	ListAgentCommands(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAgentsCommandsResponse, error)
	GetAgentByHostname(ctx context.Context, in *Hostname, opts ...grpc.CallOption) (*Agent, error)
	ListAgentsWithCommands(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAgentsResponse, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) RegisterAgent(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (*AgentResponse, error) {
	out := new(AgentResponse)
	err := c.cc.Invoke(ctx, "/agent.AgentService/RegisterAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) AgentStream(ctx context.Context, opts ...grpc.CallOption) (AgentService_AgentStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[0], "/agent.AgentService/AgentStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceAgentStreamClient{stream}
	return x, nil
}

type AgentService_AgentStreamClient interface {
	Send(*BidirectionalStream) error
	Recv() (*BidirectionalStream, error)
	grpc.ClientStream
}

type agentServiceAgentStreamClient struct {
	grpc.ClientStream
}

func (x *agentServiceAgentStreamClient) Send(m *BidirectionalStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentServiceAgentStreamClient) Recv() (*BidirectionalStream, error) {
	m := new(BidirectionalStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentServiceClient) DeleteAgent(ctx context.Context, in *AgentDelete, opts ...grpc.CallOption) (*AgentResponse, error) {
	out := new(AgentResponse)
	err := c.cc.Invoke(ctx, "/agent.AgentService/DeleteAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) Ping(ctx context.Context, opts ...grpc.CallOption) (AgentService_PingClient, error) {
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[1], "/agent.AgentService/Ping", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentServicePingClient{stream}
	return x, nil
}

type AgentService_PingClient interface {
	Send(*PingResponse) error
	Recv() (*PingRequest, error)
	grpc.ClientStream
}

type agentServicePingClient struct {
	grpc.ClientStream
}

func (x *agentServicePingClient) Send(m *PingResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentServicePingClient) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentServiceClient) ListAgents(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAgentsResponse, error) {
	out := new(ListAgentsResponse)
	err := c.cc.Invoke(ctx, "/agent.AgentService/ListAgents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) UpdateAgentType(ctx context.Context, in *AgentTypeUpdate, opts ...grpc.CallOption) (*Agent, error) {
	out := new(Agent)
	err := c.cc.Invoke(ctx, "/agent.AgentService/UpdateAgentType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) UpdateAgentGroup(ctx context.Context, in *AgentGroupUpdate, opts ...grpc.CallOption) (*Agent, error) {
	out := new(Agent)
	err := c.cc.Invoke(ctx, "/agent.AgentService/UpdateAgentGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) ListAgentCommands(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAgentsCommandsResponse, error) {
	out := new(ListAgentsCommandsResponse)
	err := c.cc.Invoke(ctx, "/agent.AgentService/ListAgentCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) GetAgentByHostname(ctx context.Context, in *Hostname, opts ...grpc.CallOption) (*Agent, error) {
	out := new(Agent)
	err := c.cc.Invoke(ctx, "/agent.AgentService/GetAgentByHostname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) ListAgentsWithCommands(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAgentsResponse, error) {
	out := new(ListAgentsResponse)
	err := c.cc.Invoke(ctx, "/agent.AgentService/ListAgentsWithCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations must embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	RegisterAgent(context.Context, *AgentRequest) (*AgentResponse, error)
	AgentStream(AgentService_AgentStreamServer) error
	DeleteAgent(context.Context, *AgentDelete) (*AgentResponse, error)
	Ping(AgentService_PingServer) error
	ListAgents(context.Context, *ListRequest) (*ListAgentsResponse, error)
	UpdateAgentType(context.Context, *AgentTypeUpdate) (*Agent, error)
	UpdateAgentGroup(context.Context, *AgentGroupUpdate) (*Agent, error)
	ListAgentCommands(context.Context, *ListRequest) (*ListAgentsCommandsResponse, error)
	GetAgentByHostname(context.Context, *Hostname) (*Agent, error)
	ListAgentsWithCommands(context.Context, *ListRequest) (*ListAgentsResponse, error)
	mustEmbedUnimplementedAgentServiceServer()
}

// UnimplementedAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (UnimplementedAgentServiceServer) RegisterAgent(context.Context, *AgentRequest) (*AgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAgent not implemented")
}
func (UnimplementedAgentServiceServer) AgentStream(AgentService_AgentStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AgentStream not implemented")
}
func (UnimplementedAgentServiceServer) DeleteAgent(context.Context, *AgentDelete) (*AgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAgent not implemented")
}
func (UnimplementedAgentServiceServer) Ping(AgentService_PingServer) error {
	return status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAgentServiceServer) ListAgents(context.Context, *ListRequest) (*ListAgentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgents not implemented")
}
func (UnimplementedAgentServiceServer) UpdateAgentType(context.Context, *AgentTypeUpdate) (*Agent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAgentType not implemented")
}
func (UnimplementedAgentServiceServer) UpdateAgentGroup(context.Context, *AgentGroupUpdate) (*Agent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAgentGroup not implemented")
}
func (UnimplementedAgentServiceServer) ListAgentCommands(context.Context, *ListRequest) (*ListAgentsCommandsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgentCommands not implemented")
}
func (UnimplementedAgentServiceServer) GetAgentByHostname(context.Context, *Hostname) (*Agent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgentByHostname not implemented")
}
func (UnimplementedAgentServiceServer) ListAgentsWithCommands(context.Context, *ListRequest) (*ListAgentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgentsWithCommands not implemented")
}
func (UnimplementedAgentServiceServer) mustEmbedUnimplementedAgentServiceServer() {}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_RegisterAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).RegisterAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.AgentService/RegisterAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).RegisterAgent(ctx, req.(*AgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_AgentStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServiceServer).AgentStream(&agentServiceAgentStreamServer{stream})
}

type AgentService_AgentStreamServer interface {
	Send(*BidirectionalStream) error
	Recv() (*BidirectionalStream, error)
	grpc.ServerStream
}

type agentServiceAgentStreamServer struct {
	grpc.ServerStream
}

func (x *agentServiceAgentStreamServer) Send(m *BidirectionalStream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentServiceAgentStreamServer) Recv() (*BidirectionalStream, error) {
	m := new(BidirectionalStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AgentService_DeleteAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).DeleteAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.AgentService/DeleteAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).DeleteAgent(ctx, req.(*AgentDelete))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_Ping_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServiceServer).Ping(&agentServicePingServer{stream})
}

type AgentService_PingServer interface {
	Send(*PingRequest) error
	Recv() (*PingResponse, error)
	grpc.ServerStream
}

type agentServicePingServer struct {
	grpc.ServerStream
}

func (x *agentServicePingServer) Send(m *PingRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentServicePingServer) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AgentService_ListAgents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).ListAgents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.AgentService/ListAgents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).ListAgents(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_UpdateAgentType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentTypeUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).UpdateAgentType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.AgentService/UpdateAgentType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).UpdateAgentType(ctx, req.(*AgentTypeUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_UpdateAgentGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentGroupUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).UpdateAgentGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.AgentService/UpdateAgentGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).UpdateAgentGroup(ctx, req.(*AgentGroupUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_ListAgentCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).ListAgentCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.AgentService/ListAgentCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).ListAgentCommands(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_GetAgentByHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hostname)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).GetAgentByHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.AgentService/GetAgentByHostname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).GetAgentByHostname(ctx, req.(*Hostname))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_ListAgentsWithCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).ListAgentsWithCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.AgentService/ListAgentsWithCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).ListAgentsWithCommands(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAgent",
			Handler:    _AgentService_RegisterAgent_Handler,
		},
		{
			MethodName: "DeleteAgent",
			Handler:    _AgentService_DeleteAgent_Handler,
		},
		{
			MethodName: "ListAgents",
			Handler:    _AgentService_ListAgents_Handler,
		},
		{
			MethodName: "UpdateAgentType",
			Handler:    _AgentService_UpdateAgentType_Handler,
		},
		{
			MethodName: "UpdateAgentGroup",
			Handler:    _AgentService_UpdateAgentGroup_Handler,
		},
		{
			MethodName: "ListAgentCommands",
			Handler:    _AgentService_ListAgentCommands_Handler,
		},
		{
			MethodName: "GetAgentByHostname",
			Handler:    _AgentService_GetAgentByHostname_Handler,
		},
		{
			MethodName: "ListAgentsWithCommands",
			Handler:    _AgentService_ListAgentsWithCommands_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AgentStream",
			Handler:       _AgentService_AgentStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Ping",
			Handler:       _AgentService_Ping_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "agent.proto",
}

// PanelServiceClient is the client API for PanelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PanelServiceClient interface {
	ProcessCommand(ctx context.Context, opts ...grpc.CallOption) (PanelService_ProcessCommandClient, error)
}

type panelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPanelServiceClient(cc grpc.ClientConnInterface) PanelServiceClient {
	return &panelServiceClient{cc}
}

func (c *panelServiceClient) ProcessCommand(ctx context.Context, opts ...grpc.CallOption) (PanelService_ProcessCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &PanelService_ServiceDesc.Streams[0], "/agent.PanelService/ProcessCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &panelServiceProcessCommandClient{stream}
	return x, nil
}

type PanelService_ProcessCommandClient interface {
	Send(*UtmCommand) error
	Recv() (*CommandResult, error)
	grpc.ClientStream
}

type panelServiceProcessCommandClient struct {
	grpc.ClientStream
}

func (x *panelServiceProcessCommandClient) Send(m *UtmCommand) error {
	return x.ClientStream.SendMsg(m)
}

func (x *panelServiceProcessCommandClient) Recv() (*CommandResult, error) {
	m := new(CommandResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PanelServiceServer is the server API for PanelService service.
// All implementations must embed UnimplementedPanelServiceServer
// for forward compatibility
type PanelServiceServer interface {
	ProcessCommand(PanelService_ProcessCommandServer) error
	mustEmbedUnimplementedPanelServiceServer()
}

// UnimplementedPanelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPanelServiceServer struct {
}

func (UnimplementedPanelServiceServer) ProcessCommand(PanelService_ProcessCommandServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessCommand not implemented")
}
func (UnimplementedPanelServiceServer) mustEmbedUnimplementedPanelServiceServer() {}

// UnsafePanelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PanelServiceServer will
// result in compilation errors.
type UnsafePanelServiceServer interface {
	mustEmbedUnimplementedPanelServiceServer()
}

func RegisterPanelServiceServer(s grpc.ServiceRegistrar, srv PanelServiceServer) {
	s.RegisterService(&PanelService_ServiceDesc, srv)
}

func _PanelService_ProcessCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PanelServiceServer).ProcessCommand(&panelServiceProcessCommandServer{stream})
}

type PanelService_ProcessCommandServer interface {
	Send(*CommandResult) error
	Recv() (*UtmCommand, error)
	grpc.ServerStream
}

type panelServiceProcessCommandServer struct {
	grpc.ServerStream
}

func (x *panelServiceProcessCommandServer) Send(m *CommandResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *panelServiceProcessCommandServer) Recv() (*UtmCommand, error) {
	m := new(UtmCommand)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PanelService_ServiceDesc is the grpc.ServiceDesc for PanelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PanelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent.PanelService",
	HandlerType: (*PanelServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProcessCommand",
			Handler:       _PanelService_ProcessCommand_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "agent.proto",
}

// AgentGroupServiceClient is the client API for AgentGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentGroupServiceClient interface {
	CreateGroup(ctx context.Context, in *AgentGroup, opts ...grpc.CallOption) (*AgentGroup, error)
	EditGroup(ctx context.Context, in *AgentGroup, opts ...grpc.CallOption) (*AgentGroup, error)
	ListGroups(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAgentsGroupResponse, error)
	DeleteGroup(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error)
}

type agentGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentGroupServiceClient(cc grpc.ClientConnInterface) AgentGroupServiceClient {
	return &agentGroupServiceClient{cc}
}

func (c *agentGroupServiceClient) CreateGroup(ctx context.Context, in *AgentGroup, opts ...grpc.CallOption) (*AgentGroup, error) {
	out := new(AgentGroup)
	err := c.cc.Invoke(ctx, "/agent.AgentGroupService/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentGroupServiceClient) EditGroup(ctx context.Context, in *AgentGroup, opts ...grpc.CallOption) (*AgentGroup, error) {
	out := new(AgentGroup)
	err := c.cc.Invoke(ctx, "/agent.AgentGroupService/EditGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentGroupServiceClient) ListGroups(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAgentsGroupResponse, error) {
	out := new(ListAgentsGroupResponse)
	err := c.cc.Invoke(ctx, "/agent.AgentGroupService/ListGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentGroupServiceClient) DeleteGroup(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/agent.AgentGroupService/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentGroupServiceServer is the server API for AgentGroupService service.
// All implementations must embed UnimplementedAgentGroupServiceServer
// for forward compatibility
type AgentGroupServiceServer interface {
	CreateGroup(context.Context, *AgentGroup) (*AgentGroup, error)
	EditGroup(context.Context, *AgentGroup) (*AgentGroup, error)
	ListGroups(context.Context, *ListRequest) (*ListAgentsGroupResponse, error)
	DeleteGroup(context.Context, *Id) (*Id, error)
	mustEmbedUnimplementedAgentGroupServiceServer()
}

// UnimplementedAgentGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentGroupServiceServer struct {
}

func (UnimplementedAgentGroupServiceServer) CreateGroup(context.Context, *AgentGroup) (*AgentGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedAgentGroupServiceServer) EditGroup(context.Context, *AgentGroup) (*AgentGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditGroup not implemented")
}
func (UnimplementedAgentGroupServiceServer) ListGroups(context.Context, *ListRequest) (*ListAgentsGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroups not implemented")
}
func (UnimplementedAgentGroupServiceServer) DeleteGroup(context.Context, *Id) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedAgentGroupServiceServer) mustEmbedUnimplementedAgentGroupServiceServer() {}

// UnsafeAgentGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentGroupServiceServer will
// result in compilation errors.
type UnsafeAgentGroupServiceServer interface {
	mustEmbedUnimplementedAgentGroupServiceServer()
}

func RegisterAgentGroupServiceServer(s grpc.ServiceRegistrar, srv AgentGroupServiceServer) {
	s.RegisterService(&AgentGroupService_ServiceDesc, srv)
}

func _AgentGroupService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentGroupServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.AgentGroupService/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentGroupServiceServer).CreateGroup(ctx, req.(*AgentGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentGroupService_EditGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentGroupServiceServer).EditGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.AgentGroupService/EditGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentGroupServiceServer).EditGroup(ctx, req.(*AgentGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentGroupService_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentGroupServiceServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.AgentGroupService/ListGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentGroupServiceServer).ListGroups(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentGroupService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentGroupServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.AgentGroupService/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentGroupServiceServer).DeleteGroup(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentGroupService_ServiceDesc is the grpc.ServiceDesc for AgentGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent.AgentGroupService",
	HandlerType: (*AgentGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _AgentGroupService_CreateGroup_Handler,
		},
		{
			MethodName: "EditGroup",
			Handler:    _AgentGroupService_EditGroup_Handler,
		},
		{
			MethodName: "ListGroups",
			Handler:    _AgentGroupService_ListGroups_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _AgentGroupService_DeleteGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}

// AgentTypeServiceClient is the client API for AgentTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentTypeServiceClient interface {
	ListAgentTypes(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAgentsTypeResponse, error)
}

type agentTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentTypeServiceClient(cc grpc.ClientConnInterface) AgentTypeServiceClient {
	return &agentTypeServiceClient{cc}
}

func (c *agentTypeServiceClient) ListAgentTypes(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListAgentsTypeResponse, error) {
	out := new(ListAgentsTypeResponse)
	err := c.cc.Invoke(ctx, "/agent.AgentTypeService/ListAgentTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentTypeServiceServer is the server API for AgentTypeService service.
// All implementations must embed UnimplementedAgentTypeServiceServer
// for forward compatibility
type AgentTypeServiceServer interface {
	ListAgentTypes(context.Context, *ListRequest) (*ListAgentsTypeResponse, error)
	mustEmbedUnimplementedAgentTypeServiceServer()
}

// UnimplementedAgentTypeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentTypeServiceServer struct {
}

func (UnimplementedAgentTypeServiceServer) ListAgentTypes(context.Context, *ListRequest) (*ListAgentsTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgentTypes not implemented")
}
func (UnimplementedAgentTypeServiceServer) mustEmbedUnimplementedAgentTypeServiceServer() {}

// UnsafeAgentTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentTypeServiceServer will
// result in compilation errors.
type UnsafeAgentTypeServiceServer interface {
	mustEmbedUnimplementedAgentTypeServiceServer()
}

func RegisterAgentTypeServiceServer(s grpc.ServiceRegistrar, srv AgentTypeServiceServer) {
	s.RegisterService(&AgentTypeService_ServiceDesc, srv)
}

func _AgentTypeService_ListAgentTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTypeServiceServer).ListAgentTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.AgentTypeService/ListAgentTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTypeServiceServer).ListAgentTypes(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentTypeService_ServiceDesc is the grpc.ServiceDesc for AgentTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent.AgentTypeService",
	HandlerType: (*AgentTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAgentTypes",
			Handler:    _AgentTypeService_ListAgentTypes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}