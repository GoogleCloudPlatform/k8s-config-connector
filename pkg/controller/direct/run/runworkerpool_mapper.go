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

package run

import (
	pb "cloud.google.com/go/run/apiv2/runpb"
	networkservicesv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	krmrunv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	apipb "google.golang.org/genproto/googleapis/api"
)

func BinaryAuthorization_UseDefault_ToProto(mapCtx *direct.MapContext, in *bool) *pb.BinaryAuthorization_UseDefault {
	if in == nil {
		return nil
	}
	return &pb.BinaryAuthorization_UseDefault{UseDefault: *in}
}

func EnvVar_Value_ToProto(mapCtx *direct.MapContext, in *string) *pb.EnvVar_Value {
	if in == nil {
		return nil
	}
	return &pb.EnvVar_Value{Value: *in}
}

func RunWorkerPoolSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmrunv1alpha1.RunWorkerPoolSpec) *pb.WorkerPool {
	if in == nil {
		return nil
	}
	out := &pb.WorkerPool{}
	// MISSING: Name (handled by controller)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: Labels (handled by controller via metadata.labels)
	out.Annotations = direct.MapStringString_ToProto(mapCtx, in.Annotations)
	out.Client = direct.ValueOf(in.Client)
	out.ClientVersion = direct.ValueOf(in.ClientVersion)
	out.LaunchStage = direct.Enum_ToProto[apipb.LaunchStage](mapCtx, in.LaunchStage)
	out.BinaryAuthorization = BinaryAuthorization_v1alpha1_ToProto(mapCtx, in.BinaryAuthorization)
	out.Template = WorkerPoolRevisionTemplate_v1alpha1_ToProto(mapCtx, in.Template)
	out.InstanceSplits = direct.Slice_ToProto(mapCtx, in.InstanceSplits, InstanceSplit_v1alpha1_ToProto)
	out.Scaling = WorkerPoolScaling_v1alpha1_ToProto(mapCtx, in.Scaling)
	out.CustomAudiences = in.CustomAudiences
	return out
}

func RunWorkerPoolObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.WorkerPool) *krmrunv1alpha1.RunWorkerPoolObservedState {
	if in == nil {
		return nil
	}
	out := &krmrunv1alpha1.RunWorkerPoolObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	out.Generation = direct.LazyPtr(in.GetGeneration())
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Creator = direct.LazyPtr(in.GetCreator())
	out.LastModifier = direct.LazyPtr(in.GetLastModifier())
	out.Template = WorkerPoolRevisionTemplateObservedState_v1alpha1_FromProto(mapCtx, in.GetTemplate())
	out.ObservedGeneration = direct.LazyPtr(in.GetObservedGeneration())
	out.TerminalCondition = RunCondition_v1alpha1_FromProto(mapCtx, in.GetTerminalCondition())
	out.ObservedConditions = direct.Slice_FromProto(mapCtx, in.Conditions, RunCondition_v1alpha1_FromProto)
	out.LatestReadyRevision = direct.LazyPtr(in.GetLatestReadyRevision())
	out.LatestCreatedRevision = direct.LazyPtr(in.GetLatestCreatedRevision())
	out.InstanceSplitStatuses = direct.Slice_FromProto(mapCtx, in.InstanceSplitStatuses, InstanceSplitStatus_v1alpha1_FromProto)
	out.CustomAudiences = in.CustomAudiences
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}

func ServiceMesh_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmrunv1alpha1.ServiceMesh) *pb.ServiceMesh {
	if in == nil {
		return nil
	}
	out := &pb.ServiceMesh{}
	if in.MeshRef != nil {
		out.Mesh = in.MeshRef.External
	}
	return out
}

func ServiceMesh_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ServiceMesh) *krmrunv1alpha1.ServiceMesh {
	if in == nil {
		return nil
	}
	out := &krmrunv1alpha1.ServiceMesh{}
	if in.GetMesh() != "" {
		out.MeshRef = &networkservicesv1alpha1.NetworkServicesMeshRef{External: in.GetMesh()}
	}
	return out
}

func RunCondition_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Condition) *krmrunv1alpha1.RunCondition {
	if in == nil {
		return nil
	}
	out := &krmrunv1alpha1.RunCondition{}
	out.Type = direct.LazyPtr(in.GetType())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.LastTransitionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastTransitionTime())
	out.Severity = direct.Enum_FromProto(mapCtx, in.GetSeverity())
	out.Reason = direct.Enum_FromProto(mapCtx, in.GetReason())
	out.RevisionReason = direct.Enum_FromProto(mapCtx, in.GetRevisionReason())
	out.ExecutionReason = direct.Enum_FromProto(mapCtx, in.GetExecutionReason())
	return out
}

func RunCondition_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmrunv1alpha1.RunCondition) *pb.Condition {
	if in == nil {
		return nil
	}
	out := &pb.Condition{}
	out.Type = direct.ValueOf(in.Type)
	out.State = direct.Enum_ToProto[pb.Condition_State](mapCtx, in.State)
	out.Message = direct.ValueOf(in.Message)
	out.LastTransitionTime = direct.StringTimestamp_ToProto(mapCtx, in.LastTransitionTime)
	out.Severity = direct.Enum_ToProto[pb.Condition_Severity](mapCtx, in.Severity)
	if oneof := Condition_Reason_ToProto(mapCtx, in.Reason); oneof != nil {
		out.Reasons = oneof
	}
	if oneof := Condition_RevisionReason_ToProto(mapCtx, in.RevisionReason); oneof != nil {
		out.Reasons = oneof
	}
	if oneof := Condition_ExecutionReason_ToProto(mapCtx, in.ExecutionReason); oneof != nil {
		out.Reasons = oneof
	}
	return out
}
