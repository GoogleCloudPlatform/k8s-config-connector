// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	addonv1alpha1 "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=validatingwebhookconfigurationcustomizations,scope=Cluster

// ValidatingWebhookConfigurationCustomization is the Schema for customizing the validating webhook
// configurations in config connector.
type ValidatingWebhookConfigurationCustomization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebhookConfigurationCustomizationSpec   `json:"spec"`
	Status WebhookConfigurationCustomizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ValidatingWebhookConfigurationCustomizationList contains a list of ValidatingWebhookConfigurationTimeout.
type ValidatingWebhookConfigurationCustomizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ValidatingWebhookConfigurationCustomization `json:"items"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=mutatingwebhookconfigurationcustomizations,scope=Cluster

// MutatingWebhookConfigurationCustomization is the Schema for customizing the mutating webhook
// configurations in config connector.
type MutatingWebhookConfigurationCustomization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebhookConfigurationCustomizationSpec   `json:"spec"`
	Status WebhookConfigurationCustomizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MutatingWebhookConfigurationCustomizationList contains a list of MutatingWebhookConfigurationTimeout.
type MutatingWebhookConfigurationCustomizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MutatingWebhookConfigurationCustomization `json:"items"`
}

// WebhookConfigurationCustomizationSpec is the specification for customizing the webhooks of a config
// connector webhook configuration.
type WebhookConfigurationCustomizationSpec struct {
	// The list of webhooks whose configuration to be customized.
	// Required
	Webhooks []WebhookCustomizationSpec `json:"webhooks"`
}

// WebhookCustomizationSpec is the specification for customizing for a specific webhook in config connector.
type WebhookCustomizationSpec struct {
	// The name of the webhook. Do not include the `.cnrm.cloud.google.com` suffix.
	// +kubebuilder:validation:Enum=abandon-on-uninstall;deny-immutable-field-updates;deny-unknown-fields;iam-validation;resource-validation;container-annotation-handler;generic-defaulter;iam-defaulter;management-conflict-annotation-defaulter
	// Required
	Name string `json:"name"`
	// TimeoutSeconds customizes the timeout of the webhook.
	// The timeout value must be between 1 and 30 seconds.
	// The default timeout in Kubernetes is 10 seconds.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=30
	// Required
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
}

// WebhookConfigurationCustomizationStatus defines the observed state of ValidatingWebhookConfigurationCustomization and
// MutatingWebhookConfigurationCustomization.
type WebhookConfigurationCustomizationStatus struct {
	addonv1alpha1.CommonStatus `json:",inline"`
}

func (c *ValidatingWebhookConfigurationCustomization) SetCommonStatus(s addonv1alpha1.CommonStatus) {
	c.Status.CommonStatus = s
}

func (c *MutatingWebhookConfigurationCustomization) SetCommonStatus(s addonv1alpha1.CommonStatus) {
	c.Status.CommonStatus = s
}

func init() {
	SchemeBuilder.Register(
		&ValidatingWebhookConfigurationCustomization{},
		&ValidatingWebhookConfigurationCustomizationList{},
		&MutatingWebhookConfigurationCustomization{},
		&MutatingWebhookConfigurationCustomizationList{})
}
