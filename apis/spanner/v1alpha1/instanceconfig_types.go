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

var SpannerInstanceConfigGVK = GroupVersion.WithKind("SpannerInstanceConfig")

// SpannerInstanceConfigSpec defines the desired state of SpannerInstanceConfig
type SpannerInstanceConfigSpec struct {
	// Immutable. The Project that this resource belongs to.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// The location of the cluster.
	Location string `json:"location,omitempty"`

	// The SpannerInstanceConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The name of this instance configuration as it appears in UIs.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The geographic placement of nodes in this instance configuration and their
	//  replication properties.
	//
	//  To create user-managed configurations, input
	//  `replicas` must include all replicas in `replicas` of the `base_config`
	//  and include one or more replicas in the `optional_replicas` of the
	//  `base_config`.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.replicas
	Replicas []ReplicaInfo `json:"replicas,omitempty"`

	// Reference to the configuration based on which this configuration is created. Only set for user-managed
	//  configurations. `base_config` must refer to a configuration of type
	//  `GOOGLE_MANAGED` in the same project as this configuration.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.base_config
	BaseConfig *InstanceConfigRef `json:"baseConfigRef,omitempty"`

	// Cloud Labels are a flexible and lightweight mechanism for organizing cloud
	//  resources into groups that reflect a customer's organizational needs and
	//  deployment strategies. Cloud Labels can be used to filter collections of
	//  resources. They can be used to control how resource metrics are aggregated.
	//  And they can be used as arguments to policy management rules (e.g. route,
	//  firewall, load balancing, etc.).
	//
	//   * Label keys must be between 1 and 63 characters long and must conform to
	//     the following regular expression: `[a-z][a-z0-9_-]{0,62}`.
	//   * Label values must be between 0 and 63 characters long and must conform
	//     to the regular expression `[a-z0-9_-]{0,63}`.
	//   * No more than 64 labels can be associated with a given resource.
	//
	//  See https://goo.gl/xmQnxf for more information on and examples of labels.
	//
	//  If you plan to use labels in your own code, please note that additional
	//  characters may be allowed in the future. Therefore, you are advised to use
	//  an internal label representation, such as JSON, which doesn't rely upon
	//  specific characters being disallowed.  For example, representing labels
	//  as the string:  name + "_" + value  would prove problematic if we were to
	//  allow "_" in a future release.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Allowed values of the "default_leader" schema option for databases in
	//  instances that use this instance configuration.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.leader_options
	LeaderOptions []string `json:"leaderOptions,omitempty"`
}

// SpannerInstanceConfigStatus defines the config connector machine state of SpannerInstanceConfig
type SpannerInstanceConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SpannerInstanceConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SpannerInstanceConfigObservedState `json:"observedState,omitempty"`
}

// SpannerInstanceConfigSpec defines the desired state of SpannerInstanceConfig
// +kcc:proto=google.spanner.admin.instance.v1.InstanceConfig
// SpannerInstanceConfigObservedState is the state of the SpannerInstanceConfig resource as most recently observed in GCP.
type SpannerInstanceConfigObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpspannerinstanceconfig;gcpspannerinstanceconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SpannerInstanceConfig is the Schema for the SpannerInstanceConfig API
// +k8s:openapi-gen=true
type SpannerInstanceConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SpannerInstanceConfigSpec   `json:"spec,omitempty"`
	Status SpannerInstanceConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SpannerInstanceConfigList contains a list of SpannerInstanceConfig
type SpannerInstanceConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpannerInstanceConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpannerInstanceConfig{}, &SpannerInstanceConfigList{})
}
