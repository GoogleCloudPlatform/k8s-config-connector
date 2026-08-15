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

package gkehub

import (
	binaryauthorizationv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/binaryauthorization/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	gkehubv1 "google.golang.org/api/gkehub/v1"
)

func GKEHubFleetBinaryAuthorizationConfig_ToAPI(mapCtx *direct.MapContext, in *krm.BinaryAuthorizationConfig) *gkehubv1.BinaryAuthorizationConfig {
	if in == nil {
		return nil
	}
	out := &gkehubv1.BinaryAuthorizationConfig{}
	out.EvaluationMode = direct.ValueOf(in.EvaluationMode)
	for _, ref := range in.PolicyBindingsRefs {
		if ref.External != "" {
			out.PolicyBindings = append(out.PolicyBindings, &gkehubv1.PolicyBinding{
				Name: ref.External,
			})
		}
	}
	return out
}

func GKEHubFleetBinaryAuthorizationConfig_FromAPI(mapCtx *direct.MapContext, in *gkehubv1.BinaryAuthorizationConfig) *krm.BinaryAuthorizationConfig {
	if in == nil {
		return nil
	}
	out := &krm.BinaryAuthorizationConfig{}
	out.EvaluationMode = direct.LazyPtr(in.EvaluationMode)
	for _, pb := range in.PolicyBindings {
		if pb != nil && pb.Name != "" {
			out.PolicyBindingsRefs = append(out.PolicyBindingsRefs, binaryauthorizationv1alpha1.BinaryAuthorizationPlatformPolicyRef{
				External: pb.Name,
			})
		}
	}
	return out
}

func GKEHubFleetSecurityPostureConfig_ToAPI(mapCtx *direct.MapContext, in *krm.SecurityPostureConfig) *gkehubv1.SecurityPostureConfig {
	if in == nil {
		return nil
	}
	out := &gkehubv1.SecurityPostureConfig{}
	out.Mode = direct.ValueOf(in.Mode)
	out.VulnerabilityMode = direct.ValueOf(in.VulnerabilityMode)
	return out
}

func GKEHubFleetSecurityPostureConfig_FromAPI(mapCtx *direct.MapContext, in *gkehubv1.SecurityPostureConfig) *krm.SecurityPostureConfig {
	if in == nil {
		return nil
	}
	out := &krm.SecurityPostureConfig{}
	out.Mode = direct.LazyPtr(in.Mode)
	out.VulnerabilityMode = direct.LazyPtr(in.VulnerabilityMode)
	return out
}

func GKEHubFleetCompliancePostureConfig_ToAPI(mapCtx *direct.MapContext, in *krm.CompliancePostureConfig) *gkehubv1.CompliancePostureConfig {
	if in == nil {
		return nil
	}
	out := &gkehubv1.CompliancePostureConfig{}
	out.Mode = direct.ValueOf(in.Mode)
	for _, cs := range in.ComplianceStandards {
		out.ComplianceStandards = append(out.ComplianceStandards, &gkehubv1.ComplianceStandard{
			Standard: direct.ValueOf(cs.Standard),
		})
	}
	return out
}

func GKEHubFleetCompliancePostureConfig_FromAPI(mapCtx *direct.MapContext, in *gkehubv1.CompliancePostureConfig) *krm.CompliancePostureConfig {
	if in == nil {
		return nil
	}
	out := &krm.CompliancePostureConfig{}
	out.Mode = direct.LazyPtr(in.Mode)
	for _, cs := range in.ComplianceStandards {
		if cs == nil {
			continue
		}
		out.ComplianceStandards = append(out.ComplianceStandards, krm.CompliancePostureConfig_ComplianceStandard{
			Standard: direct.LazyPtr(cs.Standard),
		})
	}
	return out
}

func GKEHubFleetDefaultClusterConfig_ToAPI(mapCtx *direct.MapContext, in *krm.DefaultClusterConfig) *gkehubv1.DefaultClusterConfig {
	if in == nil {
		return nil
	}
	out := &gkehubv1.DefaultClusterConfig{}
	out.BinaryAuthorizationConfig = GKEHubFleetBinaryAuthorizationConfig_ToAPI(mapCtx, in.BinaryAuthorizationConfig)
	out.CompliancePostureConfig = GKEHubFleetCompliancePostureConfig_ToAPI(mapCtx, in.CompliancePostureConfig)
	out.SecurityPostureConfig = GKEHubFleetSecurityPostureConfig_ToAPI(mapCtx, in.SecurityPostureConfig)
	return out
}

func GKEHubFleetDefaultClusterConfig_FromAPI(mapCtx *direct.MapContext, in *gkehubv1.DefaultClusterConfig) *krm.DefaultClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.DefaultClusterConfig{}
	out.BinaryAuthorizationConfig = GKEHubFleetBinaryAuthorizationConfig_FromAPI(mapCtx, in.BinaryAuthorizationConfig)
	out.CompliancePostureConfig = GKEHubFleetCompliancePostureConfig_FromAPI(mapCtx, in.CompliancePostureConfig)
	out.SecurityPostureConfig = GKEHubFleetSecurityPostureConfig_FromAPI(mapCtx, in.SecurityPostureConfig)
	return out
}

func GKEHubFleetSpec_ToAPI(mapCtx *direct.MapContext, in *krm.GKEHubFleetSpec) *gkehubv1.Fleet {
	if in == nil {
		return nil
	}
	out := &gkehubv1.Fleet{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = in.Labels
	out.DefaultClusterConfig = GKEHubFleetDefaultClusterConfig_ToAPI(mapCtx, in.DefaultClusterConfig)
	return out
}

func GKEHubFleetSpec_FromAPI(mapCtx *direct.MapContext, in *gkehubv1.Fleet, id *krm.GKEHubFleetIdentity) *krm.GKEHubFleetSpec {
	if in == nil {
		return nil
	}
	out := &krm.GKEHubFleetSpec{}
	if id == nil && in.Name != "" {
		id = &krm.GKEHubFleetIdentity{}
		if err := id.FromExternal(in.Name); err != nil {
			mapCtx.Errorf("parsing GKEHubFleet resource name %q: %v", in.Name, err)
		}
	}
	if id != nil {
		out.ProjectRef = &refs.ProjectRef{External: "projects/" + id.ProjectID}
		out.Location = direct.LazyPtr(id.Location)
		out.ResourceID = direct.LazyPtr(id.FleetID)
	}
	out.DisplayName = direct.LazyPtr(in.DisplayName)
	out.Labels = in.Labels
	out.DefaultClusterConfig = GKEHubFleetDefaultClusterConfig_FromAPI(mapCtx, in.DefaultClusterConfig)
	return out
}

func GKEHubFleetStatus_FromAPI(mapCtx *direct.MapContext, in *gkehubv1.Fleet) *krm.GKEHubFleetStatus {
	if in == nil {
		return nil
	}
	out := &krm.GKEHubFleetStatus{}
	out.ObservedState = &krm.GKEHubFleetObservedState{
		CreateTime: direct.LazyPtr(in.CreateTime),
		UpdateTime: direct.LazyPtr(in.UpdateTime),
		DeleteTime: direct.LazyPtr(in.DeleteTime),
		Uid:        direct.LazyPtr(in.Uid),
	}
	if in.State != nil {
		out.ObservedState.State = &krm.FleetLifecycleStateObservedState{
			Code: direct.LazyPtr(in.State.Code),
		}
	}
	return out
}

func GKEHubFleetStatus_ToAPI(mapCtx *direct.MapContext, in *krm.GKEHubFleetStatus) *gkehubv1.Fleet {
	if in == nil || in.ObservedState == nil {
		return nil
	}
	out := &gkehubv1.Fleet{}
	out.CreateTime = direct.ValueOf(in.ObservedState.CreateTime)
	out.UpdateTime = direct.ValueOf(in.ObservedState.UpdateTime)
	out.DeleteTime = direct.ValueOf(in.ObservedState.DeleteTime)
	out.Uid = direct.ValueOf(in.ObservedState.Uid)
	if in.ObservedState.State != nil {
		out.State = &gkehubv1.FleetLifecycleState{
			Code: direct.ValueOf(in.ObservedState.State.Code),
		}
	}
	return out
}
