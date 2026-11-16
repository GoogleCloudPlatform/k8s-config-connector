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

// API sources for BackupDRAutoProtectionPolicy, recorded by generate-types:
// +kcc:source:proto=https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/mockgcp/apis/google/cloud/backupdr/v1/autoprotectionpolicy.proto
// +kcc:source:service-docs=https://cloud.google.com/backup-disaster-recovery
// +kcc:source:resource-docs=https://docs.cloud.google.com/backup-disaster-recovery/docs/reference/rest/v1/projects.locations.autoProtectionPolicies

package v1alpha1

import (
	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BackupDRAutoProtectionPolicyGVK = GroupVersion.WithKind("BackupDRAutoProtectionPolicy")

// BackupDRAutoProtectionPolicySpec defines the desired state of BackupDRAutoProtectionPolicy
// +kcc:spec:proto=google.cloud.backupdr.v1.AutoProtectionPolicy
// +kcc:required-from-proto
type BackupDRAutoProtectionPolicySpec struct {
	*parent.ProjectAndLocationRef `json:",inline"`

	// The BackupDRAutoProtectionPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. A description of the policy. The description must be at most 2048
	//  characters.
	// +kcc:proto:field=google.cloud.backupdr.v1.AutoProtectionPolicy.description
	Description *string `json:"description,omitempty"`

	// Required. The backup plan details for matching resources.
	// +kcc:proto:field=google.cloud.backupdr.v1.AutoProtectionPolicy.backup_plan_details
	// +required
	BackupPlanDetails []BackupPlanDetail `json:"backupPlanDetails,omitempty"`

	// Required. The criteria for matching resources.
	// +kcc:proto:field=google.cloud.backupdr.v1.AutoProtectionPolicy.criteria
	// +required
	Criteria *Criteria `json:"criteria,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.BackupPlanDetail
type BackupPlanDetail struct {
	// Required. The type of resource that this policy will automatically protect.
	// Allowed values are, `compute.googleapis.com/Instance` and
	// `compute.googleapis.com/Disk` are supported.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupPlanDetail.resource_type
	ResourceType *string `json:"resourceType,omitempty"`

	// Required. The resource name of the BackupPlan to apply to resources that
	// match the policy's criteria. This backup plan must be in the same location
	// as AutoProtectionPolicy. Format:
	// projects/{project}/locations/{location}/backupPlans/{backupPlanId}
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupPlanDetail.backup_plan
	BackupPlanRef *BackupPlanRef `json:"backupPlanRef,omitempty"`
}

// BackupDRAutoProtectionPolicyStatus defines the config connector machine state of BackupDRAutoProtectionPolicy
type BackupDRAutoProtectionPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BackupDRAutoProtectionPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BackupDRAutoProtectionPolicyObservedState `json:"observedState,omitempty"`
}

// BackupDRAutoProtectionPolicyObservedState is the state of the BackupDRAutoProtectionPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.backupdr.v1.AutoProtectionPolicy
type BackupDRAutoProtectionPolicyObservedState struct {
	// Output only. The current state of the AutoProtectionPolicy.
	// +kcc:proto:field=google.cloud.backupdr.v1.AutoProtectionPolicy.state
	State *string `json:"state,omitempty"`

	// Output only. The etag for the policy.
	// If this is provided on update, it must match the server's etag.
	// +kcc:proto:field=google.cloud.backupdr.v1.AutoProtectionPolicy.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. Summary of total count of different containers covered through
	//  this policy.
	// +kcc:proto:field=google.cloud.backupdr.v1.AutoProtectionPolicy.binding_summary
	BindingSummary *BindingSummaryObservedState `json:"bindingSummary,omitempty"`

	// Output only. Summary of last policy evaluation when job was run per
	//  workload type.
	// +kcc:proto:field=google.cloud.backupdr.v1.AutoProtectionPolicy.policy_evaluation_summary
	PolicyEvaluationSummary *PolicyEvaluationSummaryObservedState `json:"policyEvaluationSummary,omitempty"`

	// Output only. The timestamp when the policy was created.
	// +kcc:proto:field=google.cloud.backupdr.v1.AutoProtectionPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the policy was last updated.
	// +kcc:proto:field=google.cloud.backupdr.v1.AutoProtectionPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbackupdrautoprotectionpolicy;gcpbackupdrautoprotectionpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BackupDRAutoProtectionPolicy is the Schema for the BackupDRAutoProtectionPolicy API
// +k8s:openapi-gen=true
type BackupDRAutoProtectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BackupDRAutoProtectionPolicySpec   `json:"spec,omitempty"`
	Status BackupDRAutoProtectionPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BackupDRAutoProtectionPolicyList contains a list of BackupDRAutoProtectionPolicy
type BackupDRAutoProtectionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupDRAutoProtectionPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupDRAutoProtectionPolicy{}, &BackupDRAutoProtectionPolicyList{})
}
