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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudNumberRegistryIpamAdminScopeGVK = GroupVersion.WithKind("CloudNumberRegistryIpamAdminScope")

// CloudNumberRegistryIpamAdminScopeSpec defines the desired state of CloudNumberRegistryIpamAdminScope
// +kcc:spec:proto=google.cloud.numberregistry.v1alpha.IpamAdminScope
type CloudNumberRegistryIpamAdminScopeSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The CloudNumberRegistryIpamAdminScope name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Add-on platforms that are enabled for this IpamAdminScope. Cloud
	//  Number Registry only discovers the IP addresses from the enabled platforms.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.IpamAdminScope.enabled_addon_platforms
	// +kubebuilder:validation:Required
	EnabledAddonPlatforms []string `json:"enabledAddonPlatforms,omitempty"`

	// Required. Administrative scopes enabled for IP address discovery and
	//  management. For example, "organizations/1234567890". Minimum of 1 scope is
	//  required. In preview, only one organization scope is allowed.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.IpamAdminScope.scopes
	// +kubebuilder:validation:Required
	Scopes []string `json:"scopes,omitempty"`

	// Optional. User-defined labels.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.IpamAdminScope.labels
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`
}

// CloudNumberRegistryIpamAdminScopeStatus defines the config connector machine state of CloudNumberRegistryIpamAdminScope
type CloudNumberRegistryIpamAdminScopeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudNumberRegistryIpamAdminScope resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudNumberRegistryIpamAdminScopeObservedState `json:"observedState,omitempty"`
}

// CloudNumberRegistryIpamAdminScopeObservedState is the state of the CloudNumberRegistryIpamAdminScope resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.numberregistry.v1alpha.IpamAdminScope
type CloudNumberRegistryIpamAdminScopeObservedState struct {
	// Output only. State of resource discovery pipeline.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.IpamAdminScope.state
	State *string `json:"state,omitempty"`

	// Output only. The time at which the IpamAdminScope was created.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.IpamAdminScope.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the IpamAdminScope was last updated.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.IpamAdminScope.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudnumberregistryipamadminscope;gcpcloudnumberregistryipamadminscopes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudNumberRegistryIpamAdminScope is the Schema for the CloudNumberRegistryIpamAdminScope API
// +k8s:openapi-gen=true
type CloudNumberRegistryIpamAdminScope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudNumberRegistryIpamAdminScopeSpec   `json:"spec,omitempty"`
	Status CloudNumberRegistryIpamAdminScopeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudNumberRegistryIpamAdminScopeList contains a list of CloudNumberRegistryIpamAdminScope
type CloudNumberRegistryIpamAdminScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudNumberRegistryIpamAdminScope `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudNumberRegistryIpamAdminScope{}, &CloudNumberRegistryIpamAdminScopeList{})
}
