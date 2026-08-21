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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

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

func EventarcEnrollmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Enrollment) *krm.EventarcEnrollmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcEnrollmentSpec{}
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.CELMatch = direct.LazyPtr(in.GetCelMatch())
	if in.GetMessageBus() != "" {
		out.MessageBusRef = &krm.EventarcMessageBusRef{External: in.GetMessageBus()}
	}
	if in.GetDestination() != "" {
		out.DestinationRef = &krm.EventarcPipelineRef{External: in.GetDestination()}
	}
	return out
}

func EventarcEnrollmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcEnrollmentSpec) *pb.Enrollment {
	if in == nil {
		return nil
	}
	out := &pb.Enrollment{}
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.CelMatch = direct.ValueOf(in.CELMatch)
	if in.MessageBusRef != nil {
		out.MessageBus = in.MessageBusRef.External
	}
	if in.DestinationRef != nil {
		out.Destination = in.DestinationRef.External
	}
	return out
}
