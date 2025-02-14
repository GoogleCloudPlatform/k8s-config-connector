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

package v1alpha1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FilestoreInstanceGVK = GroupVersion.WithKind("FilestoreInstance")

// FilestoreInstanceSpec defines the desired state of FilestoreInstance
// +kcc:proto=google.cloud.filestore.v1.Instance
type FilestoreInstanceSpec struct {
	/* Immutable. The zone for this filestore instance. */
	Location string `json:"location"`

	/* The project that this resource belongs to. */
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// The description of the instance (2048 characters or less).
	Description *string `json:"description,omitempty"`

	// The service tier of the instance.
	Tier *string `json:"tier,omitempty"`

	// NOTYET
	// Resource labels to represent user provided metadata.
	// Labels []InstanceLabel `json:"labels,omitempty"`

	// File system shares on the instance.
	//  For this version, only a single file share is supported.
	FileShares []FileShareConfig `json:"fileShares,omitempty"`

	// VPC networks to which the instance is connected.
	//  For this version, only a single network is supported.
	Networks []NetworkConfig `json:"networks,omitempty"`

	// The KMS key name used for data encryption.
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// The FilestoreInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// FilestoreInstanceStatus defines the config connector machine state of FilestoreInstance
type FilestoreInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the FilestoreInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *FilestoreInstanceObservedState `json:"observedState,omitempty"`
}

// FilestoreInstanceObservedState is the state of the FilestoreInstance resource as most recently observed in GCP.
// +kcc:proto=google.cloud.filestore.v1.Instance
type FilestoreInstanceObservedState struct {
	// Output only. The resource name of the instance, in the format
	//  `projects/{project}/locations/{location}/instances/{instance}`.
	Name *string `json:"name,omitempty"`

	// Output only. The instance state.
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the instance state, if available.
	StatusMessage *string `json:"statusMessage,omitempty"`

	// Output only. The time when the instance was created.
	CreateTime *string `json:"createTime,omitempty"`

	// VPC networks to which the instance is connected.
	//  For this version, only a single network is supported.
	Networks []NetworkConfigObservedState `json:"networks,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`

	// Output only. Field indicates all the reasons the instance is in "SUSPENDED"
	//  state.
	SuspensionReasons []string `json:"suspensionReasons,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpfilestoreinstance;gcpfilestoreinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// FilestoreInstance is the Schema for the FilestoreInstance API
// +k8s:openapi-gen=true
type FilestoreInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FilestoreInstanceSpec   `json:"spec,omitempty"`
	Status FilestoreInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// FilestoreInstanceList contains a list of FilestoreInstance
type FilestoreInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FilestoreInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FilestoreInstance{}, &FilestoreInstanceList{})
}
