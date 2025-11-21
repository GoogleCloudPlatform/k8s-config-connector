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
// krm.group: pubsublite.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.pubsublite.v1

package pubsublite

import (
	pb "cloud.google.com/go/pubsublite/apiv1/pubsublitepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsublite/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ExportConfig_FromProto(mapCtx *direct.MapContext, in *pb.ExportConfig) *krm.ExportConfig {
	if in == nil {
		return nil
	}
	out := &krm.ExportConfig{}
	out.DesiredState = direct.Enum_FromProto(mapCtx, in.GetDesiredState())
	// MISSING: CurrentState
	out.DeadLetterTopic = direct.LazyPtr(in.GetDeadLetterTopic())
	out.PubsubConfig = ExportConfig_PubSubConfig_FromProto(mapCtx, in.GetPubsubConfig())
	return out
}
func ExportConfig_ToProto(mapCtx *direct.MapContext, in *krm.ExportConfig) *pb.ExportConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExportConfig{}
	out.DesiredState = direct.Enum_ToProto[pb.ExportConfig_State](mapCtx, in.DesiredState)
	// MISSING: CurrentState
	out.DeadLetterTopic = direct.ValueOf(in.DeadLetterTopic)
	if oneof := ExportConfig_PubSubConfig_ToProto(mapCtx, in.PubsubConfig); oneof != nil {
		out.Destination = &pb.ExportConfig_PubsubConfig{PubsubConfig: oneof}
	}
	return out
}
func ExportConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ExportConfig) *krm.ExportConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ExportConfigObservedState{}
	// MISSING: DesiredState
	out.CurrentState = direct.Enum_FromProto(mapCtx, in.GetCurrentState())
	// MISSING: DeadLetterTopic
	// MISSING: PubsubConfig
	return out
}
func ExportConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ExportConfigObservedState) *pb.ExportConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExportConfig{}
	// MISSING: DesiredState
	out.CurrentState = direct.Enum_ToProto[pb.ExportConfig_State](mapCtx, in.CurrentState)
	// MISSING: DeadLetterTopic
	// MISSING: PubsubConfig
	return out
}
func ExportConfig_PubSubConfig_FromProto(mapCtx *direct.MapContext, in *pb.ExportConfig_PubSubConfig) *krm.ExportConfig_PubSubConfig {
	if in == nil {
		return nil
	}
	out := &krm.ExportConfig_PubSubConfig{}
	out.Topic = direct.LazyPtr(in.GetTopic())
	return out
}
func ExportConfig_PubSubConfig_ToProto(mapCtx *direct.MapContext, in *krm.ExportConfig_PubSubConfig) *pb.ExportConfig_PubSubConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExportConfig_PubSubConfig{}
	out.Topic = direct.ValueOf(in.Topic)
	return out
}
func PubSubLiteSubscriptionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Subscription) *krm.PubSubLiteSubscriptionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PubSubLiteSubscriptionObservedState{}
	// MISSING: Name
	// MISSING: ExportConfig
	return out
}
func PubSubLiteSubscriptionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PubSubLiteSubscriptionObservedState) *pb.Subscription {
	if in == nil {
		return nil
	}
	out := &pb.Subscription{}
	// MISSING: Name
	// MISSING: ExportConfig
	return out
}
func PubSubLiteSubscriptionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Subscription) *krm.PubSubLiteSubscriptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.PubSubLiteSubscriptionSpec{}
	// MISSING: Name
	out.Topic = direct.LazyPtr(in.GetTopic())
	out.DeliveryConfig = Subscription_DeliveryConfig_FromProto(mapCtx, in.GetDeliveryConfig())
	// MISSING: ExportConfig
	return out
}
func PubSubLiteSubscriptionSpec_ToProto(mapCtx *direct.MapContext, in *krm.PubSubLiteSubscriptionSpec) *pb.Subscription {
	if in == nil {
		return nil
	}
	out := &pb.Subscription{}
	// MISSING: Name
	out.Topic = direct.ValueOf(in.Topic)
	out.DeliveryConfig = Subscription_DeliveryConfig_ToProto(mapCtx, in.DeliveryConfig)
	// MISSING: ExportConfig
	return out
}
func Subscription_DeliveryConfig_FromProto(mapCtx *direct.MapContext, in *pb.Subscription_DeliveryConfig) *krm.Subscription_DeliveryConfig {
	if in == nil {
		return nil
	}
	out := &krm.Subscription_DeliveryConfig{}
	out.DeliveryRequirement = direct.Enum_FromProto(mapCtx, in.GetDeliveryRequirement())
	return out
}
func Subscription_DeliveryConfig_ToProto(mapCtx *direct.MapContext, in *krm.Subscription_DeliveryConfig) *pb.Subscription_DeliveryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Subscription_DeliveryConfig{}
	out.DeliveryRequirement = direct.Enum_ToProto[pb.Subscription_DeliveryConfig_DeliveryRequirement](mapCtx, in.DeliveryRequirement)
	return out
}
