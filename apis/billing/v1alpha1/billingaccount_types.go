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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BillingAccountGVK = GroupVersion.WithKind("BillingAccount")

// BillingAccountSpec defines the desired state of BillingAccount
// +kcc:spec:proto=google.cloud.billing.v1.BillingAccount
type BillingAccountSpec struct {
	// The BillingAccount name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The display name given to the billing account, such as `My Billing
	//  Account`. This name is displayed in the Google Cloud Console.
	// +kcc:proto:field=google.cloud.billing.v1.BillingAccount.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The currency in which the billing account is billed and charged,
	//  represented as an ISO 4217 code such as `USD`.
	//
	//  Billing account currency is determined at the time of billing account
	//  creation and cannot be updated subsequently, so this field should not be
	//  set on update requests. In addition, a subaccount always matches the
	//  currency of its parent billing account, so this field should not be set on
	//  subaccounts. Clients can read this field to determine the
	//  currency of an existing billing account.
	// +kcc:proto:field=google.cloud.billing.v1.BillingAccount.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// Optional. The billing account's parent resource.
	ParentBillingAccountRef *BillingAccountRef `json:"parent,omitempty"`
}

// BillingAccountStatus defines the config connector machine state of BillingAccount
type BillingAccountStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BillingAccount resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BillingAccountObservedState `json:"observedState,omitempty"`
}

// BillingAccountObservedState is the state of the BillingAccount resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.billing.v1.BillingAccount
type BillingAccountObservedState struct {
	// Output only. The resource name of the billing account. The resource name
	//  has the form `billingAccounts/{billing_account_id}`. For example,
	//  `billingAccounts/012345-567890-ABCDEF` would be the resource name for
	//  billing account `012345-567890-ABCDEF`.
	// +kcc:proto:field=google.cloud.billing.v1.BillingAccount.name
	Name *string `json:"name,omitempty"`

	// Output only. True if the billing account is open, and will therefore be
	//  charged for any usage on associated projects. False if the billing account
	//  is closed, and therefore projects associated with it are unable to use paid
	//  services.
	// +kcc:proto:field=google.cloud.billing.v1.BillingAccount.open
	Open *bool `json:"open,omitempty"`

	// Optional. The currency in which the billing account is billed and charged,
	//  represented as an ISO 4217 code such as `USD`.
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// If this account is a
	//  [subaccount](https://cloud.google.com/billing/docs/concepts), then this
	//  will be the resource name of the parent billing account that it is being
	//  resold through.
	//  Otherwise this will be empty.
	// +kcc:proto:field=google.cloud.billing.v1.BillingAccount.master_billing_account
	MasterBillingAccount *string `json:"masterBillingAccount,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbillingaccount;gcpbillingaccounts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BillingAccount is the Schema for the BillingAccount API
// +k8s:openapi-gen=true
type BillingAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BillingAccountSpec   `json:"spec,omitempty"`
	Status BillingAccountStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BillingAccountList contains a list of BillingAccount
type BillingAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BillingAccount `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BillingAccount{}, &BillingAccountList{})
}
