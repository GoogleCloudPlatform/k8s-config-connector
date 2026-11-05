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
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krmcomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkServicesLBEdgeExtensionSpec_FromProto(mapCtx *direct.MapContext, in *pb.LbEdgeExtension) *krm.NetworkServicesLBEdgeExtensionSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesLBEdgeExtensionSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())

	if v := in.GetForwardingRules(); len(v) != 0 {
		for i := range v {
			out.ForwardingRuleRefs = append(out.ForwardingRuleRefs, &krmcomputev1beta1.ForwardingRuleRef{External: v[i]})
		}
	}

	out.ExtensionChains = direct.Slice_FromProto(mapCtx, in.ExtensionChains, ExtensionChain_FromProto)
	out.LoadBalancingScheme = direct.Enum_FromProto(mapCtx, in.GetLoadBalancingScheme())
	return out
}

func NetworkServicesLBEdgeExtensionSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesLBEdgeExtensionSpec) *pb.LbEdgeExtension {
	if in == nil {
		return nil
	}
	out := &pb.LbEdgeExtension{}
	out.Description = direct.ValueOf(in.Description)

	if v := in.ForwardingRuleRefs; len(v) != 0 {
		for i := range v {
			out.ForwardingRules = append(out.ForwardingRules, v[i].External)
		}
	}

	out.ExtensionChains = direct.Slice_ToProto(mapCtx, in.ExtensionChains, ExtensionChain_ToProto)
	out.LoadBalancingScheme = direct.Enum_ToProto[pb.LoadBalancingScheme](mapCtx, in.LoadBalancingScheme)
	return out
}

func NetworkServicesLBEdgeExtensionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LbEdgeExtension) *krm.NetworkServicesLBEdgeExtensionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesLBEdgeExtensionObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func NetworkServicesLBEdgeExtensionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesLBEdgeExtensionObservedState) *pb.LbEdgeExtension {
	if in == nil {
		return nil
	}
	out := &pb.LbEdgeExtension{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
