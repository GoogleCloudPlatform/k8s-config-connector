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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputePerInstanceConfigSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputePerInstanceConfigSpec) *pb.PerInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.PerInstanceConfig{}
	out.Name = in.ResourceID
	out.PreservedState = PerinstanceconfigPreservedState_v1alpha1_ToProto(mapCtx, in.PreservedState)
	return out
}

func ComputePerInstanceConfigSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.PerInstanceConfig) *krm.ComputePerInstanceConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputePerInstanceConfigSpec{}
	out.ResourceID = in.Name
	out.PreservedState = PerinstanceconfigPreservedState_v1alpha1_FromProto(mapCtx, in.PreservedState)
	return out
}

func ComputePerInstanceConfigStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputePerInstanceConfigStatus) *pb.PerInstanceConfig {
	return nil
}

func ComputePerInstanceConfigStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.PerInstanceConfig) *krm.ComputePerInstanceConfigStatus {
	return nil
}

func PerinstanceconfigPreservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.PerinstanceconfigPreservedState) *pb.PreservedState {
	if in == nil {
		return nil
	}
	out := &pb.PreservedState{}

	if len(in.Disk) > 0 {
		out.Disks = make(map[string]*pb.PreservedStatePreservedDisk)
		for _, d := range in.Disk {
			out.Disks[d.DeviceName] = PerinstanceconfigDisk_v1alpha1_ToProto(mapCtx, &d)
		}
	}

	if len(in.ExternalIp) > 0 {
		out.ExternalIPs = make(map[string]*pb.PreservedStatePreservedNetworkIp)
		for _, ip := range in.ExternalIp {
			out.ExternalIPs[ip.InterfaceName] = PerinstanceconfigExternalIp_v1alpha1_ToProto(mapCtx, &ip)
		}
	}

	if len(in.InternalIp) > 0 {
		out.InternalIPs = make(map[string]*pb.PreservedStatePreservedNetworkIp)
		for _, ip := range in.InternalIp {
			out.InternalIPs[ip.InterfaceName] = PerinstanceconfigInternalIp_v1alpha1_ToProto(mapCtx, &ip)
		}
	}

	if len(in.Metadata) > 0 {
		out.Metadata = in.Metadata
	}

	return out
}

func PerinstanceconfigPreservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.PreservedState) *krm.PerinstanceconfigPreservedState {
	if in == nil {
		return nil
	}
	out := &krm.PerinstanceconfigPreservedState{}

	if len(in.Disks) > 0 {
		out.Disk = []krm.PerinstanceconfigDisk{}
		for k, v := range in.Disks {
			d := PerinstanceconfigDisk_v1alpha1_FromProto(mapCtx, v)
			if d != nil {
				d.DeviceName = k
				out.Disk = append(out.Disk, *d)
			}
		}
	}

	if len(in.ExternalIPs) > 0 {
		out.ExternalIp = []krm.PerinstanceconfigExternalIp{}
		for k, v := range in.ExternalIPs {
			ip := PerinstanceconfigExternalIp_v1alpha1_FromProto(mapCtx, v)
			if ip != nil {
				ip.InterfaceName = k
				out.ExternalIp = append(out.ExternalIp, *ip)
			}
		}
	}

	if len(in.InternalIPs) > 0 {
		out.InternalIp = []krm.PerinstanceconfigInternalIp{}
		for k, v := range in.InternalIPs {
			ip := PerinstanceconfigInternalIp_v1alpha1_FromProto(mapCtx, v)
			if ip != nil {
				ip.InterfaceName = k
				out.InternalIp = append(out.InternalIp, *ip)
			}
		}
	}

	if len(in.Metadata) > 0 {
		out.Metadata = in.Metadata
	}

	return out
}

func PerinstanceconfigDisk_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.PerinstanceconfigDisk) *pb.PreservedStatePreservedDisk {
	if in == nil {
		return nil
	}
	out := &pb.PreservedStatePreservedDisk{}
	out.AutoDelete = in.DeleteRule
	out.Mode = in.Mode
	if in.Source != "" {
		out.Source = &in.Source
	}
	return out
}

func PerinstanceconfigDisk_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.PreservedStatePreservedDisk) *krm.PerinstanceconfigDisk {
	if in == nil {
		return nil
	}
	out := &krm.PerinstanceconfigDisk{}
	out.DeleteRule = in.AutoDelete
	out.Mode = in.Mode
	if in.Source != nil {
		out.Source = *in.Source
	}
	return out
}

func PerinstanceconfigExternalIp_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.PerinstanceconfigExternalIp) *pb.PreservedStatePreservedNetworkIp {
	if in == nil {
		return nil
	}
	out := &pb.PreservedStatePreservedNetworkIp{}
	out.AutoDelete = in.AutoDelete
	out.IpAddress = PerinstanceconfigIpAddress_v1alpha1_ToProto(mapCtx, in.IpAddress)
	return out
}

func PerinstanceconfigExternalIp_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.PreservedStatePreservedNetworkIp) *krm.PerinstanceconfigExternalIp {
	if in == nil {
		return nil
	}
	out := &krm.PerinstanceconfigExternalIp{}
	out.AutoDelete = in.AutoDelete
	out.IpAddress = PerinstanceconfigIpAddress_v1alpha1_FromProto(mapCtx, in.IpAddress)
	return out
}

func PerinstanceconfigInternalIp_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.PerinstanceconfigInternalIp) *pb.PreservedStatePreservedNetworkIp {
	if in == nil {
		return nil
	}
	out := &pb.PreservedStatePreservedNetworkIp{}
	out.AutoDelete = in.AutoDelete
	out.IpAddress = PerinstanceconfigIpAddress_v1alpha1_ToProto(mapCtx, in.IpAddress)
	return out
}

func PerinstanceconfigInternalIp_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.PreservedStatePreservedNetworkIp) *krm.PerinstanceconfigInternalIp {
	if in == nil {
		return nil
	}
	out := &krm.PerinstanceconfigInternalIp{}
	out.AutoDelete = in.AutoDelete
	out.IpAddress = PerinstanceconfigIpAddress_v1alpha1_FromProto(mapCtx, in.IpAddress)
	return out
}

func PerinstanceconfigIpAddress_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.PerinstanceconfigIpAddress) *pb.PreservedStatePreservedNetworkIpIpAddress {
	if in == nil {
		return nil
	}
	out := &pb.PreservedStatePreservedNetworkIpIpAddress{}
	out.Address = in.Address
	return out
}

func PerinstanceconfigIpAddress_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.PreservedStatePreservedNetworkIpIpAddress) *krm.PerinstanceconfigIpAddress {
	if in == nil {
		return nil
	}
	out := &krm.PerinstanceconfigIpAddress{}
	out.Address = in.Address
	return out
}
