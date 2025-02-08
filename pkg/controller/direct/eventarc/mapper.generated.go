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

package eventarc

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
)
func Channel_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.Channel {
	if in == nil {
		return nil
	}
	out := &krm.Channel{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Provider = direct.LazyPtr(in.GetProvider())
	// MISSING: PubsubTopic
	// MISSING: State
	// MISSING: ActivationToken
	out.CryptoKeyName = direct.LazyPtr(in.GetCryptoKeyName())
	// MISSING: SatisfiesPzs
	return out
}
func Channel_ToProto(mapCtx *direct.MapContext, in *krm.Channel) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Provider = direct.ValueOf(in.Provider)
	// MISSING: PubsubTopic
	// MISSING: State
	// MISSING: ActivationToken
	out.CryptoKeyName = direct.ValueOf(in.CryptoKeyName)
	// MISSING: SatisfiesPzs
	return out
}
func ChannelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.ChannelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Provider
	out.PubsubTopic = direct.LazyPtr(in.GetPubsubTopic())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ActivationToken = direct.LazyPtr(in.GetActivationToken())
	// MISSING: CryptoKeyName
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	return out
}
func ChannelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelObservedState) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Provider
	if oneof := ChannelObservedState_PubsubTopic_ToProto(mapCtx, in.PubsubTopic); oneof != nil {
		out.Transport = oneof
	}
	out.State = direct.Enum_ToProto[pb.Channel_State](mapCtx, in.State)
	out.ActivationToken = direct.ValueOf(in.ActivationToken)
	// MISSING: CryptoKeyName
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	return out
}
func EventarcChannelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.EventarcChannelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventarcChannelObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Provider
	// MISSING: PubsubTopic
	// MISSING: State
	// MISSING: ActivationToken
	// MISSING: CryptoKeyName
	// MISSING: SatisfiesPzs
	return out
}
func EventarcChannelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventarcChannelObservedState) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Provider
	// MISSING: PubsubTopic
	// MISSING: State
	// MISSING: ActivationToken
	// MISSING: CryptoKeyName
	// MISSING: SatisfiesPzs
	return out
}
func EventarcChannelSpec_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.EventarcChannelSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcChannelSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Provider
	// MISSING: PubsubTopic
	// MISSING: State
	// MISSING: ActivationToken
	// MISSING: CryptoKeyName
	// MISSING: SatisfiesPzs
	return out
}
func EventarcChannelSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcChannelSpec) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Provider
	// MISSING: PubsubTopic
	// MISSING: State
	// MISSING: ActivationToken
	// MISSING: CryptoKeyName
	// MISSING: SatisfiesPzs
	return out
}
