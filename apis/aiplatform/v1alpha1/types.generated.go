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


// +kcc:proto=google.cloud.aiplatform.v1.Feature
type Feature struct {
	// Immutable. Name of the Feature.
	//  Format:
	//  `projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entity_type}/features/{feature}`
	//  `projects/{project}/locations/{location}/featureGroups/{feature_group}/features/{feature}`
	//
	//  The last part feature is assigned by the client. The feature can be up to
	//  64 characters long and can consist only of ASCII Latin letters A-Z and a-z,
	//  underscore(_), and ASCII digits 0-9 starting with a letter. The value will
	//  be unique given an entity type.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Feature.name
	Name *string `json:"name,omitempty"`

	// Description of the Feature.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Feature.description
	Description *string `json:"description,omitempty"`

	// Immutable. Only applicable for Vertex AI Feature Store (Legacy).
	//  Type of Feature value.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Feature.value_type
	ValueType *string `json:"valueType,omitempty"`

	// Optional. The labels with user-defined metadata to organize your Features.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information on and examples of labels.
	//  No more than 64 user labels can be associated with one Feature (System
	//  labels are excluded)."
	//  System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	//  and are immutable.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Feature.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Used to perform a consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Feature.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Only applicable for Vertex AI Feature Store (Legacy).
	//  If not set, use the monitoring_config defined for the EntityType this
	//  Feature belongs to.
	//  Only Features with type
	//  ([Feature.ValueType][google.cloud.aiplatform.v1.Feature.ValueType]) BOOL,
	//  STRING, DOUBLE or INT64 can enable monitoring.
	//
	//  If set to true, all types of data monitoring are disabled despite the
	//  config on EntityType.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Feature.disable_monitoring
	DisableMonitoring *bool `json:"disableMonitoring,omitempty"`

	// Only applicable for Vertex AI Feature Store.
	//  The name of the BigQuery Table/View column hosting data for this version.
	//  If no value is provided, will use feature_id.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Feature.version_column_name
	VersionColumnName *string `json:"versionColumnName,omitempty"`

	// Entity responsible for maintaining this feature. Can be comma separated
	//  list of email addresses or URIs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Feature.point_of_contact
	PointOfContact *string `json:"pointOfContact,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Feature.MonitoringStatsAnomaly
type Feature_MonitoringStatsAnomaly struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureStatsAnomaly
type FeatureStatsAnomaly struct {
	// Feature importance score, only populated when cross-feature monitoring is
	//  enabled. For now only used to represent feature attribution score within
	//  range [0, 1] for
	//  [ModelDeploymentMonitoringObjectiveType.FEATURE_ATTRIBUTION_SKEW][google.cloud.aiplatform.v1.ModelDeploymentMonitoringObjectiveType.FEATURE_ATTRIBUTION_SKEW]
	//  and
	//  [ModelDeploymentMonitoringObjectiveType.FEATURE_ATTRIBUTION_DRIFT][google.cloud.aiplatform.v1.ModelDeploymentMonitoringObjectiveType.FEATURE_ATTRIBUTION_DRIFT].
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureStatsAnomaly.score
	Score *float64 `json:"score,omitempty"`

	// Path of the stats file for current feature values in Cloud Storage bucket.
	//  Format: gs://<bucket_name>/<object_name>/stats.
	//  Example: gs://monitoring_bucket/feature_name/stats.
	//  Stats are stored as binary format with Protobuf message
	//  [tensorflow.metadata.v0.FeatureNameStatistics](https://github.com/tensorflow/metadata/blob/master/tensorflow_metadata/proto/v0/statistics.proto).
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureStatsAnomaly.stats_uri
	StatsURI *string `json:"statsURI,omitempty"`

	// Path of the anomaly file for current feature values in Cloud Storage
	//  bucket.
	//  Format: gs://<bucket_name>/<object_name>/anomalies.
	//  Example: gs://monitoring_bucket/feature_name/anomalies.
	//  Stats are stored as binary format with Protobuf message
	//  Anoamlies are stored as binary format with Protobuf message
	//  [tensorflow.metadata.v0.AnomalyInfo]
	//  (https://github.com/tensorflow/metadata/blob/master/tensorflow_metadata/proto/v0/anomalies.proto).
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureStatsAnomaly.anomaly_uri
	AnomalyURI *string `json:"anomalyURI,omitempty"`

	// Deviation from the current stats to baseline stats.
	//    1. For categorical feature, the distribution distance is calculated by
	//       L-inifinity norm.
	//    2. For numerical feature, the distribution distance is calculated by
	//       Jensen–Shannon divergence.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureStatsAnomaly.distribution_deviation
	DistributionDeviation *float64 `json:"distributionDeviation,omitempty"`

	// This is the threshold used when detecting anomalies.
	//  The threshold can be changed by user, so this one might be different from
	//  [ThresholdConfig.value][google.cloud.aiplatform.v1.ThresholdConfig.value].
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureStatsAnomaly.anomaly_detection_threshold
	AnomalyDetectionThreshold *float64 `json:"anomalyDetectionThreshold,omitempty"`

	// The start timestamp of window where stats were generated.
	//  For objectives where time window doesn't make sense (e.g. Featurestore
	//  Snapshot Monitoring), start_time is only used to indicate the monitoring
	//  intervals, so it always equals to (end_time - monitoring_interval).
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureStatsAnomaly.start_time
	StartTime *string `json:"startTime,omitempty"`

	// The end timestamp of window where stats were generated.
	//  For objectives where time window doesn't make sense (e.g. Featurestore
	//  Snapshot Monitoring), end_time indicates the timestamp of the data used to
	//  generate stats (e.g. timestamp we take snapshots for feature values).
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureStatsAnomaly.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Feature
type FeatureObservedState struct {
	// Output only. Only applicable for Vertex AI Feature Store (Legacy).
	//  Timestamp when this EntityType was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Feature.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Only applicable for Vertex AI Feature Store (Legacy).
	//  Timestamp when this EntityType was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Feature.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Only applicable for Vertex AI Feature Store (Legacy).
	//  The list of historical stats and anomalies with specified objectives.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Feature.monitoring_stats_anomalies
	MonitoringStatsAnomalies []Feature_MonitoringStatsAnomaly `json:"monitoringStatsAnomalies,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.Feature.MonitoringStatsAnomaly
type Feature_MonitoringStatsAnomalyObservedState struct {
	// Output only. The objective for each stats.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Feature.MonitoringStatsAnomaly.objective
	Objective *string `json:"objective,omitempty"`

	// Output only. The stats and anomalies generated at specific timestamp.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Feature.MonitoringStatsAnomaly.feature_stats_anomaly
	FeatureStatsAnomaly *FeatureStatsAnomaly `json:"featureStatsAnomaly,omitempty"`
}
