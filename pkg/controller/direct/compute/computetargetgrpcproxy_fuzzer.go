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
// proto.message: google.cloud.compute.v1.TargetGrpcProxy
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeTargetGRPCProxyFuzzer())
}

func computeTargetGRPCProxyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TargetGrpcProxy{},
		ComputeTargetGRPCProxySpec_v1beta1_FromProto, ComputeTargetGRPCProxySpec_v1beta1_ToProto,
		ComputeTargetGRPCProxyStatus_v1beta1_FromProto, ComputeTargetGRPCProxyStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".url_map")
	f.SpecField(".validate_for_proxyless")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".fingerprint")
	f.StatusField(".self_link")
	f.StatusField(".self_link_with_id")

	// Identity field
	f.Unimplemented_Identity(".name")

	// Internal fields
	f.Unimplemented_Internal(".id")
	f.Unimplemented_Internal(".kind")

	return f
}
