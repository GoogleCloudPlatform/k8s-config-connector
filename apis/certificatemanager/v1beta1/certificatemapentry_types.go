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

package v1beta1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CertificateManagerCertificateMapEntryGVK = GroupVersion.WithKind("CertificateManagerCertificateMapEntry")

// CertificateManagerCertificateMapEntrySpec defines the desired state of CertificateManagerCertificateMapEntry
// +kcc:spec:proto=google.cloud.certificatemanager.v1.CertificateMapEntry
type CertificateManagerCertificateMapEntrySpec struct {
	/* Immutable. A list of references to CertificateManagerCertificate resources that will be associated with this map entry. */
	// +required
	CertificatesRefs []CertificateManagerCertificateRef `json:"certificatesRefs"`

	/* A human-readable description of the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
	for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
	selecting a proper certificate. */
	// +optional
	Hostname *string `json:"hostname,omitempty"`

	/* A map entry that is inputted into the certificate map. */
	// +required
	MapRef CertificateManagerCertificateMapRef `json:"mapRef"`

	/* Immutable. A predefined matcher for particular cases, other than SNI selection. */
	// +optional
	Matcher *string `json:"matcher,omitempty"`

	/* The project that this resource belongs to. */
	// +required
	ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// CertificateManagerCertificateMapEntryStatus defines the config connector machine state of CertificateManagerCertificateMapEntry
type CertificateManagerCertificateMapEntryStatus struct {
	/* Conditions represent the latest available observations of the
	   CertificateManagerCertificateMapEntry's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* Creation timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
	with nanosecond resolution and up to nine fractional digits.
	Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* A serving state of this Certificate Map Entry. */
	// +optional
	State *string `json:"state,omitempty"`

	/* Update timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
	with nanosecond resolution and up to nine fractional digits.
	Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcertificatemanagercertificatemapentry;gcpcertificatemanagercertificatemapentries
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="internal.cloud.google.com/additional-versions=v1alpha1"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CertificateManagerCertificateMapEntry is the Schema for the CertificateManagerCertificateMapEntry API
// +k8s:openapi-gen=true
type CertificateManagerCertificateMapEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CertificateManagerCertificateMapEntrySpec   `json:"spec,omitempty"`
	Status CertificateManagerCertificateMapEntryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertificateManagerCertificateMapEntryList contains a list of CertificateManagerCertificateMapEntry
type CertificateManagerCertificateMapEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateManagerCertificateMapEntry `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CertificateManagerCertificateMapEntry{}, &CertificateManagerCertificateMapEntryList{})
}
