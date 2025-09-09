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

package clouddeploy

import (
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CloudDeployTargetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Target) *krm.CloudDeployTargetSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudDeployTargetSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	out.RequireApproval = direct.LazyPtr(in.GetRequireApproval())
	out.Gke = GkeCluster_FromProto(mapCtx, in.GetGke())
	out.AnthosCluster = AnthosCluster_FromProto(mapCtx, in.GetAnthosCluster())
	out.Run = CloudRunLocation_FromProto(mapCtx, in.GetRun())
	out.MultiTarget = MultiTarget_FromProto(mapCtx, in.GetMultiTarget())
	out.CustomTarget = CustomTarget_FromProto(mapCtx, in.GetCustomTarget())
	if in.AssociatedEntities != nil {
		out.AssociatedEntities = make(map[string]*krm.AssociatedEntities)
		for k, v := range in.AssociatedEntities {
			out.AssociatedEntities[k] = AssociatedEntities_FromProto(mapCtx, v)
		}
	}

	// Manually convert the slice of pointers to avoid a compilation error.
	// The direct.Slice_FromProto helper function can incorrectly infer the type
	// as a slice of values ([]krm.ExecutionConfig) instead of a slice of pointers
	// ([]*krm.ExecutionConfig), causing a type mismatch.
	if in.ExecutionConfigs != nil {
		out.ExecutionConfigs = make([]*krm.ExecutionConfig, len(in.ExecutionConfigs))
		for i, v := range in.ExecutionConfigs {
			out.ExecutionConfigs[i] = ExecutionConfig_FromProto(mapCtx, v)
		}
	}

	out.DeployParameters = in.DeployParameters
	return out
}
func CloudDeployTargetSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudDeployTargetSpec) *pb.Target {
	if in == nil {
		return nil
	}
	out := &pb.Target{}
	out.Description = direct.ValueOf(in.Description)
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	out.RequireApproval = direct.ValueOf(in.RequireApproval)
	if oneof := GkeCluster_ToProto(mapCtx, in.Gke); oneof != nil {
		out.DeploymentTarget = &pb.Target_Gke{Gke: oneof}
	}
	if oneof := AnthosCluster_ToProto(mapCtx, in.AnthosCluster); oneof != nil {
		out.DeploymentTarget = &pb.Target_AnthosCluster{AnthosCluster: oneof}
	}
	if oneof := CloudRunLocation_ToProto(mapCtx, in.Run); oneof != nil {
		out.DeploymentTarget = &pb.Target_Run{Run: oneof}
	}
	if oneof := MultiTarget_ToProto(mapCtx, in.MultiTarget); oneof != nil {
		out.DeploymentTarget = &pb.Target_MultiTarget{MultiTarget: oneof}
	}
	if oneof := CustomTarget_ToProto(mapCtx, in.CustomTarget); oneof != nil {
		out.DeploymentTarget = &pb.Target_CustomTarget{CustomTarget: oneof}
	}

	// Manually convert the map of pointers to avoid a compilation error.
	if in.AssociatedEntities != nil {
		out.AssociatedEntities = make(map[string]*pb.AssociatedEntities)
		for k, v := range in.AssociatedEntities {
			out.AssociatedEntities[k] = AssociatedEntities_ToProto(mapCtx, v)
		}
	}

	// Manually convert the slice of pointers to avoid a compilation error.
	// The direct.Slice_ToProto helper function can incorrectly infer the type
	// as a slice of values ([]pb.ExecutionConfig) instead of a slice of pointers
	// ([]*pb.ExecutionConfig), causing a type mismatch.
	if in.ExecutionConfigs != nil {
		out.ExecutionConfigs = make([]*pb.ExecutionConfig, len(in.ExecutionConfigs))
		for i, v := range in.ExecutionConfigs {
			out.ExecutionConfigs[i] = ExecutionConfig_ToProto(mapCtx, v)
		}
	}
	out.DeployParameters = in.DeployParameters
	return out
}
func CloudDeployTargetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Target) *krm.CloudDeployTargetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudDeployTargetObservedState{}
	out.TargetId = direct.LazyPtr(in.GetTargetId())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func CloudDeployTargetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudDeployTargetObservedState) *pb.Target {
	if in == nil {
		return nil
	}
	out := &pb.Target{}
	out.TargetId = direct.ValueOf(in.TargetId)
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func AssociatedEntities_FromProto(mapCtx *direct.MapContext, in *pb.AssociatedEntities) *krm.AssociatedEntities {
	if in == nil {
		return nil
	}
	out := &krm.AssociatedEntities{}
	out.GkeClusters = direct.Slice_FromProto(mapCtx, in.GkeClusters, GkeCluster_FromProto)
	out.AnthosClusters = direct.Slice_FromProto(mapCtx, in.AnthosClusters, AnthosCluster_FromProto)
	return out
}
func AssociatedEntities_ToProto(mapCtx *direct.MapContext, in *krm.AssociatedEntities) *pb.AssociatedEntities {
	if in == nil {
		return nil
	}
	out := &pb.AssociatedEntities{}
	out.GkeClusters = direct.Slice_ToProto(mapCtx, in.GkeClusters, GkeCluster_ToProto)
	out.AnthosClusters = direct.Slice_ToProto(mapCtx, in.AnthosClusters, AnthosCluster_ToProto)
	return out
}
