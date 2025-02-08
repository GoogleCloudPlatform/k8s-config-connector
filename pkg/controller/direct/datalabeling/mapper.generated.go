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
func Annotation_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.Annotation {
	if in == nil {
		return nil
	}
	out := &krm.Annotation{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AnnotationSource = direct.Enum_FromProto(mapCtx, in.GetAnnotationSource())
	out.AnnotationValue = AnnotationValue_FromProto(mapCtx, in.GetAnnotationValue())
	out.AnnotationMetadata = AnnotationMetadata_FromProto(mapCtx, in.GetAnnotationMetadata())
	out.AnnotationSentiment = direct.Enum_FromProto(mapCtx, in.GetAnnotationSentiment())
	return out
}
func Annotation_ToProto(mapCtx *direct.MapContext, in *krm.Annotation) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	out.Name = direct.ValueOf(in.Name)
	out.AnnotationSource = direct.Enum_ToProto[pb.AnnotationSource](mapCtx, in.AnnotationSource)
	out.AnnotationValue = AnnotationValue_ToProto(mapCtx, in.AnnotationValue)
	out.AnnotationMetadata = AnnotationMetadata_ToProto(mapCtx, in.AnnotationMetadata)
	out.AnnotationSentiment = direct.Enum_ToProto[pb.AnnotationSentiment](mapCtx, in.AnnotationSentiment)
	return out
}
func AnnotationSpec_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpec) *krm.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &krm.AnnotationSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func AnnotationSpec_ToProto(mapCtx *direct.MapContext, in *krm.AnnotationSpec) *pb.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpec{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func AnnotationValue_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationValue) *krm.AnnotationValue {
	if in == nil {
		return nil
	}
	out := &krm.AnnotationValue{}
	out.ImageClassificationAnnotation = ImageClassificationAnnotation_FromProto(mapCtx, in.GetImageClassificationAnnotation())
	out.ImageBoundingPolyAnnotation = ImageBoundingPolyAnnotation_FromProto(mapCtx, in.GetImageBoundingPolyAnnotation())
	out.ImagePolylineAnnotation = ImagePolylineAnnotation_FromProto(mapCtx, in.GetImagePolylineAnnotation())
	out.ImageSegmentationAnnotation = ImageSegmentationAnnotation_FromProto(mapCtx, in.GetImageSegmentationAnnotation())
	out.TextClassificationAnnotation = TextClassificationAnnotation_FromProto(mapCtx, in.GetTextClassificationAnnotation())
	out.TextEntityExtractionAnnotation = TextEntityExtractionAnnotation_FromProto(mapCtx, in.GetTextEntityExtractionAnnotation())
	out.VideoClassificationAnnotation = VideoClassificationAnnotation_FromProto(mapCtx, in.GetVideoClassificationAnnotation())
	out.VideoObjectTrackingAnnotation = VideoObjectTrackingAnnotation_FromProto(mapCtx, in.GetVideoObjectTrackingAnnotation())
	out.VideoEventAnnotation = VideoEventAnnotation_FromProto(mapCtx, in.GetVideoEventAnnotation())
	return out
}
func AnnotationValue_ToProto(mapCtx *direct.MapContext, in *krm.AnnotationValue) *pb.AnnotationValue {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationValue{}
	if oneof := ImageClassificationAnnotation_ToProto(mapCtx, in.ImageClassificationAnnotation); oneof != nil {
		out.ValueType = &pb.AnnotationValue_ImageClassificationAnnotation{ImageClassificationAnnotation: oneof}
	}
	if oneof := ImageBoundingPolyAnnotation_ToProto(mapCtx, in.ImageBoundingPolyAnnotation); oneof != nil {
		out.ValueType = &pb.AnnotationValue_ImageBoundingPolyAnnotation{ImageBoundingPolyAnnotation: oneof}
	}
	if oneof := ImagePolylineAnnotation_ToProto(mapCtx, in.ImagePolylineAnnotation); oneof != nil {
		out.ValueType = &pb.AnnotationValue_ImagePolylineAnnotation{ImagePolylineAnnotation: oneof}
	}
	if oneof := ImageSegmentationAnnotation_ToProto(mapCtx, in.ImageSegmentationAnnotation); oneof != nil {
		out.ValueType = &pb.AnnotationValue_ImageSegmentationAnnotation{ImageSegmentationAnnotation: oneof}
	}
	if oneof := TextClassificationAnnotation_ToProto(mapCtx, in.TextClassificationAnnotation); oneof != nil {
		out.ValueType = &pb.AnnotationValue_TextClassificationAnnotation{TextClassificationAnnotation: oneof}
	}
	if oneof := TextEntityExtractionAnnotation_ToProto(mapCtx, in.TextEntityExtractionAnnotation); oneof != nil {
		out.ValueType = &pb.AnnotationValue_TextEntityExtractionAnnotation{TextEntityExtractionAnnotation: oneof}
	}
	if oneof := VideoClassificationAnnotation_ToProto(mapCtx, in.VideoClassificationAnnotation); oneof != nil {
		out.ValueType = &pb.AnnotationValue_VideoClassificationAnnotation{VideoClassificationAnnotation: oneof}
	}
	if oneof := VideoObjectTrackingAnnotation_ToProto(mapCtx, in.VideoObjectTrackingAnnotation); oneof != nil {
		out.ValueType = &pb.AnnotationValue_VideoObjectTrackingAnnotation{VideoObjectTrackingAnnotation: oneof}
	}
	if oneof := VideoEventAnnotation_ToProto(mapCtx, in.VideoEventAnnotation); oneof != nil {
		out.ValueType = &pb.AnnotationValue_VideoEventAnnotation{VideoEventAnnotation: oneof}
	}
	return out
}
func BoundingPoly_FromProto(mapCtx *direct.MapContext, in *pb.BoundingPoly) *krm.BoundingPoly {
	if in == nil {
		return nil
	}
	out := &krm.BoundingPoly{}
	out.Vertices = direct.Slice_FromProto(mapCtx, in.Vertices, Vertex_FromProto)
	return out
}
func BoundingPoly_ToProto(mapCtx *direct.MapContext, in *krm.BoundingPoly) *pb.BoundingPoly {
	if in == nil {
		return nil
	}
	out := &pb.BoundingPoly{}
	out.Vertices = direct.Slice_ToProto(mapCtx, in.Vertices, Vertex_ToProto)
	return out
}
func DatalabelingExampleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Example) *krm.DatalabelingExampleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingExampleObservedState{}
	// MISSING: ImagePayload
	// MISSING: TextPayload
	// MISSING: VideoPayload
	// MISSING: Name
	// MISSING: Annotations
	return out
}
func DatalabelingExampleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingExampleObservedState) *pb.Example {
	if in == nil {
		return nil
	}
	out := &pb.Example{}
	// MISSING: ImagePayload
	// MISSING: TextPayload
	// MISSING: VideoPayload
	// MISSING: Name
	// MISSING: Annotations
	return out
}
func DatalabelingExampleSpec_FromProto(mapCtx *direct.MapContext, in *pb.Example) *krm.DatalabelingExampleSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingExampleSpec{}
	// MISSING: ImagePayload
	// MISSING: TextPayload
	// MISSING: VideoPayload
	// MISSING: Name
	// MISSING: Annotations
	return out
}
func DatalabelingExampleSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingExampleSpec) *pb.Example {
	if in == nil {
		return nil
	}
	out := &pb.Example{}
	// MISSING: ImagePayload
	// MISSING: TextPayload
	// MISSING: VideoPayload
	// MISSING: Name
	// MISSING: Annotations
	return out
}
func Example_FromProto(mapCtx *direct.MapContext, in *pb.Example) *krm.Example {
	if in == nil {
		return nil
	}
	out := &krm.Example{}
	out.ImagePayload = ImagePayload_FromProto(mapCtx, in.GetImagePayload())
	out.TextPayload = TextPayload_FromProto(mapCtx, in.GetTextPayload())
	out.VideoPayload = VideoPayload_FromProto(mapCtx, in.GetVideoPayload())
	out.Name = direct.LazyPtr(in.GetName())
	out.Annotations = direct.Slice_FromProto(mapCtx, in.Annotations, Annotation_FromProto)
	return out
}
func Example_ToProto(mapCtx *direct.MapContext, in *krm.Example) *pb.Example {
	if in == nil {
		return nil
	}
	out := &pb.Example{}
	if oneof := ImagePayload_ToProto(mapCtx, in.ImagePayload); oneof != nil {
		out.Payload = &pb.Example_ImagePayload{ImagePayload: oneof}
	}
	if oneof := TextPayload_ToProto(mapCtx, in.TextPayload); oneof != nil {
		out.Payload = &pb.Example_TextPayload{TextPayload: oneof}
	}
	if oneof := VideoPayload_ToProto(mapCtx, in.VideoPayload); oneof != nil {
		out.Payload = &pb.Example_VideoPayload{VideoPayload: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.Annotations = direct.Slice_ToProto(mapCtx, in.Annotations, Annotation_ToProto)
	return out
}
func ImageBoundingPolyAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.ImageBoundingPolyAnnotation) *krm.ImageBoundingPolyAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.ImageBoundingPolyAnnotation{}
	out.BoundingPoly = BoundingPoly_FromProto(mapCtx, in.GetBoundingPoly())
	out.NormalizedBoundingPoly = NormalizedBoundingPoly_FromProto(mapCtx, in.GetNormalizedBoundingPoly())
	out.AnnotationSpec = AnnotationSpec_FromProto(mapCtx, in.GetAnnotationSpec())
	return out
}
func ImageBoundingPolyAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.ImageBoundingPolyAnnotation) *pb.ImageBoundingPolyAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.ImageBoundingPolyAnnotation{}
	if oneof := BoundingPoly_ToProto(mapCtx, in.BoundingPoly); oneof != nil {
		out.BoundedArea = &pb.ImageBoundingPolyAnnotation_BoundingPoly{BoundingPoly: oneof}
	}
	if oneof := NormalizedBoundingPoly_ToProto(mapCtx, in.NormalizedBoundingPoly); oneof != nil {
		out.BoundedArea = &pb.ImageBoundingPolyAnnotation_NormalizedBoundingPoly{NormalizedBoundingPoly: oneof}
	}
	out.AnnotationSpec = AnnotationSpec_ToProto(mapCtx, in.AnnotationSpec)
	return out
}
func ImageClassificationAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.ImageClassificationAnnotation) *krm.ImageClassificationAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.ImageClassificationAnnotation{}
	out.AnnotationSpec = AnnotationSpec_FromProto(mapCtx, in.GetAnnotationSpec())
	return out
}
func ImageClassificationAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.ImageClassificationAnnotation) *pb.ImageClassificationAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.ImageClassificationAnnotation{}
	out.AnnotationSpec = AnnotationSpec_ToProto(mapCtx, in.AnnotationSpec)
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
func ImagePolylineAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.ImagePolylineAnnotation) *krm.ImagePolylineAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.ImagePolylineAnnotation{}
	out.Polyline = Polyline_FromProto(mapCtx, in.GetPolyline())
	out.NormalizedPolyline = NormalizedPolyline_FromProto(mapCtx, in.GetNormalizedPolyline())
	out.AnnotationSpec = AnnotationSpec_FromProto(mapCtx, in.GetAnnotationSpec())
	return out
}
func ImagePolylineAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.ImagePolylineAnnotation) *pb.ImagePolylineAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.ImagePolylineAnnotation{}
	if oneof := Polyline_ToProto(mapCtx, in.Polyline); oneof != nil {
		out.Poly = &pb.ImagePolylineAnnotation_Polyline{Polyline: oneof}
	}
	if oneof := NormalizedPolyline_ToProto(mapCtx, in.NormalizedPolyline); oneof != nil {
		out.Poly = &pb.ImagePolylineAnnotation_NormalizedPolyline{NormalizedPolyline: oneof}
	}
	out.AnnotationSpec = AnnotationSpec_ToProto(mapCtx, in.AnnotationSpec)
	return out
}
func ImageSegmentationAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.ImageSegmentationAnnotation) *krm.ImageSegmentationAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.ImageSegmentationAnnotation{}
	// MISSING: AnnotationColors
	out.MimeType = direct.LazyPtr(in.GetMimeType())
	out.ImageBytes = in.GetImageBytes()
	return out
}
func ImageSegmentationAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.ImageSegmentationAnnotation) *pb.ImageSegmentationAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.ImageSegmentationAnnotation{}
	// MISSING: AnnotationColors
	out.MimeType = direct.ValueOf(in.MimeType)
	out.ImageBytes = in.ImageBytes
	return out
}
func NormalizedBoundingPoly_FromProto(mapCtx *direct.MapContext, in *pb.NormalizedBoundingPoly) *krm.NormalizedBoundingPoly {
	if in == nil {
		return nil
	}
	out := &krm.NormalizedBoundingPoly{}
	out.NormalizedVertices = direct.Slice_FromProto(mapCtx, in.NormalizedVertices, NormalizedVertex_FromProto)
	return out
}
func NormalizedBoundingPoly_ToProto(mapCtx *direct.MapContext, in *krm.NormalizedBoundingPoly) *pb.NormalizedBoundingPoly {
	if in == nil {
		return nil
	}
	out := &pb.NormalizedBoundingPoly{}
	out.NormalizedVertices = direct.Slice_ToProto(mapCtx, in.NormalizedVertices, NormalizedVertex_ToProto)
	return out
}
func NormalizedPolyline_FromProto(mapCtx *direct.MapContext, in *pb.NormalizedPolyline) *krm.NormalizedPolyline {
	if in == nil {
		return nil
	}
	out := &krm.NormalizedPolyline{}
	out.NormalizedVertices = direct.Slice_FromProto(mapCtx, in.NormalizedVertices, NormalizedVertex_FromProto)
	return out
}
func NormalizedPolyline_ToProto(mapCtx *direct.MapContext, in *krm.NormalizedPolyline) *pb.NormalizedPolyline {
	if in == nil {
		return nil
	}
	out := &pb.NormalizedPolyline{}
	out.NormalizedVertices = direct.Slice_ToProto(mapCtx, in.NormalizedVertices, NormalizedVertex_ToProto)
	return out
}
func NormalizedVertex_FromProto(mapCtx *direct.MapContext, in *pb.NormalizedVertex) *krm.NormalizedVertex {
	if in == nil {
		return nil
	}
	out := &krm.NormalizedVertex{}
	out.X = direct.LazyPtr(in.GetX())
	out.Y = direct.LazyPtr(in.GetY())
	return out
}
func NormalizedVertex_ToProto(mapCtx *direct.MapContext, in *krm.NormalizedVertex) *pb.NormalizedVertex {
	if in == nil {
		return nil
	}
	out := &pb.NormalizedVertex{}
	out.X = direct.ValueOf(in.X)
	out.Y = direct.ValueOf(in.Y)
	return out
}
func ObjectTrackingFrame_FromProto(mapCtx *direct.MapContext, in *pb.ObjectTrackingFrame) *krm.ObjectTrackingFrame {
	if in == nil {
		return nil
	}
	out := &krm.ObjectTrackingFrame{}
	out.BoundingPoly = BoundingPoly_FromProto(mapCtx, in.GetBoundingPoly())
	out.NormalizedBoundingPoly = NormalizedBoundingPoly_FromProto(mapCtx, in.GetNormalizedBoundingPoly())
	out.TimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetTimeOffset())
	return out
}
func ObjectTrackingFrame_ToProto(mapCtx *direct.MapContext, in *krm.ObjectTrackingFrame) *pb.ObjectTrackingFrame {
	if in == nil {
		return nil
	}
	out := &pb.ObjectTrackingFrame{}
	if oneof := BoundingPoly_ToProto(mapCtx, in.BoundingPoly); oneof != nil {
		out.BoundedArea = &pb.ObjectTrackingFrame_BoundingPoly{BoundingPoly: oneof}
	}
	if oneof := NormalizedBoundingPoly_ToProto(mapCtx, in.NormalizedBoundingPoly); oneof != nil {
		out.BoundedArea = &pb.ObjectTrackingFrame_NormalizedBoundingPoly{NormalizedBoundingPoly: oneof}
	}
	out.TimeOffset = direct.StringDuration_ToProto(mapCtx, in.TimeOffset)
	return out
}
func Polyline_FromProto(mapCtx *direct.MapContext, in *pb.Polyline) *krm.Polyline {
	if in == nil {
		return nil
	}
	out := &krm.Polyline{}
	out.Vertices = direct.Slice_FromProto(mapCtx, in.Vertices, Vertex_FromProto)
	return out
}
func Polyline_ToProto(mapCtx *direct.MapContext, in *krm.Polyline) *pb.Polyline {
	if in == nil {
		return nil
	}
	out := &pb.Polyline{}
	out.Vertices = direct.Slice_ToProto(mapCtx, in.Vertices, Vertex_ToProto)
	return out
}
func SequentialSegment_FromProto(mapCtx *direct.MapContext, in *pb.SequentialSegment) *krm.SequentialSegment {
	if in == nil {
		return nil
	}
	out := &krm.SequentialSegment{}
	out.Start = direct.LazyPtr(in.GetStart())
	out.End = direct.LazyPtr(in.GetEnd())
	return out
}
func SequentialSegment_ToProto(mapCtx *direct.MapContext, in *krm.SequentialSegment) *pb.SequentialSegment {
	if in == nil {
		return nil
	}
	out := &pb.SequentialSegment{}
	out.Start = direct.ValueOf(in.Start)
	out.End = direct.ValueOf(in.End)
	return out
}
func TextClassificationAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.TextClassificationAnnotation) *krm.TextClassificationAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.TextClassificationAnnotation{}
	out.AnnotationSpec = AnnotationSpec_FromProto(mapCtx, in.GetAnnotationSpec())
	return out
}
func TextClassificationAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.TextClassificationAnnotation) *pb.TextClassificationAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.TextClassificationAnnotation{}
	out.AnnotationSpec = AnnotationSpec_ToProto(mapCtx, in.AnnotationSpec)
	return out
}
func TextEntityExtractionAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.TextEntityExtractionAnnotation) *krm.TextEntityExtractionAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.TextEntityExtractionAnnotation{}
	out.AnnotationSpec = AnnotationSpec_FromProto(mapCtx, in.GetAnnotationSpec())
	out.SequentialSegment = SequentialSegment_FromProto(mapCtx, in.GetSequentialSegment())
	return out
}
func TextEntityExtractionAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.TextEntityExtractionAnnotation) *pb.TextEntityExtractionAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.TextEntityExtractionAnnotation{}
	out.AnnotationSpec = AnnotationSpec_ToProto(mapCtx, in.AnnotationSpec)
	out.SequentialSegment = SequentialSegment_ToProto(mapCtx, in.SequentialSegment)
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
func TimeSegment_FromProto(mapCtx *direct.MapContext, in *pb.TimeSegment) *krm.TimeSegment {
	if in == nil {
		return nil
	}
	out := &krm.TimeSegment{}
	out.StartTimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetStartTimeOffset())
	out.EndTimeOffset = direct.StringDuration_FromProto(mapCtx, in.GetEndTimeOffset())
	return out
}
func TimeSegment_ToProto(mapCtx *direct.MapContext, in *krm.TimeSegment) *pb.TimeSegment {
	if in == nil {
		return nil
	}
	out := &pb.TimeSegment{}
	out.StartTimeOffset = direct.StringDuration_ToProto(mapCtx, in.StartTimeOffset)
	out.EndTimeOffset = direct.StringDuration_ToProto(mapCtx, in.EndTimeOffset)
	return out
}
func Vertex_FromProto(mapCtx *direct.MapContext, in *pb.Vertex) *krm.Vertex {
	if in == nil {
		return nil
	}
	out := &krm.Vertex{}
	out.X = direct.LazyPtr(in.GetX())
	out.Y = direct.LazyPtr(in.GetY())
	return out
}
func Vertex_ToProto(mapCtx *direct.MapContext, in *krm.Vertex) *pb.Vertex {
	if in == nil {
		return nil
	}
	out := &pb.Vertex{}
	out.X = direct.ValueOf(in.X)
	out.Y = direct.ValueOf(in.Y)
	return out
}
func VideoClassificationAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.VideoClassificationAnnotation) *krm.VideoClassificationAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.VideoClassificationAnnotation{}
	out.TimeSegment = TimeSegment_FromProto(mapCtx, in.GetTimeSegment())
	out.AnnotationSpec = AnnotationSpec_FromProto(mapCtx, in.GetAnnotationSpec())
	return out
}
func VideoClassificationAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.VideoClassificationAnnotation) *pb.VideoClassificationAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.VideoClassificationAnnotation{}
	out.TimeSegment = TimeSegment_ToProto(mapCtx, in.TimeSegment)
	out.AnnotationSpec = AnnotationSpec_ToProto(mapCtx, in.AnnotationSpec)
	return out
}
func VideoEventAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.VideoEventAnnotation) *krm.VideoEventAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.VideoEventAnnotation{}
	out.AnnotationSpec = AnnotationSpec_FromProto(mapCtx, in.GetAnnotationSpec())
	out.TimeSegment = TimeSegment_FromProto(mapCtx, in.GetTimeSegment())
	return out
}
func VideoEventAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.VideoEventAnnotation) *pb.VideoEventAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.VideoEventAnnotation{}
	out.AnnotationSpec = AnnotationSpec_ToProto(mapCtx, in.AnnotationSpec)
	out.TimeSegment = TimeSegment_ToProto(mapCtx, in.TimeSegment)
	return out
}
func VideoObjectTrackingAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.VideoObjectTrackingAnnotation) *krm.VideoObjectTrackingAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.VideoObjectTrackingAnnotation{}
	out.AnnotationSpec = AnnotationSpec_FromProto(mapCtx, in.GetAnnotationSpec())
	out.TimeSegment = TimeSegment_FromProto(mapCtx, in.GetTimeSegment())
	out.ObjectTrackingFrames = direct.Slice_FromProto(mapCtx, in.ObjectTrackingFrames, ObjectTrackingFrame_FromProto)
	return out
}
func VideoObjectTrackingAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.VideoObjectTrackingAnnotation) *pb.VideoObjectTrackingAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.VideoObjectTrackingAnnotation{}
	out.AnnotationSpec = AnnotationSpec_ToProto(mapCtx, in.AnnotationSpec)
	out.TimeSegment = TimeSegment_ToProto(mapCtx, in.TimeSegment)
	out.ObjectTrackingFrames = direct.Slice_ToProto(mapCtx, in.ObjectTrackingFrames, ObjectTrackingFrame_ToProto)
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
