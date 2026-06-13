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

package vertexaihyperparametertuningjob

import (
	"strconv"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func Int32Value_FromProto(mapCtx *direct.MapContext, in *wrapperspb.Int32Value) *krm.Int32Value {
	if in == nil {
		return nil
	}
	out := &krm.Int32Value{}
	val := in.GetValue()
	out.Value = &val
	return out
}

func Int32Value_ToProto(mapCtx *direct.MapContext, in *krm.Int32Value) *wrapperspb.Int32Value {
	if in == nil {
		return nil
	}
	out := &wrapperspb.Int32Value{}
	if in.Value != nil {
		out.Value = *in.Value
	}
	return out
}

func Status_FromProto(mapCtx *direct.MapContext, in *status.Status) *common.Status {
	if in == nil {
		return nil
	}
	out := &common.Status{}
	code := in.GetCode()
	out.Code = &code
	msg := in.GetMessage()
	out.Message = &msg
	return out
}

func Status_ToProto(mapCtx *direct.MapContext, in *common.Status) *status.Status {
	if in == nil {
		return nil
	}
	out := &status.Status{}
	if in.Code != nil {
		out.Code = *in.Code
	}
	if in.Message != nil {
		out.Message = *in.Message
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
	if in.StructValue != nil {
		out.Kind = &structpb.Value_StructValue{
			StructValue: StructValue_ToProto(mapCtx, in.StructValue),
		}
	}
	return out
}

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
		out.StructValue = StructValue_FromProto(mapCtx, in.GetStructValue())
	}
	return out
}

func StructValue_FromProto(mapCtx *direct.MapContext, in *structpb.Struct) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string)
	for key, val := range in.Fields {
		out[key] = val.GetStringValue()
	}
	return out
}

func StructValue_ToProto(mapCtx *direct.MapContext, in map[string]string) *structpb.Struct {
	if in == nil {
		return nil
	}
	out := &structpb.Struct{}
	if len(in) > 0 {
		out.Fields = make(map[string]*structpb.Value)
	}
	for key, val := range in {
		value := &structpb.Value_StringValue{
			StringValue: val,
		}
		out.Fields[key] = &structpb.Value{
			Kind: value,
		}
	}
	return out
}
