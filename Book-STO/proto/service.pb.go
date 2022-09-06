// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/service.proto

package proto

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

type FindBookByIdAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdAuthor int64 `protobuf:"varint,1,opt,name=idAuthor,proto3" json:"idAuthor,omitempty"`
}

func (x *FindBookByIdAuthorRequest) Reset() {
	*x = FindBookByIdAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindBookByIdAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindBookByIdAuthorRequest) ProtoMessage() {}

func (x *FindBookByIdAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindBookByIdAuthorRequest.ProtoReflect.Descriptor instead.
func (*FindBookByIdAuthorRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *FindBookByIdAuthorRequest) GetIdAuthor() int64 {
	if x != nil {
		return x.IdAuthor
	}
	return 0
}

type FindBookByIdAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdBook       int64  `protobuf:"varint,1,opt,name=idBook,proto3" json:"idBook,omitempty"`
	NameBook     string `protobuf:"bytes,2,opt,name=nameBook,proto3" json:"nameBook,omitempty"`
	CategoryBook string `protobuf:"bytes,3,opt,name=categoryBook,proto3" json:"categoryBook,omitempty"`
}

func (x *FindBookByIdAuthorResponse) Reset() {
	*x = FindBookByIdAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindBookByIdAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindBookByIdAuthorResponse) ProtoMessage() {}

func (x *FindBookByIdAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindBookByIdAuthorResponse.ProtoReflect.Descriptor instead.
func (*FindBookByIdAuthorResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *FindBookByIdAuthorResponse) GetIdBook() int64 {
	if x != nil {
		return x.IdBook
	}
	return 0
}

func (x *FindBookByIdAuthorResponse) GetNameBook() string {
	if x != nil {
		return x.NameBook
	}
	return ""
}

func (x *FindBookByIdAuthorResponse) GetCategoryBook() string {
	if x != nil {
		return x.CategoryBook
	}
	return ""
}

type BooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BooksResponse []*FindBookByIdAuthorResponse `protobuf:"bytes,1,rep,name=books_response,json=booksResponse,proto3" json:"books_response,omitempty"`
}

func (x *BooksResponse) Reset() {
	*x = BooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksResponse) ProtoMessage() {}

func (x *BooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooksResponse.ProtoReflect.Descriptor instead.
func (*BooksResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *BooksResponse) GetBooksResponse() []*FindBookByIdAuthorResponse {
	if x != nil {
		return x.BooksResponse
	}
	return nil
}

var File_proto_service_proto protoreflect.FileDescriptor

var file_proto_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x19,
	0x46, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x64, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x74, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x6f,
	0x6b, 0x42, 0x79, 0x49, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x59, 0x0a, 0x0d, 0x42,
	0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0e,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x60, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x12, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x6f, 0x6f,
	0x6b, 0x42, 0x79, 0x49, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_proto_rawDescOnce sync.Once
	file_proto_service_proto_rawDescData = file_proto_service_proto_rawDesc
)

func file_proto_service_proto_rawDescGZIP() []byte {
	file_proto_service_proto_rawDescOnce.Do(func() {
		file_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_proto_rawDescData)
	})
	return file_proto_service_proto_rawDescData
}

var file_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_service_proto_goTypes = []interface{}{
	(*FindBookByIdAuthorRequest)(nil),  // 0: proto.FindBookByIdAuthorRequest
	(*FindBookByIdAuthorResponse)(nil), // 1: proto.FindBookByIdAuthorResponse
	(*BooksResponse)(nil),              // 2: proto.BooksResponse
}
var file_proto_service_proto_depIdxs = []int32{
	1, // 0: proto.BooksResponse.books_response:type_name -> proto.FindBookByIdAuthorResponse
	0, // 1: proto.AddAuthorService.FindBookByIdAuthor:input_type -> proto.FindBookByIdAuthorRequest
	2, // 2: proto.AddAuthorService.FindBookByIdAuthor:output_type -> proto.BooksResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_service_proto_init() }
func file_proto_service_proto_init() {
	if File_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindBookByIdAuthorRequest); i {
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
		file_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindBookByIdAuthorResponse); i {
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
		file_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooksResponse); i {
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
			RawDescriptor: file_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_proto_goTypes,
		DependencyIndexes: file_proto_service_proto_depIdxs,
		MessageInfos:      file_proto_service_proto_msgTypes,
	}.Build()
	File_proto_service_proto = out.File
	file_proto_service_proto_rawDesc = nil
	file_proto_service_proto_goTypes = nil
	file_proto_service_proto_depIdxs = nil
}
