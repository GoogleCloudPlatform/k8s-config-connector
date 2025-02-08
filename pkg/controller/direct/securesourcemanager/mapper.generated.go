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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securesourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Instance_HostConfig_FromProto(mapCtx *direct.MapContext, in *pb.Instance_HostConfig) *krm.Instance_HostConfig {
	if in == nil {
		return nil
	}
	out := &krm.Instance_HostConfig{}
	// MISSING: HTML
	// MISSING: Api
	// MISSING: GitHTTP
	// MISSING: GitSSH
	return out
}
func Instance_HostConfig_ToProto(mapCtx *direct.MapContext, in *krm.Instance_HostConfig) *pb.Instance_HostConfig {
	if in == nil {
		return nil
	}
	out := &pb.Instance_HostConfig{}
	// MISSING: HTML
	// MISSING: Api
	// MISSING: GitHTTP
	// MISSING: GitSSH
	return out
}
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
func SecureSourceManagerInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.SecureSourceManagerInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerInstanceObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateNote = direct.Enum_FromProto(mapCtx, in.GetStateNote())
	// MISSING: KMSKey
	out.HostConfig = Instance_HostConfig_FromProto(mapCtx, in.GetHostConfig())
	return out
}
func SecureSourceManagerInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecureSourceManagerInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StateNote = direct.Enum_ToProto[pb.Instance_StateNote](mapCtx, in.StateNote)
	// MISSING: KMSKey
	out.HostConfig = Instance_HostConfig_ToProto(mapCtx, in.HostConfig)
	return out
}
func SecureSourceManagerInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.SecureSourceManagerInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerInstanceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.PrivateConfig = Instance_PrivateConfig_FromProto(mapCtx, in.GetPrivateConfig())
	// MISSING: KMSKey
	return out
}
func SecureSourceManagerInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecureSourceManagerInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.PrivateConfig = Instance_PrivateConfig_ToProto(mapCtx, in.PrivateConfig)
	// MISSING: KMSKey
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
		out.InstanceRef = &refs.*SecureSourceManagerInstanceRef{External: in.GetInstance()}
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
