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

// +generated:mapper
// krm.group: eventarc.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.eventarc.v1

package eventarc

import (
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func EventarcChannelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.EventarcChannelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventarcChannelObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.PubsubTopic = direct.LazyPtr(in.GetPubsubTopic())

	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ActivationToken = direct.LazyPtr(in.GetActivationToken())
	// MISSING: CryptoKeyName
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	return out
}

func EventarcChannelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventarcChannelObservedState) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	if oneof := EventarcChannelObservedState_PubsubTopic_ToProto(mapCtx, in.PubsubTopic); oneof != nil {
		out.Transport = oneof
	}
	out.State = direct.Enum_ToProto[pb.Channel_State](mapCtx, in.State)
	out.ActivationToken = direct.ValueOf(in.ActivationToken)
	// MISSING: CryptoKeyName
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	return out
}

func EventarcChannelObservedState_PubsubTopic_ToProto(mapCtx *direct.MapContext, in *string) *pb.Channel_PubsubTopic {
	if in == nil {
		return nil
	}
	return &pb.Channel_PubsubTopic{
		PubsubTopic: *in,
	}
}

func EventarcChannelSpec_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.EventarcChannelSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcChannelSpec{}
	// MISSING: Name
	out.Provider = direct.LazyPtr(in.GetProvider())
	// MISSING: CryptoKeyName
	return out
}
func EventarcChannelSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcChannelSpec) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	// MISSING: Name
	out.Provider = direct.ValueOf(in.Provider)
	// MISSING: CryptoKeyName
	return out
}
