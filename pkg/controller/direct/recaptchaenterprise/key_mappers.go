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

package recaptchaenterprise

import (
	pb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/recaptchaenterprise/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func RecaptchaEnterpriseKeySpec_DisplayName_ToProto(mapCtx *direct.MapContext, in string) string {
	return in
}

func KeyTestingOptions_FromProto(mapCtx *direct.MapContext, in *pb.TestingOptions) *krm.KeyTestingOptions {
	if in == nil {
		return nil
	}
	out := &krm.KeyTestingOptions{}
	if in.TestingScore != 0 {
		val := float64(in.GetTestingScore())
		out.TestingScore = &val
	}
	out.TestingChallenge = direct.Enum_FromProto(mapCtx, in.GetTestingChallenge())
	return out
}

func KeyTestingOptions_ToProto(mapCtx *direct.MapContext, in *krm.KeyTestingOptions) *pb.TestingOptions {
	if in == nil {
		return nil
	}
	out := &pb.TestingOptions{}
	if in.TestingScore != nil {
		out.TestingScore = float32(*in.TestingScore)
	}
	out.TestingChallenge = direct.Enum_ToProto[pb.TestingOptions_TestingChallenge](mapCtx, in.TestingChallenge)
	return out
}

func KeyWafSettings_FromProto(mapCtx *direct.MapContext, in *pb.WafSettings) *krm.KeyWafSettings {
	if in == nil {
		return nil
	}
	out := &krm.KeyWafSettings{}
	if val := direct.Enum_FromProto(mapCtx, in.GetWafService()); val != nil {
		out.WafService = *val
	}
	if val := direct.Enum_FromProto(mapCtx, in.GetWafFeature()); val != nil {
		out.WafFeature = *val
	}
	return out
}

func KeyWafSettings_ToProto(mapCtx *direct.MapContext, in *krm.KeyWafSettings) *pb.WafSettings {
	if in == nil {
		return nil
	}
	out := &pb.WafSettings{}
	out.WafService = direct.Enum_ToProto[pb.WafSettings_WafService](mapCtx, &in.WafService)
	out.WafFeature = direct.Enum_ToProto[pb.WafSettings_WafFeature](mapCtx, &in.WafFeature)
	return out
}

func KeyWebSettings_FromProto(mapCtx *direct.MapContext, in *pb.WebKeySettings) *krm.KeyWebSettings {
	if in == nil {
		return nil
	}
	out := &krm.KeyWebSettings{}
	out.AllowAllDomains = direct.LazyPtr(in.GetAllowAllDomains())
	out.AllowedDomains = in.AllowedDomains
	out.AllowAmpTraffic = direct.LazyPtr(in.GetAllowAmpTraffic())
	if val := direct.Enum_FromProto(mapCtx, in.GetIntegrationType()); val != nil {
		out.IntegrationType = *val
	}
	out.ChallengeSecurityPreference = direct.Enum_FromProto(mapCtx, in.GetChallengeSecurityPreference())
	return out
}

func KeyWebSettings_ToProto(mapCtx *direct.MapContext, in *krm.KeyWebSettings) *pb.WebKeySettings {
	if in == nil {
		return nil
	}
	out := &pb.WebKeySettings{}
	out.AllowAllDomains = direct.ValueOf(in.AllowAllDomains)
	out.AllowedDomains = in.AllowedDomains
	out.AllowAmpTraffic = direct.ValueOf(in.AllowAmpTraffic)
	out.IntegrationType = direct.Enum_ToProto[pb.WebKeySettings_IntegrationType](mapCtx, &in.IntegrationType)
	out.ChallengeSecurityPreference = direct.Enum_ToProto[pb.WebKeySettings_ChallengeSecurityPreference](mapCtx, in.ChallengeSecurityPreference)
	return out
}
