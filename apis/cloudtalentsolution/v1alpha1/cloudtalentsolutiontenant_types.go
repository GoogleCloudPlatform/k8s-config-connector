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

// API sources for CloudTalentSolutionTenant, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/talent/v4/tenant.proto
// +kcc:source:service-docs=https://cloud.google.com/talent-solution/job-search/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/talent-solution/job-search/docs/reference/rest/v4/projects.tenants

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudTalentSolutionTenantGVK = GroupVersion.WithKind("CloudTalentSolutionTenant")

// CloudTalentSolutionTenantSpec defines the desired state of CloudTalentSolutionTenant
// +kcc:spec:proto=google.cloud.talent.v4.Tenant
// +kcc:required-from-proto
type CloudTalentSolutionTenantSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The CloudTalentSolutionTenant name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Client side tenant identifier, used to uniquely identify the
	//  tenant.
	//
	//  The maximum number of allowed characters is 255.
	// +kcc:proto:field=google.cloud.talent.v4.Tenant.external_id
	// +required
	ExternalID *string `json:"externalID,omitempty"`
}

// CloudTalentSolutionTenantStatus defines the config connector machine state of CloudTalentSolutionTenant
type CloudTalentSolutionTenantStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudTalentSolutionTenant resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *CloudTalentSolutionTenantObservedState `json:"observedState,omitempty"`
}

// CloudTalentSolutionTenantObservedState is the state of the CloudTalentSolutionTenant resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.talent.v4.Tenant
// type CloudTalentSolutionTenantObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudtalentsolutiontenant;gcpcloudtalentsolutiontenants
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudTalentSolutionTenant is the Schema for the CloudTalentSolutionTenant API
// +k8s:openapi-gen=true
type CloudTalentSolutionTenant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudTalentSolutionTenantSpec   `json:"spec,omitempty"`
	Status CloudTalentSolutionTenantStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudTalentSolutionTenantList contains a list of CloudTalentSolutionTenant
type CloudTalentSolutionTenantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudTalentSolutionTenant `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudTalentSolutionTenant{}, &CloudTalentSolutionTenantList{})
}
