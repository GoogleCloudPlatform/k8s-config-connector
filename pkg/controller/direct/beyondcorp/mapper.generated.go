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

package beyondcorp

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/beyondcorp/appconnectors/apiv1/appconnectorspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/beyondcorp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AppConnector_FromProto(mapCtx *direct.MapContext, in *pb.AppConnector) *krm.AppConnector {
	if in == nil {
		return nil
	}
	out := &krm.AppConnector{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: State
	out.PrincipalInfo = AppConnector_PrincipalInfo_FromProto(mapCtx, in.GetPrincipalInfo())
	out.ResourceInfo = ResourceInfo_FromProto(mapCtx, in.GetResourceInfo())
	return out
}
func AppConnector_ToProto(mapCtx *direct.MapContext, in *krm.AppConnector) *pb.AppConnector {
	if in == nil {
		return nil
	}
	out := &pb.AppConnector{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: State
	out.PrincipalInfo = AppConnector_PrincipalInfo_ToProto(mapCtx, in.PrincipalInfo)
	out.ResourceInfo = ResourceInfo_ToProto(mapCtx, in.ResourceInfo)
	return out
}
func AppConnectorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AppConnector) *krm.AppConnectorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AppConnectorObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: PrincipalInfo
	// MISSING: ResourceInfo
	return out
}
func AppConnectorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AppConnectorObservedState) *pb.AppConnector {
	if in == nil {
		return nil
	}
	out := &pb.AppConnector{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	out.State = direct.Enum_ToProto[pb.AppConnector_State](mapCtx, in.State)
	// MISSING: PrincipalInfo
	// MISSING: ResourceInfo
	return out
}
func AppConnector_PrincipalInfo_FromProto(mapCtx *direct.MapContext, in *pb.AppConnector_PrincipalInfo) *krm.AppConnector_PrincipalInfo {
	if in == nil {
		return nil
	}
	out := &krm.AppConnector_PrincipalInfo{}
	out.ServiceAccount = AppConnector_PrincipalInfo_ServiceAccount_FromProto(mapCtx, in.GetServiceAccount())
	return out
}
func AppConnector_PrincipalInfo_ToProto(mapCtx *direct.MapContext, in *krm.AppConnector_PrincipalInfo) *pb.AppConnector_PrincipalInfo {
	if in == nil {
		return nil
	}
	out := &pb.AppConnector_PrincipalInfo{}
	if oneof := AppConnector_PrincipalInfo_ServiceAccount_ToProto(mapCtx, in.ServiceAccount); oneof != nil {
		out.Type = &pb.AppConnector_PrincipalInfo_ServiceAccount_{ServiceAccount: oneof}
	}
	return out
}
func AppConnector_PrincipalInfo_ServiceAccount_FromProto(mapCtx *direct.MapContext, in *pb.AppConnector_PrincipalInfo_ServiceAccount) *krm.AppConnector_PrincipalInfo_ServiceAccount {
	if in == nil {
		return nil
	}
	out := &krm.AppConnector_PrincipalInfo_ServiceAccount{}
	out.Email = direct.LazyPtr(in.GetEmail())
	return out
}
func AppConnector_PrincipalInfo_ServiceAccount_ToProto(mapCtx *direct.MapContext, in *krm.AppConnector_PrincipalInfo_ServiceAccount) *pb.AppConnector_PrincipalInfo_ServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.AppConnector_PrincipalInfo_ServiceAccount{}
	out.Email = direct.ValueOf(in.Email)
	return out
}
func BeyondcorpAppConnectorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AppConnector) *krm.BeyondcorpAppConnectorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BeyondcorpAppConnectorObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: State
	// MISSING: PrincipalInfo
	// MISSING: ResourceInfo
	return out
}
func BeyondcorpAppConnectorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BeyondcorpAppConnectorObservedState) *pb.AppConnector {
	if in == nil {
		return nil
	}
	out := &pb.AppConnector{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: State
	// MISSING: PrincipalInfo
	// MISSING: ResourceInfo
	return out
}
func BeyondcorpAppConnectorSpec_FromProto(mapCtx *direct.MapContext, in *pb.AppConnector) *krm.BeyondcorpAppConnectorSpec {
	if in == nil {
		return nil
	}
	out := &krm.BeyondcorpAppConnectorSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: State
	// MISSING: PrincipalInfo
	// MISSING: ResourceInfo
	return out
}
func BeyondcorpAppConnectorSpec_ToProto(mapCtx *direct.MapContext, in *krm.BeyondcorpAppConnectorSpec) *pb.AppConnector {
	if in == nil {
		return nil
	}
	out := &pb.AppConnector{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: State
	// MISSING: PrincipalInfo
	// MISSING: ResourceInfo
	return out
}
func ResourceInfo_FromProto(mapCtx *direct.MapContext, in *pb.ResourceInfo) *krm.ResourceInfo {
	if in == nil {
		return nil
	}
	out := &krm.ResourceInfo{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Status = direct.Enum_FromProto(mapCtx, in.GetStatus())
	out.Resource = Any_FromProto(mapCtx, in.GetResource())
	out.Time = direct.StringTimestamp_FromProto(mapCtx, in.GetTime())
	out.Sub = direct.Slice_FromProto(mapCtx, in.Sub, ResourceInfo_FromProto)
	return out
}
func ResourceInfo_ToProto(mapCtx *direct.MapContext, in *krm.ResourceInfo) *pb.ResourceInfo {
	if in == nil {
		return nil
	}
	out := &pb.ResourceInfo{}
	out.Id = direct.ValueOf(in.ID)
	out.Status = direct.Enum_ToProto[pb.HealthStatus](mapCtx, in.Status)
	out.Resource = Any_ToProto(mapCtx, in.Resource)
	out.Time = direct.StringTimestamp_ToProto(mapCtx, in.Time)
	out.Sub = direct.Slice_ToProto(mapCtx, in.Sub, ResourceInfo_ToProto)
	return out
}
