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

// API sources for FTPUser, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/befe6dc01ccbe4fc0292c5966efca2ad1fafc2e6/google/cloud/ftp/v1/service.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/cloud-ftp
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/cloud-ftp/docs/reference/rest/v1/projects.locations.servers.users

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FTPUserGVK = GroupVersion.WithKind("FTPUser")

// FTPUserSpec defines the desired state of FTPUser
// +kcc:spec:proto=google.cloud.ftp.v1.User
// +kcc:required-from-proto
type FTPUserSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/servers/{server} this resource belongs to.
	// +kcc:guess
	// ServerRef *FTPServerRef `json:"serverRef,omitempty"`

	// The FTPUser name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Labels as key value pairs
	// +kcc:proto:field=google.cloud.ftp.v1.User.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Service account in customer project attached to this SFTP User.
	// +kcc:proto:field=google.cloud.ftp.v1.User.customer_service_account
	// +required
	CustomerServiceAccount *string `json:"customerServiceAccount,omitempty"`

	// Required. User credential for the user.
	//  The maximum number of user credentials is 10.
	// +kcc:proto:field=google.cloud.ftp.v1.User.user_credentials
	// +required
	UserCredentials []UserCredential `json:"userCredentials,omitempty"`

	// Required. Mapping of Cloud Storage buckets to directories where the user
	//  will land in the SFTP server.
	// +kcc:proto:field=google.cloud.ftp.v1.User.storage_directory_mappings
	// +required
	StorageDirectoryMappings []StorageDirectoryMapping `json:"storageDirectoryMappings,omitempty"`
}

// FTPUserStatus defines the config connector machine state of FTPUser
type FTPUserStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the FTPUser resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *FTPUserObservedState `json:"observedState,omitempty"`
}

// FTPUserObservedState is the state of the FTPUser resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.ftp.v1.User
type FTPUserObservedState struct {
	// Output only. [Output only] Create time stamp
	// +kcc:proto:field=google.cloud.ftp.v1.User.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp
	// +kcc:proto:field=google.cloud.ftp.v1.User.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. [Output only] The username of the user.
	// +kcc:proto:field=google.cloud.ftp.v1.User.username
	Username *string `json:"username,omitempty"`

	// Output only. Tracks user creation.
	// +kcc:proto:field=google.cloud.ftp.v1.User.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpftpuser;gcpftpusers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// FTPUser is the Schema for the FTPUser API
// +k8s:openapi-gen=true
type FTPUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FTPUserSpec   `json:"spec,omitempty"`
	Status FTPUserStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// FTPUserList contains a list of FTPUser
type FTPUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FTPUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FTPUser{}, &FTPUserList{})
}
