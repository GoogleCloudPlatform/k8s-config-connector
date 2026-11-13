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

// API sources for MemorystoreACLPolicy, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/97f6460c3558c1435d7fd001a9251f8df058ae2f/google/cloud/memorystore/v1/memorystore.proto
// +kcc:source:service-docs=https://cloud.google.com/memorystore/docs/valkey
// +kcc:source:resource-docs=https://docs.cloud.google.com/memorystore/docs/valkey/reference/rest/v1/projects.locations.aclPolicies

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MemorystoreACLPolicyGVK = GroupVersion.WithKind("MemorystoreACLPolicy")

// MemorystoreACLPolicySpec defines the desired state of MemorystoreACLPolicy
// +kcc:spec:proto=google.cloud.memorystore.v1.AclPolicy
// +kcc:required-from-proto
type MemorystoreACLPolicySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The MemorystoreACLPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The ACL rules within the ACL policy.
	// +kcc:proto:field=google.cloud.memorystore.v1.AclPolicy.rules
	// +required
	Rules []AclRule `json:"rules,omitempty"`
}

// MemorystoreACLPolicyStatus defines the config connector machine state of MemorystoreACLPolicy
type MemorystoreACLPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MemorystoreACLPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MemorystoreACLPolicyObservedState `json:"observedState,omitempty"`
}

// MemorystoreACLPolicyObservedState is the state of the MemorystoreACLPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.memorystore.v1.AclPolicy
type MemorystoreACLPolicyObservedState struct {
	// Output only. The state of the ACL policy.
	// +kcc:proto:field=google.cloud.memorystore.v1.AclPolicy.state
	State *string `json:"state,omitempty"`

	// Output only. Etag of the ACL policy.
	// +kcc:proto:field=google.cloud.memorystore.v1.AclPolicy.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The ACL policy attachment status for each attached instance.
	// +kcc:proto:field=google.cloud.memorystore.v1.AclPolicy.instance_acl_policy_attachments
	InstanceAclPolicyAttachments []InstanceAclPolicyAttachmentObservedState `json:"instanceAclPolicyAttachments,omitempty"`

	// Output only. The timestamp that the ACL policy was created.
	// +kcc:proto:field=google.cloud.memorystore.v1.AclPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp that the ACL policy was last updated.
	// +kcc:proto:field=google.cloud.memorystore.v1.AclPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmemorystoreaclpolicy;gcpmemorystoreaclpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MemorystoreACLPolicy is the Schema for the MemorystoreACLPolicy API
// +k8s:openapi-gen=true
type MemorystoreACLPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MemorystoreACLPolicySpec   `json:"spec,omitempty"`
	Status MemorystoreACLPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MemorystoreACLPolicyList contains a list of MemorystoreACLPolicy
type MemorystoreACLPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MemorystoreACLPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MemorystoreACLPolicy{}, &MemorystoreACLPolicyList{})
}
