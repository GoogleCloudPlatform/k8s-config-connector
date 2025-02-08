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

package vision

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vision/apiv1p4beta1/visionpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vision/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BoundingPoly_FromProto(mapCtx *direct.MapContext, in *pb.BoundingPoly) *krm.BoundingPoly {
	if in == nil {
		return nil
	}
	out := &krm.BoundingPoly{}
	out.Vertices = direct.Slice_FromProto(mapCtx, in.Vertices, Vertex_FromProto)
	out.NormalizedVertices = direct.Slice_FromProto(mapCtx, in.NormalizedVertices, NormalizedVertex_FromProto)
	return out
}
func BoundingPoly_ToProto(mapCtx *direct.MapContext, in *krm.BoundingPoly) *pb.BoundingPoly {
	if in == nil {
		return nil
	}
	out := &pb.BoundingPoly{}
	out.Vertices = direct.Slice_ToProto(mapCtx, in.Vertices, Vertex_ToProto)
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
func ReferenceImage_FromProto(mapCtx *direct.MapContext, in *pb.ReferenceImage) *krm.ReferenceImage {
	if in == nil {
		return nil
	}
	out := &krm.ReferenceImage{}
	out.Name = direct.LazyPtr(in.GetName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.BoundingPolys = direct.Slice_FromProto(mapCtx, in.BoundingPolys, BoundingPoly_FromProto)
	return out
}
func ReferenceImage_ToProto(mapCtx *direct.MapContext, in *krm.ReferenceImage) *pb.ReferenceImage {
	if in == nil {
		return nil
	}
	out := &pb.ReferenceImage{}
	out.Name = direct.ValueOf(in.Name)
	out.Uri = direct.ValueOf(in.URI)
	out.BoundingPolys = direct.Slice_ToProto(mapCtx, in.BoundingPolys, BoundingPoly_ToProto)
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
