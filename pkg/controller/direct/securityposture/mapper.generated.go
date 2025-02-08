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

package securityposture

import (
	pb "cloud.google.com/go/securityposture/apiv1/securityposturepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securityposture/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func PostureDeployment_FromProto(mapCtx *direct.MapContext, in *pb.PostureDeployment) *krm.PostureDeployment {
	if in == nil {
		return nil
	}
	out := &krm.PostureDeployment{}
	out.Name = direct.LazyPtr(in.GetName())
	out.TargetResource = direct.LazyPtr(in.GetTargetResource())
	// MISSING: State
	out.PostureID = direct.LazyPtr(in.GetPostureId())
	out.PostureRevisionID = direct.LazyPtr(in.GetPostureRevisionId())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Annotations = in.Annotations
	// MISSING: Reconciling
	// MISSING: DesiredPostureID
	// MISSING: DesiredPostureRevisionID
	// MISSING: FailureMessage
	return out
}
func PostureDeployment_ToProto(mapCtx *direct.MapContext, in *krm.PostureDeployment) *pb.PostureDeployment {
	if in == nil {
		return nil
	}
	out := &pb.PostureDeployment{}
	out.Name = direct.ValueOf(in.Name)
	out.TargetResource = direct.ValueOf(in.TargetResource)
	// MISSING: State
	out.PostureId = direct.ValueOf(in.PostureID)
	out.PostureRevisionId = direct.ValueOf(in.PostureRevisionID)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.Etag = direct.ValueOf(in.Etag)
	out.Annotations = in.Annotations
	// MISSING: Reconciling
	// MISSING: DesiredPostureID
	// MISSING: DesiredPostureRevisionID
	// MISSING: FailureMessage
	return out
}
func PostureDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PostureDeployment) *krm.PostureDeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PostureDeploymentObservedState{}
	// MISSING: Name
	// MISSING: TargetResource
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: PostureID
	// MISSING: PostureRevisionID
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	// MISSING: Etag
	// MISSING: Annotations
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.DesiredPostureID = direct.LazyPtr(in.GetDesiredPostureId())
	out.DesiredPostureRevisionID = direct.LazyPtr(in.GetDesiredPostureRevisionId())
	out.FailureMessage = direct.LazyPtr(in.GetFailureMessage())
	return out
}
func PostureDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PostureDeploymentObservedState) *pb.PostureDeployment {
	if in == nil {
		return nil
	}
	out := &pb.PostureDeployment{}
	// MISSING: Name
	// MISSING: TargetResource
	out.State = direct.Enum_ToProto[pb.PostureDeployment_State](mapCtx, in.State)
	// MISSING: PostureID
	// MISSING: PostureRevisionID
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	// MISSING: Etag
	// MISSING: Annotations
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.DesiredPostureId = direct.ValueOf(in.DesiredPostureID)
	out.DesiredPostureRevisionId = direct.ValueOf(in.DesiredPostureRevisionID)
	out.FailureMessage = direct.ValueOf(in.FailureMessage)
	return out
}
func SecurityposturePostureDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PostureDeployment) *krm.SecurityposturePostureDeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecurityposturePostureDeploymentObservedState{}
	// MISSING: Name
	// MISSING: TargetResource
	// MISSING: State
	// MISSING: PostureID
	// MISSING: PostureRevisionID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: DesiredPostureID
	// MISSING: DesiredPostureRevisionID
	// MISSING: FailureMessage
	return out
}
func SecurityposturePostureDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecurityposturePostureDeploymentObservedState) *pb.PostureDeployment {
	if in == nil {
		return nil
	}
	out := &pb.PostureDeployment{}
	// MISSING: Name
	// MISSING: TargetResource
	// MISSING: State
	// MISSING: PostureID
	// MISSING: PostureRevisionID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: DesiredPostureID
	// MISSING: DesiredPostureRevisionID
	// MISSING: FailureMessage
	return out
}
func SecurityposturePostureDeploymentSpec_FromProto(mapCtx *direct.MapContext, in *pb.PostureDeployment) *krm.SecurityposturePostureDeploymentSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecurityposturePostureDeploymentSpec{}
	// MISSING: Name
	// MISSING: TargetResource
	// MISSING: State
	// MISSING: PostureID
	// MISSING: PostureRevisionID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: DesiredPostureID
	// MISSING: DesiredPostureRevisionID
	// MISSING: FailureMessage
	return out
}
func SecurityposturePostureDeploymentSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecurityposturePostureDeploymentSpec) *pb.PostureDeployment {
	if in == nil {
		return nil
	}
	out := &pb.PostureDeployment{}
	// MISSING: Name
	// MISSING: TargetResource
	// MISSING: State
	// MISSING: PostureID
	// MISSING: PostureRevisionID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: DesiredPostureID
	// MISSING: DesiredPostureRevisionID
	// MISSING: FailureMessage
	return out
}
