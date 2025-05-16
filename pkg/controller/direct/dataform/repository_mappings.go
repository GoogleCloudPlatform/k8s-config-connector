// Copyright 2024 Google LLC
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

package dataform

import (
	pb "cloud.google.com/go/dataform/apiv1beta1/dataformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataformRepositorySpec_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.DataformRepositorySpec {
	if in == nil {
		return nil
	}
	out := &krm.DataformRepositorySpec{}
	out.GitRemoteSettings = RepositoryGitRemoteSettings_FromProto(mapCtx, in.GetGitRemoteSettings())
	out.WorkspaceCompilationOverrides = RepositoryWorkspaceCompilationOverrides_FromProto(mapCtx, in.GetWorkspaceCompilationOverrides())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.GitRemoteSettings = RepositoryGitRemoteSettings_FromProto(mapCtx, in.GetGitRemoteSettings())

	if in.GetNpmrcEnvironmentVariablesSecretVersion() != "" {
		out.NpmrcEnvironmentVariablesSecretVersionRef = &secretmanagerv1beta1.SecretVersionRef{
			External: in.GetNpmrcEnvironmentVariablesSecretVersion(),
		}
	}

	out.WorkspaceCompilationOverrides = RepositoryWorkspaceCompilationOverrides_FromProto(mapCtx, in.GetWorkspaceCompilationOverrides())
	out.WorkspaceCompilationOverrides = RepositoryWorkspaceCompilationOverrides_FromProto(mapCtx, in.GetWorkspaceCompilationOverrides())
	out.SetAuthenticatedUserAdmin = in.GetSetAuthenticatedUserAdmin()

	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refs.IAMServiceAccountRef{
			External: in.GetServiceAccount(),
		}
	}

	return out
}

func DataformRepositorySpec_ToProto(mapCtx *direct.MapContext, in *krm.DataformRepositorySpec) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.GitRemoteSettings = RepositoryGitRemoteSettings_ToProto(mapCtx, in.GitRemoteSettings)
	out.WorkspaceCompilationOverrides = RepositoryWorkspaceCompilationOverrides_ToProto(mapCtx, in.WorkspaceCompilationOverrides)

	if in.NpmrcEnvironmentVariablesSecretVersionRef != nil {
		out.NpmrcEnvironmentVariablesSecretVersion = in.NpmrcEnvironmentVariablesSecretVersionRef.External
	}

	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.SetAuthenticatedUserAdmin = in.SetAuthenticatedUserAdmin

	return out
}

func RepositoryGitRemoteSettings_FromProto(mapCtx *direct.MapContext, in *pb.Repository_GitRemoteSettings) *krm.RepositoryGitRemoteSettings {
	if in == nil {
		return nil
	}
	out := &krm.RepositoryGitRemoteSettings{}
	out.Url = in.GetUrl()
	out.DefaultBranch = in.GetDefaultBranch()
	if in.GetAuthenticationTokenSecretVersion() != "" {
		out.AuthenticationTokenSecretVersionRef = &secretmanagerv1beta1.SecretVersionRef{
			External: in.GetAuthenticationTokenSecretVersion(),
		}
	}

	if inSshConfig := in.GetSshAuthenticationConfig(); inSshConfig != nil {
		out.SSHAuthenticationConfig = &krm.SSHAuthenticationConfig{}
		if inSshConfig.GetUserPrivateKeySecretVersion() != "" {
			out.SSHAuthenticationConfig.UserPrivateKeySecretVersionRef = &secretmanagerv1beta1.SecretVersionRef{
				External: inSshConfig.GetUserPrivateKeySecretVersion(),
			}
		}

		out.SSHAuthenticationConfig.HostPublicKey = inSshConfig.HostPublicKey
	}

	return out
}

func RepositoryGitRemoteSettings_ToProto(mapCtx *direct.MapContext, in *krm.RepositoryGitRemoteSettings) *pb.Repository_GitRemoteSettings {
	if in == nil {
		return nil
	}
	out := &pb.Repository_GitRemoteSettings{}
	out.Url = in.Url
	out.DefaultBranch = in.DefaultBranch

	if in.AuthenticationTokenSecretVersionRef != nil {
		out.AuthenticationTokenSecretVersion = in.AuthenticationTokenSecretVersionRef.External
	}

	if in.SSHAuthenticationConfig != nil {
		out.SshAuthenticationConfig = &pb.Repository_GitRemoteSettings_SshAuthenticationConfig{}

		if in.SSHAuthenticationConfig.UserPrivateKeySecretVersionRef != nil {
			out.SshAuthenticationConfig.UserPrivateKeySecretVersion = in.SSHAuthenticationConfig.UserPrivateKeySecretVersionRef.External
		}

		out.SshAuthenticationConfig.HostPublicKey = in.SSHAuthenticationConfig.HostPublicKey
	}

	return out
}

func RepositoryWorkspaceCompilationOverrides_FromProto(mapCtx *direct.MapContext, in *pb.Repository_WorkspaceCompilationOverrides) *krm.RepositoryWorkspaceCompilationOverrides {
	if in == nil {
		return nil
	}
	out := &krm.RepositoryWorkspaceCompilationOverrides{}
	out.DefaultDatabase = direct.LazyPtr(in.GetDefaultDatabase())
	out.SchemaSuffix = direct.LazyPtr(in.GetSchemaSuffix())
	out.TablePrefix = direct.LazyPtr(in.GetTablePrefix())
	return out
}

func RepositoryWorkspaceCompilationOverrides_ToProto(mapCtx *direct.MapContext, in *krm.RepositoryWorkspaceCompilationOverrides) *pb.Repository_WorkspaceCompilationOverrides {
	if in == nil {
		return nil
	}
	out := &pb.Repository_WorkspaceCompilationOverrides{}
	out.DefaultDatabase = direct.ValueOf(in.DefaultDatabase)
	out.SchemaSuffix = direct.ValueOf(in.SchemaSuffix)
	out.TablePrefix = direct.ValueOf(in.TablePrefix)
	return out
}
