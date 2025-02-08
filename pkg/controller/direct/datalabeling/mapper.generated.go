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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datalabeling/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
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
func AnnotationSpecSet_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpecSet) *krm.AnnotationSpecSet {
	if in == nil {
		return nil
	}
	out := &krm.AnnotationSpecSet{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.AnnotationSpecs = direct.Slice_FromProto(mapCtx, in.AnnotationSpecs, AnnotationSpec_FromProto)
	out.BlockingResources = in.BlockingResources
	return out
}
func AnnotationSpecSet_ToProto(mapCtx *direct.MapContext, in *krm.AnnotationSpecSet) *pb.AnnotationSpecSet {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpecSet{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.AnnotationSpecs = direct.Slice_ToProto(mapCtx, in.AnnotationSpecs, AnnotationSpec_ToProto)
	out.BlockingResources = in.BlockingResources
	return out
}
func DatalabelingAnnotationSpecSetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpecSet) *krm.DatalabelingAnnotationSpecSetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingAnnotationSpecSetObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: AnnotationSpecs
	// MISSING: BlockingResources
	return out
}
func DatalabelingAnnotationSpecSetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingAnnotationSpecSetObservedState) *pb.AnnotationSpecSet {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpecSet{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: AnnotationSpecs
	// MISSING: BlockingResources
	return out
}
func DatalabelingAnnotationSpecSetSpec_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpecSet) *krm.DatalabelingAnnotationSpecSetSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingAnnotationSpecSetSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: AnnotationSpecs
	// MISSING: BlockingResources
	return out
}
func DatalabelingAnnotationSpecSetSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingAnnotationSpecSetSpec) *pb.AnnotationSpecSet {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpecSet{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: AnnotationSpecs
	// MISSING: BlockingResources
	return out
}
