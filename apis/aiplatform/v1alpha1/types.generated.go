// Copyright 2025 Google LLC
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

package v1alpha1


// +kcc:proto=google.cloud.aiplatform.v1.EntityType
type EntityType struct {
	// Immutable. Name of the EntityType.
	//  Format:
	//  `projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entity_type}`
	//
	//  The last part entity_type is assigned by the client. The entity_type can be
	//  up to 64 characters long and can consist only of ASCII Latin letters A-Z
	//  and a-z and underscore(_), and ASCII digits 0-9 starting with a letter. The
	//  value will be unique given a featurestore.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EntityType.name
	Name *string `json:"name,omitempty"`

	// Optional. Description of the EntityType.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EntityType.description
	Description *string `json:"description,omitempty"`

	// Optional. The labels with user-defined metadata to organize your
	//  EntityTypes.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information on and examples of labels.
	//  No more than 64 user labels can be associated with one EntityType (System
	//  labels are excluded)."
	//  System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	//  and are immutable.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EntityType.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Used to perform a consistent read-modify-write updates. If not
	//  set, a blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EntityType.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The default monitoring configuration for all Features with value
	//  type
	//  ([Feature.ValueType][google.cloud.aiplatform.v1.Feature.ValueType]) BOOL,
	//  STRING, DOUBLE or INT64 under this EntityType.
	//
	//  If this is populated with
	//  [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot
	//  analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is
	//  disabled.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EntityType.monitoring_config
	MonitoringConfig *FeaturestoreMonitoringConfig `json:"monitoringConfig,omitempty"`

	// Optional. Config for data retention policy in offline storage.
	//  TTL in days for feature values that will be stored in offline storage.
	//  The Feature Store offline storage periodically removes obsolete feature
	//  values older than `offline_storage_ttl_days` since the feature generation
	//  time. If unset (or explicitly set to 0), default to 4000 days TTL.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EntityType.offline_storage_ttl_days
	OfflineStorageTtlDays *int32 `json:"offlineStorageTtlDays,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeaturestoreMonitoringConfig
type FeaturestoreMonitoringConfig struct {
	// The config for Snapshot Analysis Based Feature Monitoring.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeaturestoreMonitoringConfig.snapshot_analysis
	SnapshotAnalysis *FeaturestoreMonitoringConfig_SnapshotAnalysis `json:"snapshotAnalysis,omitempty"`

	// The config for ImportFeatures Analysis Based Feature Monitoring.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeaturestoreMonitoringConfig.import_features_analysis
	ImportFeaturesAnalysis *FeaturestoreMonitoringConfig_ImportFeaturesAnalysis `json:"importFeaturesAnalysis,omitempty"`

	// Threshold for numerical features of anomaly detection.
	//  This is shared by all objectives of Featurestore Monitoring for numerical
	//  features (i.e. Features with type
	//  ([Feature.ValueType][google.cloud.aiplatform.v1.Feature.ValueType]) DOUBLE
	//  or INT64).
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeaturestoreMonitoringConfig.numerical_threshold_config
	NumericalThresholdConfig *FeaturestoreMonitoringConfig_ThresholdConfig `json:"numericalThresholdConfig,omitempty"`

	// Threshold for categorical features of anomaly detection.
	//  This is shared by all types of Featurestore Monitoring for categorical
	//  features (i.e. Features with type
	//  ([Feature.ValueType][google.cloud.aiplatform.v1.Feature.ValueType]) BOOL or
	//  STRING).
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeaturestoreMonitoringConfig.categorical_threshold_config
	CategoricalThresholdConfig *FeaturestoreMonitoringConfig_ThresholdConfig `json:"categoricalThresholdConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeaturestoreMonitoringConfig.ImportFeaturesAnalysis
type FeaturestoreMonitoringConfig_ImportFeaturesAnalysis struct {
	// Whether to enable / disable / inherite default hebavior for import
	//  features analysis.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeaturestoreMonitoringConfig.ImportFeaturesAnalysis.state
	State *string `json:"state,omitempty"`

	// The baseline used to do anomaly detection for the statistics generated by
	//  import features analysis.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeaturestoreMonitoringConfig.ImportFeaturesAnalysis.anomaly_detection_baseline
	AnomalyDetectionBaseline *string `json:"anomalyDetectionBaseline,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeaturestoreMonitoringConfig.SnapshotAnalysis
type FeaturestoreMonitoringConfig_SnapshotAnalysis struct {
	// The monitoring schedule for snapshot analysis.
	//  For EntityType-level config:
	//    unset / disabled = true indicates disabled by
	//    default for Features under it; otherwise by default enable snapshot
	//    analysis monitoring with monitoring_interval for Features under it.
	//  Feature-level config:
	//    disabled = true indicates disabled regardless of the EntityType-level
	//    config; unset monitoring_interval indicates going with EntityType-level
	//    config; otherwise run snapshot analysis monitoring with
	//    monitoring_interval regardless of the EntityType-level config.
	//  Explicitly Disable the snapshot analysis based monitoring.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeaturestoreMonitoringConfig.SnapshotAnalysis.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Configuration of the snapshot analysis based monitoring pipeline
	//  running interval. The value indicates number of days.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days
	MonitoringIntervalDays *int32 `json:"monitoringIntervalDays,omitempty"`

	// Customized export features time window for snapshot analysis. Unit is one
	//  day. Default value is 3 weeks. Minimum value is 1 day. Maximum value is
	//  4000 days.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeaturestoreMonitoringConfig.SnapshotAnalysis.staleness_days
	StalenessDays *int32 `json:"stalenessDays,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeaturestoreMonitoringConfig.ThresholdConfig
type FeaturestoreMonitoringConfig_ThresholdConfig struct {
	// Specify a threshold value that can trigger the alert.
	//  1. For categorical feature, the distribution distance is calculated by
	//  L-inifinity norm.
	//  2. For numerical feature, the distribution distance is calculated by
	//  Jensen–Shannon divergence. Each feature must have a non-zero threshold
	//  if they need to be monitored. Otherwise no alert will be triggered for
	//  that feature.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeaturestoreMonitoringConfig.ThresholdConfig.value
	Value *float64 `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.EntityType
type EntityTypeObservedState struct {
	// Output only. Timestamp when this EntityType was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EntityType.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this EntityType was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EntityType.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EntityType.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.EntityType.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}
