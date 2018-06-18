// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/burrow/network.proto

package burrow

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

// Params
type PeerParam struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerParam) Reset()         { *m = PeerParam{} }
func (m *PeerParam) String() string { return proto.CompactTextString(m) }
func (*PeerParam) ProtoMessage()    {}
func (*PeerParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_549d0b14318a838d, []int{0}
}
func (m *PeerParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerParam.Unmarshal(m, b)
}
func (m *PeerParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerParam.Marshal(b, m, deterministic)
}
func (dst *PeerParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerParam.Merge(dst, src)
}
func (m *PeerParam) XXX_Size() int {
	return xxx_messageInfo_PeerParam.Size(m)
}
func (m *PeerParam) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerParam.DiscardUnknown(m)
}

var xxx_messageInfo_PeerParam proto.InternalMessageInfo

func (m *PeerParam) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

// Results
type ClientVersion struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientVersion) Reset()         { *m = ClientVersion{} }
func (m *ClientVersion) String() string { return proto.CompactTextString(m) }
func (*ClientVersion) ProtoMessage()    {}
func (*ClientVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_549d0b14318a838d, []int{1}
}
func (m *ClientVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientVersion.Unmarshal(m, b)
}
func (m *ClientVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientVersion.Marshal(b, m, deterministic)
}
func (dst *ClientVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientVersion.Merge(dst, src)
}
func (m *ClientVersion) XXX_Size() int {
	return xxx_messageInfo_ClientVersion.Size(m)
}
func (m *ClientVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientVersion.DiscardUnknown(m)
}

var xxx_messageInfo_ClientVersion proto.InternalMessageInfo

func (m *ClientVersion) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type NodeID struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,2,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeID) Reset()         { *m = NodeID{} }
func (m *NodeID) String() string { return proto.CompactTextString(m) }
func (*NodeID) ProtoMessage()    {}
func (*NodeID) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_549d0b14318a838d, []int{2}
}
func (m *NodeID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeID.Unmarshal(m, b)
}
func (m *NodeID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeID.Marshal(b, m, deterministic)
}
func (dst *NodeID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeID.Merge(dst, src)
}
func (m *NodeID) XXX_Size() int {
	return xxx_messageInfo_NodeID.Size(m)
}
func (m *NodeID) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeID.DiscardUnknown(m)
}

var xxx_messageInfo_NodeID proto.InternalMessageInfo

func (m *NodeID) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeID) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type NodeInfo struct {
	ID                   *NodeID  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ListenAddr           string   `protobuf:"bytes,2,opt,name=ListenAddr,proto3" json:"ListenAddr,omitempty"`
	Network              string   `protobuf:"bytes,3,opt,name=Network,proto3" json:"Network,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=Version,proto3" json:"Version,omitempty"`
	Channels             []byte   `protobuf:"bytes,5,opt,name=Channels,proto3" json:"Channels,omitempty"`
	Moniker              string   `protobuf:"bytes,6,opt,name=Moniker,proto3" json:"Moniker,omitempty"`
	Other                []string `protobuf:"bytes,7,rep,name=Other,proto3" json:"Other,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeInfo) Reset()         { *m = NodeInfo{} }
func (m *NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()    {}
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_549d0b14318a838d, []int{3}
}
func (m *NodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeInfo.Unmarshal(m, b)
}
func (m *NodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeInfo.Marshal(b, m, deterministic)
}
func (dst *NodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeInfo.Merge(dst, src)
}
func (m *NodeInfo) XXX_Size() int {
	return xxx_messageInfo_NodeInfo.Size(m)
}
func (m *NodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeInfo proto.InternalMessageInfo

func (m *NodeInfo) GetID() *NodeID {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *NodeInfo) GetListenAddr() string {
	if m != nil {
		return m.ListenAddr
	}
	return ""
}

func (m *NodeInfo) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *NodeInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *NodeInfo) GetChannels() []byte {
	if m != nil {
		return m.Channels
	}
	return nil
}

func (m *NodeInfo) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *NodeInfo) GetOther() []string {
	if m != nil {
		return m.Other
	}
	return nil
}

type NetworkInfo struct {
	Listening            bool     `protobuf:"varint,1,opt,name=Listening,proto3" json:"Listening,omitempty"`
	Listeners            []string `protobuf:"bytes,2,rep,name=Listeners,proto3" json:"Listeners,omitempty"`
	Peers                []*Peer  `protobuf:"bytes,3,rep,name=Peers,proto3" json:"Peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInfo) Reset()         { *m = NetworkInfo{} }
func (m *NetworkInfo) String() string { return proto.CompactTextString(m) }
func (*NetworkInfo) ProtoMessage()    {}
func (*NetworkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_549d0b14318a838d, []int{4}
}
func (m *NetworkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInfo.Unmarshal(m, b)
}
func (m *NetworkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInfo.Marshal(b, m, deterministic)
}
func (dst *NetworkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInfo.Merge(dst, src)
}
func (m *NetworkInfo) XXX_Size() int {
	return xxx_messageInfo_NetworkInfo.Size(m)
}
func (m *NetworkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInfo proto.InternalMessageInfo

func (m *NetworkInfo) GetListening() bool {
	if m != nil {
		return m.Listening
	}
	return false
}

func (m *NetworkInfo) GetListeners() []string {
	if m != nil {
		return m.Listeners
	}
	return nil
}

func (m *NetworkInfo) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

type Peer struct {
	NodeInfo             *NodeInfo `protobuf:"bytes,1,opt,name=NodeInfo,proto3" json:"NodeInfo,omitempty"`
	IsOutbound           bool      `protobuf:"varint,2,opt,name=IsOutbound,proto3" json:"IsOutbound,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_549d0b14318a838d, []int{5}
}
func (m *Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peer.Unmarshal(m, b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
}
func (dst *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(dst, src)
}
func (m *Peer) XXX_Size() int {
	return xxx_messageInfo_Peer.Size(m)
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetNodeInfo() *NodeInfo {
	if m != nil {
		return m.NodeInfo
	}
	return nil
}

func (m *Peer) GetIsOutbound() bool {
	if m != nil {
		return m.IsOutbound
	}
	return false
}

type PeerList struct {
	Peers                []*Peer  `protobuf:"bytes,1,rep,name=Peers,proto3" json:"Peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerList) Reset()         { *m = PeerList{} }
func (m *PeerList) String() string { return proto.CompactTextString(m) }
func (*PeerList) ProtoMessage()    {}
func (*PeerList) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_549d0b14318a838d, []int{6}
}
func (m *PeerList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerList.Unmarshal(m, b)
}
func (m *PeerList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerList.Marshal(b, m, deterministic)
}
func (dst *PeerList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerList.Merge(dst, src)
}
func (m *PeerList) XXX_Size() int {
	return xxx_messageInfo_PeerList.Size(m)
}
func (m *PeerList) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerList.DiscardUnknown(m)
}

var xxx_messageInfo_PeerList proto.InternalMessageInfo

func (m *PeerList) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func init() {
	proto.RegisterType((*PeerParam)(nil), "PeerParam")
	proto.RegisterType((*ClientVersion)(nil), "ClientVersion")
	proto.RegisterType((*NodeID)(nil), "NodeID")
	proto.RegisterType((*NodeInfo)(nil), "NodeInfo")
	proto.RegisterType((*NetworkInfo)(nil), "NetworkInfo")
	proto.RegisterType((*Peer)(nil), "Peer")
	proto.RegisterType((*PeerList)(nil), "PeerList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkClient is the client API for Network service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkClient interface {
	GetClientVersion(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ClientVersion, error)
	GetNetworkInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NetworkInfo, error)
	GetNodeInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NodeInfo, error)
	GetPeer(ctx context.Context, in *PeerParam, opts ...grpc.CallOption) (*Peer, error)
	GetPeers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PeerList, error)
}

type networkClient struct {
	cc *grpc.ClientConn
}

func NewNetworkClient(cc *grpc.ClientConn) NetworkClient {
	return &networkClient{cc}
}

func (c *networkClient) GetClientVersion(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ClientVersion, error) {
	out := new(ClientVersion)
	err := c.cc.Invoke(ctx, "/Network/GetClientVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) GetNetworkInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NetworkInfo, error) {
	out := new(NetworkInfo)
	err := c.cc.Invoke(ctx, "/Network/GetNetworkInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) GetNodeInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NodeInfo, error) {
	out := new(NodeInfo)
	err := c.cc.Invoke(ctx, "/Network/GetNodeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) GetPeer(ctx context.Context, in *PeerParam, opts ...grpc.CallOption) (*Peer, error) {
	out := new(Peer)
	err := c.cc.Invoke(ctx, "/Network/GetPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) GetPeers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PeerList, error) {
	out := new(PeerList)
	err := c.cc.Invoke(ctx, "/Network/GetPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServer is the server API for Network service.
type NetworkServer interface {
	GetClientVersion(context.Context, *Empty) (*ClientVersion, error)
	GetNetworkInfo(context.Context, *Empty) (*NetworkInfo, error)
	GetNodeInfo(context.Context, *Empty) (*NodeInfo, error)
	GetPeer(context.Context, *PeerParam) (*Peer, error)
	GetPeers(context.Context, *Empty) (*PeerList, error)
}

func RegisterNetworkServer(s *grpc.Server, srv NetworkServer) {
	s.RegisterService(&_Network_serviceDesc, srv)
}

func _Network_GetClientVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).GetClientVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Network/GetClientVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).GetClientVersion(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_GetNetworkInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).GetNetworkInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Network/GetNetworkInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).GetNetworkInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_GetNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).GetNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Network/GetNodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).GetNodeInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_GetPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).GetPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Network/GetPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).GetPeer(ctx, req.(*PeerParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_GetPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).GetPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Network/GetPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).GetPeers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Network_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Network",
	HandlerType: (*NetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClientVersion",
			Handler:    _Network_GetClientVersion_Handler,
		},
		{
			MethodName: "GetNetworkInfo",
			Handler:    _Network_GetNetworkInfo_Handler,
		},
		{
			MethodName: "GetNodeInfo",
			Handler:    _Network_GetNodeInfo_Handler,
		},
		{
			MethodName: "GetPeer",
			Handler:    _Network_GetPeer_Handler,
		},
		{
			MethodName: "GetPeers",
			Handler:    _Network_GetPeers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/burrow/network.proto",
}

func init() { proto.RegisterFile("rpc/burrow/network.proto", fileDescriptor_network_549d0b14318a838d) }

var fileDescriptor_network_549d0b14318a838d = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0xfa, 0x23, 0x4d, 0x5e, 0xc7, 0x84, 0x2c, 0xa4, 0x59, 0xa5, 0x4c, 0x55, 0xa4, 0x41,
	0xe1, 0x90, 0x49, 0xe5, 0xc6, 0x0d, 0x56, 0x54, 0x55, 0xb0, 0xae, 0xf2, 0x81, 0x03, 0xb7, 0xb4,
	0x79, 0x63, 0xd1, 0x1a, 0xbb, 0xb2, 0x5d, 0xa6, 0xfd, 0x6f, 0xdc, 0xf8, 0xc7, 0x90, 0xed, 0xb8,
	0x4e, 0xc5, 0xed, 0x7d, 0xdf, 0xfb, 0xf9, 0x7d, 0x89, 0x81, 0xca, 0xfd, 0xf6, 0x7a, 0x73, 0x90,
	0x52, 0x3c, 0x5d, 0x73, 0xd4, 0x4f, 0x42, 0x3e, 0xe6, 0x7b, 0x29, 0xb4, 0x18, 0x5d, 0xb4, 0x32,
	0x5b, 0x51, 0xd7, 0x82, 0xbb, 0x44, 0x76, 0x05, 0xe9, 0x1a, 0x51, 0xae, 0x0b, 0x59, 0xd4, 0x84,
	0xc2, 0xa0, 0x28, 0x4b, 0x89, 0x4a, 0xd1, 0x68, 0x12, 0x4d, 0xcf, 0x98, 0x87, 0xd9, 0x7b, 0x78,
	0x71, 0xb3, 0xab, 0x90, 0xeb, 0x1f, 0x28, 0x55, 0x25, 0xb8, 0x29, 0xfd, 0xed, 0x42, 0x5b, 0x9a,
	0x32, 0x0f, 0xb3, 0x4f, 0x10, 0xaf, 0x44, 0x89, 0xcb, 0x39, 0x21, 0xd0, 0x5b, 0x15, 0x35, 0x36,
	0x05, 0x36, 0x26, 0x63, 0x48, 0xd7, 0x87, 0xcd, 0xae, 0xda, 0x7e, 0xc3, 0x67, 0xda, 0xb1, 0x4b,
	0x02, 0x91, 0xfd, 0x8d, 0x20, 0xb1, 0xcd, 0xfc, 0x5e, 0x90, 0x0b, 0xe8, 0x2c, 0xe7, 0xb6, 0x79,
	0x38, 0x1b, 0xe4, 0x6e, 0x26, 0xeb, 0x2c, 0xe7, 0xe4, 0x12, 0xe0, 0x7b, 0xa5, 0x34, 0xf2, 0xcf,
	0x65, 0x29, 0xed, 0x90, 0x94, 0xb5, 0x18, 0x73, 0xdb, 0xca, 0xa9, 0xa7, 0x5d, 0x77, 0x5b, 0x03,
	0x4d, 0xa6, 0x11, 0x40, 0x7b, 0x2e, 0xe3, 0xf5, 0x8c, 0x20, 0xb9, 0x79, 0x28, 0x38, 0xc7, 0x9d,
	0xa2, 0x7d, 0x7b, 0xd6, 0x11, 0x9b, 0xae, 0x5b, 0xc1, 0xab, 0x47, 0x94, 0x34, 0x76, 0x5d, 0x0d,
	0x24, 0xaf, 0xa0, 0x7f, 0xa7, 0x1f, 0x50, 0xd2, 0xc1, 0xa4, 0x3b, 0x4d, 0x99, 0x03, 0xd9, 0x3d,
	0x0c, 0x9b, 0x85, 0x56, 0xc7, 0x18, 0x52, 0x77, 0x5c, 0xc5, 0x7f, 0x59, 0x39, 0x09, 0x0b, 0x44,
	0xc8, 0xa2, 0x54, 0xb4, 0x63, 0xc7, 0x04, 0x82, 0xbc, 0x86, 0xbe, 0xf9, 0x3c, 0x8a, 0x76, 0x27,
	0xdd, 0xe9, 0x70, 0xd6, 0xcf, 0x0d, 0x62, 0x8e, 0xcb, 0x6e, 0xa1, 0x67, 0x02, 0x72, 0x15, 0x4c,
	0x6b, 0xec, 0x4a, 0x73, 0x4f, 0xb0, 0xe0, 0xe7, 0x25, 0xc0, 0x52, 0xdd, 0x1d, 0xf4, 0x46, 0x1c,
	0x78, 0x69, 0x6d, 0x4b, 0x58, 0x8b, 0xc9, 0xde, 0x41, 0x62, 0xc6, 0x99, 0xe5, 0x61, 0x6f, 0xf4,
	0xff, 0xde, 0xd9, 0x9f, 0xe8, 0x68, 0x30, 0xf9, 0x00, 0x2f, 0x17, 0xa8, 0x4f, 0xff, 0x8d, 0x38,
	0xff, 0x5a, 0xef, 0xf5, 0xf3, 0xe8, 0x3c, 0x3f, 0xe5, 0xdf, 0xc2, 0xf9, 0x02, 0x75, 0xdb, 0x1a,
	0x5f, 0x79, 0x96, 0xb7, 0xd9, 0x09, 0x0c, 0x4d, 0x9d, 0xbf, 0xdb, 0x17, 0x05, 0x51, 0x64, 0x0c,
	0x83, 0x05, 0x6a, 0x2b, 0x1e, 0xf2, 0xe3, 0xff, 0x3b, 0x72, 0x67, 0x92, 0x37, 0x90, 0x34, 0x59,
	0xd5, 0x6a, 0xf6, 0xda, 0xbe, 0x24, 0x3f, 0x63, 0xf7, 0x12, 0x36, 0xb1, 0x7d, 0x03, 0x1f, 0xff,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x95, 0xeb, 0x57, 0x38, 0x03, 0x00, 0x00,
}
