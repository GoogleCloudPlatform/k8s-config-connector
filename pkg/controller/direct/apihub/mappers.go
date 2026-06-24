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

func Documentation_FromProto(mapCtx *direct.MapContext, in *pb.Documentation) *krm.Documentation {
	if in == nil {
		return nil
	}
	out := &krm.Documentation{}
	out.ExternalURI = direct.LazyPtr(in.GetExternalUri())
	return out
}

func Documentation_ToProto(mapCtx *direct.MapContext, in *krm.Documentation) *pb.Documentation {
	if in == nil {
		return nil
	}
	out := &pb.Documentation{}
	out.ExternalUri = direct.ValueOf(in.ExternalURI)
	return out
}

func Owner_FromProto(mapCtx *direct.MapContext, in *pb.Owner) *krm.Owner {
	if in == nil {
		return nil
	}
	out := &krm.Owner{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Email = direct.LazyPtr(in.GetEmail())
	return out
}

func Owner_ToProto(mapCtx *direct.MapContext, in *krm.Owner) *pb.Owner {
	if in == nil {
		return nil
	}
	out := &pb.Owner{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Email = direct.ValueOf(in.Email)
	return out
}

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

func AttributeValues_EnumAttributeValues_FromProto(mapCtx *direct.MapContext, in *pb.AttributeValues_EnumAttributeValues) *krm.AttributeValues_EnumAttributeValues {
	if in == nil {
		return nil
	}
	out := &krm.AttributeValues_EnumAttributeValues{}
	out.Values = direct.Slice_FromProto(mapCtx, in.Values, Attribute_AllowedValue_FromProto)
	return out
}

func AttributeValues_EnumAttributeValues_ToProto(mapCtx *direct.MapContext, in *krm.AttributeValues_EnumAttributeValues) *pb.AttributeValues_EnumAttributeValues {
	if in == nil {
		return nil
	}
	out := &pb.AttributeValues_EnumAttributeValues{}
	out.Values = direct.Slice_ToProto(mapCtx, in.Values, Attribute_AllowedValue_ToProto)
	return out
}

func AttributeValues_StringAttributeValues_FromProto(mapCtx *direct.MapContext, in *pb.AttributeValues_StringAttributeValues) *krm.AttributeValues_StringAttributeValues {
	if in == nil {
		return nil
	}
	out := &krm.AttributeValues_StringAttributeValues{}
	out.Values = in.Values
	return out
}

func AttributeValues_StringAttributeValues_ToProto(mapCtx *direct.MapContext, in *krm.AttributeValues_StringAttributeValues) *pb.AttributeValues_StringAttributeValues {
	if in == nil {
		return nil
	}
	out := &pb.AttributeValues_StringAttributeValues{}
	out.Values = in.Values
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

func APIHubAPISpec_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.APIHubAPISpec {
	if in == nil {
		return nil
	}
	out := &krm.APIHubAPISpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Documentation = Documentation_FromProto(mapCtx, in.GetDocumentation())
	out.Owner = Owner_FromProto(mapCtx, in.GetOwner())

	out.TargetUserRef = APIHubAttributeValueRef_FromProto(mapCtx, in.GetTargetUser())
	out.TeamRef = APIHubAttributeValueRef_FromProto(mapCtx, in.GetTeam())
	out.BusinessUnitRef = APIHubAttributeValueRef_FromProto(mapCtx, in.GetBusinessUnit())
	out.MaturityLevelRef = APIHubAttributeValueRef_FromProto(mapCtx, in.GetMaturityLevel())
	out.APIStyleRef = APIHubAttributeValueRef_FromProto(mapCtx, in.GetApiStyle())
	if in.GetSelectedVersion() != "" {
		out.SelectedVersionRef = &krm.APIHubVersionRef{External: in.GetSelectedVersion()}
	}

	if in.GetAttributes() != nil {
		out.AttributeRefs = make([]krm.APIHubAPIAttribute, 0, len(in.GetAttributes()))
		for k, v := range in.GetAttributes() {
			attr := krm.APIHubAPIAttribute{
				AttributeRef: &krm.APIHubAttributeRef{External: k},
				Values:       AttributeValues_FromProto(mapCtx, v),
			}
			out.AttributeRefs = append(out.AttributeRefs, attr)
		}
	}
	out.APIRequirements = AttributeValues_FromProto(mapCtx, in.GetApiRequirements())
	out.APIFunctionalRequirements = AttributeValues_FromProto(mapCtx, in.GetApiFunctionalRequirements())
	out.APITechnicalRequirements = AttributeValues_FromProto(mapCtx, in.GetApiTechnicalRequirements())
	out.Fingerprint = direct.LazyPtr(in.GetFingerprint())

	return out
}

func APIHubAPISpec_ToProto(mapCtx *direct.MapContext, in *krm.APIHubAPISpec) *pb.Api {
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
	if in.SelectedVersionRef != nil {
		out.SelectedVersion = in.SelectedVersionRef.External
	}

	if in.AttributeRefs != nil {
		out.Attributes = make(map[string]*pb.AttributeValues)
		for _, attr := range in.AttributeRefs {
			if attr.AttributeRef != nil {
				out.Attributes[attr.AttributeRef.External] = AttributeValues_ToProto(mapCtx, attr.Values)
			}
		}
	}
	out.ApiRequirements = AttributeValues_ToProto(mapCtx, in.APIRequirements)
	out.ApiFunctionalRequirements = AttributeValues_ToProto(mapCtx, in.APIFunctionalRequirements)
	out.ApiTechnicalRequirements = AttributeValues_ToProto(mapCtx, in.APITechnicalRequirements)
	out.Fingerprint = direct.ValueOf(in.Fingerprint)

	return out
}

func SourceMetadataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SourceMetadata) *krm.SourceMetadataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SourceMetadataObservedState{}
	if in.SourceType != pb.SourceMetadata_SOURCE_TYPE_UNSPECIFIED {
		out.SourceType = direct.LazyPtr(in.SourceType.String())
	}
	out.OriginalResourceID = direct.LazyPtr(in.GetOriginalResourceId())
	out.OriginalResourceCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetOriginalResourceCreateTime())
	out.OriginalResourceUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetOriginalResourceUpdateTime())
	return out
}

func SourceMetadataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SourceMetadataObservedState) *pb.SourceMetadata {
	if in == nil {
		return nil
	}
	out := &pb.SourceMetadata{}
	if in.SourceType != nil {
		out.SourceType = pb.SourceMetadata_SourceType(pb.SourceMetadata_SourceType_value[direct.ValueOf(in.SourceType)])
	}
	out.OriginalResourceId = direct.ValueOf(in.OriginalResourceID)
	out.OriginalResourceCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.OriginalResourceCreateTime)
	out.OriginalResourceUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.OriginalResourceUpdateTime)
	return out
}

func APIHubAPIObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.APIHubAPIObservedState {
	if in == nil {
		return nil
	}
	out := &krm.APIHubAPIObservedState{}
	out.Versions = in.Versions
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	if len(in.GetSourceMetadata()) > 0 {
		out.SourceMetadata = make([]krm.SourceMetadataObservedState, len(in.GetSourceMetadata()))
		for i, sm := range in.GetSourceMetadata() {
			if sm == nil {
				continue
			}
			out.SourceMetadata[i] = *SourceMetadataObservedState_FromProto(mapCtx, sm)
		}
	}
	return out
}

func APIHubAPIObservedState_ToProto(mapCtx *direct.MapContext, in *krm.APIHubAPIObservedState) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	out.Versions = in.Versions
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	if len(in.SourceMetadata) > 0 {
		out.SourceMetadata = make([]*pb.SourceMetadata, len(in.SourceMetadata))
		for i, sm := range in.SourceMetadata {
			val := sm
			out.SourceMetadata[i] = SourceMetadataObservedState_ToProto(mapCtx, &val)
		}
	}
	return out
}

func Attributes_FromProto(mapCtx *direct.MapContext, in map[string]*pb.AttributeValues) []krm.APIHubExternalAPIAttribute {
	if in == nil {
		return nil
	}
	out := make([]krm.APIHubExternalAPIAttribute, 0, len(in))
	for k, v := range in {
		attr := krm.APIHubExternalAPIAttribute{
			AttributeRef: &krm.APIHubAttributeRef{External: k},
			Values:       AttributeValues_FromProto(mapCtx, v),
		}
		out = append(out, attr)
	}
	return out
}

func Attributes_ToProto(mapCtx *direct.MapContext, in []krm.APIHubExternalAPIAttribute) map[string]*pb.AttributeValues {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.AttributeValues)
	for _, attr := range in {
		if attr.AttributeRef != nil {
			out[attr.AttributeRef.External] = AttributeValues_ToProto(mapCtx, attr.Values)
		}
	}
	return out
}

func APIHubExternalAPISpec_FromProto(mapCtx *direct.MapContext, in *pb.ExternalApi) *krm.APIHubExternalAPISpec {
	if in == nil {
		return nil
	}
	out := &krm.APIHubExternalAPISpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Endpoints = in.Endpoints
	out.Paths = in.Paths
	out.Documentation = Documentation_FromProto(mapCtx, in.GetDocumentation())
	out.AttributeRefs = Attributes_FromProto(mapCtx, in.GetAttributes())
	return out
}

func APIHubExternalAPISpec_ToProto(mapCtx *direct.MapContext, in *krm.APIHubExternalAPISpec) *pb.ExternalApi {
	if in == nil {
		return nil
	}
	out := &pb.ExternalApi{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Endpoints = in.Endpoints
	out.Paths = in.Paths
	out.Documentation = Documentation_ToProto(mapCtx, in.Documentation)
	out.Attributes = Attributes_ToProto(mapCtx, in.AttributeRefs)
	return out
}

func DependencyEntityReference_FromProto(mapCtx *direct.MapContext, in *pb.DependencyEntityReference) *krm.DependencyEntityReference {
	if in == nil {
		return nil
	}
	out := &krm.DependencyEntityReference{}
	if in.GetOperationResourceName() != "" {
		out.OperationRef = &krm.APIHubOperationRef{External: in.GetOperationResourceName()}
	}
	if in.GetExternalApiResourceName() != "" {
		out.ExternalAPIRef = &krm.APIHubExternalAPIRef{External: in.GetExternalApiResourceName()}
	}
	return out
}

func DependencyEntityReference_ToProto(mapCtx *direct.MapContext, in *krm.DependencyEntityReference) *pb.DependencyEntityReference {
	if in == nil {
		return nil
	}
	out := &pb.DependencyEntityReference{}
	if in.OperationRef != nil && in.OperationRef.External != "" {
		out.Identifier = &pb.DependencyEntityReference_OperationResourceName{OperationResourceName: in.OperationRef.External}
	} else if in.ExternalAPIRef != nil && in.ExternalAPIRef.External != "" {
		out.Identifier = &pb.DependencyEntityReference_ExternalApiResourceName{ExternalApiResourceName: in.ExternalAPIRef.External}
	}
	return out
}

func APIHubDependencySpec_FromProto(mapCtx *direct.MapContext, in *pb.Dependency) *krm.APIHubDependencySpec {
	if in == nil {
		return nil
	}
	out := &krm.APIHubDependencySpec{}
	out.Consumer = DependencyEntityReference_FromProto(mapCtx, in.GetConsumer())
	out.Supplier = DependencyEntityReference_FromProto(mapCtx, in.GetSupplier())
	out.Description = direct.LazyPtr(in.GetDescription())

	if len(in.Attributes) > 0 {
		out.Attributes = make([]krm.DependencyAttribute, 0, len(in.Attributes))
		for k, v := range in.Attributes {
			attr := krm.DependencyAttribute{
				AttributeRef: krm.APIHubAttributeRef{External: k},
			}
			if v != nil {
				attrVals := AttributeValues_FromProto(mapCtx, v)
				if attrVals != nil {
					attr.Values = attrVals
				}
			}
			out.Attributes = append(out.Attributes, attr)
		}
	}
	return out
}

func APIHubDependencySpec_ToProto(mapCtx *direct.MapContext, in *krm.APIHubDependencySpec) *pb.Dependency {
	if in == nil {
		return nil
	}
	out := &pb.Dependency{}
	out.Consumer = DependencyEntityReference_ToProto(mapCtx, in.Consumer)
	out.Supplier = DependencyEntityReference_ToProto(mapCtx, in.Supplier)
	out.Description = direct.ValueOf(in.Description)

	if len(in.Attributes) > 0 {
		out.Attributes = make(map[string]*pb.AttributeValues)
		for _, attr := range in.Attributes {
			if attr.AttributeRef.External != "" {
				out.Attributes[attr.AttributeRef.External] = AttributeValues_ToProto(mapCtx, attr.Values)
			}
		}
	}
	return out
}
