// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/eliza/v1/eliza.proto

package elizav1

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

type SayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sentence string `protobuf:"bytes,1,opt,name=sentence,proto3" json:"sentence,omitempty"`
}

func (x *SayRequest) Reset() {
	*x = SayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eliza_v1_eliza_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayRequest) ProtoMessage() {}

func (x *SayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eliza_v1_eliza_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayRequest.ProtoReflect.Descriptor instead.
func (*SayRequest) Descriptor() ([]byte, []int) {
	return file_proto_eliza_v1_eliza_proto_rawDescGZIP(), []int{0}
}

func (x *SayRequest) GetSentence() string {
	if x != nil {
		return x.Sentence
	}
	return ""
}

type SayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sentence string `protobuf:"bytes,1,opt,name=sentence,proto3" json:"sentence,omitempty"`
}

func (x *SayResponse) Reset() {
	*x = SayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eliza_v1_eliza_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayResponse) ProtoMessage() {}

func (x *SayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eliza_v1_eliza_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayResponse.ProtoReflect.Descriptor instead.
func (*SayResponse) Descriptor() ([]byte, []int) {
	return file_proto_eliza_v1_eliza_proto_rawDescGZIP(), []int{1}
}

func (x *SayResponse) GetSentence() string {
	if x != nil {
		return x.Sentence
	}
	return ""
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eliza_v1_eliza_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eliza_v1_eliza_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_eliza_v1_eliza_proto_rawDescGZIP(), []int{2}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age int64 `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eliza_v1_eliza_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eliza_v1_eliza_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_proto_eliza_v1_eliza_proto_rawDescGZIP(), []int{3}
}

func (x *HelloResponse) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

var File_proto_eliza_v1_eliza_proto protoreflect.FileDescriptor

var file_proto_eliza_v1_eliza_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6c, 0x69, 0x7a, 0x61, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6c, 0x69, 0x7a, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6c,
	0x69, 0x7a, 0x61, 0x2e, 0x76, 0x31, 0x22, 0x28, 0x0a, 0x0a, 0x53, 0x61, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65,
	0x22, 0x29, 0x0a, 0x0b, 0x53, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x21, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61,
	0x67, 0x65, 0x32, 0x80, 0x01, 0x0a, 0x0c, 0x45, 0x6c, 0x69, 0x7a, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x03, 0x53, 0x61, 0x79, 0x12, 0x14, 0x2e, 0x65, 0x6c, 0x69,
	0x7a, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x65, 0x6c, 0x69, 0x7a, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x05, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x12, 0x16, 0x2e, 0x65, 0x6c, 0x69, 0x7a, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x6c, 0x69,
	0x7a, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x6c, 0x69, 0x7a, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6c, 0x69, 0x7a, 0x61,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_eliza_v1_eliza_proto_rawDescOnce sync.Once
	file_proto_eliza_v1_eliza_proto_rawDescData = file_proto_eliza_v1_eliza_proto_rawDesc
)

func file_proto_eliza_v1_eliza_proto_rawDescGZIP() []byte {
	file_proto_eliza_v1_eliza_proto_rawDescOnce.Do(func() {
		file_proto_eliza_v1_eliza_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_eliza_v1_eliza_proto_rawDescData)
	})
	return file_proto_eliza_v1_eliza_proto_rawDescData
}

var file_proto_eliza_v1_eliza_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_eliza_v1_eliza_proto_goTypes = []interface{}{
	(*SayRequest)(nil),    // 0: eliza.v1.SayRequest
	(*SayResponse)(nil),   // 1: eliza.v1.SayResponse
	(*HelloRequest)(nil),  // 2: eliza.v1.HelloRequest
	(*HelloResponse)(nil), // 3: eliza.v1.HelloResponse
}
var file_proto_eliza_v1_eliza_proto_depIdxs = []int32{
	0, // 0: eliza.v1.ElizaService.Say:input_type -> eliza.v1.SayRequest
	2, // 1: eliza.v1.ElizaService.Hello:input_type -> eliza.v1.HelloRequest
	1, // 2: eliza.v1.ElizaService.Say:output_type -> eliza.v1.SayResponse
	3, // 3: eliza.v1.ElizaService.Hello:output_type -> eliza.v1.HelloResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_eliza_v1_eliza_proto_init() }
func file_proto_eliza_v1_eliza_proto_init() {
	if File_proto_eliza_v1_eliza_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_eliza_v1_eliza_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayRequest); i {
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
		file_proto_eliza_v1_eliza_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayResponse); i {
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
		file_proto_eliza_v1_eliza_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_proto_eliza_v1_eliza_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
			RawDescriptor: file_proto_eliza_v1_eliza_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_eliza_v1_eliza_proto_goTypes,
		DependencyIndexes: file_proto_eliza_v1_eliza_proto_depIdxs,
		MessageInfos:      file_proto_eliza_v1_eliza_proto_msgTypes,
	}.Build()
	File_proto_eliza_v1_eliza_proto = out.File
	file_proto_eliza_v1_eliza_proto_rawDesc = nil
	file_proto_eliza_v1_eliza_proto_goTypes = nil
	file_proto_eliza_v1_eliza_proto_depIdxs = nil
}