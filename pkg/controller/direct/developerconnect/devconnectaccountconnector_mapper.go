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

// +generated:mapper
// krm.group: developerconnect.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.developerconnect.v1

package developerconnect

import (
	pb "cloud.google.com/go/developerconnect/apiv1/developerconnectpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/developerconnect/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ProviderOAuthConfig_ToProto(mapCtx *direct.MapContext, in *krm.ProviderOAuthConfig) *pb.ProviderOAuthConfig {
	if in == nil {
		return nil
	}
	out := &pb.ProviderOAuthConfig{}
	out.Scopes = in.Scopes
	if in.SystemProviderID != nil {
		val := direct.Enum_ToProto[pb.SystemProvider](mapCtx, in.SystemProviderID)
		out.OauthProviderId = &pb.ProviderOAuthConfig_SystemProviderId{
			SystemProviderId: val,
		}
	}
	return out
}

func ProviderOAuthConfig_FromProto(mapCtx *direct.MapContext, in *pb.ProviderOAuthConfig) *krm.ProviderOAuthConfig {
	if in == nil {
		return nil
	}
	out := &krm.ProviderOAuthConfig{}
	out.Scopes = in.Scopes
	if in.GetSystemProviderId() != pb.SystemProvider_SYSTEM_PROVIDER_UNSPECIFIED {
		out.SystemProviderID = direct.Enum_FromProto(mapCtx, in.GetSystemProviderId())
	}
	return out
}

func DevConnectAccountConnectorSpec_ToProto(mapCtx *direct.MapContext, in *krm.DevConnectAccountConnectorSpec) *pb.AccountConnector {
	if in == nil {
		return nil
	}
	out := &pb.AccountConnector{}
	if in.ProviderOauthConfig != nil {
		out.AccountConnectorConfig = &pb.AccountConnector_ProviderOauthConfig{
			ProviderOauthConfig: ProviderOAuthConfig_ToProto(mapCtx, in.ProviderOauthConfig),
		}
	}
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	return out
}

func DevConnectAccountConnectorSpec_FromProto(mapCtx *direct.MapContext, in *pb.AccountConnector) *krm.DevConnectAccountConnectorSpec {
	if in == nil {
		return nil
	}
	out := &krm.DevConnectAccountConnectorSpec{}
	if config := in.GetProviderOauthConfig(); config != nil {
		out.ProviderOauthConfig = ProviderOAuthConfig_FromProto(mapCtx, config)
	}
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	return out
}
