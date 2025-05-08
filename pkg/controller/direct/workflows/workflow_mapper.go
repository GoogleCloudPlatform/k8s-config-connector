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

package workflows

import (
	pb "cloud.google.com/go/workflows/apiv1/workflowspb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workflows/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func WorkflowsWorkflowStateError_FromProto(mapCtx *direct.MapContext, in *pb.Workflow_StateError) *krm.WorkflowsWorkflow_StateError {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowsWorkflow_StateError{
		Details: direct.LazyPtr(in.GetDetails()),
		Type:    direct.LazyPtr(in.GetType().String()),
	}
	return out
}

func WorkflowsWorkflowStateError_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowsWorkflow_StateError) *pb.Workflow_StateError {
	if in == nil {
		return nil
	}
	out := &pb.Workflow_StateError{
		Details: direct.ValueOf(in.Details),
		Type:    direct.Enum_ToProto[pb.Workflow_StateError_Type](mapCtx, in.Type),
	}
	return out
}

func WorkflowsWorkflowObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Workflow) *krm.WorkflowsWorkflowObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowsWorkflowObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.RevisionId = direct.LazyPtr(in.GetRevisionId())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.RevisionCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRevisionCreateTime())
	out.StateError = WorkflowsWorkflowStateError_FromProto(mapCtx, in.GetStateError())
	out.AllKmsKeys = []refs.KMSCryptoKeyRef{}
	for _, kmsKey := range in.AllKmsKeys {
		out.AllKmsKeys = append(out.AllKmsKeys, refs.KMSCryptoKeyRef{External: kmsKey})
	}
	out.AllKmsKeysVersions = []string{}
	for _, kmsKeyVersion := range in.AllKmsKeysVersions {
		out.AllKmsKeysVersions = append(out.AllKmsKeysVersions, kmsKeyVersion)
	}
	out.CryptoKeyVersion = direct.LazyPtr(in.GetCryptoKeyVersion())
	return out
}

func WorkflowsWorkflowObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowsWorkflowObservedState) *pb.Workflow {
	if in == nil {
		return nil
	}
	out := &pb.Workflow{}
	out.State = direct.Enum_ToProto[pb.Workflow_State](mapCtx, in.State)
	out.RevisionId = direct.ValueOf(in.RevisionId)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.RevisionCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.RevisionCreateTime)
	out.StateError = WorkflowsWorkflowStateError_ToProto(mapCtx, in.StateError)
	out.AllKmsKeys = []string{}
	for _, kmsKey := range in.AllKmsKeys {
		out.AllKmsKeys = append(out.AllKmsKeys, kmsKey.External)
	}
	out.AllKmsKeysVersions = []string{}
	for _, kmsKeyVersion := range in.AllKmsKeysVersions {
		out.AllKmsKeysVersions = append(out.AllKmsKeysVersions, kmsKeyVersion)
	}
	out.CryptoKeyVersion = direct.ValueOf(in.CryptoKeyVersion)
	return out
}

func WorkflowsWorkflowSpec_FromProto(mapCtx *direct.MapContext, in *pb.Workflow) *krm.WorkflowsWorkflowSpec {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowsWorkflowSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refs.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.SourceContents = direct.LazyPtr(in.GetSourceContents())
	if in.GetCryptoKeyName() != "" {
		out.KMSCryptoKeyRef = &refs.KMSCryptoKeyRef{External: in.GetCryptoKeyName()}
	}
	out.CallLogLevel = direct.Enum_FromProto(mapCtx, in.GetCallLogLevel())
	out.UserEnvVars = in.GetUserEnvVars()
	out.ExecutionHistoryLevel = direct.Enum_FromProto(mapCtx, in.GetExecutionHistoryLevel())
	out.Tags = in.GetTags()
	return out
}

func WorkflowsWorkflowSpec_SourceContents_ToProto(mapCtx *direct.MapContext, in *string) *pb.Workflow_SourceContents {
	if in == nil {
		return nil
	}
	out := &pb.Workflow_SourceContents{
		SourceContents: direct.ValueOf(in),
	}
	return out

}
func WorkflowsWorkflowSpec_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowsWorkflowSpec) *pb.Workflow {
	if in == nil {
		return nil
	}
	out := &pb.Workflow{}
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	if oneof := WorkflowsWorkflowSpec_SourceContents_ToProto(mapCtx, in.SourceContents); oneof != nil {
		out.SourceCode = oneof
	}
	if in.KMSCryptoKeyRef != nil {
		out.CryptoKeyName = in.KMSCryptoKeyRef.External
	}
	if in.CallLogLevel != nil {
		out.CallLogLevel = direct.Enum_ToProto[pb.Workflow_CallLogLevel](mapCtx, in.CallLogLevel)
	}
	out.UserEnvVars = in.UserEnvVars
	if in.ExecutionHistoryLevel != nil {
		out.ExecutionHistoryLevel = direct.Enum_ToProto[pb.ExecutionHistoryLevel](mapCtx, in.ExecutionHistoryLevel)
	}
	out.Tags = in.Tags
	return out
}
