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
		RunWorkerPoolSpec_FromProto, RunWorkerPoolSpec_ToProto,
		RunWorkerPoolObservedState_FromProto, RunWorkerPoolObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".annotations")
	f.SpecFields.Insert(".client")
	f.SpecFields.Insert(".client_version")
	f.SpecFields.Insert(".launch_stage")
	f.SpecFields.Insert(".binary_authorization")
	f.SpecFields.Insert(".template")
	f.SpecFields.Insert(".instance_splits")
	f.SpecFields.Insert(".scaling")
	f.SpecFields.Insert(".custom_audiences")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".generation")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".delete_time")
	f.StatusFields.Insert(".expire_time")
	f.StatusFields.Insert(".creator")
	f.StatusFields.Insert(".last_modifier")
	f.StatusFields.Insert(".template")
	f.StatusFields.Insert(".observed_generation")
	f.StatusFields.Insert(".terminal_condition")
	f.StatusFields.Insert(".conditions")
	f.StatusFields.Insert(".latest_ready_revision")
	f.StatusFields.Insert(".latest_created_revision")
	f.StatusFields.Insert(".instance_split_statuses")
	f.StatusFields.Insert(".threat_detection_enabled")
	f.StatusFields.Insert(".satisfies_pzs")
	f.StatusFields.Insert(".reconciling")

		f.IdentityField(".name")
	
	    // Fields that are not yet implemented or have known issues
	    f.Unimplemented_NotYetTriaged(".template.containers[].readiness_probe")
	    f.Unimplemented_NotYetTriaged(".template.containers[].liveness_probe.grpc")
	    f.Unimplemented_NotYetTriaged(".template.containers[].startup_probe.grpc")
	
		return f
	}
