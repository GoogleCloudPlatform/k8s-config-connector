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

// IAMAuditConfigSpec defines the desired state of IAMAuditConfig.
type IAMAuditConfigSpec struct {
	// Immutable. Required. The GCP resource to set the IAMAuditConfig on
	// (e.g. project).
	ResourceReference ResourceReference `json:"resourceRef"`

	// Immutable. Required. The service for which to enable Data Access
	// audit logs. The special value 'allServices' covers all services.
	// Note that if there are audit configs covering both 'allServices' and
	// a specific service, then the union of the two audit configs is used
	// for that service: the 'logTypes' specified in each 'auditLogConfig'
	// are enabled, and the 'exemptedMembers' in each 'auditLogConfig' are
	// exempted.
	Service string `json:"service"`
	// Required. The configuration for logging of each type of permission.
	AuditLogConfigs []AuditLogConfig `json:"auditLogConfigs"`
}

// IAMAuditConfigStatus defines the observed state of IAMAuditConfig.
type IAMAuditConfigStatus struct {
	// Conditions represent the latest available observations of the
	// IAMAuditConfig's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller.
	// If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAMAuditConfig is the schema for the IAM audit logging API.
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=".status.conditions[?(@.type=='Ready')].status",description="When 'True' the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=".status.conditions[?(@.type=='Ready')].reason",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",type="date",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime"
// +kubebuilder:subresource:status
// +k8s:openapi-gen=true
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:resource:categories=gcp,shortName=gcpiamauditconfig;gcpiamauditconfigs
type IAMAuditConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMAuditConfigSpec   `json:"spec,omitempty"`
	Status IAMAuditConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAMAuditConfigList contains a list of IAMAuditConfig.
type IAMAuditConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMAuditConfig `json:"items"`
}

const IAMAuditConfigReconcileInterval = 10 * time.Minute

func init() {
	SchemeBuilder.Register(&IAMAuditConfig{}, &IAMAuditConfigList{})
}
