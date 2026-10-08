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
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var OSLoginSSHPublicKeyGVK = GroupVersion.WithKind("OSLoginSSHPublicKey")

// +kcc:spec:proto=google.cloud.oslogin.common.SshPublicKey
type OSLoginSSHPublicKeySpec struct {
	/* An expiration time in microseconds since epoch. */
	// +optional
	// +kcc:proto:field=google.cloud.oslogin.common.SshPublicKey.expiration_time_usec
	ExpirationTimeUsec *string `json:"expirationTimeUsec,omitempty"`

	/* Immutable. Public key text in SSH format, defined by RFC4253 section 6.6. */
	// +required
	// +kcc:proto:field=google.cloud.oslogin.common.SshPublicKey.key
	Key string `json:"key"`

	/* Immutable. The project ID of the Google Cloud Platform project. */
	// +optional
	Project *string `json:"project,omitempty"`

	/* Immutable. Optional. The service-generated fingerprint of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The user email. */
	// +required
	User string `json:"user"`
}

// +kcc:status:proto=google.cloud.oslogin.common.SshPublicKey
type OSLoginSSHPublicKeyStatus struct {
	/* Conditions represent the latest available observation of the
	   resource's current state. */
	// +optional
	Conditions []k8sv1alpha1.Condition `json:"conditions,omitempty"`

	/* The SHA-256 fingerprint of the SSH public key. */
	// +optional
	// +kcc:proto:field=google.cloud.oslogin.common.SshPublicKey.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcposloginsshpublickey;gcposloginsshpublickeys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// OSLoginSSHPublicKey is the Schema for the oslogin API
// +k8s:openapi-gen=true
type OSLoginSSHPublicKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   OSLoginSSHPublicKeySpec   `json:"spec,omitempty"`
	Status OSLoginSSHPublicKeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OSLoginSSHPublicKeyList contains a list of OSLoginSSHPublicKey
type OSLoginSSHPublicKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OSLoginSSHPublicKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OSLoginSSHPublicKey{}, &OSLoginSSHPublicKeyList{})
}
