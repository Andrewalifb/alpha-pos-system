// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: rewards.proto

package loyalty_service

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

// PosReward
type PosReward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RewardId       string                 `protobuf:"bytes,1,opt,name=reward_id,json=rewardId,proto3" json:"reward_id,omitempty"`
	RewardName     string                 `protobuf:"bytes,2,opt,name=reward_name,json=rewardName,proto3" json:"reward_name,omitempty"`
	PointsRequired int32                  `protobuf:"varint,3,opt,name=points_required,json=pointsRequired,proto3" json:"points_required,omitempty"`
	CompanyId      string                 `protobuf:"bytes,4,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy      string                 `protobuf:"bytes,6,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy      string                 `protobuf:"bytes,8,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
}

func (x *PosReward) Reset() {
	*x = PosReward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PosReward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PosReward) ProtoMessage() {}

func (x *PosReward) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PosReward.ProtoReflect.Descriptor instead.
func (*PosReward) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{0}
}

func (x *PosReward) GetRewardId() string {
	if x != nil {
		return x.RewardId
	}
	return ""
}

func (x *PosReward) GetRewardName() string {
	if x != nil {
		return x.RewardName
	}
	return ""
}

func (x *PosReward) GetPointsRequired() int32 {
	if x != nil {
		return x.PointsRequired
	}
	return 0
}

func (x *PosReward) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *PosReward) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PosReward) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *PosReward) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *PosReward) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

// Request and Response messages for PosReward
type CreatePosRewardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosReward  *PosReward  `protobuf:"bytes,1,opt,name=pos_reward,json=posReward,proto3" json:"pos_reward,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *CreatePosRewardRequest) Reset() {
	*x = CreatePosRewardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePosRewardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePosRewardRequest) ProtoMessage() {}

func (x *CreatePosRewardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePosRewardRequest.ProtoReflect.Descriptor instead.
func (*CreatePosRewardRequest) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePosRewardRequest) GetPosReward() *PosReward {
	if x != nil {
		return x.PosReward
	}
	return nil
}

func (x *CreatePosRewardRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type CreatePosRewardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosReward *PosReward `protobuf:"bytes,1,opt,name=pos_reward,json=posReward,proto3" json:"pos_reward,omitempty"`
}

func (x *CreatePosRewardResponse) Reset() {
	*x = CreatePosRewardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePosRewardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePosRewardResponse) ProtoMessage() {}

func (x *CreatePosRewardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePosRewardResponse.ProtoReflect.Descriptor instead.
func (*CreatePosRewardResponse) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePosRewardResponse) GetPosReward() *PosReward {
	if x != nil {
		return x.PosReward
	}
	return nil
}

type ReadPosRewardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RewardId   string      `protobuf:"bytes,1,opt,name=reward_id,json=rewardId,proto3" json:"reward_id,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *ReadPosRewardRequest) Reset() {
	*x = ReadPosRewardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPosRewardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPosRewardRequest) ProtoMessage() {}

func (x *ReadPosRewardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPosRewardRequest.ProtoReflect.Descriptor instead.
func (*ReadPosRewardRequest) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{3}
}

func (x *ReadPosRewardRequest) GetRewardId() string {
	if x != nil {
		return x.RewardId
	}
	return ""
}

func (x *ReadPosRewardRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type ReadPosRewardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosReward *PosReward `protobuf:"bytes,1,opt,name=pos_reward,json=posReward,proto3" json:"pos_reward,omitempty"`
}

func (x *ReadPosRewardResponse) Reset() {
	*x = ReadPosRewardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPosRewardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPosRewardResponse) ProtoMessage() {}

func (x *ReadPosRewardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPosRewardResponse.ProtoReflect.Descriptor instead.
func (*ReadPosRewardResponse) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{4}
}

func (x *ReadPosRewardResponse) GetPosReward() *PosReward {
	if x != nil {
		return x.PosReward
	}
	return nil
}

type UpdatePosRewardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosReward  *PosReward  `protobuf:"bytes,1,opt,name=pos_reward,json=posReward,proto3" json:"pos_reward,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *UpdatePosRewardRequest) Reset() {
	*x = UpdatePosRewardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosRewardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosRewardRequest) ProtoMessage() {}

func (x *UpdatePosRewardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosRewardRequest.ProtoReflect.Descriptor instead.
func (*UpdatePosRewardRequest) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePosRewardRequest) GetPosReward() *PosReward {
	if x != nil {
		return x.PosReward
	}
	return nil
}

func (x *UpdatePosRewardRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type UpdatePosRewardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosReward *PosReward `protobuf:"bytes,1,opt,name=pos_reward,json=posReward,proto3" json:"pos_reward,omitempty"`
}

func (x *UpdatePosRewardResponse) Reset() {
	*x = UpdatePosRewardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosRewardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosRewardResponse) ProtoMessage() {}

func (x *UpdatePosRewardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosRewardResponse.ProtoReflect.Descriptor instead.
func (*UpdatePosRewardResponse) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePosRewardResponse) GetPosReward() *PosReward {
	if x != nil {
		return x.PosReward
	}
	return nil
}

type DeletePosRewardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RewardId   string      `protobuf:"bytes,1,opt,name=reward_id,json=rewardId,proto3" json:"reward_id,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *DeletePosRewardRequest) Reset() {
	*x = DeletePosRewardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePosRewardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePosRewardRequest) ProtoMessage() {}

func (x *DeletePosRewardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePosRewardRequest.ProtoReflect.Descriptor instead.
func (*DeletePosRewardRequest) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePosRewardRequest) GetRewardId() string {
	if x != nil {
		return x.RewardId
	}
	return ""
}

func (x *DeletePosRewardRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type DeletePosRewardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeletePosRewardResponse) Reset() {
	*x = DeletePosRewardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePosRewardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePosRewardResponse) ProtoMessage() {}

func (x *DeletePosRewardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePosRewardResponse.ProtoReflect.Descriptor instead.
func (*DeletePosRewardResponse) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePosRewardResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ReadAllPosRewardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtPayload *JWTPayload `protobuf:"bytes,1,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *ReadAllPosRewardsRequest) Reset() {
	*x = ReadAllPosRewardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllPosRewardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllPosRewardsRequest) ProtoMessage() {}

func (x *ReadAllPosRewardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllPosRewardsRequest.ProtoReflect.Descriptor instead.
func (*ReadAllPosRewardsRequest) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{9}
}

func (x *ReadAllPosRewardsRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type ReadAllPosRewardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosRewards []*PosReward `protobuf:"bytes,1,rep,name=pos_rewards,json=posRewards,proto3" json:"pos_rewards,omitempty"`
}

func (x *ReadAllPosRewardsResponse) Reset() {
	*x = ReadAllPosRewardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllPosRewardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllPosRewardsResponse) ProtoMessage() {}

func (x *ReadAllPosRewardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllPosRewardsResponse.ProtoReflect.Descriptor instead.
func (*ReadAllPosRewardsResponse) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{10}
}

func (x *ReadAllPosRewardsResponse) GetPosRewards() []*PosReward {
	if x != nil {
		return x.PosRewards
	}
	return nil
}

var File_rewards_proto protoreflect.FileDescriptor

var file_rewards_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x70, 0x6f, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x79, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x5f, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x2e,
	0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e,
	0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x48, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x22, 0x65, 0x0a, 0x14, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73,
	0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x46, 0x0a, 0x15, 0x52, 0x65, 0x61, 0x64, 0x50,
	0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x22,
	0x79, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x70, 0x6f, 0x73,
	0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x09, 0x70,
	0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a,
	0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x48, 0x0a, 0x17, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x5f, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x2e,
	0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x22, 0x67, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x6a,
	0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x33, 0x0a,
	0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x4c, 0x0a, 0x18, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30,
	0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x4c, 0x0a, 0x19, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a,
	0x0b, 0x70, 0x6f, 0x73, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x32, 0x98,
	0x03, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x46, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x12, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x70,
	0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c,
	0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x73,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x73, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x64, 0x72, 0x65, 0x77, 0x61, 0x6c,
	0x69, 0x66, 0x62, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2d, 0x70, 0x6f, 0x73, 0x2d, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2f, 0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rewards_proto_rawDescOnce sync.Once
	file_rewards_proto_rawDescData = file_rewards_proto_rawDesc
)

func file_rewards_proto_rawDescGZIP() []byte {
	file_rewards_proto_rawDescOnce.Do(func() {
		file_rewards_proto_rawDescData = protoimpl.X.CompressGZIP(file_rewards_proto_rawDescData)
	})
	return file_rewards_proto_rawDescData
}

var file_rewards_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_rewards_proto_goTypes = []interface{}{
	(*PosReward)(nil),                 // 0: pos.PosReward
	(*CreatePosRewardRequest)(nil),    // 1: pos.CreatePosRewardRequest
	(*CreatePosRewardResponse)(nil),   // 2: pos.CreatePosRewardResponse
	(*ReadPosRewardRequest)(nil),      // 3: pos.ReadPosRewardRequest
	(*ReadPosRewardResponse)(nil),     // 4: pos.ReadPosRewardResponse
	(*UpdatePosRewardRequest)(nil),    // 5: pos.UpdatePosRewardRequest
	(*UpdatePosRewardResponse)(nil),   // 6: pos.UpdatePosRewardResponse
	(*DeletePosRewardRequest)(nil),    // 7: pos.DeletePosRewardRequest
	(*DeletePosRewardResponse)(nil),   // 8: pos.DeletePosRewardResponse
	(*ReadAllPosRewardsRequest)(nil),  // 9: pos.ReadAllPosRewardsRequest
	(*ReadAllPosRewardsResponse)(nil), // 10: pos.ReadAllPosRewardsResponse
	(*timestamppb.Timestamp)(nil),     // 11: google.protobuf.Timestamp
	(*JWTPayload)(nil),                // 12: pos.JWTPayload
}
var file_rewards_proto_depIdxs = []int32{
	11, // 0: pos.PosReward.created_at:type_name -> google.protobuf.Timestamp
	11, // 1: pos.PosReward.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: pos.CreatePosRewardRequest.pos_reward:type_name -> pos.PosReward
	12, // 3: pos.CreatePosRewardRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 4: pos.CreatePosRewardResponse.pos_reward:type_name -> pos.PosReward
	12, // 5: pos.ReadPosRewardRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 6: pos.ReadPosRewardResponse.pos_reward:type_name -> pos.PosReward
	0,  // 7: pos.UpdatePosRewardRequest.pos_reward:type_name -> pos.PosReward
	12, // 8: pos.UpdatePosRewardRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 9: pos.UpdatePosRewardResponse.pos_reward:type_name -> pos.PosReward
	12, // 10: pos.DeletePosRewardRequest.jwt_payload:type_name -> pos.JWTPayload
	12, // 11: pos.ReadAllPosRewardsRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 12: pos.ReadAllPosRewardsResponse.pos_rewards:type_name -> pos.PosReward
	1,  // 13: pos.PosRewardService.CreatePosReward:input_type -> pos.CreatePosRewardRequest
	3,  // 14: pos.PosRewardService.ReadPosReward:input_type -> pos.ReadPosRewardRequest
	5,  // 15: pos.PosRewardService.UpdatePosReward:input_type -> pos.UpdatePosRewardRequest
	7,  // 16: pos.PosRewardService.DeletePosReward:input_type -> pos.DeletePosRewardRequest
	9,  // 17: pos.PosRewardService.ReadAllPosRewards:input_type -> pos.ReadAllPosRewardsRequest
	2,  // 18: pos.PosRewardService.CreatePosReward:output_type -> pos.CreatePosRewardResponse
	4,  // 19: pos.PosRewardService.ReadPosReward:output_type -> pos.ReadPosRewardResponse
	6,  // 20: pos.PosRewardService.UpdatePosReward:output_type -> pos.UpdatePosRewardResponse
	8,  // 21: pos.PosRewardService.DeletePosReward:output_type -> pos.DeletePosRewardResponse
	10, // 22: pos.PosRewardService.ReadAllPosRewards:output_type -> pos.ReadAllPosRewardsResponse
	18, // [18:23] is the sub-list for method output_type
	13, // [13:18] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_rewards_proto_init() }
func file_rewards_proto_init() {
	if File_rewards_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rewards_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PosReward); i {
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
		file_rewards_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePosRewardRequest); i {
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
		file_rewards_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePosRewardResponse); i {
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
		file_rewards_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPosRewardRequest); i {
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
		file_rewards_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPosRewardResponse); i {
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
		file_rewards_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosRewardRequest); i {
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
		file_rewards_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosRewardResponse); i {
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
		file_rewards_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePosRewardRequest); i {
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
		file_rewards_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePosRewardResponse); i {
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
		file_rewards_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllPosRewardsRequest); i {
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
		file_rewards_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllPosRewardsResponse); i {
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
			RawDescriptor: file_rewards_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rewards_proto_goTypes,
		DependencyIndexes: file_rewards_proto_depIdxs,
		MessageInfos:      file_rewards_proto_msgTypes,
	}.Build()
	File_rewards_proto = out.File
	file_rewards_proto_rawDesc = nil
	file_rewards_proto_goTypes = nil
	file_rewards_proto_depIdxs = nil
}
