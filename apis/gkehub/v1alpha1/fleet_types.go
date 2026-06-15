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

var GKEHubFleetGVK = GroupVersion.WithKind("GKEHubFleet")

// GKEHubFleetSpec defines the desired state of GKEHubFleet
// +kcc:spec:proto=google.cloud.gkehub.v1.Fleet
type GKEHubFleetSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The GKEHubFleet name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A user-assigned display name of the Fleet.
	// When present, it must be between 4 to 30 characters.
	// Allowed characters are: lowercase and uppercase letters, numbers,
	// hyphen, single-quote, double-quote, space, and exclamation point.
	//
	// Example: `Production Fleet`
	// +kcc:proto:field=google.cloud.gkehub.v1.Fleet.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The default cluster configurations to apply across the fleet.
	// +kcc:proto:field=google.cloud.gkehub.v1.Fleet.default_cluster_config
	DefaultClusterConfig *DefaultClusterConfig `json:"defaultClusterConfig,omitempty"`

	// Optional. Labels for this Fleet.
	// +kcc:proto:field=google.cloud.gkehub.v1.Fleet.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// GKEHubFleetStatus defines the config connector machine state of GKEHubFleet
type GKEHubFleetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GKEHubFleet resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GKEHubFleetObservedState `json:"observedState,omitempty"`
}

// GKEHubFleetObservedState is the state of the GKEHubFleet resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.gkehub.v1.Fleet
type GKEHubFleetObservedState struct {
	// Output only. The full, unique resource name of this fleet in the format of
	// `projects/{project}/locations/{location}/fleets/{fleet}`.
	//
	// Each Google Cloud project can have at most one fleet resource, named
	// "default".
	// +kcc:proto:field=google.cloud.gkehub.v1.Fleet.name
	Name *string `json:"name,omitempty"`

	// Output only. When the Fleet was created.
	// +kcc:proto:field=google.cloud.gkehub.v1.Fleet.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. When the Fleet was last updated.
	// +kcc:proto:field=google.cloud.gkehub.v1.Fleet.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. When the Fleet was deleted.
	// +kcc:proto:field=google.cloud.gkehub.v1.Fleet.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. Google-generated UUID for this resource. This is unique across
	// all Fleet resources. If a Fleet resource is deleted and another resource
	// with the same name is created, it gets a different uid.
	// +kcc:proto:field=google.cloud.gkehub.v1.Fleet.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. State of the Fleet resource on the server.
	// +kcc:proto:field=google.cloud.gkehub.v1.Fleet.state
	State *FleetLifecycleStateObservedState `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkehubfleet;gcpgkehubfleets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEHubFleet is the Schema for the GKEHubFleet API
// +k8s:openapi-gen=true
type GKEHubFleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GKEHubFleetSpec   `json:"spec,omitempty"`
	Status GKEHubFleetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GKEHubFleetList contains a list of GKEHubFleet
type GKEHubFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEHubFleet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEHubFleet{}, &GKEHubFleetList{})
}
