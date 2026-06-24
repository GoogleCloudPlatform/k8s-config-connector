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

package networkservices

import (
	"strings"

	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krmcomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkServicesLBRouteExtensionSpec_FromProto(mapCtx *direct.MapContext, in *pb.LbRouteExtension) *krm.NetworkServicesLBRouteExtensionSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesLBRouteExtensionSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())

	if v := in.GetForwardingRules(); len(v) != 0 {
		for i := range v {
			out.ForwardingRuleRefs = append(out.ForwardingRuleRefs, &krmcomputev1beta1.ForwardingRuleRef{External: v[i]})
		}
	}

	out.ExtensionChains = direct.Slice_FromProto(mapCtx, in.ExtensionChains, ExtensionChain_FromProto)
	out.LoadBalancingScheme = direct.Enum_FromProto(mapCtx, in.GetLoadBalancingScheme())
	out.Metadata = direct.Struct_FromProto(mapCtx, in.GetMetadata())
	return out
}

func NetworkServicesLBRouteExtensionSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesLBRouteExtensionSpec) *pb.LbRouteExtension {
	if in == nil {
		return nil
	}
	out := &pb.LbRouteExtension{}
	out.Description = direct.ValueOf(in.Description)

	if v := in.ForwardingRuleRefs; len(v) != 0 {
		for i := range v {
			out.ForwardingRules = append(out.ForwardingRules, v[i].External)
		}
	}

	out.ExtensionChains = direct.Slice_ToProto(mapCtx, in.ExtensionChains, ExtensionChain_ToProto)
	out.LoadBalancingScheme = direct.Enum_ToProto[pb.LoadBalancingScheme](mapCtx, in.LoadBalancingScheme)
	out.Metadata = direct.Struct_ToProto(mapCtx, in.Metadata)
	return out
}

func NetworkServicesLBRouteExtensionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LbRouteExtension) *krm.NetworkServicesLBRouteExtensionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesLBRouteExtensionObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func NetworkServicesLBRouteExtensionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesLBRouteExtensionObservedState) *pb.LbRouteExtension {
	if in == nil {
		return nil
	}
	out := &pb.LbRouteExtension{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func ExtensionChain_FromProto(mapCtx *direct.MapContext, in *pb.ExtensionChain) *krm.ExtensionChain {
	if in == nil {
		return nil
	}
	out := &krm.ExtensionChain{}
	out.Name = direct.LazyPtr(in.GetName())
	out.MatchCondition = ExtensionChain_MatchCondition_FromProto(mapCtx, in.GetMatchCondition())
	out.Extensions = direct.Slice_FromProto(mapCtx, in.Extensions, ExtensionChain_Extension_FromProto)
	return out
}

func ExtensionChain_ToProto(mapCtx *direct.MapContext, in *krm.ExtensionChain) *pb.ExtensionChain {
	if in == nil {
		return nil
	}
	out := &pb.ExtensionChain{}
	out.Name = direct.ValueOf(in.Name)
	out.MatchCondition = ExtensionChain_MatchCondition_ToProto(mapCtx, in.MatchCondition)
	out.Extensions = direct.Slice_ToProto(mapCtx, in.Extensions, ExtensionChain_Extension_ToProto)
	return out
}

func ExtensionChain_MatchCondition_FromProto(mapCtx *direct.MapContext, in *pb.ExtensionChain_MatchCondition) *krm.ExtensionChain_MatchCondition {
	if in == nil {
		return nil
	}
	out := &krm.ExtensionChain_MatchCondition{}
	out.CelExpression = direct.LazyPtr(in.GetCelExpression())
	return out
}

func ExtensionChain_MatchCondition_ToProto(mapCtx *direct.MapContext, in *krm.ExtensionChain_MatchCondition) *pb.ExtensionChain_MatchCondition {
	if in == nil {
		return nil
	}
	out := &pb.ExtensionChain_MatchCondition{}
	out.CelExpression = direct.ValueOf(in.CelExpression)
	return out
}

func ExtensionChain_Extension_FromProto(mapCtx *direct.MapContext, in *pb.ExtensionChain_Extension) *krm.ExtensionChain_Extension {
	if in == nil {
		return nil
	}
	out := &krm.ExtensionChain_Extension{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Authority = direct.LazyPtr(in.GetAuthority())

	service := in.GetService()
	if service != "" {
		if strings.Contains(service, "/backendServices/") {
			out.BackendServiceRef = &krmcomputev1beta1.ComputeBackendServiceRef{External: service}
		} else if strings.Contains(service, "/wasmPlugins/") {
			out.WasmPluginRef = &krm.NetworkServicesWasmPluginRef{External: service}
		}
	}

	out.SupportedEvents = direct.EnumSlice_FromProto(mapCtx, in.SupportedEvents)
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.FailOpen = direct.LazyPtr(in.GetFailOpen())
	out.ForwardHeaders = in.GetForwardHeaders()
	out.Metadata = direct.Struct_FromProto(mapCtx, in.GetMetadata())
	return out
}

func ExtensionChain_Extension_ToProto(mapCtx *direct.MapContext, in *krm.ExtensionChain_Extension) *pb.ExtensionChain_Extension {
	if in == nil {
		return nil
	}
	out := &pb.ExtensionChain_Extension{}
	out.Name = direct.ValueOf(in.Name)
	out.Authority = direct.ValueOf(in.Authority)

	if in.BackendServiceRef != nil {
		out.Service = in.BackendServiceRef.External
	} else if in.WasmPluginRef != nil {
		out.Service = in.WasmPluginRef.External
	}

	out.SupportedEvents = direct.EnumSlice_ToProto[pb.EventType](mapCtx, in.SupportedEvents)
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.FailOpen = direct.ValueOf(in.FailOpen)
	out.ForwardHeaders = in.ForwardHeaders
	out.Metadata = direct.Struct_ToProto(mapCtx, in.Metadata)
	return out
}
