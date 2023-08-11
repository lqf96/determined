// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: determined/workspace/v1/workspace.proto

package workspacev1

import (
	userv1 "github.com/determined-ai/determined/proto/pkg/userv1"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

// WorkspaceState is used to track progress during a deletion.
type WorkspaceState int32

const (
	// Object deletion is not in progress.
	WorkspaceState_WORKSPACE_STATE_UNSPECIFIED WorkspaceState = 0
	// The object is being deleted.
	WorkspaceState_WORKSPACE_STATE_DELETING WorkspaceState = 1
	// The object failed to delete.
	WorkspaceState_WORKSPACE_STATE_DELETE_FAILED WorkspaceState = 2
	// The object finished deleting.
	WorkspaceState_WORKSPACE_STATE_DELETED WorkspaceState = 3
)

// Enum value maps for WorkspaceState.
var (
	WorkspaceState_name = map[int32]string{
		0: "WORKSPACE_STATE_UNSPECIFIED",
		1: "WORKSPACE_STATE_DELETING",
		2: "WORKSPACE_STATE_DELETE_FAILED",
		3: "WORKSPACE_STATE_DELETED",
	}
	WorkspaceState_value = map[string]int32{
		"WORKSPACE_STATE_UNSPECIFIED":   0,
		"WORKSPACE_STATE_DELETING":      1,
		"WORKSPACE_STATE_DELETE_FAILED": 2,
		"WORKSPACE_STATE_DELETED":       3,
	}
)

func (x WorkspaceState) Enum() *WorkspaceState {
	p := new(WorkspaceState)
	*p = x
	return p
}

func (x WorkspaceState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkspaceState) Descriptor() protoreflect.EnumDescriptor {
	return file_determined_workspace_v1_workspace_proto_enumTypes[0].Descriptor()
}

func (WorkspaceState) Type() protoreflect.EnumType {
	return &file_determined_workspace_v1_workspace_proto_enumTypes[0]
}

func (x WorkspaceState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkspaceState.Descriptor instead.
func (WorkspaceState) EnumDescriptor() ([]byte, []int) {
	return file_determined_workspace_v1_workspace_proto_rawDescGZIP(), []int{0}
}

// Workspace is a named collection of projects.
type Workspace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique id of the workspace.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The unique name of the workspace.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Whether this workspace is archived or not.
	Archived bool `protobuf:"varint,3,opt,name=archived,proto3" json:"archived,omitempty"`
	// User who created this workspace.
	Username string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	// Whether this workspace is immutable (default uncategorized workspace).
	Immutable bool `protobuf:"varint,5,opt,name=immutable,proto3" json:"immutable,omitempty"`
	// Number of projects associated with this workspace.
	NumProjects int32 `protobuf:"varint,6,opt,name=num_projects,json=numProjects,proto3" json:"num_projects,omitempty"`
	// Pin status of this workspace for the current user.
	Pinned bool `protobuf:"varint,7,opt,name=pinned,proto3" json:"pinned,omitempty"`
	// ID of the user who created this project.
	UserId int32 `protobuf:"varint,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Number of experiments associated with this workspace.
	NumExperiments int32 `protobuf:"varint,9,opt,name=num_experiments,json=numExperiments,proto3" json:"num_experiments,omitempty"`
	// State of workspace during deletion.
	State WorkspaceState `protobuf:"varint,10,opt,name=state,proto3,enum=determined.workspace.v1.WorkspaceState" json:"state,omitempty"`
	// Message stored from errors on async-deleting a workspace.
	ErrorMessage string `protobuf:"bytes,11,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	// Optional agent host uid and gid override.
	AgentUserGroup *userv1.AgentUserGroup `protobuf:"bytes,12,opt,name=agent_user_group,json=agentUserGroup,proto3,oneof" json:"agent_user_group,omitempty"`
	// Optional checkpoint storage config.
	// Expects same format as experiment config's checkpoint storage.
	CheckpointStorageConfig *_struct.Struct `protobuf:"bytes,13,opt,name=checkpoint_storage_config,json=checkpointStorageConfig,proto3,oneof" json:"checkpoint_storage_config,omitempty"`
	// Optional date when workspace was pinned.
	PinnedAt *timestamp.Timestamp `protobuf:"bytes,14,opt,name=pinned_at,json=pinnedAt,proto3,oneof" json:"pinned_at,omitempty"`
	// Name of the default compute pool.
	DefaultComputePool string `protobuf:"bytes,15,opt,name=default_compute_pool,json=defaultComputePool,proto3" json:"default_compute_pool,omitempty"`
	// Name of the default aux pool.
	DefaultAuxPool string `protobuf:"bytes,16,opt,name=default_aux_pool,json=defaultAuxPool,proto3" json:"default_aux_pool,omitempty"`
}

func (x *Workspace) Reset() {
	*x = Workspace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_workspace_v1_workspace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workspace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workspace) ProtoMessage() {}

func (x *Workspace) ProtoReflect() protoreflect.Message {
	mi := &file_determined_workspace_v1_workspace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workspace.ProtoReflect.Descriptor instead.
func (*Workspace) Descriptor() ([]byte, []int) {
	return file_determined_workspace_v1_workspace_proto_rawDescGZIP(), []int{0}
}

func (x *Workspace) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Workspace) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Workspace) GetArchived() bool {
	if x != nil {
		return x.Archived
	}
	return false
}

func (x *Workspace) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Workspace) GetImmutable() bool {
	if x != nil {
		return x.Immutable
	}
	return false
}

func (x *Workspace) GetNumProjects() int32 {
	if x != nil {
		return x.NumProjects
	}
	return 0
}

func (x *Workspace) GetPinned() bool {
	if x != nil {
		return x.Pinned
	}
	return false
}

func (x *Workspace) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Workspace) GetNumExperiments() int32 {
	if x != nil {
		return x.NumExperiments
	}
	return 0
}

func (x *Workspace) GetState() WorkspaceState {
	if x != nil {
		return x.State
	}
	return WorkspaceState_WORKSPACE_STATE_UNSPECIFIED
}

func (x *Workspace) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *Workspace) GetAgentUserGroup() *userv1.AgentUserGroup {
	if x != nil {
		return x.AgentUserGroup
	}
	return nil
}

func (x *Workspace) GetCheckpointStorageConfig() *_struct.Struct {
	if x != nil {
		return x.CheckpointStorageConfig
	}
	return nil
}

func (x *Workspace) GetPinnedAt() *timestamp.Timestamp {
	if x != nil {
		return x.PinnedAt
	}
	return nil
}

func (x *Workspace) GetDefaultComputePool() string {
	if x != nil {
		return x.DefaultComputePool
	}
	return ""
}

func (x *Workspace) GetDefaultAuxPool() string {
	if x != nil {
		return x.DefaultAuxPool
	}
	return ""
}

// PatchWorkspace is a partial update to a workspace with all optional fields.
type PatchWorkspace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The new name for the workspace.
	Name *wrappers.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional agent host uid and gid override.
	AgentUserGroup *userv1.AgentUserGroup `protobuf:"bytes,12,opt,name=agent_user_group,json=agentUserGroup,proto3,oneof" json:"agent_user_group,omitempty"`
	// Optional checkpoint storage config.
	// Expects same format as experiment config's checkpoint storage.
	CheckpointStorageConfig *_struct.Struct `protobuf:"bytes,13,opt,name=checkpoint_storage_config,json=checkpointStorageConfig,proto3,oneof" json:"checkpoint_storage_config,omitempty"`
	// Name of the default compute pool.
	//
	// Deprecated: Do not use.
	DefaultComputePool string `protobuf:"bytes,14,opt,name=default_compute_pool,json=defaultComputePool,proto3" json:"default_compute_pool,omitempty"`
	// Name of the default compute pool can be optional.
	DefaultComputeResourcePool *string `protobuf:"bytes,16,opt,name=default_compute_resource_pool,json=defaultComputeResourcePool,proto3,oneof" json:"default_compute_resource_pool,omitempty"`
	// Name of the default aux pool.
	//
	// Deprecated: Do not use.
	DefaultAuxPool string `protobuf:"bytes,15,opt,name=default_aux_pool,json=defaultAuxPool,proto3" json:"default_aux_pool,omitempty"`
	// Name of the default aux pool can be optional.
	DefaultAuxResourcePool *string `protobuf:"bytes,17,opt,name=default_aux_resource_pool,json=defaultAuxResourcePool,proto3,oneof" json:"default_aux_resource_pool,omitempty"`
}

func (x *PatchWorkspace) Reset() {
	*x = PatchWorkspace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_workspace_v1_workspace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchWorkspace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchWorkspace) ProtoMessage() {}

func (x *PatchWorkspace) ProtoReflect() protoreflect.Message {
	mi := &file_determined_workspace_v1_workspace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchWorkspace.ProtoReflect.Descriptor instead.
func (*PatchWorkspace) Descriptor() ([]byte, []int) {
	return file_determined_workspace_v1_workspace_proto_rawDescGZIP(), []int{1}
}

func (x *PatchWorkspace) GetName() *wrappers.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *PatchWorkspace) GetAgentUserGroup() *userv1.AgentUserGroup {
	if x != nil {
		return x.AgentUserGroup
	}
	return nil
}

func (x *PatchWorkspace) GetCheckpointStorageConfig() *_struct.Struct {
	if x != nil {
		return x.CheckpointStorageConfig
	}
	return nil
}

// Deprecated: Do not use.
func (x *PatchWorkspace) GetDefaultComputePool() string {
	if x != nil {
		return x.DefaultComputePool
	}
	return ""
}

func (x *PatchWorkspace) GetDefaultComputeResourcePool() string {
	if x != nil && x.DefaultComputeResourcePool != nil {
		return *x.DefaultComputeResourcePool
	}
	return ""
}

// Deprecated: Do not use.
func (x *PatchWorkspace) GetDefaultAuxPool() string {
	if x != nil {
		return x.DefaultAuxPool
	}
	return ""
}

func (x *PatchWorkspace) GetDefaultAuxResourcePool() string {
	if x != nil && x.DefaultAuxResourcePool != nil {
		return *x.DefaultAuxResourcePool
	}
	return ""
}

var File_determined_workspace_v1_workspace_proto protoreflect.FileDescriptor

var file_determined_workspace_v1_workspace_proto_rawDesc = []byte{
	0x0a, 0x27, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x64, 0x65, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1d, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7,
	0x06, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0x92, 0x41, 0x03, 0x80,
	0x01, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x75, 0x6d,
	0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x64, 0x65, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x51, 0x0a, 0x10, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x65, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x00, 0x52,
	0x0e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x88,
	0x01, 0x01, 0x12, 0x58, 0x0a, 0x19, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x01,
	0x52, 0x17, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x09,
	0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x08, 0x70,
	0x69, 0x6e, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x6f,
	0x6f, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x28, 0x0a, 0x10,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x75, 0x78, 0x5f, 0x70, 0x6f, 0x6f, 0x6c,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41,
	0x75, 0x78, 0x50, 0x6f, 0x6f, 0x6c, 0x3a, 0x7f, 0x92, 0x41, 0x7c, 0x0a, 0x7a, 0xd2, 0x01, 0x08,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0xd2, 0x01, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0xd2, 0x01, 0x02, 0x69, 0x64, 0xd2, 0x01, 0x09,
	0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0xd2, 0x01, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0xd2, 0x01, 0x0f, 0x6e, 0x75, 0x6d, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0xd2, 0x01, 0x0c, 0x6e, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0xd2, 0x01, 0x06, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0xd2, 0x01, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0xd2, 0x01, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x1c, 0x0a, 0x1a,
	0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70,
	0x69, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0xce, 0x04, 0x0a, 0x0e, 0x50, 0x61, 0x74,
	0x63, 0x68, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x51, 0x0a,
	0x10, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x65, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x00, 0x52, 0x0e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x88, 0x01, 0x01,
	0x12, 0x58, 0x0a, 0x19, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x01, 0x52, 0x17,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x14, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x6f,
	0x6f, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x12, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x50, 0x6f, 0x6f, 0x6c,
	0x12, 0x46, 0x0a, 0x1d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f,
	0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x1a, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x10, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x5f, 0x61, 0x75, 0x78, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41,
	0x75, 0x78, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x3e, 0x0a, 0x19, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x61, 0x75, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70,
	0x6f, 0x6f, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x16, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x41, 0x75, 0x78, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50,
	0x6f, 0x6f, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x1c, 0x0a, 0x1a, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x20, 0x0a, 0x1e, 0x5f, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x42, 0x1c, 0x0a, 0x1a, 0x5f,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x75, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x2a, 0x8f, 0x01, 0x0a, 0x0e, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x1b,
	0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a,
	0x18, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x57,
	0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1b,
	0x0a, 0x17, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0x3b, 0x5a, 0x39, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x65, 0x64, 0x2d, 0x61, 0x69, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_determined_workspace_v1_workspace_proto_rawDescOnce sync.Once
	file_determined_workspace_v1_workspace_proto_rawDescData = file_determined_workspace_v1_workspace_proto_rawDesc
)

func file_determined_workspace_v1_workspace_proto_rawDescGZIP() []byte {
	file_determined_workspace_v1_workspace_proto_rawDescOnce.Do(func() {
		file_determined_workspace_v1_workspace_proto_rawDescData = protoimpl.X.CompressGZIP(file_determined_workspace_v1_workspace_proto_rawDescData)
	})
	return file_determined_workspace_v1_workspace_proto_rawDescData
}

var file_determined_workspace_v1_workspace_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_determined_workspace_v1_workspace_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_determined_workspace_v1_workspace_proto_goTypes = []interface{}{
	(WorkspaceState)(0),           // 0: determined.workspace.v1.WorkspaceState
	(*Workspace)(nil),             // 1: determined.workspace.v1.Workspace
	(*PatchWorkspace)(nil),        // 2: determined.workspace.v1.PatchWorkspace
	(*userv1.AgentUserGroup)(nil), // 3: determined.user.v1.AgentUserGroup
	(*_struct.Struct)(nil),        // 4: google.protobuf.Struct
	(*timestamp.Timestamp)(nil),   // 5: google.protobuf.Timestamp
	(*wrappers.StringValue)(nil),  // 6: google.protobuf.StringValue
}
var file_determined_workspace_v1_workspace_proto_depIdxs = []int32{
	0, // 0: determined.workspace.v1.Workspace.state:type_name -> determined.workspace.v1.WorkspaceState
	3, // 1: determined.workspace.v1.Workspace.agent_user_group:type_name -> determined.user.v1.AgentUserGroup
	4, // 2: determined.workspace.v1.Workspace.checkpoint_storage_config:type_name -> google.protobuf.Struct
	5, // 3: determined.workspace.v1.Workspace.pinned_at:type_name -> google.protobuf.Timestamp
	6, // 4: determined.workspace.v1.PatchWorkspace.name:type_name -> google.protobuf.StringValue
	3, // 5: determined.workspace.v1.PatchWorkspace.agent_user_group:type_name -> determined.user.v1.AgentUserGroup
	4, // 6: determined.workspace.v1.PatchWorkspace.checkpoint_storage_config:type_name -> google.protobuf.Struct
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_determined_workspace_v1_workspace_proto_init() }
func file_determined_workspace_v1_workspace_proto_init() {
	if File_determined_workspace_v1_workspace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_determined_workspace_v1_workspace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workspace); i {
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
		file_determined_workspace_v1_workspace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchWorkspace); i {
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
	file_determined_workspace_v1_workspace_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_determined_workspace_v1_workspace_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_determined_workspace_v1_workspace_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_determined_workspace_v1_workspace_proto_goTypes,
		DependencyIndexes: file_determined_workspace_v1_workspace_proto_depIdxs,
		EnumInfos:         file_determined_workspace_v1_workspace_proto_enumTypes,
		MessageInfos:      file_determined_workspace_v1_workspace_proto_msgTypes,
	}.Build()
	File_determined_workspace_v1_workspace_proto = out.File
	file_determined_workspace_v1_workspace_proto_rawDesc = nil
	file_determined_workspace_v1_workspace_proto_goTypes = nil
	file_determined_workspace_v1_workspace_proto_depIdxs = nil
}
