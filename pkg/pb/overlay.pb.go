// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: overlay.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// NodeTransport is an enum of possible transports for the overlay network
type NodeTransport int32

const (
	NodeTransport_TCP_TLS_GRPC NodeTransport = 0
)

var NodeTransport_name = map[int32]string{
	0: "TCP_TLS_GRPC",
}
var NodeTransport_value = map[string]int32{
	"TCP_TLS_GRPC": 0,
}

func (x NodeTransport) String() string {
	return proto.EnumName(NodeTransport_name, int32(x))
}
func (NodeTransport) EnumDescriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{0} }

// NodeType is an enum of possible node types
type NodeType int32

const (
	NodeType_ADMIN   NodeType = 0
	NodeType_STORAGE NodeType = 1
)

var NodeType_name = map[int32]string{
	0: "ADMIN",
	1: "STORAGE",
}
var NodeType_value = map[string]int32{
	"ADMIN":   0,
	"STORAGE": 1,
}

func (x NodeType) String() string {
	return proto.EnumName(NodeType_name, int32(x))
}
func (NodeType) EnumDescriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{1} }

type Restriction_Operator int32

const (
	Restriction_LT  Restriction_Operator = 0
	Restriction_EQ  Restriction_Operator = 1
	Restriction_GT  Restriction_Operator = 2
	Restriction_LTE Restriction_Operator = 3
	Restriction_GTE Restriction_Operator = 4
)

var Restriction_Operator_name = map[int32]string{
	0: "LT",
	1: "EQ",
	2: "GT",
	3: "LTE",
	4: "GTE",
}
var Restriction_Operator_value = map[string]int32{
	"LT":  0,
	"EQ":  1,
	"GT":  2,
	"LTE": 3,
	"GTE": 4,
}

func (x Restriction_Operator) String() string {
	return proto.EnumName(Restriction_Operator_name, int32(x))
}
func (Restriction_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorOverlay, []int{15, 0}
}

type Restriction_Operand int32

const (
	Restriction_freeBandwidth Restriction_Operand = 0
	Restriction_freeDisk      Restriction_Operand = 1
)

var Restriction_Operand_name = map[int32]string{
	0: "freeBandwidth",
	1: "freeDisk",
}
var Restriction_Operand_value = map[string]int32{
	"freeBandwidth": 0,
	"freeDisk":      1,
}

func (x Restriction_Operand) String() string {
	return proto.EnumName(Restriction_Operand_name, int32(x))
}
func (Restriction_Operand) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorOverlay, []int{15, 1}
}

// LookupRequest is is request message for the lookup rpc call
type LookupRequest struct {
	NodeID string `protobuf:"bytes,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
}

func (m *LookupRequest) Reset()                    { *m = LookupRequest{} }
func (m *LookupRequest) String() string            { return proto.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()               {}
func (*LookupRequest) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{0} }

func (m *LookupRequest) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

// LookupResponse is is response message for the lookup rpc call
type LookupResponse struct {
	Node *Node `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
}

func (m *LookupResponse) Reset()                    { *m = LookupResponse{} }
func (m *LookupResponse) String() string            { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()               {}
func (*LookupResponse) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{1} }

func (m *LookupResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// LookupRequests is a list of LookupRequest
type LookupRequests struct {
	Lookuprequest []*LookupRequest `protobuf:"bytes,1,rep,name=lookuprequest" json:"lookuprequest,omitempty"`
}

func (m *LookupRequests) Reset()                    { *m = LookupRequests{} }
func (m *LookupRequests) String() string            { return proto.CompactTextString(m) }
func (*LookupRequests) ProtoMessage()               {}
func (*LookupRequests) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{2} }

func (m *LookupRequests) GetLookuprequest() []*LookupRequest {
	if m != nil {
		return m.Lookuprequest
	}
	return nil
}

// LookupResponse is a list of LookupResponse
type LookupResponses struct {
	Lookupresponse []*LookupResponse `protobuf:"bytes,1,rep,name=lookupresponse" json:"lookupresponse,omitempty"`
}

func (m *LookupResponses) Reset()                    { *m = LookupResponses{} }
func (m *LookupResponses) String() string            { return proto.CompactTextString(m) }
func (*LookupResponses) ProtoMessage()               {}
func (*LookupResponses) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{3} }

func (m *LookupResponses) GetLookupresponse() []*LookupResponse {
	if m != nil {
		return m.Lookupresponse
	}
	return nil
}

// FindStorageNodesResponse is is response message for the FindStorageNodes rpc call
type FindStorageNodesResponse struct {
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *FindStorageNodesResponse) Reset()                    { *m = FindStorageNodesResponse{} }
func (m *FindStorageNodesResponse) String() string            { return proto.CompactTextString(m) }
func (*FindStorageNodesResponse) ProtoMessage()               {}
func (*FindStorageNodesResponse) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{4} }

func (m *FindStorageNodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// FindStorageNodesRequest is is request message for the FindStorageNodes rpc call
type FindStorageNodesRequest struct {
	ObjectSize     int64                      `protobuf:"varint,1,opt,name=objectSize,proto3" json:"objectSize,omitempty"`
	ContractLength *google_protobuf1.Duration `protobuf:"bytes,2,opt,name=contractLength" json:"contractLength,omitempty"`
	Opts           *OverlayOptions            `protobuf:"bytes,3,opt,name=opts" json:"opts,omitempty"`
	Start          []byte                     `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	MaxNodes       int64                      `protobuf:"varint,5,opt,name=maxNodes,proto3" json:"maxNodes,omitempty"`
}

func (m *FindStorageNodesRequest) Reset()                    { *m = FindStorageNodesRequest{} }
func (m *FindStorageNodesRequest) String() string            { return proto.CompactTextString(m) }
func (*FindStorageNodesRequest) ProtoMessage()               {}
func (*FindStorageNodesRequest) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{5} }

func (m *FindStorageNodesRequest) GetObjectSize() int64 {
	if m != nil {
		return m.ObjectSize
	}
	return 0
}

func (m *FindStorageNodesRequest) GetContractLength() *google_protobuf1.Duration {
	if m != nil {
		return m.ContractLength
	}
	return nil
}

func (m *FindStorageNodesRequest) GetOpts() *OverlayOptions {
	if m != nil {
		return m.Opts
	}
	return nil
}

func (m *FindStorageNodesRequest) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *FindStorageNodesRequest) GetMaxNodes() int64 {
	if m != nil {
		return m.MaxNodes
	}
	return 0
}

// NodeAddress contains the information needed to communicate with a node on the network
type NodeAddress struct {
	Transport NodeTransport `protobuf:"varint,1,opt,name=transport,proto3,enum=overlay.NodeTransport" json:"transport,omitempty"`
	Address   string        `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *NodeAddress) Reset()                    { *m = NodeAddress{} }
func (m *NodeAddress) String() string            { return proto.CompactTextString(m) }
func (*NodeAddress) ProtoMessage()               {}
func (*NodeAddress) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{6} }

func (m *NodeAddress) GetTransport() NodeTransport {
	if m != nil {
		return m.Transport
	}
	return NodeTransport_TCP_TLS_GRPC
}

func (m *NodeAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// OverlayOptions is a set of criteria that a node must meet to be considered for a storage opportunity
type OverlayOptions struct {
	MaxLatency    *google_protobuf1.Duration `protobuf:"bytes,1,opt,name=maxLatency" json:"maxLatency,omitempty"`
	MinReputation *NodeRep                   `protobuf:"bytes,2,opt,name=minReputation" json:"minReputation,omitempty"`
	MinSpeedKbps  int64                      `protobuf:"varint,3,opt,name=minSpeedKbps,proto3" json:"minSpeedKbps,omitempty"`
	Amount        int64                      `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Restrictions  *NodeRestrictions          `protobuf:"bytes,5,opt,name=restrictions" json:"restrictions,omitempty"`
	ExcludedNodes []string                   `protobuf:"bytes,6,rep,name=excluded_nodes,json=excludedNodes" json:"excluded_nodes,omitempty"`
}

func (m *OverlayOptions) Reset()                    { *m = OverlayOptions{} }
func (m *OverlayOptions) String() string            { return proto.CompactTextString(m) }
func (*OverlayOptions) ProtoMessage()               {}
func (*OverlayOptions) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{7} }

func (m *OverlayOptions) GetMaxLatency() *google_protobuf1.Duration {
	if m != nil {
		return m.MaxLatency
	}
	return nil
}

func (m *OverlayOptions) GetMinReputation() *NodeRep {
	if m != nil {
		return m.MinReputation
	}
	return nil
}

func (m *OverlayOptions) GetMinSpeedKbps() int64 {
	if m != nil {
		return m.MinSpeedKbps
	}
	return 0
}

func (m *OverlayOptions) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *OverlayOptions) GetRestrictions() *NodeRestrictions {
	if m != nil {
		return m.Restrictions
	}
	return nil
}

func (m *OverlayOptions) GetExcludedNodes() []string {
	if m != nil {
		return m.ExcludedNodes
	}
	return nil
}

// NodeRep is the reputation characteristics of a node
type NodeRep struct {
	MinUptime       float32 `protobuf:"fixed32,1,opt,name=minUptime,proto3" json:"minUptime,omitempty"`
	MinAuditSuccess float32 `protobuf:"fixed32,2,opt,name=minAuditSuccess,proto3" json:"minAuditSuccess,omitempty"`
	MinAuditCount   int64   `protobuf:"varint,3,opt,name=minAuditCount,proto3" json:"minAuditCount,omitempty"`
}

func (m *NodeRep) Reset()                    { *m = NodeRep{} }
func (m *NodeRep) String() string            { return proto.CompactTextString(m) }
func (*NodeRep) ProtoMessage()               {}
func (*NodeRep) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{8} }

func (m *NodeRep) GetMinUptime() float32 {
	if m != nil {
		return m.MinUptime
	}
	return 0
}

func (m *NodeRep) GetMinAuditSuccess() float32 {
	if m != nil {
		return m.MinAuditSuccess
	}
	return 0
}

func (m *NodeRep) GetMinAuditCount() int64 {
	if m != nil {
		return m.MinAuditCount
	}
	return 0
}

//  NodeRestrictions contains all relevant data about a nodes ability to store data
type NodeRestrictions struct {
	FreeBandwidth int64 `protobuf:"varint,1,opt,name=freeBandwidth,proto3" json:"freeBandwidth,omitempty"`
	FreeDisk      int64 `protobuf:"varint,2,opt,name=freeDisk,proto3" json:"freeDisk,omitempty"`
}

func (m *NodeRestrictions) Reset()                    { *m = NodeRestrictions{} }
func (m *NodeRestrictions) String() string            { return proto.CompactTextString(m) }
func (*NodeRestrictions) ProtoMessage()               {}
func (*NodeRestrictions) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{9} }

func (m *NodeRestrictions) GetFreeBandwidth() int64 {
	if m != nil {
		return m.FreeBandwidth
	}
	return 0
}

func (m *NodeRestrictions) GetFreeDisk() int64 {
	if m != nil {
		return m.FreeDisk
	}
	return 0
}

// Node represents a node in the overlay network
type Node struct {
	Id           string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address      *NodeAddress      `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Type         NodeType          `protobuf:"varint,3,opt,name=type,proto3,enum=overlay.NodeType" json:"type,omitempty"`
	Restrictions *NodeRestrictions `protobuf:"bytes,4,opt,name=restrictions" json:"restrictions,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{10} }

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetAddress() *NodeAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Node) GetType() NodeType {
	if m != nil {
		return m.Type
	}
	return NodeType_ADMIN
}

func (m *Node) GetRestrictions() *NodeRestrictions {
	if m != nil {
		return m.Restrictions
	}
	return nil
}

type QueryRequest struct {
	Sender   *Node `protobuf:"bytes,1,opt,name=sender" json:"sender,omitempty"`
	Target   *Node `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	Limit    int64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Pingback bool  `protobuf:"varint,4,opt,name=pingback,proto3" json:"pingback,omitempty"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{11} }

func (m *QueryRequest) GetSender() *Node {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *QueryRequest) GetTarget() *Node {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *QueryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *QueryRequest) GetPingback() bool {
	if m != nil {
		return m.Pingback
	}
	return false
}

type QueryResponse struct {
	Sender   *Node   `protobuf:"bytes,1,opt,name=sender" json:"sender,omitempty"`
	Response []*Node `protobuf:"bytes,2,rep,name=response" json:"response,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{12} }

func (m *QueryResponse) GetSender() *Node {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *QueryResponse) GetResponse() []*Node {
	if m != nil {
		return m.Response
	}
	return nil
}

type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{13} }

type PingResponse struct {
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{14} }

type Restriction struct {
	Operator Restriction_Operator `protobuf:"varint,1,opt,name=operator,proto3,enum=overlay.Restriction_Operator" json:"operator,omitempty"`
	Operand  Restriction_Operand  `protobuf:"varint,2,opt,name=operand,proto3,enum=overlay.Restriction_Operand" json:"operand,omitempty"`
	Value    int64                `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Restriction) Reset()                    { *m = Restriction{} }
func (m *Restriction) String() string            { return proto.CompactTextString(m) }
func (*Restriction) ProtoMessage()               {}
func (*Restriction) Descriptor() ([]byte, []int) { return fileDescriptorOverlay, []int{15} }

func (m *Restriction) GetOperator() Restriction_Operator {
	if m != nil {
		return m.Operator
	}
	return Restriction_LT
}

func (m *Restriction) GetOperand() Restriction_Operand {
	if m != nil {
		return m.Operand
	}
	return Restriction_freeBandwidth
}

func (m *Restriction) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*LookupRequest)(nil), "overlay.LookupRequest")
	proto.RegisterType((*LookupResponse)(nil), "overlay.LookupResponse")
	proto.RegisterType((*LookupRequests)(nil), "overlay.LookupRequests")
	proto.RegisterType((*LookupResponses)(nil), "overlay.LookupResponses")
	proto.RegisterType((*FindStorageNodesResponse)(nil), "overlay.FindStorageNodesResponse")
	proto.RegisterType((*FindStorageNodesRequest)(nil), "overlay.FindStorageNodesRequest")
	proto.RegisterType((*NodeAddress)(nil), "overlay.NodeAddress")
	proto.RegisterType((*OverlayOptions)(nil), "overlay.OverlayOptions")
	proto.RegisterType((*NodeRep)(nil), "overlay.NodeRep")
	proto.RegisterType((*NodeRestrictions)(nil), "overlay.NodeRestrictions")
	proto.RegisterType((*Node)(nil), "overlay.Node")
	proto.RegisterType((*QueryRequest)(nil), "overlay.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "overlay.QueryResponse")
	proto.RegisterType((*PingRequest)(nil), "overlay.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "overlay.PingResponse")
	proto.RegisterType((*Restriction)(nil), "overlay.Restriction")
	proto.RegisterEnum("overlay.NodeTransport", NodeTransport_name, NodeTransport_value)
	proto.RegisterEnum("overlay.NodeType", NodeType_name, NodeType_value)
	proto.RegisterEnum("overlay.Restriction_Operator", Restriction_Operator_name, Restriction_Operator_value)
	proto.RegisterEnum("overlay.Restriction_Operand", Restriction_Operand_name, Restriction_Operand_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Overlay service

type OverlayClient interface {
	// Lookup finds a nodes address from the network
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error)
	// BulkLookup finds nodes addresses from the network
	BulkLookup(ctx context.Context, in *LookupRequests, opts ...grpc.CallOption) (*LookupResponses, error)
	// FindStorageNodes finds a list of nodes in the network that meet the specified request parameters
	FindStorageNodes(ctx context.Context, in *FindStorageNodesRequest, opts ...grpc.CallOption) (*FindStorageNodesResponse, error)
}

type overlayClient struct {
	cc *grpc.ClientConn
}

func NewOverlayClient(cc *grpc.ClientConn) OverlayClient {
	return &overlayClient{cc}
}

func (c *overlayClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error) {
	out := new(LookupResponse)
	err := grpc.Invoke(ctx, "/overlay.Overlay/Lookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *overlayClient) BulkLookup(ctx context.Context, in *LookupRequests, opts ...grpc.CallOption) (*LookupResponses, error) {
	out := new(LookupResponses)
	err := grpc.Invoke(ctx, "/overlay.Overlay/BulkLookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *overlayClient) FindStorageNodes(ctx context.Context, in *FindStorageNodesRequest, opts ...grpc.CallOption) (*FindStorageNodesResponse, error) {
	out := new(FindStorageNodesResponse)
	err := grpc.Invoke(ctx, "/overlay.Overlay/FindStorageNodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Overlay service

type OverlayServer interface {
	// Lookup finds a nodes address from the network
	Lookup(context.Context, *LookupRequest) (*LookupResponse, error)
	// BulkLookup finds nodes addresses from the network
	BulkLookup(context.Context, *LookupRequests) (*LookupResponses, error)
	// FindStorageNodes finds a list of nodes in the network that meet the specified request parameters
	FindStorageNodes(context.Context, *FindStorageNodesRequest) (*FindStorageNodesResponse, error)
}

func RegisterOverlayServer(s *grpc.Server, srv OverlayServer) {
	s.RegisterService(&_Overlay_serviceDesc, srv)
}

func _Overlay_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Overlay/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayServer).Lookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Overlay_BulkLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayServer).BulkLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Overlay/BulkLookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayServer).BulkLookup(ctx, req.(*LookupRequests))
	}
	return interceptor(ctx, in, info, handler)
}

func _Overlay_FindStorageNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindStorageNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayServer).FindStorageNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Overlay/FindStorageNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayServer).FindStorageNodes(ctx, req.(*FindStorageNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Overlay_serviceDesc = grpc.ServiceDesc{
	ServiceName: "overlay.Overlay",
	HandlerType: (*OverlayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lookup",
			Handler:    _Overlay_Lookup_Handler,
		},
		{
			MethodName: "BulkLookup",
			Handler:    _Overlay_BulkLookup_Handler,
		},
		{
			MethodName: "FindStorageNodes",
			Handler:    _Overlay_FindStorageNodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "overlay.proto",
}

// Client API for Nodes service

type NodesClient interface {
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type nodesClient struct {
	cc *grpc.ClientConn
}

func NewNodesClient(cc *grpc.ClientConn) NodesClient {
	return &nodesClient{cc}
}

func (c *nodesClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := grpc.Invoke(ctx, "/overlay.Nodes/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/overlay.Nodes/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Nodes service

type NodesServer interface {
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterNodesServer(s *grpc.Server, srv NodesServer) {
	s.RegisterService(&_Nodes_serviceDesc, srv)
}

func _Nodes_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Nodes/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nodes_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Nodes/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nodes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "overlay.Nodes",
	HandlerType: (*NodesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Nodes_Query_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Nodes_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "overlay.proto",
}

func init() { proto.RegisterFile("overlay.proto", fileDescriptorOverlay) }

var fileDescriptorOverlay = []byte{
	// 997 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x6d, 0x6f, 0xe3, 0x44,
	0x10, 0xae, 0x9d, 0xf7, 0x49, 0xe2, 0xf3, 0xad, 0xee, 0xda, 0x10, 0x1d, 0xa7, 0xd4, 0x50, 0x11,
	0x8a, 0x94, 0x93, 0xd2, 0x53, 0xa5, 0x4a, 0xa0, 0x2a, 0x7d, 0xa1, 0x3a, 0x11, 0xda, 0xde, 0x26,
	0x08, 0x09, 0x09, 0x55, 0x4e, 0xbc, 0x97, 0x33, 0x4d, 0x76, 0x8d, 0xbd, 0xbe, 0x6b, 0xf8, 0x11,
	0xfc, 0x0b, 0xbe, 0xf1, 0x8f, 0x90, 0xf8, 0x1d, 0x7c, 0x44, 0xfb, 0x62, 0x27, 0x76, 0x9b, 0x0a,
	0x3e, 0x79, 0x67, 0xe6, 0x99, 0xf1, 0xcc, 0x33, 0xb3, 0xb3, 0xd0, 0x64, 0x1f, 0x48, 0x38, 0x77,
	0x97, 0xbd, 0x20, 0x64, 0x9c, 0xa1, 0x8a, 0x16, 0xdb, 0x30, 0x63, 0x33, 0xa6, 0x94, 0xed, 0x97,
	0x33, 0xc6, 0x66, 0x73, 0xf2, 0x4a, 0x4a, 0x93, 0xf8, 0xdd, 0x2b, 0x2f, 0x0e, 0x5d, 0xee, 0x33,
	0xaa, 0xec, 0xce, 0x17, 0xd0, 0x1c, 0x32, 0x76, 0x1b, 0x07, 0x98, 0xfc, 0x1a, 0x93, 0x88, 0xa3,
	0x6d, 0x28, 0x53, 0xe6, 0x91, 0x37, 0x67, 0x2d, 0xa3, 0x63, 0x74, 0x6b, 0x58, 0x4b, 0xce, 0x01,
	0x58, 0x09, 0x30, 0x0a, 0x18, 0x8d, 0x08, 0xda, 0x85, 0xa2, 0xb0, 0x49, 0x5c, 0xbd, 0xdf, 0xec,
	0x25, 0xd9, 0x5c, 0x32, 0x8f, 0x60, 0x69, 0x72, 0x2e, 0x57, 0x4e, 0x32, 0x7a, 0x84, 0xbe, 0x86,
	0xe6, 0x5c, 0x6a, 0x42, 0xa5, 0x69, 0x19, 0x9d, 0x42, 0xb7, 0xde, 0xdf, 0x4e, 0xbd, 0x33, 0x78,
	0x9c, 0x05, 0x3b, 0x18, 0x9e, 0x64, 0x93, 0x88, 0xd0, 0x31, 0x58, 0x09, 0x46, 0xa9, 0x74, 0xc4,
	0x9d, 0x7b, 0x11, 0x95, 0x19, 0xe7, 0xe0, 0xce, 0x31, 0xb4, 0xbe, 0xf5, 0xa9, 0x37, 0xe2, 0x2c,
	0x74, 0x67, 0x44, 0x24, 0x1f, 0xa5, 0x25, 0x7e, 0x06, 0x25, 0x51, 0x47, 0xa4, 0x63, 0xe6, 0x6a,
	0x54, 0x36, 0xe7, 0x2f, 0x03, 0x76, 0xee, 0x47, 0x50, 0x6c, 0xbe, 0x04, 0x60, 0x93, 0x5f, 0xc8,
	0x94, 0x8f, 0xfc, 0xdf, 0x14, 0x53, 0x05, 0xbc, 0xa6, 0x41, 0x03, 0xb0, 0xa6, 0x8c, 0xf2, 0xd0,
	0x9d, 0xf2, 0x21, 0xa1, 0x33, 0xfe, 0xbe, 0x65, 0x4a, 0x36, 0x3f, 0xe9, 0xa9, 0xbe, 0xf5, 0x92,
	0xbe, 0xf5, 0xce, 0x74, 0xdf, 0x70, 0xce, 0x01, 0x7d, 0x05, 0x45, 0x16, 0xf0, 0xa8, 0x55, 0x90,
	0x8e, 0xab, 0xb2, 0xaf, 0xd4, 0xf7, 0x2a, 0x10, 0x5e, 0x11, 0x96, 0x20, 0xf4, 0x0c, 0x4a, 0x11,
	0x77, 0x43, 0xde, 0x2a, 0x76, 0x8c, 0x6e, 0x03, 0x2b, 0x01, 0xb5, 0xa1, 0xba, 0x70, 0xef, 0x64,
	0xe2, 0xad, 0x92, 0xcc, 0x31, 0x95, 0x9d, 0x9f, 0xa1, 0x2e, 0x0e, 0x03, 0xcf, 0x0b, 0x49, 0x14,
	0xa1, 0xd7, 0x50, 0xe3, 0xa1, 0x4b, 0xa3, 0x80, 0x85, 0x5c, 0xd6, 0x63, 0xad, 0xf5, 0x4e, 0x00,
	0xc7, 0x89, 0x15, 0xaf, 0x80, 0xa8, 0x05, 0x15, 0x57, 0x05, 0x90, 0xf5, 0xd5, 0x70, 0x22, 0x3a,
	0x7f, 0x98, 0x60, 0x65, 0x33, 0x45, 0x47, 0x00, 0x0b, 0xf7, 0x6e, 0xe8, 0x72, 0x42, 0xa7, 0x4b,
	0x3d, 0x5d, 0x8f, 0xf0, 0xb1, 0x06, 0x46, 0x87, 0xd0, 0x5c, 0xf8, 0x14, 0x93, 0x20, 0xe6, 0xd2,
	0xa8, 0xd9, 0xb4, 0xb3, 0x7d, 0x23, 0x01, 0xce, 0xc2, 0x90, 0x03, 0x8d, 0x85, 0x4f, 0x47, 0x01,
	0x21, 0xde, 0x77, 0x93, 0x40, 0x71, 0x59, 0xc0, 0x19, 0x9d, 0xb8, 0x18, 0xee, 0x82, 0xc5, 0x54,
	0x71, 0x57, 0xc0, 0x5a, 0x42, 0xdf, 0x40, 0x23, 0x24, 0x11, 0x0f, 0xfd, 0xa9, 0x4c, 0x5f, 0x12,
	0x28, 0x12, 0xce, 0xfe, 0x72, 0x05, 0xc0, 0x19, 0x38, 0xda, 0x03, 0x8b, 0xdc, 0x4d, 0xe7, 0xb1,
	0x47, 0xbc, 0x1b, 0x35, 0x6b, 0xe5, 0x4e, 0xa1, 0x5b, 0xc3, 0xcd, 0x44, 0xab, 0xda, 0xf0, 0x11,
	0x2a, 0x3a, 0x77, 0xf4, 0x02, 0x6a, 0x0b, 0x9f, 0xfe, 0x10, 0x70, 0x7f, 0xa1, 0x46, 0xca, 0xc4,
	0x2b, 0x05, 0xea, 0xc2, 0x93, 0x85, 0x4f, 0x07, 0xb1, 0xe7, 0xf3, 0x51, 0x3c, 0x9d, 0x26, 0x94,
	0x9b, 0x38, 0xaf, 0x46, 0x9f, 0x4b, 0xb2, 0xa4, 0xea, 0x54, 0xd6, 0xa5, 0xaa, 0xce, 0x2a, 0x9d,
	0x31, 0xd8, 0xf9, 0x0a, 0x84, 0xe7, 0xbb, 0x90, 0x90, 0x13, 0x97, 0x7a, 0x1f, 0x7d, 0x8f, 0xbf,
	0xd7, 0x83, 0x9d, 0x55, 0x8a, 0xa9, 0x12, 0x8a, 0x33, 0x3f, 0xba, 0x95, 0x29, 0x14, 0x70, 0x2a,
	0x3b, 0x7f, 0x1a, 0x50, 0x14, 0x61, 0x91, 0x05, 0xa6, 0xef, 0xe9, 0x55, 0x63, 0xfa, 0x1e, 0xea,
	0x65, 0x27, 0xa5, 0xde, 0x7f, 0x96, 0x21, 0x52, 0x8f, 0x61, 0x3a, 0x3f, 0x68, 0x0f, 0x8a, 0x7c,
	0x19, 0x10, 0x99, 0xbb, 0xd5, 0x7f, 0x9a, 0x1d, 0xc5, 0x65, 0x40, 0xb0, 0x34, 0xdf, 0x6b, 0x52,
	0xf1, 0x7f, 0x35, 0xc9, 0xf9, 0xdd, 0x80, 0xc6, 0xdb, 0x98, 0x84, 0xcb, 0xe4, 0x5e, 0xef, 0x41,
	0x39, 0x22, 0xd4, 0x23, 0xe1, 0xc3, 0xdb, 0x4f, 0x1b, 0x05, 0x8c, 0xbb, 0xe1, 0x8c, 0x70, 0x5d,
	0x4c, 0x1e, 0xa6, 0x8c, 0xe2, 0x56, 0xce, 0xfd, 0x85, 0x9f, 0x74, 0x40, 0x09, 0x82, 0xbf, 0xc0,
	0xa7, 0xb3, 0x89, 0x3b, 0xbd, 0x95, 0xf9, 0x56, 0x71, 0x2a, 0x3b, 0x2e, 0x34, 0x75, 0x3e, 0x7a,
	0x53, 0xfd, 0xc7, 0x84, 0xbe, 0x84, 0x6a, 0xba, 0x27, 0xcd, 0x87, 0x76, 0x5a, 0x6a, 0x76, 0x9a,
	0x50, 0xbf, 0xf6, 0xe9, 0x4c, 0x57, 0xec, 0x58, 0xd0, 0x50, 0xa2, 0x36, 0xff, 0x63, 0x40, 0x7d,
	0x8d, 0x31, 0x74, 0x04, 0x55, 0x16, 0x90, 0xd0, 0xe5, 0x2c, 0xd4, 0x7b, 0xe1, 0xd3, 0x34, 0xf2,
	0x1a, 0xae, 0x77, 0xa5, 0x41, 0x38, 0x85, 0xa3, 0x43, 0xa8, 0xc8, 0x33, 0xf5, 0x24, 0x4d, 0x56,
	0xff, 0xc5, 0x66, 0x4f, 0xea, 0xe1, 0x04, 0x2c, 0x68, 0xfb, 0xe0, 0xce, 0x63, 0x92, 0xd0, 0x26,
	0x05, 0xe7, 0x35, 0x54, 0x93, 0x7f, 0xa0, 0x32, 0x98, 0xc3, 0xb1, 0xbd, 0x25, 0xbe, 0xe7, 0x6f,
	0x6d, 0x43, 0x7c, 0x2f, 0xc6, 0xb6, 0x89, 0x2a, 0x50, 0x18, 0x8e, 0xcf, 0xed, 0x82, 0x38, 0x5c,
	0x8c, 0xcf, 0xed, 0xa2, 0xb3, 0x0f, 0x15, 0x1d, 0x1f, 0x3d, 0xcd, 0x4d, 0xb7, 0xbd, 0x85, 0x1a,
	0xab, 0x51, 0xb6, 0x8d, 0xfd, 0x5d, 0x68, 0x66, 0x36, 0x1d, 0xb2, 0xa1, 0x31, 0x3e, 0xbd, 0xbe,
	0x19, 0x0f, 0x47, 0x37, 0x17, 0xf8, 0xfa, 0xd4, 0xde, 0xda, 0x77, 0xa0, 0x9a, 0x4c, 0x20, 0xaa,
	0x41, 0x69, 0x70, 0xf6, 0xfd, 0x9b, 0x4b, 0x7b, 0x0b, 0xd5, 0xa1, 0x32, 0x1a, 0x5f, 0xe1, 0xc1,
	0xc5, 0xb9, 0x6d, 0xf4, 0xff, 0x36, 0xa0, 0xa2, 0x57, 0x1f, 0x3a, 0x82, 0xb2, 0x7a, 0xa6, 0xd0,
	0x86, 0x97, 0xb0, 0xbd, 0xe9, 0x3d, 0x43, 0xc7, 0x00, 0x27, 0xf1, 0xfc, 0x56, 0xbb, 0xef, 0x3c,
	0xec, 0x1e, 0xb5, 0x5b, 0x1b, 0xfc, 0x23, 0xf4, 0x23, 0xd8, 0xf9, 0xe7, 0x0b, 0x75, 0x52, 0xf4,
	0x86, 0x97, 0xad, 0xbd, 0xfb, 0x08, 0x42, 0x45, 0xee, 0x73, 0x28, 0xa9, 0x68, 0x87, 0x50, 0x92,
	0xd3, 0x8a, 0x9e, 0xa7, 0x4e, 0xeb, 0xb7, 0xa9, 0xbd, 0x9d, 0x57, 0xeb, 0xd2, 0x0e, 0xa0, 0x28,
	0x66, 0x0e, 0xad, 0x76, 0xc0, 0xda, 0x44, 0xb6, 0x9f, 0xe7, 0xb4, 0xca, 0xe9, 0xa4, 0xf8, 0x93,
	0x19, 0x4c, 0x26, 0x65, 0xf9, 0x50, 0x1c, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x30, 0xe2, 0x43,
	0x53, 0x24, 0x09, 0x00, 0x00,
}
