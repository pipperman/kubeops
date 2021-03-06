// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.1
// source: api/v1/kubeops.proto

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

// KubeOpsApiClient is the client API for KubeOpsApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubeOpsApiClient interface {
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error)
	ListProject(ctx context.Context, in *ListProjectRequest, opts ...grpc.CallOption) (*ListProjectResponse, error)
	GetInventory(ctx context.Context, in *GetInventoryRequest, opts ...grpc.CallOption) (*GetInventoryResponse, error)
	RunPlaybook(ctx context.Context, in *RunPlaybookRequest, opts ...grpc.CallOption) (*RunPlaybookResult, error)
	RunAdhoc(ctx context.Context, in *RunAdhocRequest, opts ...grpc.CallOption) (*RunAdhocResult, error)
	WatchResult(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (KubeOpsApi_WatchResultClient, error)
	GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*GetResultResponse, error)
	ListResult(ctx context.Context, in *ListResultRequest, opts ...grpc.CallOption) (*ListResultResponse, error)
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type kubeOpsApiClient struct {
	cc grpc.ClientConnInterface
}

func NewKubeOpsApiClient(cc grpc.ClientConnInterface) KubeOpsApiClient {
	return &kubeOpsApiClient{cc}
}

func (c *kubeOpsApiClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error) {
	out := new(CreateProjectResponse)
	err := c.cc.Invoke(ctx, "/api.KubeOpsApi/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeOpsApiClient) ListProject(ctx context.Context, in *ListProjectRequest, opts ...grpc.CallOption) (*ListProjectResponse, error) {
	out := new(ListProjectResponse)
	err := c.cc.Invoke(ctx, "/api.KubeOpsApi/ListProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeOpsApiClient) GetInventory(ctx context.Context, in *GetInventoryRequest, opts ...grpc.CallOption) (*GetInventoryResponse, error) {
	out := new(GetInventoryResponse)
	err := c.cc.Invoke(ctx, "/api.KubeOpsApi/GetInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeOpsApiClient) RunPlaybook(ctx context.Context, in *RunPlaybookRequest, opts ...grpc.CallOption) (*RunPlaybookResult, error) {
	out := new(RunPlaybookResult)
	err := c.cc.Invoke(ctx, "/api.KubeOpsApi/RunPlaybook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeOpsApiClient) RunAdhoc(ctx context.Context, in *RunAdhocRequest, opts ...grpc.CallOption) (*RunAdhocResult, error) {
	out := new(RunAdhocResult)
	err := c.cc.Invoke(ctx, "/api.KubeOpsApi/RunAdhoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeOpsApiClient) WatchResult(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (KubeOpsApi_WatchResultClient, error) {
	stream, err := c.cc.NewStream(ctx, &KubeOpsApi_ServiceDesc.Streams[0], "/api.KubeOpsApi/WatchResult", opts...)
	if err != nil {
		return nil, err
	}
	x := &kubeOpsApiWatchResultClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KubeOpsApi_WatchResultClient interface {
	Recv() (*WatchStream, error)
	grpc.ClientStream
}

type kubeOpsApiWatchResultClient struct {
	grpc.ClientStream
}

func (x *kubeOpsApiWatchResultClient) Recv() (*WatchStream, error) {
	m := new(WatchStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kubeOpsApiClient) GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*GetResultResponse, error) {
	out := new(GetResultResponse)
	err := c.cc.Invoke(ctx, "/api.KubeOpsApi/GetResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeOpsApiClient) ListResult(ctx context.Context, in *ListResultRequest, opts ...grpc.CallOption) (*ListResultResponse, error) {
	out := new(ListResultResponse)
	err := c.cc.Invoke(ctx, "/api.KubeOpsApi/ListResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeOpsApiClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/api.KubeOpsApi/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubeOpsApiServer is the server API for KubeOpsApi service.
// All implementations must embed UnimplementedKubeOpsApiServer
// for forward compatibility
type KubeOpsApiServer interface {
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
	ListProject(context.Context, *ListProjectRequest) (*ListProjectResponse, error)
	GetInventory(context.Context, *GetInventoryRequest) (*GetInventoryResponse, error)
	RunPlaybook(context.Context, *RunPlaybookRequest) (*RunPlaybookResult, error)
	RunAdhoc(context.Context, *RunAdhocRequest) (*RunAdhocResult, error)
	WatchResult(*WatchRequest, KubeOpsApi_WatchResultServer) error
	GetResult(context.Context, *GetResultRequest) (*GetResultResponse, error)
	ListResult(context.Context, *ListResultRequest) (*ListResultResponse, error)
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	mustEmbedUnimplementedKubeOpsApiServer()
}

// UnimplementedKubeOpsApiServer must be embedded to have forward compatible implementations.
type UnimplementedKubeOpsApiServer struct {
}

func (UnimplementedKubeOpsApiServer) CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedKubeOpsApiServer) ListProject(context.Context, *ListProjectRequest) (*ListProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProject not implemented")
}
func (UnimplementedKubeOpsApiServer) GetInventory(context.Context, *GetInventoryRequest) (*GetInventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInventory not implemented")
}
func (UnimplementedKubeOpsApiServer) RunPlaybook(context.Context, *RunPlaybookRequest) (*RunPlaybookResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunPlaybook not implemented")
}
func (UnimplementedKubeOpsApiServer) RunAdhoc(context.Context, *RunAdhocRequest) (*RunAdhocResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunAdhoc not implemented")
}
func (UnimplementedKubeOpsApiServer) WatchResult(*WatchRequest, KubeOpsApi_WatchResultServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchResult not implemented")
}
func (UnimplementedKubeOpsApiServer) GetResult(context.Context, *GetResultRequest) (*GetResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}
func (UnimplementedKubeOpsApiServer) ListResult(context.Context, *ListResultRequest) (*ListResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResult not implemented")
}
func (UnimplementedKubeOpsApiServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedKubeOpsApiServer) mustEmbedUnimplementedKubeOpsApiServer() {}

// UnsafeKubeOpsApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubeOpsApiServer will
// result in compilation errors.
type UnsafeKubeOpsApiServer interface {
	mustEmbedUnimplementedKubeOpsApiServer()
}

func RegisterKubeOpsApiServer(s grpc.ServiceRegistrar, srv KubeOpsApiServer) {
	s.RegisterService(&KubeOpsApi_ServiceDesc, srv)
}

func _KubeOpsApi_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeOpsApiServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KubeOpsApi/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeOpsApiServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeOpsApi_ListProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeOpsApiServer).ListProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KubeOpsApi/ListProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeOpsApiServer).ListProject(ctx, req.(*ListProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeOpsApi_GetInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeOpsApiServer).GetInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KubeOpsApi/GetInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeOpsApiServer).GetInventory(ctx, req.(*GetInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeOpsApi_RunPlaybook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunPlaybookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeOpsApiServer).RunPlaybook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KubeOpsApi/RunPlaybook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeOpsApiServer).RunPlaybook(ctx, req.(*RunPlaybookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeOpsApi_RunAdhoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunAdhocRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeOpsApiServer).RunAdhoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KubeOpsApi/RunAdhoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeOpsApiServer).RunAdhoc(ctx, req.(*RunAdhocRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeOpsApi_WatchResult_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KubeOpsApiServer).WatchResult(m, &kubeOpsApiWatchResultServer{stream})
}

type KubeOpsApi_WatchResultServer interface {
	Send(*WatchStream) error
	grpc.ServerStream
}

type kubeOpsApiWatchResultServer struct {
	grpc.ServerStream
}

func (x *kubeOpsApiWatchResultServer) Send(m *WatchStream) error {
	return x.ServerStream.SendMsg(m)
}

func _KubeOpsApi_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeOpsApiServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KubeOpsApi/GetResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeOpsApiServer).GetResult(ctx, req.(*GetResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeOpsApi_ListResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeOpsApiServer).ListResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KubeOpsApi/ListResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeOpsApiServer).ListResult(ctx, req.(*ListResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeOpsApi_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeOpsApiServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KubeOpsApi/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeOpsApiServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KubeOpsApi_ServiceDesc is the grpc.ServiceDesc for KubeOpsApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KubeOpsApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.KubeOpsApi",
	HandlerType: (*KubeOpsApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _KubeOpsApi_CreateProject_Handler,
		},
		{
			MethodName: "ListProject",
			Handler:    _KubeOpsApi_ListProject_Handler,
		},
		{
			MethodName: "GetInventory",
			Handler:    _KubeOpsApi_GetInventory_Handler,
		},
		{
			MethodName: "RunPlaybook",
			Handler:    _KubeOpsApi_RunPlaybook_Handler,
		},
		{
			MethodName: "RunAdhoc",
			Handler:    _KubeOpsApi_RunAdhoc_Handler,
		},
		{
			MethodName: "GetResult",
			Handler:    _KubeOpsApi_GetResult_Handler,
		},
		{
			MethodName: "ListResult",
			Handler:    _KubeOpsApi_ListResult_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _KubeOpsApi_Health_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchResult",
			Handler:       _KubeOpsApi_WatchResult_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/v1/kubeops.proto",
}
