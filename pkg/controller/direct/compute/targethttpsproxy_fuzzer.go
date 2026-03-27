// Copyright 2025 Google LLC
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
// proto.message: google.cloud.compute.v1.TargetHttpsProxy
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeTargetHTTPSProxyFuzzer())
}

func computeTargetHTTPSProxyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TargetHttpsProxy{},
		ComputeTargetHTTPSProxySpec_v1beta1_FromProto, ComputeTargetHTTPSProxySpec_v1beta1_ToProto,
		ComputeTargetHTTPSProxyObservedState_v1beta1_FromProto, ComputeTargetHTTPSProxyObservedState_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecFields.Insert(".authorization_policy")
	f.SpecFields.Insert(".certificate_map")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".http_keep_alive_timeout_sec")
	f.SpecFields.Insert(".proxy_bind")
	f.SpecFields.Insert(".quic_override")
	f.SpecFields.Insert(".server_tls_policy")
	f.SpecFields.Insert(".ssl_certificates")
	f.SpecFields.Insert(".ssl_policy")
	f.SpecFields.Insert(".tls_early_data")
	f.SpecFields.Insert(".url_map")

	// KRM-only spec fields
	f.SpecFields.Insert(".location")
	f.SpecFields.Insert(".resourceID")

	// Status fields
	f.StatusFields.Insert(".creation_timestamp")
	f.StatusFields.Insert(".id")
	f.StatusFields.Insert(".self_link")
	f.StatusFields.Insert(".fingerprint")

	// KRM-only status fields
	f.StatusFields.Insert(".observedGeneration")
	f.StatusFields.Insert(".externalRef")

	// Unimplemented proto fields
	f.UnimplementedFields.Insert(".kind")
	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".region")
	f.UnimplementedFields.Insert(".authorization_policy")
	f.UnimplementedFields.Insert(".tls_early_data")
	f.UnimplementedFields.Insert(".creation_timestamp")
	f.UnimplementedFields.Insert(".id")
	f.UnimplementedFields.Insert(".self_link")

	f.FilterSpec = func(in *pb.TargetHttpsProxy) {
	}
	f.FilterStatus = func(in *pb.TargetHttpsProxy) {
	}

	return f
}
