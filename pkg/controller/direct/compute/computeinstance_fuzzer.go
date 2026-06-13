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
// proto.message: google.cloud.compute.v1.Instance
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeInstanceFuzzer())
}

func computeInstanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Instance{},
		ComputeInstanceSpec_v1beta1_FromProto, ComputeInstanceSpec_v1beta1_ToProto,
		ComputeInstanceStatus_v1beta1_FromProto, ComputeInstanceStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".advanced_machine_features")
	f.SpecField(".confidential_instance_config")
	f.SpecField(".deletion_protection")
	f.SpecField(".description")
	f.SpecField(".hostname")
	f.SpecField(".machine_type")
	f.SpecField(".metadata")
	f.SpecField(".network_performance_config")
	f.SpecField(".params")
	f.SpecField(".reservation_affinity")
	f.SpecField(".resource_policies")
	f.SpecField(".scheduling")
	f.SpecField(".shielded_instance_config")
	f.SpecField(".tags")
	f.SpecField(".zone")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	return f
}
