// Copyright 2024 Google LLC
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
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	api "google.golang.org/api/gkehub/v1beta"
)

func featureMembershipSpecKRMtoMembershipFeatureSpecAPI(r *krm.GKEHubFeatureMembershipSpec) (*api.MembershipFeatureSpec, error) {
	var acm *api.ConfigManagementMembershipSpec
	var err error
	if r.Configmanagement != nil {
		acm, err = convertKRMtoAPI_ConfigManagement(r.Configmanagement)
		if err != nil {
			return nil, err
		}
	}
	var poco *api.PolicyControllerMembershipSpec
	if r.Policycontroller != nil {
		poco = convertKRMtoAPI_Policycontroller(r.Policycontroller)
	}

	var mesh *api.ServiceMeshMembershipSpec
	if r.Mesh != nil {
		mesh = convertKRMtoAPI_ServiceMesh(r.Mesh)
	}

	return &api.MembershipFeatureSpec{
		Configmanagement: acm,
		Policycontroller: poco,
		Mesh:             mesh,
	}, nil
}

func FeaturemembershipBinauthz_FromProto(mapCtx *direct.MapContext, r *api.ConfigManagementBinauthzConfig) *krm.FeaturemembershipBinauthz {
	if r == nil {
		return nil
	}

	return &krm.FeaturemembershipBinauthz{
		Enabled: &r.Enabled,
	}
}

func ConfigManagementMembershipSpec_FromProto(mapCtx *direct.MapContext, r *api.ConfigManagementMembershipSpec) *krm.FeaturemembershipConfigmanagement {
	if r == nil {
		return nil
	}
	return &krm.FeaturemembershipConfigmanagement{
		Binauthz:   FeaturemembershipBinauthz_FromProto(mapCtx, r.Binauthz),
		ConfigSync: ConfigSyncMembershipSpec_FromProto(mapCtx, r.ConfigSync),
	}
}

func OciMembershipSpec_FromProto(mapCtx *direct.MapContext, r *api.ConfigManagementOciConfig) *krm.FeaturemembershipOci {
	if r == nil {
		return nil
	}

	out := &krm.FeaturemembershipOci{}
	if r.GcpServiceAccountEmail != "" {
		out.GcpServiceAccountRef = &refs.IAMServiceAccountRef{
			External: r.GcpServiceAccountEmail,
		}
	}
	out.PolicyDir = direct.LazyPtr(r.PolicyDir)
	out.SecretType = direct.LazyPtr(r.SecretType)
	out.SyncRepo = direct.LazyPtr(r.SyncRepo)
	out.SyncWaitSecs = direct.LazyPtr(fmt.Sprint(r.SyncWaitSecs))

	return out
}

func GitMembershipSpec_FromProto(mapCtx *direct.MapContext, r *api.ConfigManagementGitConfig) *krm.FeaturemembershipGit {
	if r == nil {
		return nil
	}

	out := &krm.FeaturemembershipGit{}
	if r.GcpServiceAccountEmail != "" {
		out.GcpServiceAccountRef = &refs.IAMServiceAccountRef{
			External: r.GcpServiceAccountEmail,
		}
	}
	out.HttpsProxy = direct.LazyPtr(r.HttpsProxy)
	out.PolicyDir = direct.LazyPtr(r.PolicyDir)
	out.SecretType = direct.LazyPtr(r.SecretType)
	out.SyncBranch = direct.LazyPtr(r.SyncBranch)
	out.SyncRepo = direct.LazyPtr(r.SyncRepo)
	out.SyncRev = direct.LazyPtr(r.SyncRev)
	out.SyncWaitSecs = direct.LazyPtr(fmt.Sprint(r.SyncWaitSecs))

	return out
}

func ConfigSyncMembershipSpec_FromProto(mapCtx *direct.MapContext, r *api.ConfigManagementConfigSync) *krm.FeaturemembershipConfigSync {
	if r == nil {
		return nil
	}

	out := &krm.FeaturemembershipConfigSync{}
	out.PreventDrift = direct.LazyPtr(r.PreventDrift)
	out.SourceFormat = direct.LazyPtr(r.SourceFormat)
	out.Git = GitMembershipSpec_FromProto(mapCtx, r.Git)
	out.Oci = OciMembershipSpec_FromProto(mapCtx, r.Oci)
	if r.MetricsGcpServiceAccountEmail != "" {
		out.MetricsGcpServiceAccountRef = &refs.MetricsGcpServiceAccountRef{
			External: r.MetricsGcpServiceAccountEmail,
		}
	}

	return out
}

func TemplateLibraryMembershipSpec_FromProto(mapCtx *direct.MapContext, r *api.PolicyControllerTemplateLibraryConfig) *krm.FeaturemembershipTemplateLibrary {
	if r == nil {
		return nil
	}

	out := &krm.FeaturemembershipTemplateLibrary{}
	out.Installation = direct.LazyPtr(r.Installation)

	return out
}

func FeaturemembershipPolicyContent_FromProto(mapCtx *direct.MapContext, r *api.PolicyControllerPolicyContentSpec) *krm.FeaturemembershipPolicyContent {
	if r == nil {
		return nil
	}

	out := &krm.FeaturemembershipPolicyContent{}
	out.TemplateLibrary = TemplateLibraryMembershipSpec_FromProto(mapCtx, r.TemplateLibrary)

	return out
}

func FeaturemembershipPolicyControllerHubConfig_FromProto(mapCtx *direct.MapContext, r *api.PolicyControllerHubConfig) *krm.FeaturemembershipPolicyControllerHubConfig {
	if r == nil {
		return nil
	}

	out := &krm.FeaturemembershipPolicyControllerHubConfig{}
	out.AuditIntervalSeconds = direct.LazyPtr(r.AuditIntervalSeconds)
	out.ConstraintViolationLimit = direct.LazyPtr(r.ConstraintViolationLimit)
	out.ExemptableNamespaces = r.ExemptableNamespaces
	out.InstallSpec = direct.LazyPtr(r.InstallSpec)
	out.LogDeniesEnabled = direct.LazyPtr(r.LogDeniesEnabled)
	out.Monitoring = FeaturemembershipMonitoring_FromProto(mapCtx, r.Monitoring)
	out.MutationEnabled = direct.LazyPtr(r.MutationEnabled)
	out.PolicyContent = FeaturemembershipPolicyContent_FromProto(mapCtx, r.PolicyContent)
	out.ReferentialRulesEnabled = direct.LazyPtr(r.ReferentialRulesEnabled)

	return out
}

func FeaturemembershipMonitoring_FromProto(mapCtx *direct.MapContext, r *api.PolicyControllerMonitoringConfig) *krm.FeaturemembershipMonitoring {
	if r == nil {
		return nil
	}
	out := &krm.FeaturemembershipMonitoring{}
	out.Backends = r.Backends
	return out
}

func PolicycontrollerMembershipSpec_FromProto(mapCtx *direct.MapContext, r *api.PolicyControllerMembershipSpec) *krm.FeaturemembershipPolicycontroller {
	if r == nil {
		return nil
	}

	out := &krm.FeaturemembershipPolicycontroller{}
	out.PolicyControllerHubConfig = *FeaturemembershipPolicyControllerHubConfig_FromProto(mapCtx, r.PolicyControllerHubConfig)
	out.Version = direct.LazyPtr(r.Version)

	return out
}

func ServiceMeshMembershipSpec_FromProto(mapCtx *direct.MapContext, r *api.ServiceMeshMembershipSpec) *krm.FeaturemembershipMesh {
	if r == nil {
		return nil
	}

	out := &krm.FeaturemembershipMesh{}
	out.ControlPlane = direct.LazyPtr(r.ControlPlane)
	out.Management = direct.LazyPtr(r.Management)

	return out
}

func GKEHubFeatureMembershipSpec_FromProto(mapCtx *direct.MapContext, r *api.MembershipFeatureSpec) *krm.GKEHubFeatureMembershipSpec {
	if r == nil {
		return nil
	}
	out := &krm.GKEHubFeatureMembershipSpec{}
	out.Configmanagement = ConfigManagementMembershipSpec_FromProto(mapCtx, r.Configmanagement)
	out.Policycontroller = PolicycontrollerMembershipSpec_FromProto(mapCtx, r.Policycontroller)
	out.Mesh = ServiceMeshMembershipSpec_FromProto(mapCtx, r.Mesh)

	return out
}
