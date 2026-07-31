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

package datalabelingjob

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/genproto/googleapis/type/money"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func VertexAIDataLabelingJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataLabelingJob) *krm.VertexAIDataLabelingJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIDataLabelingJobObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.LabelingProgress = direct.LazyPtr(in.GetLabelingProgress())
	out.CurrentSpend = Money_FromProto(mapCtx, in.GetCurrentSpend())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	return out
}

func VertexAIDataLabelingJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIDataLabelingJobObservedState) *pb.DataLabelingJob {
	if in == nil {
		return nil
	}
	out := &pb.DataLabelingJob{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.JobState](mapCtx, in.State)
	out.LabelingProgress = direct.ValueOf(in.LabelingProgress)
	out.CurrentSpend = Money_ToProto(mapCtx, in.CurrentSpend)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Error = Status_ToProto(mapCtx, in.Error)
	return out
}

func VertexAIDataLabelingJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataLabelingJob) *krm.VertexAIDataLabelingJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIDataLabelingJobSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.DatasetRefs = VertexAIDataLabelingJobSpec_DatasetRefs_FromProto(mapCtx, in.GetDatasets())
	out.AnnotationLabels = in.GetAnnotationLabels()
	out.LabelerCount = direct.LazyPtr(in.GetLabelerCount())
	out.InstructionURI = direct.LazyPtr(in.GetInstructionUri())
	out.InputsSchemaURI = direct.LazyPtr(in.GetInputsSchemaUri())
	out.Inputs = JSON_FromProto(mapCtx, in.GetInputs())
	out.Labels = in.GetLabels()
	out.SpecialistPools = in.GetSpecialistPools()
	out.EncryptionSpec = EncryptionSpecV1_FromProto(mapCtx, in.GetEncryptionSpec())
	out.ActiveLearningConfig = ActiveLearningConfig_FromProto(mapCtx, in.GetActiveLearningConfig())
	return out
}

func VertexAIDataLabelingJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIDataLabelingJobSpec) *pb.DataLabelingJob {
	if in == nil {
		return nil
	}
	out := &pb.DataLabelingJob{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Datasets = VertexAIDataLabelingJobSpec_DatasetRefs_ToProto(mapCtx, in.DatasetRefs)
	out.AnnotationLabels = in.AnnotationLabels
	out.LabelerCount = direct.ValueOf(in.LabelerCount)
	out.InstructionUri = direct.ValueOf(in.InstructionURI)
	out.InputsSchemaUri = direct.ValueOf(in.InputsSchemaURI)
	out.Inputs = JSON_ToProto(mapCtx, in.Inputs)
	out.Labels = in.Labels
	out.SpecialistPools = in.SpecialistPools
	out.EncryptionSpec = EncryptionSpecV1_ToProto(mapCtx, in.EncryptionSpec)
	out.ActiveLearningConfig = ActiveLearningConfig_ToProto(mapCtx, in.ActiveLearningConfig)
	return out
}

func VertexAIDataLabelingJobSpec_DatasetRefs_FromProto(mapCtx *direct.MapContext, in []string) []v1beta1.VertexAIDatasetRef {
	if in == nil {
		return nil
	}
	out := make([]v1beta1.VertexAIDatasetRef, len(in))
	for i, v := range in {
		out[i] = v1beta1.VertexAIDatasetRef{External: v}
	}
	return out
}

func VertexAIDataLabelingJobSpec_DatasetRefs_ToProto(mapCtx *direct.MapContext, in []v1beta1.VertexAIDatasetRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = v.External
	}
	return out
}

func Money_FromProto(mapCtx *direct.MapContext, in *money.Money) *krm.Money {
	if in == nil {
		return nil
	}
	out := &krm.Money{}
	out.CurrencyCode = direct.LazyPtr(in.GetCurrencyCode())
	out.Units = direct.LazyPtr(in.GetUnits())
	out.Nanos = direct.LazyPtr(in.GetNanos())
	return out
}

func Money_ToProto(mapCtx *direct.MapContext, in *krm.Money) *money.Money {
	if in == nil {
		return nil
	}
	out := &money.Money{}
	out.CurrencyCode = direct.ValueOf(in.CurrencyCode)
	out.Units = direct.ValueOf(in.Units)
	out.Nanos = direct.ValueOf(in.Nanos)
	return out
}

func Status_FromProto(mapCtx *direct.MapContext, in *status.Status) *common.Status {
	if in == nil {
		return nil
	}
	out := &common.Status{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}

func Status_ToProto(mapCtx *direct.MapContext, in *common.Status) *status.Status {
	if in == nil {
		return nil
	}
	out := &status.Status{}
	out.Code = direct.ValueOf(in.Code)
	out.Message = direct.ValueOf(in.Message)
	return out
}

func JSON_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *apiextensionsv1.JSON {
	if in == nil {
		return nil
	}
	b, err := protojson.Marshal(in)
	if err != nil {
		return nil
	}
	return &apiextensionsv1.JSON{Raw: b}
}

func JSON_ToProto(mapCtx *direct.MapContext, in *apiextensionsv1.JSON) *structpb.Value {
	if in == nil || in.Raw == nil {
		return nil
	}
	out := &structpb.Value{}
	if err := protojson.Unmarshal(in.Raw, out); err != nil {
		return nil
	}
	return out
}

func EncryptionSpecV1_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionSpec{}
	if in.GetKmsKeyName() != "" {
		out.KMSKeyRef = &kmsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	return out
}

func EncryptionSpecV1_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	if in.KMSKeyRef != nil {
		out.KmsKeyName = in.KMSKeyRef.External
	}
	return out
}
func ActiveLearningConfig_FromProto(mapCtx *direct.MapContext, in *pb.ActiveLearningConfig) *krm.ActiveLearningConfig {
	if in == nil {
		return nil
	}
	out := &krm.ActiveLearningConfig{}
	out.MaxDataItemCount = direct.LazyPtr(in.GetMaxDataItemCount())
	out.MaxDataItemPercentage = direct.LazyPtr(in.GetMaxDataItemPercentage())
	out.SampleConfig = SampleConfig_FromProto(mapCtx, in.GetSampleConfig())
	out.TrainingConfig = TrainingConfig_FromProto(mapCtx, in.GetTrainingConfig())
	return out
}
func ActiveLearningConfig_ToProto(mapCtx *direct.MapContext, in *krm.ActiveLearningConfig) *pb.ActiveLearningConfig {
	if in == nil {
		return nil
	}
	out := &pb.ActiveLearningConfig{}
	if oneof := ActiveLearningConfig_MaxDataItemCount_ToProto(mapCtx, in.MaxDataItemCount); oneof != nil {
		out.HumanLabelingBudget = oneof
	}
	if oneof := ActiveLearningConfig_MaxDataItemPercentage_ToProto(mapCtx, in.MaxDataItemPercentage); oneof != nil {
		out.HumanLabelingBudget = oneof
	}
	out.SampleConfig = SampleConfig_ToProto(mapCtx, in.SampleConfig)
	out.TrainingConfig = TrainingConfig_ToProto(mapCtx, in.TrainingConfig)
	return out
}
func ActiveLearningConfig_MaxDataItemCount_ToProto(mapCtx *direct.MapContext, in *int64) *pb.ActiveLearningConfig_MaxDataItemCount {
	if in == nil {
		return nil
	}
	return &pb.ActiveLearningConfig_MaxDataItemCount{MaxDataItemCount: *in}
}
func ActiveLearningConfig_MaxDataItemPercentage_ToProto(mapCtx *direct.MapContext, in *int32) *pb.ActiveLearningConfig_MaxDataItemPercentage {
	if in == nil {
		return nil
	}
	return &pb.ActiveLearningConfig_MaxDataItemPercentage{MaxDataItemPercentage: *in}
}
func SampleConfig_FromProto(mapCtx *direct.MapContext, in *pb.SampleConfig) *krm.SampleConfig {
	if in == nil {
		return nil
	}
	out := &krm.SampleConfig{}
	out.InitialBatchSamplePercentage = direct.LazyPtr(in.GetInitialBatchSamplePercentage())
	out.FollowingBatchSamplePercentage = direct.LazyPtr(in.GetFollowingBatchSamplePercentage())
	out.SampleStrategy = direct.Enum_FromProto(mapCtx, in.GetSampleStrategy())
	return out
}
func SampleConfig_ToProto(mapCtx *direct.MapContext, in *krm.SampleConfig) *pb.SampleConfig {
	if in == nil {
		return nil
	}
	out := &pb.SampleConfig{}
	if oneof := SampleConfig_InitialBatchSamplePercentage_ToProto(mapCtx, in.InitialBatchSamplePercentage); oneof != nil {
		out.InitialBatchSampleSize = oneof
	}
	if oneof := SampleConfig_FollowingBatchSamplePercentage_ToProto(mapCtx, in.FollowingBatchSamplePercentage); oneof != nil {
		out.FollowingBatchSampleSize = oneof
	}
	out.SampleStrategy = direct.Enum_ToProto[pb.SampleConfig_SampleStrategy](mapCtx, in.SampleStrategy)
	return out
}
func SampleConfig_InitialBatchSamplePercentage_ToProto(mapCtx *direct.MapContext, in *int32) *pb.SampleConfig_InitialBatchSamplePercentage {
	if in == nil {
		return nil
	}
	return &pb.SampleConfig_InitialBatchSamplePercentage{InitialBatchSamplePercentage: *in}
}
func SampleConfig_FollowingBatchSamplePercentage_ToProto(mapCtx *direct.MapContext, in *int32) *pb.SampleConfig_FollowingBatchSamplePercentage {
	if in == nil {
		return nil
	}
	return &pb.SampleConfig_FollowingBatchSamplePercentage{FollowingBatchSamplePercentage: *in}
}
func TrainingConfig_FromProto(mapCtx *direct.MapContext, in *pb.TrainingConfig) *krm.TrainingConfig {
	if in == nil {
		return nil
	}
	out := &krm.TrainingConfig{}
	out.TimeoutTrainingMilliHours = direct.LazyPtr(in.GetTimeoutTrainingMilliHours())
	return out
}
func TrainingConfig_ToProto(mapCtx *direct.MapContext, in *krm.TrainingConfig) *pb.TrainingConfig {
	if in == nil {
		return nil
	}
	out := &pb.TrainingConfig{}
	out.TimeoutTrainingMilliHours = direct.ValueOf(in.TimeoutTrainingMilliHours)
	return out
}
