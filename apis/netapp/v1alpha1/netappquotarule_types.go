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

// API sources for NetappQuotaRule, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e09e85d32ca349e1b205514a412817a7692595c6/google/cloud/netapp/v1/quota_rule.proto
// +kcc:source:service-docs=https://cloud.google.com/netapp/
// +kcc:source:resource-docs=https://docs.cloud.google.com/netapp/volumes/docs/reference/rest/v1/projects.locations.volumes.quotaRules

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetappQuotaRuleGVK = GroupVersion.WithKind("NetappQuotaRule")

// NetappQuotaRuleSpec defines the desired state of NetappQuotaRule
// +kcc:spec:proto=google.cloud.netapp.v1.QuotaRule
// +kcc:required-from-proto
type NetappQuotaRuleSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The NetappQuotaRule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The quota rule applies to the specified user or group, identified
	//  by a Unix UID/GID, Windows SID, or null for default.
	// +kcc:proto:field=google.cloud.netapp.v1.QuotaRule.target
	Target *string `json:"target,omitempty"`

	// Required. The type of quota rule.
	// +kcc:proto:field=google.cloud.netapp.v1.QuotaRule.type
	// +required
	Type *string `json:"type,omitempty"`

	// Required. The maximum allowed disk space in MiB.
	// +kcc:proto:field=google.cloud.netapp.v1.QuotaRule.disk_limit_mib
	// +required
	DiskLimitMIB *int32 `json:"diskLimitMIB,omitempty"`

	// Optional. Description of the quota rule
	// +kcc:proto:field=google.cloud.netapp.v1.QuotaRule.description
	Description *string `json:"description,omitempty"`

	// Optional. Labels of the quota rule
	// +kcc:proto:field=google.cloud.netapp.v1.QuotaRule.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// NetappQuotaRuleStatus defines the config connector machine state of NetappQuotaRule
type NetappQuotaRuleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetappQuotaRule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetappQuotaRuleObservedState `json:"observedState,omitempty"`
}

// NetappQuotaRuleObservedState is the state of the NetappQuotaRule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.netapp.v1.QuotaRule
type NetappQuotaRuleObservedState struct {
	// Output only. State of the quota rule
	// +kcc:proto:field=google.cloud.netapp.v1.QuotaRule.state
	State *string `json:"state,omitempty"`

	// Output only. State details of the quota rule
	// +kcc:proto:field=google.cloud.netapp.v1.QuotaRule.state_details
	StateDetails *string `json:"stateDetails,omitempty"`

	// Output only. Create time of the quota rule
	// +kcc:proto:field=google.cloud.netapp.v1.QuotaRule.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetappquotarule;gcpnetappquotarules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetappQuotaRule is the Schema for the NetappQuotaRule API
// +k8s:openapi-gen=true
type NetappQuotaRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetappQuotaRuleSpec   `json:"spec,omitempty"`
	Status NetappQuotaRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetappQuotaRuleList contains a list of NetappQuotaRule
type NetappQuotaRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetappQuotaRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetappQuotaRule{}, &NetappQuotaRuleList{})
}
