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

var VisionProductGVK = GroupVersion.WithKind("VisionProduct")

// +kcc:proto=google.cloud.vision.v1.Product.KeyValue
type ProductKeyValue struct {
	// The key of the label attached to the product. Cannot be empty and cannot
	//  exceed 128 bytes.
	// +kubebuilder:validation:Required
	Key *string `json:"key"`

	// The value of the label attached to the product. Cannot be empty and
	//  cannot exceed 128 bytes.
	// +kubebuilder:validation:Required
	Value *string `json:"value"`
}

// VisionProductSpec defines the desired state of VisionProduct
// +kcc:spec:proto=google.cloud.vision.v1.Product
type VisionProductSpec struct {
	// The project that this resource belongs to.
	// +kubebuilder:validation:Required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The VisionProduct name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// The user-provided name for this Product. Must not be empty. Must be at most
	//  4096 characters long.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName"`

	// User-provided metadata to be stored with this product. Must be at most 4096
	//  characters long.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Immutable. The category for the product identified by the reference image.
	//  This should be one of "homegoods-v2", "apparel-v2", "toys-v2",
	//  "packagedgoods-v1" or "general-v1". The legacy categories "homegoods",
	//  "apparel", and "toys" are still supported, but these should not be used for
	//  new products.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ProductCategory field is immutable"
	// +kubebuilder:validation:Enum=homegoods-v2;apparel-v2;toys-v2;packagedgoods-v1;general-v1;homegoods;apparel;toys
	ProductCategory *string `json:"productCategory"`

	// Key-value pairs that can be attached to a product. At query time,
	//  constraints can be specified based on the product_labels.
	//
	//  Note that integer values can be provided as strings, e.g. "1199". Only
	//  strings with integer values can match a range-based restriction which is
	//  to be supported soon.
	//
	//  Multiple values can be assigned to the same key. One product may have up to
	//  500 product_labels.
	//
	//  Notice that the total number of distinct product_labels over all products
	//  in one ProductSet cannot exceed 1M, otherwise the product search pipeline
	//  will refuse to work for that ProductSet.
	// +kubebuilder:validation:Optional
	ProductLabels []ProductKeyValue `json:"productLabels,omitempty"`
}

// VisionProductStatus defines the config connector machine state of VisionProduct
type VisionProductStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VisionProduct resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VisionProductObservedState `json:"observedState,omitempty"`
}

// VisionProductObservedState is the state of the VisionProduct resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.vision.v1.Product
// +kubebuilder:validation:XPreserveUnknownFields
type VisionProductObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvisionproduct;gcpvisionproducts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VisionProduct is the Schema for the VisionProduct API
// +k8s:openapi-gen=true
type VisionProduct struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VisionProductSpec   `json:"spec,omitempty"`
	Status VisionProductStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VisionProductList contains a list of VisionProduct
type VisionProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VisionProduct `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VisionProduct{}, &VisionProductList{})
}
