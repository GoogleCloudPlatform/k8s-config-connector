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

package dialogflow

import (
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DialogflowAgentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Agent) *krm.DialogflowAgentSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowAgentSpec{}
	out.DisplayName = in.GetDisplayName()
	out.DefaultLanguageCode = in.GetDefaultLanguageCode()
	out.SupportedLanguageCodes = in.SupportedLanguageCodes
	out.TimeZone = in.GetTimeZone()
	out.Description = direct.LazyPtr(in.GetDescription())
	out.AvatarURI = direct.LazyPtr(in.GetAvatarUri())
	out.EnableLogging = direct.LazyPtr(in.GetEnableLogging())
	out.MatchMode = direct.Enum_FromProto(mapCtx, in.GetMatchMode())
	out.ClassificationThreshold = direct.LazyPtr(in.GetClassificationThreshold())
	out.APIVersion = direct.Enum_FromProto(mapCtx, in.GetApiVersion())
	out.Tier = direct.Enum_FromProto(mapCtx, in.GetTier())
	return out
}

func DialogflowAgentSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowAgentSpec) *pb.Agent {
	if in == nil {
		return nil
	}
	out := &pb.Agent{}
	out.DisplayName = in.DisplayName
	out.DefaultLanguageCode = in.DefaultLanguageCode
	out.SupportedLanguageCodes = in.SupportedLanguageCodes
	out.TimeZone = in.TimeZone
	out.Description = direct.ValueOf(in.Description)
	out.AvatarUri = direct.ValueOf(in.AvatarURI)
	out.EnableLogging = direct.ValueOf(in.EnableLogging)
	out.MatchMode = direct.Enum_ToProto[pb.Agent_MatchMode](mapCtx, in.MatchMode)
	out.ClassificationThreshold = direct.ValueOf(in.ClassificationThreshold)
	out.ApiVersion = direct.Enum_ToProto[pb.Agent_ApiVersion](mapCtx, in.APIVersion)
	out.Tier = direct.Enum_ToProto[pb.Agent_Tier](mapCtx, in.Tier)
	return out
}
