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

// API sources for SecuresourcemanagerPullRequestComment, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/securesourcemanager/v1/secure_source_manager.proto
// +kcc:source:service-docs=https://cloud.google.com/secure-source-manager
// +kcc:source:resource-docs=https://docs.cloud.google.com/secure-source-manager/docs/reference/rest/v1/projects.locations.repositories.pullRequests.pullRequestComments

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecuresourcemanagerPullRequestCommentGVK = GroupVersion.WithKind("SecuresourcemanagerPullRequestComment")

// SecuresourcemanagerPullRequestCommentSpec defines the desired state of SecuresourcemanagerPullRequestComment
// +kcc:spec:proto=google.cloud.securesourcemanager.v1.PullRequestComment
// +kcc:required-from-proto
type SecuresourcemanagerPullRequestCommentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/repositories/{repository}/pullRequests/{pull_request} this resource belongs to.
	// +kcc:guess
	// PullRequestRef *SecuresourcemanagerPullRequestRef `json:"pullRequestRef,omitempty"`

	// The SecuresourcemanagerPullRequestComment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The review summary comment.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequestComment.review
	Review *PullRequestComment_Review `json:"review,omitempty"`

	// Optional. The general pull request comment.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequestComment.comment
	Comment *PullRequestComment_Comment `json:"comment,omitempty"`

	// Optional. The comment on a code line.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequestComment.code
	Code *PullRequestComment_Code `json:"code,omitempty"`
}

// SecuresourcemanagerPullRequestCommentStatus defines the config connector machine state of SecuresourcemanagerPullRequestComment
type SecuresourcemanagerPullRequestCommentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecuresourcemanagerPullRequestComment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecuresourcemanagerPullRequestCommentObservedState `json:"observedState,omitempty"`
}

// SecuresourcemanagerPullRequestCommentObservedState is the state of the SecuresourcemanagerPullRequestComment resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.securesourcemanager.v1.PullRequestComment
type SecuresourcemanagerPullRequestCommentObservedState struct {
	// Output only. Creation timestamp.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequestComment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last updated timestamp.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequestComment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. The review summary comment.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequestComment.review
	Review *PullRequestComment_ReviewObservedState `json:"review,omitempty"`

	// Optional. The comment on a code line.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequestComment.code
	Code *PullRequestComment_CodeObservedState `json:"code,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuresourcemanagerpullrequestcomment;gcpsecuresourcemanagerpullrequestcomments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecuresourcemanagerPullRequestComment is the Schema for the SecuresourcemanagerPullRequestComment API
// +k8s:openapi-gen=true
type SecuresourcemanagerPullRequestComment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SecuresourcemanagerPullRequestCommentSpec   `json:"spec,omitempty"`
	Status SecuresourcemanagerPullRequestCommentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecuresourcemanagerPullRequestCommentList contains a list of SecuresourcemanagerPullRequestComment
type SecuresourcemanagerPullRequestCommentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecuresourcemanagerPullRequestComment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecuresourcemanagerPullRequestComment{}, &SecuresourcemanagerPullRequestCommentList{})
}
