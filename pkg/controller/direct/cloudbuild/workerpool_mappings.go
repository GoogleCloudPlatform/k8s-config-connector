/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudbuild

import (
	pb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

func CloudBuildWorkerPoolObservedState_FromProto(mapCtx *MapContext, in *pb.WorkerPool) *krm.CloudBuildWorkerPoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudBuildWorkerPoolObservedState{}
	out.ETag = LazyPtr(in.Etag)
	out.CreateTime = ToOpenAPIDateTime(in.GetCreateTime())
	out.UpdateTime = ToOpenAPIDateTime(in.GetUpdateTime())

	privateConfig := in.GetPrivatePoolV1Config()
	if privateConfig != nil {
		// privateConfig := PrivatePoolV1ConfigStatus_FromProto(mapCtx, in.GetPrivatePoolV1Config())
		out.WorkerConfig = PrivatePoolV1Config_WorkerConfig_FromProto(mapCtx, privateConfig.GetWorkerConfig())
		out.NetworkConfig = PrivatePoolV1Config_NetworkConfigStatus_FromProto(mapCtx, privateConfig.GetNetworkConfig())
	}
	return out
}

func CloudBuildWorkerPoolSpec_ToProto(mapCtx *MapContext, in *krm.CloudBuildWorkerPoolSpec) *pb.WorkerPool {
	if in == nil {
		return nil
	}
	out := &pb.WorkerPool{}
	out.DisplayName = in.DisplayName
	out.Config = &pb.WorkerPool_PrivatePoolV1Config{
		PrivatePoolV1Config: PrivatePoolV1Config_ToProto(mapCtx, in.PrivatePoolConfig),
	}
	return out
}

func PrivatePoolV1Config_NetworkConfigStatus_FromProto(mapCtx *MapContext, in *pb.PrivatePoolV1Config_NetworkConfig) *krm.PrivatePoolV1Config_NetworkConfigStatus {
	if in == nil {
		return nil
	}
	out := &krm.PrivatePoolV1Config_NetworkConfigStatus{}
	out.PeeredNetwork = LazyPtr(in.GetPeeredNetwork())
	out.EgressOption = Enum_FromProto(mapCtx, in.EgressOption)
	out.PeeredNetworkIPRange = LazyPtr(in.GetPeeredNetworkIpRange())
	return out
}

func PrivatePoolV1Config_FromProto(mapCtx *MapContext, in *pb.PrivatePoolV1Config) *krm.PrivatePoolV1Config {
	if in == nil {
		return nil
	}
	out := &krm.PrivatePoolV1Config{}
	out.WorkerConfig = PrivatePoolV1Config_WorkerConfig_FromProto(mapCtx, in.GetWorkerConfig())
	out.NetworkConfig = PrivatePoolV1Config_NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	return out
}
func PrivatePoolV1Config_ToProto(mapCtx *MapContext, in *krm.PrivatePoolV1Config) *pb.PrivatePoolV1Config {
	if in == nil {
		return nil
	}
	out := &pb.PrivatePoolV1Config{}
	out.WorkerConfig = PrivatePoolV1Config_WorkerConfig_ToProto(mapCtx, in.WorkerConfig)
	out.NetworkConfig = PrivatePoolV1Config_NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	return out
}
func PrivatePoolV1Config_NetworkConfig_FromProto(mapCtx *MapContext, in *pb.PrivatePoolV1Config_NetworkConfig) *krm.PrivatePoolV1Config_NetworkConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.PrivatePoolV1Config_NetworkConfigSpec{}
	out.PeeredNetworkRef = refv1beta1.ComputeNetworkRef{
		External: in.GetPeeredNetwork(),
	}
	out.EgressOption = Enum_FromProto(mapCtx, in.EgressOption)
	out.PeeredNetworkIPRange = LazyPtr(in.GetPeeredNetworkIpRange())
	return out
}
func PrivatePoolV1Config_NetworkConfig_ToProto(mapCtx *MapContext, in *krm.PrivatePoolV1Config_NetworkConfigSpec) *pb.PrivatePoolV1Config_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrivatePoolV1Config_NetworkConfig{}
	out.PeeredNetwork = in.PeeredNetworkRef.External
	out.EgressOption = Enum_ToProto[pb.PrivatePoolV1Config_NetworkConfig_EgressOption](mapCtx, in.EgressOption)
	out.PeeredNetworkIpRange = ValueOf(in.PeeredNetworkIPRange)
	return out
}
func PrivatePoolV1Config_WorkerConfig_FromProto(mapCtx *MapContext, in *pb.PrivatePoolV1Config_WorkerConfig) *krm.PrivatePoolV1Config_WorkerConfig {
	if in == nil {
		return nil
	}
	out := &krm.PrivatePoolV1Config_WorkerConfig{}
	out.MachineType = LazyPtr(in.GetMachineType())
	out.DiskSizeGb = LazyPtr(in.GetDiskSizeGb())
	return out
}
func PrivatePoolV1Config_WorkerConfig_ToProto(mapCtx *MapContext, in *krm.PrivatePoolV1Config_WorkerConfig) *pb.PrivatePoolV1Config_WorkerConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrivatePoolV1Config_WorkerConfig{}
	out.MachineType = ValueOf(in.MachineType)
	out.DiskSizeGb = ValueOf(in.DiskSizeGb)
	return out
}
