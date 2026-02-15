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
func PubSubLiteTopicObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Topic) *krm.PubSubLiteTopicObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PubSubLiteTopicObservedState{}
	// MISSING: Name
	return out
}
func PubSubLiteTopicObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PubSubLiteTopicObservedState) *pb.Topic {
	if in == nil {
		return nil
	}
	out := &pb.Topic{}
	// MISSING: Name
	return out
}
func PubSubLiteTopicSpec_FromProto(mapCtx *direct.MapContext, in *pb.Topic) *krm.PubSubLiteTopicSpec {
	if in == nil {
		return nil
	}
	out := &krm.PubSubLiteTopicSpec{}
	// MISSING: Name
	out.PartitionConfig = Topic_PartitionConfig_FromProto(mapCtx, in.GetPartitionConfig())
	out.RetentionConfig = Topic_RetentionConfig_FromProto(mapCtx, in.GetRetentionConfig())
	out.ReservationConfig = Topic_ReservationConfig_FromProto(mapCtx, in.GetReservationConfig())
	return out
}
func PubSubLiteTopicSpec_ToProto(mapCtx *direct.MapContext, in *krm.PubSubLiteTopicSpec) *pb.Topic {
	if in == nil {
		return nil
	}
	out := &pb.Topic{}
	// MISSING: Name
	out.PartitionConfig = Topic_PartitionConfig_ToProto(mapCtx, in.PartitionConfig)
	out.RetentionConfig = Topic_RetentionConfig_ToProto(mapCtx, in.RetentionConfig)
	out.ReservationConfig = Topic_ReservationConfig_ToProto(mapCtx, in.ReservationConfig)
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
func Topic_PartitionConfig_FromProto(mapCtx *direct.MapContext, in *pb.Topic_PartitionConfig) *krm.Topic_PartitionConfig {
	if in == nil {
		return nil
	}
	out := &krm.Topic_PartitionConfig{}
	out.Count = direct.LazyPtr(in.GetCount())
	// MISSING: Scale
	out.Capacity = Topic_PartitionConfig_Capacity_FromProto(mapCtx, in.GetCapacity())
	return out
}
func Topic_PartitionConfig_ToProto(mapCtx *direct.MapContext, in *krm.Topic_PartitionConfig) *pb.Topic_PartitionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Topic_PartitionConfig{}
	out.Count = direct.ValueOf(in.Count)
	// MISSING: Scale
	if oneof := Topic_PartitionConfig_Capacity_ToProto(mapCtx, in.Capacity); oneof != nil {
		out.Dimension = &pb.Topic_PartitionConfig_Capacity_{Capacity: oneof}
	}
	return out
}
func Topic_PartitionConfig_Capacity_FromProto(mapCtx *direct.MapContext, in *pb.Topic_PartitionConfig_Capacity) *krm.Topic_PartitionConfig_Capacity {
	if in == nil {
		return nil
	}
	out := &krm.Topic_PartitionConfig_Capacity{}
	out.PublishMIBPerSec = direct.LazyPtr(in.GetPublishMibPerSec())
	out.SubscribeMIBPerSec = direct.LazyPtr(in.GetSubscribeMibPerSec())
	return out
}
func Topic_PartitionConfig_Capacity_ToProto(mapCtx *direct.MapContext, in *krm.Topic_PartitionConfig_Capacity) *pb.Topic_PartitionConfig_Capacity {
	if in == nil {
		return nil
	}
	out := &pb.Topic_PartitionConfig_Capacity{}
	out.PublishMibPerSec = direct.ValueOf(in.PublishMIBPerSec)
	out.SubscribeMibPerSec = direct.ValueOf(in.SubscribeMIBPerSec)
	return out
}
func Topic_ReservationConfig_FromProto(mapCtx *direct.MapContext, in *pb.Topic_ReservationConfig) *krm.Topic_ReservationConfig {
	if in == nil {
		return nil
	}
	out := &krm.Topic_ReservationConfig{}
	out.ThroughputReservation = direct.LazyPtr(in.GetThroughputReservation())
	return out
}
func Topic_ReservationConfig_ToProto(mapCtx *direct.MapContext, in *krm.Topic_ReservationConfig) *pb.Topic_ReservationConfig {
	if in == nil {
		return nil
	}
	out := &pb.Topic_ReservationConfig{}
	out.ThroughputReservation = direct.ValueOf(in.ThroughputReservation)
	return out
}
func Topic_RetentionConfig_FromProto(mapCtx *direct.MapContext, in *pb.Topic_RetentionConfig) *krm.Topic_RetentionConfig {
	if in == nil {
		return nil
	}
	out := &krm.Topic_RetentionConfig{}
	out.PerPartitionBytes = direct.LazyPtr(in.GetPerPartitionBytes())
	out.Period = direct.StringDuration_FromProto(mapCtx, in.GetPeriod())
	return out
}
func Topic_RetentionConfig_ToProto(mapCtx *direct.MapContext, in *krm.Topic_RetentionConfig) *pb.Topic_RetentionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Topic_RetentionConfig{}
	out.PerPartitionBytes = direct.ValueOf(in.PerPartitionBytes)
	out.Period = direct.StringDuration_ToProto(mapCtx, in.Period)
	return out
}
