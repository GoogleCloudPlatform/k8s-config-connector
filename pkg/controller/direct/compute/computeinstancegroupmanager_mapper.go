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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func InstancegroupmanagerAutoHealingPolicies_FromProto(in *pb.InstanceGroupManagerAutoHealingPolicy) *krm.InstancegroupmanagerAutoHealingPolicies {
	if in == nil {
		return nil
	}
	out := &krm.InstancegroupmanagerAutoHealingPolicies{}
	if in.GetHealthCheck() != "" {
		out.HealthCheckRef = &krm.ComputeHealthCheckRef{External: in.GetHealthCheck()}
	}
	if in.InitialDelaySec != nil {
		val := int64(*in.InitialDelaySec)
		out.InitialDelaySec = &val
	}
	return out
}

func InstancegroupmanagerAutoHealingPolicies_ToProto(in *krm.InstancegroupmanagerAutoHealingPolicies) *pb.InstanceGroupManagerAutoHealingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerAutoHealingPolicy{}
	if in.HealthCheckRef != nil && in.HealthCheckRef.External != "" {
		out.HealthCheck = &in.HealthCheckRef.External
	}
	if in.InitialDelaySec != nil {
		val := int32(*in.InitialDelaySec)
		out.InitialDelaySec = &val
	}
	return out
}

func InstancegroupmanagerDistributionPolicy_FromProto(in *pb.DistributionPolicy) *krm.InstancegroupmanagerDistributionPolicy {
	if in == nil {
		return nil
	}
	out := &krm.InstancegroupmanagerDistributionPolicy{}
	out.TargetShape = in.TargetShape
	for _, zoneConf := range in.Zones {
		if zoneConf != nil {
			out.Zones = append(out.Zones, krm.InstancegroupmanagerZones{Zone: zoneConf.Zone})
		}
	}
	return out
}

func InstancegroupmanagerDistributionPolicy_ToProto(in *krm.InstancegroupmanagerDistributionPolicy) *pb.DistributionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.DistributionPolicy{}
	out.TargetShape = in.TargetShape
	for _, z := range in.Zones {
		out.Zones = append(out.Zones, &pb.DistributionPolicyZoneConfiguration{Zone: z.Zone})
	}
	return out
}

func InstancegroupmanagerNamedPorts_FromProto(in *pb.NamedPort) *krm.InstancegroupmanagerNamedPorts {
	if in == nil {
		return nil
	}
	out := &krm.InstancegroupmanagerNamedPorts{}
	out.Name = in.Name
	if in.Port != nil {
		val := int64(*in.Port)
		out.Port = &val
	}
	return out
}

func InstancegroupmanagerNamedPorts_ToProto(in *krm.InstancegroupmanagerNamedPorts) *pb.NamedPort {
	if in == nil {
		return nil
	}
	out := &pb.NamedPort{}
	out.Name = in.Name
	if in.Port != nil {
		val := int32(*in.Port)
		out.Port = &val
	}
	return out
}

func InstancegroupmanagerMaxSurge_FromProto(in *pb.FixedOrPercent) *krm.InstancegroupmanagerMaxSurge {
	if in == nil {
		return nil
	}
	out := &krm.InstancegroupmanagerMaxSurge{}
	if in.Fixed != nil {
		val := int64(*in.Fixed)
		out.Fixed = &val
	}
	if in.Percent != nil {
		val := int64(*in.Percent)
		out.Percent = &val
	}
	return out
}

func InstancegroupmanagerMaxSurge_ToProto(in *krm.InstancegroupmanagerMaxSurge) *pb.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &pb.FixedOrPercent{}
	if in.Fixed != nil {
		val := int32(*in.Fixed)
		out.Fixed = &val
	}
	if in.Percent != nil {
		val := int32(*in.Percent)
		out.Percent = &val
	}
	return out
}

func InstancegroupmanagerMaxUnavailable_FromProto(in *pb.FixedOrPercent) *krm.InstancegroupmanagerMaxUnavailable {
	if in == nil {
		return nil
	}
	out := &krm.InstancegroupmanagerMaxUnavailable{}
	if in.Fixed != nil {
		val := int64(*in.Fixed)
		out.Fixed = &val
	}
	if in.Percent != nil {
		val := int64(*in.Percent)
		out.Percent = &val
	}
	return out
}

func InstancegroupmanagerMaxUnavailable_ToProto(in *krm.InstancegroupmanagerMaxUnavailable) *pb.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &pb.FixedOrPercent{}
	if in.Fixed != nil {
		val := int32(*in.Fixed)
		out.Fixed = &val
	}
	if in.Percent != nil {
		val := int32(*in.Percent)
		out.Percent = &val
	}
	return out
}

func InstancegroupmanagerUpdatePolicy_FromProto(in *pb.InstanceGroupManagerUpdatePolicy) *krm.InstancegroupmanagerUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &krm.InstancegroupmanagerUpdatePolicy{}
	out.InstanceRedistributionType = in.InstanceRedistributionType
	out.MaxSurge = InstancegroupmanagerMaxSurge_FromProto(in.MaxSurge)
	out.MaxUnavailable = InstancegroupmanagerMaxUnavailable_FromProto(in.MaxUnavailable)
	out.MinimalAction = in.MinimalAction
	out.ReplacementMethod = in.ReplacementMethod
	out.Type = in.Type
	return out
}

func InstancegroupmanagerUpdatePolicy_ToProto(in *krm.InstancegroupmanagerUpdatePolicy) *pb.InstanceGroupManagerUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerUpdatePolicy{}
	out.InstanceRedistributionType = in.InstanceRedistributionType
	out.MaxSurge = InstancegroupmanagerMaxSurge_ToProto(in.MaxSurge)
	out.MaxUnavailable = InstancegroupmanagerMaxUnavailable_ToProto(in.MaxUnavailable)
	out.MinimalAction = in.MinimalAction
	out.ReplacementMethod = in.ReplacementMethod
	out.Type = in.Type
	return out
}

func InstancegroupmanagerTargetSize_FromProto(in *pb.FixedOrPercent) *krm.InstancegroupmanagerTargetSize {
	if in == nil {
		return nil
	}
	out := &krm.InstancegroupmanagerTargetSize{}
	if in.Calculated != nil {
		val := int64(*in.Calculated)
		out.Calculated = &val
	}
	if in.Fixed != nil {
		val := int64(*in.Fixed)
		out.Fixed = &val
	}
	if in.Percent != nil {
		val := int64(*in.Percent)
		out.Percent = &val
	}
	return out
}

func InstancegroupmanagerTargetSize_ToProto(in *krm.InstancegroupmanagerTargetSize) *pb.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &pb.FixedOrPercent{}
	if in.Calculated != nil {
		val := int32(*in.Calculated)
		out.Calculated = &val
	}
	if in.Fixed != nil {
		val := int32(*in.Fixed)
		out.Fixed = &val
	}
	if in.Percent != nil {
		val := int32(*in.Percent)
		out.Percent = &val
	}
	return out
}

func InstancegroupmanagerVersions_FromProto(in *pb.InstanceGroupManagerVersion) krm.InstancegroupmanagerVersions {
	out := krm.InstancegroupmanagerVersions{}
	if in == nil {
		return out
	}
	if in.GetInstanceTemplate() != "" {
		out.InstanceTemplateRef = &krm.VersionsInstanceTemplateRef{External: in.GetInstanceTemplate()}
	}
	out.Name = in.Name
	out.TargetSize = InstancegroupmanagerTargetSize_FromProto(in.TargetSize)
	return out
}

func InstancegroupmanagerVersions_ToProto(in *krm.InstancegroupmanagerVersions) *pb.InstanceGroupManagerVersion {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerVersion{}
	if in.InstanceTemplateRef != nil && in.InstanceTemplateRef.External != "" {
		out.InstanceTemplate = &in.InstanceTemplateRef.External
	}
	out.Name = in.Name
	out.TargetSize = InstancegroupmanagerTargetSize_ToProto(in.TargetSize)
	return out
}

func ComputeInstanceGroupManagerSpec_TargetPools_FromProto(mapCtx *direct.MapContext, in []string) []krm.ComputeTargetPoolRef {
	var out []krm.ComputeTargetPoolRef
	for _, tp := range in {
		out = append(out, krm.ComputeTargetPoolRef{External: tp})
	}
	return out
}

func ComputeInstanceGroupManagerSpec_TargetPools_ToProto(mapCtx *direct.MapContext, in []krm.ComputeTargetPoolRef) []string {
	var out []string
	for _, tp := range in {
		if tp.External != "" {
			out = append(out, tp.External)
		}
	}
	return out
}

func InstancegroupmanagerStatefulPolicy_FromProto(in *pb.StatefulPolicy) *krm.InstancegroupmanagerStatefulPolicy {
	if in == nil {
		return nil
	}
	out := &krm.InstancegroupmanagerStatefulPolicy{}
	if in.PreservedState != nil {
		out.PreservedState = &krm.InstancegroupmanagerPreservedState{}
		if len(in.PreservedState.Disks) > 0 {
			out.PreservedState.Disks = make(map[string]krm.InstancegroupmanagerDisks)
			for k, v := range in.PreservedState.Disks {
				if v != nil {
					out.PreservedState.Disks[k] = krm.InstancegroupmanagerDisks{AutoDelete: v.AutoDelete}
				}
			}
		}
		if len(in.PreservedState.ExternalIPs) > 0 {
			out.PreservedState.ExternalIps = make(map[string]krm.InstancegroupmanagerExternalIps)
			for k, v := range in.PreservedState.ExternalIPs {
				if v != nil {
					out.PreservedState.ExternalIps[k] = krm.InstancegroupmanagerExternalIps{AutoDelete: v.AutoDelete}
				}
			}
		}
		if len(in.PreservedState.InternalIPs) > 0 {
			out.PreservedState.InternalIps = make(map[string]krm.InstancegroupmanagerInternalIps)
			for k, v := range in.PreservedState.InternalIPs {
				if v != nil {
					out.PreservedState.InternalIps[k] = krm.InstancegroupmanagerInternalIps{AutoDelete: v.AutoDelete}
				}
			}
		}
	}
	return out
}

func InstancegroupmanagerStatefulPolicy_ToProto(in *krm.InstancegroupmanagerStatefulPolicy) *pb.StatefulPolicy {
	if in == nil {
		return nil
	}
	out := &pb.StatefulPolicy{}
	if in.PreservedState != nil {
		out.PreservedState = &pb.StatefulPolicyPreservedState{}
		if len(in.PreservedState.Disks) > 0 {
			out.PreservedState.Disks = make(map[string]*pb.StatefulPolicyPreservedStateDiskDevice)
			for k, v := range in.PreservedState.Disks {
				out.PreservedState.Disks[k] = &pb.StatefulPolicyPreservedStateDiskDevice{AutoDelete: v.AutoDelete}
			}
		}
		if len(in.PreservedState.ExternalIps) > 0 {
			out.PreservedState.ExternalIPs = make(map[string]*pb.StatefulPolicyPreservedStateNetworkIp)
			for k, v := range in.PreservedState.ExternalIps {
				out.PreservedState.ExternalIPs[k] = &pb.StatefulPolicyPreservedStateNetworkIp{AutoDelete: v.AutoDelete}
			}
		}
		if len(in.PreservedState.InternalIps) > 0 {
			out.PreservedState.InternalIPs = make(map[string]*pb.StatefulPolicyPreservedStateNetworkIp)
			for k, v := range in.PreservedState.InternalIps {
				out.PreservedState.InternalIPs[k] = &pb.StatefulPolicyPreservedStateNetworkIp{AutoDelete: v.AutoDelete}
			}
		}
	}
	return out
}

func ComputeInstanceGroupManagerSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManager) *krm.ComputeInstanceGroupManagerSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceGroupManagerSpec{}
	for _, val := range in.AutoHealingPolicies {
		out.AutoHealingPolicies = append(out.AutoHealingPolicies, *InstancegroupmanagerAutoHealingPolicies_FromProto(val))
	}
	out.BaseInstanceName = in.BaseInstanceName
	out.Description = in.Description
	out.DistributionPolicy = InstancegroupmanagerDistributionPolicy_FromProto(in.DistributionPolicy)
	if in.GetInstanceTemplate() != "" {
		out.InstanceTemplateRef = &krm.ComputeInstanceTemplateRef{External: in.GetInstanceTemplate()}
	}
	for _, val := range in.NamedPorts {
		out.NamedPorts = append(out.NamedPorts, *InstancegroupmanagerNamedPorts_FromProto(val))
	}
	out.StatefulPolicy = InstancegroupmanagerStatefulPolicy_FromProto(in.StatefulPolicy)
	out.TargetPools = ComputeInstanceGroupManagerSpec_TargetPools_FromProto(mapCtx, in.TargetPools)
	if in.TargetSize != nil {
		out.TargetSize = int64(*in.TargetSize)
	}
	out.UpdatePolicy = InstancegroupmanagerUpdatePolicy_FromProto(in.UpdatePolicy)
	for _, val := range in.Versions {
		out.Versions = append(out.Versions, InstancegroupmanagerVersions_FromProto(val))
	}
	return out
}

func ComputeInstanceGroupManagerSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceGroupManagerSpec) *pb.InstanceGroupManager {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManager{}
	for _, val := range in.AutoHealingPolicies {
		out.AutoHealingPolicies = append(out.AutoHealingPolicies, InstancegroupmanagerAutoHealingPolicies_ToProto(&val))
	}
	out.BaseInstanceName = in.BaseInstanceName
	out.Description = in.Description
	out.DistributionPolicy = InstancegroupmanagerDistributionPolicy_ToProto(in.DistributionPolicy)
	if in.InstanceTemplateRef != nil && in.InstanceTemplateRef.External != "" {
		out.InstanceTemplate = &in.InstanceTemplateRef.External
	}
	for _, val := range in.NamedPorts {
		out.NamedPorts = append(out.NamedPorts, InstancegroupmanagerNamedPorts_ToProto(&val))
	}
	out.StatefulPolicy = InstancegroupmanagerStatefulPolicy_ToProto(in.StatefulPolicy)
	out.TargetPools = ComputeInstanceGroupManagerSpec_TargetPools_ToProto(mapCtx, in.TargetPools)
	val32 := int32(in.TargetSize)
	out.TargetSize = &val32
	out.UpdatePolicy = InstancegroupmanagerUpdatePolicy_ToProto(in.UpdatePolicy)
	for _, val := range in.Versions {
		out.Versions = append(out.Versions, InstancegroupmanagerVersions_ToProto(&val))
	}
	return out
}

func InstancegroupmanagerCurrentActionsStatus_FromProto(in *pb.InstanceGroupManagerActionsSummary) *krm.InstancegroupmanagerCurrentActionsStatus {
	if in == nil {
		return nil
	}
	out := &krm.InstancegroupmanagerCurrentActionsStatus{}
	if in.Abandoning != nil {
		val := int64(*in.Abandoning)
		out.Abandoning = &val
	}
	if in.Creating != nil {
		val := int64(*in.Creating)
		out.Creating = &val
	}
	if in.CreatingWithoutRetries != nil {
		val := int64(*in.CreatingWithoutRetries)
		out.CreatingWithoutRetries = &val
	}
	if in.Deleting != nil {
		val := int64(*in.Deleting)
		out.Deleting = &val
	}
	if in.None != nil {
		val := int64(*in.None)
		out.None = &val
	}
	if in.Recreating != nil {
		val := int64(*in.Recreating)
		out.Recreating = &val
	}
	if in.Refreshing != nil {
		val := int64(*in.Refreshing)
		out.Refreshing = &val
	}
	if in.Restarting != nil {
		val := int64(*in.Restarting)
		out.Restarting = &val
	}
	if in.Verifying != nil {
		val := int64(*in.Verifying)
		out.Verifying = &val
	}
	return out
}

func InstancegroupmanagerCurrentActionsStatus_ToProto(in *krm.InstancegroupmanagerCurrentActionsStatus) *pb.InstanceGroupManagerActionsSummary {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerActionsSummary{}
	if in.Abandoning != nil {
		val := int32(*in.Abandoning)
		out.Abandoning = &val
	}
	if in.Creating != nil {
		val := int32(*in.Creating)
		out.Creating = &val
	}
	if in.CreatingWithoutRetries != nil {
		val := int32(*in.CreatingWithoutRetries)
		out.CreatingWithoutRetries = &val
	}
	if in.Deleting != nil {
		val := int32(*in.Deleting)
		out.Deleting = &val
	}
	if in.None != nil {
		val := int32(*in.None)
		out.None = &val
	}
	if in.Recreating != nil {
		val := int32(*in.Recreating)
		out.Recreating = &val
	}
	if in.Refreshing != nil {
		val := int32(*in.Refreshing)
		out.Refreshing = &val
	}
	if in.Restarting != nil {
		val := int32(*in.Restarting)
		out.Restarting = &val
	}
	if in.Verifying != nil {
		val := int32(*in.Verifying)
		out.Verifying = &val
	}
	return out
}

func InstancegroupmanagerStatusStatus_FromProto(in *pb.InstanceGroupManagerStatus) *krm.InstancegroupmanagerStatusStatus {
	if in == nil {
		return nil
	}
	out := &krm.InstancegroupmanagerStatusStatus{}
	out.Autoscaler = in.Autoscaler
	out.IsStable = in.IsStable
	if in.Stateful != nil {
		out.Stateful = &krm.InstancegroupmanagerStatefulStatus{}
		out.Stateful.HasStatefulConfig = in.Stateful.HasStatefulConfig
		if in.Stateful.PerInstanceConfigs != nil {
			out.Stateful.PerInstanceConfigs = &krm.InstancegroupmanagerPerInstanceConfigsStatus{
				AllEffective: in.Stateful.PerInstanceConfigs.AllEffective,
			}
		}
	}
	if in.VersionTarget != nil {
		out.VersionTarget = &krm.InstancegroupmanagerVersionTargetStatus{
			IsReached: in.VersionTarget.IsReached,
		}
	}
	return out
}

func InstancegroupmanagerStatusStatus_ToProto(in *krm.InstancegroupmanagerStatusStatus) *pb.InstanceGroupManagerStatus {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerStatus{}
	out.Autoscaler = in.Autoscaler
	out.IsStable = in.IsStable
	if in.Stateful != nil {
		out.Stateful = &pb.InstanceGroupManagerStatusStateful{}
		out.Stateful.HasStatefulConfig = in.Stateful.HasStatefulConfig
		if in.Stateful.PerInstanceConfigs != nil {
			out.Stateful.PerInstanceConfigs = &pb.InstanceGroupManagerStatusStatefulPerInstanceConfigs{
				AllEffective: in.Stateful.PerInstanceConfigs.AllEffective,
			}
		}
	}
	if in.VersionTarget != nil {
		out.VersionTarget = &pb.InstanceGroupManagerStatusVersionTarget{
			IsReached: in.VersionTarget.IsReached,
		}
	}
	return out
}

func ComputeInstanceGroupManagerStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManager) krm.ComputeInstanceGroupManagerStatus {
	out := krm.ComputeInstanceGroupManagerStatus{}
	if in == nil {
		return out
	}
	out.CreationTimestamp = in.CreationTimestamp
	out.CurrentActions = InstancegroupmanagerCurrentActionsStatus_FromProto(in.CurrentActions)
	out.Fingerprint = in.Fingerprint
	if in.Id != nil {
		val := int64(*in.Id)
		out.Id = &val
	}
	out.InstanceGroup = in.InstanceGroup
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	out.Status = InstancegroupmanagerStatusStatus_FromProto(in.Status)
	out.Zone = in.Zone
	return out
}

func ComputeInstanceGroupManagerStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceGroupManagerStatus) *pb.InstanceGroupManager {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManager{}
	out.CreationTimestamp = in.CreationTimestamp
	out.CurrentActions = InstancegroupmanagerCurrentActionsStatus_ToProto(in.CurrentActions)
	out.Fingerprint = in.Fingerprint
	if in.Id != nil {
		val := uint64(*in.Id)
		out.Id = &val
	}
	out.InstanceGroup = in.InstanceGroup
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	out.Status = InstancegroupmanagerStatusStatus_ToProto(in.Status)
	out.Zone = in.Zone
	return out
}

func ComputeInstanceGroupManager_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManager) *krm.ComputeInstanceGroupManager {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceGroupManager{}
	out.Spec = *ComputeInstanceGroupManagerSpec_v1beta1_FromProto(mapCtx, in)
	out.Status = ComputeInstanceGroupManagerStatus_v1beta1_FromProto(mapCtx, in)
	return out
}

func ComputeInstanceGroupManager_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceGroupManager) *pb.InstanceGroupManager {
	if in == nil {
		return nil
	}
	out := ComputeInstanceGroupManagerSpec_v1beta1_ToProto(mapCtx, &in.Spec)
	return out
}
