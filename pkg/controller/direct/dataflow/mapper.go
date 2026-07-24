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

package dataflow

import (
	"encoding/json"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"

	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"

	pb "cloud.google.com/go/dataflow/apiv1beta3/dataflowpb"
	krmcomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataflow/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"k8s.io/apimachinery/pkg/runtime"
)

func DataflowJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.DataflowJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataflowJobSpec{}

	out.TransformNameMapping = mapToRawExtension(in.GetTransformNameMapping())

	env := in.GetEnvironment()
	if env != nil {
		out.TempGcsLocation = env.GetTempStoragePrefix()
		if env.GetServiceKmsKeyName() != "" {
			out.KmsKeyRef = &kmsv1beta1.KMSCryptoKeyRef{External: env.GetServiceKmsKeyName()}
		}
		if len(env.GetExperiments()) > 0 {
			out.AdditionalExperiments = env.GetExperiments()
		}

		pools := env.GetWorkerPools()
		if len(pools) > 0 {
			wp := pools[0]
			if wp.GetMachineType() != "" {
				out.MachineType = direct.LazyPtr(wp.GetMachineType())
			}
			if wp.GetAutoscalingSettings() != nil && wp.GetAutoscalingSettings().GetMaxNumWorkers() != 0 {
				out.MaxWorkers = direct.LazyPtr(int64(wp.GetAutoscalingSettings().GetMaxNumWorkers()))
			}
			if wp.GetNetwork() != "" {
				out.NetworkRef = &computerefs.ComputeNetworkRef{External: wp.GetNetwork()}
			}
			if wp.GetSubnetwork() != "" {
				out.SubnetworkRef = &krmcomputev1beta1.ComputeSubnetworkRef{External: wp.GetSubnetwork()}
			}
			if wp.GetIpConfiguration() != pb.WorkerIPAddressConfiguration_WORKER_IP_UNSPECIFIED {
				out.IpConfiguration = direct.Enum_FromProto(mapCtx, wp.GetIpConfiguration())
			}
			if wp.GetZone() != "" {
				out.Zone = direct.LazyPtr(wp.GetZone())
			}
		}
	}

	if in.GetLocation() != "" {
		out.Region = direct.LazyPtr(in.GetLocation())
	}
	if in.GetName() != "" {
		out.ResourceID = direct.LazyPtr(in.GetName())
	}

	return out
}

func DataflowJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataflowJobSpec) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}

	out.TransformNameMapping = rawExtensionToMap(in.TransformNameMapping)

	env := &pb.Environment{}
	hasEnv := false

	if in.TempGcsLocation != "" {
		env.TempStoragePrefix = in.TempGcsLocation
		hasEnv = true
	}
	if in.KmsKeyRef != nil {
		env.ServiceKmsKeyName = in.KmsKeyRef.External
		hasEnv = true
	}
	if len(in.AdditionalExperiments) > 0 {
		env.Experiments = in.AdditionalExperiments
		hasEnv = true
	}

	wp := &pb.WorkerPool{}
	hasWorkerPool := false

	if in.MachineType != nil {
		wp.MachineType = *in.MachineType
		hasWorkerPool = true
	}
	if in.MaxWorkers != nil {
		wp.AutoscalingSettings = &pb.AutoscalingSettings{
			MaxNumWorkers: int32(*in.MaxWorkers),
		}
		hasWorkerPool = true
	}
	if in.NetworkRef != nil {
		wp.Network = in.NetworkRef.External
		hasWorkerPool = true
	}
	if in.SubnetworkRef != nil {
		wp.Subnetwork = in.SubnetworkRef.External
		hasWorkerPool = true
	}
	if in.IpConfiguration != nil {
		wp.IpConfiguration = direct.Enum_ToProto[pb.WorkerIPAddressConfiguration](mapCtx, in.IpConfiguration)
		hasWorkerPool = true
	}
	if in.Zone != nil {
		wp.Zone = *in.Zone
		hasWorkerPool = true
	}

	if hasWorkerPool {
		env.WorkerPools = []*pb.WorkerPool{wp}
		hasEnv = true
	}

	if hasEnv {
		out.Environment = env
	}

	if in.Region != nil {
		out.Location = *in.Region
	}
	if in.ResourceID != nil {
		out.Name = *in.ResourceID
	}

	return out
}

func DataflowJobStatus_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.DataflowJobStatus {
	if in == nil {
		return nil
	}
	out := &krm.DataflowJobStatus{}

	if in.GetId() != "" {
		out.JobId = direct.LazyPtr(in.GetId())
	}
	if in.GetType() != pb.JobType_JOB_TYPE_UNKNOWN {
		out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	}
	if in.GetCurrentState() != pb.JobState_JOB_STATE_UNKNOWN {
		out.State = direct.Enum_FromProto(mapCtx, in.GetCurrentState())
	}

	return out
}

func DataflowJobStatus_ToProto(mapCtx *direct.MapContext, in *krm.DataflowJobStatus) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}

	if in.JobId != nil {
		out.Id = *in.JobId
	}
	if in.Type != nil {
		out.Type = direct.Enum_ToProto[pb.JobType](mapCtx, in.Type)
	}
	if in.State != nil {
		out.CurrentState = direct.Enum_ToProto[pb.JobState](mapCtx, in.State)
	}

	return out
}

func mapToRawExtension(in map[string]string) *runtime.RawExtension {
	if in == nil {
		return nil
	}
	raw, err := json.Marshal(in)
	if err != nil {
		return nil
	}
	return &runtime.RawExtension{Raw: raw}
}

func rawExtensionToMap(in *runtime.RawExtension) map[string]string {
	if in == nil || in.Raw == nil {
		return nil
	}
	var out map[string]string
	if err := json.Unmarshal(in.Raw, &out); err != nil {
		return nil
	}
	return out
}
