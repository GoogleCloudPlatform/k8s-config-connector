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

package ces

import (
	"strconv"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ces/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
)

func Value_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *krm.Value {
	if in == nil {
		return nil
	}
	out := &krm.Value{}
	switch in.GetKind().(type) {
	case *structpb.Value_StringValue:
		value := in.GetStringValue()
		out.StringValue = &value
	case *structpb.Value_NumberValue:
		value := in.GetNumberValue()
		out.NumberValue = &value
	case *structpb.Value_NullValue:
		value := in.GetNullValue().String()
		out.NullValue = &value
	case *structpb.Value_BoolValue:
		value := in.GetBoolValue()
		out.BoolValue = &value
	case *structpb.Value_StructValue:
		if val := direct.Struct_FromProto(mapCtx, in.GetStructValue()); val != nil {
			out.StructValue = *val
		}
	}
	return out
}

func Value_ToProto(mapCtx *direct.MapContext, in *krm.Value) *structpb.Value {
	if in == nil {
		return nil
	}
	out := &structpb.Value{}
	if in.BoolValue != nil {
		out.Kind = &structpb.Value_BoolValue{
			BoolValue: direct.ValueOf(in.BoolValue),
		}
	}
	if in.NullValue != nil {
		value, err := strconv.Atoi(direct.ValueOf(in.NullValue))
		if err != nil {
			mapCtx.Errorf("error converting value %s from string to int", direct.ValueOf(in.NullValue))
		}
		out.Kind = &structpb.Value_NullValue{
			NullValue: structpb.NullValue(value),
		}
	}
	if in.NumberValue != nil {
		out.Kind = &structpb.Value_NumberValue{
			NumberValue: direct.ValueOf(in.NumberValue),
		}
	}
	if in.StringValue != nil {
		out.Kind = &structpb.Value_StringValue{
			StringValue: direct.ValueOf(in.StringValue),
		}
	}
	if len(in.StructValue.Raw) > 0 {
		out.Kind = &structpb.Value_StructValue{
			StructValue: direct.Struct_ToProto(mapCtx, &in.StructValue),
		}
	}
	return out
}
