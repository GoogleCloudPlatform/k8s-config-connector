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
	computev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	workflowsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workflows/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var EventarcPipelineGVK = GroupVersion.WithKind("EventarcPipeline")

// EventarcPipelineSpec defines the desired state of EventarcPipeline
// +kcc:spec:proto=google.cloud.eventarc.v1.Pipeline
type EventarcPipelineSpec struct {
	// The project that this resource belongs to.
	// +kubebuilder:validation:Required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location,omitempty"`

	// The EventarcPipeline name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. User labels attached to the Pipeline that can be used to group
	//  resources.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.labels
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. User-defined annotations.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.annotations
	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Display name of resource.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.display_name
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty"`

	// Required. List of destinations to which messages will be forwarded.
	//  Currently, exactly one destination is supported per Pipeline.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.destinations
	// +kubebuilder:validation:Required
	Destinations []Pipeline_Destination `json:"destinations,omitempty"`

	// Optional. List of mediation operations to be performed on the message.
	//  Currently, only one Transformation operation is allowed in each Pipeline.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.mediations
	// +kubebuilder:validation:Optional
	Mediations []Pipeline_Mediation `json:"mediations,omitempty"`

	// Optional. Resource name of a KMS crypto key (managed by the user) used to
	//  encrypt/decrypt the event data. If not set, an internal Google-owned key
	//  will be used to encrypt messages. It must match the pattern
	//  "projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{key}".
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.crypto_key_name
	// +kubebuilder:validation:Optional
	CryptoKeyRef *kmsv1beta1.KMSCryptoKeyRef `json:"cryptoKeyRef,omitempty"`

	// Optional. The payload format expected for the messages received by the
	//  Pipeline. If input_payload_format is set then any messages not matching
	//  this format will be treated as persistent errors. If input_payload_format
	//  is not set, then the message data will be treated as an opaque binary and
	//  no output format can be set on the Pipeline through the
	//  Pipeline.Destination.output_payload_format field. Any Mediations on the
	//  Pipeline that involve access to the data field will fail as persistent
	//  errors.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.input_payload_format
	// +kubebuilder:validation:Optional
	InputPayloadFormat *Pipeline_MessagePayloadFormat `json:"inputPayloadFormat,omitempty"`

	// Optional. Config to control Platform Logging for Pipelines.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.logging_config
	// +kubebuilder:validation:Optional
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	// Optional. The retry policy to use in the pipeline.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.retry_policy
	// +kubebuilder:validation:Optional
	RetryPolicy *Pipeline_RetryPolicy `json:"retryPolicy,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.Destination
type Pipeline_Destination struct {
	// Optional. Network config is used to configure how Pipeline resolves and
	//  connects to a destination.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.network_config
	// +kubebuilder:validation:Optional
	NetworkConfig *Pipeline_Destination_NetworkConfig `json:"networkConfig,omitempty"`

	// Optional. An HTTP endpoint destination described by an URI.
	//  If a DNS FQDN is provided as the endpoint, Pipeline will create a
	//  peering zone to the consumer VPC and forward DNS requests to the VPC
	//  specified by network config to resolve the service endpoint.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.http_endpoint
	// +kubebuilder:validation:Optional
	HTTPEndpoint *Pipeline_Destination_HTTPEndpoint `json:"httpEndpoint,omitempty"`

	// Optional. The resource name of the Workflow whose Executions are
	//  triggered by the events. The Workflow resource should be deployed in
	//  the same project as the Pipeline. Format:
	//  `projects/{project}/locations/{location}/workflows/{workflow}`
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.workflow
	// +kubebuilder:validation:Optional
	WorkflowRef *workflowsv1alpha1.WorkflowsWorkflowRef `json:"workflowRef,omitempty"`

	// Optional. The resource name of the Message Bus to which events should
	//  be published. The Message Bus resource should exist in the same project
	//  as the Pipeline. Format:
	//  `projects/{project}/locations/{location}/messageBuses/{message_bus}`
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.message_bus
	// +kubebuilder:validation:Optional
	MessageBusRef *EventarcMessageBusRef `json:"messageBusRef,omitempty"`

	// Optional. The resource name of the Pub/Sub topic to which events should
	//  be published. Format:
	//  `projects/{project}/locations/{location}/topics/{topic}`
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.topic
	// +kubebuilder:validation:Optional
	TopicRef *pubsubv1beta1.PubSubTopicRef `json:"topicRef,omitempty"`

	// Optional. An authentication config used to authenticate message requests,
	//  such that destinations can verify the source. For example, this can be
	//  used with private GCP destinations that require GCP credentials to access
	//  like Cloud Run. This field is optional and should be set only by users
	//  interested in authenticated push
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.authentication_config
	// +kubebuilder:validation:Optional
	AuthenticationConfig *Pipeline_Destination_AuthenticationConfig `json:"authenticationConfig,omitempty"`

	// Optional. The message format before it is delivered to the destination.
	//  If not set, the message will be delivered in the format it was originally
	//  delivered to the Pipeline. This field can only be set if
	//  Pipeline.input_payload_format is also set.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.output_payload_format
	// +kubebuilder:validation:Optional
	OutputPayloadFormat *Pipeline_MessagePayloadFormat `json:"outputPayloadFormat,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.Destination.NetworkConfig
type Pipeline_Destination_NetworkConfig struct {
	// Required. Name of the NetworkAttachment that allows access to the
	//  consumer VPC. Format:
	//  `projects/{PROJECT_ID}/regions/{REGION}/networkAttachments/{NETWORK_ATTACHMENT_NAME}`
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.NetworkConfig.network_attachment
	// +kubebuilder:validation:Required
	NetworkAttachmentRef *computev1alpha1.ComputeNetworkAttachmentRef `json:"networkAttachmentRef,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig
type Pipeline_Destination_AuthenticationConfig struct {
	// Optional. This authenticate method will apply Google OIDC tokens
	//  signed by a GCP service account to the requests.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.google_oidc
	// +kubebuilder:validation:Optional
	GoogleOIDC *Pipeline_Destination_AuthenticationConfig_OIDCToken `json:"googleOIDC,omitempty"`

	// Optional. If specified, an OAuth
	//  token will be generated and attached as an `Authorization` header in the HTTP
	//  request.
	//
	//  This type of authorization should generally only be used when calling
	//  Google APIs hosted on *.googleapis.com.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.oauth_token
	// +kubebuilder:validation:Optional
	OauthToken *Pipeline_Destination_AuthenticationConfig_OAuthToken `json:"oauthToken,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.OAuthToken
type Pipeline_Destination_AuthenticationConfig_OAuthToken struct {
	// Required. Service account email used to generate the OAuth
	//  token.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.OAuthToken.service_account
	// +kubebuilder:validation:Required
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Optional. OAuth scope to be used for generating OAuth access token.
	//  If not specified, "https://www.googleapis.com/auth/cloud-platform"
	//  will be used.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.OAuthToken.scope
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.OidcToken
type Pipeline_Destination_AuthenticationConfig_OIDCToken struct {
	// Required. Service account email used to generate the OIDC Token.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.OidcToken.service_account
	// +kubebuilder:validation:Required
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Optional. Audience to be used to generate the OIDC Token. The
	//  audience claim identifies the recipient that the JWT is intended for.
	//  If unspecified, the destination URI will be used.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.OidcToken.audience
	// +kubebuilder:validation:Optional
	Audience *string `json:"audience,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.MessagePayloadFormat.JsonFormat
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:Schemaless
type Pipeline_MessagePayloadFormat_JsonFormat struct {
}

// EventarcPipelineStatus defines the config connector machine state of EventarcPipeline
type EventarcPipelineStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the EventarcPipeline resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *EventarcPipelineObservedState `json:"observedState,omitempty"`
}

// EventarcPipelineObservedState is the state of the EventarcPipeline resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.eventarc.v1.Pipeline
type EventarcPipelineObservedState struct {
	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last-modified time.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Server-assigned unique identifier for the Pipeline. The value
	//  is a UUID4 string and guaranteed to remain unchanged until the resource is
	//  deleted.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. This checksum is computed by the server based on the value of
	//  other fields, and might be sent only on create requests to ensure that the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpeventarcpipeline;gcpeventarcpipelines
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// EventarcPipeline is the Schema for the EventarcPipeline API
// +k8s:openapi-gen=true
type EventarcPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   EventarcPipelineSpec   `json:"spec,omitempty"`
	Status EventarcPipelineStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EventarcPipelineList contains a list of EventarcPipeline
type EventarcPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventarcPipeline `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EventarcPipeline{}, &EventarcPipelineList{})
}
