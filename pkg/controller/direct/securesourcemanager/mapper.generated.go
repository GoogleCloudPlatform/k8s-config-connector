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

package securesourcemanager

import (
	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securesourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Instance_HostConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_HostConfig) *krm.Instance_HostConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Instance_HostConfigObservedState{}
	out.HTML = direct.LazyPtr(in.GetHtml())
	out.Api = direct.LazyPtr(in.GetApi())
	out.GitHTTP = direct.LazyPtr(in.GetGitHttp())
	out.GitSSH = direct.LazyPtr(in.GetGitSsh())
	return out
}
func Instance_HostConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Instance_HostConfigObservedState) *pb.Instance_HostConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_HostConfig{}
	out.Html = direct.ValueOf(in.HTML)
	out.Api = direct.ValueOf(in.Api)
	out.GitHttp = direct.ValueOf(in.GitHTTP)
	out.GitSsh = direct.ValueOf(in.GitSSH)
	return out
}
func Instance_PrivateConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PrivateConfig) *krm.Instance_PrivateConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_PrivateConfig{}
	out.IsPrivate = direct.LazyPtr(in.GetIsPrivate())
	if in.GetCaPool() != "" {
		out.CaPoolRef = &refs.PrivateCACAPoolRef{External: in.GetCaPool()}
	}
	// MISSING: PscAllowedProjects
	return out
}
func Instance_PrivateConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_PrivateConfig) *pb.Instance_PrivateConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PrivateConfig{}
	out.IsPrivate = direct.ValueOf(in.IsPrivate)
	if in.CaPoolRef != nil {
		out.CaPool = in.CaPoolRef.External
	}
	// MISSING: PscAllowedProjects
	return out
}
func Instance_PrivateConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_PrivateConfig) *krm.Instance_PrivateConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Instance_PrivateConfigObservedState{}
	out.HTTPServiceAttachment = direct.LazyPtr(in.GetHttpServiceAttachment())
	out.SSHServiceAttachment = direct.LazyPtr(in.GetSshServiceAttachment())
	// MISSING: PscAllowedProjects
	return out
}
func Instance_PrivateConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Instance_PrivateConfigObservedState) *pb.Instance_PrivateConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_PrivateConfig{}
	out.HttpServiceAttachment = direct.ValueOf(in.HTTPServiceAttachment)
	out.SshServiceAttachment = direct.ValueOf(in.SSHServiceAttachment)
	// MISSING: PscAllowedProjects
	return out
}
func Repository_InitialConfig_FromProto(mapCtx *direct.MapContext, in *pb.Repository_InitialConfig) *krm.Repository_InitialConfig {
	if in == nil {
		return nil
	}
	out := &krm.Repository_InitialConfig{}
	out.DefaultBranch = direct.LazyPtr(in.GetDefaultBranch())
	out.Gitignores = in.Gitignores
	out.License = direct.LazyPtr(in.GetLicense())
	out.Readme = direct.LazyPtr(in.GetReadme())
	return out
}
func Repository_InitialConfig_ToProto(mapCtx *direct.MapContext, in *krm.Repository_InitialConfig) *pb.Repository_InitialConfig {
	if in == nil {
		return nil
	}
	out := &pb.Repository_InitialConfig{}
	out.DefaultBranch = direct.ValueOf(in.DefaultBranch)
	out.Gitignores = in.Gitignores
	out.License = direct.ValueOf(in.License)
	out.Readme = direct.ValueOf(in.Readme)
	return out
}
func Repository_URIsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Repository_URIs) *krm.Repository_URIsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Repository_URIsObservedState{}
	out.HTML = direct.LazyPtr(in.GetHtml())
	out.GitHTTPS = direct.LazyPtr(in.GetGitHttps())
	out.Api = direct.LazyPtr(in.GetApi())
	return out
}
func Repository_URIsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Repository_URIsObservedState) *pb.Repository_URIs {
	if in == nil {
		return nil
	}
	out := &pb.Repository_URIs{}
	out.Html = direct.ValueOf(in.HTML)
	out.GitHttps = direct.ValueOf(in.GitHTTPS)
	out.Api = direct.ValueOf(in.Api)
	return out
}
func SecureSourceManagerInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.SecureSourceManagerInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerInstanceObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.PrivateConfig = Instance_PrivateConfigObservedState_FromProto(mapCtx, in.GetPrivateConfig())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateNote = direct.Enum_FromProto(mapCtx, in.GetStateNote())
	out.HostConfig = Instance_HostConfigObservedState_FromProto(mapCtx, in.GetHostConfig())
	return out
}
func SecureSourceManagerInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecureSourceManagerInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.PrivateConfig = Instance_PrivateConfigObservedState_ToProto(mapCtx, in.PrivateConfig)
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StateNote = direct.Enum_ToProto[pb.Instance_StateNote](mapCtx, in.StateNote)
	out.HostConfig = Instance_HostConfigObservedState_ToProto(mapCtx, in.HostConfig)
	return out
}
func SecureSourceManagerInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.SecureSourceManagerInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerInstanceSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	out.PrivateConfig = Instance_PrivateConfig_FromProto(mapCtx, in.GetPrivateConfig())
	if in.GetKmsKey() != "" {
		out.KmsKeyRef = &kmsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKey()}
	}
	return out
}
func SecureSourceManagerInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecureSourceManagerInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	out.Labels = in.Labels
	out.PrivateConfig = Instance_PrivateConfig_ToProto(mapCtx, in.PrivateConfig)
	if in.KmsKeyRef != nil {
		out.KmsKey = in.KmsKeyRef.External
	}
	return out
}
func SecureSourceManagerRepositorySpec_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.SecureSourceManagerRepositorySpec {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerRepositorySpec{}
	// MISSING: Name
	// MISSING: Description
	if in.GetInstance() != "" {
		out.InstanceRef = &krm.SecureSourceManagerInstanceRef{External: in.GetInstance()}
	}
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Uris
	out.InitialConfig = Repository_InitialConfig_FromProto(mapCtx, in.GetInitialConfig())
	return out
}
func SecureSourceManagerRepositorySpec_ToProto(mapCtx *direct.MapContext, in *krm.SecureSourceManagerRepositorySpec) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	// MISSING: Name
	// MISSING: Description
	if in.InstanceRef != nil {
		out.Instance = in.InstanceRef.External
	}
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Uris
	out.InitialConfig = Repository_InitialConfig_ToProto(mapCtx, in.InitialConfig)
	return out
}
