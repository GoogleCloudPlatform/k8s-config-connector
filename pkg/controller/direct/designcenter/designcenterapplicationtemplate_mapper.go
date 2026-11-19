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
	"strconv"

	pb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/designcenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
)

func Value_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *krm.Value {
	if in == nil {
		return nil
	}
	out := &krm.Value{}
	switch kind := in.GetKind().(type) {
	case *structpb.Value_NullValue:
		v := kind.NullValue.String()
		out.NullValue = &v
	case *structpb.Value_NumberValue:
		v := kind.NumberValue
		out.NumberValue = &v
	case *structpb.Value_StringValue:
		v := kind.StringValue
		out.StringValue = &v
	case *structpb.Value_BoolValue:
		v := kind.BoolValue
		out.BoolValue = &v
	case *structpb.Value_StructValue:
		if val := direct.Struct_FromProto(mapCtx, kind.StructValue); val != nil {
			out.StructValue = *val
		}
	case *structpb.Value_ListValue:
		if kind.ListValue != nil {
			out.ListValue = &krm.ListValue{}
			for _, elem := range kind.ListValue.GetValues() {
				if elemVal := Value_FromProto(mapCtx, elem); elemVal != nil {
					out.ListValue.Values = append(out.ListValue.Values, *elemVal)
				}
			}
		}
	}
	return out
}

func Value_ToProto(mapCtx *direct.MapContext, in *krm.Value) *structpb.Value {
	if in == nil {
		return nil
	}
	out := &structpb.Value{}
	if in.NullValue != nil {
		strVal := direct.ValueOf(in.NullValue)
		var val int
		if v, ok := structpb.NullValue_value[strVal]; ok {
			val = int(v)
		} else if parsed, err := strconv.Atoi(strVal); err == nil {
			val = parsed
		}
		out.Kind = &structpb.Value_NullValue{NullValue: structpb.NullValue(val)}
	} else if in.NumberValue != nil {
		out.Kind = &structpb.Value_NumberValue{NumberValue: direct.ValueOf(in.NumberValue)}
	} else if in.StringValue != nil {
		out.Kind = &structpb.Value_StringValue{StringValue: direct.ValueOf(in.StringValue)}
	} else if in.BoolValue != nil {
		out.Kind = &structpb.Value_BoolValue{BoolValue: direct.ValueOf(in.BoolValue)}
	} else if len(in.StructValue.Raw) > 0 {
		out.Kind = &structpb.Value_StructValue{StructValue: direct.Struct_ToProto(mapCtx, &in.StructValue)}
	} else if in.ListValue != nil {
		lv := &structpb.ListValue{}
		for _, elem := range in.ListValue.Values {
			if elemProto := Value_ToProto(mapCtx, &elem); elemProto != nil {
				lv.Values = append(lv.Values, elemProto)
			}
		}
		out.Kind = &structpb.Value_ListValue{ListValue: lv}
	}
	return out
}

func TfBlueprintMetadataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TFBlueprintMetadata) *krm.TfBlueprintMetadataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TfBlueprintMetadataObservedState{}
	out.TerraformInput = direct.Slice_FromProto(mapCtx, in.GetTerraformInput(), TerraformInputObservedState_FromProto)
	out.TerraformOutput = direct.Slice_FromProto(mapCtx, in.GetTerraformOutput(), TerraformOutputObservedState_FromProto)
	out.APIs = in.GetApis()
	out.Roles = in.GetRoles()
	out.DeploymentDuration = direct.StringDuration_FromProto(mapCtx, in.GetDeploymentDuration())
	out.IconURI = direct.LazyPtr(in.GetIconUri())
	out.Connections = direct.Slice_FromProto(mapCtx, in.GetConnections(), TerraformInputConnectionsObservedState_FromProto)
	out.UiMetadata = TerraformBlueprintUiMetadataObservedState_FromProto(mapCtx, in.GetUiMetadata())
	out.ProviderVersions = direct.Slice_FromProto(mapCtx, in.GetProviderVersions(), ProviderVersionObservedState_FromProto)
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
