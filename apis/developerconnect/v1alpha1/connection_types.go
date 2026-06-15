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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	servicedirectoryv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DevConnectConnectionGVK = GroupVersion.WithKind("DevConnectConnection")

// DevConnectConnectionSpec defines the desired state of DevConnectConnection
// +kcc:spec:proto=google.cloud.developerconnect.v1.Connection
type DevConnectConnectionSpec struct {
	// The project that this resource belongs to.
	// +kubebuilder:validation:Required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The DevConnectConnection name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Configuration for connections to github.com.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.github_config
	GithubConfig *GitHubConfig `json:"githubConfig,omitempty"`

	// Configuration for connections to an instance of GitHub Enterprise.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.github_enterprise_config
	GithubEnterpriseConfig *GitHubEnterpriseConfig `json:"githubEnterpriseConfig,omitempty"`

	// Configuration for connections to gitlab.com.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.gitlab_config
	GitlabConfig *GitLabConfig `json:"gitlabConfig,omitempty"`

	// Configuration for connections to an instance of GitLab Enterprise.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.gitlab_enterprise_config
	GitlabEnterpriseConfig *GitLabEnterpriseConfig `json:"gitlabEnterpriseConfig,omitempty"`

	// Configuration for connections to an instance of Bitbucket Data Center.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.bitbucket_data_center_config
	BitbucketDataCenterConfig *BitbucketDataCenterConfig `json:"bitbucketDataCenterConfig,omitempty"`

	// Configuration for connections to an instance of Bitbucket Clouds.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.bitbucket_cloud_config
	BitbucketCloudConfig *BitbucketCloudConfig `json:"bitbucketCloudConfig,omitempty"`

	// Optional. If disabled is set to true, functionality is disabled for this
	//  connection. Repository based API methods and webhooks processing for
	//  repositories in this connection will be disabled.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Optional. Allows clients to store small amounts of arbitrary data.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. The crypto key configuration. This field is used by the
	//  Customer-Managed Encryption Keys (CMEK) feature.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.crypto_key_config
	CryptoKeyConfig *CryptoKeyConfig `json:"cryptoKeyConfig,omitempty"`

	// Optional. Configuration for the git proxy feature. Enabling the git proxy
	//  allows clients to perform git operations on the repositories linked in the
	//  connection.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.git_proxy_config
	GitProxyConfig *GitProxyConfig `json:"gitProxyConfig,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.ServiceDirectoryConfig
type ServiceDirectoryConfig struct {
	// The Service Directory service.
	// +kcc:proto:field=google.cloud.developerconnect.v1.ServiceDirectoryConfig.service
	ServiceRef *servicedirectoryv1alpha1.ServiceDirectoryServiceRef `json:"serviceRef,omitempty"`
}

// DevConnectConnectionStatus defines the config connector machine state of DevConnectConnection
type DevConnectConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DevConnectConnection resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DevConnectConnectionObservedState `json:"observedState,omitempty"`
}

// DevConnectConnectionObservedState is the state of the DevConnectConnection resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.developerconnect.v1.Connection
type DevConnectConnectionObservedState struct {
	// Output only. [Output only] Create timestamp
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update timestamp
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. [Output only] Delete timestamp
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. Installation state of the Connection.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.installation_state
	InstallationState *InstallationStateObservedState `json:"installationState,omitempty"`

	// Output only. Set to true when the connection is being set up or updated in
	//  the background.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. A system-assigned unique identifier for the Connection.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.uid
	Uid *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdevconnectconnection;gcpdevconnectconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DevConnectConnection is the Schema for the DevConnectConnection API
// +k8s:openapi-gen=true
type DevConnectConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DevConnectConnectionSpec   `json:"spec,omitempty"`
	Status DevConnectConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DevConnectConnectionList contains a list of DevConnectConnection
type DevConnectConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevConnectConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DevConnectConnection{}, &DevConnectConnectionList{})
}
