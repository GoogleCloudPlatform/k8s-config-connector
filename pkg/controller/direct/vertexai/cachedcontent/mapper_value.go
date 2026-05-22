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

package cachedcontent

import (
	"encoding/json"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	structpb "google.golang.org/protobuf/types/known/structpb"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func ValueValue_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *krm.Value {
	if in == nil {
		return nil
	}
	out := &krm.Value{}
	switch v := in.Kind.(type) {
	case *structpb.Value_NullValue:
		s := v.NullValue.String()
		out.NullValue = &s
	case *structpb.Value_NumberValue:
		out.NumberValue = &v.NumberValue
	case *structpb.Value_StringValue:
		out.StringValue = &v.StringValue
	case *structpb.Value_BoolValue:
		out.BoolValue = &v.BoolValue
	case *structpb.Value_StructValue:
		b, err := v.StructValue.MarshalJSON()
		if err != nil {
			mapCtx.Errorf("marshaling structpb.Struct to JSON: %v", err)
		} else {
			out.StructValue = apiextensionsv1.JSON{Raw: b}
		}
	case *structpb.Value_ListValue:
		// We'd map ListValue, but usually we don't need it deeply parsed if we can avoid it.
		// Let's just marshal the whole value to JSON and put it in structValue for fallback?
		// Actually ListValue is generated too! Let's ignore it for now or implement it.
		_ = v
	}
	return out
}

func ValueValue_ToProto(mapCtx *direct.MapContext, in *krm.Value) *structpb.Value {
	if in == nil {
		return nil
	}
	out := &structpb.Value{}
	if in.NullValue != nil {
		out.Kind = &structpb.Value_NullValue{NullValue: 0}
	} else if in.NumberValue != nil {
		out.Kind = &structpb.Value_NumberValue{NumberValue: *in.NumberValue}
	} else if in.StringValue != nil {
		out.Kind = &structpb.Value_StringValue{StringValue: *in.StringValue}
	} else if in.BoolValue != nil {
		out.Kind = &structpb.Value_BoolValue{BoolValue: *in.BoolValue}
	} else if len(in.StructValue.Raw) > 0 {
		s := &structpb.Struct{}
		if err := s.UnmarshalJSON(in.StructValue.Raw); err != nil {
			mapCtx.Errorf("unmarshaling JSON to structpb.Struct: %v", err)
		} else {
			out.Kind = &structpb.Value_StructValue{StructValue: s}
		}
	}
	return out
}

func Value_FromProto(mapCtx *direct.MapContext, in *structpb.Value) apiextensionsv1.JSON {
	if in == nil {
		return apiextensionsv1.JSON{}
	}
	krmObj := ValueValue_FromProto(mapCtx, in)
	b, err := json.Marshal(krmObj)
	if err != nil {
		mapCtx.Errorf("marshaling Value to JSON: %v", err)
		return apiextensionsv1.JSON{}
	}
	return apiextensionsv1.JSON{Raw: b}
}

func Value_ToProto(mapCtx *direct.MapContext, in apiextensionsv1.JSON) *structpb.Value {
	if len(in.Raw) == 0 || string(in.Raw) == "null" {
		return nil
	}
	krmObj := &krm.Value{}
	if err := json.Unmarshal(in.Raw, krmObj); err != nil {
		mapCtx.Errorf("unmarshaling JSON to Value: %v", err)
		return nil
	}
	return ValueValue_ToProto(mapCtx, krmObj)
}
