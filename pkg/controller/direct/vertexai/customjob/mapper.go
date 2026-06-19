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

package customjob

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VertexAIDataLabelingJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataLabelingJob) *krm.VertexAIDataLabelingJobObservedState {
	return nil
}

func VertexAIDataLabelingJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIDataLabelingJobObservedState) *pb.DataLabelingJob {
	return nil
}

func VertexAIDataLabelingJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataLabelingJob) *krm.VertexAIDataLabelingJobSpec {
	return nil
}

func VertexAIDataLabelingJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIDataLabelingJobSpec) *pb.DataLabelingJob {
	return nil
}

// CustomJob manual mappers to handle name casing differences and nested structures

func VertexAICustomJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomJob) *krm.VertexAICustomJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAICustomJobObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Error = direct.Status_FromProto(mapCtx, in.GetError())
	out.WebAccessURIs = in.GetWebAccessUris()
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}

func VertexAICustomJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAICustomJobObservedState) *pb.CustomJob {
	if in == nil {
		return nil
	}
	out := &pb.CustomJob{}
	out.State = direct.Enum_ToProto[pb.JobState](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Error = direct.Status_ToProto(mapCtx, in.Error)
	out.WebAccessUris = in.WebAccessURIs
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}

func CustomJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomJobSpec) *krm.CustomJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.CustomJobSpec{}
	if in.GetPersistentResourceId() != "" {
		out.PersistentResourceRef = &refsv1beta1.VertexAIPersistentResourceRef{External: in.GetPersistentResourceId()}
	}
	out.CustomJobWorkerPoolSpecs = direct.Slice_FromProto(mapCtx, in.GetWorkerPoolSpecs(), CustomJobWorkerPoolSpec_FromProto)
	out.CustomJobScheduling = CustomJobScheduling_FromProto(mapCtx, in.GetScheduling())
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	if in.GetNetwork() != "" {
		out.NetworkRef = &refsv1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.ReservedIPRanges = in.ReservedIpRanges
	out.PscInterfaceConfig = PSCInterfaceConfig_FromProto(mapCtx, in.GetPscInterfaceConfig())
	out.BaseOutputDirectory = CustomJobGcsDestination_FromProto(mapCtx, in.GetBaseOutputDirectory())
	out.ProtectedArtifactLocationID = direct.LazyPtr(in.GetProtectedArtifactLocationId())
	if in.GetTensorboard() != "" {
		out.TensorboardRef = &refsv1beta1.VertexAITensorboardRef{External: in.GetTensorboard()}
	}
	out.EnableWebAccess = direct.LazyPtr(in.GetEnableWebAccess())
	out.EnableDashboardAccess = direct.LazyPtr(in.GetEnableDashboardAccess())
	out.Experiment = direct.LazyPtr(in.GetExperiment())
	out.ExperimentRun = direct.LazyPtr(in.GetExperimentRun())
	out.Models = in.Models
	return out
}

func CustomJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.CustomJobSpec) *pb.CustomJobSpec {
	if in == nil {
		return nil
	}
	out := &pb.CustomJobSpec{}
	if in.PersistentResourceRef != nil {
		out.PersistentResourceId = in.PersistentResourceRef.External
	}
	out.WorkerPoolSpecs = direct.Slice_ToProto(mapCtx, in.CustomJobWorkerPoolSpecs, CustomJobWorkerPoolSpec_ToProto)
	out.Scheduling = CustomJobScheduling_ToProto(mapCtx, in.CustomJobScheduling)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	out.ReservedIpRanges = in.ReservedIPRanges
	out.PscInterfaceConfig = PSCInterfaceConfig_ToProto(mapCtx, in.PscInterfaceConfig)
	out.BaseOutputDirectory = CustomJobGcsDestination_ToProto(mapCtx, in.BaseOutputDirectory)
	out.ProtectedArtifactLocationId = direct.ValueOf(in.ProtectedArtifactLocationID)
	if in.TensorboardRef != nil {
		out.Tensorboard = in.TensorboardRef.External
	}
	out.EnableWebAccess = direct.ValueOf(in.EnableWebAccess)
	out.EnableDashboardAccess = direct.ValueOf(in.EnableDashboardAccess)
	out.Experiment = direct.ValueOf(in.Experiment)
	out.ExperimentRun = direct.ValueOf(in.ExperimentRun)
	out.Models = in.Models
	return out
}

func CustomJobWorkerPoolSpec_FromProto(mapCtx *direct.MapContext, in *pb.WorkerPoolSpec) *krm.CustomJobWorkerPoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.CustomJobWorkerPoolSpec{}
	out.CustomJobContainerSpec = CustomJobContainerSpec_FromProto(mapCtx, in.GetContainerSpec())
	out.CustomJobPythonPackageSpec = CustomJobPythonPackageSpec_FromProto(mapCtx, in.GetPythonPackageSpec())
	out.CustomJobMachineSpec = CustomJobMachineSpec_FromProto(mapCtx, in.GetMachineSpec())
	out.ReplicaCount = direct.LazyPtr(in.GetReplicaCount())
	out.CustomJobNfsMounts = direct.Slice_FromProto(mapCtx, in.GetNfsMounts(), CustomJobNfsMount_FromProto)
	out.CustomJobDiskSpec = CustomJobDiskSpec_FromProto(mapCtx, in.GetDiskSpec())
	return out
}

func CustomJobWorkerPoolSpec_ToProto(mapCtx *direct.MapContext, in *krm.CustomJobWorkerPoolSpec) *pb.WorkerPoolSpec {
	if in == nil {
		return nil
	}
	out := &pb.WorkerPoolSpec{}
	if in.CustomJobContainerSpec != nil {
		out.Task = &pb.WorkerPoolSpec_ContainerSpec{
			ContainerSpec: CustomJobContainerSpec_ToProto(mapCtx, in.CustomJobContainerSpec),
		}
	}
	if in.CustomJobPythonPackageSpec != nil {
		out.Task = &pb.WorkerPoolSpec_PythonPackageSpec{
			PythonPackageSpec: CustomJobPythonPackageSpec_ToProto(mapCtx, in.CustomJobPythonPackageSpec),
		}
	}
	out.MachineSpec = CustomJobMachineSpec_ToProto(mapCtx, in.CustomJobMachineSpec)
	out.ReplicaCount = direct.ValueOf(in.ReplicaCount)
	out.NfsMounts = direct.Slice_ToProto(mapCtx, in.CustomJobNfsMounts, CustomJobNfsMount_ToProto)
	out.DiskSpec = CustomJobDiskSpec_ToProto(mapCtx, in.CustomJobDiskSpec)
	return out
}

func CustomJobPythonPackageSpec_FromProto(mapCtx *direct.MapContext, in *pb.PythonPackageSpec) *krm.CustomJobPythonPackageSpec {
	if in == nil {
		return nil
	}
	out := &krm.CustomJobPythonPackageSpec{}
	out.ExecutorImageURI = direct.LazyPtr(in.GetExecutorImageUri())
	out.PackageURIs = in.PackageUris
	out.PythonModule = direct.LazyPtr(in.GetPythonModule())
	out.Args = in.Args
	out.Env = direct.Slice_FromProto(mapCtx, in.Env, CustomJobEnvVar_FromProto)
	return out
}

func CustomJobPythonPackageSpec_ToProto(mapCtx *direct.MapContext, in *krm.CustomJobPythonPackageSpec) *pb.PythonPackageSpec {
	if in == nil {
		return nil
	}
	out := &pb.PythonPackageSpec{}
	out.ExecutorImageUri = direct.ValueOf(in.ExecutorImageURI)
	out.PackageUris = in.PackageURIs
	out.PythonModule = direct.ValueOf(in.PythonModule)
	out.Args = in.Args
	out.Env = direct.Slice_ToProto(mapCtx, in.Env, CustomJobEnvVar_ToProto)
	return out
}

func CustomJobEncryptionSpec_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.CustomJobEncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.CustomJobEncryptionSpec{}
	if in.KmsKeyName != "" {
		out.KMSKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: in.KmsKeyName}
	}
	return out
}

func CustomJobEncryptionSpec_ToProto(mapCtx *direct.MapContext, in *krm.CustomJobEncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	if in.KMSKeyRef != nil {
		out.KmsKeyName = in.KMSKeyRef.External
	}
	return out
}
