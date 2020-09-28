// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: proto/messenger.proto

package messenger

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMessagesRequest) Reset() {
	*x = GetMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messenger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesRequest) ProtoMessage() {}

func (x *GetMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messenger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetMessagesRequest) Descriptor() ([]byte, []int) {
	return file_proto_messenger_proto_rawDescGZIP(), []int{0}
}

type GetMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetMessagesResponse) Reset() {
	*x = GetMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messenger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesResponse) ProtoMessage() {}

func (x *GetMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messenger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesResponse.ProtoReflect.Descriptor instead.
func (*GetMessagesResponse) Descriptor() ([]byte, []int) {
	return file_proto_messenger_proto_rawDescGZIP(), []int{1}
}

func (x *GetMessagesResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetMessagesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateMessageRequest) Reset() {
	*x = CreateMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messenger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageRequest) ProtoMessage() {}

func (x *CreateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messenger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageRequest.ProtoReflect.Descriptor instead.
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_messenger_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMessageRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *CreateMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateMessageResponse) Reset() {
	*x = CreateMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messenger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageResponse) ProtoMessage() {}

func (x *CreateMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messenger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageResponse.ProtoReflect.Descriptor instead.
func (*CreateMessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_messenger_proto_rawDescGZIP(), []int{3}
}

func (x *CreateMessageResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *CreateMessageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_messenger_proto protoreflect.FileDescriptor

var file_proto_messenger_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67,
	0x65, 0x72, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x4d, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0xb3, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72,
	0x12, 0x50, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x1d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x54, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_messenger_proto_rawDescOnce sync.Once
	file_proto_messenger_proto_rawDescData = file_proto_messenger_proto_rawDesc
)

func file_proto_messenger_proto_rawDescGZIP() []byte {
	file_proto_messenger_proto_rawDescOnce.Do(func() {
		file_proto_messenger_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_messenger_proto_rawDescData)
	})
	return file_proto_messenger_proto_rawDescData
}

var file_proto_messenger_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_messenger_proto_goTypes = []interface{}{
	(*GetMessagesRequest)(nil),    // 0: messenger.GetMessagesRequest
	(*GetMessagesResponse)(nil),   // 1: messenger.GetMessagesResponse
	(*CreateMessageRequest)(nil),  // 2: messenger.CreateMessageRequest
	(*CreateMessageResponse)(nil), // 3: messenger.CreateMessageResponse
}
var file_proto_messenger_proto_depIdxs = []int32{
	0, // 0: messenger.Messenger.GetMessages:input_type -> messenger.GetMessagesRequest
	2, // 1: messenger.Messenger.CreateMessage:input_type -> messenger.CreateMessageRequest
	1, // 2: messenger.Messenger.GetMessages:output_type -> messenger.GetMessagesResponse
	3, // 3: messenger.Messenger.CreateMessage:output_type -> messenger.CreateMessageResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_messenger_proto_init() }
func file_proto_messenger_proto_init() {
	if File_proto_messenger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_messenger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessagesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_messenger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessagesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_messenger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_messenger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_messenger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_messenger_proto_goTypes,
		DependencyIndexes: file_proto_messenger_proto_depIdxs,
		MessageInfos:      file_proto_messenger_proto_msgTypes,
	}.Build()
	File_proto_messenger_proto = out.File
	file_proto_messenger_proto_rawDesc = nil
	file_proto_messenger_proto_goTypes = nil
	file_proto_messenger_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessengerClient is the client API for Messenger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessengerClient interface {
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (Messenger_GetMessagesClient, error)
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error)
}

type messengerClient struct {
	cc grpc.ClientConnInterface
}

func NewMessengerClient(cc grpc.ClientConnInterface) MessengerClient {
	return &messengerClient{cc}
}

func (c *messengerClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (Messenger_GetMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Messenger_serviceDesc.Streams[0], "/messenger.Messenger/GetMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &messengerGetMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Messenger_GetMessagesClient interface {
	Recv() (*GetMessagesResponse, error)
	grpc.ClientStream
}

type messengerGetMessagesClient struct {
	grpc.ClientStream
}

func (x *messengerGetMessagesClient) Recv() (*GetMessagesResponse, error) {
	m := new(GetMessagesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messengerClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error) {
	out := new(CreateMessageResponse)
	err := c.cc.Invoke(ctx, "/messenger.Messenger/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessengerServer is the server API for Messenger service.
type MessengerServer interface {
	GetMessages(*GetMessagesRequest, Messenger_GetMessagesServer) error
	CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error)
}

// UnimplementedMessengerServer can be embedded to have forward compatible implementations.
type UnimplementedMessengerServer struct {
}

func (*UnimplementedMessengerServer) GetMessages(*GetMessagesRequest, Messenger_GetMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (*UnimplementedMessengerServer) CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}

func RegisterMessengerServer(s *grpc.Server, srv MessengerServer) {
	s.RegisterService(&_Messenger_serviceDesc, srv)
}

func _Messenger_GetMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetMessagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessengerServer).GetMessages(m, &messengerGetMessagesServer{stream})
}

type Messenger_GetMessagesServer interface {
	Send(*GetMessagesResponse) error
	grpc.ServerStream
}

type messengerGetMessagesServer struct {
	grpc.ServerStream
}

func (x *messengerGetMessagesServer) Send(m *GetMessagesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Messenger_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messenger.Messenger/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Messenger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messenger.Messenger",
	HandlerType: (*MessengerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _Messenger_CreateMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMessages",
			Handler:       _Messenger_GetMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/messenger.proto",
}
