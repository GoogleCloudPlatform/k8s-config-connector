// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +generated:mapper
// krm.group: memorystore.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.memorystore.v1

package memorystore

import (
	pb "cloud.google.com/go/memorystore/apiv1/memorystorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AclPolicyRevisionStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AclPolicyRevisionStatus) *krm.AclPolicyRevisionStatusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AclPolicyRevisionStatusObservedState{}
	out.AclPolicyRevision = direct.LazyPtr(in.GetAclPolicyRevision())
	out.AclPolicyRevisionNumber = direct.LazyPtr(in.GetAclPolicyRevisionNumber())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ErrorMessage = direct.LazyPtr(in.GetErrorMessage())
	return out
}

func AclPolicyRevisionStatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AclPolicyRevisionStatusObservedState) *pb.AclPolicyRevisionStatus {
	if in == nil {
		return nil
	}
	out := &pb.AclPolicyRevisionStatus{}
	out.AclPolicyRevision = direct.ValueOf(in.AclPolicyRevision)
	out.AclPolicyRevisionNumber = direct.ValueOf(in.AclPolicyRevisionNumber)
	out.State = direct.Enum_ToProto[pb.AclPolicyRevisionStatus_State](mapCtx, in.State)
	out.ErrorMessage = direct.ValueOf(in.ErrorMessage)
	return out
}

func AclRule_FromProto(mapCtx *direct.MapContext, in *pb.AclRule) *krm.AclRule {
	if in == nil {
		return nil
	}
	out := &krm.AclRule{}
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Rule = direct.LazyPtr(in.GetRule())
	return out
}

func AclRule_ToProto(mapCtx *direct.MapContext, in *krm.AclRule) *pb.AclRule {
	if in == nil {
		return nil
	}
	out := &pb.AclRule{}
	out.Username = direct.ValueOf(in.Username)
	out.Rule = direct.ValueOf(in.Rule)
	return out
}

func InstanceAclPolicyAttachmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceAclPolicyAttachment) *krm.InstanceAclPolicyAttachmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceAclPolicyAttachmentObservedState{}
	out.Instance = direct.LazyPtr(in.GetInstance())
	out.AclPolicyRevisionStatuses = direct.Slice_FromProto(mapCtx, in.AclPolicyRevisionStatuses, AclPolicyRevisionStatusObservedState_FromProto)
	return out
}

func InstanceAclPolicyAttachmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceAclPolicyAttachmentObservedState) *pb.InstanceAclPolicyAttachment {
	if in == nil {
		return nil
	}
	out := &pb.InstanceAclPolicyAttachment{}
	out.Instance = direct.ValueOf(in.Instance)
	out.AclPolicyRevisionStatuses = direct.Slice_ToProto(mapCtx, in.AclPolicyRevisionStatuses, AclPolicyRevisionStatusObservedState_ToProto)
	return out
}

func MemorystoreACLPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AclPolicy) *krm.MemorystoreACLPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MemorystoreACLPolicyObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Etag = in.Etag
	out.InstanceAclPolicyAttachments = direct.Slice_FromProto(mapCtx, in.InstanceAclPolicyAttachments, InstanceAclPolicyAttachmentObservedState_FromProto)
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func MemorystoreACLPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MemorystoreACLPolicyObservedState) *pb.AclPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AclPolicy{}
	out.State = direct.Enum_ToProto[pb.AclPolicy_State](mapCtx, in.State)
	out.Etag = in.Etag
	out.InstanceAclPolicyAttachments = direct.Slice_ToProto(mapCtx, in.InstanceAclPolicyAttachments, InstanceAclPolicyAttachmentObservedState_ToProto)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func MemorystoreACLPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.AclPolicy) *krm.MemorystoreACLPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.MemorystoreACLPolicySpec{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, AclRule_FromProto)
	return out
}

func MemorystoreACLPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.MemorystoreACLPolicySpec) *pb.AclPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AclPolicy{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, AclRule_ToProto)
	return out
}
