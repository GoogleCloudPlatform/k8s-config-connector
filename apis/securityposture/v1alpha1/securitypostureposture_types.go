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

// API sources for SecurityposturePosture, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/securityposture/v1/securityposture.proto
// +kcc:source:service-docs=https://cloud.google.com/security-command-center
// +kcc:source:resource-docs=https://docs.cloud.google.com/security-command-center/docs/reference/securityposture/rest/v1/organizations.locations.postures

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecurityposturePostureGVK = GroupVersion.WithKind("SecurityposturePosture")

// SecurityposturePostureSpec defines the desired state of SecurityposturePosture
// +kcc:spec:proto=google.cloud.securityposture.v1.Posture
// +kcc:required-from-proto
type SecurityposturePostureSpec struct {
	// The organization that this resource belongs to.
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The SecurityposturePosture name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. State of Posture resource.
	// +kcc:proto:field=google.cloud.securityposture.v1.Posture.state
	// +required
	State *string `json:"state,omitempty"`

	// Optional. User provided description of the posture.
	// +kcc:proto:field=google.cloud.securityposture.v1.Posture.description
	Description *string `json:"description,omitempty"`

	// Required. List of Policy sets.
	// +kcc:proto:field=google.cloud.securityposture.v1.Posture.policy_sets
	// +required
	PolicySets []PolicySet `json:"policySets,omitempty"`

	// Optional. An opaque tag indicating the current version of the Posture, used
	//  for concurrency control. When the `Posture` is returned from either a
	//  `GetPosture` or a `ListPostures` request, this `etag` indicates the version
	//  of the current `Posture` to use when executing a read-modify-write loop.
	//
	//  When the `Posture` is used in a `UpdatePosture` method, use the `etag`
	//  value that was returned from a `GetPosture` request as part of a
	//  read-modify-write loop for concurrency control. Not setting the `etag` in a
	//  `UpdatePosture` request will result in an unconditional write of the
	//  `Posture`.
	// +kcc:proto:field=google.cloud.securityposture.v1.Posture.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. User annotations. These attributes can only be set and used by
	//  the user, and not by Google Security Postures.
	//  .
	// +kcc:proto:field=google.cloud.securityposture.v1.Posture.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// SecurityposturePostureStatus defines the config connector machine state of SecurityposturePosture
type SecurityposturePostureStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecurityposturePosture resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecurityposturePostureObservedState `json:"observedState,omitempty"`
}

// SecurityposturePostureObservedState is the state of the SecurityposturePosture resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.securityposture.v1.Posture
type SecurityposturePostureObservedState struct {
	// Output only. Immutable. The revision ID of the posture.
	//  The format is an 8-character hexadecimal string.
	//  https://google.aip.dev/162
	// +kcc:proto:field=google.cloud.securityposture.v1.Posture.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. The timestamp that the posture was created.
	// +kcc:proto:field=google.cloud.securityposture.v1.Posture.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp that the posture was updated.
	// +kcc:proto:field=google.cloud.securityposture.v1.Posture.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Required. List of Policy sets.
	// +kcc:proto:field=google.cloud.securityposture.v1.Posture.policy_sets
	PolicySets []PolicySetObservedState `json:"policySets,omitempty"`

	// Output only. Whether or not this Posture is in the process of being
	//  updated.
	// +kcc:proto:field=google.cloud.securityposture.v1.Posture.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuritypostureposture;gcpsecurityposturepostures
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecurityposturePosture is the Schema for the SecurityposturePosture API
// +k8s:openapi-gen=true
type SecurityposturePosture struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SecurityposturePostureSpec   `json:"spec,omitempty"`
	Status SecurityposturePostureStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecurityposturePostureList contains a list of SecurityposturePosture
type SecurityposturePostureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityposturePosture `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecurityposturePosture{}, &SecurityposturePostureList{})
}
