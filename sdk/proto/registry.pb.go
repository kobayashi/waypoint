// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registry.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Push struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Push) Reset()         { *m = Push{} }
func (m *Push) String() string { return proto.CompactTextString(m) }
func (*Push) ProtoMessage()    {}
func (*Push) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{0}
}

func (m *Push) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Push.Unmarshal(m, b)
}
func (m *Push) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Push.Marshal(b, m, deterministic)
}
func (m *Push) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Push.Merge(m, src)
}
func (m *Push) XXX_Size() int {
	return xxx_messageInfo_Push.Size(m)
}
func (m *Push) XXX_DiscardUnknown() {
	xxx_messageInfo_Push.DiscardUnknown(m)
}

var xxx_messageInfo_Push proto.InternalMessageInfo

type Push_Resp struct {
	// result is the resulting opaque data type
	Result               *any.Any `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Push_Resp) Reset()         { *m = Push_Resp{} }
func (m *Push_Resp) String() string { return proto.CompactTextString(m) }
func (*Push_Resp) ProtoMessage()    {}
func (*Push_Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_41af05d40a615591, []int{0, 0}
}

func (m *Push_Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Push_Resp.Unmarshal(m, b)
}
func (m *Push_Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Push_Resp.Marshal(b, m, deterministic)
}
func (m *Push_Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Push_Resp.Merge(m, src)
}
func (m *Push_Resp) XXX_Size() int {
	return xxx_messageInfo_Push_Resp.Size(m)
}
func (m *Push_Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Push_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Push_Resp proto.InternalMessageInfo

func (m *Push_Resp) GetResult() *any.Any {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Push)(nil), "proto.Push")
	proto.RegisterType((*Push_Resp)(nil), "proto.Push.Resp")
}

func init() {
	proto.RegisterFile("registry.proto", fileDescriptor_41af05d40a615591)
}

var fileDescriptor_41af05d40a615591 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x4b, 0xfb, 0x30,
	0x18, 0x67, 0xd0, 0xff, 0xd8, 0xff, 0xb1, 0xb8, 0x11, 0xa6, 0xcc, 0x0a, 0x22, 0x9e, 0x14, 0x24,
	0x83, 0x6d, 0x37, 0x77, 0xd8, 0x18, 0x0a, 0xbb, 0x49, 0x07, 0xde, 0xbb, 0xfa, 0xac, 0x2b, 0x74,
	0x69, 0x4d, 0x9e, 0x1c, 0xf6, 0x3d, 0xfc, 0xc0, 0x92, 0x97, 0x89, 0x4e, 0x72, 0xd0, 0x53, 0xc9,
	0xef, 0xbd, 0x69, 0xe1, 0x54, 0x62, 0x51, 0x2a, 0x92, 0x7b, 0xde, 0xc8, 0x9a, 0x6a, 0xf6, 0xcf,
	0x3e, 0x92, 0x8b, 0xa2, 0xae, 0x8b, 0x0a, 0x87, 0xf6, 0xb4, 0xd6, 0x9b, 0x61, 0x26, 0xbc, 0x22,
	0xb9, 0x3c, 0xa6, 0x70, 0xd7, 0xd0, 0x81, 0x8c, 0x9b, 0x4a, 0x17, 0xa5, 0x70, 0xa7, 0x9b, 0x29,
	0x44, 0xcf, 0x5a, 0x6d, 0x93, 0x09, 0x44, 0x29, 0xaa, 0x86, 0xdd, 0x43, 0x5b, 0xa2, 0xd2, 0x15,
	0x0d, 0x5a, 0xd7, 0xad, 0xdb, 0x93, 0x51, 0x9f, 0xbb, 0x2c, 0x7e, 0xc8, 0xe2, 0x73, 0xb1, 0x4f,
	0xbd, 0x66, 0xf4, 0x1e, 0x41, 0x27, 0xf5, 0xeb, 0xd8, 0x0c, 0xba, 0x4b, 0x35, 0xd7, 0xb4, 0x45,
	0x41, 0x65, 0x9e, 0x51, 0x2d, 0xd9, 0xf9, 0x0f, 0xf7, 0xa3, 0x59, 0x92, 0x9c, 0x39, 0x80, 0x2f,
	0x77, 0x4d, 0x85, 0x3b, 0x14, 0xa4, 0x6c, 0xf9, 0x04, 0x22, 0xe3, 0x67, 0x7d, 0x4f, 0x3f, 0x69,
	0x91, 0xaf, 0x1a, 0xcc, 0xf9, 0x5c, 0x16, 0x2a, 0x09, 0x84, 0xb1, 0x31, 0x74, 0x8c, 0xcb, 0x08,
	0x83, 0x85, 0xdd, 0xa3, 0x44, 0x36, 0x85, 0xf8, 0x25, 0xab, 0xca, 0xd7, 0x8c, 0xf0, 0x0f, 0x95,
	0x0f, 0xd0, 0xfb, 0xea, 0xfe, 0x5d, 0xf5, 0x0c, 0xe2, 0x45, 0x2d, 0x36, 0x65, 0xb1, 0x22, 0xa9,
	0x73, 0x0a, 0x1a, 0x07, 0xde, 0xe8, 0xc4, 0xdc, 0xa9, 0xed, 0x3d, 0x2d, 0xe0, 0xbf, 0x03, 0xb5,
	0x44, 0x76, 0xf5, 0x5d, 0xf6, 0x49, 0xa4, 0xf8, 0xa6, 0x51, 0x51, 0xf0, 0x1d, 0xee, 0xa0, 0x63,
	0xbe, 0xbc, 0x9d, 0x14, 0xfb, 0x8c, 0xc0, 0x62, 0xee, 0x7e, 0x92, 0xc0, 0x25, 0xf5, 0x3c, 0x6a,
	0x24, 0xdc, 0xec, 0x5b, 0xb7, 0x2d, 0x30, 0xfe, 0x08, 0x00, 0x00, 0xff, 0xff, 0xba, 0x6c, 0xb6,
	0x65, 0xba, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistryClient interface {
	IsAuthenticator(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ImplementsResp, error)
	Auth(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error)
	AuthSpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error)
	ValidateAuth(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error)
	ValidateAuthSpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error)
	ConfigStruct(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Config_StructResp, error)
	Configure(ctx context.Context, in *Config_ConfigureRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	PushSpec(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FuncSpec, error)
	Push(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*Push_Resp, error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) IsAuthenticator(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ImplementsResp, error) {
	out := new(ImplementsResp)
	err := c.cc.Invoke(ctx, "/proto.Registry/IsAuthenticator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Auth(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Registry/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) AuthSpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error) {
	out := new(FuncSpec)
	err := c.cc.Invoke(ctx, "/proto.Registry/AuthSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ValidateAuth(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Registry/ValidateAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ValidateAuthSpec(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FuncSpec, error) {
	out := new(FuncSpec)
	err := c.cc.Invoke(ctx, "/proto.Registry/ValidateAuthSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ConfigStruct(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Config_StructResp, error) {
	out := new(Config_StructResp)
	err := c.cc.Invoke(ctx, "/proto.Registry/ConfigStruct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Configure(ctx context.Context, in *Config_ConfigureRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Registry/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) PushSpec(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FuncSpec, error) {
	out := new(FuncSpec)
	err := c.cc.Invoke(ctx, "/proto.Registry/PushSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Push(ctx context.Context, in *FuncSpec_Args, opts ...grpc.CallOption) (*Push_Resp, error) {
	out := new(Push_Resp)
	err := c.cc.Invoke(ctx, "/proto.Registry/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
type RegistryServer interface {
	IsAuthenticator(context.Context, *empty.Empty) (*ImplementsResp, error)
	Auth(context.Context, *FuncSpec_Args) (*empty.Empty, error)
	AuthSpec(context.Context, *empty.Empty) (*FuncSpec, error)
	ValidateAuth(context.Context, *FuncSpec_Args) (*empty.Empty, error)
	ValidateAuthSpec(context.Context, *empty.Empty) (*FuncSpec, error)
	ConfigStruct(context.Context, *empty.Empty) (*Config_StructResp, error)
	Configure(context.Context, *Config_ConfigureRequest) (*empty.Empty, error)
	PushSpec(context.Context, *Empty) (*FuncSpec, error)
	Push(context.Context, *FuncSpec_Args) (*Push_Resp, error)
}

// UnimplementedRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedRegistryServer struct {
}

func (*UnimplementedRegistryServer) IsAuthenticator(ctx context.Context, req *empty.Empty) (*ImplementsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAuthenticator not implemented")
}
func (*UnimplementedRegistryServer) Auth(ctx context.Context, req *FuncSpec_Args) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (*UnimplementedRegistryServer) AuthSpec(ctx context.Context, req *empty.Empty) (*FuncSpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthSpec not implemented")
}
func (*UnimplementedRegistryServer) ValidateAuth(ctx context.Context, req *FuncSpec_Args) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAuth not implemented")
}
func (*UnimplementedRegistryServer) ValidateAuthSpec(ctx context.Context, req *empty.Empty) (*FuncSpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAuthSpec not implemented")
}
func (*UnimplementedRegistryServer) ConfigStruct(ctx context.Context, req *empty.Empty) (*Config_StructResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigStruct not implemented")
}
func (*UnimplementedRegistryServer) Configure(ctx context.Context, req *Config_ConfigureRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (*UnimplementedRegistryServer) PushSpec(ctx context.Context, req *Empty) (*FuncSpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushSpec not implemented")
}
func (*UnimplementedRegistryServer) Push(ctx context.Context, req *FuncSpec_Args) (*Push_Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}

func RegisterRegistryServer(s *grpc.Server, srv RegistryServer) {
	s.RegisterService(&_Registry_serviceDesc, srv)
}

func _Registry_IsAuthenticator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).IsAuthenticator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Registry/IsAuthenticator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).IsAuthenticator(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuncSpec_Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Registry/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Auth(ctx, req.(*FuncSpec_Args))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_AuthSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).AuthSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Registry/AuthSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).AuthSpec(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ValidateAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuncSpec_Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ValidateAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Registry/ValidateAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ValidateAuth(ctx, req.(*FuncSpec_Args))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ValidateAuthSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ValidateAuthSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Registry/ValidateAuthSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ValidateAuthSpec(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ConfigStruct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ConfigStruct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Registry/ConfigStruct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ConfigStruct(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config_ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Registry/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Configure(ctx, req.(*Config_ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_PushSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).PushSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Registry/PushSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).PushSpec(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuncSpec_Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Registry/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Push(ctx, req.(*FuncSpec_Args))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAuthenticator",
			Handler:    _Registry_IsAuthenticator_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _Registry_Auth_Handler,
		},
		{
			MethodName: "AuthSpec",
			Handler:    _Registry_AuthSpec_Handler,
		},
		{
			MethodName: "ValidateAuth",
			Handler:    _Registry_ValidateAuth_Handler,
		},
		{
			MethodName: "ValidateAuthSpec",
			Handler:    _Registry_ValidateAuthSpec_Handler,
		},
		{
			MethodName: "ConfigStruct",
			Handler:    _Registry_ConfigStruct_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Registry_Configure_Handler,
		},
		{
			MethodName: "PushSpec",
			Handler:    _Registry_PushSpec_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _Registry_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry.proto",
}
