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

// API sources for ComputeLicense, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/compute/v1/compute.proto
// +kcc:source:service-docs=https://developers.google.com/compute/docs/reference/latest/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/compute/docs/reference/rest

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeLicenseGVK = GroupVersion.WithKind("ComputeLicense")

// ComputeLicenseSpec defines the desired state of ComputeLicense
// +kcc:spec:proto=google.cloud.compute.v1.License
// +kcc:required-from-proto
type ComputeLicenseSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ComputeLicense name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Specifies licenseCodes of licenses that can replace this license. Note: such replacements are allowed even if removable_from_disk is false.
	// +kcc:proto:field=google.cloud.compute.v1.License.allowed_replacement_licenses
	AllowedReplacementLicenses []string `json:"allowedReplacementLicenses,omitempty"`

	// If true, this license can be appended to an existing disk's set of licenses.
	// +kcc:proto:field=google.cloud.compute.v1.License.appendable_to_disk
	AppendableToDisk *bool `json:"appendableToDisk,omitempty"`

	// [Output Only] Deprecated. This field no longer reflects whether a license charges a usage fee.
	// +kcc:proto:field=google.cloud.compute.v1.License.charges_use_fee
	ChargesUseFee *bool `json:"chargesUseFee,omitempty"`

	// An optional textual description of the resource; provided by the client when the resource is created.
	// +kcc:proto:field=google.cloud.compute.v1.License.description
	Description *string `json:"description,omitempty"`

	// Specifies licenseCodes of licenses that are incompatible with this license. If a license is incompatible with this license, it cannot be attached to the same disk or image.
	// +kcc:proto:field=google.cloud.compute.v1.License.incompatible_licenses
	IncompatibleLicenses []string `json:"incompatibleLicenses,omitempty"`

	// [Output Only] The unique code used to attach this license to images, snapshots, and disks.
	// +kcc:proto:field=google.cloud.compute.v1.License.license_code
	LicenseCode *uint64 `json:"licenseCode,omitempty"`

	// If set, this license will be unable to be removed or replaced once attached to a disk until the minimum_retention period has passed.
	// +kcc:proto:field=google.cloud.compute.v1.License.minimum_retention
	MinimumRetention *Duration `json:"minimumRetention,omitempty"`

	// If true, this license can only be used on VMs on multi tenant nodes.
	// +kcc:proto:field=google.cloud.compute.v1.License.multi_tenant_only
	MultiTenantOnly *bool `json:"multiTenantOnly,omitempty"`

	// If true, indicates this is an OS license. Only one OS license can be attached to a disk or image at a time.
	// +kcc:proto:field=google.cloud.compute.v1.License.os_license
	OSLicense *bool `json:"osLicense,omitempty"`

	// If true, this license can be removed from a disk's set of licenses, with no replacement license needed.
	// +kcc:proto:field=google.cloud.compute.v1.License.removable_from_disk
	RemovableFromDisk *bool `json:"removableFromDisk,omitempty"`

	// Specifies the set of permissible coattached licenseCodes of licenses that satisfy the coattachment requirement of this license. At least one license from the set must be attached to the same disk or image as this license.
	// +kcc:proto:field=google.cloud.compute.v1.License.required_coattached_licenses
	RequiredCoattachedLicenses []string `json:"requiredCoattachedLicenses,omitempty"`

	// [Input Only] Deprecated.
	// +kcc:proto:field=google.cloud.compute.v1.License.resource_requirements
	ResourceRequirements *LicenseResourceRequirements `json:"resourceRequirements,omitempty"`

	// If true, this license can only be used on VMs on sole tenant nodes.
	// +kcc:proto:field=google.cloud.compute.v1.License.sole_tenant_only
	SoleTenantOnly *bool `json:"soleTenantOnly,omitempty"`

	// If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
	// +kcc:proto:field=google.cloud.compute.v1.License.transferable
	Transferable *bool `json:"transferable,omitempty"`

	// [Output Only] Last update timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.License.update_timestamp
	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`
}

// ComputeLicenseStatus defines the config connector machine state of ComputeLicense
type ComputeLicenseStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeLicense resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeLicenseObservedState `json:"observedState,omitempty"`
}

// ComputeLicenseObservedState is the state of the ComputeLicense resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.License
type ComputeLicenseObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.License.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.License.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Type of resource. Always compute#license for licenses.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.License.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.License.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output Only] Server-defined URL for this resource with the resource id.
	// +kcc:guess=placement reason=no-field-behavior-on-message
	// +kcc:proto:field=google.cloud.compute.v1.License.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputelicense;gcpcomputelicenses
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeLicense is the Schema for the ComputeLicense API
// +k8s:openapi-gen=true
type ComputeLicense struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeLicenseSpec   `json:"spec,omitempty"`
	Status ComputeLicenseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeLicenseList contains a list of ComputeLicense
type ComputeLicenseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeLicense `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeLicense{}, &ComputeLicenseList{})
}
