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
	fuzztesting.RegisterKRMSpecFuzzer(computeRegionPerInstanceConfigFuzzer())
}

func computeRegionPerInstanceConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.PerInstanceConfig{},
		ComputeRegionPerInstanceConfigSpec_v1alpha1_FromProto, ComputeRegionPerInstanceConfigSpec_v1alpha1_ToProto,
	)

	// Field Comparison & Mapping Documentation:
	//
	// KRM Spec Fields:
	// - minimalAction                -> Not in pb.PerInstanceConfig (action param in API requests)
	// - mostDisruptiveAllowedAction  -> Not in pb.PerInstanceConfig (action param in API requests)
	// - preservedState.disk          -> f.SpecField(".preserved_state.disks")
	//   - deleteRule                 -> f.SpecField(".preserved_state.disks[].auto_delete")
	//   - deviceName                 -> Map key for .preserved_state.disks
	//   - mode                       -> f.SpecField(".preserved_state.disks[].mode")
	//   - source                     -> f.SpecField(".preserved_state.disks[].source")
	// - preservedState.externalIp    -> f.SpecField(".preserved_state.external_i_ps")
	//   - autoDelete                 -> f.SpecField(".preserved_state.external_i_ps[].auto_delete")
	//   - interfaceName              -> Map key for .preserved_state.external_i_ps
	//   - ipAddress.address          -> f.SpecField(".preserved_state.external_i_ps[].ip_address.address")
	// - preservedState.internalIp    -> f.SpecField(".preserved_state.internal_i_ps")
	//   - autoDelete                 -> f.SpecField(".preserved_state.internal_i_ps[].auto_delete")
	//   - interfaceName              -> Map key for .preserved_state.internal_i_ps
	//   - ipAddress.address          -> f.SpecField(".preserved_state.internal_i_ps[].ip_address.address")
	// - preservedState.metadata      -> f.SpecField(".preserved_state.metadata")
	// - projectRef                   -> Not in pb.PerInstanceConfig (resolved from identity/parent URL)
	// - region                       -> Not in pb.PerInstanceConfig (resolved from identity/parent URL)
	// - regionInstanceGroupManagerRef-> Not in pb.PerInstanceConfig (resolved from identity/parent URL)
	// - removeInstanceStateOnDestroy -> Not in pb.PerInstanceConfig (client-side deletion behavior)
	// - resourceID                   -> f.Unimplemented_Identity(".name")
	//
	// KRM Status Fields:
	// - conditions                   -> Standard KCC conditions (no Status in pb.PerInstanceConfig mapped)
	// - observedGeneration           -> Standard KCC generation
	//
	// Unimplemented Proto Fields:
	// - fingerprint                  -> f.Unimplemented_NotYetTriaged(".fingerprint")
	// - status                       -> f.Unimplemented_NotYetTriaged(".status")

	// Spec fields
	f.SpecField(".preserved_state")
	f.SpecField(".preserved_state.disks")
	f.SpecField(".preserved_state.disks[].auto_delete")
	f.SpecField(".preserved_state.disks[].mode")
	f.SpecField(".preserved_state.disks[].source")

	f.SpecField(".preserved_state.external_i_ps")
	f.SpecField(".preserved_state.external_i_ps[].auto_delete")
	f.SpecField(".preserved_state.external_i_ps[].ip_address")
	f.SpecField(".preserved_state.external_i_ps[].ip_address.address")

	f.SpecField(".preserved_state.internal_i_ps")
	f.SpecField(".preserved_state.internal_i_ps[].auto_delete")
	f.SpecField(".preserved_state.internal_i_ps[].ip_address")
	f.SpecField(".preserved_state.internal_i_ps[].ip_address.address")

	f.SpecField(".preserved_state.metadata")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".fingerprint")
	f.Unimplemented_NotYetTriaged(".status")

	f.FilterSpec = func(in *pb.PerInstanceConfig) {
		if in.PreservedState != nil {
			for _, disk := range in.PreservedState.Disks {
				if disk != nil {
					if disk.Source != nil && *disk.Source == "" {
						disk.Source = nil
					}
				}
			}
			for _, ip := range in.PreservedState.ExternalIPs {
				if ip != nil {
					if ip.IpAddress != nil {
						if ip.IpAddress.Address == nil || *ip.IpAddress.Address == "" {
							ip.IpAddress = nil
						} else {
							ip.IpAddress.Literal = nil
						}
					}
				}
			}
			for _, ip := range in.PreservedState.InternalIPs {
				if ip != nil {
					if ip.IpAddress != nil {
						if ip.IpAddress.Address == nil || *ip.IpAddress.Address == "" {
							ip.IpAddress = nil
						} else {
							ip.IpAddress.Literal = nil
						}
					}
				}
			}
		}
	}

	return f
}
