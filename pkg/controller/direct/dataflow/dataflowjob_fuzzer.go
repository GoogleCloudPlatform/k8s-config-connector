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
	f.Unimplemented_Identity(".project_id")
	f.Unimplemented_NotYetTriaged(".steps")
	f.Unimplemented_NotYetTriaged(".steps_location")
	f.Unimplemented_NotYetTriaged(".requested_state")
	f.Unimplemented_NotYetTriaged(".execution_info")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".replace_job_id")
	f.Unimplemented_NotYetTriaged(".client_request_id")
	f.Unimplemented_NotYetTriaged(".replaced_by_job_id")
	f.Unimplemented_NotYetTriaged(".temp_files")
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_NotYetTriaged(".pipeline_description")
	f.Unimplemented_NotYetTriaged(".stage_states")
	f.Unimplemented_NotYetTriaged(".job_metadata")
	f.Unimplemented_NotYetTriaged(".start_time")
	f.Unimplemented_NotYetTriaged(".created_from_snapshot_id")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".runtime_updatable_params")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")
	f.Unimplemented_NotYetTriaged(".service_resources")
	f.Unimplemented_NotYetTriaged(".current_state_time")
	f.Unimplemented_NotYetTriaged(".pausable")

	// Unimplemented Job.Environment fields
	f.Unimplemented_NotYetTriaged(".environment.cluster_manager_api_service")
	f.Unimplemented_NotYetTriaged(".environment.service_options")
	f.Unimplemented_NotYetTriaged(".environment.user_agent")
	f.Unimplemented_NotYetTriaged(".environment.version")
	f.Unimplemented_NotYetTriaged(".environment.dataset")
	f.Unimplemented_NotYetTriaged(".environment.sdk_pipeline_options")
	f.Unimplemented_NotYetTriaged(".environment.internal_experiments")
	f.Unimplemented_NotYetTriaged(".environment.worker_region")
	f.Unimplemented_NotYetTriaged(".environment.worker_zone")
	f.Unimplemented_NotYetTriaged(".environment.service_account_email")
	f.Unimplemented_NotYetTriaged(".environment.shuffle_mode")
	f.Unimplemented_NotYetTriaged(".environment.streaming_mode")
	f.Unimplemented_NotYetTriaged(".environment.flex_resource_scheduling_goal")
	f.Unimplemented_NotYetTriaged(".environment.use_streaming_engine_resource_based_billing")
	f.Unimplemented_NotYetTriaged(".environment.debug_options")
	f.Unimplemented_NotYetTriaged(".environment.use_public_ips")

	// Unimplemented Job.Environment.WorkerPools[] fields
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].kind")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].num_workers")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].packages")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].default_package_set")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].teardown_policy")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].disk_size_gb")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].disk_type")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].disk_source_image")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].taskrunner_settings")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].on_host_maintenance")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].data_disks")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].metadata")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].pool_args")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].worker_harness_container_image")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].num_threads_per_worker")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].sdk_harness_container_images")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].autoscaling_settings.algorithm")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].disk_provisioned_iops")
	f.Unimplemented_NotYetTriaged(".environment.worker_pools[].disk_provisioned_throughput_mibps")

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
