// Copyright 2024 Google LLC
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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var PrivilegedAccessManagerEntitlementGVK = GroupVersion.WithKind("PrivilegedAccessManagerEntitlement")

// PrivilegedAccessManagerEntitlementSpec defines the desired state of
// PrivilegedAccessManagerEntitlement.
// +kcc:spec:proto=google.cloud.privilegedaccessmanager.v1.Entitlement
type PrivilegedAccessManagerEntitlementSpec struct {
	// Immutable. The Project that this resource belongs to.
	// One and only one of 'projectRef', 'folderRef', or 'organizationRef' must
	// be set.
	// +optional
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// Immutable. The Folder that this resource belongs to.
	// One and only one of 'projectRef', 'folderRef', or 'organizationRef' must
	// be set.
	// +optional
	FolderRef *refs.FolderRef `json:"folderRef,omitempty"`

	// Immutable. The Organization that this resource belongs to.
	// One and only one of 'projectRef', 'folderRef', or 'organizationRef' must
	// be set.
	// +optional
	OrganizationRef *refs.OrganizationRef `json:"organizationRef,omitempty"`

	// Immutable. Location of the resource.
	// +required
	Location *string `json:"location"`

	// The PrivilegedAccessManagerEntitlement name. If not given, the
	// 'metadata.name' will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Who can create grants using this entitlement. This list should
	// contain at most one entry.
	// +required
	EligibleUsers []AccessControlEntry `json:"eligibleUsers,omitempty"`

	// Optional. The approvals needed before access are granted to a requester.
	// No approvals are needed if this field is null.
	// +optional
	ApprovalWorkflow *ApprovalWorkflow `json:"approvalWorkflow,omitempty"`

	// Required. The maximum amount of time that access is granted for a request.
	// A requester can ask for a duration less than this, but never more.
	// +required
	MaxRequestDuration *string `json:"maxRequestDuration,omitempty"`

	// The access granted to a requester on successful approval.
	// +required
	PrivilegedAccess *PrivilegedAccess `json:"privilegedAccess,omitempty"`

	// Required. The manner in which the requester should provide a justification
	// for requesting access.
	// +required
	RequesterJustificationConfig *RequesterJustificationConfig `json:"requesterJustificationConfig,omitempty"`

	// Optional. Additional email addresses to be notified based on actions taken.
	// +optional
	AdditionalNotificationTargets *AdditionalNotificationTargets `json:"additionalNotificationTargets,omitempty"`
}

// PrivilegedAccessManagerEntitlementStatus defines the config connector machine
// state of PrivilegedAccessManagerEntitlement.
type PrivilegedAccessManagerEntitlementStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most
	// recently observed by the Config Connector controller. If this is equal to
	// 'metadata.generation', then that means that the current reported status
	// reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the PrivilegedAccessManagerEntitlement resource in
	// GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in
	// GCP.
	// +optional
	ObservedState *PrivilegedAccessManagerEntitlementObservedState `json:"observedState,omitempty"`
}

// PrivilegedAccessManagerEntitlementSpec defines the desired state of
// PrivilegedAccessManagerEntitlement.
// +kcc:observedstate:proto=google.privilegedaccessmanager.v1.Entitlement
type PrivilegedAccessManagerEntitlementObservedState struct {
	// Output only. Create time stamp.
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of this entitlement.
	// +optional
	State *string `json:"state,omitempty"`

	// An 'etag' is used for optimistic concurrency control as a way to prevent
	// simultaneous updates to the same entitlement. An 'etag' is returned in the
	// response to 'GetEntitlement' and the caller should put the 'etag' in the
	// request to 'UpdateEntitlement' so that their change is applied on
	// the same version. If this field is omitted or if there is a mismatch while
	// updating an entitlement, then the server rejects the request.
	// +optional
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpprivilegedaccessmanagerentitlement;gcpprivilegedaccessmanagerentitlements
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// PrivilegedAccessManagerEntitlement is the Schema for the
// PrivilegedAccessManagerEntitlement API.
// +k8s:openapi-gen=true
type PrivilegedAccessManagerEntitlement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrivilegedAccessManagerEntitlementSpec   `json:"spec,omitempty"`
	Status PrivilegedAccessManagerEntitlementStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// PrivilegedAccessManagerEntitlementList contains a list of
// PrivilegedAccessManagerEntitlement.
type PrivilegedAccessManagerEntitlementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivilegedAccessManagerEntitlement `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrivilegedAccessManagerEntitlement{}, &PrivilegedAccessManagerEntitlementList{})
}
