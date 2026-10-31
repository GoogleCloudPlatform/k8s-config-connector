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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeBackendBucketSignedURLKeyGVK = GroupVersion.WithKind("ComputeBackendBucketSignedURLKey")

// ComputeBackendBucketSignedURLKeySpec defines the desired state of ComputeBackendBucketSignedURLKey
// +kcc:proto=google.cloud.compute.v1.SignedUrlKey
type ComputeBackendBucketSignedURLKeySpec struct {
	BackendBucketRef refsv1beta1.ComputeBackendBucketRef `json:"backendBucketRef"`

	/* Immutable. 128-bit key value used for signing the URL. The key value must be a
	valid RFC 4648 Section 5 base64url encoded string. */
	KeyValue refsv1beta1secret.Legacy `json:"keyValue"`

	/* The project that this resource belongs to. */
	ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeBackendBucketSignedURLKeyStatus defines the config connector machine state of ComputeBackendBucketSignedURLKey
type ComputeBackendBucketSignedURLKeyStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeBackendBucketSignedURLKey's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputebackendbucketsignedurlkey;gcpcomputebackendbucketsignedurlkeys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeBackendBucketSignedURLKey is the Schema for the compute API
// +k8s:openapi-gen=true
// +kcc:proto=google.cloud.compute.v1.SignedUrlKey
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
