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

var NetworkSecurityInterceptDeploymentGVK = GroupVersion.WithKind("NetworkSecurityInterceptDeployment")

// NetworkSecurityInterceptDeploymentSpec defines the desired state of NetworkSecurityInterceptDeployment
// +kcc:spec:proto=google.cloud.networksecurity.v1.InterceptDeployment
type NetworkSecurityInterceptDeploymentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The NetworkSecurityInterceptDeployment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. User-provided description of the deployment.
	// Used as additional context for the deployment.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Required. Immutable. The regional forwarding rule that fronts the interceptors, for example: `projects/123456789/regions/us-central1/forwardingRules/my-rule`.
	// +kubebuilder:validation:Required
	ForwardingRuleRef *refsv1beta1.ComputeForwardingRuleRef `json:"forwardingRuleRef"`

	// Required. Immutable. The deployment group that this deployment is a part of, for example: `projects/123456789/locations/global/interceptDeploymentGroups/my-dg`.
	// +kubebuilder:validation:Required
	InterceptDeploymentGroupRef *refsv1beta1.NetworkSecurityInterceptDeploymentGroupRef `json:"interceptDeploymentGroupRef"`

	// Optional. Labels are key/value pairs that help to organize and filter resources.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`
}

// NetworkSecurityInterceptDeploymentStatus defines the config connector machine state of NetworkSecurityInterceptDeployment
type NetworkSecurityInterceptDeploymentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityInterceptDeployment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityInterceptDeploymentObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityInterceptDeploymentObservedState is the state of the NetworkSecurityInterceptDeployment resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.InterceptDeployment
type NetworkSecurityInterceptDeploymentObservedState struct {
	// Output only. The timestamp when the resource was created.
	// See https://google.aip.dev/148#timestamps.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was most recently updated.
	// See https://google.aip.dev/148#timestamps.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The current state of the deployment.
	// See https://google.aip.dev/216.
	State *string `json:"state,omitempty"`

	// Output only. The current state of the resource does not match the user's intended state, and the system is working to reconcile them.
	Reconciling *bool `json:"reconciling,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecurityinterceptdeployment;gcpnetworksecurityinterceptdeployments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityInterceptDeployment is the Schema for the NetworkSecurityInterceptDeployment API
// +k8s:openapi-gen=true
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
type NetworkSecurityInterceptDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityInterceptDeploymentSpec   `json:"spec,omitempty"`
	Status NetworkSecurityInterceptDeploymentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityInterceptDeploymentList contains a list of NetworkSecurityInterceptDeployment
type NetworkSecurityInterceptDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityInterceptDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityInterceptDeployment{}, &NetworkSecurityInterceptDeploymentList{})
}
