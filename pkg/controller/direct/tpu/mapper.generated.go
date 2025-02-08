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
	pb "cloud.google.com/go/tpu/apiv1/tpupb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tpu/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AcceleratorType_FromProto(mapCtx *direct.MapContext, in *pb.AcceleratorType) *krm.AcceleratorType {
	if in == nil {
		return nil
	}
	out := &krm.AcceleratorType{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.LazyPtr(in.GetType())
	return out
}
func AcceleratorType_ToProto(mapCtx *direct.MapContext, in *krm.AcceleratorType) *pb.AcceleratorType {
	if in == nil {
		return nil
	}
	out := &pb.AcceleratorType{}
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.ValueOf(in.Type)
	return out
}
func TpuAcceleratorTypeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AcceleratorType) *krm.TpuAcceleratorTypeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TpuAcceleratorTypeObservedState{}
	// MISSING: Name
	// MISSING: Type
	return out
}
func TpuAcceleratorTypeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TpuAcceleratorTypeObservedState) *pb.AcceleratorType {
	if in == nil {
		return nil
	}
	out := &pb.AcceleratorType{}
	// MISSING: Name
	// MISSING: Type
	return out
}
func TpuAcceleratorTypeSpec_FromProto(mapCtx *direct.MapContext, in *pb.AcceleratorType) *krm.TpuAcceleratorTypeSpec {
	if in == nil {
		return nil
	}
	out := &krm.TpuAcceleratorTypeSpec{}
	// MISSING: Name
	// MISSING: Type
	return out
}
func TpuAcceleratorTypeSpec_ToProto(mapCtx *direct.MapContext, in *krm.TpuAcceleratorTypeSpec) *pb.AcceleratorType {
	if in == nil {
		return nil
	}
	out := &pb.AcceleratorType{}
	// MISSING: Name
	// MISSING: Type
	return out
}
