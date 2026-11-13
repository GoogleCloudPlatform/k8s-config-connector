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

// API sources for FTPServer, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/ftp/v1/service.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/cloud-ftp
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/cloud-ftp/docs/reference/rest/v1/projects.locations.servers

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FTPServerGVK = GroupVersion.WithKind("FTPServer")

// FTPServerSpec defines the desired state of FTPServer
// +kcc:spec:proto=google.cloud.ftp.v1.Server
// +kcc:required-from-proto
type FTPServerSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The FTPServer name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Configuration for external access.
	// +kcc:proto:field=google.cloud.ftp.v1.Server.external_config
	ExternalConfig *ExternalServerConfig `json:"externalConfig,omitempty"`

	// Configuration for internal access.
	// +kcc:proto:field=google.cloud.ftp.v1.Server.internal_config
	InternalConfig *InternalServerConfig `json:"internalConfig,omitempty"`

	// Optional. Labels as key value pairs
	// +kcc:proto:field=google.cloud.ftp.v1.Server.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Display name of the Server
	// +kcc:proto:field=google.cloud.ftp.v1.Server.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The access type of the Server.
	// +kcc:proto:field=google.cloud.ftp.v1.Server.access_type
	// +required
	AccessType *string `json:"accessType,omitempty"`
}

// FTPServerStatus defines the config connector machine state of FTPServer
type FTPServerStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the FTPServer resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *FTPServerObservedState `json:"observedState,omitempty"`
}

// FTPServerObservedState is the state of the FTPServer resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.ftp.v1.Server
type FTPServerObservedState struct {
	// Configuration for external access.
	// +kcc:proto:field=google.cloud.ftp.v1.Server.external_config
	ExternalConfig *ExternalServerConfigObservedState `json:"externalConfig,omitempty"`

	// Configuration for internal access.
	// +kcc:proto:field=google.cloud.ftp.v1.Server.internal_config
	InternalConfig *InternalServerConfigObservedState `json:"internalConfig,omitempty"`

	// Output only. [Output only] Create time stamp
	// +kcc:proto:field=google.cloud.ftp.v1.Server.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp
	// +kcc:proto:field=google.cloud.ftp.v1.Server.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the server.
	// +kcc:proto:field=google.cloud.ftp.v1.Server.state
	State *string `json:"state,omitempty"`

	// Output only. Credentials of the FTP Server.
	// +kcc:proto:field=google.cloud.ftp.v1.Server.google_managed_server_credential
	GoogleManagedServerCredential *ServerCredentialObservedState `json:"googleManagedServerCredential,omitempty"`

	// Output only. Service agent used to access the customer bucket.
	// +kcc:proto:field=google.cloud.ftp.v1.Server.service_agent
	ServiceAgent *string `json:"serviceAgent,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpftpserver;gcpftpservers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// FTPServer is the Schema for the FTPServer API
// +k8s:openapi-gen=true
type FTPServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FTPServerSpec   `json:"spec,omitempty"`
	Status FTPServerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// FTPServerList contains a list of FTPServer
type FTPServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FTPServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FTPServer{}, &FTPServerList{})
}
