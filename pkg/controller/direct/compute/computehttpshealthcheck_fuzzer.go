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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeHTTPSHealthCheckFuzzer())
}

func computeHTTPSHealthCheckFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.HealthCheck{},
		ComputeHTTPSHealthCheckSpec_v1beta1_FromProto, ComputeHTTPSHealthCheckSpec_v1beta1_ToProto,
		ComputeHTTPSHealthCheckStatus_v1beta1_FromProto, ComputeHTTPSHealthCheckStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".check_interval_sec")
	f.SpecField(".description")
	f.SpecField(".healthy_threshold")
	f.SpecField(".timeout_sec")
	f.SpecField(".unhealthy_threshold")
	f.SpecField(".https_health_check")
	f.SpecField(".https_health_check.host")
	f.SpecField(".https_health_check.port")
	f.SpecField(".https_health_check.request_path")

	// KRM-only Spec fields
	f.SpecField(".resourceID")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".self_link")

	// KRM-only status fields
	f.StatusField(".observedGeneration")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".region")
	f.Unimplemented_NotYetTriaged(".type")
	f.Unimplemented_NotYetTriaged(".log_config")
	f.Unimplemented_NotYetTriaged(".log_config.enable")
	f.Unimplemented_NotYetTriaged(".grpc_health_check")
	f.Unimplemented_NotYetTriaged(".grpc_health_check.grpc_service_name")
	f.Unimplemented_NotYetTriaged(".grpc_health_check.port")
	f.Unimplemented_NotYetTriaged(".grpc_health_check.port_name")
	f.Unimplemented_NotYetTriaged(".grpc_health_check.port_specification")
	f.Unimplemented_NotYetTriaged(".http2_health_check")
	f.Unimplemented_NotYetTriaged(".http2_health_check.host")
	f.Unimplemented_NotYetTriaged(".http2_health_check.port")
	f.Unimplemented_NotYetTriaged(".http2_health_check.port_name")
	f.Unimplemented_NotYetTriaged(".http2_health_check.port_specification")
	f.Unimplemented_NotYetTriaged(".http2_health_check.proxy_header")
	f.Unimplemented_NotYetTriaged(".http2_health_check.request_path")
	f.Unimplemented_NotYetTriaged(".http2_health_check.response")
	f.Unimplemented_NotYetTriaged(".http_health_check")
	f.Unimplemented_NotYetTriaged(".http_health_check.host")
	f.Unimplemented_NotYetTriaged(".http_health_check.port")
	f.Unimplemented_NotYetTriaged(".http_health_check.port_name")
	f.Unimplemented_NotYetTriaged(".http_health_check.port_specification")
	f.Unimplemented_NotYetTriaged(".http_health_check.proxy_header")
	f.Unimplemented_NotYetTriaged(".http_health_check.request_path")
	f.Unimplemented_NotYetTriaged(".http_health_check.response")
	f.Unimplemented_NotYetTriaged(".ssl_health_check")
	f.Unimplemented_NotYetTriaged(".ssl_health_check.port")
	f.Unimplemented_NotYetTriaged(".ssl_health_check.port_name")
	f.Unimplemented_NotYetTriaged(".ssl_health_check.port_specification")
	f.Unimplemented_NotYetTriaged(".ssl_health_check.proxy_header")
	f.Unimplemented_NotYetTriaged(".ssl_health_check.request")
	f.Unimplemented_NotYetTriaged(".ssl_health_check.response")
	f.Unimplemented_NotYetTriaged(".tcp_health_check")
	f.Unimplemented_NotYetTriaged(".tcp_health_check.port")
	f.Unimplemented_NotYetTriaged(".tcp_health_check.port_name")
	f.Unimplemented_NotYetTriaged(".tcp_health_check.port_specification")
	f.Unimplemented_NotYetTriaged(".tcp_health_check.proxy_header")
	f.Unimplemented_NotYetTriaged(".tcp_health_check.request")
	f.Unimplemented_NotYetTriaged(".tcp_health_check.response")

	f.Unimplemented_NotYetTriaged(".grpc_tls_health_check")
	f.Unimplemented_NotYetTriaged(".grpc_tls_health_check.grpc_service_name")
	f.Unimplemented_NotYetTriaged(".grpc_tls_health_check.port")
	f.Unimplemented_NotYetTriaged(".grpc_tls_health_check.port_name")
	f.Unimplemented_NotYetTriaged(".grpc_tls_health_check.port_specification")
	f.Unimplemented_NotYetTriaged(".source_regions")

	f.Unimplemented_NotYetTriaged(".https_health_check.port_name")
	f.Unimplemented_NotYetTriaged(".https_health_check.port_specification")
	f.Unimplemented_NotYetTriaged(".https_health_check.proxy_header")
	f.Unimplemented_NotYetTriaged(".https_health_check.response")

	f.FilterSpec = func(in *pb.HealthCheck) {
		in.GrpcHealthCheck = nil
		in.GrpcTlsHealthCheck = nil
		in.Http2HealthCheck = nil
		in.HttpHealthCheck = nil
		in.LogConfig = nil
		in.SslHealthCheck = nil
		in.TcpHealthCheck = nil
		in.Type = nil
		in.SourceRegions = nil
		if in.HttpsHealthCheck != nil {
			in.HttpsHealthCheck.PortName = nil
			in.HttpsHealthCheck.PortSpecification = nil
			in.HttpsHealthCheck.ProxyHeader = nil
			in.HttpsHealthCheck.Response = nil
			if in.HttpsHealthCheck.Host == nil && in.HttpsHealthCheck.Port == nil && in.HttpsHealthCheck.RequestPath == nil {
				in.HttpsHealthCheck = nil
			}
		}
	}

	f.FilterStatus = func(in *pb.HealthCheck) {
		in.GrpcHealthCheck = nil
		in.GrpcTlsHealthCheck = nil
		in.Http2HealthCheck = nil
		in.HttpHealthCheck = nil
		in.HttpsHealthCheck = nil
		in.LogConfig = nil
		in.SslHealthCheck = nil
		in.TcpHealthCheck = nil
		in.Type = nil
		in.SourceRegions = nil
	}

	return f
}
