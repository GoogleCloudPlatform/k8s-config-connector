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
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudIdentityGroupGVK = GroupVersion.WithKind("CloudIdentityGroup")

// +kcc:proto=google.apps.cloudidentity.groups.v1beta1.EntityKey
type EntityKey struct {
	// Immutable. The ID of the entity. For Google-managed entities, the `id` must be the email address of an existing group or user. For external-identity-mapped entities, the `id` must be a string conforming to the Identity Source's requirements. Must be unique within a `namespace`.
	// +required
	ID string `json:"id"`

	// Immutable. The namespace in which the entity exists. If not specified, the `EntityKey` represents a Google-managed entity such as a Google user or a Google Group. If specified, the `EntityKey` represents an external-identity-mapped group. The namespace must correspond to an identity source created in Admin Console and must be in the form of `identitysources/{identity_source_id}`.
	Namespace *string `json:"namespace,omitempty"`
}

// CloudIdentityGroupSpec defines the desired state of CloudIdentityGroup
// +kcc:proto=google.apps.cloudidentity.groups.v1beta1.Group
type CloudIdentityGroupSpec struct {
	// An extended description to help users determine the purpose of a `Group`. Must not be longer than 4,096 characters.
	Description *string `json:"description,omitempty"`

	// The display name of the `Group`.
	DisplayName *string `json:"displayName,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="GroupKey field is immutable"
	// Immutable. EntityKey of the Group.
	// +required
	GroupKey *EntityKey `json:"groupKey"`

	// Immutable. The initial configuration options for creating a Group. See the [API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig) for possible values. Default value: "EMPTY" Possible values: ["INITIAL_GROUP_CONFIG_UNSPECIFIED", "WITH_INITIAL_OWNER", "EMPTY"].
	InitialGroupConfig *string `json:"initialGroupConfig,omitempty"`

	// One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value. Google Groups are the default type of group and have a label with a key of cloudidentity.googleapis.com/groups.discussion_forum and an empty value. Existing Google Groups can have an additional label with a key of cloudidentity.googleapis.com/groups.security and an empty value added to them. This is an immutable change and the security label cannot be removed once added. Dynamic groups have a label with a key of cloudidentity.googleapis.com/groups.dynamic. Identity-mapped groups for Cloud Search have a label with a key of system/groups/external and an empty value.
	// +required
	Labels map[string]string `json:"labels"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Parent field is immutable"
	// Immutable. The resource name of the entity under which this `Group` resides in the Cloud Identity resource hierarchy. Must be of the form `identitysources/{identity_source}` for external [identity-mapped groups](https://support.google.com/a/answer/9039510) or `customers/{customer_id}` for Google Groups. The `customer_id` must begin with "C" (for example, 'C046psxkn'). [Find your customer ID.] (https://support.google.com/cloudidentity/answer/10070793)
	// +required
	Parent *string `json:"parent,omitempty"`

	// Immutable. The CloudIdentityGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// CloudIdentityGroupStatus defines the config connector machine state of CloudIdentityGroup
// +kcc:proto=google.apps.cloudidentity.groups.v1beta1.Group
type CloudIdentityGroupStatus struct {
	commonv1alpha1.CommonStatus `json:",inline"`

	// The time when the `Group` was created.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.Group.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Group`. Shall be of the form `groups/{group_id}`.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.Group.name
	Name *string `json:"name,omitempty"`

	// The time when the `Group` was last updated.
	// +kcc:proto:field=mockgcp.cloud.cloudidentity.groups.v1beta1.Group.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudIdentityGroupObservedState `json:"observedState,omitempty"`
}

// CloudIdentityGroupObservedState is the state of the CloudIdentityGroup resource as most recently observed in GCP.
// +kcc:proto=google.apps.cloudidentity.groups.v1beta1.Group
type CloudIdentityGroupObservedState struct {
	// Additional group keys associated with the Group.
	AdditionalGroupKeys []EntityKey `json:"additionalGroupKeys,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudidentitygroup;gcpcloudidentitygroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudIdentityGroup is the Schema for the CloudIdentityGroup API
// +k8s:openapi-gen=true
type CloudIdentityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudIdentityGroupSpec   `json:"spec,omitempty"`
	Status CloudIdentityGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudIdentityGroupList contains a list of CloudIdentityGroup
type CloudIdentityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudIdentityGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudIdentityGroup{}, &CloudIdentityGroupList{})
}
