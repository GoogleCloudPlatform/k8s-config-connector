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

package rapidmigrationassessment

import (
	pb "cloud.google.com/go/rapidmigrationassessment/apiv1/rapidmigrationassessmentpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/rapidmigrationassessment/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Annotation_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.Annotation {
	if in == nil {
		return nil
	}
	out := &krm.Annotation{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Annotation_ToProto(mapCtx *direct.MapContext, in *krm.Annotation) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Type = direct.Enum_ToProto[pb.Annotation_Type](mapCtx, in.Type)
	return out
}
func AnnotationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.AnnotationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnnotationObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Type
	return out
}
func AnnotationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnnotationObservedState) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Type
	return out
}
func RapidmigrationassessmentAnnotationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.RapidmigrationassessmentAnnotationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RapidmigrationassessmentAnnotationObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	return out
}
func RapidmigrationassessmentAnnotationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RapidmigrationassessmentAnnotationObservedState) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	return out
}
func RapidmigrationassessmentAnnotationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.RapidmigrationassessmentAnnotationSpec {
	if in == nil {
		return nil
	}
	out := &krm.RapidmigrationassessmentAnnotationSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	return out
}
func RapidmigrationassessmentAnnotationSpec_ToProto(mapCtx *direct.MapContext, in *krm.RapidmigrationassessmentAnnotationSpec) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	return out
}
