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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AutomationGVK = GroupVersion.WithKind("CloudDeployAutomation")

// AutomationSpec defines the desired state of CloudDeployAutomation
// +kcc:proto=google.cloud.deploy.v1.Automation
type AutomationSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// +required
	Location string `json:"location"`

	// The CloudDeployAutomation name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Description of the `Automation`. Max length is 255 characters.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.description
	Description *string `json:"description,omitempty"`

	// NOT YET
	// // Optional. User annotations. These attributes can only be set and used by
	// //  the user, and not by Cloud Deploy. Annotations must meet the following
	// //  constraints:
	// //
	// //  * Annotations are key/value pairs.
	// //  * Valid annotation keys have two segments: an optional prefix and name,
	// //  separated by a slash (`/`).
	// //  * The name segment is required and must be 63 characters or less,
	// //  beginning and ending with an alphanumeric character (`[a-z0-9A-Z]`) with
	// //  dashes (`-`), underscores (`_`), dots (`.`), and alphanumerics between.
	// //  * The prefix is optional. If specified, the prefix must be a DNS subdomain:
	// //  a series of DNS labels separated by dots(`.`), not longer than 253
	// //  characters in total, followed by a slash (`/`).
	// //
	// // +kcc:proto:field=google.cloud.deploy.v1.Automation.annotations
	// Annotations map[string]string `json:"annotations,omitempty"`

	// // Optional. Labels are attributes that can be set and used by both the
	// //  user and by Cloud Deploy. Labels must meet the following constraints:
	// //
	// //  * Keys and values can contain only lowercase letters, numeric characters,
	// //  underscores, and dashes.
	// //  * All characters must use UTF-8 encoding, and international characters are
	// //  allowed.
	// //  * Keys must start with a lowercase letter or international character.
	// //  * Each resource is limited to a maximum of 64 labels.
	// //
	// //  Both keys and values are additionally constrained to be <= 63 characters.
	// // +kcc:proto:field=google.cloud.deploy.v1.Automation.labels
	// Labels map[string]string `json:"labels,omitempty"`

	// // Optional. The weak etag of the `Automation` resource.
	// //  This checksum is computed by the server based on the value of other
	// //  fields, and may be sent on update and delete requests to ensure the
	// //  client has an up-to-date value before proceeding.
	// // +kcc:proto:field=google.cloud.deploy.v1.Automation.etag
	// Etag *string `json:"etag,omitempty"`

	// Optional. When Suspended, automation is deactivated from execution.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.suspended
	Suspended *bool `json:"suspended,omitempty"`

	// Required. Email address of the user-managed IAM service account that
	//  creates Cloud Deploy release and rollout resources.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.service_account
	ServiceAccountRef *refs.IAMServiceAccountRef `json:"serviceAccount,omitempty"`

	// Required. Selected resources to which the automation will be applied.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.selector
	Selector *AutomationResourceSelector `json:"selector,omitempty"`

	// Required. List of Automation rules associated with the Automation resource.
	//  Must have at least one rule and limited to 250 rules per Delivery Pipeline.
	//  Note: the order of the rules here is not the same as the order of
	//  execution.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.rules
	Rules []AutomationRule `json:"rules,omitempty"`
}

// AutomationStatus defines the config connector machine state of CloudDeployAutomation
type AutomationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudDeployAutomation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudDeployAutomationObservedState `json:"observedState,omitempty"`
}

// CloudDeployAutomationObservedState is the state of the CloudDeployAutomation resource as most recently observed in GCP.
// +kcc:proto=google.cloud.deploy.v1.Automation
type CloudDeployAutomationObservedState struct {
	// Output only. Time at which the automation was created.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time at which the automation was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.Automation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// NOT YET
	// // Required. List of Automation rules associated with the Automation resource.
	// //  Must have at least one rule and limited to 250 rules per Delivery Pipeline.
	// //  Note: the order of the rules here is not the same as the order of
	// //  execution.
	// // +kcc:proto:field=google.cloud.deploy.v1.Automation.rules
	// Rules []AutomationRuleObservedState `json:"rules,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddeployautomation;gcpclouddeployautomations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudDeployAutomation is the Schema for the CloudDeployAutomation API
// +k8s:openapi-gen=true
type CloudDeployAutomation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AutomationSpec   `json:"spec,omitempty"`
	Status AutomationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudDeployAutomationList contains a list of CloudDeployAutomation
type CloudDeployAutomationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudDeployAutomation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudDeployAutomation{}, &CloudDeployAutomationList{})
}
