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

var LicenseManagerConfigurationGVK = GroupVersion.WithKind("LicenseManagerConfiguration")

// LicenseManagerConfigurationSpec defines the desired state of LicenseManagerConfiguration
// +kcc:spec:proto=google.cloud.licensemanager.v1.Configuration
type LicenseManagerConfigurationSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The LicenseManagerConfiguration name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// +kubebuilder:validation:Required
	// Required. User given name.
	DisplayName *string `json:"displayName,omitempty"`

	// +kubebuilder:validation:Required
	// Required. Name field (with URL) of the Product offered for SPLA.
	Product *string `json:"product,omitempty"`

	// +kubebuilder:validation:Required
	// Required. LicenseType to be applied for billing.
	LicenseType *string `json:"licenseType,omitempty"`

	// +kubebuilder:validation:Required
	// Required. Billing information applicable till end of the current month.
	CurrentBillingInfo *BillingInfo `json:"currentBillingInfo,omitempty"`

	// +kubebuilder:validation:Required
	// Required. Billing information applicable for next month.
	NextBillingInfo *BillingInfo `json:"nextBillingInfo,omitempty"`

	// +kubebuilder:validation:Optional
	// Optional. Labels as key value pairs.
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.licensemanager.v1.BillingInfo
type BillingInfo struct {
	// +kubebuilder:validation:Required
	// Required. This type of billing uses user count for computing total charge.
	UserCountBilling *UserCountBillingInfo `json:"userCountBilling,omitempty"`
}

// +kcc:proto=google.cloud.licensemanager.v1.UserCountBillingInfo
type UserCountBillingInfo struct {
	// +kubebuilder:validation:Required
	// Required. Number of users to bill for.
	UserCount *int32 `json:"userCount,omitempty"`
}

// LicenseManagerConfigurationStatus defines the config connector machine state of LicenseManagerConfiguration
type LicenseManagerConfigurationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the LicenseManagerConfiguration resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *LicenseManagerConfigurationObservedState `json:"observedState,omitempty"`
}

// LicenseManagerConfigurationObservedState is the state of the LicenseManagerConfiguration resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.licensemanager.v1.Configuration
type LicenseManagerConfigurationObservedState struct {
	// Required. Billing information applicable till end of the current month.
	CurrentBillingInfo *BillingInfoObservedState `json:"currentBillingInfo,omitempty"`

	// Required. Billing information applicable for next month.
	NextBillingInfo *BillingInfoObservedState `json:"nextBillingInfo,omitempty"`

	// Output only. [Output only] Create time stamp.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the configuration.
	State *string `json:"state,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.licensemanager.v1.BillingInfo
type BillingInfoObservedState struct {
	// Output only. When the billing starts.
	StartTime *string `json:"startTime,omitempty"`

	// Output only. When the billing ends.
	EndTime *string `json:"endTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcplicensemanagerconfiguration;gcplicensemanagerconfigurations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// LicenseManagerConfiguration is the Schema for the LicenseManagerConfiguration API
// +k8s:openapi-gen=true
type LicenseManagerConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   LicenseManagerConfigurationSpec   `json:"spec,omitempty"`
	Status LicenseManagerConfigurationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// LicenseManagerConfigurationList contains a list of LicenseManagerConfiguration
type LicenseManagerConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LicenseManagerConfiguration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LicenseManagerConfiguration{}, &LicenseManagerConfigurationList{})
}
