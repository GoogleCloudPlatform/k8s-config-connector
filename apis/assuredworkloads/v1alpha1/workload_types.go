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
	billingv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AssuredWorkloadsWorkloadGVK = GroupVersion.WithKind("AssuredWorkloadsWorkload")

// AssuredWorkloadsWorkloadSpec defines the desired state of AssuredWorkloadsWorkload
// +kcc:spec:proto=google.cloud.assuredworkloads.v1.Workload
type AssuredWorkloadsWorkloadSpec struct {
	// The AssuredWorkloadsWorkload name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The user-assigned display name of the Workload.
	//  When present it must be between 4 to 30 characters.
	//  Allowed characters are: lowercase and uppercase letters, numbers,
	//  hyphen, and spaces.
	//
	//  Example: My Workload
	// +kcc:proto:field=google.cloud.assuredworkloads.v1.Workload.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Immutable. Compliance Regime associated with this workload.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1.Workload.compliance_regime
	ComplianceRegime *string `json:"complianceRegime,omitempty"`

	// Optional. The billing account used for the resources which are
	//  direct children of workload. This billing account is initially associated
	//  with the resources created as part of Workload creation.
	//  After the initial creation of these resources, the customer can change
	//  the assigned billing account.
	//  The resource name has the form
	//  `billingAccounts/{billing_account_id}`. For example,
	//  `billingAccounts/012345-567890-ABCDEF`.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1.Workload.billing_account
	BillingAccountRef *billingv1alpha1.BillingAccountRef `json:"billingAccountRef,omitempty"`

	// Optional. ETag of the workload, it is calculated on the basis
	//  of the Workload contents. It will be used in Update & Delete operations.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1.Workload.etag
	// Etag *string `json:"etag,omitempty"`

	// Optional. Labels applied to the workload.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1.Workload.labels
	// Labels map[string]string `json:"labels,omitempty"`

	// Input only. The parent resource for the resources managed by this Assured Workload. May
	//  be either empty or a folder resource which is a child of the
	//  Workload parent. If not specified all resources are created under the
	//  parent organization.
	//  Format:
	//  folders/{folder_id}
	// +kcc:proto:field=google.cloud.assuredworkloads.v1.Workload.provisioned_resources_parent
	// ProvisionedResourcesParent *string `json:"provisionedResourcesParent,omitempty"`

	// DEPRECATED
	// Input only. Settings used to create a CMEK crypto key. When set, a project with a KMS
	//  CMEK key is provisioned.
	//  This field is deprecated as of Feb 28, 2022.
	//  In order to create a Keyring, callers should specify,
	//  ENCRYPTION_KEYS_PROJECT or KEYRING in ResourceSettings.resource_type field.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1.Workload.kms_settings
	// KMSSettings *Workload_KMSSettings `json:"kmsSettings,omitempty"`

	// Input only. Resource properties that are used to customize workload resources.
	//  These properties (such as custom project id) will be used to create
	//  workload resources if possible. This field is optional.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1.Workload.resource_settings
	ResourceSettings []Workload_ResourceSettings `json:"resourceSettings,omitempty"`

	// Optional. Indicates the sovereignty status of the given workload.
	//  Currently meant to be used by Europe/Canada customers.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1.Workload.enable_sovereign_controls
	EnableSovereignControls *bool `json:"enableSovereignControls,omitempty"`

	// Optional. Compliance Regime associated with this workload.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1.Workload.partner
	Partner *string `json:"partner,omitempty"`
}

// AssuredWorkloadsWorkloadStatus defines the config connector machine state of AssuredWorkloadsWorkload
type AssuredWorkloadsWorkloadStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AssuredWorkloadsWorkload resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AssuredWorkloadsWorkloadObservedState `json:"observedState,omitempty"`
}

// AssuredWorkloadsWorkloadObservedState is the state of the AssuredWorkloadsWorkload resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.assuredworkloads.v1.Workload
type AssuredWorkloadsWorkloadObservedState struct {
	// Output only. The resources associated with this workload.
	//  These resources will be created when creating the workload.
	//  If any of the projects already exist, the workload creation will fail.
	//  Always read only.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1.Workload.resources
	Resources []Workload_ResourceInfo `json:"resources,omitempty"`

	// Output only. Immutable. The Workload creation timestamp.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1.Workload.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Represents the KAJ enrollment state of the given workload.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1.Workload.kaj_enrollment_state
	KajEnrollmentState *string `json:"kajEnrollmentState,omitempty"`

	// Output only. Represents the SAA enrollment response of the given workload.
	//  SAA enrollment response is queried during GetWorkload call.
	//  In failure cases, user friendly error message is shown in SAA details page.
	// +kcc:proto:field=google.cloud.assuredworkloads.v1.Workload.saa_enrollment_response
	SaaEnrollmentResponse *Workload_SaaEnrollmentResponse `json:"saaEnrollmentResponse,omitempty"`

	// Output only. Urls for services which are compliant for this Assured Workload, but which
	//  are currently disallowed by the ResourceUsageRestriction org policy.
	//  Invoke RestrictAllowedResources endpoint to allow your project developers
	//  to use these services in their environment."
	// +kcc:proto:field=google.cloud.assuredworkloads.v1.Workload.compliant_but_disallowed_services
	CompliantButDisallowedServices []string `json:"compliantButDisallowedServices,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpassuredworkloadsworkload;gcpassuredworkloadsworkloads
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AssuredWorkloadsWorkload is the Schema for the AssuredWorkloadsWorkload API
// +k8s:openapi-gen=true
type AssuredWorkloadsWorkload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AssuredWorkloadsWorkloadSpec   `json:"spec,omitempty"`
	Status AssuredWorkloadsWorkloadStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AssuredWorkloadsWorkloadList contains a list of AssuredWorkloadsWorkload
type AssuredWorkloadsWorkloadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AssuredWorkloadsWorkload `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AssuredWorkloadsWorkload{}, &AssuredWorkloadsWorkloadList{})
}
