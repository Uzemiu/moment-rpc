// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: moment.proto

package pb

import (
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

type Moment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	CatId     string   `protobuf:"bytes,3,opt,name=catId,proto3" json:"catId,omitempty"`
	Text      string   `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	ImageUrls []string `protobuf:"bytes,5,rep,name=imageUrls,proto3" json:"imageUrls,omitempty"`
	UserId    string   `protobuf:"bytes,6,opt,name=userId,proto3" json:"userId,omitempty"`
	CreateAt  int64    `protobuf:"varint,7,opt,name=createAt,proto3" json:"createAt,omitempty"`
	Likes     int32    `protobuf:"varint,8,opt,name=likes,proto3" json:"likes,omitempty"`
}

func (x *Moment) Reset() {
	*x = Moment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Moment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Moment) ProtoMessage() {}

func (x *Moment) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Moment.ProtoReflect.Descriptor instead.
func (*Moment) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{0}
}

func (x *Moment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Moment) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Moment) GetCatId() string {
	if x != nil {
		return x.CatId
	}
	return ""
}

func (x *Moment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Moment) GetImageUrls() []string {
	if x != nil {
		return x.ImageUrls
	}
	return nil
}

func (x *Moment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Moment) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Moment) GetLikes() int32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

type GetMomentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMomentReq) Reset() {
	*x = GetMomentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMomentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMomentReq) ProtoMessage() {}

func (x *GetMomentReq) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMomentReq.ProtoReflect.Descriptor instead.
func (*GetMomentReq) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{1}
}

func (x *GetMomentReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMomentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Moment *Moment `protobuf:"bytes,1,opt,name=moment,proto3" json:"moment,omitempty"`
}

func (x *GetMomentResp) Reset() {
	*x = GetMomentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMomentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMomentResp) ProtoMessage() {}

func (x *GetMomentResp) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMomentResp.ProtoReflect.Descriptor instead.
func (*GetMomentResp) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{2}
}

func (x *GetMomentResp) GetMoment() *Moment {
	if x != nil {
		return x.Moment
	}
	return nil
}

type QueryMomentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *QueryMomentReq) Reset() {
	*x = QueryMomentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMomentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMomentReq) ProtoMessage() {}

func (x *QueryMomentReq) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMomentReq.ProtoReflect.Descriptor instead.
func (*QueryMomentReq) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{3}
}

func (x *QueryMomentReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *QueryMomentReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type QueryMomentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Moments []*Moment `protobuf:"bytes,1,rep,name=moments,proto3" json:"moments,omitempty"`
}

func (x *QueryMomentResp) Reset() {
	*x = QueryMomentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMomentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMomentResp) ProtoMessage() {}

func (x *QueryMomentResp) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMomentResp.ProtoReflect.Descriptor instead.
func (*QueryMomentResp) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{4}
}

func (x *QueryMomentResp) GetMoments() []*Moment {
	if x != nil {
		return x.Moments
	}
	return nil
}

type InsertMomentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	CatId     string   `protobuf:"bytes,2,opt,name=catId,proto3" json:"catId,omitempty"`
	Text      string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	ImageUrls []string `protobuf:"bytes,4,rep,name=imageUrls,proto3" json:"imageUrls,omitempty"`
	UserId    string   `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *InsertMomentReq) Reset() {
	*x = InsertMomentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertMomentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertMomentReq) ProtoMessage() {}

func (x *InsertMomentReq) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertMomentReq.ProtoReflect.Descriptor instead.
func (*InsertMomentReq) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{5}
}

func (x *InsertMomentReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *InsertMomentReq) GetCatId() string {
	if x != nil {
		return x.CatId
	}
	return ""
}

func (x *InsertMomentReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *InsertMomentReq) GetImageUrls() []string {
	if x != nil {
		return x.ImageUrls
	}
	return nil
}

func (x *InsertMomentReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type InsertMomentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InsertMomentResp) Reset() {
	*x = InsertMomentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertMomentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertMomentResp) ProtoMessage() {}

func (x *InsertMomentResp) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertMomentResp.ProtoReflect.Descriptor instead.
func (*InsertMomentResp) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{6}
}

func (x *InsertMomentResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateMomentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	CatId     string   `protobuf:"bytes,3,opt,name=catId,proto3" json:"catId,omitempty"`
	Text      string   `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	ImageUrls []string `protobuf:"bytes,5,rep,name=imageUrls,proto3" json:"imageUrls,omitempty"`
	// 考虑到一般情况下用户只能更新自己的动态，增加userId字段
	// 如果userId为空，则更新时不会校验用户id
	UserId string `protobuf:"bytes,6,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UpdateMomentReq) Reset() {
	*x = UpdateMomentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMomentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMomentReq) ProtoMessage() {}

func (x *UpdateMomentReq) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMomentReq.ProtoReflect.Descriptor instead.
func (*UpdateMomentReq) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateMomentReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateMomentReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateMomentReq) GetCatId() string {
	if x != nil {
		return x.CatId
	}
	return ""
}

func (x *UpdateMomentReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *UpdateMomentReq) GetImageUrls() []string {
	if x != nil {
		return x.ImageUrls
	}
	return nil
}

func (x *UpdateMomentReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdateMomentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateMomentResp) Reset() {
	*x = UpdateMomentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMomentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMomentResp) ProtoMessage() {}

func (x *UpdateMomentResp) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMomentResp.ProtoReflect.Descriptor instead.
func (*UpdateMomentResp) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{8}
}

type DeleteMomentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 考虑到一般情况下用户只能删除自己的动态，增加userId字段
	// 如果userId为空，则删除时不会校验用户id
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *DeleteMomentReq) Reset() {
	*x = DeleteMomentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMomentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMomentReq) ProtoMessage() {}

func (x *DeleteMomentReq) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMomentReq.ProtoReflect.Descriptor instead.
func (*DeleteMomentReq) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteMomentReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteMomentReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeleteMomentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteMomentResp) Reset() {
	*x = DeleteMomentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moment_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMomentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMomentResp) ProtoMessage() {}

func (x *DeleteMomentResp) ProtoReflect() protoreflect.Message {
	mi := &file_moment_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMomentResp.ProtoReflect.Descriptor instead.
func (*DeleteMomentResp) Descriptor() ([]byte, []int) {
	return file_moment_proto_rawDescGZIP(), []int{10}
}

var File_moment_proto protoreflect.FileDescriptor

var file_moment_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xc0, 0x01, 0x0a, 0x06, 0x4d, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x22, 0x1e, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x26, 0x0a, 0x06, 0x6d, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x6d, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x38, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x3b, 0x0a, 0x0f,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x28, 0x0a, 0x07, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x07, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x0f, 0x49, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4d, 0x6f, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x39, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x32, 0xce, 0x02, 0x0a, 0x09, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x70, 0x63, 0x12, 0x38, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x14, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3e, 0x0a, 0x0b,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x6d, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x0c,
	0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x6d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4d, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x41, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x17, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x41, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x17, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x6d, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moment_proto_rawDescOnce sync.Once
	file_moment_proto_rawDescData = file_moment_proto_rawDesc
)

func file_moment_proto_rawDescGZIP() []byte {
	file_moment_proto_rawDescOnce.Do(func() {
		file_moment_proto_rawDescData = protoimpl.X.CompressGZIP(file_moment_proto_rawDescData)
	})
	return file_moment_proto_rawDescData
}

var file_moment_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_moment_proto_goTypes = []interface{}{
	(*Moment)(nil),           // 0: moment.Moment
	(*GetMomentReq)(nil),     // 1: moment.GetMomentReq
	(*GetMomentResp)(nil),    // 2: moment.GetMomentResp
	(*QueryMomentReq)(nil),   // 3: moment.QueryMomentReq
	(*QueryMomentResp)(nil),  // 4: moment.QueryMomentResp
	(*InsertMomentReq)(nil),  // 5: moment.InsertMomentReq
	(*InsertMomentResp)(nil), // 6: moment.InsertMomentResp
	(*UpdateMomentReq)(nil),  // 7: moment.UpdateMomentReq
	(*UpdateMomentResp)(nil), // 8: moment.UpdateMomentResp
	(*DeleteMomentReq)(nil),  // 9: moment.DeleteMomentReq
	(*DeleteMomentResp)(nil), // 10: moment.DeleteMomentResp
}
var file_moment_proto_depIdxs = []int32{
	0,  // 0: moment.GetMomentResp.moment:type_name -> moment.Moment
	0,  // 1: moment.QueryMomentResp.moments:type_name -> moment.Moment
	1,  // 2: moment.MomentRpc.getMoment:input_type -> moment.GetMomentReq
	3,  // 3: moment.MomentRpc.queryMoment:input_type -> moment.QueryMomentReq
	5,  // 4: moment.MomentRpc.insertMoment:input_type -> moment.InsertMomentReq
	7,  // 5: moment.MomentRpc.updateMoment:input_type -> moment.UpdateMomentReq
	9,  // 6: moment.MomentRpc.deleteMoment:input_type -> moment.DeleteMomentReq
	2,  // 7: moment.MomentRpc.getMoment:output_type -> moment.GetMomentResp
	4,  // 8: moment.MomentRpc.queryMoment:output_type -> moment.QueryMomentResp
	6,  // 9: moment.MomentRpc.insertMoment:output_type -> moment.InsertMomentResp
	8,  // 10: moment.MomentRpc.updateMoment:output_type -> moment.UpdateMomentResp
	10, // 11: moment.MomentRpc.deleteMoment:output_type -> moment.DeleteMomentResp
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_moment_proto_init() }
func file_moment_proto_init() {
	if File_moment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Moment); i {
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
		file_moment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMomentReq); i {
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
		file_moment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMomentResp); i {
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
		file_moment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMomentReq); i {
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
		file_moment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMomentResp); i {
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
		file_moment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertMomentReq); i {
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
		file_moment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertMomentResp); i {
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
		file_moment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMomentReq); i {
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
		file_moment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMomentResp); i {
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
		file_moment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMomentReq); i {
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
		file_moment_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMomentResp); i {
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
			RawDescriptor: file_moment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_moment_proto_goTypes,
		DependencyIndexes: file_moment_proto_depIdxs,
		MessageInfos:      file_moment_proto_msgTypes,
	}.Build()
	File_moment_proto = out.File
	file_moment_proto_rawDesc = nil
	file_moment_proto_goTypes = nil
	file_moment_proto_depIdxs = nil
}
