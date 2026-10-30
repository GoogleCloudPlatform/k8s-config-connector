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

// API sources for ArtifactRegistryAttachment, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/devtools/artifactregistry/v1/attachment.proto
// +kcc:source:service-docs=https://cloud.google.com/artifacts/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.attachments

package v1alpha1

import (
	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ArtifactRegistryAttachmentGVK = GroupVersion.WithKind("ArtifactRegistryAttachment")

// ArtifactRegistryAttachmentSpec defines the desired state of ArtifactRegistryAttachment
// +kcc:spec:proto=google.devtools.artifactregistry.v1.Attachment
// +kcc:required-from-proto
type ArtifactRegistryAttachmentSpec struct {
	// Required. A reference to the ArtifactRegistryRepository that this attachment belongs to.
	// +required
	RepositoryRef *v1beta1.ArtifactRegistryRepositoryRef `json:"repositoryRef,omitempty"`

	// The ArtifactRegistryAttachment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The target the attachment is for, can be a Version, Package or
	//  Repository. E.g.
	//  `projects/p1/locations/us-central1/repositories/repo1/packages/p1/versions/v1`.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Attachment.target
	// +required
	Target *string `json:"target,omitempty"`

	// Type of attachment.
	//  E.g. `application/vnd.spdx+json`
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Attachment.type
	Type *string `json:"type,omitempty"`

	// The namespace this attachment belongs to.
	//  E.g. If an attachment is created by artifact analysis, namespace is set
	//  to `artifactanalysis.googleapis.com`.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Attachment.attachment_namespace
	AttachmentNamespace *string `json:"attachmentNamespace,omitempty"`

	// Optional. User annotations. These attributes can only be set and used by
	//  the user, and not by Artifact Registry. See
	//  https://google.aip.dev/128#annotations for more details such as format and
	//  size limitations.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Attachment.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Required. The files that belong to this attachment.
	//  If the file ID part contains slashes, they are escaped. E.g.
	//  `projects/p1/locations/us-central1/repositories/repo1/files/sha:<sha-of-file>`.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Attachment.files
	// +required
	Files []string `json:"files,omitempty"`
}

// ArtifactRegistryAttachmentStatus defines the config connector machine state of ArtifactRegistryAttachment
type ArtifactRegistryAttachmentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ArtifactRegistryAttachment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ArtifactRegistryAttachmentObservedState `json:"observedState,omitempty"`
}

// ArtifactRegistryAttachmentObservedState is the state of the ArtifactRegistryAttachment resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.devtools.artifactregistry.v1.Attachment
type ArtifactRegistryAttachmentObservedState struct {
	// Output only. The time when the attachment was created.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Attachment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the attachment was last updated.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Attachment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The name of the OCI version that this attachment created. Only
	//  populated for Docker attachments. E.g.
	//  `projects/p1/locations/us-central1/repositories/repo1/packages/p1/versions/v1`.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Attachment.oci_version_name
	OciVersionName *string `json:"ociVersionName,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpartifactregistryattachment;gcpartifactregistryattachments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ArtifactRegistryAttachment is the Schema for the ArtifactRegistryAttachment API
// +k8s:openapi-gen=true
type ArtifactRegistryAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ArtifactRegistryAttachmentSpec   `json:"spec,omitempty"`
	Status ArtifactRegistryAttachmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ArtifactRegistryAttachmentList contains a list of ArtifactRegistryAttachment
type ArtifactRegistryAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArtifactRegistryAttachment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArtifactRegistryAttachment{}, &ArtifactRegistryAttachmentList{})
}
