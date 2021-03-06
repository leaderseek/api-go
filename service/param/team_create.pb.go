// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: service/param/team_create.proto

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

type TeamCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team    *message.Team     `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	Players []*message.Player `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *TeamCreateRequest) Reset() {
	*x = TeamCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_param_team_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamCreateRequest) ProtoMessage() {}

func (x *TeamCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_param_team_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamCreateRequest.ProtoReflect.Descriptor instead.
func (*TeamCreateRequest) Descriptor() ([]byte, []int) {
	return file_service_param_team_create_proto_rawDescGZIP(), []int{0}
}

func (x *TeamCreateRequest) GetTeam() *message.Team {
	if x != nil {
		return x.Team
	}
	return nil
}

func (x *TeamCreateRequest) GetPlayers() []*message.Player {
	if x != nil {
		return x.Players
	}
	return nil
}

type TeamCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamID string `protobuf:"bytes,1,opt,name=team_i_d,json=teamID,proto3" json:"team_i_d,omitempty"`
}

func (x *TeamCreateResponse) Reset() {
	*x = TeamCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_param_team_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamCreateResponse) ProtoMessage() {}

func (x *TeamCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_param_team_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamCreateResponse.ProtoReflect.Descriptor instead.
func (*TeamCreateResponse) Descriptor() ([]byte, []int) {
	return file_service_param_team_create_proto_rawDescGZIP(), []int{1}
}

func (x *TeamCreateResponse) GetTeamID() string {
	if x != nil {
		return x.TeamID
	}
	return ""
}

var File_service_param_team_create_proto protoreflect.FileDescriptor

var file_service_param_team_create_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x2f,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6f, 0x0a, 0x11, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d,
	0x12, 0x30, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x6d,
	0x5f, 0x69, 0x5f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x44, 0x42, 0x68, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x65, 0x65, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2d,
	0x64, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x42, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x65, 0x65, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_param_team_create_proto_rawDescOnce sync.Once
	file_service_param_team_create_proto_rawDescData = file_service_param_team_create_proto_rawDesc
)

func file_service_param_team_create_proto_rawDescGZIP() []byte {
	file_service_param_team_create_proto_rawDescOnce.Do(func() {
		file_service_param_team_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_param_team_create_proto_rawDescData)
	})
	return file_service_param_team_create_proto_rawDescData
}

var file_service_param_team_create_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_param_team_create_proto_goTypes = []interface{}{
	(*TeamCreateRequest)(nil),  // 0: api.v0.service.param.TeamCreateRequest
	(*TeamCreateResponse)(nil), // 1: api.v0.service.param.TeamCreateResponse
	(*message.Team)(nil),       // 2: api.v0.message.Team
	(*message.Player)(nil),     // 3: api.v0.message.Player
}
var file_service_param_team_create_proto_depIdxs = []int32{
	2, // 0: api.v0.service.param.TeamCreateRequest.team:type_name -> api.v0.message.Team
	3, // 1: api.v0.service.param.TeamCreateRequest.players:type_name -> api.v0.message.Player
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_param_team_create_proto_init() }
func file_service_param_team_create_proto_init() {
	if File_service_param_team_create_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_param_team_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamCreateRequest); i {
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
		file_service_param_team_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamCreateResponse); i {
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
			RawDescriptor: file_service_param_team_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_param_team_create_proto_goTypes,
		DependencyIndexes: file_service_param_team_create_proto_depIdxs,
		MessageInfos:      file_service_param_team_create_proto_msgTypes,
	}.Build()
	File_service_param_team_create_proto = out.File
	file_service_param_team_create_proto_rawDesc = nil
	file_service_param_team_create_proto_goTypes = nil
	file_service_param_team_create_proto_depIdxs = nil
}
