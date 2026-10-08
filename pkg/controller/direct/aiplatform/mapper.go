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

package aiplatform

import (
	"encoding/json"

	aiplatformpb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/wrapperspb"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func Int32Value_FromProto(mapCtx *direct.MapContext, in *wrapperspb.Int32Value) *krm.Int32Value {
	if in == nil {
		return nil
	}
	out := in.Value
	return &krm.Int32Value{
		Value: &out,
	}
}

func Int32Value_ToProto(mapCtx *direct.MapContext, in *krm.Int32Value) *wrapperspb.Int32Value {
	if in == nil || in.Value == nil {
		return nil
	}
	return wrapperspb.Int32(*in.Value)
}

func Schema_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.Schema) *krm.Schema {
	if in == nil {
		return nil
	}
	out := &krm.Schema{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Format = direct.LazyPtr(in.GetFormat())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Nullable = direct.LazyPtr(in.GetNullable())
	out.Default = Value_FromProto(mapCtx, in.GetDefault())
	if in.GetItems() != nil {
		nested := Schema_FromProto(mapCtx, in.GetItems())
		if nested != nil {
			b, err := json.Marshal(nested)
			if err != nil {
				mapCtx.Errorf("error marshalling nested schema to JSON: %v", err)
			} else {
				out.Items = &apiextensionsv1.JSON{Raw: b}
			}
		}
	}
	out.MinItems = direct.LazyPtr(in.GetMinItems())
	out.MaxItems = direct.LazyPtr(in.GetMaxItems())
	out.Enum = in.Enum
	out.PropertyOrdering = in.PropertyOrdering
	out.Required = in.Required
	out.MinProperties = direct.LazyPtr(in.GetMinProperties())
	out.MaxProperties = direct.LazyPtr(in.GetMaxProperties())
	out.Minimum = direct.LazyPtr(in.GetMinimum())
	out.Maximum = direct.LazyPtr(in.GetMaximum())
	out.MinLength = direct.LazyPtr(in.GetMinLength())
	out.MaxLength = direct.LazyPtr(in.GetMaxLength())
	out.Pattern = direct.LazyPtr(in.GetPattern())
	out.Example = Value_FromProto(mapCtx, in.GetExample())
	if in.AnyOf != nil {
		out.AnyOf = make([]apiextensionsv1.JSON, 0, len(in.AnyOf))
		for _, x := range in.AnyOf {
			nested := Schema_FromProto(mapCtx, x)
			if nested != nil {
				b, err := json.Marshal(nested)
				if err != nil {
					mapCtx.Errorf("error marshalling nested anyOf schema to JSON: %v", err)
				} else {
					out.AnyOf = append(out.AnyOf, apiextensionsv1.JSON{Raw: b})
				}
			}
		}
	}
	out.AdditionalProperties = Value_FromProto(mapCtx, in.GetAdditionalProperties())
	out.Ref = direct.LazyPtr(in.GetRef())
	return out
}

func Schema_ToProto(mapCtx *direct.MapContext, in *krm.Schema) *aiplatformpb.Schema {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.Schema{}
	out.Type = direct.Enum_ToProto[aiplatformpb.Type](mapCtx, in.Type)
	out.Format = direct.ValueOf(in.Format)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Nullable = direct.ValueOf(in.Nullable)
	out.Default = Value_ToProto(mapCtx, in.Default)
	if in.Items != nil {
		var nested krm.Schema
		if err := json.Unmarshal(in.Items.Raw, &nested); err != nil {
			mapCtx.Errorf("error unmarshalling nested schema from JSON: %v", err)
		} else {
			out.Items = Schema_ToProto(mapCtx, &nested)
		}
	}
	out.MinItems = direct.ValueOf(in.MinItems)
	out.MaxItems = direct.ValueOf(in.MaxItems)
	out.Enum = in.Enum
	out.PropertyOrdering = in.PropertyOrdering
	out.Required = in.Required
	out.MinProperties = direct.ValueOf(in.MinProperties)
	out.MaxProperties = direct.ValueOf(in.MaxProperties)
	out.Minimum = direct.ValueOf(in.Minimum)
	out.Maximum = direct.ValueOf(in.Maximum)
	out.MinLength = direct.ValueOf(in.MinLength)
	out.MaxLength = direct.ValueOf(in.MaxLength)
	out.Pattern = direct.ValueOf(in.Pattern)
	out.Example = Value_ToProto(mapCtx, in.Example)
	if len(in.AnyOf) > 0 {
		out.AnyOf = make([]*aiplatformpb.Schema, 0, len(in.AnyOf))
		for _, x := range in.AnyOf {
			var nested krm.Schema
			if err := json.Unmarshal(x.Raw, &nested); err != nil {
				mapCtx.Errorf("error unmarshalling anyOf schema from JSON: %v", err)
			} else {
				out.AnyOf = append(out.AnyOf, Schema_ToProto(mapCtx, &nested))
			}
		}
	}
	out.AdditionalProperties = Value_ToProto(mapCtx, in.AdditionalProperties)
	out.Ref = direct.ValueOf(in.Ref)
	return out
}
