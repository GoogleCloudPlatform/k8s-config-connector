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

package aiplatform

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AiplatformAnnotationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.AiplatformAnnotationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformAnnotationObservedState{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func AiplatformAnnotationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformAnnotationObservedState) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func AiplatformAnnotationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.AiplatformAnnotationSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformAnnotationSpec{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func AiplatformAnnotationSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformAnnotationSpec) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func Annotation_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.Annotation {
	if in == nil {
		return nil
	}
	out := &krm.Annotation{}
	// MISSING: Name
	out.PayloadSchemaURI = direct.LazyPtr(in.GetPayloadSchemaUri())
	out.Payload = Value_FromProto(mapCtx, in.GetPayload())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: AnnotationSource
	out.Labels = in.Labels
	return out
}
func Annotation_ToProto(mapCtx *direct.MapContext, in *krm.Annotation) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	// MISSING: Name
	out.PayloadSchemaUri = direct.ValueOf(in.PayloadSchemaURI)
	out.Payload = Value_ToProto(mapCtx, in.Payload)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: AnnotationSource
	out.Labels = in.Labels
	return out
}
func AnnotationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.AnnotationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnnotationObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Etag
	out.AnnotationSource = UserActionReference_FromProto(mapCtx, in.GetAnnotationSource())
	// MISSING: Labels
	return out
}
func AnnotationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnnotationObservedState) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Etag
	out.AnnotationSource = UserActionReference_ToProto(mapCtx, in.AnnotationSource)
	// MISSING: Labels
	return out
}
func UserActionReference_FromProto(mapCtx *direct.MapContext, in *pb.UserActionReference) *krm.UserActionReference {
	if in == nil {
		return nil
	}
	out := &krm.UserActionReference{}
	out.Operation = direct.LazyPtr(in.GetOperation())
	out.DataLabelingJob = direct.LazyPtr(in.GetDataLabelingJob())
	out.Method = direct.LazyPtr(in.GetMethod())
	return out
}
func UserActionReference_ToProto(mapCtx *direct.MapContext, in *krm.UserActionReference) *pb.UserActionReference {
	if in == nil {
		return nil
	}
	out := &pb.UserActionReference{}
	if oneof := UserActionReference_Operation_ToProto(mapCtx, in.Operation); oneof != nil {
		out.Reference = oneof
	}
	if oneof := UserActionReference_DataLabelingJob_ToProto(mapCtx, in.DataLabelingJob); oneof != nil {
		out.Reference = oneof
	}
	out.Method = direct.ValueOf(in.Method)
	return out
}
