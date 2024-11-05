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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	featureapi "google.golang.org/api/gkehub/v1beta"
)

// standalone Poco
func convertKRMtoAPI_Policycontroller(r *krm.FeaturemembershipPolicycontroller) *featureapi.PolicyControllerMembershipSpec {
	apiObj := &featureapi.PolicyControllerMembershipSpec{}
	apiObj.PolicyControllerHubConfig = convertKRMtoAPI_PolicycontrollerHubConfig(r.PolicyControllerHubConfig)
	if r.Version != nil {
		apiObj.Version = *r.Version
	}
	return apiObj
}

func convertKRMtoAPI_PolicycontrollerHubConfig(r krm.FeaturemembershipPolicyControllerHubConfig) *featureapi.PolicyControllerHubConfig {
	apiObj := &featureapi.PolicyControllerHubConfig{}
	if r.AuditIntervalSeconds != nil {
		apiObj.AuditIntervalSeconds = int64(direct.ValueOf(r.AuditIntervalSeconds))
	}
	if r.ConstraintViolationLimit != nil {
		apiObj.ConstraintViolationLimit = int64(direct.ValueOf(r.ConstraintViolationLimit))
	}
	if r.ExemptableNamespaces != nil {
		apiObj.ExemptableNamespaces = r.ExemptableNamespaces
	}
	if r.InstallSpec != nil {
		apiObj.InstallSpec = *r.InstallSpec
	}
	if r.LogDeniesEnabled != nil {
		apiObj.LogDeniesEnabled = *r.LogDeniesEnabled
	}
	if r.Monitoring != nil {
		apiObj.Monitoring = convertKRMtoAPI_FeaturemembershipMonitoring(r.Monitoring)
	}
	if r.MutationEnabled != nil {
		apiObj.MutationEnabled = *r.MutationEnabled
	}
	if r.PolicyContent != nil {
		apiObj.PolicyContent = convertKRMtoAPI_PolicyContent(r.PolicyContent)
	}
	if r.ReferentialRulesEnabled != nil {
		apiObj.ReferentialRulesEnabled = *r.ReferentialRulesEnabled
	}
	return apiObj
}

func convertKRMtoAPI_FeaturemembershipMonitoring(r *krm.FeaturemembershipMonitoring) *featureapi.PolicyControllerMonitoringConfig {
	apiObj := &featureapi.PolicyControllerMonitoringConfig{}
	if r.Backends != nil {
		apiObj.Backends = r.Backends
	}
	return apiObj
}

func convertKRMtoAPI_PolicyContent(r *krm.FeaturemembershipPolicyContent) *featureapi.PolicyControllerPolicyContentSpec {
	apiObj := &featureapi.PolicyControllerPolicyContentSpec{}
	if r.TemplateLibrary != nil {
		apiObj.TemplateLibrary = convertKRMtoAPI_TemplateLibrary(r.TemplateLibrary)
	}
	return apiObj
}

func convertKRMtoAPI_TemplateLibrary(r *krm.FeaturemembershipTemplateLibrary) *featureapi.PolicyControllerTemplateLibraryConfig {
	apiObj := &featureapi.PolicyControllerTemplateLibraryConfig{}
	if r.Installation != nil {
		apiObj.Installation = *r.Installation
	}
	return apiObj
}
