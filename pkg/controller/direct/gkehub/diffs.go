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
	}
	if left.Policycontroller != nil {
		if right.Policycontroller == nil {
			report.AddField("policycontroller", nil, left.Policycontroller)
		} else {
			diffPolicycontroller(report, "policycontroller", left.Policycontroller, right.Policycontroller)
		}
	}
	if left.Mesh != nil {
		if right.Mesh == nil {
			report.AddField("mesh", nil, left.Mesh)
		} else {
			diffMesh(report, "mesh", left.Mesh, right.Mesh)
		}
	}
}

func diffConfigManagement(report *structuredreporting.Diff, path string, left *krm.FeaturemembershipConfigmanagement, right *featureapi.ConfigManagementMembershipSpec) {
	if left.Binauthz != nil {
		if right.Binauthz == nil {
			report.AddField(path+".binauthz", nil, left.Binauthz)
		} else {
			diffBinauthz(report, path+".binauthz", left.Binauthz, right.Binauthz)
		}
	}
	if left.ConfigSync != nil {
		if right.ConfigSync == nil {
			report.AddField(path+".configsync", nil, left.ConfigSync)
		} else {
			diffConfigSync(report, path+".configsync", left.ConfigSync, right.ConfigSync)
		}
	}
	if left.HierarchyController != nil {
		if right.HierarchyController == nil {
			report.AddField(path+".hierarchycontroller", nil, left.HierarchyController)
		} else {
			diffHierarchyController(report, path+".hierarchycontroller", left.HierarchyController, right.HierarchyController)
		}
	}
	if left.PolicyController != nil {
		if right.PolicyController == nil {
			report.AddField(path+".policyController", nil, left.PolicyController)
		} else {
			diffPolicyController(report, path+".policyController", left.PolicyController, right.PolicyController)
		}
	}
	if left.Version != nil && !reflect.DeepEqual(*left.Version, right.Version) {
		report.AddField(path+".version", right.Version, *left.Version)
	}
	if left.Management != nil && !reflect.DeepEqual(*left.Management, right.Management) {
		report.AddField(path+".management", right.Management, *left.Management)
	}
}

func diffBinauthz(report *structuredreporting.Diff, path string, left *krm.FeaturemembershipBinauthz, right *featureapi.ConfigManagementBinauthzConfig) {
	if left.Enabled != nil && !reflect.DeepEqual(*left.Enabled, right.Enabled) {
		report.AddField(path+".enabled", right.Enabled, *left.Enabled)
	}
}

func diffConfigSync(report *structuredreporting.Diff, path string, left *krm.FeaturemembershipConfigSync, right *featureapi.ConfigManagementConfigSync) {
	if left.Git != nil {
		if right.Git == nil {
			report.AddField(path+".git", nil, left.Git)
		} else {
			diffGit(report, path+".git", left.Git, right.Git)
		}
	}
	if left.MetricsGcpServiceAccountRef != nil && !reflect.DeepEqual(left.MetricsGcpServiceAccountRef.External, right.MetricsGcpServiceAccountEmail) {
		report.AddField(path+".metricsGcpServiceAccountEmail", right.MetricsGcpServiceAccountEmail, left.MetricsGcpServiceAccountRef.External)
	}
	if left.Oci != nil {
		if right.Oci == nil {
			report.AddField(path+".oci", nil, left.Oci)
		} else {
			diffOci(report, path+".oci", left.Oci, right.Oci)
		}
	}
	if left.PreventDrift != nil && !reflect.DeepEqual(*left.PreventDrift, right.PreventDrift) {
		report.AddField(path+".preventDrift", right.PreventDrift, *left.PreventDrift)
	}
	if left.StopSyncing != nil && !reflect.DeepEqual(*left.StopSyncing, right.StopSyncing) {
		report.AddField(path+".stopSyncing", right.StopSyncing, *left.StopSyncing)
	}
	if left.SourceFormat != nil && !reflect.DeepEqual(*left.SourceFormat, right.SourceFormat) {
		report.AddField(path+".sourceFormat", right.SourceFormat, *left.SourceFormat)
	}
}

func diffGit(report *structuredreporting.Diff, path string, left *krm.FeaturemembershipGit, right *featureapi.ConfigManagementGitConfig) {
	if left.GcpServiceAccountRef != nil && !reflect.DeepEqual(left.GcpServiceAccountRef.External, right.GcpServiceAccountEmail) {
		report.AddField(path+".gcpServiceAccountEmail", right.GcpServiceAccountEmail, left.GcpServiceAccountRef.External)
	}
	if left.HttpsProxy != nil && !reflect.DeepEqual(*left.HttpsProxy, right.HttpsProxy) {
		report.AddField(path+".httpsProxy", right.HttpsProxy, *left.HttpsProxy)
	}
	if left.PolicyDir != nil && !reflect.DeepEqual(*left.PolicyDir, right.PolicyDir) {
		report.AddField(path+".policyDir", right.PolicyDir, *left.PolicyDir)
	}
	if left.SecretType != nil && !reflect.DeepEqual(*left.SecretType, right.SecretType) {
		report.AddField(path+".secretType", right.SecretType, *left.SecretType)
	}
	if left.SyncBranch != nil && !reflect.DeepEqual(*left.SyncBranch, right.SyncBranch) {
		report.AddField(path+".syncBranch", right.SyncBranch, *left.SyncBranch)
	}
	if left.SyncRepo != nil && !reflect.DeepEqual(*left.SyncRepo, right.SyncRepo) {
		report.AddField(path+".syncRepo", right.SyncRepo, *left.SyncRepo)
	}
	if left.SyncRev != nil && !reflect.DeepEqual(*left.SyncRev, right.SyncRev) {
		report.AddField(path+".syncRev", right.SyncRev, *left.SyncRev)
	}
	if left.SyncWaitSecs != nil && !reflect.DeepEqual(*left.SyncWaitSecs, fmt.Sprintf("%d", right.SyncWaitSecs)) {
		report.AddField(path+".syncWaitSecs", right.SyncWaitSecs, *left.SyncWaitSecs)
	}
}

func diffOci(report *structuredreporting.Diff, path string, left *krm.FeaturemembershipOci, right *featureapi.ConfigManagementOciConfig) {
	if left.GcpServiceAccountRef != nil && !reflect.DeepEqual(left.GcpServiceAccountRef.External, right.GcpServiceAccountEmail) {
		report.AddField(path+".gcpServiceAccountEmail", right.GcpServiceAccountEmail, left.GcpServiceAccountRef.External)
	}
	if left.PolicyDir != nil && !reflect.DeepEqual(*left.PolicyDir, right.PolicyDir) {
		report.AddField(path+".policyDir", right.PolicyDir, *left.PolicyDir)
	}
	if left.SecretType != nil && !reflect.DeepEqual(*left.SecretType, right.SecretType) {
		report.AddField(path+".secretType", right.SecretType, *left.SecretType)
	}
	if left.SyncRepo != nil && !reflect.DeepEqual(*left.SyncRepo, right.SyncRepo) {
		report.AddField(path+".syncRepo", right.SyncRepo, *left.SyncRepo)
	}
	if left.SyncWaitSecs != nil && !reflect.DeepEqual(*left.SyncWaitSecs, fmt.Sprintf("%d", right.SyncWaitSecs)) {
		report.AddField(path+".syncWaitSecs", right.SyncWaitSecs, *left.SyncWaitSecs)
	}
}

func diffHierarchyController(report *structuredreporting.Diff, path string, left *krm.FeaturemembershipHierarchyController, right *featureapi.ConfigManagementHierarchyControllerConfig) {
	if left.EnableHierarchicalResourceQuota != nil && !reflect.DeepEqual(*left.EnableHierarchicalResourceQuota, right.EnableHierarchicalResourceQuota) {
		report.AddField(path+".enableHierarchicalResourceQuota", right.EnableHierarchicalResourceQuota, *left.EnableHierarchicalResourceQuota)
	}
	if left.EnablePodTreeLabels != nil && !reflect.DeepEqual(*left.EnablePodTreeLabels, right.EnablePodTreeLabels) {
		report.AddField(path+".enablePodTreeLabels", right.EnablePodTreeLabels, *left.EnablePodTreeLabels)
	}
	if left.Enabled != nil && !reflect.DeepEqual(*left.Enabled, right.Enabled) {
		report.AddField(path+".enabled", right.Enabled, *left.Enabled)
	}
}

func diffPolicyController(report *structuredreporting.Diff, path string, left *krm.FeaturemembershipPolicyController, right *featureapi.ConfigManagementPolicyController) {
	if left.AuditIntervalSeconds != nil && !reflect.DeepEqual(*left.AuditIntervalSeconds, fmt.Sprintf("%d", right.AuditIntervalSeconds)) {
		report.AddField(path+".auditIntervalSeconds", right.AuditIntervalSeconds, *left.AuditIntervalSeconds)
	}
	if left.Enabled != nil && !reflect.DeepEqual(*left.Enabled, right.Enabled) {
		report.AddField(path+".enabled", right.Enabled, *left.Enabled)
	}
	if left.ExemptableNamespaces != nil && !reflect.DeepEqual(left.ExemptableNamespaces, right.ExemptableNamespaces) {
		report.AddField(path+".exemptableNamespaces", right.ExemptableNamespaces, left.ExemptableNamespaces)
	}
	if left.LogDeniesEnabled != nil && !reflect.DeepEqual(*left.LogDeniesEnabled, right.LogDeniesEnabled) {
		report.AddField(path+".logDeniesEnabled", right.LogDeniesEnabled, *left.LogDeniesEnabled)
	}
	if left.Monitoring != nil {
		if right.Monitoring == nil {
			report.AddField(path+".monitoring", nil, left.Monitoring)
		} else {
			diffMonitoring(report, path+".monitoring", left.Monitoring, right.Monitoring)
		}
	}
	if left.MutationEnabled != nil && !reflect.DeepEqual(*left.MutationEnabled, right.MutationEnabled) {
		report.AddField(path+".mutationEnabled", right.MutationEnabled, *left.MutationEnabled)
	}
	if left.ReferentialRulesEnabled != nil && !reflect.DeepEqual(*left.ReferentialRulesEnabled, right.ReferentialRulesEnabled) {
		report.AddField(path+".referentialRulesEnabled", right.ReferentialRulesEnabled, *left.ReferentialRulesEnabled)
	}
	if left.TemplateLibraryInstalled != nil && !reflect.DeepEqual(*left.TemplateLibraryInstalled, right.TemplateLibraryInstalled) {
		report.AddField(path+".templateLibraryInstalled", right.TemplateLibraryInstalled, *left.TemplateLibraryInstalled)
	}
}

func diffMonitoring(report *structuredreporting.Diff, path string, left *krm.FeaturemembershipMonitoring, right *featureapi.ConfigManagementPolicyControllerMonitoring) {
	if left.Backends != nil && !reflect.DeepEqual(left.Backends, right.Backends) {
		report.AddField(path+".backends", right.Backends, left.Backends)
	}
}

func diffMesh(report *structuredreporting.Diff, path string, left *krm.FeaturemembershipMesh, right *featureapi.ServiceMeshMembershipSpec) {
	if left.ControlPlane != nil && !reflect.DeepEqual(*left.ControlPlane, right.ControlPlane) {
		report.AddField(path+".controlPlane", right.ControlPlane, *left.ControlPlane)
	}
	if left.Management != nil && !reflect.DeepEqual(*left.Management, right.Management) {
		report.AddField(path+".management", right.Management, *left.Management)
	}
}

func diffPolicycontroller(report *structuredreporting.Diff, path string, left *krm.FeaturemembershipPolicycontroller, right *featureapi.PolicyControllerMembershipSpec) {
	if right.PolicyControllerHubConfig == nil {
		report.AddField(path+".policyControllerHubConfig", nil, left.PolicyControllerHubConfig)
	} else {
		diffPolicyControllerHubConfig(report, path+".policyControllerHubConfig", left.PolicyControllerHubConfig, right.PolicyControllerHubConfig)
	}
	if left.Version != nil && !reflect.DeepEqual(*left.Version, right.Version) {
		report.AddField(path+".version", right.Version, *left.Version)
	}
}

func diffPolicyControllerHubConfig(report *structuredreporting.Diff, path string, left *krm.HubConfig, right *featureapi.PolicyControllerHubConfig) {
	if left == nil && right == nil {
		return
	}
	if left == nil {
		report.AddField(path, right, nil)
		return
	}
	if right == nil {
		report.AddField(path, nil, left)
		return
	}

	if left.AuditIntervalSeconds != nil && *left.AuditIntervalSeconds != right.AuditIntervalSeconds {
		report.AddField(path+".auditIntervalSeconds", right.AuditIntervalSeconds, *left.AuditIntervalSeconds)
	}
	if left.ConstraintViolationLimit != nil && *left.ConstraintViolationLimit != right.ConstraintViolationLimit {
		report.AddField(path+".constraintViolationLimit", right.ConstraintViolationLimit, *left.ConstraintViolationLimit)
	}
	if left.ExemptableNamespaces != nil && !reflect.DeepEqual(left.ExemptableNamespaces, right.ExemptableNamespaces) {
		report.AddField(path+".exemptableNamespaces", right.ExemptableNamespaces, left.ExemptableNamespaces)
	}
	if left.InstallSpec != nil && *left.InstallSpec != right.InstallSpec {
		report.AddField(path+".installSpec", right.InstallSpec, *left.InstallSpec)
	}
	if left.LogDeniesEnabled != nil && *left.LogDeniesEnabled != right.LogDeniesEnabled {
		report.AddField(path+".logDeniesEnabled", right.LogDeniesEnabled, *left.LogDeniesEnabled)
	}
	if left.Monitoring != nil {
		if right.Monitoring == nil {
			report.AddField(path+".monitoring", nil, left.Monitoring)
		} else {
			if !reflect.DeepEqual(left.Monitoring.Backends, right.Monitoring.Backends) {
				report.AddField(path+".monitoring.backends", right.Monitoring.Backends, left.Monitoring.Backends)
			}
		}
	}
	if left.MutationEnabled != nil && *left.MutationEnabled != right.MutationEnabled {
		report.AddField(path+".mutationEnabled", right.MutationEnabled, *left.MutationEnabled)
	}
	if left.ReferentialRulesEnabled != nil && *left.ReferentialRulesEnabled != right.ReferentialRulesEnabled {
		report.AddField(path+".referentialRulesEnabled", right.ReferentialRulesEnabled, *left.ReferentialRulesEnabled)
	}
	if left.PolicyContent != nil {
		if right.PolicyContent == nil {
			report.AddField(path+".policyContent", nil, left.PolicyContent)
		} else {
			diffPolicyContent(report, path+".policyContent", left.PolicyContent, right.PolicyContent)
		}
	}
	if left.DeploymentConfigs != nil {
		if right.DeploymentConfigs == nil {
			report.AddField(path+".deploymentConfigs", nil, left.DeploymentConfigs)
		} else {
			diffDeploymentConfigs(report, path+".deploymentConfigs", left.DeploymentConfigs, right.DeploymentConfigs)
		}
	}
}

func diffPolicyContent(report *structuredreporting.Diff, path string, left *krm.FeaturemembershipPolicyContent, right *featureapi.PolicyControllerPolicyContentSpec) {
	if left == nil && right == nil {
		return
	}
	if left == nil {
		report.AddField(path, right, nil)
		return
	}
	if right == nil {
		report.AddField(path, nil, left)
		return
	}

	if left.TemplateLibrary != nil {
		if right.TemplateLibrary == nil {
			report.AddField(path+".templateLibrary", nil, left.TemplateLibrary)
		} else {
			if left.TemplateLibrary.Installation != nil && *left.TemplateLibrary.Installation != right.TemplateLibrary.Installation {
				report.AddField(path+".templateLibrary.installation", right.TemplateLibrary.Installation, *left.TemplateLibrary.Installation)
			}
		}
	}
}

func diffDeploymentConfigs(report *structuredreporting.Diff, path string, left *krm.PolicyControllerDeploymentConfigs, right map[string]featureapi.PolicyControllerPolicyControllerDeploymentConfig) {
	if left == nil && right == nil {
		return
	}
	if left == nil {
		report.AddField(path, right, nil)
		return
	}
	if right == nil {
		report.AddField(path, nil, left)
		return
	}

	if left.Admission != nil {
		if val, ok := right["admission"]; !ok {
			report.AddField(path+".admission", nil, left.Admission)
		} else {
			diffDeploymentConfig(report, path+".admission", left.Admission, &val)
		}
	}
	if left.Audit != nil {
		if val, ok := right["audit"]; !ok {
			report.AddField(path+".audit", nil, left.Audit)
		} else {
			diffDeploymentConfig(report, path+".audit", left.Audit, &val)
		}
	}
	if left.Mutation != nil {
		if val, ok := right["mutation"]; !ok {
			report.AddField(path+".mutation", nil, left.Mutation)
		} else {
			diffDeploymentConfig(report, path+".mutation", left.Mutation, &val)
		}
	}
}

func diffDeploymentConfig(report *structuredreporting.Diff, path string, left *krm.PolicyControllerDeploymentConfig, right *featureapi.PolicyControllerPolicyControllerDeploymentConfig) {
	if left == nil && right == nil {
		return
	}
	if left == nil {
		report.AddField(path, right, nil)
		return
	}
	if right == nil {
		report.AddField(path, nil, left)
		return
	}

	if left.ReplicaCount != nil && *left.ReplicaCount != right.ReplicaCount {
		report.AddField(path+".replicaCount", right.ReplicaCount, *left.ReplicaCount)
	}
	if left.PodAffinity != nil && *left.PodAffinity != right.PodAffinity {
		report.AddField(path+".podAffinity", right.PodAffinity, *left.PodAffinity)
	}
	if left.ContainerResources != nil {
		if right.ContainerResources == nil {
			report.AddField(path+".containerResources", nil, left.ContainerResources)
		} else {
			diffContainerResources(report, path+".containerResources", left.ContainerResources, right.ContainerResources)
		}
	}
	if left.PodTolerations != nil {
		diffPodTolerations(report, path+".podTolerations", left.PodTolerations, right.PodTolerations)
	}
}

func diffContainerResources(report *structuredreporting.Diff, path string, left *krm.ResourceRequirements, right *featureapi.PolicyControllerResourceRequirements) {
	if left == nil && right == nil {
		return
	}
	if left == nil {
		report.AddField(path, right, nil)
		return
	}
	if right == nil {
		report.AddField(path, nil, left)
		return
	}

	if left.Limits != nil {
		if right.Limits == nil {
			report.AddField(path+".limits", nil, left.Limits)
		} else {
			diffResourceList(report, path+".limits", left.Limits, right.Limits)
		}
	}
	if left.Requests != nil {
		if right.Requests == nil {
			report.AddField(path+".requests", nil, left.Requests)
		} else {
			diffResourceList(report, path+".requests", left.Requests, right.Requests)
		}
	}
}

func diffResourceList(report *structuredreporting.Diff, path string, left *krm.ResourceList, right *featureapi.PolicyControllerResourceList) {
	if left == nil && right == nil {
		return
	}
	if left == nil {
		report.AddField(path, right, nil)
		return
	}
	if right == nil {
		report.AddField(path, nil, left)
		return
	}

	if left.CPU != nil && *left.CPU != right.Cpu {
		report.AddField(path+".cpu", right.Cpu, *left.CPU)
	}
	if left.Memory != nil && *left.Memory != right.Memory {
		report.AddField(path+".memory", right.Memory, *left.Memory)
	}
}

func diffPodTolerations(report *structuredreporting.Diff, path string, left []krm.PolicyControllerDeploymentConfig_Toleration, right []*featureapi.PolicyControllerToleration) {
	if len(left) != len(right) {
		report.AddField(path, right, left)
		return
	}
	for i := range left {
		l := left[i]
		r := right[i]
		if l.Effect != nil && *l.Effect != r.Effect {
			report.AddField(path, right, left)
			return
		}
		if l.Key != nil && *l.Key != r.Key {
			report.AddField(path, right, left)
			return
		}
		if l.Operator != nil && *l.Operator != r.Operator {
			report.AddField(path, right, left)
			return
		}
		if l.Value != nil && *l.Value != r.Value {
			report.AddField(path, right, left)
			return
		}
	}
}
