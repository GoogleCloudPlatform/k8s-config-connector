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

package vmwareengine

import (
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	krmcomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VMwareEnginePrivateConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.VMwareEnginePrivateConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEnginePrivateConnectionSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.GetVmwareEngineNetwork() != "" {
		out.VMwareEngineNetworkRef = &krm.VmwareEngineNetworkRef{
			External: in.GetVmwareEngineNetwork(),
		}
	}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.RoutingMode = direct.Enum_FromProto(mapCtx, in.GetRoutingMode())
	if in.GetServiceNetwork() != "" {
		out.ServiceNetworkRef = &krmcomputev1beta1.ComputeNetworkRef{External: in.GetServiceNetwork()}
	}
	return out
}

func VMwareEnginePrivateConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEnginePrivateConnectionSpec) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	out.Description = direct.ValueOf(in.Description)
	if in.VMwareEngineNetworkRef != nil {
		out.VmwareEngineNetwork = in.VMwareEngineNetworkRef.External
	}
	out.Type = direct.Enum_ToProto[pb.PrivateConnection_Type](mapCtx, in.Type)
	out.RoutingMode = direct.Enum_ToProto[pb.PrivateConnection_RoutingMode](mapCtx, in.RoutingMode)
	if in.ServiceNetworkRef != nil {
		out.ServiceNetwork = in.ServiceNetworkRef.External
	}
	return out
}
