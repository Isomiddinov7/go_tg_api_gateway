// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: images.proto

package coins_service

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

type ImagePrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ImagePrimaryKey) Reset() {
	*x = ImagePrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_images_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImagePrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImagePrimaryKey) ProtoMessage() {}

func (x *ImagePrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_images_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImagePrimaryKey.ProtoReflect.Descriptor instead.
func (*ImagePrimaryKey) Descriptor() ([]byte, []int) {
	return file_images_proto_rawDescGZIP(), []int{0}
}

func (x *ImagePrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ImageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ImageLink string `protobuf:"bytes,2,opt,name=image_link,json=imageLink,proto3" json:"image_link,omitempty"`
}

func (x *ImageData) Reset() {
	*x = ImageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_images_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageData) ProtoMessage() {}

func (x *ImageData) ProtoReflect() protoreflect.Message {
	mi := &file_images_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageData.ProtoReflect.Descriptor instead.
func (*ImageData) Descriptor() ([]byte, []int) {
	return file_images_proto_rawDescGZIP(), []int{1}
}

func (x *ImageData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ImageData) GetImageLink() string {
	if x != nil {
		return x.ImageLink
	}
	return ""
}

type ImageUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageUrl string `protobuf:"bytes,1,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *ImageUrl) Reset() {
	*x = ImageUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_images_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUrl) ProtoMessage() {}

func (x *ImageUrl) ProtoReflect() protoreflect.Message {
	mi := &file_images_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUrl.ProtoReflect.Descriptor instead.
func (*ImageUrl) Descriptor() ([]byte, []int) {
	return file_images_proto_rawDescGZIP(), []int{2}
}

func (x *ImageUrl) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_images_proto protoreflect.FileDescriptor

var file_images_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0a, 0x63,
	0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x09,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x27, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72,
	0x6c, 0x32, 0xe7, 0x01, 0x0a, 0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1e, 0x2e, 0x63,
	0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x0b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x14, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_images_proto_rawDescOnce sync.Once
	file_images_proto_rawDescData = file_images_proto_rawDesc
)

func file_images_proto_rawDescGZIP() []byte {
	file_images_proto_rawDescOnce.Do(func() {
		file_images_proto_rawDescData = protoimpl.X.CompressGZIP(file_images_proto_rawDescData)
	})
	return file_images_proto_rawDescData
}

var file_images_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_images_proto_goTypes = []interface{}{
	(*ImagePrimaryKey)(nil), // 0: coins_service.ImagePrimaryKey
	(*ImageData)(nil),       // 1: coins_service.ImageData
	(*ImageUrl)(nil),        // 2: coins_service.ImageUrl
	(*Empty)(nil),           // 3: coins_service.Empty
}
var file_images_proto_depIdxs = []int32{
	1, // 0: coins_service.ImagesService.ImageUpload:input_type -> coins_service.ImageData
	0, // 1: coins_service.ImagesService.ImageDelete:input_type -> coins_service.ImagePrimaryKey
	0, // 2: coins_service.ImagesService.GetFile:input_type -> coins_service.ImagePrimaryKey
	0, // 3: coins_service.ImagesService.ImageUpload:output_type -> coins_service.ImagePrimaryKey
	3, // 4: coins_service.ImagesService.ImageDelete:output_type -> coins_service.Empty
	2, // 5: coins_service.ImagesService.GetFile:output_type -> coins_service.ImageUrl
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_images_proto_init() }
func file_images_proto_init() {
	if File_images_proto != nil {
		return
	}
	file_coin_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_images_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImagePrimaryKey); i {
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
		file_images_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageData); i {
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
		file_images_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageUrl); i {
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
			RawDescriptor: file_images_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_images_proto_goTypes,
		DependencyIndexes: file_images_proto_depIdxs,
		MessageInfos:      file_images_proto_msgTypes,
	}.Build()
	File_images_proto = out.File
	file_images_proto_rawDesc = nil
	file_images_proto_goTypes = nil
	file_images_proto_depIdxs = nil
}
