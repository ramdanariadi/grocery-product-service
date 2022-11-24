// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: proto/ cart.proto

package cart

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

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Total     uint32 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	ProductId string `protobuf:"bytes,3,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	UserId    string `protobuf:"bytes,4,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto__cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_proto__cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_proto__cart_proto_rawDescGZIP(), []int{0}
}

func (x *Cart) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cart) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Cart) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Cart) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CartId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *CartId) Reset() {
	*x = CartId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto__cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartId) ProtoMessage() {}

func (x *CartId) ProtoReflect() protoreflect.Message {
	mi := &file_proto__cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartId.ProtoReflect.Descriptor instead.
func (*CartId) Descriptor() ([]byte, []int) {
	return file_proto__cart_proto_rawDescGZIP(), []int{1}
}

func (x *CartId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CartEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CartEmpty) Reset() {
	*x = CartEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto__cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartEmpty) ProtoMessage() {}

func (x *CartEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_proto__cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartEmpty.ProtoReflect.Descriptor instead.
func (*CartEmpty) Descriptor() ([]byte, []int) {
	return file_proto__cart_proto_rawDescGZIP(), []int{2}
}

type MultipleCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Cart `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *MultipleCartResponse) Reset() {
	*x = MultipleCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto__cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleCartResponse) ProtoMessage() {}

func (x *MultipleCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto__cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleCartResponse.ProtoReflect.Descriptor instead.
func (*MultipleCartResponse) Descriptor() ([]byte, []int) {
	return file_proto__cart_proto_rawDescGZIP(), []int{3}
}

func (x *MultipleCartResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MultipleCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MultipleCartResponse) GetData() []*Cart {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto__cart_proto protoreflect.FileDescriptor

var file_proto__cart_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x20, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x62, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1c,
	0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x06, 0x43, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x0b,
	0x0a, 0x09, 0x43, 0x61, 0x72, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x69, 0x0a, 0x14, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x97, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x29, 0x0a, 0x16, 0x69, 0x64, 0x2e, 0x67, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x2e, 0x74,
	0x75, 0x6e, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0d, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto__cart_proto_rawDescOnce sync.Once
	file_proto__cart_proto_rawDescData = file_proto__cart_proto_rawDesc
)

func file_proto__cart_proto_rawDescGZIP() []byte {
	file_proto__cart_proto_rawDescOnce.Do(func() {
		file_proto__cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto__cart_proto_rawDescData)
	})
	return file_proto__cart_proto_rawDescData
}

var file_proto__cart_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto__cart_proto_goTypes = []interface{}{
	(*Cart)(nil),                 // 0: proto.Cart
	(*CartId)(nil),               // 1: proto.CartId
	(*CartEmpty)(nil),            // 2: proto.CartEmpty
	(*MultipleCartResponse)(nil), // 3: proto.MultipleCartResponse
	(*response.Response)(nil),    // 4: proto.Response
}
var file_proto__cart_proto_depIdxs = []int32{
	0, // 0: proto.MultipleCartResponse.data:type_name -> proto.Cart
	0, // 1: proto.CartService.Save:input_type -> proto.Cart
	1, // 2: proto.CartService.Delete:input_type -> proto.CartId
	2, // 3: proto.CartService.FindAll:input_type -> proto.CartEmpty
	4, // 4: proto.CartService.Save:output_type -> proto.Response
	4, // 5: proto.CartService.Delete:output_type -> proto.Response
	3, // 6: proto.CartService.FindAll:output_type -> proto.MultipleCartResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto__cart_proto_init() }
func file_proto__cart_proto_init() {
	if File_proto__cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto__cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
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
		file_proto__cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartId); i {
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
		file_proto__cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartEmpty); i {
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
		file_proto__cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleCartResponse); i {
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
			RawDescriptor: file_proto__cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto__cart_proto_goTypes,
		DependencyIndexes: file_proto__cart_proto_depIdxs,
		MessageInfos:      file_proto__cart_proto_msgTypes,
	}.Build()
	File_proto__cart_proto = out.File
	file_proto__cart_proto_rawDesc = nil
	file_proto__cart_proto_goTypes = nil
	file_proto__cart_proto_depIdxs = nil
}
