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

package securitycenter

import (
	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func OrganizationSettings_FromProto(mapCtx *direct.MapContext, in *pb.OrganizationSettings) *krm.OrganizationSettings {
	if in == nil {
		return nil
	}
	out := &krm.OrganizationSettings{}
	out.Name = direct.LazyPtr(in.GetName())
	out.EnableAssetDiscovery = direct.LazyPtr(in.GetEnableAssetDiscovery())
	out.AssetDiscoveryConfig = OrganizationSettings_AssetDiscoveryConfig_FromProto(mapCtx, in.GetAssetDiscoveryConfig())
	return out
}
func OrganizationSettings_ToProto(mapCtx *direct.MapContext, in *krm.OrganizationSettings) *pb.OrganizationSettings {
	if in == nil {
		return nil
	}
	out := &pb.OrganizationSettings{}
	out.Name = direct.ValueOf(in.Name)
	out.EnableAssetDiscovery = direct.ValueOf(in.EnableAssetDiscovery)
	out.AssetDiscoveryConfig = OrganizationSettings_AssetDiscoveryConfig_ToProto(mapCtx, in.AssetDiscoveryConfig)
	return out
}
func OrganizationSettings_AssetDiscoveryConfig_FromProto(mapCtx *direct.MapContext, in *pb.OrganizationSettings_AssetDiscoveryConfig) *krm.OrganizationSettings_AssetDiscoveryConfig {
	if in == nil {
		return nil
	}
	out := &krm.OrganizationSettings_AssetDiscoveryConfig{}
	out.ProjectIds = in.ProjectIds
	out.InclusionMode = direct.Enum_FromProto(mapCtx, in.GetInclusionMode())
	out.FolderIds = in.FolderIds
	return out
}
func OrganizationSettings_AssetDiscoveryConfig_ToProto(mapCtx *direct.MapContext, in *krm.OrganizationSettings_AssetDiscoveryConfig) *pb.OrganizationSettings_AssetDiscoveryConfig {
	if in == nil {
		return nil
	}
	out := &pb.OrganizationSettings_AssetDiscoveryConfig{}
	out.ProjectIds = in.ProjectIds
	out.InclusionMode = direct.Enum_ToProto[pb.OrganizationSettings_AssetDiscoveryConfig_InclusionMode](mapCtx, in.InclusionMode)
	out.FolderIds = in.FolderIds
	return out
}
func SecuritycenterOrganizationSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OrganizationSettings) *krm.SecuritycenterOrganizationSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterOrganizationSettingsObservedState{}
	// MISSING: Name
	// MISSING: EnableAssetDiscovery
	// MISSING: AssetDiscoveryConfig
	return out
}
func SecuritycenterOrganizationSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterOrganizationSettingsObservedState) *pb.OrganizationSettings {
	if in == nil {
		return nil
	}
	out := &pb.OrganizationSettings{}
	// MISSING: Name
	// MISSING: EnableAssetDiscovery
	// MISSING: AssetDiscoveryConfig
	return out
}
func SecuritycenterOrganizationSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.OrganizationSettings) *krm.SecuritycenterOrganizationSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterOrganizationSettingsSpec{}
	// MISSING: Name
	// MISSING: EnableAssetDiscovery
	// MISSING: AssetDiscoveryConfig
	return out
}
func SecuritycenterOrganizationSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterOrganizationSettingsSpec) *pb.OrganizationSettings {
	if in == nil {
		return nil
	}
	out := &pb.OrganizationSettings{}
	// MISSING: Name
	// MISSING: EnableAssetDiscovery
	// MISSING: AssetDiscoveryConfig
	return out
}
