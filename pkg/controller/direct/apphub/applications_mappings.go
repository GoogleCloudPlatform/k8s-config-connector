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

package apphub

import (
	pb "cloud.google.com/go/apphub/apiv1/apphubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apphub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AppHubApplicationStatus_FromProto(mapCtx *direct.MapContext, in *pb.Application) *krm.AppHubApplicationStatus {
	if in == nil {
		return nil
	}
	out := &krm.AppHubApplicationStatus{}
	out.ObservedState = AppHubApplicationObservedState_FromProto(mapCtx, in)
	return out
}

func AppHubApplicationStatus_ToProto(mapCtx *direct.MapContext, in *krm.AppHubApplicationStatus) *pb.Application {
	if in == nil {
		return nil
	}
	out := &pb.Application{}
	out = AppHubApplicationObservedState_ToProto(mapCtx, in.ObservedState)
	return out
}

func AppHubApplicationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Application) *krm.AppHubApplicationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AppHubApplicationObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func AppHubApplicationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AppHubApplicationObservedState) *pb.Application {
	if in == nil {
		return nil
	}
	out := &pb.Application{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Uid = direct.ValueOf(in.Uid)
	out.State = direct.Enum_ToProto[pb.Application_State](mapCtx, in.State)
	return out
}
func AppHubApplicationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Application) *krm.AppHubApplicationSpec {
	if in == nil {
		return nil
	}
	out := &krm.AppHubApplicationSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Attributes = Attributes_FromProto(mapCtx, in.GetAttributes())
	out.Scope = Scope_FromProto(mapCtx, in.GetScope())
	return out
}
func AppHubApplicationSpec_ToProto(mapCtx *direct.MapContext, in *krm.AppHubApplicationSpec) *pb.Application {
	if in == nil {
		return nil
	}
	out := &pb.Application{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Attributes = Attributes_ToProto(mapCtx, in.Attributes)
	out.Scope = Scope_ToProto(mapCtx, in.Scope)
	return out
}
func Attributes_FromProto(mapCtx *direct.MapContext, in *pb.Attributes) *krm.Attributes {
	if in == nil {
		return nil
	}
	out := &krm.Attributes{}
	out.Criticality = Criticality_FromProto(mapCtx, in.GetCriticality())
	out.Environment = Environment_FromProto(mapCtx, in.GetEnvironment())
	out.DeveloperOwners = direct.Slice_FromProto(mapCtx, in.DeveloperOwners, ContactInfo_FromProto)
	out.OperatorOwners = direct.Slice_FromProto(mapCtx, in.OperatorOwners, ContactInfo_FromProto)
	out.BusinessOwners = direct.Slice_FromProto(mapCtx, in.BusinessOwners, ContactInfo_FromProto)
	return out
}
func Attributes_ToProto(mapCtx *direct.MapContext, in *krm.Attributes) *pb.Attributes {
	if in == nil {
		return nil
	}
	out := &pb.Attributes{}
	out.Criticality = Criticality_ToProto(mapCtx, in.Criticality)
	out.Environment = Environment_ToProto(mapCtx, in.Environment)
	out.DeveloperOwners = direct.Slice_ToProto(mapCtx, in.DeveloperOwners, ContactInfo_ToProto)
	out.OperatorOwners = direct.Slice_ToProto(mapCtx, in.OperatorOwners, ContactInfo_ToProto)
	out.BusinessOwners = direct.Slice_ToProto(mapCtx, in.BusinessOwners, ContactInfo_ToProto)
	return out
}
func ContactInfo_FromProto(mapCtx *direct.MapContext, in *pb.ContactInfo) *krm.ContactInfo {
	if in == nil {
		return nil
	}
	out := &krm.ContactInfo{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Email = direct.LazyPtr(in.GetEmail())
	return out
}
func ContactInfo_ToProto(mapCtx *direct.MapContext, in *krm.ContactInfo) *pb.ContactInfo {
	if in == nil {
		return nil
	}
	out := &pb.ContactInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Email = direct.ValueOf(in.Email)
	return out
}
func Criticality_FromProto(mapCtx *direct.MapContext, in *pb.Criticality) *krm.Criticality {
	if in == nil {
		return nil
	}
	out := &krm.Criticality{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Criticality_ToProto(mapCtx *direct.MapContext, in *krm.Criticality) *pb.Criticality {
	if in == nil {
		return nil
	}
	out := &pb.Criticality{}
	out.Type = direct.Enum_ToProto[pb.Criticality_Type](mapCtx, in.Type)
	return out
}
func Environment_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.Environment {
	if in == nil {
		return nil
	}
	out := &krm.Environment{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Environment_ToProto(mapCtx *direct.MapContext, in *krm.Environment) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	out.Type = direct.Enum_ToProto[pb.Environment_Type](mapCtx, in.Type)
	return out
}
func Scope_FromProto(mapCtx *direct.MapContext, in *pb.Scope) *krm.Scope {
	if in == nil {
		return nil
	}
	out := &krm.Scope{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Scope_ToProto(mapCtx *direct.MapContext, in *krm.Scope) *pb.Scope {
	if in == nil {
		return nil
	}
	out := &pb.Scope{}
	out.Type = direct.Enum_ToProto[pb.Scope_Type](mapCtx, in.Type)
	return out
}
