// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: instagram/instagram.proto

package instagram

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

type GetFollowersInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetFollowersInfoRequest) Reset() {
	*x = GetFollowersInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instagram_instagram_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowersInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowersInfoRequest) ProtoMessage() {}

func (x *GetFollowersInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_instagram_instagram_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowersInfoRequest.ProtoReflect.Descriptor instead.
func (*GetFollowersInfoRequest) Descriptor() ([]byte, []int) {
	return file_instagram_instagram_proto_rawDescGZIP(), []int{0}
}

func (x *GetFollowersInfoRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetFollowersInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Followers int32  `protobuf:"varint,2,opt,name=followers,proto3" json:"followers,omitempty"`
	Following int32  `protobuf:"varint,3,opt,name=following,proto3" json:"following,omitempty"`
}

func (x *GetFollowersInfoResponse) Reset() {
	*x = GetFollowersInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_instagram_instagram_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowersInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowersInfoResponse) ProtoMessage() {}

func (x *GetFollowersInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_instagram_instagram_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowersInfoResponse.ProtoReflect.Descriptor instead.
func (*GetFollowersInfoResponse) Descriptor() ([]byte, []int) {
	return file_instagram_instagram_proto_rawDescGZIP(), []int{1}
}

func (x *GetFollowersInfoResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetFollowersInfoResponse) GetFollowers() int32 {
	if x != nil {
		return x.Followers
	}
	return 0
}

func (x *GetFollowersInfoResponse) GetFollowing() int32 {
	if x != nil {
		return x.Following
	}
	return 0
}

var File_instagram_instagram_proto protoreflect.FileDescriptor

var file_instagram_instagram_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x22, 0x35, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x72, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x32, 0x68, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x5b,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x22, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72,
	0x61, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x23, 0x5a, 0x21, 0x67,
	0x72, 0x70, 0x63, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x63,
	0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_instagram_instagram_proto_rawDescOnce sync.Once
	file_instagram_instagram_proto_rawDescData = file_instagram_instagram_proto_rawDesc
)

func file_instagram_instagram_proto_rawDescGZIP() []byte {
	file_instagram_instagram_proto_rawDescOnce.Do(func() {
		file_instagram_instagram_proto_rawDescData = protoimpl.X.CompressGZIP(file_instagram_instagram_proto_rawDescData)
	})
	return file_instagram_instagram_proto_rawDescData
}

var file_instagram_instagram_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_instagram_instagram_proto_goTypes = []interface{}{
	(*GetFollowersInfoRequest)(nil),  // 0: instagram.GetFollowersInfoRequest
	(*GetFollowersInfoResponse)(nil), // 1: instagram.GetFollowersInfoResponse
}
var file_instagram_instagram_proto_depIdxs = []int32{
	0, // 0: instagram.Instagram.GetFollowersInfo:input_type -> instagram.GetFollowersInfoRequest
	1, // 1: instagram.Instagram.GetFollowersInfo:output_type -> instagram.GetFollowersInfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_instagram_instagram_proto_init() }
func file_instagram_instagram_proto_init() {
	if File_instagram_instagram_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_instagram_instagram_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowersInfoRequest); i {
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
		file_instagram_instagram_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowersInfoResponse); i {
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
			RawDescriptor: file_instagram_instagram_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_instagram_instagram_proto_goTypes,
		DependencyIndexes: file_instagram_instagram_proto_depIdxs,
		MessageInfos:      file_instagram_instagram_proto_msgTypes,
	}.Build()
	File_instagram_instagram_proto = out.File
	file_instagram_instagram_proto_rawDesc = nil
	file_instagram_instagram_proto_goTypes = nil
	file_instagram_instagram_proto_depIdxs = nil
}
