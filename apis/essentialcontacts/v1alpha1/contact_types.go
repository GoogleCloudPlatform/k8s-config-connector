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

var EssentialContactsContactGVK = GroupVersion.WithKind("EssentialContactsContact")

type Parent struct {
	// +optional
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef,omitempty"`
	// +optional
	FolderRef *refv1beta1.FolderRef `json:"folderRef,omitempty"`
	// +optional
	OrganizationRef *refv1beta1.OrganizationRef `json:"organizationRef,omitempty"`
}

// EssentialContactsContactSpec defines the desired state of EssentialContactsContact
// +kcc:proto=google.cloud.essentialcontacts.v1.Contact
type EssentialContactsContactSpec struct {
	Parent `json:",inline"`
	// The EssentialContactsContact name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The email address to send notifications to. The email address
	//  does not need to be a Google Account.
	// +kcc:proto:field=google.cloud.essentialcontacts.v1.Contact.email
	Email *string `json:"email,omitempty"`

	// Required. The categories of notifications that the contact will receive
	//  communications for.
	// +kcc:proto:field=google.cloud.essentialcontacts.v1.Contact.notification_category_subscriptions
	NotificationCategorySubscriptions []string `json:"notificationCategorySubscriptions,omitempty"`

	// Required. The preferred language for notifications, as a ISO 639-1 language
	//  code. See [Supported
	//  languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages)
	//  for a list of supported languages.
	// +kcc:proto:field=google.cloud.essentialcontacts.v1.Contact.language_tag
	LanguageTag *string `json:"languageTag,omitempty"`

	// The last time the validation_state was updated, either manually or
	//  automatically. A contact is considered stale if its validation state was
	//  updated more than 1 year ago.
	// +kcc:proto:field=google.cloud.essentialcontacts.v1.Contact.validate_time
	ValidateTime *string `json:"validateTime,omitempty"`
}

// EssentialContactsContactStatus defines the config connector machine state of EssentialContactsContact
type EssentialContactsContactStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the EssentialContactsContact resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *EssentialContactsContactObservedState `json:"observedState,omitempty"`
}

// EssentialContactsContactObservedState is the state of the EssentialContactsContact resource as most recently observed in GCP.
// +kcc:proto=google.cloud.essentialcontacts.v1.Contact
type EssentialContactsContactObservedState struct {
	// Output only. The identifier for the contact.
	//  Format: {resource_type}/{resource_id}/contacts/{contact_id}
	// +kcc:proto:field=google.cloud.essentialcontacts.v1.Contact.name
	Name *string `json:"name,omitempty"`

	// Output only. The validity of the contact. A contact is considered valid if
	//  it is the correct recipient for notifications for a particular resource.
	// +kcc:proto:field=google.cloud.essentialcontacts.v1.Contact.validation_state
	ValidationState *string `json:"validationState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpessentialcontactscontact;gcpessentialcontactscontacts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// EssentialContactsContact is the Schema for the EssentialContactsContact API
// +k8s:openapi-gen=true
type EssentialContactsContact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   EssentialContactsContactSpec   `json:"spec,omitempty"`
	Status EssentialContactsContactStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EssentialContactsContactList contains a list of EssentialContactsContact
type EssentialContactsContactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EssentialContactsContact `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EssentialContactsContact{}, &EssentialContactsContactList{})
}
