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

package pubsublite

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/pubsublite/apiv1/pubsublitepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsublite/v1alpha1"
)
func PubsubliteTopicObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Topic) *krm.PubsubliteTopicObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PubsubliteTopicObservedState{}
	// MISSING: Name
	// MISSING: PartitionConfig
	// MISSING: RetentionConfig
	// MISSING: ReservationConfig
	return out
}
func PubsubliteTopicObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PubsubliteTopicObservedState) *pb.Topic {
	if in == nil {
		return nil
	}
	out := &pb.Topic{}
	// MISSING: Name
	// MISSING: PartitionConfig
	// MISSING: RetentionConfig
	// MISSING: ReservationConfig
	return out
}
func PubsubliteTopicSpec_FromProto(mapCtx *direct.MapContext, in *pb.Topic) *krm.PubsubliteTopicSpec {
	if in == nil {
		return nil
	}
	out := &krm.PubsubliteTopicSpec{}
	// MISSING: Name
	// MISSING: PartitionConfig
	// MISSING: RetentionConfig
	// MISSING: ReservationConfig
	return out
}
func PubsubliteTopicSpec_ToProto(mapCtx *direct.MapContext, in *krm.PubsubliteTopicSpec) *pb.Topic {
	if in == nil {
		return nil
	}
	out := &pb.Topic{}
	// MISSING: Name
	// MISSING: PartitionConfig
	// MISSING: RetentionConfig
	// MISSING: ReservationConfig
	return out
}
func Topic_FromProto(mapCtx *direct.MapContext, in *pb.Topic) *krm.Topic {
	if in == nil {
		return nil
	}
	out := &krm.Topic{}
	out.Name = direct.LazyPtr(in.GetName())
	out.PartitionConfig = Topic_PartitionConfig_FromProto(mapCtx, in.GetPartitionConfig())
	out.RetentionConfig = Topic_RetentionConfig_FromProto(mapCtx, in.GetRetentionConfig())
	out.ReservationConfig = Topic_ReservationConfig_FromProto(mapCtx, in.GetReservationConfig())
	return out
}
func Topic_ToProto(mapCtx *direct.MapContext, in *krm.Topic) *pb.Topic {
	if in == nil {
		return nil
	}
	out := &pb.Topic{}
	out.Name = direct.ValueOf(in.Name)
	out.PartitionConfig = Topic_PartitionConfig_ToProto(mapCtx, in.PartitionConfig)
	out.RetentionConfig = Topic_RetentionConfig_ToProto(mapCtx, in.RetentionConfig)
	out.ReservationConfig = Topic_ReservationConfig_ToProto(mapCtx, in.ReservationConfig)
	return out
}
func Topic_PartitionConfig_FromProto(mapCtx *direct.MapContext, in *pb.Topic_PartitionConfig) *krm.Topic_PartitionConfig {
	if in == nil {
		return nil
	}
	out := &krm.Topic_PartitionConfig{}
	out.Count = direct.LazyPtr(in.GetCount())
	out.Scale = direct.LazyPtr(in.GetScale())
	out.Capacity = Topic_PartitionConfig_Capacity_FromProto(mapCtx, in.GetCapacity())
	return out
}
func Topic_PartitionConfig_ToProto(mapCtx *direct.MapContext, in *krm.Topic_PartitionConfig) *pb.Topic_PartitionConfig {
	if in == nil {
		return nil
	}
	out := &pb.Topic_PartitionConfig{}
	out.Count = direct.ValueOf(in.Count)
	if oneof := Topic_PartitionConfig_Scale_ToProto(mapCtx, in.Scale); oneof != nil {
		out.Dimension = oneof
	}
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
	out.PublishMibPerSec = direct.LazyPtr(in.GetPublishMibPerSec())
	out.SubscribeMibPerSec = direct.LazyPtr(in.GetSubscribeMibPerSec())
	return out
}
func Topic_PartitionConfig_Capacity_ToProto(mapCtx *direct.MapContext, in *krm.Topic_PartitionConfig_Capacity) *pb.Topic_PartitionConfig_Capacity {
	if in == nil {
		return nil
	}
	out := &pb.Topic_PartitionConfig_Capacity{}
	out.PublishMibPerSec = direct.ValueOf(in.PublishMibPerSec)
	out.SubscribeMibPerSec = direct.ValueOf(in.SubscribeMibPerSec)
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
