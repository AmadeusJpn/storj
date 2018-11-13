// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: piecestore.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PayerBandwidthAllocation_Action int32

const (
	PayerBandwidthAllocation_PUT PayerBandwidthAllocation_Action = 0
	PayerBandwidthAllocation_GET PayerBandwidthAllocation_Action = 1
)

var PayerBandwidthAllocation_Action_name = map[int32]string{
	0: "PUT",
	1: "GET",
}
var PayerBandwidthAllocation_Action_value = map[string]int32{
	"PUT": 0,
	"GET": 1,
}

func (x PayerBandwidthAllocation_Action) String() string {
	return proto.EnumName(PayerBandwidthAllocation_Action_name, int32(x))
}
func (PayerBandwidthAllocation_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorPiecestore, []int{0, 0}
}

type PayerBandwidthAllocation struct {
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Data      []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PayerBandwidthAllocation) Reset()         { *m = PayerBandwidthAllocation{} }
func (m *PayerBandwidthAllocation) String() string { return proto.CompactTextString(m) }
func (*PayerBandwidthAllocation) ProtoMessage()    {}
func (*PayerBandwidthAllocation) Descriptor() ([]byte, []int) {
	return fileDescriptorPiecestore, []int{0}
}

func (m *PayerBandwidthAllocation) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PayerBandwidthAllocation) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PayerBandwidthAllocation_Data struct {
	SatelliteId       []byte                          `protobuf:"bytes,1,opt,name=satellite_id,json=satelliteId,proto3" json:"satellite_id,omitempty"`
	UplinkId          []byte                          `protobuf:"bytes,2,opt,name=uplink_id,json=uplinkId,proto3" json:"uplink_id,omitempty"`
	MaxSize           int64                           `protobuf:"varint,3,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
	ExpirationUnixSec int64                           `protobuf:"varint,4,opt,name=expiration_unix_sec,json=expirationUnixSec,proto3" json:"expiration_unix_sec,omitempty"`
	SerialNumber      string                          `protobuf:"bytes,5,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	Action            PayerBandwidthAllocation_Action `protobuf:"varint,6,opt,name=action,proto3,enum=piecestoreroutes.PayerBandwidthAllocation_Action" json:"action,omitempty"`
	CreatedUnixSec    int64                           `protobuf:"varint,7,opt,name=created_unix_sec,json=createdUnixSec,proto3" json:"created_unix_sec,omitempty"`
}

func (m *PayerBandwidthAllocation_Data) Reset()         { *m = PayerBandwidthAllocation_Data{} }
func (m *PayerBandwidthAllocation_Data) String() string { return proto.CompactTextString(m) }
func (*PayerBandwidthAllocation_Data) ProtoMessage()    {}
func (*PayerBandwidthAllocation_Data) Descriptor() ([]byte, []int) {
	return fileDescriptorPiecestore, []int{0, 0}
}

func (m *PayerBandwidthAllocation_Data) GetSatelliteId() []byte {
	if m != nil {
		return m.SatelliteId
	}
	return nil
}

func (m *PayerBandwidthAllocation_Data) GetUplinkId() []byte {
	if m != nil {
		return m.UplinkId
	}
	return nil
}

func (m *PayerBandwidthAllocation_Data) GetMaxSize() int64 {
	if m != nil {
		return m.MaxSize
	}
	return 0
}

func (m *PayerBandwidthAllocation_Data) GetExpirationUnixSec() int64 {
	if m != nil {
		return m.ExpirationUnixSec
	}
	return 0
}

func (m *PayerBandwidthAllocation_Data) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *PayerBandwidthAllocation_Data) GetAction() PayerBandwidthAllocation_Action {
	if m != nil {
		return m.Action
	}
	return PayerBandwidthAllocation_PUT
}

func (m *PayerBandwidthAllocation_Data) GetCreatedUnixSec() int64 {
	if m != nil {
		return m.CreatedUnixSec
	}
	return 0
}

type RenterBandwidthAllocation struct {
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Data      []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *RenterBandwidthAllocation) Reset()         { *m = RenterBandwidthAllocation{} }
func (m *RenterBandwidthAllocation) String() string { return proto.CompactTextString(m) }
func (*RenterBandwidthAllocation) ProtoMessage()    {}
func (*RenterBandwidthAllocation) Descriptor() ([]byte, []int) {
	return fileDescriptorPiecestore, []int{1}
}

func (m *RenterBandwidthAllocation) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *RenterBandwidthAllocation) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type RenterBandwidthAllocation_Data struct {
	PayerAllocation *PayerBandwidthAllocation `protobuf:"bytes,1,opt,name=payer_allocation,json=payerAllocation" json:"payer_allocation,omitempty"`
	Total           int64                     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	StorageNodeId   []byte                    `protobuf:"bytes,3,opt,name=storage_node_id,json=storageNodeId,proto3" json:"storage_node_id,omitempty"`
}

func (m *RenterBandwidthAllocation_Data) Reset()         { *m = RenterBandwidthAllocation_Data{} }
func (m *RenterBandwidthAllocation_Data) String() string { return proto.CompactTextString(m) }
func (*RenterBandwidthAllocation_Data) ProtoMessage()    {}
func (*RenterBandwidthAllocation_Data) Descriptor() ([]byte, []int) {
	return fileDescriptorPiecestore, []int{1, 0}
}

func (m *RenterBandwidthAllocation_Data) GetPayerAllocation() *PayerBandwidthAllocation {
	if m != nil {
		return m.PayerAllocation
	}
	return nil
}

func (m *RenterBandwidthAllocation_Data) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RenterBandwidthAllocation_Data) GetStorageNodeId() []byte {
	if m != nil {
		return m.StorageNodeId
	}
	return nil
}

type PieceStore struct {
	Bandwidthallocation *RenterBandwidthAllocation `protobuf:"bytes,1,opt,name=bandwidthallocation" json:"bandwidthallocation,omitempty"`
	Piecedata           *PieceStore_PieceData      `protobuf:"bytes,2,opt,name=piecedata" json:"piecedata,omitempty"`
	Authorization       *SignedMessage             `protobuf:"bytes,3,opt,name=authorization" json:"authorization,omitempty"`
}

func (m *PieceStore) Reset()                    { *m = PieceStore{} }
func (m *PieceStore) String() string            { return proto.CompactTextString(m) }
func (*PieceStore) ProtoMessage()               {}
func (*PieceStore) Descriptor() ([]byte, []int) { return fileDescriptorPiecestore, []int{2} }

func (m *PieceStore) GetBandwidthallocation() *RenterBandwidthAllocation {
	if m != nil {
		return m.Bandwidthallocation
	}
	return nil
}

func (m *PieceStore) GetPiecedata() *PieceStore_PieceData {
	if m != nil {
		return m.Piecedata
	}
	return nil
}

func (m *PieceStore) GetAuthorization() *SignedMessage {
	if m != nil {
		return m.Authorization
	}
	return nil
}

type PieceStore_PieceData struct {
	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExpirationUnixSec int64  `protobuf:"varint,2,opt,name=expiration_unix_sec,json=expirationUnixSec,proto3" json:"expiration_unix_sec,omitempty"`
	Content           []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *PieceStore_PieceData) Reset()         { *m = PieceStore_PieceData{} }
func (m *PieceStore_PieceData) String() string { return proto.CompactTextString(m) }
func (*PieceStore_PieceData) ProtoMessage()    {}
func (*PieceStore_PieceData) Descriptor() ([]byte, []int) {
	return fileDescriptorPiecestore, []int{2, 0}
}

func (m *PieceStore_PieceData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PieceStore_PieceData) GetExpirationUnixSec() int64 {
	if m != nil {
		return m.ExpirationUnixSec
	}
	return 0
}

func (m *PieceStore_PieceData) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type PieceId struct {
	Id            string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Authorization *SignedMessage `protobuf:"bytes,2,opt,name=authorization" json:"authorization,omitempty"`
}

func (m *PieceId) Reset()                    { *m = PieceId{} }
func (m *PieceId) String() string            { return proto.CompactTextString(m) }
func (*PieceId) ProtoMessage()               {}
func (*PieceId) Descriptor() ([]byte, []int) { return fileDescriptorPiecestore, []int{3} }

func (m *PieceId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PieceId) GetAuthorization() *SignedMessage {
	if m != nil {
		return m.Authorization
	}
	return nil
}

type PieceSummary struct {
	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Size_             int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	ExpirationUnixSec int64  `protobuf:"varint,3,opt,name=expiration_unix_sec,json=expirationUnixSec,proto3" json:"expiration_unix_sec,omitempty"`
}

func (m *PieceSummary) Reset()                    { *m = PieceSummary{} }
func (m *PieceSummary) String() string            { return proto.CompactTextString(m) }
func (*PieceSummary) ProtoMessage()               {}
func (*PieceSummary) Descriptor() ([]byte, []int) { return fileDescriptorPiecestore, []int{4} }

func (m *PieceSummary) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PieceSummary) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *PieceSummary) GetExpirationUnixSec() int64 {
	if m != nil {
		return m.ExpirationUnixSec
	}
	return 0
}

type PieceRetrieval struct {
	Bandwidthallocation *RenterBandwidthAllocation `protobuf:"bytes,1,opt,name=bandwidthallocation" json:"bandwidthallocation,omitempty"`
	PieceData           *PieceRetrieval_PieceData  `protobuf:"bytes,2,opt,name=pieceData" json:"pieceData,omitempty"`
	Authorization       *SignedMessage             `protobuf:"bytes,3,opt,name=authorization" json:"authorization,omitempty"`
}

func (m *PieceRetrieval) Reset()                    { *m = PieceRetrieval{} }
func (m *PieceRetrieval) String() string            { return proto.CompactTextString(m) }
func (*PieceRetrieval) ProtoMessage()               {}
func (*PieceRetrieval) Descriptor() ([]byte, []int) { return fileDescriptorPiecestore, []int{5} }

func (m *PieceRetrieval) GetBandwidthallocation() *RenterBandwidthAllocation {
	if m != nil {
		return m.Bandwidthallocation
	}
	return nil
}

func (m *PieceRetrieval) GetPieceData() *PieceRetrieval_PieceData {
	if m != nil {
		return m.PieceData
	}
	return nil
}

func (m *PieceRetrieval) GetAuthorization() *SignedMessage {
	if m != nil {
		return m.Authorization
	}
	return nil
}

type PieceRetrieval_PieceData struct {
	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Size_  int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Offset int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (m *PieceRetrieval_PieceData) Reset()         { *m = PieceRetrieval_PieceData{} }
func (m *PieceRetrieval_PieceData) String() string { return proto.CompactTextString(m) }
func (*PieceRetrieval_PieceData) ProtoMessage()    {}
func (*PieceRetrieval_PieceData) Descriptor() ([]byte, []int) {
	return fileDescriptorPiecestore, []int{5, 0}
}

func (m *PieceRetrieval_PieceData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PieceRetrieval_PieceData) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *PieceRetrieval_PieceData) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type PieceRetrievalStream struct {
	Size_   int64  `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *PieceRetrievalStream) Reset()                    { *m = PieceRetrievalStream{} }
func (m *PieceRetrievalStream) String() string            { return proto.CompactTextString(m) }
func (*PieceRetrievalStream) ProtoMessage()               {}
func (*PieceRetrievalStream) Descriptor() ([]byte, []int) { return fileDescriptorPiecestore, []int{6} }

func (m *PieceRetrievalStream) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *PieceRetrievalStream) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type PieceDelete struct {
	Id            string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Authorization *SignedMessage `protobuf:"bytes,3,opt,name=authorization" json:"authorization,omitempty"`
}

func (m *PieceDelete) Reset()                    { *m = PieceDelete{} }
func (m *PieceDelete) String() string            { return proto.CompactTextString(m) }
func (*PieceDelete) ProtoMessage()               {}
func (*PieceDelete) Descriptor() ([]byte, []int) { return fileDescriptorPiecestore, []int{7} }

func (m *PieceDelete) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PieceDelete) GetAuthorization() *SignedMessage {
	if m != nil {
		return m.Authorization
	}
	return nil
}

type PieceDeleteSummary struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *PieceDeleteSummary) Reset()                    { *m = PieceDeleteSummary{} }
func (m *PieceDeleteSummary) String() string            { return proto.CompactTextString(m) }
func (*PieceDeleteSummary) ProtoMessage()               {}
func (*PieceDeleteSummary) Descriptor() ([]byte, []int) { return fileDescriptorPiecestore, []int{8} }

func (m *PieceDeleteSummary) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PieceStoreSummary struct {
	Message       string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	TotalReceived int64  `protobuf:"varint,2,opt,name=totalReceived,proto3" json:"totalReceived,omitempty"`
}

func (m *PieceStoreSummary) Reset()                    { *m = PieceStoreSummary{} }
func (m *PieceStoreSummary) String() string            { return proto.CompactTextString(m) }
func (*PieceStoreSummary) ProtoMessage()               {}
func (*PieceStoreSummary) Descriptor() ([]byte, []int) { return fileDescriptorPiecestore, []int{9} }

func (m *PieceStoreSummary) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PieceStoreSummary) GetTotalReceived() int64 {
	if m != nil {
		return m.TotalReceived
	}
	return 0
}

type StatsReq struct {
}

func (m *StatsReq) Reset()                    { *m = StatsReq{} }
func (m *StatsReq) String() string            { return proto.CompactTextString(m) }
func (*StatsReq) ProtoMessage()               {}
func (*StatsReq) Descriptor() ([]byte, []int) { return fileDescriptorPiecestore, []int{10} }

type StatSummary struct {
	UsedSpace          int64 `protobuf:"varint,1,opt,name=usedSpace,proto3" json:"usedSpace,omitempty"`
	AvailableSpace     int64 `protobuf:"varint,2,opt,name=availableSpace,proto3" json:"availableSpace,omitempty"`
	UsedBandwidth      int64 `protobuf:"varint,3,opt,name=usedBandwidth,proto3" json:"usedBandwidth,omitempty"`
	AvailableBandwidth int64 `protobuf:"varint,4,opt,name=availableBandwidth,proto3" json:"availableBandwidth,omitempty"`
}

func (m *StatSummary) Reset()                    { *m = StatSummary{} }
func (m *StatSummary) String() string            { return proto.CompactTextString(m) }
func (*StatSummary) ProtoMessage()               {}
func (*StatSummary) Descriptor() ([]byte, []int) { return fileDescriptorPiecestore, []int{11} }

func (m *StatSummary) GetUsedSpace() int64 {
	if m != nil {
		return m.UsedSpace
	}
	return 0
}

func (m *StatSummary) GetAvailableSpace() int64 {
	if m != nil {
		return m.AvailableSpace
	}
	return 0
}

func (m *StatSummary) GetUsedBandwidth() int64 {
	if m != nil {
		return m.UsedBandwidth
	}
	return 0
}

func (m *StatSummary) GetAvailableBandwidth() int64 {
	if m != nil {
		return m.AvailableBandwidth
	}
	return 0
}

type SignedMessage struct {
	Data      []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	PublicKey []byte `protobuf:"bytes,3,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
}

func (m *SignedMessage) Reset()                    { *m = SignedMessage{} }
func (m *SignedMessage) String() string            { return proto.CompactTextString(m) }
func (*SignedMessage) ProtoMessage()               {}
func (*SignedMessage) Descriptor() ([]byte, []int) { return fileDescriptorPiecestore, []int{12} }

func (m *SignedMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SignedMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignedMessage) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func init() {
	proto.RegisterType((*PayerBandwidthAllocation)(nil), "piecestoreroutes.PayerBandwidthAllocation")
	proto.RegisterType((*PayerBandwidthAllocation_Data)(nil), "piecestoreroutes.PayerBandwidthAllocation.Data")
	proto.RegisterType((*RenterBandwidthAllocation)(nil), "piecestoreroutes.RenterBandwidthAllocation")
	proto.RegisterType((*RenterBandwidthAllocation_Data)(nil), "piecestoreroutes.RenterBandwidthAllocation.Data")
	proto.RegisterType((*PieceStore)(nil), "piecestoreroutes.PieceStore")
	proto.RegisterType((*PieceStore_PieceData)(nil), "piecestoreroutes.PieceStore.PieceData")
	proto.RegisterType((*PieceId)(nil), "piecestoreroutes.PieceId")
	proto.RegisterType((*PieceSummary)(nil), "piecestoreroutes.PieceSummary")
	proto.RegisterType((*PieceRetrieval)(nil), "piecestoreroutes.PieceRetrieval")
	proto.RegisterType((*PieceRetrieval_PieceData)(nil), "piecestoreroutes.PieceRetrieval.PieceData")
	proto.RegisterType((*PieceRetrievalStream)(nil), "piecestoreroutes.PieceRetrievalStream")
	proto.RegisterType((*PieceDelete)(nil), "piecestoreroutes.PieceDelete")
	proto.RegisterType((*PieceDeleteSummary)(nil), "piecestoreroutes.PieceDeleteSummary")
	proto.RegisterType((*PieceStoreSummary)(nil), "piecestoreroutes.PieceStoreSummary")
	proto.RegisterType((*StatsReq)(nil), "piecestoreroutes.StatsReq")
	proto.RegisterType((*StatSummary)(nil), "piecestoreroutes.StatSummary")
	proto.RegisterType((*SignedMessage)(nil), "piecestoreroutes.SignedMessage")
	proto.RegisterEnum("piecestoreroutes.PayerBandwidthAllocation_Action", PayerBandwidthAllocation_Action_name, PayerBandwidthAllocation_Action_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PieceStoreRoutes service

type PieceStoreRoutesClient interface {
	Piece(ctx context.Context, in *PieceId, opts ...grpc.CallOption) (*PieceSummary, error)
	Retrieve(ctx context.Context, opts ...grpc.CallOption) (PieceStoreRoutes_RetrieveClient, error)
	Store(ctx context.Context, opts ...grpc.CallOption) (PieceStoreRoutes_StoreClient, error)
	Delete(ctx context.Context, in *PieceDelete, opts ...grpc.CallOption) (*PieceDeleteSummary, error)
	Stats(ctx context.Context, in *StatsReq, opts ...grpc.CallOption) (*StatSummary, error)
}

type pieceStoreRoutesClient struct {
	cc *grpc.ClientConn
}

func NewPieceStoreRoutesClient(cc *grpc.ClientConn) PieceStoreRoutesClient {
	return &pieceStoreRoutesClient{cc}
}

func (c *pieceStoreRoutesClient) Piece(ctx context.Context, in *PieceId, opts ...grpc.CallOption) (*PieceSummary, error) {
	out := new(PieceSummary)
	err := grpc.Invoke(ctx, "/piecestoreroutes.PieceStoreRoutes/Piece", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pieceStoreRoutesClient) Retrieve(ctx context.Context, opts ...grpc.CallOption) (PieceStoreRoutes_RetrieveClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PieceStoreRoutes_serviceDesc.Streams[0], c.cc, "/piecestoreroutes.PieceStoreRoutes/Retrieve", opts...)
	if err != nil {
		return nil, err
	}
	x := &pieceStoreRoutesRetrieveClient{stream}
	return x, nil
}

type PieceStoreRoutes_RetrieveClient interface {
	Send(*PieceRetrieval) error
	Recv() (*PieceRetrievalStream, error)
	grpc.ClientStream
}

type pieceStoreRoutesRetrieveClient struct {
	grpc.ClientStream
}

func (x *pieceStoreRoutesRetrieveClient) Send(m *PieceRetrieval) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pieceStoreRoutesRetrieveClient) Recv() (*PieceRetrievalStream, error) {
	m := new(PieceRetrievalStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pieceStoreRoutesClient) Store(ctx context.Context, opts ...grpc.CallOption) (PieceStoreRoutes_StoreClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PieceStoreRoutes_serviceDesc.Streams[1], c.cc, "/piecestoreroutes.PieceStoreRoutes/Store", opts...)
	if err != nil {
		return nil, err
	}
	x := &pieceStoreRoutesStoreClient{stream}
	return x, nil
}

type PieceStoreRoutes_StoreClient interface {
	Send(*PieceStore) error
	CloseAndRecv() (*PieceStoreSummary, error)
	grpc.ClientStream
}

type pieceStoreRoutesStoreClient struct {
	grpc.ClientStream
}

func (x *pieceStoreRoutesStoreClient) Send(m *PieceStore) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pieceStoreRoutesStoreClient) CloseAndRecv() (*PieceStoreSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PieceStoreSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pieceStoreRoutesClient) Delete(ctx context.Context, in *PieceDelete, opts ...grpc.CallOption) (*PieceDeleteSummary, error) {
	out := new(PieceDeleteSummary)
	err := grpc.Invoke(ctx, "/piecestoreroutes.PieceStoreRoutes/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pieceStoreRoutesClient) Stats(ctx context.Context, in *StatsReq, opts ...grpc.CallOption) (*StatSummary, error) {
	out := new(StatSummary)
	err := grpc.Invoke(ctx, "/piecestoreroutes.PieceStoreRoutes/Stats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PieceStoreRoutes service

type PieceStoreRoutesServer interface {
	Piece(context.Context, *PieceId) (*PieceSummary, error)
	Retrieve(PieceStoreRoutes_RetrieveServer) error
	Store(PieceStoreRoutes_StoreServer) error
	Delete(context.Context, *PieceDelete) (*PieceDeleteSummary, error)
	Stats(context.Context, *StatsReq) (*StatSummary, error)
}

func RegisterPieceStoreRoutesServer(s *grpc.Server, srv PieceStoreRoutesServer) {
	s.RegisterService(&_PieceStoreRoutes_serviceDesc, srv)
}

func _PieceStoreRoutes_Piece_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PieceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PieceStoreRoutesServer).Piece(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/piecestoreroutes.PieceStoreRoutes/Piece",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PieceStoreRoutesServer).Piece(ctx, req.(*PieceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PieceStoreRoutes_Retrieve_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PieceStoreRoutesServer).Retrieve(&pieceStoreRoutesRetrieveServer{stream})
}

type PieceStoreRoutes_RetrieveServer interface {
	Send(*PieceRetrievalStream) error
	Recv() (*PieceRetrieval, error)
	grpc.ServerStream
}

type pieceStoreRoutesRetrieveServer struct {
	grpc.ServerStream
}

func (x *pieceStoreRoutesRetrieveServer) Send(m *PieceRetrievalStream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pieceStoreRoutesRetrieveServer) Recv() (*PieceRetrieval, error) {
	m := new(PieceRetrieval)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PieceStoreRoutes_Store_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PieceStoreRoutesServer).Store(&pieceStoreRoutesStoreServer{stream})
}

type PieceStoreRoutes_StoreServer interface {
	SendAndClose(*PieceStoreSummary) error
	Recv() (*PieceStore, error)
	grpc.ServerStream
}

type pieceStoreRoutesStoreServer struct {
	grpc.ServerStream
}

func (x *pieceStoreRoutesStoreServer) SendAndClose(m *PieceStoreSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pieceStoreRoutesStoreServer) Recv() (*PieceStore, error) {
	m := new(PieceStore)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PieceStoreRoutes_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PieceDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PieceStoreRoutesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/piecestoreroutes.PieceStoreRoutes/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PieceStoreRoutesServer).Delete(ctx, req.(*PieceDelete))
	}
	return interceptor(ctx, in, info, handler)
}

func _PieceStoreRoutes_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PieceStoreRoutesServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/piecestoreroutes.PieceStoreRoutes/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PieceStoreRoutesServer).Stats(ctx, req.(*StatsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PieceStoreRoutes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "piecestoreroutes.PieceStoreRoutes",
	HandlerType: (*PieceStoreRoutesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Piece",
			Handler:    _PieceStoreRoutes_Piece_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PieceStoreRoutes_Delete_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _PieceStoreRoutes_Stats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Retrieve",
			Handler:       _PieceStoreRoutes_Retrieve_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Store",
			Handler:       _PieceStoreRoutes_Store_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "piecestore.proto",
}

func init() { proto.RegisterFile("piecestore.proto", fileDescriptorPiecestore) }

var fileDescriptorPiecestore = []byte{
	// 885 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xe1, 0x6e, 0xdb, 0x36,
	0x10, 0x8e, 0x24, 0xc7, 0x8e, 0xcf, 0xb1, 0xeb, 0xb2, 0xc5, 0xa0, 0x68, 0xe9, 0xe6, 0xa9, 0x45,
	0x60, 0x74, 0x80, 0xb1, 0x65, 0x4f, 0xd0, 0xc2, 0x45, 0x67, 0x0c, 0xcb, 0x02, 0xb9, 0xf9, 0x53,
	0x60, 0xd0, 0x68, 0xf1, 0xea, 0x12, 0x93, 0x25, 0x4d, 0xa2, 0x32, 0x27, 0x0f, 0xb2, 0x3f, 0x7d,
	0x82, 0x61, 0x4f, 0xb2, 0xa7, 0xd8, 0x03, 0xec, 0x25, 0x06, 0x91, 0xb4, 0x64, 0xc7, 0x56, 0x32,
	0x14, 0xe9, 0x3f, 0xf1, 0x8e, 0xf7, 0xdd, 0x77, 0xc7, 0xef, 0x48, 0x41, 0x3f, 0xe1, 0x18, 0x60,
	0x26, 0xe2, 0x14, 0x47, 0x49, 0x1a, 0x8b, 0x98, 0xac, 0x59, 0xd2, 0x38, 0x17, 0x98, 0x39, 0x30,
	0x8f, 0xe7, 0xb1, 0xf2, 0xba, 0x7f, 0x58, 0x60, 0x9f, 0xd3, 0x2b, 0x4c, 0x5f, 0xd2, 0x88, 0xfd,
	0xce, 0x99, 0x78, 0xff, 0x22, 0x0c, 0xe3, 0x80, 0x0a, 0x1e, 0x47, 0xe4, 0x18, 0xda, 0x19, 0x9f,
	0x47, 0x54, 0xe4, 0x29, 0xda, 0xc6, 0xc0, 0x18, 0x1e, 0x7a, 0x95, 0x81, 0x10, 0x68, 0x30, 0x2a,
	0xa8, 0x6d, 0x4a, 0x87, 0xfc, 0x76, 0xfe, 0x32, 0xa1, 0x31, 0xa6, 0x82, 0x92, 0xaf, 0xe0, 0x30,
	0xa3, 0x02, 0xc3, 0x90, 0x0b, 0xf4, 0x39, 0xd3, 0xd1, 0x9d, 0xd2, 0x36, 0x61, 0xe4, 0x73, 0x68,
	0xe7, 0x49, 0xc8, 0xa3, 0x5f, 0x0b, 0xbf, 0x02, 0x39, 0x50, 0x86, 0x09, 0x23, 0x47, 0x70, 0xb0,
	0xa0, 0x4b, 0x3f, 0xe3, 0xd7, 0x68, 0x5b, 0x03, 0x63, 0x68, 0x79, 0xad, 0x05, 0x5d, 0x4e, 0xf9,
	0x35, 0x92, 0x11, 0x3c, 0xc2, 0x65, 0xc2, 0x53, 0xc9, 0xd1, 0xcf, 0x23, 0xbe, 0xf4, 0x33, 0x0c,
	0xec, 0x86, 0xdc, 0xf5, 0xb0, 0x72, 0x5d, 0x44, 0x7c, 0x39, 0xc5, 0x80, 0x3c, 0x85, 0x6e, 0x86,
	0x29, 0xa7, 0xa1, 0x1f, 0xe5, 0x8b, 0x19, 0xa6, 0xf6, 0xfe, 0xc0, 0x18, 0xb6, 0xbd, 0x43, 0x65,
	0x3c, 0x93, 0x36, 0x32, 0x81, 0x26, 0x0d, 0x8a, 0x28, 0xbb, 0x39, 0x30, 0x86, 0xbd, 0xd3, 0x6f,
	0x47, 0x37, 0xdb, 0x36, 0xaa, 0x6b, 0xd3, 0xe8, 0x85, 0x0c, 0xf4, 0x34, 0x00, 0x19, 0x42, 0x3f,
	0x48, 0x91, 0x0a, 0x64, 0x15, 0xb9, 0x96, 0x24, 0xd7, 0xd3, 0x76, 0xcd, 0xcc, 0x75, 0xa0, 0xa9,
	0x62, 0x49, 0x0b, 0xac, 0xf3, 0x8b, 0x37, 0xfd, 0xbd, 0xe2, 0xe3, 0xf5, 0xab, 0x37, 0x7d, 0xc3,
	0xfd, 0xd7, 0x80, 0x23, 0x0f, 0x23, 0x71, 0x5f, 0x27, 0xf3, 0xc1, 0xd0, 0x27, 0x73, 0x01, 0xfd,
	0xa4, 0xa8, 0xc4, 0xa7, 0x25, 0x9c, 0x44, 0xe8, 0x9c, 0x3e, 0xff, 0xff, 0x35, 0x7b, 0x0f, 0x24,
	0xc6, 0x1a, 0xa3, 0xc7, 0xb0, 0x2f, 0x62, 0x41, 0x43, 0x99, 0xd4, 0xf2, 0xd4, 0x82, 0x9c, 0xc0,
	0x83, 0x02, 0x8e, 0xce, 0xd1, 0x8f, 0x62, 0x26, 0x95, 0x60, 0x49, 0x52, 0x5d, 0x6d, 0x3e, 0x8b,
	0x19, 0x4e, 0x98, 0xfb, 0x8f, 0x09, 0x70, 0x5e, 0x24, 0x9f, 0x16, 0xc9, 0xc9, 0xcf, 0xf0, 0x68,
	0xb6, 0x4a, 0xba, 0x45, 0xf3, 0xeb, 0x6d, 0x9a, 0xb5, 0x8d, 0xf2, 0x76, 0xe1, 0x90, 0x31, 0xb4,
	0x25, 0x44, 0xd9, 0xa4, 0xce, 0xe9, 0xc9, 0x8e, 0xda, 0x4b, 0x3e, 0xea, 0xb3, 0xe8, 0x9e, 0x57,
	0x05, 0x92, 0x57, 0xd0, 0xa5, 0xb9, 0x78, 0x1f, 0xa7, 0xfc, 0x5a, 0xd1, 0xb3, 0x24, 0xd2, 0x97,
	0xdb, 0x48, 0x53, 0x3e, 0x8f, 0x90, 0xfd, 0x88, 0x59, 0x46, 0xe7, 0xe8, 0x6d, 0x46, 0x39, 0x08,
	0xed, 0x12, 0x9e, 0xf4, 0xc0, 0xd4, 0xc3, 0xd2, 0xf6, 0x4c, 0xce, 0xea, 0xb4, 0x6e, 0xd6, 0x69,
	0xdd, 0x86, 0x56, 0x10, 0x47, 0x02, 0x23, 0xa1, 0xfb, 0xbc, 0x5a, 0xba, 0xbf, 0x40, 0x4b, 0xa6,
	0x99, 0xb0, 0xad, 0x24, 0x5b, 0x85, 0x98, 0x1f, 0x53, 0x88, 0x3b, 0x83, 0x43, 0xd5, 0xb2, 0x7c,
	0xb1, 0xa0, 0xe9, 0xd5, 0x56, 0x1a, 0x02, 0x0d, 0x39, 0xce, 0x8a, 0xbc, 0xfc, 0xae, 0xab, 0xcf,
	0xaa, 0xa9, 0xcf, 0xfd, 0xdb, 0x84, 0x9e, 0x4c, 0xe2, 0xa1, 0x48, 0x39, 0x5e, 0xd2, 0xf0, 0x53,
	0x6b, 0xe5, 0x7b, 0xad, 0x95, 0x71, 0xa5, 0x95, 0xe7, 0x35, 0x5a, 0x29, 0x39, 0x6d, 0xe9, 0x65,
	0x7c, 0x8f, 0x7a, 0x79, 0x7d, 0x9b, 0x5e, 0x76, 0xf5, 0xf8, 0x33, 0x68, 0xc6, 0xef, 0xde, 0x65,
	0x28, 0x74, 0x5b, 0xf5, 0xca, 0x1d, 0xc3, 0xe3, 0x4d, 0xda, 0x53, 0x91, 0x22, 0x5d, 0x94, 0x18,
	0xc6, 0x1a, 0xc6, 0x9a, 0xae, 0xcc, 0x4d, 0x5d, 0x31, 0xe8, 0x28, 0x3a, 0x18, 0xa2, 0xc0, 0xbb,
	0xb5, 0xf5, 0x51, 0x45, 0xbb, 0x23, 0x20, 0x6b, 0x59, 0x56, 0x0a, 0xb3, 0xa1, 0xb5, 0x50, 0xfb,
	0x75, 0xc6, 0xd5, 0xd2, 0x9d, 0xc2, 0xc3, 0x6a, 0x7c, 0xef, 0xdc, 0x4e, 0x9e, 0x41, 0x57, 0xde,
	0x57, 0x1e, 0x06, 0xc8, 0x2f, 0x91, 0xe9, 0xfe, 0x6d, 0x1a, 0x5d, 0x80, 0x83, 0xa9, 0xa0, 0x22,
	0xf3, 0xf0, 0x37, 0xf7, 0x4f, 0x03, 0x3a, 0xc5, 0x62, 0x85, 0x7d, 0x0c, 0xed, 0x3c, 0x43, 0x36,
	0x4d, 0x68, 0xb0, 0xea, 0x5c, 0x65, 0x20, 0x27, 0xd0, 0xa3, 0x97, 0x94, 0x87, 0x74, 0x16, 0xa2,
	0xda, 0xa2, 0x12, 0xdc, 0xb0, 0x16, 0x3c, 0x8a, 0xa0, 0x52, 0x9c, 0xfa, 0xc4, 0x36, 0x8d, 0x64,
	0x04, 0xa4, 0x8c, 0xab, 0xb6, 0xaa, 0xf7, 0x6f, 0x87, 0xc7, 0xf5, 0xa1, 0xbb, 0xd1, 0xdc, 0xf2,
	0x7d, 0x30, 0xaa, 0xf7, 0x61, 0xf3, 0x45, 0x31, 0x6f, 0xbe, 0x28, 0xc7, 0xd0, 0x4e, 0xf2, 0x59,
	0xc8, 0x83, 0x1f, 0xf0, 0x4a, 0xdf, 0x2c, 0x95, 0xe1, 0xf4, 0x83, 0x05, 0xfd, 0xaa, 0xdd, 0x9e,
	0x3c, 0x4f, 0x32, 0x86, 0x7d, 0x69, 0x23, 0x47, 0x35, 0xe3, 0x32, 0x61, 0xce, 0x17, 0x75, 0xb7,
	0xae, 0xea, 0xaa, 0xbb, 0x47, 0xde, 0xc2, 0x81, 0xd6, 0x27, 0x92, 0xc1, 0x5d, 0x73, 0xe7, 0x9c,
	0xdc, 0xb5, 0x43, 0x49, 0xdc, 0xdd, 0x1b, 0x1a, 0xdf, 0x18, 0xe4, 0x0c, 0xf6, 0xd5, 0x73, 0x73,
	0x7c, 0xdb, 0xe5, 0xef, 0x3c, 0xbd, 0xcd, 0x5b, 0x32, 0x1d, 0x1a, 0xe4, 0x27, 0x68, 0xea, 0x29,
	0x78, 0x52, 0x13, 0xa2, 0xdc, 0xce, 0xb3, 0x5b, 0xdd, 0x55, 0xf1, 0xe3, 0x82, 0x20, 0x15, 0x19,
	0x71, 0x76, 0x8c, 0x8b, 0x56, 0xa2, 0xf3, 0x64, 0xb7, 0xaf, 0x44, 0x79, 0xd9, 0x78, 0x6b, 0x26,
	0xb3, 0x59, 0x53, 0xfe, 0xef, 0x7d, 0xf7, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0x49, 0x5f,
	0x0e, 0x21, 0x0a, 0x00, 0x00,
}
