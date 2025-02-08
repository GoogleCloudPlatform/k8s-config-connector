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

package parametermanager

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/parametermanager/apiv1/parametermanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/parametermanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ParameterVersion_FromProto(mapCtx *direct.MapContext, in *pb.ParameterVersion) *krm.ParameterVersion {
	if in == nil {
		return nil
	}
	out := &krm.ParameterVersion{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.Payload = ParameterVersionPayload_FromProto(mapCtx, in.GetPayload())
	return out
}
func ParameterVersion_ToProto(mapCtx *direct.MapContext, in *krm.ParameterVersion) *pb.ParameterVersion {
	if in == nil {
		return nil
	}
	out := &pb.ParameterVersion{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Disabled = direct.ValueOf(in.Disabled)
	out.Payload = ParameterVersionPayload_ToProto(mapCtx, in.Payload)
	return out
}
func ParameterVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ParameterVersion) *krm.ParameterVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ParameterVersionObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Disabled
	// MISSING: Payload
	return out
}
func ParameterVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ParameterVersionObservedState) *pb.ParameterVersion {
	if in == nil {
		return nil
	}
	out := &pb.ParameterVersion{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Disabled
	// MISSING: Payload
	return out
}
func ParameterVersionPayload_FromProto(mapCtx *direct.MapContext, in *pb.ParameterVersionPayload) *krm.ParameterVersionPayload {
	if in == nil {
		return nil
	}
	out := &krm.ParameterVersionPayload{}
	out.Data = in.GetData()
	return out
}
func ParameterVersionPayload_ToProto(mapCtx *direct.MapContext, in *krm.ParameterVersionPayload) *pb.ParameterVersionPayload {
	if in == nil {
		return nil
	}
	out := &pb.ParameterVersionPayload{}
	out.Data = in.Data
	return out
}
func ParametermanagerParameterVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ParameterVersion) *krm.ParametermanagerParameterVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ParametermanagerParameterVersionObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Disabled
	// MISSING: Payload
	return out
}
func ParametermanagerParameterVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ParametermanagerParameterVersionObservedState) *pb.ParameterVersion {
	if in == nil {
		return nil
	}
	out := &pb.ParameterVersion{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Disabled
	// MISSING: Payload
	return out
}
func ParametermanagerParameterVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.ParameterVersion) *krm.ParametermanagerParameterVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.ParametermanagerParameterVersionSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Disabled
	// MISSING: Payload
	return out
}
func ParametermanagerParameterVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.ParametermanagerParameterVersionSpec) *pb.ParameterVersion {
	if in == nil {
		return nil
	}
	out := &pb.ParameterVersion{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Disabled
	// MISSING: Payload
	return out
}
