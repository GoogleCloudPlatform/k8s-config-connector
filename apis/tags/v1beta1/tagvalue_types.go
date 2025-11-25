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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var TagsTagValueGVK = GroupVersion.WithKind("TagsTagValue")

// TagsTagValueSpec defines the desired state of TagsTagValue
// +kcc:spec:proto=google.cloud.resourcemanager.v3.TagValue
type TagsTagValueSpec struct {
	// // Required. Defines the parent path of the resource.
	// *parent.ProjectAndLocationRef `json:",inline"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The TagValue's parent TagKey.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.parent
	// +required
	ParentRef *TagsTagKeyRef `json:"parentRef,omitempty"`

	// Required. Immutable. User-assigned short name for TagValue. The short name
	//  should be unique for TagValues within the same parent TagKey.
	//
	//  The short name must be 63 characters or less, beginning and ending with
	//  an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_),
	//  dots (.), and alphanumerics between.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.short_name
	// +required
	ShortName *string `json:"shortName,omitempty"`

	// Optional. User-assigned description of the TagValue.
	//  Must not exceed 256 characters.
	//
	//  Read-write.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.description
	Description *string `json:"description,omitempty"`
}

// TagsTagValueStatus defines the config connector machine state of TagsTagValue
// +kcc:status:proto=google.cloud.resourcemanager.v3.TagValue
type TagsTagValueStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TagsTagValue resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// NOTYET: TERRAFORM COMPATIBILITY
	// // ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *TagsTagValueObservedState `json:"observedState,omitempty"`

	// Immutable. Resource name for TagValue in the format `tagValues/456`.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.name
	Name *string `json:"name,omitempty"`

	// Output only. The namespaced name of the TagValue. Can be in the form
	//  `{organization_id}/{tag_key_short_name}/{tag_value_short_name}` or
	//  `{project_id}/{tag_key_short_name}/{tag_value_short_name}` or
	//  `{project_number}/{tag_key_short_name}/{tag_value_short_name}`.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.namespaced_name
	NamespacedName *string `json:"namespacedName,omitempty"`

	// Output only. Creation time.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// TagsTagValueObservedState is the state of the TagsTagValue resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.resourcemanager.v3.TagValue
type TagsTagValueObservedState struct {
	// NOTYET: ETAG
	// // Optional. Entity tag which users can pass to prevent race conditions. This
	// //  field is always set in server responses. See UpdateTagValueRequest for
	// //  details.
	// // +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.etag
	// Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptagstagvalue;gcptagstagvalues
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TagsTagValue is the Schema for the TagsTagValue API
// +k8s:openapi-gen=true
type TagsTagValue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TagsTagValueSpec   `json:"spec,omitempty"`
	Status TagsTagValueStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TagsTagValueList contains a list of TagsTagValue
type TagsTagValueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TagsTagValue `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TagsTagValue{}, &TagsTagValueList{})
}
