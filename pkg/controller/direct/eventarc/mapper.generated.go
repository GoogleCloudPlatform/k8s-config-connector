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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ChannelConnection_FromProto(mapCtx *direct.MapContext, in *pb.ChannelConnection) *krm.ChannelConnection {
	if in == nil {
		return nil
	}
	out := &krm.ChannelConnection{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Uid
	out.Channel = direct.LazyPtr(in.GetChannel())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ActivationToken = direct.LazyPtr(in.GetActivationToken())
	return out
}
func ChannelConnection_ToProto(mapCtx *direct.MapContext, in *krm.ChannelConnection) *pb.ChannelConnection {
	if in == nil {
		return nil
	}
	out := &pb.ChannelConnection{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Uid
	out.Channel = direct.ValueOf(in.Channel)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ActivationToken = direct.ValueOf(in.ActivationToken)
	return out
}
func ChannelConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ChannelConnection) *krm.ChannelConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelConnectionObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Channel
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: ActivationToken
	return out
}
func ChannelConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelConnectionObservedState) *pb.ChannelConnection {
	if in == nil {
		return nil
	}
	out := &pb.ChannelConnection{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Channel
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: ActivationToken
	return out
}
func EventarcChannelConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ChannelConnection) *krm.EventarcChannelConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventarcChannelConnectionObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Channel
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ActivationToken
	return out
}
func EventarcChannelConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventarcChannelConnectionObservedState) *pb.ChannelConnection {
	if in == nil {
		return nil
	}
	out := &pb.ChannelConnection{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Channel
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ActivationToken
	return out
}
func EventarcChannelConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.ChannelConnection) *krm.EventarcChannelConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcChannelConnectionSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Channel
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ActivationToken
	return out
}
func EventarcChannelConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcChannelConnectionSpec) *pb.ChannelConnection {
	if in == nil {
		return nil
	}
	out := &pb.ChannelConnection{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Channel
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ActivationToken
	return out
}
