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

package vertexaimodeldeploymentmonitoringjob

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform"
	rpc "google.golang.org/genproto/googleapis/rpc/status"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

func Value_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *krm.Value {
	return aiplatform.Value_FromProto(mapCtx, in)
}

func Value_ToProto(mapCtx *direct.MapContext, in *krm.Value) *structpb.Value {
	return aiplatform.Value_ToProto(mapCtx, in)
}

func Status_FromProto(mapCtx *direct.MapContext, in *rpc.Status) *common.Status {
	if in == nil {
		return nil
	}
	return &common.Status{
		Code:    direct.LazyPtr(in.GetCode()),
		Message: direct.LazyPtr(in.GetMessage()),
	}
}

func Status_ToProto(mapCtx *direct.MapContext, in *common.Status) *rpc.Status {
	if in == nil {
		return nil
	}
	return &rpc.Status{
		Code:    direct.ValueOf(in.Code),
		Message: direct.ValueOf(in.Message),
	}
}

func ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata_FromProto(mapCtx *direct.MapContext, in *pb.ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata) *krm.ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata {
	if in == nil {
		return nil
	}
	out := &krm.ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata{}
	out.RunTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRunTime())
	out.Status = Status_FromProto(mapCtx, in.GetStatus())
	return out
}

func ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata_ToProto(mapCtx *direct.MapContext, in *krm.ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata) *pb.ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ModelDeploymentMonitoringJob_LatestMonitoringPipelineMetadata{}
	out.RunTime = direct.StringTimestamp_ToProto(mapCtx, in.RunTime)
	out.Status = Status_ToProto(mapCtx, in.Status)
	return out
}

func EncryptionSpec_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionSpec{}
	if in.GetKmsKeyName() != "" {
		out.KmsKeyRef = &refs.KMSCryptoKeyRef{
			External: in.GetKmsKeyName(),
		}
	}
	return out
}

func EncryptionSpec_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	if in.KmsKeyRef != nil {
		out.KmsKeyName = in.KmsKeyRef.External
	}
	return out
}
