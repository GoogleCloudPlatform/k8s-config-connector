// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	cloudbuildpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudbuild/cloudbuild_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild"
)

// Server implements the gRPC interface for BuildTrigger.
type BuildTriggerServer struct{}

// ProtoToBuildTriggerGithubPullRequestCommentControlEnum converts a BuildTriggerGithubPullRequestCommentControlEnum enum from its proto representation.
func ProtoToCloudbuildBuildTriggerGithubPullRequestCommentControlEnum(e cloudbuildpb.CloudbuildBuildTriggerGithubPullRequestCommentControlEnum) *cloudbuild.BuildTriggerGithubPullRequestCommentControlEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudbuildpb.CloudbuildBuildTriggerGithubPullRequestCommentControlEnum_name[int32(e)]; ok {
		e := cloudbuild.BuildTriggerGithubPullRequestCommentControlEnum(n[len("CloudbuildBuildTriggerGithubPullRequestCommentControlEnum"):])
		return &e
	}
	return nil
}

// ProtoToBuildTriggerBuildStepsStatusEnum converts a BuildTriggerBuildStepsStatusEnum enum from its proto representation.
func ProtoToCloudbuildBuildTriggerBuildStepsStatusEnum(e cloudbuildpb.CloudbuildBuildTriggerBuildStepsStatusEnum) *cloudbuild.BuildTriggerBuildStepsStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudbuildpb.CloudbuildBuildTriggerBuildStepsStatusEnum_name[int32(e)]; ok {
		e := cloudbuild.BuildTriggerBuildStepsStatusEnum(n[len("CloudbuildBuildTriggerBuildStepsStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToBuildTriggerTriggerTemplate converts a BuildTriggerTriggerTemplate resource from its proto representation.
func ProtoToCloudbuildBuildTriggerTriggerTemplate(p *cloudbuildpb.CloudbuildBuildTriggerTriggerTemplate) *cloudbuild.BuildTriggerTriggerTemplate {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.BuildTriggerTriggerTemplate{
		ProjectId:   dcl.StringOrNil(p.ProjectId),
		RepoName:    dcl.StringOrNil(p.RepoName),
		BranchName:  dcl.StringOrNil(p.BranchName),
		TagName:     dcl.StringOrNil(p.TagName),
		CommitSha:   dcl.StringOrNil(p.CommitSha),
		Dir:         dcl.StringOrNil(p.Dir),
		InvertRegex: dcl.Bool(p.InvertRegex),
	}
	return obj
}

// ProtoToBuildTriggerGithub converts a BuildTriggerGithub resource from its proto representation.
func ProtoToCloudbuildBuildTriggerGithub(p *cloudbuildpb.CloudbuildBuildTriggerGithub) *cloudbuild.BuildTriggerGithub {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.BuildTriggerGithub{
		Owner:       dcl.StringOrNil(p.Owner),
		Name:        dcl.StringOrNil(p.Name),
		PullRequest: ProtoToCloudbuildBuildTriggerGithubPullRequest(p.GetPullRequest()),
		Push:        ProtoToCloudbuildBuildTriggerGithubPush(p.GetPush()),
	}
	return obj
}

// ProtoToBuildTriggerGithubPullRequest converts a BuildTriggerGithubPullRequest resource from its proto representation.
func ProtoToCloudbuildBuildTriggerGithubPullRequest(p *cloudbuildpb.CloudbuildBuildTriggerGithubPullRequest) *cloudbuild.BuildTriggerGithubPullRequest {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.BuildTriggerGithubPullRequest{
		Branch:         dcl.StringOrNil(p.Branch),
		CommentControl: ProtoToCloudbuildBuildTriggerGithubPullRequestCommentControlEnum(p.GetCommentControl()),
		InvertRegex:    dcl.Bool(p.InvertRegex),
	}
	return obj
}

// ProtoToBuildTriggerGithubPush converts a BuildTriggerGithubPush resource from its proto representation.
func ProtoToCloudbuildBuildTriggerGithubPush(p *cloudbuildpb.CloudbuildBuildTriggerGithubPush) *cloudbuild.BuildTriggerGithubPush {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.BuildTriggerGithubPush{
		Branch:      dcl.StringOrNil(p.Branch),
		Tag:         dcl.StringOrNil(p.Tag),
		InvertRegex: dcl.Bool(p.InvertRegex),
	}
	return obj
}

// ProtoToBuildTriggerBuild converts a BuildTriggerBuild resource from its proto representation.
func ProtoToCloudbuildBuildTriggerBuild(p *cloudbuildpb.CloudbuildBuildTriggerBuild) *cloudbuild.BuildTriggerBuild {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.BuildTriggerBuild{
		QueueTtl:   dcl.StringOrNil(p.QueueTtl),
		LogsBucket: dcl.StringOrNil(p.LogsBucket),
		Timeout:    dcl.StringOrNil(p.Timeout),
		Source:     ProtoToCloudbuildBuildTriggerBuildSource(p.GetSource()),
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	for _, r := range p.GetImages() {
		obj.Images = append(obj.Images, r)
	}
	for _, r := range p.GetSecrets() {
		obj.Secrets = append(obj.Secrets, *ProtoToCloudbuildBuildTriggerBuildSecrets(r))
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToCloudbuildBuildTriggerBuildSteps(r))
	}
	return obj
}

// ProtoToBuildTriggerBuildSecrets converts a BuildTriggerBuildSecrets resource from its proto representation.
func ProtoToCloudbuildBuildTriggerBuildSecrets(p *cloudbuildpb.CloudbuildBuildTriggerBuildSecrets) *cloudbuild.BuildTriggerBuildSecrets {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.BuildTriggerBuildSecrets{
		KmsKeyName: dcl.StringOrNil(p.KmsKeyName),
	}
	return obj
}

// ProtoToBuildTriggerBuildSteps converts a BuildTriggerBuildSteps resource from its proto representation.
func ProtoToCloudbuildBuildTriggerBuildSteps(p *cloudbuildpb.CloudbuildBuildTriggerBuildSteps) *cloudbuild.BuildTriggerBuildSteps {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.BuildTriggerBuildSteps{
		Name:       dcl.StringOrNil(p.Name),
		Dir:        dcl.StringOrNil(p.Dir),
		Id:         dcl.StringOrNil(p.Id),
		Entrypoint: dcl.StringOrNil(p.Entrypoint),
		Timing:     ProtoToCloudbuildBuildTriggerBuildStepsTiming(p.GetTiming()),
		PullTiming: ProtoToCloudbuildBuildTriggerBuildStepsPullTiming(p.GetPullTiming()),
		Timeout:    dcl.StringOrNil(p.Timeout),
		Status:     ProtoToCloudbuildBuildTriggerBuildStepsStatusEnum(p.GetStatus()),
	}
	for _, r := range p.GetEnv() {
		obj.Env = append(obj.Env, r)
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetWaitFor() {
		obj.WaitFor = append(obj.WaitFor, r)
	}
	for _, r := range p.GetSecretEnv() {
		obj.SecretEnv = append(obj.SecretEnv, r)
	}
	for _, r := range p.GetVolumes() {
		obj.Volumes = append(obj.Volumes, *ProtoToCloudbuildBuildTriggerBuildStepsVolumes(r))
	}
	return obj
}

// ProtoToBuildTriggerBuildStepsVolumes converts a BuildTriggerBuildStepsVolumes resource from its proto representation.
func ProtoToCloudbuildBuildTriggerBuildStepsVolumes(p *cloudbuildpb.CloudbuildBuildTriggerBuildStepsVolumes) *cloudbuild.BuildTriggerBuildStepsVolumes {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.BuildTriggerBuildStepsVolumes{
		Name: dcl.StringOrNil(p.Name),
		Path: dcl.StringOrNil(p.Path),
	}
	return obj
}

// ProtoToBuildTriggerBuildStepsTiming converts a BuildTriggerBuildStepsTiming resource from its proto representation.
func ProtoToCloudbuildBuildTriggerBuildStepsTiming(p *cloudbuildpb.CloudbuildBuildTriggerBuildStepsTiming) *cloudbuild.BuildTriggerBuildStepsTiming {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.BuildTriggerBuildStepsTiming{
		StartTime: dcl.StringOrNil(p.StartTime),
		EndTime:   dcl.StringOrNil(p.EndTime),
	}
	return obj
}

// ProtoToBuildTriggerBuildStepsPullTiming converts a BuildTriggerBuildStepsPullTiming resource from its proto representation.
func ProtoToCloudbuildBuildTriggerBuildStepsPullTiming(p *cloudbuildpb.CloudbuildBuildTriggerBuildStepsPullTiming) *cloudbuild.BuildTriggerBuildStepsPullTiming {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.BuildTriggerBuildStepsPullTiming{
		StartTime: dcl.StringOrNil(p.StartTime),
		EndTime:   dcl.StringOrNil(p.EndTime),
	}
	return obj
}

// ProtoToBuildTriggerBuildSource converts a BuildTriggerBuildSource resource from its proto representation.
func ProtoToCloudbuildBuildTriggerBuildSource(p *cloudbuildpb.CloudbuildBuildTriggerBuildSource) *cloudbuild.BuildTriggerBuildSource {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.BuildTriggerBuildSource{
		StorageSource: ProtoToCloudbuildBuildTriggerBuildSourceStorageSource(p.GetStorageSource()),
		RepoSource:    ProtoToCloudbuildBuildTriggerBuildSourceRepoSource(p.GetRepoSource()),
	}
	return obj
}

// ProtoToBuildTriggerBuildSourceStorageSource converts a BuildTriggerBuildSourceStorageSource resource from its proto representation.
func ProtoToCloudbuildBuildTriggerBuildSourceStorageSource(p *cloudbuildpb.CloudbuildBuildTriggerBuildSourceStorageSource) *cloudbuild.BuildTriggerBuildSourceStorageSource {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.BuildTriggerBuildSourceStorageSource{
		Bucket:     dcl.StringOrNil(p.Bucket),
		Object:     dcl.StringOrNil(p.Object),
		Generation: dcl.StringOrNil(p.Generation),
	}
	return obj
}

// ProtoToBuildTriggerBuildSourceRepoSource converts a BuildTriggerBuildSourceRepoSource resource from its proto representation.
func ProtoToCloudbuildBuildTriggerBuildSourceRepoSource(p *cloudbuildpb.CloudbuildBuildTriggerBuildSourceRepoSource) *cloudbuild.BuildTriggerBuildSourceRepoSource {
	if p == nil {
		return nil
	}
	obj := &cloudbuild.BuildTriggerBuildSourceRepoSource{
		ProjectId:   dcl.StringOrNil(p.ProjectId),
		RepoName:    dcl.StringOrNil(p.RepoName),
		BranchName:  dcl.StringOrNil(p.BranchName),
		TagName:     dcl.StringOrNil(p.TagName),
		CommitSha:   dcl.StringOrNil(p.CommitSha),
		Dir:         dcl.StringOrNil(p.Dir),
		InvertRegex: dcl.Bool(p.InvertRegex),
	}
	return obj
}

// ProtoToBuildTrigger converts a BuildTrigger resource from its proto representation.
func ProtoToBuildTrigger(p *cloudbuildpb.CloudbuildBuildTrigger) *cloudbuild.BuildTrigger {
	obj := &cloudbuild.BuildTrigger{
		Name:            dcl.StringOrNil(p.Name),
		Description:     dcl.StringOrNil(p.Description),
		Disabled:        dcl.Bool(p.Disabled),
		Filename:        dcl.StringOrNil(p.Filename),
		TriggerTemplate: ProtoToCloudbuildBuildTriggerTriggerTemplate(p.GetTriggerTemplate()),
		Github:          ProtoToCloudbuildBuildTriggerGithub(p.GetGithub()),
		Project:         dcl.StringOrNil(p.Project),
		Build:           ProtoToCloudbuildBuildTriggerBuild(p.GetBuild()),
		Id:              dcl.StringOrNil(p.Id),
		CreateTime:      dcl.StringOrNil(p.CreateTime),
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	for _, r := range p.GetIgnoredFiles() {
		obj.IgnoredFiles = append(obj.IgnoredFiles, r)
	}
	for _, r := range p.GetIncludedFiles() {
		obj.IncludedFiles = append(obj.IncludedFiles, r)
	}
	return obj
}

// BuildTriggerGithubPullRequestCommentControlEnumToProto converts a BuildTriggerGithubPullRequestCommentControlEnum enum to its proto representation.
func CloudbuildBuildTriggerGithubPullRequestCommentControlEnumToProto(e *cloudbuild.BuildTriggerGithubPullRequestCommentControlEnum) cloudbuildpb.CloudbuildBuildTriggerGithubPullRequestCommentControlEnum {
	if e == nil {
		return cloudbuildpb.CloudbuildBuildTriggerGithubPullRequestCommentControlEnum(0)
	}
	if v, ok := cloudbuildpb.CloudbuildBuildTriggerGithubPullRequestCommentControlEnum_value["BuildTriggerGithubPullRequestCommentControlEnum"+string(*e)]; ok {
		return cloudbuildpb.CloudbuildBuildTriggerGithubPullRequestCommentControlEnum(v)
	}
	return cloudbuildpb.CloudbuildBuildTriggerGithubPullRequestCommentControlEnum(0)
}

// BuildTriggerBuildStepsStatusEnumToProto converts a BuildTriggerBuildStepsStatusEnum enum to its proto representation.
func CloudbuildBuildTriggerBuildStepsStatusEnumToProto(e *cloudbuild.BuildTriggerBuildStepsStatusEnum) cloudbuildpb.CloudbuildBuildTriggerBuildStepsStatusEnum {
	if e == nil {
		return cloudbuildpb.CloudbuildBuildTriggerBuildStepsStatusEnum(0)
	}
	if v, ok := cloudbuildpb.CloudbuildBuildTriggerBuildStepsStatusEnum_value["BuildTriggerBuildStepsStatusEnum"+string(*e)]; ok {
		return cloudbuildpb.CloudbuildBuildTriggerBuildStepsStatusEnum(v)
	}
	return cloudbuildpb.CloudbuildBuildTriggerBuildStepsStatusEnum(0)
}

// BuildTriggerTriggerTemplateToProto converts a BuildTriggerTriggerTemplate resource to its proto representation.
func CloudbuildBuildTriggerTriggerTemplateToProto(o *cloudbuild.BuildTriggerTriggerTemplate) *cloudbuildpb.CloudbuildBuildTriggerTriggerTemplate {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildBuildTriggerTriggerTemplate{
		ProjectId:   dcl.ValueOrEmptyString(o.ProjectId),
		RepoName:    dcl.ValueOrEmptyString(o.RepoName),
		BranchName:  dcl.ValueOrEmptyString(o.BranchName),
		TagName:     dcl.ValueOrEmptyString(o.TagName),
		CommitSha:   dcl.ValueOrEmptyString(o.CommitSha),
		Dir:         dcl.ValueOrEmptyString(o.Dir),
		InvertRegex: dcl.ValueOrEmptyBool(o.InvertRegex),
	}
	return p
}

// BuildTriggerGithubToProto converts a BuildTriggerGithub resource to its proto representation.
func CloudbuildBuildTriggerGithubToProto(o *cloudbuild.BuildTriggerGithub) *cloudbuildpb.CloudbuildBuildTriggerGithub {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildBuildTriggerGithub{
		Owner:       dcl.ValueOrEmptyString(o.Owner),
		Name:        dcl.ValueOrEmptyString(o.Name),
		PullRequest: CloudbuildBuildTriggerGithubPullRequestToProto(o.PullRequest),
		Push:        CloudbuildBuildTriggerGithubPushToProto(o.Push),
	}
	return p
}

// BuildTriggerGithubPullRequestToProto converts a BuildTriggerGithubPullRequest resource to its proto representation.
func CloudbuildBuildTriggerGithubPullRequestToProto(o *cloudbuild.BuildTriggerGithubPullRequest) *cloudbuildpb.CloudbuildBuildTriggerGithubPullRequest {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildBuildTriggerGithubPullRequest{
		Branch:         dcl.ValueOrEmptyString(o.Branch),
		CommentControl: CloudbuildBuildTriggerGithubPullRequestCommentControlEnumToProto(o.CommentControl),
		InvertRegex:    dcl.ValueOrEmptyBool(o.InvertRegex),
	}
	return p
}

// BuildTriggerGithubPushToProto converts a BuildTriggerGithubPush resource to its proto representation.
func CloudbuildBuildTriggerGithubPushToProto(o *cloudbuild.BuildTriggerGithubPush) *cloudbuildpb.CloudbuildBuildTriggerGithubPush {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildBuildTriggerGithubPush{
		Branch:      dcl.ValueOrEmptyString(o.Branch),
		Tag:         dcl.ValueOrEmptyString(o.Tag),
		InvertRegex: dcl.ValueOrEmptyBool(o.InvertRegex),
	}
	return p
}

// BuildTriggerBuildToProto converts a BuildTriggerBuild resource to its proto representation.
func CloudbuildBuildTriggerBuildToProto(o *cloudbuild.BuildTriggerBuild) *cloudbuildpb.CloudbuildBuildTriggerBuild {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildBuildTriggerBuild{
		QueueTtl:   dcl.ValueOrEmptyString(o.QueueTtl),
		LogsBucket: dcl.ValueOrEmptyString(o.LogsBucket),
		Timeout:    dcl.ValueOrEmptyString(o.Timeout),
		Source:     CloudbuildBuildTriggerBuildSourceToProto(o.Source),
	}
	for _, r := range o.Tags {
		p.Tags = append(p.Tags, r)
	}
	for _, r := range o.Images {
		p.Images = append(p.Images, r)
	}
	p.Substitutions = make(map[string]string)
	for k, r := range o.Substitutions {
		p.Substitutions[k] = r
	}
	for _, r := range o.Secrets {
		p.Secrets = append(p.Secrets, CloudbuildBuildTriggerBuildSecretsToProto(&r))
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, CloudbuildBuildTriggerBuildStepsToProto(&r))
	}
	return p
}

// BuildTriggerBuildSecretsToProto converts a BuildTriggerBuildSecrets resource to its proto representation.
func CloudbuildBuildTriggerBuildSecretsToProto(o *cloudbuild.BuildTriggerBuildSecrets) *cloudbuildpb.CloudbuildBuildTriggerBuildSecrets {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildBuildTriggerBuildSecrets{
		KmsKeyName: dcl.ValueOrEmptyString(o.KmsKeyName),
	}
	p.SecretEnv = make(map[string]string)
	for k, r := range o.SecretEnv {
		p.SecretEnv[k] = r
	}
	return p
}

// BuildTriggerBuildStepsToProto converts a BuildTriggerBuildSteps resource to its proto representation.
func CloudbuildBuildTriggerBuildStepsToProto(o *cloudbuild.BuildTriggerBuildSteps) *cloudbuildpb.CloudbuildBuildTriggerBuildSteps {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildBuildTriggerBuildSteps{
		Name:       dcl.ValueOrEmptyString(o.Name),
		Dir:        dcl.ValueOrEmptyString(o.Dir),
		Id:         dcl.ValueOrEmptyString(o.Id),
		Entrypoint: dcl.ValueOrEmptyString(o.Entrypoint),
		Timing:     CloudbuildBuildTriggerBuildStepsTimingToProto(o.Timing),
		PullTiming: CloudbuildBuildTriggerBuildStepsPullTimingToProto(o.PullTiming),
		Timeout:    dcl.ValueOrEmptyString(o.Timeout),
		Status:     CloudbuildBuildTriggerBuildStepsStatusEnumToProto(o.Status),
	}
	for _, r := range o.Env {
		p.Env = append(p.Env, r)
	}
	for _, r := range o.Args {
		p.Args = append(p.Args, r)
	}
	for _, r := range o.WaitFor {
		p.WaitFor = append(p.WaitFor, r)
	}
	for _, r := range o.SecretEnv {
		p.SecretEnv = append(p.SecretEnv, r)
	}
	for _, r := range o.Volumes {
		p.Volumes = append(p.Volumes, CloudbuildBuildTriggerBuildStepsVolumesToProto(&r))
	}
	return p
}

// BuildTriggerBuildStepsVolumesToProto converts a BuildTriggerBuildStepsVolumes resource to its proto representation.
func CloudbuildBuildTriggerBuildStepsVolumesToProto(o *cloudbuild.BuildTriggerBuildStepsVolumes) *cloudbuildpb.CloudbuildBuildTriggerBuildStepsVolumes {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildBuildTriggerBuildStepsVolumes{
		Name: dcl.ValueOrEmptyString(o.Name),
		Path: dcl.ValueOrEmptyString(o.Path),
	}
	return p
}

// BuildTriggerBuildStepsTimingToProto converts a BuildTriggerBuildStepsTiming resource to its proto representation.
func CloudbuildBuildTriggerBuildStepsTimingToProto(o *cloudbuild.BuildTriggerBuildStepsTiming) *cloudbuildpb.CloudbuildBuildTriggerBuildStepsTiming {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildBuildTriggerBuildStepsTiming{
		StartTime: dcl.ValueOrEmptyString(o.StartTime),
		EndTime:   dcl.ValueOrEmptyString(o.EndTime),
	}
	return p
}

// BuildTriggerBuildStepsPullTimingToProto converts a BuildTriggerBuildStepsPullTiming resource to its proto representation.
func CloudbuildBuildTriggerBuildStepsPullTimingToProto(o *cloudbuild.BuildTriggerBuildStepsPullTiming) *cloudbuildpb.CloudbuildBuildTriggerBuildStepsPullTiming {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildBuildTriggerBuildStepsPullTiming{
		StartTime: dcl.ValueOrEmptyString(o.StartTime),
		EndTime:   dcl.ValueOrEmptyString(o.EndTime),
	}
	return p
}

// BuildTriggerBuildSourceToProto converts a BuildTriggerBuildSource resource to its proto representation.
func CloudbuildBuildTriggerBuildSourceToProto(o *cloudbuild.BuildTriggerBuildSource) *cloudbuildpb.CloudbuildBuildTriggerBuildSource {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildBuildTriggerBuildSource{
		StorageSource: CloudbuildBuildTriggerBuildSourceStorageSourceToProto(o.StorageSource),
		RepoSource:    CloudbuildBuildTriggerBuildSourceRepoSourceToProto(o.RepoSource),
	}
	return p
}

// BuildTriggerBuildSourceStorageSourceToProto converts a BuildTriggerBuildSourceStorageSource resource to its proto representation.
func CloudbuildBuildTriggerBuildSourceStorageSourceToProto(o *cloudbuild.BuildTriggerBuildSourceStorageSource) *cloudbuildpb.CloudbuildBuildTriggerBuildSourceStorageSource {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildBuildTriggerBuildSourceStorageSource{
		Bucket:     dcl.ValueOrEmptyString(o.Bucket),
		Object:     dcl.ValueOrEmptyString(o.Object),
		Generation: dcl.ValueOrEmptyString(o.Generation),
	}
	return p
}

// BuildTriggerBuildSourceRepoSourceToProto converts a BuildTriggerBuildSourceRepoSource resource to its proto representation.
func CloudbuildBuildTriggerBuildSourceRepoSourceToProto(o *cloudbuild.BuildTriggerBuildSourceRepoSource) *cloudbuildpb.CloudbuildBuildTriggerBuildSourceRepoSource {
	if o == nil {
		return nil
	}
	p := &cloudbuildpb.CloudbuildBuildTriggerBuildSourceRepoSource{
		ProjectId:   dcl.ValueOrEmptyString(o.ProjectId),
		RepoName:    dcl.ValueOrEmptyString(o.RepoName),
		BranchName:  dcl.ValueOrEmptyString(o.BranchName),
		TagName:     dcl.ValueOrEmptyString(o.TagName),
		CommitSha:   dcl.ValueOrEmptyString(o.CommitSha),
		Dir:         dcl.ValueOrEmptyString(o.Dir),
		InvertRegex: dcl.ValueOrEmptyBool(o.InvertRegex),
	}
	p.Substitutions = make(map[string]string)
	for k, r := range o.Substitutions {
		p.Substitutions[k] = r
	}
	return p
}

// BuildTriggerToProto converts a BuildTrigger resource to its proto representation.
func BuildTriggerToProto(resource *cloudbuild.BuildTrigger) *cloudbuildpb.CloudbuildBuildTrigger {
	p := &cloudbuildpb.CloudbuildBuildTrigger{
		Name:            dcl.ValueOrEmptyString(resource.Name),
		Description:     dcl.ValueOrEmptyString(resource.Description),
		Disabled:        dcl.ValueOrEmptyBool(resource.Disabled),
		Filename:        dcl.ValueOrEmptyString(resource.Filename),
		TriggerTemplate: CloudbuildBuildTriggerTriggerTemplateToProto(resource.TriggerTemplate),
		Github:          CloudbuildBuildTriggerGithubToProto(resource.Github),
		Project:         dcl.ValueOrEmptyString(resource.Project),
		Build:           CloudbuildBuildTriggerBuildToProto(resource.Build),
		Id:              dcl.ValueOrEmptyString(resource.Id),
		CreateTime:      dcl.ValueOrEmptyString(resource.CreateTime),
	}
	for _, r := range resource.Tags {
		p.Tags = append(p.Tags, r)
	}
	for _, r := range resource.IgnoredFiles {
		p.IgnoredFiles = append(p.IgnoredFiles, r)
	}
	for _, r := range resource.IncludedFiles {
		p.IncludedFiles = append(p.IncludedFiles, r)
	}

	return p
}

// ApplyBuildTrigger handles the gRPC request by passing it to the underlying BuildTrigger Apply() method.
func (s *BuildTriggerServer) applyBuildTrigger(ctx context.Context, c *cloudbuild.Client, request *cloudbuildpb.ApplyCloudbuildBuildTriggerRequest) (*cloudbuildpb.CloudbuildBuildTrigger, error) {
	p := ProtoToBuildTrigger(request.GetResource())
	res, err := c.ApplyBuildTrigger(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BuildTriggerToProto(res)
	return r, nil
}

// ApplyBuildTrigger handles the gRPC request by passing it to the underlying BuildTrigger Apply() method.
func (s *BuildTriggerServer) ApplyCloudbuildBuildTrigger(ctx context.Context, request *cloudbuildpb.ApplyCloudbuildBuildTriggerRequest) (*cloudbuildpb.CloudbuildBuildTrigger, error) {
	cl, err := createConfigBuildTrigger(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyBuildTrigger(ctx, cl, request)
}

// DeleteBuildTrigger handles the gRPC request by passing it to the underlying BuildTrigger Delete() method.
func (s *BuildTriggerServer) DeleteCloudbuildBuildTrigger(ctx context.Context, request *cloudbuildpb.DeleteCloudbuildBuildTriggerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigBuildTrigger(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteBuildTrigger(ctx, ProtoToBuildTrigger(request.GetResource()))

}

// ListCloudbuildBuildTrigger handles the gRPC request by passing it to the underlying BuildTriggerList() method.
func (s *BuildTriggerServer) ListCloudbuildBuildTrigger(ctx context.Context, request *cloudbuildpb.ListCloudbuildBuildTriggerRequest) (*cloudbuildpb.ListCloudbuildBuildTriggerResponse, error) {
	cl, err := createConfigBuildTrigger(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBuildTrigger(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*cloudbuildpb.CloudbuildBuildTrigger
	for _, r := range resources.Items {
		rp := BuildTriggerToProto(r)
		protos = append(protos, rp)
	}
	return &cloudbuildpb.ListCloudbuildBuildTriggerResponse{Items: protos}, nil
}

func createConfigBuildTrigger(ctx context.Context, service_account_file string) (*cloudbuild.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudbuild.NewClient(conf), nil
}
