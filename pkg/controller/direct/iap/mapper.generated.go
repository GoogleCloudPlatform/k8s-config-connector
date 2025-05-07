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

// +generated:mapper
// krm.group: iap.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.iap.v1

package iap

import (
	pb "cloud.google.com/go/iap/apiv1/iappb"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iap/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AccessDeniedPageSettings_FromProto(mapCtx *direct.MapContext, in *pb.AccessDeniedPageSettings) *krmv1beta1.AccessDeniedPageSettings {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AccessDeniedPageSettings{}
	out.AccessDeniedPageURI = direct.StringValue_FromProto(mapCtx, in.GetAccessDeniedPageUri())
	out.GenerateTroubleshootingURI = direct.BoolValue_FromProto(mapCtx, in.GetGenerateTroubleshootingUri())
	out.RemediationTokenGenerationEnabled = direct.BoolValue_FromProto(mapCtx, in.GetRemediationTokenGenerationEnabled())
	return out
}
func AccessSettings_FromProto(mapCtx *direct.MapContext, in *pb.AccessSettings) *krmv1beta1.AccessSettings {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AccessSettings{}
	out.GcipSettings = GcipSettings_FromProto(mapCtx, in.GetGcipSettings())
	out.CorsSettings = CorsSettings_FromProto(mapCtx, in.GetCorsSettings())
	out.OauthSettings = OAuthSettings_FromProto(mapCtx, in.GetOauthSettings())
	out.ReauthSettings = ReauthSettings_FromProto(mapCtx, in.GetReauthSettings())
	out.AllowedDomainsSettings = AllowedDomainsSettings_FromProto(mapCtx, in.GetAllowedDomainsSettings())
	return out
}
func AccessSettings_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AccessSettings) *pb.AccessSettings {
	if in == nil {
		return nil
	}
	out := &pb.AccessSettings{}
	out.GcipSettings = GcipSettings_ToProto(mapCtx, in.GcipSettings)
	out.CorsSettings = CorsSettings_ToProto(mapCtx, in.CorsSettings)
	out.OauthSettings = OAuthSettings_ToProto(mapCtx, in.OauthSettings)
	out.ReauthSettings = ReauthSettings_ToProto(mapCtx, in.ReauthSettings)
	out.AllowedDomainsSettings = AllowedDomainsSettings_ToProto(mapCtx, in.AllowedDomainsSettings)
	return out
}
func AllowedDomainsSettings_FromProto(mapCtx *direct.MapContext, in *pb.AllowedDomainsSettings) *krmv1beta1.AllowedDomainsSettings {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AllowedDomainsSettings{}
	out.Enable = in.Enable
	out.Domains = in.Domains
	return out
}
func AllowedDomainsSettings_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AllowedDomainsSettings) *pb.AllowedDomainsSettings {
	if in == nil {
		return nil
	}
	out := &pb.AllowedDomainsSettings{}
	out.Enable = in.Enable
	out.Domains = in.Domains
	return out
}
func ApplicationSettings_FromProto(mapCtx *direct.MapContext, in *pb.ApplicationSettings) *krmv1beta1.ApplicationSettings {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ApplicationSettings{}
	out.CsmSettings = CsmSettings_FromProto(mapCtx, in.GetCsmSettings())
	out.AccessDeniedPageSettings = AccessDeniedPageSettings_FromProto(mapCtx, in.GetAccessDeniedPageSettings())
	out.CookieDomain = direct.StringValue_FromProto(mapCtx, in.GetCookieDomain())
	out.AttributePropagationSettings = AttributePropagationSettings_FromProto(mapCtx, in.GetAttributePropagationSettings())
	return out
}
func ApplicationSettings_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ApplicationSettings) *pb.ApplicationSettings {
	if in == nil {
		return nil
	}
	out := &pb.ApplicationSettings{}
	out.CsmSettings = CsmSettings_ToProto(mapCtx, in.CsmSettings)
	out.AccessDeniedPageSettings = AccessDeniedPageSettings_ToProto(mapCtx, in.AccessDeniedPageSettings)
	out.CookieDomain = direct.StringValue_ToProto(mapCtx, in.CookieDomain)
	out.AttributePropagationSettings = AttributePropagationSettings_ToProto(mapCtx, in.AttributePropagationSettings)
	return out
}
func AttributePropagationSettings_FromProto(mapCtx *direct.MapContext, in *pb.AttributePropagationSettings) *krmv1beta1.AttributePropagationSettings {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AttributePropagationSettings{}
	out.Expression = in.Expression
	out.OutputCredentials = direct.EnumSlice_FromProto(mapCtx, in.OutputCredentials)
	out.Enable = in.Enable
	return out
}
func AttributePropagationSettings_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AttributePropagationSettings) *pb.AttributePropagationSettings {
	if in == nil {
		return nil
	}
	out := &pb.AttributePropagationSettings{}
	out.Expression = in.Expression
	out.OutputCredentials = direct.EnumSlice_ToProto[pb.AttributePropagationSettings_OutputCredentials](mapCtx, in.OutputCredentials)
	out.Enable = in.Enable
	return out
}
func CorsSettings_FromProto(mapCtx *direct.MapContext, in *pb.CorsSettings) *krmv1beta1.CorsSettings {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.CorsSettings{}
	out.AllowHTTPOptions = direct.BoolValue_FromProto(mapCtx, in.GetAllowHttpOptions())
	return out
}
func CorsSettings_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.CorsSettings) *pb.CorsSettings {
	if in == nil {
		return nil
	}
	out := &pb.CorsSettings{}
	out.AllowHttpOptions = direct.BoolValue_ToProto(mapCtx, in.AllowHTTPOptions)
	return out
}
func CsmSettings_FromProto(mapCtx *direct.MapContext, in *pb.CsmSettings) *krmv1beta1.CsmSettings {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.CsmSettings{}
	out.RctokenAud = direct.StringValue_FromProto(mapCtx, in.GetRctokenAud())
	return out
}
func CsmSettings_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.CsmSettings) *pb.CsmSettings {
	if in == nil {
		return nil
	}
	out := &pb.CsmSettings{}
	out.RctokenAud = direct.StringValue_ToProto(mapCtx, in.RctokenAud)
	return out
}
func IAPSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.IapSettings) *krmv1beta1.IAPSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.IAPSettingsSpec{}
	// MISSING: Name
	out.AccessSettings = AccessSettings_FromProto(mapCtx, in.GetAccessSettings())
	out.ApplicationSettings = ApplicationSettings_FromProto(mapCtx, in.GetApplicationSettings())
	return out
}
func IAPSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.IAPSettingsSpec) *pb.IapSettings {
	if in == nil {
		return nil
	}
	out := &pb.IapSettings{}
	// MISSING: Name
	out.AccessSettings = AccessSettings_ToProto(mapCtx, in.AccessSettings)
	out.ApplicationSettings = ApplicationSettings_ToProto(mapCtx, in.ApplicationSettings)
	return out
}
func OAuthSettings_FromProto(mapCtx *direct.MapContext, in *pb.OAuthSettings) *krmv1beta1.OAuthSettings {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.OAuthSettings{}
	out.LoginHint = direct.StringValue_FromProto(mapCtx, in.GetLoginHint())
	out.ProgrammaticClients = in.ProgrammaticClients
	return out
}
func OAuthSettings_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.OAuthSettings) *pb.OAuthSettings {
	if in == nil {
		return nil
	}
	out := &pb.OAuthSettings{}
	out.LoginHint = direct.StringValue_ToProto(mapCtx, in.LoginHint)
	out.ProgrammaticClients = in.ProgrammaticClients
	return out
}
func ReauthSettings_FromProto(mapCtx *direct.MapContext, in *pb.ReauthSettings) *krmv1beta1.ReauthSettings {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ReauthSettings{}
	out.Method = direct.Enum_FromProto(mapCtx, in.GetMethod())
	out.MaxAge = direct.StringDuration_FromProto(mapCtx, in.GetMaxAge())
	out.PolicyType = direct.Enum_FromProto(mapCtx, in.GetPolicyType())
	return out
}
func ReauthSettings_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ReauthSettings) *pb.ReauthSettings {
	if in == nil {
		return nil
	}
	out := &pb.ReauthSettings{}
	out.Method = direct.Enum_ToProto[pb.ReauthSettings_Method](mapCtx, in.Method)
	out.MaxAge = direct.StringDuration_ToProto(mapCtx, in.MaxAge)
	out.PolicyType = direct.Enum_ToProto[pb.ReauthSettings_PolicyType](mapCtx, in.PolicyType)
	return out
}
