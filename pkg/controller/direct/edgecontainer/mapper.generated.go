// Copyright 2025 Google LLC
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

// +generated:mapper
// proto.service: google.cloud.edgecontainer.v1
// krm.group: edgecontainer.cnrm.cloud.google.com
// krm.version: v1alpha1

package edgecontainer

import (
	pb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/edgecontainer/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func EdgeContainerMachineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Machine) *krm.EdgeContainerMachineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EdgeContainerMachineObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.HostedNode = direct.LazyPtr(in.GetHostedNode())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}
func EdgeContainerMachineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EdgeContainerMachineObservedState) *pb.Machine {
	if in == nil {
		return nil
	}
	out := &pb.Machine{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.HostedNode = direct.ValueOf(in.HostedNode)
	out.Version = direct.ValueOf(in.Version)
	out.Disabled = direct.ValueOf(in.Disabled)
	return out
}
func EdgeContainerMachineSpec_FromProto(mapCtx *direct.MapContext, in *pb.Machine) *krm.EdgeContainerMachineSpec {
	if in == nil {
		return nil
	}
	out := &krm.EdgeContainerMachineSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	out.Zone = direct.LazyPtr(in.GetZone())
	return out
}
func EdgeContainerMachineSpec_ToProto(mapCtx *direct.MapContext, in *krm.EdgeContainerMachineSpec) *pb.Machine {
	if in == nil {
		return nil
	}
	out := &pb.Machine{}
	// MISSING: Name
	out.Labels = in.Labels
	out.Zone = direct.ValueOf(in.Zone)
	return out
}
