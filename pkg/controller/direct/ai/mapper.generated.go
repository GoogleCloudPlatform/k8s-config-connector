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

package ai

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/ai/generativelanguage/apiv1beta/generativelanguagepb"
)
func AiChunkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Chunk) *krm.AiChunkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiChunkObservedState{}
	// MISSING: Name
	// MISSING: Data
	// MISSING: CustomMetadata
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	return out
}
func AiChunkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiChunkObservedState) *pb.Chunk {
	if in == nil {
		return nil
	}
	out := &pb.Chunk{}
	// MISSING: Name
	// MISSING: Data
	// MISSING: CustomMetadata
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	return out
}
func AiChunkSpec_FromProto(mapCtx *direct.MapContext, in *pb.Chunk) *krm.AiChunkSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiChunkSpec{}
	// MISSING: Name
	// MISSING: Data
	// MISSING: CustomMetadata
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	return out
}
func AiChunkSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiChunkSpec) *pb.Chunk {
	if in == nil {
		return nil
	}
	out := &pb.Chunk{}
	// MISSING: Name
	// MISSING: Data
	// MISSING: CustomMetadata
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	return out
}
func AiCorpusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Corpus) *krm.AiCorpusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiCorpusObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AiCorpusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiCorpusObservedState) *pb.Corpus {
	if in == nil {
		return nil
	}
	out := &pb.Corpus{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AiCorpusSpec_FromProto(mapCtx *direct.MapContext, in *pb.Corpus) *krm.AiCorpusSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiCorpusSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AiCorpusSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiCorpusSpec) *pb.Corpus {
	if in == nil {
		return nil
	}
	out := &pb.Corpus{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AiDocumentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.AiDocumentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiDocumentObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomMetadata
	// MISSING: UpdateTime
	// MISSING: CreateTime
	return out
}
func AiDocumentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiDocumentObservedState) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomMetadata
	// MISSING: UpdateTime
	// MISSING: CreateTime
	return out
}
func AiDocumentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.AiDocumentSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiDocumentSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomMetadata
	// MISSING: UpdateTime
	// MISSING: CreateTime
	return out
}
func AiDocumentSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiDocumentSpec) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomMetadata
	// MISSING: UpdateTime
	// MISSING: CreateTime
	return out
}
func AiPermissionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Permission) *krm.AiPermissionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiPermissionObservedState{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func AiPermissionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiPermissionObservedState) *pb.Permission {
	if in == nil {
		return nil
	}
	out := &pb.Permission{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func AiPermissionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Permission) *krm.AiPermissionSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiPermissionSpec{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func AiPermissionSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiPermissionSpec) *pb.Permission {
	if in == nil {
		return nil
	}
	out := &pb.Permission{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func AiTunedModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TunedModel) *krm.AiTunedModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiTunedModelObservedState{}
	// MISSING: TunedModelSource
	// MISSING: BaseModel
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Temperature
	// MISSING: TopP
	// MISSING: TopK
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: TuningTask
	// MISSING: ReaderProjectNumbers
	return out
}
func AiTunedModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiTunedModelObservedState) *pb.TunedModel {
	if in == nil {
		return nil
	}
	out := &pb.TunedModel{}
	// MISSING: TunedModelSource
	// MISSING: BaseModel
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Temperature
	// MISSING: TopP
	// MISSING: TopK
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: TuningTask
	// MISSING: ReaderProjectNumbers
	return out
}
func AiTunedModelSpec_FromProto(mapCtx *direct.MapContext, in *pb.TunedModel) *krm.AiTunedModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiTunedModelSpec{}
	// MISSING: TunedModelSource
	// MISSING: BaseModel
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Temperature
	// MISSING: TopP
	// MISSING: TopK
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: TuningTask
	// MISSING: ReaderProjectNumbers
	return out
}
func AiTunedModelSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiTunedModelSpec) *pb.TunedModel {
	if in == nil {
		return nil
	}
	out := &pb.TunedModel{}
	// MISSING: TunedModelSource
	// MISSING: BaseModel
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Temperature
	// MISSING: TopP
	// MISSING: TopK
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: TuningTask
	// MISSING: ReaderProjectNumbers
	return out
}
func Dataset_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.Dataset {
	if in == nil {
		return nil
	}
	out := &krm.Dataset{}
	out.Examples = TuningExamples_FromProto(mapCtx, in.GetExamples())
	return out
}
func Dataset_ToProto(mapCtx *direct.MapContext, in *krm.Dataset) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	if oneof := TuningExamples_ToProto(mapCtx, in.Examples); oneof != nil {
		out.Dataset = &pb.Dataset_Examples{Examples: oneof}
	}
	return out
}
func Hyperparameters_FromProto(mapCtx *direct.MapContext, in *pb.Hyperparameters) *krm.Hyperparameters {
	if in == nil {
		return nil
	}
	out := &krm.Hyperparameters{}
	out.LearningRate = direct.LazyPtr(in.GetLearningRate())
	out.LearningRateMultiplier = direct.LazyPtr(in.GetLearningRateMultiplier())
	out.EpochCount = in.EpochCount
	out.BatchSize = in.BatchSize
	return out
}
func Hyperparameters_ToProto(mapCtx *direct.MapContext, in *krm.Hyperparameters) *pb.Hyperparameters {
	if in == nil {
		return nil
	}
	out := &pb.Hyperparameters{}
	if oneof := Hyperparameters_LearningRate_ToProto(mapCtx, in.LearningRate); oneof != nil {
		out.LearningRateOption = oneof
	}
	if oneof := Hyperparameters_LearningRateMultiplier_ToProto(mapCtx, in.LearningRateMultiplier); oneof != nil {
		out.LearningRateOption = oneof
	}
	out.EpochCount = in.EpochCount
	out.BatchSize = in.BatchSize
	return out
}
func TunedModel_FromProto(mapCtx *direct.MapContext, in *pb.TunedModel) *krm.TunedModel {
	if in == nil {
		return nil
	}
	out := &krm.TunedModel{}
	out.TunedModelSource = TunedModelSource_FromProto(mapCtx, in.GetTunedModelSource())
	out.BaseModel = direct.LazyPtr(in.GetBaseModel())
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Temperature = in.Temperature
	out.TopP = in.TopP
	out.TopK = in.TopK
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.TuningTask = TuningTask_FromProto(mapCtx, in.GetTuningTask())
	out.ReaderProjectNumbers = in.ReaderProjectNumbers
	return out
}
func TunedModel_ToProto(mapCtx *direct.MapContext, in *krm.TunedModel) *pb.TunedModel {
	if in == nil {
		return nil
	}
	out := &pb.TunedModel{}
	if oneof := TunedModelSource_ToProto(mapCtx, in.TunedModelSource); oneof != nil {
		out.SourceModel = &pb.TunedModel_TunedModelSource{TunedModelSource: oneof}
	}
	if oneof := TunedModel_BaseModel_ToProto(mapCtx, in.BaseModel); oneof != nil {
		out.SourceModel = oneof
	}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Temperature = in.Temperature
	out.TopP = in.TopP
	out.TopK = in.TopK
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.TuningTask = TuningTask_ToProto(mapCtx, in.TuningTask)
	out.ReaderProjectNumbers = in.ReaderProjectNumbers
	return out
}
func TunedModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TunedModel) *krm.TunedModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TunedModelObservedState{}
	out.TunedModelSource = TunedModelSourceObservedState_FromProto(mapCtx, in.GetTunedModelSource())
	// MISSING: BaseModel
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Temperature
	// MISSING: TopP
	// MISSING: TopK
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.TuningTask = TuningTaskObservedState_FromProto(mapCtx, in.GetTuningTask())
	// MISSING: ReaderProjectNumbers
	return out
}
func TunedModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TunedModelObservedState) *pb.TunedModel {
	if in == nil {
		return nil
	}
	out := &pb.TunedModel{}
	if oneof := TunedModelSourceObservedState_ToProto(mapCtx, in.TunedModelSource); oneof != nil {
		out.SourceModel = &pb.TunedModel_TunedModelSource{TunedModelSource: oneof}
	}
	// MISSING: BaseModel
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Temperature
	// MISSING: TopP
	// MISSING: TopK
	out.State = direct.Enum_ToProto[pb.TunedModel_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.TuningTask = TuningTaskObservedState_ToProto(mapCtx, in.TuningTask)
	// MISSING: ReaderProjectNumbers
	return out
}
func TunedModelSource_FromProto(mapCtx *direct.MapContext, in *pb.TunedModelSource) *krm.TunedModelSource {
	if in == nil {
		return nil
	}
	out := &krm.TunedModelSource{}
	out.TunedModel = direct.LazyPtr(in.GetTunedModel())
	// MISSING: BaseModel
	return out
}
func TunedModelSource_ToProto(mapCtx *direct.MapContext, in *krm.TunedModelSource) *pb.TunedModelSource {
	if in == nil {
		return nil
	}
	out := &pb.TunedModelSource{}
	out.TunedModel = direct.ValueOf(in.TunedModel)
	// MISSING: BaseModel
	return out
}
func TunedModelSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TunedModelSource) *krm.TunedModelSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TunedModelSourceObservedState{}
	// MISSING: TunedModel
	out.BaseModel = direct.LazyPtr(in.GetBaseModel())
	return out
}
func TunedModelSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TunedModelSourceObservedState) *pb.TunedModelSource {
	if in == nil {
		return nil
	}
	out := &pb.TunedModelSource{}
	// MISSING: TunedModel
	out.BaseModel = direct.ValueOf(in.BaseModel)
	return out
}
func TuningExample_FromProto(mapCtx *direct.MapContext, in *pb.TuningExample) *krm.TuningExample {
	if in == nil {
		return nil
	}
	out := &krm.TuningExample{}
	out.TextInput = direct.LazyPtr(in.GetTextInput())
	out.Output = direct.LazyPtr(in.GetOutput())
	return out
}
func TuningExample_ToProto(mapCtx *direct.MapContext, in *krm.TuningExample) *pb.TuningExample {
	if in == nil {
		return nil
	}
	out := &pb.TuningExample{}
	if oneof := TuningExample_TextInput_ToProto(mapCtx, in.TextInput); oneof != nil {
		out.ModelInput = oneof
	}
	out.Output = direct.ValueOf(in.Output)
	return out
}
func TuningExamples_FromProto(mapCtx *direct.MapContext, in *pb.TuningExamples) *krm.TuningExamples {
	if in == nil {
		return nil
	}
	out := &krm.TuningExamples{}
	out.Examples = direct.Slice_FromProto(mapCtx, in.Examples, TuningExample_FromProto)
	return out
}
func TuningExamples_ToProto(mapCtx *direct.MapContext, in *krm.TuningExamples) *pb.TuningExamples {
	if in == nil {
		return nil
	}
	out := &pb.TuningExamples{}
	out.Examples = direct.Slice_ToProto(mapCtx, in.Examples, TuningExample_ToProto)
	return out
}
func TuningSnapshot_FromProto(mapCtx *direct.MapContext, in *pb.TuningSnapshot) *krm.TuningSnapshot {
	if in == nil {
		return nil
	}
	out := &krm.TuningSnapshot{}
	// MISSING: Step
	// MISSING: Epoch
	// MISSING: MeanLoss
	// MISSING: ComputeTime
	return out
}
func TuningSnapshot_ToProto(mapCtx *direct.MapContext, in *krm.TuningSnapshot) *pb.TuningSnapshot {
	if in == nil {
		return nil
	}
	out := &pb.TuningSnapshot{}
	// MISSING: Step
	// MISSING: Epoch
	// MISSING: MeanLoss
	// MISSING: ComputeTime
	return out
}
func TuningSnapshotObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TuningSnapshot) *krm.TuningSnapshotObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TuningSnapshotObservedState{}
	out.Step = direct.LazyPtr(in.GetStep())
	out.Epoch = direct.LazyPtr(in.GetEpoch())
	out.MeanLoss = direct.LazyPtr(in.GetMeanLoss())
	out.ComputeTime = direct.StringTimestamp_FromProto(mapCtx, in.GetComputeTime())
	return out
}
func TuningSnapshotObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TuningSnapshotObservedState) *pb.TuningSnapshot {
	if in == nil {
		return nil
	}
	out := &pb.TuningSnapshot{}
	out.Step = direct.ValueOf(in.Step)
	out.Epoch = direct.ValueOf(in.Epoch)
	out.MeanLoss = direct.ValueOf(in.MeanLoss)
	out.ComputeTime = direct.StringTimestamp_ToProto(mapCtx, in.ComputeTime)
	return out
}
func TuningTask_FromProto(mapCtx *direct.MapContext, in *pb.TuningTask) *krm.TuningTask {
	if in == nil {
		return nil
	}
	out := &krm.TuningTask{}
	// MISSING: StartTime
	// MISSING: CompleteTime
	// MISSING: Snapshots
	out.TrainingData = Dataset_FromProto(mapCtx, in.GetTrainingData())
	out.Hyperparameters = Hyperparameters_FromProto(mapCtx, in.GetHyperparameters())
	return out
}
func TuningTask_ToProto(mapCtx *direct.MapContext, in *krm.TuningTask) *pb.TuningTask {
	if in == nil {
		return nil
	}
	out := &pb.TuningTask{}
	// MISSING: StartTime
	// MISSING: CompleteTime
	// MISSING: Snapshots
	out.TrainingData = Dataset_ToProto(mapCtx, in.TrainingData)
	out.Hyperparameters = Hyperparameters_ToProto(mapCtx, in.Hyperparameters)
	return out
}
func TuningTaskObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TuningTask) *krm.TuningTaskObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TuningTaskObservedState{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.CompleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCompleteTime())
	out.Snapshots = direct.Slice_FromProto(mapCtx, in.Snapshots, TuningSnapshot_FromProto)
	// MISSING: TrainingData
	// MISSING: Hyperparameters
	return out
}
func TuningTaskObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TuningTaskObservedState) *pb.TuningTask {
	if in == nil {
		return nil
	}
	out := &pb.TuningTask{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.CompleteTime = direct.StringTimestamp_ToProto(mapCtx, in.CompleteTime)
	out.Snapshots = direct.Slice_ToProto(mapCtx, in.Snapshots, TuningSnapshot_ToProto)
	// MISSING: TrainingData
	// MISSING: Hyperparameters
	return out
}
