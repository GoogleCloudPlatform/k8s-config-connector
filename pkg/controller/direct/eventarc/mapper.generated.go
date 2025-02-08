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
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func EventarcPipelineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline) *krm.EventarcPipelineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventarcPipelineObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: Destinations
	// MISSING: Mediations
	// MISSING: CryptoKeyName
	// MISSING: InputPayloadFormat
	// MISSING: LoggingConfig
	// MISSING: RetryPolicy
	// MISSING: Etag
	return out
}
func EventarcPipelineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventarcPipelineObservedState) *pb.Pipeline {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: Destinations
	// MISSING: Mediations
	// MISSING: CryptoKeyName
	// MISSING: InputPayloadFormat
	// MISSING: LoggingConfig
	// MISSING: RetryPolicy
	// MISSING: Etag
	return out
}
func EventarcPipelineSpec_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline) *krm.EventarcPipelineSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcPipelineSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: Destinations
	// MISSING: Mediations
	// MISSING: CryptoKeyName
	// MISSING: InputPayloadFormat
	// MISSING: LoggingConfig
	// MISSING: RetryPolicy
	// MISSING: Etag
	return out
}
func EventarcPipelineSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcPipelineSpec) *pb.Pipeline {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Uid
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: Destinations
	// MISSING: Mediations
	// MISSING: CryptoKeyName
	// MISSING: InputPayloadFormat
	// MISSING: LoggingConfig
	// MISSING: RetryPolicy
	// MISSING: Etag
	return out
}
func LoggingConfig_FromProto(mapCtx *direct.MapContext, in *pb.LoggingConfig) *krm.LoggingConfig {
	if in == nil {
		return nil
	}
	out := &krm.LoggingConfig{}
	out.LogSeverity = direct.Enum_FromProto(mapCtx, in.GetLogSeverity())
	return out
}
func LoggingConfig_ToProto(mapCtx *direct.MapContext, in *krm.LoggingConfig) *pb.LoggingConfig {
	if in == nil {
		return nil
	}
	out := &pb.LoggingConfig{}
	out.LogSeverity = direct.Enum_ToProto[pb.LoggingConfig_LogSeverity](mapCtx, in.LogSeverity)
	return out
}
func Pipeline_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline) *krm.Pipeline {
	if in == nil {
		return nil
	}
	out := &krm.Pipeline{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: Uid
	out.Annotations = in.Annotations
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Destinations = direct.Slice_FromProto(mapCtx, in.Destinations, Pipeline_Destination_FromProto)
	out.Mediations = direct.Slice_FromProto(mapCtx, in.Mediations, Pipeline_Mediation_FromProto)
	out.CryptoKeyName = direct.LazyPtr(in.GetCryptoKeyName())
	out.InputPayloadFormat = Pipeline_MessagePayloadFormat_FromProto(mapCtx, in.GetInputPayloadFormat())
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	out.RetryPolicy = Pipeline_RetryPolicy_FromProto(mapCtx, in.GetRetryPolicy())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func Pipeline_ToProto(mapCtx *direct.MapContext, in *krm.Pipeline) *pb.Pipeline {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: Uid
	out.Annotations = in.Annotations
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Destinations = direct.Slice_ToProto(mapCtx, in.Destinations, Pipeline_Destination_ToProto)
	out.Mediations = direct.Slice_ToProto(mapCtx, in.Mediations, Pipeline_Mediation_ToProto)
	out.CryptoKeyName = direct.ValueOf(in.CryptoKeyName)
	out.InputPayloadFormat = Pipeline_MessagePayloadFormat_ToProto(mapCtx, in.InputPayloadFormat)
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	out.RetryPolicy = Pipeline_RetryPolicy_ToProto(mapCtx, in.RetryPolicy)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func PipelineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline) *krm.PipelineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PipelineObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: Destinations
	// MISSING: Mediations
	// MISSING: CryptoKeyName
	// MISSING: InputPayloadFormat
	// MISSING: LoggingConfig
	// MISSING: RetryPolicy
	// MISSING: Etag
	return out
}
func PipelineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PipelineObservedState) *pb.Pipeline {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: Destinations
	// MISSING: Mediations
	// MISSING: CryptoKeyName
	// MISSING: InputPayloadFormat
	// MISSING: LoggingConfig
	// MISSING: RetryPolicy
	// MISSING: Etag
	return out
}
func Pipeline_Destination_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline_Destination) *krm.Pipeline_Destination {
	if in == nil {
		return nil
	}
	out := &krm.Pipeline_Destination{}
	out.NetworkConfig = Pipeline_Destination_NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	out.HTTPEndpoint = Pipeline_Destination_HttpEndpoint_FromProto(mapCtx, in.GetHttpEndpoint())
	out.Workflow = direct.LazyPtr(in.GetWorkflow())
	out.MessageBus = direct.LazyPtr(in.GetMessageBus())
	out.Topic = direct.LazyPtr(in.GetTopic())
	out.AuthenticationConfig = Pipeline_Destination_AuthenticationConfig_FromProto(mapCtx, in.GetAuthenticationConfig())
	out.OutputPayloadFormat = Pipeline_MessagePayloadFormat_FromProto(mapCtx, in.GetOutputPayloadFormat())
	return out
}
func Pipeline_Destination_ToProto(mapCtx *direct.MapContext, in *krm.Pipeline_Destination) *pb.Pipeline_Destination {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline_Destination{}
	out.NetworkConfig = Pipeline_Destination_NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	if oneof := Pipeline_Destination_HttpEndpoint_ToProto(mapCtx, in.HTTPEndpoint); oneof != nil {
		out.DestinationDescriptor = &pb.Pipeline_Destination_HttpEndpoint_{HttpEndpoint: oneof}
	}
	if oneof := Pipeline_Destination_Workflow_ToProto(mapCtx, in.Workflow); oneof != nil {
		out.DestinationDescriptor = oneof
	}
	if oneof := Pipeline_Destination_MessageBus_ToProto(mapCtx, in.MessageBus); oneof != nil {
		out.DestinationDescriptor = oneof
	}
	if oneof := Pipeline_Destination_Topic_ToProto(mapCtx, in.Topic); oneof != nil {
		out.DestinationDescriptor = oneof
	}
	out.AuthenticationConfig = Pipeline_Destination_AuthenticationConfig_ToProto(mapCtx, in.AuthenticationConfig)
	out.OutputPayloadFormat = Pipeline_MessagePayloadFormat_ToProto(mapCtx, in.OutputPayloadFormat)
	return out
}
func Pipeline_Destination_AuthenticationConfig_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline_Destination_AuthenticationConfig) *krm.Pipeline_Destination_AuthenticationConfig {
	if in == nil {
		return nil
	}
	out := &krm.Pipeline_Destination_AuthenticationConfig{}
	out.GoogleOidc = Pipeline_Destination_AuthenticationConfig_OidcToken_FromProto(mapCtx, in.GetGoogleOidc())
	out.OauthToken = Pipeline_Destination_AuthenticationConfig_OAuthToken_FromProto(mapCtx, in.GetOauthToken())
	return out
}
func Pipeline_Destination_AuthenticationConfig_ToProto(mapCtx *direct.MapContext, in *krm.Pipeline_Destination_AuthenticationConfig) *pb.Pipeline_Destination_AuthenticationConfig {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline_Destination_AuthenticationConfig{}
	if oneof := Pipeline_Destination_AuthenticationConfig_OidcToken_ToProto(mapCtx, in.GoogleOidc); oneof != nil {
		out.AuthenticationMethodDescriptor = &pb.Pipeline_Destination_AuthenticationConfig_GoogleOidc{GoogleOidc: oneof}
	}
	if oneof := Pipeline_Destination_AuthenticationConfig_OAuthToken_ToProto(mapCtx, in.OauthToken); oneof != nil {
		out.AuthenticationMethodDescriptor = &pb.Pipeline_Destination_AuthenticationConfig_OauthToken{OauthToken: oneof}
	}
	return out
}
func Pipeline_Destination_AuthenticationConfig_OAuthToken_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline_Destination_AuthenticationConfig_OAuthToken) *krm.Pipeline_Destination_AuthenticationConfig_OAuthToken {
	if in == nil {
		return nil
	}
	out := &krm.Pipeline_Destination_AuthenticationConfig_OAuthToken{}
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.Scope = direct.LazyPtr(in.GetScope())
	return out
}
func Pipeline_Destination_AuthenticationConfig_OAuthToken_ToProto(mapCtx *direct.MapContext, in *krm.Pipeline_Destination_AuthenticationConfig_OAuthToken) *pb.Pipeline_Destination_AuthenticationConfig_OAuthToken {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline_Destination_AuthenticationConfig_OAuthToken{}
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.Scope = direct.ValueOf(in.Scope)
	return out
}
func Pipeline_Destination_AuthenticationConfig_OidcToken_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline_Destination_AuthenticationConfig_OidcToken) *krm.Pipeline_Destination_AuthenticationConfig_OidcToken {
	if in == nil {
		return nil
	}
	out := &krm.Pipeline_Destination_AuthenticationConfig_OidcToken{}
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.Audience = direct.LazyPtr(in.GetAudience())
	return out
}
func Pipeline_Destination_AuthenticationConfig_OidcToken_ToProto(mapCtx *direct.MapContext, in *krm.Pipeline_Destination_AuthenticationConfig_OidcToken) *pb.Pipeline_Destination_AuthenticationConfig_OidcToken {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline_Destination_AuthenticationConfig_OidcToken{}
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.Audience = direct.ValueOf(in.Audience)
	return out
}
func Pipeline_Destination_HttpEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline_Destination_HttpEndpoint) *krm.Pipeline_Destination_HttpEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.Pipeline_Destination_HttpEndpoint{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.MessageBindingTemplate = direct.LazyPtr(in.GetMessageBindingTemplate())
	return out
}
func Pipeline_Destination_HttpEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.Pipeline_Destination_HttpEndpoint) *pb.Pipeline_Destination_HttpEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline_Destination_HttpEndpoint{}
	out.Uri = direct.ValueOf(in.URI)
	out.MessageBindingTemplate = direct.ValueOf(in.MessageBindingTemplate)
	return out
}
func Pipeline_Destination_NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline_Destination_NetworkConfig) *krm.Pipeline_Destination_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.Pipeline_Destination_NetworkConfig{}
	out.NetworkAttachment = direct.LazyPtr(in.GetNetworkAttachment())
	return out
}
func Pipeline_Destination_NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.Pipeline_Destination_NetworkConfig) *pb.Pipeline_Destination_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline_Destination_NetworkConfig{}
	out.NetworkAttachment = direct.ValueOf(in.NetworkAttachment)
	return out
}
func Pipeline_Mediation_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline_Mediation) *krm.Pipeline_Mediation {
	if in == nil {
		return nil
	}
	out := &krm.Pipeline_Mediation{}
	out.Transformation = Pipeline_Mediation_Transformation_FromProto(mapCtx, in.GetTransformation())
	return out
}
func Pipeline_Mediation_ToProto(mapCtx *direct.MapContext, in *krm.Pipeline_Mediation) *pb.Pipeline_Mediation {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline_Mediation{}
	if oneof := Pipeline_Mediation_Transformation_ToProto(mapCtx, in.Transformation); oneof != nil {
		out.MediationDescriptor = &pb.Pipeline_Mediation_Transformation_{Transformation: oneof}
	}
	return out
}
func Pipeline_Mediation_Transformation_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline_Mediation_Transformation) *krm.Pipeline_Mediation_Transformation {
	if in == nil {
		return nil
	}
	out := &krm.Pipeline_Mediation_Transformation{}
	out.TransformationTemplate = direct.LazyPtr(in.GetTransformationTemplate())
	return out
}
func Pipeline_Mediation_Transformation_ToProto(mapCtx *direct.MapContext, in *krm.Pipeline_Mediation_Transformation) *pb.Pipeline_Mediation_Transformation {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline_Mediation_Transformation{}
	out.TransformationTemplate = direct.ValueOf(in.TransformationTemplate)
	return out
}
func Pipeline_MessagePayloadFormat_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline_MessagePayloadFormat) *krm.Pipeline_MessagePayloadFormat {
	if in == nil {
		return nil
	}
	out := &krm.Pipeline_MessagePayloadFormat{}
	out.Protobuf = Pipeline_MessagePayloadFormat_ProtobufFormat_FromProto(mapCtx, in.GetProtobuf())
	out.Avro = Pipeline_MessagePayloadFormat_AvroFormat_FromProto(mapCtx, in.GetAvro())
	out.Json = Pipeline_MessagePayloadFormat_JsonFormat_FromProto(mapCtx, in.GetJson())
	return out
}
func Pipeline_MessagePayloadFormat_ToProto(mapCtx *direct.MapContext, in *krm.Pipeline_MessagePayloadFormat) *pb.Pipeline_MessagePayloadFormat {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline_MessagePayloadFormat{}
	if oneof := Pipeline_MessagePayloadFormat_ProtobufFormat_ToProto(mapCtx, in.Protobuf); oneof != nil {
		out.Kind = &pb.Pipeline_MessagePayloadFormat_Protobuf{Protobuf: oneof}
	}
	if oneof := Pipeline_MessagePayloadFormat_AvroFormat_ToProto(mapCtx, in.Avro); oneof != nil {
		out.Kind = &pb.Pipeline_MessagePayloadFormat_Avro{Avro: oneof}
	}
	if oneof := Pipeline_MessagePayloadFormat_JsonFormat_ToProto(mapCtx, in.Json); oneof != nil {
		out.Kind = &pb.Pipeline_MessagePayloadFormat_Json{Json: oneof}
	}
	return out
}
func Pipeline_MessagePayloadFormat_AvroFormat_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline_MessagePayloadFormat_AvroFormat) *krm.Pipeline_MessagePayloadFormat_AvroFormat {
	if in == nil {
		return nil
	}
	out := &krm.Pipeline_MessagePayloadFormat_AvroFormat{}
	out.SchemaDefinition = direct.LazyPtr(in.GetSchemaDefinition())
	return out
}
func Pipeline_MessagePayloadFormat_AvroFormat_ToProto(mapCtx *direct.MapContext, in *krm.Pipeline_MessagePayloadFormat_AvroFormat) *pb.Pipeline_MessagePayloadFormat_AvroFormat {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline_MessagePayloadFormat_AvroFormat{}
	out.SchemaDefinition = direct.ValueOf(in.SchemaDefinition)
	return out
}
func Pipeline_MessagePayloadFormat_JsonFormat_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline_MessagePayloadFormat_JsonFormat) *krm.Pipeline_MessagePayloadFormat_JsonFormat {
	if in == nil {
		return nil
	}
	out := &krm.Pipeline_MessagePayloadFormat_JsonFormat{}
	return out
}
func Pipeline_MessagePayloadFormat_JsonFormat_ToProto(mapCtx *direct.MapContext, in *krm.Pipeline_MessagePayloadFormat_JsonFormat) *pb.Pipeline_MessagePayloadFormat_JsonFormat {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline_MessagePayloadFormat_JsonFormat{}
	return out
}
func Pipeline_MessagePayloadFormat_ProtobufFormat_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline_MessagePayloadFormat_ProtobufFormat) *krm.Pipeline_MessagePayloadFormat_ProtobufFormat {
	if in == nil {
		return nil
	}
	out := &krm.Pipeline_MessagePayloadFormat_ProtobufFormat{}
	out.SchemaDefinition = direct.LazyPtr(in.GetSchemaDefinition())
	return out
}
func Pipeline_MessagePayloadFormat_ProtobufFormat_ToProto(mapCtx *direct.MapContext, in *krm.Pipeline_MessagePayloadFormat_ProtobufFormat) *pb.Pipeline_MessagePayloadFormat_ProtobufFormat {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline_MessagePayloadFormat_ProtobufFormat{}
	out.SchemaDefinition = direct.ValueOf(in.SchemaDefinition)
	return out
}
func Pipeline_RetryPolicy_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline_RetryPolicy) *krm.Pipeline_RetryPolicy {
	if in == nil {
		return nil
	}
	out := &krm.Pipeline_RetryPolicy{}
	out.MaxAttempts = direct.LazyPtr(in.GetMaxAttempts())
	out.MinRetryDelay = direct.StringDuration_FromProto(mapCtx, in.GetMinRetryDelay())
	out.MaxRetryDelay = direct.StringDuration_FromProto(mapCtx, in.GetMaxRetryDelay())
	return out
}
func Pipeline_RetryPolicy_ToProto(mapCtx *direct.MapContext, in *krm.Pipeline_RetryPolicy) *pb.Pipeline_RetryPolicy {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline_RetryPolicy{}
	out.MaxAttempts = direct.ValueOf(in.MaxAttempts)
	out.MinRetryDelay = direct.StringDuration_ToProto(mapCtx, in.MinRetryDelay)
	out.MaxRetryDelay = direct.StringDuration_ToProto(mapCtx, in.MaxRetryDelay)
	return out
}
