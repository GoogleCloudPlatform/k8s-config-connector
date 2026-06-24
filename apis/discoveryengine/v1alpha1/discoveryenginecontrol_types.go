// Copyright 2026 Google LLC
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

var DiscoveryEngineControlGVK = GroupVersion.WithKind("DiscoveryEngineControl")

// DiscoveryEngineControlSpec defines the desired state of DiscoveryEngineControl
// +kcc:spec:proto=google.cloud.discoveryengine.v1.Control
type DiscoveryEngineControlSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// +required
	Location *string `json:"location"`

	// Immutable. The DataStore this control belongs to.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="DataStoreRef field is immutable"
	// +required
	DataStoreRef *DiscoveryEngineDataStoreRef `json:"dataStoreRef"`

	// The DiscoveryEngineControl name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Human readable name. The identifier used in UI views.
	//
	// Must be UTF-8 encoded string. Length limit is 128 characters.
	// Otherwise an INVALID ARGUMENT error is thrown.
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. What solution the control belongs to.
	//
	// Must be compatible with vertical of resource.
	// Otherwise an INVALID ARGUMENT error is thrown.
	SolutionType *string `json:"solutionType,omitempty"`

	// Specifies the use case for the control.
	// Affects what condition fields can be set.
	// Only applies to
	// [SOLUTION_TYPE_SEARCH][google.cloud.discoveryengine.v1.SolutionType.SOLUTION_TYPE_SEARCH].
	// Currently only allow one use case per control.
	// Must be set when solution_type is
	// [SolutionType.SOLUTION_TYPE_SEARCH][google.cloud.discoveryengine.v1.SolutionType.SOLUTION_TYPE_SEARCH].
	UseCases []string `json:"useCases,omitempty"`

	// Determines when the associated action will trigger.
	//
	// Omit to always apply the action.
	// Currently only a single condition may be specified.
	// Otherwise an INVALID ARGUMENT error is thrown.
	Conditions []Condition `json:"conditions,omitempty"`

	// Defines a boost-type control
	BoostAction *Control_BoostAction `json:"boostAction,omitempty"`

	// Defines a filter-type control
	// Currently not supported by Recommendation
	FilterAction *Control_FilterAction `json:"filterAction,omitempty"`

	// Defines a redirect-type control.
	RedirectAction *Control_RedirectAction `json:"redirectAction,omitempty"`

	// Treats a group of terms as synonyms of one another.
	SynonymsAction *Control_SynonymsAction `json:"synonymsAction,omitempty"`

	// Promote certain links based on predefined trigger queries.
	PromoteAction *Control_PromoteAction `json:"promoteAction,omitempty"`
}

// DiscoveryEngineControlStatus defines the config connector machine state of DiscoveryEngineControl
type DiscoveryEngineControlStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineControl resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineControlObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineControlObservedState is the state of the DiscoveryEngineControl resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1.Control
type DiscoveryEngineControlObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginecontrol;gcpdiscoveryenginecontrols
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineControl is the Schema for the DiscoveryEngineControl API
// +k8s:openapi-gen=true
type DiscoveryEngineControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineControlSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineControlStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineControlList contains a list of DiscoveryEngineControl
type DiscoveryEngineControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineControl `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineControl{}, &DiscoveryEngineControlList{})
}

// +kcc:proto=google.cloud.discoveryengine.v1.Control.BoostAction
type Control_BoostAction struct {
	// Optional. Strength of the boost, which should be in [-1, 1]. Negative
	//  boost means demotion. Default is 0.0 (No-op).
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Control.BoostAction.fixed_boost
	FixedBoost *float32 `json:"fixedBoost,omitempty"`

	// Optional. Complex specification for custom ranking based on customer
	//  defined attribute value.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Control.BoostAction.interpolation_boost_spec
	InterpolationBoostSpec *Control_BoostAction_InterpolationBoostSpec `json:"interpolationBoostSpec,omitempty"`

	// Strength of the boost, which should be in [-1, 1]. Negative
	//  boost means demotion. Default is 0.0 (No-op).
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Control.BoostAction.boost
	Boost *float32 `json:"boost,omitempty"`

	// Required. Specifies which products to apply the boost to.
	//
	//  If no filter is provided all products will be boosted (No-op).
	//  Syntax documentation:
	//  https://cloud.google.com/retail/docs/filter-and-order
	//  Maximum length is 5000 characters.
	//  Otherwise an INVALID ARGUMENT error is thrown.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Control.BoostAction.filter
	Filter *string `json:"filter,omitempty"`

	// Required. Specifies which data store's documents can be boosted by this
	//  control. Full data store name e.g.
	//  projects/123/locations/global/collections/default_collection/dataStores/default_data_store
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Control.BoostAction.data_store
	DataStoreRef *DiscoveryEngineDataStoreRef `json:"dataStoreRef,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Control.FilterAction
type Control_FilterAction struct {
	// Required. A filter to apply on the matching condition results.
	//
	//  Required
	//  Syntax documentation:
	//  https://cloud.google.com/retail/docs/filter-and-order
	//  Maximum length is 5000 characters. Otherwise an INVALID
	//  ARGUMENT error is thrown.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Control.FilterAction.filter
	Filter *string `json:"filter,omitempty"`

	// Required. Specifies which data store's documents can be filtered by this
	//  control. Full data store name e.g.
	//  projects/123/locations/global/collections/default_collection/dataStores/default_data_store
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Control.FilterAction.data_store
	DataStoreRef *DiscoveryEngineDataStoreRef `json:"dataStoreRef,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.Control.PromoteAction
type Control_PromoteAction struct {
	// Required. Data store with which this promotion is attached to.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Control.PromoteAction.data_store
	DataStoreRef *DiscoveryEngineDataStoreRef `json:"dataStoreRef,omitempty"`

	// Required. Promotion attached to this action.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Control.PromoteAction.search_link_promotion
	SearchLinkPromotion *SearchLinkPromotion `json:"searchLinkPromotion,omitempty"`
}
