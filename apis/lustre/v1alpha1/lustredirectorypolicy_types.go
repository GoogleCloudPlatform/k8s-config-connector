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

// API sources for LustreDirectoryPolicy, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/6c345acd31e060c46c74cc82c1d28af451bddd86/google/cloud/lustre/v1/directory_policy.proto
// +kcc:source:service-docs=https://cloud.google.com/managed-lustre/docs
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/managed-lustre/docs/reference/rest/v1/projects.locations.instances.directoryPolicies

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var LustreDirectoryPolicyGVK = GroupVersion.WithKind("LustreDirectoryPolicy")

// LustreDirectoryPolicySpec defines the desired state of LustreDirectoryPolicy
// +kcc:spec:proto=google.cloud.lustre.v1.DirectoryPolicy
// +kcc:required-from-proto
type LustreDirectoryPolicySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/instances/{instance} this resource belongs to.
	// +kcc:guess
	// InstanceRef *LustreInstanceRef `json:"instanceRef,omitempty"`

	// The LustreDirectoryPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Immutable. The lustre instance filesystem full path of the
	//  directory. It must start with a slash. e.g. /lustre/testFolder
	// +kcc:proto:field=google.cloud.lustre.v1.DirectoryPolicy.directory_path
	// +required
	DirectoryPath *string `json:"directoryPath,omitempty"`
}

// LustreDirectoryPolicyStatus defines the config connector machine state of LustreDirectoryPolicy
type LustreDirectoryPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the LustreDirectoryPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *LustreDirectoryPolicyObservedState `json:"observedState,omitempty"`
}

// LustreDirectoryPolicyObservedState is the state of the LustreDirectoryPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.lustre.v1.DirectoryPolicy
type LustreDirectoryPolicyObservedState struct {
	// Output only. The lustre project ID assigned for the directory by the
	//  service. This read-only ID can be used to manage quotas. See more details
	//  in
	//  https://docs.cloud.google.com/managed-lustre/docs/file-system-quotas#set_quotas
	// +kcc:proto:field=google.cloud.lustre.v1.DirectoryPolicy.lustre_project_id
	LustreProjectID *int64 `json:"lustreProjectID,omitempty"`

	// Output only. Unique ID of the resource.
	// +kcc:proto:field=google.cloud.lustre.v1.DirectoryPolicy.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The current state of the directory policy.
	// +kcc:proto:field=google.cloud.lustre.v1.DirectoryPolicy.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcplustredirectorypolicy;gcplustredirectorypolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// LustreDirectoryPolicy is the Schema for the LustreDirectoryPolicy API
// +k8s:openapi-gen=true
type LustreDirectoryPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   LustreDirectoryPolicySpec   `json:"spec,omitempty"`
	Status LustreDirectoryPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// LustreDirectoryPolicyList contains a list of LustreDirectoryPolicy
type LustreDirectoryPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LustreDirectoryPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LustreDirectoryPolicy{}, &LustreDirectoryPolicyList{})
}
