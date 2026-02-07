// Copyright 2025 Google LLC
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
// krm.group: workflows.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.workflows.v1

package workflows

import (
	pb "cloud.google.com/go/workflows/apiv1/workflowspb"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workflows/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Workflow_StateError_FromProto(mapCtx *direct.MapContext, in *pb.Workflow_StateError) *krm.Workflow_StateError {
	if in == nil {
		return nil
	}
	out := &krm.Workflow_StateError{}
	out.Details = direct.LazyPtr(in.GetDetails())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Workflow_StateError_ToProto(mapCtx *direct.MapContext, in *krm.Workflow_StateError) *pb.Workflow_StateError {
	if in == nil {
		return nil
	}
	out := &pb.Workflow_StateError{}
	out.Details = direct.ValueOf(in.Details)
	out.Type = direct.Enum_ToProto[pb.Workflow_StateError_Type](mapCtx, in.Type)
	return out
}
func WorkflowsWorkflowObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Workflow) *krm.WorkflowsWorkflowObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowsWorkflowObservedState{}
	// MISSING: Name
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: RevisionID
	// (near miss): "RevisionID" vs "RevisionId"
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.RevisionCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRevisionCreateTime())
	// MISSING: Labels
	out.StateError = WorkflowsWorkflow_StateError_FromProto(mapCtx, in.GetStateError())
	// MISSING: AllKMSKeys
	// MISSING: AllKMSKeysVersions
	// MISSING: CryptoKeyVersion
	return out
}
func WorkflowsWorkflowObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowsWorkflowObservedState) *pb.Workflow {
	if in == nil {
		return nil
	}
	out := &pb.Workflow{}
	// MISSING: Name
	out.State = direct.Enum_ToProto[pb.Workflow_State](mapCtx, in.State)
	// MISSING: RevisionID
	// (near miss): "RevisionID" vs "RevisionId"
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.RevisionCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.RevisionCreateTime)
	// MISSING: Labels
	out.StateError = WorkflowsWorkflow_StateError_ToProto(mapCtx, in.StateError)
	// MISSING: AllKMSKeys
	// MISSING: AllKMSKeysVersions
	// MISSING: CryptoKeyVersion
	return out
}
func WorkflowsWorkflowSpec_FromProto(mapCtx *direct.MapContext, in *pb.Workflow) *krm.WorkflowsWorkflowSpec {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowsWorkflowSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: RevisionID
	// MISSING: Labels
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.SourceContents = direct.LazyPtr(in.GetSourceContents())
	if in.GetCryptoKeyName() != "" {
		out.CryptoKeyNameRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetCryptoKeyName()}
	}
	out.CallLogLevel = direct.Enum_FromProto(mapCtx, in.GetCallLogLevel())
	out.UserEnvVars = in.UserEnvVars
	out.ExecutionHistoryLevel = direct.Enum_FromProto(mapCtx, in.GetExecutionHistoryLevel())
	// MISSING: AllKMSKeys
	// MISSING: AllKMSKeysVersions
	// MISSING: CryptoKeyVersion
	out.Tags = in.Tags
	return out
}
func WorkflowsWorkflowSpec_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowsWorkflowSpec) *pb.Workflow {
	if in == nil {
		return nil
	}
	out := &pb.Workflow{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	// MISSING: RevisionID
	// MISSING: Labels
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	if oneof := WorkflowsWorkflowSpec_SourceContents_ToProto(mapCtx, in.SourceContents); oneof != nil {
		out.SourceCode = oneof
	}
	if in.CryptoKeyNameRef != nil {
		out.CryptoKeyName = in.CryptoKeyNameRef.External
	}
	out.CallLogLevel = direct.Enum_ToProto[pb.Workflow_CallLogLevel](mapCtx, in.CallLogLevel)
	out.UserEnvVars = in.UserEnvVars
	out.ExecutionHistoryLevel = direct.Enum_ToProto[pb.ExecutionHistoryLevel](mapCtx, in.ExecutionHistoryLevel)
	// MISSING: AllKMSKeys
	// MISSING: AllKMSKeysVersions
	// MISSING: CryptoKeyVersion
	out.Tags = in.Tags
	return out
}
func WorkflowsWorkflowSpec_SourceContents_ToProto(mapCtx *direct.MapContext, in *string) *pb.Workflow_SourceContents {
	if in == nil {
		return nil
	}
	return &pb.Workflow_SourceContents{SourceContents: *in}
}
