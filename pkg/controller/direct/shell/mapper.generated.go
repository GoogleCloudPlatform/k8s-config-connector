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

package shell

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/shell/apiv1/shellpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/shell/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Environment_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.Environment {
	if in == nil {
		return nil
	}
	out := &krm.Environment{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ID
	out.DockerImage = direct.LazyPtr(in.GetDockerImage())
	// MISSING: State
	// MISSING: WebHost
	// MISSING: SSHUsername
	// MISSING: SSHHost
	// MISSING: SSHPort
	// MISSING: PublicKeys
	return out
}
func Environment_ToProto(mapCtx *direct.MapContext, in *krm.Environment) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ID
	out.DockerImage = direct.ValueOf(in.DockerImage)
	// MISSING: State
	// MISSING: WebHost
	// MISSING: SSHUsername
	// MISSING: SSHHost
	// MISSING: SSHPort
	// MISSING: PublicKeys
	return out
}
func EnvironmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.EnvironmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EnvironmentObservedState{}
	// MISSING: Name
	out.ID = direct.LazyPtr(in.GetId())
	// MISSING: DockerImage
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.WebHost = direct.LazyPtr(in.GetWebHost())
	out.SSHUsername = direct.LazyPtr(in.GetSshUsername())
	out.SSHHost = direct.LazyPtr(in.GetSshHost())
	out.SSHPort = direct.LazyPtr(in.GetSshPort())
	out.PublicKeys = in.PublicKeys
	return out
}
func EnvironmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EnvironmentObservedState) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	// MISSING: Name
	out.Id = direct.ValueOf(in.ID)
	// MISSING: DockerImage
	out.State = direct.Enum_ToProto[pb.Environment_State](mapCtx, in.State)
	out.WebHost = direct.ValueOf(in.WebHost)
	out.SshUsername = direct.ValueOf(in.SSHUsername)
	out.SshHost = direct.ValueOf(in.SSHHost)
	out.SshPort = direct.ValueOf(in.SSHPort)
	out.PublicKeys = in.PublicKeys
	return out
}
func ShellEnvironmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.ShellEnvironmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ShellEnvironmentObservedState{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: DockerImage
	// MISSING: State
	// MISSING: WebHost
	// MISSING: SSHUsername
	// MISSING: SSHHost
	// MISSING: SSHPort
	// MISSING: PublicKeys
	return out
}
func ShellEnvironmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ShellEnvironmentObservedState) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: DockerImage
	// MISSING: State
	// MISSING: WebHost
	// MISSING: SSHUsername
	// MISSING: SSHHost
	// MISSING: SSHPort
	// MISSING: PublicKeys
	return out
}
func ShellEnvironmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.ShellEnvironmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ShellEnvironmentSpec{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: DockerImage
	// MISSING: State
	// MISSING: WebHost
	// MISSING: SSHUsername
	// MISSING: SSHHost
	// MISSING: SSHPort
	// MISSING: PublicKeys
	return out
}
func ShellEnvironmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.ShellEnvironmentSpec) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: DockerImage
	// MISSING: State
	// MISSING: WebHost
	// MISSING: SSHUsername
	// MISSING: SSHHost
	// MISSING: SSHPort
	// MISSING: PublicKeys
	return out
}
