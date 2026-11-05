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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowCXWebhookGVK = GroupVersion.WithKind("DialogflowCXWebhook")

// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.Webhook
type DialogflowCXWebhookSpec struct {
	// Indicates whether the webhook is disabled.
	// +optional
	Disabled *bool `json:"disabled,omitempty"`

	// The human-readable name of the webhook, unique within the agent.
	// +required
	DisplayName string `json:"displayName"`

	// Indicates if automatic spell correction is enabled in detect intent requests.
	// +optional
	EnableSpellCorrection *bool `json:"enableSpellCorrection,omitempty"`

	// Determines whether this agent should log conversation queries.
	// +optional
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty"`

	// Configuration for a generic web service.
	// +optional
	GenericWebService *GenericWebService `json:"genericWebService,omitempty"`

	/* Immutable. The agent to create a webhook for.
	Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>. */
	// +optional
	Parent *string `json:"parent,omitempty"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Name of the SecuritySettings reference for the agent.
	Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>. */
	// +optional
	SecuritySettings *string `json:"securitySettings,omitempty"`

	// Configuration for a Service Directory service.
	// +optional
	ServiceDirectory *ServiceDirectoryConfig `json:"serviceDirectory,omitempty"`

	// Webhook execution timeout.
	// +optional
	Timeout *string `json:"timeout,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Webhook.GenericWebService
type GenericWebService struct {
	// Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification.
	// +optional
	AllowedCACerts []string `json:"allowedCaCerts,omitempty"`

	// Immutable. The HTTP request headers to send together with webhook requests.
	// +optional
	RequestHeaders map[string]string `json:"requestHeaders,omitempty"`

	// Whether to use speech adaptation for speech recognition.
	// +required
	URI string `json:"uri"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Webhook.ServiceDirectoryConfig
type ServiceDirectoryConfig struct {
	// The name of Service Directory service.
	// +required
	GenericWebService *GenericWebService `json:"genericWebService"`

	// The name of Service Directory service.
	// +required
	Service string `json:"service"`
}

type DialogflowCXWebhookStatus struct {
	// Conditions represent the latest available observation of the resource's current state.
	// +optional
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// A unique specifier for the DialogflowCXWebhook resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	/* The unique identifier of the webhook.
	Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>. */
	// +optional
	Name *string `json:"name,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>. */
	// +optional
	StartFlow *string `json:"startFlow,omitempty"`
}

// DialogflowCXWebhookObservedState is the state of the DialogflowCXWebhook resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.Webhook
type DialogflowCXWebhookObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowcxwebhook;gcpdialogflowcxwebhooks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowCXWebhook is the Schema for the DialogflowCXWebhook API
// +k8s:openapi-gen=true
type DialogflowCXWebhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowCXWebhookSpec   `json:"spec,omitempty"`
	Status DialogflowCXWebhookStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowCXWebhookList contains a list of DialogflowCXWebhook
type DialogflowCXWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowCXWebhook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowCXWebhook{}, &DialogflowCXWebhookList{})
}
