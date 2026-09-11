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

package apiextensionsv1

import (
	"encoding/json"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type JSON = v1.JSON

func JSON_FromProto(mapCtx *direct.MapContext, in *structpb.Struct) *JSON {
	if in == nil {
		return nil
	}
	b, err := json.Marshal(in)
	if err != nil {
		mapCtx.Errorf("error marshalling structpb.Struct to JSON: %v", err)
		return nil
	}
	return &JSON{Raw: b}
}

func JSON_ToProto(mapCtx *direct.MapContext, in *JSON) *structpb.Struct {
	if in == nil {
		return nil
	}
	out := &structpb.Struct{}
	if err := json.Unmarshal(in.Raw, out); err != nil {
		mapCtx.Errorf("error unmarshalling JSON to structpb.Struct: %v", err)
		return nil
	}
	return out
}
