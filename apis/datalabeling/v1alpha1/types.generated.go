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

package v1alpha1


// +kcc:proto=google.cloud.datalabeling.v1beta1.Annotation
type Annotation struct {
	// Output only. Unique name of this annotation, format is:
	//
	//  projects/{project_id}/datasets/{dataset_id}/annotatedDatasets/{annotated_dataset}/examples/{example_id}/annotations/{annotation_id}
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Annotation.name
	Name *string `json:"name,omitempty"`

	// Output only. The source of the annotation.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Annotation.annotation_source
	AnnotationSource *string `json:"annotationSource,omitempty"`

	// Output only. This is the actual annotation value, e.g classification,
	//  bounding box values are stored here.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Annotation.annotation_value
	AnnotationValue *AnnotationValue `json:"annotationValue,omitempty"`

	// Output only. Annotation metadata, including information like votes
	//  for labels.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Annotation.annotation_metadata
	AnnotationMetadata *AnnotationMetadata `json:"annotationMetadata,omitempty"`

	// Output only. Sentiment for this annotation.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Annotation.annotation_sentiment
	AnnotationSentiment *string `json:"annotationSentiment,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.AnnotationMetadata
type AnnotationMetadata struct {
	// Metadata related to human labeling.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationMetadata.operator_metadata
	OperatorMetadata *OperatorMetadata `json:"operatorMetadata,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.AnnotationSpec
type AnnotationSpec struct {
	// Required. The display name of the AnnotationSpec. Maximum of 64 characters.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationSpec.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-provided description of the annotation specification.
	//  The description can be up to 10,000 characters long.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationSpec.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.AnnotationValue
type AnnotationValue struct {
	// Annotation value for image classification case.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationValue.image_classification_annotation
	ImageClassificationAnnotation *ImageClassificationAnnotation `json:"imageClassificationAnnotation,omitempty"`

	// Annotation value for image bounding box, oriented bounding box
	//  and polygon cases.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationValue.image_bounding_poly_annotation
	ImageBoundingPolyAnnotation *ImageBoundingPolyAnnotation `json:"imageBoundingPolyAnnotation,omitempty"`

	// Annotation value for image polyline cases.
	//  Polyline here is different from BoundingPoly. It is formed by
	//  line segments connected to each other but not closed form(Bounding Poly).
	//  The line segments can cross each other.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationValue.image_polyline_annotation
	ImagePolylineAnnotation *ImagePolylineAnnotation `json:"imagePolylineAnnotation,omitempty"`

	// Annotation value for image segmentation.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationValue.image_segmentation_annotation
	ImageSegmentationAnnotation *ImageSegmentationAnnotation `json:"imageSegmentationAnnotation,omitempty"`

	// Annotation value for text classification case.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationValue.text_classification_annotation
	TextClassificationAnnotation *TextClassificationAnnotation `json:"textClassificationAnnotation,omitempty"`

	// Annotation value for text entity extraction case.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationValue.text_entity_extraction_annotation
	TextEntityExtractionAnnotation *TextEntityExtractionAnnotation `json:"textEntityExtractionAnnotation,omitempty"`

	// Annotation value for video classification case.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationValue.video_classification_annotation
	VideoClassificationAnnotation *VideoClassificationAnnotation `json:"videoClassificationAnnotation,omitempty"`

	// Annotation value for video object detection and tracking case.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationValue.video_object_tracking_annotation
	VideoObjectTrackingAnnotation *VideoObjectTrackingAnnotation `json:"videoObjectTrackingAnnotation,omitempty"`

	// Annotation value for video event case.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationValue.video_event_annotation
	VideoEventAnnotation *VideoEventAnnotation `json:"videoEventAnnotation,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.BoundingPoly
type BoundingPoly struct {
	// The bounding polygon vertices.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.BoundingPoly.vertices
	Vertices []Vertex `json:"vertices,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.Example
type Example struct {
	// The image payload, a container of the image bytes/uri.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Example.image_payload
	ImagePayload *ImagePayload `json:"imagePayload,omitempty"`

	// The text payload, a container of the text content.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Example.text_payload
	TextPayload *TextPayload `json:"textPayload,omitempty"`

	// The video payload, a container of the video uri.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Example.video_payload
	VideoPayload *VideoPayload `json:"videoPayload,omitempty"`

	// Output only. Name of the example, in format of:
	//  projects/{project_id}/datasets/{dataset_id}/annotatedDatasets/
	//  {annotated_dataset_id}/examples/{example_id}
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Example.name
	Name *string `json:"name,omitempty"`

	// Output only. Annotations for the piece of data in Example.
	//  One piece of data can have multiple annotations.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Example.annotations
	Annotations []Annotation `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ImageBoundingPolyAnnotation
type ImageBoundingPolyAnnotation struct {
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImageBoundingPolyAnnotation.bounding_poly
	BoundingPoly *BoundingPoly `json:"boundingPoly,omitempty"`

	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImageBoundingPolyAnnotation.normalized_bounding_poly
	NormalizedBoundingPoly *NormalizedBoundingPoly `json:"normalizedBoundingPoly,omitempty"`

	// Label of object in this bounding polygon.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImageBoundingPolyAnnotation.annotation_spec
	AnnotationSpec *AnnotationSpec `json:"annotationSpec,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ImageClassificationAnnotation
type ImageClassificationAnnotation struct {
	// Label of image.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImageClassificationAnnotation.annotation_spec
	AnnotationSpec *AnnotationSpec `json:"annotationSpec,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ImagePayload
type ImagePayload struct {
	// Image format.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImagePayload.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// A byte string of a thumbnail image.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImagePayload.image_thumbnail
	ImageThumbnail []byte `json:"imageThumbnail,omitempty"`

	// Image uri from the user bucket.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImagePayload.image_uri
	ImageURI *string `json:"imageURI,omitempty"`

	// Signed uri of the image file in the service bucket.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImagePayload.signed_uri
	SignedURI *string `json:"signedURI,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ImagePolylineAnnotation
type ImagePolylineAnnotation struct {
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImagePolylineAnnotation.polyline
	Polyline *Polyline `json:"polyline,omitempty"`

	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImagePolylineAnnotation.normalized_polyline
	NormalizedPolyline *NormalizedPolyline `json:"normalizedPolyline,omitempty"`

	// Label of this polyline.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImagePolylineAnnotation.annotation_spec
	AnnotationSpec *AnnotationSpec `json:"annotationSpec,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ImageSegmentationAnnotation
type ImageSegmentationAnnotation struct {

	// TODO: unsupported map type with key string and value message


	// Image format.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImageSegmentationAnnotation.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// A byte string of a full image's color map.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ImageSegmentationAnnotation.image_bytes
	ImageBytes []byte `json:"imageBytes,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.NormalizedBoundingPoly
type NormalizedBoundingPoly struct {
	// The bounding polygon normalized vertices.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.NormalizedBoundingPoly.normalized_vertices
	NormalizedVertices []NormalizedVertex `json:"normalizedVertices,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.NormalizedPolyline
type NormalizedPolyline struct {
	// The normalized polyline vertices.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.NormalizedPolyline.normalized_vertices
	NormalizedVertices []NormalizedVertex `json:"normalizedVertices,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.NormalizedVertex
type NormalizedVertex struct {
	// X coordinate.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.NormalizedVertex.x
	X *float32 `json:"x,omitempty"`

	// Y coordinate.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.NormalizedVertex.y
	Y *float32 `json:"y,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.ObjectTrackingFrame
type ObjectTrackingFrame struct {
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ObjectTrackingFrame.bounding_poly
	BoundingPoly *BoundingPoly `json:"boundingPoly,omitempty"`

	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ObjectTrackingFrame.normalized_bounding_poly
	NormalizedBoundingPoly *NormalizedBoundingPoly `json:"normalizedBoundingPoly,omitempty"`

	// The time offset of this frame relative to the beginning of the video.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.ObjectTrackingFrame.time_offset
	TimeOffset *string `json:"timeOffset,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.OperatorMetadata
type OperatorMetadata struct {
	// Confidence score corresponding to a label. For examle, if 3 contributors
	//  have answered the question and 2 of them agree on the final label, the
	//  confidence score will be 0.67 (2/3).
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.OperatorMetadata.score
	Score *float32 `json:"score,omitempty"`

	// The total number of contributors that answer this question.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.OperatorMetadata.total_votes
	TotalVotes *int32 `json:"totalVotes,omitempty"`

	// The total number of contributors that choose this label.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.OperatorMetadata.label_votes
	LabelVotes *int32 `json:"labelVotes,omitempty"`

	// Comments from contributors.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.OperatorMetadata.comments
	Comments []string `json:"comments,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.Polyline
type Polyline struct {
	// The polyline vertices.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Polyline.vertices
	Vertices []Vertex `json:"vertices,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.SequentialSegment
type SequentialSegment struct {
	// Start position (inclusive).
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.SequentialSegment.start
	Start *int32 `json:"start,omitempty"`

	// End position (exclusive).
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.SequentialSegment.end
	End *int32 `json:"end,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.TextClassificationAnnotation
type TextClassificationAnnotation struct {
	// Label of the text.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.TextClassificationAnnotation.annotation_spec
	AnnotationSpec *AnnotationSpec `json:"annotationSpec,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.TextEntityExtractionAnnotation
type TextEntityExtractionAnnotation struct {
	// Label of the text entities.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.TextEntityExtractionAnnotation.annotation_spec
	AnnotationSpec *AnnotationSpec `json:"annotationSpec,omitempty"`

	// Position of the entity.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.TextEntityExtractionAnnotation.sequential_segment
	SequentialSegment *SequentialSegment `json:"sequentialSegment,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.TextPayload
type TextPayload struct {
	// Text content.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.TextPayload.text_content
	TextContent *string `json:"textContent,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.TimeSegment
type TimeSegment struct {
	// Start of the time segment (inclusive), represented as the duration since
	//  the example start.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.TimeSegment.start_time_offset
	StartTimeOffset *string `json:"startTimeOffset,omitempty"`

	// End of the time segment (exclusive), represented as the duration since the
	//  example start.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.TimeSegment.end_time_offset
	EndTimeOffset *string `json:"endTimeOffset,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.Vertex
type Vertex struct {
	// X coordinate.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Vertex.x
	X *int32 `json:"x,omitempty"`

	// Y coordinate.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Vertex.y
	Y *int32 `json:"y,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.VideoClassificationAnnotation
type VideoClassificationAnnotation struct {
	// The time segment of the video to which the annotation applies.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoClassificationAnnotation.time_segment
	TimeSegment *TimeSegment `json:"timeSegment,omitempty"`

	// Label of the segment specified by time_segment.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoClassificationAnnotation.annotation_spec
	AnnotationSpec *AnnotationSpec `json:"annotationSpec,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.VideoEventAnnotation
type VideoEventAnnotation struct {
	// Label of the event in this annotation.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoEventAnnotation.annotation_spec
	AnnotationSpec *AnnotationSpec `json:"annotationSpec,omitempty"`

	// The time segment of the video to which the annotation applies.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoEventAnnotation.time_segment
	TimeSegment *TimeSegment `json:"timeSegment,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.VideoObjectTrackingAnnotation
type VideoObjectTrackingAnnotation struct {
	// Label of the object tracked in this annotation.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoObjectTrackingAnnotation.annotation_spec
	AnnotationSpec *AnnotationSpec `json:"annotationSpec,omitempty"`

	// The time segment of the video to which object tracking applies.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoObjectTrackingAnnotation.time_segment
	TimeSegment *TimeSegment `json:"timeSegment,omitempty"`

	// The list of frames where this object track appears.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoObjectTrackingAnnotation.object_tracking_frames
	ObjectTrackingFrames []ObjectTrackingFrame `json:"objectTrackingFrames,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.VideoPayload
type VideoPayload struct {
	// Video format.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoPayload.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Video uri from the user bucket.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoPayload.video_uri
	VideoURI *string `json:"videoURI,omitempty"`

	// The list of video thumbnails.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoPayload.video_thumbnails
	VideoThumbnails []VideoThumbnail `json:"videoThumbnails,omitempty"`

	// FPS of the video.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoPayload.frame_rate
	FrameRate *float32 `json:"frameRate,omitempty"`

	// Signed uri of the video file in the service bucket.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoPayload.signed_uri
	SignedURI *string `json:"signedURI,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.VideoThumbnail
type VideoThumbnail struct {
	// A byte string of the video frame.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoThumbnail.thumbnail
	Thumbnail []byte `json:"thumbnail,omitempty"`

	// Time offset relative to the beginning of the video, corresponding to the
	//  video frame where the thumbnail has been extracted from.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.VideoThumbnail.time_offset
	TimeOffset *string `json:"timeOffset,omitempty"`
}
