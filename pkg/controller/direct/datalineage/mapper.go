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

package datalineage

import (
	"encoding/json"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func Attributes_FromProto(mapCtx *direct.MapContext, in map[string]*structpb.Value) map[string]apiextensionsv1.JSON {
	if in == nil {
		return nil
	}
	out := make(map[string]apiextensionsv1.JSON)
	for k, v := range in {
		b, err := json.Marshal(v.AsInterface())
		if err != nil {
			mapCtx.Errorf("marshalling structpb.Value to json: %v", err)
			continue
		}
		out[k] = apiextensionsv1.JSON{Raw: b}
	}
	return out
}

func Attributes_ToProto(mapCtx *direct.MapContext, in map[string]apiextensionsv1.JSON) map[string]*structpb.Value {
	if in == nil {
		return nil
	}
	out := make(map[string]*structpb.Value)
	for k, v := range in {
		if len(v.Raw) == 0 {
			continue
		}
		var val any
		if err := json.Unmarshal(v.Raw, &val); err != nil {
			mapCtx.Errorf("unmarshalling json to value: %v", err)
			continue
		}
		pbVal, err := structpb.NewValue(val)
		if err != nil {
			mapCtx.Errorf("error converting value to structpb.Value: %v", err)
			continue
		}
		out[k] = pbVal
	}
	return out
}
