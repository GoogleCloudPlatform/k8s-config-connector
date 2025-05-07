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

package vmwareengine

import (
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VMwareEngineNetworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork) *krm.VMwareEngineNetworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEngineNetworkObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.VPCNetworks = direct.Slice_FromProto(mapCtx, in.VpcNetworks, VmwareEngineNetwork_VpcNetworkObservedState_FromProto)
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UID = direct.LazyPtr(in.GetUid())
	return out
}
func VMwareEngineNetworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEngineNetworkObservedState) *pb.VmwareEngineNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.VpcNetworks = direct.Slice_ToProto(mapCtx, in.VPCNetworks, VmwareEngineNetwork_VpcNetworkObservedState_ToProto)
	out.State = direct.Enum_ToProto[pb.VmwareEngineNetwork_State](mapCtx, in.State)
	out.Uid = direct.ValueOf(in.UID)
	return out
}
