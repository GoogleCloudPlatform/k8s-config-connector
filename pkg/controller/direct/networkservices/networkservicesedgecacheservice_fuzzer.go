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
// proto.message: mockgcp.cloud.networkservices.v1.EdgeCacheService
// api.group: networkservices.cnrm.cloud.google.com

package networkservices

import (
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkservices/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(edgeCacheServiceFuzzer())
}

func edgeCacheServiceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.EdgeCacheService{},
		EdgeCacheServiceSpec_FromProto, EdgeCacheServiceSpec_ToProto,
		EdgeCacheServiceStatus_FromProto, EdgeCacheServiceStatus_ToProto,
	)

	f.SpecFields.Insert(".projectRef")
	f.SpecFields.Insert(".resourceID")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")
	f.Unimplemented_LabelsAnnotations(".labels")

	f.SpecField(".description")
	f.SpecField(".disable_quic")
	f.SpecField(".disable_http2")
	f.SpecField(".require_tls")
	f.SpecField(".edge_security_policy")
	f.SpecField(".ssl_policy")
	f.SpecField(".edge_ssl_certificates")
	f.StatusField(".ipv4_addresses")
	f.StatusField(".ipv6_addresses")
	f.SpecField(".log_config")
	f.SpecField(".routing")

	f.FilterSpec = func(in *pb.EdgeCacheService) {
		if in.Routing == nil {
			in.Routing = &pb.EdgeCacheServiceRouting{}
		}
	}

	return f
}
