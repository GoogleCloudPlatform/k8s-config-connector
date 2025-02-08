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
func BigQueryConfig_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryConfig) *krm.BigQueryConfig {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryConfig{}
	out.Table = direct.LazyPtr(in.GetTable())
	out.UseTopicSchema = direct.LazyPtr(in.GetUseTopicSchema())
	out.WriteMetadata = direct.LazyPtr(in.GetWriteMetadata())
	out.DropUnknownFields = direct.LazyPtr(in.GetDropUnknownFields())
	// MISSING: State
	out.UseTableSchema = direct.LazyPtr(in.GetUseTableSchema())
	out.ServiceAccountEmail = direct.LazyPtr(in.GetServiceAccountEmail())
	return out
}
func BigQueryConfig_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryConfig) *pb.BigQueryConfig {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryConfig{}
	out.Table = direct.ValueOf(in.Table)
	out.UseTopicSchema = direct.ValueOf(in.UseTopicSchema)
	out.WriteMetadata = direct.ValueOf(in.WriteMetadata)
	out.DropUnknownFields = direct.ValueOf(in.DropUnknownFields)
	// MISSING: State
	out.UseTableSchema = direct.ValueOf(in.UseTableSchema)
	out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountEmail)
	return out
}
func BigQueryConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryConfig) *krm.BigQueryConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryConfigObservedState{}
	// MISSING: Table
	// MISSING: UseTopicSchema
	// MISSING: WriteMetadata
	// MISSING: DropUnknownFields
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: UseTableSchema
	// MISSING: ServiceAccountEmail
	return out
}
func BigQueryConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryConfigObservedState) *pb.BigQueryConfig {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryConfig{}
	// MISSING: Table
	// MISSING: UseTopicSchema
	// MISSING: WriteMetadata
	// MISSING: DropUnknownFields
	out.State = direct.Enum_ToProto[pb.BigQueryConfig_State](mapCtx, in.State)
	// MISSING: UseTableSchema
	// MISSING: ServiceAccountEmail
	return out
}
func CloudStorageConfig_FromProto(mapCtx *direct.MapContext, in *pb.CloudStorageConfig) *krm.CloudStorageConfig {
	if in == nil {
		return nil
	}
	out := &krm.CloudStorageConfig{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	out.FilenamePrefix = direct.LazyPtr(in.GetFilenamePrefix())
	out.FilenameSuffix = direct.LazyPtr(in.GetFilenameSuffix())
	out.FilenameDatetimeFormat = direct.LazyPtr(in.GetFilenameDatetimeFormat())
	out.TextConfig = CloudStorageConfig_TextConfig_FromProto(mapCtx, in.GetTextConfig())
	out.AvroConfig = CloudStorageConfig_AvroConfig_FromProto(mapCtx, in.GetAvroConfig())
	out.MaxDuration = direct.StringDuration_FromProto(mapCtx, in.GetMaxDuration())
	out.MaxBytes = direct.LazyPtr(in.GetMaxBytes())
	out.MaxMessages = direct.LazyPtr(in.GetMaxMessages())
	// MISSING: State
	out.ServiceAccountEmail = direct.LazyPtr(in.GetServiceAccountEmail())
	return out
}
func CloudStorageConfig_ToProto(mapCtx *direct.MapContext, in *krm.CloudStorageConfig) *pb.CloudStorageConfig {
	if in == nil {
		return nil
	}
	out := &pb.CloudStorageConfig{}
	out.Bucket = direct.ValueOf(in.Bucket)
	out.FilenamePrefix = direct.ValueOf(in.FilenamePrefix)
	out.FilenameSuffix = direct.ValueOf(in.FilenameSuffix)
	out.FilenameDatetimeFormat = direct.ValueOf(in.FilenameDatetimeFormat)
	if oneof := CloudStorageConfig_TextConfig_ToProto(mapCtx, in.TextConfig); oneof != nil {
		out.OutputFormat = &pb.CloudStorageConfig_TextConfig_{TextConfig: oneof}
	}
	if oneof := CloudStorageConfig_AvroConfig_ToProto(mapCtx, in.AvroConfig); oneof != nil {
		out.OutputFormat = &pb.CloudStorageConfig_AvroConfig_{AvroConfig: oneof}
	}
	out.MaxDuration = direct.StringDuration_ToProto(mapCtx, in.MaxDuration)
	out.MaxBytes = direct.ValueOf(in.MaxBytes)
	out.MaxMessages = direct.ValueOf(in.MaxMessages)
	// MISSING: State
	out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountEmail)
	return out
}
func CloudStorageConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudStorageConfig) *krm.CloudStorageConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudStorageConfigObservedState{}
	// MISSING: Bucket
	// MISSING: FilenamePrefix
	// MISSING: FilenameSuffix
	// MISSING: FilenameDatetimeFormat
	// MISSING: TextConfig
	// MISSING: AvroConfig
	// MISSING: MaxDuration
	// MISSING: MaxBytes
	// MISSING: MaxMessages
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: ServiceAccountEmail
	return out
}
func CloudStorageConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudStorageConfigObservedState) *pb.CloudStorageConfig {
	if in == nil {
		return nil
	}
	out := &pb.CloudStorageConfig{}
	// MISSING: Bucket
	// MISSING: FilenamePrefix
	// MISSING: FilenameSuffix
	// MISSING: FilenameDatetimeFormat
	// MISSING: TextConfig
	// MISSING: AvroConfig
	// MISSING: MaxDuration
	// MISSING: MaxBytes
	// MISSING: MaxMessages
	out.State = direct.Enum_ToProto[pb.CloudStorageConfig_State](mapCtx, in.State)
	// MISSING: ServiceAccountEmail
	return out
}
func CloudStorageConfig_AvroConfig_FromProto(mapCtx *direct.MapContext, in *pb.CloudStorageConfig_AvroConfig) *krm.CloudStorageConfig_AvroConfig {
	if in == nil {
		return nil
	}
	out := &krm.CloudStorageConfig_AvroConfig{}
	out.WriteMetadata = direct.LazyPtr(in.GetWriteMetadata())
	out.UseTopicSchema = direct.LazyPtr(in.GetUseTopicSchema())
	return out
}
func CloudStorageConfig_AvroConfig_ToProto(mapCtx *direct.MapContext, in *krm.CloudStorageConfig_AvroConfig) *pb.CloudStorageConfig_AvroConfig {
	if in == nil {
		return nil
	}
	out := &pb.CloudStorageConfig_AvroConfig{}
	out.WriteMetadata = direct.ValueOf(in.WriteMetadata)
	out.UseTopicSchema = direct.ValueOf(in.UseTopicSchema)
	return out
}
func CloudStorageConfig_TextConfig_FromProto(mapCtx *direct.MapContext, in *pb.CloudStorageConfig_TextConfig) *krm.CloudStorageConfig_TextConfig {
	if in == nil {
		return nil
	}
	out := &krm.CloudStorageConfig_TextConfig{}
	return out
}
func CloudStorageConfig_TextConfig_ToProto(mapCtx *direct.MapContext, in *krm.CloudStorageConfig_TextConfig) *pb.CloudStorageConfig_TextConfig {
	if in == nil {
		return nil
	}
	out := &pb.CloudStorageConfig_TextConfig{}
	return out
}
func DeadLetterPolicy_FromProto(mapCtx *direct.MapContext, in *pb.DeadLetterPolicy) *krm.DeadLetterPolicy {
	if in == nil {
		return nil
	}
	out := &krm.DeadLetterPolicy{}
	out.DeadLetterTopic = direct.LazyPtr(in.GetDeadLetterTopic())
	out.MaxDeliveryAttempts = direct.LazyPtr(in.GetMaxDeliveryAttempts())
	return out
}
func DeadLetterPolicy_ToProto(mapCtx *direct.MapContext, in *krm.DeadLetterPolicy) *pb.DeadLetterPolicy {
	if in == nil {
		return nil
	}
	out := &pb.DeadLetterPolicy{}
	out.DeadLetterTopic = direct.ValueOf(in.DeadLetterTopic)
	out.MaxDeliveryAttempts = direct.ValueOf(in.MaxDeliveryAttempts)
	return out
}
func ExpirationPolicy_FromProto(mapCtx *direct.MapContext, in *pb.ExpirationPolicy) *krm.ExpirationPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ExpirationPolicy{}
	out.Ttl = direct.StringDuration_FromProto(mapCtx, in.GetTtl())
	return out
}
func ExpirationPolicy_ToProto(mapCtx *direct.MapContext, in *krm.ExpirationPolicy) *pb.ExpirationPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ExpirationPolicy{}
	out.Ttl = direct.StringDuration_ToProto(mapCtx, in.Ttl)
	return out
}
func PubsubSubscriptionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Subscription) *krm.PubsubSubscriptionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PubsubSubscriptionObservedState{}
	// MISSING: Name
	// MISSING: Topic
	// MISSING: PushConfig
	// MISSING: BigqueryConfig
	// MISSING: CloudStorageConfig
	// MISSING: AckDeadlineSeconds
	// MISSING: RetainAckedMessages
	// MISSING: MessageRetentionDuration
	// MISSING: Labels
	// MISSING: EnableMessageOrdering
	// MISSING: ExpirationPolicy
	// MISSING: Filter
	// MISSING: DeadLetterPolicy
	// MISSING: RetryPolicy
	// MISSING: Detached
	// MISSING: EnableExactlyOnceDelivery
	// MISSING: TopicMessageRetentionDuration
	// MISSING: State
	// MISSING: AnalyticsHubSubscriptionInfo
	return out
}
func PubsubSubscriptionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PubsubSubscriptionObservedState) *pb.Subscription {
	if in == nil {
		return nil
	}
	out := &pb.Subscription{}
	// MISSING: Name
	// MISSING: Topic
	// MISSING: PushConfig
	// MISSING: BigqueryConfig
	// MISSING: CloudStorageConfig
	// MISSING: AckDeadlineSeconds
	// MISSING: RetainAckedMessages
	// MISSING: MessageRetentionDuration
	// MISSING: Labels
	// MISSING: EnableMessageOrdering
	// MISSING: ExpirationPolicy
	// MISSING: Filter
	// MISSING: DeadLetterPolicy
	// MISSING: RetryPolicy
	// MISSING: Detached
	// MISSING: EnableExactlyOnceDelivery
	// MISSING: TopicMessageRetentionDuration
	// MISSING: State
	// MISSING: AnalyticsHubSubscriptionInfo
	return out
}
func PubsubSubscriptionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Subscription) *krm.PubsubSubscriptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.PubsubSubscriptionSpec{}
	// MISSING: Name
	// MISSING: Topic
	// MISSING: PushConfig
	// MISSING: BigqueryConfig
	// MISSING: CloudStorageConfig
	// MISSING: AckDeadlineSeconds
	// MISSING: RetainAckedMessages
	// MISSING: MessageRetentionDuration
	// MISSING: Labels
	// MISSING: EnableMessageOrdering
	// MISSING: ExpirationPolicy
	// MISSING: Filter
	// MISSING: DeadLetterPolicy
	// MISSING: RetryPolicy
	// MISSING: Detached
	// MISSING: EnableExactlyOnceDelivery
	// MISSING: TopicMessageRetentionDuration
	// MISSING: State
	// MISSING: AnalyticsHubSubscriptionInfo
	return out
}
func PubsubSubscriptionSpec_ToProto(mapCtx *direct.MapContext, in *krm.PubsubSubscriptionSpec) *pb.Subscription {
	if in == nil {
		return nil
	}
	out := &pb.Subscription{}
	// MISSING: Name
	// MISSING: Topic
	// MISSING: PushConfig
	// MISSING: BigqueryConfig
	// MISSING: CloudStorageConfig
	// MISSING: AckDeadlineSeconds
	// MISSING: RetainAckedMessages
	// MISSING: MessageRetentionDuration
	// MISSING: Labels
	// MISSING: EnableMessageOrdering
	// MISSING: ExpirationPolicy
	// MISSING: Filter
	// MISSING: DeadLetterPolicy
	// MISSING: RetryPolicy
	// MISSING: Detached
	// MISSING: EnableExactlyOnceDelivery
	// MISSING: TopicMessageRetentionDuration
	// MISSING: State
	// MISSING: AnalyticsHubSubscriptionInfo
	return out
}
func PushConfig_FromProto(mapCtx *direct.MapContext, in *pb.PushConfig) *krm.PushConfig {
	if in == nil {
		return nil
	}
	out := &krm.PushConfig{}
	out.PushEndpoint = direct.LazyPtr(in.GetPushEndpoint())
	out.Attributes = in.Attributes
	out.OidcToken = PushConfig_OidcToken_FromProto(mapCtx, in.GetOidcToken())
	out.PubsubWrapper = PushConfig_PubsubWrapper_FromProto(mapCtx, in.GetPubsubWrapper())
	out.NoWrapper = PushConfig_NoWrapper_FromProto(mapCtx, in.GetNoWrapper())
	return out
}
func PushConfig_ToProto(mapCtx *direct.MapContext, in *krm.PushConfig) *pb.PushConfig {
	if in == nil {
		return nil
	}
	out := &pb.PushConfig{}
	out.PushEndpoint = direct.ValueOf(in.PushEndpoint)
	out.Attributes = in.Attributes
	if oneof := PushConfig_OidcToken_ToProto(mapCtx, in.OidcToken); oneof != nil {
		out.AuthenticationMethod = &pb.PushConfig_OidcToken_{OidcToken: oneof}
	}
	if oneof := PushConfig_PubsubWrapper_ToProto(mapCtx, in.PubsubWrapper); oneof != nil {
		out.Wrapper = &pb.PushConfig_PubsubWrapper_{PubsubWrapper: oneof}
	}
	if oneof := PushConfig_NoWrapper_ToProto(mapCtx, in.NoWrapper); oneof != nil {
		out.Wrapper = &pb.PushConfig_NoWrapper_{NoWrapper: oneof}
	}
	return out
}
func PushConfig_NoWrapper_FromProto(mapCtx *direct.MapContext, in *pb.PushConfig_NoWrapper) *krm.PushConfig_NoWrapper {
	if in == nil {
		return nil
	}
	out := &krm.PushConfig_NoWrapper{}
	out.WriteMetadata = direct.LazyPtr(in.GetWriteMetadata())
	return out
}
func PushConfig_NoWrapper_ToProto(mapCtx *direct.MapContext, in *krm.PushConfig_NoWrapper) *pb.PushConfig_NoWrapper {
	if in == nil {
		return nil
	}
	out := &pb.PushConfig_NoWrapper{}
	out.WriteMetadata = direct.ValueOf(in.WriteMetadata)
	return out
}
func PushConfig_OidcToken_FromProto(mapCtx *direct.MapContext, in *pb.PushConfig_OidcToken) *krm.PushConfig_OidcToken {
	if in == nil {
		return nil
	}
	out := &krm.PushConfig_OidcToken{}
	out.ServiceAccountEmail = direct.LazyPtr(in.GetServiceAccountEmail())
	out.Audience = direct.LazyPtr(in.GetAudience())
	return out
}
func PushConfig_OidcToken_ToProto(mapCtx *direct.MapContext, in *krm.PushConfig_OidcToken) *pb.PushConfig_OidcToken {
	if in == nil {
		return nil
	}
	out := &pb.PushConfig_OidcToken{}
	out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountEmail)
	out.Audience = direct.ValueOf(in.Audience)
	return out
}
func PushConfig_PubsubWrapper_FromProto(mapCtx *direct.MapContext, in *pb.PushConfig_PubsubWrapper) *krm.PushConfig_PubsubWrapper {
	if in == nil {
		return nil
	}
	out := &krm.PushConfig_PubsubWrapper{}
	return out
}
func PushConfig_PubsubWrapper_ToProto(mapCtx *direct.MapContext, in *krm.PushConfig_PubsubWrapper) *pb.PushConfig_PubsubWrapper {
	if in == nil {
		return nil
	}
	out := &pb.PushConfig_PubsubWrapper{}
	return out
}
func RetryPolicy_FromProto(mapCtx *direct.MapContext, in *pb.RetryPolicy) *krm.RetryPolicy {
	if in == nil {
		return nil
	}
	out := &krm.RetryPolicy{}
	out.MinimumBackoff = direct.StringDuration_FromProto(mapCtx, in.GetMinimumBackoff())
	out.MaximumBackoff = direct.StringDuration_FromProto(mapCtx, in.GetMaximumBackoff())
	return out
}
func RetryPolicy_ToProto(mapCtx *direct.MapContext, in *krm.RetryPolicy) *pb.RetryPolicy {
	if in == nil {
		return nil
	}
	out := &pb.RetryPolicy{}
	out.MinimumBackoff = direct.StringDuration_ToProto(mapCtx, in.MinimumBackoff)
	out.MaximumBackoff = direct.StringDuration_ToProto(mapCtx, in.MaximumBackoff)
	return out
}
func Subscription_FromProto(mapCtx *direct.MapContext, in *pb.Subscription) *krm.Subscription {
	if in == nil {
		return nil
	}
	out := &krm.Subscription{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Topic = direct.LazyPtr(in.GetTopic())
	out.PushConfig = PushConfig_FromProto(mapCtx, in.GetPushConfig())
	out.BigqueryConfig = BigQueryConfig_FromProto(mapCtx, in.GetBigqueryConfig())
	out.CloudStorageConfig = CloudStorageConfig_FromProto(mapCtx, in.GetCloudStorageConfig())
	out.AckDeadlineSeconds = direct.LazyPtr(in.GetAckDeadlineSeconds())
	out.RetainAckedMessages = direct.LazyPtr(in.GetRetainAckedMessages())
	out.MessageRetentionDuration = direct.StringDuration_FromProto(mapCtx, in.GetMessageRetentionDuration())
	out.Labels = in.Labels
	out.EnableMessageOrdering = direct.LazyPtr(in.GetEnableMessageOrdering())
	out.ExpirationPolicy = ExpirationPolicy_FromProto(mapCtx, in.GetExpirationPolicy())
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.DeadLetterPolicy = DeadLetterPolicy_FromProto(mapCtx, in.GetDeadLetterPolicy())
	out.RetryPolicy = RetryPolicy_FromProto(mapCtx, in.GetRetryPolicy())
	out.Detached = direct.LazyPtr(in.GetDetached())
	out.EnableExactlyOnceDelivery = direct.LazyPtr(in.GetEnableExactlyOnceDelivery())
	// MISSING: TopicMessageRetentionDuration
	// MISSING: State
	// MISSING: AnalyticsHubSubscriptionInfo
	return out
}
func Subscription_ToProto(mapCtx *direct.MapContext, in *krm.Subscription) *pb.Subscription {
	if in == nil {
		return nil
	}
	out := &pb.Subscription{}
	out.Name = direct.ValueOf(in.Name)
	out.Topic = direct.ValueOf(in.Topic)
	out.PushConfig = PushConfig_ToProto(mapCtx, in.PushConfig)
	out.BigqueryConfig = BigQueryConfig_ToProto(mapCtx, in.BigqueryConfig)
	out.CloudStorageConfig = CloudStorageConfig_ToProto(mapCtx, in.CloudStorageConfig)
	out.AckDeadlineSeconds = direct.ValueOf(in.AckDeadlineSeconds)
	out.RetainAckedMessages = direct.ValueOf(in.RetainAckedMessages)
	out.MessageRetentionDuration = direct.StringDuration_ToProto(mapCtx, in.MessageRetentionDuration)
	out.Labels = in.Labels
	out.EnableMessageOrdering = direct.ValueOf(in.EnableMessageOrdering)
	out.ExpirationPolicy = ExpirationPolicy_ToProto(mapCtx, in.ExpirationPolicy)
	out.Filter = direct.ValueOf(in.Filter)
	out.DeadLetterPolicy = DeadLetterPolicy_ToProto(mapCtx, in.DeadLetterPolicy)
	out.RetryPolicy = RetryPolicy_ToProto(mapCtx, in.RetryPolicy)
	out.Detached = direct.ValueOf(in.Detached)
	out.EnableExactlyOnceDelivery = direct.ValueOf(in.EnableExactlyOnceDelivery)
	// MISSING: TopicMessageRetentionDuration
	// MISSING: State
	// MISSING: AnalyticsHubSubscriptionInfo
	return out
}
func SubscriptionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Subscription) *krm.SubscriptionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SubscriptionObservedState{}
	// MISSING: Name
	// MISSING: Topic
	// MISSING: PushConfig
	out.BigqueryConfig = BigQueryConfigObservedState_FromProto(mapCtx, in.GetBigqueryConfig())
	out.CloudStorageConfig = CloudStorageConfigObservedState_FromProto(mapCtx, in.GetCloudStorageConfig())
	// MISSING: AckDeadlineSeconds
	// MISSING: RetainAckedMessages
	// MISSING: MessageRetentionDuration
	// MISSING: Labels
	// MISSING: EnableMessageOrdering
	// MISSING: ExpirationPolicy
	// MISSING: Filter
	// MISSING: DeadLetterPolicy
	// MISSING: RetryPolicy
	// MISSING: Detached
	// MISSING: EnableExactlyOnceDelivery
	out.TopicMessageRetentionDuration = direct.StringDuration_FromProto(mapCtx, in.GetTopicMessageRetentionDuration())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.AnalyticsHubSubscriptionInfo = Subscription_AnalyticsHubSubscriptionInfo_FromProto(mapCtx, in.GetAnalyticsHubSubscriptionInfo())
	return out
}
func SubscriptionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SubscriptionObservedState) *pb.Subscription {
	if in == nil {
		return nil
	}
	out := &pb.Subscription{}
	// MISSING: Name
	// MISSING: Topic
	// MISSING: PushConfig
	out.BigqueryConfig = BigQueryConfigObservedState_ToProto(mapCtx, in.BigqueryConfig)
	out.CloudStorageConfig = CloudStorageConfigObservedState_ToProto(mapCtx, in.CloudStorageConfig)
	// MISSING: AckDeadlineSeconds
	// MISSING: RetainAckedMessages
	// MISSING: MessageRetentionDuration
	// MISSING: Labels
	// MISSING: EnableMessageOrdering
	// MISSING: ExpirationPolicy
	// MISSING: Filter
	// MISSING: DeadLetterPolicy
	// MISSING: RetryPolicy
	// MISSING: Detached
	// MISSING: EnableExactlyOnceDelivery
	out.TopicMessageRetentionDuration = direct.StringDuration_ToProto(mapCtx, in.TopicMessageRetentionDuration)
	out.State = direct.Enum_ToProto[pb.Subscription_State](mapCtx, in.State)
	out.AnalyticsHubSubscriptionInfo = Subscription_AnalyticsHubSubscriptionInfo_ToProto(mapCtx, in.AnalyticsHubSubscriptionInfo)
	return out
}
func Subscription_AnalyticsHubSubscriptionInfo_FromProto(mapCtx *direct.MapContext, in *pb.Subscription_AnalyticsHubSubscriptionInfo) *krm.Subscription_AnalyticsHubSubscriptionInfo {
	if in == nil {
		return nil
	}
	out := &krm.Subscription_AnalyticsHubSubscriptionInfo{}
	out.Listing = direct.LazyPtr(in.GetListing())
	out.Subscription = direct.LazyPtr(in.GetSubscription())
	return out
}
func Subscription_AnalyticsHubSubscriptionInfo_ToProto(mapCtx *direct.MapContext, in *krm.Subscription_AnalyticsHubSubscriptionInfo) *pb.Subscription_AnalyticsHubSubscriptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.Subscription_AnalyticsHubSubscriptionInfo{}
	out.Listing = direct.ValueOf(in.Listing)
	out.Subscription = direct.ValueOf(in.Subscription)
	return out
}
