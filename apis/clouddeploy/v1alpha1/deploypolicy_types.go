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

var DeployDeployPolicyGVK = GroupVersion.WithKind("CloudDeployDeployPolicy")

type Parent struct {
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`
}

// DeployPolicySpec defines the desired state of DeployDeployPolicy
// +kcc:spec:proto=google.cloud.deploy.v1.DeployPolicy
type DeployPolicySpec struct {
	Parent `json:",inline"`

	// The DeployDeployPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Description of the `DeployPolicy`. Max length is 255 characters.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.description
	Description *string `json:"description,omitempty"`

	// NOT YET
	// // User annotations. These attributes can only be set and used by the
	// //  user, and not by Cloud Deploy. Annotations must meet the following
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
	// //  See
	// //  https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/#syntax-and-character-set
	// //  for more details.
	// // +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.annotations
	// Annotations map[string]string `json:"annotations,omitempty"`

	// // Labels are attributes that can be set and used by both the
	// //  user and by Cloud Deploy. Labels must meet the following constraints:
	// //
	// //  * Keys and values can contain only lowercase letters, numeric characters,
	// //  underscores, and dashes.
	// //  * All characters must use UTF-8 encoding, and international characters are
	// //  allowed.
	// //  * Keys must start with a lowercase letter or international character.
	// //  * Each resource is limited to a maximum of 64 labels.
	// //
	// //  Both keys and values are additionally constrained to be <= 128 bytes.
	// // +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.labels
	// Labels map[string]string `json:"labels,omitempty"`

	// When suspended, the policy will not prevent actions from occurring, even
	//  if the action violates the policy.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.suspended
	Suspended *bool `json:"suspended,omitempty"`

	// Required. Selected resources to which the policy will be applied. At least
	//  one selector is required. If one selector matches the resource the policy
	//  applies. For example, if there are two selectors and the action being
	//  attempted matches one of them, the policy will apply to that action.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.selectors
	Selectors []DeployPolicyResourceSelector `json:"selectors,omitempty"`

	// Required. Rules to apply. At least one rule must be present.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.rules
	Rules []PolicyRule `json:"rules,omitempty"`

	// NOT YET
	// // The weak etag of the `Automation` resource.
	// //  This checksum is computed by the server based on the value of other
	// //  fields, and may be sent on update and delete requests to ensure the
	// //  client has an up-to-date value before proceeding.
	// // +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.etag
	// Etag *string `json:"etag,omitempty"`
}

// DeployPolicyStatus defines the config connector machine state of DeployDeployPolicy
type DeployPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DeployDeployPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DeployPolicyObservedState `json:"observedState,omitempty"`
}

// DeployPolicyObservedState is the state of the DeployDeployPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.deploy.v1.DeployPolicy
type DeployPolicyObservedState struct {
	// Output only. Name of the `DeployPolicy`. Format is
	//  `projects/{project}/locations/{location}/deployPolicies/{deployPolicy}`.
	//  The `deployPolicy` component must match `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.name
	Name *string `json:"name,omitempty"`

	// Output only. Unique identifier of the `DeployPolicy`.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Time at which the deploy policy was created.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Most recent time at which the deploy policy was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.DeployPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcodedeploydeploypolicy;gcpcodedeploydeploypolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudDeployDeployPolicy is the Schema for the CloudDeployDeployPolicy API
// +k8s:openapi-gen=true
type CloudDeployDeployPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DeployPolicySpec   `json:"spec,omitempty"`
	Status DeployPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudDeployDeployPolicyList contains a list of DeployDeployPolicy
type CloudDeployDeployPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudDeployDeployPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudDeployDeployPolicy{}, &CloudDeployDeployPolicyList{})
}
