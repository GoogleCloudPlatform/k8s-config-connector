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

package visionai

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AIEnabledDevicesInputConfig_FromProto(mapCtx *direct.MapContext, in *pb.AIEnabledDevicesInputConfig) *krm.AIEnabledDevicesInputConfig {
	if in == nil {
		return nil
	}
	out := &krm.AIEnabledDevicesInputConfig{}
	return out
}
func AIEnabledDevicesInputConfig_ToProto(mapCtx *direct.MapContext, in *krm.AIEnabledDevicesInputConfig) *pb.AIEnabledDevicesInputConfig {
	if in == nil {
		return nil
	}
	out := &pb.AIEnabledDevicesInputConfig{}
	return out
}
func ApplicationConfigs_FromProto(mapCtx *direct.MapContext, in *pb.ApplicationConfigs) *krm.ApplicationConfigs {
	if in == nil {
		return nil
	}
	out := &krm.ApplicationConfigs{}
	out.Nodes = direct.Slice_FromProto(mapCtx, in.Nodes, Node_FromProto)
	out.EventDeliveryConfig = ApplicationConfigs_EventDeliveryConfig_FromProto(mapCtx, in.GetEventDeliveryConfig())
	return out
}
func ApplicationConfigs_ToProto(mapCtx *direct.MapContext, in *krm.ApplicationConfigs) *pb.ApplicationConfigs {
	if in == nil {
		return nil
	}
	out := &pb.ApplicationConfigs{}
	out.Nodes = direct.Slice_ToProto(mapCtx, in.Nodes, Node_ToProto)
	out.EventDeliveryConfig = ApplicationConfigs_EventDeliveryConfig_ToProto(mapCtx, in.EventDeliveryConfig)
	return out
}
func ApplicationConfigs_EventDeliveryConfig_FromProto(mapCtx *direct.MapContext, in *pb.ApplicationConfigs_EventDeliveryConfig) *krm.ApplicationConfigs_EventDeliveryConfig {
	if in == nil {
		return nil
	}
	out := &krm.ApplicationConfigs_EventDeliveryConfig{}
	out.Channel = direct.LazyPtr(in.GetChannel())
	out.MinimalDeliveryInterval = direct.StringDuration_FromProto(mapCtx, in.GetMinimalDeliveryInterval())
	return out
}
func ApplicationConfigs_EventDeliveryConfig_ToProto(mapCtx *direct.MapContext, in *krm.ApplicationConfigs_EventDeliveryConfig) *pb.ApplicationConfigs_EventDeliveryConfig {
	if in == nil {
		return nil
	}
	out := &pb.ApplicationConfigs_EventDeliveryConfig{}
	out.Channel = direct.ValueOf(in.Channel)
	out.MinimalDeliveryInterval = direct.StringDuration_ToProto(mapCtx, in.MinimalDeliveryInterval)
	return out
}
func AutoscalingMetricSpec_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingMetricSpec) *krm.AutoscalingMetricSpec {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingMetricSpec{}
	out.MetricName = direct.LazyPtr(in.GetMetricName())
	out.Target = direct.LazyPtr(in.GetTarget())
	return out
}
func AutoscalingMetricSpec_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingMetricSpec) *pb.AutoscalingMetricSpec {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingMetricSpec{}
	out.MetricName = direct.ValueOf(in.MetricName)
	out.Target = direct.ValueOf(in.Target)
	return out
}
func BigQueryConfig_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryConfig) *krm.BigQueryConfig {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryConfig{}
	out.Table = direct.LazyPtr(in.GetTable())
	out.CloudFunctionMapping = in.CloudFunctionMapping
	out.CreateDefaultTableIfNotExists = direct.LazyPtr(in.GetCreateDefaultTableIfNotExists())
	return out
}
func BigQueryConfig_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryConfig) *pb.BigQueryConfig {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryConfig{}
	out.Table = direct.ValueOf(in.Table)
	out.CloudFunctionMapping = in.CloudFunctionMapping
	out.CreateDefaultTableIfNotExists = direct.ValueOf(in.CreateDefaultTableIfNotExists)
	return out
}
func DedicatedResources_FromProto(mapCtx *direct.MapContext, in *pb.DedicatedResources) *krm.DedicatedResources {
	if in == nil {
		return nil
	}
	out := &krm.DedicatedResources{}
	out.MachineSpec = MachineSpec_FromProto(mapCtx, in.GetMachineSpec())
	out.MinReplicaCount = direct.LazyPtr(in.GetMinReplicaCount())
	out.MaxReplicaCount = direct.LazyPtr(in.GetMaxReplicaCount())
	out.AutoscalingMetricSpecs = direct.Slice_FromProto(mapCtx, in.AutoscalingMetricSpecs, AutoscalingMetricSpec_FromProto)
	return out
}
func DedicatedResources_ToProto(mapCtx *direct.MapContext, in *krm.DedicatedResources) *pb.DedicatedResources {
	if in == nil {
		return nil
	}
	out := &pb.DedicatedResources{}
	out.MachineSpec = MachineSpec_ToProto(mapCtx, in.MachineSpec)
	out.MinReplicaCount = direct.ValueOf(in.MinReplicaCount)
	out.MaxReplicaCount = direct.ValueOf(in.MaxReplicaCount)
	out.AutoscalingMetricSpecs = direct.Slice_ToProto(mapCtx, in.AutoscalingMetricSpecs, AutoscalingMetricSpec_ToProto)
	return out
}
func Draft_FromProto(mapCtx *direct.MapContext, in *pb.Draft) *krm.Draft {
	if in == nil {
		return nil
	}
	out := &krm.Draft{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DraftApplicationConfigs = ApplicationConfigs_FromProto(mapCtx, in.GetDraftApplicationConfigs())
	return out
}
func Draft_ToProto(mapCtx *direct.MapContext, in *krm.Draft) *pb.Draft {
	if in == nil {
		return nil
	}
	out := &pb.Draft{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.DraftApplicationConfigs = ApplicationConfigs_ToProto(mapCtx, in.DraftApplicationConfigs)
	return out
}
func DraftObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Draft) *krm.DraftObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DraftObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DraftApplicationConfigs
	return out
}
func DraftObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DraftObservedState) *pb.Draft {
	if in == nil {
		return nil
	}
	out := &pb.Draft{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DraftApplicationConfigs
	return out
}
func GcsOutputConfig_FromProto(mapCtx *direct.MapContext, in *pb.GcsOutputConfig) *krm.GcsOutputConfig {
	if in == nil {
		return nil
	}
	out := &krm.GcsOutputConfig{}
	out.GcsPath = direct.LazyPtr(in.GetGcsPath())
	return out
}
func GcsOutputConfig_ToProto(mapCtx *direct.MapContext, in *krm.GcsOutputConfig) *pb.GcsOutputConfig {
	if in == nil {
		return nil
	}
	out := &pb.GcsOutputConfig{}
	out.GcsPath = direct.ValueOf(in.GcsPath)
	return out
}
func GeneralObjectDetectionConfig_FromProto(mapCtx *direct.MapContext, in *pb.GeneralObjectDetectionConfig) *krm.GeneralObjectDetectionConfig {
	if in == nil {
		return nil
	}
	out := &krm.GeneralObjectDetectionConfig{}
	return out
}
func GeneralObjectDetectionConfig_ToProto(mapCtx *direct.MapContext, in *krm.GeneralObjectDetectionConfig) *pb.GeneralObjectDetectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.GeneralObjectDetectionConfig{}
	return out
}
func MachineSpec_FromProto(mapCtx *direct.MapContext, in *pb.MachineSpec) *krm.MachineSpec {
	if in == nil {
		return nil
	}
	out := &krm.MachineSpec{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.AcceleratorType = direct.Enum_FromProto(mapCtx, in.GetAcceleratorType())
	out.AcceleratorCount = direct.LazyPtr(in.GetAcceleratorCount())
	return out
}
func MachineSpec_ToProto(mapCtx *direct.MapContext, in *krm.MachineSpec) *pb.MachineSpec {
	if in == nil {
		return nil
	}
	out := &pb.MachineSpec{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.AcceleratorType = direct.Enum_ToProto[pb.AcceleratorType](mapCtx, in.AcceleratorType)
	out.AcceleratorCount = direct.ValueOf(in.AcceleratorCount)
	return out
}
func MediaWarehouseConfig_FromProto(mapCtx *direct.MapContext, in *pb.MediaWarehouseConfig) *krm.MediaWarehouseConfig {
	if in == nil {
		return nil
	}
	out := &krm.MediaWarehouseConfig{}
	out.Corpus = direct.LazyPtr(in.GetCorpus())
	out.Region = direct.LazyPtr(in.GetRegion())
	out.Ttl = direct.StringDuration_FromProto(mapCtx, in.GetTtl())
	return out
}
func MediaWarehouseConfig_ToProto(mapCtx *direct.MapContext, in *krm.MediaWarehouseConfig) *pb.MediaWarehouseConfig {
	if in == nil {
		return nil
	}
	out := &pb.MediaWarehouseConfig{}
	out.Corpus = direct.ValueOf(in.Corpus)
	out.Region = direct.ValueOf(in.Region)
	out.Ttl = direct.StringDuration_ToProto(mapCtx, in.Ttl)
	return out
}
func Node_FromProto(mapCtx *direct.MapContext, in *pb.Node) *krm.Node {
	if in == nil {
		return nil
	}
	out := &krm.Node{}
	out.OutputAllOutputChannelsToStream = direct.LazyPtr(in.GetOutputAllOutputChannelsToStream())
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.NodeConfig = ProcessorConfig_FromProto(mapCtx, in.GetNodeConfig())
	out.Processor = direct.LazyPtr(in.GetProcessor())
	out.Parents = direct.Slice_FromProto(mapCtx, in.Parents, Node_InputEdge_FromProto)
	return out
}
func Node_ToProto(mapCtx *direct.MapContext, in *krm.Node) *pb.Node {
	if in == nil {
		return nil
	}
	out := &pb.Node{}
	if oneof := Node_OutputAllOutputChannelsToStream_ToProto(mapCtx, in.OutputAllOutputChannelsToStream); oneof != nil {
		out.StreamOutputConfig = oneof
	}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.NodeConfig = ProcessorConfig_ToProto(mapCtx, in.NodeConfig)
	out.Processor = direct.ValueOf(in.Processor)
	out.Parents = direct.Slice_ToProto(mapCtx, in.Parents, Node_InputEdge_ToProto)
	return out
}
func Node_InputEdge_FromProto(mapCtx *direct.MapContext, in *pb.Node_InputEdge) *krm.Node_InputEdge {
	if in == nil {
		return nil
	}
	out := &krm.Node_InputEdge{}
	out.ParentNode = direct.LazyPtr(in.GetParentNode())
	out.ParentOutputChannel = direct.LazyPtr(in.GetParentOutputChannel())
	out.ConnectedInputChannel = direct.LazyPtr(in.GetConnectedInputChannel())
	return out
}
func Node_InputEdge_ToProto(mapCtx *direct.MapContext, in *krm.Node_InputEdge) *pb.Node_InputEdge {
	if in == nil {
		return nil
	}
	out := &pb.Node_InputEdge{}
	out.ParentNode = direct.ValueOf(in.ParentNode)
	out.ParentOutputChannel = direct.ValueOf(in.ParentOutputChannel)
	out.ConnectedInputChannel = direct.ValueOf(in.ConnectedInputChannel)
	return out
}
func NormalizedPolygon_FromProto(mapCtx *direct.MapContext, in *pb.NormalizedPolygon) *krm.NormalizedPolygon {
	if in == nil {
		return nil
	}
	out := &krm.NormalizedPolygon{}
	out.NormalizedVertices = direct.Slice_FromProto(mapCtx, in.NormalizedVertices, NormalizedVertex_FromProto)
	return out
}
func NormalizedPolygon_ToProto(mapCtx *direct.MapContext, in *krm.NormalizedPolygon) *pb.NormalizedPolygon {
	if in == nil {
		return nil
	}
	out := &pb.NormalizedPolygon{}
	out.NormalizedVertices = direct.Slice_ToProto(mapCtx, in.NormalizedVertices, NormalizedVertex_ToProto)
	return out
}
func NormalizedPolyline_FromProto(mapCtx *direct.MapContext, in *pb.NormalizedPolyline) *krm.NormalizedPolyline {
	if in == nil {
		return nil
	}
	out := &krm.NormalizedPolyline{}
	out.NormalizedVertices = direct.Slice_FromProto(mapCtx, in.NormalizedVertices, NormalizedVertex_FromProto)
	return out
}
func NormalizedPolyline_ToProto(mapCtx *direct.MapContext, in *krm.NormalizedPolyline) *pb.NormalizedPolyline {
	if in == nil {
		return nil
	}
	out := &pb.NormalizedPolyline{}
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
func OccupancyCountConfig_FromProto(mapCtx *direct.MapContext, in *pb.OccupancyCountConfig) *krm.OccupancyCountConfig {
	if in == nil {
		return nil
	}
	out := &krm.OccupancyCountConfig{}
	out.EnablePeopleCounting = direct.LazyPtr(in.GetEnablePeopleCounting())
	out.EnableVehicleCounting = direct.LazyPtr(in.GetEnableVehicleCounting())
	out.EnableDwellingTimeTracking = direct.LazyPtr(in.GetEnableDwellingTimeTracking())
	return out
}
func OccupancyCountConfig_ToProto(mapCtx *direct.MapContext, in *krm.OccupancyCountConfig) *pb.OccupancyCountConfig {
	if in == nil {
		return nil
	}
	out := &pb.OccupancyCountConfig{}
	out.EnablePeopleCounting = direct.ValueOf(in.EnablePeopleCounting)
	out.EnableVehicleCounting = direct.ValueOf(in.EnableVehicleCounting)
	out.EnableDwellingTimeTracking = direct.ValueOf(in.EnableDwellingTimeTracking)
	return out
}
func PersonBlurConfig_FromProto(mapCtx *direct.MapContext, in *pb.PersonBlurConfig) *krm.PersonBlurConfig {
	if in == nil {
		return nil
	}
	out := &krm.PersonBlurConfig{}
	out.PersonBlurType = direct.Enum_FromProto(mapCtx, in.GetPersonBlurType())
	out.FacesOnly = direct.LazyPtr(in.GetFacesOnly())
	return out
}
func PersonBlurConfig_ToProto(mapCtx *direct.MapContext, in *krm.PersonBlurConfig) *pb.PersonBlurConfig {
	if in == nil {
		return nil
	}
	out := &pb.PersonBlurConfig{}
	out.PersonBlurType = direct.Enum_ToProto[pb.PersonBlurConfig_PersonBlurType](mapCtx, in.PersonBlurType)
	out.FacesOnly = direct.ValueOf(in.FacesOnly)
	return out
}
func PersonVehicleDetectionConfig_FromProto(mapCtx *direct.MapContext, in *pb.PersonVehicleDetectionConfig) *krm.PersonVehicleDetectionConfig {
	if in == nil {
		return nil
	}
	out := &krm.PersonVehicleDetectionConfig{}
	out.EnablePeopleCounting = direct.LazyPtr(in.GetEnablePeopleCounting())
	out.EnableVehicleCounting = direct.LazyPtr(in.GetEnableVehicleCounting())
	return out
}
func PersonVehicleDetectionConfig_ToProto(mapCtx *direct.MapContext, in *krm.PersonVehicleDetectionConfig) *pb.PersonVehicleDetectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.PersonVehicleDetectionConfig{}
	out.EnablePeopleCounting = direct.ValueOf(in.EnablePeopleCounting)
	out.EnableVehicleCounting = direct.ValueOf(in.EnableVehicleCounting)
	return out
}
func PersonalProtectiveEquipmentDetectionConfig_FromProto(mapCtx *direct.MapContext, in *pb.PersonalProtectiveEquipmentDetectionConfig) *krm.PersonalProtectiveEquipmentDetectionConfig {
	if in == nil {
		return nil
	}
	out := &krm.PersonalProtectiveEquipmentDetectionConfig{}
	out.EnableFaceCoverageDetection = direct.LazyPtr(in.GetEnableFaceCoverageDetection())
	out.EnableHeadCoverageDetection = direct.LazyPtr(in.GetEnableHeadCoverageDetection())
	out.EnableHandsCoverageDetection = direct.LazyPtr(in.GetEnableHandsCoverageDetection())
	return out
}
func PersonalProtectiveEquipmentDetectionConfig_ToProto(mapCtx *direct.MapContext, in *krm.PersonalProtectiveEquipmentDetectionConfig) *pb.PersonalProtectiveEquipmentDetectionConfig {
	if in == nil {
		return nil
	}
	out := &pb.PersonalProtectiveEquipmentDetectionConfig{}
	out.EnableFaceCoverageDetection = direct.ValueOf(in.EnableFaceCoverageDetection)
	out.EnableHeadCoverageDetection = direct.ValueOf(in.EnableHeadCoverageDetection)
	out.EnableHandsCoverageDetection = direct.ValueOf(in.EnableHandsCoverageDetection)
	return out
}
func ProcessorConfig_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorConfig) *krm.ProcessorConfig {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorConfig{}
	out.VideoStreamInputConfig = VideoStreamInputConfig_FromProto(mapCtx, in.GetVideoStreamInputConfig())
	out.AiEnabledDevicesInputConfig = AIEnabledDevicesInputConfig_FromProto(mapCtx, in.GetAiEnabledDevicesInputConfig())
	out.MediaWarehouseConfig = MediaWarehouseConfig_FromProto(mapCtx, in.GetMediaWarehouseConfig())
	out.PersonBlurConfig = PersonBlurConfig_FromProto(mapCtx, in.GetPersonBlurConfig())
	out.OccupancyCountConfig = OccupancyCountConfig_FromProto(mapCtx, in.GetOccupancyCountConfig())
	out.PersonVehicleDetectionConfig = PersonVehicleDetectionConfig_FromProto(mapCtx, in.GetPersonVehicleDetectionConfig())
	out.VertexAutomlVisionConfig = VertexAutoMLVisionConfig_FromProto(mapCtx, in.GetVertexAutomlVisionConfig())
	out.VertexAutomlVideoConfig = VertexAutoMLVideoConfig_FromProto(mapCtx, in.GetVertexAutomlVideoConfig())
	out.VertexCustomConfig = VertexCustomConfig_FromProto(mapCtx, in.GetVertexCustomConfig())
	out.GeneralObjectDetectionConfig = GeneralObjectDetectionConfig_FromProto(mapCtx, in.GetGeneralObjectDetectionConfig())
	out.BigQueryConfig = BigQueryConfig_FromProto(mapCtx, in.GetBigQueryConfig())
	out.GcsOutputConfig = GcsOutputConfig_FromProto(mapCtx, in.GetGcsOutputConfig())
	out.ProductRecognizerConfig = ProductRecognizerConfig_FromProto(mapCtx, in.GetProductRecognizerConfig())
	out.PersonalProtectiveEquipmentDetectionConfig = PersonalProtectiveEquipmentDetectionConfig_FromProto(mapCtx, in.GetPersonalProtectiveEquipmentDetectionConfig())
	out.TagRecognizerConfig = TagRecognizerConfig_FromProto(mapCtx, in.GetTagRecognizerConfig())
	out.UniversalInputConfig = UniversalInputConfig_FromProto(mapCtx, in.GetUniversalInputConfig())
	out.ExperimentalConfig = ExperimentalConfig_FromProto(mapCtx, in.GetExperimentalConfig())
	return out
}
func ProcessorConfig_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorConfig) *pb.ProcessorConfig {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorConfig{}
	if oneof := VideoStreamInputConfig_ToProto(mapCtx, in.VideoStreamInputConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_VideoStreamInputConfig{VideoStreamInputConfig: oneof}
	}
	if oneof := AIEnabledDevicesInputConfig_ToProto(mapCtx, in.AiEnabledDevicesInputConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_AiEnabledDevicesInputConfig{AiEnabledDevicesInputConfig: oneof}
	}
	if oneof := MediaWarehouseConfig_ToProto(mapCtx, in.MediaWarehouseConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_MediaWarehouseConfig{MediaWarehouseConfig: oneof}
	}
	if oneof := PersonBlurConfig_ToProto(mapCtx, in.PersonBlurConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_PersonBlurConfig{PersonBlurConfig: oneof}
	}
	if oneof := OccupancyCountConfig_ToProto(mapCtx, in.OccupancyCountConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_OccupancyCountConfig{OccupancyCountConfig: oneof}
	}
	if oneof := PersonVehicleDetectionConfig_ToProto(mapCtx, in.PersonVehicleDetectionConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_PersonVehicleDetectionConfig{PersonVehicleDetectionConfig: oneof}
	}
	if oneof := VertexAutoMLVisionConfig_ToProto(mapCtx, in.VertexAutomlVisionConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_VertexAutomlVisionConfig{VertexAutomlVisionConfig: oneof}
	}
	if oneof := VertexAutoMLVideoConfig_ToProto(mapCtx, in.VertexAutomlVideoConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_VertexAutomlVideoConfig{VertexAutomlVideoConfig: oneof}
	}
	if oneof := VertexCustomConfig_ToProto(mapCtx, in.VertexCustomConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_VertexCustomConfig{VertexCustomConfig: oneof}
	}
	if oneof := GeneralObjectDetectionConfig_ToProto(mapCtx, in.GeneralObjectDetectionConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_GeneralObjectDetectionConfig{GeneralObjectDetectionConfig: oneof}
	}
	if oneof := BigQueryConfig_ToProto(mapCtx, in.BigQueryConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_BigQueryConfig{BigQueryConfig: oneof}
	}
	if oneof := GcsOutputConfig_ToProto(mapCtx, in.GcsOutputConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_GcsOutputConfig{GcsOutputConfig: oneof}
	}
	if oneof := ProductRecognizerConfig_ToProto(mapCtx, in.ProductRecognizerConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_ProductRecognizerConfig{ProductRecognizerConfig: oneof}
	}
	if oneof := PersonalProtectiveEquipmentDetectionConfig_ToProto(mapCtx, in.PersonalProtectiveEquipmentDetectionConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_PersonalProtectiveEquipmentDetectionConfig{PersonalProtectiveEquipmentDetectionConfig: oneof}
	}
	if oneof := TagRecognizerConfig_ToProto(mapCtx, in.TagRecognizerConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_TagRecognizerConfig{TagRecognizerConfig: oneof}
	}
	if oneof := UniversalInputConfig_ToProto(mapCtx, in.UniversalInputConfig); oneof != nil {
		out.ProcessorConfig = &pb.ProcessorConfig_UniversalInputConfig{UniversalInputConfig: oneof}
	}
	out.ExperimentalConfig = ExperimentalConfig_ToProto(mapCtx, in.ExperimentalConfig)
	return out
}
func ProductRecognizerConfig_FromProto(mapCtx *direct.MapContext, in *pb.ProductRecognizerConfig) *krm.ProductRecognizerConfig {
	if in == nil {
		return nil
	}
	out := &krm.ProductRecognizerConfig{}
	out.RetailEndpoint = direct.LazyPtr(in.GetRetailEndpoint())
	out.RecognitionConfidenceThreshold = direct.LazyPtr(in.GetRecognitionConfidenceThreshold())
	return out
}
func ProductRecognizerConfig_ToProto(mapCtx *direct.MapContext, in *krm.ProductRecognizerConfig) *pb.ProductRecognizerConfig {
	if in == nil {
		return nil
	}
	out := &pb.ProductRecognizerConfig{}
	out.RetailEndpoint = direct.ValueOf(in.RetailEndpoint)
	out.RecognitionConfidenceThreshold = direct.ValueOf(in.RecognitionConfidenceThreshold)
	return out
}
func StreamAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.StreamAnnotation) *krm.StreamAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.StreamAnnotation{}
	out.ActiveZone = NormalizedPolygon_FromProto(mapCtx, in.GetActiveZone())
	out.CrossingLine = NormalizedPolyline_FromProto(mapCtx, in.GetCrossingLine())
	out.ID = direct.LazyPtr(in.GetId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.SourceStream = direct.LazyPtr(in.GetSourceStream())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func StreamAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.StreamAnnotation) *pb.StreamAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.StreamAnnotation{}
	if oneof := NormalizedPolygon_ToProto(mapCtx, in.ActiveZone); oneof != nil {
		out.AnnotationPayload = &pb.StreamAnnotation_ActiveZone{ActiveZone: oneof}
	}
	if oneof := NormalizedPolyline_ToProto(mapCtx, in.CrossingLine); oneof != nil {
		out.AnnotationPayload = &pb.StreamAnnotation_CrossingLine{CrossingLine: oneof}
	}
	out.Id = direct.ValueOf(in.ID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.SourceStream = direct.ValueOf(in.SourceStream)
	out.Type = direct.Enum_ToProto[pb.StreamAnnotationType](mapCtx, in.Type)
	return out
}
func StreamWithAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.StreamWithAnnotation) *krm.StreamWithAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.StreamWithAnnotation{}
	out.Stream = direct.LazyPtr(in.GetStream())
	out.ApplicationAnnotations = direct.Slice_FromProto(mapCtx, in.ApplicationAnnotations, StreamAnnotation_FromProto)
	out.NodeAnnotations = direct.Slice_FromProto(mapCtx, in.NodeAnnotations, StreamWithAnnotation_NodeAnnotation_FromProto)
	return out
}
func StreamWithAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.StreamWithAnnotation) *pb.StreamWithAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.StreamWithAnnotation{}
	out.Stream = direct.ValueOf(in.Stream)
	out.ApplicationAnnotations = direct.Slice_ToProto(mapCtx, in.ApplicationAnnotations, StreamAnnotation_ToProto)
	out.NodeAnnotations = direct.Slice_ToProto(mapCtx, in.NodeAnnotations, StreamWithAnnotation_NodeAnnotation_ToProto)
	return out
}
func StreamWithAnnotation_NodeAnnotation_FromProto(mapCtx *direct.MapContext, in *pb.StreamWithAnnotation_NodeAnnotation) *krm.StreamWithAnnotation_NodeAnnotation {
	if in == nil {
		return nil
	}
	out := &krm.StreamWithAnnotation_NodeAnnotation{}
	out.Node = direct.LazyPtr(in.GetNode())
	out.Annotations = direct.Slice_FromProto(mapCtx, in.Annotations, StreamAnnotation_FromProto)
	return out
}
func StreamWithAnnotation_NodeAnnotation_ToProto(mapCtx *direct.MapContext, in *krm.StreamWithAnnotation_NodeAnnotation) *pb.StreamWithAnnotation_NodeAnnotation {
	if in == nil {
		return nil
	}
	out := &pb.StreamWithAnnotation_NodeAnnotation{}
	out.Node = direct.ValueOf(in.Node)
	out.Annotations = direct.Slice_ToProto(mapCtx, in.Annotations, StreamAnnotation_ToProto)
	return out
}
func TagParsingConfig_FromProto(mapCtx *direct.MapContext, in *pb.TagParsingConfig) *krm.TagParsingConfig {
	if in == nil {
		return nil
	}
	out := &krm.TagParsingConfig{}
	out.EntityParsingConfigs = direct.Slice_FromProto(mapCtx, in.EntityParsingConfigs, TagParsingConfig_EntityParsingConfig_FromProto)
	return out
}
func TagParsingConfig_ToProto(mapCtx *direct.MapContext, in *krm.TagParsingConfig) *pb.TagParsingConfig {
	if in == nil {
		return nil
	}
	out := &pb.TagParsingConfig{}
	out.EntityParsingConfigs = direct.Slice_ToProto(mapCtx, in.EntityParsingConfigs, TagParsingConfig_EntityParsingConfig_ToProto)
	return out
}
func TagParsingConfig_EntityParsingConfig_FromProto(mapCtx *direct.MapContext, in *pb.TagParsingConfig_EntityParsingConfig) *krm.TagParsingConfig_EntityParsingConfig {
	if in == nil {
		return nil
	}
	out := &krm.TagParsingConfig_EntityParsingConfig{}
	out.EntityClass = direct.LazyPtr(in.GetEntityClass())
	out.Regex = direct.LazyPtr(in.GetRegex())
	out.EntityMatchingStrategy = direct.Enum_FromProto(mapCtx, in.GetEntityMatchingStrategy())
	return out
}
func TagParsingConfig_EntityParsingConfig_ToProto(mapCtx *direct.MapContext, in *krm.TagParsingConfig_EntityParsingConfig) *pb.TagParsingConfig_EntityParsingConfig {
	if in == nil {
		return nil
	}
	out := &pb.TagParsingConfig_EntityParsingConfig{}
	out.EntityClass = direct.ValueOf(in.EntityClass)
	out.Regex = direct.ValueOf(in.Regex)
	out.EntityMatchingStrategy = direct.Enum_ToProto[pb.TagParsingConfig_EntityParsingConfig_EntityMatchingStrategy](mapCtx, in.EntityMatchingStrategy)
	return out
}
func TagRecognizerConfig_FromProto(mapCtx *direct.MapContext, in *pb.TagRecognizerConfig) *krm.TagRecognizerConfig {
	if in == nil {
		return nil
	}
	out := &krm.TagRecognizerConfig{}
	out.EntityDetectionConfidenceThreshold = direct.LazyPtr(in.GetEntityDetectionConfidenceThreshold())
	out.TagParsingConfig = TagParsingConfig_FromProto(mapCtx, in.GetTagParsingConfig())
	return out
}
func TagRecognizerConfig_ToProto(mapCtx *direct.MapContext, in *krm.TagRecognizerConfig) *pb.TagRecognizerConfig {
	if in == nil {
		return nil
	}
	out := &pb.TagRecognizerConfig{}
	out.EntityDetectionConfidenceThreshold = direct.ValueOf(in.EntityDetectionConfidenceThreshold)
	out.TagParsingConfig = TagParsingConfig_ToProto(mapCtx, in.TagParsingConfig)
	return out
}
func UniversalInputConfig_FromProto(mapCtx *direct.MapContext, in *pb.UniversalInputConfig) *krm.UniversalInputConfig {
	if in == nil {
		return nil
	}
	out := &krm.UniversalInputConfig{}
	return out
}
func UniversalInputConfig_ToProto(mapCtx *direct.MapContext, in *krm.UniversalInputConfig) *pb.UniversalInputConfig {
	if in == nil {
		return nil
	}
	out := &pb.UniversalInputConfig{}
	return out
}
func VertexAutoMLVideoConfig_FromProto(mapCtx *direct.MapContext, in *pb.VertexAutoMLVideoConfig) *krm.VertexAutoMLVideoConfig {
	if in == nil {
		return nil
	}
	out := &krm.VertexAutoMLVideoConfig{}
	out.ConfidenceThreshold = direct.LazyPtr(in.GetConfidenceThreshold())
	out.BlockedLabels = in.BlockedLabels
	out.MaxPredictions = direct.LazyPtr(in.GetMaxPredictions())
	out.BoundingBoxSizeLimit = direct.LazyPtr(in.GetBoundingBoxSizeLimit())
	return out
}
func VertexAutoMLVideoConfig_ToProto(mapCtx *direct.MapContext, in *krm.VertexAutoMLVideoConfig) *pb.VertexAutoMLVideoConfig {
	if in == nil {
		return nil
	}
	out := &pb.VertexAutoMLVideoConfig{}
	out.ConfidenceThreshold = direct.ValueOf(in.ConfidenceThreshold)
	out.BlockedLabels = in.BlockedLabels
	out.MaxPredictions = direct.ValueOf(in.MaxPredictions)
	out.BoundingBoxSizeLimit = direct.ValueOf(in.BoundingBoxSizeLimit)
	return out
}
func VertexAutoMLVisionConfig_FromProto(mapCtx *direct.MapContext, in *pb.VertexAutoMLVisionConfig) *krm.VertexAutoMLVisionConfig {
	if in == nil {
		return nil
	}
	out := &krm.VertexAutoMLVisionConfig{}
	out.ConfidenceThreshold = direct.LazyPtr(in.GetConfidenceThreshold())
	out.MaxPredictions = direct.LazyPtr(in.GetMaxPredictions())
	return out
}
func VertexAutoMLVisionConfig_ToProto(mapCtx *direct.MapContext, in *krm.VertexAutoMLVisionConfig) *pb.VertexAutoMLVisionConfig {
	if in == nil {
		return nil
	}
	out := &pb.VertexAutoMLVisionConfig{}
	out.ConfidenceThreshold = direct.ValueOf(in.ConfidenceThreshold)
	out.MaxPredictions = direct.ValueOf(in.MaxPredictions)
	return out
}
func VertexCustomConfig_FromProto(mapCtx *direct.MapContext, in *pb.VertexCustomConfig) *krm.VertexCustomConfig {
	if in == nil {
		return nil
	}
	out := &krm.VertexCustomConfig{}
	out.MaxPredictionFps = direct.LazyPtr(in.GetMaxPredictionFps())
	out.DedicatedResources = DedicatedResources_FromProto(mapCtx, in.GetDedicatedResources())
	out.PostProcessingCloudFunction = direct.LazyPtr(in.GetPostProcessingCloudFunction())
	out.AttachApplicationMetadata = direct.LazyPtr(in.GetAttachApplicationMetadata())
	out.DynamicConfigInputTopic = in.DynamicConfigInputTopic
	return out
}
func VertexCustomConfig_ToProto(mapCtx *direct.MapContext, in *krm.VertexCustomConfig) *pb.VertexCustomConfig {
	if in == nil {
		return nil
	}
	out := &pb.VertexCustomConfig{}
	out.MaxPredictionFps = direct.ValueOf(in.MaxPredictionFps)
	out.DedicatedResources = DedicatedResources_ToProto(mapCtx, in.DedicatedResources)
	out.PostProcessingCloudFunction = direct.ValueOf(in.PostProcessingCloudFunction)
	out.AttachApplicationMetadata = direct.ValueOf(in.AttachApplicationMetadata)
	out.DynamicConfigInputTopic = in.DynamicConfigInputTopic
	return out
}
func VideoStreamInputConfig_FromProto(mapCtx *direct.MapContext, in *pb.VideoStreamInputConfig) *krm.VideoStreamInputConfig {
	if in == nil {
		return nil
	}
	out := &krm.VideoStreamInputConfig{}
	out.Streams = in.Streams
	out.StreamsWithAnnotation = direct.Slice_FromProto(mapCtx, in.StreamsWithAnnotation, StreamWithAnnotation_FromProto)
	return out
}
func VideoStreamInputConfig_ToProto(mapCtx *direct.MapContext, in *krm.VideoStreamInputConfig) *pb.VideoStreamInputConfig {
	if in == nil {
		return nil
	}
	out := &pb.VideoStreamInputConfig{}
	out.Streams = in.Streams
	out.StreamsWithAnnotation = direct.Slice_ToProto(mapCtx, in.StreamsWithAnnotation, StreamWithAnnotation_ToProto)
	return out
}
func VisionaiDraftObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Draft) *krm.VisionaiDraftObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiDraftObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DraftApplicationConfigs
	return out
}
func VisionaiDraftObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiDraftObservedState) *pb.Draft {
	if in == nil {
		return nil
	}
	out := &pb.Draft{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DraftApplicationConfigs
	return out
}
func VisionaiDraftSpec_FromProto(mapCtx *direct.MapContext, in *pb.Draft) *krm.VisionaiDraftSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiDraftSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DraftApplicationConfigs
	return out
}
func VisionaiDraftSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiDraftSpec) *pb.Draft {
	if in == nil {
		return nil
	}
	out := &pb.Draft{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DraftApplicationConfigs
	return out
}
