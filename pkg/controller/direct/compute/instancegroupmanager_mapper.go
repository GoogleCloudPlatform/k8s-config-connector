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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeInstanceGroupManagerSpec_TargetPools_FromProto(mapCtx *direct.MapContext, in []string) []computev1beta1.InstanceResourceRef {
	if in == nil {
		return nil
	}
	var out []computev1beta1.InstanceResourceRef
	for _, i := range in {
		out = append(out, computev1beta1.InstanceResourceRef{
			External: i,
		})
	}
	return out
}

func ComputeInstanceGroupManagerSpec_TargetPools_ToProto(mapCtx *direct.MapContext, in []computev1beta1.InstanceResourceRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, i := range in {
		if i.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", i.Name)
		}
		out = append(out, i.External)
	}
	return out
}

func StatefulPolicyPreservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.StatefulPolicyPreservedState) *computev1beta1.StatefulPolicyPreservedState {
	if in == nil {
		return nil
	}
	out := &computev1beta1.StatefulPolicyPreservedState{}
	if in.Disks != nil {
		out.Disks = make(map[string]computev1beta1.StatefulPolicyPreservedStateDiskDevice)
		for k, v := range in.Disks {
			if v != nil {
				out.Disks[k] = computev1beta1.StatefulPolicyPreservedStateDiskDevice{
					AutoDelete: v.AutoDelete,
				}
			}
		}
	}
	if in.ExternalIPs != nil {
		out.ExternalIPs = make(map[string]computev1beta1.StatefulPolicyPreservedStateNetworkIP)
		for k, v := range in.ExternalIPs {
			if v != nil {
				out.ExternalIPs[k] = computev1beta1.StatefulPolicyPreservedStateNetworkIP{
					AutoDelete: v.AutoDelete,
				}
			}
		}
	}
	if in.InternalIPs != nil {
		out.InternalIPs = make(map[string]computev1beta1.StatefulPolicyPreservedStateNetworkIP)
		for k, v := range in.InternalIPs {
			if v != nil {
				out.InternalIPs[k] = computev1beta1.StatefulPolicyPreservedStateNetworkIP{
					AutoDelete: v.AutoDelete,
				}
			}
		}
	}
	return out
}

func StatefulPolicyPreservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *computev1beta1.StatefulPolicyPreservedState) *pb.StatefulPolicyPreservedState {
	if in == nil {
		return nil
	}
	out := &pb.StatefulPolicyPreservedState{}
	if in.Disks != nil {
		out.Disks = make(map[string]*pb.StatefulPolicyPreservedStateDiskDevice)
		for k, v := range in.Disks {
			out.Disks[k] = &pb.StatefulPolicyPreservedStateDiskDevice{
				AutoDelete: v.AutoDelete,
			}
		}
	}
	if in.ExternalIPs != nil {
		out.ExternalIPs = make(map[string]*pb.StatefulPolicyPreservedStateNetworkIp)
		for k, v := range in.ExternalIPs {
			out.ExternalIPs[k] = &pb.StatefulPolicyPreservedStateNetworkIp{
				AutoDelete: v.AutoDelete,
			}
		}
	}
	if in.InternalIPs != nil {
		out.InternalIPs = make(map[string]*pb.StatefulPolicyPreservedStateNetworkIp)
		for k, v := range in.InternalIPs {
			out.InternalIPs[k] = &pb.StatefulPolicyPreservedStateNetworkIp{
				AutoDelete: v.AutoDelete,
			}
		}
	}
	return out
}
