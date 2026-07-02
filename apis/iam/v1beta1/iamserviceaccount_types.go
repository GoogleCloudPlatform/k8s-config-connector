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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var IAMServiceAccountGVK = GroupVersion.WithKind("IAMServiceAccount")

// IAMServiceAccountSpec defines the desired state of IAMServiceAccount
// +kcc:spec:proto=google.iam.admin.v1.ServiceAccount
type IAMServiceAccountSpec struct {
	/* A text description of the service account. Must be less than or equal to 256 UTF-8 bytes. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Whether the service account is disabled. Defaults to false. */
	// +optional
	Disabled *bool `json:"disabled,omitempty"`

	/* The display name for the service account. Can be updated without creating a new resource. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* Immutable. Optional. The accountId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// IAMServiceAccountStatus defines the config connector machine state of IAMServiceAccount
type IAMServiceAccountStatus struct {
	/* Conditions represent the latest available observations of the
	   IAMServiceAccount's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* The e-mail address of the service account. This value should be referenced from any google_iam_policy data sources that would grant the service account privileges. */
	// +optional
	Email *string `json:"email,omitempty"`

	/* The Identity of the service account in the form 'serviceAccount:{email}'. This value is often used to refer to the service account in order to grant IAM permissions. */
	// +optional
	Member *string `json:"member,omitempty"`

	/* The fully-qualified name of the service account. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* A unique specifier for the IAMServiceAccount resource in GCP. */
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	/* The unique id of the service account. */
	// +optional
	UniqueId *string `json:"uniqueId,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpiamserviceaccount;gcpiamserviceaccounts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// IAMServiceAccount is the Schema for the IAMServiceAccount API
// +k8s:openapi-gen=true
type IAMServiceAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMServiceAccountSpec   `json:"spec,omitempty"`
	Status IAMServiceAccountStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// IAMServiceAccountList contains a list of IAMServiceAccount
type IAMServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMServiceAccount `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IAMServiceAccount{}, &IAMServiceAccountList{})
}
