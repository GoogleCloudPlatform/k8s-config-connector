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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1alpha1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataformRepositorySpec_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.DataformRepositorySpec {
	if in == nil {
		return nil
	}
	out := &krm.DataformRepositorySpec{}
	out.GitRemoteSettings = RepositoryGitRemoteSettings_FromProto(mapCtx, in.GetGitRemoteSettings())
	out.WorkspaceCompilationOverrides = RepositoryWorkspaceCompilationOverrides_FromProto(mapCtx, in.GetWorkspaceCompilationOverrides())

	return out
}
func DataformRepositorySpec_ToProto(mapCtx *direct.MapContext, in *krm.DataformRepositorySpec) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	out.GitRemoteSettings = RepositoryGitRemoteSettings_ToProto(mapCtx, in.GitRemoteSettings)
	out.WorkspaceCompilationOverrides = RepositoryWorkspaceCompilationOverrides_ToProto(mapCtx, in.WorkspaceCompilationOverrides)

	return out
}
func RepositoryGitRemoteSettings_FromProto(mapCtx *direct.MapContext, in *pb.Repository_GitRemoteSettings) *krm.RepositoryGitRemoteSettings {
	if in == nil {
		return nil
	}
	out := &krm.RepositoryGitRemoteSettings{}
	out.Url = in.GetUrl()
	out.DefaultBranch = in.GetDefaultBranch()
	// out.AuthenticationTokenSecretVersion = in.GetAuthenticationTokenSecretVersion() // todo acpana

	return out
}
func RepositoryGitRemoteSettings_ToProto(mapCtx *direct.MapContext, in *krm.RepositoryGitRemoteSettings) *pb.Repository_GitRemoteSettings {
	if in == nil {
		return nil
	}
	out := &pb.Repository_GitRemoteSettings{}
	out.Url = in.Url
	out.DefaultBranch = in.DefaultBranch
	// out.AuthenticationTokenSecretVersion = in.AuthenticationTokenSecretVersion // todo acpana

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
