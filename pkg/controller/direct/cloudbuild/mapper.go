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

package cloudbuild

import (
	pb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	krmcloudbuildv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1alpha1"
	krmcloudbuildv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func PrivatePoolV1Config_NetworkConfigSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PrivatePoolV1Config_NetworkConfig) *krmcloudbuildv1beta1.PrivatePoolV1Config_NetworkConfigSpec {
	return PrivatePoolV1Config_NetworkConfigSpec_FromProto(mapCtx, in)
}

func PrivatePoolV1Config_NetworkConfigSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmcloudbuildv1beta1.PrivatePoolV1Config_NetworkConfigSpec) *pb.PrivatePoolV1Config_NetworkConfig {
	return PrivatePoolV1Config_NetworkConfigSpec_ToProto(mapCtx, in)
}

func CloudBuildWorkerPoolSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmcloudbuildv1beta1.CloudBuildWorkerPoolSpec) *pb.WorkerPool {
	return CloudBuildWorkerPoolSpec_ToProto(mapCtx, in)
}

func Dependency_GitSourceRepository_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcloudbuildv1alpha1.Dependency_GitSourceRepository) *pb.Dependency_GitSourceRepository {
	if in == nil {
		return nil
	}
	out := &pb.Dependency_GitSourceRepository{}
	if oneof := Dependency_GitSourceRepository_Url_ToProto(mapCtx, in.URL); oneof != nil {
		out.Repotype = oneof
	}
	if in.DeveloperConnectRef != nil {
		out.Repotype = &pb.Dependency_GitSourceRepository_DeveloperConnect{DeveloperConnect: in.DeveloperConnectRef.External}
	}
	return out
}

func Artifacts_GoModule_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Artifacts_GoModule) *krmcloudbuildv1alpha1.Artifacts_GoModule {
	if in == nil {
		return nil
	}
	out := &krmcloudbuildv1alpha1.Artifacts_GoModule{}
	out.RepositoryName = direct.LazyPtr(in.GetRepositoryName())
	out.RepositoryLocation = direct.LazyPtr(in.GetRepositoryLocation())
	if in.GetRepositoryProjectId() != "" {
		out.RepositoryProjectRef = &refsv1beta1.ProjectRef{External: in.GetRepositoryProjectId()}
	}
	out.SourcePath = direct.LazyPtr(in.GetSourcePath())
	out.ModulePath = direct.LazyPtr(in.GetModulePath())
	out.ModuleVersion = direct.LazyPtr(in.GetModuleVersion())
	return out
}

func Artifacts_GoModule_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcloudbuildv1alpha1.Artifacts_GoModule) *pb.Artifacts_GoModule {
	if in == nil {
		return nil
	}
	out := &pb.Artifacts_GoModule{}
	out.RepositoryName = direct.ValueOf(in.RepositoryName)
	out.RepositoryLocation = direct.ValueOf(in.RepositoryLocation)
	if in.RepositoryProjectRef != nil {
		out.RepositoryProjectId = in.RepositoryProjectRef.External
	}
	out.SourcePath = direct.ValueOf(in.SourcePath)
	out.ModulePath = direct.ValueOf(in.ModulePath)
	out.ModuleVersion = direct.ValueOf(in.ModuleVersion)
	return out
}

func Build_RepoSource_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.RepoSource) *krmcloudbuildv1alpha1.Build_RepoSource {
	if in == nil {
		return nil
	}
	out := &krmcloudbuildv1alpha1.Build_RepoSource{}
	if in.GetProjectId() != "" {
		out.ProjectRef = &refsv1beta1.ProjectRef{External: in.GetProjectId()}
	}
	out.RepoName = direct.LazyPtr(in.GetRepoName())
	out.BranchName = direct.LazyPtr(in.GetBranchName())
	out.TagName = direct.LazyPtr(in.GetTagName())
	out.CommitSha = direct.LazyPtr(in.GetCommitSha())
	out.Dir = direct.LazyPtr(in.GetDir())
	out.InvertRegex = direct.LazyPtr(in.GetInvertRegex())
	out.Substitutions = in.Substitutions
	return out
}

func Build_RepoSource_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcloudbuildv1alpha1.Build_RepoSource) *pb.RepoSource {
	if in == nil {
		return nil
	}
	out := &pb.RepoSource{}
	if in.ProjectRef != nil {
		out.ProjectId = in.ProjectRef.External
	}
	out.RepoName = direct.ValueOf(in.RepoName)
	if in.BranchName != nil {
		out.Revision = &pb.RepoSource_BranchName{BranchName: *in.BranchName}
	}
	if in.TagName != nil {
		out.Revision = &pb.RepoSource_TagName{TagName: *in.TagName}
	}
	if in.CommitSha != nil {
		out.Revision = &pb.RepoSource_CommitSha{CommitSha: *in.CommitSha}
	}
	out.Dir = direct.ValueOf(in.Dir)
	out.InvertRegex = direct.ValueOf(in.InvertRegex)
	out.Substitutions = in.Substitutions
	return out
}

func GitConfig_HTTPConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.GitConfig_HttpConfig) *krmcloudbuildv1alpha1.GitConfig_HTTPConfig {
	if in == nil {
		return nil
	}
	out := &krmcloudbuildv1alpha1.GitConfig_HTTPConfig{}
	if in.GetProxySecretVersionName() != "" {
		out.ProxySecretVersionRef = &secretmanagerv1beta1.SecretManagerSecretVersionRef{External: in.GetProxySecretVersionName()}
	}
	return out
}

func GitConfig_HTTPConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcloudbuildv1alpha1.GitConfig_HTTPConfig) *pb.GitConfig_HttpConfig {
	if in == nil {
		return nil
	}
	out := &pb.GitConfig_HttpConfig{}
	if in.ProxySecretVersionRef != nil {
		out.ProxySecretVersionName = in.ProxySecretVersionRef.External
	}
	return out
}

func InlineSecret_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.InlineSecret) *krmcloudbuildv1alpha1.InlineSecret {
	if in == nil {
		return nil
	}
	out := &krmcloudbuildv1alpha1.InlineSecret{}
	if in.GetKmsKeyName() != "" {
		out.KMSKeyRef = &kmsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	return out
}

func InlineSecret_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcloudbuildv1alpha1.InlineSecret) *pb.InlineSecret {
	if in == nil {
		return nil
	}
	out := &pb.InlineSecret{}
	if in.KMSKeyRef != nil {
		out.KmsKeyName = in.KMSKeyRef.External
	}
	return out
}

func Secret_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Secret) *krmcloudbuildv1alpha1.Secret {
	if in == nil {
		return nil
	}
	out := &krmcloudbuildv1alpha1.Secret{}
	if in.GetKmsKeyName() != "" {
		out.KMSKeyRef = &kmsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	return out
}

func Secret_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcloudbuildv1alpha1.Secret) *pb.Secret {
	if in == nil {
		return nil
	}
	out := &pb.Secret{}
	if in.KMSKeyRef != nil {
		out.KmsKeyName = in.KMSKeyRef.External
	}
	return out
}

func SecretManagerSecret_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SecretManagerSecret) *krmcloudbuildv1alpha1.SecretManagerSecret {
	if in == nil {
		return nil
	}
	out := &krmcloudbuildv1alpha1.SecretManagerSecret{}
	if in.GetVersionName() != "" {
		out.VersionRef = &secretmanagerv1beta1.SecretManagerSecretVersionRef{External: in.GetVersionName()}
	}
	out.Env = direct.LazyPtr(in.GetEnv())
	return out
}

func SecretManagerSecret_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcloudbuildv1alpha1.SecretManagerSecret) *pb.SecretManagerSecret {
	if in == nil {
		return nil
	}
	out := &pb.SecretManagerSecret{}
	if in.VersionRef != nil {
		out.VersionName = in.VersionRef.External
	}
	out.Env = direct.ValueOf(in.Env)
	return out
}
