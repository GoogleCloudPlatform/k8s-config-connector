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

package tpu

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/tpu/apiv2/tpupb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tpu/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func RuntimeVersion_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeVersion) *krm.RuntimeVersion {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeVersion{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func RuntimeVersion_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeVersion) *pb.RuntimeVersion {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeVersion{}
	out.Name = direct.ValueOf(in.Name)
	out.Version = direct.ValueOf(in.Version)
	return out
}
func TpuRuntimeVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeVersion) *krm.TpuRuntimeVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TpuRuntimeVersionObservedState{}
	// MISSING: Name
	// MISSING: Version
	return out
}
func TpuRuntimeVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TpuRuntimeVersionObservedState) *pb.RuntimeVersion {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeVersion{}
	// MISSING: Name
	// MISSING: Version
	return out
}
func TpuRuntimeVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeVersion) *krm.TpuRuntimeVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.TpuRuntimeVersionSpec{}
	// MISSING: Name
	// MISSING: Version
	return out
}
func TpuRuntimeVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.TpuRuntimeVersionSpec) *pb.RuntimeVersion {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeVersion{}
	// MISSING: Name
	// MISSING: Version
	return out
}
