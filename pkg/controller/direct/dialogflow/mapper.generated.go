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

package dialogflow

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Webhook_FromProto(mapCtx *direct.MapContext, in *pb.Webhook) *krm.Webhook {
	if in == nil {
		return nil
	}
	out := &krm.Webhook{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.GenericWebService = Webhook_GenericWebService_FromProto(mapCtx, in.GetGenericWebService())
	out.ServiceDirectory = Webhook_ServiceDirectoryConfig_FromProto(mapCtx, in.GetServiceDirectory())
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}
func Webhook_ToProto(mapCtx *direct.MapContext, in *krm.Webhook) *pb.Webhook {
	if in == nil {
		return nil
	}
	out := &pb.Webhook{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if oneof := Webhook_GenericWebService_ToProto(mapCtx, in.GenericWebService); oneof != nil {
		out.Webhook = &pb.Webhook_GenericWebService_{GenericWebService: oneof}
	}
	if oneof := Webhook_ServiceDirectoryConfig_ToProto(mapCtx, in.ServiceDirectory); oneof != nil {
		out.Webhook = &pb.Webhook_ServiceDirectory{ServiceDirectory: oneof}
	}
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.Disabled = direct.ValueOf(in.Disabled)
	return out
}
func Webhook_GenericWebService_FromProto(mapCtx *direct.MapContext, in *pb.Webhook_GenericWebService) *krm.Webhook_GenericWebService {
	if in == nil {
		return nil
	}
	out := &krm.Webhook_GenericWebService{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.RequestHeaders = in.RequestHeaders
	out.AllowedCaCerts = in.AllowedCaCerts
	out.OauthConfig = Webhook_GenericWebService_OAuthConfig_FromProto(mapCtx, in.GetOauthConfig())
	out.ServiceAgentAuth = direct.Enum_FromProto(mapCtx, in.GetServiceAgentAuth())
	out.WebhookType = direct.Enum_FromProto(mapCtx, in.GetWebhookType())
	out.HTTPMethod = direct.Enum_FromProto(mapCtx, in.GetHttpMethod())
	out.RequestBody = direct.LazyPtr(in.GetRequestBody())
	out.ParameterMapping = in.ParameterMapping
	return out
}
func Webhook_GenericWebService_ToProto(mapCtx *direct.MapContext, in *krm.Webhook_GenericWebService) *pb.Webhook_GenericWebService {
	if in == nil {
		return nil
	}
	out := &pb.Webhook_GenericWebService{}
	out.Uri = direct.ValueOf(in.URI)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.RequestHeaders = in.RequestHeaders
	out.AllowedCaCerts = in.AllowedCaCerts
	out.OauthConfig = Webhook_GenericWebService_OAuthConfig_ToProto(mapCtx, in.OauthConfig)
	out.ServiceAgentAuth = direct.Enum_ToProto[pb.Webhook_GenericWebService_ServiceAgentAuth](mapCtx, in.ServiceAgentAuth)
	out.WebhookType = direct.Enum_ToProto[pb.Webhook_GenericWebService_WebhookType](mapCtx, in.WebhookType)
	out.HttpMethod = direct.Enum_ToProto[pb.Webhook_GenericWebService_HttpMethod](mapCtx, in.HTTPMethod)
	out.RequestBody = direct.ValueOf(in.RequestBody)
	out.ParameterMapping = in.ParameterMapping
	return out
}
func Webhook_GenericWebService_OAuthConfig_FromProto(mapCtx *direct.MapContext, in *pb.Webhook_GenericWebService_OAuthConfig) *krm.Webhook_GenericWebService_OAuthConfig {
	if in == nil {
		return nil
	}
	out := &krm.Webhook_GenericWebService_OAuthConfig{}
	out.ClientID = direct.LazyPtr(in.GetClientId())
	out.ClientSecret = direct.LazyPtr(in.GetClientSecret())
	out.TokenEndpoint = direct.LazyPtr(in.GetTokenEndpoint())
	out.Scopes = in.Scopes
	return out
}
func Webhook_GenericWebService_OAuthConfig_ToProto(mapCtx *direct.MapContext, in *krm.Webhook_GenericWebService_OAuthConfig) *pb.Webhook_GenericWebService_OAuthConfig {
	if in == nil {
		return nil
	}
	out := &pb.Webhook_GenericWebService_OAuthConfig{}
	out.ClientId = direct.ValueOf(in.ClientID)
	out.ClientSecret = direct.ValueOf(in.ClientSecret)
	out.TokenEndpoint = direct.ValueOf(in.TokenEndpoint)
	out.Scopes = in.Scopes
	return out
}
func Webhook_ServiceDirectoryConfig_FromProto(mapCtx *direct.MapContext, in *pb.Webhook_ServiceDirectoryConfig) *krm.Webhook_ServiceDirectoryConfig {
	if in == nil {
		return nil
	}
	out := &krm.Webhook_ServiceDirectoryConfig{}
	out.Service = direct.LazyPtr(in.GetService())
	out.GenericWebService = Webhook_GenericWebService_FromProto(mapCtx, in.GetGenericWebService())
	return out
}
func Webhook_ServiceDirectoryConfig_ToProto(mapCtx *direct.MapContext, in *krm.Webhook_ServiceDirectoryConfig) *pb.Webhook_ServiceDirectoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Webhook_ServiceDirectoryConfig{}
	out.Service = direct.ValueOf(in.Service)
	out.GenericWebService = Webhook_GenericWebService_ToProto(mapCtx, in.GenericWebService)
	return out
}
