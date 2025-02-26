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
	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ApigeeEndpointAttachmentGVK = GroupVersion.WithKind("ApigeeEndpointAttachment")

// ApigeeEndpointAttachmentSpec defines the desired state of ApigeeEndpointAttachment
// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1EndpointAttachment
type ApigeeEndpointAttachmentSpec struct {
	// Reference to parent Apigee Organization.
	// +required
	OrganizationRef *apigeev1beta1.ApigeeOrganizationRef `json:"organizationRef"`

	// Required. Location of the endpoint attachment.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1EndpointAttachment.location
	// +rquired
	Location *string `json:"location,omitempty"`

	// Reference to the ServiceAttachment for the EndpointAttachment.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1EndpointAttachment.service_attachment
	ServiceAttachmentRef *refs.ComputeServiceAttachmentRef `json:"serviceAttachmentRef,omitempty"`

	// The ApigeeEndpointAttachment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// ApigeeEndpointAttachmentStatus defines the config connector machine state of ApigeeEndpointAttachment
type ApigeeEndpointAttachmentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ApigeeEndpointAttachment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ApigeeEndpointAttachmentObservedState `json:"observedState,omitempty"`
}

// ApigeeEndpointAttachmentObservedState is the state of the ApigeeEndpointAttachment resource as most recently observed in GCP.
// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1EndpointAttachment
type ApigeeEndpointAttachmentObservedState struct {
	// Output only. State of the endpoint attachment connection to the service attachment.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1EndpointAttachment.connection_state
	ConnectionState *string `json:"connectionState,omitempty"`

	// Output only. Host that can be used in either the HTTP target endpoint directly or as the host in target server.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1EndpointAttachment.host
	Host *string `json:"host,omitempty"`

	// Output only. State of the endpoint attachment. Values other than `ACTIVE` mean the resource is not ready to use.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1EndpointAttachment.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeendpointattachment;gcpapigeeendpointattachments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ApigeeEndpointAttachment is the Schema for the ApigeeEndpointAttachment API
// +k8s:openapi-gen=true
type ApigeeEndpointAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ApigeeEndpointAttachmentSpec   `json:"spec,omitempty"`
	Status ApigeeEndpointAttachmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ApigeeEndpointAttachmentList contains a list of ApigeeEndpointAttachment
type ApigeeEndpointAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeEndpointAttachment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeEndpointAttachment{}, &ApigeeEndpointAttachmentList{})
}
