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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIFeaturestoreEntityTypeGVK = GroupVersion.WithKind("VertexAIFeaturestoreEntityType")

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.EntityType
type VertexAIFeaturestoreEntityTypeSpec struct {
	// Optional. Description of the EntityType.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EntityType.description
	Description *string `json:"description,omitempty"`

	// Immutable. The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.
	// +required
	Featurestore string `json:"featurestore"`

	// The default monitoring configuration for all Features under this EntityType.
	//
	// If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EntityType.monitoring_config
	MonitoringConfig *FeaturestoreMonitoringConfig `json:"monitoringConfig,omitempty"`

	// Config for data retention policy in offline storage.
	// TTL in days for feature values that will be stored in offline storage.
	// The Feature Store offline storage periodically removes obsolete feature values older than offlineStorageTtlDays since the feature generation time. If unset (or explicitly set to 0), default to 4000 days TTL.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EntityType.offline_storage_ttl_days
	OfflineStorageTTLDays *int `json:"offlineStorageTtlDays,omitempty"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.SnapshotAnalysis
type FeaturestoreMonitoringConfig_SnapshotAnalysis struct {
	// The monitoring schedule for snapshot analysis.
	// For EntityType-level config: unset / disabled = true indicates disabled by
	// default for Features under it; otherwise by default enable snapshot
	// analysis monitoring with monitoring_interval for Features under it.
	// Feature-level config: disabled = true indicates disabled regardless of the
	// EntityType-level config; unset monitoring_interval indicates going with
	// EntityType-level config; otherwise run snapshot analysis monitoring with
	// monitoring_interval regardless of the EntityType-level config.
	// Explicitly Disable the snapshot analysis based monitoring.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.SnapshotAnalysis.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Configuration of the snapshot analysis based monitoring pipeline running
	// interval. The value is rolled up to full day.
	// If both
	// [monitoring_interval_days][google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days]
	// and the deprecated `monitoring_interval` field are set when creating/updating
	// EntityTypes/Features,
	// [monitoring_interval_days][google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days]
	// will be used.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval
	MonitoringInterval *string `json:"monitoringInterval,omitempty"`

	// Configuration of the snapshot analysis based monitoring pipeline running
	// interval. The value indicates number of days.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days
	MonitoringIntervalDays *int `json:"monitoringIntervalDays,omitempty"`

	// Customized export features time window for snapshot analysis. Unit is one
	// day. Default value is 3 weeks. Minimum value is 1 day. Maximum value is
	// 4000 days.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.SnapshotAnalysis.staleness_days
	StalenessDays *int `json:"stalenessDays,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.ThresholdConfig
type FeaturestoreMonitoringConfig_ThresholdConfig struct {
	// Specify a threshold value that can trigger the alert.
	//  1. For categorical feature, the distribution distance is calculated by
	//  L-inifinity norm.
	//  2. For numerical feature, the distribution distance is calculated by
	//  Jensen–Shannon divergence. Each feature must have a non-zero threshold
	//  if they need to be monitored. Otherwise no alert will be triggered for
	//  that feature.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.ThresholdConfig.value
	// +required
	Value float64 `json:"value"`
}

// +kcc:status:proto=google.cloud.aiplatform.v1beta1.EntityType
type VertexAIFeaturestoreEntityTypeStatus struct {
	// Conditions represent the latest available observation of the resource's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EntityType.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Used to perform consistent read-modify-write updates.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EntityType.etag
	Etag *string `json:"etag,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// The region of the EntityType.
	Region *string `json:"region,omitempty"`

	// The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EntityType.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaifeaturestoreentitytype;gcpvertexaifeaturestoreentitytypes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIFeaturestoreEntityType is the Schema for the VertexAIFeaturestoreEntityType API
// +k8s:openapi-gen=true
type VertexAIFeaturestoreEntityType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIFeaturestoreEntityTypeSpec   `json:"spec,omitempty"`
	Status VertexAIFeaturestoreEntityTypeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIFeaturestoreEntityTypeList contains a list of VertexAIFeaturestoreEntityType
type VertexAIFeaturestoreEntityTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIFeaturestoreEntityType `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIFeaturestoreEntityType{}, &VertexAIFeaturestoreEntityTypeList{})
}
