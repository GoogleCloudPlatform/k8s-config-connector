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
// proto.message: google.cloud.compute.v1.NetworkAttachment
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeNetworkAttachmentFuzzer())
}

func computeNetworkAttachmentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NetworkAttachment{},
		ComputeNetworkAttachmentSpec_v1alpha1_FromProto, ComputeNetworkAttachmentSpec_v1alpha1_ToProto,
		ComputeNetworkAttachmentObservedState_v1alpha1_FromProto, ComputeNetworkAttachmentObservedState_v1alpha1_ToProto,
	)

	f.SpecField(".connection_preference")
	f.SpecField(".description")
	f.SpecField(".fingerprint")
	f.SpecField(".producer_accept_lists")
	f.SpecField(".producer_reject_lists")
	f.SpecField(".subnetworks")

	f.StatusField(".connection_endpoints")
	f.StatusField(".creation_timestamp")
	f.StatusField(".id")
	f.StatusField(".kind")
	f.StatusField(".region")
	f.StatusField(".self_link")
	f.StatusField(".self_link_with_id")
	f.StatusField(".network")

	f.Unimplemented_Identity(".name")

	return f
}
