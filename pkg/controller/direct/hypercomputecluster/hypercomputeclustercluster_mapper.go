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

package hypercomputecluster

import (
	pb "cloud.google.com/go/hypercomputecluster/apiv1/hypercomputeclusterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/hypercomputecluster/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkResources_FromProto(mapCtx *direct.MapContext, in map[string]*pb.NetworkResource) map[string]*krm.NetworkResource {
	if in == nil {
		return nil
	}
	out := make(map[string]*krm.NetworkResource)
	for k, v := range in {
		out[k] = NetworkResource_FromProto(mapCtx, v)
	}
	return out
}

func NetworkResources_ToProto(mapCtx *direct.MapContext, in map[string]*krm.NetworkResource) map[string]*pb.NetworkResource {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.NetworkResource)
	for k, v := range in {
		out[k] = NetworkResource_ToProto(mapCtx, v)
	}
	return out
}

func StorageResources_FromProto(mapCtx *direct.MapContext, in map[string]*pb.StorageResource) map[string]*krm.StorageResource {
	if in == nil {
		return nil
	}
	out := make(map[string]*krm.StorageResource)
	for k, v := range in {
		out[k] = StorageResource_FromProto(mapCtx, v)
	}
	return out
}

func StorageResources_ToProto(mapCtx *direct.MapContext, in map[string]*krm.StorageResource) map[string]*pb.StorageResource {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.StorageResource)
	for k, v := range in {
		out[k] = StorageResource_ToProto(mapCtx, v)
	}
	return out
}

func ComputeResources_FromProto(mapCtx *direct.MapContext, in map[string]*pb.ComputeResource) map[string]*krm.ComputeResource {
	if in == nil {
		return nil
	}
	out := make(map[string]*krm.ComputeResource)
	for k, v := range in {
		out[k] = ComputeResource_FromProto(mapCtx, v)
	}
	return out
}

func ComputeResources_ToProto(mapCtx *direct.MapContext, in map[string]*krm.ComputeResource) map[string]*pb.ComputeResource {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.ComputeResource)
	for k, v := range in {
		out[k] = ComputeResource_ToProto(mapCtx, v)
	}
	return out
}

func SlurmLoginNodes_FromProto(mapCtx *direct.MapContext, in *pb.SlurmLoginNodes) *krm.SlurmLoginNodes {
	if in == nil {
		return nil
	}
	out := &krm.SlurmLoginNodes{}
	out.Count = direct.LazyPtr(in.GetCount())
	out.Zone = direct.LazyPtr(in.GetZone())
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.StartupScript = direct.LazyPtr(in.GetStartupScript())
	out.EnableOSLogin = direct.LazyPtr(in.GetEnableOsLogin())
	out.EnablePublicIPs = direct.LazyPtr(in.GetEnablePublicIps())
	out.Labels = in.Labels
	out.StorageConfigs = direct.Slice_FromProto(mapCtx, in.StorageConfigs, StorageConfig_FromProto)
	out.BootDisk = BootDisk_FromProto(mapCtx, in.GetBootDisk())
	return out
}

func SlurmLoginNodes_ToProto(mapCtx *direct.MapContext, in *krm.SlurmLoginNodes) *pb.SlurmLoginNodes {
	if in == nil {
		return nil
	}
	out := &pb.SlurmLoginNodes{}
	out.Count = direct.ValueOf(in.Count)
	out.Zone = direct.ValueOf(in.Zone)
	out.MachineType = direct.ValueOf(in.MachineType)
	out.StartupScript = direct.ValueOf(in.StartupScript)
	out.EnableOsLogin = direct.ValueOf(in.EnableOSLogin)
	out.EnablePublicIps = direct.ValueOf(in.EnablePublicIPs)
	out.Labels = in.Labels
	out.StorageConfigs = direct.Slice_ToProto(mapCtx, in.StorageConfigs, StorageConfig_ToProto)
	out.BootDisk = BootDisk_ToProto(mapCtx, in.BootDisk)
	return out
}
