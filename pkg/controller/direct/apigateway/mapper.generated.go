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

package apigateway

import (
	pb "cloud.google.com/go/apigateway/apiv1/apigatewaypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigateway/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

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
func ApiConfig_FromProto(mapCtx *direct.MapContext, in *pb.ApiConfig) *krm.ApiConfig {
	if in == nil {
		return nil
	}
	out := &krm.ApiConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.GatewayServiceAccount = direct.LazyPtr(in.GetGatewayServiceAccount())
	// MISSING: ServiceConfigID
	// MISSING: State
	out.OpenapiDocuments = direct.Slice_FromProto(mapCtx, in.OpenapiDocuments, ApiConfig_OpenApiDocument_FromProto)
	out.GrpcServices = direct.Slice_FromProto(mapCtx, in.GrpcServices, ApiConfig_GrpcServiceDefinition_FromProto)
	out.ManagedServiceConfigs = direct.Slice_FromProto(mapCtx, in.ManagedServiceConfigs, ApiConfig_File_FromProto)
	return out
}
func ApiConfig_ToProto(mapCtx *direct.MapContext, in *krm.ApiConfig) *pb.ApiConfig {
	if in == nil {
		return nil
	}
	out := &pb.ApiConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.GatewayServiceAccount = direct.ValueOf(in.GatewayServiceAccount)
	// MISSING: ServiceConfigID
	// MISSING: State
	out.OpenapiDocuments = direct.Slice_ToProto(mapCtx, in.OpenapiDocuments, ApiConfig_OpenApiDocument_ToProto)
	out.GrpcServices = direct.Slice_ToProto(mapCtx, in.GrpcServices, ApiConfig_GrpcServiceDefinition_ToProto)
	out.ManagedServiceConfigs = direct.Slice_ToProto(mapCtx, in.ManagedServiceConfigs, ApiConfig_File_ToProto)
	return out
}
func ApiConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ApiConfig) *krm.ApiConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApiConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: GatewayServiceAccount
	out.ServiceConfigID = direct.LazyPtr(in.GetServiceConfigId())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: OpenapiDocuments
	// MISSING: GrpcServices
	// MISSING: ManagedServiceConfigs
	return out
}
func ApiConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApiConfigObservedState) *pb.ApiConfig {
	if in == nil {
		return nil
	}
	out := &pb.ApiConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: GatewayServiceAccount
	out.ServiceConfigId = direct.ValueOf(in.ServiceConfigID)
	out.State = direct.Enum_ToProto[pb.ApiConfig_State](mapCtx, in.State)
	// MISSING: OpenapiDocuments
	// MISSING: GrpcServices
	// MISSING: ManagedServiceConfigs
	return out
}
func ApiConfig_File_FromProto(mapCtx *direct.MapContext, in *pb.ApiConfig_File) *krm.ApiConfig_File {
	if in == nil {
		return nil
	}
	out := &krm.ApiConfig_File{}
	out.Path = direct.LazyPtr(in.GetPath())
	out.Contents = in.GetContents()
	return out
}
func ApiConfig_File_ToProto(mapCtx *direct.MapContext, in *krm.ApiConfig_File) *pb.ApiConfig_File {
	if in == nil {
		return nil
	}
	out := &pb.ApiConfig_File{}
	out.Path = direct.ValueOf(in.Path)
	out.Contents = in.Contents
	return out
}
func ApiConfig_GrpcServiceDefinition_FromProto(mapCtx *direct.MapContext, in *pb.ApiConfig_GrpcServiceDefinition) *krm.ApiConfig_GrpcServiceDefinition {
	if in == nil {
		return nil
	}
	out := &krm.ApiConfig_GrpcServiceDefinition{}
	out.FileDescriptorSet = ApiConfig_File_FromProto(mapCtx, in.GetFileDescriptorSet())
	out.Source = direct.Slice_FromProto(mapCtx, in.Source, ApiConfig_File_FromProto)
	return out
}
func ApiConfig_GrpcServiceDefinition_ToProto(mapCtx *direct.MapContext, in *krm.ApiConfig_GrpcServiceDefinition) *pb.ApiConfig_GrpcServiceDefinition {
	if in == nil {
		return nil
	}
	out := &pb.ApiConfig_GrpcServiceDefinition{}
	out.FileDescriptorSet = ApiConfig_File_ToProto(mapCtx, in.FileDescriptorSet)
	out.Source = direct.Slice_ToProto(mapCtx, in.Source, ApiConfig_File_ToProto)
	return out
}
func ApiConfig_OpenApiDocument_FromProto(mapCtx *direct.MapContext, in *pb.ApiConfig_OpenApiDocument) *krm.ApiConfig_OpenApiDocument {
	if in == nil {
		return nil
	}
	out := &krm.ApiConfig_OpenApiDocument{}
	out.Document = ApiConfig_File_FromProto(mapCtx, in.GetDocument())
	return out
}
func ApiConfig_OpenApiDocument_ToProto(mapCtx *direct.MapContext, in *krm.ApiConfig_OpenApiDocument) *pb.ApiConfig_OpenApiDocument {
	if in == nil {
		return nil
	}
	out := &pb.ApiConfig_OpenApiDocument{}
	out.Document = ApiConfig_File_ToProto(mapCtx, in.Document)
	return out
}
