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

package designcenter

import (
	pb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/designcenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
)

func Value_ToProto(mapCtx *direct.MapContext, in *krm.Value) *structpb.Value {
	if in == nil {
		return nil
	}
	out := &structpb.Value{}
	if oneof := Value_NullValue_ToProto(mapCtx, in.NullValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := Value_NumberValue_ToProto(mapCtx, in.NumberValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := Value_StringValue_ToProto(mapCtx, in.StringValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := Value_BoolValue_ToProto(mapCtx, in.BoolValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := direct.Struct_ToProto(mapCtx, &in.StructValue); oneof != nil {
		out.Kind = &structpb.Value_StructValue{StructValue: oneof}
	}
	if oneof := ListValue_ToProto(mapCtx, in.ListValue); oneof != nil {
		out.Kind = &structpb.Value_ListValue{ListValue: oneof}
	}
	return out
}

func TfBlueprintMetadataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TFBlueprintMetadata) *krm.TfBlueprintMetadataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TfBlueprintMetadataObservedState{}
	out.TerraformInput = direct.Slice_FromProto(mapCtx, in.TerraformInput, TerraformInputObservedState_FromProto)
	out.TerraformOutput = direct.Slice_FromProto(mapCtx, in.TerraformOutput, TerraformOutputObservedState_FromProto)
	out.APIs = in.Apis
	out.Roles = in.Roles
	out.DeploymentDuration = direct.StringDuration_FromProto(mapCtx, in.GetDeploymentDuration())
	out.IconURI = direct.LazyPtr(in.GetIconUri())
	out.Connections = direct.Slice_FromProto(mapCtx, in.Connections, TerraformInputConnectionsObservedState_FromProto)
	out.UiMetadata = TerraformBlueprintUiMetadataObservedState_FromProto(mapCtx, in.GetUiMetadata())
	out.ProviderVersions = direct.Slice_FromProto(mapCtx, in.ProviderVersions, ProviderVersionObservedState_FromProto)
	return out
}

func TerraformBlueprintUiMetadataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TerraformBlueprintUiMetadata) *krm.TerraformBlueprintUiMetadataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TerraformBlueprintUiMetadataObservedState{}
	out.TerraformUiInput = TerraformUiInputObservedState_FromProto(mapCtx, in.GetTerraformUiInput())
	out.TerraformUiOutput = TerraformUiOutputObservedState_FromProto(mapCtx, in.GetTerraformUiOutput())
	return out
}

func TerraformBlueprintUiMetadataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TerraformBlueprintUiMetadataObservedState) *pb.TerraformBlueprintUiMetadata {
	if in == nil {
		return nil
	}
	out := &pb.TerraformBlueprintUiMetadata{}
	out.TerraformUiInput = TerraformUiInputObservedState_ToProto(mapCtx, in.TerraformUiInput)
	out.TerraformUiOutput = TerraformUiOutputObservedState_ToProto(mapCtx, in.TerraformUiOutput)
	return out
}

func TfBlueprintMetadataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TfBlueprintMetadataObservedState) *pb.TFBlueprintMetadata {
	if in == nil {
		return nil
	}
	out := &pb.TFBlueprintMetadata{}
	out.TerraformInput = direct.Slice_ToProto(mapCtx, in.TerraformInput, TerraformInputObservedState_ToProto)
	out.TerraformOutput = direct.Slice_ToProto(mapCtx, in.TerraformOutput, TerraformOutputObservedState_ToProto)
	out.Apis = in.APIs
	out.Roles = in.Roles
	out.DeploymentDuration = direct.StringDuration_ToProto(mapCtx, in.DeploymentDuration)
	out.IconUri = direct.ValueOf(in.IconURI)
	out.Connections = direct.Slice_ToProto(mapCtx, in.Connections, TerraformInputConnectionsObservedState_ToProto)
	out.UiMetadata = TerraformBlueprintUiMetadataObservedState_ToProto(mapCtx, in.UiMetadata)
	out.ProviderVersions = direct.Slice_ToProto(mapCtx, in.ProviderVersions, ProviderVersionObservedState_ToProto)
	return out
}

func HelmChartMetadataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.HelmChartMetadata) *krm.HelmChartMetadataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.HelmChartMetadataObservedState{}
	out.HelmInput = direct.Slice_FromProto(mapCtx, in.HelmInput, HelmChartInputObservedState_FromProto)
	out.HelmOutput = direct.Slice_FromProto(mapCtx, in.HelmOutput, HelmChartOutputObservedState_FromProto)
	return out
}

func HelmChartMetadataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.HelmChartMetadataObservedState) *pb.HelmChartMetadata {
	if in == nil {
		return nil
	}
	out := &pb.HelmChartMetadata{}
	out.HelmInput = direct.Slice_ToProto(mapCtx, in.HelmInput, HelmChartInputObservedState_ToProto)
	out.HelmOutput = direct.Slice_ToProto(mapCtx, in.HelmOutput, HelmChartOutputObservedState_ToProto)
	return out
}

func InferredMetadataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InferredMetadata) *krm.InferredMetadataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InferredMetadataObservedState{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.TemplateMetadata = TfBlueprintMetadataObservedState_FromProto(mapCtx, in.GetTemplateMetadata())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func InferredMetadataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InferredMetadataObservedState) *pb.InferredMetadata {
	if in == nil {
		return nil
	}
	out := &pb.InferredMetadata{}
	out.Description = direct.ValueOf(in.Description)
	out.TemplateMetadata = TfBlueprintMetadataObservedState_ToProto(mapCtx, in.TemplateMetadata)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
