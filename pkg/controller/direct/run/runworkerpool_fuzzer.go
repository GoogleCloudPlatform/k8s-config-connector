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
// proto.message: google.cloud.run.v2.WorkerPool
// api.group: run.cnrm.cloud.google.com

package run

import (
	pb "cloud.google.com/go/run/apiv2/runpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(runWorkerPoolFuzzer())
}

func runWorkerPoolFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.WorkerPool{},
		RunWorkerPoolSpec_v1alpha1_FromProto, RunWorkerPoolSpec_v1alpha1_ToProto,
		RunWorkerPoolObservedState_v1alpha1_FromProto, RunWorkerPoolObservedState_v1alpha1_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".client")
	f.SpecField(".client_version")
	f.SpecField(".launch_stage")
	f.SpecField(".binary_authorization")
	f.SpecField(".template")
	f.SpecField(".instance_splits")
	f.SpecField(".scaling")
	f.SpecField(".custom_audiences")

	f.StatusField(".uid")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".delete_time")
	f.StatusField(".expire_time")
	f.StatusField(".creator")
	f.StatusField(".last_modifier")
	f.StatusField(".terminal_condition")
	f.StatusField(".latest_ready_revision")
	f.StatusField(".latest_created_revision")
	f.StatusField(".instance_split_statuses")
	f.StatusField(".satisfies_pzs")
	f.StatusField(".reconciling")
	f.StatusField(".etag")

	f.IdentityField(".name")

	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_LabelsAnnotations(".annotations")

	f.Unimplemented_NotYetTriaged(".generation")
	f.Unimplemented_NotYetTriaged(".observed_generation")
	f.Unimplemented_NotYetTriaged(".conditions")
	f.Unimplemented_NotYetTriaged(".threat_detection_enabled")

	f.Unimplemented_NotYetTriaged(".binary_authorization.policy")
	f.Unimplemented_NotYetTriaged(".binary_authorization.use_default")
	f.Unimplemented_NotYetTriaged(".template.labels")
	f.Unimplemented_NotYetTriaged(".template.containers[].base_image_uri")
	f.Unimplemented_NotYetTriaged(".template.containers[].build_info")
	f.Unimplemented_NotYetTriaged(".template.containers[].depends_on")
	f.Unimplemented_NotYetTriaged(".template.containers[].resources.startup_cpu_boost")
	f.Unimplemented_NotYetTriaged(".template.containers[].resources.cpu_idle")
	f.Unimplemented_NotYetTriaged(".template.containers[].startup_probe.grpc")
	f.Unimplemented_NotYetTriaged(".template.containers[].startup_probe.http_get")
	f.Unimplemented_NotYetTriaged(".template.containers[].liveness_probe.grpc")
	f.Unimplemented_NotYetTriaged(".template.containers[].liveness_probe.http_get")
	f.Unimplemented_NotYetTriaged(".template.containers[].source_code")
	f.Unimplemented_NotYetTriaged(".template.containers[].volume_mounts[].sub_path")
	f.Unimplemented_NotYetTriaged(".template.gpu_zonal_redundancy_disabled")
	f.Unimplemented_NotYetTriaged(".template.node_selector")
	f.Unimplemented_NotYetTriaged(".terminal_condition.execution_reason")
	f.Unimplemented_NotYetTriaged(".terminal_condition.reason")
	f.Unimplemented_NotYetTriaged(".terminal_condition.revision_reason")
	f.Unimplemented_NotYetTriaged(".template.containers[].readiness_probe")

	return f
}
