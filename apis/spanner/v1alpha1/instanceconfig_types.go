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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Parent struct {
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SpannerInstanceConfig is the Schema for the spanner API
// +k8s:openapi-gen=true
type SpannerInstanceConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Parent Parent `json:",inline"`

	Spec   SpannerInstanceConfigSpec   `json:"spec,omitempty"`
	Status SpannerInstanceConfigStatus `json:"status,omitempty"`
}

// +k8s:openapi-gen=true
type SpannerInstanceConfigSpec struct {
	// Immutable. The geographic placement of nodes in this instance configuration and their
	// replication properties.
	// +optional
	Replicas []SpannerInstanceConfigReplicas `json:"replicas,omitempty"`
	// Base configuration name, e.g. projects/<project_name>/instanceConfigs/nam3,
	// based on which this configuration is created. Only set for user-managed
	// configurations. `base_config` must refer to a configuration of type
	// `GOOGLE_MANAGED` in the same project as this configuration.
	// +optional
	BaseConfig *string `json:"baseConfig,omitempty"`
	// The name of this instance configuration as it appears in UIs.
	DisplayName string `json:"displayName"`
	// etag is used for optimistic concurrency control as a way
	// to help prevent simultaneous updates of a instance configuration from
	// overwriting each other. It is strongly suggested that systems make use of
	// the etag in the read-modify-write cycle to perform instance configuration
	// updates in order to avoid race conditions: An etag is returned in the
	// response which contains instance configurations, and systems are expected
	// to put that etag in the request to update instance configuration to ensure
	// that their change is applied to the same version of the instance
	// configuration. If no etag is provided in the call to update the instance
	// configuration, then the existing instance configuration is overwritten
	// blindly.
	// +optional
	Etag *string `json:"etag,omitempty"`
	// Allowed values of the "default_leader" schema option for databases in
	// instances that use this instance configuration.
	// +optional
	LeaderOptions []string `json:"leaderOptions,omitempty"`
	// Some guided coding here please
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

// SpannerInstanceConfigReplicas is a replica in a Spanner Instance Config.
type SpannerInstanceConfigReplicas struct {
	// If true, this location is designated as the default leader location where
	// leader replicas are placed. See the [region types
	// documentation](https://cloud.google.com/spanner/docs/instances#region_types)
	// for more details.
	// +optional
	DefaultLeaderLocation *bool `json:"defaultLeaderLocation,omitempty"`
	// The location of the serving resources, e.g. "us-central1".
	// +optional
	Location *string `json:"location,omitempty"`
	// The type of replica.
	// +optional
	Type *string `json:"type,omitempty"`
}

// +k8s:openapi-gen=true
type SpannerInstanceConfigStatus struct {
	// Conditions represent the latest available observation of the
	// SpannerInstanceConfig's current state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// Output only. The available optional replicas to choose from for
	// user-managed configurations. Populated for Google-managed configurations.
	// +optional
	OptionalReplicas []SpannerInstanceConfigOptionalReplicas `json:"optionalReplicas,omitempty"`
	// If true, the instance configuration is being created or updated.
	// If false, there are no ongoing operations for the instance configuration.
	// +optional
	Reconciling *bool `json:"reconciling,omitempty"`
	// Output only. The current instance configuration state.
	// +optional
	State *string `json:"state,omitempty"`
}

// SpannerInstanceConfigOptionalReplicas is an optional replica in a Spanner Instance Config.
type SpannerInstanceConfigOptionalReplicas struct {
	// If true, this location is designated as the default leader location where
	// leader replicas are placed. See the [region types
	// documentation](https://cloud.google.com/spanner/docs/instances#region_types)
	// for more details.
	// +optional
	DefaultLeaderLocation *bool `json:"defaultLeaderLocation,omitempty"`
	// The location of the serving resources, e.g. "us-central1".
	// +optional
	Location *string `json:"location,omitempty"`
	// The type of replica.
	// +optional
	Type *string `json:"type,omitempty"`
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
