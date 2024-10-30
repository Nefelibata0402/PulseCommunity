// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: ranking_service.proto

package rankingGrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	article "pulseCommunity/idl/articleGrpc"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetTopNRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetTopNRequest) Reset() {
	*x = GetTopNRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ranking_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopNRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopNRequest) ProtoMessage() {}

func (x *GetTopNRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ranking_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopNRequest.ProtoReflect.Descriptor instead.
func (*GetTopNRequest) Descriptor() ([]byte, []int) {
	return file_ranking_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetTopNRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetTopNResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode  int32              `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg   string             `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	ArticleList []*article.Article `protobuf:"bytes,3,rep,name=article_list,json=articleList,proto3" json:"article_list,omitempty"`
}

func (x *GetTopNResponse) Reset() {
	*x = GetTopNResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ranking_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopNResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopNResponse) ProtoMessage() {}

func (x *GetTopNResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ranking_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopNResponse.ProtoReflect.Descriptor instead.
func (*GetTopNResponse) Descriptor() ([]byte, []int) {
	return file_ranking_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetTopNResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetTopNResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *GetTopNResponse) GetArticleList() []*article.Article {
	if x != nil {
		return x.ArticleList
	}
	return nil
}

type TopNRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *TopNRequest) Reset() {
	*x = TopNRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ranking_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopNRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopNRequest) ProtoMessage() {}

func (x *TopNRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ranking_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopNRequest.ProtoReflect.Descriptor instead.
func (*TopNRequest) Descriptor() ([]byte, []int) {
	return file_ranking_service_proto_rawDescGZIP(), []int{2}
}

func (x *TopNRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type TopNResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
}

func (x *TopNResponse) Reset() {
	*x = TopNResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ranking_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopNResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopNResponse) ProtoMessage() {}

func (x *TopNResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ranking_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopNResponse.ProtoReflect.Descriptor instead.
func (*TopNResponse) Descriptor() ([]byte, []int) {
	return file_ranking_service_proto_rawDescGZIP(), []int{3}
}

func (x *TopNResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *TopNResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

var File_ranking_service_proto protoreflect.FileDescriptor

var file_ranking_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x47, 0x72, 0x70, 0x63, 0x1a, 0x1f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x47, 0x72, 0x70,
	0x63, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x8a, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x73, 0x67, 0x12, 0x37, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x26, 0x0a,
	0x0b, 0x54, 0x6f, 0x70, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x0c, 0x54, 0x6f, 0x70, 0x4e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x73, 0x67, 0x32, 0x93, 0x01, 0x0a, 0x0e, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x70, 0x4e, 0x12, 0x1b, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x04, 0x54, 0x6f, 0x70, 0x4e, 0x12, 0x18, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x47, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x6f, 0x70, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x54,
	0x6f, 0x70, 0x4e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x72,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x69,
	0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ranking_service_proto_rawDescOnce sync.Once
	file_ranking_service_proto_rawDescData = file_ranking_service_proto_rawDesc
)

func file_ranking_service_proto_rawDescGZIP() []byte {
	file_ranking_service_proto_rawDescOnce.Do(func() {
		file_ranking_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_ranking_service_proto_rawDescData)
	})
	return file_ranking_service_proto_rawDescData
}

var file_ranking_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ranking_service_proto_goTypes = []interface{}{
	(*GetTopNRequest)(nil),  // 0: rankingGrpc.GetTopNRequest
	(*GetTopNResponse)(nil), // 1: rankingGrpc.GetTopNResponse
	(*TopNRequest)(nil),     // 2: rankingGrpc.TopNRequest
	(*TopNResponse)(nil),    // 3: rankingGrpc.TopNResponse
	(*article.Article)(nil), // 4: articleGrpc.Article
}
var file_ranking_service_proto_depIdxs = []int32{
	4, // 0: rankingGrpc.GetTopNResponse.article_list:type_name -> articleGrpc.Article
	0, // 1: rankingGrpc.RankingService.GetTopN:input_type -> rankingGrpc.GetTopNRequest
	2, // 2: rankingGrpc.RankingService.TopN:input_type -> rankingGrpc.TopNRequest
	1, // 3: rankingGrpc.RankingService.GetTopN:output_type -> rankingGrpc.GetTopNResponse
	3, // 4: rankingGrpc.RankingService.TopN:output_type -> rankingGrpc.TopNResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ranking_service_proto_init() }
func file_ranking_service_proto_init() {
	if File_ranking_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ranking_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopNRequest); i {
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
		file_ranking_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopNResponse); i {
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
		file_ranking_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopNRequest); i {
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
		file_ranking_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopNResponse); i {
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
			RawDescriptor: file_ranking_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ranking_service_proto_goTypes,
		DependencyIndexes: file_ranking_service_proto_depIdxs,
		MessageInfos:      file_ranking_service_proto_msgTypes,
	}.Build()
	File_ranking_service_proto = out.File
	file_ranking_service_proto_rawDesc = nil
	file_ranking_service_proto_goTypes = nil
	file_ranking_service_proto_depIdxs = nil
}
