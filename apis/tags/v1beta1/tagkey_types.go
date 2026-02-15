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

	// "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var TagsTagKeyGVK = GroupVersion.WithKind("TagsTagKey")

// TagsTagKeySpec defines the desired state of TagsTagKey
// +kcc:spec:proto=google.cloud.resourcemanager.v3.TagKey
type TagsTagKeySpec struct {
	// // Required. Defines the parent path of the resource.
	// *parent.ProjectAndLocationRef `json:",inline"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The resource name of the TagKey's parent. A TagKey can be
	//  parented by an Organization or a Project. For a TagKey parented by an
	//  Organization, its parent must be in the form `organizations/{org_id}`. For
	//  a TagKey parented by a Project, its parent can be in the form
	//  `projects/{project_id}` or `projects/{project_number}`.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.parent
	// +required
	Parent *string `json:"parent,omitempty"`

	// Required. Immutable. The user friendly name for a TagKey. The short name
	//  should be unique for TagKeys within the same tag namespace.
	//
	//  The short name must be 1-63 characters, beginning and ending with
	//  an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_),
	//  dots (.), and alphanumerics between.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.short_name
	// +required
	ShortName *string `json:"shortName,omitempty"`

	// Optional. User-assigned description of the TagKey. Must not exceed 256
	//  characters.
	//
	//  Read-write.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.description
	Description *string `json:"description,omitempty"`

	// Optional. A purpose denotes that this Tag is intended for use in policies
	//  of a specific policy engine, and will involve that policy engine in
	//  management operations involving this Tag. A purpose does not grant a
	//  policy engine exclusive rights to the Tag, and it may be referenced by
	//  other policy engines.
	//
	//  A purpose cannot be changed once set.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.purpose
	Purpose *string `json:"purpose,omitempty"`

	// Optional. Purpose data corresponds to the policy system that the tag is
	//  intended for. See documentation for `Purpose` for formatting of this field.
	//
	//  Purpose data cannot be changed once set.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.purpose_data
	PurposeData map[string]string `json:"purposeData,omitempty"`
}

// TagsTagKeyStatus defines the config connector machine state of TagsTagKey
// +kcc:status:proto=google.cloud.resourcemanager.v3.TagKey
type TagsTagKeyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TagsTagKey resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// NOTYET: Stay compatible with terraform
	// // ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *TagsTagKeyObservedState `json:"observedState,omitempty"`

	// Output only. Creation time.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Immutable. Namespaced name of the TagKey.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.namespaced_name
	NamespacedName *string `json:"namespacedName,omitempty"`

	// The generated numeric id for the TagKey.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.name
	Name *string `json:"name,omitempty"`
}

// TagsTagKeyObservedState is the state of the TagsTagKey resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.resourcemanager.v3.TagKey
type TagsTagKeyObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptagstagkey;gcptagstagkeys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TagsTagKey is the Schema for the TagsTagKey API
// +k8s:openapi-gen=true
type TagsTagKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TagsTagKeySpec   `json:"spec,omitempty"`
	Status TagsTagKeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TagsTagKeyList contains a list of TagsTagKey
type TagsTagKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TagsTagKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TagsTagKey{}, &TagsTagKeyList{})
}
