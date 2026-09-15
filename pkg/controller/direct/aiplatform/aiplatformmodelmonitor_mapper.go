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
	aiplatformpb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/encoding/protojson"
)

func ExplanationSpecBetaToV1(in *aiplatformpb.ExplanationSpec) (*pb.ExplanationSpec, error) {
	if in == nil {
		return nil, nil
	}
	b, err := protojson.Marshal(in)
	if err != nil {
		return nil, err
	}
	out := &pb.ExplanationSpec{}
	if err := protojson.Unmarshal(b, out); err != nil {
		return nil, err
	}
	return out, nil
}

func ExplanationSpecV1ToBeta(in *pb.ExplanationSpec) (*aiplatformpb.ExplanationSpec, error) {
	if in == nil {
		return nil, nil
	}
	b, err := protojson.Marshal(in)
	if err != nil {
		return nil, err
	}
	out := &aiplatformpb.ExplanationSpec{}
	if err := protojson.Unmarshal(b, out); err != nil {
		return nil, err
	}
	return out, nil
}

func EncryptionSpecBetaToV1(in *aiplatformpb.EncryptionSpec) (*pb.EncryptionSpec, error) {
	if in == nil {
		return nil, nil
	}
	b, err := protojson.Marshal(in)
	if err != nil {
		return nil, err
	}
	out := &pb.EncryptionSpec{}
	if err := protojson.Unmarshal(b, out); err != nil {
		return nil, err
	}
	return out, nil
}

func EncryptionSpecV1ToBeta(in *pb.EncryptionSpec) (*aiplatformpb.EncryptionSpec, error) {
	if in == nil {
		return nil, nil
	}
	b, err := protojson.Marshal(in)
	if err != nil {
		return nil, err
	}
	out := &aiplatformpb.EncryptionSpec{}
	if err := protojson.Unmarshal(b, out); err != nil {
		return nil, err
	}
	return out, nil
}

func ModelMonitor_ModelMonitoringTarget_VertexModelSource_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.ModelMonitor_ModelMonitoringTarget_VertexModelSource) *krm.ModelMonitor_ModelMonitoringTarget_VertexModelSource {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitor_ModelMonitoringTarget_VertexModelSource{}
	if in.GetModel() != "" {
		out.ModelRef = &krm.AIPlatformModelRef{External: in.GetModel()}
	}
	out.ModelVersionID = direct.LazyPtr(in.GetModelVersionId())
	return out
}

func ModelMonitor_ModelMonitoringTarget_VertexModelSource_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitor_ModelMonitoringTarget_VertexModelSource) *aiplatformpb.ModelMonitor_ModelMonitoringTarget_VertexModelSource {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.ModelMonitor_ModelMonitoringTarget_VertexModelSource{}
	if in.ModelRef != nil {
		out.Model = in.ModelRef.External
	}
	out.ModelVersionId = direct.ValueOf(in.ModelVersionID)
	return out
}

func AIPlatformModelMonitorObservedState_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.ModelMonitor) *krm.AIPlatformModelMonitorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AIPlatformModelMonitorObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}

func AIPlatformModelMonitorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AIPlatformModelMonitorObservedState) *aiplatformpb.ModelMonitor {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.ModelMonitor{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}

func AIPlatformModelMonitorSpec_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.ModelMonitor) *krm.AIPlatformModelMonitorSpec {
	if in == nil {
		return nil
	}
	out := &krm.AIPlatformModelMonitorSpec{}
	out.TabularObjective = ModelMonitoringObjectiveSpec_TabularObjective_FromProto(mapCtx, in.GetTabularObjective())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ModelMonitoringTarget = ModelMonitor_ModelMonitoringTarget_FromProto(mapCtx, in.GetModelMonitoringTarget())
	out.TrainingDataset = ModelMonitoringInput_FromProto(mapCtx, in.GetTrainingDataset())
	out.NotificationSpec = ModelMonitoringNotificationSpec_FromProto(mapCtx, in.GetNotificationSpec())
	out.OutputSpec = ModelMonitoringOutputSpec_FromProto(mapCtx, in.GetOutputSpec())

	if in.GetExplanationSpec() != nil {
		v1Spec, err := ExplanationSpecBetaToV1(in.GetExplanationSpec())
		if err != nil {
			mapCtx.Errorf("failed to convert explanation spec: %v", err)
		} else {
			out.ExplanationSpec = ExplanationSpec_FromProto(mapCtx, v1Spec)
		}
	}

	out.ModelMonitoringSchema = ModelMonitoringSchema_FromProto(mapCtx, in.GetModelMonitoringSchema())

	if in.GetEncryptionSpec() != nil {
		v1Spec, err := EncryptionSpecBetaToV1(in.GetEncryptionSpec())
		if err != nil {
			mapCtx.Errorf("failed to convert encryption spec: %v", err)
		} else {
			out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, v1Spec)
		}
	}

	return out
}

func AIPlatformModelMonitorSpec_ToProto(mapCtx *direct.MapContext, in *krm.AIPlatformModelMonitorSpec) *aiplatformpb.ModelMonitor {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.ModelMonitor{}
	if oneof := ModelMonitoringObjectiveSpec_TabularObjective_ToProto(mapCtx, in.TabularObjective); oneof != nil {
		out.DefaultObjective = &aiplatformpb.ModelMonitor_TabularObjective{TabularObjective: oneof}
	}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ModelMonitoringTarget = ModelMonitor_ModelMonitoringTarget_ToProto(mapCtx, in.ModelMonitoringTarget)
	out.TrainingDataset = ModelMonitoringInput_ToProto(mapCtx, in.TrainingDataset)
	out.NotificationSpec = ModelMonitoringNotificationSpec_ToProto(mapCtx, in.NotificationSpec)
	out.OutputSpec = ModelMonitoringOutputSpec_ToProto(mapCtx, in.OutputSpec)

	if in.ExplanationSpec != nil {
		v1Spec := ExplanationSpec_ToProto(mapCtx, in.ExplanationSpec)
		if v1Spec != nil {
			betaSpec, err := ExplanationSpecV1ToBeta(v1Spec)
			if err != nil {
				mapCtx.Errorf("failed to convert explanation spec to beta: %v", err)
			} else {
				out.ExplanationSpec = betaSpec
			}
		}
	}

	out.ModelMonitoringSchema = ModelMonitoringSchema_ToProto(mapCtx, in.ModelMonitoringSchema)

	if in.EncryptionSpec != nil {
		v1Spec := EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
		if v1Spec != nil {
			betaSpec, err := EncryptionSpecV1ToBeta(v1Spec)
			if err != nil {
				mapCtx.Errorf("failed to convert encryption spec to beta: %v", err)
			} else {
				out.EncryptionSpec = betaSpec
			}
		}
	}

	return out
}
