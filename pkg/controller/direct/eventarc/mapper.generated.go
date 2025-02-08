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
func CloudRun_FromProto(mapCtx *direct.MapContext, in *pb.CloudRun) *krm.CloudRun {
	if in == nil {
		return nil
	}
	out := &krm.CloudRun{}
	out.Service = direct.LazyPtr(in.GetService())
	out.Path = direct.LazyPtr(in.GetPath())
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func CloudRun_ToProto(mapCtx *direct.MapContext, in *krm.CloudRun) *pb.CloudRun {
	if in == nil {
		return nil
	}
	out := &pb.CloudRun{}
	out.Service = direct.ValueOf(in.Service)
	out.Path = direct.ValueOf(in.Path)
	out.Region = direct.ValueOf(in.Region)
	return out
}
func Destination_FromProto(mapCtx *direct.MapContext, in *pb.Destination) *krm.Destination {
	if in == nil {
		return nil
	}
	out := &krm.Destination{}
	out.CloudRun = CloudRun_FromProto(mapCtx, in.GetCloudRun())
	out.CloudFunction = direct.LazyPtr(in.GetCloudFunction())
	out.Gke = GKE_FromProto(mapCtx, in.GetGke())
	out.Workflow = direct.LazyPtr(in.GetWorkflow())
	out.HTTPEndpoint = HttpEndpoint_FromProto(mapCtx, in.GetHttpEndpoint())
	out.NetworkConfig = NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	return out
}
func Destination_ToProto(mapCtx *direct.MapContext, in *krm.Destination) *pb.Destination {
	if in == nil {
		return nil
	}
	out := &pb.Destination{}
	if oneof := CloudRun_ToProto(mapCtx, in.CloudRun); oneof != nil {
		out.Descriptor = &pb.Destination_CloudRun{CloudRun: oneof}
	}
	if oneof := Destination_CloudFunction_ToProto(mapCtx, in.CloudFunction); oneof != nil {
		out.Descriptor = oneof
	}
	if oneof := GKE_ToProto(mapCtx, in.Gke); oneof != nil {
		out.Descriptor = &pb.Destination_Gke{Gke: oneof}
	}
	if oneof := Destination_Workflow_ToProto(mapCtx, in.Workflow); oneof != nil {
		out.Descriptor = oneof
	}
	if oneof := HttpEndpoint_ToProto(mapCtx, in.HTTPEndpoint); oneof != nil {
		out.Descriptor = &pb.Destination_HttpEndpoint{HttpEndpoint: oneof}
	}
	out.NetworkConfig = NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	return out
}
func EventFilter_FromProto(mapCtx *direct.MapContext, in *pb.EventFilter) *krm.EventFilter {
	if in == nil {
		return nil
	}
	out := &krm.EventFilter{}
	out.Attribute = direct.LazyPtr(in.GetAttribute())
	out.Value = direct.LazyPtr(in.GetValue())
	out.Operator = direct.LazyPtr(in.GetOperator())
	return out
}
func EventFilter_ToProto(mapCtx *direct.MapContext, in *krm.EventFilter) *pb.EventFilter {
	if in == nil {
		return nil
	}
	out := &pb.EventFilter{}
	out.Attribute = direct.ValueOf(in.Attribute)
	out.Value = direct.ValueOf(in.Value)
	out.Operator = direct.ValueOf(in.Operator)
	return out
}
func EventarcTriggerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Trigger) *krm.EventarcTriggerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventarcTriggerObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EventFilters
	// MISSING: ServiceAccount
	// MISSING: Destination
	// MISSING: Transport
	// MISSING: Labels
	// MISSING: Channel
	// MISSING: Conditions
	// MISSING: EventDataContentType
	// MISSING: SatisfiesPzs
	// MISSING: Etag
	return out
}
func EventarcTriggerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventarcTriggerObservedState) *pb.Trigger {
	if in == nil {
		return nil
	}
	out := &pb.Trigger{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EventFilters
	// MISSING: ServiceAccount
	// MISSING: Destination
	// MISSING: Transport
	// MISSING: Labels
	// MISSING: Channel
	// MISSING: Conditions
	// MISSING: EventDataContentType
	// MISSING: SatisfiesPzs
	// MISSING: Etag
	return out
}
func EventarcTriggerSpec_FromProto(mapCtx *direct.MapContext, in *pb.Trigger) *krm.EventarcTriggerSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcTriggerSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EventFilters
	// MISSING: ServiceAccount
	// MISSING: Destination
	// MISSING: Transport
	// MISSING: Labels
	// MISSING: Channel
	// MISSING: Conditions
	// MISSING: EventDataContentType
	// MISSING: SatisfiesPzs
	// MISSING: Etag
	return out
}
func EventarcTriggerSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcTriggerSpec) *pb.Trigger {
	if in == nil {
		return nil
	}
	out := &pb.Trigger{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EventFilters
	// MISSING: ServiceAccount
	// MISSING: Destination
	// MISSING: Transport
	// MISSING: Labels
	// MISSING: Channel
	// MISSING: Conditions
	// MISSING: EventDataContentType
	// MISSING: SatisfiesPzs
	// MISSING: Etag
	return out
}
func GKE_FromProto(mapCtx *direct.MapContext, in *pb.GKE) *krm.GKE {
	if in == nil {
		return nil
	}
	out := &krm.GKE{}
	out.Cluster = direct.LazyPtr(in.GetCluster())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Namespace = direct.LazyPtr(in.GetNamespace())
	out.Service = direct.LazyPtr(in.GetService())
	out.Path = direct.LazyPtr(in.GetPath())
	return out
}
func GKE_ToProto(mapCtx *direct.MapContext, in *krm.GKE) *pb.GKE {
	if in == nil {
		return nil
	}
	out := &pb.GKE{}
	out.Cluster = direct.ValueOf(in.Cluster)
	out.Location = direct.ValueOf(in.Location)
	out.Namespace = direct.ValueOf(in.Namespace)
	out.Service = direct.ValueOf(in.Service)
	out.Path = direct.ValueOf(in.Path)
	return out
}
func HttpEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.HttpEndpoint) *krm.HttpEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.HttpEndpoint{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func HttpEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.HttpEndpoint) *pb.HttpEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.HttpEndpoint{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig{}
	out.NetworkAttachment = direct.LazyPtr(in.GetNetworkAttachment())
	return out
}
func NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	out.NetworkAttachment = direct.ValueOf(in.NetworkAttachment)
	return out
}
func Pubsub_FromProto(mapCtx *direct.MapContext, in *pb.Pubsub) *krm.Pubsub {
	if in == nil {
		return nil
	}
	out := &krm.Pubsub{}
	out.Topic = direct.LazyPtr(in.GetTopic())
	// MISSING: Subscription
	return out
}
func Pubsub_ToProto(mapCtx *direct.MapContext, in *krm.Pubsub) *pb.Pubsub {
	if in == nil {
		return nil
	}
	out := &pb.Pubsub{}
	out.Topic = direct.ValueOf(in.Topic)
	// MISSING: Subscription
	return out
}
func PubsubObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Pubsub) *krm.PubsubObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PubsubObservedState{}
	// MISSING: Topic
	out.Subscription = direct.LazyPtr(in.GetSubscription())
	return out
}
func PubsubObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PubsubObservedState) *pb.Pubsub {
	if in == nil {
		return nil
	}
	out := &pb.Pubsub{}
	// MISSING: Topic
	out.Subscription = direct.ValueOf(in.Subscription)
	return out
}
func StateCondition_FromProto(mapCtx *direct.MapContext, in *pb.StateCondition) *krm.StateCondition {
	if in == nil {
		return nil
	}
	out := &krm.StateCondition{}
	out.Code = direct.Enum_FromProto(mapCtx, in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
func StateCondition_ToProto(mapCtx *direct.MapContext, in *krm.StateCondition) *pb.StateCondition {
	if in == nil {
		return nil
	}
	out := &pb.StateCondition{}
	out.Code = direct.Enum_ToProto[pb.Code](mapCtx, in.Code)
	out.Message = direct.ValueOf(in.Message)
	return out
}
func Transport_FromProto(mapCtx *direct.MapContext, in *pb.Transport) *krm.Transport {
	if in == nil {
		return nil
	}
	out := &krm.Transport{}
	out.Pubsub = Pubsub_FromProto(mapCtx, in.GetPubsub())
	return out
}
func Transport_ToProto(mapCtx *direct.MapContext, in *krm.Transport) *pb.Transport {
	if in == nil {
		return nil
	}
	out := &pb.Transport{}
	if oneof := Pubsub_ToProto(mapCtx, in.Pubsub); oneof != nil {
		out.Intermediary = &pb.Transport_Pubsub{Pubsub: oneof}
	}
	return out
}
func TransportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Transport) *krm.TransportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TransportObservedState{}
	out.Pubsub = PubsubObservedState_FromProto(mapCtx, in.GetPubsub())
	return out
}
func TransportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TransportObservedState) *pb.Transport {
	if in == nil {
		return nil
	}
	out := &pb.Transport{}
	if oneof := PubsubObservedState_ToProto(mapCtx, in.Pubsub); oneof != nil {
		out.Intermediary = &pb.Transport_Pubsub{Pubsub: oneof}
	}
	return out
}
func Trigger_FromProto(mapCtx *direct.MapContext, in *pb.Trigger) *krm.Trigger {
	if in == nil {
		return nil
	}
	out := &krm.Trigger{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.EventFilters = direct.Slice_FromProto(mapCtx, in.EventFilters, EventFilter_FromProto)
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.Destination = Destination_FromProto(mapCtx, in.GetDestination())
	out.Transport = Transport_FromProto(mapCtx, in.GetTransport())
	out.Labels = in.Labels
	out.Channel = direct.LazyPtr(in.GetChannel())
	// MISSING: Conditions
	out.EventDataContentType = direct.LazyPtr(in.GetEventDataContentType())
	// MISSING: SatisfiesPzs
	// MISSING: Etag
	return out
}
func Trigger_ToProto(mapCtx *direct.MapContext, in *krm.Trigger) *pb.Trigger {
	if in == nil {
		return nil
	}
	out := &pb.Trigger{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.EventFilters = direct.Slice_ToProto(mapCtx, in.EventFilters, EventFilter_ToProto)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.Destination = Destination_ToProto(mapCtx, in.Destination)
	out.Transport = Transport_ToProto(mapCtx, in.Transport)
	out.Labels = in.Labels
	out.Channel = direct.ValueOf(in.Channel)
	// MISSING: Conditions
	out.EventDataContentType = direct.ValueOf(in.EventDataContentType)
	// MISSING: SatisfiesPzs
	// MISSING: Etag
	return out
}
func TriggerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Trigger) *krm.TriggerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TriggerObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: EventFilters
	// MISSING: ServiceAccount
	// MISSING: Destination
	out.Transport = TransportObservedState_FromProto(mapCtx, in.GetTransport())
	// MISSING: Labels
	// MISSING: Channel
	// MISSING: Conditions
	// MISSING: EventDataContentType
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func TriggerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TriggerObservedState) *pb.Trigger {
	if in == nil {
		return nil
	}
	out := &pb.Trigger{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: EventFilters
	// MISSING: ServiceAccount
	// MISSING: Destination
	out.Transport = TransportObservedState_ToProto(mapCtx, in.Transport)
	// MISSING: Labels
	// MISSING: Channel
	// MISSING: Conditions
	// MISSING: EventDataContentType
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
