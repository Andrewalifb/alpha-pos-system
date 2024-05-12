// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: supplier.proto

package product_service

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

// PosSupplier
type PosSupplier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupplierId   string                 `protobuf:"bytes,1,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
	SupplierName string                 `protobuf:"bytes,2,opt,name=supplier_name,json=supplierName,proto3" json:"supplier_name,omitempty"`
	ContactName  string                 `protobuf:"bytes,3,opt,name=contact_name,json=contactName,proto3" json:"contact_name,omitempty"`
	ContactEmail string                 `protobuf:"bytes,4,opt,name=contact_email,json=contactEmail,proto3" json:"contact_email,omitempty"`
	ContactPhone string                 `protobuf:"bytes,5,opt,name=contact_phone,json=contactPhone,proto3" json:"contact_phone,omitempty"`
	CompanyId    string                 `protobuf:"bytes,6,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy    string                 `protobuf:"bytes,8,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy    string                 `protobuf:"bytes,10,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
}

func (x *PosSupplier) Reset() {
	*x = PosSupplier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PosSupplier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PosSupplier) ProtoMessage() {}

func (x *PosSupplier) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PosSupplier.ProtoReflect.Descriptor instead.
func (*PosSupplier) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{0}
}

func (x *PosSupplier) GetSupplierId() string {
	if x != nil {
		return x.SupplierId
	}
	return ""
}

func (x *PosSupplier) GetSupplierName() string {
	if x != nil {
		return x.SupplierName
	}
	return ""
}

func (x *PosSupplier) GetContactName() string {
	if x != nil {
		return x.ContactName
	}
	return ""
}

func (x *PosSupplier) GetContactEmail() string {
	if x != nil {
		return x.ContactEmail
	}
	return ""
}

func (x *PosSupplier) GetContactPhone() string {
	if x != nil {
		return x.ContactPhone
	}
	return ""
}

func (x *PosSupplier) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *PosSupplier) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PosSupplier) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *PosSupplier) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *PosSupplier) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

// Request and Response messages
type CreatePosSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosSupplier *PosSupplier `protobuf:"bytes,1,opt,name=pos_supplier,json=posSupplier,proto3" json:"pos_supplier,omitempty"`
	JwtPayload  *JWTPayload  `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *CreatePosSupplierRequest) Reset() {
	*x = CreatePosSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePosSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePosSupplierRequest) ProtoMessage() {}

func (x *CreatePosSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePosSupplierRequest.ProtoReflect.Descriptor instead.
func (*CreatePosSupplierRequest) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePosSupplierRequest) GetPosSupplier() *PosSupplier {
	if x != nil {
		return x.PosSupplier
	}
	return nil
}

func (x *CreatePosSupplierRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type CreatePosSupplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosSupplier *PosSupplier `protobuf:"bytes,1,opt,name=pos_supplier,json=posSupplier,proto3" json:"pos_supplier,omitempty"`
}

func (x *CreatePosSupplierResponse) Reset() {
	*x = CreatePosSupplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePosSupplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePosSupplierResponse) ProtoMessage() {}

func (x *CreatePosSupplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePosSupplierResponse.ProtoReflect.Descriptor instead.
func (*CreatePosSupplierResponse) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePosSupplierResponse) GetPosSupplier() *PosSupplier {
	if x != nil {
		return x.PosSupplier
	}
	return nil
}

type ReadPosSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupplierId string      `protobuf:"bytes,1,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *ReadPosSupplierRequest) Reset() {
	*x = ReadPosSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPosSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPosSupplierRequest) ProtoMessage() {}

func (x *ReadPosSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPosSupplierRequest.ProtoReflect.Descriptor instead.
func (*ReadPosSupplierRequest) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{3}
}

func (x *ReadPosSupplierRequest) GetSupplierId() string {
	if x != nil {
		return x.SupplierId
	}
	return ""
}

func (x *ReadPosSupplierRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type ReadPosSupplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosSupplier *PosSupplier `protobuf:"bytes,1,opt,name=pos_supplier,json=posSupplier,proto3" json:"pos_supplier,omitempty"`
}

func (x *ReadPosSupplierResponse) Reset() {
	*x = ReadPosSupplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPosSupplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPosSupplierResponse) ProtoMessage() {}

func (x *ReadPosSupplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPosSupplierResponse.ProtoReflect.Descriptor instead.
func (*ReadPosSupplierResponse) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{4}
}

func (x *ReadPosSupplierResponse) GetPosSupplier() *PosSupplier {
	if x != nil {
		return x.PosSupplier
	}
	return nil
}

type UpdatePosSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosSupplier *PosSupplier `protobuf:"bytes,1,opt,name=pos_supplier,json=posSupplier,proto3" json:"pos_supplier,omitempty"`
	JwtPayload  *JWTPayload  `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *UpdatePosSupplierRequest) Reset() {
	*x = UpdatePosSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosSupplierRequest) ProtoMessage() {}

func (x *UpdatePosSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosSupplierRequest.ProtoReflect.Descriptor instead.
func (*UpdatePosSupplierRequest) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePosSupplierRequest) GetPosSupplier() *PosSupplier {
	if x != nil {
		return x.PosSupplier
	}
	return nil
}

func (x *UpdatePosSupplierRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type UpdatePosSupplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosSupplier *PosSupplier `protobuf:"bytes,1,opt,name=pos_supplier,json=posSupplier,proto3" json:"pos_supplier,omitempty"`
}

func (x *UpdatePosSupplierResponse) Reset() {
	*x = UpdatePosSupplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosSupplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosSupplierResponse) ProtoMessage() {}

func (x *UpdatePosSupplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosSupplierResponse.ProtoReflect.Descriptor instead.
func (*UpdatePosSupplierResponse) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePosSupplierResponse) GetPosSupplier() *PosSupplier {
	if x != nil {
		return x.PosSupplier
	}
	return nil
}

type DeletePosSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupplierId string      `protobuf:"bytes,1,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *DeletePosSupplierRequest) Reset() {
	*x = DeletePosSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePosSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePosSupplierRequest) ProtoMessage() {}

func (x *DeletePosSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePosSupplierRequest.ProtoReflect.Descriptor instead.
func (*DeletePosSupplierRequest) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePosSupplierRequest) GetSupplierId() string {
	if x != nil {
		return x.SupplierId
	}
	return ""
}

func (x *DeletePosSupplierRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type DeletePosSupplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeletePosSupplierResponse) Reset() {
	*x = DeletePosSupplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePosSupplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePosSupplierResponse) ProtoMessage() {}

func (x *DeletePosSupplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePosSupplierResponse.ProtoReflect.Descriptor instead.
func (*DeletePosSupplierResponse) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePosSupplierResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ReadAllPosSuppliersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtPayload *JWTPayload `protobuf:"bytes,1,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *ReadAllPosSuppliersRequest) Reset() {
	*x = ReadAllPosSuppliersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllPosSuppliersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllPosSuppliersRequest) ProtoMessage() {}

func (x *ReadAllPosSuppliersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllPosSuppliersRequest.ProtoReflect.Descriptor instead.
func (*ReadAllPosSuppliersRequest) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{9}
}

func (x *ReadAllPosSuppliersRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type ReadAllPosSuppliersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosSuppliers []*PosSupplier `protobuf:"bytes,1,rep,name=pos_suppliers,json=posSuppliers,proto3" json:"pos_suppliers,omitempty"`
}

func (x *ReadAllPosSuppliersResponse) Reset() {
	*x = ReadAllPosSuppliersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllPosSuppliersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllPosSuppliersResponse) ProtoMessage() {}

func (x *ReadAllPosSuppliersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllPosSuppliersResponse.ProtoReflect.Descriptor instead.
func (*ReadAllPosSuppliersResponse) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{10}
}

func (x *ReadAllPosSuppliersResponse) GetPosSuppliers() []*PosSupplier {
	if x != nil {
		return x.PosSuppliers
	}
	return nil
}

var File_supplier_proto protoreflect.FileDescriptor

var file_supplier_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x70, 0x6f, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x03, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x81, 0x01, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x5f, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52,
	0x0b, 0x70, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0b,
	0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x50,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x70,
	0x6f, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x22, 0x6b, 0x0a, 0x16, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x6a,
	0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x4e, 0x0a,
	0x17, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x5f,
	0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x52, 0x0b, 0x70, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x22, 0x81, 0x01,
	0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0c, 0x70, 0x6f,
	0x73, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12,
	0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x50, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x53, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x53, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x22, 0x6d, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x35, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4e, 0x0a, 0x1a, 0x52, 0x65, 0x61,
	0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a,
	0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x54, 0x0a, 0x1b, 0x52, 0x65, 0x61,
	0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x5f,
	0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x32,
	0xb8, 0x03, 0x0a, 0x12, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x70, 0x6f,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x52, 0x65,
	0x61, 0x64, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x1b, 0x2e,
	0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x73,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x2e,
	0x70, 0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70,
	0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x58, 0x0a, 0x13, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x53, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x64, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x69, 0x66, 0x62, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2d, 0x70, 0x6f, 0x73, 0x2d, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_supplier_proto_rawDescOnce sync.Once
	file_supplier_proto_rawDescData = file_supplier_proto_rawDesc
)

func file_supplier_proto_rawDescGZIP() []byte {
	file_supplier_proto_rawDescOnce.Do(func() {
		file_supplier_proto_rawDescData = protoimpl.X.CompressGZIP(file_supplier_proto_rawDescData)
	})
	return file_supplier_proto_rawDescData
}

var file_supplier_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_supplier_proto_goTypes = []interface{}{
	(*PosSupplier)(nil),                 // 0: pos.PosSupplier
	(*CreatePosSupplierRequest)(nil),    // 1: pos.CreatePosSupplierRequest
	(*CreatePosSupplierResponse)(nil),   // 2: pos.CreatePosSupplierResponse
	(*ReadPosSupplierRequest)(nil),      // 3: pos.ReadPosSupplierRequest
	(*ReadPosSupplierResponse)(nil),     // 4: pos.ReadPosSupplierResponse
	(*UpdatePosSupplierRequest)(nil),    // 5: pos.UpdatePosSupplierRequest
	(*UpdatePosSupplierResponse)(nil),   // 6: pos.UpdatePosSupplierResponse
	(*DeletePosSupplierRequest)(nil),    // 7: pos.DeletePosSupplierRequest
	(*DeletePosSupplierResponse)(nil),   // 8: pos.DeletePosSupplierResponse
	(*ReadAllPosSuppliersRequest)(nil),  // 9: pos.ReadAllPosSuppliersRequest
	(*ReadAllPosSuppliersResponse)(nil), // 10: pos.ReadAllPosSuppliersResponse
	(*timestamppb.Timestamp)(nil),       // 11: google.protobuf.Timestamp
	(*JWTPayload)(nil),                  // 12: pos.JWTPayload
}
var file_supplier_proto_depIdxs = []int32{
	11, // 0: pos.PosSupplier.created_at:type_name -> google.protobuf.Timestamp
	11, // 1: pos.PosSupplier.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: pos.CreatePosSupplierRequest.pos_supplier:type_name -> pos.PosSupplier
	12, // 3: pos.CreatePosSupplierRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 4: pos.CreatePosSupplierResponse.pos_supplier:type_name -> pos.PosSupplier
	12, // 5: pos.ReadPosSupplierRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 6: pos.ReadPosSupplierResponse.pos_supplier:type_name -> pos.PosSupplier
	0,  // 7: pos.UpdatePosSupplierRequest.pos_supplier:type_name -> pos.PosSupplier
	12, // 8: pos.UpdatePosSupplierRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 9: pos.UpdatePosSupplierResponse.pos_supplier:type_name -> pos.PosSupplier
	12, // 10: pos.DeletePosSupplierRequest.jwt_payload:type_name -> pos.JWTPayload
	12, // 11: pos.ReadAllPosSuppliersRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 12: pos.ReadAllPosSuppliersResponse.pos_suppliers:type_name -> pos.PosSupplier
	1,  // 13: pos.PosSupplierService.CreatePosSupplier:input_type -> pos.CreatePosSupplierRequest
	3,  // 14: pos.PosSupplierService.ReadPosSupplier:input_type -> pos.ReadPosSupplierRequest
	5,  // 15: pos.PosSupplierService.UpdatePosSupplier:input_type -> pos.UpdatePosSupplierRequest
	7,  // 16: pos.PosSupplierService.DeletePosSupplier:input_type -> pos.DeletePosSupplierRequest
	9,  // 17: pos.PosSupplierService.ReadAllPosSuppliers:input_type -> pos.ReadAllPosSuppliersRequest
	2,  // 18: pos.PosSupplierService.CreatePosSupplier:output_type -> pos.CreatePosSupplierResponse
	4,  // 19: pos.PosSupplierService.ReadPosSupplier:output_type -> pos.ReadPosSupplierResponse
	6,  // 20: pos.PosSupplierService.UpdatePosSupplier:output_type -> pos.UpdatePosSupplierResponse
	8,  // 21: pos.PosSupplierService.DeletePosSupplier:output_type -> pos.DeletePosSupplierResponse
	10, // 22: pos.PosSupplierService.ReadAllPosSuppliers:output_type -> pos.ReadAllPosSuppliersResponse
	18, // [18:23] is the sub-list for method output_type
	13, // [13:18] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_supplier_proto_init() }
func file_supplier_proto_init() {
	if File_supplier_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_supplier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PosSupplier); i {
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
		file_supplier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePosSupplierRequest); i {
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
		file_supplier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePosSupplierResponse); i {
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
		file_supplier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPosSupplierRequest); i {
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
		file_supplier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPosSupplierResponse); i {
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
		file_supplier_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosSupplierRequest); i {
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
		file_supplier_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosSupplierResponse); i {
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
		file_supplier_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePosSupplierRequest); i {
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
		file_supplier_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePosSupplierResponse); i {
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
		file_supplier_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllPosSuppliersRequest); i {
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
		file_supplier_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllPosSuppliersResponse); i {
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
			RawDescriptor: file_supplier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_supplier_proto_goTypes,
		DependencyIndexes: file_supplier_proto_depIdxs,
		MessageInfos:      file_supplier_proto_msgTypes,
	}.Build()
	File_supplier_proto = out.File
	file_supplier_proto_rawDesc = nil
	file_supplier_proto_goTypes = nil
	file_supplier_proto_depIdxs = nil
}
