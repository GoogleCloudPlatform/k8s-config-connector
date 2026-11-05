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

// API sources for VisionReferenceImage, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/vision/v1/product_search_service.proto
// +kcc:source:service-docs=https://cloud.google.com/vision/
// +kcc:source:resource-docs=https://docs.cloud.google.com/vision/docs/reference/rest/v1/projects.locations.products.referenceImages

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VisionReferenceImageGVK = GroupVersion.WithKind("VisionReferenceImage")

// VisionReferenceImageSpec defines the desired state of VisionReferenceImage
// +kcc:spec:proto=google.cloud.vision.v1.ReferenceImage
// +kcc:required-from-proto
type VisionReferenceImageSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The VisionReferenceImage name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The Google Cloud Storage URI of the reference image.
	//
	//  The URI must start with `gs://`.
	// +kcc:proto:field=google.cloud.vision.v1.ReferenceImage.uri
	// +required
	URI *string `json:"uri,omitempty"`

	// Optional. Bounding polygons around the areas of interest in the reference
	//  image. If this field is empty, the system will try to detect regions of
	//  interest. At most 10 bounding polygons will be used.
	//
	//  The provided shape is converted into a non-rotated rectangle. Once
	//  converted, the small edge of the rectangle must be greater than or equal
	//  to 300 pixels. The aspect ratio must be 1:4 or less (i.e. 1:3 is ok; 1:5
	//  is not).
	// +kcc:proto:field=google.cloud.vision.v1.ReferenceImage.bounding_polys
	BoundingPolys []BoundingPoly `json:"boundingPolys,omitempty"`
}

// VisionReferenceImageStatus defines the config connector machine state of VisionReferenceImage
type VisionReferenceImageStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VisionReferenceImage resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *VisionReferenceImageObservedState `json:"observedState,omitempty"`
}

// VisionReferenceImageObservedState is the state of the VisionReferenceImage resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.vision.v1.ReferenceImage
// type VisionReferenceImageObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvisionreferenceimage;gcpvisionreferenceimages
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VisionReferenceImage is the Schema for the VisionReferenceImage API
// +k8s:openapi-gen=true
type VisionReferenceImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VisionReferenceImageSpec   `json:"spec,omitempty"`
	Status VisionReferenceImageStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VisionReferenceImageList contains a list of VisionReferenceImage
type VisionReferenceImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VisionReferenceImage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VisionReferenceImage{}, &VisionReferenceImageList{})
}
