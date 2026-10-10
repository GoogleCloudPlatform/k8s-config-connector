// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// API sources for FinancialServicesEngineConfig, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/financialservices/v1/engine_config.proto
// +kcc:source:service-docs=https://cloud.google.com/financial-services/anti-money-laundering/docs/concepts/overview
// +kcc:source:resource-docs=https://docs.cloud.google.com/financial-services/anti-money-laundering/docs/reference/rest/v1/projects.locations.instances.engineConfigs

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FinancialServicesEngineConfigGVK = GroupVersion.WithKind("FinancialServicesEngineConfig")

// FinancialServicesEngineConfigSpec defines the desired state of FinancialServicesEngineConfig
// +kcc:spec:proto=google.cloud.financialservices.v1.EngineConfig
// +kcc:required-from-proto
type FinancialServicesEngineConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// A reference to the projects/{project_num}/locations/{location}/instances/{instance} this resource belongs to.
	// +kcc:guess
	// InstanceRef *FinancialServicesInstanceRef `json:"instanceRef,omitempty"`

	// The FinancialServicesEngineConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Labels
	// +kcc:proto:field=google.cloud.financialservices.v1.EngineConfig.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The resource name of the EngineVersion used in this model tuning.
	//  format:
	//  `/projects/{project_num}/locations/{location}/instances/{instance}/engineVersions/{engine_version}`
	// +kcc:proto:field=google.cloud.financialservices.v1.EngineConfig.engine_version
	// +required
	EngineVersion *string `json:"engineVersion,omitempty"`

	// Optional. Configuration for tuning in creation of the EngineConfig.
	//  This field is required if `hyperparameter_source.type` is not `INHERITED`,
	//  and output-only otherwise.
	// +kcc:proto:field=google.cloud.financialservices.v1.EngineConfig.tuning
	Tuning *EngineConfig_Tuning `json:"tuning,omitempty"`

	// Optional. PerformanceTarget gives information on how the tuning and
	//  training will be evaluated. This field is required if
	//  `hyperparameter_source.type` is not `INHERITED`, and output-only otherwise.
	// +kcc:proto:field=google.cloud.financialservices.v1.EngineConfig.performance_target
	PerformanceTarget *EngineConfig_PerformanceTarget `json:"performanceTarget,omitempty"`

	// Optional. The origin of hyperparameters for the created EngineConfig. The
	//  default is `TUNING`. In this case, the hyperparameters are selected as a
	//  result of a
	//   tuning run.
	// +kcc:proto:field=google.cloud.financialservices.v1.EngineConfig.hyperparameter_source_type
	HyperparameterSourceType *string `json:"hyperparameterSourceType,omitempty"`

	// Optional. Configuration of hyperparameters source EngineConfig.
	// +kcc:proto:field=google.cloud.financialservices.v1.EngineConfig.hyperparameter_source
	HyperparameterSource *EngineConfig_HyperparameterSource `json:"hyperparameterSource,omitempty"`
}

// FinancialServicesEngineConfigStatus defines the config connector machine state of FinancialServicesEngineConfig
type FinancialServicesEngineConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the FinancialServicesEngineConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *FinancialServicesEngineConfigObservedState `json:"observedState,omitempty"`
}

// FinancialServicesEngineConfigObservedState is the state of the FinancialServicesEngineConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.financialservices.v1.EngineConfig
type FinancialServicesEngineConfigObservedState struct {
	// Output only. The timestamp of creation of this resource.
	// +kcc:proto:field=google.cloud.financialservices.v1.EngineConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp of the most recent update of this resource.
	// +kcc:proto:field=google.cloud.financialservices.v1.EngineConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the EngineConfig (creating, active, deleting, etc.)
	// +kcc:proto:field=google.cloud.financialservices.v1.EngineConfig.state
	State *string `json:"state,omitempty"`

	// Output only. The line of business (Retail/Commercial) this engine config is
	//  used for. Determined by EngineVersion, cannot be set by user.
	// +kcc:proto:field=google.cloud.financialservices.v1.EngineConfig.line_of_business
	LineOfBusiness *string `json:"lineOfBusiness,omitempty"`

	// Optional. Configuration of hyperparameters source EngineConfig.
	// +kcc:proto:field=google.cloud.financialservices.v1.EngineConfig.hyperparameter_source
	HyperparameterSource *EngineConfig_HyperparameterSourceObservedState `json:"hyperparameterSource,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpfinancialservicesengineconfig;gcpfinancialservicesengineconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// FinancialServicesEngineConfig is the Schema for the FinancialServicesEngineConfig API
// +k8s:openapi-gen=true
type FinancialServicesEngineConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FinancialServicesEngineConfigSpec   `json:"spec,omitempty"`
	Status FinancialServicesEngineConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// FinancialServicesEngineConfigList contains a list of FinancialServicesEngineConfig
type FinancialServicesEngineConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FinancialServicesEngineConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FinancialServicesEngineConfig{}, &FinancialServicesEngineConfigList{})
}
