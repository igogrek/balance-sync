// Code generated by protoc-gen-go. DO NOT EDIT.
// source: balance.proto

package balance

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BalanceRequest struct {
	Did                  string   `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceRequest) Reset()         { *m = BalanceRequest{} }
func (m *BalanceRequest) String() string { return proto.CompactTextString(m) }
func (*BalanceRequest) ProtoMessage()    {}
func (*BalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_balance_3d5da835e355dd94, []int{0}
}
func (m *BalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceRequest.Unmarshal(m, b)
}
func (m *BalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceRequest.Marshal(b, m, deterministic)
}
func (dst *BalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceRequest.Merge(dst, src)
}
func (m *BalanceRequest) XXX_Size() int {
	return xxx_messageInfo_BalanceRequest.Size(m)
}
func (m *BalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceRequest proto.InternalMessageInfo

func (m *BalanceRequest) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

type BalanceResponse struct {
	Balance              int32     `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Actions              []*Action `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BalanceResponse) Reset()         { *m = BalanceResponse{} }
func (m *BalanceResponse) String() string { return proto.CompactTextString(m) }
func (*BalanceResponse) ProtoMessage()    {}
func (*BalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_balance_3d5da835e355dd94, []int{1}
}
func (m *BalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceResponse.Unmarshal(m, b)
}
func (m *BalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceResponse.Marshal(b, m, deterministic)
}
func (dst *BalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceResponse.Merge(dst, src)
}
func (m *BalanceResponse) XXX_Size() int {
	return xxx_messageInfo_BalanceResponse.Size(m)
}
func (m *BalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceResponse proto.InternalMessageInfo

func (m *BalanceResponse) GetBalance() int32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *BalanceResponse) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

type BalanceUpdate struct {
	Did                  string    `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	Actions              []*Action `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	UpdatedBy            string    `protobuf:"bytes,3,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BalanceUpdate) Reset()         { *m = BalanceUpdate{} }
func (m *BalanceUpdate) String() string { return proto.CompactTextString(m) }
func (*BalanceUpdate) ProtoMessage()    {}
func (*BalanceUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_balance_3d5da835e355dd94, []int{2}
}
func (m *BalanceUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceUpdate.Unmarshal(m, b)
}
func (m *BalanceUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceUpdate.Marshal(b, m, deterministic)
}
func (dst *BalanceUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceUpdate.Merge(dst, src)
}
func (m *BalanceUpdate) XXX_Size() int {
	return xxx_messageInfo_BalanceUpdate.Size(m)
}
func (m *BalanceUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceUpdate proto.InternalMessageInfo

func (m *BalanceUpdate) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *BalanceUpdate) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *BalanceUpdate) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

type BalanceUpdateResponse struct {
	Message              string         `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Update               *BalanceUpdate `protobuf:"bytes,2,opt,name=update,proto3" json:"update,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BalanceUpdateResponse) Reset()         { *m = BalanceUpdateResponse{} }
func (m *BalanceUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*BalanceUpdateResponse) ProtoMessage()    {}
func (*BalanceUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_balance_3d5da835e355dd94, []int{3}
}
func (m *BalanceUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceUpdateResponse.Unmarshal(m, b)
}
func (m *BalanceUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceUpdateResponse.Marshal(b, m, deterministic)
}
func (dst *BalanceUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceUpdateResponse.Merge(dst, src)
}
func (m *BalanceUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_BalanceUpdateResponse.Size(m)
}
func (m *BalanceUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceUpdateResponse proto.InternalMessageInfo

func (m *BalanceUpdateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BalanceUpdateResponse) GetUpdate() *BalanceUpdate {
	if m != nil {
		return m.Update
	}
	return nil
}

type SubscribeRequest struct {
	SubscriberId         string   `protobuf:"bytes,1,opt,name=subscriberId,proto3" json:"subscriberId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_balance_3d5da835e355dd94, []int{4}
}
func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (dst *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(dst, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetSubscriberId() string {
	if m != nil {
		return m.SubscriberId
	}
	return ""
}

type Action struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp            string   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_balance_3d5da835e355dd94, []int{5}
}
func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (dst *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(dst, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Action) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

type MilestoneResponse struct {
	Milestones           []string `protobuf:"bytes,1,rep,name=milestones,proto3" json:"milestones,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MilestoneResponse) Reset()         { *m = MilestoneResponse{} }
func (m *MilestoneResponse) String() string { return proto.CompactTextString(m) }
func (*MilestoneResponse) ProtoMessage()    {}
func (*MilestoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_balance_3d5da835e355dd94, []int{6}
}
func (m *MilestoneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MilestoneResponse.Unmarshal(m, b)
}
func (m *MilestoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MilestoneResponse.Marshal(b, m, deterministic)
}
func (dst *MilestoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MilestoneResponse.Merge(dst, src)
}
func (m *MilestoneResponse) XXX_Size() int {
	return xxx_messageInfo_MilestoneResponse.Size(m)
}
func (m *MilestoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MilestoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MilestoneResponse proto.InternalMessageInfo

func (m *MilestoneResponse) GetMilestones() []string {
	if m != nil {
		return m.Milestones
	}
	return nil
}

type ClaimPrizeResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClaimPrizeResponse) Reset()         { *m = ClaimPrizeResponse{} }
func (m *ClaimPrizeResponse) String() string { return proto.CompactTextString(m) }
func (*ClaimPrizeResponse) ProtoMessage()    {}
func (*ClaimPrizeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_balance_3d5da835e355dd94, []int{7}
}
func (m *ClaimPrizeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClaimPrizeResponse.Unmarshal(m, b)
}
func (m *ClaimPrizeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClaimPrizeResponse.Marshal(b, m, deterministic)
}
func (dst *ClaimPrizeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimPrizeResponse.Merge(dst, src)
}
func (m *ClaimPrizeResponse) XXX_Size() int {
	return xxx_messageInfo_ClaimPrizeResponse.Size(m)
}
func (m *ClaimPrizeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimPrizeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimPrizeResponse proto.InternalMessageInfo

func (m *ClaimPrizeResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*BalanceRequest)(nil), "BalanceRequest")
	proto.RegisterType((*BalanceResponse)(nil), "BalanceResponse")
	proto.RegisterType((*BalanceUpdate)(nil), "BalanceUpdate")
	proto.RegisterType((*BalanceUpdateResponse)(nil), "BalanceUpdateResponse")
	proto.RegisterType((*SubscribeRequest)(nil), "SubscribeRequest")
	proto.RegisterType((*Action)(nil), "Action")
	proto.RegisterType((*MilestoneResponse)(nil), "MilestoneResponse")
	proto.RegisterType((*ClaimPrizeResponse)(nil), "ClaimPrizeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BalanceClient is the client API for Balance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BalanceClient interface {
	GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	UpdateBalance(ctx context.Context, in *BalanceUpdate, opts ...grpc.CallOption) (*BalanceUpdateResponse, error)
	UpdateBalanceSubscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Balance_UpdateBalanceSubscribeClient, error)
	GetMilestones(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*MilestoneResponse, error)
	ClaimPrize(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*ClaimPrizeResponse, error)
}

type balanceClient struct {
	cc *grpc.ClientConn
}

func NewBalanceClient(cc *grpc.ClientConn) BalanceClient {
	return &balanceClient{cc}
}

func (c *balanceClient) GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/balance/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balanceClient) UpdateBalance(ctx context.Context, in *BalanceUpdate, opts ...grpc.CallOption) (*BalanceUpdateResponse, error) {
	out := new(BalanceUpdateResponse)
	err := c.cc.Invoke(ctx, "/balance/UpdateBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balanceClient) UpdateBalanceSubscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Balance_UpdateBalanceSubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Balance_serviceDesc.Streams[0], "/balance/UpdateBalanceSubscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &balanceUpdateBalanceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Balance_UpdateBalanceSubscribeClient interface {
	Recv() (*BalanceUpdateResponse, error)
	grpc.ClientStream
}

type balanceUpdateBalanceSubscribeClient struct {
	grpc.ClientStream
}

func (x *balanceUpdateBalanceSubscribeClient) Recv() (*BalanceUpdateResponse, error) {
	m := new(BalanceUpdateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *balanceClient) GetMilestones(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*MilestoneResponse, error) {
	out := new(MilestoneResponse)
	err := c.cc.Invoke(ctx, "/balance/GetMilestones", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balanceClient) ClaimPrize(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*ClaimPrizeResponse, error) {
	out := new(ClaimPrizeResponse)
	err := c.cc.Invoke(ctx, "/balance/ClaimPrize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalanceServer is the server API for Balance service.
type BalanceServer interface {
	GetBalance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	UpdateBalance(context.Context, *BalanceUpdate) (*BalanceUpdateResponse, error)
	UpdateBalanceSubscribe(*SubscribeRequest, Balance_UpdateBalanceSubscribeServer) error
	GetMilestones(context.Context, *BalanceRequest) (*MilestoneResponse, error)
	ClaimPrize(context.Context, *BalanceRequest) (*ClaimPrizeResponse, error)
}

func RegisterBalanceServer(s *grpc.Server, srv BalanceServer) {
	s.RegisterService(&_Balance_serviceDesc, srv)
}

func _Balance_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/balance/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServer).GetBalance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balance_UpdateBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServer).UpdateBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/balance/UpdateBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServer).UpdateBalance(ctx, req.(*BalanceUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balance_UpdateBalanceSubscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BalanceServer).UpdateBalanceSubscribe(m, &balanceUpdateBalanceSubscribeServer{stream})
}

type Balance_UpdateBalanceSubscribeServer interface {
	Send(*BalanceUpdateResponse) error
	grpc.ServerStream
}

type balanceUpdateBalanceSubscribeServer struct {
	grpc.ServerStream
}

func (x *balanceUpdateBalanceSubscribeServer) Send(m *BalanceUpdateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Balance_GetMilestones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServer).GetMilestones(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/balance/GetMilestones",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServer).GetMilestones(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balance_ClaimPrize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServer).ClaimPrize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/balance/ClaimPrize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServer).ClaimPrize(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Balance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "balance",
	HandlerType: (*BalanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _Balance_GetBalance_Handler,
		},
		{
			MethodName: "UpdateBalance",
			Handler:    _Balance_UpdateBalance_Handler,
		},
		{
			MethodName: "GetMilestones",
			Handler:    _Balance_GetMilestones_Handler,
		},
		{
			MethodName: "ClaimPrize",
			Handler:    _Balance_ClaimPrize_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateBalanceSubscribe",
			Handler:       _Balance_UpdateBalanceSubscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "balance.proto",
}

func init() { proto.RegisterFile("balance.proto", fileDescriptor_balance_3d5da835e355dd94) }

var fileDescriptor_balance_3d5da835e355dd94 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4b, 0x6b, 0xdb, 0x40,
	0x18, 0x44, 0x12, 0xb5, 0xab, 0xcf, 0x95, 0x1f, 0x5b, 0x6a, 0x84, 0x29, 0x45, 0xdd, 0x43, 0xd1,
	0x69, 0x5b, 0xe4, 0xe2, 0x9c, 0xe3, 0x10, 0x4c, 0x0e, 0x09, 0x41, 0x21, 0x87, 0x1c, 0xf5, 0x58,
	0xc2, 0x82, 0xf5, 0x88, 0xbe, 0xd5, 0x21, 0xf9, 0x3d, 0xf9, 0xa1, 0xc1, 0x92, 0x56, 0xb2, 0xac,
	0x04, 0x72, 0xd3, 0x37, 0xda, 0x99, 0xd9, 0x99, 0xdd, 0x05, 0x2b, 0x0c, 0xf6, 0x41, 0x1a, 0x71,
	0x96, 0x17, 0x99, 0xcc, 0x28, 0x85, 0xe9, 0xb6, 0x06, 0x7c, 0xfe, 0x54, 0x72, 0x94, 0x64, 0x0e,
	0x46, 0x2c, 0x62, 0x5b, 0x73, 0x34, 0xd7, 0xf4, 0x0f, 0x9f, 0xf4, 0x06, 0x66, 0xed, 0x1a, 0xcc,
	0xb3, 0x14, 0x39, 0xb1, 0x61, 0xdc, 0xe8, 0x54, 0x0b, 0xbf, 0xf8, 0x6a, 0x24, 0xbf, 0x61, 0x1c,
	0x44, 0x52, 0x64, 0x29, 0xda, 0xba, 0x63, 0xb8, 0x13, 0x6f, 0xcc, 0xce, 0xab, 0xd9, 0x57, 0x38,
	0x0d, 0xc1, 0x6a, 0xf4, 0xee, 0xf3, 0x38, 0x90, 0x7c, 0x68, 0xf9, 0x09, 0x15, 0xf2, 0x13, 0xcc,
	0xb2, 0xa2, 0xc7, 0xdb, 0x67, 0xdb, 0xa8, 0xa8, 0x1d, 0x40, 0x1f, 0xe0, 0x47, 0xcf, 0xe3, 0x78,
	0xe7, 0x09, 0x47, 0x0c, 0x1e, 0x79, 0xe3, 0xa7, 0x46, 0xf2, 0x07, 0x46, 0x35, 0xdf, 0xd6, 0x1d,
	0xcd, 0x9d, 0x78, 0x53, 0xd6, 0x57, 0x68, 0xfe, 0xd2, 0x0d, 0xcc, 0xef, 0xca, 0x10, 0xa3, 0x42,
	0x84, 0x6d, 0x69, 0x14, 0xbe, 0xa1, 0xc2, 0x8a, 0x2b, 0x15, 0xa5, 0x87, 0xd1, 0x0d, 0x8c, 0xea,
	0x0c, 0x64, 0x0a, 0x7a, 0x1b, 0x57, 0x17, 0xf1, 0x21, 0x8a, 0x14, 0x09, 0x47, 0x19, 0x24, 0x79,
	0x65, 0x6e, 0xfa, 0x1d, 0x40, 0xd7, 0xb0, 0xb8, 0x16, 0x7b, 0x8e, 0x32, 0x4b, 0xbb, 0x18, 0xbf,
	0x00, 0x12, 0x05, 0xa2, 0xad, 0x39, 0x86, 0x6b, 0xfa, 0x47, 0x08, 0x65, 0x40, 0x2e, 0xf6, 0x81,
	0x48, 0x6e, 0x0b, 0xf1, 0xd2, 0x0b, 0x8f, 0x65, 0x14, 0x71, 0xc4, 0xca, 0xfd, 0xab, 0xaf, 0x46,
	0xef, 0x55, 0x6f, 0x4f, 0x94, 0xfc, 0x05, 0xd8, 0x71, 0xd9, 0x84, 0x27, 0x33, 0xd6, 0xbf, 0x20,
	0xab, 0x39, 0x3b, 0xbd, 0x0d, 0x67, 0x60, 0xd5, 0x1d, 0x29, 0xce, 0x49, 0x75, 0xab, 0x25, 0x7b,
	0xff, 0x30, 0x2e, 0x61, 0xd9, 0x23, 0xb6, 0xbd, 0x92, 0x05, 0x3b, 0xed, 0xf8, 0x23, 0x91, 0x7f,
	0x1a, 0xf9, 0x0f, 0xd6, 0x8e, 0xcb, 0xb6, 0x24, 0x1c, 0xee, 0x99, 0xb0, 0x61, 0x85, 0x1e, 0x40,
	0x57, 0xd1, 0x90, 0xf2, 0x9d, 0x0d, 0x0b, 0x0c, 0x47, 0xd5, 0xab, 0x59, 0xbf, 0x05, 0x00, 0x00,
	0xff, 0xff, 0xe0, 0xc9, 0x4d, 0x0a, 0x46, 0x03, 0x00, 0x00,
}
