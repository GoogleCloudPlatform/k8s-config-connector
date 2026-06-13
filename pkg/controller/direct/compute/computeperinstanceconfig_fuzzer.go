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
// proto.message: google.cloud.compute.v1.PerInstanceConfig
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(computePerInstanceConfigFuzzer())
}

func computePerInstanceConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.PerInstanceConfig{},
		ComputePerInstanceConfigSpec_v1alpha1_FromProto, ComputePerInstanceConfigSpec_v1alpha1_ToProto,
	)

	f.SpecField(".preserved_state")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_Internal(".fingerprint")
	f.Unimplemented_Internal(".status")

	f.FilterSpec = func(in *pb.PerInstanceConfig) {
		if in.PreservedState != nil {
			for _, ip := range in.PreservedState.ExternalIPs {
				if ip.IpAddress != nil {
					ip.IpAddress.Literal = nil
				}
			}
			for _, ip := range in.PreservedState.InternalIPs {
				if ip.IpAddress != nil {
					ip.IpAddress.Literal = nil
				}
			}
			for _, disk := range in.PreservedState.Disks {
				if disk.Source == nil || *disk.Source == "" {
					disk.Source = nil
				}
			}
		}
	}

	return f
}
