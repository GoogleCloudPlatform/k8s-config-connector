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

// API sources for ApigeeregistryAPISpec, recorded by generate-types:
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

var ApigeeregistryAPISpecGVK = GroupVersion.WithKind("ApigeeregistryAPISpec")

// ApigeeregistryAPISpecSpec defines the desired state of ApigeeregistryAPISpec
// +kcc:spec:proto=google.cloud.apigeeregistry.v1.ApiSpec
// +kcc:required-from-proto
type ApigeeregistryAPISpecSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ApigeeregistryAPISpec name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// A possibly-hierarchical name used to refer to the spec from other specs.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.filename
	Filename *string `json:"filename,omitempty"`

	// A detailed description.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.description
	Description *string `json:"description,omitempty"`

	// A style (format) descriptor for this spec that is specified as a Media Type
	//  (https://en.wikipedia.org/wiki/Media_type). Possible values include
	//  `application/vnd.apigee.proto`, `application/vnd.apigee.openapi`, and
	//  `application/vnd.apigee.graphql`, with possible suffixes representing
	//  compression types. These hypothetical names are defined in the vendor tree
	//  defined in RFC6838 (https://tools.ietf.org/html/rfc6838) and are not final.
	//  Content types can specify compression. Currently only GZip compression is
	//  supported (indicated with "+gzip").
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// The original source URI of the spec (if one exists).
	//  This is an external location that can be used for reference purposes
	//  but which may not be authoritative since this external resource may
	//  change after the spec is retrieved.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.source_uri
	SourceURI *string `json:"sourceURI,omitempty"`

	// Input only. The contents of the spec.
	//  Provided by API callers when specs are created or updated.
	//  To access the contents of a spec, use GetApiSpecContents.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.contents
	Contents []byte `json:"contents,omitempty"`

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
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations attach non-identifying metadata to resources.
	//
	//  Annotation keys and values are less restricted than those of labels, but
	//  should be generally used for small values of broad interest. Larger, topic-
	//  specific metadata should be stored in Artifacts.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// ApigeeregistryAPISpecStatus defines the config connector machine state of ApigeeregistryAPISpec
type ApigeeregistryAPISpecStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ApigeeregistryAPISpec resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ApigeeregistryAPISpecObservedState `json:"observedState,omitempty"`
}

// ApigeeregistryAPISpecObservedState is the state of the ApigeeregistryAPISpec resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apigeeregistry.v1.ApiSpec
type ApigeeregistryAPISpecObservedState struct {
	// Output only. Immutable. The revision ID of the spec.
	//  A new revision is committed whenever the spec contents are changed.
	//  The format is an 8-character hexadecimal string.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. Creation timestamp; when the spec resource was created.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Revision creation timestamp; when the represented revision was created.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`

	// Output only. Last update timestamp: when the represented revision was last modified.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.revision_update_time
	RevisionUpdateTime *string `json:"revisionUpdateTime,omitempty"`

	// Output only. The size of the spec file in bytes. If the spec is gzipped, this is the
	//  size of the uncompressed spec.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.size_bytes
	SizeBytes *int32 `json:"sizeBytes,omitempty"`

	// Output only. A SHA-256 hash of the spec's contents. If the spec is gzipped, this is
	//  the hash of the uncompressed spec.
	// +kcc:proto:field=google.cloud.apigeeregistry.v1.ApiSpec.hash
	Hash *string `json:"hash,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeregistryapispec;gcpapigeeregistryapispecs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ApigeeregistryAPISpec is the Schema for the ApigeeregistryAPISpec API
// +k8s:openapi-gen=true
type ApigeeregistryAPISpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ApigeeregistryAPISpecSpec   `json:"spec,omitempty"`
	Status ApigeeregistryAPISpecStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ApigeeregistryAPISpecList contains a list of ApigeeregistryAPISpec
type ApigeeregistryAPISpecList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeregistryAPISpec `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeregistryAPISpec{}, &ApigeeregistryAPISpecList{})
}
