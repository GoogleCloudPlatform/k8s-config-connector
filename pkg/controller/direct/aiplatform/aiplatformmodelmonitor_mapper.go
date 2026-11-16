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
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	monitoringv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	vertexaiv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/genproto/googleapis/type/interval"
	"google.golang.org/protobuf/proto"
)

func Interval_FromProto(mapCtx *direct.MapContext, in *interval.Interval) *krm.Interval {
	if in == nil {
		return nil
	}
	out := &krm.Interval{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}

func Interval_ToProto(mapCtx *direct.MapContext, in *krm.Interval) *interval.Interval {
	if in == nil {
		return nil
	}
	out := &interval.Interval{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}

func ExplanationSpecWithRequired_ToProto(mapCtx *direct.MapContext, in *krm.ExplanationSpecWithRequired) *aiplatformpb.ExplanationSpec {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.ExplanationSpec{}
	if in.Parameters != nil {
		v1Params := ExplanationParameters_ToProto(mapCtx, in.Parameters)
		if v1Params != nil {
			bytes, err := proto.Marshal(v1Params)
			if err != nil {
				mapCtx.Errorf("marshaling explanation parameters: %w", err)
			} else {
				v1beta1Params := &aiplatformpb.ExplanationParameters{}
				if err := proto.Unmarshal(bytes, v1beta1Params); err != nil {
					mapCtx.Errorf("unmarshaling explanation parameters: %w", err)
				} else {
					out.Parameters = v1beta1Params
				}
			}
		}
	}
	if in.Metadata != nil {
		v1Meta := ExplanationMetadata_ToProto(mapCtx, in.Metadata)
		if v1Meta != nil {
			bytes, err := proto.Marshal(v1Meta)
			if err != nil {
				mapCtx.Errorf("marshaling explanation metadata: %w", err)
			} else {
				v1beta1Meta := &aiplatformpb.ExplanationMetadata{}
				if err := proto.Unmarshal(bytes, v1beta1Meta); err != nil {
					mapCtx.Errorf("unmarshaling explanation metadata: %w", err)
				} else {
					out.Metadata = v1beta1Meta
				}
			}
		}
	}
	return out
}

func ExplanationSpecWithRequired_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.ExplanationSpec) *krm.ExplanationSpecWithRequired {
	if in == nil {
		return nil
	}
	out := &krm.ExplanationSpecWithRequired{}
	if in.GetParameters() != nil {
		bytes, err := proto.Marshal(in.GetParameters())
		if err != nil {
			mapCtx.Errorf("marshaling explanation parameters: %w", err)
		} else {
			v1Params := &pb.ExplanationParameters{}
			if err := proto.Unmarshal(bytes, v1Params); err != nil {
				mapCtx.Errorf("unmarshaling explanation parameters: %w", err)
			} else {
				out.Parameters = ExplanationParameters_FromProto(mapCtx, v1Params)
			}
		}
	}
	if in.GetMetadata() != nil {
		bytes, err := proto.Marshal(in.GetMetadata())
		if err != nil {
			mapCtx.Errorf("marshaling explanation metadata: %w", err)
		} else {
			v1Meta := &pb.ExplanationMetadata{}
			if err := proto.Unmarshal(bytes, v1Meta); err != nil {
				mapCtx.Errorf("unmarshaling explanation metadata: %w", err)
			} else {
				out.Metadata = ExplanationMetadata_FromProto(mapCtx, v1Meta)
			}
		}
	}
	return out
}

func BatchDedicatedResources_ToProto(mapCtx *direct.MapContext, in *krm.BatchDedicatedResources) *aiplatformpb.BatchDedicatedResources {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.BatchDedicatedResources{}
	if in.MachineSpec != nil {
		v1MachineSpec := MachineSpec_ToProto(mapCtx, in.MachineSpec)
		if v1MachineSpec != nil {
			bytes, err := proto.Marshal(v1MachineSpec)
			if err != nil {
				mapCtx.Errorf("marshaling machineSpec: %w", err)
			} else {
				v1beta1MachineSpec := &aiplatformpb.MachineSpec{}
				if err := proto.Unmarshal(bytes, v1beta1MachineSpec); err != nil {
					mapCtx.Errorf("unmarshaling machineSpec: %w", err)
				} else {
					out.MachineSpec = v1beta1MachineSpec
				}
			}
		}
	}
	out.StartingReplicaCount = direct.ValueOf(in.StartingReplicaCount)
	out.MaxReplicaCount = direct.ValueOf(in.MaxReplicaCount)
	out.FlexStart = FlexStart_ToProto(mapCtx, in.FlexStart)
	out.Spot = direct.ValueOf(in.Spot)
	return out
}

func BatchDedicatedResources_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.BatchDedicatedResources) *krm.BatchDedicatedResources {
	if in == nil {
		return nil
	}
	out := &krm.BatchDedicatedResources{}
	if in.GetMachineSpec() != nil {
		bytes, err := proto.Marshal(in.GetMachineSpec())
		if err != nil {
			mapCtx.Errorf("marshaling machineSpec: %w", err)
		} else {
			v1MachineSpec := &pb.MachineSpec{}
			if err := proto.Unmarshal(bytes, v1MachineSpec); err != nil {
				mapCtx.Errorf("unmarshaling machineSpec: %w", err)
			} else {
				out.MachineSpec = MachineSpec_FromProto(mapCtx, v1MachineSpec)
			}
		}
	}
	out.StartingReplicaCount = direct.LazyPtr(in.GetStartingReplicaCount())
	out.MaxReplicaCount = direct.LazyPtr(in.GetMaxReplicaCount())
	out.FlexStart = FlexStart_FromProto(mapCtx, in.GetFlexStart())
	out.Spot = direct.LazyPtr(in.GetSpot())
	return out
}

func ModelMonitoringObjectiveSpec_DataDriftSpec_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringObjectiveSpec_DataDriftSpec) *aiplatformpb.ModelMonitoringObjectiveSpec_DataDriftSpec {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.ModelMonitoringObjectiveSpec_DataDriftSpec{}
	out.Features = in.Features
	out.CategoricalMetricType = direct.ValueOf(in.CategoricalMetricType)
	out.NumericMetricType = direct.ValueOf(in.NumericMetricType)
	out.DefaultCategoricalAlertCondition = ModelMonitoringAlertCondition_ToProto(mapCtx, in.DefaultCategoricalAlertCondition)
	out.DefaultNumericAlertCondition = ModelMonitoringAlertCondition_ToProto(mapCtx, in.DefaultNumericAlertCondition)
	if in.FeatureAlertConditions != nil {
		out.FeatureAlertConditions = make(map[string]*aiplatformpb.ModelMonitoringAlertCondition, len(in.FeatureAlertConditions))
		for k, v := range in.FeatureAlertConditions {
			out.FeatureAlertConditions[k] = ModelMonitoringAlertCondition_ToProto(mapCtx, &v)
		}
	}
	return out
}

func ModelMonitoringObjectiveSpec_FeatureAttributionSpec_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringObjectiveSpec_FeatureAttributionSpec) *aiplatformpb.ModelMonitoringObjectiveSpec_FeatureAttributionSpec {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.ModelMonitoringObjectiveSpec_FeatureAttributionSpec{}
	out.Features = in.Features
	out.DefaultAlertCondition = ModelMonitoringAlertCondition_ToProto(mapCtx, in.DefaultAlertCondition)
	if in.FeatureAlertConditions != nil {
		out.FeatureAlertConditions = make(map[string]*aiplatformpb.ModelMonitoringAlertCondition, len(in.FeatureAlertConditions))
		for k, v := range in.FeatureAlertConditions {
			out.FeatureAlertConditions[k] = ModelMonitoringAlertCondition_ToProto(mapCtx, &v)
		}
	}
	out.BatchExplanationDedicatedResources = BatchDedicatedResources_ToProto(mapCtx, in.BatchExplanationDedicatedResources)
	return out
}

func ModelMonitoringOutputSpec_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.ModelMonitoringOutputSpec) *krm.ModelMonitoringOutputSpec {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitoringOutputSpec{}
	if in.GetGcsBaseDirectory() != nil {
		out.GCSBaseDirectory = &krm.GCSDestinationWithRequired{
			OutputURIPrefix: direct.LazyPtr(in.GetGcsBaseDirectory().GetOutputUriPrefix()),
		}
	}
	return out
}

func ModelMonitoringOutputSpec_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringOutputSpec) *aiplatformpb.ModelMonitoringOutputSpec {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.ModelMonitoringOutputSpec{}
	if in.GCSBaseDirectory != nil {
		out.GcsBaseDirectory = &aiplatformpb.GcsDestination{
			OutputUriPrefix: direct.ValueOf(in.GCSBaseDirectory.OutputURIPrefix),
		}
	}
	return out
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

func ModelMonitoringNotificationSpec_NotificationChannelConfig_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.ModelMonitoringNotificationSpec_NotificationChannelConfig) *krm.ModelMonitoringNotificationSpec_NotificationChannelConfig {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitoringNotificationSpec_NotificationChannelConfig{}
	if in.GetNotificationChannel() != "" {
		out.NotificationChannelRef = &monitoringv1beta1.MonitoringNotificationChannelRef{External: in.GetNotificationChannel()}
	}
	return out
}

func ModelMonitoringNotificationSpec_NotificationChannelConfig_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringNotificationSpec_NotificationChannelConfig) *aiplatformpb.ModelMonitoringNotificationSpec_NotificationChannelConfig {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.ModelMonitoringNotificationSpec_NotificationChannelConfig{}
	if in.NotificationChannelRef != nil {
		out.NotificationChannel = in.NotificationChannelRef.External
	}
	return out
}

func ModelMonitoringInput_ModelMonitoringDataset_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.ModelMonitoringInput_ModelMonitoringDataset) *krm.ModelMonitoringInput_ModelMonitoringDataset {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitoringInput_ModelMonitoringDataset{}
	if in.GetVertexDataset() != "" {
		out.VertexDatasetRef = &vertexaiv1beta1.VertexAIDatasetRef{External: in.GetVertexDataset()}
	}
	out.GCSSource = ModelMonitoringInput_ModelMonitoringDataset_ModelMonitoringGCSSource_FromProto(mapCtx, in.GetGcsSource())
	out.BigquerySource = ModelMonitoringInput_ModelMonitoringDataset_ModelMonitoringBigQuerySource_FromProto(mapCtx, in.GetBigquerySource())
	out.TimestampField = direct.LazyPtr(in.GetTimestampField())
	return out
}

func ModelMonitoringInput_ModelMonitoringDataset_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringInput_ModelMonitoringDataset) *aiplatformpb.ModelMonitoringInput_ModelMonitoringDataset {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.ModelMonitoringInput_ModelMonitoringDataset{}
	if in.VertexDatasetRef != nil {
		out.DataLocation = &aiplatformpb.ModelMonitoringInput_ModelMonitoringDataset_VertexDataset{VertexDataset: in.VertexDatasetRef.External}
	}
	if in.GCSSource != nil {
		if oneof := ModelMonitoringInput_ModelMonitoringDataset_ModelMonitoringGCSSource_ToProto(mapCtx, in.GCSSource); oneof != nil {
			out.DataLocation = &aiplatformpb.ModelMonitoringInput_ModelMonitoringDataset_GcsSource{GcsSource: oneof}
		}
	}
	if in.BigquerySource != nil {
		if oneof := ModelMonitoringInput_ModelMonitoringDataset_ModelMonitoringBigQuerySource_ToProto(mapCtx, in.BigquerySource); oneof != nil {
			out.DataLocation = &aiplatformpb.ModelMonitoringInput_ModelMonitoringDataset_BigquerySource{BigquerySource: oneof}
		}
	}
	out.TimestampField = direct.ValueOf(in.TimestampField)
	return out
}

func ModelMonitoringInput_VertexEndpointLogs_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.ModelMonitoringInput_VertexEndpointLogs) *krm.ModelMonitoringInput_VertexEndpointLogs {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitoringInput_VertexEndpointLogs{}
	for _, e := range in.GetEndpoints() {
		out.EndpointRefs = append(out.EndpointRefs, vertexaiv1beta1.VertexAIEndpointRef{External: e})
	}
	return out
}

func ModelMonitoringInput_VertexEndpointLogs_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringInput_VertexEndpointLogs) *aiplatformpb.ModelMonitoringInput_VertexEndpointLogs {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.ModelMonitoringInput_VertexEndpointLogs{}
	for _, ref := range in.EndpointRefs {
		out.Endpoints = append(out.Endpoints, ref.External)
	}
	return out
}

func ModelMonitoringInput_BatchPredictionOutput_FromProto(mapCtx *direct.MapContext, in *aiplatformpb.ModelMonitoringInput_BatchPredictionOutput) *krm.ModelMonitoringInput_BatchPredictionOutput {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitoringInput_BatchPredictionOutput{}
	if in.GetBatchPredictionJob() != "" {
		out.BatchPredictionJobRef = &krm.AIPlatformBatchPredictionJobRef{External: in.GetBatchPredictionJob()}
	}
	return out
}

func ModelMonitoringInput_BatchPredictionOutput_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringInput_BatchPredictionOutput) *aiplatformpb.ModelMonitoringInput_BatchPredictionOutput {
	if in == nil {
		return nil
	}
	out := &aiplatformpb.ModelMonitoringInput_BatchPredictionOutput{}
	if in.BatchPredictionJobRef != nil {
		out.BatchPredictionJob = in.BatchPredictionJobRef.External
	}
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
	out.ExplanationSpec = ExplanationSpecWithRequired_FromProto(mapCtx, in.GetExplanationSpec())
	out.ModelMonitoringSchema = ModelMonitoringSchema_FromProto(mapCtx, in.GetModelMonitoringSchema())
	if in.GetEncryptionSpec() != nil {
		out.EncryptionSpec = &krm.EncryptionSpec{}
		if in.GetEncryptionSpec().GetKmsKeyName() != "" {
			out.EncryptionSpec.KmsKeyRef = &kmsv1beta1.KMSCryptoKeyRef{External: in.GetEncryptionSpec().GetKmsKeyName()}
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
	out.ExplanationSpec = ExplanationSpecWithRequired_ToProto(mapCtx, in.ExplanationSpec)
	out.ModelMonitoringSchema = ModelMonitoringSchema_ToProto(mapCtx, in.ModelMonitoringSchema)
	if in.EncryptionSpec != nil && in.EncryptionSpec.KmsKeyRef != nil {
		out.EncryptionSpec = &aiplatformpb.EncryptionSpec{
			KmsKeyName: in.EncryptionSpec.KmsKeyRef.External,
		}
	}
	return out
}

func FeatureAlertConditions_FromProto(mapCtx *direct.MapContext, in map[string]*aiplatformpb.ModelMonitoringAlertCondition) map[string]krm.ModelMonitoringAlertCondition {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.ModelMonitoringAlertCondition, len(in))
	for k, v := range in {
		if val := ModelMonitoringAlertCondition_FromProto(mapCtx, v); val != nil {
			out[k] = *val
		}
	}
	return out
}

func FeatureAlertConditions_ToProto(mapCtx *direct.MapContext, in map[string]krm.ModelMonitoringAlertCondition) map[string]*aiplatformpb.ModelMonitoringAlertCondition {
	if in == nil {
		return nil
	}
	out := make(map[string]*aiplatformpb.ModelMonitoringAlertCondition, len(in))
	for k, v := range in {
		out[k] = ModelMonitoringAlertCondition_ToProto(mapCtx, &v)
	}
	return out
}
