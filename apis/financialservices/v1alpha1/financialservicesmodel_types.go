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

// API sources for FinancialServicesModel, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/financialservices/v1/model.proto
// +kcc:source:service-docs=https://cloud.google.com/financial-services/anti-money-laundering/docs/concepts/overview
// +kcc:source:resource-docs=https://docs.cloud.google.com/financial-services/anti-money-laundering/docs/reference/rest/v1/projects.locations.instances.models

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FinancialServicesModelGVK = GroupVersion.WithKind("FinancialServicesModel")

// FinancialServicesModelSpec defines the desired state of FinancialServicesModel
// +kcc:spec:proto=google.cloud.financialservices.v1.Model
// +kcc:required-from-proto
type FinancialServicesModelSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// A reference to the projects/{project_num}/locations/{location}/instances/{instance} this resource belongs to.
	// +kcc:guess
	// InstanceRef *FinancialServicesInstanceRef `json:"instanceRef,omitempty"`

	// The FinancialServicesModel name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Labels
	// +kcc:proto:field=google.cloud.financialservices.v1.Model.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The resource name of the EngineConfig the model training will be
	//  based on. Format:
	//  `/projects/{project_num}/locations/{location}/instances/{instance}/engineConfigs/{engineConfig}`
	// +kcc:guess=possible-reference target=FinancialServicesEngineConfig
	// +kcc:proto:field=google.cloud.financialservices.v1.Model.engine_config
	// +required
	EngineConfig *string `json:"engineConfig,omitempty"`

	// Required. The resource name of the Primary Dataset used in this model
	//  training. For information about how primary and auxiliary datasets are
	//  used, refer to the engine version's documentation.  Format:
	//  `/projects/{project_num}/locations/{location}/instances/{instance}/datasets/{dataset}`
	// +kcc:proto:field=google.cloud.financialservices.v1.Model.primary_dataset
	// +required
	PrimaryDataset *string `json:"primaryDataset,omitempty"`

	// Required. End_time specifies the latest time from which labels are used and
	//  from which data is used to generate features for training.  End_time should
	//  be no later than the end of the date_range of the dataset.
	// +kcc:proto:field=google.cloud.financialservices.v1.Model.end_time
	// +required
	EndTime *string `json:"endTime,omitempty"`
}

// FinancialServicesModelStatus defines the config connector machine state of FinancialServicesModel
type FinancialServicesModelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the FinancialServicesModel resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *FinancialServicesModelObservedState `json:"observedState,omitempty"`
}

// FinancialServicesModelObservedState is the state of the FinancialServicesModel resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.financialservices.v1.Model
type FinancialServicesModelObservedState struct {
	// Output only. The timestamp of creation of this resource.
	// +kcc:proto:field=google.cloud.financialservices.v1.Model.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp of the most recent update of this resource.
	// +kcc:proto:field=google.cloud.financialservices.v1.Model.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the model (creating, active, deleting, etc.)
	// +kcc:proto:field=google.cloud.financialservices.v1.Model.state
	State *string `json:"state,omitempty"`

	// Output only. The EngineVersion used in training this model.  This is output
	//  only, and is determined from the EngineConfig used.
	// +kcc:proto:field=google.cloud.financialservices.v1.Model.engine_version
	EngineVersion *string `json:"engineVersion,omitempty"`

	// Output only. The line of business (Retail/Commercial) this model is used
	//  for. Determined by EngineConfig, cannot be set by user.
	// +kcc:proto:field=google.cloud.financialservices.v1.Model.line_of_business
	LineOfBusiness *string `json:"lineOfBusiness,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpfinancialservicesmodel;gcpfinancialservicesmodels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// FinancialServicesModel is the Schema for the FinancialServicesModel API
// +k8s:openapi-gen=true
type FinancialServicesModel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FinancialServicesModelSpec   `json:"spec,omitempty"`
	Status FinancialServicesModelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// FinancialServicesModelList contains a list of FinancialServicesModel
type FinancialServicesModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FinancialServicesModel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FinancialServicesModel{}, &FinancialServicesModelList{})
}
