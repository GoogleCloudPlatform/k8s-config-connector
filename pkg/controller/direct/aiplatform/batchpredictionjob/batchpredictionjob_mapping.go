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

package batchpredictionjob

import (
	"strconv"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
)

func BatchPredictionJobEncryptionSpec_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.BatchPredictionJobEncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.BatchPredictionJobEncryptionSpec{}
	if in.KmsKeyName != "" {
		out.KMSKeyRef = &refs.KMSCryptoKeyRef{External: in.KmsKeyName}
	}
	return out
}

func BatchPredictionJobEncryptionSpec_ToProto(mapCtx *direct.MapContext, in *krm.BatchPredictionJobEncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	if in.KMSKeyRef != nil {
		out.KmsKeyName = in.KMSKeyRef.External
	}
	return out
}

func ListValue_FromProto(mapCtx *direct.MapContext, in *structpb.ListValue) *krm.ListValue {
	if in == nil {
		return nil
	}
	out := &krm.ListValue{}
	for _, val := range in.Values {
		mappedVal := Value_FromProto(mapCtx, val)
		if mappedVal != nil {
			out.Values = append(out.Values, *mappedVal)
		}
	}
	return out
}

func ListValue_ToProto(mapCtx *direct.MapContext, in *krm.ListValue) *structpb.ListValue {
	if in == nil {
		return nil
	}
	out := &structpb.ListValue{}
	for _, val := range in.Values {
		mappedVal := Value_ToProto(mapCtx, &val)
		if mappedVal != nil {
			out.Values = append(out.Values, mappedVal)
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
