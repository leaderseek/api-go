// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: service/param/team_add_members.proto

package param

import (
	message "github.com/leaderseek/api-go/message"
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

type TeamAddMembersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamID  string            `protobuf:"bytes,1,opt,name=team_i_d,json=teamID,proto3" json:"team_i_d,omitempty"`
	Players []*message.Player `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *TeamAddMembersRequest) Reset() {
	*x = TeamAddMembersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_param_team_add_members_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamAddMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamAddMembersRequest) ProtoMessage() {}

func (x *TeamAddMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_param_team_add_members_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamAddMembersRequest.ProtoReflect.Descriptor instead.
func (*TeamAddMembersRequest) Descriptor() ([]byte, []int) {
	return file_service_param_team_add_members_proto_rawDescGZIP(), []int{0}
}

func (x *TeamAddMembersRequest) GetTeamID() string {
	if x != nil {
		return x.TeamID
	}
	return ""
}

func (x *TeamAddMembersRequest) GetPlayers() []*message.Player {
	if x != nil {
		return x.Players
	}
	return nil
}

var File_service_param_team_add_members_proto protoreflect.FileDescriptor

var file_service_param_team_add_members_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x2f,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x14, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x63, 0x0a, 0x15, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x08, 0x74,
	0x65, 0x61, 0x6d, 0x5f, 0x69, 0x5f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x65, 0x61, 0x6d, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x68, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x65, 0x65, 0x6b,
	0x2e, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x42, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x65, 0x65, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_param_team_add_members_proto_rawDescOnce sync.Once
	file_service_param_team_add_members_proto_rawDescData = file_service_param_team_add_members_proto_rawDesc
)

func file_service_param_team_add_members_proto_rawDescGZIP() []byte {
	file_service_param_team_add_members_proto_rawDescOnce.Do(func() {
		file_service_param_team_add_members_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_param_team_add_members_proto_rawDescData)
	})
	return file_service_param_team_add_members_proto_rawDescData
}

var file_service_param_team_add_members_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_service_param_team_add_members_proto_goTypes = []interface{}{
	(*TeamAddMembersRequest)(nil), // 0: api.v0.service.param.TeamAddMembersRequest
	(*message.Player)(nil),        // 1: api.v0.message.Player
}
var file_service_param_team_add_members_proto_depIdxs = []int32{
	1, // 0: api.v0.service.param.TeamAddMembersRequest.players:type_name -> api.v0.message.Player
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_param_team_add_members_proto_init() }
func file_service_param_team_add_members_proto_init() {
	if File_service_param_team_add_members_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_param_team_add_members_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamAddMembersRequest); i {
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
			RawDescriptor: file_service_param_team_add_members_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_param_team_add_members_proto_goTypes,
		DependencyIndexes: file_service_param_team_add_members_proto_depIdxs,
		MessageInfos:      file_service_param_team_add_members_proto_msgTypes,
	}.Build()
	File_service_param_team_add_members_proto = out.File
	file_service_param_team_add_members_proto_rawDesc = nil
	file_service_param_team_add_members_proto_goTypes = nil
	file_service_param_team_add_members_proto_depIdxs = nil
}
