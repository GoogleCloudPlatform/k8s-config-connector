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
// krm.group: configdeployment.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.config.v1

package configdeployment

import (
	pb "cloud.google.com/go/config/apiv1/configpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/configdeployment/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DeploymentUnit_FromProto(mapCtx *direct.MapContext, in *pb.DeploymentUnit) *krm.DeploymentUnit {
	if in == nil {
		return nil
	}
	out := &krm.DeploymentUnit{}
	out.ID = direct.LazyPtr(in.GetId())
	if in.GetDeployment() != "" {
		out.DeploymentRef = &krm.ConfigDeploymentRef{External: in.GetDeployment()}
	}
	out.Dependencies = in.GetDependencies()
	return out
}

func DeploymentUnit_ToProto(mapCtx *direct.MapContext, in *krm.DeploymentUnit) *pb.DeploymentUnit {
	if in == nil {
		return nil
	}
	out := &pb.DeploymentUnit{}
	out.Id = direct.ValueOf(in.ID)
	if in.DeploymentRef != nil {
		out.Deployment = direct.LazyPtr(in.DeploymentRef.External)
	}
	out.Dependencies = in.Dependencies
	return out
}

func ConfigDeploymentGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.DeploymentGroup) *krm.ConfigDeploymentGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConfigDeploymentGroupSpec{}
	out.Labels = in.GetLabels()
	out.Annotations = in.GetAnnotations()
	out.DeploymentUnits = direct.Slice_FromProto(mapCtx, in.GetDeploymentUnits(), DeploymentUnit_FromProto)
	return out
}

func ConfigDeploymentGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConfigDeploymentGroupSpec) *pb.DeploymentGroup {
	if in == nil {
		return nil
	}
	out := &pb.DeploymentGroup{}
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DeploymentUnits = direct.Slice_ToProto(mapCtx, in.DeploymentUnits, DeploymentUnit_ToProto)
	return out
}

func ConfigDeploymentGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeploymentGroup) *krm.ConfigDeploymentGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConfigDeploymentGroupObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateDescription = direct.LazyPtr(in.GetStateDescription())
	out.ProvisioningState = direct.Enum_FromProto(mapCtx, in.GetProvisioningState())
	out.ProvisioningStateDescription = direct.LazyPtr(in.GetProvisioningStateDescription())
	out.ProvisioningError = direct.Status_FromProto(mapCtx, in.GetProvisioningError())
	return out
}

func ConfigDeploymentGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConfigDeploymentGroupObservedState) *pb.DeploymentGroup {
	if in == nil {
		return nil
	}
	out := &pb.DeploymentGroup{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.DeploymentGroup_State](mapCtx, in.State)
	out.StateDescription = direct.ValueOf(in.StateDescription)
	out.ProvisioningState = direct.Enum_ToProto[pb.DeploymentGroup_ProvisioningState](mapCtx, in.ProvisioningState)
	out.ProvisioningStateDescription = direct.ValueOf(in.ProvisioningStateDescription)
	out.ProvisioningError = direct.Status_ToProto(mapCtx, in.ProvisioningError)
	return out
}
