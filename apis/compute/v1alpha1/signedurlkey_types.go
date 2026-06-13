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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = apiextensionsv1.JSON{}

var ComputeBackendServiceSignedURLKeyGVK = GroupVersion.WithKind("ComputeBackendServiceSignedURLKey")

type BackendservicesignedurlkeyKeyValue struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *BackendservicesignedurlkeyValueFrom `json:"valueFrom,omitempty"`
}

type BackendservicesignedurlkeyValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *k8sv1alpha1.SecretKeyRef `json:"secretKeyRef,omitempty"`
}

// ComputeBackendServiceSignedURLKeySpec defines the desired state of ComputeBackendServiceSignedURLKey
// +kcc:proto=google.cloud.compute.v1.SignedUrlKey
type ComputeBackendServiceSignedURLKeySpec struct {
	// +required
	BackendServiceRef computev1beta1.ComputeBackendServiceRef `json:"backendServiceRef"`

	/* Immutable. 128-bit key value used for signing the URL. The key value must be a
	valid RFC 4648 Section 5 base64url encoded string. */
	// +required
	KeyValue BackendservicesignedurlkeyKeyValue `json:"keyValue"`

	/* The project that this resource belongs to. */
	// +required
	ProjectRef refsv1beta1.ProjectRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeBackendServiceSignedURLKeyStatus defines the config connector machine state of ComputeBackendServiceSignedURLKey
type ComputeBackendServiceSignedURLKeyStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeBackendServiceSignedURLKey's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputebackendservicesignedurlkey;gcpcomputebackendservicesignedurlkeys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeBackendServiceSignedURLKey is the Schema for the ComputeBackendServiceSignedURLKey API
// +k8s:openapi-gen=true
type ComputeBackendServiceSignedURLKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeBackendServiceSignedURLKeySpec   `json:"spec,omitempty"`
	Status ComputeBackendServiceSignedURLKeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeBackendServiceSignedURLKeyList contains a list of ComputeBackendServiceSignedURLKey
type ComputeBackendServiceSignedURLKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeBackendServiceSignedURLKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeBackendServiceSignedURLKey{}, &ComputeBackendServiceSignedURLKeyList{})
}
