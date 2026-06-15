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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GrafeasNoteGVK = GroupVersion.WithKind("GrafeasNote")

// GrafeasNoteSpec defines the desired state of GrafeasNote
// +kcc:spec:proto=grafeas.v1.Note
type GrafeasNoteSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The GrafeasNote name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// A one sentence description of this note.
	// +kcc:proto:field=grafeas.v1.Note.short_description
	// +optional
	ShortDescription *string `json:"shortDescription,omitempty"`

	// A detailed description of this note.
	// +kcc:proto:field=grafeas.v1.Note.long_description
	// +optional
	LongDescription *string `json:"longDescription,omitempty"`

	// URLs associated with this note.
	// +kcc:proto:field=grafeas.v1.Note.related_url
	// +optional
	RelatedURL []RelatedURL `json:"relatedURL,omitempty"`

	// Time of expiration for this note. Empty if note does not expire.
	// +kcc:proto:field=grafeas.v1.Note.expiration_time
	// +optional
	ExpirationTime *string `json:"expirationTime,omitempty"`

	// Other notes related to this note.
	// +kcc:proto:field=grafeas.v1.Note.related_note_names
	// +optional
	RelatedNoteNames []string `json:"relatedNoteNames,omitempty"`

	// A note describing a package vulnerability.
	// +kcc:proto:field=grafeas.v1.Note.vulnerability
	// +optional
	Vulnerability *VulnerabilityNote `json:"vulnerability,omitempty"`

	// A note describing build provenance for a verifiable build.
	// +kcc:proto:field=grafeas.v1.Note.build
	// +optional
	Build *BuildNote `json:"build,omitempty"`

	// A note describing a base image.
	// +kcc:proto:field=grafeas.v1.Note.image
	// +optional
	Image *ImageNote `json:"image,omitempty"`

	// A note describing a package hosted by various package managers.
	// +kcc:proto:field=grafeas.v1.Note.package
	// +optional
	Package *PackageNote `json:"package,omitempty"`

	// A note describing something that can be deployed.
	// +kcc:proto:field=grafeas.v1.Note.deployment
	// +optional
	Deployment *DeploymentNote `json:"deployment,omitempty"`

	// A note describing the initial analysis of a resource.
	// +kcc:proto:field=grafeas.v1.Note.discovery
	// +optional
	Discovery *DiscoveryNote `json:"discovery,omitempty"`

	// A note describing an attestation role.
	// +kcc:proto:field=grafeas.v1.Note.attestation
	// +optional
	Attestation *AttestationNote `json:"attestation,omitempty"`

	// A note describing available package upgrades.
	// +kcc:proto:field=grafeas.v1.Note.upgrade
	// +optional
	Upgrade *UpgradeNote `json:"upgrade,omitempty"`

	// A note describing a compliance check.
	// +kcc:proto:field=grafeas.v1.Note.compliance
	// +optional
	Compliance *ComplianceNote `json:"compliance,omitempty"`

	// A note describing a dsse attestation note.
	// +kcc:proto:field=grafeas.v1.Note.dsse_attestation
	// +optional
	DsseAttestation *DsseAttestationNote `json:"dsseAttestation,omitempty"`

	// A note describing a vulnerability assessment.
	// +kcc:proto:field=grafeas.v1.Note.vulnerability_assessment
	// +optional
	VulnerabilityAssessment *VulnerabilityAssessmentNote `json:"vulnerabilityAssessment,omitempty"`

	// A note describing an SBOM reference.
	// +kcc:proto:field=grafeas.v1.Note.sbom_reference
	// +optional
	SbomReference *SbomReferenceNote `json:"sbomReference,omitempty"`

	// A note describing a secret.
	// +kcc:proto:field=grafeas.v1.Note.secret
	// +optional
	Secret *SecretNote `json:"secret,omitempty"`
}

// GrafeasNoteStatus defines the config connector machine state of GrafeasNote
type GrafeasNoteStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GrafeasNote resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GrafeasNoteObservedState `json:"observedState,omitempty"`
}

// GrafeasNoteObservedState is the state of the GrafeasNote resource as most recently observed in GCP.
// +kcc:observedstate:proto=grafeas.v1.Note
type GrafeasNoteObservedState struct {
	// Output only. The name of the note in the form of
	//  `projects/[PROVIDER_ID]/notes/[NOTE_ID]`.
	// +kcc:proto:field=grafeas.v1.Note.name
	// +optional
	Name *string `json:"name,omitempty"`

	// Output only. The type of analysis. This field can be used as a filter in
	//  list requests.
	// +kcc:proto:field=grafeas.v1.Note.kind
	// +optional
	Kind *string `json:"kind,omitempty"`

	// Output only. The time this note was created. This field can be used as a
	//  filter in list requests.
	// +kcc:proto:field=grafeas.v1.Note.create_time
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time this note was last updated. This field can be used as
	//  a filter in list requests.
	// +kcc:proto:field=grafeas.v1.Note.update_time
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgrafeasnote;gcpgrafeasnotes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GrafeasNote is the Schema for the GrafeasNote API
// +k8s:openapi-gen=true
type GrafeasNote struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GrafeasNoteSpec   `json:"spec,omitempty"`
	Status GrafeasNoteStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GrafeasNoteList contains a list of GrafeasNote
type GrafeasNoteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GrafeasNote `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GrafeasNote{}, &GrafeasNoteList{})
}
