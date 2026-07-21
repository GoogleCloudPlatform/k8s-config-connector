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
	"strings"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	dataformv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CreatePipelineJobRequestObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CreatePipelineJobRequest) *krm.CreatePipelineJobRequestObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CreatePipelineJobRequestObservedState{}
	out.PipelineJob = VertexAIPipelineJobObservedState_FromProto(mapCtx, in.GetPipelineJob())
	return out
}

func CreatePipelineJobRequestObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CreatePipelineJobRequestObservedState) *pb.CreatePipelineJobRequest {
	if in == nil {
		return nil
	}
	out := &pb.CreatePipelineJobRequest{}
	out.PipelineJob = VertexAIPipelineJobObservedState_ToProto(mapCtx, in.PipelineJob)
	return out
}

func CreateNotebookExecutionJobRequestObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CreateNotebookExecutionJobRequest) *krm.CreateNotebookExecutionJobRequestObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CreateNotebookExecutionJobRequestObservedState{}
	out.NotebookExecutionJob = NotebookExecutionJobObservedState_FromProto(mapCtx, in.GetNotebookExecutionJob())
	return out
}

func CreateNotebookExecutionJobRequestObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CreateNotebookExecutionJobRequestObservedState) *pb.CreateNotebookExecutionJobRequest {
	if in == nil {
		return nil
	}
	out := &pb.CreateNotebookExecutionJobRequest{}
	out.NotebookExecutionJob = NotebookExecutionJobObservedState_ToProto(mapCtx, in.NotebookExecutionJob)
	return out
}

func ContextObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Context) *krm.ContextObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContextObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Labels = in.Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ParentContexts = in.ParentContexts
	out.SchemaTitle = direct.LazyPtr(in.GetSchemaTitle())
	out.SchemaVersion = direct.LazyPtr(in.GetSchemaVersion())
	out.Metadata = direct.Struct_FromProto(mapCtx, in.GetMetadata())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}

func ContextObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContextObservedState) *pb.Context {
	if in == nil {
		return nil
	}
	out := &pb.Context{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Etag = direct.ValueOf(in.Etag)
	out.Labels = in.Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ParentContexts = in.ParentContexts
	out.SchemaTitle = direct.ValueOf(in.SchemaTitle)
	out.SchemaVersion = direct.ValueOf(in.SchemaVersion)
	out.Metadata = direct.Struct_ToProto(mapCtx, in.Metadata)
	out.Description = direct.ValueOf(in.Description)
	return out
}

func ExecutionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Execution) *krm.ExecutionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Labels = in.Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.SchemaTitle = direct.LazyPtr(in.GetSchemaTitle())
	out.SchemaVersion = direct.LazyPtr(in.GetSchemaVersion())
	out.Metadata = direct.Struct_FromProto(mapCtx, in.GetMetadata())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}

func ExecutionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionObservedState) *pb.Execution {
	if in == nil {
		return nil
	}
	out := &pb.Execution{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.State = direct.Enum_ToProto[pb.Execution_State](mapCtx, in.State)
	out.Etag = direct.ValueOf(in.Etag)
	out.Labels = in.Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.SchemaTitle = direct.ValueOf(in.SchemaTitle)
	out.SchemaVersion = direct.ValueOf(in.SchemaVersion)
	out.Metadata = direct.Struct_ToProto(mapCtx, in.Metadata)
	out.Description = direct.ValueOf(in.Description)
	return out
}

func PipelineJob_FromProto(mapCtx *direct.MapContext, in *pb.PipelineJob) *krm.PipelineJob {
	return VertexAIPipelineJobSpec_FromProto(mapCtx, in)
}

func PipelineJob_ToProto(mapCtx *direct.MapContext, in *krm.PipelineJob) *pb.PipelineJob {
	return VertexAIPipelineJobSpec_ToProto(mapCtx, in)
}

func Schedule_RunResponse_FromProto(mapCtx *direct.MapContext, in *pb.Schedule_RunResponse) *krm.Schedule_RunResponse {
	if in == nil {
		return nil
	}
	out := &krm.Schedule_RunResponse{}
	out.ScheduledRunTime = direct.StringTimestamp_FromProto(mapCtx, in.GetScheduledRunTime())
	out.RunResponse = direct.LazyPtr(in.GetRunResponse())
	return out
}

func Schedule_RunResponse_ToProto(mapCtx *direct.MapContext, in *krm.Schedule_RunResponse) *pb.Schedule_RunResponse {
	if in == nil {
		return nil
	}
	out := &pb.Schedule_RunResponse{}
	out.ScheduledRunTime = direct.StringTimestamp_ToProto(mapCtx, in.ScheduledRunTime)
	out.RunResponse = direct.ValueOf(in.RunResponse)
	return out
}

func CreatePipelineJobRequest_FromProto(mapCtx *direct.MapContext, in *pb.CreatePipelineJobRequest) *krm.CreatePipelineJobRequest {
	if in == nil {
		return nil
	}
	out := &krm.CreatePipelineJobRequest{}
	out.Parent = direct.LazyPtr(in.GetParent())
	out.PipelineJob = PipelineJob_FromProto(mapCtx, in.GetPipelineJob())
	return out
}

func CreatePipelineJobRequest_ToProto(mapCtx *direct.MapContext, in *krm.CreatePipelineJobRequest) *pb.CreatePipelineJobRequest {
	if in == nil {
		return nil
	}
	out := &pb.CreatePipelineJobRequest{}
	out.Parent = direct.ValueOf(in.Parent)
	out.PipelineJob = PipelineJob_ToProto(mapCtx, in.PipelineJob)
	return out
}

func CreateNotebookExecutionJobRequest_FromProto(mapCtx *direct.MapContext, in *pb.CreateNotebookExecutionJobRequest) *krm.CreateNotebookExecutionJobRequest {
	if in == nil {
		return nil
	}
	out := &krm.CreateNotebookExecutionJobRequest{}
	out.NotebookExecutionJob = NotebookExecutionJob_FromProto(mapCtx, in.GetNotebookExecutionJob())
	out.NotebookExecutionJobID = direct.LazyPtr(in.GetNotebookExecutionJobId())
	return out
}

func CreateNotebookExecutionJobRequest_ToProto(mapCtx *direct.MapContext, in *krm.CreateNotebookExecutionJobRequest) *pb.CreateNotebookExecutionJobRequest {
	if in == nil {
		return nil
	}
	out := &pb.CreateNotebookExecutionJobRequest{}
	out.NotebookExecutionJob = NotebookExecutionJob_ToProto(mapCtx, in.NotebookExecutionJob)
	out.NotebookExecutionJobId = direct.ValueOf(in.NotebookExecutionJobID)
	return out
}

func NetworkSpec_FromProto(mapCtx *direct.MapContext, in *pb.NetworkSpec) *krm.NetworkSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSpec{}
	out.EnableInternetAccess = direct.LazyPtr(in.GetEnableInternetAccess())
	if in.GetNetwork() != "" {
		out.NetworkRef = &computev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &computev1beta1.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	return out
}

func NetworkSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSpec) *pb.NetworkSpec {
	if in == nil {
		return nil
	}
	out := &pb.NetworkSpec{}
	out.EnableInternetAccess = direct.ValueOf(in.EnableInternetAccess)
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}
	return out
}

func NotebookExecutionJob_DataformRepositorySource_FromProto(mapCtx *direct.MapContext, in *pb.NotebookExecutionJob_DataformRepositorySource) *krm.NotebookExecutionJob_DataformRepositorySource {
	if in == nil {
		return nil
	}
	out := &krm.NotebookExecutionJob_DataformRepositorySource{}
	if in.GetDataformRepositoryResourceName() != "" {
		out.DataformRepositoryRef = &dataformv1beta1.DataformRepositoryRef{External: in.GetDataformRepositoryResourceName()}
	}
	out.CommitSha = direct.LazyPtr(in.GetCommitSha())
	return out
}

func NotebookExecutionJob_DataformRepositorySource_ToProto(mapCtx *direct.MapContext, in *krm.NotebookExecutionJob_DataformRepositorySource) *pb.NotebookExecutionJob_DataformRepositorySource {
	if in == nil {
		return nil
	}
	out := &pb.NotebookExecutionJob_DataformRepositorySource{}
	if in.DataformRepositoryRef != nil {
		out.DataformRepositoryResourceName = in.DataformRepositoryRef.External
	}
	out.CommitSha = direct.ValueOf(in.CommitSha)
	return out
}

func NotebookExecutionJob_FromProto(mapCtx *direct.MapContext, in *pb.NotebookExecutionJob) *krm.NotebookExecutionJob {
	if in == nil {
		return nil
	}
	out := &krm.NotebookExecutionJob{}
	out.DataformRepositorySource = NotebookExecutionJob_DataformRepositorySource_FromProto(mapCtx, in.GetDataformRepositorySource())
	out.GCSNotebookSource = NotebookExecutionJob_GCSNotebookSource_FromProto(mapCtx, in.GetGcsNotebookSource())
	out.DirectNotebookSource = NotebookExecutionJob_DirectNotebookSource_FromProto(mapCtx, in.GetDirectNotebookSource())
	out.NotebookRuntimeTemplateResourceName = direct.LazyPtr(in.GetNotebookRuntimeTemplateResourceName())
	out.CustomEnvironmentSpec = NotebookExecutionJob_CustomEnvironmentSpec_FromProto(mapCtx, in.GetCustomEnvironmentSpec())
	if in.GetGcsOutputUri() != "" {
		out.GCSOutputRef = &storagev1beta1.StorageBucketRef{
			External: in.GetGcsOutputUri(),
		}
	}
	out.ExecutionUser = direct.LazyPtr(in.GetExecutionUser())
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.WorkbenchRuntime = NotebookExecutionJob_WorkbenchRuntime_FromProto(mapCtx, in.GetWorkbenchRuntime())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ExecutionTimeout = direct.StringDuration_FromProto(mapCtx, in.GetExecutionTimeout())
	if in.GetScheduleResourceName() != "" {
		out.ScheduleRef = &krm.VertexAIScheduleRef{External: in.GetScheduleResourceName()}
	}
	out.Labels = in.Labels
	out.KernelName = direct.LazyPtr(in.GetKernelName())
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	return out
}

func NotebookExecutionJob_ToProto(mapCtx *direct.MapContext, in *krm.NotebookExecutionJob) *pb.NotebookExecutionJob {
	if in == nil {
		return nil
	}
	out := &pb.NotebookExecutionJob{}
	if oneof := NotebookExecutionJob_DataformRepositorySource_ToProto(mapCtx, in.DataformRepositorySource); oneof != nil {
		out.NotebookSource = &pb.NotebookExecutionJob_DataformRepositorySource_{DataformRepositorySource: oneof}
	}
	if oneof := NotebookExecutionJob_GCSNotebookSource_ToProto(mapCtx, in.GCSNotebookSource); oneof != nil {
		out.NotebookSource = &pb.NotebookExecutionJob_GcsNotebookSource_{GcsNotebookSource: oneof}
	}
	if oneof := NotebookExecutionJob_DirectNotebookSource_ToProto(mapCtx, in.DirectNotebookSource); oneof != nil {
		out.NotebookSource = &pb.NotebookExecutionJob_DirectNotebookSource_{DirectNotebookSource: oneof}
	}
	if oneof := NotebookExecutionJob_NotebookRuntimeTemplateResourceName_ToProto(mapCtx, in.NotebookRuntimeTemplateResourceName); oneof != nil {
		out.EnvironmentSpec = oneof
	}
	if oneof := NotebookExecutionJob_CustomEnvironmentSpec_ToProto(mapCtx, in.CustomEnvironmentSpec); oneof != nil {
		out.EnvironmentSpec = &pb.NotebookExecutionJob_CustomEnvironmentSpec_{CustomEnvironmentSpec: oneof}
	}
	if in.GCSOutputRef != nil {
		id := &storagev1beta1.StorageBucketIdentity{}
		if err := id.FromExternal(in.GCSOutputRef.External); err != nil {
			mapCtx.Errorf("gcsOutputRef: %v", err)
		} else {
			out.ExecutionSink = &pb.NotebookExecutionJob_GcsOutputUri{
				GcsOutputUri: "gs://" + id.Bucket,
			}
		}
	}
	if oneof := NotebookExecutionJob_ExecutionUser_ToProto(mapCtx, in.ExecutionUser); oneof != nil {
		out.ExecutionIdentity = oneof
	}
	if in.ServiceAccountRef != nil {
		out.ExecutionIdentity = &pb.NotebookExecutionJob_ServiceAccount{ServiceAccount: in.ServiceAccountRef.External}
	}
	if oneof := NotebookExecutionJob_WorkbenchRuntime_ToProto(mapCtx, in.WorkbenchRuntime); oneof != nil {
		out.RuntimeEnvironment = &pb.NotebookExecutionJob_WorkbenchRuntime_{WorkbenchRuntime: oneof}
	}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ExecutionTimeout = direct.StringDuration_ToProto(mapCtx, in.ExecutionTimeout)
	if in.ScheduleRef != nil {
		out.ScheduleResourceName = in.ScheduleRef.External
	}
	out.Labels = in.Labels
	out.KernelName = direct.ValueOf(in.KernelName)
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	return out
}

func NotebookExecutionJob_GCSNotebookSource_FromProto(mapCtx *direct.MapContext, in *pb.NotebookExecutionJob_GcsNotebookSource) *krm.NotebookExecutionJob_GCSNotebookSource {
	if in == nil {
		return nil
	}
	out := &krm.NotebookExecutionJob_GCSNotebookSource{}
	if in.GetUri() != "" {
		uri := in.GetUri()
		if strings.HasPrefix(uri, "gs://") {
			trimmed := strings.TrimPrefix(uri, "gs://")
			parts := strings.SplitN(trimmed, "/", 2)
			if len(parts) > 0 && parts[0] != "" {
				out.BucketRef = &storagev1beta1.StorageBucketRef{
					External: parts[0],
				}
				if len(parts) > 1 {
					out.Object = direct.LazyPtr(parts[1])
				}
			}
		} else {
			out.BucketRef = &storagev1beta1.StorageBucketRef{
				External: uri,
			}
		}
	}
	out.Generation = direct.LazyPtr(in.GetGeneration())
	return out
}

func NotebookExecutionJob_GCSNotebookSource_ToProto(mapCtx *direct.MapContext, in *krm.NotebookExecutionJob_GCSNotebookSource) *pb.NotebookExecutionJob_GcsNotebookSource {
	if in == nil {
		return nil
	}
	out := &pb.NotebookExecutionJob_GcsNotebookSource{}
	if in.BucketRef != nil {
		id := &storagev1beta1.StorageBucketIdentity{}
		if err := id.FromExternal(in.BucketRef.External); err != nil {
			mapCtx.Errorf("gcsNotebookSource.bucketRef: %v", err)
		} else {
			bucket := id.Bucket
			object := direct.ValueOf(in.Object)
			out.Uri = "gs://" + bucket + "/" + object
		}
	}
	out.Generation = direct.ValueOf(in.Generation)
	return out
}

func EncryptionSpec_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionSpec{}
	if in.GetKmsKeyName() != "" {
		out.KmsKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
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
