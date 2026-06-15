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
	analyticsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/analytics/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MarketingPlatformAdminAnalyticsAccountLinkGVK = GroupVersion.WithKind("MarketingPlatformAdminAnalyticsAccountLink")

// MarketingPlatformAdminAnalyticsAccountLinkSpec defines the desired state of MarketingPlatformAdminAnalyticsAccountLink
// +kcc:spec:proto=google.marketingplatform.admin.v1alpha.AnalyticsAccountLink
type MarketingPlatformAdminAnalyticsAccountLinkSpec struct {
	// The Organization that this resource belongs to.
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// Immutable. Required. The resource name of the AnalyticsAdmin API account.
	// The account ID will be used as the ID of this AnalyticsAccountLink
	// resource, which will become the final component of the resource name.
	//
	// Format: analyticsadmin.googleapis.com/accounts/{account_id}
	// +required
	AnalyticsAccountRef *analyticsv1alpha1.AccountRef `json:"analyticsAccountRef"`

	// The MarketingPlatformAdminAnalyticsAccountLink name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// MarketingPlatformAdminAnalyticsAccountLinkStatus defines the config connector machine state of MarketingPlatformAdminAnalyticsAccountLink
type MarketingPlatformAdminAnalyticsAccountLinkStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MarketingPlatformAdminAnalyticsAccountLink resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MarketingPlatformAdminAnalyticsAccountLinkObservedState `json:"observedState,omitempty"`
}

// MarketingPlatformAdminAnalyticsAccountLinkObservedState is the state of the MarketingPlatformAdminAnalyticsAccountLink resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.marketingplatform.admin.v1alpha.AnalyticsAccountLink
type MarketingPlatformAdminAnalyticsAccountLinkObservedState struct {
	// Output only. The human-readable name for the Analytics account.
	// +kcc:proto:field=google.marketingplatform.admin.v1alpha.AnalyticsAccountLink.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. The verification state of the link between the Analytics
	// account and the parent organization.
	// +kcc:proto:field=google.marketingplatform.admin.v1alpha.AnalyticsAccountLink.link_verification_state
	LinkVerificationState *string `json:"linkVerificationState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmarketingplatformadminanalyticsaccountlink;gcpmarketingplatformadminanalyticsaccountlinks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MarketingPlatformAdminAnalyticsAccountLink is the Schema for the MarketingPlatformAdminAnalyticsAccountLink API
// +k8s:openapi-gen=true
type MarketingPlatformAdminAnalyticsAccountLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MarketingPlatformAdminAnalyticsAccountLinkSpec   `json:"spec,omitempty"`
	Status MarketingPlatformAdminAnalyticsAccountLinkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MarketingPlatformAdminAnalyticsAccountLinkList contains a list of MarketingPlatformAdminAnalyticsAccountLink
type MarketingPlatformAdminAnalyticsAccountLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MarketingPlatformAdminAnalyticsAccountLink `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MarketingPlatformAdminAnalyticsAccountLink{}, &MarketingPlatformAdminAnalyticsAccountLinkList{})
}
