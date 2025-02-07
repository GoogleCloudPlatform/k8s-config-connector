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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func ApphubApplicationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Application) *krm.ApphubApplicationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApphubApplicationObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Scope
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApphubApplicationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApphubApplicationObservedState) *pb.Application {
	if in == nil {
		return nil
	}
	out := &pb.Application{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Scope
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApphubApplicationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Application) *krm.ApphubApplicationSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApphubApplicationSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Scope
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApphubApplicationSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApphubApplicationSpec) *pb.Application {
	if in == nil {
		return nil
	}
	out := &pb.Application{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Scope
	// MISSING: Uid
	// MISSING: State
	return out
}
func Application_FromProto(mapCtx *direct.MapContext, in *pb.Application) *krm.Application {
	if in == nil {
		return nil
	}
	out := &krm.Application{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Attributes = Attributes_FromProto(mapCtx, in.GetAttributes())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Scope = Scope_FromProto(mapCtx, in.GetScope())
	// MISSING: Uid
	// MISSING: State
	return out
}
func Application_ToProto(mapCtx *direct.MapContext, in *krm.Application) *pb.Application {
	if in == nil {
		return nil
	}
	out := &pb.Application{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Attributes = Attributes_ToProto(mapCtx, in.Attributes)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Scope = Scope_ToProto(mapCtx, in.Scope)
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApplicationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Application) *krm.ApplicationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApplicationObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Attributes
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Scope
	out.Uid = direct.LazyPtr(in.GetUid())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func ApplicationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApplicationObservedState) *pb.Application {
	if in == nil {
		return nil
	}
	out := &pb.Application{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Attributes
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Scope
	out.Uid = direct.ValueOf(in.Uid)
	out.State = direct.Enum_ToProto[pb.Application_State](mapCtx, in.State)
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
