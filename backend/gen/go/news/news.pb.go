// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.1
// source: news.proto

package news

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Запрос всех новостей
type GetNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNewsRequest) Reset() {
	*x = GetNewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsRequest) ProtoMessage() {}

func (x *GetNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsRequest.ProtoReflect.Descriptor instead.
func (*GetNewsRequest) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{0}
}

// Запрос новости по Id
type GetNewsByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetNewsByIdRequest) Reset() {
	*x = GetNewsByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewsByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsByIdRequest) ProtoMessage() {}

func (x *GetNewsByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsByIdRequest.ProtoReflect.Descriptor instead.
func (*GetNewsByIdRequest) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{1}
}

func (x *GetNewsByIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Запрос новостей по категориям
type GetNewsByCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories string `protobuf:"bytes,1,opt,name=categories,proto3" json:"categories,omitempty"`
}

func (x *GetNewsByCategoryRequest) Reset() {
	*x = GetNewsByCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewsByCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsByCategoryRequest) ProtoMessage() {}

func (x *GetNewsByCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsByCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetNewsByCategoryRequest) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{2}
}

func (x *GetNewsByCategoryRequest) GetCategories() string {
	if x != nil {
		return x.Categories
	}
	return ""
}

// Запрос на добавление новости
type AddNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title      string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text       string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Datetime   string `protobuf:"bytes,4,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Categories string `protobuf:"bytes,5,opt,name=categories,proto3" json:"categories,omitempty"`
}

func (x *AddNewsRequest) Reset() {
	*x = AddNewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNewsRequest) ProtoMessage() {}

func (x *AddNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNewsRequest.ProtoReflect.Descriptor instead.
func (*AddNewsRequest) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{3}
}

func (x *AddNewsRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddNewsRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddNewsRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *AddNewsRequest) GetDatetime() string {
	if x != nil {
		return x.Datetime
	}
	return ""
}

func (x *AddNewsRequest) GetCategories() string {
	if x != nil {
		return x.Categories
	}
	return ""
}

// Запрос на удаление новости
type DelNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DelNewsRequest) Reset() {
	*x = DelNewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelNewsRequest) ProtoMessage() {}

func (x *DelNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelNewsRequest.ProtoReflect.Descriptor instead.
func (*DelNewsRequest) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{4}
}

func (x *DelNewsRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Вывод всех/категоризированных новостей
type GetNewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	News map[string]*GetNewsItem `protobuf:"bytes,1,rep,name=news,proto3" json:"news,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetNewsResponse) Reset() {
	*x = GetNewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsResponse) ProtoMessage() {}

func (x *GetNewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsResponse.ProtoReflect.Descriptor instead.
func (*GetNewsResponse) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{5}
}

func (x *GetNewsResponse) GetNews() map[string]*GetNewsItem {
	if x != nil {
		return x.News
	}
	return nil
}

type GetNewsItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title      string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text       string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Datetime   string `protobuf:"bytes,4,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Categories string `protobuf:"bytes,5,opt,name=categories,proto3" json:"categories,omitempty"`
}

func (x *GetNewsItem) Reset() {
	*x = GetNewsItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewsItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsItem) ProtoMessage() {}

func (x *GetNewsItem) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsItem.ProtoReflect.Descriptor instead.
func (*GetNewsItem) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{6}
}

func (x *GetNewsItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetNewsItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetNewsItem) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *GetNewsItem) GetDatetime() string {
	if x != nil {
		return x.Datetime
	}
	return ""
}

func (x *GetNewsItem) GetCategories() string {
	if x != nil {
		return x.Categories
	}
	return ""
}

// Вывод новости по Id
type GetNewsByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text       string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Datetime   string `protobuf:"bytes,4,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Categories string `protobuf:"bytes,5,opt,name=categories,proto3" json:"categories,omitempty"`
}

func (x *GetNewsByIdResponse) Reset() {
	*x = GetNewsByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewsByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsByIdResponse) ProtoMessage() {}

func (x *GetNewsByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsByIdResponse.ProtoReflect.Descriptor instead.
func (*GetNewsByIdResponse) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{7}
}

func (x *GetNewsByIdResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetNewsByIdResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *GetNewsByIdResponse) GetDatetime() string {
	if x != nil {
		return x.Datetime
	}
	return ""
}

func (x *GetNewsByIdResponse) GetCategories() string {
	if x != nil {
		return x.Categories
	}
	return ""
}

// Добавление новости
type AddNewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *AddNewsResponse) Reset() {
	*x = AddNewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNewsResponse) ProtoMessage() {}

func (x *AddNewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNewsResponse.ProtoReflect.Descriptor instead.
func (*AddNewsResponse) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{8}
}

func (x *AddNewsResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddNewsResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

// Удаление новости
type DelNewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *DelNewsResponse) Reset() {
	*x = DelNewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelNewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelNewsResponse) ProtoMessage() {}

func (x *DelNewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_news_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelNewsResponse.ProtoReflect.Descriptor instead.
func (*DelNewsResponse) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{9}
}

func (x *DelNewsResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_news_proto protoreflect.FileDescriptor

var file_news_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x65,
	0x77, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x77, 0x73, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x20, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x92, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x1a, 0x4a, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x4e, 0x65, 0x77, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x83, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x7b, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x4e, 0x65, 0x77, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x33, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4e, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x23, 0x0a, 0x0f,
	0x44, 0x65, 0x6c, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72,
	0x72, 0x32, 0xbc, 0x03, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x45, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x14, 0x2e, 0x6e,
	0x65, 0x77, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x07, 0x12, 0x05, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x12, 0x59, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x77, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x69, 0x64, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x6d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x42, 0x79,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1e, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x7b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x7d, 0x12, 0x4c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x14, 0x2e,
	0x6e, 0x65, 0x77, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x61, 0x64, 0x64,
	0x12, 0x4e, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x14, 0x2e, 0x6e, 0x65,
	0x77, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x4e, 0x65, 0x77, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10,
	0x2a, 0x0e, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_news_proto_rawDescOnce sync.Once
	file_news_proto_rawDescData = file_news_proto_rawDesc
)

func file_news_proto_rawDescGZIP() []byte {
	file_news_proto_rawDescOnce.Do(func() {
		file_news_proto_rawDescData = protoimpl.X.CompressGZIP(file_news_proto_rawDescData)
	})
	return file_news_proto_rawDescData
}

var file_news_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_news_proto_goTypes = []interface{}{
	(*GetNewsRequest)(nil),           // 0: news.GetNewsRequest
	(*GetNewsByIdRequest)(nil),       // 1: news.GetNewsByIdRequest
	(*GetNewsByCategoryRequest)(nil), // 2: news.GetNewsByCategoryRequest
	(*AddNewsRequest)(nil),           // 3: news.AddNewsRequest
	(*DelNewsRequest)(nil),           // 4: news.DelNewsRequest
	(*GetNewsResponse)(nil),          // 5: news.GetNewsResponse
	(*GetNewsItem)(nil),              // 6: news.GetNewsItem
	(*GetNewsByIdResponse)(nil),      // 7: news.GetNewsByIdResponse
	(*AddNewsResponse)(nil),          // 8: news.AddNewsResponse
	(*DelNewsResponse)(nil),          // 9: news.DelNewsResponse
	nil,                              // 10: news.GetNewsResponse.NewsEntry
}
var file_news_proto_depIdxs = []int32{
	10, // 0: news.GetNewsResponse.news:type_name -> news.GetNewsResponse.NewsEntry
	6,  // 1: news.GetNewsResponse.NewsEntry.value:type_name -> news.GetNewsItem
	0,  // 2: news.NewsService.GetNews:input_type -> news.GetNewsRequest
	1,  // 3: news.NewsService.GetNewsById:input_type -> news.GetNewsByIdRequest
	2,  // 4: news.NewsService.GetNewsByCategory:input_type -> news.GetNewsByCategoryRequest
	3,  // 5: news.NewsService.AddNews:input_type -> news.AddNewsRequest
	4,  // 6: news.NewsService.DelNews:input_type -> news.DelNewsRequest
	5,  // 7: news.NewsService.GetNews:output_type -> news.GetNewsResponse
	7,  // 8: news.NewsService.GetNewsById:output_type -> news.GetNewsByIdResponse
	5,  // 9: news.NewsService.GetNewsByCategory:output_type -> news.GetNewsResponse
	8,  // 10: news.NewsService.AddNews:output_type -> news.AddNewsResponse
	9,  // 11: news.NewsService.DelNews:output_type -> news.DelNewsResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_news_proto_init() }
func file_news_proto_init() {
	if File_news_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_news_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewsRequest); i {
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
		file_news_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewsByIdRequest); i {
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
		file_news_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewsByCategoryRequest); i {
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
		file_news_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNewsRequest); i {
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
		file_news_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelNewsRequest); i {
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
		file_news_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewsResponse); i {
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
		file_news_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewsItem); i {
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
		file_news_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewsByIdResponse); i {
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
		file_news_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNewsResponse); i {
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
		file_news_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelNewsResponse); i {
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
			RawDescriptor: file_news_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_news_proto_goTypes,
		DependencyIndexes: file_news_proto_depIdxs,
		MessageInfos:      file_news_proto_msgTypes,
	}.Build()
	File_news_proto = out.File
	file_news_proto_rawDesc = nil
	file_news_proto_goTypes = nil
	file_news_proto_depIdxs = nil
}
