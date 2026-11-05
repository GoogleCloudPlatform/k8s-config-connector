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

// +tool:fuzz-gen
// proto.message: google.appengine.v1.Version
// api.group: appengine.cnrm.cloud.google.com

package appengine

import (
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(appEngineFlexibleAppVersionFuzzer())
}

func appEngineFlexibleAppVersionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Version{},
		AppEngineFlexibleAppVersionSpec_v1alpha1_FromProto, AppEngineFlexibleAppVersionSpec_v1alpha1_ToProto,
		AppEngineFlexibleAppVersionObservedState_v1alpha1_FromProto, AppEngineFlexibleAppVersionObservedState_v1alpha1_ToProto,
	)

	// +-----------------------------------------------------------+-----------------------------------+-------------------------------------------------------------+
	// | KRM Spec Field                                            | Proto Field                       | Fuzzer/Mapping Status                                       |
	// +-----------------------------------------------------------+-----------------------------------+-------------------------------------------------------------+
	// | apiConfig                                                 | api_config                        | SpecField (".api_config")                                    |
	// | apiConfig.authFailAction                                  | api_config.auth_fail_action       | SpecField (via .api_config)                                  |
	// | apiConfig.login                                           | api_config.login                  | SpecField (via .api_config)                                  |
	// | apiConfig.script                                          | api_config.script                 | SpecField (via .api_config)                                  |
	// | apiConfig.securityLevel                                   | api_config.security_level         | SpecField (via .api_config)                                  |
	// | apiConfig.url                                             | api_config.url                    | SpecField (via .api_config)                                  |
	// | automaticScaling                                          | automatic_scaling                 | SpecField (".automatic_scaling")                            |
	// | automaticScaling.coolDownPeriod                           | automatic_scaling.cool_down_period | SpecField (via .automatic_scaling)                         |
	// | automaticScaling.cpuUtilization                           | automatic_scaling.cpu_utilization | SpecField (via .automatic_scaling)                          |
	// | automaticScaling.cpuUtilization.aggregationWindowLength   | automatic_scaling.cpu_utilization.aggregation_window_length | SpecField (via .automatic_scaling) |
	// | automaticScaling.cpuUtilization.targetUtilization         | automatic_scaling.cpu_utilization.target_utilization | SpecField (via .automatic_scaling)         |
	// | automaticScaling.diskUtilization                          | automatic_scaling.disk_utilization | SpecField (via .automatic_scaling)                         |
	// | automaticScaling.diskUtilization.targetReadBytesPerSecond | automatic_scaling.disk_utilization.target_read_bytes_per_second | SpecField (via .automatic_scaling) |
	// | automaticScaling.diskUtilization.targetReadOpsPerSecond   | automatic_scaling.disk_utilization.target_read_ops_per_second | SpecField (via .automatic_scaling)   |
	// | automaticScaling.diskUtilization.targetWriteBytesPerSecond| automatic_scaling.disk_utilization.target_write_bytes_per_second | SpecField (via .automatic_scaling)|
	// | automaticScaling.diskUtilization.targetWriteOpsPerSecond  | automatic_scaling.disk_utilization.target_write_ops_per_second | SpecField (via .automatic_scaling)  |
	// | automaticScaling.maxConcurrentRequests                     | automatic_scaling.max_concurrent_requests | SpecField (via .automatic_scaling)                   |
	// | automaticScaling.maxIdleInstances                         | automatic_scaling.max_idle_instances | SpecField (via .automatic_scaling)                      |
	// | automaticScaling.maxPendingLatency                        | automatic_scaling.max_pending_latency | SpecField (via .automatic_scaling)                     |
	// | automaticScaling.maxTotalInstances                        | automatic_scaling.max_total_instances | SpecField (via .automatic_scaling)                     |
	// | automaticScaling.minIdleInstances                         | automatic_scaling.min_idle_instances | SpecField (via .automatic_scaling)                      |
	// | automaticScaling.minPendingLatency                        | automatic_scaling.min_pending_latency | SpecField (via .automatic_scaling)                     |
	// | automaticScaling.minTotalInstances                        | automatic_scaling.min_total_instances | SpecField (via .automatic_scaling)                     |
	// | automaticScaling.networkUtilization                       | automatic_scaling.network_utilization | SpecField (via .automatic_scaling)                      |
	// | automaticScaling.networkUtilization.targetReceivedBytesPerSecond | automatic_scaling.network_utilization.target_received_bytes_per_second | SpecField (via .automatic_scaling) |
	// | automaticScaling.networkUtilization.targetReceivedPacketsPerSecond | automatic_scaling.network_utilization.target_received_packets_per_second | SpecField (via .automatic_scaling) |
	// | automaticScaling.networkUtilization.targetSentBytesPerSecond | automatic_scaling.network_utilization.target_sent_bytes_per_second | SpecField (via .automatic_scaling) |
	// | automaticScaling.networkUtilization.targetSentPacketsPerSecond | automatic_scaling.network_utilization.target_sent_packets_per_second | SpecField (via .automatic_scaling) |
	// | automaticScaling.requestUtilization                       | automatic_scaling.request_utilization | SpecField (via .automatic_scaling)                      |
	// | automaticScaling.requestUtilization.targetConcurrentRequests | automatic_scaling.request_utilization.target_concurrent_requests | SpecField (via .automatic_scaling) |
	// | automaticScaling.requestUtilization.targetRequestCountPerSecond | automatic_scaling.request_utilization.target_request_count_per_second | SpecField (via .automatic_scaling) |
	// | betaSettings                                              | beta_settings                     | SpecField (".beta_settings")                                |
	// | defaultExpiration                                         | default_expiration                | SpecField (".default_expiration")                           |
	// | deleteServiceOnDestroy                                    | N/A                               | KRM spec-only field                                         |
	// | deployment                                                | deployment                        | SpecField (".deployment")                                   |
	// | deployment.cloudBuildOptions                              | deployment.cloud_build_options    | SpecField (via .deployment)                                 |
	// | deployment.cloudBuildOptions.appYamlPath                  | deployment.cloud_build_options.app_yaml_path | SpecField (via .deployment)                        |
	// | deployment.cloudBuildOptions.cloudBuildTimeout            | deployment.cloud_build_options.cloud_build_timeout | SpecField (via .deployment)                    |
	// | deployment.container                                      | deployment.container              | SpecField (via .deployment)                                 |
	// | deployment.container.image                                | deployment.container.image        | SpecField (via .deployment)                                 |
	// | deployment.files[].name                                   | deployment.files (map key)        | SpecField (via .deployment)                                 |
	// | deployment.files[].sha1Sum                                | deployment.files[].sha1_sum       | SpecField (via .deployment)                                 |
	// | deployment.files[].sourceUrl                              | deployment.files[].source_url     | SpecField (via .deployment)                                 |
	// | deployment.zip.filesCount                                 | deployment.zip.files_count        | SpecField (via .deployment)                                 |
	// | deployment.zip.sourceUrl                                  | deployment.zip.source_url         | SpecField (via .deployment)                                 |
	// | endpointsApiService                                       | endpoints_api_service             | SpecField (".endpoints_api_service")                         |
	// | endpointsApiService.configId                              | endpoints_api_service.config_id   | SpecField (via .endpoints_api_service)                      |
	// | endpointsApiService.disableTraceSampling                  | endpoints_api_service.disable_trace_sampling | SpecField (via .endpoints_api_service)              |
	// | endpointsApiService.name                                  | endpoints_api_service.name        | SpecField (via .endpoints_api_service)                      |
	// | endpointsApiService.rolloutStrategy                       | endpoints_api_service.rollout_strategy | SpecField (via .endpoints_api_service)                  |
	// | entrypoint                                                | entrypoint                        | SpecField (".entrypoint")                                   |
	// | entrypoint.shell                                          | entrypoint.shell                  | SpecField (via .entrypoint)                                 |
	// | envVariables                                              | env_variables                     | SpecField (".env_variables")                                |
	// | handlers                                                  | handlers                          | SpecField (".handlers")                                     |
	// | handlers[].authFailAction                                 | handlers[].auth_fail_action       | SpecField (via .handlers)                                   |
	// | handlers[].login                                          | handlers[].login                  | SpecField (via .handlers)                                   |
	// | handlers[].redirectHttpResponseCode                       | handlers[].redirect_http_response_code | SpecField (via .handlers)                              |
	// | handlers[].script.scriptPath                              | handlers[].script.script_path     | SpecField (via .handlers)                                   |
	// | handlers[].securityLevel                                  | handlers[].security_level         | SpecField (via .handlers)                                   |
	// | handlers[].staticFiles.applicationReadable                | handlers[].static_files.application_readable | SpecField (via .handlers)                           |
	// | handlers[].staticFiles.expiration                         | handlers[].static_files.expiration | SpecField (via .handlers)                                  |
	// | handlers[].staticFiles.httpHeaders                        | handlers[].static_files.http_headers | SpecField (via .handlers)                                |
	// | handlers[].staticFiles.mimeType                           | handlers[].static_files.mime_type | SpecField (via .handlers)                                   |
	// | handlers[].staticFiles.path                               | handlers[].static_files.path     | SpecField (via .handlers)                                   |
	// | handlers[].staticFiles.requireMatchingFile                | handlers[].static_files.require_matching_file | SpecField (via .handlers)                           |
	// | handlers[].staticFiles.uploadPathRegex                    | handlers[].static_files.upload_path_regex | SpecField (via .handlers)                               |
	// | handlers[].urlRegex                                       | handlers[].url_regex              | SpecField (via .handlers)                                   |
	// | inboundServices                                           | inbound_services                  | SpecField (".inbound_services")                             |
	// | instanceClass                                             | instance_class                    | SpecField (".instance_class")                               |
	// | livenessCheck                                             | liveness_check                    | SpecField (".liveness_check")                               |
	// | livenessCheck.checkInterval                               | liveness_check.check_interval     | SpecField (via .liveness_check)                             |
	// | livenessCheck.failureThreshold                            | liveness_check.failure_threshold  | SpecField (via .liveness_check)                             |
	// | livenessCheck.host                                        | liveness_check.host               | SpecField (via .liveness_check)                             |
	// | livenessCheck.initialDelay                                | liveness_check.initial_delay      | SpecField (via .liveness_check)                             |
	// | livenessCheck.path                                        | liveness_check.path               | SpecField (via .liveness_check)                             |
	// | livenessCheck.successThreshold                            | liveness_check.success_threshold  | SpecField (via .liveness_check)                             |
	// | livenessCheck.timeout                                     | liveness_check.timeout             | SpecField (via .liveness_check)                             |
	// | manualScaling                                             | manual_scaling                    | SpecField (".manual_scaling")                               |
	// | manualScaling.instances                                   | manual_scaling.instances          | SpecField (via .manual_scaling)                             |
	// | network                                                   | network                           | SpecField (".network")                                      |
	// | network.forwardedPorts                                    | network.forwarded_ports           | SpecField (via .network)                                    |
	// | network.instanceTag                                       | network.instance_tag              | SpecField (via .network)                                    |
	// | network.name                                              | network.name                      | SpecField (via .network)                                    |
	// | network.sessionAffinity                                   | network.session_affinity          | SpecField (via .network)                                    |
	// | network.subnetwork                                        | network.subnetwork_name           | SpecField (via .network)                                    |
	// | nobuildFilesRegex                                         | nobuild_files_regex               | SpecField (".nobuild_files_regex")                          |
	// | noopOnDestroy                                             | N/A                               | KRM spec-only field                                         |
	// | project                                                   | N/A                               | KRM spec-only GCP project reference                         |
	// | readinessCheck                                            | readiness_check                   | SpecField (".readiness_check")                              |
	// | readinessCheck.appStartTimeout                            | readiness_check.app_start_timeout | SpecField (via .readiness_check)                            |
	// | readinessCheck.checkInterval                              | readiness_check.check_interval    | SpecField (via .readiness_check)                            |
	// | readinessCheck.failureThreshold                           | readiness_check.failure_threshold | SpecField (via .readiness_check)                            |
	// | readinessCheck.host                                       | readiness_check.host              | SpecField (via .readiness_check)                            |
	// | readinessCheck.path                                       | readiness_check.path              | SpecField (via .readiness_check)                            |
	// | readinessCheck.successThreshold                           | readiness_check.success_threshold | SpecField (via .readiness_check)                            |
	// | readinessCheck.timeout                                    | readiness_check.timeout            | SpecField (via .readiness_check)                            |
	// | resourceID                                                | id                                | Unimplemented_Identity (".id")                              |
	// | resources                                                 | resources                         | SpecField (".resources")                                    |
	// | resources.cpu                                             | resources.cpu                     | SpecField (via .resources)                                  |
	// | resources.diskGb                                          | resources.disk_gb                 | SpecField (via .resources)                                  |
	// | resources.memoryGb                                        | resources.memory_gb               | SpecField (via .resources)                                  |
	// | resources.volumes[].name                                  | resources.volumes[].name          | SpecField (via .resources)                                  |
	// | resources.volumes[].sizeGb                                | resources.volumes[].size_gb       | SpecField (via .resources)                                  |
	// | resources.volumes[].volumeType                            | resources.volumes[].volume_type   | SpecField (via .resources)                                  |
	// | runtime                                                   | runtime                           | SpecField (".runtime")                                      |
	// | runtimeApiVersion                                         | runtime_api_version               | SpecField (".runtime_api_version")                          |
	// | runtimeChannel                                            | runtime_channel                   | SpecField (".runtime_channel")                              |
	// | runtimeMainExecutablePath                                 | runtime_main_executable_path      | SpecField (".runtime_main_executable_path")                 |
	// | serviceAccount                                            | service_account                   | SpecField (".service_account")                              |
	// | serviceRef                                                | N/A                               | KRM spec-only GCP service reference                         |
	// | servingStatus                                             | serving_status                    | SpecField (".serving_status")                               |
	// | vpcAccessConnector                                        | vpc_access_connector              | SpecField (".vpc_access_connector")                         |
	// | vpcAccessConnector.name                                   | vpc_access_connector.name         | SpecField (via .vpc_access_connector)                       |
	// +-----------------------------------------------------------+-----------------------------------+-------------------------------------------------------------+

	// Spec fields
	f.SpecField(".api_config")
	f.SpecField(".automatic_scaling")
	f.SpecField(".beta_settings")
	f.SpecField(".default_expiration")
	f.SpecField(".deployment")
	f.SpecField(".endpoints_api_service")
	f.SpecField(".entrypoint")
	f.SpecField(".env_variables")
	f.SpecField(".handlers")
	f.SpecField(".inbound_services")
	f.SpecField(".instance_class")
	f.SpecField(".liveness_check")
	f.SpecField(".manual_scaling")
	f.SpecField(".network")
	f.SpecField(".nobuild_files_regex")
	f.SpecField(".readiness_check")
	f.SpecField(".resources")
	f.SpecField(".runtime")
	f.SpecField(".runtime_api_version")
	f.SpecField(".runtime_channel")
	f.SpecField(".runtime_main_executable_path")
	f.SpecField(".service_account")
	f.SpecField(".serving_status")
	f.SpecField(".vpc_access_connector")

	// Resources cpu, disk_gb, and volume size_gb are float64 in proto but integer (*int/int) in KRM.
	// FilterSpec normalizes them to integer values to prevent fractional float precision loss during roundtrip.
	f.FilterSpec = func(in *pb.Version) {
		if in.Resources != nil {
			in.Resources.Cpu = float64(int(in.Resources.Cpu))
			in.Resources.DiskGb = float64(int(in.Resources.DiskGb))
			for _, v := range in.Resources.Volumes {
				v.SizeGb = float64(int(v.SizeGb))
			}
		}
	}

	// Identity fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".id")

	// Unimplemented / Not yet triaged fields in pb.Version
	f.Unimplemented_NotYetTriaged(".basic_scaling")
	f.Unimplemented_NotYetTriaged(".zones")
	f.Unimplemented_NotYetTriaged(".threadsafe")
	f.Unimplemented_NotYetTriaged(".vm")
	f.Unimplemented_NotYetTriaged(".app_engine_apis")
	f.Unimplemented_NotYetTriaged(".env")
	f.Unimplemented_NotYetTriaged(".created_by")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".disk_usage_bytes")
	f.Unimplemented_NotYetTriaged(".error_handlers")
	f.Unimplemented_NotYetTriaged(".libraries")
	f.Unimplemented_NotYetTriaged(".build_env_variables")
	f.Unimplemented_NotYetTriaged(".health_check")
	f.Unimplemented_NotYetTriaged(".version_url")

	// Unimplemented nested fields in automatic_scaling
	f.Unimplemented_NotYetTriaged(".automatic_scaling.standard_scheduler_settings")

	// Unimplemented nested fields in deployment
	f.Unimplemented_NotYetTriaged(".deployment.files[].mime_type")

	// Unimplemented nested fields in handlers
	f.Unimplemented_NotYetTriaged(".handlers[].api_endpoint")

	// Unimplemented nested fields in resources
	f.Unimplemented_NotYetTriaged(".resources.kms_key_reference")

	// Unimplemented nested fields in vpc_access_connector
	f.Unimplemented_NotYetTriaged(".vpc_access_connector.egress_setting")

	return f
}
