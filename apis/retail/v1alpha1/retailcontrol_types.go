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

// API sources for RetailControl, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/retail/v2/control.proto
// +kcc:source:service-docs=https://cloud.google.com/recommendations
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/recommendations/docs/reference/rest/v2/projects.locations.catalogs.controls

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var RetailControlGVK = GroupVersion.WithKind("RetailControl")

// RetailControlSpec defines the desired state of RetailControl
// +kcc:spec:proto=google.cloud.retail.v2.Control
// +kcc:required-from-proto
type RetailControlSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The RetailControl name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// A rule control - a condition-action pair.
	//  Enacts a set action when the condition is triggered.
	//  For example: Boost "gShoe" when query full matches "Running Shoes".
	// +kcc:proto:field=google.cloud.retail.v2.Control.rule
	Rule *Rule `json:"rule,omitempty"`

	// Required. The human readable control display name. Used in Retail UI.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 128
	//  characters. Otherwise, an INVALID_ARGUMENT error is thrown.
	// +kcc:proto:field=google.cloud.retail.v2.Control.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. The solution types that the control is used for.
	//  Currently we support setting only one type of solution at creation time.
	//
	//  Only `SOLUTION_TYPE_SEARCH` value is supported at the moment.
	//  If no solution type is provided at creation time, will default to
	//  [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2.SolutionType.SOLUTION_TYPE_SEARCH].
	// +kcc:proto:field=google.cloud.retail.v2.Control.solution_types
	// +required
	SolutionTypes []string `json:"solutionTypes,omitempty"`

	// Specifies the use case for the control.
	//  Affects what condition fields can be set.
	//  Only settable by search controls.
	//  Will default to
	//  [SEARCH_SOLUTION_USE_CASE_SEARCH][google.cloud.retail.v2.SearchSolutionUseCase.SEARCH_SOLUTION_USE_CASE_SEARCH]
	//  if not specified. Currently only allow one search_solution_use_case per
	//  control.
	// +kcc:proto:field=google.cloud.retail.v2.Control.search_solution_use_case
	SearchSolutionUseCase []string `json:"searchSolutionUseCase,omitempty"`
}

// RetailControlStatus defines the config connector machine state of RetailControl
type RetailControlStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the RetailControl resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *RetailControlObservedState `json:"observedState,omitempty"`
}

// RetailControlObservedState is the state of the RetailControl resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.retail.v2.Control
type RetailControlObservedState struct {
	// Output only. List of [serving config][google.cloud.retail.v2.ServingConfig]
	//  ids that are associated with this control in the same
	//  [Catalog][google.cloud.retail.v2.Catalog].
	//
	//  Note the association is managed via the
	//  [ServingConfig][google.cloud.retail.v2.ServingConfig], this is an output
	//  only denormalized view.
	// +kcc:proto:field=google.cloud.retail.v2.Control.associated_serving_config_ids
	AssociatedServingConfigIDs []string `json:"associatedServingConfigIDs,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpretailcontrol;gcpretailcontrols
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// RetailControl is the Schema for the RetailControl API
// +k8s:openapi-gen=true
type RetailControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   RetailControlSpec   `json:"spec,omitempty"`
	Status RetailControlStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RetailControlList contains a list of RetailControl
type RetailControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RetailControl `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RetailControl{}, &RetailControlList{})
}
