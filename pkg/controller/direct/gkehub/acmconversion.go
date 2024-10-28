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
	"strconv"

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
		apiObj.HierarchyController = convertKRMtoAPI_HierarchyController(r.HierarchyController)
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

func convertKRMtoAPI_Monitoring(r *krm.FeaturemembershipMonitoring) *featureapi.ConfigManagementPolicyControllerMonitoring {
	apiObj := featureapi.ConfigManagementPolicyControllerMonitoring{}
	if r.Backends != nil {
		apiObj.Backends = r.Backends
	}
	return &apiObj
}

func convertKRMtoAPI_HierarchyController(r *krm.FeaturemembershipHierarchyController) *featureapi.ConfigManagementHierarchyControllerConfig {
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

func convertKRMtoAPI_Binauthz(r *krm.FeaturemembershipBinauthz) *featureapi.ConfigManagementBinauthzConfig {
	apiObj := featureapi.ConfigManagementBinauthzConfig{}
	if r.Enabled != nil {
		apiObj.Enabled = *r.Enabled
	}
	return &apiObj
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
	if r.MetricsGcpServiceAccountRef != nil {
		// the IAM references have been resolved as external as this point.
		apiObj.MetricsGcpServiceAccountEmail = r.MetricsGcpServiceAccountRef.External
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

func convertKRMtoAPI_Git(r *krm.FeaturemembershipGit) (*featureapi.ConfigManagementGitConfig, error) {
	apiObj := featureapi.ConfigManagementGitConfig{}
	if r.GcpServiceAccountRef != nil {
		// the IAM references have been resolved as external as this point.
		apiObj.GcpServiceAccountEmail = r.GcpServiceAccountRef.External
	}
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
		apiObj.SyncBranch = *r.SyncBranch
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

func convertKRMtoAPI_Oci(r *krm.FeaturemembershipOci) (*featureapi.ConfigManagementOciConfig, error) {
	apiObj := featureapi.ConfigManagementOciConfig{}
	if r.GcpServiceAccountRef != nil {
		// the IAM references have been resolved as external as this point.
		apiObj.GcpServiceAccountEmail = r.GcpServiceAccountRef.External
	}
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

func convertStringToInt64(s string) (int64, error) {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}
