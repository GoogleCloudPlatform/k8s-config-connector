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

// API sources for SecureSourceManagerIssue, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/securesourcemanager/v1/secure_source_manager.proto
// +kcc:source:service-docs=https://cloud.google.com/secure-source-manager
// +kcc:source:resource-docs=https://docs.cloud.google.com/secure-source-manager/docs/reference/rest/v1/projects.locations.repositories.issues

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecureSourceManagerIssueGVK = GroupVersion.WithKind("SecureSourceManagerIssue")

// SecureSourceManagerIssueSpec defines the desired state of SecureSourceManagerIssue
// +kcc:spec:proto=google.cloud.securesourcemanager.v1.Issue
// +kcc:required-from-proto
type SecureSourceManagerIssueSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The SecureSourceManagerIssue name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Issue title.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.Issue.title
	// +required
	Title *string `json:"title,omitempty"`

	// Optional. Issue body. Provides a detailed description of the issue.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.Issue.body
	Body *string `json:"body,omitempty"`

	// Optional. This checksum is computed by the server based on the value of
	//  other fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.Issue.etag
	Etag *string `json:"etag,omitempty"`
}

// SecureSourceManagerIssueStatus defines the config connector machine state of SecureSourceManagerIssue
type SecureSourceManagerIssueStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecureSourceManagerIssue resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecureSourceManagerIssueObservedState `json:"observedState,omitempty"`
}

// SecureSourceManagerIssueObservedState is the state of the SecureSourceManagerIssue resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.securesourcemanager.v1.Issue
type SecureSourceManagerIssueObservedState struct {
	// Output only. State of the issue.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.Issue.state
	State *string `json:"state,omitempty"`

	// Output only. Creation timestamp.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.Issue.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last updated timestamp.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.Issue.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Close timestamp (if closed). Cleared when is re-opened.
	// +kcc:proto:field=google.cloud.securesourcemanager.v1.Issue.close_time
	CloseTime *string `json:"closeTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuresourcemanagerissue;gcpsecuresourcemanagerissues
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecureSourceManagerIssue is the Schema for the SecureSourceManagerIssue API
// +k8s:openapi-gen=true
type SecureSourceManagerIssue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SecureSourceManagerIssueSpec   `json:"spec,omitempty"`
	Status SecureSourceManagerIssueStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecureSourceManagerIssueList contains a list of SecureSourceManagerIssue
type SecureSourceManagerIssueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecureSourceManagerIssue `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecureSourceManagerIssue{}, &SecureSourceManagerIssueList{})
}
