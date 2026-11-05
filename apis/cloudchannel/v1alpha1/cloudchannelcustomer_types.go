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

// API sources for CloudchannelCustomer, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/channel/v1/customers.proto
// +kcc:source:service-docs=https://cloud.google.com/channel
// +kcc:source:resource-docs=https://docs.cloud.google.com/channel/docs/reference/rest/v1/accounts.customers

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudchannelCustomerGVK = GroupVersion.WithKind("CloudchannelCustomer")

// CloudchannelCustomerSpec defines the desired state of CloudchannelCustomer
// +kcc:spec:proto=google.cloud.channel.v1.Customer
// +kcc:required-from-proto
type CloudchannelCustomerSpec struct {

	// The CloudchannelCustomer name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Name of the organization that the customer entity represents.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.org_display_name
	// +required
	OrgDisplayName *string `json:"orgDisplayName,omitempty"`

	// Required. The organization address for the customer. To enforce US laws and
	//  embargoes, we require a region, postal code, and address lines. You must
	//  provide valid addresses for every customer. To set the customer's
	//  language, use the Customer-level language code.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.org_postal_address
	// +required
	OrgPostalAddress *PostalAddress `json:"orgPostalAddress,omitempty"`

	// Primary contact info.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.primary_contact_info
	PrimaryContactInfo *ContactInfo `json:"primaryContactInfo,omitempty"`

	// Secondary contact email. You need to provide an alternate email to create
	//  different domains if a primary contact email already exists. Users will
	//  receive a notification with credentials when you create an admin.google.com
	//  account. Secondary emails are also recovery email addresses. Alternate
	//  emails are optional when you create Team customers.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.alternate_email
	AlternateEmail *string `json:"alternateEmail,omitempty"`

	// Required. The customer's primary domain. Must match the primary contact
	//  email's domain.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.domain
	// +required
	Domain *string `json:"domain,omitempty"`

	// Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more
	//  information, see
	//  https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Cloud Identity ID of the customer's channel partner.
	//  Populated only if a channel partner exists for this customer.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.channel_partner_id
	ChannelPartnerID *string `json:"channelPartnerID,omitempty"`

	// Optional. External CRM ID for the customer.
	//  Populated only if a CRM ID exists for this customer.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.correlation_id
	CorrelationID *string `json:"correlationID,omitempty"`

	// Optional. Indicate whether a customer is attesting about the correctness of
	//  provided information. Only required if creating a GCP Entitlement.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.customer_attestation_state
	CustomerAttestationState *string `json:"customerAttestationState,omitempty"`
}

// CloudchannelCustomerStatus defines the config connector machine state of CloudchannelCustomer
type CloudchannelCustomerStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudchannelCustomer resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudchannelCustomerObservedState `json:"observedState,omitempty"`
}

// CloudchannelCustomerObservedState is the state of the CloudchannelCustomer resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.channel.v1.Customer
type CloudchannelCustomerObservedState struct {
	// Primary contact info.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.primary_contact_info
	PrimaryContactInfo *ContactInfoObservedState `json:"primaryContactInfo,omitempty"`

	// Output only. Time when the customer was created.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the customer was updated.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The customer's Cloud Identity ID if the customer has a Cloud
	//  Identity resource.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.cloud_identity_id
	CloudIdentityID *string `json:"cloudIdentityID,omitempty"`

	// Output only. Cloud Identity information for the customer.
	//  Populated only if a Cloud Identity account exists for this customer.
	// +kcc:proto:field=google.cloud.channel.v1.Customer.cloud_identity_info
	CloudIdentityInfo *CloudIdentityInfoObservedState `json:"cloudIdentityInfo,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudchannelcustomer;gcpcloudchannelcustomers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudchannelCustomer is the Schema for the CloudchannelCustomer API
// +k8s:openapi-gen=true
type CloudchannelCustomer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudchannelCustomerSpec   `json:"spec,omitempty"`
	Status CloudchannelCustomerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudchannelCustomerList contains a list of CloudchannelCustomer
type CloudchannelCustomerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudchannelCustomer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudchannelCustomer{}, &CloudchannelCustomerList{})
}
