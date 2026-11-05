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

// API sources for ApigeeregistryAPIDeployment, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/apigeeregistry/v1/registry_models.proto
// +kcc:guess=source-link reason=verify-service-docs-link
// +kcc:source:service-docs=
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ApigeeregistryAPIDeploymentGVK = GroupVersion.WithKind("ApigeeregistryAPIDeployment")

// ApigeeregistryAPIDeploymentSpec defines the desired state of ApigeeregistryAPIDeployment
// +kcc:spec:proto=google.cloud.apigeeregistry.v1.ApiDeployment
// +kcc:required-from-proto
type ApigeeregistryAPIDeploymentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ApigeeregistryAPIDeployment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Human-meaningful name.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A detailed description.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.description
	Description *string `json:"description,omitempty"`

	// The full resource name (including revision ID) of the spec of the API being
	//  served by the deployment. Changes to this value will update the revision.
	//  Format: `apis/{api}/deployments/{deployment}`
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.api_spec_revision
	APISpecRevision *string `json:"apiSpecRevision,omitempty"`

	// The address where the deployment is serving. Changes to this value will
	//  update the revision.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.endpoint_uri
	EndpointURI *string `json:"endpointURI,omitempty"`

	// The address of the external channel of the API (e.g., the Developer
	//  Portal). Changes to this value will not affect the revision.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.external_channel_uri
	ExternalChannelURI *string `json:"externalChannelURI,omitempty"`

	// Text briefly identifying the intended audience of the API. Changes to this
	//  value will not affect the revision.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.intended_audience
	IntendedAudience *string `json:"intendedAudience,omitempty"`

	// Text briefly describing how to access the endpoint. Changes to this value
	//  will not affect the revision.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.access_guidance
	AccessGuidance *string `json:"accessGuidance,omitempty"`

	// Labels attach identifying metadata to resources. Identifying metadata can
	//  be used to filter list operations.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one resource (System
	//  labels are excluded).
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	//  System reserved label keys are prefixed with
	//  `apigeeregistry.googleapis.com/` and cannot be changed.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations attach non-identifying metadata to resources.
	//
	//  Annotation keys and values are less restricted than those of labels, but
	//  should be generally used for small values of broad interest. Larger, topic-
	//  specific metadata should be stored in Artifacts.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// ApigeeregistryAPIDeploymentStatus defines the config connector machine state of ApigeeregistryAPIDeployment
type ApigeeregistryAPIDeploymentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ApigeeregistryAPIDeployment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ApigeeregistryAPIDeploymentObservedState `json:"observedState,omitempty"`
}

// ApigeeregistryAPIDeploymentObservedState is the state of the ApigeeregistryAPIDeployment resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apigeeregistry.v1.ApiDeployment
type ApigeeregistryAPIDeploymentObservedState struct {
	// Output only. Immutable. The revision ID of the deployment.
	//  A new revision is committed whenever the deployment contents are changed.
	//  The format is an 8-character hexadecimal string.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. Creation timestamp; when the deployment resource was created.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Revision creation timestamp; when the represented revision was created.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`

	// Output only. Last update timestamp: when the represented revision was last modified.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiDeployment.revision_update_time
	RevisionUpdateTime *string `json:"revisionUpdateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeregistryapideployment;gcpapigeeregistryapideployments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ApigeeregistryAPIDeployment is the Schema for the ApigeeregistryAPIDeployment API
// +k8s:openapi-gen=true
type ApigeeregistryAPIDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ApigeeregistryAPIDeploymentSpec   `json:"spec,omitempty"`
	Status ApigeeregistryAPIDeploymentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ApigeeregistryAPIDeploymentList contains a list of ApigeeregistryAPIDeployment
type ApigeeregistryAPIDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeregistryAPIDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeregistryAPIDeployment{}, &ApigeeregistryAPIDeploymentList{})
}
