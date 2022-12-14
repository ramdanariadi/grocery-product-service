// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: proto/category.proto

package category

import (
	"github.com/ramdanariadi/grocery-product-service/main/service/response"
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

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ImageUrl string `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_proto_category_proto_rawDescGZIP(), []int{0}
}

func (x *Category) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type CategoryId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CategoryId) Reset() {
	*x = CategoryId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryId) ProtoMessage() {}

func (x *CategoryId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryId.ProtoReflect.Descriptor instead.
func (*CategoryId) Descriptor() ([]byte, []int) {
	return file_proto_category_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EmptyCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyCategory) Reset() {
	*x = EmptyCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyCategory) ProtoMessage() {}

func (x *EmptyCategory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyCategory.ProtoReflect.Descriptor instead.
func (*EmptyCategory) Descriptor() ([]byte, []int) {
	return file_proto_category_proto_rawDescGZIP(), []int{2}
}

type CategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Category `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CategoryResponse) Reset() {
	*x = CategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryResponse) ProtoMessage() {}

func (x *CategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryResponse.ProtoReflect.Descriptor instead.
func (*CategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_category_proto_rawDescGZIP(), []int{3}
}

func (x *CategoryResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CategoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CategoryResponse) GetData() *Category {
	if x != nil {
		return x.Data
	}
	return nil
}

type MultipleCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string      `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Category `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *MultipleCategoryResponse) Reset() {
	*x = MultipleCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleCategoryResponse) ProtoMessage() {}

func (x *MultipleCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleCategoryResponse.ProtoReflect.Descriptor instead.
func (*MultipleCategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_category_proto_rawDescGZIP(), []int{4}
}

func (x *MultipleCategoryResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MultipleCategoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MultipleCategoryResponse) GetData() []*Category {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_category_proto protoreflect.FileDescriptor

var file_proto_category_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x1c, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x69, 0x0a, 0x10, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x71, 0x0a, 0x18, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x8f, 0x02, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x46, 0x69, 0x6e,
	0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x40, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x0a, 0x16, 0x69, 0x64, 0x2e, 0x67, 0x72,
	0x6f, 0x63, 0x65, 0x72, 0x79, 0x2e, 0x74, 0x75, 0x6e, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x11, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_category_proto_rawDescOnce sync.Once
	file_proto_category_proto_rawDescData = file_proto_category_proto_rawDesc
)

func file_proto_category_proto_rawDescGZIP() []byte {
	file_proto_category_proto_rawDescOnce.Do(func() {
		file_proto_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_category_proto_rawDescData)
	})
	return file_proto_category_proto_rawDescData
}

var file_proto_category_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_category_proto_goTypes = []interface{}{
	(*Category)(nil),                 // 0: proto.Category
	(*CategoryId)(nil),               // 1: proto.CategoryId
	(*EmptyCategory)(nil),            // 2: proto.EmptyCategory
	(*CategoryResponse)(nil),         // 3: proto.CategoryResponse
	(*MultipleCategoryResponse)(nil), // 4: proto.MultipleCategoryResponse
	(*response.Response)(nil),        // 5: proto.Response
}
var file_proto_category_proto_depIdxs = []int32{
	0, // 0: proto.CategoryResponse.data:type_name -> proto.Category
	0, // 1: proto.MultipleCategoryResponse.data:type_name -> proto.Category
	1, // 2: proto.CategoryService.FindById:input_type -> proto.CategoryId
	2, // 3: proto.CategoryService.FindAll:input_type -> proto.EmptyCategory
	0, // 4: proto.CategoryService.Save:input_type -> proto.Category
	0, // 5: proto.CategoryService.Update:input_type -> proto.Category
	1, // 6: proto.CategoryService.Delete:input_type -> proto.CategoryId
	3, // 7: proto.CategoryService.FindById:output_type -> proto.CategoryResponse
	4, // 8: proto.CategoryService.FindAll:output_type -> proto.MultipleCategoryResponse
	5, // 9: proto.CategoryService.Save:output_type -> proto.Response
	5, // 10: proto.CategoryService.Update:output_type -> proto.Response
	5, // 11: proto.CategoryService.Delete:output_type -> proto.Response
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_category_proto_init() }
func file_proto_category_proto_init() {
	if File_proto_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_proto_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryId); i {
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
		file_proto_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyCategory); i {
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
		file_proto_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryResponse); i {
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
		file_proto_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleCategoryResponse); i {
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
			RawDescriptor: file_proto_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_category_proto_goTypes,
		DependencyIndexes: file_proto_category_proto_depIdxs,
		MessageInfos:      file_proto_category_proto_msgTypes,
	}.Build()
	File_proto_category_proto = out.File
	file_proto_category_proto_rawDesc = nil
	file_proto_category_proto_goTypes = nil
	file_proto_category_proto_depIdxs = nil
}
