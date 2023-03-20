// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: pagination/v1/pagination.proto

package v1

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

// PageIdentifier represents a Page Identifier.
type PageIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is  the name of the PageIdentifier, this could be a ID (database) or a resource name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// filter represents filtering option for the pagination
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *PageIdentifier) Reset() {
	*x = PageIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pagination_v1_pagination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageIdentifier) ProtoMessage() {}

func (x *PageIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_pagination_v1_pagination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageIdentifier.ProtoReflect.Descriptor instead.
func (*PageIdentifier) Descriptor() ([]byte, []int) {
	return file_pagination_v1_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *PageIdentifier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PageIdentifier) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

var File_pagination_v1_pagination_proto protoreflect.FileDescriptor

var file_pagination_v1_pagination_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22,
	0x3c, 0x0a, 0x0e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x2e, 0x5a,
	0x2c, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x62,
	0x69, 0x74, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x63, 0x2f, 0x6b, 0x69, 0x74, 0x2f,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pagination_v1_pagination_proto_rawDescOnce sync.Once
	file_pagination_v1_pagination_proto_rawDescData = file_pagination_v1_pagination_proto_rawDesc
)

func file_pagination_v1_pagination_proto_rawDescGZIP() []byte {
	file_pagination_v1_pagination_proto_rawDescOnce.Do(func() {
		file_pagination_v1_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_pagination_v1_pagination_proto_rawDescData)
	})
	return file_pagination_v1_pagination_proto_rawDescData
}

var file_pagination_v1_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pagination_v1_pagination_proto_goTypes = []interface{}{
	(*PageIdentifier)(nil), // 0: pagination.v1.PageIdentifier
}
var file_pagination_v1_pagination_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pagination_v1_pagination_proto_init() }
func file_pagination_v1_pagination_proto_init() {
	if File_pagination_v1_pagination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pagination_v1_pagination_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageIdentifier); i {
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
			RawDescriptor: file_pagination_v1_pagination_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pagination_v1_pagination_proto_goTypes,
		DependencyIndexes: file_pagination_v1_pagination_proto_depIdxs,
		MessageInfos:      file_pagination_v1_pagination_proto_msgTypes,
	}.Build()
	File_pagination_v1_pagination_proto = out.File
	file_pagination_v1_pagination_proto_rawDesc = nil
	file_pagination_v1_pagination_proto_goTypes = nil
	file_pagination_v1_pagination_proto_depIdxs = nil
}
