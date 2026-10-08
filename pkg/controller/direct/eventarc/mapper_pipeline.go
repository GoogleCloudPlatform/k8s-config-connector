// Copyright 2026 Google LLC
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
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func EventarcPipelineSpec_FromProto(mapCtx *direct.MapContext, in *pb.Pipeline) *krm.EventarcPipelineSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcPipelineSpec{}
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Destinations = direct.Slice_FromProto(mapCtx, in.Destinations, Pipeline_Destination_FromProto)
	out.Mediations = direct.Slice_FromProto(mapCtx, in.Mediations, Pipeline_Mediation_FromProto)
	if in.GetCryptoKeyName() != "" {
		out.CryptoKeyRef = &kmsv1beta1.KMSCryptoKeyRef{
			External: in.GetCryptoKeyName(),
		}
	}
	out.InputPayloadFormat = Pipeline_MessagePayloadFormat_FromProto(mapCtx, in.GetInputPayloadFormat())
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	out.RetryPolicy = Pipeline_RetryPolicy_FromProto(mapCtx, in.GetRetryPolicy())
	return out
}

func EventarcPipelineSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcPipelineSpec) *pb.Pipeline {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline{}
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Destinations = direct.Slice_ToProto(mapCtx, in.Destinations, Pipeline_Destination_ToProto)
	out.Mediations = direct.Slice_ToProto(mapCtx, in.Mediations, Pipeline_Mediation_ToProto)
	if in.CryptoKeyRef != nil {
		out.CryptoKeyName = in.CryptoKeyRef.External
	}
	out.InputPayloadFormat = Pipeline_MessagePayloadFormat_ToProto(mapCtx, in.InputPayloadFormat)
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	out.RetryPolicy = Pipeline_RetryPolicy_ToProto(mapCtx, in.RetryPolicy)
	return out
}

func Pipeline_Destination_ToProto(mapCtx *direct.MapContext, in *krm.Pipeline_Destination) *pb.Pipeline_Destination {
	if in == nil {
		return nil
	}
	out := &pb.Pipeline_Destination{}
	out.NetworkConfig = Pipeline_Destination_NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	if oneof := Pipeline_Destination_HTTPEndpoint_ToProto(mapCtx, in.HTTPEndpoint); oneof != nil {
		out.DestinationDescriptor = &pb.Pipeline_Destination_HttpEndpoint_{HttpEndpoint: oneof}
	}
	if in.WorkflowRef != nil {
		out.DestinationDescriptor = &pb.Pipeline_Destination_Workflow{Workflow: in.WorkflowRef.External}
	}
	if in.MessageBusRef != nil {
		out.DestinationDescriptor = &pb.Pipeline_Destination_MessageBus{MessageBus: in.MessageBusRef.External}
	}
	if in.TopicRef != nil {
		out.DestinationDescriptor = &pb.Pipeline_Destination_Topic{Topic: in.TopicRef.External}
	}
	out.AuthenticationConfig = Pipeline_Destination_AuthenticationConfig_ToProto(mapCtx, in.AuthenticationConfig)
	out.OutputPayloadFormat = Pipeline_MessagePayloadFormat_ToProto(mapCtx, in.OutputPayloadFormat)
	return out
}
