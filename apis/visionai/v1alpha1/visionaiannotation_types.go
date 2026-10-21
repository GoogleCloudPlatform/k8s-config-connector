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

// API sources for VisionAIAnnotation, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/visionai/v1/warehouse.proto
// +kcc:source:service-docs=https://cloud.google.com/vision-ai/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/vision-ai/docs/reference/rest/v1/projects.locations.corpora.assets.annotations

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VisionAIAnnotationGVK = GroupVersion.WithKind("VisionAIAnnotation")

// VisionAIAnnotationSpec defines the desired state of VisionAIAnnotation
// +kcc:spec:proto=google.cloud.visionai.v1.Annotation
// +kcc:required-from-proto
type VisionAIAnnotationSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The VisionAIAnnotation name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// User provided annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.Annotation.user_specified_annotation
	UserSpecifiedAnnotation *UserSpecifiedAnnotation `json:"userSpecifiedAnnotation,omitempty"`
}

// VisionAIAnnotationStatus defines the config connector machine state of VisionAIAnnotation
type VisionAIAnnotationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VisionAIAnnotation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *VisionAIAnnotationObservedState `json:"observedState,omitempty"`
}

// VisionAIAnnotationObservedState is the state of the VisionAIAnnotation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.visionai.v1.Annotation
// type VisionAIAnnotationObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvisionaiannotation;gcpvisionaiannotations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VisionAIAnnotation is the Schema for the VisionAIAnnotation API
// +k8s:openapi-gen=true
type VisionAIAnnotation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VisionAIAnnotationSpec   `json:"spec,omitempty"`
	Status VisionAIAnnotationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VisionAIAnnotationList contains a list of VisionAIAnnotation
type VisionAIAnnotationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VisionAIAnnotation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VisionAIAnnotation{}, &VisionAIAnnotationList{})
}

// +kcc:proto=google.cloud.visionai.v1.AnnotationValue
type AnnotationValue struct {
	// Value of int type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.int_value
	IntValue *int64 `json:"intValue,omitempty"`

	// Value of float type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.float_value
	FloatValue *float32 `json:"floatValue,omitempty"`

	// Value of string type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.str_value
	StrValue *string `json:"strValue,omitempty"`

	// Value of date time type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.datetime_value
	DatetimeValue *string `json:"datetimeValue,omitempty"`

	// Value of geo coordinate type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.geo_coordinate
	GeoCoordinate *GeoCoordinate `json:"geoCoordinate,omitempty"`

	// Value of any proto value.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.proto_any_value
	ProtoAnyValue *Any `json:"protoAnyValue,omitempty"`

	// Value of boolean type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Value of customized struct annotation. This field does not have effects.
	//  Use customized_struct_value instead for customized struct annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.customized_struct_data_value
	CustomizedStructDataValue apiextensionsv1.JSON `json:"customizedStructDataValue,omitempty"`

	// Value of list type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.list_value
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:validation:Type=object
	ListValue *AnnotationList `json:"listValue,omitempty"`

	// Value of custom struct type annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.AnnotationValue.customized_struct_value
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:XPreserveUnknownFields
	// +kubebuilder:validation:Type=object
	CustomizedStructValue *AnnotationCustomizedStruct `json:"customizedStructValue,omitempty"`
}
