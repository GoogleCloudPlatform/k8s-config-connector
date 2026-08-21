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
// proto.message: google.cloud.compute.v1.TargetInstance
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeTargetInstanceFuzzer())
}

func computeTargetInstanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TargetInstance{},
		ComputeTargetInstanceSpec_v1beta1_FromProto, ComputeTargetInstanceSpec_v1beta1_ToProto,
		ComputeTargetInstanceStatus_v1beta1_FromProto, ComputeTargetInstanceStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".instance")
	f.SpecField(".nat_policy")
	f.SpecField(".network")
	f.SpecField(".security_policy")
	f.SpecField(".zone")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".self_link")

	// Unimplemented fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Internal(".id")
	f.Unimplemented_Internal(".kind")

	return f
}
