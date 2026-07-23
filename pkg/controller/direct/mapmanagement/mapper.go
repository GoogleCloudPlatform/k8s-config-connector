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

package mapmanagement

import (
	pb "cloud.google.com/go/maps/mapmanagement/apiv2beta/mapmanagementpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/mapmanagement/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func MapFeatures_FromProto(mapCtx *direct.MapContext, in *pb.MapFeatures) *krm.MapFeatures {
	if in == nil {
		return nil
	}
	out := &krm.MapFeatures{}
	if in.PoiBoostLevel != nil {
		out.PoiBoostLevel = direct.LazyPtr(in.GetPoiBoostLevel())
	}
	out.SimpleFeatures = direct.EnumSlice_FromProto(mapCtx, in.GetSimpleFeatures())
	return out
}

func MapFeatures_ToProto(mapCtx *direct.MapContext, in *krm.MapFeatures) *pb.MapFeatures {
	if in == nil {
		return nil
	}
	out := &pb.MapFeatures{}
	if in.PoiBoostLevel != nil {
		out.PoiBoostLevel = direct.LazyPtr(direct.ValueOf(in.PoiBoostLevel))
	}
	out.SimpleFeatures = direct.EnumSlice_ToProto[pb.MapFeatures_SimpleFeature](mapCtx, in.SimpleFeatures)
	return out
}
