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

// API sources for SecureSourceManagerBranchRule, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/securesourcemanager/v1/secure_source_manager.proto
// +kcc:source:service-docs=https://cloud.google.com/secure-source-manager
// +kcc:source:resource-docs=https://docs.cloud.google.com/secure-source-manager/docs/reference/rest/v1/projects.locations.repositories.branchRules

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecureSourceManagerBranchRuleGVK = GroupVersion.WithKind("SecureSourceManagerBranchRule")

// SecureSourceManagerBranchRuleSpec defines the desired state of SecureSourceManagerBranchRule
// +kcc:spec:proto=google.cloud.securesourcemanager.v1.BranchRule
// +kcc:required-from-proto
type SecureSourceManagerBranchRuleSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The SecureSourceManagerBranchRule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. User annotations. These attributes can only be set and used by
	//  the user. See https://google.aip.dev/128#annotations for more details such
	//  as format and size limitations.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. This checksum is computed by the server based on the value of
	//  other fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The pattern of the branch that can match to this BranchRule.
	//  Specified as regex.
	//  .* for all branches. Examples: main, (main|release.*).
	//  Current MVP phase only support `.*` for wildcard.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.include_pattern
	IncludePattern *string `json:"includePattern,omitempty"`

	// Optional. Determines if the branch rule is disabled or not.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Optional. Determines if the branch rule requires a pull request or not.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.require_pull_request
	RequirePullRequest *bool `json:"requirePullRequest,omitempty"`

	// Optional. The minimum number of reviews required for the branch rule to be
	//  matched.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.minimum_reviews_count
	MinimumReviewsCount *int32 `json:"minimumReviewsCount,omitempty"`

	// Optional. The minimum number of approvals required for the branch rule to
	//  be matched.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.minimum_approvals_count
	MinimumApprovalsCount *int32 `json:"minimumApprovalsCount,omitempty"`

	// Optional. Determines if require comments resolved before merging to the
	//  branch.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.require_comments_resolved
	RequireCommentsResolved *bool `json:"requireCommentsResolved,omitempty"`

	// Optional. Determines if allow stale reviews or approvals before merging to
	//  the branch.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.allow_stale_reviews
	AllowStaleReviews *bool `json:"allowStaleReviews,omitempty"`

	// Optional. Determines if require linear history before merging to the
	//  branch.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.require_linear_history
	RequireLinearHistory *bool `json:"requireLinearHistory,omitempty"`

	// Optional. List of required status checks before merging to the branch.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.required_status_checks
	RequiredStatusChecks []BranchRule_Check `json:"requiredStatusChecks,omitempty"`
}

// SecureSourceManagerBranchRuleStatus defines the config connector machine state of SecureSourceManagerBranchRule
type SecureSourceManagerBranchRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecureSourceManagerBranchRule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecureSourceManagerBranchRuleObservedState `json:"observedState,omitempty"`
}

// SecureSourceManagerBranchRuleObservedState is the state of the SecureSourceManagerBranchRule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.securesourcemanager.v1.BranchRule
type SecureSourceManagerBranchRuleObservedState struct {
	// Output only. Unique identifier of the repository.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Create timestamp.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update timestamp.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.BranchRule.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuresourcemanagerbranchrule;gcpsecuresourcemanagerbranchrules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecureSourceManagerBranchRule is the Schema for the SecureSourceManagerBranchRule API
// +k8s:openapi-gen=true
type SecureSourceManagerBranchRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SecureSourceManagerBranchRuleSpec   `json:"spec,omitempty"`
	Status SecureSourceManagerBranchRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecureSourceManagerBranchRuleList contains a list of SecureSourceManagerBranchRule
type SecureSourceManagerBranchRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecureSourceManagerBranchRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecureSourceManagerBranchRule{}, &SecureSourceManagerBranchRuleList{})
}
