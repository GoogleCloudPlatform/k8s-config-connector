// Copyright 2024 Google LLC
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
// proto.message: google.cloud.compute.v1.NetworkEdgeSecurityService
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeNetworkEdgeSecurityServiceFuzzer())
}

func computeNetworkEdgeSecurityServiceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NetworkEdgeSecurityService{},
		ComputeNetworkEdgeSecurityServiceSpec_FromProto, ComputeNetworkEdgeSecurityServiceSpec_ToProto,
		ComputeNetworkEdgeSecurityServiceObservedState_FromProto, ComputeNetworkEdgeSecurityServiceObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".fingerprint")
	f.SpecFields.Insert(".security_policy")

	f.StatusFields.Insert(".creation_timestamp")
	f.StatusFields.Insert(".id")
	f.StatusFields.Insert(".kind")
	f.StatusFields.Insert(".region")
	f.StatusFields.Insert(".self_link")
	f.StatusFields.Insert(".self_link_with_id")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
