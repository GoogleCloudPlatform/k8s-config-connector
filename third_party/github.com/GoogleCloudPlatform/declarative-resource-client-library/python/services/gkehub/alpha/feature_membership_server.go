// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/alpha/gkehub_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/alpha"
)

// FeatureMembershipServer implements the gRPC interface for FeatureMembership.
type FeatureMembershipServer struct{}

// ProtoToFeatureMembershipMeshManagementEnum converts a FeatureMembershipMeshManagementEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipMeshManagementEnum(e alphapb.GkehubAlphaFeatureMembershipMeshManagementEnum) *alpha.FeatureMembershipMeshManagementEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureMembershipMeshManagementEnum_name[int32(e)]; ok {
		e := alpha.FeatureMembershipMeshManagementEnum(n[len("GkehubAlphaFeatureMembershipMeshManagementEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipMeshControlPlaneEnum converts a FeatureMembershipMeshControlPlaneEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipMeshControlPlaneEnum(e alphapb.GkehubAlphaFeatureMembershipMeshControlPlaneEnum) *alpha.FeatureMembershipMeshControlPlaneEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureMembershipMeshControlPlaneEnum_name[int32(e)]; ok {
		e := alpha.FeatureMembershipMeshControlPlaneEnum(n[len("GkehubAlphaFeatureMembershipMeshControlPlaneEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum converts a FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(e alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum) *alpha.FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum_name[int32(e)]; ok {
		e := alpha.FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(n[len("GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(e alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum) *alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum_name[int32(e)]; ok {
		e := alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(n[len("GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(e alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum) *alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum_name[int32(e)]; ok {
		e := alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(n[len("GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(e alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum) *alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum_name[int32(e)]; ok {
		e := alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(n[len("GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipMesh converts a FeatureMembershipMesh object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipMesh(p *alphapb.GkehubAlphaFeatureMembershipMesh) *alpha.FeatureMembershipMesh {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipMesh{
		Management:   ProtoToGkehubAlphaFeatureMembershipMeshManagementEnum(p.GetManagement()),
		ControlPlane: ProtoToGkehubAlphaFeatureMembershipMeshControlPlaneEnum(p.GetControlPlane()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagement converts a FeatureMembershipConfigmanagement object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagement(p *alphapb.GkehubAlphaFeatureMembershipConfigmanagement) *alpha.FeatureMembershipConfigmanagement {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipConfigmanagement{
		ConfigSync:          ProtoToGkehubAlphaFeatureMembershipConfigmanagementConfigSync(p.GetConfigSync()),
		PolicyController:    ProtoToGkehubAlphaFeatureMembershipConfigmanagementPolicyController(p.GetPolicyController()),
		Binauthz:            ProtoToGkehubAlphaFeatureMembershipConfigmanagementBinauthz(p.GetBinauthz()),
		HierarchyController: ProtoToGkehubAlphaFeatureMembershipConfigmanagementHierarchyController(p.GetHierarchyController()),
		Version:             dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSync converts a FeatureMembershipConfigmanagementConfigSync object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagementConfigSync(p *alphapb.GkehubAlphaFeatureMembershipConfigmanagementConfigSync) *alpha.FeatureMembershipConfigmanagementConfigSync {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipConfigmanagementConfigSync{
		Git:                           ProtoToGkehubAlphaFeatureMembershipConfigmanagementConfigSyncGit(p.GetGit()),
		SourceFormat:                  dcl.StringOrNil(p.GetSourceFormat()),
		PreventDrift:                  dcl.Bool(p.GetPreventDrift()),
		MetricsGcpServiceAccountEmail: dcl.StringOrNil(p.GetMetricsGcpServiceAccountEmail()),
		Oci:                           ProtoToGkehubAlphaFeatureMembershipConfigmanagementConfigSyncOci(p.GetOci()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSyncGit converts a FeatureMembershipConfigmanagementConfigSyncGit object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagementConfigSyncGit(p *alphapb.GkehubAlphaFeatureMembershipConfigmanagementConfigSyncGit) *alpha.FeatureMembershipConfigmanagementConfigSyncGit {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipConfigmanagementConfigSyncGit{
		SyncRepo:               dcl.StringOrNil(p.GetSyncRepo()),
		SyncBranch:             dcl.StringOrNil(p.GetSyncBranch()),
		PolicyDir:              dcl.StringOrNil(p.GetPolicyDir()),
		SyncWaitSecs:           dcl.StringOrNil(p.GetSyncWaitSecs()),
		SyncRev:                dcl.StringOrNil(p.GetSyncRev()),
		SecretType:             dcl.StringOrNil(p.GetSecretType()),
		HttpsProxy:             dcl.StringOrNil(p.GetHttpsProxy()),
		GcpServiceAccountEmail: dcl.StringOrNil(p.GetGcpServiceAccountEmail()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSyncOci converts a FeatureMembershipConfigmanagementConfigSyncOci object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagementConfigSyncOci(p *alphapb.GkehubAlphaFeatureMembershipConfigmanagementConfigSyncOci) *alpha.FeatureMembershipConfigmanagementConfigSyncOci {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipConfigmanagementConfigSyncOci{
		SyncRepo:               dcl.StringOrNil(p.GetSyncRepo()),
		PolicyDir:              dcl.StringOrNil(p.GetPolicyDir()),
		SyncWaitSecs:           dcl.StringOrNil(p.GetSyncWaitSecs()),
		SecretType:             dcl.StringOrNil(p.GetSecretType()),
		GcpServiceAccountEmail: dcl.StringOrNil(p.GetGcpServiceAccountEmail()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementPolicyController converts a FeatureMembershipConfigmanagementPolicyController object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagementPolicyController(p *alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyController) *alpha.FeatureMembershipConfigmanagementPolicyController {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipConfigmanagementPolicyController{
		Enabled:                  dcl.Bool(p.GetEnabled()),
		ReferentialRulesEnabled:  dcl.Bool(p.GetReferentialRulesEnabled()),
		LogDeniesEnabled:         dcl.Bool(p.GetLogDeniesEnabled()),
		MutationEnabled:          dcl.Bool(p.GetMutationEnabled()),
		Monitoring:               ProtoToGkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoring(p.GetMonitoring()),
		TemplateLibraryInstalled: dcl.Bool(p.GetTemplateLibraryInstalled()),
		AuditIntervalSeconds:     dcl.StringOrNil(p.GetAuditIntervalSeconds()),
	}
	for _, r := range p.GetExemptableNamespaces() {
		obj.ExemptableNamespaces = append(obj.ExemptableNamespaces, r)
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementPolicyControllerMonitoring converts a FeatureMembershipConfigmanagementPolicyControllerMonitoring object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoring(p *alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoring) *alpha.FeatureMembershipConfigmanagementPolicyControllerMonitoring {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipConfigmanagementPolicyControllerMonitoring{}
	for _, r := range p.GetBackends() {
		obj.Backends = append(obj.Backends, *ProtoToGkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(r))
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementBinauthz converts a FeatureMembershipConfigmanagementBinauthz object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagementBinauthz(p *alphapb.GkehubAlphaFeatureMembershipConfigmanagementBinauthz) *alpha.FeatureMembershipConfigmanagementBinauthz {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipConfigmanagementBinauthz{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementHierarchyController converts a FeatureMembershipConfigmanagementHierarchyController object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagementHierarchyController(p *alphapb.GkehubAlphaFeatureMembershipConfigmanagementHierarchyController) *alpha.FeatureMembershipConfigmanagementHierarchyController {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipConfigmanagementHierarchyController{
		Enabled:                         dcl.Bool(p.GetEnabled()),
		EnablePodTreeLabels:             dcl.Bool(p.GetEnablePodTreeLabels()),
		EnableHierarchicalResourceQuota: dcl.Bool(p.GetEnableHierarchicalResourceQuota()),
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontroller converts a FeatureMembershipPolicycontroller object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipPolicycontroller(p *alphapb.GkehubAlphaFeatureMembershipPolicycontroller) *alpha.FeatureMembershipPolicycontroller {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipPolicycontroller{
		Version:                   dcl.StringOrNil(p.GetVersion()),
		PolicyControllerHubConfig: ProtoToGkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfig(p.GetPolicyControllerHubConfig()),
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfig converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfig object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfig(p *alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfig) *alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfig{
		InstallSpec:              ProtoToGkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(p.GetInstallSpec()),
		ReferentialRulesEnabled:  dcl.Bool(p.GetReferentialRulesEnabled()),
		LogDeniesEnabled:         dcl.Bool(p.GetLogDeniesEnabled()),
		MutationEnabled:          dcl.Bool(p.GetMutationEnabled()),
		Monitoring:               ProtoToGkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(p.GetMonitoring()),
		AuditIntervalSeconds:     dcl.Int64OrNil(p.GetAuditIntervalSeconds()),
		ConstraintViolationLimit: dcl.Int64OrNil(p.GetConstraintViolationLimit()),
		PolicyContent:            ProtoToGkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(p.GetPolicyContent()),
	}
	for _, r := range p.GetExemptableNamespaces() {
		obj.ExemptableNamespaces = append(obj.ExemptableNamespaces, r)
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(p *alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) *alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring{}
	for _, r := range p.GetBackends() {
		obj.Backends = append(obj.Backends, *ProtoToGkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(r))
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(p *alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) *alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent{
		TemplateLibrary: ProtoToGkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(p.GetTemplateLibrary()),
	}
	return obj
}

// ProtoToFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(p *alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary) *alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary{
		Installation: ProtoToGkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(p.GetInstallation()),
	}
	return obj
}

// ProtoToFeatureMembership converts a FeatureMembership resource from its proto representation.
func ProtoToFeatureMembership(p *alphapb.GkehubAlphaFeatureMembership) *alpha.FeatureMembership {
	obj := &alpha.FeatureMembership{
		Mesh:               ProtoToGkehubAlphaFeatureMembershipMesh(p.GetMesh()),
		Configmanagement:   ProtoToGkehubAlphaFeatureMembershipConfigmanagement(p.GetConfigmanagement()),
		Policycontroller:   ProtoToGkehubAlphaFeatureMembershipPolicycontroller(p.GetPolicycontroller()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
		Feature:            dcl.StringOrNil(p.GetFeature()),
		Membership:         dcl.StringOrNil(p.GetMembership()),
		MembershipLocation: dcl.StringOrNil(p.GetMembershipLocation()),
	}
	return obj
}

// FeatureMembershipMeshManagementEnumToProto converts a FeatureMembershipMeshManagementEnum enum to its proto representation.
func GkehubAlphaFeatureMembershipMeshManagementEnumToProto(e *alpha.FeatureMembershipMeshManagementEnum) alphapb.GkehubAlphaFeatureMembershipMeshManagementEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureMembershipMeshManagementEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureMembershipMeshManagementEnum_value["FeatureMembershipMeshManagementEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureMembershipMeshManagementEnum(v)
	}
	return alphapb.GkehubAlphaFeatureMembershipMeshManagementEnum(0)
}

// FeatureMembershipMeshControlPlaneEnumToProto converts a FeatureMembershipMeshControlPlaneEnum enum to its proto representation.
func GkehubAlphaFeatureMembershipMeshControlPlaneEnumToProto(e *alpha.FeatureMembershipMeshControlPlaneEnum) alphapb.GkehubAlphaFeatureMembershipMeshControlPlaneEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureMembershipMeshControlPlaneEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureMembershipMeshControlPlaneEnum_value["FeatureMembershipMeshControlPlaneEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureMembershipMeshControlPlaneEnum(v)
	}
	return alphapb.GkehubAlphaFeatureMembershipMeshControlPlaneEnum(0)
}

// FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnumToProto converts a FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum enum to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnumToProto(e *alpha.FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum) alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum_value["FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(v)
	}
	return alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(0)
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnumToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum enum to its proto representation.
func GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnumToProto(e *alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum) alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum_value["FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(v)
	}
	return alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(0)
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnumToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum enum to its proto representation.
func GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnumToProto(e *alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum) alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum_value["FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(v)
	}
	return alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(0)
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnumToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum enum to its proto representation.
func GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnumToProto(e *alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum) alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum_value["FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(v)
	}
	return alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(0)
}

// FeatureMembershipMeshToProto converts a FeatureMembershipMesh object to its proto representation.
func GkehubAlphaFeatureMembershipMeshToProto(o *alpha.FeatureMembershipMesh) *alphapb.GkehubAlphaFeatureMembershipMesh {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipMesh{}
	p.SetManagement(GkehubAlphaFeatureMembershipMeshManagementEnumToProto(o.Management))
	p.SetControlPlane(GkehubAlphaFeatureMembershipMeshControlPlaneEnumToProto(o.ControlPlane))
	return p
}

// FeatureMembershipConfigmanagementToProto converts a FeatureMembershipConfigmanagement object to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementToProto(o *alpha.FeatureMembershipConfigmanagement) *alphapb.GkehubAlphaFeatureMembershipConfigmanagement {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipConfigmanagement{}
	p.SetConfigSync(GkehubAlphaFeatureMembershipConfigmanagementConfigSyncToProto(o.ConfigSync))
	p.SetPolicyController(GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerToProto(o.PolicyController))
	p.SetBinauthz(GkehubAlphaFeatureMembershipConfigmanagementBinauthzToProto(o.Binauthz))
	p.SetHierarchyController(GkehubAlphaFeatureMembershipConfigmanagementHierarchyControllerToProto(o.HierarchyController))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// FeatureMembershipConfigmanagementConfigSyncToProto converts a FeatureMembershipConfigmanagementConfigSync object to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementConfigSyncToProto(o *alpha.FeatureMembershipConfigmanagementConfigSync) *alphapb.GkehubAlphaFeatureMembershipConfigmanagementConfigSync {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipConfigmanagementConfigSync{}
	p.SetGit(GkehubAlphaFeatureMembershipConfigmanagementConfigSyncGitToProto(o.Git))
	p.SetSourceFormat(dcl.ValueOrEmptyString(o.SourceFormat))
	p.SetPreventDrift(dcl.ValueOrEmptyBool(o.PreventDrift))
	p.SetMetricsGcpServiceAccountEmail(dcl.ValueOrEmptyString(o.MetricsGcpServiceAccountEmail))
	p.SetOci(GkehubAlphaFeatureMembershipConfigmanagementConfigSyncOciToProto(o.Oci))
	return p
}

// FeatureMembershipConfigmanagementConfigSyncGitToProto converts a FeatureMembershipConfigmanagementConfigSyncGit object to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementConfigSyncGitToProto(o *alpha.FeatureMembershipConfigmanagementConfigSyncGit) *alphapb.GkehubAlphaFeatureMembershipConfigmanagementConfigSyncGit {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipConfigmanagementConfigSyncGit{}
	p.SetSyncRepo(dcl.ValueOrEmptyString(o.SyncRepo))
	p.SetSyncBranch(dcl.ValueOrEmptyString(o.SyncBranch))
	p.SetPolicyDir(dcl.ValueOrEmptyString(o.PolicyDir))
	p.SetSyncWaitSecs(dcl.ValueOrEmptyString(o.SyncWaitSecs))
	p.SetSyncRev(dcl.ValueOrEmptyString(o.SyncRev))
	p.SetSecretType(dcl.ValueOrEmptyString(o.SecretType))
	p.SetHttpsProxy(dcl.ValueOrEmptyString(o.HttpsProxy))
	p.SetGcpServiceAccountEmail(dcl.ValueOrEmptyString(o.GcpServiceAccountEmail))
	return p
}

// FeatureMembershipConfigmanagementConfigSyncOciToProto converts a FeatureMembershipConfigmanagementConfigSyncOci object to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementConfigSyncOciToProto(o *alpha.FeatureMembershipConfigmanagementConfigSyncOci) *alphapb.GkehubAlphaFeatureMembershipConfigmanagementConfigSyncOci {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipConfigmanagementConfigSyncOci{}
	p.SetSyncRepo(dcl.ValueOrEmptyString(o.SyncRepo))
	p.SetPolicyDir(dcl.ValueOrEmptyString(o.PolicyDir))
	p.SetSyncWaitSecs(dcl.ValueOrEmptyString(o.SyncWaitSecs))
	p.SetSecretType(dcl.ValueOrEmptyString(o.SecretType))
	p.SetGcpServiceAccountEmail(dcl.ValueOrEmptyString(o.GcpServiceAccountEmail))
	return p
}

// FeatureMembershipConfigmanagementPolicyControllerToProto converts a FeatureMembershipConfigmanagementPolicyController object to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerToProto(o *alpha.FeatureMembershipConfigmanagementPolicyController) *alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyController {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyController{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetReferentialRulesEnabled(dcl.ValueOrEmptyBool(o.ReferentialRulesEnabled))
	p.SetLogDeniesEnabled(dcl.ValueOrEmptyBool(o.LogDeniesEnabled))
	p.SetMutationEnabled(dcl.ValueOrEmptyBool(o.MutationEnabled))
	p.SetMonitoring(GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringToProto(o.Monitoring))
	p.SetTemplateLibraryInstalled(dcl.ValueOrEmptyBool(o.TemplateLibraryInstalled))
	p.SetAuditIntervalSeconds(dcl.ValueOrEmptyString(o.AuditIntervalSeconds))
	sExemptableNamespaces := make([]string, len(o.ExemptableNamespaces))
	for i, r := range o.ExemptableNamespaces {
		sExemptableNamespaces[i] = r
	}
	p.SetExemptableNamespaces(sExemptableNamespaces)
	return p
}

// FeatureMembershipConfigmanagementPolicyControllerMonitoringToProto converts a FeatureMembershipConfigmanagementPolicyControllerMonitoring object to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringToProto(o *alpha.FeatureMembershipConfigmanagementPolicyControllerMonitoring) *alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoring {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoring{}
	sBackends := make([]alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum, len(o.Backends))
	for i, r := range o.Backends {
		sBackends[i] = alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum_value[string(r)])
	}
	p.SetBackends(sBackends)
	return p
}

// FeatureMembershipConfigmanagementBinauthzToProto converts a FeatureMembershipConfigmanagementBinauthz object to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementBinauthzToProto(o *alpha.FeatureMembershipConfigmanagementBinauthz) *alphapb.GkehubAlphaFeatureMembershipConfigmanagementBinauthz {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipConfigmanagementBinauthz{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// FeatureMembershipConfigmanagementHierarchyControllerToProto converts a FeatureMembershipConfigmanagementHierarchyController object to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementHierarchyControllerToProto(o *alpha.FeatureMembershipConfigmanagementHierarchyController) *alphapb.GkehubAlphaFeatureMembershipConfigmanagementHierarchyController {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipConfigmanagementHierarchyController{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetEnablePodTreeLabels(dcl.ValueOrEmptyBool(o.EnablePodTreeLabels))
	p.SetEnableHierarchicalResourceQuota(dcl.ValueOrEmptyBool(o.EnableHierarchicalResourceQuota))
	return p
}

// FeatureMembershipPolicycontrollerToProto converts a FeatureMembershipPolicycontroller object to its proto representation.
func GkehubAlphaFeatureMembershipPolicycontrollerToProto(o *alpha.FeatureMembershipPolicycontroller) *alphapb.GkehubAlphaFeatureMembershipPolicycontroller {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipPolicycontroller{}
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetPolicyControllerHubConfig(GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigToProto(o.PolicyControllerHubConfig))
	return p
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfig object to its proto representation.
func GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigToProto(o *alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfig) *alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfig{}
	p.SetInstallSpec(GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnumToProto(o.InstallSpec))
	p.SetReferentialRulesEnabled(dcl.ValueOrEmptyBool(o.ReferentialRulesEnabled))
	p.SetLogDeniesEnabled(dcl.ValueOrEmptyBool(o.LogDeniesEnabled))
	p.SetMutationEnabled(dcl.ValueOrEmptyBool(o.MutationEnabled))
	p.SetMonitoring(GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringToProto(o.Monitoring))
	p.SetAuditIntervalSeconds(dcl.ValueOrEmptyInt64(o.AuditIntervalSeconds))
	p.SetConstraintViolationLimit(dcl.ValueOrEmptyInt64(o.ConstraintViolationLimit))
	p.SetPolicyContent(GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentToProto(o.PolicyContent))
	sExemptableNamespaces := make([]string, len(o.ExemptableNamespaces))
	for i, r := range o.ExemptableNamespaces {
		sExemptableNamespaces[i] = r
	}
	p.SetExemptableNamespaces(sExemptableNamespaces)
	return p
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring object to its proto representation.
func GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringToProto(o *alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) *alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring{}
	sBackends := make([]alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum, len(o.Backends))
	for i, r := range o.Backends {
		sBackends[i] = alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum_value[string(r)])
	}
	p.SetBackends(sBackends)
	return p
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent object to its proto representation.
func GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentToProto(o *alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) *alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent{}
	p.SetTemplateLibrary(GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryToProto(o.TemplateLibrary))
	return p
}

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryToProto converts a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary object to its proto representation.
func GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryToProto(o *alpha.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary) *alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary{}
	p.SetInstallation(GkehubAlphaFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnumToProto(o.Installation))
	return p
}

// FeatureMembershipToProto converts a FeatureMembership resource to its proto representation.
func FeatureMembershipToProto(resource *alpha.FeatureMembership) *alphapb.GkehubAlphaFeatureMembership {
	p := &alphapb.GkehubAlphaFeatureMembership{}
	p.SetMesh(GkehubAlphaFeatureMembershipMeshToProto(resource.Mesh))
	p.SetConfigmanagement(GkehubAlphaFeatureMembershipConfigmanagementToProto(resource.Configmanagement))
	p.SetPolicycontroller(GkehubAlphaFeatureMembershipPolicycontrollerToProto(resource.Policycontroller))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetFeature(dcl.ValueOrEmptyString(resource.Feature))
	p.SetMembership(dcl.ValueOrEmptyString(resource.Membership))
	p.SetMembershipLocation(dcl.ValueOrEmptyString(resource.MembershipLocation))

	return p
}

// applyFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembership Apply() method.
func (s *FeatureMembershipServer) applyFeatureMembership(ctx context.Context, c *alpha.Client, request *alphapb.ApplyGkehubAlphaFeatureMembershipRequest) (*alphapb.GkehubAlphaFeatureMembership, error) {
	p := ProtoToFeatureMembership(request.GetResource())
	res, err := c.ApplyFeatureMembership(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FeatureMembershipToProto(res)
	return r, nil
}

// applyGkehubAlphaFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembership Apply() method.
func (s *FeatureMembershipServer) ApplyGkehubAlphaFeatureMembership(ctx context.Context, request *alphapb.ApplyGkehubAlphaFeatureMembershipRequest) (*alphapb.GkehubAlphaFeatureMembership, error) {
	cl, err := createConfigFeatureMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFeatureMembership(ctx, cl, request)
}

// DeleteFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembership Delete() method.
func (s *FeatureMembershipServer) DeleteGkehubAlphaFeatureMembership(ctx context.Context, request *alphapb.DeleteGkehubAlphaFeatureMembershipRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFeatureMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFeatureMembership(ctx, ProtoToFeatureMembership(request.GetResource()))

}

// ListGkehubAlphaFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembershipList() method.
func (s *FeatureMembershipServer) ListGkehubAlphaFeatureMembership(ctx context.Context, request *alphapb.ListGkehubAlphaFeatureMembershipRequest) (*alphapb.ListGkehubAlphaFeatureMembershipResponse, error) {
	cl, err := createConfigFeatureMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFeatureMembership(ctx, request.GetProject(), request.GetLocation(), request.GetFeature())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.GkehubAlphaFeatureMembership
	for _, r := range resources.Items {
		rp := FeatureMembershipToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListGkehubAlphaFeatureMembershipResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFeatureMembership(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
