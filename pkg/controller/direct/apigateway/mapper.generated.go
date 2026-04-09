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

// +generated:mapper
// krm.group: apigateway.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.apigateway.v1

package apigateway

import (
	pb "cloud.google.com/go/apigateway/apiv1/apigatewaypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigateway/v1alpha1"
	krmapigatewayv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigateway/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func APIConfig_File_FromProto(mapCtx *direct.MapContext, in *pb.ApiConfig_File) *krm.APIConfig_File {
	if in == nil {
		return nil
	}
	out := &krm.APIConfig_File{}
	out.Path = direct.LazyPtr(in.GetPath())
	out.Contents = []krm.byte{direct.LazyPtr(in.GetContents())}
	return out
}
func APIConfig_File_ToProto(mapCtx *direct.MapContext, in *krm.APIConfig_File) *pb.ApiConfig_File {
	if in == nil {
		return nil
	}
	out := &pb.ApiConfig_File{}
	out.Path = direct.ValueOf(in.Path)
	if len(in.Contents) > 0 && in.Contents[0] != nil {
		out.Contents = direct.ValueOf(in.Contents[0])
	}
	return out
}
func APIConfig_GrpcServiceDefinition_FromProto(mapCtx *direct.MapContext, in *pb.ApiConfig_GrpcServiceDefinition) *krm.APIConfig_GrpcServiceDefinition {
	if in == nil {
		return nil
	}
	out := &krm.APIConfig_GrpcServiceDefinition{}
	out.FileDescriptorSet = APIConfig_File_FromProto(mapCtx, in.GetFileDescriptorSet())
	out.Source = direct.Slice_FromProto(mapCtx, in.Source, APIConfig_File_FromProto)
	return out
}
func APIConfig_GrpcServiceDefinition_ToProto(mapCtx *direct.MapContext, in *krm.APIConfig_GrpcServiceDefinition) *pb.ApiConfig_GrpcServiceDefinition {
	if in == nil {
		return nil
	}
	out := &pb.ApiConfig_GrpcServiceDefinition{}
	out.FileDescriptorSet = APIConfig_File_ToProto(mapCtx, in.FileDescriptorSet)
	out.Source = direct.Slice_ToProto(mapCtx, in.Source, APIConfig_File_ToProto)
	return out
}
func APIConfig_OpenAPIDocument_FromProto(mapCtx *direct.MapContext, in *pb.ApiConfig_OpenApiDocument) *krm.APIConfig_OpenAPIDocument {
	if in == nil {
		return nil
	}
	out := &krm.APIConfig_OpenAPIDocument{}
	out.Document = APIConfig_File_FromProto(mapCtx, in.GetDocument())
	return out
}
func APIConfig_OpenAPIDocument_ToProto(mapCtx *direct.MapContext, in *krm.APIConfig_OpenAPIDocument) *pb.ApiConfig_OpenApiDocument {
	if in == nil {
		return nil
	}
	out := &pb.ApiConfig_OpenApiDocument{}
	out.Document = APIConfig_File_ToProto(mapCtx, in.Document)
	return out
}
func APIGatewayAPIConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ApiConfig) *krm.APIGatewayAPIConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.APIGatewayAPIConfigObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: GatewayServiceAccount
	// MISSING: ServiceConfigID
	// MISSING: State
	// MISSING: OpenapiDocuments
	// MISSING: GrpcServices
	// MISSING: ManagedServiceConfigs
	return out
}
func APIGatewayAPIConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.APIGatewayAPIConfigObservedState) *pb.ApiConfig {
	if in == nil {
		return nil
	}
	out := &pb.ApiConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: GatewayServiceAccount
	// MISSING: ServiceConfigID
	// MISSING: State
	// MISSING: OpenapiDocuments
	// MISSING: GrpcServices
	// MISSING: ManagedServiceConfigs
	return out
}
func APIGatewayAPIConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.ApiConfig) *krm.APIGatewayAPIConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.APIGatewayAPIConfigSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: GatewayServiceAccount
	// MISSING: ServiceConfigID
	// MISSING: State
	// MISSING: OpenapiDocuments
	// MISSING: GrpcServices
	// MISSING: ManagedServiceConfigs
	return out
}
func APIGatewayAPIConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.APIGatewayAPIConfigSpec) *pb.ApiConfig {
	if in == nil {
		return nil
	}
	out := &pb.ApiConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: GatewayServiceAccount
	// MISSING: ServiceConfigID
	// MISSING: State
	// MISSING: OpenapiDocuments
	// MISSING: GrpcServices
	// MISSING: ManagedServiceConfigs
	return out
}
func APIGatewayAPIObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.APIGatewayAPIObservedState {
	if in == nil {
		return nil
	}
	out := &krm.APIGatewayAPIObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func APIGatewayAPIObservedState_ToProto(mapCtx *direct.MapContext, in *krm.APIGatewayAPIObservedState) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Api_State](mapCtx, in.State)
	return out
}
func APIGatewayAPIObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krmapigatewayv1beta1.APIGatewayAPIObservedState {
	if in == nil {
		return nil
	}
	out := &krmapigatewayv1beta1.APIGatewayAPIObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func APIGatewayAPIObservedState_ToProto(mapCtx *direct.MapContext, in *krmapigatewayv1beta1.APIGatewayAPIObservedState) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Api_State](mapCtx, in.State)
	return out
}
func APIGatewayAPISpec_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.APIGatewayAPISpec {
	if in == nil {
		return nil
	}
	out := &krm.APIGatewayAPISpec{}
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ManagedService = direct.LazyPtr(in.GetManagedService())
	return out
}
func APIGatewayAPISpec_ToProto(mapCtx *direct.MapContext, in *krm.APIGatewayAPISpec) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ManagedService = direct.ValueOf(in.ManagedService)
	return out
}
func APIGatewayAPISpec_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krmapigatewayv1beta1.APIGatewayAPISpec {
	if in == nil {
		return nil
	}
	out := &krmapigatewayv1beta1.APIGatewayAPISpec{}
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ManagedService = direct.LazyPtr(in.GetManagedService())
	return out
}
func APIGatewayAPISpec_ToProto(mapCtx *direct.MapContext, in *krmapigatewayv1beta1.APIGatewayAPISpec) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ManagedService = direct.ValueOf(in.ManagedService)
	return out
}
