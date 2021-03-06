// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/proto/getpreview.proto

package api

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

type PreviewUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoUrl string `protobuf:"bytes,1,opt,name=videoUrl,proto3" json:"videoUrl,omitempty"`
}

func (x *PreviewUrl) Reset() {
	*x = PreviewUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_getpreview_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreviewUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviewUrl) ProtoMessage() {}

func (x *PreviewUrl) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_getpreview_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviewUrl.ProtoReflect.Descriptor instead.
func (*PreviewUrl) Descriptor() ([]byte, []int) {
	return file_api_proto_getpreview_proto_rawDescGZIP(), []int{0}
}

func (x *PreviewUrl) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

type PreviewImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *PreviewImage) Reset() {
	*x = PreviewImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_getpreview_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreviewImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviewImage) ProtoMessage() {}

func (x *PreviewImage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_getpreview_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviewImage.ProtoReflect.Descriptor instead.
func (*PreviewImage) Descriptor() ([]byte, []int) {
	return file_api_proto_getpreview_proto_rawDescGZIP(), []int{1}
}

func (x *PreviewImage) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_api_proto_getpreview_proto protoreflect.FileDescriptor

var file_api_proto_getpreview_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x74, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0a,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x55, 0x72, 0x6c, 0x22, 0x24, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x32, 0x37, 0x0a, 0x09,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0b, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x55, 0x72, 0x6c, 0x1a, 0x0d, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_getpreview_proto_rawDescOnce sync.Once
	file_api_proto_getpreview_proto_rawDescData = file_api_proto_getpreview_proto_rawDesc
)

func file_api_proto_getpreview_proto_rawDescGZIP() []byte {
	file_api_proto_getpreview_proto_rawDescOnce.Do(func() {
		file_api_proto_getpreview_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_getpreview_proto_rawDescData)
	})
	return file_api_proto_getpreview_proto_rawDescData
}

var file_api_proto_getpreview_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_getpreview_proto_goTypes = []interface{}{
	(*PreviewUrl)(nil),   // 0: PreviewUrl
	(*PreviewImage)(nil), // 1: PreviewImage
}
var file_api_proto_getpreview_proto_depIdxs = []int32{
	0, // 0: Previewer.GetPreview:input_type -> PreviewUrl
	1, // 1: Previewer.GetPreview:output_type -> PreviewImage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_getpreview_proto_init() }
func file_api_proto_getpreview_proto_init() {
	if File_api_proto_getpreview_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_getpreview_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreviewUrl); i {
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
		file_api_proto_getpreview_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreviewImage); i {
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
			RawDescriptor: file_api_proto_getpreview_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_getpreview_proto_goTypes,
		DependencyIndexes: file_api_proto_getpreview_proto_depIdxs,
		MessageInfos:      file_api_proto_getpreview_proto_msgTypes,
	}.Build()
	File_api_proto_getpreview_proto = out.File
	file_api_proto_getpreview_proto_rawDesc = nil
	file_api_proto_getpreview_proto_goTypes = nil
	file_api_proto_getpreview_proto_depIdxs = nil
}
