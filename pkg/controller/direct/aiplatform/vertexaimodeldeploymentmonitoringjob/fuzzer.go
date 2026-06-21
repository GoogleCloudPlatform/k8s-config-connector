// Copyright 2026 Google LLC
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

package vertexaimodeldeploymentmonitoringjob

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexaimodeldeploymentmonitoringjobFuzzer())
}

func vertexaimodeldeploymentmonitoringjobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ModelDeploymentMonitoringJob{},
		VertexAIModelDeploymentMonitoringJobSpec_FromProto, VertexAIModelDeploymentMonitoringJobSpec_ToProto,
		VertexAIModelDeploymentMonitoringJobObservedState_FromProto, VertexAIModelDeploymentMonitoringJobObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".analysis_instance_schema_uri")
	f.Unimplemented_NotYetTriaged(".model_deployment_monitoring_schedule_config.monitor_window.nanos")
	f.Unimplemented_NotYetTriaged(".enable_monitoring_pipeline_logs")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".next_schedule_time")
	f.Unimplemented_NotYetTriaged(".model_deployment_monitoring_schedule_config")
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".next_schedule_time.seconds")
	f.Unimplemented_NotYetTriaged(".next_schedule_time.nanos")
	f.Unimplemented_NotYetTriaged(".encryption_spec")
	f.Unimplemented_NotYetTriaged(".encryption_spec.kms_key_name")
	f.Unimplemented_NotYetTriaged(".update_time")
	f.Unimplemented_NotYetTriaged(".endpoint")
	f.Unimplemented_NotYetTriaged(".error")
	f.Unimplemented_NotYetTriaged(".model_deployment_monitoring_schedule_config.monitor_interval.nanos")
	f.Unimplemented_NotYetTriaged(".model_deployment_monitoring_objective_configs[].objective_config.explanation_config")
	f.Unimplemented_NotYetTriaged(".model_deployment_monitoring_objective_configs")
	f.Unimplemented_NotYetTriaged(".display_name")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")
	f.Unimplemented_NotYetTriaged(".sample_predict_instance")
	f.Unimplemented_NotYetTriaged(".update_time.seconds")
	f.Unimplemented_NotYetTriaged(".stats_anomalies_base_directory")
	f.Unimplemented_NotYetTriaged(".schedule_state")
	f.Unimplemented_NotYetTriaged(".model_monitoring_alert_config")
	f.Unimplemented_NotYetTriaged(".log_ttl.seconds")
	f.Unimplemented_NotYetTriaged(".log_ttl")
	f.Unimplemented_NotYetTriaged(".logging_sampling_strategy")
	f.Unimplemented_NotYetTriaged(".latest_monitoring_pipeline_metadata")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".endpoint")
	f.SpecFields.Insert(".model_deployment_monitoring_objective_configs")
	f.SpecFields.Insert(".model_deployment_monitoring_schedule_config")
	f.SpecFields.Insert(".logging_sampling_strategy")
	f.SpecFields.Insert(".model_monitoring_alert_config")
	f.Unimplemented_NotYetTriaged(".predict_instance_schema_uri")
	f.SpecFields.Insert(".sample_predict_instance")
	f.SpecFields.Insert(".analysis_instance_schema_uri")
	f.SpecFields.Insert(".log_ttl")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".stats_anomalies_base_directory")
	f.SpecFields.Insert(".encryption_spec")
	f.SpecFields.Insert(".enable_monitoring_pipeline_logs")

	f.Unimplemented_NotYetTriaged(".state")
	f.StatusFields.Insert(".schedule_state")
	f.StatusFields.Insert(".latest_monitoring_pipeline_metadata")
	f.Unimplemented_NotYetTriaged(".bigquery_tables")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".next_schedule_time")
	f.StatusFields.Insert(".error")

	return f
}
