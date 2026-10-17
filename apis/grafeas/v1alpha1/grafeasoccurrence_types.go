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

// API sources for GrafeasOccurrence, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/grafeas/v1/grafeas.proto
// +kcc:source:service-docs=https://cloud.google.com/container-analysis/api/reference/rest/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/container-analysis/api/reference/rest/docs/reference/rest/v1/projects.occurrences

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GrafeasOccurrenceGVK = GroupVersion.WithKind("GrafeasOccurrence")

// GrafeasOccurrenceSpec defines the desired state of GrafeasOccurrence
// +kcc:spec:proto=grafeas.v1.Occurrence
// +kcc:required-from-proto
type GrafeasOccurrenceSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The GrafeasOccurrence name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Immutable. A URI that represents the resource for which the
	//  occurrence applies. For example,
	//  `https://gcr.io/project/image@sha256:123abc` for a Docker image.
	// +kcc:proto:field=grafeas.v1.Occurrence.resource_uri
	ResourceURI *string `json:"resourceURI,omitempty"`

	// Required. Immutable. The analysis note associated with this occurrence, in
	//  the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`. This field can be
	//  used as a filter in list requests.
	// +kcc:proto:field=grafeas.v1.Occurrence.note_name
	NoteName *string `json:"noteName,omitempty"`

	// A description of actions that can be taken to remedy the note.
	// +kcc:proto:field=grafeas.v1.Occurrence.remediation
	Remediation *string `json:"remediation,omitempty"`

	// Describes a security vulnerability.
	// +kcc:proto:field=grafeas.v1.Occurrence.vulnerability
	Vulnerability *VulnerabilityOccurrence `json:"vulnerability,omitempty"`

	// Describes a verifiable build.
	// +kcc:proto:field=grafeas.v1.Occurrence.build
	Build *BuildOccurrence `json:"build,omitempty"`

	// Describes how this resource derives from the basis in the associated
	//  note.
	// +kcc:proto:field=grafeas.v1.Occurrence.image
	Image *ImageOccurrence `json:"image,omitempty"`

	// Describes the installation of a package on the linked resource.
	// +kcc:proto:field=grafeas.v1.Occurrence.package
	Package *PackageOccurrence `json:"package,omitempty"`

	// Describes the deployment of an artifact on a runtime.
	// +kcc:proto:field=grafeas.v1.Occurrence.deployment
	Deployment *DeploymentOccurrence `json:"deployment,omitempty"`

	// Describes when a resource was discovered.
	// +kcc:proto:field=grafeas.v1.Occurrence.discovery
	Discovery *DiscoveryOccurrence `json:"discovery,omitempty"`

	// Describes an attestation of an artifact.
	// +kcc:proto:field=grafeas.v1.Occurrence.attestation
	Attestation *AttestationOccurrence `json:"attestation,omitempty"`

	// Describes an available package upgrade on the linked resource.
	// +kcc:proto:field=grafeas.v1.Occurrence.upgrade
	Upgrade *UpgradeOccurrence `json:"upgrade,omitempty"`

	// Describes a compliance violation on a linked resource.
	// +kcc:proto:field=grafeas.v1.Occurrence.compliance
	Compliance *ComplianceOccurrence `json:"compliance,omitempty"`

	// Describes an attestation of an artifact using dsse.
	// +kcc:proto:field=grafeas.v1.Occurrence.dsse_attestation
	DsseAttestation *DsseAttestationOccurrence `json:"dsseAttestation,omitempty"`

	// Describes a specific SBOM reference occurrences.
	// +kcc:proto:field=grafeas.v1.Occurrence.sbom_reference
	SbomReference *SbomReferenceOccurrence `json:"sbomReference,omitempty"`

	// Describes a secret.
	// +kcc:proto:field=grafeas.v1.Occurrence.secret
	Secret *SecretOccurrence `json:"secret,omitempty"`

	// https://github.com/secure-systems-lab/dsse
	// +kcc:proto:field=grafeas.v1.Occurrence.envelope
	Envelope *Envelope `json:"envelope,omitempty"`
}

// GrafeasOccurrenceStatus defines the config connector machine state of GrafeasOccurrence
type GrafeasOccurrenceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GrafeasOccurrence resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GrafeasOccurrenceObservedState `json:"observedState,omitempty"`
}

// GrafeasOccurrenceObservedState is the state of the GrafeasOccurrence resource as most recently observed in GCP.
// +kcc:observedstate:proto=grafeas.v1.Occurrence
type GrafeasOccurrenceObservedState struct {
	// Output only. This explicitly denotes which of the occurrence details are
	//  specified. This field can be used as a filter in list requests.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=grafeas.v1.Occurrence.kind
	Kind *string `json:"kind,omitempty"`

	// Output only. The time this occurrence was created.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=grafeas.v1.Occurrence.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time this occurrence was last updated.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=grafeas.v1.Occurrence.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Describes a security vulnerability.
	// +kcc:proto:field=grafeas.v1.Occurrence.vulnerability
	Vulnerability *VulnerabilityOccurrenceObservedState `json:"vulnerability,omitempty"`

	// Describes the installation of a package on the linked resource.
	// +kcc:proto:field=grafeas.v1.Occurrence.package
	Package *PackageOccurrenceObservedState `json:"package,omitempty"`

	// Describes when a resource was discovered.
	// +kcc:proto:field=grafeas.v1.Occurrence.discovery
	Discovery *DiscoveryOccurrenceObservedState `json:"discovery,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgrafeasoccurrence;gcpgrafeasoccurrences
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GrafeasOccurrence is the Schema for the GrafeasOccurrence API
// +k8s:openapi-gen=true
type GrafeasOccurrence struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GrafeasOccurrenceSpec   `json:"spec,omitempty"`
	Status GrafeasOccurrenceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GrafeasOccurrenceList contains a list of GrafeasOccurrence
type GrafeasOccurrenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GrafeasOccurrence `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GrafeasOccurrence{}, &GrafeasOccurrenceList{})
}
