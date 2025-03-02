// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: inventory.proto

package inventory

import (
	context "context"
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

// 查询库存相关消息
type QueryStockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkuId []uint64 `protobuf:"varint,1,rep,packed,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"` // SKU ID
}

func (x *QueryStockReq) Reset() {
	*x = QueryStockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryStockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryStockReq) ProtoMessage() {}

func (x *QueryStockReq) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryStockReq.ProtoReflect.Descriptor instead.
func (*QueryStockReq) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *QueryStockReq) GetSkuId() []uint64 {
	if x != nil {
		return x.SkuId
	}
	return nil
}

type QueryStockResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentStock map[uint64]uint32 `protobuf:"bytes,3,rep,name=current_stock,json=currentStock,proto3" json:"current_stock,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` //map[sku_id]stock
}

func (x *QueryStockResp) Reset() {
	*x = QueryStockResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryStockResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryStockResp) ProtoMessage() {}

func (x *QueryStockResp) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryStockResp.ProtoReflect.Descriptor instead.
func (*QueryStockResp) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *QueryStockResp) GetCurrentStock() map[uint64]uint32 {
	if x != nil {
		return x.CurrentStock
	}
	return nil
}

type InventoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string               `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"` // 订单号（全局唯一）
	Items   []*InventoryReq_Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	Force   bool                 `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"` //管理员强制操作
}

func (x *InventoryReq) Reset() {
	*x = InventoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryReq) ProtoMessage() {}

func (x *InventoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryReq.ProtoReflect.Descriptor instead.
func (*InventoryReq) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *InventoryReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *InventoryReq) GetItems() []*InventoryReq_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *InventoryReq) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

type InventoryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Errors  []*InventoryResp_Error `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *InventoryResp) Reset() {
	*x = InventoryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryResp) ProtoMessage() {}

func (x *InventoryResp) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryResp.ProtoReflect.Descriptor instead.
func (*InventoryResp) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *InventoryResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *InventoryResp) GetErrors() []*InventoryResp_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

// 商品服务 → 库存服务
type ProductCreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkuId        uint64 `protobuf:"varint,1,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	SkuName      string `protobuf:"bytes,2,opt,name=sku_name,json=skuName,proto3" json:"sku_name,omitempty"`
	InitialStock uint32 `protobuf:"varint,3,opt,name=initial_stock,json=initialStock,proto3" json:"initial_stock,omitempty"`
	Operator     string `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"` //  google.protobuf.Timestamp event_time = 5;
}

func (x *ProductCreatedEvent) Reset() {
	*x = ProductCreatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCreatedEvent) ProtoMessage() {}

func (x *ProductCreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCreatedEvent.ProtoReflect.Descriptor instead.
func (*ProductCreatedEvent) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{4}
}

func (x *ProductCreatedEvent) GetSkuId() uint64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *ProductCreatedEvent) GetSkuName() string {
	if x != nil {
		return x.SkuName
	}
	return ""
}

func (x *ProductCreatedEvent) GetInitialStock() uint32 {
	if x != nil {
		return x.InitialStock
	}
	return 0
}

func (x *ProductCreatedEvent) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

type InventoryReq_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkuId    uint64 `protobuf:"varint,1,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"` // SKU ID
	Quantity int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`        // 数量
}

func (x *InventoryReq_Item) Reset() {
	*x = InventoryReq_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryReq_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryReq_Item) ProtoMessage() {}

func (x *InventoryReq_Item) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryReq_Item.ProtoReflect.Descriptor instead.
func (*InventoryReq_Item) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{2, 0}
}

func (x *InventoryReq_Item) GetSkuId() uint64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *InventoryReq_Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type InventoryResp_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkuId   uint64 `protobuf:"varint,1,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"` // 失败条目
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`           // 错误详情
}

func (x *InventoryResp_Error) Reset() {
	*x = InventoryResp_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryResp_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryResp_Error) ProtoMessage() {}

func (x *InventoryResp_Error) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryResp_Error.ProtoReflect.Descriptor instead.
func (*InventoryResp_Error) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{3, 0}
}

func (x *InventoryResp_Error) GetSkuId() uint64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *InventoryResp_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_inventory_proto protoreflect.FileDescriptor

var file_inventory_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x26, 0x0a, 0x0d,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x15, 0x0a,
	0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x05, 0x73,
	0x6b, 0x75, 0x49, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x50, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x1a, 0x3f, 0x0a, 0x11, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xae, 0x01, 0x0a, 0x0c, 0x49,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x1a, 0x39, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x9b, 0x01, 0x0a, 0x0d,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a,
	0x38, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x13, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x75, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x75, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x32, 0x9e, 0x02, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x1a, 0x19, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x0c,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x41, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12,
	0x17, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x41, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x12, 0x17, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x7a, 0x63, 0x32, 0x31, 0x36, 0x2f, 0x67, 0x6f, 0x6d, 0x61, 0x6c,
	0x6c, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f,
	0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventory_proto_rawDescOnce sync.Once
	file_inventory_proto_rawDescData = file_inventory_proto_rawDesc
)

func file_inventory_proto_rawDescGZIP() []byte {
	file_inventory_proto_rawDescOnce.Do(func() {
		file_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_proto_rawDescData)
	})
	return file_inventory_proto_rawDescData
}

var file_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_inventory_proto_goTypes = []interface{}{
	(*QueryStockReq)(nil),       // 0: inventory.QueryStockReq
	(*QueryStockResp)(nil),      // 1: inventory.QueryStockResp
	(*InventoryReq)(nil),        // 2: inventory.InventoryReq
	(*InventoryResp)(nil),       // 3: inventory.InventoryResp
	(*ProductCreatedEvent)(nil), // 4: inventory.ProductCreatedEvent
	nil,                         // 5: inventory.QueryStockResp.CurrentStockEntry
	(*InventoryReq_Item)(nil),   // 6: inventory.InventoryReq.Item
	(*InventoryResp_Error)(nil), // 7: inventory.InventoryResp.Error
}
var file_inventory_proto_depIdxs = []int32{
	5, // 0: inventory.QueryStockResp.current_stock:type_name -> inventory.QueryStockResp.CurrentStockEntry
	6, // 1: inventory.InventoryReq.items:type_name -> inventory.InventoryReq.Item
	7, // 2: inventory.InventoryResp.errors:type_name -> inventory.InventoryResp.Error
	0, // 3: inventory.InventoryService.QueryStock:input_type -> inventory.QueryStockReq
	2, // 4: inventory.InventoryService.ReserveStock:input_type -> inventory.InventoryReq
	2, // 5: inventory.InventoryService.ConfirmStock:input_type -> inventory.InventoryReq
	2, // 6: inventory.InventoryService.ReleaseStock:input_type -> inventory.InventoryReq
	1, // 7: inventory.InventoryService.QueryStock:output_type -> inventory.QueryStockResp
	3, // 8: inventory.InventoryService.ReserveStock:output_type -> inventory.InventoryResp
	3, // 9: inventory.InventoryService.ConfirmStock:output_type -> inventory.InventoryResp
	3, // 10: inventory.InventoryService.ReleaseStock:output_type -> inventory.InventoryResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_inventory_proto_init() }
func file_inventory_proto_init() {
	if File_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryStockReq); i {
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
		file_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryStockResp); i {
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
		file_inventory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryReq); i {
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
		file_inventory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryResp); i {
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
		file_inventory_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCreatedEvent); i {
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
		file_inventory_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryReq_Item); i {
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
		file_inventory_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryResp_Error); i {
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
			RawDescriptor: file_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_proto_goTypes,
		DependencyIndexes: file_inventory_proto_depIdxs,
		MessageInfos:      file_inventory_proto_msgTypes,
	}.Build()
	File_inventory_proto = out.File
	file_inventory_proto_rawDesc = nil
	file_inventory_proto_goTypes = nil
	file_inventory_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.1. DO NOT EDIT.

type InventoryService interface {
	QueryStock(ctx context.Context, req *QueryStockReq) (res *QueryStockResp, err error)
	ReserveStock(ctx context.Context, req *InventoryReq) (res *InventoryResp, err error)
	ConfirmStock(ctx context.Context, req *InventoryReq) (res *InventoryResp, err error)
	ReleaseStock(ctx context.Context, req *InventoryReq) (res *InventoryResp, err error)
}
