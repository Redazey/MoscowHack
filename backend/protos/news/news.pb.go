// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
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

type NewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NewsRequest) Reset() {
	*x = NewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsRequest) ProtoMessage() {}

func (x *NewsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use NewsRequest.ProtoReflect.Descriptor instead.
func (*NewsRequest) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{0}
}

func (x *NewsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type NewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	News map[string]*NewsItem `protobuf:"bytes,1,rep,name=news,proto3" json:"news,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NewsResponse) Reset() {
	*x = NewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsResponse) ProtoMessage() {}

func (x *NewsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use NewsResponse.ProtoReflect.Descriptor instead.
func (*NewsResponse) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{1}
}

func (x *NewsResponse) GetNews() map[string]*NewsItem {
	if x != nil {
		return x.News
	}
	return nil
}

type NewsItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewsItem map[string]*NewsContent `protobuf:"bytes,1,rep,name=newsItem,proto3" json:"newsItem,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NewsItem) Reset() {
	*x = NewsItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsItem) ProtoMessage() {}

func (x *NewsItem) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use NewsItem.ProtoReflect.Descriptor instead.
func (*NewsItem) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{2}
}

func (x *NewsItem) GetNewsItem() map[string]*NewsContent {
	if x != nil {
		return x.NewsItem
	}
	return nil
}

type NewsContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title      string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text       string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Datetime   string `protobuf:"bytes,4,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Categories string `protobuf:"bytes,5,opt,name=categories,proto3" json:"categories,omitempty"`
}

func (x *NewsContent) Reset() {
	*x = NewsContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_news_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsContent) ProtoMessage() {}

func (x *NewsContent) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use NewsContent.ProtoReflect.Descriptor instead.
func (*NewsContent) Descriptor() ([]byte, []int) {
	return file_news_proto_rawDescGZIP(), []int{3}
}

func (x *NewsContent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NewsContent) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewsContent) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *NewsContent) GetDatetime() string {
	if x != nil {
		return x.Datetime
	}
	return ""
}

func (x *NewsContent) GetCategories() string {
	if x != nil {
		return x.Categories
	}
	return ""
}

var File_news_proto protoreflect.FileDescriptor

var file_news_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x65,
	0x77, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1d, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x89, 0x01, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6e, 0x65,
	0x77, 0x73, 0x1a, 0x47, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x94, 0x01, 0x0a, 0x08,
	0x4e, 0x65, 0x77, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x38, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x73,
	0x49, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x77,
	0x73, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x49,
	0x74, 0x65, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x73, 0x49, 0x74,
	0x65, 0x6d, 0x1a, 0x4e, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x83, 0x01, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x32, 0x53, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x04, 0x4e, 0x65, 0x77, 0x73, 0x12,
	0x11, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d,
	0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x69, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_news_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_news_proto_goTypes = []interface{}{
	(*NewsRequest)(nil),  // 0: news.NewsRequest
	(*NewsResponse)(nil), // 1: news.NewsResponse
	(*NewsItem)(nil),     // 2: news.NewsItem
	(*NewsContent)(nil),  // 3: news.NewsContent
	nil,                  // 4: news.NewsResponse.NewsEntry
	nil,                  // 5: news.NewsItem.NewsItemEntry
}
var file_news_proto_depIdxs = []int32{
	4, // 0: news.NewsResponse.news:type_name -> news.NewsResponse.NewsEntry
	5, // 1: news.NewsItem.newsItem:type_name -> news.NewsItem.NewsItemEntry
	2, // 2: news.NewsResponse.NewsEntry.value:type_name -> news.NewsItem
	3, // 3: news.NewsItem.NewsItemEntry.value:type_name -> news.NewsContent
	0, // 4: news.NewsService.News:input_type -> news.NewsRequest
	1, // 5: news.NewsService.News:output_type -> news.NewsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_news_proto_init() }
func file_news_proto_init() {
	if File_news_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_news_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsRequest); i {
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
			switch v := v.(*NewsResponse); i {
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
			switch v := v.(*NewsItem); i {
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
			switch v := v.(*NewsContent); i {
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
			NumMessages:   6,
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
