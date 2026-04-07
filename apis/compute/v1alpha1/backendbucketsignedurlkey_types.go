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
	refsecret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeBackendBucketSignedURLKeyGVK = GroupVersion.WithKind("ComputeBackendBucketSignedURLKey")

// BackendBucketRef is a reference to a ComputeBackendBucket resource.
type BackendBucketRef struct {
	// Allowed value: The `name` field of a `ComputeBackendBucket` resource.
	// +optional
	External string `json:"external,omitempty"`
	// Name of the referent.
	// +optional
	Name string `json:"name,omitempty"`
	// Namespace of the referent.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// ComputeBackendBucketSignedURLKeySpec defines the desired state of ComputeBackendBucketSignedURLKey.
type ComputeBackendBucketSignedURLKeySpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The backend bucket this signed URL key is associated with.
	// +required
	BackendBucketRef BackendBucketRef `json:"backendBucketRef"`

	// Immutable. 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// +required
	KeyValue refsecret.Legacy `json:"keyValue"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition.
	// When unset, the value of `metadata.name` is used as the default.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeBackendBucketSignedURLKeyStatus defines the config connector machine state of ComputeBackendBucketSignedURLKey.
type ComputeBackendBucketSignedURLKeyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller.
	// If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeBackendBucketSignedURLKey resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputebackendbucketsignedurlkey;gcpcomputebackendbucketsignedurlkeys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeBackendBucketSignedURLKey is the Schema for the ComputeBackendBucketSignedURLKey API
// +k8s:openapi-gen=true
type ComputeBackendBucketSignedURLKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeBackendBucketSignedURLKeySpec   `json:"spec,omitempty"`
	Status ComputeBackendBucketSignedURLKeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeBackendBucketSignedURLKeyList contains a list of ComputeBackendBucketSignedURLKey
type ComputeBackendBucketSignedURLKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeBackendBucketSignedURLKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeBackendBucketSignedURLKey{}, &ComputeBackendBucketSignedURLKeyList{})
}
