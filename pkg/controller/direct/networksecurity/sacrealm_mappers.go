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

package networksecurity

import (
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkSecuritySACRealmObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SACRealm) *krm.NetworkSecuritySACRealmObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecuritySACRealmObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.PairingKey = SACRealm_PairingKeyObservedState_v1alpha1_FromProto(mapCtx, in.GetPairingKey())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}

func NetworkSecuritySACRealmObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecuritySACRealmObservedState) *pb.SACRealm {
	if in == nil {
		return nil
	}
	out := &pb.SACRealm{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.PairingKey = SACRealm_PairingKeyObservedState_v1alpha1_ToProto(mapCtx, in.PairingKey)
	out.State = direct.Enum_ToProto[pb.SACRealm_State](mapCtx, in.State)
	return out
}

func NetworkSecuritySACRealmSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SACRealm) *krm.NetworkSecuritySACRealmSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecuritySACRealmSpec{}
	out.Labels = in.Labels
	out.SecurityService = direct.Enum_FromProto(mapCtx, in.GetSecurityService())
	return out
}

func NetworkSecuritySACRealmSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecuritySACRealmSpec) *pb.SACRealm {
	if in == nil {
		return nil
	}
	out := &pb.SACRealm{}
	out.Labels = in.Labels
	out.SecurityService = direct.Enum_ToProto[pb.SACRealm_SecurityService](mapCtx, in.SecurityService)
	return out
}

func SACRealm_PairingKeyObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SACRealm_PairingKey) *krm.SACRealm_PairingKeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SACRealm_PairingKeyObservedState{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	return out
}

func SACRealm_PairingKeyObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.SACRealm_PairingKeyObservedState) *pb.SACRealm_PairingKey {
	if in == nil {
		return nil
	}
	out := &pb.SACRealm_PairingKey{}
	out.Key = direct.ValueOf(in.Key)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	return out
}
