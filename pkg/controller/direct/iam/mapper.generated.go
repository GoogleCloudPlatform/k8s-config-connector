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

package iam

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/iam/apiv1beta/iampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func IamWorkloadIdentityPoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadIdentityPool) *krm.IamWorkloadIdentityPoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IamWorkloadIdentityPoolObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: State
	// MISSING: Disabled
	return out
}
func IamWorkloadIdentityPoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IamWorkloadIdentityPoolObservedState) *pb.WorkloadIdentityPool {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadIdentityPool{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: State
	// MISSING: Disabled
	return out
}
func IamWorkloadIdentityPoolSpec_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadIdentityPool) *krm.IamWorkloadIdentityPoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.IamWorkloadIdentityPoolSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: State
	// MISSING: Disabled
	return out
}
func IamWorkloadIdentityPoolSpec_ToProto(mapCtx *direct.MapContext, in *krm.IamWorkloadIdentityPoolSpec) *pb.WorkloadIdentityPool {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadIdentityPool{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: State
	// MISSING: Disabled
	return out
}
func WorkloadIdentityPool_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadIdentityPool) *krm.WorkloadIdentityPool {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadIdentityPool{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}
func WorkloadIdentityPool_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadIdentityPool) *pb.WorkloadIdentityPool {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadIdentityPool{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	out.Disabled = direct.ValueOf(in.Disabled)
	return out
}
func WorkloadIdentityPoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadIdentityPool) *krm.WorkloadIdentityPoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadIdentityPoolObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Disabled
	return out
}
func WorkloadIdentityPoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadIdentityPoolObservedState) *pb.WorkloadIdentityPool {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadIdentityPool{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.WorkloadIdentityPool_State](mapCtx, in.State)
	// MISSING: Disabled
	return out
}
