// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: invoices.proto

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

// PosInvoice
type PosInvoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvoiceId string                 `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	SaleId    string                 `protobuf:"bytes,2,opt,name=sale_id,json=saleId,proto3" json:"sale_id,omitempty"`
	Date      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Total     float64                `protobuf:"fixed64,4,opt,name=total,proto3" json:"total,omitempty"`
	Discounts float64                `protobuf:"fixed64,5,opt,name=discounts,proto3" json:"discounts,omitempty"`
	Taxes     float64                `protobuf:"fixed64,6,opt,name=taxes,proto3" json:"taxes,omitempty"`
	CompanyId string                 `protobuf:"bytes,7,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy string                 `protobuf:"bytes,9,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy string                 `protobuf:"bytes,11,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
}

func (x *PosInvoice) Reset() {
	*x = PosInvoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoices_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PosInvoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PosInvoice) ProtoMessage() {}

func (x *PosInvoice) ProtoReflect() protoreflect.Message {
	mi := &file_invoices_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PosInvoice.ProtoReflect.Descriptor instead.
func (*PosInvoice) Descriptor() ([]byte, []int) {
	return file_invoices_proto_rawDescGZIP(), []int{0}
}

func (x *PosInvoice) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *PosInvoice) GetSaleId() string {
	if x != nil {
		return x.SaleId
	}
	return ""
}

func (x *PosInvoice) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *PosInvoice) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PosInvoice) GetDiscounts() float64 {
	if x != nil {
		return x.Discounts
	}
	return 0
}

func (x *PosInvoice) GetTaxes() float64 {
	if x != nil {
		return x.Taxes
	}
	return 0
}

func (x *PosInvoice) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *PosInvoice) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PosInvoice) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *PosInvoice) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *PosInvoice) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

// Request and Response messages
type CreatePosInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosInvoice *PosInvoice `protobuf:"bytes,1,opt,name=pos_invoice,json=posInvoice,proto3" json:"pos_invoice,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *CreatePosInvoiceRequest) Reset() {
	*x = CreatePosInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoices_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePosInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePosInvoiceRequest) ProtoMessage() {}

func (x *CreatePosInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invoices_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePosInvoiceRequest.ProtoReflect.Descriptor instead.
func (*CreatePosInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_invoices_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePosInvoiceRequest) GetPosInvoice() *PosInvoice {
	if x != nil {
		return x.PosInvoice
	}
	return nil
}

func (x *CreatePosInvoiceRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type CreatePosInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosInvoice *PosInvoice `protobuf:"bytes,1,opt,name=pos_invoice,json=posInvoice,proto3" json:"pos_invoice,omitempty"`
}

func (x *CreatePosInvoiceResponse) Reset() {
	*x = CreatePosInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoices_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePosInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePosInvoiceResponse) ProtoMessage() {}

func (x *CreatePosInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invoices_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePosInvoiceResponse.ProtoReflect.Descriptor instead.
func (*CreatePosInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_invoices_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePosInvoiceResponse) GetPosInvoice() *PosInvoice {
	if x != nil {
		return x.PosInvoice
	}
	return nil
}

type ReadPosInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvoiceId  string      `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *ReadPosInvoiceRequest) Reset() {
	*x = ReadPosInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoices_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPosInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPosInvoiceRequest) ProtoMessage() {}

func (x *ReadPosInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invoices_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPosInvoiceRequest.ProtoReflect.Descriptor instead.
func (*ReadPosInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_invoices_proto_rawDescGZIP(), []int{3}
}

func (x *ReadPosInvoiceRequest) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *ReadPosInvoiceRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type ReadPosInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosInvoice *PosInvoice `protobuf:"bytes,1,opt,name=pos_invoice,json=posInvoice,proto3" json:"pos_invoice,omitempty"`
}

func (x *ReadPosInvoiceResponse) Reset() {
	*x = ReadPosInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoices_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPosInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPosInvoiceResponse) ProtoMessage() {}

func (x *ReadPosInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invoices_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPosInvoiceResponse.ProtoReflect.Descriptor instead.
func (*ReadPosInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_invoices_proto_rawDescGZIP(), []int{4}
}

func (x *ReadPosInvoiceResponse) GetPosInvoice() *PosInvoice {
	if x != nil {
		return x.PosInvoice
	}
	return nil
}

type UpdatePosInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosInvoice *PosInvoice `protobuf:"bytes,1,opt,name=pos_invoice,json=posInvoice,proto3" json:"pos_invoice,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *UpdatePosInvoiceRequest) Reset() {
	*x = UpdatePosInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoices_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosInvoiceRequest) ProtoMessage() {}

func (x *UpdatePosInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invoices_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosInvoiceRequest.ProtoReflect.Descriptor instead.
func (*UpdatePosInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_invoices_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePosInvoiceRequest) GetPosInvoice() *PosInvoice {
	if x != nil {
		return x.PosInvoice
	}
	return nil
}

func (x *UpdatePosInvoiceRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type UpdatePosInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosInvoice *PosInvoice `protobuf:"bytes,1,opt,name=pos_invoice,json=posInvoice,proto3" json:"pos_invoice,omitempty"`
}

func (x *UpdatePosInvoiceResponse) Reset() {
	*x = UpdatePosInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoices_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosInvoiceResponse) ProtoMessage() {}

func (x *UpdatePosInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invoices_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosInvoiceResponse.ProtoReflect.Descriptor instead.
func (*UpdatePosInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_invoices_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePosInvoiceResponse) GetPosInvoice() *PosInvoice {
	if x != nil {
		return x.PosInvoice
	}
	return nil
}

type DeletePosInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvoiceId  string      `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *DeletePosInvoiceRequest) Reset() {
	*x = DeletePosInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoices_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePosInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePosInvoiceRequest) ProtoMessage() {}

func (x *DeletePosInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invoices_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePosInvoiceRequest.ProtoReflect.Descriptor instead.
func (*DeletePosInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_invoices_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePosInvoiceRequest) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *DeletePosInvoiceRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type DeletePosInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeletePosInvoiceResponse) Reset() {
	*x = DeletePosInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoices_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePosInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePosInvoiceResponse) ProtoMessage() {}

func (x *DeletePosInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invoices_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePosInvoiceResponse.ProtoReflect.Descriptor instead.
func (*DeletePosInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_invoices_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePosInvoiceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ReadAllPosInvoicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtPayload *JWTPayload `protobuf:"bytes,1,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *ReadAllPosInvoicesRequest) Reset() {
	*x = ReadAllPosInvoicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoices_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllPosInvoicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllPosInvoicesRequest) ProtoMessage() {}

func (x *ReadAllPosInvoicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invoices_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllPosInvoicesRequest.ProtoReflect.Descriptor instead.
func (*ReadAllPosInvoicesRequest) Descriptor() ([]byte, []int) {
	return file_invoices_proto_rawDescGZIP(), []int{9}
}

func (x *ReadAllPosInvoicesRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type ReadAllPosInvoicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosInvoices []*PosInvoice `protobuf:"bytes,1,rep,name=pos_invoices,json=posInvoices,proto3" json:"pos_invoices,omitempty"`
}

func (x *ReadAllPosInvoicesResponse) Reset() {
	*x = ReadAllPosInvoicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoices_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllPosInvoicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllPosInvoicesResponse) ProtoMessage() {}

func (x *ReadAllPosInvoicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invoices_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllPosInvoicesResponse.ProtoReflect.Descriptor instead.
func (*ReadAllPosInvoicesResponse) Descriptor() ([]byte, []int) {
	return file_invoices_proto_rawDescGZIP(), []int{10}
}

func (x *ReadAllPosInvoicesResponse) GetPosInvoices() []*PosInvoice {
	if x != nil {
		return x.PosInvoices
	}
	return nil
}

var File_invoices_proto protoreflect.FileDescriptor

var file_invoices_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x70, 0x6f, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x03, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x61, 0x78, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x74, 0x61, 0x78, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x7d, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50,
	0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73,
	0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x4c, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50,
	0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x22, 0x68, 0x0a, 0x15, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x30, 0x0a,
	0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x4a, 0x0a, 0x16, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x70, 0x6f, 0x73,
	0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52,
	0x0a, 0x70, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x22, 0x7d, 0x0a, 0x17, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x5f, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f,
	0x73, 0x2e, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x0a, 0x70, 0x6f,
	0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a,
	0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x4c, 0x0a, 0x18, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x5f, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f,
	0x73, 0x2e, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x0a, 0x70, 0x6f,
	0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x22, 0x6a, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57,
	0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x34, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4d, 0x0a, 0x19, 0x52, 0x65,
	0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a,
	0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x50, 0x0a, 0x1a, 0x52, 0x65, 0x61,
	0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x5f, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x0b,
	0x70, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x32, 0xa8, 0x03, 0x0a, 0x11,
	0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4f, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50,
	0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a,
	0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x12, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x55, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x64, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x69, 0x66, 0x62,
	0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2d, 0x70, 0x6f, 0x73, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_invoices_proto_rawDescOnce sync.Once
	file_invoices_proto_rawDescData = file_invoices_proto_rawDesc
)

func file_invoices_proto_rawDescGZIP() []byte {
	file_invoices_proto_rawDescOnce.Do(func() {
		file_invoices_proto_rawDescData = protoimpl.X.CompressGZIP(file_invoices_proto_rawDescData)
	})
	return file_invoices_proto_rawDescData
}

var file_invoices_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_invoices_proto_goTypes = []interface{}{
	(*PosInvoice)(nil),                 // 0: pos.PosInvoice
	(*CreatePosInvoiceRequest)(nil),    // 1: pos.CreatePosInvoiceRequest
	(*CreatePosInvoiceResponse)(nil),   // 2: pos.CreatePosInvoiceResponse
	(*ReadPosInvoiceRequest)(nil),      // 3: pos.ReadPosInvoiceRequest
	(*ReadPosInvoiceResponse)(nil),     // 4: pos.ReadPosInvoiceResponse
	(*UpdatePosInvoiceRequest)(nil),    // 5: pos.UpdatePosInvoiceRequest
	(*UpdatePosInvoiceResponse)(nil),   // 6: pos.UpdatePosInvoiceResponse
	(*DeletePosInvoiceRequest)(nil),    // 7: pos.DeletePosInvoiceRequest
	(*DeletePosInvoiceResponse)(nil),   // 8: pos.DeletePosInvoiceResponse
	(*ReadAllPosInvoicesRequest)(nil),  // 9: pos.ReadAllPosInvoicesRequest
	(*ReadAllPosInvoicesResponse)(nil), // 10: pos.ReadAllPosInvoicesResponse
	(*timestamppb.Timestamp)(nil),      // 11: google.protobuf.Timestamp
	(*JWTPayload)(nil),                 // 12: pos.JWTPayload
}
var file_invoices_proto_depIdxs = []int32{
	11, // 0: pos.PosInvoice.date:type_name -> google.protobuf.Timestamp
	11, // 1: pos.PosInvoice.created_at:type_name -> google.protobuf.Timestamp
	11, // 2: pos.PosInvoice.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 3: pos.CreatePosInvoiceRequest.pos_invoice:type_name -> pos.PosInvoice
	12, // 4: pos.CreatePosInvoiceRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 5: pos.CreatePosInvoiceResponse.pos_invoice:type_name -> pos.PosInvoice
	12, // 6: pos.ReadPosInvoiceRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 7: pos.ReadPosInvoiceResponse.pos_invoice:type_name -> pos.PosInvoice
	0,  // 8: pos.UpdatePosInvoiceRequest.pos_invoice:type_name -> pos.PosInvoice
	12, // 9: pos.UpdatePosInvoiceRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 10: pos.UpdatePosInvoiceResponse.pos_invoice:type_name -> pos.PosInvoice
	12, // 11: pos.DeletePosInvoiceRequest.jwt_payload:type_name -> pos.JWTPayload
	12, // 12: pos.ReadAllPosInvoicesRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 13: pos.ReadAllPosInvoicesResponse.pos_invoices:type_name -> pos.PosInvoice
	1,  // 14: pos.PosInvoiceService.CreatePosInvoice:input_type -> pos.CreatePosInvoiceRequest
	3,  // 15: pos.PosInvoiceService.ReadPosInvoice:input_type -> pos.ReadPosInvoiceRequest
	5,  // 16: pos.PosInvoiceService.UpdatePosInvoice:input_type -> pos.UpdatePosInvoiceRequest
	7,  // 17: pos.PosInvoiceService.DeletePosInvoice:input_type -> pos.DeletePosInvoiceRequest
	9,  // 18: pos.PosInvoiceService.ReadAllPosInvoices:input_type -> pos.ReadAllPosInvoicesRequest
	2,  // 19: pos.PosInvoiceService.CreatePosInvoice:output_type -> pos.CreatePosInvoiceResponse
	4,  // 20: pos.PosInvoiceService.ReadPosInvoice:output_type -> pos.ReadPosInvoiceResponse
	6,  // 21: pos.PosInvoiceService.UpdatePosInvoice:output_type -> pos.UpdatePosInvoiceResponse
	8,  // 22: pos.PosInvoiceService.DeletePosInvoice:output_type -> pos.DeletePosInvoiceResponse
	10, // 23: pos.PosInvoiceService.ReadAllPosInvoices:output_type -> pos.ReadAllPosInvoicesResponse
	19, // [19:24] is the sub-list for method output_type
	14, // [14:19] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_invoices_proto_init() }
func file_invoices_proto_init() {
	if File_invoices_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_invoices_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PosInvoice); i {
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
		file_invoices_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePosInvoiceRequest); i {
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
		file_invoices_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePosInvoiceResponse); i {
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
		file_invoices_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPosInvoiceRequest); i {
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
		file_invoices_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPosInvoiceResponse); i {
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
		file_invoices_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosInvoiceRequest); i {
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
		file_invoices_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosInvoiceResponse); i {
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
		file_invoices_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePosInvoiceRequest); i {
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
		file_invoices_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePosInvoiceResponse); i {
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
		file_invoices_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllPosInvoicesRequest); i {
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
		file_invoices_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllPosInvoicesResponse); i {
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
			RawDescriptor: file_invoices_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_invoices_proto_goTypes,
		DependencyIndexes: file_invoices_proto_depIdxs,
		MessageInfos:      file_invoices_proto_msgTypes,
	}.Build()
	File_invoices_proto = out.File
	file_invoices_proto_rawDesc = nil
	file_invoices_proto_goTypes = nil
	file_invoices_proto_depIdxs = nil
}
