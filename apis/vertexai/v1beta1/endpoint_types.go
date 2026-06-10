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

package v1beta1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ apiextensionsv1.JSON

var VertexAIEndpointGVK = GroupVersion.WithKind("VertexAIEndpoint")

// VertexAIEndpointSpec defines the desired state of VertexAIEndpoint
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.Endpoint
type VertexAIEndpointSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The region of this resource.
	Region string `json:"region"`

	// The VertexAIEndpoint name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the Endpoint.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Endpoint.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the Endpoint.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Endpoint.description
	Description *string `json:"description,omitempty"`

	// The labels with user-defined metadata to organize your Endpoints.
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Endpoint.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The full name of the Google Compute Engine
	//  [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered.
	//  Private services access must already be configured for the network. If left
	//  unspecified, the Endpoint is not peered with any network.
	//  Only one of the fields,
	//  [network][google.cloud.aiplatform.v1beta1.Endpoint.network] or
	//  [enable_private_service_connect][google.cloud.aiplatform.v1beta1.Endpoint.enable_private_service_connect],
	//  can be set.
	//  [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/get):
	//  `projects/{project}/global/networks/{network}`.
	//  Where {project} is a project number, as in '12345', and {network} is
	//  network name.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Endpoint.network
	NetworkRef *refsv1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Deprecated: If true, expose the Endpoint via private service connect.
	//  Only one of the fields,
	//  [network][google.cloud.aiplatform.v1beta1.Endpoint.network] or
	//  [enable_private_service_connect][google.cloud.aiplatform.v1beta1.Endpoint.enable_private_service_connect],
	//  can be set.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Endpoint.enable_private_service_connect
	EnablePrivateServiceConnect *bool `json:"enablePrivateServiceConnect,omitempty"`

	// Configuration for private service connect.
	//  [network][google.cloud.aiplatform.v1beta1.Endpoint.network] and
	//  [private_service_connect_config][google.cloud.aiplatform.v1beta1.Endpoint.private_service_connect_config]
	//  are mutually exclusive.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Endpoint.private_service_connect_config
	// PrivateServiceConnectConfig *PrivateServiceConnectConfig `json:"privateServiceConnectConfig,omitempty"`

	// Configures the request-response logging for online prediction.	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Endpoint.predict_request_response_logging_config
	PredictRequestResponseLoggingConfig *PredictRequestResponseLoggingConfig `json:"predictRequestResponseLoggingConfig,omitempty"`

	// Customer-managed encryption key spec for an Endpoint. If set, this
	// Endpoint and all sub-resources of this Endpoint will be secured by
	// this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Endpoint.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// If true, the endpoint will be exposed through a dedicated DNS	//  [Endpoint.dedicated_endpoint_dns]. Your request to the dedicated DNS will
	//  be isolated from other users' traffic and will have better performance and
	//  routing capabilities.
	//  Note: Once you enabled dedicated endpoint, you won't be able to send
	//  request to the shared DNS {region}-aiplatform.googleapis.com. The
	//  limitation will be removed soon.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Endpoint.dedicated_endpoint_enabled
	DedicatedEndpointEnabled *bool `json:"dedicatedEndpointEnabled,omitempty"`

	// Configurations that are applied to the endpoint for online prediction.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Endpoint.client_connection_config
	ClientConnectionConfig *ClientConnectionConfig `json:"clientConnectionConfig,omitempty"`
}

// VertexAIEndpointStatus defines the config connector machine state of VertexAIEndpoint
type VertexAIEndpointStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIEndpoint resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIEndpointObservedState `json:"observedState,omitempty"`
}

// VertexAIEndpointObservedState is the state of the VertexAIEndpoint resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.Endpoint
type VertexAIEndpointObservedState struct {
	// Output only. Timestamp when this Endpoint was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Endpoint.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Endpoint was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Endpoint.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. DNS of the dedicated endpoint. Will only be populated if
	//  dedicated_endpoint_enabled is true.
	//  Format:
	//  `https://{endpoint_id}.{region}-{project_number}.prediction.vertexai.goog`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Endpoint.dedicated_endpoint_dns
	DedicatedEndpointDNS *string `json:"dedicatedEndpointDNS,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaiendpoint;gcpvertexaiendpoints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIEndpoint is the Schema for the VertexAIEndpoint API
// +k8s:openapi-gen=true
type VertexAIEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIEndpointSpec   `json:"spec,omitempty"`
	Status VertexAIEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIEndpointList contains a list of VertexAIEndpoint
type VertexAIEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIEndpoint{}, &VertexAIEndpointList{})
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PrivateServiceConnectConfig
type PrivateServiceConnectConfig struct {
	// Required. If true, expose the IndexEndpoint via private service connect.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PrivateServiceConnectConfig.enable_private_service_connect
	EnablePrivateServiceConnect *bool `json:"enablePrivateServiceConnect,omitempty"`

	// A list of Projects from which the forwarding rule will target the service
	//  attachment.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PrivateServiceConnectConfig.project_allowlist
	ProjectAllowlist []refsv1beta1.ProjectRef `json:"projectAllowlist,omitempty"`

	// Optional. List of projects and networks where the PSC endpoints will be
	//  created. This field is used by Online Inference(Prediction) only.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PrivateServiceConnectConfig.psc_automation_configs
	PSCAutomationConfigs []PSCAutomationConfig `json:"pscAutomationConfigs,omitempty"`

	// Optional. If set to true, enable secure private service connect with IAM
	//  authorization. Otherwise, private service connect will be done without
	//  authorization. Note latency will be slightly increased if authorization is
	//  enabled.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PrivateServiceConnectConfig.enable_secure_private_service_connect
	EnableSecurePrivateServiceConnect *bool `json:"enableSecurePrivateServiceConnect,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PSCAutomationConfig
type PSCAutomationConfig struct {
	// Required. Project id used to create forwarding rule.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PSCAutomationConfig.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Required. The full name of the Google Compute Engine
	//  [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks).
	//  [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/get):
	//  `projects/{project}/global/networks/{network}`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PSCAutomationConfig.network
	NetworkRef *refsv1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`
}
