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

// API sources for GKEMultiCloudAzureClient, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/gkemulticloud/v1/azure_resources.proto
// +kcc:source:service-docs=https://cloud.google.com/kubernetes-engine/multi-cloud/docs
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.azureClients

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GKEMultiCloudAzureClientGVK = GroupVersion.WithKind("GKEMultiCloudAzureClient")

// GKEMultiCloudAzureClientSpec defines the desired state of GKEMultiCloudAzureClient
// +kcc:spec:proto=google.cloud.gkemulticloud.v1.AzureClient
// +kcc:required-from-proto
type GKEMultiCloudAzureClientSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The GKEMultiCloudAzureClient name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The Azure Active Directory Tenant ID.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.tenant_id
	// +required
	TenantID *string `json:"tenantID,omitempty"`

	// Required. The Azure Active Directory Application ID.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.application_id
	// +required
	ApplicationID *string `json:"applicationID,omitempty"`

	// Optional. Annotations on the resource.
	//
	//  This field has the same restrictions as Kubernetes annotations.
	//  The total size of all keys and values combined is limited to 256k.
	//  Keys can have 2 segments: prefix (optional) and name (required),
	//  separated by a slash (/).
	//  Prefix must be a DNS subdomain.
	//  Name must be 63 characters or less, begin and end with alphanumerics,
	//  with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// GKEMultiCloudAzureClientStatus defines the config connector machine state of GKEMultiCloudAzureClient
type GKEMultiCloudAzureClientStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GKEMultiCloudAzureClient resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GKEMultiCloudAzureClientObservedState `json:"observedState,omitempty"`
}

// GKEMultiCloudAzureClientObservedState is the state of the GKEMultiCloudAzureClient resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.gkemulticloud.v1.AzureClient
type GKEMultiCloudAzureClientObservedState struct {
	// Output only. If set, there are currently pending changes to the client.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The PEM encoded x509 certificate.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.pem_certificate
	PemCertificate *string `json:"pemCertificate,omitempty"`

	// Output only. A globally unique identifier for the client.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time at which this resource was created.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this client was last updated.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureClient.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkemulticloudazureclient;gcpgkemulticloudazureclients
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEMultiCloudAzureClient is the Schema for the GKEMultiCloudAzureClient API
// +k8s:openapi-gen=true
type GKEMultiCloudAzureClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GKEMultiCloudAzureClientSpec   `json:"spec,omitempty"`
	Status GKEMultiCloudAzureClientStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GKEMultiCloudAzureClientList contains a list of GKEMultiCloudAzureClient
type GKEMultiCloudAzureClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEMultiCloudAzureClient `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEMultiCloudAzureClient{}, &GKEMultiCloudAzureClientList{})
}
