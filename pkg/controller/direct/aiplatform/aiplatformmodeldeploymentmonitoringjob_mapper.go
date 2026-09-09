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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	monitoringv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	vertexaiv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata_FromProto(mapCtx *direct.MapContext, in *pb.ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata) *krm.ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata {
	if in == nil {
		return nil
	}
	out := &krm.ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata{}
	out.RunTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRunTime())
	out.Status = direct.Status_FromProto(mapCtx, in.GetStatus())
	return out
}

func ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata_ToProto(mapCtx *direct.MapContext, in *krm.ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata) *pb.ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata{}
	out.RunTime = direct.StringTimestamp_ToProto(mapCtx, in.RunTime)
	out.Status = direct.Status_ToProto(mapCtx, in.Status)
	return out
}

func ModelMonitoringAlertConfig_FromProto(mapCtx *direct.MapContext, in *pb.ModelMonitoringAlertConfig) *krm.ModelMonitoringAlertConfig {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitoringAlertConfig{}
	if in.GetEmailAlertConfig() != nil {
		out.EmailAlertConfig = ModelMonitoringAlertConfig_EmailAlertConfig_FromProto(mapCtx, in.GetEmailAlertConfig())
	}
	out.EnableLogging = direct.LazyPtr(in.GetEnableLogging())
	for _, channel := range in.GetNotificationChannels() {
		out.NotificationChannelRefs = append(out.NotificationChannelRefs, monitoringv1beta1.MonitoringNotificationChannelRef{External: channel})
	}
	return out
}

func ModelMonitoringAlertConfig_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringAlertConfig) *pb.ModelMonitoringAlertConfig {
	if in == nil {
		return nil
	}
	out := &pb.ModelMonitoringAlertConfig{}
	if in.EmailAlertConfig != nil {
		out.Alert = &pb.ModelMonitoringAlertConfig_EmailAlertConfig_{
			EmailAlertConfig: ModelMonitoringAlertConfig_EmailAlertConfig_ToProto(mapCtx, in.EmailAlertConfig),
		}
	}
	out.EnableLogging = direct.ValueOf(in.EnableLogging)
	for _, ref := range in.NotificationChannelRefs {
		out.NotificationChannels = append(out.NotificationChannels, ref.External)
	}
	return out
}

func ModelMonitoringObjectiveConfig_TrainingDataset_FromProto(mapCtx *direct.MapContext, in *pb.ModelMonitoringObjectiveConfig_TrainingDataset) *krm.ModelMonitoringObjectiveConfig_TrainingDataset {
	if in == nil {
		return nil
	}
	out := &krm.ModelMonitoringObjectiveConfig_TrainingDataset{}
	if in.GetDataset() != "" {
		out.DatasetRef = &vertexaiv1beta1.VertexAIDatasetRef{External: in.GetDataset()}
	}
	out.GCSSource = GCSSource_FromProto(mapCtx, in.GetGcsSource())
	out.BigquerySource = BigQuerySource_FromProto(mapCtx, in.GetBigquerySource())
	out.DataFormat = direct.LazyPtr(in.GetDataFormat())
	out.TargetField = direct.LazyPtr(in.GetTargetField())
	out.LoggingSamplingStrategy = SamplingStrategy_FromProto(mapCtx, in.GetLoggingSamplingStrategy())
	return out
}

func ModelMonitoringObjectiveConfig_TrainingDataset_ToProto(mapCtx *direct.MapContext, in *krm.ModelMonitoringObjectiveConfig_TrainingDataset) *pb.ModelMonitoringObjectiveConfig_TrainingDataset {
	if in == nil {
		return nil
	}
	out := &pb.ModelMonitoringObjectiveConfig_TrainingDataset{}
	if in.DatasetRef != nil {
		out.DataSource = &pb.ModelMonitoringObjectiveConfig_TrainingDataset_Dataset{Dataset: in.DatasetRef.External}
	}
	if in.GCSSource != nil {
		out.DataSource = &pb.ModelMonitoringObjectiveConfig_TrainingDataset_GcsSource{GcsSource: GCSSource_ToProto(mapCtx, in.GCSSource)}
	}
	if in.BigquerySource != nil {
		out.DataSource = &pb.ModelMonitoringObjectiveConfig_TrainingDataset_BigquerySource{BigquerySource: BigQuerySource_ToProto(mapCtx, in.BigquerySource)}
	}
	out.DataFormat = direct.ValueOf(in.DataFormat)
	out.TargetField = direct.ValueOf(in.TargetField)
	out.LoggingSamplingStrategy = SamplingStrategy_ToProto(mapCtx, in.LoggingSamplingStrategy)
	return out
}
