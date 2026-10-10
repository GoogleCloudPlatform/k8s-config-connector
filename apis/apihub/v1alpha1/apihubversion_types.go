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

// API sources for APIHubVersion, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/apihub/v1/common_fields.proto
// +kcc:source:service-docs=https://cloud.google.com/apigee/docs/api-hub/what-is-api-hub
// +kcc:source:resource-docs=https://docs.cloud.google.com/apigee/docs/reference/apis/apihub/rest/v1/projects.locations.apis.versions

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var APIHubVersionGVK = GroupVersion.WithKind("APIHubVersion")

// APIHubVersionSpec defines the desired state of APIHubVersion
// +kcc:spec:proto=google.cloud.apihub.v1.Version
// +kcc:required-from-proto
type APIHubVersionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The APIHubVersion name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The display name of the version.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the version.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.description
	Description *string `json:"description,omitempty"`

	// Optional. The documentation of the version.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.documentation
	Documentation *Documentation `json:"documentation,omitempty"`

	// Optional. The deployments linked to this API version.
	//  Note: A particular API version could be deployed to multiple deployments
	//  (for dev deployment, UAT deployment, etc)
	//  Format is
	//  `projects/{project}/locations/{location}/deployments/{deployment}`
	// +kcc:guess=possible-reference target=APIHubDeployment
	// +kcc:proto:field=google.cloud.apihub.v1.Version.deployments
	Deployments []string `json:"deployments,omitempty"`

	// Optional. The lifecycle of the API version.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-lifecycle`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.lifecycle
	Lifecycle *AttributeValuesRequired `json:"lifecycle,omitempty"`

	// Optional. The compliance associated with the API version.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-compliance`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.compliance
	Compliance *AttributeValuesRequired `json:"compliance,omitempty"`

	// Optional. The accreditations associated with the API version.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-accreditation`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.accreditation
	Accreditation *AttributeValuesRequired `json:"accreditation,omitempty"`

	// Optional. The list of user defined attributes associated with the Version
	//  resource. The key is the attribute name. It will be of the format:
	//  `projects/{project}/locations/{location}/attributes/{attribute}`.
	//  The value is the attribute values associated with the resource.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.attributes
	Attributes map[string]AttributeValuesRequired `json:"attributes,omitempty"`

	// Optional. The selected deployment for a Version resource.
	//  This can be used when special handling is needed on client side for a
	//  particular deployment linked to the version.
	//  Format is
	//  `projects/{project}/locations/{location}/deployments/{deployment}`
	// +kcc:proto:field=google.cloud.apihub.v1.Version.selected_deployment
	SelectedDeployment *string `json:"selectedDeployment,omitempty"`
}

// APIHubVersionStatus defines the config connector machine state of APIHubVersion
type APIHubVersionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the APIHubVersion resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *APIHubVersionObservedState `json:"observedState,omitempty"`
}

// APIHubVersionObservedState is the state of the APIHubVersion resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apihub.v1.Version
type APIHubVersionObservedState struct {
	// Output only. The specs associated with this version.
	//  Note that an API version can be associated with multiple specs.
	//  Format is
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}/specs/{spec}`
	// +kcc:proto:field=google.cloud.apihub.v1.Version.specs
	Specs []string `json:"specs,omitempty"`

	// Output only. The operations contained in the API version.
	//  These operations will be added to the version when a new spec is
	//  added or when an existing spec is updated. Format is
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}/operations/{operation}`
	// +kcc:proto:field=google.cloud.apihub.v1.Version.api_operations
	APIOperations []string `json:"apiOperations,omitempty"`

	// Output only. The definitions contained in the API version.
	//  These definitions will be added to the version when a new spec is
	//  added or when an existing spec is updated. Format is
	//  `projects/{project}/locations/{location}/apis/{api}/versions/{version}/definitions/{definition}`
	// +kcc:proto:field=google.cloud.apihub.v1.Version.definitions
	Definitions []string `json:"definitions,omitempty"`

	// Output only. The time at which the version was created.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the version was last updated.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. The lifecycle of the API version.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-lifecycle`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.lifecycle
	Lifecycle *AttributeValuesObservedState `json:"lifecycle,omitempty"`

	// Optional. The compliance associated with the API version.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-compliance`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.compliance
	Compliance *AttributeValuesObservedState `json:"compliance,omitempty"`

	// Optional. The accreditations associated with the API version.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-accreditation`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.accreditation
	Accreditation *AttributeValuesObservedState `json:"accreditation,omitempty"`

	// Output only. The list of sources and metadata from the sources of the
	//  version.
	// +kcc:proto:field=google.cloud.apihub.v1.Version.source_metadata
	SourceMetadata []SourceMetadataObservedState `json:"sourceMetadata,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapihubversion;gcpapihubversions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// APIHubVersion is the Schema for the APIHubVersion API
// +k8s:openapi-gen=true
type APIHubVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   APIHubVersionSpec   `json:"spec,omitempty"`
	Status APIHubVersionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// APIHubVersionList contains a list of APIHubVersion
type APIHubVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIHubVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIHubVersion{}, &APIHubVersionList{})
}
