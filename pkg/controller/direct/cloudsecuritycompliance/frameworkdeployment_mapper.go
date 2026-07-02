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
	out.CloudControlDeploymentReferences = direct.Slice_FromProto(mapCtx, in.CloudControlDeploymentReferences, CloudControlDeploymentReferenceObservedState_FromProto)
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
	out.CloudControlDeploymentReferences = direct.Slice_ToProto(mapCtx, in.CloudControlDeploymentReferences, CloudControlDeploymentReferenceObservedState_ToProto)
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
	out.Etag = direct.LazyPtr(in.GetEtag())
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
	out.Etag = direct.ValueOf(in.Etag)
	return out
}

func TargetResourceConfig_FromProto(mapCtx *direct.MapContext, in *pb.TargetResourceConfig) *krm.TargetResourceConfig {
	if in == nil {
		return nil
	}
	out := &krm.TargetResourceConfig{}
	out.ExistingTargetResource = direct.LazyPtr(in.GetExistingTargetResource())
	out.TargetResourceCreationConfig = TargetResourceCreationConfig_FromProto(mapCtx, in.GetTargetResourceCreationConfig())
	return out
}

func TargetResourceConfig_ToProto(mapCtx *direct.MapContext, in *krm.TargetResourceConfig) *pb.TargetResourceConfig {
	if in == nil {
		return nil
	}
	out := &pb.TargetResourceConfig{}
	if in.ExistingTargetResource != nil {
		out.ResourceConfig = &pb.TargetResourceConfig_ExistingTargetResource{
			ExistingTargetResource: *in.ExistingTargetResource,
		}
	} else if in.TargetResourceCreationConfig != nil {
		out.ResourceConfig = &pb.TargetResourceConfig_TargetResourceCreationConfig{
			TargetResourceCreationConfig: TargetResourceCreationConfig_ToProto(mapCtx, in.TargetResourceCreationConfig),
		}
	}
	return out
}

func TargetResourceCreationConfig_FromProto(mapCtx *direct.MapContext, in *pb.TargetResourceCreationConfig) *krm.TargetResourceCreationConfig {
	if in == nil {
		return nil
	}
	out := &krm.TargetResourceCreationConfig{}
	out.FolderCreationConfig = FolderCreationConfig_FromProto(mapCtx, in.GetFolderCreationConfig())
	out.ProjectCreationConfig = ProjectCreationConfig_FromProto(mapCtx, in.GetProjectCreationConfig())
	return out
}

func TargetResourceCreationConfig_ToProto(mapCtx *direct.MapContext, in *krm.TargetResourceCreationConfig) *pb.TargetResourceCreationConfig {
	if in == nil {
		return nil
	}
	out := &pb.TargetResourceCreationConfig{}
	if in.FolderCreationConfig != nil {
		out.ResourceCreationConfig = &pb.TargetResourceCreationConfig_FolderCreationConfig{
			FolderCreationConfig: FolderCreationConfig_ToProto(mapCtx, in.FolderCreationConfig),
		}
	} else if in.ProjectCreationConfig != nil {
		out.ResourceCreationConfig = &pb.TargetResourceCreationConfig_ProjectCreationConfig{
			ProjectCreationConfig: ProjectCreationConfig_ToProto(mapCtx, in.ProjectCreationConfig),
		}
	}
	return out
}

func FolderCreationConfig_FromProto(mapCtx *direct.MapContext, in *pb.FolderCreationConfig) *krm.FolderCreationConfig {
	if in == nil {
		return nil
	}
	out := &krm.FolderCreationConfig{}
	out.Parent = direct.LazyPtr(in.GetParent())
	out.FolderDisplayName = direct.LazyPtr(in.GetFolderDisplayName())
	return out
}

func FolderCreationConfig_ToProto(mapCtx *direct.MapContext, in *krm.FolderCreationConfig) *pb.FolderCreationConfig {
	if in == nil {
		return nil
	}
	out := &pb.FolderCreationConfig{}
	out.Parent = direct.ValueOf(in.Parent)
	out.FolderDisplayName = direct.ValueOf(in.FolderDisplayName)
	return out
}

func ProjectCreationConfig_FromProto(mapCtx *direct.MapContext, in *pb.ProjectCreationConfig) *krm.ProjectCreationConfig {
	if in == nil {
		return nil
	}
	out := &krm.ProjectCreationConfig{}
	out.Parent = direct.LazyPtr(in.GetParent())
	out.ProjectDisplayName = direct.LazyPtr(in.GetProjectDisplayName())
	out.BillingAccountID = direct.LazyPtr(in.GetBillingAccountId())
	return out
}

func ProjectCreationConfig_ToProto(mapCtx *direct.MapContext, in *krm.ProjectCreationConfig) *pb.ProjectCreationConfig {
	if in == nil {
		return nil
	}
	out := &pb.ProjectCreationConfig{}
	out.Parent = direct.ValueOf(in.Parent)
	out.ProjectDisplayName = direct.ValueOf(in.ProjectDisplayName)
	out.BillingAccountId = direct.ValueOf(in.BillingAccountID)
	return out
}

func FrameworkReference_FromProto(mapCtx *direct.MapContext, in *pb.FrameworkReference) *krm.FrameworkReference {
	if in == nil {
		return nil
	}
	out := &krm.FrameworkReference{}
	if in.GetFramework() != "" {
		out.FrameworkRef = &krm.CloudSecurityComplianceFrameworkRef{
			External: in.GetFramework(),
		}
	}
	out.MajorRevisionID = direct.LazyPtr(in.GetMajorRevisionId())
	return out
}

func FrameworkReference_ToProto(mapCtx *direct.MapContext, in *krm.FrameworkReference) *pb.FrameworkReference {
	if in == nil {
		return nil
	}
	out := &pb.FrameworkReference{}
	if in.FrameworkRef != nil {
		out.Framework = in.FrameworkRef.External
	}
	out.MajorRevisionId = in.MajorRevisionID
	return out
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

func CloudControlDetails_FromProto(mapCtx *direct.MapContext, in *pb.CloudControlDetails) *krm.CloudControlDetails {
	if in == nil {
		return nil
	}
	out := &krm.CloudControlDetails{}
	if in.GetName() != "" {
		out.CloudControlRef = &krm.CloudSecurityComplianceCloudControlRef{
			External: in.GetName(),
		}
	}
	out.MajorRevisionID = direct.LazyPtr(in.GetMajorRevisionId())
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, Parameter_FromProto)
	return out
}

func CloudControlDetails_ToProto(mapCtx *direct.MapContext, in *krm.CloudControlDetails) *pb.CloudControlDetails {
	if in == nil {
		return nil
	}
	out := &pb.CloudControlDetails{}
	if in.CloudControlRef != nil {
		out.Name = in.CloudControlRef.External
	}
	out.MajorRevisionId = direct.ValueOf(in.MajorRevisionID)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, Parameter_ToProto)
	return out
}

func CloudControlDeploymentReferenceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudControlDeploymentReference) *krm.CloudControlDeploymentReferenceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudControlDeploymentReferenceObservedState{}
	out.CloudControlDeployment = direct.LazyPtr(in.GetCloudControlDeployment())
	return out
}

func CloudControlDeploymentReferenceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudControlDeploymentReferenceObservedState) *pb.CloudControlDeploymentReference {
	if in == nil {
		return nil
	}
	out := &pb.CloudControlDeploymentReference{}
	out.CloudControlDeployment = direct.ValueOf(in.CloudControlDeployment)
	return out
}
