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

// API sources for FinancialServicesPredictionResult, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/financialservices/v1/prediction_result.proto
// +kcc:source:service-docs=https://cloud.google.com/financial-services/anti-money-laundering/docs/concepts/overview
// +kcc:source:resource-docs=https://docs.cloud.google.com/financial-services/anti-money-laundering/docs/reference/rest/v1/projects.locations.instances.predictionResults

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FinancialServicesPredictionResultGVK = GroupVersion.WithKind("FinancialServicesPredictionResult")

// FinancialServicesPredictionResultSpec defines the desired state of FinancialServicesPredictionResult
// +kcc:spec:proto=google.cloud.financialservices.v1.PredictionResult
// +kcc:required-from-proto
type FinancialServicesPredictionResultSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// A reference to the projects/{project_num}/locations/{location}/instances/{instance} this resource belongs to.
	// +kcc:guess
	// InstanceRef *FinancialServicesInstanceRef `json:"instanceRef,omitempty"`

	// The FinancialServicesPredictionResult name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Labels
	// +kcc:proto:field=google.cloud.financialservices.v1.PredictionResult.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The resource name of the Dataset to do predictions on
	//  Format:
	//  `/projects/{project_num}/locations/{location}/instances/{instance}/dataset/{dataset_id}`
	// +kcc:guess=possible-reference target=FinancialServicesDataset
	// +kcc:proto:field=google.cloud.financialservices.v1.PredictionResult.dataset
	// +required
	Dataset *string `json:"dataset,omitempty"`

	// Required. The resource name of the Model to use to use to make predictions
	//  Format:
	//  `/projects/{project_num}/locations/{location}/instances/{instance}/models/{model}`
	// +kcc:guess=possible-reference target=FinancialServicesModel
	// +kcc:proto:field=google.cloud.financialservices.v1.PredictionResult.model
	// +required
	Model *string `json:"model,omitempty"`

	// Required. Specifies the latest time from which data is used to generate
	//  features for predictions.  This time should be no later than the end of the
	//  [date_range][google.cloud.financialservices.v1.Dataset.date_range] of the
	//  dataset.
	// +kcc:proto:field=google.cloud.financialservices.v1.PredictionResult.end_time
	// +required
	EndTime *string `json:"endTime,omitempty"`

	// The number of consecutive months to produce predictions for, ending with
	//  the last full month prior to
	//  [end_time][google.cloud.financialservices.v1.PredictionResult.end_time]
	//  according to the dataset's timezone.
	// +kcc:proto:field=google.cloud.financialservices.v1.PredictionResult.prediction_periods
	PredictionPeriods *int32 `json:"predictionPeriods,omitempty"`

	// Required. Where to write the output of the predictions.
	// +kcc:proto:field=google.cloud.financialservices.v1.PredictionResult.outputs
	// +required
	Outputs *PredictionResult_Outputs `json:"outputs,omitempty"`
}

// FinancialServicesPredictionResultStatus defines the config connector machine state of FinancialServicesPredictionResult
type FinancialServicesPredictionResultStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the FinancialServicesPredictionResult resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *FinancialServicesPredictionResultObservedState `json:"observedState,omitempty"`
}

// FinancialServicesPredictionResultObservedState is the state of the FinancialServicesPredictionResult resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.financialservices.v1.PredictionResult
type FinancialServicesPredictionResultObservedState struct {
	// Output only. The timestamp of creation of this resource.
	// +kcc:proto:field=google.cloud.financialservices.v1.PredictionResult.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp of the most recent update of this resource.
	// +kcc:proto:field=google.cloud.financialservices.v1.PredictionResult.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the PredictionResult (creating, active, deleting,
	//  etc.)
	// +kcc:proto:field=google.cloud.financialservices.v1.PredictionResult.state
	State *string `json:"state,omitempty"`

	// Output only. The line of business (Retail/Commercial) this prediction is
	//  for. Determined by Model, cannot be set by user.
	// +kcc:proto:field=google.cloud.financialservices.v1.PredictionResult.line_of_business
	LineOfBusiness *string `json:"lineOfBusiness,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpfinancialservicespredictionresult;gcpfinancialservicespredictionresults
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// FinancialServicesPredictionResult is the Schema for the FinancialServicesPredictionResult API
// +k8s:openapi-gen=true
type FinancialServicesPredictionResult struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FinancialServicesPredictionResultSpec   `json:"spec,omitempty"`
	Status FinancialServicesPredictionResultStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// FinancialServicesPredictionResultList contains a list of FinancialServicesPredictionResult
type FinancialServicesPredictionResultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FinancialServicesPredictionResult `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FinancialServicesPredictionResult{}, &FinancialServicesPredictionResultList{})
}
