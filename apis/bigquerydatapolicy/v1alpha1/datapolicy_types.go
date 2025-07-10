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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigQueryDataPolicyDataPolicyGVK = GroupVersion.WithKind("BigQueryDataPolicyDataPolicy")

// +kcc:proto=google.cloud.bigquery.datapolicies.v1beta1.DataMaskingPolicy
type DataMaskingPolicy struct {
	// A predefined masking expression.
	// Learn more here: https://cloud.google.com/bigquery/docs/reference/bigquerydatapolicy/rest/v1beta1/projects.locations.dataPolicies#predefinedexpression
	// +kcc:proto:field=google.cloud.bigquery.datapolicies.v1beta1.DataMaskingPolicy.predefined_expression
	PredefinedExpression *string `json:"predefinedExpression,omitempty"`
}

type Parent struct {
	// Required. The location of the application.
	// +required
	Location string `json:"location,omitempty"`

	// Required. The host project of the application.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`
}

// BigQueryDataPolicyDataPolicySpec defines the desired state of BigQueryDataPolicy
// +kcc:spec:proto=google.cloud.bigquery.datapolicies.v1beta1.DataPolicy
type BigQueryDataPolicyDataPolicySpec struct {
	// The BigQueryDataPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Defines the parent path of the resource.
	*Parent `json:",inline"`

	/* Policy tag resource name, in the format of projects/{project_number}/locations/{locationId}/taxonomies/{taxonomyId}/policyTags/{policyTag_id}. */
	PolicyTag *string `json:"policyTag"`

	// The data masking policy that specifies the data masking rule to use.
	// +kcc:proto:field=google.cloud.bigquery.datapolicies.v1beta1.DataPolicy.data_masking_policy
	DataMaskingPolicy *DataMaskingPolicy `json:"dataMaskingPolicy,omitempty"`

	// Required. Data policy type. Type of data policy.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.datapolicies.v1beta1.DataPolicy.data_policy_type
	DataPolicyType *string `json:"dataPolicyType,omitempty"`
}

// BigQueryDataPolicyDataPolicyStatus defines the config connector machine state of BigQueryDataPolicy
type BigQueryDataPolicyDataPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigQueryDataPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryDataPolicyDataPolicyObservedState `json:"observedState,omitempty"`
}

// BigQueryDataPolicyDataPolicyObservedState is the state of the BigQueryDataPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.bigquery.datapolicies.v1beta1.DataPolicy
type BigQueryDataPolicyDataPolicyObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigquerydatapolicydatapolicy;gcpbigquerydatapolicydatapolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryDataPolicyDataPolicy is the Schema for the BigQueryDataPolicy API
// +k8s:openapi-gen=true
type BigQueryDataPolicyDataPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigQueryDataPolicyDataPolicySpec   `json:"spec,omitempty"`
	Status BigQueryDataPolicyDataPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryDataPolicyDataPolicyList contains a list of BigQueryDataPolicyDataPolicy
type BigQueryDataPolicyDataPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryDataPolicyDataPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryDataPolicyDataPolicy{}, &BigQueryDataPolicyDataPolicyList{})
}
