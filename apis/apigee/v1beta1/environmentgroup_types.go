// Copyright 2024 Google LLC
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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ApigeeEnvgroupGVK = GroupVersion.WithKind("ApigeeEnvgroup")

type Parent struct {
	// +required
	OrganizationRef *OrganizationRef `json:"organizationRef"`
}

// ApigeeEnvgroupSpec defines the desired state of ApigeeEnvgroup
// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1EnvironmentGroup
type ApigeeEnvgroupSpec struct {
	Parent `json:",inline"`
	// The ApigeeEnvgroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Host names for this environment group.
	Hostnames []string `json:"hostnames,omitempty"`
}

// ApigeeEnvgroupStatus defines the config connector machine state of ApigeeEnvgroup
type ApigeeEnvgroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ApigeeEnvgroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ApigeeEnvgroupObservedState `json:"observedState,omitempty"`
}

// ApigeeEnvgroupSpec defines the desired state of ApigeeEnvgroup
// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1EnvironmentGroup
// ApigeeEnvgroupObservedState is the state of the ApigeeEnvgroup resource as most recently observed in GCP.
type ApigeeEnvgroupObservedState struct {
	// ID of the environment group.
	Name *string `json:"name,omitempty"`
	// Output only. The time at which the environment group was created as milliseconds since epoch.
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// Output only. The time at which the environment group was last updated as milliseconds since epoch.
	LastModifiedAt *int64 `json:"lastModifiedAt,omitempty"`
	// Output only. State of the environment group. Values other than ACTIVE means the resource is not ready to use.
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeenvgroup;gcpapigeeenvgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// +kubebuilder:storageversion

// ApigeeEnvgroup is the Schema for the ApigeeEnvgroup API
// +k8s:openapi-gen=true
type ApigeeEnvgroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ApigeeEnvgroupSpec   `json:"spec,omitempty"`
	Status ApigeeEnvgroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ApigeeEnvgroupList contains a list of ApigeeEnvgroup
type ApigeeEnvgroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeEnvgroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeEnvgroup{}, &ApigeeEnvgroupList{})
}
