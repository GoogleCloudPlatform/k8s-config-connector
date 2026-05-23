// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.aiplatform.v1beta1.DeploymentResourcePool
// api.group: vertexai.cnrm.cloud.google.com

package deploymentresourcepool

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(deploymentResourcePoolFuzzer())
}

func deploymentResourcePoolFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DeploymentResourcePool{},
		VertexAIDeploymentResourcePoolSpec_FromProto, VertexAIDeploymentResourcePoolSpec_ToProto,
		VertexAIDeploymentResourcePoolObservedState_FromProto, VertexAIDeploymentResourcePoolObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".dedicated_resources")
	f.SpecFields.Insert(".encryption_spec")
	f.SpecFields.Insert(".service_account")
	f.SpecFields.Insert(".disable_container_logging")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".satisfies_pzs")
	f.StatusFields.Insert(".satisfies_pzi")

	// Untriaged / Unimplemented fields not in KRM
	f.Unimplemented_NotYetTriaged(".dedicated_resources.initial_replica_count")
	f.Unimplemented_NotYetTriaged(".dedicated_resources.machine_spec.min_gpu_driver_version")
	f.Unimplemented_NotYetTriaged(".dedicated_resources.machine_spec.reservation_affinity")
	f.Unimplemented_NotYetTriaged(".dedicated_resources.machine_spec.reservation_affinity.reservation_affinity_type")
	f.Unimplemented_NotYetTriaged(".dedicated_resources.machine_spec.reservation_affinity.key")
	f.Unimplemented_NotYetTriaged(".dedicated_resources.machine_spec.reservation_affinity.values")
	f.Unimplemented_NotYetTriaged(".dedicated_resources.spot")
	f.Unimplemented_NotYetTriaged(".dedicated_resources.flex_start")
	f.Unimplemented_NotYetTriaged(".dedicated_resources.flex_start.max_runtime_duration")
	f.Unimplemented_NotYetTriaged(".dedicated_resources.scale_to_zero_spec")
	f.Unimplemented_NotYetTriaged(".dedicated_resources.scale_to_zero_spec.min_scaleup_period")
	f.Unimplemented_NotYetTriaged(".dedicated_resources.scale_to_zero_spec.idle_scaledown_period")
	f.Unimplemented_NotYetTriaged(".dedicated_resources.autoscaling_metric_specs[].monitored_resource_labels")

	return f
}
