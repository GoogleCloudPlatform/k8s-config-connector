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

package vectorsearch

import (
	pb "cloud.google.com/go/vectorsearch/apiv1/vectorsearchpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vectorsearch/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VectorSchema_FromProto(mapCtx *direct.MapContext, in map[string]*pb.VectorField) map[string]krm.VectorField {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.VectorField)
	for k, v := range in {
		if v != nil {
			mapped := VectorField_FromProto(mapCtx, v)
			if mapped != nil {
				out[k] = *mapped
			}
		}
	}
	return out
}

func VectorSchema_ToProto(mapCtx *direct.MapContext, in map[string]krm.VectorField) map[string]*pb.VectorField {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.VectorField)
	for k, v := range in {
		mapped := VectorField_ToProto(mapCtx, &v)
		if mapped != nil {
			out[k] = mapped
		}
	}
	return out
}
