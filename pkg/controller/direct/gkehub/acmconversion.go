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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	featureapi "google.golang.org/api/gkehub/v1beta"
)

func convertKRMtoAPI_ConfigManagement(r *krm.FeaturemembershipConfigmanagement) (*featureapi.ConfigManagementMembershipSpec, error) {
	apiObj := featureapi.ConfigManagementMembershipSpec{}
	if r.Binauthz != nil {
		apiObj.Binauthz = convertKRMtoAPI_Binauthz(r.Binauthz)

	}
	if r.ConfigSync != nil {
		val, err := convertKRMtoAPI_ConfigSync(r.ConfigSync)
		if err != nil {
			return nil, err
		}
		apiObj.ConfigSync = val
	}
	if r.HierarchyController != nil {
		apiObj.HierarchyController = convertKRMtoAPI_HierachyController(r.HierarchyController)
	}
	if r.PolicyController != nil {
		val, err := convertKRMtoAPI_ConfigManagementPolicyController(r.PolicyController)
		if err != nil {
			return nil, err
		}
		apiObj.PolicyController = val
	}
	if r.Version != nil {
		apiObj.Version = *r.Version
	}
	return &apiObj, nil

}

func convertAPItoKRM_ConfigManagement(r *featureapi.ConfigManagementMembershipSpec) *krm.FeaturemembershipConfigmanagement {
	krmObj := krm.FeaturemembershipConfigmanagement{}
	if r.Binauthz != nil {
		krmObj.Binauthz = convertAPItoKRM_Binauthz(r.Binauthz)
	}
	if r.ConfigSync != nil {
		krmObj.ConfigSync = convertAPItoKRM_ConfigSync(r.ConfigSync)
	}
	if r.HierarchyController != nil {
		krmObj.HierarchyController = convertAPItoKRM_HierachyController(r.HierarchyController)
	}
	if r.PolicyController != nil {
		krmObj.PolicyController = convertAPItoKRM_ConfigManagementPolicyController(r.PolicyController)
	}
	krmObj.Version = LazyPtr(r.Version)
	return &krmObj
}

func convertKRMtoAPI_ConfigManagementPolicyController(r *krm.FeaturemembershipPolicyController) (*featureapi.ConfigManagementPolicyController, error) {
	apiObj := featureapi.ConfigManagementPolicyController{}
	if r.AuditIntervalSeconds != nil {
		val, err := convertStringToInt64(*r.AuditIntervalSeconds)
		if err != nil {
			return nil, err
		}
		apiObj.AuditIntervalSeconds = val
	}
	if r.Enabled != nil {
		apiObj.Enabled = *r.Enabled
	}
	if r.ExemptableNamespaces != nil {
		apiObj.ExemptableNamespaces = r.ExemptableNamespaces
	}
	if r.LogDeniesEnabled != nil {
		apiObj.LogDeniesEnabled = *r.LogDeniesEnabled
	}
	if r.Monitoring != nil {
		apiObj.Monitoring = convertKRMtoAPI_Monitoring(r.Monitoring)
	}
	if r.MutationEnabled != nil {
		apiObj.MutationEnabled = *r.MutationEnabled
	}
	if r.ReferentialRulesEnabled != nil {
		apiObj.ReferentialRulesEnabled = *r.ReferentialRulesEnabled
	}
	if r.TemplateLibraryInstalled != nil {
		apiObj.TemplateLibraryInstalled = *r.TemplateLibraryInstalled
	}
	return &apiObj, nil
}

func convertAPItoKRM_ConfigManagementPolicyController(r *featureapi.ConfigManagementPolicyController) *krm.FeaturemembershipPolicyController {
	krmObj := krm.FeaturemembershipPolicyController{}
	if r.AuditIntervalSeconds != 0 {
		krmObj.AuditIntervalSeconds = PtrTo(convertInt64toString(r.AuditIntervalSeconds))
	}
	krmObj.Enabled = LazyPtr(r.Enabled)
	if r.ExemptableNamespaces != nil {
		krmObj.ExemptableNamespaces = r.ExemptableNamespaces
	}
	krmObj.LogDeniesEnabled = LazyPtr(r.LogDeniesEnabled)
	if r.Monitoring != nil {
		krmObj.Monitoring = convertAPItoKRM_Monitoring(r.Monitoring)
	}
	krmObj.MutationEnabled = LazyPtr(r.MutationEnabled)
	krmObj.ReferentialRulesEnabled = LazyPtr(r.ReferentialRulesEnabled)
	krmObj.TemplateLibraryInstalled = LazyPtr(r.TemplateLibraryInstalled)
	return &krmObj
}

func convertKRMtoAPI_Monitoring(r *krm.FeaturemembershipMonitoring) *featureapi.ConfigManagementPolicyControllerMonitoring {
	apiObj := featureapi.ConfigManagementPolicyControllerMonitoring{}
	if r.Backends != nil {
		apiObj.Backends = r.Backends
	}
	return &apiObj
}

func convertAPItoKRM_Monitoring(r *featureapi.ConfigManagementPolicyControllerMonitoring) *krm.FeaturemembershipMonitoring {
	krmObj := krm.FeaturemembershipMonitoring{}
	if r.Backends != nil {
		krmObj.Backends = r.Backends
	}
	return &krmObj
}

func convertKRMtoAPI_HierachyController(r *krm.FeaturemembershipHierarchyController) *featureapi.ConfigManagementHierarchyControllerConfig {
	apiObj := featureapi.ConfigManagementHierarchyControllerConfig{}
	if r.EnableHierarchicalResourceQuota != nil {
		apiObj.EnableHierarchicalResourceQuota = *r.EnableHierarchicalResourceQuota
	}
	if r.EnablePodTreeLabels != nil {
		apiObj.EnablePodTreeLabels = *r.EnablePodTreeLabels
	}
	if r.Enabled != nil {
		apiObj.Enabled = *r.Enabled
	}
	return &apiObj
}

func convertAPItoKRM_HierachyController(r *featureapi.ConfigManagementHierarchyControllerConfig) *krm.FeaturemembershipHierarchyController {
	krmObj := krm.FeaturemembershipHierarchyController{}
	krmObj.EnableHierarchicalResourceQuota = LazyPtr(r.EnableHierarchicalResourceQuota)
	krmObj.EnablePodTreeLabels = LazyPtr(r.EnablePodTreeLabels)
	krmObj.Enabled = LazyPtr(r.Enabled)
	return &krmObj
}

func convertKRMtoAPI_Binauthz(r *krm.FeaturemembershipBinauthz) *featureapi.ConfigManagementBinauthzConfig {
	apiObj := featureapi.ConfigManagementBinauthzConfig{}
	if r.Enabled != nil {
		apiObj.Enabled = *r.Enabled
	}
	return &apiObj
}

func convertAPItoKRM_Binauthz(r *featureapi.ConfigManagementBinauthzConfig) *krm.FeaturemembershipBinauthz {
	krmObj := krm.FeaturemembershipBinauthz{}
	krmObj.Enabled = LazyPtr(r.Enabled)
	return &krmObj
}

func convertKRMtoAPI_ConfigSync(r *krm.FeaturemembershipConfigSync) (*featureapi.ConfigManagementConfigSync, error) {
	apiObj := featureapi.ConfigManagementConfigSync{}
	if r.Git != nil {
		val, err := convertKRMtoAPI_Git(r.Git)
		if err != nil {
			return nil, err
		}
		apiObj.Git = val
	}
	if r.Oci != nil {
		val, err := convertKRMtoAPI_Oci(r.Oci)
		if err != nil {
			return nil, err
		}
		apiObj.Oci = val
	}
	if r.PreventDrift != nil {
		apiObj.PreventDrift = *r.PreventDrift
	}
	if r.SourceFormat != nil {
		apiObj.SourceFormat = *r.SourceFormat
	}
	return &apiObj, nil
}

func convertAPItoKRM_ConfigSync(r *featureapi.ConfigManagementConfigSync) *krm.FeaturemembershipConfigSync {
	krmObj := krm.FeaturemembershipConfigSync{}
	if r.Git != nil {
		krmObj.Git = convertAPItoKRM_Git(r.Git)
	}
	if r.MetricsGcpServiceAccountEmail != "" {
		krmObj.MetricsGcpServiceAccountRef.External = r.MetricsGcpServiceAccountEmail
	}
	if r.Oci != nil {
		krmObj.Oci = convertAPItoKRM_Oci(r.Oci)
	}
	krmObj.PreventDrift = LazyPtr(r.PreventDrift)
	if r.SourceFormat != "" {
		krmObj.SourceFormat = &r.SourceFormat
	}
	return &krmObj
}

func convertKRMtoAPI_Git(r *krm.FeaturemembershipGit) (*featureapi.ConfigManagementGitConfig, error) {
	apiObj := featureapi.ConfigManagementGitConfig{}
	if r.HttpsProxy != nil {
		apiObj.HttpsProxy = *r.HttpsProxy
	}
	if r.PolicyDir != nil {
		apiObj.PolicyDir = *r.PolicyDir
	}
	if r.SecretType != nil {
		apiObj.SecretType = *r.SecretType
	}
	if r.SyncBranch != nil {
		apiObj.SecretType = *r.SyncBranch
	}
	if r.SyncRepo != nil {
		apiObj.SyncRepo = *r.SyncRepo
	}
	if r.SyncRev != nil {
		apiObj.SyncRev = *r.SyncRev
	}
	if r.SyncWaitSecs != nil {
		val, err := convertStringToInt64(*r.SyncWaitSecs)
		if err != nil {
			return nil, err
		}
		apiObj.SyncWaitSecs = val
	}
	return &apiObj, nil
}

func convertAPItoKRM_Git(r *featureapi.ConfigManagementGitConfig) *krm.FeaturemembershipGit {
	krmObj := krm.FeaturemembershipGit{}
	if r.GcpServiceAccountEmail != "" {
		krmObj.GcpServiceAccountRef.External = r.GcpServiceAccountEmail
	}
	krmObj.HttpsProxy = LazyPtr(r.HttpsProxy)
	krmObj.PolicyDir = LazyPtr(r.PolicyDir)
	krmObj.SecretType = LazyPtr(r.SecretType)
	krmObj.SyncBranch = LazyPtr(r.SyncBranch)
	krmObj.SyncRepo = LazyPtr(r.SyncRepo)
	krmObj.SyncRev = LazyPtr(r.SyncRev)
	// treat as unset if the value is 0
	if r.SyncWaitSecs != 0 {
		krmObj.SyncWaitSecs = PtrTo(convertInt64toString(r.SyncWaitSecs))
	}
	return &krmObj
}

func convertKRMtoAPI_Oci(r *krm.FeaturemembershipOci) (*featureapi.ConfigManagementOciConfig, error) {
	apiObj := featureapi.ConfigManagementOciConfig{}
	if r.PolicyDir != nil {
		apiObj.PolicyDir = *r.PolicyDir
	}
	if r.SecretType != nil {
		apiObj.SecretType = *r.SecretType
	}
	if r.SyncRepo != nil {
		apiObj.SyncRepo = *r.SyncRepo
	}
	if r.SyncWaitSecs != nil {
		val, err := convertStringToInt64(*r.SyncWaitSecs)
		if err != nil {
			return nil, err
		}
		apiObj.SyncWaitSecs = val
	}
	return &apiObj, nil
}

func convertAPItoKRM_Oci(r *featureapi.ConfigManagementOciConfig) *krm.FeaturemembershipOci {
	krmObj := krm.FeaturemembershipOci{}
	if r.GcpServiceAccountEmail != "" {
		krmObj.GcpServiceAccountRef.External = r.GcpServiceAccountEmail
	}
	krmObj.PolicyDir = LazyPtr(r.PolicyDir)
	krmObj.SecretType = LazyPtr(r.SecretType)
	krmObj.SyncRepo = LazyPtr(r.SyncRepo)

	if r.SyncWaitSecs != 0 {
		krmObj.SyncWaitSecs = PtrTo(convertInt64toString(r.SyncWaitSecs))
	}
	return &krmObj
}
