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

package firestore

import (
	"encoding/json"

	pb "cloud.google.com/go/firestore/apiv1/firestorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func FirestoreDocumentSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.FirestoreDocumentSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreDocumentSpec{}

	if in.Fields != nil {
		out.Fields = make(map[string]apiextensionsv1.JSON, len(in.Fields))
	}
	for k, v := range in.Fields {
		outV := Field_FromProto(mapCtx, v)

		j, err := json.Marshal(outV)
		if err != nil {
			return nil
		}
		out.Fields[k] = apiextensionsv1.JSON{Raw: j}
	}

	return out
}

func Field_FromProto(mapCtx *direct.MapContext, in *pb.Value) any {
	if in == nil {
		return nil
	}
	switch v := in.ValueType.(type) {
	case *pb.Value_NullValue:
		return nil
	case *pb.Value_BooleanValue:
		return v.BooleanValue
	case *pb.Value_IntegerValue:
		return v.IntegerValue
	case *pb.Value_DoubleValue:
		return v.DoubleValue
	case *pb.Value_StringValue:
		return v.StringValue

		// These types do not easily round-trip to JSON, so we omit them for now.
	// case *pb.Value_BytesValue:
	// 	return v.BytesValue
	// case *pb.Value_TimestampValue:
	// 	return v.TimestampValue
	// case *pb.Value_ReferenceValue:
	// 	return v.ReferenceValue
	// case *pb.Value_GeoPointValue:
	// 	return v.GeoPointValue

	case *pb.Value_ArrayValue:
		arr := make([]any, len(v.ArrayValue.Values))
		for i, elem := range v.ArrayValue.Values {
			arr[i] = Field_FromProto(mapCtx, elem)
		}
		return arr
	case *pb.Value_MapValue:
		m := make(map[string]any)
		for k, elem := range v.MapValue.Fields {
			m[k] = Field_FromProto(mapCtx, elem)
		}
		return m
	default:
		// Unknown type
		mapCtx.Errorf("unknown type in FirestoreDocument: %T", v)
		return nil
	}
}

func FirestoreDocumentSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreDocumentSpec) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	if in.Fields != nil {
		out.Fields = make(map[string]*pb.Value, len(in.Fields))
	}
	for k, v := range in.Fields {
		if len(v.Raw) == 0 {
			out.Fields[k] = &pb.Value{ValueType: &pb.Value_NullValue{}}
			continue
		}

		var wrappedValue any
		if err := json.Unmarshal(v.Raw, &wrappedValue); err != nil {
			mapCtx.Errorf("failed to unmarshal JSON field in FirestoreDocument: %v", err)
			out.Fields[k] = &pb.Value{ValueType: &pb.Value_NullValue{}}
			continue
		}

		outV := Field_ToProto(mapCtx, wrappedValue)
		out.Fields[k] = outV
	}

	return out
}

func Field_ToProto(mapCtx *direct.MapContext, in any) *pb.Value {
	if in == nil {
		return &pb.Value{ValueType: &pb.Value_NullValue{}}
	}

	switch in := in.(type) {
	case bool:
		return &pb.Value{ValueType: &pb.Value_BooleanValue{BooleanValue: in}}
	case string:
		return &pb.Value{ValueType: &pb.Value_StringValue{StringValue: in}}
	case int64:
		return &pb.Value{ValueType: &pb.Value_IntegerValue{IntegerValue: in}}
	case float64:
		return &pb.Value{ValueType: &pb.Value_DoubleValue{DoubleValue: in}}
	case []any:
		arr := make([]*pb.Value, len(in))
		for i, elem := range in {
			arr[i] = Field_ToProto(mapCtx, elem)
		}
		return &pb.Value{ValueType: &pb.Value_ArrayValue{ArrayValue: &pb.ArrayValue{Values: arr}}}
	case map[string]any:
		m := make(map[string]*pb.Value, len(in))
		for k, elem := range in {
			m[k] = Field_ToProto(mapCtx, elem)
		}
		return &pb.Value{ValueType: &pb.Value_MapValue{MapValue: &pb.MapValue{Fields: m}}}
	default:
		// Unknown type
		mapCtx.Errorf("unknown type in FirestoreDocument: %T", in)
		return nil
	}
}

func Value_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Value) *krm.Value {
	mapCtx.NotImplemented()
	return nil
}

func Value_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Value) *pb.Value {
	mapCtx.NotImplemented()
	return nil
}

// This mapper is unused and the default generator generates incorrect code (that does not compile).
// We provide a stub implementation to avoid the generation of incorrect code.
func Value_BytesValue_ToProto(mapCtx *direct.MapContext, in []byte) *pb.Value_BytesValue {
	mapCtx.NotImplemented()
	return nil
}
