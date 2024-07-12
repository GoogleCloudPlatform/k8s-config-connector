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

package securesourcemanager

import (
	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securesourcemanager/v1alpha1"
)

func Instance_HostConfig_FromProto(mapCtx *MapContext, in *pb.Instance_HostConfig) *krm.Instance_HostConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_HostConfig{}
	out.Html = LazyPtr(in.GetHtml())
	out.Api = LazyPtr(in.GetApi())
	out.GitHttp = LazyPtr(in.GetGitHttp())
	out.GitSsh = LazyPtr(in.GetGitSsh())
	return out
}
func Instance_HostConfig_ToProto(mapCtx *MapContext, in *krm.Instance_HostConfig) *pb.Instance_HostConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_HostConfig{}
	out.Html = ValueOf(in.Html)
	out.Api = ValueOf(in.Api)
	out.GitHttp = ValueOf(in.GitHttp)
	out.GitSsh = ValueOf(in.GitSsh)
	return out
}
func OperationMetadata_FromProto(mapCtx *MapContext, in *pb.OperationMetadata) *krm.OperationMetadata {
	if in == nil {
		return nil
	}
	out := &krm.OperationMetadata{}
	out.CreateTime = OperationMetadata_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.EndTime = OperationMetadata_EndTime_FromProto(mapCtx, in.GetEndTime())
	out.Target = LazyPtr(in.GetTarget())
	out.Verb = LazyPtr(in.GetVerb())
	out.StatusMessage = LazyPtr(in.GetStatusMessage())
	out.RequestedCancellation = LazyPtr(in.GetRequestedCancellation())
	out.ApiVersion = LazyPtr(in.GetApiVersion())
	return out
}
func OperationMetadata_ToProto(mapCtx *MapContext, in *krm.OperationMetadata) *pb.OperationMetadata {
	if in == nil {
		return nil
	}
	out := &pb.OperationMetadata{}
	out.CreateTime = OperationMetadata_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.EndTime = OperationMetadata_EndTime_ToProto(mapCtx, in.EndTime)
	out.Target = ValueOf(in.Target)
	out.Verb = ValueOf(in.Verb)
	out.StatusMessage = ValueOf(in.StatusMessage)
	out.RequestedCancellation = ValueOf(in.RequestedCancellation)
	out.ApiVersion = ValueOf(in.ApiVersion)
	return out
}
func Repository_FromProto(mapCtx *MapContext, in *pb.Repository) *krm.Repository {
	if in == nil {
		return nil
	}
	out := &krm.Repository{}
	out.Name = LazyPtr(in.GetName())
	out.Description = LazyPtr(in.GetDescription())
	out.Instance = LazyPtr(in.GetInstance())
	out.Uid = LazyPtr(in.GetUid())
	out.CreateTime = Repository_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = Repository_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	out.Etag = LazyPtr(in.GetEtag())
	out.Uris = Repository_URIs_FromProto(mapCtx, in.GetUris())
	out.InitialConfig = Repository_InitialConfig_FromProto(mapCtx, in.GetInitialConfig())
	return out
}
func Repository_ToProto(mapCtx *MapContext, in *krm.Repository) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	out.Name = ValueOf(in.Name)
	out.Description = ValueOf(in.Description)
	out.Instance = ValueOf(in.Instance)
	out.Uid = ValueOf(in.Uid)
	out.CreateTime = Repository_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = Repository_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	out.Etag = ValueOf(in.Etag)
	out.Uris = Repository_URIs_ToProto(mapCtx, in.Uris)
	out.InitialConfig = Repository_InitialConfig_ToProto(mapCtx, in.InitialConfig)
	return out
}
func Repository_InitialConfig_FromProto(mapCtx *MapContext, in *pb.Repository_InitialConfig) *krm.Repository_InitialConfig {
	if in == nil {
		return nil
	}
	out := &krm.Repository_InitialConfig{}
	out.DefaultBranch = LazyPtr(in.GetDefaultBranch())
	out.Gitignores = in.Gitignores
	out.License = LazyPtr(in.GetLicense())
	out.Readme = LazyPtr(in.GetReadme())
	return out
}
func Repository_InitialConfig_ToProto(mapCtx *MapContext, in *krm.Repository_InitialConfig) *pb.Repository_InitialConfig {
	if in == nil {
		return nil
	}
	out := &pb.Repository_InitialConfig{}
	out.DefaultBranch = ValueOf(in.DefaultBranch)
	out.Gitignores = in.Gitignores
	out.License = ValueOf(in.License)
	out.Readme = ValueOf(in.Readme)
	return out
}
func Repository_URIs_FromProto(mapCtx *MapContext, in *pb.Repository_URIs) *krm.Repository_URIs {
	if in == nil {
		return nil
	}
	out := &krm.Repository_URIs{}
	out.Html = LazyPtr(in.GetHtml())
	out.GitHttps = LazyPtr(in.GetGitHttps())
	out.Api = LazyPtr(in.GetApi())
	return out
}
func Repository_URIs_ToProto(mapCtx *MapContext, in *krm.Repository_URIs) *pb.Repository_URIs {
	if in == nil {
		return nil
	}
	out := &pb.Repository_URIs{}
	out.Html = ValueOf(in.Html)
	out.GitHttps = ValueOf(in.GitHttps)
	out.Api = ValueOf(in.Api)
	return out
}
func SecureSourceManagerInstanceObservedState_FromProto(mapCtx *MapContext, in *pb.Instance) *krm.SecureSourceManagerInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerInstanceObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.State = Enum_FromProto(mapCtx, in.State)
	out.StateNote = Enum_FromProto(mapCtx, in.StateNote)
	// MISSING: KmsKey
	out.HostConfig = Instance_HostConfig_FromProto(mapCtx, in.GetHostConfig())
	return out
}
func SecureSourceManagerInstanceObservedState_ToProto(mapCtx *MapContext, in *krm.SecureSourceManagerInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.State = Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StateNote = Enum_ToProto[pb.Instance_StateNote](mapCtx, in.StateNote)
	// MISSING: KmsKey
	out.HostConfig = Instance_HostConfig_ToProto(mapCtx, in.HostConfig)
	return out
}
func SecureSourceManagerInstanceSpec_FromProto(mapCtx *MapContext, in *pb.Instance) *krm.SecureSourceManagerInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerInstanceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: State
	// MISSING: StateNote
	out.KmsKey = LazyPtr(in.GetKmsKey())
	// MISSING: HostConfig
	return out
}
func SecureSourceManagerInstanceSpec_ToProto(mapCtx *MapContext, in *krm.SecureSourceManagerInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: State
	// MISSING: StateNote
	out.KmsKey = ValueOf(in.KmsKey)
	// MISSING: HostConfig
	return out
}
