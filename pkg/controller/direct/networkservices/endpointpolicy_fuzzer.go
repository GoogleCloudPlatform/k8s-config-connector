// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.networkservices.v1.EndpointPolicy
// api.group: networkservices.cnrm.cloud.google.com

package networkservices

import (
	pb "google.golang.org/genproto/googleapis/cloud/networkservices/v1"

	// Assuming ref types are needed
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(networkServicesEndpointPolicyFuzzer())
}

func networkServicesEndpointPolicyFuzzer() fuzztesting.KRMFuzzer {
	// NetworkServicesEndpointPolicyObservedState is empty, so no status mappers needed.
	f := fuzztesting.NewKRMTypedFuzzer(&pb.EndpointPolicy{},
		NetworkServicesEndpointPolicySpec_FromProto, NetworkServicesEndpointPolicySpec_ToProto,
		nil, nil,
	)

	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".type")
	f.SpecFields.Insert(".authorization_policy")
	f.SpecFields.Insert(".endpoint_matcher")
	f.SpecFields.Insert(".traffic_port_selector")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".server_tls_policy")
	f.SpecFields.Insert(".client_tls_policy")

	// Status fields are not mapped in KRM observed state as it's empty
	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".create_time")
	f.UnimplementedFields.Insert(".update_time")

	return f
}

// Placeholder for the actual FromProto function
// Replace with the real implementation from your mappers package
func NetworkServicesEndpointPolicySpec_FromProto(ctx *direct.MapContext, in *pb.EndpointPolicy) *krm.NetworkServicesEndpointPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesEndpointPolicySpec{}
	out.Labels = in.Labels
	out.Type = direct.Enum_FromProto(ctx, in.GetType())
	out.AuthorizationPolicy = direct.LazyPtr(in.GetAuthorizationPolicy())
	out.EndpointMatcher = EndpointMatcher_FromProto(ctx, in.GetEndpointMatcher())             // Assuming nested mapper exists
	out.TrafficPortSelector = TrafficPortSelector_FromProto(ctx, in.GetTrafficPortSelector()) // Assuming nested mapper exists
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ServerTLSPolicy = direct.LazyPtr(in.GetServerTlsPolicy())
	out.ClientTLSPolicy = direct.LazyPtr(in.GetClientTlsPolicy())
	// ResourceID/name is handled separately
	return out
}

// Placeholder for the actual ToProto function
// Replace with the real implementation from your mappers package
func NetworkServicesEndpointPolicySpec_ToProto(ctx *direct.MapContext, in *krm.NetworkServicesEndpointPolicySpec) *pb.EndpointPolicy {
	if in == nil {
		return nil
	}
	out := &pb.EndpointPolicy{}
	// Name/ResourceID is handled by the controller/resource mapping
	out.Labels = in.Labels
	out.Type = direct.Enum_ToProto[pb.EndpointPolicy_EndpointPolicyType](ctx, in.Type)
	out.AuthorizationPolicy = direct.ValueOf(in.AuthorizationPolicy)
	out.EndpointMatcher = EndpointMatcher_ToProto(ctx, in.EndpointMatcher)             // Assuming nested mapper exists
	out.TrafficPortSelector = TrafficPortSelector_ToProto(ctx, in.TrafficPortSelector) // Assuming nested mapper exists
	out.Description = direct.ValueOf(in.Description)
	out.ServerTlsPolicy = direct.ValueOf(in.ServerTLSPolicy)
	out.ClientTlsPolicy = direct.ValueOf(in.ClientTLSPolicy)
	return out
}

// Placeholders for nested struct mappers (replace with actual ones if they exist)
func EndpointMatcher_FromProto(ctx *direct.MapContext, in *pb.EndpointMatcher) *krm.EndpointMatcher {
	// Replace with actual implementation
	if in == nil {
		return nil
	}
	return &krm.EndpointMatcher{}
}

func EndpointMatcher_ToProto(ctx *direct.MapContext, in *krm.EndpointMatcher) *pb.EndpointMatcher {
	// Replace with actual implementation
	if in == nil {
		return nil
	}
	return &pb.EndpointMatcher{}
}

func TrafficPortSelector_FromProto(ctx *direct.MapContext, in *pb.TrafficPortSelector) *krm.TrafficPortSelector {
	// Replace with actual implementation
	if in == nil {
		return nil
	}
	return &krm.TrafficPortSelector{
		Ports: in.GetPorts(),
	}
}

func TrafficPortSelector_ToProto(ctx *direct.MapContext, in *krm.TrafficPortSelector) *pb.TrafficPortSelector {
	// Replace with actual implementation
	if in == nil {
		return nil
	}
	return &pb.TrafficPortSelector{
		Ports: in.Ports,
	}
}
