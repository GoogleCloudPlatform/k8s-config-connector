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
// proto.message: google.cloud.compute.v1.Metadata
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	"sort"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeProjectMetadataFuzzer())
}

func computeProjectMetadataFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Metadata{},
		ComputeProjectMetadataSpec_FromProto, ComputeProjectMetadataSpec_ToProto,
		ComputeProjectMetadataStatus_FromProto, ComputeProjectMetadataStatus_ToProto,
	)

	// Spec fields
	f.SpecField(".items")

	// FilterSpec to handle map-like repeated fields
	f.FilterSpec = func(in *pb.Metadata) {
		if in == nil {
			return
		}
		var cleanItems []*pb.Items
		seen := make(map[string]bool)
		for _, item := range in.Items {
			if item == nil {
				continue
			}
			k := item.GetKey()
			if k == "" {
				continue
			}
			if seen[k] {
				continue
			}
			seen[k] = true
			if item.Value == nil {
				var empty string
				item.Value = &empty
			}
			cleanItems = append(cleanItems, item)
		}
		in.Items = cleanItems
		sort.Slice(in.Items, func(i, j int) bool {
			return in.Items[i].GetKey() < in.Items[j].GetKey()
		})
	}

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".fingerprint")
	f.Unimplemented_NotYetTriaged(".kind")

	return f
}
