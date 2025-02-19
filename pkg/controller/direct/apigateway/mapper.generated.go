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
