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
