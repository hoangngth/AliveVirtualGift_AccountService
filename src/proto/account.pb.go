// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	account.proto

It has these top-level messages:
	Status
	LoginRequest
	LoginResponse
	LogoutRequest
	CreateRequest
	UpdateRequest
	ShowRequest
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Status struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto1.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Status) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto1.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto1.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto1.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LogoutRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CreateRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UpdateRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Avatar   string `protobuf:"bytes,3,opt,name=avatar" json:"avatar,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto1.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UpdateRequest) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UpdateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ShowRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *ShowRequest) Reset()                    { *m = ShowRequest{} }
func (m *ShowRequest) String() string            { return proto1.CompactTextString(m) }
func (*ShowRequest) ProtoMessage()               {}
func (*ShowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ShowRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto1.RegisterType((*Status)(nil), "proto.Status")
	proto1.RegisterType((*LoginRequest)(nil), "proto.LoginRequest")
	proto1.RegisterType((*LoginResponse)(nil), "proto.LoginResponse")
	proto1.RegisterType((*LogoutRequest)(nil), "proto.LogoutRequest")
	proto1.RegisterType((*CreateRequest)(nil), "proto.CreateRequest")
	proto1.RegisterType((*UpdateRequest)(nil), "proto.UpdateRequest")
	proto1.RegisterType((*ShowRequest)(nil), "proto.ShowRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AccountService service

type AccountServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*Status, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Status, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Status, error)
	Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*Status, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/proto.AccountService/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.AccountService/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.AccountService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.AccountService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/proto.AccountService/Show", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccountService service

type AccountServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*Status, error)
	Create(context.Context, *CreateRequest) (*Status, error)
	Update(context.Context, *UpdateRequest) (*Status, error)
	Show(context.Context, *ShowRequest) (*Status, error)
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AccountService/Show",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Show(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AccountService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AccountService_Logout_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AccountService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AccountService_Update_Handler,
		},
		{
			MethodName: "Show",
			Handler:    _AccountService_Show_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}

func init() { proto1.RegisterFile("account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0x31, 0x8d, 0x75, 0x34, 0x3d, 0x8c, 0x45, 0x42, 0x41, 0x90, 0x80, 0xa0, 0x87, 0xf6,
	0x50, 0x7f, 0x81, 0x08, 0x9e, 0x3c, 0xa5, 0xf8, 0x03, 0xd6, 0x64, 0xd0, 0xa0, 0xcd, 0xa6, 0xfb,
	0xd1, 0xfe, 0x79, 0x0f, 0xd2, 0xfd, 0x48, 0x93, 0x90, 0x9b, 0xa7, 0xe4, 0xcd, 0xbc, 0xd9, 0xb7,
	0xf3, 0xde, 0x42, 0xc2, 0x8a, 0x82, 0xeb, 0x5a, 0xad, 0x1a, 0xc1, 0x15, 0xc7, 0x89, 0xf9, 0x64,
	0x19, 0xc4, 0x1b, 0xc5, 0x94, 0x96, 0x98, 0xc2, 0xb9, 0xd4, 0x45, 0x41, 0x52, 0xa6, 0xc1, 0x5d,
	0xf0, 0x30, 0xcd, 0x3d, 0xcc, 0x5e, 0xe1, 0xea, 0x8d, 0x7f, 0x56, 0x75, 0x4e, 0x3b, 0x4d, 0x52,
	0xe1, 0x02, 0xa6, 0x5a, 0x92, 0xa8, 0xd9, 0x96, 0x0c, 0xf5, 0x22, 0x6f, 0xf1, 0xb1, 0xd7, 0x30,
	0x29, 0x0f, 0x5c, 0x94, 0x69, 0x68, 0x7b, 0x1e, 0x67, 0xf7, 0x90, 0xb8, 0x73, 0x64, 0xc3, 0x6b,
	0x49, 0x38, 0x87, 0x89, 0xe2, 0xdf, 0x54, 0xbb, 0x53, 0x2c, 0x70, 0x34, 0xae, 0x95, 0xd7, 0x1b,
	0xa7, 0xed, 0x20, 0x79, 0x11, 0xc4, 0x14, 0xfd, 0xf3, 0x5a, 0x88, 0x10, 0x99, 0x99, 0x33, 0x53,
	0x37, 0xff, 0x47, 0x49, 0xda, 0xb2, 0xea, 0x27, 0x8d, 0xac, 0xa4, 0x01, 0xd9, 0x16, 0x92, 0xf7,
	0xa6, 0xec, 0x48, 0xfa, 0xd1, 0x60, 0x6c, 0x34, 0xec, 0x8c, 0xe2, 0x0d, 0xc4, 0x6c, 0xcf, 0x14,
	0x13, 0x4e, 0xc6, 0xa1, 0xde, 0xc5, 0xa2, 0x81, 0x5f, 0xb7, 0x70, 0xb9, 0xf9, 0xe2, 0x07, 0x2f,
	0x36, 0x83, 0xb0, 0x2a, 0x8d, 0x54, 0x94, 0x87, 0x55, 0xb9, 0xfe, 0x0d, 0x60, 0xf6, 0x6c, 0x33,
	0xdd, 0x90, 0xd8, 0x57, 0x05, 0xe1, 0x1a, 0x26, 0xc6, 0x61, 0xbc, 0xb6, 0x29, 0xaf, 0xba, 0xb9,
	0x2d, 0xe6, 0xfd, 0xa2, 0x0b, 0x61, 0x09, 0xb1, 0xb5, 0x1b, 0x3b, 0xfd, 0x93, 0xfb, 0x8b, 0xc4,
	0x55, 0xdd, 0x33, 0x59, 0x42, 0x6c, 0x6d, 0x6f, 0xe9, 0xbd, 0x14, 0x46, 0xe8, 0xd6, 0xb2, 0x96,
	0xde, 0x73, 0x70, 0x48, 0x7f, 0x84, 0xe8, 0xb8, 0x32, 0xa2, 0x2f, 0x9f, 0xf6, 0x1f, 0x50, 0x3f,
	0x62, 0x83, 0x9e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xb5, 0x03, 0xe5, 0xd8, 0x02, 0x00,
	0x00,
}
