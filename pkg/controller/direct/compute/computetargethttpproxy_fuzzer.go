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
// proto.message: google.cloud.compute.v1.TargetHttpProxy
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeTargetHTTPProxyFuzzer())
}

func computeTargetHTTPProxyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TargetHttpProxy{},
		ComputeTargetHTTPProxySpec_v1beta1_FromProto, ComputeTargetHTTPProxySpec_v1beta1_ToProto,
		ComputeTargetHTTPProxyStatus_v1beta1_FromProto, func(ctx *direct.MapContext, in *krm.ComputeTargetHTTPProxyStatus) *pb.TargetHttpProxy {
			if in == nil {
				return nil
			}
			out := &pb.TargetHttpProxy{}
			out.CreationTimestamp = in.CreationTimestamp
			if in.ProxyId != nil {
				id := uint64(*in.ProxyId)
				out.Id = &id
			}
			out.SelfLink = in.SelfLink
			return out
		},
	)

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".http_keep_alive_timeout_sec")
	f.SpecField(".proxy_bind")
	f.SpecField(".url_map")

	// KRM-only spec fields
	f.SpecField(".location")
	f.SpecField(".resourceID")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".id")
	f.StatusField(".self_link")

	// KRM-only status fields
	f.StatusField(".observedGeneration")

	// Unimplemented proto fields
	f.Unimplemented_Identity(".kind")
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".region")
	f.Unimplemented_Identity(".creation_timestamp")
	f.Unimplemented_Identity(".id")
	f.Unimplemented_Identity(".self_link")
	f.Unimplemented_Identity(".fingerprint")

	f.FilterSpec = func(in *pb.TargetHttpProxy) {
	}
	f.FilterStatus = func(in *pb.TargetHttpProxy) {
	}

	return f
}
