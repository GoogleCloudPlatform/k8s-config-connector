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

func NetworkServicesAuthzExtensionSpec_FromProto(mapCtx *direct.MapContext, in *pb.AuthzExtension) *krm.NetworkServicesAuthzExtensionSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesAuthzExtensionSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	out.LoadBalancingScheme = direct.Enum_FromProto(mapCtx, in.GetLoadBalancingScheme())
	out.Authority = direct.LazyPtr(in.GetAuthority())

	if in.GetService() != "" {
		out.BackendServiceRef = &krmcomputev1beta1.ComputeBackendServiceRef{External: in.GetService()}
	}

	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.FailOpen = direct.LazyPtr(in.GetFailOpen())
	out.Metadata = direct.Struct_FromProto(mapCtx, in.GetMetadata())
	out.ForwardHeaders = in.ForwardHeaders
	out.WireFormat = direct.Enum_FromProto(mapCtx, in.GetWireFormat())
	return out
}

func NetworkServicesAuthzExtensionSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesAuthzExtensionSpec) *pb.AuthzExtension {
	if in == nil {
		return nil
	}
	out := &pb.AuthzExtension{}
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.LoadBalancingScheme = direct.Enum_ToProto[pb.LoadBalancingScheme](mapCtx, in.LoadBalancingScheme)
	out.Authority = direct.ValueOf(in.Authority)

	if in.BackendServiceRef != nil {
		out.Service = in.BackendServiceRef.External
	}

	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.FailOpen = direct.ValueOf(in.FailOpen)
	out.Metadata = direct.Struct_ToProto(mapCtx, in.Metadata)
	out.ForwardHeaders = in.ForwardHeaders
	out.WireFormat = direct.Enum_ToProto[pb.WireFormat](mapCtx, in.WireFormat)
	return out
}

func NetworkServicesAuthzExtensionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AuthzExtension) *krm.NetworkServicesAuthzExtensionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesAuthzExtensionObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func NetworkServicesAuthzExtensionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesAuthzExtensionObservedState) *pb.AuthzExtension {
	if in == nil {
		return nil
	}
	out := &pb.AuthzExtension{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
