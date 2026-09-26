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

var ChronicleWatchlistGVK = GroupVersion.WithKind("ChronicleWatchlist")

// ChronicleWatchlistSpec defines the desired state of ChronicleWatchlist
// +kcc:spec:proto=google.cloud.chronicle.v1.Watchlist
type ChronicleWatchlistSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ChronicleWatchlist name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Display name of the watchlist.
	//  Note that it must be at least one character and less than 63 characters
	//  (https://google.aip.dev/148).
	// +kcc:proto:field=google.cloud.chronicle.v1.Watchlist.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the watchlist.
	// +kcc:proto:field=google.cloud.chronicle.v1.Watchlist.description
	Description *string `json:"description,omitempty"`

	// Optional. Weight applied to the risk score for entities
	//  in this watchlist.
	//  The default is 1.0 if it is not specified.
	// +kcc:proto:field=google.cloud.chronicle.v1.Watchlist.multiplying_factor
	MultiplyingFactor *float32 `json:"multiplyingFactor,omitempty"`

	// Required. Mechanism to populate entities in the watchlist.
	// +kcc:proto:field=google.cloud.chronicle.v1.Watchlist.entity_population_mechanism
	// +required
	EntityPopulationMechanism *Watchlist_EntityPopulationMechanism `json:"entityPopulationMechanism,omitempty"`

	// Optional. User preferences for watchlist configuration.
	// +kcc:proto:field=google.cloud.chronicle.v1.Watchlist.watchlist_user_preferences
	WatchlistUserPreferences *WatchlistUserPreferences `json:"watchlistUserPreferences,omitempty"`
}

// ChronicleWatchlistStatus defines the config connector machine state of ChronicleWatchlist
type ChronicleWatchlistStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ChronicleWatchlist resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ChronicleWatchlistObservedState `json:"observedState,omitempty"`
}

// ChronicleWatchlistObservedState is the state of the ChronicleWatchlist resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.chronicle.v1.Watchlist
type ChronicleWatchlistObservedState struct {
	// Output only. Entity count in the watchlist.
	// +kcc:proto:field=google.cloud.chronicle.v1.Watchlist.entity_count
	EntityCount *Watchlist_EntityCountObservedState `json:"entityCount,omitempty"`

	// Output only. Time the watchlist was created.
	// +kcc:proto:field=google.cloud.chronicle.v1.Watchlist.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time the watchlist was last updated.
	// +kcc:proto:field=google.cloud.chronicle.v1.Watchlist.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpchroniclewatchlist;gcpchroniclewatchlists
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ChronicleWatchlist is the Schema for the ChronicleWatchlist API
// +k8s:openapi-gen=true
type ChronicleWatchlist struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ChronicleWatchlistSpec   `json:"spec,omitempty"`
	Status ChronicleWatchlistStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ChronicleWatchlistList contains a list of ChronicleWatchlist
type ChronicleWatchlistList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChronicleWatchlist `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChronicleWatchlist{}, &ChronicleWatchlistList{})
}
