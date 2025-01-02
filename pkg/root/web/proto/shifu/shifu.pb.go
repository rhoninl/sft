// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.28.0--rc1
// source: pkg/root/web/proto/shifu/shifu.proto

package shifu

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

type CheckInstallationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckInstallationRequest) Reset() {
	*x = CheckInstallationRequest{}
	mi := &file_pkg_root_web_proto_shifu_shifu_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckInstallationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInstallationRequest) ProtoMessage() {}

func (x *CheckInstallationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_root_web_proto_shifu_shifu_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInstallationRequest.ProtoReflect.Descriptor instead.
func (*CheckInstallationRequest) Descriptor() ([]byte, []int) {
	return file_pkg_root_web_proto_shifu_shifu_proto_rawDescGZIP(), []int{0}
}

type CheckInstallationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Installed     bool                   `protobuf:"varint,1,opt,name=installed,proto3" json:"installed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckInstallationResponse) Reset() {
	*x = CheckInstallationResponse{}
	mi := &file_pkg_root_web_proto_shifu_shifu_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckInstallationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInstallationResponse) ProtoMessage() {}

func (x *CheckInstallationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_root_web_proto_shifu_shifu_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInstallationResponse.ProtoReflect.Descriptor instead.
func (*CheckInstallationResponse) Descriptor() ([]byte, []int) {
	return file_pkg_root_web_proto_shifu_shifu_proto_rawDescGZIP(), []int{1}
}

func (x *CheckInstallationResponse) GetInstalled() bool {
	if x != nil {
		return x.Installed
	}
	return false
}

type InstallShifuRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InstallShifuRequest) Reset() {
	*x = InstallShifuRequest{}
	mi := &file_pkg_root_web_proto_shifu_shifu_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InstallShifuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallShifuRequest) ProtoMessage() {}

func (x *InstallShifuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_root_web_proto_shifu_shifu_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallShifuRequest.ProtoReflect.Descriptor instead.
func (*InstallShifuRequest) Descriptor() ([]byte, []int) {
	return file_pkg_root_web_proto_shifu_shifu_proto_rawDescGZIP(), []int{2}
}

func (x *InstallShifuRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type InstallShifuResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InstallShifuResponse) Reset() {
	*x = InstallShifuResponse{}
	mi := &file_pkg_root_web_proto_shifu_shifu_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InstallShifuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallShifuResponse) ProtoMessage() {}

func (x *InstallShifuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_root_web_proto_shifu_shifu_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallShifuResponse.ProtoReflect.Descriptor instead.
func (*InstallShifuResponse) Descriptor() ([]byte, []int) {
	return file_pkg_root_web_proto_shifu_shifu_proto_rawDescGZIP(), []int{3}
}

func (x *InstallShifuResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *InstallShifuResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_root_web_proto_shifu_shifu_proto protoreflect.FileDescriptor

var file_pkg_root_web_proto_shifu_shifu_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x6f, 0x74, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x69, 0x66, 0x75, 0x2f, 0x73, 0x68, 0x69, 0x66, 0x75,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x68, 0x69, 0x66, 0x75, 0x22, 0x1a, 0x0a,
	0x18, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x19, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x65, 0x64, 0x22, 0x2f, 0x0a, 0x13, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53,
	0x68, 0x69, 0x66, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x46, 0x0a, 0x14, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x53, 0x68, 0x69, 0x66, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xb3, 0x01,
	0x0a, 0x0c, 0x53, 0x68, 0x69, 0x66, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58,
	0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x73, 0x68, 0x69, 0x66, 0x75, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x68, 0x69, 0x66, 0x75, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x53, 0x68, 0x69, 0x66, 0x75, 0x12, 0x1a, 0x2e, 0x73, 0x68, 0x69, 0x66, 0x75,
	0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x68, 0x69, 0x66, 0x75, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x68, 0x69, 0x66, 0x75, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x53, 0x68, 0x69, 0x66, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x68, 0x6f, 0x6e, 0x69, 0x6e, 0x6c, 0x2f, 0x73, 0x66, 0x74, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x72, 0x6f, 0x6f, 0x74, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x68, 0x69, 0x66, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_root_web_proto_shifu_shifu_proto_rawDescOnce sync.Once
	file_pkg_root_web_proto_shifu_shifu_proto_rawDescData = file_pkg_root_web_proto_shifu_shifu_proto_rawDesc
)

func file_pkg_root_web_proto_shifu_shifu_proto_rawDescGZIP() []byte {
	file_pkg_root_web_proto_shifu_shifu_proto_rawDescOnce.Do(func() {
		file_pkg_root_web_proto_shifu_shifu_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_root_web_proto_shifu_shifu_proto_rawDescData)
	})
	return file_pkg_root_web_proto_shifu_shifu_proto_rawDescData
}

var file_pkg_root_web_proto_shifu_shifu_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_root_web_proto_shifu_shifu_proto_goTypes = []any{
	(*CheckInstallationRequest)(nil),  // 0: shifu.CheckInstallationRequest
	(*CheckInstallationResponse)(nil), // 1: shifu.CheckInstallationResponse
	(*InstallShifuRequest)(nil),       // 2: shifu.InstallShifuRequest
	(*InstallShifuResponse)(nil),      // 3: shifu.InstallShifuResponse
}
var file_pkg_root_web_proto_shifu_shifu_proto_depIdxs = []int32{
	0, // 0: shifu.ShifuService.CheckInstallation:input_type -> shifu.CheckInstallationRequest
	2, // 1: shifu.ShifuService.InstallShifu:input_type -> shifu.InstallShifuRequest
	1, // 2: shifu.ShifuService.CheckInstallation:output_type -> shifu.CheckInstallationResponse
	3, // 3: shifu.ShifuService.InstallShifu:output_type -> shifu.InstallShifuResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_root_web_proto_shifu_shifu_proto_init() }
func file_pkg_root_web_proto_shifu_shifu_proto_init() {
	if File_pkg_root_web_proto_shifu_shifu_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_root_web_proto_shifu_shifu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_root_web_proto_shifu_shifu_proto_goTypes,
		DependencyIndexes: file_pkg_root_web_proto_shifu_shifu_proto_depIdxs,
		MessageInfos:      file_pkg_root_web_proto_shifu_shifu_proto_msgTypes,
	}.Build()
	File_pkg_root_web_proto_shifu_shifu_proto = out.File
	file_pkg_root_web_proto_shifu_shifu_proto_rawDesc = nil
	file_pkg_root_web_proto_shifu_shifu_proto_goTypes = nil
	file_pkg_root_web_proto_shifu_shifu_proto_depIdxs = nil
}