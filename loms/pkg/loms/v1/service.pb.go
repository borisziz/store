// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: service.proto

package loms_v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderStatus int32

const (
	OrderStatus_Undefined       OrderStatus = 0
	OrderStatus_New             OrderStatus = 1
	OrderStatus_AwaitingPayment OrderStatus = 2
	OrderStatus_Failed          OrderStatus = 3
	OrderStatus_Payed           OrderStatus = 4
	OrderStatus_Cancelled       OrderStatus = 5
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "Undefined",
		1: "New",
		2: "AwaitingPayment",
		3: "Failed",
		4: "Payed",
		5: "Cancelled",
	}
	OrderStatus_value = map[string]int32{
		"Undefined":       0,
		"New":             1,
		"AwaitingPayment": 2,
		"Failed":          3,
		"Payed":           4,
		"Cancelled":       5,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku   uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *Item) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  int64   `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Items []*Item `protobuf:"bytes,2,rep,name=items,json=user,proto3" json:"items,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *CreateOrderRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID int64 `protobuf:"varint,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderResponse) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

type ListOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID int64 `protobuf:"varint,1,opt,name=orderID,json=order_id,proto3" json:"orderID,omitempty"`
}

func (x *ListOrderRequest) Reset() {
	*x = ListOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderRequest) ProtoMessage() {}

func (x *ListOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderRequest.ProtoReflect.Descriptor instead.
func (*ListOrderRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListOrderRequest) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status OrderStatus `protobuf:"varint,2,opt,name=status,proto3,enum=loms_v1.OrderStatus" json:"status,omitempty"`
	User   int64       `protobuf:"varint,3,opt,name=user,proto3" json:"user,omitempty"`
	Items  []*Item     `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *Order) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_Undefined
}

func (x *Order) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *Order) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status OrderStatus `protobuf:"varint,1,opt,name=status,proto3,enum=loms_v1.OrderStatus" json:"status,omitempty"`
	User   int64       `protobuf:"varint,2,opt,name=user,proto3" json:"user,omitempty"`
	Items  []*Item     `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListOrderResponse) Reset() {
	*x = ListOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderResponse) ProtoMessage() {}

func (x *ListOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderResponse.ProtoReflect.Descriptor instead.
func (*ListOrderResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListOrderResponse) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_Undefined
}

func (x *ListOrderResponse) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *ListOrderResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type OrderPayedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID int64 `protobuf:"varint,1,opt,name=orderID,json=order_id,proto3" json:"orderID,omitempty"`
}

func (x *OrderPayedRequest) Reset() {
	*x = OrderPayedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPayedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPayedRequest) ProtoMessage() {}

func (x *OrderPayedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPayedRequest.ProtoReflect.Descriptor instead.
func (*OrderPayedRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *OrderPayedRequest) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

type CancelOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID int64 `protobuf:"varint,1,opt,name=orderID,json=order_id,proto3" json:"orderID,omitempty"`
}

func (x *CancelOrderRequest) Reset() {
	*x = CancelOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderRequest) ProtoMessage() {}

func (x *CancelOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderRequest.ProtoReflect.Descriptor instead.
func (*CancelOrderRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{7}
}

func (x *CancelOrderRequest) GetOrderID() int64 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

type StocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *StocksRequest) Reset() {
	*x = StocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StocksRequest) ProtoMessage() {}

func (x *StocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StocksRequest.ProtoReflect.Descriptor instead.
func (*StocksRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{8}
}

func (x *StocksRequest) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

type Stock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WarehouseID int64  `protobuf:"varint,1,opt,name=warehouseID,proto3" json:"warehouseID,omitempty"`
	Count       uint64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Stock) Reset() {
	*x = Stock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stock) ProtoMessage() {}

func (x *Stock) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stock.ProtoReflect.Descriptor instead.
func (*Stock) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{9}
}

func (x *Stock) GetWarehouseID() int64 {
	if x != nil {
		return x.WarehouseID
	}
	return 0
}

func (x *Stock) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type StocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stocks []*Stock `protobuf:"bytes,1,rep,name=stocks,proto3" json:"stocks,omitempty"`
}

func (x *StocksResponse) Reset() {
	*x = StocksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StocksResponse) ProtoMessage() {}

func (x *StocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StocksResponse.ProtoReflect.Descriptor instead.
func (*StocksResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{10}
}

func (x *StocksResponse) GetStocks() []*Stock {
	if x != nil {
		return x.Stocks
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x04,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x19, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12,
	0x1d, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x55,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x36, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x7e,
	0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76,
	0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f,
	0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x7a,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x37, 0x0a, 0x11, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22,
	0x02, 0x20, 0x00, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x2a, 0x0a,
	0x0d, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x22, 0x3f, 0x0a, 0x05, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x38, 0x0a, 0x0e, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c,
	0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x73, 0x2a, 0x60, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x65, 0x77, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x41,
	0x77, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05,
	0x50, 0x61, 0x79, 0x65, 0x64, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x6c, 0x65, 0x64, 0x10, 0x05, 0x32, 0xf8, 0x03, 0x0a, 0x06, 0x4c, 0x4f, 0x4d, 0x53, 0x56,
	0x31, 0x12, 0x6a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x6c, 0x6f, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x62, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x6c, 0x6f, 0x6d,
	0x73, 0x5f, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x6c, 0x6f, 0x6d, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x3a, 0x01,
	0x2a, 0x12, 0x61, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x65, 0x64, 0x12,
	0x1a, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x61, 0x79, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x6c, 0x6f,
	0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x79, 0x65,
	0x64, 0x3a, 0x01, 0x2a, 0x12, 0x64, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a,
	0x22, 0x15, 0x2f, 0x6c, 0x6f, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x55, 0x0a, 0x06, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6c,
	0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f,
	0x6c, 0x6f, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x3a, 0x01,
	0x2a, 0x42, 0x23, 0x5a, 0x21, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2f, 0x6c, 0x6f,
	0x6d, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x3b, 0x6c,
	0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_service_proto_goTypes = []interface{}{
	(OrderStatus)(0),            // 0: loms_v1.OrderStatus
	(*Item)(nil),                // 1: loms_v1.Item
	(*CreateOrderRequest)(nil),  // 2: loms_v1.CreateOrderRequest
	(*CreateOrderResponse)(nil), // 3: loms_v1.CreateOrderResponse
	(*ListOrderRequest)(nil),    // 4: loms_v1.ListOrderRequest
	(*Order)(nil),               // 5: loms_v1.Order
	(*ListOrderResponse)(nil),   // 6: loms_v1.ListOrderResponse
	(*OrderPayedRequest)(nil),   // 7: loms_v1.OrderPayedRequest
	(*CancelOrderRequest)(nil),  // 8: loms_v1.CancelOrderRequest
	(*StocksRequest)(nil),       // 9: loms_v1.StocksRequest
	(*Stock)(nil),               // 10: loms_v1.Stock
	(*StocksResponse)(nil),      // 11: loms_v1.StocksResponse
	(*emptypb.Empty)(nil),       // 12: google.protobuf.Empty
}
var file_service_proto_depIdxs = []int32{
	1,  // 0: loms_v1.CreateOrderRequest.items:type_name -> loms_v1.Item
	0,  // 1: loms_v1.Order.status:type_name -> loms_v1.OrderStatus
	1,  // 2: loms_v1.Order.items:type_name -> loms_v1.Item
	0,  // 3: loms_v1.ListOrderResponse.status:type_name -> loms_v1.OrderStatus
	1,  // 4: loms_v1.ListOrderResponse.items:type_name -> loms_v1.Item
	10, // 5: loms_v1.StocksResponse.stocks:type_name -> loms_v1.Stock
	2,  // 6: loms_v1.LOMSV1.CreateOrder:input_type -> loms_v1.CreateOrderRequest
	4,  // 7: loms_v1.LOMSV1.ListOrder:input_type -> loms_v1.ListOrderRequest
	7,  // 8: loms_v1.LOMSV1.OrderPayed:input_type -> loms_v1.OrderPayedRequest
	8,  // 9: loms_v1.LOMSV1.CancelOrder:input_type -> loms_v1.CancelOrderRequest
	9,  // 10: loms_v1.LOMSV1.Stocks:input_type -> loms_v1.StocksRequest
	3,  // 11: loms_v1.LOMSV1.CreateOrder:output_type -> loms_v1.CreateOrderResponse
	6,  // 12: loms_v1.LOMSV1.ListOrder:output_type -> loms_v1.ListOrderResponse
	12, // 13: loms_v1.LOMSV1.OrderPayed:output_type -> google.protobuf.Empty
	12, // 14: loms_v1.LOMSV1.CancelOrder:output_type -> google.protobuf.Empty
	11, // 15: loms_v1.LOMSV1.Stocks:output_type -> loms_v1.StocksResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderRequest); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderResponse); i {
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
		file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPayedRequest); i {
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
		file_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOrderRequest); i {
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
		file_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StocksRequest); i {
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
		file_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stock); i {
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
		file_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StocksResponse); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		EnumInfos:         file_service_proto_enumTypes,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
