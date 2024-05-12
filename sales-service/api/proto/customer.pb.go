// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: customer.proto

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

// PosCustomer
type PosCustomer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId       string                 `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	FirstName        string                 `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName         string                 `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email            string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber      string                 `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	DateOfBirth      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	RegistrationDate *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=registration_date,json=registrationDate,proto3" json:"registration_date,omitempty"`
	Address          string                 `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	City             string                 `protobuf:"bytes,9,opt,name=city,proto3" json:"city,omitempty"`
	Country          string                 `protobuf:"bytes,10,opt,name=country,proto3" json:"country,omitempty"`
	CompanyId        string                 `protobuf:"bytes,11,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy        string                 `protobuf:"bytes,13,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt        *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy        string                 `protobuf:"bytes,15,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
}

func (x *PosCustomer) Reset() {
	*x = PosCustomer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PosCustomer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PosCustomer) ProtoMessage() {}

func (x *PosCustomer) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PosCustomer.ProtoReflect.Descriptor instead.
func (*PosCustomer) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{0}
}

func (x *PosCustomer) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *PosCustomer) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *PosCustomer) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *PosCustomer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PosCustomer) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *PosCustomer) GetDateOfBirth() *timestamppb.Timestamp {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

func (x *PosCustomer) GetRegistrationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.RegistrationDate
	}
	return nil
}

func (x *PosCustomer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PosCustomer) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *PosCustomer) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *PosCustomer) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *PosCustomer) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PosCustomer) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *PosCustomer) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *PosCustomer) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

// Request and Response messages
type CreatePosCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosCustomer *PosCustomer `protobuf:"bytes,1,opt,name=pos_customer,json=posCustomer,proto3" json:"pos_customer,omitempty"`
	JwtPayload  *JWTPayload  `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *CreatePosCustomerRequest) Reset() {
	*x = CreatePosCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePosCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePosCustomerRequest) ProtoMessage() {}

func (x *CreatePosCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePosCustomerRequest.ProtoReflect.Descriptor instead.
func (*CreatePosCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePosCustomerRequest) GetPosCustomer() *PosCustomer {
	if x != nil {
		return x.PosCustomer
	}
	return nil
}

func (x *CreatePosCustomerRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type CreatePosCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosCustomer *PosCustomer `protobuf:"bytes,1,opt,name=pos_customer,json=posCustomer,proto3" json:"pos_customer,omitempty"`
}

func (x *CreatePosCustomerResponse) Reset() {
	*x = CreatePosCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePosCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePosCustomerResponse) ProtoMessage() {}

func (x *CreatePosCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePosCustomerResponse.ProtoReflect.Descriptor instead.
func (*CreatePosCustomerResponse) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePosCustomerResponse) GetPosCustomer() *PosCustomer {
	if x != nil {
		return x.PosCustomer
	}
	return nil
}

type ReadPosCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string      `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *ReadPosCustomerRequest) Reset() {
	*x = ReadPosCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPosCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPosCustomerRequest) ProtoMessage() {}

func (x *ReadPosCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPosCustomerRequest.ProtoReflect.Descriptor instead.
func (*ReadPosCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{3}
}

func (x *ReadPosCustomerRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *ReadPosCustomerRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type ReadPosCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosCustomer *PosCustomer `protobuf:"bytes,1,opt,name=pos_customer,json=posCustomer,proto3" json:"pos_customer,omitempty"`
}

func (x *ReadPosCustomerResponse) Reset() {
	*x = ReadPosCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPosCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPosCustomerResponse) ProtoMessage() {}

func (x *ReadPosCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPosCustomerResponse.ProtoReflect.Descriptor instead.
func (*ReadPosCustomerResponse) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{4}
}

func (x *ReadPosCustomerResponse) GetPosCustomer() *PosCustomer {
	if x != nil {
		return x.PosCustomer
	}
	return nil
}

type UpdatePosCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosCustomer *PosCustomer `protobuf:"bytes,1,opt,name=pos_customer,json=posCustomer,proto3" json:"pos_customer,omitempty"`
	JwtPayload  *JWTPayload  `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *UpdatePosCustomerRequest) Reset() {
	*x = UpdatePosCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosCustomerRequest) ProtoMessage() {}

func (x *UpdatePosCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosCustomerRequest.ProtoReflect.Descriptor instead.
func (*UpdatePosCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePosCustomerRequest) GetPosCustomer() *PosCustomer {
	if x != nil {
		return x.PosCustomer
	}
	return nil
}

func (x *UpdatePosCustomerRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type UpdatePosCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosCustomer *PosCustomer `protobuf:"bytes,1,opt,name=pos_customer,json=posCustomer,proto3" json:"pos_customer,omitempty"`
}

func (x *UpdatePosCustomerResponse) Reset() {
	*x = UpdatePosCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosCustomerResponse) ProtoMessage() {}

func (x *UpdatePosCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosCustomerResponse.ProtoReflect.Descriptor instead.
func (*UpdatePosCustomerResponse) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePosCustomerResponse) GetPosCustomer() *PosCustomer {
	if x != nil {
		return x.PosCustomer
	}
	return nil
}

type DeletePosCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string      `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *DeletePosCustomerRequest) Reset() {
	*x = DeletePosCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePosCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePosCustomerRequest) ProtoMessage() {}

func (x *DeletePosCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePosCustomerRequest.ProtoReflect.Descriptor instead.
func (*DeletePosCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePosCustomerRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *DeletePosCustomerRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type DeletePosCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeletePosCustomerResponse) Reset() {
	*x = DeletePosCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePosCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePosCustomerResponse) ProtoMessage() {}

func (x *DeletePosCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePosCustomerResponse.ProtoReflect.Descriptor instead.
func (*DeletePosCustomerResponse) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePosCustomerResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ReadAllPosCustomersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtPayload *JWTPayload `protobuf:"bytes,1,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *ReadAllPosCustomersRequest) Reset() {
	*x = ReadAllPosCustomersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllPosCustomersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllPosCustomersRequest) ProtoMessage() {}

func (x *ReadAllPosCustomersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllPosCustomersRequest.ProtoReflect.Descriptor instead.
func (*ReadAllPosCustomersRequest) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{9}
}

func (x *ReadAllPosCustomersRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type ReadAllPosCustomersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosCustomers []*PosCustomer `protobuf:"bytes,1,rep,name=pos_customers,json=posCustomers,proto3" json:"pos_customers,omitempty"`
}

func (x *ReadAllPosCustomersResponse) Reset() {
	*x = ReadAllPosCustomersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllPosCustomersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllPosCustomersResponse) ProtoMessage() {}

func (x *ReadAllPosCustomersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllPosCustomersResponse.ProtoReflect.Descriptor instead.
func (*ReadAllPosCustomersResponse) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{10}
}

func (x *ReadAllPosCustomersResponse) GetPosCustomers() []*PosCustomer {
	if x != nil {
		return x.PosCustomers
	}
	return nil
}

var File_customer_proto protoreflect.FileDescriptor

var file_customer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x70, 0x6f, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x04, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0d, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x47, 0x0a, 0x11, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x81,
	0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0c, 0x70,
	0x6f, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x50, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x33, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x22, 0x6b, 0x0a, 0x16, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x4e, 0x0a, 0x17, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0c,
	0x70, 0x6f, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x22, 0x81, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33,
	0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a,
	0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x50, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50,
	0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x6d, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e,
	0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x35, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4e, 0x0a,
	0x1a, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0b, 0x6a,
	0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x54, 0x0a,
	0x1b, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0d,
	0x70, 0x6f, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x32, 0xb8, 0x03, 0x0a, 0x12, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12,
	0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c,
	0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x52, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x13, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50,
	0x6f, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x6f,
	0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70,
	0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x64,
	0x72, 0x65, 0x77, 0x61, 0x6c, 0x69, 0x66, 0x62, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2d, 0x70,
	0x6f, 0x73, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customer_proto_rawDescOnce sync.Once
	file_customer_proto_rawDescData = file_customer_proto_rawDesc
)

func file_customer_proto_rawDescGZIP() []byte {
	file_customer_proto_rawDescOnce.Do(func() {
		file_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_customer_proto_rawDescData)
	})
	return file_customer_proto_rawDescData
}

var file_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_customer_proto_goTypes = []interface{}{
	(*PosCustomer)(nil),                 // 0: pos.PosCustomer
	(*CreatePosCustomerRequest)(nil),    // 1: pos.CreatePosCustomerRequest
	(*CreatePosCustomerResponse)(nil),   // 2: pos.CreatePosCustomerResponse
	(*ReadPosCustomerRequest)(nil),      // 3: pos.ReadPosCustomerRequest
	(*ReadPosCustomerResponse)(nil),     // 4: pos.ReadPosCustomerResponse
	(*UpdatePosCustomerRequest)(nil),    // 5: pos.UpdatePosCustomerRequest
	(*UpdatePosCustomerResponse)(nil),   // 6: pos.UpdatePosCustomerResponse
	(*DeletePosCustomerRequest)(nil),    // 7: pos.DeletePosCustomerRequest
	(*DeletePosCustomerResponse)(nil),   // 8: pos.DeletePosCustomerResponse
	(*ReadAllPosCustomersRequest)(nil),  // 9: pos.ReadAllPosCustomersRequest
	(*ReadAllPosCustomersResponse)(nil), // 10: pos.ReadAllPosCustomersResponse
	(*timestamppb.Timestamp)(nil),       // 11: google.protobuf.Timestamp
	(*JWTPayload)(nil),                  // 12: pos.JWTPayload
}
var file_customer_proto_depIdxs = []int32{
	11, // 0: pos.PosCustomer.date_of_birth:type_name -> google.protobuf.Timestamp
	11, // 1: pos.PosCustomer.registration_date:type_name -> google.protobuf.Timestamp
	11, // 2: pos.PosCustomer.created_at:type_name -> google.protobuf.Timestamp
	11, // 3: pos.PosCustomer.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 4: pos.CreatePosCustomerRequest.pos_customer:type_name -> pos.PosCustomer
	12, // 5: pos.CreatePosCustomerRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 6: pos.CreatePosCustomerResponse.pos_customer:type_name -> pos.PosCustomer
	12, // 7: pos.ReadPosCustomerRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 8: pos.ReadPosCustomerResponse.pos_customer:type_name -> pos.PosCustomer
	0,  // 9: pos.UpdatePosCustomerRequest.pos_customer:type_name -> pos.PosCustomer
	12, // 10: pos.UpdatePosCustomerRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 11: pos.UpdatePosCustomerResponse.pos_customer:type_name -> pos.PosCustomer
	12, // 12: pos.DeletePosCustomerRequest.jwt_payload:type_name -> pos.JWTPayload
	12, // 13: pos.ReadAllPosCustomersRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 14: pos.ReadAllPosCustomersResponse.pos_customers:type_name -> pos.PosCustomer
	1,  // 15: pos.PosCustomerService.CreatePosCustomer:input_type -> pos.CreatePosCustomerRequest
	3,  // 16: pos.PosCustomerService.ReadPosCustomer:input_type -> pos.ReadPosCustomerRequest
	5,  // 17: pos.PosCustomerService.UpdatePosCustomer:input_type -> pos.UpdatePosCustomerRequest
	7,  // 18: pos.PosCustomerService.DeletePosCustomer:input_type -> pos.DeletePosCustomerRequest
	9,  // 19: pos.PosCustomerService.ReadAllPosCustomers:input_type -> pos.ReadAllPosCustomersRequest
	2,  // 20: pos.PosCustomerService.CreatePosCustomer:output_type -> pos.CreatePosCustomerResponse
	4,  // 21: pos.PosCustomerService.ReadPosCustomer:output_type -> pos.ReadPosCustomerResponse
	6,  // 22: pos.PosCustomerService.UpdatePosCustomer:output_type -> pos.UpdatePosCustomerResponse
	8,  // 23: pos.PosCustomerService.DeletePosCustomer:output_type -> pos.DeletePosCustomerResponse
	10, // 24: pos.PosCustomerService.ReadAllPosCustomers:output_type -> pos.ReadAllPosCustomersResponse
	20, // [20:25] is the sub-list for method output_type
	15, // [15:20] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_customer_proto_init() }
func file_customer_proto_init() {
	if File_customer_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PosCustomer); i {
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
		file_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePosCustomerRequest); i {
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
		file_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePosCustomerResponse); i {
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
		file_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPosCustomerRequest); i {
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
		file_customer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPosCustomerResponse); i {
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
		file_customer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosCustomerRequest); i {
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
		file_customer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosCustomerResponse); i {
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
		file_customer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePosCustomerRequest); i {
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
		file_customer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePosCustomerResponse); i {
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
		file_customer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllPosCustomersRequest); i {
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
		file_customer_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllPosCustomersResponse); i {
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
			RawDescriptor: file_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customer_proto_goTypes,
		DependencyIndexes: file_customer_proto_depIdxs,
		MessageInfos:      file_customer_proto_msgTypes,
	}.Build()
	File_customer_proto = out.File
	file_customer_proto_rawDesc = nil
	file_customer_proto_goTypes = nil
	file_customer_proto_depIdxs = nil
}
