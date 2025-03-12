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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BackupDRManagementServerGVK = GroupVersion.WithKind("BackupDRManagementServer")

type Parent struct {
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`
}

// +kcc:proto=google.cloud.backupdr.v1.NetworkConfig
type NetworkConfig struct {
	// Optional. The resource name of the Google Compute Engine VPC network to
	//  which the ManagementServer instance is connected.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkConfig.network
	NetworkRef *refsv1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. The network connect mode of the ManagementServer instance. For
	//  this version, only PRIVATE_SERVICE_ACCESS is supported.
	// +kcc:proto:field=google.cloud.backupdr.v1.NetworkConfig.peering_mode
	PeeringMode *string `json:"peeringMode,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.WorkforceIdentityBasedOAuth2ClientID
type WorkforceIdentityBasedOAuth2ClientIDObservedState struct {
	// Output only. First party OAuth Client ID for Google Identities.
	// +kcc:proto:field=google.cloud.backupdr.v1.WorkforceIdentityBasedOAuth2ClientID.first_party_oauth2_client_id
	FirstPartyOAuth2ClientID *string `json:"firstPartyOAuth2ClientID,omitempty"`

	// Output only. Third party OAuth Client ID for External Identity Providers.
	// +kcc:proto:field=google.cloud.backupdr.v1.WorkforceIdentityBasedOAuth2ClientID.third_party_oauth2_client_id
	ThirdPartyOAuth2ClientID *string `json:"thirdPartyOAuth2ClientID,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.ManagementURI
type ManagementURIObservedState struct {
	// Output only. The ManagementServer AGM/RD WebUI URL.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementURI.web_ui
	WebUI *string `json:"webUI,omitempty"`

	// Output only. The ManagementServer AGM/RD API URL.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementURI.api
	API *string `json:"api,omitempty"`
}

// BackupDRManagementServerSpec defines the desired state of BackupDRManagementServer
// +kcc:proto=google.cloud.backupdr.v1.ManagementServer
type BackupDRManagementServerSpec struct {
	Parent `json:",inline"`

	// The BackupDRManagementServer name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The description of the ManagementServer instance (2048 characters
	//  or less).
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.description
	Description *string `json:"description,omitempty"`

	// Optional. Resource labels to represent user provided metadata.
	//  Labels currently defined:
	//  1. migrate_from_go=<false|true>
	//     If set to true, the MS is created in migration ready mode.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The type of the ManagementServer resource.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.type
	Type *string `json:"type,omitempty"`

	// Optional. VPC networks to which the ManagementServer instance is connected.
	//  For this version, only a single network is supported. This field is
	//  optional if MS is created without PSA
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.networks
	Networks []NetworkConfig `json:"networks,omitempty"`

	// Optional. Server specified ETag for the ManagementServer resource to
	//  prevent simultaneous updates from overwiting each other.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.etag
	Etag *string `json:"etag,omitempty"`
}

// BackupDRManagementServerStatus defines the config connector machine state of BackupDRManagementServer
type BackupDRManagementServerStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BackupDRManagementServer resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BackupDRManagementServerObservedState `json:"observedState,omitempty"`
}

// BackupDRManagementServerObservedState is the state of the BackupDRManagementServer resource as most recently observed in GCP.
// +kcc:proto=google.cloud.backupdr.v1.ManagementServer
type BackupDRManagementServerObservedState struct {
	// Output only. Identifier. The resource name.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. The time when the instance was created.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the instance was updated.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The hostname or ip address of the exposed AGM endpoints, used
	//  by clients to connect to AGM/RD graphical user interface and APIs.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.management_uri
	ManagementURI *ManagementURIObservedState `json:"managementURI,omitempty"`

	// Output only. The hostnames of the exposed AGM endpoints for both types of
	//  user i.e. 1p and 3p, used to connect AGM/RM UI.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.workforce_identity_based_management_uri
	WorkforceIdentityBasedManagementURI *WorkforceIdentityBasedManagementURIObservedState `json:"workforceIdentityBasedManagementURI,omitempty"`

	// Output only. The ManagementServer state.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.state
	State *string `json:"state,omitempty"`

	// Output only. The OAuth 2.0 client id is required to make API calls to the
	//  BackupDR instance API of this ManagementServer. This is the value that
	//  should be provided in the 'aud' field of the OIDC ID Token (see openid
	//  specification
	//  https://openid.net/specs/openid-connect-core-1_0.html#IDToken).
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.oauth2_client_id
	OAuth2ClientID *string `json:"oauth2ClientID,omitempty"`

	// Output only. The OAuth client IDs for both types of user i.e. 1p and 3p.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.workforce_identity_based_oauth2_client_id
	WorkforceIdentityBasedOAuth2ClientID *WorkforceIdentityBasedOAuth2ClientIDObservedState `json:"workforceIdentityBasedOAuth2ClientID,omitempty"`

	// Output only. The hostname or ip address of the exposed AGM endpoints, used
	//  by BAs to connect to BA proxy.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.ba_proxy_uri
	BAProxyURIs []string `json:"baProxyURIs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.satisfies_pzs
	// NOTYET
	// SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.backupdr.v1.ManagementServer.satisfies_pzi
	// NOTYET
	// SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbackupdrmanagementserver;gcpbackupdrmanagementservers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BackupDRManagementServer is the Schema for the BackupDRManagementServer API
// +k8s:openapi-gen=true
type BackupDRManagementServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BackupDRManagementServerSpec   `json:"spec,omitempty"`
	Status BackupDRManagementServerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BackupDRManagementServerList contains a list of BackupDRManagementServer
type BackupDRManagementServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupDRManagementServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupDRManagementServer{}, &BackupDRManagementServerList{})
}
