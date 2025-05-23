// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proxyqueue/v1/proxyqueue.proto

package proxyqueuev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type SubmitURLRequest struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Url                string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	UserAgent          *string                `protobuf:"bytes,2,opt,name=user_agent,json=userAgent,proto3,oneof" json:"user_agent,omitempty"`
	UseRandomUserAgent bool                   `protobuf:"varint,3,opt,name=use_random_user_agent,json=useRandomUserAgent,proto3" json:"use_random_user_agent,omitempty"`
	BinaryContent      bool                   `protobuf:"varint,4,opt,name=binary_content,json=binaryContent,proto3" json:"binary_content,omitempty"`
	Headers            map[string]string      `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *SubmitURLRequest) Reset() {
	*x = SubmitURLRequest{}
	mi := &file_proxyqueue_v1_proxyqueue_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitURLRequest) ProtoMessage() {}

func (x *SubmitURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxyqueue_v1_proxyqueue_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitURLRequest.ProtoReflect.Descriptor instead.
func (*SubmitURLRequest) Descriptor() ([]byte, []int) {
	return file_proxyqueue_v1_proxyqueue_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitURLRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SubmitURLRequest) GetUserAgent() string {
	if x != nil && x.UserAgent != nil {
		return *x.UserAgent
	}
	return ""
}

func (x *SubmitURLRequest) GetUseRandomUserAgent() bool {
	if x != nil {
		return x.UseRandomUserAgent
	}
	return false
}

func (x *SubmitURLRequest) GetBinaryContent() bool {
	if x != nil {
		return x.BinaryContent
	}
	return false
}

func (x *SubmitURLRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

type SubmitURLResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	HtmlContent   string                 `protobuf:"bytes,5,opt,name=html_content,json=htmlContent,proto3" json:"html_content,omitempty"`
	BinaryContent []byte                 `protobuf:"bytes,6,opt,name=binary_content,json=binaryContent,proto3" json:"binary_content,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	StartedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	FinishedAt    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitURLResponse) Reset() {
	*x = SubmitURLResponse{}
	mi := &file_proxyqueue_v1_proxyqueue_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitURLResponse) ProtoMessage() {}

func (x *SubmitURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proxyqueue_v1_proxyqueue_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitURLResponse.ProtoReflect.Descriptor instead.
func (*SubmitURLResponse) Descriptor() ([]byte, []int) {
	return file_proxyqueue_v1_proxyqueue_proto_rawDescGZIP(), []int{1}
}

func (x *SubmitURLResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubmitURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SubmitURLResponse) GetHtmlContent() string {
	if x != nil {
		return x.HtmlContent
	}
	return ""
}

func (x *SubmitURLResponse) GetBinaryContent() []byte {
	if x != nil {
		return x.BinaryContent
	}
	return nil
}

func (x *SubmitURLResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *SubmitURLResponse) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *SubmitURLResponse) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

var File_proxyqueue_v1_proxyqueue_proto protoreflect.FileDescriptor

const file_proxyqueue_v1_proxyqueue_proto_rawDesc = "" +
	"\n" +
	"\x1eproxyqueue/v1/proxyqueue.proto\x12\rproxyqueue.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\xb5\x02\n" +
	"\x10SubmitURLRequest\x12\x10\n" +
	"\x03url\x18\x01 \x01(\tR\x03url\x12\"\n" +
	"\n" +
	"user_agent\x18\x02 \x01(\tH\x00R\tuserAgent\x88\x01\x01\x121\n" +
	"\x15use_random_user_agent\x18\x03 \x01(\bR\x12useRandomUserAgent\x12%\n" +
	"\x0ebinary_content\x18\x04 \x01(\bR\rbinaryContent\x12F\n" +
	"\aheaders\x18\x05 \x03(\v2,.proxyqueue.v1.SubmitURLRequest.HeadersEntryR\aheaders\x1a:\n" +
	"\fHeadersEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\r\n" +
	"\v_user_agent\"\xb2\x02\n" +
	"\x11SubmitURLResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x10\n" +
	"\x03url\x18\x02 \x01(\tR\x03url\x12!\n" +
	"\fhtml_content\x18\x05 \x01(\tR\vhtmlContent\x12%\n" +
	"\x0ebinary_content\x18\x06 \x01(\fR\rbinaryContent\x129\n" +
	"\n" +
	"created_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"started_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tstartedAt\x12;\n" +
	"\vfinished_at\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"finishedAt2e\n" +
	"\x11ProxyQueueService\x12P\n" +
	"\tSubmitURL\x12\x1f.proxyqueue.v1.SubmitURLRequest\x1a .proxyqueue.v1.SubmitURLResponse\"\x00B7Z5go.sazak.io/proxyqueue/gen/proxyqueue/v1;proxyqueuev1b\x06proto3"

var (
	file_proxyqueue_v1_proxyqueue_proto_rawDescOnce sync.Once
	file_proxyqueue_v1_proxyqueue_proto_rawDescData []byte
)

func file_proxyqueue_v1_proxyqueue_proto_rawDescGZIP() []byte {
	file_proxyqueue_v1_proxyqueue_proto_rawDescOnce.Do(func() {
		file_proxyqueue_v1_proxyqueue_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proxyqueue_v1_proxyqueue_proto_rawDesc), len(file_proxyqueue_v1_proxyqueue_proto_rawDesc)))
	})
	return file_proxyqueue_v1_proxyqueue_proto_rawDescData
}

var file_proxyqueue_v1_proxyqueue_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proxyqueue_v1_proxyqueue_proto_goTypes = []any{
	(*SubmitURLRequest)(nil),      // 0: proxyqueue.v1.SubmitURLRequest
	(*SubmitURLResponse)(nil),     // 1: proxyqueue.v1.SubmitURLResponse
	nil,                           // 2: proxyqueue.v1.SubmitURLRequest.HeadersEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_proxyqueue_v1_proxyqueue_proto_depIdxs = []int32{
	2, // 0: proxyqueue.v1.SubmitURLRequest.headers:type_name -> proxyqueue.v1.SubmitURLRequest.HeadersEntry
	3, // 1: proxyqueue.v1.SubmitURLResponse.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: proxyqueue.v1.SubmitURLResponse.started_at:type_name -> google.protobuf.Timestamp
	3, // 3: proxyqueue.v1.SubmitURLResponse.finished_at:type_name -> google.protobuf.Timestamp
	0, // 4: proxyqueue.v1.ProxyQueueService.SubmitURL:input_type -> proxyqueue.v1.SubmitURLRequest
	1, // 5: proxyqueue.v1.ProxyQueueService.SubmitURL:output_type -> proxyqueue.v1.SubmitURLResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proxyqueue_v1_proxyqueue_proto_init() }
func file_proxyqueue_v1_proxyqueue_proto_init() {
	if File_proxyqueue_v1_proxyqueue_proto != nil {
		return
	}
	file_proxyqueue_v1_proxyqueue_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proxyqueue_v1_proxyqueue_proto_rawDesc), len(file_proxyqueue_v1_proxyqueue_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proxyqueue_v1_proxyqueue_proto_goTypes,
		DependencyIndexes: file_proxyqueue_v1_proxyqueue_proto_depIdxs,
		MessageInfos:      file_proxyqueue_v1_proxyqueue_proto_msgTypes,
	}.Build()
	File_proxyqueue_v1_proxyqueue_proto = out.File
	file_proxyqueue_v1_proxyqueue_proto_goTypes = nil
	file_proxyqueue_v1_proxyqueue_proto_depIdxs = nil
}
