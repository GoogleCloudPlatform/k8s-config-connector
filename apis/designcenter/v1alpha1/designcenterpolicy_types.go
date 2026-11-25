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

// API sources for DesignCenterPolicy, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/547f8fdb4b91dbac70d540b3757520f0608ae356/google/cloud/designcenter/v1/best_practices.proto
// +kcc:guess=source-link reason=verify-service-docs-link
// +kcc:source:service-docs=
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DesignCenterPolicyGVK = GroupVersion.WithKind("DesignCenterPolicy")

// DesignCenterPolicySpec defines the desired state of DesignCenterPolicy
// +kcc:spec:proto=google.cloud.designcenter.v1.Policy
// +kcc:required-from-proto
type DesignCenterPolicySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/spaces/{space}/applicationTemplates/{application_template} this resource belongs to.
	// +kcc:guess
	// ApplicationTemplateRef *DesignCenterApplicationTemplateRef `json:"applicationTemplateRef,omitempty"`

	// The DesignCenterPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. URI of the policy.
	// +kcc:proto:field=google.cloud.designcenter.v1.Policy.policy_uri
	PolicyURI *string `json:"policyURI,omitempty"`

	// Optional. URI of the policy revision.
	// +kcc:proto:field=google.cloud.designcenter.v1.Policy.policy_revision_uri
	PolicyRevisionURI *string `json:"policyRevisionURI,omitempty"`

	// Optional. The policy display name.
	// +kcc:proto:field=google.cloud.designcenter.v1.Policy.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The policy description.
	// +kcc:proto:field=google.cloud.designcenter.v1.Policy.description
	Description *string `json:"description,omitempty"`

	// Required. The policy type.
	//  Policy type is used to determine the domain of the policy.
	// +kcc:proto:field=google.cloud.designcenter.v1.Policy.policy_type
	// +required
	PolicyType *string `json:"policyType,omitempty"`

	// Optional. Conditions for applying the policy. Condition accepts the
	//  `criticality` and `environment` parameters from the application template.
	// +kcc:proto:field=google.cloud.designcenter.v1.Policy.apply_conditions
	ApplyConditions []ApplyCondition `json:"applyConditions,omitempty"`

	// Optional. Input only. Parameter values for the cloud controls of the
	//  policy's framework, grouped by cloud control. Only honored when framework
	//  parameter support is enabled for the project; ignored otherwise.
	//
	//  The values are saved with the application template and are not returned on
	//  the policy. On update, they replace the saved values. If they are empty,
	//  the saved values are cleared if `cloud_control_configs` is named in the
	//  update mask. Otherwise, the saved values are kept, unless the policy's
	//  framework changes.
	// +kcc:proto:field=google.cloud.designcenter.v1.Policy.cloud_control_configs
	CloudControlConfigs []CloudControlConfig `json:"cloudControlConfigs,omitempty"`
}

// DesignCenterPolicyStatus defines the config connector machine state of DesignCenterPolicy
type DesignCenterPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DesignCenterPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DesignCenterPolicyObservedState `json:"observedState,omitempty"`
}

// DesignCenterPolicyObservedState is the state of the DesignCenterPolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.designcenter.v1.Policy
type DesignCenterPolicyObservedState struct {
	// Output only. The policy creation timestamp.
	// +kcc:proto:field=google.cloud.designcenter.v1.Policy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The policy update timestamp.
	// +kcc:proto:field=google.cloud.designcenter.v1.Policy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Input only. Parameter values for the cloud controls of the
	//  policy's framework, grouped by cloud control. Only honored when framework
	//  parameter support is enabled for the project; ignored otherwise.
	//
	//  The values are saved with the application template and are not returned on
	//  the policy. On update, they replace the saved values. If they are empty,
	//  the saved values are cleared if `cloud_control_configs` is named in the
	//  update mask. Otherwise, the saved values are kept, unless the policy's
	//  framework changes.
	// +kcc:proto:field=google.cloud.designcenter.v1.Policy.cloud_control_configs
	CloudControlConfigs []CloudControlConfigObservedState `json:"cloudControlConfigs,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdesigncenterpolicy;gcpdesigncenterpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DesignCenterPolicy is the Schema for the DesignCenterPolicy API
// +k8s:openapi-gen=true
type DesignCenterPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DesignCenterPolicySpec   `json:"spec,omitempty"`
	Status DesignCenterPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DesignCenterPolicyList contains a list of DesignCenterPolicy
type DesignCenterPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DesignCenterPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DesignCenterPolicy{}, &DesignCenterPolicyList{})
}
