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

// API sources for SecuresourcemanagerPullRequest, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/securesourcemanager/v1/secure_source_manager.proto
// +kcc:source:service-docs=https://cloud.google.com/secure-source-manager
// +kcc:source:resource-docs=https://docs.cloud.google.com/secure-source-manager/docs/reference/rest/v1/projects.locations.repositories.pullRequests

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecuresourcemanagerPullRequestGVK = GroupVersion.WithKind("SecuresourcemanagerPullRequest")

// SecuresourcemanagerPullRequestSpec defines the desired state of SecuresourcemanagerPullRequest
// +kcc:spec:proto=google.cloud.securesourcemanager.v1.PullRequest
// +kcc:required-from-proto
type SecuresourcemanagerPullRequestSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The SecuresourcemanagerPullRequest name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The pull request title.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequest.title
	// +required
	Title *string `json:"title,omitempty"`

	// Optional. The pull request body. Provides a detailed description of the
	//  changes.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequest.body
	Body *string `json:"body,omitempty"`

	// Required. The branch to merge changes in.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequest.base
	// +required
	Base *PullRequest_Branch `json:"base,omitempty"`

	// Immutable. The branch containing the changes to be merged.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequest.head
	Head *PullRequest_Branch `json:"head,omitempty"`
}

// SecuresourcemanagerPullRequestStatus defines the config connector machine state of SecuresourcemanagerPullRequest
type SecuresourcemanagerPullRequestStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecuresourcemanagerPullRequest resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecuresourcemanagerPullRequestObservedState `json:"observedState,omitempty"`
}

// SecuresourcemanagerPullRequestObservedState is the state of the SecuresourcemanagerPullRequest resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.securesourcemanager.v1.PullRequest
type SecuresourcemanagerPullRequestObservedState struct {
	// Required. The branch to merge changes in.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequest.base
	Base *PullRequest_BranchObservedState `json:"base,omitempty"`

	// Immutable. The branch containing the changes to be merged.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequest.head
	Head *PullRequest_BranchObservedState `json:"head,omitempty"`

	// Output only. State of the pull request (open, closed or merged).
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequest.state
	State *string `json:"state,omitempty"`

	// Output only. Creation timestamp.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequest.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last updated timestamp.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequest.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Close timestamp (if closed or merged). Cleared when pull
	//  request is re-opened.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.PullRequest.close_time
	CloseTime *string `json:"closeTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuresourcemanagerpullrequest;gcpsecuresourcemanagerpullrequests
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecuresourcemanagerPullRequest is the Schema for the SecuresourcemanagerPullRequest API
// +k8s:openapi-gen=true
type SecuresourcemanagerPullRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SecuresourcemanagerPullRequestSpec   `json:"spec,omitempty"`
	Status SecuresourcemanagerPullRequestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecuresourcemanagerPullRequestList contains a list of SecuresourcemanagerPullRequest
type SecuresourcemanagerPullRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecuresourcemanagerPullRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecuresourcemanagerPullRequest{}, &SecuresourcemanagerPullRequestList{})
}
