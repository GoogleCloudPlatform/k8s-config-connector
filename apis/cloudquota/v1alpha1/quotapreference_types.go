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

package v1alpha1

import (
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var APIQuotaPreferenceGVK = GroupVersion.WithKind("APIQuotaPreference")

// APIQuotaPreferenceSpec defines the desired state of APIQuotaPreference

// Parent holds the parent object reference
type Parent struct {
	// +optional
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef,omitempty"`
	// +optional
	OrganizationRef *refv1beta1.OrganizationRef `json:"organizationRef,omitempty"`
	// +optional
	FolderRef *refv1beta1.FolderRef `json:"folderRef,omitempty"`
}

// +kcc:proto=google.api.cloudquotas.v1beta.QuotaPreference
type APIQuotaPreferenceSpec struct {
	// Parent reference
	Parent Parent `json:",inline"`
	// The APIQuotaPreference name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The dimensions that this quota preference applies to. The key of
	//  the map entry is the name of a dimension, such as "region", "zone",
	//  "network_id", and the value of the map entry is the dimension value.
	//
	//  If a dimension is missing from the map of dimensions, the quota preference
	//  applies to all the dimension values except for those that have other quota
	//  preferences configured for the specific value.
	//
	//  NOTE: QuotaPreferences can only be applied across all values of "user" and
	//  "resource" dimension. Do not set values for "user" or "resource" in the
	//  dimension map.
	//
	//  Example: {"provider", "Foo Inc"} where "provider" is a service specific
	//  dimension.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.dimensions
	Dimensions map[string]string `json:"dimensions,omitempty"`

	// Required. Preferred quota configuration.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.quota_config
	//+required
	QuotaConfig *QuotaConfig `json:"quotaConfig,omitempty"`

	// Required. The name of the service to which the quota preference is applied.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.service
	//+required
	Service *string `json:"service,omitempty"`

	// Required. The id of the quota to which the quota preference is applied. A
	//  quota name is unique in the service. Example: `CpusPerProjectPerRegion`
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.quota_id
	//+required
	QuotaID *string `json:"quotaID,omitempty"`

	// The reason / justification for this quota preference.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.justification
	Justification *string `json:"justification,omitempty"`

	// Input only. An email address that can be used to contact the the user, in
	//  case Google Cloud needs more information to make a decision before
	//  additional quota can be granted.
	//
	//  When requesting a quota increase, the email address is required.
	//  When requesting a quota decrease, the email address is optional.
	//  For example, the email address is optional when the
	//  `QuotaConfig.preferred_value` is smaller than the
	//  `QuotaDetails.reset_value`.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.contact_email
	ContactEmail *string `json:"contactEmail,omitempty"`
}

// APIQuotaPreferenceStatus defines the config connector machine state of APIQuotaPreference
type APIQuotaPreferenceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the APIQuotaPreference resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *APIQuotaPreferenceObservedState `json:"observedState,omitempty"`
}

// APIQuotaPreferenceObservedState is the state of the APIQuotaPreference resource as most recently observed in GCP.
// +kcc:proto=google.api.cloudquotas.v1beta.QuotaPreference
type APIQuotaPreferenceObservedState struct {
	// Required. Preferred quota configuration.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.quota_config
	QuotaConfig *QuotaConfigObservedState `json:"quotaConfig,omitempty"`

	// Output only. Create time stamp
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Is the quota preference pending Google Cloud approval and
	//  fulfillment.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Optional. The current etag of the quota preference. If an etag is provided
	//  on update and does not match the current server's etag of the quota
	//  preference, the request will be blocked and an ABORTED error will be
	//  returned. See https://google.aip.dev/134#etags for more details on etags.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaPreference.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpapiquotapreference;gcpapiquotapreferences
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// APIQuotaPreference is the Schema for the APIQuotaPreference API
// +k8s:openapi-gen=true
type APIQuotaPreference struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   APIQuotaPreferenceSpec   `json:"spec,omitempty"`
	Status APIQuotaPreferenceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// APIQuotaPreferenceList contains a list of APIQuotaPreference
type APIQuotaPreferenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIQuotaPreference `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIQuotaPreference{}, &APIQuotaPreferenceList{})
}
