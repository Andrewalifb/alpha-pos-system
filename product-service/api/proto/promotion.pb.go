// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: promotion.proto

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

// PosPromotion
type PosPromotion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PromotionId  string                 `protobuf:"bytes,1,opt,name=promotion_id,json=promotionId,proto3" json:"promotion_id,omitempty"`
	ProductId    string                 `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	StartDate    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	DiscountRate float64                `protobuf:"fixed64,5,opt,name=discount_rate,json=discountRate,proto3" json:"discount_rate,omitempty"`
	CompanyId    string                 `protobuf:"bytes,6,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy    string                 `protobuf:"bytes,8,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy    string                 `protobuf:"bytes,10,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
}

func (x *PosPromotion) Reset() {
	*x = PosPromotion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PosPromotion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PosPromotion) ProtoMessage() {}

func (x *PosPromotion) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PosPromotion.ProtoReflect.Descriptor instead.
func (*PosPromotion) Descriptor() ([]byte, []int) {
	return file_promotion_proto_rawDescGZIP(), []int{0}
}

func (x *PosPromotion) GetPromotionId() string {
	if x != nil {
		return x.PromotionId
	}
	return ""
}

func (x *PosPromotion) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *PosPromotion) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *PosPromotion) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *PosPromotion) GetDiscountRate() float64 {
	if x != nil {
		return x.DiscountRate
	}
	return 0
}

func (x *PosPromotion) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *PosPromotion) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PosPromotion) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *PosPromotion) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *PosPromotion) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

// Request and Response messages
type CreatePosPromotionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosPromotion *PosPromotion `protobuf:"bytes,1,opt,name=pos_promotion,json=posPromotion,proto3" json:"pos_promotion,omitempty"`
	JwtPayload   *JWTPayload   `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *CreatePosPromotionRequest) Reset() {
	*x = CreatePosPromotionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePosPromotionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePosPromotionRequest) ProtoMessage() {}

func (x *CreatePosPromotionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePosPromotionRequest.ProtoReflect.Descriptor instead.
func (*CreatePosPromotionRequest) Descriptor() ([]byte, []int) {
	return file_promotion_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePosPromotionRequest) GetPosPromotion() *PosPromotion {
	if x != nil {
		return x.PosPromotion
	}
	return nil
}

func (x *CreatePosPromotionRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type CreatePosPromotionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosPromotion *PosPromotion `protobuf:"bytes,1,opt,name=pos_promotion,json=posPromotion,proto3" json:"pos_promotion,omitempty"`
}

func (x *CreatePosPromotionResponse) Reset() {
	*x = CreatePosPromotionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePosPromotionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePosPromotionResponse) ProtoMessage() {}

func (x *CreatePosPromotionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePosPromotionResponse.ProtoReflect.Descriptor instead.
func (*CreatePosPromotionResponse) Descriptor() ([]byte, []int) {
	return file_promotion_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePosPromotionResponse) GetPosPromotion() *PosPromotion {
	if x != nil {
		return x.PosPromotion
	}
	return nil
}

type ReadPosPromotionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PromotionId string      `protobuf:"bytes,1,opt,name=promotion_id,json=promotionId,proto3" json:"promotion_id,omitempty"`
	JwtPayload  *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *ReadPosPromotionRequest) Reset() {
	*x = ReadPosPromotionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPosPromotionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPosPromotionRequest) ProtoMessage() {}

func (x *ReadPosPromotionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPosPromotionRequest.ProtoReflect.Descriptor instead.
func (*ReadPosPromotionRequest) Descriptor() ([]byte, []int) {
	return file_promotion_proto_rawDescGZIP(), []int{3}
}

func (x *ReadPosPromotionRequest) GetPromotionId() string {
	if x != nil {
		return x.PromotionId
	}
	return ""
}

func (x *ReadPosPromotionRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type ReadPosPromotionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosPromotion *PosPromotion `protobuf:"bytes,1,opt,name=pos_promotion,json=posPromotion,proto3" json:"pos_promotion,omitempty"`
}

func (x *ReadPosPromotionResponse) Reset() {
	*x = ReadPosPromotionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPosPromotionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPosPromotionResponse) ProtoMessage() {}

func (x *ReadPosPromotionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPosPromotionResponse.ProtoReflect.Descriptor instead.
func (*ReadPosPromotionResponse) Descriptor() ([]byte, []int) {
	return file_promotion_proto_rawDescGZIP(), []int{4}
}

func (x *ReadPosPromotionResponse) GetPosPromotion() *PosPromotion {
	if x != nil {
		return x.PosPromotion
	}
	return nil
}

type UpdatePosPromotionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosPromotion *PosPromotion `protobuf:"bytes,1,opt,name=pos_promotion,json=posPromotion,proto3" json:"pos_promotion,omitempty"`
	JwtPayload   *JWTPayload   `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *UpdatePosPromotionRequest) Reset() {
	*x = UpdatePosPromotionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosPromotionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosPromotionRequest) ProtoMessage() {}

func (x *UpdatePosPromotionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosPromotionRequest.ProtoReflect.Descriptor instead.
func (*UpdatePosPromotionRequest) Descriptor() ([]byte, []int) {
	return file_promotion_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePosPromotionRequest) GetPosPromotion() *PosPromotion {
	if x != nil {
		return x.PosPromotion
	}
	return nil
}

func (x *UpdatePosPromotionRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type UpdatePosPromotionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosPromotion *PosPromotion `protobuf:"bytes,1,opt,name=pos_promotion,json=posPromotion,proto3" json:"pos_promotion,omitempty"`
}

func (x *UpdatePosPromotionResponse) Reset() {
	*x = UpdatePosPromotionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosPromotionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosPromotionResponse) ProtoMessage() {}

func (x *UpdatePosPromotionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosPromotionResponse.ProtoReflect.Descriptor instead.
func (*UpdatePosPromotionResponse) Descriptor() ([]byte, []int) {
	return file_promotion_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePosPromotionResponse) GetPosPromotion() *PosPromotion {
	if x != nil {
		return x.PosPromotion
	}
	return nil
}

type DeletePosPromotionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PromotionId string      `protobuf:"bytes,1,opt,name=promotion_id,json=promotionId,proto3" json:"promotion_id,omitempty"`
	JwtPayload  *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *DeletePosPromotionRequest) Reset() {
	*x = DeletePosPromotionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePosPromotionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePosPromotionRequest) ProtoMessage() {}

func (x *DeletePosPromotionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePosPromotionRequest.ProtoReflect.Descriptor instead.
func (*DeletePosPromotionRequest) Descriptor() ([]byte, []int) {
	return file_promotion_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePosPromotionRequest) GetPromotionId() string {
	if x != nil {
		return x.PromotionId
	}
	return ""
}

func (x *DeletePosPromotionRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type DeletePosPromotionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeletePosPromotionResponse) Reset() {
	*x = DeletePosPromotionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePosPromotionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePosPromotionResponse) ProtoMessage() {}

func (x *DeletePosPromotionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePosPromotionResponse.ProtoReflect.Descriptor instead.
func (*DeletePosPromotionResponse) Descriptor() ([]byte, []int) {
	return file_promotion_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePosPromotionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ReadAllPosPromotionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtPayload *JWTPayload `protobuf:"bytes,1,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *ReadAllPosPromotionsRequest) Reset() {
	*x = ReadAllPosPromotionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllPosPromotionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllPosPromotionsRequest) ProtoMessage() {}

func (x *ReadAllPosPromotionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllPosPromotionsRequest.ProtoReflect.Descriptor instead.
func (*ReadAllPosPromotionsRequest) Descriptor() ([]byte, []int) {
	return file_promotion_proto_rawDescGZIP(), []int{9}
}

func (x *ReadAllPosPromotionsRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type ReadAllPosPromotionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosPromotions []*PosPromotion `protobuf:"bytes,1,rep,name=pos_promotions,json=posPromotions,proto3" json:"pos_promotions,omitempty"`
}

func (x *ReadAllPosPromotionsResponse) Reset() {
	*x = ReadAllPosPromotionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllPosPromotionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllPosPromotionsResponse) ProtoMessage() {}

func (x *ReadAllPosPromotionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllPosPromotionsResponse.ProtoReflect.Descriptor instead.
func (*ReadAllPosPromotionsResponse) Descriptor() ([]byte, []int) {
	return file_promotion_proto_rawDescGZIP(), []int{10}
}

func (x *ReadAllPosPromotionsResponse) GetPosPromotions() []*PosPromotion {
	if x != nil {
		return x.PosPromotions
	}
	return nil
}

var File_promotion_proto protoreflect.FileDescriptor

var file_promotion_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x70, 0x6f, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x03, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x22, 0x85, 0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f,
	0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a,
	0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x54, 0x0a, 0x1a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x6e, 0x0a, 0x17, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x30,
	0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x52, 0x0a, 0x18, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0d,
	0x70, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x85, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f, 0x73, 0x2e,
	0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x70, 0x6f,
	0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77,
	0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x54, 0x0a, 0x1a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x70, 0x6f,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x70, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57,
	0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x36, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4f, 0x0a, 0x1b,
	0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0b, 0x6a,
	0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x58, 0x0a,
	0x1c, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x0e, 0x70, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x6f, 0x73, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xc8, 0x03, 0x0a, 0x13, 0x50, 0x6f, 0x73, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x55, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f,
	0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x6f, 0x73,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e,
	0x70, 0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x70, 0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x14, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c,
	0x50, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x2e,
	0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x50, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73,
	0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x6e, 0x64, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x69, 0x66, 0x62, 0x2f, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2d, 0x70, 0x6f, 0x73, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_promotion_proto_rawDescOnce sync.Once
	file_promotion_proto_rawDescData = file_promotion_proto_rawDesc
)

func file_promotion_proto_rawDescGZIP() []byte {
	file_promotion_proto_rawDescOnce.Do(func() {
		file_promotion_proto_rawDescData = protoimpl.X.CompressGZIP(file_promotion_proto_rawDescData)
	})
	return file_promotion_proto_rawDescData
}

var file_promotion_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_promotion_proto_goTypes = []interface{}{
	(*PosPromotion)(nil),                 // 0: pos.PosPromotion
	(*CreatePosPromotionRequest)(nil),    // 1: pos.CreatePosPromotionRequest
	(*CreatePosPromotionResponse)(nil),   // 2: pos.CreatePosPromotionResponse
	(*ReadPosPromotionRequest)(nil),      // 3: pos.ReadPosPromotionRequest
	(*ReadPosPromotionResponse)(nil),     // 4: pos.ReadPosPromotionResponse
	(*UpdatePosPromotionRequest)(nil),    // 5: pos.UpdatePosPromotionRequest
	(*UpdatePosPromotionResponse)(nil),   // 6: pos.UpdatePosPromotionResponse
	(*DeletePosPromotionRequest)(nil),    // 7: pos.DeletePosPromotionRequest
	(*DeletePosPromotionResponse)(nil),   // 8: pos.DeletePosPromotionResponse
	(*ReadAllPosPromotionsRequest)(nil),  // 9: pos.ReadAllPosPromotionsRequest
	(*ReadAllPosPromotionsResponse)(nil), // 10: pos.ReadAllPosPromotionsResponse
	(*timestamppb.Timestamp)(nil),        // 11: google.protobuf.Timestamp
	(*JWTPayload)(nil),                   // 12: pos.JWTPayload
}
var file_promotion_proto_depIdxs = []int32{
	11, // 0: pos.PosPromotion.start_date:type_name -> google.protobuf.Timestamp
	11, // 1: pos.PosPromotion.end_date:type_name -> google.protobuf.Timestamp
	11, // 2: pos.PosPromotion.created_at:type_name -> google.protobuf.Timestamp
	11, // 3: pos.PosPromotion.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 4: pos.CreatePosPromotionRequest.pos_promotion:type_name -> pos.PosPromotion
	12, // 5: pos.CreatePosPromotionRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 6: pos.CreatePosPromotionResponse.pos_promotion:type_name -> pos.PosPromotion
	12, // 7: pos.ReadPosPromotionRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 8: pos.ReadPosPromotionResponse.pos_promotion:type_name -> pos.PosPromotion
	0,  // 9: pos.UpdatePosPromotionRequest.pos_promotion:type_name -> pos.PosPromotion
	12, // 10: pos.UpdatePosPromotionRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 11: pos.UpdatePosPromotionResponse.pos_promotion:type_name -> pos.PosPromotion
	12, // 12: pos.DeletePosPromotionRequest.jwt_payload:type_name -> pos.JWTPayload
	12, // 13: pos.ReadAllPosPromotionsRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 14: pos.ReadAllPosPromotionsResponse.pos_promotions:type_name -> pos.PosPromotion
	1,  // 15: pos.PosPromotionService.CreatePosPromotion:input_type -> pos.CreatePosPromotionRequest
	3,  // 16: pos.PosPromotionService.ReadPosPromotion:input_type -> pos.ReadPosPromotionRequest
	5,  // 17: pos.PosPromotionService.UpdatePosPromotion:input_type -> pos.UpdatePosPromotionRequest
	7,  // 18: pos.PosPromotionService.DeletePosPromotion:input_type -> pos.DeletePosPromotionRequest
	9,  // 19: pos.PosPromotionService.ReadAllPosPromotions:input_type -> pos.ReadAllPosPromotionsRequest
	2,  // 20: pos.PosPromotionService.CreatePosPromotion:output_type -> pos.CreatePosPromotionResponse
	4,  // 21: pos.PosPromotionService.ReadPosPromotion:output_type -> pos.ReadPosPromotionResponse
	6,  // 22: pos.PosPromotionService.UpdatePosPromotion:output_type -> pos.UpdatePosPromotionResponse
	8,  // 23: pos.PosPromotionService.DeletePosPromotion:output_type -> pos.DeletePosPromotionResponse
	10, // 24: pos.PosPromotionService.ReadAllPosPromotions:output_type -> pos.ReadAllPosPromotionsResponse
	20, // [20:25] is the sub-list for method output_type
	15, // [15:20] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_promotion_proto_init() }
func file_promotion_proto_init() {
	if File_promotion_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_promotion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PosPromotion); i {
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
		file_promotion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePosPromotionRequest); i {
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
		file_promotion_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePosPromotionResponse); i {
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
		file_promotion_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPosPromotionRequest); i {
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
		file_promotion_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPosPromotionResponse); i {
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
		file_promotion_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosPromotionRequest); i {
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
		file_promotion_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosPromotionResponse); i {
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
		file_promotion_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePosPromotionRequest); i {
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
		file_promotion_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePosPromotionResponse); i {
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
		file_promotion_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllPosPromotionsRequest); i {
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
		file_promotion_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllPosPromotionsResponse); i {
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
			RawDescriptor: file_promotion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_promotion_proto_goTypes,
		DependencyIndexes: file_promotion_proto_depIdxs,
		MessageInfos:      file_promotion_proto_msgTypes,
	}.Build()
	File_promotion_proto = out.File
	file_promotion_proto_rawDesc = nil
	file_promotion_proto_goTypes = nil
	file_promotion_proto_depIdxs = nil
}
