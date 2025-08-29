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
// proto.message: google.cloud.compute.v1.TargetTcpProxy
// api.group: compute.cnrm.cloud.google.com

package targettcpproxy

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeTargetTCPProxyFuzzer())
}

func computeTargetTCPProxyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TargetTcpProxy{},
		ComputeTargetTCPProxySpec_FromProto, ComputeTargetTCPProxySpec_ToProto,
		ComputeTargetTCPProxyStatus_FromProto, ComputeTargetTCPProxyStatus_ToProto,
	)

	// Spec fields
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".proxy_bind")
	f.SpecFields.Insert(".proxy_header")
	f.SpecFields.Insert(".region")
	f.SpecFields.Insert(".service")

	// Status fields
	f.StatusFields.Insert(".creation_timestamp")
	f.StatusFields.Insert(".id")
	f.StatusFields.Insert(".self_link")

	// Unimplemented fields
	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".kind")

	return f
}
