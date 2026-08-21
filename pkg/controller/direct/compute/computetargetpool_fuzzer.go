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
// proto.message: google.cloud.compute.v1.TargetPool
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(computeTargetPoolFuzzer())
}

func computeTargetPoolFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer[*pb.TargetPool, krm.ComputeTargetPoolSpec, krm.ComputeTargetPoolStatus](&pb.TargetPool{},
		ComputeTargetPoolSpec_v1beta1_FromProto, ComputeTargetPoolSpec_v1beta1_ToProto,
		nil, nil,
	)

	// Spec fields
	f.SpecField(".backup_pool")
	f.SpecField(".description")
	f.SpecField(".failover_ratio")
	f.SpecField(".health_checks")
	f.SpecField(".instances")
	f.SpecField(".region")
	f.SpecField(".security_policy")
	f.SpecField(".session_affinity")

	// Status fields
	f.StatusField(".self_link")

	// Unimplemented fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Internal(".kind")
	f.Unimplemented_Internal(".creation_timestamp")
	f.Unimplemented_Internal(".id")

	return f
}
