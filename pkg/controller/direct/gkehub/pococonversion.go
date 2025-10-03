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

func convertKRMtoAPI_PolicycontrollerHubConfig(r *krm.HubConfig) *featureapi.PolicyControllerHubConfig {
	if r == nil {
		return nil
	}

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
	if r.DeploymentConfigs != nil {
		apiObj.DeploymentConfigs = convertKRMtoAPI_DeploymentConfigs(r.DeploymentConfigs)
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

func convertKRMtoAPI_DeploymentConfigs(r *krm.PolicyControllerDeploymentConfigs) map[string]featureapi.PolicyControllerPolicyControllerDeploymentConfig {
	if r == nil {
		return nil
	}
	out := make(map[string]featureapi.PolicyControllerPolicyControllerDeploymentConfig)
	if r.Admission != nil {
		out["admission"] = convertKRMtoAPI_PolicyControllerDeploymentConfig(r.Admission)
	}
	if r.Audit != nil {
		out["audit"] = convertKRMtoAPI_PolicyControllerDeploymentConfig(r.Audit)
	}
	if r.Mutation != nil {
		out["mutation"] = convertKRMtoAPI_PolicyControllerDeploymentConfig(r.Mutation)
	}
	return out
}

func convertKRMtoAPI_PolicyControllerDeploymentConfig(in *krm.PolicyControllerDeploymentConfig) featureapi.PolicyControllerPolicyControllerDeploymentConfig {
	out := featureapi.PolicyControllerPolicyControllerDeploymentConfig{}
	if in == nil {
		return out
	}
	if in.ContainerResources != nil {
		out.ContainerResources = &featureapi.PolicyControllerResourceRequirements{
			Limits:   convertKRMToAPI_PolicyControllerResourceList(in.ContainerResources.Limits),
			Requests: convertKRMToAPI_PolicyControllerResourceList(in.ContainerResources.Requests),
		}
	}

	if in.PodAffinity != nil {
		out.PodAffinity = direct.ValueOf(in.PodAffinity)
	}

	if in.ReplicaCount != nil {
		out.ReplicaCount = direct.ValueOf(in.ReplicaCount)
	}

	for _, v := range in.PodTolerations {
		out.PodTolerations = append(out.PodTolerations, &featureapi.PolicyControllerToleration{
			Effect:   direct.ValueOf(v.Effect),
			Key:      direct.ValueOf(v.Key),
			Operator: direct.ValueOf(v.Operator),
			Value:    direct.ValueOf(v.Value),
		})
	}

	return out
}

func convertKRMToAPI_PolicyControllerResourceList(in *krm.ResourceList) *featureapi.PolicyControllerResourceList {
	if in == nil {
		return nil
	}
	out := &featureapi.PolicyControllerResourceList{}
	out.Cpu = direct.ValueOf(in.CPU)
	out.Memory = direct.ValueOf(in.Memory)
	return out
}

func DeploymentConfigs_FromProto(mapCtx *direct.MapContext, in map[string]featureapi.PolicyControllerPolicyControllerDeploymentConfig) *krm.PolicyControllerDeploymentConfigs {
	if in == nil {
		return nil
	}
	out := &krm.PolicyControllerDeploymentConfigs{}
	if val, ok := in["admission"]; ok {
		out.Admission = PolicyControllerDeploymentConfig_FromProto(mapCtx, &val)
	}
	if val, ok := in["audit"]; ok {
		out.Audit = PolicyControllerDeploymentConfig_FromProto(mapCtx, &val)
	}
	if val, ok := in["mutation"]; ok {
		out.Mutation = PolicyControllerDeploymentConfig_FromProto(mapCtx, &val)
	}
	return out
}

func PolicyControllerDeploymentConfig_FromProto(mapCtx *direct.MapContext, in *featureapi.PolicyControllerPolicyControllerDeploymentConfig) *krm.PolicyControllerDeploymentConfig {
	if in == nil {
		return nil
	}
	out := &krm.PolicyControllerDeploymentConfig{}

	if in.ContainerResources != nil {
		out.ContainerResources = &krm.ResourceRequirements{
			Limits:   ResourceList_FromProto(in.ContainerResources.Limits),
			Requests: ResourceList_FromProto(in.ContainerResources.Requests),
		}
	}
	out.PodAffinity = direct.LazyPtr(in.PodAffinity)
	out.ReplicaCount = direct.LazyPtr(in.ReplicaCount)
	for _, v := range in.PodTolerations {
		out.PodTolerations = append(out.PodTolerations, krm.PolicyControllerDeploymentConfig_Toleration{
			Effect:   direct.LazyPtr(v.Effect),
			Key:      direct.LazyPtr(v.Key),
			Operator: direct.LazyPtr(v.Operator),
			Value:    direct.LazyPtr(v.Value),
		})
	}

	return out
}

func ResourceList_FromProto(in *featureapi.PolicyControllerResourceList) *krm.ResourceList {
	if in == nil {
		return nil
	}
	out := &krm.ResourceList{}
	out.Memory = direct.LazyPtr(in.Memory)
	out.CPU = direct.LazyPtr(in.Cpu)
	return out
}
