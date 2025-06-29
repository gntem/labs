// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: proto/leaderboard.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListLeaderboardsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PageSize      int32                  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNumber    int32                  `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListLeaderboardsRequest) Reset() {
	*x = ListLeaderboardsRequest{}
	mi := &file_proto_leaderboard_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLeaderboardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLeaderboardsRequest) ProtoMessage() {}

func (x *ListLeaderboardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_leaderboard_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLeaderboardsRequest.ProtoReflect.Descriptor instead.
func (*ListLeaderboardsRequest) Descriptor() ([]byte, []int) {
	return file_proto_leaderboard_proto_rawDescGZIP(), []int{0}
}

func (x *ListLeaderboardsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListLeaderboardsRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

type ListLeaderboardsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Leaderboards  []*Leaderboard         `protobuf:"bytes,1,rep,name=leaderboards,proto3" json:"leaderboards,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	PageNumber    int32                  `protobuf:"varint,3,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListLeaderboardsResponse) Reset() {
	*x = ListLeaderboardsResponse{}
	mi := &file_proto_leaderboard_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLeaderboardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLeaderboardsResponse) ProtoMessage() {}

func (x *ListLeaderboardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_leaderboard_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLeaderboardsResponse.ProtoReflect.Descriptor instead.
func (*ListLeaderboardsResponse) Descriptor() ([]byte, []int) {
	return file_proto_leaderboard_proto_rawDescGZIP(), []int{1}
}

func (x *ListLeaderboardsResponse) GetLeaderboards() []*Leaderboard {
	if x != nil {
		return x.Leaderboards
	}
	return nil
}

func (x *ListLeaderboardsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListLeaderboardsResponse) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

type Leaderboard struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Score         int64                  `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Leaderboard) Reset() {
	*x = Leaderboard{}
	mi := &file_proto_leaderboard_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Leaderboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Leaderboard) ProtoMessage() {}

func (x *Leaderboard) ProtoReflect() protoreflect.Message {
	mi := &file_proto_leaderboard_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Leaderboard.ProtoReflect.Descriptor instead.
func (*Leaderboard) Descriptor() ([]byte, []int) {
	return file_proto_leaderboard_proto_rawDescGZIP(), []int{2}
}

func (x *Leaderboard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Leaderboard) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Leaderboard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Leaderboard) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_proto_leaderboard_proto protoreflect.FileDescriptor

const file_proto_leaderboard_proto_rawDesc = "" +
	"\n" +
	"\x17proto/leaderboard.proto\x12\vleaderboard\"W\n" +
	"\x17ListLeaderboardsRequest\x12\x1b\n" +
	"\tpage_size\x18\x01 \x01(\x05R\bpageSize\x12\x1f\n" +
	"\vpage_number\x18\x02 \x01(\x05R\n" +
	"pageNumber\"\x9a\x01\n" +
	"\x18ListLeaderboardsResponse\x12<\n" +
	"\fleaderboards\x18\x01 \x03(\v2\x18.leaderboard.LeaderboardR\fleaderboards\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\x12\x1f\n" +
	"\vpage_number\x18\x03 \x01(\x05R\n" +
	"pageNumber\"`\n" +
	"\vLeaderboard\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x14\n" +
	"\x05score\x18\x04 \x01(\x03R\x05score2w\n" +
	"\x12LeaderboardService\x12a\n" +
	"\x10ListLeaderboards\x12$.leaderboard.ListLeaderboardsRequest\x1a%.leaderboard.ListLeaderboardsResponse\"\x00B\"Z workshop4-grpc-leaderboard/protob\x06proto3"

var (
	file_proto_leaderboard_proto_rawDescOnce sync.Once
	file_proto_leaderboard_proto_rawDescData []byte
)

func file_proto_leaderboard_proto_rawDescGZIP() []byte {
	file_proto_leaderboard_proto_rawDescOnce.Do(func() {
		file_proto_leaderboard_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_leaderboard_proto_rawDesc), len(file_proto_leaderboard_proto_rawDesc)))
	})
	return file_proto_leaderboard_proto_rawDescData
}

var file_proto_leaderboard_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_leaderboard_proto_goTypes = []any{
	(*ListLeaderboardsRequest)(nil),  // 0: leaderboard.ListLeaderboardsRequest
	(*ListLeaderboardsResponse)(nil), // 1: leaderboard.ListLeaderboardsResponse
	(*Leaderboard)(nil),              // 2: leaderboard.Leaderboard
}
var file_proto_leaderboard_proto_depIdxs = []int32{
	2, // 0: leaderboard.ListLeaderboardsResponse.leaderboards:type_name -> leaderboard.Leaderboard
	0, // 1: leaderboard.LeaderboardService.ListLeaderboards:input_type -> leaderboard.ListLeaderboardsRequest
	1, // 2: leaderboard.LeaderboardService.ListLeaderboards:output_type -> leaderboard.ListLeaderboardsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_leaderboard_proto_init() }
func file_proto_leaderboard_proto_init() {
	if File_proto_leaderboard_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_leaderboard_proto_rawDesc), len(file_proto_leaderboard_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_leaderboard_proto_goTypes,
		DependencyIndexes: file_proto_leaderboard_proto_depIdxs,
		MessageInfos:      file_proto_leaderboard_proto_msgTypes,
	}.Build()
	File_proto_leaderboard_proto = out.File
	file_proto_leaderboard_proto_goTypes = nil
	file_proto_leaderboard_proto_depIdxs = nil
}
