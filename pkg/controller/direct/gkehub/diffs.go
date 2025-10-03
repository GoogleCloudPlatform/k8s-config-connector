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
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	featureapi "google.golang.org/api/gkehub/v1beta"
)

func diffFeatureMembership(left *krm.GKEHubFeatureMembershipSpec, right *featureapi.MembershipFeatureSpec) []string {
	var diffs []string
	if left.Configmanagement != nil {
		if right.Configmanagement == nil {
			diffs = append(diffs, "configmanagement")
		} else {
			diffs = append(diffs, diffConfigManagement(left.Configmanagement, right.Configmanagement)...)
		}
	}
	if left.Policycontroller != nil {
		if right.Policycontroller == nil {
			diffs = append(diffs, "policycontroller")
		} else {
			diffs = append(diffs, diffPolicycontroller(left.Policycontroller, right.Policycontroller)...)
		}
	}
	if left.Mesh != nil {
		if right.Mesh == nil {
			diffs = append(diffs, "mesh")
		} else {
			diffs = append(diffs, diffMesh(left.Mesh, right.Mesh)...)
		}
	}
	return diffs
}

func diffConfigManagement(left *krm.FeaturemembershipConfigmanagement, right *featureapi.ConfigManagementMembershipSpec) []string {
	var diffs []string
	if left.Binauthz != nil {
		if right.Binauthz == nil {
			diffs = append(diffs, "binauthz")
		} else {
			diffs = append(diffs, diffBinauthz(left.Binauthz, right.Binauthz)...)
		}
	}
	if left.ConfigSync != nil {
		if right.ConfigSync == nil {
			diffs = append(diffs, "configsync")
		} else {
			diffs = append(diffs, diffConfigSync(left.ConfigSync, right.ConfigSync)...)
		}
	}
	if left.HierarchyController != nil {
		if right.HierarchyController == nil {
			diffs = append(diffs, "hierachycontroller")
		} else {
			diffs = append(diffs, diffHierarchyController(left.HierarchyController, right.HierarchyController)...)
		}
	}
	if left.PolicyController != nil {
		if right.PolicyController == nil {
			diffs = append(diffs, "policyController")
		} else {
			diffs = append(diffs, diffPolicyController(left.PolicyController, right.PolicyController)...)
		}
	}
	if left.Version != nil && !reflect.DeepEqual(left.Version, right.Version) {
		diffs = append(diffs, "version")
	}
	if left.Management != nil && !reflect.DeepEqual(left.Management, right.Management) {
		diffs = append(diffs, "management")
	}
	return diffs
}

func diffBinauthz(left *krm.FeaturemembershipBinauthz, right *featureapi.ConfigManagementBinauthzConfig) []string {
	var diffs []string
	if left.Enabled != nil && !reflect.DeepEqual(left.Enabled, right.Enabled) {
		diffs = append(diffs, "enabled")
	}
	return diffs
}

func diffConfigSync(left *krm.FeaturemembershipConfigSync, right *featureapi.ConfigManagementConfigSync) []string {
	var diffs []string
	if left.Git != nil {
		if right.Git == nil {
			diffs = append(diffs, "git")
		} else {
			diffs = append(diffs, diffGit(left.Git, right.Git)...)
		}
	}
	if left.MetricsGcpServiceAccountRef != nil && !reflect.DeepEqual(left.MetricsGcpServiceAccountRef.External, right.MetricsGcpServiceAccountEmail) {
		diffs = append(diffs, "metricsGcpServiceAccountEmail")
	}
	if left.Oci != nil {
		if right.Oci == nil {
			diffs = append(diffs, "oci")
		} else {
			diffs = append(diffs, diffOci(left.Oci, right.Oci)...)
		}
	}
	if left.PreventDrift != nil && !reflect.DeepEqual(left.PreventDrift, right.PreventDrift) {
		diffs = append(diffs, "preventDrift")
	}
	if left.StopSyncing != nil && !reflect.DeepEqual(left.StopSyncing, right.StopSyncing) {
		diffs = append(diffs, "stopSyncing")
	}
	if left.SourceFormat != nil && !reflect.DeepEqual(left.SourceFormat, right.SourceFormat) {
		diffs = append(diffs, "sourceFormat")
	}
	return diffs
}

func diffGit(left *krm.FeaturemembershipGit, right *featureapi.ConfigManagementGitConfig) []string {
	var diffs []string
	if left.GcpServiceAccountRef != nil && !reflect.DeepEqual(left.GcpServiceAccountRef.External, right.GcpServiceAccountEmail) {
		diffs = append(diffs, "gcpServiceAccountEmail")
	}
	if left.HttpsProxy != nil && !reflect.DeepEqual(left.HttpsProxy, right.HttpsProxy) {
		diffs = append(diffs, "httpsProxy")
	}
	if left.PolicyDir != nil && !reflect.DeepEqual(left.PolicyDir, right.PolicyDir) {
		diffs = append(diffs, "policyDir")
	}
	if left.SecretType != nil && !reflect.DeepEqual(left.SecretType, right.SecretType) {
		diffs = append(diffs, "secretType")
	}
	if left.SyncBranch != nil && !reflect.DeepEqual(left.SyncBranch, right.SyncBranch) {
		diffs = append(diffs, "syncBranch")
	}
	if left.SyncRepo != nil && !reflect.DeepEqual(left.SyncRepo, right.SyncRepo) {
		diffs = append(diffs, "syncRepo")
	}
	if left.SyncRev != nil && !reflect.DeepEqual(left.SyncRev, right.SyncRev) {
		diffs = append(diffs, "syncRev")
	}
	if left.SyncWaitSecs != nil && !reflect.DeepEqual(left.SyncWaitSecs, fmt.Sprintf("%d", right.SyncWaitSecs)) {
		diffs = append(diffs, "syncWaitSecss")
	}
	return diffs
}

func diffOci(left *krm.FeaturemembershipOci, right *featureapi.ConfigManagementOciConfig) []string {
	var diffs []string
	if left.GcpServiceAccountRef != nil && !reflect.DeepEqual(left.GcpServiceAccountRef.External, right.GcpServiceAccountEmail) {
		diffs = append(diffs, "gcpServiceAccountEmail")
	}
	if left.PolicyDir != nil && !reflect.DeepEqual(left.PolicyDir, right.PolicyDir) {
		diffs = append(diffs, "policyDir")
	}
	if left.SecretType != nil && !reflect.DeepEqual(left.SecretType, right.SecretType) {
		diffs = append(diffs, "secretType")
	}
	if left.SyncRepo != nil && !reflect.DeepEqual(left.SyncRepo, right.SyncRepo) {
		diffs = append(diffs, "syncRepo")
	}
	if left.SyncWaitSecs != nil && !reflect.DeepEqual(left.SyncWaitSecs, fmt.Sprintf("%d", right.SyncWaitSecs)) {
		diffs = append(diffs, "syncWaitSecs")
	}
	return diffs
}

func diffHierarchyController(left *krm.FeaturemembershipHierarchyController, right *featureapi.ConfigManagementHierarchyControllerConfig) []string {
	var diffs []string
	if left.EnableHierarchicalResourceQuota != nil && !reflect.DeepEqual(left.EnableHierarchicalResourceQuota, right.EnableHierarchicalResourceQuota) {
		diffs = append(diffs, "enableHierarchicalResourceQuota")
	}
	if left.EnablePodTreeLabels != nil && !reflect.DeepEqual(left.EnablePodTreeLabels, right.EnablePodTreeLabels) {
		diffs = append(diffs, "enablePodTreeLabels")
	}
	if left.Enabled != nil && !reflect.DeepEqual(left.Enabled, right.Enabled) {
		diffs = append(diffs, "enabled")
	}
	return diffs
}

func diffPolicyController(left *krm.FeaturemembershipPolicyController, right *featureapi.ConfigManagementPolicyController) []string {
	var diffs []string
	if left.AuditIntervalSeconds != nil && !reflect.DeepEqual(left.AuditIntervalSeconds, fmt.Sprintf("%d", right.AuditIntervalSeconds)) {
		diffs = append(diffs, "auditIntervalSeconds")
	}
	if left.Enabled != nil && !reflect.DeepEqual(left.Enabled, right.Enabled) {
		diffs = append(diffs, "enabled")
	}
	if left.ExemptableNamespaces != nil && !reflect.DeepEqual(left.ExemptableNamespaces, right.ExemptableNamespaces) {
		diffs = append(diffs, "exemptableNamespaces")
	}
	if left.LogDeniesEnabled != nil && !reflect.DeepEqual(left.LogDeniesEnabled, right.LogDeniesEnabled) {
		diffs = append(diffs, "logDeniesEnabled")
	}
	if left.Monitoring != nil {
		if right.Monitoring == nil {
			diffs = append(diffs, "monitoring")
		} else {
			diffs = append(diffs, diffMonitoring(left.Monitoring, right.Monitoring)...)
		}
	}
	if left.MutationEnabled != nil && !reflect.DeepEqual(left.MutationEnabled, right.MutationEnabled) {
		diffs = append(diffs, "mutationEnabled")
	}
	if left.ReferentialRulesEnabled != nil && !reflect.DeepEqual(left.ReferentialRulesEnabled, right.ReferentialRulesEnabled) {
		diffs = append(diffs, "referentialRulesEnabled")
	}
	if left.TemplateLibraryInstalled != nil && !reflect.DeepEqual(left.TemplateLibraryInstalled, right.TemplateLibraryInstalled) {
		diffs = append(diffs, "templateLibraryInstalled")
	}
	return diffs
}

func diffMonitoring(left *krm.FeaturemembershipMonitoring, right *featureapi.ConfigManagementPolicyControllerMonitoring) []string {
	var diffs []string
	if left.Backends != nil && !reflect.DeepEqual(left.Backends, right.Backends) {
		diffs = append(diffs, "backends")
	}
	return diffs
}

func diffMesh(left *krm.FeaturemembershipMesh, right *featureapi.ServiceMeshMembershipSpec) []string {
	var diffs []string
	if left.ControlPlane != nil && !reflect.DeepEqual(left.ControlPlane, right.ControlPlane) {
		diffs = append(diffs, "controlPlane")
	}
	if left.Management != nil && !reflect.DeepEqual(left.Management, right.Management) {
		diffs = append(diffs, "management")
	}
	return diffs
}

func diffPolicycontroller(left *krm.FeaturemembershipPolicycontroller, right *featureapi.PolicyControllerMembershipSpec) []string {
	var diffs []string
	if right.PolicyControllerHubConfig == nil {
		diffs = append(diffs, "policyControllerHubConfig")
	} else if !reflect.DeepEqual(left.PolicyControllerHubConfig, right.PolicyControllerHubConfig) {
		diffs = append(diffs, "policyControllerHubConfig")
	}
	if left.Version != nil && !reflect.DeepEqual(left.Version, right.Version) {
		diffs = append(diffs, "version")
	}
	return diffs
}
