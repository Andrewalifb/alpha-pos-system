// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: return.proto

package sales_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PosReturn
type PosReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReturnId   string                 `protobuf:"bytes,1,opt,name=return_id,json=returnId,proto3" json:"return_id,omitempty"`
	SaleId     string                 `protobuf:"bytes,2,opt,name=sale_id,json=saleId,proto3" json:"sale_id,omitempty"`
	ProductId  string                 `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity   int32                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ReturnDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
	Reason     string                 `protobuf:"bytes,6,opt,name=reason,proto3" json:"reason,omitempty"`
	CompanyId  string                 `protobuf:"bytes,7,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy  string                 `protobuf:"bytes,9,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy  string                 `protobuf:"bytes,11,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
}

func (x *PosReturn) Reset() {
	*x = PosReturn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_return_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PosReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PosReturn) ProtoMessage() {}

func (x *PosReturn) ProtoReflect() protoreflect.Message {
	mi := &file_return_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PosReturn.ProtoReflect.Descriptor instead.
func (*PosReturn) Descriptor() ([]byte, []int) {
	return file_return_proto_rawDescGZIP(), []int{0}
}

func (x *PosReturn) GetReturnId() string {
	if x != nil {
		return x.ReturnId
	}
	return ""
}

func (x *PosReturn) GetSaleId() string {
	if x != nil {
		return x.SaleId
	}
	return ""
}

func (x *PosReturn) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *PosReturn) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *PosReturn) GetReturnDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ReturnDate
	}
	return nil
}

func (x *PosReturn) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *PosReturn) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *PosReturn) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PosReturn) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *PosReturn) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *PosReturn) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

// Request and Response messages
type CreatePosReturnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosReturn  *PosReturn  `protobuf:"bytes,1,opt,name=pos_return,json=posReturn,proto3" json:"pos_return,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *CreatePosReturnRequest) Reset() {
	*x = CreatePosReturnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_return_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePosReturnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePosReturnRequest) ProtoMessage() {}

func (x *CreatePosReturnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_return_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePosReturnRequest.ProtoReflect.Descriptor instead.
func (*CreatePosReturnRequest) Descriptor() ([]byte, []int) {
	return file_return_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePosReturnRequest) GetPosReturn() *PosReturn {
	if x != nil {
		return x.PosReturn
	}
	return nil
}

func (x *CreatePosReturnRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type CreatePosReturnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosReturn *PosReturn `protobuf:"bytes,1,opt,name=pos_return,json=posReturn,proto3" json:"pos_return,omitempty"`
}

func (x *CreatePosReturnResponse) Reset() {
	*x = CreatePosReturnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_return_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePosReturnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePosReturnResponse) ProtoMessage() {}

func (x *CreatePosReturnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_return_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePosReturnResponse.ProtoReflect.Descriptor instead.
func (*CreatePosReturnResponse) Descriptor() ([]byte, []int) {
	return file_return_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePosReturnResponse) GetPosReturn() *PosReturn {
	if x != nil {
		return x.PosReturn
	}
	return nil
}

type ReadPosReturnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReturnId   string      `protobuf:"bytes,1,opt,name=return_id,json=returnId,proto3" json:"return_id,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *ReadPosReturnRequest) Reset() {
	*x = ReadPosReturnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_return_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPosReturnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPosReturnRequest) ProtoMessage() {}

func (x *ReadPosReturnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_return_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPosReturnRequest.ProtoReflect.Descriptor instead.
func (*ReadPosReturnRequest) Descriptor() ([]byte, []int) {
	return file_return_proto_rawDescGZIP(), []int{3}
}

func (x *ReadPosReturnRequest) GetReturnId() string {
	if x != nil {
		return x.ReturnId
	}
	return ""
}

func (x *ReadPosReturnRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type ReadPosReturnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosReturn *PosReturn `protobuf:"bytes,1,opt,name=pos_return,json=posReturn,proto3" json:"pos_return,omitempty"`
}

func (x *ReadPosReturnResponse) Reset() {
	*x = ReadPosReturnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_return_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPosReturnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPosReturnResponse) ProtoMessage() {}

func (x *ReadPosReturnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_return_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPosReturnResponse.ProtoReflect.Descriptor instead.
func (*ReadPosReturnResponse) Descriptor() ([]byte, []int) {
	return file_return_proto_rawDescGZIP(), []int{4}
}

func (x *ReadPosReturnResponse) GetPosReturn() *PosReturn {
	if x != nil {
		return x.PosReturn
	}
	return nil
}

type UpdatePosReturnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosReturn  *PosReturn  `protobuf:"bytes,1,opt,name=pos_return,json=posReturn,proto3" json:"pos_return,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *UpdatePosReturnRequest) Reset() {
	*x = UpdatePosReturnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_return_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosReturnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosReturnRequest) ProtoMessage() {}

func (x *UpdatePosReturnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_return_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosReturnRequest.ProtoReflect.Descriptor instead.
func (*UpdatePosReturnRequest) Descriptor() ([]byte, []int) {
	return file_return_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePosReturnRequest) GetPosReturn() *PosReturn {
	if x != nil {
		return x.PosReturn
	}
	return nil
}

func (x *UpdatePosReturnRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type UpdatePosReturnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosReturn *PosReturn `protobuf:"bytes,1,opt,name=pos_return,json=posReturn,proto3" json:"pos_return,omitempty"`
}

func (x *UpdatePosReturnResponse) Reset() {
	*x = UpdatePosReturnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_return_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosReturnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosReturnResponse) ProtoMessage() {}

func (x *UpdatePosReturnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_return_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosReturnResponse.ProtoReflect.Descriptor instead.
func (*UpdatePosReturnResponse) Descriptor() ([]byte, []int) {
	return file_return_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePosReturnResponse) GetPosReturn() *PosReturn {
	if x != nil {
		return x.PosReturn
	}
	return nil
}

type DeletePosReturnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReturnId   string      `protobuf:"bytes,1,opt,name=return_id,json=returnId,proto3" json:"return_id,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *DeletePosReturnRequest) Reset() {
	*x = DeletePosReturnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_return_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePosReturnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePosReturnRequest) ProtoMessage() {}

func (x *DeletePosReturnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_return_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePosReturnRequest.ProtoReflect.Descriptor instead.
func (*DeletePosReturnRequest) Descriptor() ([]byte, []int) {
	return file_return_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePosReturnRequest) GetReturnId() string {
	if x != nil {
		return x.ReturnId
	}
	return ""
}

func (x *DeletePosReturnRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type DeletePosReturnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeletePosReturnResponse) Reset() {
	*x = DeletePosReturnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_return_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePosReturnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePosReturnResponse) ProtoMessage() {}

func (x *DeletePosReturnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_return_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePosReturnResponse.ProtoReflect.Descriptor instead.
func (*DeletePosReturnResponse) Descriptor() ([]byte, []int) {
	return file_return_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePosReturnResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ReadAllPosReturnsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtPayload *JWTPayload `protobuf:"bytes,1,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *ReadAllPosReturnsRequest) Reset() {
	*x = ReadAllPosReturnsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_return_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllPosReturnsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllPosReturnsRequest) ProtoMessage() {}

func (x *ReadAllPosReturnsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_return_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllPosReturnsRequest.ProtoReflect.Descriptor instead.
func (*ReadAllPosReturnsRequest) Descriptor() ([]byte, []int) {
	return file_return_proto_rawDescGZIP(), []int{9}
}

func (x *ReadAllPosReturnsRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type ReadAllPosReturnsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosReturns []*PosReturn `protobuf:"bytes,1,rep,name=pos_returns,json=posReturns,proto3" json:"pos_returns,omitempty"`
}

func (x *ReadAllPosReturnsResponse) Reset() {
	*x = ReadAllPosReturnsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_return_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllPosReturnsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllPosReturnsResponse) ProtoMessage() {}

func (x *ReadAllPosReturnsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_return_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllPosReturnsResponse.ProtoReflect.Descriptor instead.
func (*ReadAllPosReturnsResponse) Descriptor() ([]byte, []int) {
	return file_return_proto_rawDescGZIP(), []int{10}
}

func (x *ReadAllPosReturnsResponse) GetPosReturns() []*PosReturn {
	if x != nil {
		return x.PosReturns
	}
	return nil
}

var File_return_proto protoreflect.FileDescriptor

var file_return_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x70, 0x6f, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa4, 0x03, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x61, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x79, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f,
	0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57,
	0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x48, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x65,
	0x0a, 0x14, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a,
	0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x46, 0x0a, 0x15, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d,
	0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x79, 0x0a,
	0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x5f, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f,
	0x73, 0x2e, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f,
	0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77,
	0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x48, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f,
	0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x22, 0x67, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74,
	0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x33, 0x0a, 0x17, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x4c, 0x0a, 0x18, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0b,
	0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x4c,
	0x0a, 0x19, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x70,
	0x6f, 0x73, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x52, 0x0a, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x32, 0x98, 0x03, 0x0a,
	0x10, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x46, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x12, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6f,
	0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f,
	0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x64, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x69, 0x66,
	0x62, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2d, 0x70, 0x6f, 0x73, 0x2d, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_return_proto_rawDescOnce sync.Once
	file_return_proto_rawDescData = file_return_proto_rawDesc
)

func file_return_proto_rawDescGZIP() []byte {
	file_return_proto_rawDescOnce.Do(func() {
		file_return_proto_rawDescData = protoimpl.X.CompressGZIP(file_return_proto_rawDescData)
	})
	return file_return_proto_rawDescData
}

var file_return_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_return_proto_goTypes = []interface{}{
	(*PosReturn)(nil),                 // 0: pos.PosReturn
	(*CreatePosReturnRequest)(nil),    // 1: pos.CreatePosReturnRequest
	(*CreatePosReturnResponse)(nil),   // 2: pos.CreatePosReturnResponse
	(*ReadPosReturnRequest)(nil),      // 3: pos.ReadPosReturnRequest
	(*ReadPosReturnResponse)(nil),     // 4: pos.ReadPosReturnResponse
	(*UpdatePosReturnRequest)(nil),    // 5: pos.UpdatePosReturnRequest
	(*UpdatePosReturnResponse)(nil),   // 6: pos.UpdatePosReturnResponse
	(*DeletePosReturnRequest)(nil),    // 7: pos.DeletePosReturnRequest
	(*DeletePosReturnResponse)(nil),   // 8: pos.DeletePosReturnResponse
	(*ReadAllPosReturnsRequest)(nil),  // 9: pos.ReadAllPosReturnsRequest
	(*ReadAllPosReturnsResponse)(nil), // 10: pos.ReadAllPosReturnsResponse
	(*timestamppb.Timestamp)(nil),     // 11: google.protobuf.Timestamp
	(*JWTPayload)(nil),                // 12: pos.JWTPayload
}
var file_return_proto_depIdxs = []int32{
	11, // 0: pos.PosReturn.return_date:type_name -> google.protobuf.Timestamp
	11, // 1: pos.PosReturn.created_at:type_name -> google.protobuf.Timestamp
	11, // 2: pos.PosReturn.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 3: pos.CreatePosReturnRequest.pos_return:type_name -> pos.PosReturn
	12, // 4: pos.CreatePosReturnRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 5: pos.CreatePosReturnResponse.pos_return:type_name -> pos.PosReturn
	12, // 6: pos.ReadPosReturnRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 7: pos.ReadPosReturnResponse.pos_return:type_name -> pos.PosReturn
	0,  // 8: pos.UpdatePosReturnRequest.pos_return:type_name -> pos.PosReturn
	12, // 9: pos.UpdatePosReturnRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 10: pos.UpdatePosReturnResponse.pos_return:type_name -> pos.PosReturn
	12, // 11: pos.DeletePosReturnRequest.jwt_payload:type_name -> pos.JWTPayload
	12, // 12: pos.ReadAllPosReturnsRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 13: pos.ReadAllPosReturnsResponse.pos_returns:type_name -> pos.PosReturn
	1,  // 14: pos.PosReturnService.CreatePosReturn:input_type -> pos.CreatePosReturnRequest
	3,  // 15: pos.PosReturnService.ReadPosReturn:input_type -> pos.ReadPosReturnRequest
	5,  // 16: pos.PosReturnService.UpdatePosReturn:input_type -> pos.UpdatePosReturnRequest
	7,  // 17: pos.PosReturnService.DeletePosReturn:input_type -> pos.DeletePosReturnRequest
	9,  // 18: pos.PosReturnService.ReadAllPosReturns:input_type -> pos.ReadAllPosReturnsRequest
	2,  // 19: pos.PosReturnService.CreatePosReturn:output_type -> pos.CreatePosReturnResponse
	4,  // 20: pos.PosReturnService.ReadPosReturn:output_type -> pos.ReadPosReturnResponse
	6,  // 21: pos.PosReturnService.UpdatePosReturn:output_type -> pos.UpdatePosReturnResponse
	8,  // 22: pos.PosReturnService.DeletePosReturn:output_type -> pos.DeletePosReturnResponse
	10, // 23: pos.PosReturnService.ReadAllPosReturns:output_type -> pos.ReadAllPosReturnsResponse
	19, // [19:24] is the sub-list for method output_type
	14, // [14:19] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_return_proto_init() }
func file_return_proto_init() {
	if File_return_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_return_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PosReturn); i {
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
		file_return_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePosReturnRequest); i {
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
		file_return_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePosReturnResponse); i {
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
		file_return_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPosReturnRequest); i {
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
		file_return_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPosReturnResponse); i {
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
		file_return_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosReturnRequest); i {
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
		file_return_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosReturnResponse); i {
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
		file_return_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePosReturnRequest); i {
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
		file_return_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePosReturnResponse); i {
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
		file_return_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllPosReturnsRequest); i {
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
		file_return_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllPosReturnsResponse); i {
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
			RawDescriptor: file_return_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_return_proto_goTypes,
		DependencyIndexes: file_return_proto_depIdxs,
		MessageInfos:      file_return_proto_msgTypes,
	}.Build()
	File_return_proto = out.File
	file_return_proto_rawDesc = nil
	file_return_proto_goTypes = nil
	file_return_proto_depIdxs = nil
}
