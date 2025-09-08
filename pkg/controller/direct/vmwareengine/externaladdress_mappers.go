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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// func VMwareEngineExternalAddressSpec_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAddress) *krm.VMwareEngineExternalAddressSpec {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.VMwareEngineExternalAddressSpec{}
// 	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
// 	out.Description = direct.LazyPtr(in.GetDescription())
// 	return out
// }

// func VMwareEngineExternalAddressSpec_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEngineExternalAddressSpec) *pb.ExternalAddress {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.ExternalAddress{}
// 	out.InternalIp = direct.ValueOf(in.InternalIP)
// 	out.Description = direct.ValueOf(in.Description)
// 	return out
// }

func VMwareEngineExternalAddressObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAddress) *krm.VMwareEngineExternalAddressObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEngineExternalAddressObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ExternalIP = direct.LazyPtr(in.GetExternalIp())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UID = direct.LazyPtr(in.GetUid())
	return out
}

func VMwareEngineExternalAddressObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEngineExternalAddressObservedState) *pb.ExternalAddress {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAddress{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ExternalIp = direct.ValueOf(in.ExternalIP)
	out.State = direct.Enum_ToProto[pb.ExternalAddress_State](mapCtx, in.State)
	// MISSING: Uid
	// (near miss): "Uid" vs "UID"
	return out
}
