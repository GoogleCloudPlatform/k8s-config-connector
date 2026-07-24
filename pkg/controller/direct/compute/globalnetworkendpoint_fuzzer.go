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
// proto.message: google.cloud.compute.v1.NetworkEndpoint
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(computeGlobalNetworkEndpointFuzzer())
}

func computeGlobalNetworkEndpointFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.NetworkEndpoint{},
		computeGlobalNetworkEndpointSpec_FromProto,
		computeGlobalNetworkEndpointSpec_ToProto,
	)

	// Spec fields: the three endpoint-identifying fields mapped to KRM spec.
	f.SpecFields.Insert(".port")
	f.SpecFields.Insert(".fqdn")
	f.SpecFields.Insert(".ip_address")

	// Fields not mapped: only relevant for other NEG types (GCE_VM_IP, etc.)
	// or are internal service fields not exposed in the KRM spec.
	f.UnimplementedFields.Insert(".annotations")
	f.UnimplementedFields.Insert(".instance")
	f.UnimplementedFields.Insert(".ipv6_address")
	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".network_attachment")
	f.UnimplementedFields.Insert(".network_endpoint_type")
	f.UnimplementedFields.Insert(".region")
	f.UnimplementedFields.Insert(".subnetwork")
	f.UnimplementedFields.Insert(".subnetwork_ipv6_address")
	f.UnimplementedFields.Insert(".zone")

	return f
}

func computeGlobalNetworkEndpointSpec_FromProto(_ *direct.MapContext, ep *pb.NetworkEndpoint) *krm.ComputeGlobalNetworkEndpointSpec {
	port, fqdn, ip := NetworkEndpoint_FromProto(ep)
	spec := &krm.ComputeGlobalNetworkEndpointSpec{Port: port}
	if fqdn != "" {
		spec.Fqdn = &fqdn
	}
	if ip != "" {
		spec.IPAddress = &ip
	}
	return spec
}

func computeGlobalNetworkEndpointSpec_ToProto(_ *direct.MapContext, spec *krm.ComputeGlobalNetworkEndpointSpec) *pb.NetworkEndpoint {
	return NetworkEndpoint_ToProto(spec)
}
