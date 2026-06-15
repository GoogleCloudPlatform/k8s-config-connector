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
// proto.message: google.dataflow.v1beta3.Job
// api.group: dataflow.cnrm.cloud.google.com

package dataflow

import (
	pb "cloud.google.com/go/dataflow/apiv1beta3/dataflowpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataflowJobFuzzer())
}

func dataflowJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Job{},
		DataflowJobSpec_FromProto, DataflowJobSpec_ToProto,
		DataflowJobStatus_FromProto, DataflowJobStatus_ToProto,
	)

	// Mapped KRM Spec Fields
	f.SpecField(".transform_name_mapping")
	f.SpecField(".environment.temp_storage_prefix")
	f.SpecField(".environment.service_kms_key_name")
	f.SpecField(".environment.experiments")
	f.SpecField(".environment.worker_pools[].machine_type")
	f.SpecField(".environment.worker_pools[].autoscaling_settings.max_num_workers")
	f.SpecField(".environment.worker_pools[].network")
	f.SpecField(".environment.worker_pools[].subnetwork")
	f.SpecField(".environment.worker_pools[].ip_configuration")
	f.SpecField(".environment.worker_pools[].zone")
	f.SpecField(".location")
	f.SpecField(".name")

	// Mapped KRM Status Fields
	f.StatusField(".id")
	f.StatusField(".type")
	f.StatusField(".current_state")

	// Unimplemented Job fields
	f.UnimplementedFields.Insert(".project_id")
	f.UnimplementedFields.Insert(".steps")
	f.UnimplementedFields.Insert(".steps_location")
	f.UnimplementedFields.Insert(".requested_state")
	f.UnimplementedFields.Insert(".execution_info")
	f.UnimplementedFields.Insert(".create_time")
	f.UnimplementedFields.Insert(".replace_job_id")
	f.UnimplementedFields.Insert(".client_request_id")
	f.UnimplementedFields.Insert(".replaced_by_job_id")
	f.UnimplementedFields.Insert(".temp_files")
	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".pipeline_description")
	f.UnimplementedFields.Insert(".stage_states")
	f.UnimplementedFields.Insert(".job_metadata")
	f.UnimplementedFields.Insert(".start_time")
	f.UnimplementedFields.Insert(".created_from_snapshot_id")
	f.UnimplementedFields.Insert(".satisfies_pzs")
	f.UnimplementedFields.Insert(".runtime_updatable_params")
	f.UnimplementedFields.Insert(".satisfies_pzi")
	f.UnimplementedFields.Insert(".service_resources")
	f.UnimplementedFields.Insert(".current_state_time")

	// Unimplemented Job.Environment fields
	f.UnimplementedFields.Insert(".environment.cluster_manager_api_service")
	f.UnimplementedFields.Insert(".environment.service_options")
	f.UnimplementedFields.Insert(".environment.user_agent")
	f.UnimplementedFields.Insert(".environment.version")
	f.UnimplementedFields.Insert(".environment.dataset")
	f.UnimplementedFields.Insert(".environment.sdk_pipeline_options")
	f.UnimplementedFields.Insert(".environment.internal_experiments")
	f.UnimplementedFields.Insert(".environment.worker_region")
	f.UnimplementedFields.Insert(".environment.worker_zone")
	f.UnimplementedFields.Insert(".environment.service_account_email")
	f.UnimplementedFields.Insert(".environment.shuffle_mode")
	f.UnimplementedFields.Insert(".environment.streaming_mode")
	f.UnimplementedFields.Insert(".environment.flex_resource_scheduling_goal")
	f.UnimplementedFields.Insert(".environment.use_streaming_engine_resource_based_billing")
	f.UnimplementedFields.Insert(".environment.debug_options")

	// Unimplemented Job.Environment.WorkerPools[] fields
	f.UnimplementedFields.Insert(".environment.worker_pools[].kind")
	f.UnimplementedFields.Insert(".environment.worker_pools[].num_workers")
	f.UnimplementedFields.Insert(".environment.worker_pools[].packages")
	f.UnimplementedFields.Insert(".environment.worker_pools[].default_package_set")
	f.UnimplementedFields.Insert(".environment.worker_pools[].teardown_policy")
	f.UnimplementedFields.Insert(".environment.worker_pools[].disk_size_gb")
	f.UnimplementedFields.Insert(".environment.worker_pools[].disk_type")
	f.UnimplementedFields.Insert(".environment.worker_pools[].disk_source_image")
	f.UnimplementedFields.Insert(".environment.worker_pools[].taskrunner_settings")
	f.UnimplementedFields.Insert(".environment.worker_pools[].on_host_maintenance")
	f.UnimplementedFields.Insert(".environment.worker_pools[].data_disks")
	f.UnimplementedFields.Insert(".environment.worker_pools[].metadata")
	f.UnimplementedFields.Insert(".environment.worker_pools[].pool_args")
	f.UnimplementedFields.Insert(".environment.worker_pools[].worker_harness_container_image")
	f.UnimplementedFields.Insert(".environment.worker_pools[].num_threads_per_worker")
	f.UnimplementedFields.Insert(".environment.worker_pools[].sdk_harness_container_images")
	f.UnimplementedFields.Insert(".environment.worker_pools[].autoscaling_settings.algorithm")

	f.FilterSpec = func(in *pb.Job) {
		if in.GetEnvironment() != nil {
			env := in.GetEnvironment()
			pools := env.GetWorkerPools()
			if len(pools) > 0 {
				wp := pools[0]

				// Normalize/Clear autoscaling settings if max_num_workers is 0
				if wp.GetAutoscalingSettings() != nil && wp.GetAutoscalingSettings().GetMaxNumWorkers() == 0 {
					wp.AutoscalingSettings = nil
				}

				// If all mapped fields are empty, remove the worker pool entirely
				if wp.GetMachineType() == "" &&
					wp.GetAutoscalingSettings() == nil &&
					wp.GetNetwork() == "" &&
					wp.GetSubnetwork() == "" &&
					wp.GetIpConfiguration() == pb.WorkerIPAddressConfiguration_WORKER_IP_UNSPECIFIED &&
					wp.GetZone() == "" {
					env.WorkerPools = nil
				} else {
					// Otherwise, keep only the first one
					env.WorkerPools = pools[:1]
				}
			}

			// If the entire environment has no mapped fields populated, remove it entirely
			if env.GetTempStoragePrefix() == "" &&
				env.GetServiceKmsKeyName() == "" &&
				len(env.GetExperiments()) == 0 &&
				len(env.GetWorkerPools()) == 0 {
				in.Environment = nil
			}
		}
	}

	f.FilterStatus = func(in *pb.Job) {
		in.Environment = nil
		in.TransformNameMapping = nil
		in.Location = ""
		in.Name = ""
		in.CurrentStateTime = nil
	}

	return f
}
