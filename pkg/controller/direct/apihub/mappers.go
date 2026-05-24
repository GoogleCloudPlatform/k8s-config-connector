// Copyright 2026 Google LLC
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

package apihub

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Attribute_AllowedValue_FromProto(mapCtx *direct.MapContext, in *pb.Attribute_AllowedValue) *krm.Attribute_AllowedValue {
	if in == nil {
		return nil
	}
	out := &krm.Attribute_AllowedValue{}
	out.ID = direct.LazyPtr(in.GetId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Immutable = direct.LazyPtr(in.GetImmutable())
	return out
}

func Attribute_AllowedValue_ToProto(mapCtx *direct.MapContext, in *krm.Attribute_AllowedValue) *pb.Attribute_AllowedValue {
	if in == nil {
		return nil
	}
	out := &pb.Attribute_AllowedValue{}
	out.Id = direct.ValueOf(in.ID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Immutable = direct.ValueOf(in.Immutable)
	return out
}

func APIHubAttributeValueRef_FromProto(mapCtx *direct.MapContext, in *pb.AttributeValues) *krm.APIHubAttributeValueRef {
	if in == nil {
		return nil
	}
	out := &krm.APIHubAttributeValueRef{}
	if in.GetEnumValues() != nil && len(in.GetEnumValues().Values) > 0 {
		out.External = in.GetEnumValues().Values[0].Id
	}
	// Note: We only support enum values for now as per the requirements for deploymentType, slo, and environment
	return out
}

func APIHubAttributeValueRef_ToProto(mapCtx *direct.MapContext, in *krm.APIHubAttributeValueRef) *pb.AttributeValues {
	if in == nil {
		return nil
	}
	if in.External == "" {
		return nil
	}
	out := &pb.AttributeValues{}
	out.Value = &pb.AttributeValues_EnumValues{
		EnumValues: &pb.AttributeValues_EnumAttributeValues{
			Values: []*pb.Attribute_AllowedValue{
				{
					Id: in.External,
				},
			},
		},
	}
	return out
}

func APIHubDeploymentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.APIHubDeploymentSpec {
	if in == nil {
		return nil
	}
	out := &krm.APIHubDeploymentSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Documentation = Documentation_FromProto(mapCtx, in.GetDocumentation())
	out.DeploymentTypeRef = APIHubAttributeValueRef_FromProto(mapCtx, in.GetDeploymentType())
	out.ResourceURI = direct.LazyPtr(in.GetResourceUri())
	out.Endpoints = in.Endpoints
	out.SloRef = APIHubAttributeValueRef_FromProto(mapCtx, in.GetSlo())
	out.EnvironmentRef = APIHubAttributeValueRef_FromProto(mapCtx, in.GetEnvironment())
	return out
}

func APIHubDeploymentSpec_ToProto(mapCtx *direct.MapContext, in *krm.APIHubDeploymentSpec) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Documentation = Documentation_ToProto(mapCtx, in.Documentation)
	out.DeploymentType = APIHubAttributeValueRef_ToProto(mapCtx, in.DeploymentTypeRef)
	out.ResourceUri = direct.ValueOf(in.ResourceURI)
	out.Endpoints = in.Endpoints
	out.Slo = APIHubAttributeValueRef_ToProto(mapCtx, in.SloRef)
	out.Environment = APIHubAttributeValueRef_ToProto(mapCtx, in.EnvironmentRef)
	return out
}

func ApiHubApiSpec_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.ApiHubApiSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApiHubApiSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Documentation = Documentation_FromProto(mapCtx, in.GetDocumentation())
	out.Owner = Owner_FromProto(mapCtx, in.GetOwner())

	out.TargetUserRef = APIHubAttributeValueRef_FromProto(mapCtx, in.GetTargetUser())
	out.TeamRef = APIHubAttributeValueRef_FromProto(mapCtx, in.GetTeam())
	out.BusinessUnitRef = APIHubAttributeValueRef_FromProto(mapCtx, in.GetBusinessUnit())
	out.MaturityLevelRef = APIHubAttributeValueRef_FromProto(mapCtx, in.GetMaturityLevel())
	out.APIStyleRef = APIHubAttributeValueRef_FromProto(mapCtx, in.GetApiStyle())
	out.SelectedVersion = direct.LazyPtr(in.GetSelectedVersion())
	return out
}

func ApiHubApiSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApiHubApiSpec) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Documentation = Documentation_ToProto(mapCtx, in.Documentation)
	out.Owner = Owner_ToProto(mapCtx, in.Owner)

	out.TargetUser = APIHubAttributeValueRef_ToProto(mapCtx, in.TargetUserRef)
	out.Team = APIHubAttributeValueRef_ToProto(mapCtx, in.TeamRef)
	out.BusinessUnit = APIHubAttributeValueRef_ToProto(mapCtx, in.BusinessUnitRef)
	out.MaturityLevel = APIHubAttributeValueRef_ToProto(mapCtx, in.MaturityLevelRef)
	out.ApiStyle = APIHubAttributeValueRef_ToProto(mapCtx, in.APIStyleRef)
	out.SelectedVersion = direct.ValueOf(in.SelectedVersion)
	return out
}
