// Copyright 2024 Google LLC
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

package alloydb

import (
	"fmt"

	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AlloyDBInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.AlloyDBInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBInstanceSpec{}
	out.Annotations = in.GetAnnotations()
	//out.AvailabilityType = direct.LazyPtr(in.GetAvailabilityType().String())
	out.AvailabilityType = direct.Enum_FromProto(mapCtx, in.GetAvailabilityType())
	out.DatabaseFlags = in.GetDatabaseFlags()
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.GceZone = direct.LazyPtr(in.GetGceZone())
	//out.InstanceType = direct.LazyPtr(in.GetInstanceType().String())
	out.InstanceType = direct.Enum_FromProto(mapCtx, in.GetInstanceType())
	// how to handle the labels?
	out.MachineConfig = Instance_MachineConfig_FromProto(mapCtx, in.GetMachineConfig())
	out.NetworkConfig = Instance_InstanceNetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	out.ReadPoolConfig = Instance_ReadPoolConfig_FromProto(mapCtx, in.GetReadPoolConfig())

	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels???
	// MISSING: State
	// MISSING: WritableNode
	// MISSING: Nodes
	// MISSING: QueryInsightsConfig
	// MISSING: ObservabilityConfig
	// MISSING: IpAddress
	// MISSING: PublicIpAddress
	// MISSING: Reconciling
	// MISSING: Etag
	// MISSING: UpdatePolicy
	// MISSING: ClientConnectionConfig
	// MISSING: SatisfiesPzs
	// MISSING: PscInstanceConfig
	// MISSING: GeminiConfig
	// MISSING: OutboundPublicIpAddresses
	fmt.Printf("maqiuyu...AlloyDBInstanceSpec_FromProto: %+v\n", out)
	return out
}

func AlloyDBInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDBInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Annotations = in.Annotations
	out.AvailabilityType = direct.Enum_ToProto[pb.Instance_AvailabilityType](mapCtx, in.AvailabilityType)
	out.DatabaseFlags = in.DatabaseFlags
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.GceZone = direct.ValueOf(in.GceZone)
	out.InstanceType = direct.Enum_ToProto[pb.Instance_InstanceType](mapCtx, in.InstanceType)
	// how to handle the labels?
	out.MachineConfig = Instance_MachineConfig_ToProto(mapCtx, in.MachineConfig)
	out.NetworkConfig = Instance_InstanceNetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	out.ReadPoolConfig = Instance_ReadPoolConfig_ToProto(mapCtx, in.ReadPoolConfig)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels should be an internal field for to map metadata.labels
	// MISSING: State
	// MISSING: WritableNode
	// MISSING: Nodes
	// MISSING: QueryInsightsConfig
	// MISSING: ObservabilityConfig
	// MISSING: IpAddress
	// MISSING: PublicIpAddress
	// MISSING: Reconciling
	// MISSING: Etag
	// MISSING: UpdatePolicy
	// MISSING: ClientConnectionConfig
	// MISSING: SatisfiesPzs
	// MISSING: PscInstanceConfig
	// MISSING: GeminiConfig
	// MISSING: OutboundPublicIpAddresses
	fmt.Printf("maqiuyu...AlloyDBInstanceSpec_ToProto: %+v\n", out)
	return out
}

func AlloyDBInstanceStatus_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.AlloyDBInstanceStatus {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBInstanceStatus{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.IpAddress = direct.LazyPtr(in.GetIpAddress())
	out.Name = direct.LazyPtr(in.GetName())
	out.OutboundPublicIpAddresses = in.GetOutboundPublicIpAddresses()
	out.PublicIpAddress = direct.LazyPtr(in.GetPublicIpAddress())
	out.Reconciling = direct.LazyPtr(in.Reconciling)
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Uid = direct.LazyPtr(in.Uid)
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())

	fmt.Printf("maqiuyu...AlloyDBInstanceStatus_FromProto: %+v", out)
	return out
}
