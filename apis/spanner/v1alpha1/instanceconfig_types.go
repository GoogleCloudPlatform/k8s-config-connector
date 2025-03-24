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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SpannerInstanceConfigGVK = GroupVersion.WithKind("SpannerInstanceConfig")

// SpannerInstanceConfigSpec defines the desired state of SpannerInstanceConfig
// +kcc:proto=google.spanner.admin.instance.v1.InstanceConfig
type SpannerInstanceConfigSpec struct {
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

	// Base configuration name, e.g. projects/<project_name>/instanceConfigs/nam3,
	//  based on which this configuration is created. Only set for user-managed
	//  configurations. `base_config` must refer to a configuration of type
	//  `GOOGLE_MANAGED` in the same project as this configuration.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.base_config
	BaseConfig *string `json:"baseConfig,omitempty"`

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
	//     to the regular expression: `[a-z0-9_-]{0,63}`.
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

	// etag is used for optimistic concurrency control as a way
	//  to help prevent simultaneous updates of a instance configuration from
	//  overwriting each other. It is strongly suggested that systems make use of
	//  the etag in the read-modify-write cycle to perform instance configuration
	//  updates in order to avoid race conditions: An etag is returned in the
	//  response which contains instance configurations, and systems are expected
	//  to put that etag in the request to update instance configuration to ensure
	//  that their change is applied to the same version of the instance
	//  configuration. If no etag is provided in the call to update the instance
	//  configuration, then the existing instance configuration is overwritten
	//  blindly.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.etag
	Etag *string `json:"etag,omitempty"`

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

// SpannerInstanceConfigObservedState is the state of the SpannerInstanceConfig resource as most recently observed in GCP.
// +kcc:proto=google.spanner.admin.instance.v1.InstanceConfig
type SpannerInstanceConfigObservedState struct {
	// Output only. Whether this instance configuration is a Google-managed or
	//  user-managed configuration.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.config_type
	ConfigType *string `json:"configType,omitempty"`

	// Output only. The available optional replicas to choose from for
	//  user-managed configurations. Populated for Google-managed configurations.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.optional_replicas
	OptionalReplicas []ReplicaInfo `json:"optionalReplicas,omitempty"`

	// Output only. If true, the instance configuration is being created or
	//  updated. If false, there are no ongoing operations for the instance
	//  configuration.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The current instance configuration state. Applicable only for
	//  `USER_MANAGED` configurations.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.state
	State *string `json:"state,omitempty"`

	// Output only. Describes whether free instances are available to be created
	//  in this instance configuration.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.free_instance_availability
	FreeInstanceAvailability *string `json:"freeInstanceAvailability,omitempty"`

	// Output only. The `QuorumType` of the instance configuration.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.quorum_type
	QuorumType *string `json:"quorumType,omitempty"`

	// Output only. The storage limit in bytes per processing unit.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.storage_limit_per_processing_unit
	StorageLimitPerProcessingUnit *int64 `json:"storageLimitPerProcessingUnit,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
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
