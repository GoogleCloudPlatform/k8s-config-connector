# Copyright 2026 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import re
import os

filepath = "apis/aiplatform/v1alpha1/types.generated.go"
if not os.path.exists(filepath):
    filepath = "types.generated.go"

with open(filepath, "r") as f:
    content = f.read()

content = content.replace(
    "type ModelMonitoringObjectiveConfig_PredictionDriftDetectionConfig struct {\n\n\t// TODO: unsupported map type with key string and value message\n\n\t// TODO: unsupported map type with key string and value message",
    "type ModelMonitoringObjectiveConfig_PredictionDriftDetectionConfig struct {\n\t// Key is the feature name and value is the threshold.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.PredictionDriftDetectionConfig.drift_thresholds\n\tDriftThresholds map[string]ThresholdConfig `json:\"driftThresholds,omitempty\"`\n\n\t// Key is the feature name and value is the threshold.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.PredictionDriftDetectionConfig.attribution_score_drift_thresholds\n\tAttributionScoreDriftThresholds map[string]ThresholdConfig `json:\"attributionScoreDriftThresholds,omitempty\"`"
)

content = content.replace(
    "type ModelMonitoringObjectiveConfig_TrainingPredictionSkewDetectionConfig struct {\n\n\t// TODO: unsupported map type with key string and value message\n\n\t// TODO: unsupported map type with key string and value message",
    "type ModelMonitoringObjectiveConfig_TrainingPredictionSkewDetectionConfig struct {\n\t// Key is the feature name and value is the threshold.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.TrainingPredictionSkewDetectionConfig.skew_thresholds\n\tSkewThresholds map[string]ThresholdConfig `json:\"skewThresholds,omitempty\"`\n\n\t// Key is the feature name and value is the threshold.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.ModelMonitoringObjectiveConfig.TrainingPredictionSkewDetectionConfig.attribution_score_skew_thresholds\n\tAttributionScoreSkewThresholds map[string]ThresholdConfig `json:\"attributionScoreSkewThresholds,omitempty\"`"
)

content = content.replace(
    "type ModelDeploymentMonitoringBigQueryTable struct {\n\t// The source of log.",
    "type ModelDeploymentMonitoringBigQueryTable struct {\n\t// Output only. The schema version of the request/response logging BigQuery table. Default to v1 if unset.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringBigQueryTable.request_response_logging_schema_version\n\tRequestResponseLoggingSchemaVersion *string `json:\"requestResponseLoggingSchemaVersion,omitempty\"`\n\n\t// The source of log."
)

if 'import refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"' not in content:
    content = content.replace(
        'package v1alpha1',
        'package v1alpha1\n\nimport (\n\trefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"\n)'
    )

content = content.replace(
    'type EncryptionSpec struct {\n\t// Required. The Cloud KMS resource identifier of the customer managed\n\t//  encryption key used to protect a resource. Has the form:\n\t//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.\n\t//  The key needs to be in the same region as where the compute resource is\n\t//  created.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.EncryptionSpec.kms_key_name\n\tKMSKeyName *string `json:"kmsKeyName,omitempty"`\n}',
    'type EncryptionSpec struct {\n\t// Required. The Cloud KMS resource identifier of the customer managed\n\t//  encryption key used to protect a resource.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.EncryptionSpec.kms_key_name\n\tKmsKeyRef *refs.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`\n}'
)

content = content.replace(
    "Values []Value `json:\"values,omitempty\"`",
    "Values []apiextensionsv1.JSON `json:\"values,omitempty\"`"
)

if "// +kubebuilder:validation:Type=object\n// +kubebuilder:pruning:PreserveUnknownFields\ntype Value struct {" not in content:
    content = content.replace(
        "type Value struct {",
        "// +kubebuilder:validation:Type=object\n// +kubebuilder:pruning:PreserveUnknownFields\ntype Value struct {"
    )

if "// +kubebuilder:validation:Type=object\n// +kubebuilder:pruning:PreserveUnknownFields\ntype ListValue struct {" not in content:
    content = content.replace(
        "type ListValue struct {",
        "// +kubebuilder:validation:Type=object\n// +kubebuilder:pruning:PreserveUnknownFields\ntype ListValue struct {"
    )

content = content.replace(
    "type ModelDeploymentMonitoringBigQueryTableObservedState struct {\n\t// Output only. The schema version of the request/response logging BigQuery\n\t//  table. Default to v1 if unset.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringBigQueryTable.request_response_logging_schema_version\n\tRequestResponseLoggingSchemaVersion *string `json:\"requestResponseLoggingSchemaVersion,omitempty\"`\n}",
    "type ModelDeploymentMonitoringBigQueryTableObservedState struct {\n\t// Output only. The schema version of the request/response logging BigQuery\n\t//  table. Default to v1 if unset.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringBigQueryTable.request_response_logging_schema_version\n\tRequestResponseLoggingSchemaVersion *string `json:\"requestResponseLoggingSchemaVersion,omitempty\"`\n\n\t// Log source.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringBigQueryTable.log_source\n\tLogSource *string `json:\"logSource,omitempty\"`\n\n\t// Log type.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringBigQueryTable.log_type\n\tLogType *string `json:\"logType,omitempty\"`\n\n\t// Bigquery table path.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.ModelDeploymentMonitoringBigQueryTable.bigquery_table_path\n\tBigqueryTablePath *string `json:\"bigqueryTablePath,omitempty\"`\n}"
)

content = content.replace(
    "type DeployedModelRef struct {\n\t// Immutable. A resource name of an Endpoint.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModelRef.endpoint\n\tEndpoint *string `json:\"endpoint,omitempty\"`\n\n\t// Immutable. An ID of a DeployedModel in the above Endpoint.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModelRef.deployed_model_id\n\tDeployedModelID *string `json:\"deployedModelID,omitempty\"`\n}",
    "type DeployedModelRef struct {\n\t// Immutable. A resource name of an Endpoint.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModelRef.endpoint\n\tEndpointRef *VertexAIEndpointRef `json:\"endpointRef,omitempty\"`\n\n\t// Immutable. An ID of a DeployedModel in the above Endpoint.\n\t// +kcc:proto:field=google.cloud.aiplatform.v1.DeployedModelRef.deployed_model_id\n\tDeployedModelID *string `json:\"deployedModelID,omitempty\"`\n}"
)

with open(filepath, "w") as f:
    f.write(content)
