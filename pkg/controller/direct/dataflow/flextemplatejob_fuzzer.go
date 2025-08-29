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

// +tool:fuzz-gen
// proto.message: google.dataflow.v1beta3.FlexTemplateRuntimeEnvironment
// api.group: dataflow.cnrm.cloud.google.com

package dataflow

import (
	pb "cloud.google.com/go/dataflow/apiv1beta3/dataflowpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(dataflowFlexTemplateJobFuzzer())
}

func dataflowFlexTemplateJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.FlexTemplateRuntimeEnvironment{},
		DataflowFlexTemplateJobSpec_FromProto, DataflowFlexTemplateJobSpec_ToProto,
	)

	f.SpecFields.Insert(".num_workers")
	f.SpecFields.Insert(".max_workers")
	f.SpecFields.Insert(".temp_location")
	f.SpecFields.Insert(".machine_type")
	f.SpecFields.Insert(".additional_experiments")
	f.SpecFields.Insert(".network")
	f.SpecFields.Insert(".subnetwork")
	f.SpecFields.Insert(".kms_key_name")
	f.SpecFields.Insert(".service_account_email")
	f.SpecFields.Insert(".ip_configuration")
	f.SpecFields.Insert(".enable_streaming_engine")
	f.SpecFields.Insert(".staging_location")
	f.SpecFields.Insert(".sdk_container_image")
	f.SpecFields.Insert(".autoscaling_algorithm")
	f.SpecFields.Insert(".launcher_machine_type")

	f.UnimplementedFields.Insert(".zone")
	f.UnimplementedFields.Insert(".additional_user_labels")
	f.UnimplementedFields.Insert(".worker_region")
	f.UnimplementedFields.Insert(".worker_zone")
	f.UnimplementedFields.Insert(".flexrs_goal")
	f.UnimplementedFields.Insert(".disk_size_gb")
	f.UnimplementedFields.Insert(".dump_heap_on_oom")
	f.UnimplementedFields.Insert(".save_heap_dumps_to_gcs_path")
	f.UnimplementedFields.Insert(".streaming_mode")

	return f
}
