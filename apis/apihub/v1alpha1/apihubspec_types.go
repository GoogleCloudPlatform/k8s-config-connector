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

// API sources for ApihubSpec, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/apihub/v1/common_fields.proto
// +kcc:source:service-docs=https://cloud.google.com/apigee/docs/api-hub/what-is-api-hub
// +kcc:source:resource-docs=https://docs.cloud.google.com/apigee/docs/reference/apis/apihub/rest/v1/projects.locations.apis.versions.specs

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ApihubSpecGVK = GroupVersion.WithKind("ApihubSpec")

// ApihubSpecSpec defines the desired state of ApihubSpec
// +kcc:spec:proto=google.cloud.apihub.v1.Spec
// +kcc:required-from-proto
type ApihubSpecSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/apis/{api}/versions/{version} this resource belongs to.
	// +kcc:guess
	// VersionRef *APIHubVersionRef `json:"versionRef,omitempty"`

	// The ApihubSpec name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The display name of the spec.
	//  This can contain the file name of the spec.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The type of spec.
	//  The value should be one of the allowed values defined for
	//  `projects/{project}/locations/{location}/attributes/system-spec-type`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API.
	//
	//  Note, this field is mandatory if content is provided.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.spec_type
	// +required
	SpecType *AttributeValuesWithRequired `json:"specType,omitempty"`

	// Optional. Input only. The contents of the uploaded spec.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.contents
	Contents *SpecContents `json:"contents,omitempty"`

	// Optional. The URI of the spec source in case file is uploaded
	//  from an external version control system.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.source_uri
	SourceURI *string `json:"sourceURI,omitempty"`

	// Optional. The lint response for the spec.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.lint_response
	LintResponse *LintResponse `json:"lintResponse,omitempty"`

	// Optional. The list of user defined attributes associated with the spec.
	//  The key is the attribute name. It will be of the format:
	//  `projects/{project}/locations/{location}/attributes/{attribute}`.
	//  The value is the attribute values associated with the resource.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.attributes
	Attributes map[string]AttributeValuesWithRequired `json:"attributes,omitempty"`

	// Optional. The documentation of the spec.
	//  For OpenAPI spec, this will be populated from `externalDocs` in OpenAPI
	//  spec.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.documentation
	Documentation *Documentation `json:"documentation,omitempty"`

	// Optional. Input only. Enum specifying the parsing mode for OpenAPI
	//  Specification (OAS) parsing.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.parsing_mode
	ParsingMode *string `json:"parsingMode,omitempty"`
}

// ApihubSpecStatus defines the config connector machine state of ApihubSpec
type ApihubSpecStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ApihubSpec resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ApihubSpecObservedState `json:"observedState,omitempty"`
}

// ApihubSpecObservedState is the state of the ApihubSpec resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apihub.v1.Spec
type ApihubSpecObservedState struct {
	// Required. The type of spec.
	//  The value should be one of the allowed values defined for
	//  `projects/{project}/locations/{location}/attributes/system-spec-type`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API.
	//
	//  Note, this field is mandatory if content is provided.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.spec_type
	SpecType *AttributeValuesObservedState `json:"specType,omitempty"`

	// Output only. Details parsed from the spec.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.details
	Details *SpecDetailsObservedState `json:"details,omitempty"`

	// Output only. The time at which the spec was created.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the spec was last updated.
	// +kcc:proto:field=google.cloud.apihub.v1.Spec.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapihubspec;gcpapihubspecs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ApihubSpec is the Schema for the ApihubSpec API
// +k8s:openapi-gen=true
type ApihubSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ApihubSpecSpec   `json:"spec,omitempty"`
	Status ApihubSpecStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ApihubSpecList contains a list of ApihubSpec
type ApihubSpecList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApihubSpec `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApihubSpec{}, &ApihubSpecList{})
}
