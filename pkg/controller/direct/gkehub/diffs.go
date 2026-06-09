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
	"path/filepath"
	"reflect"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	featureapi "google.golang.org/api/gkehub/v1beta"
)

func diffFeatureMembership(report *structuredreporting.Diff, left *krm.GKEHubFeatureMembershipSpec, right *featureapi.MembershipFeatureSpec) {
	if left.Configmanagement != nil {
		if right.Configmanagement == nil {
			report.AddField("configmanagement", nil, left.Configmanagement)
		} else {
			diffConfigManagement(report, "configmanagement", left.Configmanagement, right.Configmanagement)
		}
	} else if right.Configmanagement != nil {
		report.AddField("configmanagement", right.Configmanagement, nil)
	}

	if left.Policycontroller != nil {
		if right.Policycontroller == nil {
			report.AddField("policycontroller", nil, left.Policycontroller)
		} else {
			diffPolicycontroller(report, "policycontroller", left.Policycontroller, right.Policycontroller)
		}
	} else if right.Policycontroller != nil {
		report.AddField("policycontroller", right.Policycontroller, nil)
	}

	if left.Mesh != nil {
		if right.Mesh == nil {
			report.AddField("mesh", nil, left.Mesh)
		} else {
			diffMesh(report, "mesh", left.Mesh, right.Mesh)
		}
	} else if right.Mesh != nil {
		report.AddField("mesh", right.Mesh, nil)
	}
}

func diffConfigManagement(report *structuredreporting.Diff, prefix string, left *krm.FeaturemembershipConfigmanagement, right *featureapi.ConfigManagementMembershipSpec) {
	if left.Binauthz != nil {
		if right.Binauthz == nil {
			report.AddField(prefix+".binauthz", nil, left.Binauthz)
		} else {
			diffBinauthz(report, prefix+".binauthz", left.Binauthz, right.Binauthz)
		}
	} else if right.Binauthz != nil {
		report.AddField(prefix+".binauthz", right.Binauthz, nil)
	}

	if left.ConfigSync != nil {
		if right.ConfigSync == nil {
			report.AddField(prefix+".configSync", nil, left.ConfigSync)
		} else {
			diffConfigSync(report, prefix+".configSync", left.ConfigSync, right.ConfigSync)
		}
	} else if right.ConfigSync != nil {
		report.AddField(prefix+".configSync", right.ConfigSync, nil)
	}

	if left.HierarchyController != nil {
		if right.HierarchyController == nil {
			report.AddField(prefix+".hierarchyController", nil, left.HierarchyController)
		} else {
			diffHierarchyController(report, prefix+".hierarchyController", left.HierarchyController, right.HierarchyController)
		}
	} else if right.HierarchyController != nil {
		report.AddField(prefix+".hierarchyController", right.HierarchyController, nil)
	}

	if left.PolicyController != nil {
		if right.PolicyController == nil {
			report.AddField(prefix+".policyController", nil, left.PolicyController)
		} else {
			diffPolicyController(report, prefix+".policyController", left.PolicyController, right.PolicyController)
		}
	} else if right.PolicyController != nil {
		report.AddField(prefix+".policyController", right.PolicyController, nil)
	}

	if left.Version != nil && direct.ValueOf(left.Version) != right.Version {
		report.AddField(prefix+".version", right.Version, direct.ValueOf(left.Version))
	}
	if left.Management != nil && direct.ValueOf(left.Management) != right.Management {
		report.AddField(prefix+".management", right.Management, direct.ValueOf(left.Management))
	}
}

func diffBinauthz(report *structuredreporting.Diff, prefix string, left *krm.FeaturemembershipBinauthz, right *featureapi.ConfigManagementBinauthzConfig) {
	if left.Enabled != nil && direct.ValueOf(left.Enabled) != right.Enabled {
		report.AddField(prefix+".enabled", right.Enabled, direct.ValueOf(left.Enabled))
	}
}

func diffConfigSync(report *structuredreporting.Diff, prefix string, left *krm.FeaturemembershipConfigSync, right *featureapi.ConfigManagementConfigSync) {
	if left.Git != nil {
		if right.Git == nil {
			report.AddField(prefix+".git", nil, left.Git)
		} else {
			diffGit(report, prefix+".git", left.Git, right.Git)
		}
	} else if right.Git != nil {
		report.AddField(prefix+".git", right.Git, nil)
	}

	if left.MetricsGcpServiceAccountRef != nil && left.MetricsGcpServiceAccountRef.External != "" && left.MetricsGcpServiceAccountRef.External != right.MetricsGcpServiceAccountEmail {
		report.AddField(prefix+".metricsGcpServiceAccountEmail", right.MetricsGcpServiceAccountEmail, left.MetricsGcpServiceAccountRef.External)
	}
	if left.Oci != nil {
		if right.Oci == nil {
			report.AddField(prefix+".oci", nil, left.Oci)
		} else {
			diffOci(report, prefix+".oci", left.Oci, right.Oci)
		}
	} else if right.Oci != nil {
		report.AddField(prefix+".oci", right.Oci, nil)
	}

	if left.PreventDrift != nil && direct.ValueOf(left.PreventDrift) != right.PreventDrift {
		report.AddField(prefix+".preventDrift", right.PreventDrift, direct.ValueOf(left.PreventDrift))
	}
	if left.StopSyncing != nil && direct.ValueOf(left.StopSyncing) != right.StopSyncing {
		report.AddField(prefix+".stopSyncing", right.StopSyncing, direct.ValueOf(left.StopSyncing))
	}
	if left.SourceFormat != nil && !strings.EqualFold(direct.ValueOf(left.SourceFormat), right.SourceFormat) {
		report.AddField(prefix+".sourceFormat", right.SourceFormat, direct.ValueOf(left.SourceFormat))
	}
}

func diffGit(report *structuredreporting.Diff, prefix string, left *krm.FeaturemembershipGit, right *featureapi.ConfigManagementGitConfig) {
	if left.GcpServiceAccountRef != nil && left.GcpServiceAccountRef.External != "" && left.GcpServiceAccountRef.External != right.GcpServiceAccountEmail {
		report.AddField(prefix+".gcpServiceAccountEmail", right.GcpServiceAccountEmail, left.GcpServiceAccountRef.External)
	}
	if left.HttpsProxy != nil && direct.ValueOf(left.HttpsProxy) != right.HttpsProxy {
		report.AddField(prefix+".httpsProxy", right.HttpsProxy, direct.ValueOf(left.HttpsProxy))
	}
	if left.PolicyDir != nil && filepath.Clean(direct.ValueOf(left.PolicyDir)) != filepath.Clean(right.PolicyDir) {
		report.AddField(prefix+".policyDir", right.PolicyDir, direct.ValueOf(left.PolicyDir))
	}
	if left.SecretType != nil && !strings.EqualFold(direct.ValueOf(left.SecretType), right.SecretType) {
		report.AddField(prefix+".secretType", right.SecretType, direct.ValueOf(left.SecretType))
	}
	if left.SyncBranch != nil && direct.ValueOf(left.SyncBranch) != "" && direct.ValueOf(left.SyncBranch) != right.SyncBranch {
		report.AddField(prefix+".syncBranch", right.SyncBranch, direct.ValueOf(left.SyncBranch))
	}
	if left.SyncRepo != nil && direct.ValueOf(left.SyncRepo) != right.SyncRepo {
		report.AddField(prefix+".syncRepo", right.SyncRepo, direct.ValueOf(left.SyncRepo))
	}
	if left.SyncRev != nil && direct.ValueOf(left.SyncRev) != right.SyncRev {
		report.AddField(prefix+".syncRev", right.SyncRev, direct.ValueOf(left.SyncRev))
	}
	if left.SyncWaitSecs != nil && direct.ValueOf(left.SyncWaitSecs) != fmt.Sprintf("%d", right.SyncWaitSecs) {
		report.AddField(prefix+".syncWaitSecs", fmt.Sprintf("%d", right.SyncWaitSecs), direct.ValueOf(left.SyncWaitSecs))
	}
}

func diffOci(report *structuredreporting.Diff, prefix string, left *krm.FeaturemembershipOci, right *featureapi.ConfigManagementOciConfig) {
	if left.GcpServiceAccountRef != nil && left.GcpServiceAccountRef.External != "" && left.GcpServiceAccountRef.External != right.GcpServiceAccountEmail {
		report.AddField(prefix+".gcpServiceAccountEmail", right.GcpServiceAccountEmail, left.GcpServiceAccountRef.External)
	}
	if left.PolicyDir != nil && filepath.Clean(direct.ValueOf(left.PolicyDir)) != filepath.Clean(right.PolicyDir) {
		report.AddField(prefix+".policyDir", right.PolicyDir, direct.ValueOf(left.PolicyDir))
	}
	if left.SecretType != nil && !strings.EqualFold(direct.ValueOf(left.SecretType), right.SecretType) {
		report.AddField(prefix+".secretType", right.SecretType, direct.ValueOf(left.SecretType))
	}
	if left.SyncRepo != nil && direct.ValueOf(left.SyncRepo) != right.SyncRepo {
		report.AddField(prefix+".syncRepo", right.SyncRepo, direct.ValueOf(left.SyncRepo))
	}
	if left.SyncWaitSecs != nil && direct.ValueOf(left.SyncWaitSecs) != fmt.Sprintf("%d", right.SyncWaitSecs) {
		report.AddField(prefix+".syncWaitSecs", fmt.Sprintf("%d", right.SyncWaitSecs), direct.ValueOf(left.SyncWaitSecs))
	}
}

func diffHierarchyController(report *structuredreporting.Diff, prefix string, left *krm.FeaturemembershipHierarchyController, right *featureapi.ConfigManagementHierarchyControllerConfig) {
	if left.EnableHierarchicalResourceQuota != nil && direct.ValueOf(left.EnableHierarchicalResourceQuota) != right.EnableHierarchicalResourceQuota {
		report.AddField(prefix+".enableHierarchicalResourceQuota", right.EnableHierarchicalResourceQuota, direct.ValueOf(left.EnableHierarchicalResourceQuota))
	}
	if left.EnablePodTreeLabels != nil && direct.ValueOf(left.EnablePodTreeLabels) != right.EnablePodTreeLabels {
		report.AddField(prefix+".enablePodTreeLabels", right.EnablePodTreeLabels, direct.ValueOf(left.EnablePodTreeLabels))
	}
	if left.Enabled != nil && direct.ValueOf(left.Enabled) != right.Enabled {
		report.AddField(prefix+".enabled", right.Enabled, direct.ValueOf(left.Enabled))
	}
}

func diffPolicyController(report *structuredreporting.Diff, prefix string, left *krm.FeaturemembershipPolicyController, right *featureapi.ConfigManagementPolicyController) {
	if left.AuditIntervalSeconds != nil && direct.ValueOf(left.AuditIntervalSeconds) != fmt.Sprintf("%d", right.AuditIntervalSeconds) {
		report.AddField(prefix+".auditIntervalSeconds", fmt.Sprintf("%d", right.AuditIntervalSeconds), direct.ValueOf(left.AuditIntervalSeconds))
	}
	if left.Enabled != nil && direct.ValueOf(left.Enabled) != right.Enabled {
		report.AddField(prefix+".enabled", right.Enabled, direct.ValueOf(left.Enabled))
	}
	if left.ExemptableNamespaces != nil {
		if (len(left.ExemptableNamespaces) != 0 || len(right.ExemptableNamespaces) != 0) && !reflect.DeepEqual(left.ExemptableNamespaces, right.ExemptableNamespaces) {
			report.AddField(prefix+".exemptableNamespaces", right.ExemptableNamespaces, left.ExemptableNamespaces)
		}
	}
	if left.LogDeniesEnabled != nil && direct.ValueOf(left.LogDeniesEnabled) != right.LogDeniesEnabled {
		report.AddField(prefix+".logDeniesEnabled", right.LogDeniesEnabled, direct.ValueOf(left.LogDeniesEnabled))
	}
	if left.Monitoring != nil {
		if right.Monitoring == nil {
			report.AddField(prefix+".monitoring", nil, left.Monitoring)
		} else {
			diffMonitoring(report, prefix+".monitoring", left.Monitoring, right.Monitoring)
		}
	} else if right.Monitoring != nil {
		report.AddField(prefix+".monitoring", right.Monitoring, nil)
	}
	if left.MutationEnabled != nil && direct.ValueOf(left.MutationEnabled) != right.MutationEnabled {
		report.AddField(prefix+".mutationEnabled", right.MutationEnabled, direct.ValueOf(left.MutationEnabled))
	}
	if left.ReferentialRulesEnabled != nil && direct.ValueOf(left.ReferentialRulesEnabled) != right.ReferentialRulesEnabled {
		report.AddField(prefix+".referentialRulesEnabled", right.ReferentialRulesEnabled, direct.ValueOf(left.ReferentialRulesEnabled))
	}
	if left.TemplateLibraryInstalled != nil && direct.ValueOf(left.TemplateLibraryInstalled) != right.TemplateLibraryInstalled {
		report.AddField(prefix+".templateLibraryInstalled", right.TemplateLibraryInstalled, direct.ValueOf(left.TemplateLibraryInstalled))
	}
}

func diffMonitoring(report *structuredreporting.Diff, prefix string, left *krm.FeaturemembershipMonitoring, right *featureapi.ConfigManagementPolicyControllerMonitoring) {
	if left.Backends != nil {
		if (len(left.Backends) != 0 || len(right.Backends) != 0) && !reflect.DeepEqual(left.Backends, right.Backends) {
			report.AddField(prefix+".backends", right.Backends, left.Backends)
		}
	}
}

func diffMesh(report *structuredreporting.Diff, prefix string, left *krm.FeaturemembershipMesh, right *featureapi.ServiceMeshMembershipSpec) {
	if left.ControlPlane != nil && !strings.EqualFold(direct.ValueOf(left.ControlPlane), right.ControlPlane) {
		report.AddField(prefix+".controlPlane", right.ControlPlane, direct.ValueOf(left.ControlPlane))
	}
	if left.Management != nil && direct.ValueOf(left.Management) != right.Management {
		report.AddField(prefix+".management", right.Management, direct.ValueOf(left.Management))
	}
}

func diffPolicycontroller(report *structuredreporting.Diff, prefix string, left *krm.FeaturemembershipPolicycontroller, right *featureapi.PolicyControllerMembershipSpec) {
	if left.PolicyControllerHubConfig != nil {
		if right.PolicyControllerHubConfig == nil {
			report.AddField(prefix+".policyControllerHubConfig", nil, left.PolicyControllerHubConfig)
		} else {
			actual := FeaturemembershipPolicyControllerHubConfig_FromProto(&direct.MapContext{}, right.PolicyControllerHubConfig)
			if !reflect.DeepEqual(left.PolicyControllerHubConfig, actual) {
				report.AddField(prefix+".policyControllerHubConfig", actual, left.PolicyControllerHubConfig)
			}
		}
	}
	if left.Version != nil && direct.ValueOf(left.Version) != right.Version {
		report.AddField(prefix+".version", right.Version, direct.ValueOf(left.Version))
	}
}
