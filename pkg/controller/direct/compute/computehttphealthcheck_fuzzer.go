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
// proto.message: google.cloud.compute.v1.HealthCheck
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeHTTPHealthCheckFuzzer())
}

func computeHTTPHealthCheckFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.HealthCheck{},
		ComputeHTTPHealthCheckSpec_v1beta1_FromProto, ComputeHTTPHealthCheckSpec_v1beta1_ToProto,
		ComputeHTTPHealthCheckStatus_v1beta1_FromProto, ComputeHTTPHealthCheckStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".check_interval_sec")
	f.SpecField(".description")
	f.SpecField(".healthy_threshold")
	f.SpecField(".timeout_sec")
	f.SpecField(".unhealthy_threshold")
	f.SpecField(".http_health_check")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".self_link")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".region")
	f.Unimplemented_NotYetTriaged(".type")
	f.Unimplemented_NotYetTriaged(".grpc_health_check")
	f.Unimplemented_NotYetTriaged(".grpc_tls_health_check")
	f.Unimplemented_NotYetTriaged(".http2_health_check")
	f.Unimplemented_NotYetTriaged(".https_health_check")
	f.Unimplemented_NotYetTriaged(".log_config")
	f.Unimplemented_NotYetTriaged(".ssl_health_check")
	f.Unimplemented_NotYetTriaged(".tcp_health_check")
	f.Unimplemented_NotYetTriaged(".source_regions")

	f.Unimplemented_NotYetTriaged(".http_health_check.port_name")
	f.Unimplemented_NotYetTriaged(".http_health_check.port_specification")
	f.Unimplemented_NotYetTriaged(".http_health_check.proxy_header")
	f.Unimplemented_NotYetTriaged(".http_health_check.response")

	f.FilterSpec = func(in *pb.HealthCheck) {
		in.Type = direct.PtrTo("HTTP")
		if in.HttpHealthCheck != nil {
			if in.HttpHealthCheck.Host == nil && in.HttpHealthCheck.Port == nil && in.HttpHealthCheck.RequestPath == nil {
				in.HttpHealthCheck = nil
			}
		}
	}

	return f
}
