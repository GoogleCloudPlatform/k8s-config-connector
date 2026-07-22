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
