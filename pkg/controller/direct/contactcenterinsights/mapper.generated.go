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

package contactcenterinsights

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ContactcenterinsightsFeedbackLabelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FeedbackLabel) *krm.ContactcenterinsightsFeedbackLabelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsFeedbackLabelObservedState{}
	// MISSING: Label
	// MISSING: QaAnswerLabel
	// MISSING: Name
	// MISSING: LabeledResource
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ContactcenterinsightsFeedbackLabelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsFeedbackLabelObservedState) *pb.FeedbackLabel {
	if in == nil {
		return nil
	}
	out := &pb.FeedbackLabel{}
	// MISSING: Label
	// MISSING: QaAnswerLabel
	// MISSING: Name
	// MISSING: LabeledResource
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ContactcenterinsightsFeedbackLabelSpec_FromProto(mapCtx *direct.MapContext, in *pb.FeedbackLabel) *krm.ContactcenterinsightsFeedbackLabelSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsFeedbackLabelSpec{}
	// MISSING: Label
	// MISSING: QaAnswerLabel
	// MISSING: Name
	// MISSING: LabeledResource
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ContactcenterinsightsFeedbackLabelSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsFeedbackLabelSpec) *pb.FeedbackLabel {
	if in == nil {
		return nil
	}
	out := &pb.FeedbackLabel{}
	// MISSING: Label
	// MISSING: QaAnswerLabel
	// MISSING: Name
	// MISSING: LabeledResource
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func FeedbackLabel_FromProto(mapCtx *direct.MapContext, in *pb.FeedbackLabel) *krm.FeedbackLabel {
	if in == nil {
		return nil
	}
	out := &krm.FeedbackLabel{}
	out.Label = direct.LazyPtr(in.GetLabel())
	out.QaAnswerLabel = QaAnswer_AnswerValue_FromProto(mapCtx, in.GetQaAnswerLabel())
	out.Name = direct.LazyPtr(in.GetName())
	out.LabeledResource = direct.LazyPtr(in.GetLabeledResource())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func FeedbackLabel_ToProto(mapCtx *direct.MapContext, in *krm.FeedbackLabel) *pb.FeedbackLabel {
	if in == nil {
		return nil
	}
	out := &pb.FeedbackLabel{}
	if oneof := FeedbackLabel_Label_ToProto(mapCtx, in.Label); oneof != nil {
		out.LabelType = oneof
	}
	if oneof := QaAnswer_AnswerValue_ToProto(mapCtx, in.QaAnswerLabel); oneof != nil {
		out.LabelType = &pb.FeedbackLabel_QaAnswerLabel{QaAnswerLabel: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.LabeledResource = direct.ValueOf(in.LabeledResource)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func FeedbackLabelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FeedbackLabel) *krm.FeedbackLabelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FeedbackLabelObservedState{}
	// MISSING: Label
	out.QaAnswerLabel = QaAnswer_AnswerValueObservedState_FromProto(mapCtx, in.GetQaAnswerLabel())
	// MISSING: Name
	// MISSING: LabeledResource
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func FeedbackLabelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FeedbackLabelObservedState) *pb.FeedbackLabel {
	if in == nil {
		return nil
	}
	out := &pb.FeedbackLabel{}
	// MISSING: Label
	if oneof := QaAnswer_AnswerValueObservedState_ToProto(mapCtx, in.QaAnswerLabel); oneof != nil {
		out.LabelType = &pb.FeedbackLabel_QaAnswerLabel{QaAnswerLabel: oneof}
	}
	// MISSING: Name
	// MISSING: LabeledResource
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
