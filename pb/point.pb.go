// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.2
// source: point.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Value     float32                `protobuf:"fixed32,4,opt,name=value,proto3" json:"value,omitempty"`
	Time      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	Index     float32                `protobuf:"fixed32,13,opt,name=index,proto3" json:"index,omitempty"`
	Text      string                 `protobuf:"bytes,8,opt,name=text,proto3" json:"text,omitempty"`
	Key       string                 `protobuf:"bytes,11,opt,name=key,proto3" json:"key,omitempty"`
	Tombstone int32                  `protobuf:"varint,12,opt,name=tombstone,proto3" json:"tombstone,omitempty"`
	Data      []byte                 `protobuf:"bytes,14,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_point_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Point) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Point) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Point) GetIndex() float32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Point) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Point) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Point) GetTombstone() int32 {
	if x != nil {
		return x.Tombstone
	}
	return 0
}

func (x *Point) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Points struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *Points) Reset() {
	*x = Points{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Points) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Points) ProtoMessage() {}

func (x *Points) ProtoReflect() protoreflect.Message {
	mi := &file_point_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Points.ProtoReflect.Descriptor instead.
func (*Points) Descriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{1}
}

func (x *Points) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

var File_point_proto protoreflect.FileDescriptor

var file_point_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x6d, 0x62, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x6d, 0x62, 0x73, 0x74, 0x6f, 0x6e, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x2b, 0x0a, 0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x21,
	0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_point_proto_rawDescOnce sync.Once
	file_point_proto_rawDescData = file_point_proto_rawDesc
)

func file_point_proto_rawDescGZIP() []byte {
	file_point_proto_rawDescOnce.Do(func() {
		file_point_proto_rawDescData = protoimpl.X.CompressGZIP(file_point_proto_rawDescData)
	})
	return file_point_proto_rawDescData
}

var file_point_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_point_proto_goTypes = []interface{}{
	(*Point)(nil),                 // 0: pb.Point
	(*Points)(nil),                // 1: pb.Points
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_point_proto_depIdxs = []int32{
	2, // 0: pb.Point.time:type_name -> google.protobuf.Timestamp
	0, // 1: pb.Points.points:type_name -> pb.Point
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_point_proto_init() }
func file_point_proto_init() {
	if File_point_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_point_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_point_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Points); i {
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
			RawDescriptor: file_point_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_point_proto_goTypes,
		DependencyIndexes: file_point_proto_depIdxs,
		MessageInfos:      file_point_proto_msgTypes,
	}.Build()
	File_point_proto = out.File
	file_point_proto_rawDesc = nil
	file_point_proto_goTypes = nil
	file_point_proto_depIdxs = nil
}