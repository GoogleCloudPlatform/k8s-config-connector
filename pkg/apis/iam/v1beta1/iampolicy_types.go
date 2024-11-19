// Copyright 2022 Google LLC
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
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Member string

// Specifies the members to bind to an IAM role.
type IAMPolicyBinding struct {
	// Optional. The list of IAM users to be bound to the role.
	Members []Member `json:"members,omitempty"`
	// Required. The role to bind the users to.
	// +kubebuilder:validation:Pattern=^((projects|organizations)/[^/]+/)?roles/[\w_\.]+$
	Role string `json:"role"`
	// Optional. The condition under which the binding applies.
	Condition *IAMCondition `json:"condition,omitempty"`
}

// Specifies the Cloud Audit Logs configuration for the
// IAM policy.
type IAMPolicyAuditConfig struct {
	// Required. The service for which to enable Data Access audit logs. The
	// special value 'allServices' covers all services. Note that if there are
	// audit configs covering both 'allServices' and a specific service, then
	// the union of the two audit configs is used for that service: the
	// 'logTypes' specified in each 'auditLogConfig' are enabled, and the
	// 'exemptedMembers' in each 'auditLogConfig' are exempted.
	Service string `json:"service"`
	// Required. The configuration for logging of each type of permission.
	AuditLogConfigs []AuditLogConfig `json:"auditLogConfigs"`
}

// IAMPolicySpec defines the desired state of IAMPolicy
type IAMPolicySpec struct {
	// Immutable. Required. The GCP resource to set the IAM policy on (e.g.
	// organization, project...)
	ResourceReference ResourceReference `json:"resourceRef"`
	// Optional. The list of IAM bindings.
	Bindings []IAMPolicyBinding `json:"bindings,omitempty"`
	// Optional. The list of IAM audit configs.
	AuditConfigs []IAMPolicyAuditConfig `json:"auditConfigs,omitempty"`
	// Etag is used for concurrency control, and ensures that policies are updated consistently.
	// Note that this field is not exposed in the CRD's OpenAPI schema.
	Etag string `json:"-"`
}

// IAMPolicyStatus defines the observed state of IAMPolicy
type IAMPolicyStatus struct {
	// Conditions represent the latest available observations of the IAM
	// policy's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller.
	// If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAMPolicy is the Schema for the iampolicies API
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=".status.conditions[?(@.type=='Ready')].status",description="When 'True' the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=".status.conditions[?(@.type=='Ready')].reason",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",type="date",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime"
// +kubebuilder:subresource:status
// +k8s:openapi-gen=true
type IAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMPolicySpec   `json:"spec,omitempty"`
	Status IAMPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAMPolicyList contains a list of IAMPolicy
type IAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMPolicy `json:"items"`
}

const IAMPolicyReconcileInterval = 10 * time.Minute

func init() {
	SchemeBuilder.Register(&IAMPolicy{}, &IAMPolicyList{})
}
