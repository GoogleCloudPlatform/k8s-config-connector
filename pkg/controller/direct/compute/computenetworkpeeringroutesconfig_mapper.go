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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeNetworkPeeringRoutesConfigSpec_v1alpha1_FromProto maps a proto NetworkPeering message to a KRM ComputeNetworkPeeringRoutesConfigSpec.
// This is hand-coded because of custom reference types and pointer-to-value field conversions (bool vs *bool).
func ComputeNetworkPeeringRoutesConfigSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPeering) *krm.ComputeNetworkPeeringRoutesConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNetworkPeeringRoutesConfigSpec{}
	out.ExportCustomRoutes = direct.ValueOf(in.ExportCustomRoutes)
	out.ImportCustomRoutes = direct.ValueOf(in.ImportCustomRoutes)
	if in.GetNetwork() != "" {
		out.NetworkRef = computev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.ResourceID = in.Name
	return out
}

// ComputeNetworkPeeringRoutesConfigSpec_v1alpha1_ToProto maps a KRM ComputeNetworkPeeringRoutesConfigSpec to a proto NetworkPeering message.
// This is hand-coded because of custom reference types and pointer-to-value field conversions (bool vs *bool).
func ComputeNetworkPeeringRoutesConfigSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNetworkPeeringRoutesConfigSpec) *pb.NetworkPeering {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPeering{}
	out.ExportCustomRoutes = direct.PtrTo(in.ExportCustomRoutes)
	out.ImportCustomRoutes = direct.PtrTo(in.ImportCustomRoutes)
	if in.NetworkRef.External != "" {
		out.Network = direct.LazyPtr(in.NetworkRef.External)
	}
	out.Name = in.ResourceID
	return out
}
