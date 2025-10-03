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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudbuild/beta/cloudbuild_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/beta"
)

// Server implements the gRPC interface for BuildTrigger.
type BuildTriggerServer struct{}

// ProtoToBuildTriggerGithubPullRequestCommentControlEnum converts a BuildTriggerGithubPullRequestCommentControlEnum enum from its proto representation.
func ProtoToCloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnum(e betapb.CloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnum) *beta.BuildTriggerGithubPullRequestCommentControlEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnum_name[int32(e)]; ok {
		e := beta.BuildTriggerGithubPullRequestCommentControlEnum(n[len("CloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnum"):])
		return &e
	}
	return nil
}

// ProtoToBuildTriggerBuildStepsStatusEnum converts a BuildTriggerBuildStepsStatusEnum enum from its proto representation.
func ProtoToCloudbuildBetaBuildTriggerBuildStepsStatusEnum(e betapb.CloudbuildBetaBuildTriggerBuildStepsStatusEnum) *beta.BuildTriggerBuildStepsStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudbuildBetaBuildTriggerBuildStepsStatusEnum_name[int32(e)]; ok {
		e := beta.BuildTriggerBuildStepsStatusEnum(n[len("CloudbuildBetaBuildTriggerBuildStepsStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToBuildTriggerTriggerTemplate converts a BuildTriggerTriggerTemplate resource from its proto representation.
func ProtoToCloudbuildBetaBuildTriggerTriggerTemplate(p *betapb.CloudbuildBetaBuildTriggerTriggerTemplate) *beta.BuildTriggerTriggerTemplate {
	if p == nil {
		return nil
	}
	obj := &beta.BuildTriggerTriggerTemplate{
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
func ProtoToCloudbuildBetaBuildTriggerGithub(p *betapb.CloudbuildBetaBuildTriggerGithub) *beta.BuildTriggerGithub {
	if p == nil {
		return nil
	}
	obj := &beta.BuildTriggerGithub{
		Owner:       dcl.StringOrNil(p.Owner),
		Name:        dcl.StringOrNil(p.Name),
		PullRequest: ProtoToCloudbuildBetaBuildTriggerGithubPullRequest(p.GetPullRequest()),
		Push:        ProtoToCloudbuildBetaBuildTriggerGithubPush(p.GetPush()),
	}
	return obj
}

// ProtoToBuildTriggerGithubPullRequest converts a BuildTriggerGithubPullRequest resource from its proto representation.
func ProtoToCloudbuildBetaBuildTriggerGithubPullRequest(p *betapb.CloudbuildBetaBuildTriggerGithubPullRequest) *beta.BuildTriggerGithubPullRequest {
	if p == nil {
		return nil
	}
	obj := &beta.BuildTriggerGithubPullRequest{
		Branch:         dcl.StringOrNil(p.Branch),
		CommentControl: ProtoToCloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnum(p.GetCommentControl()),
		InvertRegex:    dcl.Bool(p.InvertRegex),
	}
	return obj
}

// ProtoToBuildTriggerGithubPush converts a BuildTriggerGithubPush resource from its proto representation.
func ProtoToCloudbuildBetaBuildTriggerGithubPush(p *betapb.CloudbuildBetaBuildTriggerGithubPush) *beta.BuildTriggerGithubPush {
	if p == nil {
		return nil
	}
	obj := &beta.BuildTriggerGithubPush{
		Branch:      dcl.StringOrNil(p.Branch),
		Tag:         dcl.StringOrNil(p.Tag),
		InvertRegex: dcl.Bool(p.InvertRegex),
	}
	return obj
}

// ProtoToBuildTriggerBuild converts a BuildTriggerBuild resource from its proto representation.
func ProtoToCloudbuildBetaBuildTriggerBuild(p *betapb.CloudbuildBetaBuildTriggerBuild) *beta.BuildTriggerBuild {
	if p == nil {
		return nil
	}
	obj := &beta.BuildTriggerBuild{
		QueueTtl:   dcl.StringOrNil(p.QueueTtl),
		LogsBucket: dcl.StringOrNil(p.LogsBucket),
		Timeout:    dcl.StringOrNil(p.Timeout),
		Source:     ProtoToCloudbuildBetaBuildTriggerBuildSource(p.GetSource()),
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	for _, r := range p.GetImages() {
		obj.Images = append(obj.Images, r)
	}
	for _, r := range p.GetSecrets() {
		obj.Secrets = append(obj.Secrets, *ProtoToCloudbuildBetaBuildTriggerBuildSecrets(r))
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToCloudbuildBetaBuildTriggerBuildSteps(r))
	}
	return obj
}

// ProtoToBuildTriggerBuildSecrets converts a BuildTriggerBuildSecrets resource from its proto representation.
func ProtoToCloudbuildBetaBuildTriggerBuildSecrets(p *betapb.CloudbuildBetaBuildTriggerBuildSecrets) *beta.BuildTriggerBuildSecrets {
	if p == nil {
		return nil
	}
	obj := &beta.BuildTriggerBuildSecrets{
		KmsKeyName: dcl.StringOrNil(p.KmsKeyName),
	}
	return obj
}

// ProtoToBuildTriggerBuildSteps converts a BuildTriggerBuildSteps resource from its proto representation.
func ProtoToCloudbuildBetaBuildTriggerBuildSteps(p *betapb.CloudbuildBetaBuildTriggerBuildSteps) *beta.BuildTriggerBuildSteps {
	if p == nil {
		return nil
	}
	obj := &beta.BuildTriggerBuildSteps{
		Name:       dcl.StringOrNil(p.Name),
		Dir:        dcl.StringOrNil(p.Dir),
		Id:         dcl.StringOrNil(p.Id),
		Entrypoint: dcl.StringOrNil(p.Entrypoint),
		Timing:     ProtoToCloudbuildBetaBuildTriggerBuildStepsTiming(p.GetTiming()),
		PullTiming: ProtoToCloudbuildBetaBuildTriggerBuildStepsPullTiming(p.GetPullTiming()),
		Timeout:    dcl.StringOrNil(p.Timeout),
		Status:     ProtoToCloudbuildBetaBuildTriggerBuildStepsStatusEnum(p.GetStatus()),
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
		obj.Volumes = append(obj.Volumes, *ProtoToCloudbuildBetaBuildTriggerBuildStepsVolumes(r))
	}
	return obj
}

// ProtoToBuildTriggerBuildStepsVolumes converts a BuildTriggerBuildStepsVolumes resource from its proto representation.
func ProtoToCloudbuildBetaBuildTriggerBuildStepsVolumes(p *betapb.CloudbuildBetaBuildTriggerBuildStepsVolumes) *beta.BuildTriggerBuildStepsVolumes {
	if p == nil {
		return nil
	}
	obj := &beta.BuildTriggerBuildStepsVolumes{
		Name: dcl.StringOrNil(p.Name),
		Path: dcl.StringOrNil(p.Path),
	}
	return obj
}

// ProtoToBuildTriggerBuildStepsTiming converts a BuildTriggerBuildStepsTiming resource from its proto representation.
func ProtoToCloudbuildBetaBuildTriggerBuildStepsTiming(p *betapb.CloudbuildBetaBuildTriggerBuildStepsTiming) *beta.BuildTriggerBuildStepsTiming {
	if p == nil {
		return nil
	}
	obj := &beta.BuildTriggerBuildStepsTiming{
		StartTime: dcl.StringOrNil(p.StartTime),
		EndTime:   dcl.StringOrNil(p.EndTime),
	}
	return obj
}

// ProtoToBuildTriggerBuildStepsPullTiming converts a BuildTriggerBuildStepsPullTiming resource from its proto representation.
func ProtoToCloudbuildBetaBuildTriggerBuildStepsPullTiming(p *betapb.CloudbuildBetaBuildTriggerBuildStepsPullTiming) *beta.BuildTriggerBuildStepsPullTiming {
	if p == nil {
		return nil
	}
	obj := &beta.BuildTriggerBuildStepsPullTiming{
		StartTime: dcl.StringOrNil(p.StartTime),
		EndTime:   dcl.StringOrNil(p.EndTime),
	}
	return obj
}

// ProtoToBuildTriggerBuildSource converts a BuildTriggerBuildSource resource from its proto representation.
func ProtoToCloudbuildBetaBuildTriggerBuildSource(p *betapb.CloudbuildBetaBuildTriggerBuildSource) *beta.BuildTriggerBuildSource {
	if p == nil {
		return nil
	}
	obj := &beta.BuildTriggerBuildSource{
		StorageSource: ProtoToCloudbuildBetaBuildTriggerBuildSourceStorageSource(p.GetStorageSource()),
		RepoSource:    ProtoToCloudbuildBetaBuildTriggerBuildSourceRepoSource(p.GetRepoSource()),
	}
	return obj
}

// ProtoToBuildTriggerBuildSourceStorageSource converts a BuildTriggerBuildSourceStorageSource resource from its proto representation.
func ProtoToCloudbuildBetaBuildTriggerBuildSourceStorageSource(p *betapb.CloudbuildBetaBuildTriggerBuildSourceStorageSource) *beta.BuildTriggerBuildSourceStorageSource {
	if p == nil {
		return nil
	}
	obj := &beta.BuildTriggerBuildSourceStorageSource{
		Bucket:     dcl.StringOrNil(p.Bucket),
		Object:     dcl.StringOrNil(p.Object),
		Generation: dcl.StringOrNil(p.Generation),
	}
	return obj
}

// ProtoToBuildTriggerBuildSourceRepoSource converts a BuildTriggerBuildSourceRepoSource resource from its proto representation.
func ProtoToCloudbuildBetaBuildTriggerBuildSourceRepoSource(p *betapb.CloudbuildBetaBuildTriggerBuildSourceRepoSource) *beta.BuildTriggerBuildSourceRepoSource {
	if p == nil {
		return nil
	}
	obj := &beta.BuildTriggerBuildSourceRepoSource{
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
func ProtoToBuildTrigger(p *betapb.CloudbuildBetaBuildTrigger) *beta.BuildTrigger {
	obj := &beta.BuildTrigger{
		Name:            dcl.StringOrNil(p.Name),
		Description:     dcl.StringOrNil(p.Description),
		Disabled:        dcl.Bool(p.Disabled),
		Filename:        dcl.StringOrNil(p.Filename),
		TriggerTemplate: ProtoToCloudbuildBetaBuildTriggerTriggerTemplate(p.GetTriggerTemplate()),
		Github:          ProtoToCloudbuildBetaBuildTriggerGithub(p.GetGithub()),
		Project:         dcl.StringOrNil(p.Project),
		Build:           ProtoToCloudbuildBetaBuildTriggerBuild(p.GetBuild()),
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
func CloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnumToProto(e *beta.BuildTriggerGithubPullRequestCommentControlEnum) betapb.CloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnum {
	if e == nil {
		return betapb.CloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnum(0)
	}
	if v, ok := betapb.CloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnum_value["BuildTriggerGithubPullRequestCommentControlEnum"+string(*e)]; ok {
		return betapb.CloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnum(v)
	}
	return betapb.CloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnum(0)
}

// BuildTriggerBuildStepsStatusEnumToProto converts a BuildTriggerBuildStepsStatusEnum enum to its proto representation.
func CloudbuildBetaBuildTriggerBuildStepsStatusEnumToProto(e *beta.BuildTriggerBuildStepsStatusEnum) betapb.CloudbuildBetaBuildTriggerBuildStepsStatusEnum {
	if e == nil {
		return betapb.CloudbuildBetaBuildTriggerBuildStepsStatusEnum(0)
	}
	if v, ok := betapb.CloudbuildBetaBuildTriggerBuildStepsStatusEnum_value["BuildTriggerBuildStepsStatusEnum"+string(*e)]; ok {
		return betapb.CloudbuildBetaBuildTriggerBuildStepsStatusEnum(v)
	}
	return betapb.CloudbuildBetaBuildTriggerBuildStepsStatusEnum(0)
}

// BuildTriggerTriggerTemplateToProto converts a BuildTriggerTriggerTemplate resource to its proto representation.
func CloudbuildBetaBuildTriggerTriggerTemplateToProto(o *beta.BuildTriggerTriggerTemplate) *betapb.CloudbuildBetaBuildTriggerTriggerTemplate {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaBuildTriggerTriggerTemplate{
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
func CloudbuildBetaBuildTriggerGithubToProto(o *beta.BuildTriggerGithub) *betapb.CloudbuildBetaBuildTriggerGithub {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaBuildTriggerGithub{
		Owner:       dcl.ValueOrEmptyString(o.Owner),
		Name:        dcl.ValueOrEmptyString(o.Name),
		PullRequest: CloudbuildBetaBuildTriggerGithubPullRequestToProto(o.PullRequest),
		Push:        CloudbuildBetaBuildTriggerGithubPushToProto(o.Push),
	}
	return p
}

// BuildTriggerGithubPullRequestToProto converts a BuildTriggerGithubPullRequest resource to its proto representation.
func CloudbuildBetaBuildTriggerGithubPullRequestToProto(o *beta.BuildTriggerGithubPullRequest) *betapb.CloudbuildBetaBuildTriggerGithubPullRequest {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaBuildTriggerGithubPullRequest{
		Branch:         dcl.ValueOrEmptyString(o.Branch),
		CommentControl: CloudbuildBetaBuildTriggerGithubPullRequestCommentControlEnumToProto(o.CommentControl),
		InvertRegex:    dcl.ValueOrEmptyBool(o.InvertRegex),
	}
	return p
}

// BuildTriggerGithubPushToProto converts a BuildTriggerGithubPush resource to its proto representation.
func CloudbuildBetaBuildTriggerGithubPushToProto(o *beta.BuildTriggerGithubPush) *betapb.CloudbuildBetaBuildTriggerGithubPush {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaBuildTriggerGithubPush{
		Branch:      dcl.ValueOrEmptyString(o.Branch),
		Tag:         dcl.ValueOrEmptyString(o.Tag),
		InvertRegex: dcl.ValueOrEmptyBool(o.InvertRegex),
	}
	return p
}

// BuildTriggerBuildToProto converts a BuildTriggerBuild resource to its proto representation.
func CloudbuildBetaBuildTriggerBuildToProto(o *beta.BuildTriggerBuild) *betapb.CloudbuildBetaBuildTriggerBuild {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaBuildTriggerBuild{
		QueueTtl:   dcl.ValueOrEmptyString(o.QueueTtl),
		LogsBucket: dcl.ValueOrEmptyString(o.LogsBucket),
		Timeout:    dcl.ValueOrEmptyString(o.Timeout),
		Source:     CloudbuildBetaBuildTriggerBuildSourceToProto(o.Source),
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
		p.Secrets = append(p.Secrets, CloudbuildBetaBuildTriggerBuildSecretsToProto(&r))
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, CloudbuildBetaBuildTriggerBuildStepsToProto(&r))
	}
	return p
}

// BuildTriggerBuildSecretsToProto converts a BuildTriggerBuildSecrets resource to its proto representation.
func CloudbuildBetaBuildTriggerBuildSecretsToProto(o *beta.BuildTriggerBuildSecrets) *betapb.CloudbuildBetaBuildTriggerBuildSecrets {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaBuildTriggerBuildSecrets{
		KmsKeyName: dcl.ValueOrEmptyString(o.KmsKeyName),
	}
	p.SecretEnv = make(map[string]string)
	for k, r := range o.SecretEnv {
		p.SecretEnv[k] = r
	}
	return p
}

// BuildTriggerBuildStepsToProto converts a BuildTriggerBuildSteps resource to its proto representation.
func CloudbuildBetaBuildTriggerBuildStepsToProto(o *beta.BuildTriggerBuildSteps) *betapb.CloudbuildBetaBuildTriggerBuildSteps {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaBuildTriggerBuildSteps{
		Name:       dcl.ValueOrEmptyString(o.Name),
		Dir:        dcl.ValueOrEmptyString(o.Dir),
		Id:         dcl.ValueOrEmptyString(o.Id),
		Entrypoint: dcl.ValueOrEmptyString(o.Entrypoint),
		Timing:     CloudbuildBetaBuildTriggerBuildStepsTimingToProto(o.Timing),
		PullTiming: CloudbuildBetaBuildTriggerBuildStepsPullTimingToProto(o.PullTiming),
		Timeout:    dcl.ValueOrEmptyString(o.Timeout),
		Status:     CloudbuildBetaBuildTriggerBuildStepsStatusEnumToProto(o.Status),
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
		p.Volumes = append(p.Volumes, CloudbuildBetaBuildTriggerBuildStepsVolumesToProto(&r))
	}
	return p
}

// BuildTriggerBuildStepsVolumesToProto converts a BuildTriggerBuildStepsVolumes resource to its proto representation.
func CloudbuildBetaBuildTriggerBuildStepsVolumesToProto(o *beta.BuildTriggerBuildStepsVolumes) *betapb.CloudbuildBetaBuildTriggerBuildStepsVolumes {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaBuildTriggerBuildStepsVolumes{
		Name: dcl.ValueOrEmptyString(o.Name),
		Path: dcl.ValueOrEmptyString(o.Path),
	}
	return p
}

// BuildTriggerBuildStepsTimingToProto converts a BuildTriggerBuildStepsTiming resource to its proto representation.
func CloudbuildBetaBuildTriggerBuildStepsTimingToProto(o *beta.BuildTriggerBuildStepsTiming) *betapb.CloudbuildBetaBuildTriggerBuildStepsTiming {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaBuildTriggerBuildStepsTiming{
		StartTime: dcl.ValueOrEmptyString(o.StartTime),
		EndTime:   dcl.ValueOrEmptyString(o.EndTime),
	}
	return p
}

// BuildTriggerBuildStepsPullTimingToProto converts a BuildTriggerBuildStepsPullTiming resource to its proto representation.
func CloudbuildBetaBuildTriggerBuildStepsPullTimingToProto(o *beta.BuildTriggerBuildStepsPullTiming) *betapb.CloudbuildBetaBuildTriggerBuildStepsPullTiming {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaBuildTriggerBuildStepsPullTiming{
		StartTime: dcl.ValueOrEmptyString(o.StartTime),
		EndTime:   dcl.ValueOrEmptyString(o.EndTime),
	}
	return p
}

// BuildTriggerBuildSourceToProto converts a BuildTriggerBuildSource resource to its proto representation.
func CloudbuildBetaBuildTriggerBuildSourceToProto(o *beta.BuildTriggerBuildSource) *betapb.CloudbuildBetaBuildTriggerBuildSource {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaBuildTriggerBuildSource{
		StorageSource: CloudbuildBetaBuildTriggerBuildSourceStorageSourceToProto(o.StorageSource),
		RepoSource:    CloudbuildBetaBuildTriggerBuildSourceRepoSourceToProto(o.RepoSource),
	}
	return p
}

// BuildTriggerBuildSourceStorageSourceToProto converts a BuildTriggerBuildSourceStorageSource resource to its proto representation.
func CloudbuildBetaBuildTriggerBuildSourceStorageSourceToProto(o *beta.BuildTriggerBuildSourceStorageSource) *betapb.CloudbuildBetaBuildTriggerBuildSourceStorageSource {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaBuildTriggerBuildSourceStorageSource{
		Bucket:     dcl.ValueOrEmptyString(o.Bucket),
		Object:     dcl.ValueOrEmptyString(o.Object),
		Generation: dcl.ValueOrEmptyString(o.Generation),
	}
	return p
}

// BuildTriggerBuildSourceRepoSourceToProto converts a BuildTriggerBuildSourceRepoSource resource to its proto representation.
func CloudbuildBetaBuildTriggerBuildSourceRepoSourceToProto(o *beta.BuildTriggerBuildSourceRepoSource) *betapb.CloudbuildBetaBuildTriggerBuildSourceRepoSource {
	if o == nil {
		return nil
	}
	p := &betapb.CloudbuildBetaBuildTriggerBuildSourceRepoSource{
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
func BuildTriggerToProto(resource *beta.BuildTrigger) *betapb.CloudbuildBetaBuildTrigger {
	p := &betapb.CloudbuildBetaBuildTrigger{
		Name:            dcl.ValueOrEmptyString(resource.Name),
		Description:     dcl.ValueOrEmptyString(resource.Description),
		Disabled:        dcl.ValueOrEmptyBool(resource.Disabled),
		Filename:        dcl.ValueOrEmptyString(resource.Filename),
		TriggerTemplate: CloudbuildBetaBuildTriggerTriggerTemplateToProto(resource.TriggerTemplate),
		Github:          CloudbuildBetaBuildTriggerGithubToProto(resource.Github),
		Project:         dcl.ValueOrEmptyString(resource.Project),
		Build:           CloudbuildBetaBuildTriggerBuildToProto(resource.Build),
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
func (s *BuildTriggerServer) applyBuildTrigger(ctx context.Context, c *beta.Client, request *betapb.ApplyCloudbuildBetaBuildTriggerRequest) (*betapb.CloudbuildBetaBuildTrigger, error) {
	p := ProtoToBuildTrigger(request.GetResource())
	res, err := c.ApplyBuildTrigger(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BuildTriggerToProto(res)
	return r, nil
}

// ApplyBuildTrigger handles the gRPC request by passing it to the underlying BuildTrigger Apply() method.
func (s *BuildTriggerServer) ApplyCloudbuildBetaBuildTrigger(ctx context.Context, request *betapb.ApplyCloudbuildBetaBuildTriggerRequest) (*betapb.CloudbuildBetaBuildTrigger, error) {
	cl, err := createConfigBuildTrigger(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyBuildTrigger(ctx, cl, request)
}

// DeleteBuildTrigger handles the gRPC request by passing it to the underlying BuildTrigger Delete() method.
func (s *BuildTriggerServer) DeleteCloudbuildBetaBuildTrigger(ctx context.Context, request *betapb.DeleteCloudbuildBetaBuildTriggerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigBuildTrigger(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteBuildTrigger(ctx, ProtoToBuildTrigger(request.GetResource()))

}

// ListCloudbuildBetaBuildTrigger handles the gRPC request by passing it to the underlying BuildTriggerList() method.
func (s *BuildTriggerServer) ListCloudbuildBetaBuildTrigger(ctx context.Context, request *betapb.ListCloudbuildBetaBuildTriggerRequest) (*betapb.ListCloudbuildBetaBuildTriggerResponse, error) {
	cl, err := createConfigBuildTrigger(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBuildTrigger(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.CloudbuildBetaBuildTrigger
	for _, r := range resources.Items {
		rp := BuildTriggerToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListCloudbuildBetaBuildTriggerResponse{Items: protos}, nil
}

func createConfigBuildTrigger(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
