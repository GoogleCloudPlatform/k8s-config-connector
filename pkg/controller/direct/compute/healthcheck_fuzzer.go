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
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeHealthCheckFuzzer())
}

func computeHealthCheckFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.HealthCheck{},
		ComputeHealthCheckSpec_v1beta1_FromProto, ComputeHealthCheckSpec_v1beta1_ToProto,
		ComputeHealthCheckStatus_v1beta1_FromProto, ComputeHealthCheckStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".check_interval_sec")
	f.SpecField(".description")
	f.SpecField(".grpc_health_check")
	f.SpecField(".grpc_health_check.grpc_service_name")
	f.SpecField(".grpc_health_check.port")
	f.SpecField(".grpc_health_check.port_name")
	f.SpecField(".grpc_health_check.port_specification")
	f.SpecField(".healthy_threshold")
	f.SpecField(".http2_health_check")
	f.SpecField(".http2_health_check.host")
	f.SpecField(".http2_health_check.port")
	f.SpecField(".http2_health_check.port_name")
	f.SpecField(".http2_health_check.port_specification")
	f.SpecField(".http2_health_check.proxy_header")
	f.SpecField(".http2_health_check.request_path")
	f.SpecField(".http2_health_check.response")
	f.SpecField(".http_health_check")
	f.SpecField(".http_health_check.host")
	f.SpecField(".http_health_check.port")
	f.SpecField(".http_health_check.port_name")
	f.SpecField(".http_health_check.port_specification")
	f.SpecField(".http_health_check.proxy_header")
	f.SpecField(".http_health_check.request_path")
	f.SpecField(".http_health_check.response")
	f.SpecField(".https_health_check")
	f.SpecField(".https_health_check.host")
	f.SpecField(".https_health_check.port")
	f.SpecField(".https_health_check.port_name")
	f.SpecField(".https_health_check.port_specification")
	f.SpecField(".https_health_check.proxy_header")
	f.SpecField(".https_health_check.request_path")
	f.SpecField(".https_health_check.response")
	f.SpecField(".log_config")
	f.SpecField(".log_config.enable")
	f.SpecField(".ssl_health_check")
	f.SpecField(".ssl_health_check.port")
	f.SpecField(".ssl_health_check.port_name")
	f.SpecField(".ssl_health_check.port_specification")
	f.SpecField(".ssl_health_check.proxy_header")
	f.SpecField(".ssl_health_check.request")
	f.SpecField(".ssl_health_check.response")
	f.SpecField(".tcp_health_check")
	f.SpecField(".tcp_health_check.port")
	f.SpecField(".tcp_health_check.port_name")
	f.SpecField(".tcp_health_check.port_specification")
	f.SpecField(".tcp_health_check.proxy_header")
	f.SpecField(".tcp_health_check.request")
	f.SpecField(".tcp_health_check.response")
	f.SpecField(".timeout_sec")
	f.SpecField(".unhealthy_threshold")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".self_link")
	f.StatusField(".type")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".region")
	f.Unimplemented_NotYetTriaged(".source_regions")
	f.Unimplemented_NotYetTriaged(".grpc_tls_health_check")

	f.FilterSpec = func(in *pb.HealthCheck) {
		clearIfEmpty(&in.GrpcHealthCheck)
		clearIfEmpty(&in.Http2HealthCheck)
		clearIfEmpty(&in.HttpHealthCheck)
		clearIfEmpty(&in.HttpsHealthCheck)
		clearIfEmpty(&in.LogConfig)
		clearIfEmpty(&in.SslHealthCheck)
		clearIfEmpty(&in.TcpHealthCheck)
	}

	return f
}

func clearIfEmpty[T proto.Message](p *T) {
	if p == nil {
		return
	}
	var zero T
	if any(*p) == any(zero) {
		return
	}
	if isMessageEmpty((*p).ProtoReflect()) {
		*p = zero
	}
}

func isMessageEmpty(m protoreflect.Message) bool {
	if m == nil {
		return true
	}
	isEmpty := true
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.IsList() || fd.IsMap() {
			isEmpty = false
			return false
		}
		if fd.Kind() == protoreflect.MessageKind {
			if !isMessageEmpty(v.Message()) {
				isEmpty = false
				return false
			}
		} else {
			isEmpty = false
			return false
		}
		return true
	})
	return isEmpty
}
