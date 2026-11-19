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

// API sources for ComputePublicAdvertisedPrefix, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/compute/v1/compute.proto
// +kcc:source:service-docs=https://developers.google.com/compute/docs/reference/latest/
// +kcc:source:resource-docs=https://cloud.google.com/compute/docs/reference/rest/v1/publicAdvertisedPrefixes

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputePublicAdvertisedPrefixGVK = GroupVersion.WithKind("ComputePublicAdvertisedPrefix")

// ComputePublicAdvertisedPrefixSpec defines the desired state of ComputePublicAdvertisedPrefix
// +kcc:spec:proto=google.cloud.compute.v1.PublicAdvertisedPrefix
// +kcc:required-from-proto
type ComputePublicAdvertisedPrefixSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The ComputePublicAdvertisedPrefix name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.PublicAdvertisedPrefix.description
	Description *string `json:"description,omitempty"`

	// The address to be used for reverse DNS verification.
	// +kcc:proto:field=google.cloud.compute.v1.PublicAdvertisedPrefix.dns_verification_ip
	DNSVerificationIP *string `json:"dnsVerificationIP,omitempty"`

	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicAdvertisedPrefix. An up-to-date fingerprint must be provided in order to update the PublicAdvertisedPrefix, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a PublicAdvertisedPrefix.
	// +kcc:proto:field=google.cloud.compute.v1.PublicAdvertisedPrefix.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// The address range, in CIDR format, represented by this public advertised prefix.
	// +kcc:proto:field=google.cloud.compute.v1.PublicAdvertisedPrefix.ip_cidr_range
	IPCIDRRange *string `json:"ipCIDRRange,omitempty"`

	// Specifies how child public delegated prefix will be scoped. It could be one of following values: - `REGIONAL`: The public delegated prefix is regional only. The provisioning will take a few minutes. - `GLOBAL`: The public delegated prefix is global only. The provisioning will take ~4 weeks. - `GLOBAL_AND_REGIONAL` [output only]: The public delegated prefixes is BYOIP V1 legacy prefix. This is output only value and no longer supported in BYOIP V2.
	//  Check the PdpScope enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.PublicAdvertisedPrefix.pdp_scope
	PdpScope *string `json:"pdpScope,omitempty"`
}

// ComputePublicAdvertisedPrefixStatus defines the config connector machine state of ComputePublicAdvertisedPrefix
type ComputePublicAdvertisedPrefixStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputePublicAdvertisedPrefix resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputePublicAdvertisedPrefixObservedState `json:"observedState,omitempty"`
}

// ComputePublicAdvertisedPrefixObservedState is the state of the ComputePublicAdvertisedPrefix resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.PublicAdvertisedPrefix
type ComputePublicAdvertisedPrefixObservedState struct {
	// [Output Only] The version of BYOIP API.
	//  Check the ByoipApiVersion enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.PublicAdvertisedPrefix.byoip_api_version
	ByoipAPIVersion *string `json:"byoipAPIVersion,omitempty"`

	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.PublicAdvertisedPrefix.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] The unique identifier for the resource type. The server generates this identifier.
	// +kcc:proto:field=google.cloud.compute.v1.PublicAdvertisedPrefix.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Type of the resource. Always compute#publicAdvertisedPrefix for public advertised prefixes.
	// +kcc:proto:field=google.cloud.compute.v1.PublicAdvertisedPrefix.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] The list of public delegated prefixes that exist for this public advertised prefix.
	// +kcc:proto:field=google.cloud.compute.v1.PublicAdvertisedPrefix.public_delegated_prefixs
	PublicDelegatedPrefixs []PublicAdvertisedPrefixPublicDelegatedPrefix `json:"publicDelegatedPrefixs,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.PublicAdvertisedPrefix.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output Only] The shared secret to be used for reverse DNS verification.
	// +kcc:proto:field=google.cloud.compute.v1.PublicAdvertisedPrefix.shared_secret
	SharedSecret *string `json:"sharedSecret,omitempty"`

	// The status of the public advertised prefix. Possible values include: - `INITIAL`: RPKI validation is complete. - `PTR_CONFIGURED`: User has configured the PTR. - `VALIDATED`: Reverse DNS lookup is successful. - `REVERSE_DNS_LOOKUP_FAILED`: Reverse DNS lookup failed. - `PREFIX_CONFIGURATION_IN_PROGRESS`: The prefix is being configured. - `PREFIX_CONFIGURATION_COMPLETE`: The prefix is fully configured. - `PREFIX_REMOVAL_IN_PROGRESS`: The prefix is being removed.
	//  Check the Status enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.PublicAdvertisedPrefix.status
	Status *string `json:"status,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputepublicadvertisedprefix;gcpcomputepublicadvertisedprefixes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputePublicAdvertisedPrefix is the Schema for the ComputePublicAdvertisedPrefix API
// +k8s:openapi-gen=true
type ComputePublicAdvertisedPrefix struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputePublicAdvertisedPrefixSpec   `json:"spec,omitempty"`
	Status ComputePublicAdvertisedPrefixStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputePublicAdvertisedPrefixList contains a list of ComputePublicAdvertisedPrefix
type ComputePublicAdvertisedPrefixList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputePublicAdvertisedPrefix `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputePublicAdvertisedPrefix{}, &ComputePublicAdvertisedPrefixList{})
}
