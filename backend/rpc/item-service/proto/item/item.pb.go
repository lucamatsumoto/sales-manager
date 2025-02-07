// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/item/item.proto

package go_micro_srv_item

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Item struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string   `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Category             string   `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_685d1a6429ed79c2, []int{0}
}
func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (dst *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(dst, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Item) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

type GetItemByNameRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetItemByNameRequest) Reset()         { *m = GetItemByNameRequest{} }
func (m *GetItemByNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetItemByNameRequest) ProtoMessage()    {}
func (*GetItemByNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_685d1a6429ed79c2, []int{1}
}
func (m *GetItemByNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetItemByNameRequest.Unmarshal(m, b)
}
func (m *GetItemByNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetItemByNameRequest.Marshal(b, m, deterministic)
}
func (dst *GetItemByNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetItemByNameRequest.Merge(dst, src)
}
func (m *GetItemByNameRequest) XXX_Size() int {
	return xxx_messageInfo_GetItemByNameRequest.Size(m)
}
func (m *GetItemByNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetItemByNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetItemByNameRequest proto.InternalMessageInfo

func (m *GetItemByNameRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_685d1a6429ed79c2, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	Created              bool     `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Item                 *Item    `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Items                []*Item  `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_685d1a6429ed79c2, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Response) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "go.micro.srv.item.Item")
	proto.RegisterType((*GetItemByNameRequest)(nil), "go.micro.srv.item.GetItemByNameRequest")
	proto.RegisterType((*Request)(nil), "go.micro.srv.item.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.item.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ItemService service

type ItemServiceClient interface {
	CreateItem(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error)
	GetItems(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetItemsByName(ctx context.Context, in *GetItemByNameRequest, opts ...client.CallOption) (*Response, error)
}

type itemServiceClient struct {
	c           client.Client
	serviceName string
}

func NewItemServiceClient(serviceName string, c client.Client) ItemServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.item"
	}
	return &itemServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *itemServiceClient) CreateItem(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ItemService.CreateItem", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetItems(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ItemService.GetItems", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetItemsByName(ctx context.Context, in *GetItemByNameRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ItemService.GetItemsByName", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ItemService service

type ItemServiceHandler interface {
	CreateItem(context.Context, *Item, *Response) error
	GetItems(context.Context, *Request, *Response) error
	GetItemsByName(context.Context, *GetItemByNameRequest, *Response) error
}

func RegisterItemServiceHandler(s server.Server, hdlr ItemServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ItemService{hdlr}, opts...))
}

type ItemService struct {
	ItemServiceHandler
}

func (h *ItemService) CreateItem(ctx context.Context, in *Item, out *Response) error {
	return h.ItemServiceHandler.CreateItem(ctx, in, out)
}

func (h *ItemService) GetItems(ctx context.Context, in *Request, out *Response) error {
	return h.ItemServiceHandler.GetItems(ctx, in, out)
}

func (h *ItemService) GetItemsByName(ctx context.Context, in *GetItemByNameRequest, out *Response) error {
	return h.ItemServiceHandler.GetItemsByName(ctx, in, out)
}

func init() { proto.RegisterFile("proto/item/item.proto", fileDescriptor_item_685d1a6429ed79c2) }

var fileDescriptor_item_685d1a6429ed79c2 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xbd, 0x4e, 0xf3, 0x40,
	0x10, 0xfc, 0xfc, 0x93, 0x2f, 0xce, 0x46, 0x8a, 0xc4, 0x2a, 0x88, 0x93, 0x69, 0x22, 0x17, 0x80,
	0x84, 0x38, 0xa4, 0xf0, 0x06, 0xfc, 0x08, 0xd1, 0x50, 0x98, 0x86, 0x82, 0xc6, 0x38, 0xab, 0xe8,
	0x8a, 0xf3, 0x85, 0xbb, 0x23, 0x28, 0x1d, 0x0f, 0xcd, 0x03, 0xa0, 0x5b, 0xc7, 0x14, 0x60, 0x48,
	0x63, 0x7b, 0x76, 0x66, 0x77, 0x76, 0x47, 0x86, 0xfd, 0x95, 0x35, 0xde, 0x9c, 0x2b, 0x4f, 0x9a,
	0x1f, 0x92, 0x31, 0xee, 0x2d, 0x8d, 0xd4, 0xaa, 0xb6, 0x46, 0x3a, 0xbb, 0x96, 0x81, 0x28, 0x9e,
	0x20, 0xbd, 0xf3, 0xa4, 0x71, 0x02, 0xb1, 0x5a, 0x88, 0x68, 0x16, 0x9d, 0x8c, 0xca, 0x58, 0x2d,
	0x10, 0x21, 0x6d, 0x2a, 0x4d, 0x22, 0xe6, 0x0a, 0x7f, 0xe3, 0x14, 0x06, 0xe6, 0xad, 0x21, 0x2b,
	0x12, 0x2e, 0xb6, 0x00, 0x73, 0xc8, 0xea, 0xca, 0xd3, 0xd2, 0xd8, 0x8d, 0x48, 0x99, 0xf8, 0xc2,
	0xc5, 0x11, 0x4c, 0x6f, 0xc9, 0x07, 0x83, 0xcb, 0xcd, 0x7d, 0xa5, 0xa9, 0xa4, 0x97, 0x57, 0x72,
	0xfe, 0xbb, 0x5b, 0x31, 0x82, 0xe1, 0x96, 0x2a, 0xde, 0x23, 0xc8, 0x4a, 0x72, 0x2b, 0xd3, 0x38,
	0x42, 0x01, 0xc3, 0xda, 0x52, 0xe5, 0xa9, 0x15, 0x67, 0x65, 0x07, 0xf1, 0x14, 0xd2, 0xb0, 0x3f,
	0xef, 0x37, 0x9e, 0x1f, 0xc8, 0x1f, 0x97, 0xc9, 0xe0, 0x5a, 0xb2, 0x08, 0xcf, 0x60, 0x10, 0xde,
	0x4e, 0x24, 0xb3, 0xe4, 0x2f, 0x75, 0xab, 0x9a, 0x7f, 0x44, 0x30, 0x0e, 0xf8, 0x81, 0xec, 0x5a,
	0xd5, 0x84, 0xd7, 0x00, 0x57, 0x6c, 0xcb, 0x49, 0xfd, 0xd6, 0x9d, 0x1f, 0xf6, 0x10, 0xdd, 0x25,
	0xc5, 0x3f, 0xbc, 0x81, 0x6c, 0x9b, 0x85, 0xc3, 0xbc, 0x57, 0xca, 0x01, 0xec, 0x1a, 0xf3, 0x08,
	0x93, 0x6e, 0x4c, 0x9b, 0x29, 0x1e, 0xf7, 0x34, 0xf4, 0xa5, 0xbe, 0x63, 0xf2, 0xf3, 0x7f, 0xfe,
	0x49, 0x2e, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xef, 0x03, 0x82, 0xdc, 0x3d, 0x02, 0x00, 0x00,
}
