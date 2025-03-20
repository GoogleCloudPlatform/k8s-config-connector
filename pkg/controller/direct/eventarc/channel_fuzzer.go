// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.eventarc.v1.Channel
// api.group: eventarc.cnrm.cloud.google.com

package eventarc

import (
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(eventarcChannelFuzzer())
}

func eventarcChannelFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Channel{},
		EventarcChannelSpec_FromProto, EventarcChannelSpec_ToProto,
		EventarcChannelObservedState_FromProto, EventarcChannelObservedState_ToProto,
	)
	f.SpecFields.Insert(".provider")
	f.SpecFields.Insert(".crypto_key_name")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".pubsub_topic")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".activation_token")
	f.StatusFields.Insert(".satisfies_pzs")

	f.UnimplementedFields.Insert(".name")

	return f
}

func EventarcChannelSpec_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.EventarcChannelSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcChannelSpec{}
	out.Provider = direct.LazyPtr(in.GetProvider())
	if in.GetCryptoKeyName() != "" {
		out.KmsKeyRef = &refv1beta1.KMSCryptoKeyRef{External: in.GetCryptoKeyName()}
	}
	return out
}
func EventarcChannelSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcChannelSpec) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	out.Provider = direct.ValueOf(in.Provider)
	if in.KmsKeyRef != nil {
		out.CryptoKeyName = in.KmsKeyRef.External
	}
	return out
}
func EventarcChannelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.EventarcChannelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventarcChannelObservedState{}
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.PubsubTopic = direct.LazyPtr(in.GetPubsubTopic())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ActivationToken = direct.LazyPtr(in.GetActivationToken())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	return out
}
func EventarcChannelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventarcChannelObservedState) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.PubsubTopic = direct.ValueOf(in.PubsubTopic)
	out.State = direct.Enum_ToProto[pb.Channel_State](mapCtx, in.State)
	out.ActivationToken = direct.ValueOf(in.ActivationToken)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	return out
}

```
</out>


