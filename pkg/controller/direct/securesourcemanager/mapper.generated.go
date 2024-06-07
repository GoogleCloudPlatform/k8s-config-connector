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
