// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: pesan.proto

package pesan_backend

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type CredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserHandle  string  `protobuf:"bytes,1,opt,name=user_handle,json=userHandle,proto3" json:"user_handle,omitempty"`
	DisplayName *string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3,oneof" json:"display_name,omitempty"`
}

func (x *CredentialRequest) Reset() {
	*x = CredentialRequest{}
	mi := &file_pesan_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialRequest) ProtoMessage() {}

func (x *CredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pesan_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialRequest.ProtoReflect.Descriptor instead.
func (*CredentialRequest) Descriptor() ([]byte, []int) {
	return file_pesan_proto_rawDescGZIP(), []int{0}
}

func (x *CredentialRequest) GetUserHandle() string {
	if x != nil {
		return x.UserHandle
	}
	return ""
}

func (x *CredentialRequest) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

type ChallengeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Challenge []byte `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
}

func (x *ChallengeReply) Reset() {
	*x = ChallengeReply{}
	mi := &file_pesan_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChallengeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeReply) ProtoMessage() {}

func (x *ChallengeReply) ProtoReflect() protoreflect.Message {
	mi := &file_pesan_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeReply.ProtoReflect.Descriptor instead.
func (*ChallengeReply) Descriptor() ([]byte, []int) {
	return file_pesan_proto_rawDescGZIP(), []int{1}
}

func (x *ChallengeReply) GetChallenge() []byte {
	if x != nil {
		return x.Challenge
	}
	return nil
}

type AssertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signed []byte `protobuf:"bytes,1,opt,name=signed,proto3" json:"signed,omitempty"`
}

func (x *AssertRequest) Reset() {
	*x = AssertRequest{}
	mi := &file_pesan_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssertRequest) ProtoMessage() {}

func (x *AssertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pesan_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssertRequest.ProtoReflect.Descriptor instead.
func (*AssertRequest) Descriptor() ([]byte, []int) {
	return file_pesan_proto_rawDescGZIP(), []int{2}
}

func (x *AssertRequest) GetSigned() []byte {
	if x != nil {
		return x.Signed
	}
	return nil
}

type AssertReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  *string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3,oneof" json:"access_token,omitempty"`
	RefreshToken *string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3,oneof" json:"refresh_token,omitempty"`
}

func (x *AssertReply) Reset() {
	*x = AssertReply{}
	mi := &file_pesan_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssertReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssertReply) ProtoMessage() {}

func (x *AssertReply) ProtoReflect() protoreflect.Message {
	mi := &file_pesan_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssertReply.ProtoReflect.Descriptor instead.
func (*AssertReply) Descriptor() ([]byte, []int) {
	return file_pesan_proto_rawDescGZIP(), []int{3}
}

func (x *AssertReply) GetAccessToken() string {
	if x != nil && x.AccessToken != nil {
		return *x.AccessToken
	}
	return ""
}

func (x *AssertReply) GetRefreshToken() string {
	if x != nil && x.RefreshToken != nil {
		return *x.RefreshToken
	}
	return ""
}

type NewProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description *string     `protobuf:"bytes,2,opt,name=description,proto3,oneof" json:"description,omitempty"`
	UnitLabel   string      `protobuf:"bytes,3,opt,name=unit_label,json=unitLabel,proto3" json:"unit_label,omitempty"`
	UnitPrice   float64     `protobuf:"fixed64,4,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	Categories  []*Category `protobuf:"bytes,5,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *NewProductRequest) Reset() {
	*x = NewProductRequest{}
	mi := &file_pesan_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewProductRequest) ProtoMessage() {}

func (x *NewProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pesan_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewProductRequest.ProtoReflect.Descriptor instead.
func (*NewProductRequest) Descriptor() ([]byte, []int) {
	return file_pesan_proto_rawDescGZIP(), []int{4}
}

func (x *NewProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewProductRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *NewProductRequest) GetUnitLabel() string {
	if x != nil {
		return x.UnitLabel
	}
	return ""
}

func (x *NewProductRequest) GetUnitPrice() float64 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

func (x *NewProductRequest) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// for existing category selection
	CategoryId *string `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3,oneof" json:"category_id,omitempty"`
	// for new category
	Name            *string                `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Description     *string                `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	AvailableFrom   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=available_from,json=availableFrom,proto3,oneof" json:"available_from,omitempty"`
	AvailableUntil  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=available_until,json=availableUntil,proto3,oneof" json:"available_until,omitempty"`
	AvailableWeekly []string               `protobuf:"bytes,6,rep,name=available_weekly,json=availableWeekly,proto3" json:"available_weekly,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	mi := &file_pesan_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_pesan_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_pesan_proto_rawDescGZIP(), []int{5}
}

func (x *Category) GetCategoryId() string {
	if x != nil && x.CategoryId != nil {
		return *x.CategoryId
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Category) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Category) GetAvailableFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.AvailableFrom
	}
	return nil
}

func (x *Category) GetAvailableUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.AvailableUntil
	}
	return nil
}

func (x *Category) GetAvailableWeekly() []string {
	if x != nil {
		return x.AvailableWeekly
	}
	return nil
}

type NewProductReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewProductId string `protobuf:"bytes,1,opt,name=new_product_id,json=newProductId,proto3" json:"new_product_id,omitempty"`
}

func (x *NewProductReply) Reset() {
	*x = NewProductReply{}
	mi := &file_pesan_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewProductReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewProductReply) ProtoMessage() {}

func (x *NewProductReply) ProtoReflect() protoreflect.Message {
	mi := &file_pesan_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewProductReply.ProtoReflect.Descriptor instead.
func (*NewProductReply) Descriptor() ([]byte, []int) {
	return file_pesan_proto_rawDescGZIP(), []int{6}
}

func (x *NewProductReply) GetNewProductId() string {
	if x != nil {
		return x.NewProductId
	}
	return ""
}

type NewPhoto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *NewPhoto) Reset() {
	*x = NewPhoto{}
	mi := &file_pesan_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewPhoto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPhoto) ProtoMessage() {}

func (x *NewPhoto) ProtoReflect() protoreflect.Message {
	mi := &file_pesan_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPhoto.ProtoReflect.Descriptor instead.
func (*NewPhoto) Descriptor() ([]byte, []int) {
	return file_pesan_proto_rawDescGZIP(), []int{7}
}

func (x *NewPhoto) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

var File_pesan_proto protoreflect.FileDescriptor

var file_pesan_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x65, 0x73, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x65, 0x73, 0x61, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x2e, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x22, 0x27, 0x0a, 0x0d, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x0b, 0x41,
	0x73, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0xcd, 0x01, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2f,
	0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x65, 0x73, 0x61, 0x6e, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xfd, 0x02, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x0b,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x46, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x03, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a, 0x0f, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48,
	0x04, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x6e, 0x74, 0x69,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x42, 0x12, 0x0a, 0x10, 0x5f,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x22,
	0x37, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0x8a, 0x02, 0x0a, 0x05, 0x50,
	0x65, 0x73, 0x61, 0x6e, 0x12, 0x3a, 0x0a, 0x07, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12,
	0x18, 0x2e, 0x70, 0x65, 0x73, 0x61, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x65, 0x73, 0x61,
	0x6e, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x3d, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x2e, 0x70, 0x65, 0x73, 0x61, 0x6e, 0x2e, 0x41, 0x73,
	0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x65,
	0x73, 0x61, 0x6e, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x44, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x65, 0x73, 0x61, 0x6e, 0x2e, 0x4e, 0x65, 0x77, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x65, 0x73, 0x61, 0x6e, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x0f, 0x2e, 0x70,
	0x65, 0x73, 0x61, 0x6e, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x28, 0x01, 0x42, 0x50, 0x0a, 0x1f, 0x6d, 0x79, 0x2e, 0x6d, 0x69,
	0x6c, 0x64, 0x6f, 0x74, 0x64, 0x65, 0x76, 0x2e, 0x70, 0x65, 0x73, 0x61, 0x6e, 0x2e, 0x61, 0x6e,
	0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x75, 0x62, 0x42, 0x0a, 0x50, 0x65, 0x73, 0x61,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x6c, 0x74, 0x73, 0x6d, 0x2f, 0x70, 0x65, 0x73, 0x61,
	0x6e, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pesan_proto_rawDescOnce sync.Once
	file_pesan_proto_rawDescData = file_pesan_proto_rawDesc
)

func file_pesan_proto_rawDescGZIP() []byte {
	file_pesan_proto_rawDescOnce.Do(func() {
		file_pesan_proto_rawDescData = protoimpl.X.CompressGZIP(file_pesan_proto_rawDescData)
	})
	return file_pesan_proto_rawDescData
}

var file_pesan_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pesan_proto_goTypes = []any{
	(*CredentialRequest)(nil),     // 0: pesan.CredentialRequest
	(*ChallengeReply)(nil),        // 1: pesan.ChallengeReply
	(*AssertRequest)(nil),         // 2: pesan.AssertRequest
	(*AssertReply)(nil),           // 3: pesan.AssertReply
	(*NewProductRequest)(nil),     // 4: pesan.NewProductRequest
	(*Category)(nil),              // 5: pesan.Category
	(*NewProductReply)(nil),       // 6: pesan.NewProductReply
	(*NewPhoto)(nil),              // 7: pesan.NewPhoto
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_pesan_proto_depIdxs = []int32{
	5, // 0: pesan.NewProductRequest.categories:type_name -> pesan.Category
	8, // 1: pesan.Category.available_from:type_name -> google.protobuf.Timestamp
	8, // 2: pesan.Category.available_until:type_name -> google.protobuf.Timestamp
	0, // 3: pesan.Pesan.Onboard:input_type -> pesan.CredentialRequest
	2, // 4: pesan.Pesan.RegisterPublicKey:input_type -> pesan.AssertRequest
	4, // 5: pesan.Pesan.CreateNewProduct:input_type -> pesan.NewProductRequest
	7, // 6: pesan.Pesan.UploadProductPhotos:input_type -> pesan.NewPhoto
	1, // 7: pesan.Pesan.Onboard:output_type -> pesan.ChallengeReply
	3, // 8: pesan.Pesan.RegisterPublicKey:output_type -> pesan.AssertReply
	6, // 9: pesan.Pesan.CreateNewProduct:output_type -> pesan.NewProductReply
	9, // 10: pesan.Pesan.UploadProductPhotos:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pesan_proto_init() }
func file_pesan_proto_init() {
	if File_pesan_proto != nil {
		return
	}
	file_pesan_proto_msgTypes[0].OneofWrappers = []any{}
	file_pesan_proto_msgTypes[3].OneofWrappers = []any{}
	file_pesan_proto_msgTypes[4].OneofWrappers = []any{}
	file_pesan_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pesan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pesan_proto_goTypes,
		DependencyIndexes: file_pesan_proto_depIdxs,
		MessageInfos:      file_pesan_proto_msgTypes,
	}.Build()
	File_pesan_proto = out.File
	file_pesan_proto_rawDesc = nil
	file_pesan_proto_goTypes = nil
	file_pesan_proto_depIdxs = nil
}
