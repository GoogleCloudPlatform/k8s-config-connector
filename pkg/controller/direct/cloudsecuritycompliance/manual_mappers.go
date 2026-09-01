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

package cloudsecuritycompliance

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudsecuritycompliance/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	pb "cloud.google.com/go/cloudsecuritycompliance/apiv1/cloudsecuritycompliancepb"
)

// We define manual mapping functions for Framework to skip the undefined CloudControlGroupDetails fields in pb package.

func CloudSecurityComplianceFrameworkSpec_FromProto(mapCtx *direct.MapContext, in *pb.Framework) *krm.CloudSecurityComplianceFrameworkSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudSecurityComplianceFrameworkSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CloudControlDetails = direct.Slice_FromProto(mapCtx, in.CloudControlDetails, CloudControlDetails_FromProto)
	out.Category = direct.EnumSlice_FromProto(mapCtx, in.Category)
	return out
}

func CloudSecurityComplianceFrameworkSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudSecurityComplianceFrameworkSpec) *pb.Framework {
	if in == nil {
		return nil
	}
	out := &pb.Framework{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.CloudControlDetails = direct.Slice_ToProto(mapCtx, in.CloudControlDetails, CloudControlDetails_ToProto)
	out.Category = direct.EnumSlice_ToProto[pb.FrameworkCategory](mapCtx, in.Category)
	return out
}

func CloudSecurityComplianceFrameworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Framework) *krm.CloudSecurityComplianceFrameworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudSecurityComplianceFrameworkObservedState{}
	out.MajorRevisionID = direct.LazyPtr(in.GetMajorRevisionId())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.SupportedCloudProviders = direct.EnumSlice_FromProto(mapCtx, in.SupportedCloudProviders)
	out.SupportedTargetResourceTypes = direct.EnumSlice_FromProto(mapCtx, in.SupportedTargetResourceTypes)
	return out
}

func CloudSecurityComplianceFrameworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudSecurityComplianceFrameworkObservedState) *pb.Framework {
	if in == nil {
		return nil
	}
	out := &pb.Framework{}
	out.MajorRevisionId = direct.ValueOf(in.MajorRevisionID)
	out.Type = direct.Enum_ToProto[pb.Framework_FrameworkType](mapCtx, in.Type)
	out.SupportedCloudProviders = direct.EnumSlice_ToProto[pb.CloudProvider](mapCtx, in.SupportedCloudProviders)
	out.SupportedTargetResourceTypes = direct.EnumSlice_ToProto[pb.TargetResourceType](mapCtx, in.SupportedTargetResourceTypes)
	return out
}

// Since CloudControlGroup and Framework_CloudControlGroupDetails are missing in the public pb package,
// we define dummy mapping functions to prevent the generator from generating invalid references to them.

func CloudControlGroup_FromProto(mapCtx *direct.MapContext, in interface{}) *krm.CloudControlGroup {
	return nil
}

func CloudControlGroup_ToProto(mapCtx *direct.MapContext, in *krm.CloudControlGroup) interface{} {
	return nil
}

func CloudControlGroupObservedState_FromProto(mapCtx *direct.MapContext, in interface{}) *krm.CloudControlGroupObservedState {
	return nil
}

func CloudControlGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudControlGroupObservedState) interface{} {
	return nil
}

func Framework_CloudControlGroupDetails_FromProto(mapCtx *direct.MapContext, in interface{}) *krm.Framework_CloudControlGroupDetails {
	return nil
}

func Framework_CloudControlGroupDetails_ToProto(mapCtx *direct.MapContext, in *krm.Framework_CloudControlGroupDetails) interface{} {
	return nil
}

func Framework_CloudControlGroupDetailsObservedState_FromProto(mapCtx *direct.MapContext, in interface{}) *krm.Framework_CloudControlGroupDetailsObservedState {
	return nil
}

func Framework_CloudControlGroupDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Framework_CloudControlGroupDetailsObservedState) interface{} {
	return nil
}

func CloudControlMetadata_FromProto(mapCtx *direct.MapContext, in *pb.CloudControlMetadata) *krm.CloudControlMetadata {
	if in == nil {
		return nil
	}
	out := &krm.CloudControlMetadata{}
	out.CloudControlDetails = CloudControlDetails_FromProto(mapCtx, in.GetCloudControlDetails())
	out.EnforcementMode = direct.Enum_FromProto(mapCtx, in.GetEnforcementMode())
	return out
}

func CloudControlMetadata_ToProto(mapCtx *direct.MapContext, in *krm.CloudControlMetadata) *pb.CloudControlMetadata {
	if in == nil {
		return nil
	}
	out := &pb.CloudControlMetadata{}
	out.CloudControlDetails = CloudControlDetails_ToProto(mapCtx, in.CloudControlDetails)
	out.EnforcementMode = direct.Enum_ToProto[pb.EnforcementMode](mapCtx, in.EnforcementMode)
	return out
}

func CloudSecurityComplianceFrameworkDeploymentSpec_FromProto(mapCtx *direct.MapContext, in *pb.FrameworkDeployment) *krm.CloudSecurityComplianceFrameworkDeploymentSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudSecurityComplianceFrameworkDeploymentSpec{}
	out.TargetResourceConfig = TargetResourceConfig_FromProto(mapCtx, in.GetTargetResourceConfig())
	out.Framework = FrameworkReference_FromProto(mapCtx, in.GetFramework())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CloudControlMetadata = direct.Slice_FromProto(mapCtx, in.CloudControlMetadata, CloudControlMetadata_FromProto)
	return out
}

func CloudSecurityComplianceFrameworkDeploymentSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudSecurityComplianceFrameworkDeploymentSpec) *pb.FrameworkDeployment {
	if in == nil {
		return nil
	}
	out := &pb.FrameworkDeployment{}
	out.TargetResourceConfig = TargetResourceConfig_ToProto(mapCtx, in.TargetResourceConfig)
	out.Framework = FrameworkReference_ToProto(mapCtx, in.Framework)
	out.Description = direct.ValueOf(in.Description)
	out.CloudControlMetadata = direct.Slice_ToProto(mapCtx, in.CloudControlMetadata, CloudControlMetadata_ToProto)
	return out
}

func CloudSecurityComplianceFrameworkDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FrameworkDeployment) *krm.CloudSecurityComplianceFrameworkDeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudSecurityComplianceFrameworkDeploymentObservedState{}
	out.ComputedTargetResource = direct.LazyPtr(in.GetComputedTargetResource())
	out.DeploymentState = direct.Enum_FromProto(mapCtx, in.GetDeploymentState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.TargetResourceDisplayName = direct.LazyPtr(in.GetTargetResourceDisplayName())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}

func CloudSecurityComplianceFrameworkDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudSecurityComplianceFrameworkDeploymentObservedState) *pb.FrameworkDeployment {
	if in == nil {
		return nil
	}
	out := &pb.FrameworkDeployment{}
	out.ComputedTargetResource = direct.ValueOf(in.ComputedTargetResource)
	out.DeploymentState = direct.Enum_ToProto[pb.DeploymentState](mapCtx, in.DeploymentState)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.TargetResourceDisplayName = direct.ValueOf(in.TargetResourceDisplayName)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}

func CloudControlDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in interface{}) *krm.CloudControlDeploymentObservedState {
	return nil
}

func CloudControlDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudControlDeploymentObservedState) interface{} {
	return nil
}

func CloudControlGroupDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in interface{}) *krm.CloudControlGroupDeploymentObservedState {
	return nil
}

func CloudControlGroupDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudControlGroupDeploymentObservedState) interface{} {
	return nil
}

func CloudControlDeploymentReferenceObservedState_FromProto(mapCtx *direct.MapContext, in interface{}) *krm.CloudControlDeploymentReferenceObservedState {
	return nil
}

func CloudControlDeploymentReferenceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudControlDeploymentReferenceObservedState) interface{} {
	return nil
}
