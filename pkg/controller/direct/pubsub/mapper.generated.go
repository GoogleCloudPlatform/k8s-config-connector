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

package pubsub

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func IngestionDataSourceSettings_FromProto(mapCtx *direct.MapContext, in *pb.IngestionDataSourceSettings) *krm.IngestionDataSourceSettings {
	if in == nil {
		return nil
	}
	out := &krm.IngestionDataSourceSettings{}
	out.AwsKinesis = IngestionDataSourceSettings_AwsKinesis_FromProto(mapCtx, in.GetAwsKinesis())
	out.CloudStorage = IngestionDataSourceSettings_CloudStorage_FromProto(mapCtx, in.GetCloudStorage())
	out.AzureEventHubs = IngestionDataSourceSettings_AzureEventHubs_FromProto(mapCtx, in.GetAzureEventHubs())
	out.AwsMsk = IngestionDataSourceSettings_AwsMsk_FromProto(mapCtx, in.GetAwsMsk())
	out.ConfluentCloud = IngestionDataSourceSettings_ConfluentCloud_FromProto(mapCtx, in.GetConfluentCloud())
	out.PlatformLogsSettings = PlatformLogsSettings_FromProto(mapCtx, in.GetPlatformLogsSettings())
	return out
}
func IngestionDataSourceSettings_ToProto(mapCtx *direct.MapContext, in *krm.IngestionDataSourceSettings) *pb.IngestionDataSourceSettings {
	if in == nil {
		return nil
	}
	out := &pb.IngestionDataSourceSettings{}
	if oneof := IngestionDataSourceSettings_AwsKinesis_ToProto(mapCtx, in.AwsKinesis); oneof != nil {
		out.Source = &pb.IngestionDataSourceSettings_AwsKinesis_{AwsKinesis: oneof}
	}
	if oneof := IngestionDataSourceSettings_CloudStorage_ToProto(mapCtx, in.CloudStorage); oneof != nil {
		out.Source = &pb.IngestionDataSourceSettings_CloudStorage_{CloudStorage: oneof}
	}
	if oneof := IngestionDataSourceSettings_AzureEventHubs_ToProto(mapCtx, in.AzureEventHubs); oneof != nil {
		out.Source = &pb.IngestionDataSourceSettings_AzureEventHubs_{AzureEventHubs: oneof}
	}
	if oneof := IngestionDataSourceSettings_AwsMsk_ToProto(mapCtx, in.AwsMsk); oneof != nil {
		out.Source = &pb.IngestionDataSourceSettings_AwsMsk_{AwsMsk: oneof}
	}
	if oneof := IngestionDataSourceSettings_ConfluentCloud_ToProto(mapCtx, in.ConfluentCloud); oneof != nil {
		out.Source = &pb.IngestionDataSourceSettings_ConfluentCloud_{ConfluentCloud: oneof}
	}
	out.PlatformLogsSettings = PlatformLogsSettings_ToProto(mapCtx, in.PlatformLogsSettings)
	return out
}
func IngestionDataSourceSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IngestionDataSourceSettings) *krm.IngestionDataSourceSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IngestionDataSourceSettingsObservedState{}
	out.AwsKinesis = IngestionDataSourceSettings_AwsKinesisObservedState_FromProto(mapCtx, in.GetAwsKinesis())
	out.CloudStorage = IngestionDataSourceSettings_CloudStorageObservedState_FromProto(mapCtx, in.GetCloudStorage())
	out.AzureEventHubs = IngestionDataSourceSettings_AzureEventHubsObservedState_FromProto(mapCtx, in.GetAzureEventHubs())
	out.AwsMsk = IngestionDataSourceSettings_AwsMskObservedState_FromProto(mapCtx, in.GetAwsMsk())
	out.ConfluentCloud = IngestionDataSourceSettings_ConfluentCloudObservedState_FromProto(mapCtx, in.GetConfluentCloud())
	// MISSING: PlatformLogsSettings
	return out
}
func IngestionDataSourceSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IngestionDataSourceSettingsObservedState) *pb.IngestionDataSourceSettings {
	if in == nil {
		return nil
	}
	out := &pb.IngestionDataSourceSettings{}
	if oneof := IngestionDataSourceSettings_AwsKinesisObservedState_ToProto(mapCtx, in.AwsKinesis); oneof != nil {
		out.Source = &pb.IngestionDataSourceSettings_AwsKinesis_{AwsKinesis: oneof}
	}
	if oneof := IngestionDataSourceSettings_CloudStorageObservedState_ToProto(mapCtx, in.CloudStorage); oneof != nil {
		out.Source = &pb.IngestionDataSourceSettings_CloudStorage_{CloudStorage: oneof}
	}
	if oneof := IngestionDataSourceSettings_AzureEventHubsObservedState_ToProto(mapCtx, in.AzureEventHubs); oneof != nil {
		out.Source = &pb.IngestionDataSourceSettings_AzureEventHubs_{AzureEventHubs: oneof}
	}
	if oneof := IngestionDataSourceSettings_AwsMskObservedState_ToProto(mapCtx, in.AwsMsk); oneof != nil {
		out.Source = &pb.IngestionDataSourceSettings_AwsMsk_{AwsMsk: oneof}
	}
	if oneof := IngestionDataSourceSettings_ConfluentCloudObservedState_ToProto(mapCtx, in.ConfluentCloud); oneof != nil {
		out.Source = &pb.IngestionDataSourceSettings_ConfluentCloud_{ConfluentCloud: oneof}
	}
	// MISSING: PlatformLogsSettings
	return out
}
func IngestionDataSourceSettings_AwsKinesis_FromProto(mapCtx *direct.MapContext, in *pb.IngestionDataSourceSettings_AwsKinesis) *krm.IngestionDataSourceSettings_AwsKinesis {
	if in == nil {
		return nil
	}
	out := &krm.IngestionDataSourceSettings_AwsKinesis{}
	// MISSING: State
	out.StreamArn = direct.LazyPtr(in.GetStreamArn())
	out.ConsumerArn = direct.LazyPtr(in.GetConsumerArn())
	out.AwsRoleArn = direct.LazyPtr(in.GetAwsRoleArn())
	out.GcpServiceAccount = direct.LazyPtr(in.GetGcpServiceAccount())
	return out
}
func IngestionDataSourceSettings_AwsKinesis_ToProto(mapCtx *direct.MapContext, in *krm.IngestionDataSourceSettings_AwsKinesis) *pb.IngestionDataSourceSettings_AwsKinesis {
	if in == nil {
		return nil
	}
	out := &pb.IngestionDataSourceSettings_AwsKinesis{}
	// MISSING: State
	out.StreamArn = direct.ValueOf(in.StreamArn)
	out.ConsumerArn = direct.ValueOf(in.ConsumerArn)
	out.AwsRoleArn = direct.ValueOf(in.AwsRoleArn)
	out.GcpServiceAccount = direct.ValueOf(in.GcpServiceAccount)
	return out
}
func IngestionDataSourceSettings_AwsKinesisObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IngestionDataSourceSettings_AwsKinesis) *krm.IngestionDataSourceSettings_AwsKinesisObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IngestionDataSourceSettings_AwsKinesisObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: StreamArn
	// MISSING: ConsumerArn
	// MISSING: AwsRoleArn
	// MISSING: GcpServiceAccount
	return out
}
func IngestionDataSourceSettings_AwsKinesisObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IngestionDataSourceSettings_AwsKinesisObservedState) *pb.IngestionDataSourceSettings_AwsKinesis {
	if in == nil {
		return nil
	}
	out := &pb.IngestionDataSourceSettings_AwsKinesis{}
	out.State = direct.Enum_ToProto[pb.IngestionDataSourceSettings_AwsKinesis_State](mapCtx, in.State)
	// MISSING: StreamArn
	// MISSING: ConsumerArn
	// MISSING: AwsRoleArn
	// MISSING: GcpServiceAccount
	return out
}
func IngestionDataSourceSettings_AwsMsk_FromProto(mapCtx *direct.MapContext, in *pb.IngestionDataSourceSettings_AwsMsk) *krm.IngestionDataSourceSettings_AwsMsk {
	if in == nil {
		return nil
	}
	out := &krm.IngestionDataSourceSettings_AwsMsk{}
	// MISSING: State
	out.ClusterArn = direct.LazyPtr(in.GetClusterArn())
	out.Topic = direct.LazyPtr(in.GetTopic())
	out.AwsRoleArn = direct.LazyPtr(in.GetAwsRoleArn())
	out.GcpServiceAccount = direct.LazyPtr(in.GetGcpServiceAccount())
	return out
}
func IngestionDataSourceSettings_AwsMsk_ToProto(mapCtx *direct.MapContext, in *krm.IngestionDataSourceSettings_AwsMsk) *pb.IngestionDataSourceSettings_AwsMsk {
	if in == nil {
		return nil
	}
	out := &pb.IngestionDataSourceSettings_AwsMsk{}
	// MISSING: State
	out.ClusterArn = direct.ValueOf(in.ClusterArn)
	out.Topic = direct.ValueOf(in.Topic)
	out.AwsRoleArn = direct.ValueOf(in.AwsRoleArn)
	out.GcpServiceAccount = direct.ValueOf(in.GcpServiceAccount)
	return out
}
func IngestionDataSourceSettings_AwsMskObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IngestionDataSourceSettings_AwsMsk) *krm.IngestionDataSourceSettings_AwsMskObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IngestionDataSourceSettings_AwsMskObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: ClusterArn
	// MISSING: Topic
	// MISSING: AwsRoleArn
	// MISSING: GcpServiceAccount
	return out
}
func IngestionDataSourceSettings_AwsMskObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IngestionDataSourceSettings_AwsMskObservedState) *pb.IngestionDataSourceSettings_AwsMsk {
	if in == nil {
		return nil
	}
	out := &pb.IngestionDataSourceSettings_AwsMsk{}
	out.State = direct.Enum_ToProto[pb.IngestionDataSourceSettings_AwsMsk_State](mapCtx, in.State)
	// MISSING: ClusterArn
	// MISSING: Topic
	// MISSING: AwsRoleArn
	// MISSING: GcpServiceAccount
	return out
}
func IngestionDataSourceSettings_AzureEventHubs_FromProto(mapCtx *direct.MapContext, in *pb.IngestionDataSourceSettings_AzureEventHubs) *krm.IngestionDataSourceSettings_AzureEventHubs {
	if in == nil {
		return nil
	}
	out := &krm.IngestionDataSourceSettings_AzureEventHubs{}
	// MISSING: State
	out.ResourceGroup = direct.LazyPtr(in.GetResourceGroup())
	out.Namespace = direct.LazyPtr(in.GetNamespace())
	out.EventHub = direct.LazyPtr(in.GetEventHub())
	out.ClientID = direct.LazyPtr(in.GetClientId())
	out.TenantID = direct.LazyPtr(in.GetTenantId())
	out.SubscriptionID = direct.LazyPtr(in.GetSubscriptionId())
	out.GcpServiceAccount = direct.LazyPtr(in.GetGcpServiceAccount())
	return out
}
func IngestionDataSourceSettings_AzureEventHubs_ToProto(mapCtx *direct.MapContext, in *krm.IngestionDataSourceSettings_AzureEventHubs) *pb.IngestionDataSourceSettings_AzureEventHubs {
	if in == nil {
		return nil
	}
	out := &pb.IngestionDataSourceSettings_AzureEventHubs{}
	// MISSING: State
	out.ResourceGroup = direct.ValueOf(in.ResourceGroup)
	out.Namespace = direct.ValueOf(in.Namespace)
	out.EventHub = direct.ValueOf(in.EventHub)
	out.ClientId = direct.ValueOf(in.ClientID)
	out.TenantId = direct.ValueOf(in.TenantID)
	out.SubscriptionId = direct.ValueOf(in.SubscriptionID)
	out.GcpServiceAccount = direct.ValueOf(in.GcpServiceAccount)
	return out
}
func IngestionDataSourceSettings_AzureEventHubsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IngestionDataSourceSettings_AzureEventHubs) *krm.IngestionDataSourceSettings_AzureEventHubsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IngestionDataSourceSettings_AzureEventHubsObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: ResourceGroup
	// MISSING: Namespace
	// MISSING: EventHub
	// MISSING: ClientID
	// MISSING: TenantID
	// MISSING: SubscriptionID
	// MISSING: GcpServiceAccount
	return out
}
func IngestionDataSourceSettings_AzureEventHubsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IngestionDataSourceSettings_AzureEventHubsObservedState) *pb.IngestionDataSourceSettings_AzureEventHubs {
	if in == nil {
		return nil
	}
	out := &pb.IngestionDataSourceSettings_AzureEventHubs{}
	out.State = direct.Enum_ToProto[pb.IngestionDataSourceSettings_AzureEventHubs_State](mapCtx, in.State)
	// MISSING: ResourceGroup
	// MISSING: Namespace
	// MISSING: EventHub
	// MISSING: ClientID
	// MISSING: TenantID
	// MISSING: SubscriptionID
	// MISSING: GcpServiceAccount
	return out
}
func IngestionDataSourceSettings_CloudStorage_FromProto(mapCtx *direct.MapContext, in *pb.IngestionDataSourceSettings_CloudStorage) *krm.IngestionDataSourceSettings_CloudStorage {
	if in == nil {
		return nil
	}
	out := &krm.IngestionDataSourceSettings_CloudStorage{}
	// MISSING: State
	out.Bucket = direct.LazyPtr(in.GetBucket())
	out.TextFormat = IngestionDataSourceSettings_CloudStorage_TextFormat_FromProto(mapCtx, in.GetTextFormat())
	out.AvroFormat = IngestionDataSourceSettings_CloudStorage_AvroFormat_FromProto(mapCtx, in.GetAvroFormat())
	out.PubsubAvroFormat = IngestionDataSourceSettings_CloudStorage_PubSubAvroFormat_FromProto(mapCtx, in.GetPubsubAvroFormat())
	out.MinimumObjectCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetMinimumObjectCreateTime())
	out.MatchGlob = direct.LazyPtr(in.GetMatchGlob())
	return out
}
func IngestionDataSourceSettings_CloudStorage_ToProto(mapCtx *direct.MapContext, in *krm.IngestionDataSourceSettings_CloudStorage) *pb.IngestionDataSourceSettings_CloudStorage {
	if in == nil {
		return nil
	}
	out := &pb.IngestionDataSourceSettings_CloudStorage{}
	// MISSING: State
	out.Bucket = direct.ValueOf(in.Bucket)
	if oneof := IngestionDataSourceSettings_CloudStorage_TextFormat_ToProto(mapCtx, in.TextFormat); oneof != nil {
		out.InputFormat = &pb.IngestionDataSourceSettings_CloudStorage_TextFormat_{TextFormat: oneof}
	}
	if oneof := IngestionDataSourceSettings_CloudStorage_AvroFormat_ToProto(mapCtx, in.AvroFormat); oneof != nil {
		out.InputFormat = &pb.IngestionDataSourceSettings_CloudStorage_AvroFormat_{AvroFormat: oneof}
	}
	if oneof := IngestionDataSourceSettings_CloudStorage_PubSubAvroFormat_ToProto(mapCtx, in.PubsubAvroFormat); oneof != nil {
		out.InputFormat = &pb.IngestionDataSourceSettings_CloudStorage_PubsubAvroFormat{PubsubAvroFormat: oneof}
	}
	out.MinimumObjectCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.MinimumObjectCreateTime)
	out.MatchGlob = direct.ValueOf(in.MatchGlob)
	return out
}
func IngestionDataSourceSettings_CloudStorageObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IngestionDataSourceSettings_CloudStorage) *krm.IngestionDataSourceSettings_CloudStorageObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IngestionDataSourceSettings_CloudStorageObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Bucket
	// MISSING: TextFormat
	// MISSING: AvroFormat
	// MISSING: PubsubAvroFormat
	// MISSING: MinimumObjectCreateTime
	// MISSING: MatchGlob
	return out
}
func IngestionDataSourceSettings_CloudStorageObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IngestionDataSourceSettings_CloudStorageObservedState) *pb.IngestionDataSourceSettings_CloudStorage {
	if in == nil {
		return nil
	}
	out := &pb.IngestionDataSourceSettings_CloudStorage{}
	out.State = direct.Enum_ToProto[pb.IngestionDataSourceSettings_CloudStorage_State](mapCtx, in.State)
	// MISSING: Bucket
	// MISSING: TextFormat
	// MISSING: AvroFormat
	// MISSING: PubsubAvroFormat
	// MISSING: MinimumObjectCreateTime
	// MISSING: MatchGlob
	return out
}
func IngestionDataSourceSettings_CloudStorage_AvroFormat_FromProto(mapCtx *direct.MapContext, in *pb.IngestionDataSourceSettings_CloudStorage_AvroFormat) *krm.IngestionDataSourceSettings_CloudStorage_AvroFormat {
	if in == nil {
		return nil
	}
	out := &krm.IngestionDataSourceSettings_CloudStorage_AvroFormat{}
	return out
}
func IngestionDataSourceSettings_CloudStorage_AvroFormat_ToProto(mapCtx *direct.MapContext, in *krm.IngestionDataSourceSettings_CloudStorage_AvroFormat) *pb.IngestionDataSourceSettings_CloudStorage_AvroFormat {
	if in == nil {
		return nil
	}
	out := &pb.IngestionDataSourceSettings_CloudStorage_AvroFormat{}
	return out
}
func IngestionDataSourceSettings_CloudStorage_PubSubAvroFormat_FromProto(mapCtx *direct.MapContext, in *pb.IngestionDataSourceSettings_CloudStorage_PubSubAvroFormat) *krm.IngestionDataSourceSettings_CloudStorage_PubSubAvroFormat {
	if in == nil {
		return nil
	}
	out := &krm.IngestionDataSourceSettings_CloudStorage_PubSubAvroFormat{}
	return out
}
func IngestionDataSourceSettings_CloudStorage_PubSubAvroFormat_ToProto(mapCtx *direct.MapContext, in *krm.IngestionDataSourceSettings_CloudStorage_PubSubAvroFormat) *pb.IngestionDataSourceSettings_CloudStorage_PubSubAvroFormat {
	if in == nil {
		return nil
	}
	out := &pb.IngestionDataSourceSettings_CloudStorage_PubSubAvroFormat{}
	return out
}
func IngestionDataSourceSettings_CloudStorage_TextFormat_FromProto(mapCtx *direct.MapContext, in *pb.IngestionDataSourceSettings_CloudStorage_TextFormat) *krm.IngestionDataSourceSettings_CloudStorage_TextFormat {
	if in == nil {
		return nil
	}
	out := &krm.IngestionDataSourceSettings_CloudStorage_TextFormat{}
	out.Delimiter = in.Delimiter
	return out
}
func IngestionDataSourceSettings_CloudStorage_TextFormat_ToProto(mapCtx *direct.MapContext, in *krm.IngestionDataSourceSettings_CloudStorage_TextFormat) *pb.IngestionDataSourceSettings_CloudStorage_TextFormat {
	if in == nil {
		return nil
	}
	out := &pb.IngestionDataSourceSettings_CloudStorage_TextFormat{}
	out.Delimiter = in.Delimiter
	return out
}
func IngestionDataSourceSettings_ConfluentCloud_FromProto(mapCtx *direct.MapContext, in *pb.IngestionDataSourceSettings_ConfluentCloud) *krm.IngestionDataSourceSettings_ConfluentCloud {
	if in == nil {
		return nil
	}
	out := &krm.IngestionDataSourceSettings_ConfluentCloud{}
	// MISSING: State
	out.BootstrapServer = direct.LazyPtr(in.GetBootstrapServer())
	out.ClusterID = direct.LazyPtr(in.GetClusterId())
	out.Topic = direct.LazyPtr(in.GetTopic())
	out.IdentityPoolID = direct.LazyPtr(in.GetIdentityPoolId())
	out.GcpServiceAccount = direct.LazyPtr(in.GetGcpServiceAccount())
	return out
}
func IngestionDataSourceSettings_ConfluentCloud_ToProto(mapCtx *direct.MapContext, in *krm.IngestionDataSourceSettings_ConfluentCloud) *pb.IngestionDataSourceSettings_ConfluentCloud {
	if in == nil {
		return nil
	}
	out := &pb.IngestionDataSourceSettings_ConfluentCloud{}
	// MISSING: State
	out.BootstrapServer = direct.ValueOf(in.BootstrapServer)
	out.ClusterId = direct.ValueOf(in.ClusterID)
	out.Topic = direct.ValueOf(in.Topic)
	out.IdentityPoolId = direct.ValueOf(in.IdentityPoolID)
	out.GcpServiceAccount = direct.ValueOf(in.GcpServiceAccount)
	return out
}
func IngestionDataSourceSettings_ConfluentCloudObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IngestionDataSourceSettings_ConfluentCloud) *krm.IngestionDataSourceSettings_ConfluentCloudObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IngestionDataSourceSettings_ConfluentCloudObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: BootstrapServer
	// MISSING: ClusterID
	// MISSING: Topic
	// MISSING: IdentityPoolID
	// MISSING: GcpServiceAccount
	return out
}
func IngestionDataSourceSettings_ConfluentCloudObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IngestionDataSourceSettings_ConfluentCloudObservedState) *pb.IngestionDataSourceSettings_ConfluentCloud {
	if in == nil {
		return nil
	}
	out := &pb.IngestionDataSourceSettings_ConfluentCloud{}
	out.State = direct.Enum_ToProto[pb.IngestionDataSourceSettings_ConfluentCloud_State](mapCtx, in.State)
	// MISSING: BootstrapServer
	// MISSING: ClusterID
	// MISSING: Topic
	// MISSING: IdentityPoolID
	// MISSING: GcpServiceAccount
	return out
}
func MessageStoragePolicy_FromProto(mapCtx *direct.MapContext, in *pb.MessageStoragePolicy) *krm.MessageStoragePolicy {
	if in == nil {
		return nil
	}
	out := &krm.MessageStoragePolicy{}
	out.AllowedPersistenceRegions = in.AllowedPersistenceRegions
	out.EnforceInTransit = direct.LazyPtr(in.GetEnforceInTransit())
	return out
}
func MessageStoragePolicy_ToProto(mapCtx *direct.MapContext, in *krm.MessageStoragePolicy) *pb.MessageStoragePolicy {
	if in == nil {
		return nil
	}
	out := &pb.MessageStoragePolicy{}
	out.AllowedPersistenceRegions = in.AllowedPersistenceRegions
	out.EnforceInTransit = direct.ValueOf(in.EnforceInTransit)
	return out
}
func PlatformLogsSettings_FromProto(mapCtx *direct.MapContext, in *pb.PlatformLogsSettings) *krm.PlatformLogsSettings {
	if in == nil {
		return nil
	}
	out := &krm.PlatformLogsSettings{}
	out.Severity = direct.Enum_FromProto(mapCtx, in.GetSeverity())
	return out
}
func PlatformLogsSettings_ToProto(mapCtx *direct.MapContext, in *krm.PlatformLogsSettings) *pb.PlatformLogsSettings {
	if in == nil {
		return nil
	}
	out := &pb.PlatformLogsSettings{}
	out.Severity = direct.Enum_ToProto[pb.PlatformLogsSettings_Severity](mapCtx, in.Severity)
	return out
}
func PubsubTopicObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Topic) *krm.PubsubTopicObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PubsubTopicObservedState{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: MessageStoragePolicy
	// MISSING: KMSKeyName
	// MISSING: SchemaSettings
	// MISSING: SatisfiesPzs
	// MISSING: MessageRetentionDuration
	// MISSING: State
	// MISSING: IngestionDataSourceSettings
	return out
}
func PubsubTopicObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PubsubTopicObservedState) *pb.Topic {
	if in == nil {
		return nil
	}
	out := &pb.Topic{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: MessageStoragePolicy
	// MISSING: KMSKeyName
	// MISSING: SchemaSettings
	// MISSING: SatisfiesPzs
	// MISSING: MessageRetentionDuration
	// MISSING: State
	// MISSING: IngestionDataSourceSettings
	return out
}
func PubsubTopicSpec_FromProto(mapCtx *direct.MapContext, in *pb.Topic) *krm.PubsubTopicSpec {
	if in == nil {
		return nil
	}
	out := &krm.PubsubTopicSpec{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: MessageStoragePolicy
	// MISSING: KMSKeyName
	// MISSING: SchemaSettings
	// MISSING: SatisfiesPzs
	// MISSING: MessageRetentionDuration
	// MISSING: State
	// MISSING: IngestionDataSourceSettings
	return out
}
func PubsubTopicSpec_ToProto(mapCtx *direct.MapContext, in *krm.PubsubTopicSpec) *pb.Topic {
	if in == nil {
		return nil
	}
	out := &pb.Topic{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: MessageStoragePolicy
	// MISSING: KMSKeyName
	// MISSING: SchemaSettings
	// MISSING: SatisfiesPzs
	// MISSING: MessageRetentionDuration
	// MISSING: State
	// MISSING: IngestionDataSourceSettings
	return out
}
func SchemaSettings_FromProto(mapCtx *direct.MapContext, in *pb.SchemaSettings) *krm.SchemaSettings {
	if in == nil {
		return nil
	}
	out := &krm.SchemaSettings{}
	out.Schema = direct.LazyPtr(in.GetSchema())
	out.Encoding = direct.Enum_FromProto(mapCtx, in.GetEncoding())
	out.FirstRevisionID = direct.LazyPtr(in.GetFirstRevisionId())
	out.LastRevisionID = direct.LazyPtr(in.GetLastRevisionId())
	return out
}
func SchemaSettings_ToProto(mapCtx *direct.MapContext, in *krm.SchemaSettings) *pb.SchemaSettings {
	if in == nil {
		return nil
	}
	out := &pb.SchemaSettings{}
	out.Schema = direct.ValueOf(in.Schema)
	out.Encoding = direct.Enum_ToProto[pb.Encoding](mapCtx, in.Encoding)
	out.FirstRevisionId = direct.ValueOf(in.FirstRevisionID)
	out.LastRevisionId = direct.ValueOf(in.LastRevisionID)
	return out
}
func Topic_FromProto(mapCtx *direct.MapContext, in *pb.Topic) *krm.Topic {
	if in == nil {
		return nil
	}
	out := &krm.Topic{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Labels = in.Labels
	out.MessageStoragePolicy = MessageStoragePolicy_FromProto(mapCtx, in.GetMessageStoragePolicy())
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	out.SchemaSettings = SchemaSettings_FromProto(mapCtx, in.GetSchemaSettings())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.MessageRetentionDuration = direct.StringDuration_FromProto(mapCtx, in.GetMessageRetentionDuration())
	// MISSING: State
	out.IngestionDataSourceSettings = IngestionDataSourceSettings_FromProto(mapCtx, in.GetIngestionDataSourceSettings())
	return out
}
func Topic_ToProto(mapCtx *direct.MapContext, in *krm.Topic) *pb.Topic {
	if in == nil {
		return nil
	}
	out := &pb.Topic{}
	out.Name = direct.ValueOf(in.Name)
	out.Labels = in.Labels
	out.MessageStoragePolicy = MessageStoragePolicy_ToProto(mapCtx, in.MessageStoragePolicy)
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	out.SchemaSettings = SchemaSettings_ToProto(mapCtx, in.SchemaSettings)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.MessageRetentionDuration = direct.StringDuration_ToProto(mapCtx, in.MessageRetentionDuration)
	// MISSING: State
	out.IngestionDataSourceSettings = IngestionDataSourceSettings_ToProto(mapCtx, in.IngestionDataSourceSettings)
	return out
}
func TopicObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Topic) *krm.TopicObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TopicObservedState{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: MessageStoragePolicy
	// MISSING: KMSKeyName
	// MISSING: SchemaSettings
	// MISSING: SatisfiesPzs
	// MISSING: MessageRetentionDuration
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.IngestionDataSourceSettings = IngestionDataSourceSettingsObservedState_FromProto(mapCtx, in.GetIngestionDataSourceSettings())
	return out
}
func TopicObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TopicObservedState) *pb.Topic {
	if in == nil {
		return nil
	}
	out := &pb.Topic{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: MessageStoragePolicy
	// MISSING: KMSKeyName
	// MISSING: SchemaSettings
	// MISSING: SatisfiesPzs
	// MISSING: MessageRetentionDuration
	out.State = direct.Enum_ToProto[pb.Topic_State](mapCtx, in.State)
	out.IngestionDataSourceSettings = IngestionDataSourceSettingsObservedState_ToProto(mapCtx, in.IngestionDataSourceSettings)
	return out
}
