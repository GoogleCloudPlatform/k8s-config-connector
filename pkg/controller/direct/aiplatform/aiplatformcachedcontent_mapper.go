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
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"google.golang.org/genproto/googleapis/type/latlng"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// mapper.generated.go calls these helpers for AIPlatformCachedContent, but
// generate-mapper does not write them:
//   - google.type.LatLng is not in the aiplatform proto packages.
//     ToolConfig.retrieval_config uses it.
//   - CachedContent.UsageMetadata is a message nested in CachedContent.
//     generate-mapper writes the ObservedState mapper that calls it, but not
//     the function itself.

func LatLng_FromProto(mapCtx *direct.MapContext, in *latlng.LatLng) *krm.LatLng {
	if in == nil {
		return nil
	}
	out := &krm.LatLng{}
	out.Latitude = direct.LazyPtr(in.GetLatitude())
	out.Longitude = direct.LazyPtr(in.GetLongitude())
	return out
}

func LatLng_ToProto(mapCtx *direct.MapContext, in *krm.LatLng) *latlng.LatLng {
	if in == nil {
		return nil
	}
	out := &latlng.LatLng{}
	out.Latitude = direct.ValueOf(in.Latitude)
	out.Longitude = direct.ValueOf(in.Longitude)
	return out
}

func CachedContent_UsageMetadata_FromProto(mapCtx *direct.MapContext, in *pb.CachedContent_UsageMetadata) *krm.CachedContent_UsageMetadata {
	if in == nil {
		return nil
	}
	out := &krm.CachedContent_UsageMetadata{}
	out.TotalTokenCount = direct.LazyPtr(in.GetTotalTokenCount())
	out.TextCount = direct.LazyPtr(in.GetTextCount())
	out.ImageCount = direct.LazyPtr(in.GetImageCount())
	out.VideoDurationSeconds = direct.LazyPtr(in.GetVideoDurationSeconds())
	out.AudioDurationSeconds = direct.LazyPtr(in.GetAudioDurationSeconds())
	return out
}

func CachedContent_UsageMetadata_ToProto(mapCtx *direct.MapContext, in *krm.CachedContent_UsageMetadata) *pb.CachedContent_UsageMetadata {
	if in == nil {
		return nil
	}
	out := &pb.CachedContent_UsageMetadata{}
	out.TotalTokenCount = direct.ValueOf(in.TotalTokenCount)
	out.TextCount = direct.ValueOf(in.TextCount)
	out.ImageCount = direct.ValueOf(in.ImageCount)
	out.VideoDurationSeconds = direct.ValueOf(in.VideoDurationSeconds)
	out.AudioDurationSeconds = direct.ValueOf(in.AudioDurationSeconds)
	return out
}
