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
	"strconv"

	pb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
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

func CloudBuildTriggerSpec_FromProto(mapCtx *direct.MapContext, in *pb.BuildTrigger) *krm.CloudBuildTriggerSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerSpec{}
	out.Description = direct.LazyPtr(in.Description)
	out.Disabled = direct.LazyPtr(in.Disabled)
	out.Filename = direct.LazyPtr(in.GetFilename())
	out.Filter = direct.LazyPtr(in.Filter)
	out.IgnoredFiles = in.IgnoredFiles
	out.IncludedFiles = in.IncludedFiles
	out.Substitutions = in.Substitutions
	out.Tags = in.Tags

	// out.ApprovalConfig is missing from pb.BuildTrigger in this version of the library
	out.Build = CloudBuildTriggerSpec_Build_FromProto(mapCtx, in.GetBuild())
	out.Github = CloudBuildTriggerSpec_Github_FromProto(mapCtx, in.GetGithub())
	out.PubsubConfig = CloudBuildTriggerSpec_PubsubConfig_FromProto(mapCtx, in.GetPubsubConfig())
	out.TriggerTemplate = CloudBuildTriggerSpec_TriggerTemplate_FromProto(mapCtx, in.GetTriggerTemplate())
	out.WebhookConfig = CloudBuildTriggerSpec_WebhookConfig_FromProto(mapCtx, in.GetWebhookConfig())
	if in.ServiceAccount != "" {
		out.ServiceAccountRef = &krm.CloudBuildTriggerSpec_ServiceAccountRef{
			External: direct.LazyPtr(in.ServiceAccount),
		}
	}
	out.GitFileSource = CloudBuildTriggerSpec_GitFileSource_FromProto(mapCtx, in.GetGitFileSource())
	out.RepositoryEventConfig = CloudBuildTriggerSpec_RepositoryEventConfig_FromProto(mapCtx, in.GetRepositoryEventConfig())
	out.SourceToBuild = CloudBuildTriggerSpec_SourceToBuild_FromProto(mapCtx, in.GetSourceToBuild())

	return out
}

func CloudBuildTriggerSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildTriggerSpec) *pb.BuildTrigger {
	if in == nil {
		return nil
	}
	out := &pb.BuildTrigger{}
	out.Description = direct.ValueOf(in.Description)
	out.Disabled = direct.ValueOf(in.Disabled)
	out.Filter = direct.ValueOf(in.Filter)
	out.IgnoredFiles = in.IgnoredFiles
	out.IncludedFiles = in.IncludedFiles
	out.Substitutions = in.Substitutions
	out.Tags = in.Tags

	if in.Filename != nil {
		out.BuildTemplate = &pb.BuildTrigger_Filename{Filename: *in.Filename}
	}
	if in.Build != nil {
		out.BuildTemplate = &pb.BuildTrigger_Build{Build: CloudBuildTriggerSpec_Build_ToProto(mapCtx, in.Build)}
	}
	if in.GitFileSource != nil {
		out.BuildTemplate = &pb.BuildTrigger_GitFileSource{GitFileSource: CloudBuildTriggerSpec_GitFileSource_ToProto(mapCtx, in.GitFileSource)}
	}

	out.Github = CloudBuildTriggerSpec_Github_ToProto(mapCtx, in.Github)
	out.PubsubConfig = CloudBuildTriggerSpec_PubsubConfig_ToProto(mapCtx, in.PubsubConfig)
	out.TriggerTemplate = CloudBuildTriggerSpec_TriggerTemplate_ToProto(mapCtx, in.TriggerTemplate)
	out.WebhookConfig = CloudBuildTriggerSpec_WebhookConfig_ToProto(mapCtx, in.WebhookConfig)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = direct.ValueOf(in.ServiceAccountRef.External)
	}
	out.RepositoryEventConfig = CloudBuildTriggerSpec_RepositoryEventConfig_ToProto(mapCtx, in.RepositoryEventConfig)
	out.SourceToBuild = CloudBuildTriggerSpec_SourceToBuild_ToProto(mapCtx, in.SourceToBuild)

	return out
}

func CloudBuildTriggerSpec_Build_FromProto(mapCtx *direct.MapContext, in *pb.Build) *krm.CloudBuildTriggerSpec_Build {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerSpec_Build{}
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.Images = in.Images
	out.QueueTtl = direct.StringDuration_FromProto(mapCtx, in.GetQueueTtl())
	out.Substitutions = in.Substitutions
	out.Tags = in.Tags
	if in.LogsBucket != "" {
		out.LogsBucketRef = &krm.CloudBuildTriggerSpec_Build_LogsBucketRef{
			External: direct.LazyPtr(in.LogsBucket),
		}
	}
	out.Step = make([]krm.CloudBuildTriggerSpec_Build_StepItem, len(in.Steps))
	for i, step := range in.Steps {
		out.Step[i] = krm.CloudBuildTriggerSpec_Build_StepItem{
			Name:         direct.LazyPtr(step.Name),
			Args:         step.Args,
			Env:          step.Env,
			Dir:          direct.LazyPtr(step.Dir),
			Id:           direct.LazyPtr(step.Id),
			WaitFor:      step.WaitFor,
			Entrypoint:   direct.LazyPtr(step.Entrypoint),
			SecretEnv:    step.SecretEnv,
			Timeout:      direct.StringDuration_FromProto(mapCtx, step.Timeout),
			Script:       direct.LazyPtr(step.Script),
			AllowFailure: direct.LazyPtr(step.AllowFailure),
		}
		if len(step.AllowExitCodes) > 0 {
			out.Step[i].AllowExitCodes = make([]int, len(step.AllowExitCodes))
			for j, code := range step.AllowExitCodes {
				out.Step[i].AllowExitCodes[j] = int(code)
			}
		}
		if len(step.Volumes) > 0 {
			out.Step[i].Volumes = make([]krm.CloudBuildTriggerSpec_Build_StepItem_VolumesItem, len(step.Volumes))
			for j, vol := range step.Volumes {
				out.Step[i].Volumes[j] = krm.CloudBuildTriggerSpec_Build_StepItem_VolumesItem{
					Name: direct.LazyPtr(vol.Name),
					Path: direct.LazyPtr(vol.Path),
				}
			}
		}
	}

	out.Artifacts = CloudBuildTriggerSpec_Build_Artifacts_FromProto(mapCtx, in.GetArtifacts())
	out.Options = CloudBuildTriggerSpec_Build_Options_FromProto(mapCtx, in.GetOptions())
	out.Source = CloudBuildTriggerSpec_Build_Source_FromProto(mapCtx, in.GetSource())
	out.AvailableSecrets = CloudBuildTriggerSpec_Build_AvailableSecrets_FromProto(mapCtx, in.GetAvailableSecrets())
	if len(in.GetSecrets()) > 0 {
		out.Secret = make([]krm.CloudBuildTriggerSpec_Build_SecretItem, len(in.GetSecrets()))
		for i, s := range in.GetSecrets() {
			out.Secret[i] = CloudBuildTriggerSpec_Build_SecretItem_FromProto(mapCtx, s)
		}
	}

	return out
}

func CloudBuildTriggerSpec_Build_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildTriggerSpec_Build) *pb.Build {
	if in == nil {
		return nil
	}
	out := &pb.Build{}
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.Images = in.Images
	out.QueueTtl = direct.StringDuration_ToProto(mapCtx, in.QueueTtl)
	out.Substitutions = in.Substitutions
	out.Tags = in.Tags
	if in.LogsBucketRef != nil {
		out.LogsBucket = direct.ValueOf(in.LogsBucketRef.External)
	}
	out.Steps = make([]*pb.BuildStep, len(in.Step))
	for i, step := range in.Step {
		out.Steps[i] = &pb.BuildStep{
			Name:         direct.ValueOf(step.Name),
			Args:         step.Args,
			Env:          step.Env,
			Dir:          direct.ValueOf(step.Dir),
			Id:           direct.ValueOf(step.Id),
			WaitFor:      step.WaitFor,
			Entrypoint:   direct.ValueOf(step.Entrypoint),
			SecretEnv:    step.SecretEnv,
			Timeout:      direct.StringDuration_ToProto(mapCtx, step.Timeout),
			Script:       direct.ValueOf(step.Script),
			AllowFailure: direct.ValueOf(step.AllowFailure),
		}
		if len(step.AllowExitCodes) > 0 {
			out.Steps[i].AllowExitCodes = make([]int32, len(step.AllowExitCodes))
			for j, code := range step.AllowExitCodes {
				out.Steps[i].AllowExitCodes[j] = int32(code)
			}
		}
		if len(step.Volumes) > 0 {
			out.Steps[i].Volumes = make([]*pb.Volume, len(step.Volumes))
			for j, vol := range step.Volumes {
				out.Steps[i].Volumes[j] = &pb.Volume{
					Name: direct.ValueOf(vol.Name),
					Path: direct.ValueOf(vol.Path),
				}
			}
		}
	}

	out.Artifacts = CloudBuildTriggerSpec_Build_Artifacts_ToProto(mapCtx, in.Artifacts)
	out.Options = CloudBuildTriggerSpec_Build_Options_ToProto(mapCtx, in.Options)
	out.Source = CloudBuildTriggerSpec_Build_Source_ToProto(mapCtx, in.Source)
	out.AvailableSecrets = CloudBuildTriggerSpec_Build_AvailableSecrets_ToProto(mapCtx, in.AvailableSecrets)
	if len(in.Secret) > 0 {
		out.Secrets = make([]*pb.Secret, len(in.Secret))
		for i, s := range in.Secret {
			out.Secrets[i] = CloudBuildTriggerSpec_Build_SecretItem_ToProto(mapCtx, s)
		}
	}

	return out
}

func CloudBuildTriggerSpec_Build_Artifacts_FromProto(mapCtx *direct.MapContext, in *pb.Artifacts) *krm.CloudBuildTriggerSpec_Build_Artifacts {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerSpec_Build_Artifacts{}
	out.Images = in.Images
	if in.Objects != nil {
		out.Objects = &krm.CloudBuildTriggerSpec_Build_Artifacts_Objects{
			Location: direct.LazyPtr(in.Objects.Location),
			Paths:    in.Objects.Paths,
		}
	}
	return out
}

func CloudBuildTriggerSpec_Build_Artifacts_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildTriggerSpec_Build_Artifacts) *pb.Artifacts {
	if in == nil {
		return nil
	}
	out := &pb.Artifacts{}
	out.Images = in.Images
	if in.Objects != nil {
		out.Objects = &pb.Artifacts_ArtifactObjects{
			Location: direct.ValueOf(in.Objects.Location),
			Paths:    in.Objects.Paths,
		}
	}
	return out
}

func CloudBuildTriggerSpec_Build_Options_FromProto(mapCtx *direct.MapContext, in *pb.BuildOptions) *krm.CloudBuildTriggerSpec_Build_Options {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerSpec_Build_Options{}
	if len(in.SourceProvenanceHash) > 0 {
		out.SourceProvenanceHash = make([]string, len(in.SourceProvenanceHash))
		for i, h := range in.SourceProvenanceHash {
			out.SourceProvenanceHash[i] = direct.ValueOf(direct.Enum_FromProto(mapCtx, h))
		}
	}
	out.RequestedVerifyOption = direct.Enum_FromProto(mapCtx, in.RequestedVerifyOption)
	out.MachineType = direct.Enum_FromProto(mapCtx, in.MachineType)
	out.DiskSizeGb = direct.LazyPtr(int(in.DiskSizeGb))
	out.SubstitutionOption = direct.Enum_FromProto(mapCtx, in.SubstitutionOption)
	out.DynamicSubstitutions = direct.LazyPtr(in.DynamicSubstitutions)
	out.LogStreamingOption = direct.Enum_FromProto(mapCtx, in.LogStreamingOption)
	out.WorkerPool = direct.LazyPtr(in.WorkerPool)
	out.Logging = direct.Enum_FromProto(mapCtx, in.Logging)
	out.Env = in.Env
	out.SecretEnv = in.SecretEnv
	if len(in.Volumes) > 0 {
		out.Volumes = make([]krm.CloudBuildTriggerSpec_Build_Options_VolumesItem, len(in.Volumes))
		for i, vol := range in.Volumes {
			out.Volumes[i] = krm.CloudBuildTriggerSpec_Build_Options_VolumesItem{
				Name: direct.LazyPtr(vol.Name),
				Path: direct.LazyPtr(vol.Path),
			}
		}
	}
	return out
}

func CloudBuildTriggerSpec_Build_Options_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildTriggerSpec_Build_Options) *pb.BuildOptions {
	if in == nil {
		return nil
	}
	out := &pb.BuildOptions{}
	if len(in.SourceProvenanceHash) > 0 {
		out.SourceProvenanceHash = make([]pb.Hash_HashType, len(in.SourceProvenanceHash))
		for i, h := range in.SourceProvenanceHash {
			out.SourceProvenanceHash[i] = direct.Enum_ToProto[pb.Hash_HashType](mapCtx, direct.LazyPtr(h))
		}
	}
	out.RequestedVerifyOption = direct.Enum_ToProto[pb.BuildOptions_VerifyOption](mapCtx, in.RequestedVerifyOption)
	out.MachineType = direct.Enum_ToProto[pb.BuildOptions_MachineType](mapCtx, in.MachineType)
	out.DiskSizeGb = int64(direct.ValueOf(in.DiskSizeGb))
	out.SubstitutionOption = direct.Enum_ToProto[pb.BuildOptions_SubstitutionOption](mapCtx, in.SubstitutionOption)
	out.DynamicSubstitutions = direct.ValueOf(in.DynamicSubstitutions)
	out.LogStreamingOption = direct.Enum_ToProto[pb.BuildOptions_LogStreamingOption](mapCtx, in.LogStreamingOption)
	out.WorkerPool = direct.ValueOf(in.WorkerPool)
	out.Logging = direct.Enum_ToProto[pb.BuildOptions_LoggingMode](mapCtx, in.Logging)
	out.Env = in.Env
	out.SecretEnv = in.SecretEnv
	if len(in.Volumes) > 0 {
		out.Volumes = make([]*pb.Volume, len(in.Volumes))
		for i, vol := range in.Volumes {
			out.Volumes[i] = &pb.Volume{
				Name: direct.ValueOf(vol.Name),
				Path: direct.ValueOf(vol.Path),
			}
		}
	}
	return out
}

func CloudBuildTriggerSpec_Build_Source_FromProto(mapCtx *direct.MapContext, in *pb.Source) *krm.CloudBuildTriggerSpec_Build_Source {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerSpec_Build_Source{}
	if repo := in.GetRepoSource(); repo != nil {
		out.RepoSource = &krm.CloudBuildTriggerSpec_Build_Source_RepoSource{
			ProjectId:     direct.LazyPtr(repo.ProjectId),
			RepoRef:       &krm.CloudBuildTriggerSpec_Build_Source_RepoSource_RepoRef{External: direct.LazyPtr(repo.RepoName)},
			Dir:           direct.LazyPtr(repo.Dir),
			InvertRegex:   direct.LazyPtr(repo.InvertRegex),
			Substitutions: repo.Substitutions,
		}
		switch r := repo.Revision.(type) {
		case *pb.RepoSource_BranchName:
			out.RepoSource.BranchName = direct.LazyPtr(r.BranchName)
		case *pb.RepoSource_TagName:
			out.RepoSource.TagName = direct.LazyPtr(r.TagName)
		case *pb.RepoSource_CommitSha:
			out.RepoSource.CommitSha = direct.LazyPtr(r.CommitSha)
		}
	}
	if storage := in.GetStorageSource(); storage != nil {
		out.StorageSource = &krm.CloudBuildTriggerSpec_Build_Source_StorageSource{
			BucketRef:  &krm.CloudBuildTriggerSpec_Build_Source_StorageSource_BucketRef{External: direct.LazyPtr(storage.Bucket)},
			Object:     direct.LazyPtr(storage.Object),
			Generation: direct.LazyPtr(strconv.FormatInt(storage.Generation, 10)),
		}
	}
	return out
}

func CloudBuildTriggerSpec_Build_Source_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildTriggerSpec_Build_Source) *pb.Source {
	if in == nil {
		return nil
	}
	out := &pb.Source{}
	if in.RepoSource != nil {
		repo := &pb.RepoSource{
			ProjectId:     direct.ValueOf(in.RepoSource.ProjectId),
			RepoName:      direct.ValueOf(in.RepoSource.RepoRef.External),
			Dir:           direct.ValueOf(in.RepoSource.Dir),
			InvertRegex:   direct.ValueOf(in.RepoSource.InvertRegex),
			Substitutions: in.RepoSource.Substitutions,
		}
		if in.RepoSource.BranchName != nil {
			repo.Revision = &pb.RepoSource_BranchName{BranchName: *in.RepoSource.BranchName}
		} else if in.RepoSource.TagName != nil {
			repo.Revision = &pb.RepoSource_TagName{TagName: *in.RepoSource.TagName}
		} else if in.RepoSource.CommitSha != nil {
			repo.Revision = &pb.RepoSource_CommitSha{CommitSha: *in.RepoSource.CommitSha}
		}
		out.Source = &pb.Source_RepoSource{RepoSource: repo}
	}
	if in.StorageSource != nil {
		storage := &pb.StorageSource{
			Bucket: direct.ValueOf(in.StorageSource.BucketRef.External),
			Object: direct.ValueOf(in.StorageSource.Object),
		}
		if in.StorageSource.Generation != nil {
			gen, err := strconv.ParseInt(*in.StorageSource.Generation, 10, 64)
			if err != nil {
				mapCtx.Errorf("invalid generation %q: %v", *in.StorageSource.Generation, err)
			} else {
				storage.Generation = gen
			}
		}
		out.Source = &pb.Source_StorageSource{StorageSource: storage}
	}
	return out
}

func CloudBuildTriggerSpec_Build_AvailableSecrets_FromProto(mapCtx *direct.MapContext, in *pb.Secrets) *krm.CloudBuildTriggerSpec_Build_AvailableSecrets {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerSpec_Build_AvailableSecrets{}
	if len(in.SecretManager) > 0 {
		out.SecretManager = make([]krm.CloudBuildTriggerSpec_Build_AvailableSecrets_SecretManagerItem, len(in.SecretManager))
		for i, s := range in.SecretManager {
			out.SecretManager[i] = krm.CloudBuildTriggerSpec_Build_AvailableSecrets_SecretManagerItem{
				VersionRef: &krm.CloudBuildTriggerSpec_Build_AvailableSecrets_SecretManagerItem_VersionRef{External: direct.LazyPtr(s.VersionName)},
				Env:        direct.LazyPtr(s.Env),
			}
		}
	}
	return out
}

func CloudBuildTriggerSpec_Build_AvailableSecrets_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildTriggerSpec_Build_AvailableSecrets) *pb.Secrets {
	if in == nil {
		return nil
	}
	out := &pb.Secrets{}
	if len(in.SecretManager) > 0 {
		out.SecretManager = make([]*pb.SecretManagerSecret, len(in.SecretManager))
		for i, s := range in.SecretManager {
			out.SecretManager[i] = &pb.SecretManagerSecret{
				VersionName: direct.ValueOf(s.VersionRef.External),
				Env:         direct.ValueOf(s.Env),
			}
		}
	}
	return out
}

func CloudBuildTriggerSpec_Build_SecretItem_FromProto(mapCtx *direct.MapContext, in *pb.Secret) krm.CloudBuildTriggerSpec_Build_SecretItem {
	out := krm.CloudBuildTriggerSpec_Build_SecretItem{}
	out.KmsKeyRef = &krm.CloudBuildTriggerSpec_Build_SecretItem_KmsKeyRef{External: direct.LazyPtr(in.KmsKeyName)}
	if in.SecretEnv != nil {
		out.SecretEnv = make(map[string]string)
		for k, v := range in.SecretEnv {
			out.SecretEnv[k] = string(v)
		}
	}
	return out
}

func CloudBuildTriggerSpec_Build_SecretItem_ToProto(mapCtx *direct.MapContext, in krm.CloudBuildTriggerSpec_Build_SecretItem) *pb.Secret {
	out := &pb.Secret{}
	if in.KmsKeyRef != nil {
		out.KmsKeyName = direct.ValueOf(in.KmsKeyRef.External)
	}
	if in.SecretEnv != nil {
		out.SecretEnv = make(map[string][]byte)
		for k, v := range in.SecretEnv {
			out.SecretEnv[k] = []byte(v)
		}
	}
	return out
}

func CloudBuildTriggerSpec_Github_FromProto(mapCtx *direct.MapContext, in *pb.GitHubEventsConfig) *krm.CloudBuildTriggerSpec_Github {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerSpec_Github{}
	out.Name = direct.LazyPtr(in.Name)
	out.Owner = direct.LazyPtr(in.Owner)
	if pr := in.GetPullRequest(); pr != nil {
		out.PullRequest = &krm.CloudBuildTriggerSpec_Github_PullRequest{
			CommentControl: direct.Enum_FromProto(mapCtx, pr.CommentControl),
		}
		if prBranch, ok := pr.GetGitRef().(*pb.PullRequestFilter_Branch); ok {
			out.PullRequest.Branch = direct.LazyPtr(prBranch.Branch)
		}
	}
	if push := in.GetPush(); push != nil {
		out.Push = &krm.CloudBuildTriggerSpec_Github_Push{
			InvertRegex: direct.LazyPtr(push.InvertRegex),
		}
		if pushBranch, ok := push.GetGitRef().(*pb.PushFilter_Branch); ok {
			out.Push.Branch = direct.LazyPtr(pushBranch.Branch)
		}
		if pushTag, ok := push.GetGitRef().(*pb.PushFilter_Tag); ok {
			out.Push.Tag = direct.LazyPtr(pushTag.Tag)
		}
	}
	return out
}

func CloudBuildTriggerSpec_Github_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildTriggerSpec_Github) *pb.GitHubEventsConfig {
	if in == nil {
		return nil
	}
	out := &pb.GitHubEventsConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.Owner = direct.ValueOf(in.Owner)
	if in.PullRequest != nil {
		pr := &pb.PullRequestFilter{
			CommentControl: direct.Enum_ToProto[pb.PullRequestFilter_CommentControl](mapCtx, in.PullRequest.CommentControl),
		}
		if in.PullRequest.Branch != nil {
			pr.GitRef = &pb.PullRequestFilter_Branch{Branch: *in.PullRequest.Branch}
		}
		out.Event = &pb.GitHubEventsConfig_PullRequest{PullRequest: pr}
	}
	if in.Push != nil {
		push := &pb.PushFilter{
			InvertRegex: direct.ValueOf(in.Push.InvertRegex),
		}
		if in.Push.Branch != nil {
			push.GitRef = &pb.PushFilter_Branch{Branch: *in.Push.Branch}
		}
		if in.Push.Tag != nil {
			push.GitRef = &pb.PushFilter_Tag{Tag: *in.Push.Tag}
		}
		out.Event = &pb.GitHubEventsConfig_Push{Push: push}
	}
	return out
}

func CloudBuildTriggerSpec_PubsubConfig_FromProto(mapCtx *direct.MapContext, in *pb.PubsubConfig) *krm.CloudBuildTriggerSpec_PubsubConfig {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerSpec_PubsubConfig{}
	out.Subscription = direct.LazyPtr(in.Subscription)
	out.State = direct.Enum_FromProto(mapCtx, in.State)
	if in.Topic != "" {
		out.TopicRef = &krm.CloudBuildTriggerSpec_PubsubConfig_TopicRef{
			External: direct.LazyPtr(in.Topic),
		}
	}
	if in.ServiceAccountEmail != "" {
		out.ServiceAccountRef = &krm.CloudBuildTriggerSpec_PubsubConfig_ServiceAccountRef{
			External: direct.LazyPtr(in.ServiceAccountEmail),
		}
	}
	return out
}

func CloudBuildTriggerSpec_PubsubConfig_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildTriggerSpec_PubsubConfig) *pb.PubsubConfig {
	if in == nil {
		return nil
	}
	out := &pb.PubsubConfig{}
	out.Subscription = direct.ValueOf(in.Subscription)
	out.State = direct.Enum_ToProto[pb.PubsubConfig_State](mapCtx, in.State)
	if in.TopicRef != nil {
		out.Topic = direct.ValueOf(in.TopicRef.External)
	}
	if in.ServiceAccountRef != nil {
		out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountRef.External)
	}
	return out
}

func CloudBuildTriggerSpec_TriggerTemplate_FromProto(mapCtx *direct.MapContext, in *pb.RepoSource) *krm.CloudBuildTriggerSpec_TriggerTemplate {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerSpec_TriggerTemplate{}
	out.Dir = direct.LazyPtr(in.Dir)
	out.InvertRegex = direct.LazyPtr(in.InvertRegex)
	if in.RepoName != "" {
		out.RepoRef = &krm.CloudBuildTriggerSpec_TriggerTemplate_RepoRef{
			External: direct.LazyPtr(in.RepoName),
		}
	}
	if in.Revision != nil {
		switch r := in.Revision.(type) {
		case *pb.RepoSource_BranchName:
			out.BranchName = direct.LazyPtr(r.BranchName)
		case *pb.RepoSource_TagName:
			out.TagName = direct.LazyPtr(r.TagName)
		case *pb.RepoSource_CommitSha:
			out.CommitSha = direct.LazyPtr(r.CommitSha)
		}
	}
	return out
}

func CloudBuildTriggerSpec_TriggerTemplate_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildTriggerSpec_TriggerTemplate) *pb.RepoSource {
	if in == nil {
		return nil
	}
	out := &pb.RepoSource{}
	out.Dir = direct.ValueOf(in.Dir)
	out.InvertRegex = direct.ValueOf(in.InvertRegex)
	if in.RepoRef != nil {
		out.RepoName = direct.ValueOf(in.RepoRef.External)
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
	return out
}

func CloudBuildTriggerSpec_WebhookConfig_FromProto(mapCtx *direct.MapContext, in *pb.WebhookConfig) *krm.CloudBuildTriggerSpec_WebhookConfig {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerSpec_WebhookConfig{}
	out.State = direct.Enum_FromProto(mapCtx, in.State)
	return out
}

func CloudBuildTriggerSpec_WebhookConfig_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildTriggerSpec_WebhookConfig) *pb.WebhookConfig {
	if in == nil {
		return nil
	}
	out := &pb.WebhookConfig{}
	out.State = direct.Enum_ToProto[pb.WebhookConfig_State](mapCtx, in.State)
	if in.SecretRef != nil {
		out.AuthMethod = &pb.WebhookConfig_Secret{
			Secret: direct.ValueOf(in.SecretRef.External),
		}
	}
	return out
}

func CloudBuildTriggerSpec_GitFileSource_FromProto(mapCtx *direct.MapContext, in *pb.GitFileSource) *krm.CloudBuildTriggerSpec_GitFileSource {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerSpec_GitFileSource{}
	out.Path = direct.LazyPtr(in.Path)
	out.Uri = direct.LazyPtr(in.Uri)
	out.Revision = direct.LazyPtr(in.Revision)
	out.RepoType = direct.Enum_FromProto(mapCtx, in.RepoType)
	if in.Source != nil {
		switch s := in.Source.(type) {
		case *pb.GitFileSource_Repository:
			out.RepositoryRef = &krm.CloudBuildTriggerSpec_GitFileSource_RepositoryRef{
				External: direct.LazyPtr(s.Repository),
			}
		}
	}
	if in.EnterpriseConfig != nil {
		switch e := in.EnterpriseConfig.(type) {
		case *pb.GitFileSource_GithubEnterpriseConfig:
			out.GithubEnterpriseConfigRef = &krm.CloudBuildTriggerSpec_GitFileSource_GithubEnterpriseConfigRef{
				External: direct.LazyPtr(e.GithubEnterpriseConfig),
			}
		}
	}
	return out
}

func CloudBuildTriggerSpec_GitFileSource_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildTriggerSpec_GitFileSource) *pb.GitFileSource {
	if in == nil {
		return nil
	}
	out := &pb.GitFileSource{}
	out.Path = direct.ValueOf(in.Path)
	out.Uri = direct.ValueOf(in.Uri)
	out.Revision = direct.ValueOf(in.Revision)
	out.RepoType = direct.Enum_ToProto[pb.GitFileSource_RepoType](mapCtx, in.RepoType)
	if in.RepositoryRef != nil {
		out.Source = &pb.GitFileSource_Repository{Repository: direct.ValueOf(in.RepositoryRef.External)}
	}
	if in.GithubEnterpriseConfigRef != nil {
		out.EnterpriseConfig = &pb.GitFileSource_GithubEnterpriseConfig{GithubEnterpriseConfig: direct.ValueOf(in.GithubEnterpriseConfigRef.External)}
	}
	return out
}

func CloudBuildTriggerSpec_RepositoryEventConfig_FromProto(mapCtx *direct.MapContext, in *pb.RepositoryEventConfig) *krm.CloudBuildTriggerSpec_RepositoryEventConfig {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerSpec_RepositoryEventConfig{}
	out.Repository = direct.LazyPtr(in.Repository)
	if pr := in.GetPullRequest(); pr != nil {
		out.PullRequest = &krm.CloudBuildTriggerSpec_RepositoryEventConfig_PullRequest{
			CommentControl: direct.Enum_FromProto(mapCtx, pr.CommentControl),
			InvertRegex:    direct.LazyPtr(pr.InvertRegex),
		}
		if prBranch, ok := pr.GetGitRef().(*pb.PullRequestFilter_Branch); ok {
			out.PullRequest.Branch = direct.LazyPtr(prBranch.Branch)
		}
	}
	if push := in.GetPush(); push != nil {
		out.Push = &krm.CloudBuildTriggerSpec_RepositoryEventConfig_Push{
			InvertRegex: direct.LazyPtr(push.InvertRegex),
		}
		if pushBranch, ok := push.GetGitRef().(*pb.PushFilter_Branch); ok {
			out.Push.Branch = direct.LazyPtr(pushBranch.Branch)
		}
		if pushTag, ok := push.GetGitRef().(*pb.PushFilter_Tag); ok {
			out.Push.Tag = direct.LazyPtr(pushTag.Tag)
		}
	}
	return out
}

func CloudBuildTriggerSpec_RepositoryEventConfig_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildTriggerSpec_RepositoryEventConfig) *pb.RepositoryEventConfig {
	if in == nil {
		return nil
	}
	out := &pb.RepositoryEventConfig{}
	out.Repository = direct.ValueOf(in.Repository)
	if in.PullRequest != nil {
		pr := &pb.PullRequestFilter{
			CommentControl: direct.Enum_ToProto[pb.PullRequestFilter_CommentControl](mapCtx, in.PullRequest.CommentControl),
			InvertRegex:    direct.ValueOf(in.PullRequest.InvertRegex),
		}
		if in.PullRequest.Branch != nil {
			pr.GitRef = &pb.PullRequestFilter_Branch{Branch: *in.PullRequest.Branch}
		}
		out.Filter = &pb.RepositoryEventConfig_PullRequest{PullRequest: pr}
	}
	if in.Push != nil {
		push := &pb.PushFilter{
			InvertRegex: direct.ValueOf(in.Push.InvertRegex),
		}
		if in.Push.Branch != nil {
			push.GitRef = &pb.PushFilter_Branch{Branch: *in.Push.Branch}
		}
		if in.Push.Tag != nil {
			push.GitRef = &pb.PushFilter_Tag{Tag: *in.Push.Tag}
		}
		out.Filter = &pb.RepositoryEventConfig_Push{Push: push}
	}
	return out
}

func CloudBuildTriggerSpec_SourceToBuild_FromProto(mapCtx *direct.MapContext, in *pb.GitRepoSource) *krm.CloudBuildTriggerSpec_SourceToBuild {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerSpec_SourceToBuild{}
	out.Uri = direct.LazyPtr(in.Uri)
	out.Ref = direct.LazyPtr(in.Ref)
	out.RepoType = direct.Enum_FromProto(mapCtx, in.RepoType)
	return out
}

func CloudBuildTriggerSpec_SourceToBuild_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildTriggerSpec_SourceToBuild) *pb.GitRepoSource {
	if in == nil {
		return nil
	}
	out := &pb.GitRepoSource{}
	out.Uri = direct.ValueOf(in.Uri)
	out.Ref = direct.ValueOf(in.Ref)
	out.RepoType = direct.Enum_ToProto[pb.GitFileSource_RepoType](mapCtx, in.RepoType)
	return out
}

func CloudBuildTriggerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BuildTrigger) *krm.CloudBuildTriggerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildTriggerObservedState{}
	out.Id = direct.LazyPtr(in.Id)
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.CreateTime)
	return out
}

func CloudBuildTriggerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildTriggerObservedState) *pb.BuildTrigger {
	if in == nil {
		return nil
	}
	out := &pb.BuildTrigger{}
	out.Id = direct.ValueOf(in.Id)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
