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
	fuzztesting.RegisterKRMFuzzer(appEngineStandardAppVersionFuzzer())
}

func appEngineStandardAppVersionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Version{},
		AppEngineStandardAppVersionSpec_v1alpha1_FromProto, AppEngineStandardAppVersionSpec_v1alpha1_ToProto,
		AppEngineStandardAppVersionObservedState_v1alpha1_FromProto, AppEngineStandardAppVersionObservedState_v1alpha1_ToProto,
	)

	// +-----------------------------------------------------------+-----------------------------------+-------------------------------------------------------------+
	// | KRM Spec Field                                            | Proto Field                       | Fuzzer/Mapping Status                                       |
	// +-----------------------------------------------------------+-----------------------------------+-------------------------------------------------------------+
	// | appEngineApis                                             | app_engine_apis                   | SpecField (".app_engine_apis")                              |
	// | automaticScaling                                          | automatic_scaling                 | SpecField (".automatic_scaling")                            |
	// | automaticScaling.maxConcurrentRequests                     | automatic_scaling.max_concurrent_requests | SpecField (via .automatic_scaling)                   |
	// | automaticScaling.maxIdleInstances                         | automatic_scaling.max_idle_instances | SpecField (via .automatic_scaling)                      |
	// | automaticScaling.maxPendingLatency                        | automatic_scaling.max_pending_latency | SpecField (via .automatic_scaling)                     |
	// | automaticScaling.minIdleInstances                         | automatic_scaling.min_idle_instances | SpecField (via .automatic_scaling)                      |
	// | automaticScaling.minPendingLatency                        | automatic_scaling.min_pending_latency | SpecField (via .automatic_scaling)                     |
	// | automaticScaling.standardSchedulerSettings                | automatic_scaling.standard_scheduler_settings | SpecField (via .automatic_scaling)             |
	// | automaticScaling.standardSchedulerSettings.maxInstances   | automatic_scaling.standard_scheduler_settings.max_instances | SpecField (via .automatic_scaling)     |
	// | automaticScaling.standardSchedulerSettings.minInstances   | automatic_scaling.standard_scheduler_settings.min_instances | SpecField (via .automatic_scaling)     |
	// | automaticScaling.standardSchedulerSettings.targetCpuUtilization | automatic_scaling.standard_scheduler_settings.target_cpu_utilization | SpecField (via .automatic_scaling) |
	// | automaticScaling.standardSchedulerSettings.targetThroughputUtilization | automatic_scaling.standard_scheduler_settings.target_throughput_utilization | SpecField (via .automatic_scaling) |
	// | basicScaling                                              | basic_scaling                     | SpecField (".basic_scaling")                                |
	// | basicScaling.idleTimeout                                  | basic_scaling.idle_timeout         | SpecField (via .basic_scaling)                              |
	// | basicScaling.maxInstances                                 | basic_scaling.max_instances        | SpecField (via .basic_scaling)                              |
	// | deleteServiceOnDestroy                                    | N/A                               | KRM spec-only field                                         |
	// | deployment                                                | deployment                        | SpecField (".deployment")                                   |
	// | deployment.files[].name                                   | deployment.files (map key)        | SpecField (via .deployment)                                 |
	// | deployment.files[].sha1Sum                                | deployment.files[].sha1_sum       | SpecField (via .deployment)                                 |
	// | deployment.files[].sourceUrl                              | deployment.files[].source_url     | SpecField (via .deployment)                                 |
	// | deployment.zip.filesCount                                 | deployment.zip.files_count        | SpecField (via .deployment)                                 |
	// | deployment.zip.sourceUrl                                  | deployment.zip.source_url         | SpecField (via .deployment)                                 |
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
	// | libraries                                                 | libraries                         | SpecField (".libraries")                                    |
	// | libraries[].name                                          | libraries[].name                  | SpecField (via .libraries)                                  |
	// | libraries[].version                                       | libraries[].version               | SpecField (via .libraries)                                  |
	// | manualScaling                                             | manual_scaling                    | SpecField (".manual_scaling")                               |
	// | manualScaling.instances                                   | manual_scaling.instances          | SpecField (via .manual_scaling)                             |
	// | noopOnDestroy                                             | N/A                               | KRM spec-only field                                         |
	// | project                                                   | N/A                               | KRM spec-only GCP project reference                         |
	// | resourceID                                                | id                                | Unimplemented_Identity (".id")                              |
	// | runtime                                                   | runtime                           | SpecField (".runtime")                                      |
	// | runtimeApiVersion                                         | runtime_api_version               | SpecField (".runtime_api_version")                          |
	// | serviceAccount                                            | service_account                   | SpecField (".service_account")                              |
	// | serviceRef                                                | N/A                               | KRM spec-only GCP service reference                         |
	// | threadsafe                                                | threadsafe                        | SpecField (".threadsafe")                                   |
	// | vpcAccessConnector                                        | vpc_access_connector              | SpecField (".vpc_access_connector")                         |
	// | vpcAccessConnector.egressSetting                          | vpc_access_connector.egress_setting | SpecField (via .vpc_access_connector)                     |
	// | vpcAccessConnector.name                                   | vpc_access_connector.name         | SpecField (via .vpc_access_connector)                       |
	// +-----------------------------------------------------------+-----------------------------------+-------------------------------------------------------------+

	// Spec fields
	f.SpecField(".app_engine_apis")
	f.SpecField(".automatic_scaling")
	f.SpecField(".basic_scaling")
	f.SpecField(".deployment")
	f.SpecField(".entrypoint")
	f.SpecField(".env_variables")
	f.SpecField(".handlers")
	f.SpecField(".inbound_services")
	f.SpecField(".instance_class")
	f.SpecField(".libraries")
	f.SpecField(".manual_scaling")
	f.SpecField(".runtime")
	f.SpecField(".runtime_api_version")
	f.SpecField(".service_account")
	f.SpecField(".threadsafe")
	f.SpecField(".vpc_access_connector")

	// Identity fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".id")

	// Unimplemented / Not yet triaged fields in pb.Version
	f.Unimplemented_NotYetTriaged(".network")
	f.Unimplemented_NotYetTriaged(".zones")
	f.Unimplemented_NotYetTriaged(".resources")
	f.Unimplemented_NotYetTriaged(".runtime_channel")
	f.Unimplemented_NotYetTriaged(".vm")
	f.Unimplemented_NotYetTriaged(".beta_settings")
	f.Unimplemented_NotYetTriaged(".env")
	f.Unimplemented_NotYetTriaged(".serving_status")
	f.Unimplemented_NotYetTriaged(".created_by")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".disk_usage_bytes")
	f.Unimplemented_NotYetTriaged(".runtime_main_executable_path")
	f.Unimplemented_NotYetTriaged(".error_handlers")
	f.Unimplemented_NotYetTriaged(".api_config")
	f.Unimplemented_NotYetTriaged(".build_env_variables")
	f.Unimplemented_NotYetTriaged(".default_expiration")
	f.Unimplemented_NotYetTriaged(".health_check")
	f.Unimplemented_NotYetTriaged(".readiness_check")
	f.Unimplemented_NotYetTriaged(".liveness_check")
	f.Unimplemented_NotYetTriaged(".nobuild_files_regex")
	f.Unimplemented_NotYetTriaged(".version_url")
	f.Unimplemented_NotYetTriaged(".endpoints_api_service")

	// Unimplemented nested fields in deployment
	f.Unimplemented_NotYetTriaged(".deployment.cloud_build_options")
	f.Unimplemented_NotYetTriaged(".deployment.container")
	f.Unimplemented_NotYetTriaged(".deployment.files[].mime_type")

	// Unimplemented nested fields in automatic_scaling
	f.Unimplemented_NotYetTriaged(".automatic_scaling.cool_down_period")
	f.Unimplemented_NotYetTriaged(".automatic_scaling.cpu_utilization")
	f.Unimplemented_NotYetTriaged(".automatic_scaling.max_total_instances")
	f.Unimplemented_NotYetTriaged(".automatic_scaling.min_total_instances")
	f.Unimplemented_NotYetTriaged(".automatic_scaling.request_utilization")
	f.Unimplemented_NotYetTriaged(".automatic_scaling.disk_utilization")
	f.Unimplemented_NotYetTriaged(".automatic_scaling.network_utilization")

	// Unimplemented nested fields in handlers
	f.Unimplemented_NotYetTriaged(".handlers[].api_endpoint")

	return f
}
