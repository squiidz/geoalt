// Code generated by protoc-gen-go. DO NOT EDIT.
// source: geoaltsvc.proto

package geoaltsvc

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

type RegisterReq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_geoaltsvc_5e512dd99f0f8398, []int{0}
}
func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (dst *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(dst, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterReq) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *RegisterReq) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *RegisterReq) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type RegisterResp struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResp) Reset()         { *m = RegisterResp{} }
func (m *RegisterResp) String() string { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()    {}
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_geoaltsvc_5e512dd99f0f8398, []int{1}
}
func (m *RegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResp.Unmarshal(m, b)
}
func (m *RegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResp.Marshal(b, m, deterministic)
}
func (dst *RegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResp.Merge(dst, src)
}
func (m *RegisterResp) XXX_Size() int {
	return xxx_messageInfo_RegisterResp.Size(m)
}
func (m *RegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResp proto.InternalMessageInfo

func (m *RegisterResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LoginReq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_geoaltsvc_5e512dd99f0f8398, []int{2}
}
func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (dst *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(dst, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResp struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_geoaltsvc_5e512dd99f0f8398, []int{3}
}
func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResp.Unmarshal(m, b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
}
func (dst *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(dst, src)
}
func (m *LoginResp) XXX_Size() int {
	return xxx_messageInfo_LoginResp.Size(m)
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func (m *LoginResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AddAlertReq struct {
	Lat                  float64  `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng                  float64  `protobuf:"fixed64,2,opt,name=lng,proto3" json:"lng,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Ephemeral            bool     `protobuf:"varint,4,opt,name=ephemeral,proto3" json:"ephemeral,omitempty"`
	Resolution           uint32   `protobuf:"varint,5,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Delay                int64    `protobuf:"varint,6,opt,name=delay,proto3" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddAlertReq) Reset()         { *m = AddAlertReq{} }
func (m *AddAlertReq) String() string { return proto.CompactTextString(m) }
func (*AddAlertReq) ProtoMessage()    {}
func (*AddAlertReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_geoaltsvc_5e512dd99f0f8398, []int{4}
}
func (m *AddAlertReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAlertReq.Unmarshal(m, b)
}
func (m *AddAlertReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAlertReq.Marshal(b, m, deterministic)
}
func (dst *AddAlertReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAlertReq.Merge(dst, src)
}
func (m *AddAlertReq) XXX_Size() int {
	return xxx_messageInfo_AddAlertReq.Size(m)
}
func (m *AddAlertReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAlertReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddAlertReq proto.InternalMessageInfo

func (m *AddAlertReq) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *AddAlertReq) GetLng() float64 {
	if m != nil {
		return m.Lng
	}
	return 0
}

func (m *AddAlertReq) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AddAlertReq) GetEphemeral() bool {
	if m != nil {
		return m.Ephemeral
	}
	return false
}

func (m *AddAlertReq) GetResolution() uint32 {
	if m != nil {
		return m.Resolution
	}
	return 0
}

func (m *AddAlertReq) GetDelay() int64 {
	if m != nil {
		return m.Delay
	}
	return 0
}

type AddAlertResp struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddAlertResp) Reset()         { *m = AddAlertResp{} }
func (m *AddAlertResp) String() string { return proto.CompactTextString(m) }
func (*AddAlertResp) ProtoMessage()    {}
func (*AddAlertResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_geoaltsvc_5e512dd99f0f8398, []int{5}
}
func (m *AddAlertResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAlertResp.Unmarshal(m, b)
}
func (m *AddAlertResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAlertResp.Marshal(b, m, deterministic)
}
func (dst *AddAlertResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAlertResp.Merge(dst, src)
}
func (m *AddAlertResp) XXX_Size() int {
	return xxx_messageInfo_AddAlertResp.Size(m)
}
func (m *AddAlertResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAlertResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddAlertResp proto.InternalMessageInfo

func (m *AddAlertResp) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type GetAlertsReq struct {
	Lat                  float64  `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng                  float64  `protobuf:"fixed64,2,opt,name=lng,proto3" json:"lng,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAlertsReq) Reset()         { *m = GetAlertsReq{} }
func (m *GetAlertsReq) String() string { return proto.CompactTextString(m) }
func (*GetAlertsReq) ProtoMessage()    {}
func (*GetAlertsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_geoaltsvc_5e512dd99f0f8398, []int{6}
}
func (m *GetAlertsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAlertsReq.Unmarshal(m, b)
}
func (m *GetAlertsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAlertsReq.Marshal(b, m, deterministic)
}
func (dst *GetAlertsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAlertsReq.Merge(dst, src)
}
func (m *GetAlertsReq) XXX_Size() int {
	return xxx_messageInfo_GetAlertsReq.Size(m)
}
func (m *GetAlertsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAlertsReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetAlertsReq proto.InternalMessageInfo

func (m *GetAlertsReq) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *GetAlertsReq) GetLng() float64 {
	if m != nil {
		return m.Lng
	}
	return 0
}

type GetAlertsResp struct {
	Alerts               []*Alert `protobuf:"bytes,1,rep,name=alerts,proto3" json:"alerts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAlertsResp) Reset()         { *m = GetAlertsResp{} }
func (m *GetAlertsResp) String() string { return proto.CompactTextString(m) }
func (*GetAlertsResp) ProtoMessage()    {}
func (*GetAlertsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_geoaltsvc_5e512dd99f0f8398, []int{7}
}
func (m *GetAlertsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAlertsResp.Unmarshal(m, b)
}
func (m *GetAlertsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAlertsResp.Marshal(b, m, deterministic)
}
func (dst *GetAlertsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAlertsResp.Merge(dst, src)
}
func (m *GetAlertsResp) XXX_Size() int {
	return xxx_messageInfo_GetAlertsResp.Size(m)
}
func (m *GetAlertsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAlertsResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetAlertsResp proto.InternalMessageInfo

func (m *GetAlertsResp) GetAlerts() []*Alert {
	if m != nil {
		return m.Alerts
	}
	return nil
}

type Alert struct {
	Center               *Coord   `protobuf:"bytes,1,opt,name=center,proto3" json:"center,omitempty"`
	Borders              []*Coord `protobuf:"bytes,2,rep,name=borders,proto3" json:"borders,omitempty"`
	Cell                 *Cell    `protobuf:"bytes,3,opt,name=cell,proto3" json:"cell,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp            int64    `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ReadAt               int64    `protobuf:"varint,6,opt,name=read_at,json=readAt,proto3" json:"read_at,omitempty"`
	Delay                int64    `protobuf:"varint,7,opt,name=delay,proto3" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Alert) Reset()         { *m = Alert{} }
func (m *Alert) String() string { return proto.CompactTextString(m) }
func (*Alert) ProtoMessage()    {}
func (*Alert) Descriptor() ([]byte, []int) {
	return fileDescriptor_geoaltsvc_5e512dd99f0f8398, []int{8}
}
func (m *Alert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alert.Unmarshal(m, b)
}
func (m *Alert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alert.Marshal(b, m, deterministic)
}
func (dst *Alert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alert.Merge(dst, src)
}
func (m *Alert) XXX_Size() int {
	return xxx_messageInfo_Alert.Size(m)
}
func (m *Alert) XXX_DiscardUnknown() {
	xxx_messageInfo_Alert.DiscardUnknown(m)
}

var xxx_messageInfo_Alert proto.InternalMessageInfo

func (m *Alert) GetCenter() *Coord {
	if m != nil {
		return m.Center
	}
	return nil
}

func (m *Alert) GetBorders() []*Coord {
	if m != nil {
		return m.Borders
	}
	return nil
}

func (m *Alert) GetCell() *Cell {
	if m != nil {
		return m.Cell
	}
	return nil
}

func (m *Alert) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Alert) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Alert) GetReadAt() int64 {
	if m != nil {
		return m.ReadAt
	}
	return 0
}

func (m *Alert) GetDelay() int64 {
	if m != nil {
		return m.Delay
	}
	return 0
}

type Coord struct {
	Lat                  float64  `protobuf:"fixed64,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng                  float64  `protobuf:"fixed64,3,opt,name=lng,proto3" json:"lng,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coord) Reset()         { *m = Coord{} }
func (m *Coord) String() string { return proto.CompactTextString(m) }
func (*Coord) ProtoMessage()    {}
func (*Coord) Descriptor() ([]byte, []int) {
	return fileDescriptor_geoaltsvc_5e512dd99f0f8398, []int{9}
}
func (m *Coord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coord.Unmarshal(m, b)
}
func (m *Coord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coord.Marshal(b, m, deterministic)
}
func (dst *Coord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coord.Merge(dst, src)
}
func (m *Coord) XXX_Size() int {
	return xxx_messageInfo_Coord.Size(m)
}
func (m *Coord) XXX_DiscardUnknown() {
	xxx_messageInfo_Coord.DiscardUnknown(m)
}

var xxx_messageInfo_Coord proto.InternalMessageInfo

func (m *Coord) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Coord) GetLng() float64 {
	if m != nil {
		return m.Lng
	}
	return 0
}

type Cell struct {
	// Smallest cell resolution (15)
	BaseCell uint64 `protobuf:"varint,1,opt,name=base_cell,json=baseCell,proto3" json:"base_cell,omitempty"`
	// Cell id used for indexing
	IndexCell uint64 `protobuf:"varint,2,opt,name=index_cell,json=indexCell,proto3" json:"index_cell,omitempty"`
	// Cell id with the resolution
	RealCell             uint64   `protobuf:"varint,3,opt,name=real_cell,json=realCell,proto3" json:"real_cell,omitempty"`
	Resolution           uint32   `protobuf:"varint,4,opt,name=resolution,proto3" json:"resolution,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cell) Reset()         { *m = Cell{} }
func (m *Cell) String() string { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()    {}
func (*Cell) Descriptor() ([]byte, []int) {
	return fileDescriptor_geoaltsvc_5e512dd99f0f8398, []int{10}
}
func (m *Cell) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cell.Unmarshal(m, b)
}
func (m *Cell) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cell.Marshal(b, m, deterministic)
}
func (dst *Cell) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cell.Merge(dst, src)
}
func (m *Cell) XXX_Size() int {
	return xxx_messageInfo_Cell.Size(m)
}
func (m *Cell) XXX_DiscardUnknown() {
	xxx_messageInfo_Cell.DiscardUnknown(m)
}

var xxx_messageInfo_Cell proto.InternalMessageInfo

func (m *Cell) GetBaseCell() uint64 {
	if m != nil {
		return m.BaseCell
	}
	return 0
}

func (m *Cell) GetIndexCell() uint64 {
	if m != nil {
		return m.IndexCell
	}
	return 0
}

func (m *Cell) GetRealCell() uint64 {
	if m != nil {
		return m.RealCell
	}
	return 0
}

func (m *Cell) GetResolution() uint32 {
	if m != nil {
		return m.Resolution
	}
	return 0
}

func init() {
	proto.RegisterType((*RegisterReq)(nil), "RegisterReq")
	proto.RegisterType((*RegisterResp)(nil), "RegisterResp")
	proto.RegisterType((*LoginReq)(nil), "LoginReq")
	proto.RegisterType((*LoginResp)(nil), "LoginResp")
	proto.RegisterType((*AddAlertReq)(nil), "AddAlertReq")
	proto.RegisterType((*AddAlertResp)(nil), "AddAlertResp")
	proto.RegisterType((*GetAlertsReq)(nil), "GetAlertsReq")
	proto.RegisterType((*GetAlertsResp)(nil), "GetAlertsResp")
	proto.RegisterType((*Alert)(nil), "Alert")
	proto.RegisterType((*Coord)(nil), "Coord")
	proto.RegisterType((*Cell)(nil), "Cell")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GeoAltClient is the client API for GeoAlt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeoAltClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	GeoFeed(ctx context.Context, opts ...grpc.CallOption) (GeoAlt_GeoFeedClient, error)
	AddAlert(ctx context.Context, in *AddAlertReq, opts ...grpc.CallOption) (*AddAlertResp, error)
	GetAlerts(ctx context.Context, in *GetAlertsReq, opts ...grpc.CallOption) (*GetAlertsResp, error)
	GetActiveAlerts(ctx context.Context, in *GetAlertsReq, opts ...grpc.CallOption) (*GetAlertsResp, error)
}

type geoAltClient struct {
	cc *grpc.ClientConn
}

func NewGeoAltClient(cc *grpc.ClientConn) GeoAltClient {
	return &geoAltClient{cc}
}

func (c *geoAltClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/GeoAlt/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoAltClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, "/GeoAlt/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoAltClient) GeoFeed(ctx context.Context, opts ...grpc.CallOption) (GeoAlt_GeoFeedClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GeoAlt_serviceDesc.Streams[0], "/GeoAlt/GeoFeed", opts...)
	if err != nil {
		return nil, err
	}
	x := &geoAltGeoFeedClient{stream}
	return x, nil
}

type GeoAlt_GeoFeedClient interface {
	Send(*GetAlertsReq) error
	Recv() (*GetAlertsResp, error)
	grpc.ClientStream
}

type geoAltGeoFeedClient struct {
	grpc.ClientStream
}

func (x *geoAltGeoFeedClient) Send(m *GetAlertsReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *geoAltGeoFeedClient) Recv() (*GetAlertsResp, error) {
	m := new(GetAlertsResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *geoAltClient) AddAlert(ctx context.Context, in *AddAlertReq, opts ...grpc.CallOption) (*AddAlertResp, error) {
	out := new(AddAlertResp)
	err := c.cc.Invoke(ctx, "/GeoAlt/AddAlert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoAltClient) GetAlerts(ctx context.Context, in *GetAlertsReq, opts ...grpc.CallOption) (*GetAlertsResp, error) {
	out := new(GetAlertsResp)
	err := c.cc.Invoke(ctx, "/GeoAlt/GetAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoAltClient) GetActiveAlerts(ctx context.Context, in *GetAlertsReq, opts ...grpc.CallOption) (*GetAlertsResp, error) {
	out := new(GetAlertsResp)
	err := c.cc.Invoke(ctx, "/GeoAlt/GetActiveAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeoAltServer is the server API for GeoAlt service.
type GeoAltServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	GeoFeed(GeoAlt_GeoFeedServer) error
	AddAlert(context.Context, *AddAlertReq) (*AddAlertResp, error)
	GetAlerts(context.Context, *GetAlertsReq) (*GetAlertsResp, error)
	GetActiveAlerts(context.Context, *GetAlertsReq) (*GetAlertsResp, error)
}

func RegisterGeoAltServer(s *grpc.Server, srv GeoAltServer) {
	s.RegisterService(&_GeoAlt_serviceDesc, srv)
}

func _GeoAlt_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoAltServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoAlt/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoAltServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoAlt_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoAltServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoAlt/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoAltServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoAlt_GeoFeed_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GeoAltServer).GeoFeed(&geoAltGeoFeedServer{stream})
}

type GeoAlt_GeoFeedServer interface {
	Send(*GetAlertsResp) error
	Recv() (*GetAlertsReq, error)
	grpc.ServerStream
}

type geoAltGeoFeedServer struct {
	grpc.ServerStream
}

func (x *geoAltGeoFeedServer) Send(m *GetAlertsResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *geoAltGeoFeedServer) Recv() (*GetAlertsReq, error) {
	m := new(GetAlertsReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GeoAlt_AddAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAlertReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoAltServer).AddAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoAlt/AddAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoAltServer).AddAlert(ctx, req.(*AddAlertReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoAlt_GetAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoAltServer).GetAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoAlt/GetAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoAltServer).GetAlerts(ctx, req.(*GetAlertsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoAlt_GetActiveAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoAltServer).GetActiveAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoAlt/GetActiveAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoAltServer).GetActiveAlerts(ctx, req.(*GetAlertsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GeoAlt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GeoAlt",
	HandlerType: (*GeoAltServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _GeoAlt_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _GeoAlt_Register_Handler,
		},
		{
			MethodName: "AddAlert",
			Handler:    _GeoAlt_AddAlert_Handler,
		},
		{
			MethodName: "GetAlerts",
			Handler:    _GeoAlt_GetAlerts_Handler,
		},
		{
			MethodName: "GetActiveAlerts",
			Handler:    _GeoAlt_GetActiveAlerts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GeoFeed",
			Handler:       _GeoAlt_GeoFeed_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "geoaltsvc.proto",
}

func init() { proto.RegisterFile("geoaltsvc.proto", fileDescriptor_geoaltsvc_5e512dd99f0f8398) }

var fileDescriptor_geoaltsvc_5e512dd99f0f8398 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x95, 0x1d, 0xc7, 0xb1, 0x27, 0x49, 0xfb, 0x69, 0xf5, 0x49, 0x98, 0x40, 0xab, 0x60, 0x21,
	0x11, 0x01, 0x32, 0x55, 0xb8, 0x72, 0x89, 0x2a, 0xd1, 0x0b, 0xe2, 0xb0, 0x7f, 0x20, 0xda, 0xc6,
	0x83, 0xb1, 0xba, 0xf6, 0x9a, 0xdd, 0xa5, 0xc0, 0x89, 0x1b, 0xff, 0x80, 0x33, 0x3f, 0x88, 0x3f,
	0x85, 0x76, 0x1c, 0xd7, 0x0e, 0x15, 0xa2, 0xe2, 0xe6, 0x79, 0x6f, 0x67, 0xf7, 0xcd, 0xd3, 0x1b,
	0xc3, 0x71, 0x81, 0x4a, 0x48, 0x6b, 0xae, 0x77, 0x59, 0xa3, 0x95, 0x55, 0xe9, 0x77, 0x0f, 0xa6,
	0x1c, 0x8b, 0xd2, 0x58, 0xd4, 0x1c, 0x3f, 0xb0, 0xff, 0x61, 0x8c, 0x95, 0x28, 0x65, 0xe2, 0x2d,
	0xbd, 0x55, 0xcc, 0xdb, 0x82, 0x2d, 0x20, 0x6a, 0x84, 0x31, 0x9f, 0x94, 0xce, 0x13, 0x9f, 0x88,
	0x9b, 0x9a, 0x9d, 0x00, 0xbc, 0x2b, 0xb5, 0xb1, 0xdb, 0x5a, 0x54, 0x98, 0x8c, 0x88, 0x8d, 0x09,
	0x79, 0x2b, 0x2a, 0x64, 0x0f, 0x20, 0x96, 0xa2, 0x63, 0x83, 0xb6, 0xd7, 0x01, 0x44, 0x26, 0x30,
	0x11, 0x79, 0xae, 0xd1, 0x98, 0x64, 0x4c, 0x54, 0x57, 0xa6, 0x8f, 0x61, 0xd6, 0xcb, 0x32, 0x8d,
	0xd3, 0x65, 0xd5, 0x15, 0xd6, 0x9d, 0x2e, 0x2a, 0xd2, 0x57, 0x10, 0xbd, 0x51, 0x45, 0x59, 0xff,
	0x93, 0xf2, 0xf4, 0x11, 0xc4, 0xfb, 0xee, 0x3f, 0x3e, 0xf0, 0xc3, 0x83, 0xe9, 0x26, 0xcf, 0x37,
	0x12, 0xb5, 0x75, 0x8f, 0xfc, 0x07, 0x23, 0x29, 0x2c, 0x9d, 0xf1, 0xb8, 0xfb, 0x24, 0xa4, 0x2e,
	0xe8, 0x6e, 0x87, 0xd4, 0x85, 0x1b, 0xaa, 0x42, 0x63, 0x44, 0xd1, 0xb9, 0xd1, 0x95, 0xec, 0x21,
	0xc4, 0xd8, 0xbc, 0xc7, 0x0a, 0xb5, 0x90, 0xe4, 0x45, 0xc4, 0x7b, 0x80, 0x9d, 0x02, 0x68, 0x34,
	0x4a, 0x7e, 0xb4, 0xa5, 0xaa, 0xc9, 0x8f, 0x39, 0x1f, 0x20, 0x4e, 0x61, 0x8e, 0x52, 0x7c, 0x49,
	0xc2, 0xa5, 0xb7, 0x1a, 0xf1, 0xb6, 0x48, 0x4f, 0x61, 0xd6, 0x0b, 0x34, 0x0d, 0x3b, 0x02, 0x5f,
	0x5d, 0x91, 0xc0, 0x88, 0xfb, 0xea, 0x2a, 0x5d, 0xc3, 0xec, 0x02, 0x2d, 0xf1, 0xe6, 0x8e, 0x13,
	0xa4, 0x2f, 0x60, 0x3e, 0xe8, 0x31, 0x0d, 0x3b, 0x85, 0x50, 0x50, 0x95, 0x78, 0xcb, 0xd1, 0x6a,
	0xba, 0x0e, 0xb3, 0xf6, 0xc1, 0x3d, 0x9a, 0xfe, 0xf4, 0x60, 0x4c, 0x88, 0x3b, 0xb9, 0xc3, 0xda,
	0xa2, 0xa6, 0x17, 0xdc, 0xc9, 0x73, 0xa5, 0x74, 0xce, 0xf7, 0x28, 0x5b, 0xc2, 0xe4, 0x52, 0xe9,
	0x1c, 0xb5, 0x49, 0xfc, 0xfd, 0x55, 0xed, 0x81, 0x0e, 0x66, 0xf7, 0x21, 0xd8, 0xa1, 0x94, 0xe4,
	0xdd, 0x74, 0x3d, 0xce, 0xce, 0x51, 0x4a, 0x4e, 0xd0, 0xd0, 0xd9, 0xe0, 0x96, 0xb3, 0xb6, 0xac,
	0xd0, 0x58, 0x51, 0x35, 0x64, 0xdd, 0x88, 0xf7, 0x00, 0xbb, 0x07, 0x13, 0x8d, 0x22, 0xdf, 0x0a,
	0xbb, 0xf7, 0x2e, 0x74, 0xe5, 0xc6, 0xf6, 0x96, 0x4e, 0x86, 0x96, 0x3e, 0x83, 0x31, 0x69, 0xea,
	0xbc, 0xf2, 0x6f, 0x79, 0x35, 0xea, 0xbd, 0xfa, 0x0a, 0x81, 0x53, 0xe8, 0x72, 0x7e, 0x29, 0x0c,
	0x6e, 0x49, 0xbb, 0x9b, 0x3d, 0xe0, 0x91, 0x03, 0x88, 0x3c, 0x01, 0x28, 0xeb, 0x1c, 0x3f, 0xb7,
	0xac, 0x4f, 0x6c, 0x4c, 0x48, 0xd7, 0xab, 0x51, 0xc8, 0xed, 0xcd, 0xdc, 0x01, 0x8f, 0x1c, 0x40,
	0xe4, 0x61, 0x2c, 0x82, 0xdf, 0x63, 0xb1, 0xfe, 0xe6, 0x43, 0x78, 0x81, 0x6a, 0x23, 0x9d, 0xf9,
	0x63, 0x0a, 0x34, 0x8b, 0xb3, 0x6e, 0x2d, 0x16, 0x90, 0xf5, 0x19, 0x7f, 0x02, 0x51, 0xb7, 0x54,
	0x6c, 0x96, 0x0d, 0xd6, 0x7e, 0x31, 0xcf, 0x0e, 0xb6, 0xed, 0x39, 0x4c, 0x2e, 0x50, 0xbd, 0x46,
	0xcc, 0xd9, 0x3c, 0x1b, 0xc6, 0x67, 0x71, 0x94, 0x1d, 0x24, 0x63, 0xe5, 0x9d, 0x79, 0xee, 0xda,
	0x2e, 0x82, 0x6c, 0x96, 0x0d, 0xd6, 0x65, 0x31, 0xcf, 0x0e, 0xb2, 0xf9, 0x14, 0xe2, 0x9b, 0xee,
	0xbf, 0x5c, 0xcc, 0xce, 0xe0, 0xd8, 0x01, 0x3b, 0x5b, 0x5e, 0xe3, 0x9d, 0x3a, 0x2e, 0x43, 0xfa,
	0xa3, 0xbd, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0xda, 0x66, 0x30, 0x26, 0xe4, 0x04, 0x00, 0x00,
}
