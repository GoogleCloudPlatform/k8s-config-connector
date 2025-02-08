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


// +kcc:proto=google.cloud.vision.v1.BoundingPoly
type BoundingPoly struct {
	// The bounding polygon vertices.
	// +kcc:proto:field=google.cloud.vision.v1.BoundingPoly.vertices
	Vertices []Vertex `json:"vertices,omitempty"`

	// The bounding polygon normalized vertices.
	// +kcc:proto:field=google.cloud.vision.v1.BoundingPoly.normalized_vertices
	NormalizedVertices []NormalizedVertex `json:"normalizedVertices,omitempty"`
}

// +kcc:proto=google.cloud.vision.v1.NormalizedVertex
type NormalizedVertex struct {
	// X coordinate.
	// +kcc:proto:field=google.cloud.vision.v1.NormalizedVertex.x
	X *float32 `json:"x,omitempty"`

	// Y coordinate.
	// +kcc:proto:field=google.cloud.vision.v1.NormalizedVertex.y
	Y *float32 `json:"y,omitempty"`
}

// +kcc:proto=google.cloud.vision.v1.ReferenceImage
type ReferenceImage struct {
	// The resource name of the reference image.
	//
	//  Format is:
	//  `projects/PROJECT_ID/locations/LOC_ID/products/PRODUCT_ID/referenceImages/IMAGE_ID`.
	//
	//  This field is ignored when creating a reference image.
	// +kcc:proto:field=google.cloud.vision.v1.ReferenceImage.name
	Name *string `json:"name,omitempty"`

	// Required. The Google Cloud Storage URI of the reference image.
	//
	//  The URI must start with `gs://`.
	// +kcc:proto:field=google.cloud.vision.v1.ReferenceImage.uri
	URI *string `json:"uri,omitempty"`

	// Optional. Bounding polygons around the areas of interest in the reference
	//  image. If this field is empty, the system will try to detect regions of
	//  interest. At most 10 bounding polygons will be used.
	//
	//  The provided shape is converted into a non-rotated rectangle. Once
	//  converted, the small edge of the rectangle must be greater than or equal
	//  to 300 pixels. The aspect ratio must be 1:4 or less (i.e. 1:3 is ok; 1:5
	//  is not).
	// +kcc:proto:field=google.cloud.vision.v1.ReferenceImage.bounding_polys
	BoundingPolys []BoundingPoly `json:"boundingPolys,omitempty"`
}

// +kcc:proto=google.cloud.vision.v1.Vertex
type Vertex struct {
	// X coordinate.
	// +kcc:proto:field=google.cloud.vision.v1.Vertex.x
	X *int32 `json:"x,omitempty"`

	// Y coordinate.
	// +kcc:proto:field=google.cloud.vision.v1.Vertex.y
	Y *int32 `json:"y,omitempty"`
}
