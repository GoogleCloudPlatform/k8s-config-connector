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
// proto.message: google.cloud.compute.v1.Firewall
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeFirewallFuzzer())
}

func computeFirewallFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Firewall{},
		ComputeFirewallSpec_v1beta1_FromProto, ComputeFirewallSpec_v1beta1_ToProto,
		ComputeFirewallStatus_v1beta1_FromProto, ComputeFirewallStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".allowed")
	f.SpecField(".denied")
	f.SpecField(".description")
	f.SpecField(".destination_ranges")
	f.SpecField(".direction")
	f.SpecField(".disabled")
	f.SpecField(".log_config")
	f.SpecField(".log_config.metadata")
	f.SpecField(".network")
	f.SpecField(".priority")
	f.SpecField(".source_ranges")
	f.SpecField(".source_service_accounts")
	f.SpecField(".source_tags")
	f.SpecField(".target_service_accounts")
	f.SpecField(".target_tags")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".self_link")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".log_config.enable")
	f.Unimplemented_NotYetTriaged(".params")

	f.FilterSpec = func(in *pb.Firewall) {
		if in.LogConfig != nil {
			in.LogConfig.Enable = direct.LazyPtr(true)
		}
	}

	return f
}
