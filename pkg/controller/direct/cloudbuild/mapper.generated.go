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

package cloudbuild

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1alpha1"
)
func CloudBuildWorkerPoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildWorkerPoolObservedState) *pb.WorkerPool {
	if in == nil {
		return nil
	}
	out := &pb.WorkerPool{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Annotations
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: DeleteTime
	// MISSING: State
	// MISSING: PrivatePoolV1Config
	// MISSING: Etag
	// (near miss): "Etag" vs "ETag"
	return out
}
func CloudbuildGitHubEnterpriseConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GitHubEnterpriseConfig) *krm.CloudbuildGitHubEnterpriseConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudbuildGitHubEnterpriseConfigObservedState{}
	// MISSING: Name
	// MISSING: HostURL
	// MISSING: AppID
	// MISSING: CreateTime
	// MISSING: WebhookKey
	// MISSING: PeeredNetwork
	// MISSING: Secrets
	// MISSING: DisplayName
	// MISSING: SslCa
	return out
}
func CloudbuildGitHubEnterpriseConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudbuildGitHubEnterpriseConfigObservedState) *pb.GitHubEnterpriseConfig {
	if in == nil {
		return nil
	}
	out := &pb.GitHubEnterpriseConfig{}
	// MISSING: Name
	// MISSING: HostURL
	// MISSING: AppID
	// MISSING: CreateTime
	// MISSING: WebhookKey
	// MISSING: PeeredNetwork
	// MISSING: Secrets
	// MISSING: DisplayName
	// MISSING: SslCa
	return out
}
func CloudbuildGitHubEnterpriseConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.GitHubEnterpriseConfig) *krm.CloudbuildGitHubEnterpriseConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudbuildGitHubEnterpriseConfigSpec{}
	// MISSING: Name
	// MISSING: HostURL
	// MISSING: AppID
	// MISSING: CreateTime
	// MISSING: WebhookKey
	// MISSING: PeeredNetwork
	// MISSING: Secrets
	// MISSING: DisplayName
	// MISSING: SslCa
	return out
}
func CloudbuildGitHubEnterpriseConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudbuildGitHubEnterpriseConfigSpec) *pb.GitHubEnterpriseConfig {
	if in == nil {
		return nil
	}
	out := &pb.GitHubEnterpriseConfig{}
	// MISSING: Name
	// MISSING: HostURL
	// MISSING: AppID
	// MISSING: CreateTime
	// MISSING: WebhookKey
	// MISSING: PeeredNetwork
	// MISSING: Secrets
	// MISSING: DisplayName
	// MISSING: SslCa
	return out
}
func GitHubEnterpriseConfig_FromProto(mapCtx *direct.MapContext, in *pb.GitHubEnterpriseConfig) *krm.GitHubEnterpriseConfig {
	if in == nil {
		return nil
	}
	out := &krm.GitHubEnterpriseConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.HostURL = direct.LazyPtr(in.GetHostUrl())
	out.AppID = direct.LazyPtr(in.GetAppId())
	// MISSING: CreateTime
	out.WebhookKey = direct.LazyPtr(in.GetWebhookKey())
	out.PeeredNetwork = direct.LazyPtr(in.GetPeeredNetwork())
	out.Secrets = GitHubEnterpriseSecrets_FromProto(mapCtx, in.GetSecrets())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.SslCa = direct.LazyPtr(in.GetSslCa())
	return out
}
func GitHubEnterpriseConfig_ToProto(mapCtx *direct.MapContext, in *krm.GitHubEnterpriseConfig) *pb.GitHubEnterpriseConfig {
	if in == nil {
		return nil
	}
	out := &pb.GitHubEnterpriseConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.HostUrl = direct.ValueOf(in.HostURL)
	out.AppId = direct.ValueOf(in.AppID)
	// MISSING: CreateTime
	out.WebhookKey = direct.ValueOf(in.WebhookKey)
	out.PeeredNetwork = direct.ValueOf(in.PeeredNetwork)
	out.Secrets = GitHubEnterpriseSecrets_ToProto(mapCtx, in.Secrets)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.SslCa = direct.ValueOf(in.SslCa)
	return out
}
func GitHubEnterpriseConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GitHubEnterpriseConfig) *krm.GitHubEnterpriseConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GitHubEnterpriseConfigObservedState{}
	// MISSING: Name
	// MISSING: HostURL
	// MISSING: AppID
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: WebhookKey
	// MISSING: PeeredNetwork
	// MISSING: Secrets
	// MISSING: DisplayName
	// MISSING: SslCa
	return out
}
func GitHubEnterpriseConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GitHubEnterpriseConfigObservedState) *pb.GitHubEnterpriseConfig {
	if in == nil {
		return nil
	}
	out := &pb.GitHubEnterpriseConfig{}
	// MISSING: Name
	// MISSING: HostURL
	// MISSING: AppID
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: WebhookKey
	// MISSING: PeeredNetwork
	// MISSING: Secrets
	// MISSING: DisplayName
	// MISSING: SslCa
	return out
}
func GitHubEnterpriseSecrets_FromProto(mapCtx *direct.MapContext, in *pb.GitHubEnterpriseSecrets) *krm.GitHubEnterpriseSecrets {
	if in == nil {
		return nil
	}
	out := &krm.GitHubEnterpriseSecrets{}
	out.PrivateKeyVersionName = direct.LazyPtr(in.GetPrivateKeyVersionName())
	out.WebhookSecretVersionName = direct.LazyPtr(in.GetWebhookSecretVersionName())
	out.OauthSecretVersionName = direct.LazyPtr(in.GetOauthSecretVersionName())
	out.OauthClientIDVersionName = direct.LazyPtr(in.GetOauthClientIdVersionName())
	return out
}
func GitHubEnterpriseSecrets_ToProto(mapCtx *direct.MapContext, in *krm.GitHubEnterpriseSecrets) *pb.GitHubEnterpriseSecrets {
	if in == nil {
		return nil
	}
	out := &pb.GitHubEnterpriseSecrets{}
	out.PrivateKeyVersionName = direct.ValueOf(in.PrivateKeyVersionName)
	out.WebhookSecretVersionName = direct.ValueOf(in.WebhookSecretVersionName)
	out.OauthSecretVersionName = direct.ValueOf(in.OauthSecretVersionName)
	out.OauthClientIdVersionName = direct.ValueOf(in.OauthClientIDVersionName)
	return out
}
func PrivatePoolV1Config_NetworkConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.PrivatePoolV1Config_NetworkConfig) *krm.PrivatePoolV1Config_NetworkConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.PrivatePoolV1Config_NetworkConfigSpec{}
	if in.GetPeeredNetwork() != "" {
		out.PeeredNetworkRef = &refs.refv1beta1.ComputeNetworkRef{External: in.GetPeeredNetwork()}
	}
	out.EgressOption = direct.Enum_FromProto(mapCtx, in.GetEgressOption())
	out.PeeredNetworkIPRange = direct.LazyPtr(in.GetPeeredNetworkIpRange())
	return out
}
func PrivatePoolV1Config_NetworkConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.PrivatePoolV1Config_NetworkConfigSpec) *pb.PrivatePoolV1Config_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrivatePoolV1Config_NetworkConfig{}
	if in.PeeredNetworkRef != nil {
		out.PeeredNetwork = in.PeeredNetworkRef.External
	}
	out.EgressOption = direct.Enum_ToProto[pb.PrivatePoolV1Config_NetworkConfig_EgressOption](mapCtx, in.EgressOption)
	out.PeeredNetworkIpRange = direct.ValueOf(in.PeeredNetworkIPRange)
	return out
}
func PrivatePoolV1Config_NetworkConfigStatus_ToProto(mapCtx *direct.MapContext, in *krm.PrivatePoolV1Config_NetworkConfigStatus) *pb.PrivatePoolV1Config_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrivatePoolV1Config_NetworkConfig{}
	out.PeeredNetwork = direct.ValueOf(in.PeeredNetwork)
	out.EgressOption = direct.Enum_ToProto[pb.PrivatePoolV1Config_NetworkConfig_EgressOption](mapCtx, in.EgressOption)
	out.PeeredNetworkIpRange = direct.ValueOf(in.PeeredNetworkIPRange)
	return out
}
