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
	pb "cloud.google.com/go/cloudsecuritycompliance/apiv1/cloudsecuritycompliancepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudsecuritycompliance/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CloudControlDetails_CloudControlRef_ToProto(mapCtx *direct.MapContext, in *krm.CloudSecurityComplianceCloudControlRef) string {
	if in == nil {
		return ""
	}
	return in.External
}

func CloudControlDetails_CloudControlRef_FromProto(mapCtx *direct.MapContext, in string) *krm.CloudSecurityComplianceCloudControlRef {
	if in == "" {
		return nil
	}
	return &krm.CloudSecurityComplianceCloudControlRef{
		External: in,
	}
}

// Override CloudControlDetails_FromProto manually to use the CloudControlRef field.
func CloudControlDetails_FromProto(mapCtx *direct.MapContext, in *pb.CloudControlDetails) *krm.CloudControlDetails {
	if in == nil {
		return nil
	}
	out := &krm.CloudControlDetails{}
	out.CloudControlRef = CloudControlDetails_CloudControlRef_FromProto(mapCtx, in.GetName())
	out.MajorRevisionID = direct.LazyPtr(in.GetMajorRevisionId())
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, Parameter_FromProto)
	return out
}

// Override CloudControlDetails_ToProto manually to use the CloudControlRef field.
func CloudControlDetails_ToProto(mapCtx *direct.MapContext, in *krm.CloudControlDetails) *pb.CloudControlDetails {
	if in == nil {
		return nil
	}
	out := &pb.CloudControlDetails{}
	out.Name = CloudControlDetails_CloudControlRef_ToProto(mapCtx, in.CloudControlRef)
	out.MajorRevisionId = direct.ValueOf(in.MajorRevisionID)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, Parameter_ToProto)
	return out
}

// Override CloudSecurityFrameworkSpec_FromProto and CloudSecurityFrameworkSpec_ToProto manually
// to map the actual fields supported by the Go client library SDK Framework proto.

func CloudSecurityFrameworkSpec_FromProto(mapCtx *direct.MapContext, in *pb.Framework) *krm.CloudSecurityFrameworkSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudSecurityFrameworkSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CloudControlDetails = direct.Slice_FromProto(mapCtx, in.GetCloudControlDetails(), CloudControlDetails_FromProto)
	out.Category = direct.EnumSlice_FromProto(mapCtx, in.GetCategory())
	return out
}

func CloudSecurityFrameworkSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudSecurityFrameworkSpec) *pb.Framework {
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

// Override CloudSecurityFrameworkObservedState_FromProto and CloudSecurityFrameworkObservedState_ToProto
// to map the actual status/observed state fields supported by the Go client library SDK Framework proto.

func CloudSecurityFrameworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Framework) *krm.CloudSecurityFrameworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudSecurityFrameworkObservedState{}
	out.MajorRevisionID = direct.LazyPtr(in.GetMajorRevisionId())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.SupportedCloudProviders = direct.EnumSlice_FromProto(mapCtx, in.GetSupportedCloudProviders())
	out.SupportedTargetResourceTypes = direct.EnumSlice_FromProto(mapCtx, in.GetSupportedTargetResourceTypes())
	out.SupportedEnforcementModes = direct.EnumSlice_FromProto(mapCtx, in.GetSupportedEnforcementModes())
	return out
}

func CloudSecurityFrameworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudSecurityFrameworkObservedState) *pb.Framework {
	if in == nil {
		return nil
	}
	out := &pb.Framework{}
	out.MajorRevisionId = direct.ValueOf(in.MajorRevisionID)
	out.Type = direct.Enum_ToProto[pb.Framework_FrameworkType](mapCtx, in.Type)
	out.SupportedCloudProviders = direct.EnumSlice_ToProto[pb.CloudProvider](mapCtx, in.SupportedCloudProviders)
	out.SupportedTargetResourceTypes = direct.EnumSlice_ToProto[pb.TargetResourceType](mapCtx, in.SupportedTargetResourceTypes)
	out.SupportedEnforcementModes = direct.EnumSlice_ToProto[pb.EnforcementMode](mapCtx, in.SupportedEnforcementModes)
	return out
}

// Override CloudSecurityComplianceFrameworkSpec_FromProto and CloudSecurityComplianceFrameworkSpec_ToProto
// to prevent compilation errors regarding the missing CloudControlGroupDetails field in the SDK package.

func CloudSecurityComplianceFrameworkSpec_FromProto(mapCtx *direct.MapContext, in *pb.Framework) *krm.CloudSecurityComplianceFrameworkSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudSecurityComplianceFrameworkSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CloudControlDetails = direct.Slice_FromProto(mapCtx, in.GetCloudControlDetails(), CloudControlDetails_FromProto)
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
	return out
}

// Override CloudSecurityComplianceFrameworkObservedState_FromProto and CloudSecurityComplianceFrameworkObservedState_ToProto
// to prevent compilation errors regarding the missing CloudControlGroupDetails field in the SDK package.

func CloudSecurityComplianceFrameworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Framework) *krm.CloudSecurityComplianceFrameworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudSecurityComplianceFrameworkObservedState{}
	return out
}

func CloudSecurityComplianceFrameworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudSecurityComplianceFrameworkObservedState) *pb.Framework {
	if in == nil {
		return nil
	}
	out := &pb.Framework{}
	return out
}

// Define manual stubs for types that do not exist in the GCP Go Client library SDK pb package
// but are defined in framework_types.go from master.

func CloudControlGroup_FromProto(mapCtx *direct.MapContext, in any) *krm.CloudControlGroup {
	return nil
}

func CloudControlGroup_ToProto(mapCtx *direct.MapContext, in *krm.CloudControlGroup) any {
	return nil
}

func CloudControlGroupObservedState_FromProto(mapCtx *direct.MapContext, in any) *krm.CloudControlGroupObservedState {
	return nil
}

func CloudControlGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudControlGroupObservedState) any {
	return nil
}

func Framework_CloudControlGroupDetails_FromProto(mapCtx *direct.MapContext, in any) *krm.Framework_CloudControlGroupDetails {
	return nil
}

func Framework_CloudControlGroupDetails_ToProto(mapCtx *direct.MapContext, in *krm.Framework_CloudControlGroupDetails) any {
	return nil
}

func Framework_CloudControlGroupDetailsObservedState_FromProto(mapCtx *direct.MapContext, in any) *krm.Framework_CloudControlGroupDetailsObservedState {
	return nil
}

func Framework_CloudControlGroupDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Framework_CloudControlGroupDetailsObservedState) any {
	return nil
}
