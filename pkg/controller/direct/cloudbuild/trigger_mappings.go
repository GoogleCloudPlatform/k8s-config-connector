// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudbuild

import (
	pb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Hash_FromProto(mapCtx *direct.MapContext, in *pb.Hash) *krm.Hash {
	if in == nil {
		return nil
	}
	out := &krm.Hash{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Value = in.GetValue()
	return out
}

func Hash_ToProto(mapCtx *direct.MapContext, in *krm.Hash) *pb.Hash {
	if in == nil {
		return nil
	}
	out := &pb.Hash{}
	out.Type = direct.Enum_ToProto[pb.Hash_HashType](mapCtx, in.Type)
	out.Value = in.Value
	return out
}

func RepoSource_FromProto(mapCtx *direct.MapContext, in *pb.RepoSource) *krm.RepoSource {
	if in == nil {
		return nil
	}
	out := &krm.RepoSource{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.RepoName = direct.LazyPtr(in.GetRepoName())
	if in.GetRepoName() != "" {
		out.RepoRef = &krm.SourceRepoRepositoryRef{External: in.GetRepoName()}
	}
	out.BranchName = direct.LazyPtr(in.GetBranchName())
	out.TagName = direct.LazyPtr(in.GetTagName())
	out.CommitSha = direct.LazyPtr(in.GetCommitSha())
	out.Dir = direct.LazyPtr(in.GetDir())
	out.InvertRegex = direct.LazyPtr(in.GetInvertRegex())
	out.Substitutions = in.Substitutions
	return out
}

func RepoSource_ToProto(mapCtx *direct.MapContext, in *krm.RepoSource) *pb.RepoSource {
	if in == nil {
		return nil
	}
	out := &pb.RepoSource{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.RepoName = direct.ValueOf(in.RepoName)
	if in.RepoRef != nil {
		out.RepoName = in.RepoRef.External
	}
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

func PubsubConfig_FromProto(mapCtx *direct.MapContext, in *pb.PubsubConfig) *krm.PubsubConfig {
	if in == nil {
		return nil
	}
	out := &krm.PubsubConfig{}
	out.Subscription = direct.LazyPtr(in.GetSubscription())
	out.Topic = direct.LazyPtr(in.GetTopic())
	if in.GetTopic() != "" {
		out.TopicRef = &pubsubv1beta1.PubSubTopicRef{External: in.GetTopic()}
	}
	out.ServiceAccountEmail = direct.LazyPtr(in.GetServiceAccountEmail())
	if in.GetServiceAccountEmail() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccountEmail()}
	}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}

func PubsubConfig_ToProto(mapCtx *direct.MapContext, in *krm.PubsubConfig) *pb.PubsubConfig {
	if in == nil {
		return nil
	}
	out := &pb.PubsubConfig{}
	out.Topic = direct.ValueOf(in.Topic)
	if in.TopicRef != nil {
		out.Topic = in.TopicRef.External
	}
	out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountEmail)
	if in.ServiceAccountRef != nil {
		out.ServiceAccountEmail = in.ServiceAccountRef.External
	}
	return out
}

func WebhookConfig_FromProto(mapCtx *direct.MapContext, in *pb.WebhookConfig) *krm.WebhookConfig {
	if in == nil {
		return nil
	}
	out := &krm.WebhookConfig{}
	if secret := in.GetSecret(); secret != "" {
		out.Secret = direct.LazyPtr(secret)
		out.SecretRef = &refsv1beta1.SecretManagerSecretVersionRef{External: secret}
	}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}

func WebhookConfig_ToProto(mapCtx *direct.MapContext, in *krm.WebhookConfig) *pb.WebhookConfig {
	if in == nil {
		return nil
	}
	out := &pb.WebhookConfig{}
	secret := direct.ValueOf(in.Secret)
	if in.SecretRef != nil {
		secret = in.SecretRef.External
	}
	if secret != "" {
		out.AuthMethod = &pb.WebhookConfig_Secret{Secret: secret}
	}
	out.State = direct.Enum_ToProto[pb.WebhookConfig_State](mapCtx, in.State)
	return out
}

func GitFileSource_FromProto(mapCtx *direct.MapContext, in *pb.GitFileSource) *krm.GitFileSource {
	if in == nil {
		return nil
	}
	out := &krm.GitFileSource{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.Repository = direct.LazyPtr(in.GetRepository())
	if in.GetRepository() != "" {
		out.RepositoryRef = &krm.CloudBuildV2RepositoryRef{External: in.GetRepository()}
	}
	out.Revision = direct.LazyPtr(in.GetRevision())
	out.RepoType = direct.Enum_FromProto(mapCtx, in.GetRepoType())
	out.GithubEnterpriseConfig = direct.LazyPtr(in.GetGithubEnterpriseConfig())
	return out
}

func GitFileSource_ToProto(mapCtx *direct.MapContext, in *krm.GitFileSource) *pb.GitFileSource {
	if in == nil {
		return nil
	}
	out := &pb.GitFileSource{}
	out.Uri = direct.ValueOf(in.URI)
	if in.Repository != nil {
		out.Source = &pb.GitFileSource_Repository{Repository: *in.Repository}
	}
	if in.RepositoryRef != nil {
		out.Source = &pb.GitFileSource_Repository{Repository: in.RepositoryRef.External}
	}
	out.Revision = direct.ValueOf(in.Revision)
	out.RepoType = direct.Enum_ToProto[pb.GitFileSource_RepoType](mapCtx, in.RepoType)
	if in.GithubEnterpriseConfig != nil {
		out.EnterpriseConfig = &pb.GitFileSource_GithubEnterpriseConfig{GithubEnterpriseConfig: *in.GithubEnterpriseConfig}
	}
	return out
}

func CloudBuildTriggerSpec_FromProto(mapCtx *direct.MapContext, in *pb.BuildTrigger) *krm.CloudBuildTriggerSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Tags = in.Tags
	out.TriggerTemplate = RepoSource_FromProto(mapCtx, in.GetTriggerTemplate())
	out.Github = GitHubEventsConfig_FromProto(mapCtx, in.GetGithub())
	out.PubsubConfig = PubsubConfig_FromProto(mapCtx, in.GetPubsubConfig())
	out.WebhookConfig = WebhookConfig_FromProto(mapCtx, in.GetWebhookConfig())
	out.Build = Build_FromProto(mapCtx, in.GetBuild())
	out.Filename = direct.LazyPtr(in.GetFilename())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.Substitutions = in.Substitutions
	out.IgnoredFiles = in.IgnoredFiles
	out.IncludedFiles = in.IncludedFiles
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.SourceToBuild = GitRepoSource_FromProto(mapCtx, in.GetSourceToBuild())
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.RepositoryEventConfig = RepositoryEventConfig_FromProto(mapCtx, in.GetRepositoryEventConfig())
	out.GitFileSource = GitFileSource_FromProto(mapCtx, in.GetGitFileSource())
	// MISSING: includeBuildLogs - not found in Proto
	// MISSING: approvalConfig - not found in Proto
	return out
}

func CloudBuildTriggerSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildTriggerSpec) *pb.BuildTrigger {
	if in == nil {
		return nil
	}
	out := &pb.BuildTrigger{}
	out.Description = direct.ValueOf(in.Description)
	out.Tags = in.Tags
	out.TriggerTemplate = RepoSource_ToProto(mapCtx, in.TriggerTemplate)
	out.Github = GitHubEventsConfig_ToProto(mapCtx, in.Github)
	out.PubsubConfig = PubsubConfig_ToProto(mapCtx, in.PubsubConfig)
	out.WebhookConfig = WebhookConfig_ToProto(mapCtx, in.WebhookConfig)
	if in.Build != nil {
		out.BuildTemplate = &pb.BuildTrigger_Build{Build: Build_ToProto(mapCtx, in.Build)}
	}
	if in.Filename != nil {
		out.BuildTemplate = &pb.BuildTrigger_Filename{Filename: *in.Filename}
	}
	if in.GitFileSource != nil {
		out.BuildTemplate = &pb.BuildTrigger_GitFileSource{GitFileSource: GitFileSource_ToProto(mapCtx, in.GitFileSource)}
	}
	out.Disabled = direct.ValueOf(in.Disabled)
	out.Substitutions = in.Substitutions
	out.IgnoredFiles = in.IgnoredFiles
	out.IncludedFiles = in.IncludedFiles
	out.Filter = direct.ValueOf(in.Filter)
	out.SourceToBuild = GitRepoSource_ToProto(mapCtx, in.SourceToBuild)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.RepositoryEventConfig = RepositoryEventConfig_ToProto(mapCtx, in.RepositoryEventConfig)
	return out
}

func CloudBuildTriggerStatus_FromProto(mapCtx *direct.MapContext, in *pb.BuildTrigger) *krm.CloudBuildTriggerStatus {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerStatus{}
	out.TriggerId = direct.LazyPtr(in.GetId())
	return out
}
