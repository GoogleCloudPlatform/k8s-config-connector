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
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	servicedirectoryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowWebhookGVK = GroupVersion.WithKind("DialogflowWebhook")

// DialogflowWebhookSpec defines the desired state of DialogflowWebhook
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.Webhook
type DialogflowWebhookSpec struct {
	// The DialogflowCXAgent that this webhook belongs to.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="agentRef is immutable"
	AgentRef *DialogflowCXAgentRef `json:"agentRef"`

	// Required. The human-readable name of the webhook, unique within the agent.
	// +required
	DisplayName *string `json:"displayName"`

	// Configuration for a generic web service.
	// +optional
	GenericWebService *GenericWebService `json:"genericWebService,omitempty"`

	// Configuration for a Service Directory service.
	// +optional
	ServiceDirectory *ServiceDirectoryConfig `json:"serviceDirectory,omitempty"`

	// Webhook execution timeout. Execution is considered failed if Dialogflow
	// doesn't receive a response from webhook at the end of the timeout period.
	// Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
	// +optional
	Timeout *string `json:"timeout,omitempty"`

	// Indicates whether the webhook is disabled.
	// +optional
	Disabled *bool `json:"disabled,omitempty"`

	// The DialogflowWebhook name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Webhook.GenericWebService
type GenericWebService struct {
	// Required. The webhook URI for receiving POST requests. It must use https protocol.
	// +required
	URI *string `json:"uri"`

	// The user name for HTTP Basic authentication.
	// +optional
	Username *string `json:"username,omitempty"`

	// The password for HTTP Basic authentication.
	// +optional
	Password *refsv1beta1secret.Legacy `json:"password,omitempty"`

	// The HTTP request headers to send together with webhook requests.
	// +optional
	RequestHeaders map[string]string `json:"requestHeaders,omitempty"`

	// Optional. Specifies a list of allowed custom CA certificates (in DER format)
	// for HTTPS verification.
	// +optional
	AllowedCACerts [][]byte `json:"allowedCACerts,omitempty"`

	// Optional. The OAuth configuration of the webhook.
	// +optional
	OauthConfig *GenericWebServiceOAuthConfig `json:"oauthConfig,omitempty"`

	// Optional. Indicate the auth token type generated from the Dialogflow service agent.
	// +kubebuilder:validation:Enum=SERVICE_AGENT_AUTH_UNSPECIFIED;NONE;ID_TOKEN;ACCESS_TOKEN
	// +optional
	ServiceAgentAuth *string `json:"serviceAgentAuth,omitempty"`

	// Optional. Type of the webhook.
	// +kubebuilder:validation:Enum=WEBHOOK_TYPE_UNSPECIFIED;STANDARD;FLEXIBLE
	// +optional
	WebhookType *string `json:"webhookType,omitempty"`

	// Optional. HTTP method for the flexible webhook calls. Standard webhook always uses POST.
	// +kubebuilder:validation:Enum=HTTP_METHOD_UNSPECIFIED;POST;GET;HEAD;PUT;DELETE;PATCH;OPTIONS
	// +optional
	HTTPMethod *string `json:"httpMethod,omitempty"`

	// Optional. Defines a custom JSON object as request body to send to flexible webhook.
	// +optional
	RequestBody *string `json:"requestBody,omitempty"`

	// Optional. Maps the values extracted from specific fields of the flexible
	// webhook response into session parameters.
	// +optional
	ParameterMapping map[string]string `json:"parameterMapping,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Webhook.GenericWebService.OAuthConfig
type GenericWebServiceOAuthConfig struct {
	// Required. The client ID provided by the 3rd party platform.
	// +required
	ClientID *string `json:"clientID"`

	// Optional. The client secret provided by the 3rd party platform.
	// +optional
	ClientSecret *string `json:"clientSecret,omitempty"`

	// Required. The token endpoint provided by the 3rd party platform to exchange an access token.
	// +required
	TokenEndpoint *string `json:"tokenEndpoint"`

	// Optional. The OAuth scopes to grant.
	// +optional
	Scopes []string `json:"scopes,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Webhook.ServiceDirectoryConfig
type ServiceDirectoryConfig struct {
	// Required. The Service Directory service to call.
	// +required
	ServiceRef *servicedirectoryv1beta1.ServiceDirectoryServiceRef `json:"serviceRef"`

	// Generic Service configuration of this webhook.
	// +optional
	GenericWebService *GenericWebService `json:"genericWebService,omitempty"`
}

// DialogflowWebhookStatus defines the config connector machine state of DialogflowWebhook
type DialogflowWebhookStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowWebhook resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *DialogflowWebhookObservedState `json:"observedState,omitempty"`
}

// DialogflowWebhookObservedState is the state of the DialogflowWebhook resource as most recently observed in GCP.
// // +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.Webhook
// type DialogflowWebhookObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowwebhook;gcpdialogflowwebhooks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowWebhook is the Schema for the DialogflowWebhook API
// +k8s:openapi-gen=true
type DialogflowWebhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowWebhookSpec   `json:"spec,omitempty"`
	Status DialogflowWebhookStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowWebhookList contains a list of DialogflowWebhook
type DialogflowWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowWebhook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowWebhook{}, &DialogflowWebhookList{})
}
