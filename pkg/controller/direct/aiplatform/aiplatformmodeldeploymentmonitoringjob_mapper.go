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
