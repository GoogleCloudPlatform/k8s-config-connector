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
// proto.message: google.cloud.compute.v1.Autoscaler
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeAutoscalerFuzzer())
}

func computeAutoscalerFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Autoscaler{},
		ComputeAutoscalerSpec_v1alpha1_FromProto, ComputeAutoscalerSpec_v1alpha1_ToProto,
		ComputeAutoscalerStatus_v1alpha1_FromProto, ComputeAutoscalerStatus_v1alpha1_ToProto,
	)

	f.SpecField(".autoscaling_policy")
	f.SpecField(".description")
	f.SpecField(".target")
	f.SpecField(".zone")

	f.StatusField(".creation_timestamp")
	f.StatusField(".self_link")

	f.Unimplemented_Identity(".name")

	f.Unimplemented_Internal(".id")
	f.Unimplemented_Internal(".kind")
	f.Unimplemented_Internal(".recommended_size")
	f.Unimplemented_Internal(".region")
	f.Unimplemented_Internal(".status")
	f.Unimplemented_Internal(".status_details")
	f.Unimplemented_Internal(".scaling_schedule_status")
	f.Unimplemented_Internal(".autoscaling_policy.scale_in_control.max_scaled_in_replicas.calculated")

	return f
}
