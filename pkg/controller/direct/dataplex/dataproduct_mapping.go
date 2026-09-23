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

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AccessGroups_FromProto(mapCtx *direct.MapContext, in map[string]*pb.DataProduct_AccessGroup) map[string]v1alpha1.DataProduct_AccessGroup {
	if in == nil {
		return nil
	}
	out := make(map[string]v1alpha1.DataProduct_AccessGroup)
	for k, v := range in {
		mapped := DataProduct_AccessGroup_FromProto(mapCtx, v)
		if mapped != nil {
			out[k] = *mapped
		}
	}
	return out
}

func AccessGroups_ToProto(mapCtx *direct.MapContext, in map[string]v1alpha1.DataProduct_AccessGroup) map[string]*pb.DataProduct_AccessGroup {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.DataProduct_AccessGroup)
	for k, v := range in {
		mapped := DataProduct_AccessGroup_ToProto(mapCtx, &v)
		if mapped != nil {
			out[k] = mapped
		}
	}
	return out
}
