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

package datalabeling

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datalabeling/v1alpha1"
)
func DataItem_FromProto(mapCtx *direct.MapContext, in *pb.DataItem) *krm.DataItem {
	if in == nil {
		return nil
	}
	out := &krm.DataItem{}
	out.ImagePayload = ImagePayload_FromProto(mapCtx, in.GetImagePayload())
	out.TextPayload = TextPayload_FromProto(mapCtx, in.GetTextPayload())
	out.VideoPayload = VideoPayload_FromProto(mapCtx, in.GetVideoPayload())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func DataItem_ToProto(mapCtx *direct.MapContext, in *krm.DataItem) *pb.DataItem {
	if in == nil {
		return nil
	}
	out := &pb.DataItem{}
	if oneof := ImagePayload_ToProto(mapCtx, in.ImagePayload); oneof != nil {
		out.Payload = &pb.DataItem_ImagePayload{ImagePayload: oneof}
	}
	if oneof := TextPayload_ToProto(mapCtx, in.TextPayload); oneof != nil {
		out.Payload = &pb.DataItem_TextPayload{TextPayload: oneof}
	}
	if oneof := VideoPayload_ToProto(mapCtx, in.VideoPayload); oneof != nil {
		out.Payload = &pb.DataItem_VideoPayload{VideoPayload: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func DatalabelingDataItemObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataItem) *krm.DatalabelingDataItemObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingDataItemObservedState{}
	// MISSING: ImagePayload
	// MISSING: TextPayload
	// MISSING: VideoPayload
	// MISSING: Name
	return out
}
func DatalabelingDataItemObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingDataItemObservedState) *pb.DataItem {
	if in == nil {
		return nil
	}
	out := &pb.DataItem{}
	// MISSING: ImagePayload
	// MISSING: TextPayload
	// MISSING: VideoPayload
	// MISSING: Name
	return out
}
func DatalabelingDataItemSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataItem) *krm.DatalabelingDataItemSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingDataItemSpec{}
	// MISSING: ImagePayload
	// MISSING: TextPayload
	// MISSING: VideoPayload
	// MISSING: Name
	return out
}
func DatalabelingDataItemSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingDataItemSpec) *pb.DataItem {
	if in == nil {
		return nil
	}
	out := &pb.DataItem{}
	// MISSING: ImagePayload
	// MISSING: TextPayload
	// MISSING: VideoPayload
	// MISSING: Name
	return out
}
func ImagePayload_FromProto(mapCtx *direct.MapContext, in *pb.ImagePayload) *krm.ImagePayload {
	if in == nil {
		return nil
	}
	out := &krm.ImagePayload{}
	out.MimeType = direct.LazyPtr(in.GetMimeType())
	out.ImageThumbnail = in.GetImageThumbnail()
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	out.SignedURI = direct.LazyPtr(in.GetSignedUri())
	return out
}
func ImagePayload_ToProto(mapCtx *direct.MapContext, in *krm.ImagePayload) *pb.ImagePayload {
	if in == nil {
		return nil
	}
	out := &pb.ImagePayload{}
	out.MimeType = direct.ValueOf(in.MimeType)
	out.ImageThumbnail = in.ImageThumbnail
	out.ImageUri = direct.ValueOf(in.ImageURI)
	out.SignedUri = direct.ValueOf(in.SignedURI)
	return out
}
func TextPayload_FromProto(mapCtx *direct.MapContext, in *pb.TextPayload) *krm.TextPayload {
	if in == nil {
		return nil
	}
	out := &krm.TextPayload{}
	out.TextContent = direct.LazyPtr(in.GetTextContent())
	return out
}
func TextPayload_ToProto(mapCtx *direct.MapContext, in *krm.TextPayload) *pb.TextPayload {
	if in == nil {
		return nil
	}
	out := &pb.TextPayload{}
	out.TextContent = direct.ValueOf(in.TextContent)
	return out
}
func VideoPayload_FromProto(mapCtx *direct.MapContext, in *pb.VideoPayload) *krm.VideoPayload {
	if in == nil {
		return nil
	}
	out := &krm.VideoPayload{}
	out.MimeType = direct.LazyPtr(in.GetMimeType())
	out.VideoURI = direct.LazyPtr(in.GetVideoUri())
	out.VideoThumbnails = direct.Slice_FromProto(mapCtx, in.VideoThumbnails, VideoThumbnail_FromProto)
	out.FrameRate = direct.LazyPtr(in.GetFrameRate())
	out.SignedURI = direct.LazyPtr(in.GetSignedUri())
	return out
}
func VideoPayload_ToProto(mapCtx *direct.MapContext, in *krm.VideoPayload) *pb.VideoPayload {
	if in == nil {
		return nil
	}
	out := &pb.VideoPayload{}
	out.MimeType = direct.ValueOf(in.MimeType)
	out.VideoUri = direct.ValueOf(in.VideoURI)
	out.VideoThumbnails = direct.Slice_ToProto(mapCtx, in.VideoThumbnails, VideoThumbnail_ToProto)
	out.FrameRate = direct.ValueOf(in.FrameRate)
	out.SignedUri = direct.ValueOf(in.SignedURI)
	return out
}
func VideoThumbnail_FromProto(mapCtx *direct.MapContext, in *pb.VideoThumbnail) *krm.VideoThumbnail {
	if in == nil {
		return nil
	}
	out := &krm.VideoThumbnail{}
	out.Thumbnail = in.GetThumbnail()
	out.TimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetTimeOffset())
	return out
}
func VideoThumbnail_ToProto(mapCtx *direct.MapContext, in *krm.VideoThumbnail) *pb.VideoThumbnail {
	if in == nil {
		return nil
	}
	out := &pb.VideoThumbnail{}
	out.Thumbnail = in.Thumbnail
	out.TimeOffset = direct.StringDuration_ToProto(mapCtx, in.TimeOffset)
	return out
}
