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

package v1beta1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var OSConfigOSPolicyAssignmentGVK = GroupVersion.WithKind("OSConfigOSPolicyAssignment")

// OSConfigOSPolicyAssignmentSpec defines the desired state of OSConfigOSPolicyAssignment
// +kcc:spec:proto=google.cloud.osconfig.v1.OSPolicyAssignment
type OSConfigOSPolicyAssignmentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The OSConfigOSPolicyAssignment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// OS policy assignment description. Length of the description is limited to 1024 characters.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.description
	Description *string `json:"description,omitempty"`

	// Required. List of OS policies to be applied to the VMs.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.os_policies
	OSPolicies []OSPolicy `json:"osPolicies"`

	// Required. Filter to select VMs.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.instance_filter
	InstanceFilter *OSPolicyAssignment_InstanceFilter `json:"instanceFilter"`

	// Required. Rollout to deploy the OS policy assignment.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.rollout
	Rollout *OSPolicyAssignment_Rollout `json:"rollout"`

	// Set to true to skip awaiting rollout during resource creation and update.
	SkipAwaitRollout *bool `json:"skipAwaitRollout,omitempty"`
}

// OSConfigOSPolicyAssignmentStatus defines the config connector machine state of OSConfigOSPolicyAssignment
// +kcc:status:proto=google.cloud.osconfig.v1.OSPolicyAssignment
type OSConfigOSPolicyAssignmentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Output only. Indicates that this revision has been successfully rolled out in this zone and new VMs will be assigned OS policies from this revision. For a given OS policy assignment, there is only one revision with a value of `true` for this field.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.baseline
	Baseline *bool `json:"baseline,omitempty"`

	// Output only. Indicates that this revision deletes the OS policy assignment.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.deleted
	Deleted *bool `json:"deleted,omitempty"`

	// The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. Indicates that reconciliation is in progress for the revision. This value is `true` when the `rollout_state` is one of: * IN_PROGRESS * CANCELLING
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The timestamp that the revision was created.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`

	// Output only. The assignment revision ID A new revision is committed whenever a rollout is triggered for a OS policy assignment
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. OS policy assignment rollout state Possible values: ROLLOUT_STATE_UNSPECIFIED, IN_PROGRESS, CANCELLING, CANCELLED, SUCCEEDED
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.rollout_state
	RolloutState *string `json:"rolloutState,omitempty"`

	// Output only. Server generated unique id for the OS policy assignment resource.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.uid
	Uid *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcposconfigospolicyassignment;gcposconfigospolicyassignments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// OSConfigOSPolicyAssignment is the Schema for the OSConfigOSPolicyAssignment API
// +k8s:openapi-gen=true
type OSConfigOSPolicyAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   OSConfigOSPolicyAssignmentSpec   `json:"spec,omitempty"`
	Status OSConfigOSPolicyAssignmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// OSConfigOSPolicyAssignmentList contains a list of OSConfigOSPolicyAssignment
type OSConfigOSPolicyAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OSConfigOSPolicyAssignment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OSConfigOSPolicyAssignment{}, &OSConfigOSPolicyAssignmentList{})
}
