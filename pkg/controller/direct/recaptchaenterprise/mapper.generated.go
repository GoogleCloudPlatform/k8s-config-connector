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

package recaptchaenterprise

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/recaptchaenterprise/v1alpha1"
)
func AndroidKeySettings_FromProto(mapCtx *direct.MapContext, in *pb.AndroidKeySettings) *krm.AndroidKeySettings {
	if in == nil {
		return nil
	}
	out := &krm.AndroidKeySettings{}
	out.AllowAllPackageNames = direct.LazyPtr(in.GetAllowAllPackageNames())
	out.AllowedPackageNames = in.AllowedPackageNames
	out.SupportNonGoogleAppStoreDistribution = direct.LazyPtr(in.GetSupportNonGoogleAppStoreDistribution())
	return out
}
func AndroidKeySettings_ToProto(mapCtx *direct.MapContext, in *krm.AndroidKeySettings) *pb.AndroidKeySettings {
	if in == nil {
		return nil
	}
	out := &pb.AndroidKeySettings{}
	out.AllowAllPackageNames = direct.ValueOf(in.AllowAllPackageNames)
	out.AllowedPackageNames = in.AllowedPackageNames
	out.SupportNonGoogleAppStoreDistribution = direct.ValueOf(in.SupportNonGoogleAppStoreDistribution)
	return out
}
func AppleDeveloperId_FromProto(mapCtx *direct.MapContext, in *pb.AppleDeveloperId) *krm.AppleDeveloperId {
	if in == nil {
		return nil
	}
	out := &krm.AppleDeveloperId{}
	out.PrivateKey = direct.LazyPtr(in.GetPrivateKey())
	out.KeyID = direct.LazyPtr(in.GetKeyId())
	out.TeamID = direct.LazyPtr(in.GetTeamId())
	return out
}
func AppleDeveloperId_ToProto(mapCtx *direct.MapContext, in *krm.AppleDeveloperId) *pb.AppleDeveloperId {
	if in == nil {
		return nil
	}
	out := &pb.AppleDeveloperId{}
	out.PrivateKey = direct.ValueOf(in.PrivateKey)
	out.KeyId = direct.ValueOf(in.KeyID)
	out.TeamId = direct.ValueOf(in.TeamID)
	return out
}
func ExpressKeySettings_FromProto(mapCtx *direct.MapContext, in *pb.ExpressKeySettings) *krm.ExpressKeySettings {
	if in == nil {
		return nil
	}
	out := &krm.ExpressKeySettings{}
	return out
}
func ExpressKeySettings_ToProto(mapCtx *direct.MapContext, in *krm.ExpressKeySettings) *pb.ExpressKeySettings {
	if in == nil {
		return nil
	}
	out := &pb.ExpressKeySettings{}
	return out
}
func IOSKeySettings_FromProto(mapCtx *direct.MapContext, in *pb.IOSKeySettings) *krm.IOSKeySettings {
	if in == nil {
		return nil
	}
	out := &krm.IOSKeySettings{}
	out.AllowAllBundleIds = direct.LazyPtr(in.GetAllowAllBundleIds())
	out.AllowedBundleIds = in.AllowedBundleIds
	out.AppleDeveloperID = AppleDeveloperId_FromProto(mapCtx, in.GetAppleDeveloperId())
	return out
}
func IOSKeySettings_ToProto(mapCtx *direct.MapContext, in *krm.IOSKeySettings) *pb.IOSKeySettings {
	if in == nil {
		return nil
	}
	out := &pb.IOSKeySettings{}
	out.AllowAllBundleIds = direct.ValueOf(in.AllowAllBundleIds)
	out.AllowedBundleIds = in.AllowedBundleIds
	out.AppleDeveloperId = AppleDeveloperId_ToProto(mapCtx, in.AppleDeveloperID)
	return out
}
func Key_FromProto(mapCtx *direct.MapContext, in *pb.Key) *krm.Key {
	if in == nil {
		return nil
	}
	out := &krm.Key{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.WebSettings = WebKeySettings_FromProto(mapCtx, in.GetWebSettings())
	out.AndroidSettings = AndroidKeySettings_FromProto(mapCtx, in.GetAndroidSettings())
	out.IosSettings = IOSKeySettings_FromProto(mapCtx, in.GetIosSettings())
	out.ExpressSettings = ExpressKeySettings_FromProto(mapCtx, in.GetExpressSettings())
	out.Labels = in.Labels
	// MISSING: CreateTime
	out.TestingOptions = TestingOptions_FromProto(mapCtx, in.GetTestingOptions())
	out.WafSettings = WafSettings_FromProto(mapCtx, in.GetWafSettings())
	return out
}
func Key_ToProto(mapCtx *direct.MapContext, in *krm.Key) *pb.Key {
	if in == nil {
		return nil
	}
	out := &pb.Key{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if oneof := WebKeySettings_ToProto(mapCtx, in.WebSettings); oneof != nil {
		out.PlatformSettings = &pb.Key_WebSettings{WebSettings: oneof}
	}
	if oneof := AndroidKeySettings_ToProto(mapCtx, in.AndroidSettings); oneof != nil {
		out.PlatformSettings = &pb.Key_AndroidSettings{AndroidSettings: oneof}
	}
	if oneof := IOSKeySettings_ToProto(mapCtx, in.IosSettings); oneof != nil {
		out.PlatformSettings = &pb.Key_IosSettings{IosSettings: oneof}
	}
	if oneof := ExpressKeySettings_ToProto(mapCtx, in.ExpressSettings); oneof != nil {
		out.PlatformSettings = &pb.Key_ExpressSettings{ExpressSettings: oneof}
	}
	out.Labels = in.Labels
	// MISSING: CreateTime
	out.TestingOptions = TestingOptions_ToProto(mapCtx, in.TestingOptions)
	out.WafSettings = WafSettings_ToProto(mapCtx, in.WafSettings)
	return out
}
func KeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Key) *krm.KeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KeyObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: WebSettings
	// MISSING: AndroidSettings
	// MISSING: IosSettings
	// MISSING: ExpressSettings
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: TestingOptions
	// MISSING: WafSettings
	return out
}
func KeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KeyObservedState) *pb.Key {
	if in == nil {
		return nil
	}
	out := &pb.Key{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: WebSettings
	// MISSING: AndroidSettings
	// MISSING: IosSettings
	// MISSING: ExpressSettings
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: TestingOptions
	// MISSING: WafSettings
	return out
}
func RecaptchaenterpriseKeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Key) *krm.RecaptchaenterpriseKeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RecaptchaenterpriseKeyObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: WebSettings
	// MISSING: AndroidSettings
	// MISSING: IosSettings
	// MISSING: ExpressSettings
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: TestingOptions
	// MISSING: WafSettings
	return out
}
func RecaptchaenterpriseKeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RecaptchaenterpriseKeyObservedState) *pb.Key {
	if in == nil {
		return nil
	}
	out := &pb.Key{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: WebSettings
	// MISSING: AndroidSettings
	// MISSING: IosSettings
	// MISSING: ExpressSettings
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: TestingOptions
	// MISSING: WafSettings
	return out
}
func RecaptchaenterpriseKeySpec_FromProto(mapCtx *direct.MapContext, in *pb.Key) *krm.RecaptchaenterpriseKeySpec {
	if in == nil {
		return nil
	}
	out := &krm.RecaptchaenterpriseKeySpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: WebSettings
	// MISSING: AndroidSettings
	// MISSING: IosSettings
	// MISSING: ExpressSettings
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: TestingOptions
	// MISSING: WafSettings
	return out
}
func RecaptchaenterpriseKeySpec_ToProto(mapCtx *direct.MapContext, in *krm.RecaptchaenterpriseKeySpec) *pb.Key {
	if in == nil {
		return nil
	}
	out := &pb.Key{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: WebSettings
	// MISSING: AndroidSettings
	// MISSING: IosSettings
	// MISSING: ExpressSettings
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: TestingOptions
	// MISSING: WafSettings
	return out
}
func TestingOptions_FromProto(mapCtx *direct.MapContext, in *pb.TestingOptions) *krm.TestingOptions {
	if in == nil {
		return nil
	}
	out := &krm.TestingOptions{}
	out.TestingScore = direct.LazyPtr(in.GetTestingScore())
	out.TestingChallenge = direct.Enum_FromProto(mapCtx, in.GetTestingChallenge())
	return out
}
func TestingOptions_ToProto(mapCtx *direct.MapContext, in *krm.TestingOptions) *pb.TestingOptions {
	if in == nil {
		return nil
	}
	out := &pb.TestingOptions{}
	out.TestingScore = direct.ValueOf(in.TestingScore)
	out.TestingChallenge = direct.Enum_ToProto[pb.TestingOptions_TestingChallenge](mapCtx, in.TestingChallenge)
	return out
}
func WafSettings_FromProto(mapCtx *direct.MapContext, in *pb.WafSettings) *krm.WafSettings {
	if in == nil {
		return nil
	}
	out := &krm.WafSettings{}
	out.WafService = direct.Enum_FromProto(mapCtx, in.GetWafService())
	out.WafFeature = direct.Enum_FromProto(mapCtx, in.GetWafFeature())
	return out
}
func WafSettings_ToProto(mapCtx *direct.MapContext, in *krm.WafSettings) *pb.WafSettings {
	if in == nil {
		return nil
	}
	out := &pb.WafSettings{}
	out.WafService = direct.Enum_ToProto[pb.WafSettings_WafService](mapCtx, in.WafService)
	out.WafFeature = direct.Enum_ToProto[pb.WafSettings_WafFeature](mapCtx, in.WafFeature)
	return out
}
func WebKeySettings_FromProto(mapCtx *direct.MapContext, in *pb.WebKeySettings) *krm.WebKeySettings {
	if in == nil {
		return nil
	}
	out := &krm.WebKeySettings{}
	out.AllowAllDomains = direct.LazyPtr(in.GetAllowAllDomains())
	out.AllowedDomains = in.AllowedDomains
	out.AllowAmpTraffic = direct.LazyPtr(in.GetAllowAmpTraffic())
	out.IntegrationType = direct.Enum_FromProto(mapCtx, in.GetIntegrationType())
	out.ChallengeSecurityPreference = direct.Enum_FromProto(mapCtx, in.GetChallengeSecurityPreference())
	return out
}
func WebKeySettings_ToProto(mapCtx *direct.MapContext, in *krm.WebKeySettings) *pb.WebKeySettings {
	if in == nil {
		return nil
	}
	out := &pb.WebKeySettings{}
	out.AllowAllDomains = direct.ValueOf(in.AllowAllDomains)
	out.AllowedDomains = in.AllowedDomains
	out.AllowAmpTraffic = direct.ValueOf(in.AllowAmpTraffic)
	out.IntegrationType = direct.Enum_ToProto[pb.WebKeySettings_IntegrationType](mapCtx, in.IntegrationType)
	out.ChallengeSecurityPreference = direct.Enum_ToProto[pb.WebKeySettings_ChallengeSecurityPreference](mapCtx, in.ChallengeSecurityPreference)
	return out
}
