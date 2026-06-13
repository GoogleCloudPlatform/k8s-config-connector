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

func ComputeInstanceGroupManagerSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManager) *krm.ComputeInstanceGroupManagerSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceGroupManagerSpec{}
	out.AutoHealingPolicies = direct.Slice_FromProto(mapCtx, in.AutoHealingPolicies, InstanceGroupManagerAutoHealingPolicy_v1beta1_FromProto)
	out.BaseInstanceName = in.BaseInstanceName
	out.Description = in.Description
	out.DistributionPolicy = InstanceGroupManagerDistributionPolicy_v1beta1_FromProto(mapCtx, in.DistributionPolicy)
	if in.InstanceTemplate != nil {
		out.InstanceTemplateRef = &krm.ComputeInstanceGroupManagerInstanceTemplateRef{External: *in.InstanceTemplate}
	}
	out.NamedPorts = direct.Slice_FromProto(mapCtx, in.NamedPorts, InstanceGroupManagerNamedPort_v1beta1_FromProto)
	out.StatefulPolicy = InstanceGroupManagerStatefulPolicy_v1beta1_FromProto(mapCtx, in.StatefulPolicy)
	out.TargetPools = ComputeInstanceGroupManagerSpec_TargetPools_FromProto(mapCtx, in.TargetPools)
	out.TargetSize = int64(direct.ValueOf(in.TargetSize))
	out.UpdatePolicy = InstanceGroupManagerUpdatePolicy_v1beta1_FromProto(mapCtx, in.UpdatePolicy)
	out.Versions = direct.Slice_FromProto(mapCtx, in.Versions, InstanceGroupManagerVersion_v1beta1_FromProto)
	return out
}

func ComputeInstanceGroupManagerSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceGroupManagerSpec) *pb.InstanceGroupManager {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManager{}
	out.AutoHealingPolicies = direct.Slice_ToProto(mapCtx, in.AutoHealingPolicies, InstanceGroupManagerAutoHealingPolicy_v1beta1_ToProto)
	out.BaseInstanceName = in.BaseInstanceName
	out.Description = in.Description
	out.DistributionPolicy = InstanceGroupManagerDistributionPolicy_v1beta1_ToProto(mapCtx, in.DistributionPolicy)
	if in.InstanceTemplateRef != nil {
		out.InstanceTemplate = &in.InstanceTemplateRef.External
	}
	out.NamedPorts = direct.Slice_ToProto(mapCtx, in.NamedPorts, InstanceGroupManagerNamedPort_v1beta1_ToProto)
	out.StatefulPolicy = InstanceGroupManagerStatefulPolicy_v1beta1_ToProto(mapCtx, in.StatefulPolicy)
	out.TargetPools = ComputeInstanceGroupManagerSpec_TargetPools_ToProto(mapCtx, in.TargetPools)
	out.TargetSize = direct.PtrTo(int32(in.TargetSize))
	out.UpdatePolicy = InstanceGroupManagerUpdatePolicy_v1beta1_ToProto(mapCtx, in.UpdatePolicy)
	out.Versions = direct.Slice_ToProto(mapCtx, in.Versions, InstanceGroupManagerVersion_v1beta1_ToProto)
	return out
}

func ComputeInstanceGroupManagerStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManager) *krm.ComputeInstanceGroupManagerStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInstanceGroupManagerStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.CurrentActions = InstanceGroupManagerCurrentActionsStatus_v1beta1_FromProto(mapCtx, in.CurrentActions)
	out.Fingerprint = in.Fingerprint
	out.ID = direct.PtrUint64ToPtrInt64(in.Id)
	out.InstanceGroup = in.InstanceGroup
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	out.Status = InstanceGroupManagerStatusStatus_v1beta1_FromProto(mapCtx, in.Status)
	out.UpdatePolicy = InstanceGroupManagerUpdatePolicyStatus_v1beta1_FromProto(mapCtx, in.UpdatePolicy)
	out.Zone = in.Zone
	return out
}

func ComputeInstanceGroupManagerStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceGroupManagerStatus) *pb.InstanceGroupManager {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManager{}
	out.CreationTimestamp = in.CreationTimestamp
	out.CurrentActions = InstanceGroupManagerCurrentActionsStatus_v1beta1_ToProto(mapCtx, in.CurrentActions)
	out.Fingerprint = in.Fingerprint
	out.Id = direct.PtrInt64ToPtrUint64(in.ID)
	out.InstanceGroup = in.InstanceGroup
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	out.Status = InstanceGroupManagerStatusStatus_v1beta1_ToProto(mapCtx, in.Status)
	out.UpdatePolicy = InstanceGroupManagerUpdatePolicyStatus_v1beta1_ToProto(mapCtx, in.UpdatePolicy)
	out.Zone = in.Zone
	return out
}

func InstanceGroupManagerAutoHealingPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManagerAutoHealingPolicy) *krm.InstanceGroupManagerAutoHealingPolicy {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerAutoHealingPolicy{}
	if in.HealthCheck != nil {
		out.HealthCheckRef = &krm.ComputeInstanceGroupManagerHealthCheckRef{External: *in.HealthCheck}
	}
	out.InitialDelaySec = direct.PtrInt32ToPtrInt64(in.InitialDelaySec)
	return out
}

func InstanceGroupManagerAutoHealingPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerAutoHealingPolicy) *pb.InstanceGroupManagerAutoHealingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerAutoHealingPolicy{}
	if in.HealthCheckRef != nil {
		out.HealthCheck = &in.HealthCheckRef.External
	}
	out.InitialDelaySec = direct.PtrInt64ToPtrInt32(in.InitialDelaySec)
	return out
}

func InstanceGroupManagerDistributionPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.DistributionPolicy) *krm.InstanceGroupManagerDistributionPolicy {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerDistributionPolicy{}
	out.TargetShape = in.TargetShape
	out.Zones = direct.Slice_FromProto(mapCtx, in.Zones, DistributionPolicyZoneConfiguration_v1beta1_FromProto)
	return out
}

func InstanceGroupManagerDistributionPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerDistributionPolicy) *pb.DistributionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.DistributionPolicy{}
	out.TargetShape = in.TargetShape
	out.Zones = direct.Slice_ToProto(mapCtx, in.Zones, DistributionPolicyZoneConfiguration_v1beta1_ToProto)
	return out
}

func DistributionPolicyZoneConfiguration_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.DistributionPolicyZoneConfiguration) *krm.InstanceGroupManagerZone {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerZone{}
	out.Zone = in.Zone
	return out
}

func DistributionPolicyZoneConfiguration_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerZone) *pb.DistributionPolicyZoneConfiguration {
	if in == nil {
		return nil
	}
	out := &pb.DistributionPolicyZoneConfiguration{}
	out.Zone = in.Zone
	return out
}

func InstanceGroupManagerNamedPort_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NamedPort) *krm.InstanceGroupManagerNamedPort {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerNamedPort{}
	out.Name = in.Name
	out.Port = direct.PtrInt32ToPtrInt64(in.Port)
	return out
}

func InstanceGroupManagerNamedPort_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerNamedPort) *pb.NamedPort {
	if in == nil {
		return nil
	}
	out := &pb.NamedPort{}
	out.Name = in.Name
	out.Port = direct.PtrInt64ToPtrInt32(in.Port)
	return out
}

func InstanceGroupManagerStatefulPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.StatefulPolicy) *krm.InstanceGroupManagerStatefulPolicy {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerStatefulPolicy{}
	out.PreservedState = StatefulPolicyPreservedState_v1beta1_FromProto(mapCtx, in.PreservedState)
	return out
}

func InstanceGroupManagerStatefulPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerStatefulPolicy) *pb.StatefulPolicy {
	if in == nil {
		return nil
	}
	out := &pb.StatefulPolicy{}
	out.PreservedState = StatefulPolicyPreservedState_v1beta1_ToProto(mapCtx, in.PreservedState)
	return out
}

func StatefulPolicyPreservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.StatefulPolicyPreservedState) *krm.InstanceGroupManagerPreservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerPreservedState{}
	out.Disks = StatefulPolicyPreservedStateDiskDeviceMap_FromProto(mapCtx, in.Disks)
	out.ExternalIps = StatefulPolicyPreservedStateExternalIPMap_FromProto(mapCtx, in.ExternalIPs)
	out.InternalIps = StatefulPolicyPreservedStateInternalIPMap_FromProto(mapCtx, in.InternalIPs)
	return out
}

func StatefulPolicyPreservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerPreservedState) *pb.StatefulPolicyPreservedState {
	if in == nil {
		return nil
	}
	out := &pb.StatefulPolicyPreservedState{}
	out.Disks = StatefulPolicyPreservedStateDiskDeviceMap_ToProto(mapCtx, in.Disks)
	out.ExternalIPs = StatefulPolicyPreservedStateExternalIPMap_ToProto(mapCtx, in.ExternalIps)
	out.InternalIPs = StatefulPolicyPreservedStateInternalIPMap_ToProto(mapCtx, in.InternalIps)
	return out
}

func StatefulPolicyPreservedStateDiskDeviceMap_FromProto(mapCtx *direct.MapContext, in map[string]*pb.StatefulPolicyPreservedStateDiskDevice) map[string]krm.InstanceGroupManagerPreservedStateDisk {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.InstanceGroupManagerPreservedStateDisk)
	for k, v := range in {
		if v == nil {
			continue
		}
		out[k] = krm.InstanceGroupManagerPreservedStateDisk{
			AutoDelete: v.AutoDelete,
		}
	}
	return out
}

func StatefulPolicyPreservedStateDiskDeviceMap_ToProto(mapCtx *direct.MapContext, in map[string]krm.InstanceGroupManagerPreservedStateDisk) map[string]*pb.StatefulPolicyPreservedStateDiskDevice {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.StatefulPolicyPreservedStateDiskDevice)
	for k, v := range in {
		out[k] = &pb.StatefulPolicyPreservedStateDiskDevice{
			AutoDelete: v.AutoDelete,
		}
	}
	return out
}

func StatefulPolicyPreservedStateExternalIPMap_FromProto(mapCtx *direct.MapContext, in map[string]*pb.StatefulPolicyPreservedStateNetworkIp) map[string]krm.InstanceGroupManagerPreservedStateExternalIp {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.InstanceGroupManagerPreservedStateExternalIp)
	for k, v := range in {
		if v == nil {
			continue
		}
		out[k] = krm.InstanceGroupManagerPreservedStateExternalIp{
			AutoDelete: v.AutoDelete,
		}
	}
	return out
}

func StatefulPolicyPreservedStateExternalIPMap_ToProto(mapCtx *direct.MapContext, in map[string]krm.InstanceGroupManagerPreservedStateExternalIp) map[string]*pb.StatefulPolicyPreservedStateNetworkIp {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.StatefulPolicyPreservedStateNetworkIp)
	for k, v := range in {
		out[k] = &pb.StatefulPolicyPreservedStateNetworkIp{
			AutoDelete: v.AutoDelete,
		}
	}
	return out
}

func StatefulPolicyPreservedStateInternalIPMap_FromProto(mapCtx *direct.MapContext, in map[string]*pb.StatefulPolicyPreservedStateNetworkIp) map[string]krm.InstanceGroupManagerPreservedStateInternalIp {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.InstanceGroupManagerPreservedStateInternalIp)
	for k, v := range in {
		if v == nil {
			continue
		}
		out[k] = krm.InstanceGroupManagerPreservedStateInternalIp{
			AutoDelete: v.AutoDelete,
		}
	}
	return out
}

func StatefulPolicyPreservedStateInternalIPMap_ToProto(mapCtx *direct.MapContext, in map[string]krm.InstanceGroupManagerPreservedStateInternalIp) map[string]*pb.StatefulPolicyPreservedStateNetworkIp {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.StatefulPolicyPreservedStateNetworkIp)
	for k, v := range in {
		out[k] = &pb.StatefulPolicyPreservedStateNetworkIp{
			AutoDelete: v.AutoDelete,
		}
	}
	return out
}

func InstanceGroupManagerUpdatePolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManagerUpdatePolicy) *krm.InstanceGroupManagerUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerUpdatePolicy{}
	out.InstanceRedistributionType = in.InstanceRedistributionType
	out.MaxSurge = FixedOrPercentMaxSurge_FromProto(mapCtx, in.MaxSurge)
	out.MaxUnavailable = FixedOrPercentMaxUnavailable_FromProto(mapCtx, in.MaxUnavailable)
	out.MinimalAction = in.MinimalAction
	out.MostDisruptiveAllowedAction = in.MostDisruptiveAllowedAction
	out.ReplacementMethod = in.ReplacementMethod
	out.Type = in.Type
	return out
}

func InstanceGroupManagerUpdatePolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerUpdatePolicy) *pb.InstanceGroupManagerUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerUpdatePolicy{}
	out.InstanceRedistributionType = in.InstanceRedistributionType
	out.MaxSurge = FixedOrPercentMaxSurge_ToProto(mapCtx, in.MaxSurge)
	out.MaxUnavailable = FixedOrPercentMaxUnavailable_ToProto(mapCtx, in.MaxUnavailable)
	out.MinimalAction = in.MinimalAction
	out.MostDisruptiveAllowedAction = in.MostDisruptiveAllowedAction
	out.ReplacementMethod = in.ReplacementMethod
	out.Type = in.Type
	return out
}

func FixedOrPercentMaxSurge_FromProto(mapCtx *direct.MapContext, in *pb.FixedOrPercent) *krm.InstanceGroupManagerMaxSurge {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerMaxSurge{}
	out.Fixed = direct.PtrInt32ToPtrInt64(in.Fixed)
	out.Percent = direct.PtrInt32ToPtrInt64(in.Percent)
	return out
}

func FixedOrPercentMaxSurge_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerMaxSurge) *pb.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &pb.FixedOrPercent{}
	out.Fixed = direct.PtrInt64ToPtrInt32(in.Fixed)
	out.Percent = direct.PtrInt64ToPtrInt32(in.Percent)
	return out
}

func FixedOrPercentMaxUnavailable_FromProto(mapCtx *direct.MapContext, in *pb.FixedOrPercent) *krm.InstanceGroupManagerMaxUnavailable {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerMaxUnavailable{}
	out.Fixed = direct.PtrInt32ToPtrInt64(in.Fixed)
	out.Percent = direct.PtrInt32ToPtrInt64(in.Percent)
	return out
}

func FixedOrPercentMaxUnavailable_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerMaxUnavailable) *pb.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &pb.FixedOrPercent{}
	out.Fixed = direct.PtrInt64ToPtrInt32(in.Fixed)
	out.Percent = direct.PtrInt64ToPtrInt32(in.Percent)
	return out
}

func FixedOrPercentTargetSize_FromProto(mapCtx *direct.MapContext, in *pb.FixedOrPercent) *krm.InstanceGroupManagerTargetSize {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerTargetSize{}
	out.Calculated = direct.PtrInt32ToPtrInt64(in.Calculated)
	out.Fixed = direct.PtrInt32ToPtrInt64(in.Fixed)
	out.Percent = direct.PtrInt32ToPtrInt64(in.Percent)
	return out
}

func FixedOrPercentTargetSize_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerTargetSize) *pb.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &pb.FixedOrPercent{}
	out.Calculated = direct.PtrInt64ToPtrInt32(in.Calculated)
	out.Fixed = direct.PtrInt64ToPtrInt32(in.Fixed)
	out.Percent = direct.PtrInt64ToPtrInt32(in.Percent)
	return out
}

func InstanceGroupManagerVersion_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManagerVersion) *krm.InstanceGroupManagerVersion {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerVersion{}
	if in.InstanceTemplate != nil {
		out.InstanceTemplateRef = &krm.ComputeInstanceGroupManagerInstanceTemplateRef{External: *in.InstanceTemplate}
	}
	out.Name = in.Name
	out.TargetSize = FixedOrPercentTargetSize_FromProto(mapCtx, in.TargetSize)
	return out
}

func InstanceGroupManagerVersion_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerVersion) *pb.InstanceGroupManagerVersion {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerVersion{}
	if in.InstanceTemplateRef != nil {
		out.InstanceTemplate = &in.InstanceTemplateRef.External
	}
	out.Name = in.Name
	out.TargetSize = FixedOrPercentTargetSize_ToProto(mapCtx, in.TargetSize)
	return out
}

func InstanceGroupManagerCurrentActionsStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManagerActionsSummary) *krm.InstanceGroupManagerCurrentActionsStatus {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerCurrentActionsStatus{}
	out.Abandoning = direct.PtrInt32ToPtrInt64(in.Abandoning)
	out.Creating = direct.PtrInt32ToPtrInt64(in.Creating)
	out.CreatingWithoutRetries = direct.PtrInt32ToPtrInt64(in.CreatingWithoutRetries)
	out.Deleting = direct.PtrInt32ToPtrInt64(in.Deleting)
	out.None = direct.PtrInt32ToPtrInt64(in.None)
	out.Recreating = direct.PtrInt32ToPtrInt64(in.Recreating)
	out.Refreshing = direct.PtrInt32ToPtrInt64(in.Refreshing)
	out.Restarting = direct.PtrInt32ToPtrInt64(in.Restarting)
	out.Verifying = direct.PtrInt32ToPtrInt64(in.Verifying)
	return out
}

func InstanceGroupManagerCurrentActionsStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerCurrentActionsStatus) *pb.InstanceGroupManagerActionsSummary {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerActionsSummary{}
	out.Abandoning = direct.PtrInt64ToPtrInt32(in.Abandoning)
	out.Creating = direct.PtrInt64ToPtrInt32(in.Creating)
	out.CreatingWithoutRetries = direct.PtrInt64ToPtrInt32(in.CreatingWithoutRetries)
	out.Deleting = direct.PtrInt64ToPtrInt32(in.Deleting)
	out.None = direct.PtrInt64ToPtrInt32(in.None)
	out.Recreating = direct.PtrInt64ToPtrInt32(in.Recreating)
	out.Refreshing = direct.PtrInt64ToPtrInt32(in.Refreshing)
	out.Restarting = direct.PtrInt64ToPtrInt32(in.Restarting)
	out.Verifying = direct.PtrInt64ToPtrInt32(in.Verifying)
	return out
}

func InstanceGroupManagerStatusStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManagerStatus) *krm.InstanceGroupManagerStatusStatus {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerStatusStatus{}
	out.Autoscaler = in.Autoscaler
	out.IsStable = in.IsStable
	out.Stateful = InstanceGroupManagerStatusStateful_v1beta1_FromProto(mapCtx, in.Stateful)
	out.VersionTarget = InstanceGroupManagerStatusVersionTarget_v1beta1_FromProto(mapCtx, in.VersionTarget)
	return out
}

func InstanceGroupManagerStatusStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerStatusStatus) *pb.InstanceGroupManagerStatus {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerStatus{}
	out.Autoscaler = in.Autoscaler
	out.IsStable = in.IsStable
	out.Stateful = InstanceGroupManagerStatusStateful_v1beta1_ToProto(mapCtx, in.Stateful)
	out.VersionTarget = InstanceGroupManagerStatusVersionTarget_v1beta1_ToProto(mapCtx, in.VersionTarget)
	return out
}

func InstanceGroupManagerStatusStateful_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManagerStatusStateful) *krm.InstanceGroupManagerStatefulStatus {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerStatefulStatus{}
	out.HasStatefulConfig = in.HasStatefulConfig
	out.PerInstanceConfigs = InstanceGroupManagerStatusStatefulPerInstanceConfigs_v1beta1_FromProto(mapCtx, in.PerInstanceConfigs)
	return out
}

func InstanceGroupManagerStatusStateful_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerStatefulStatus) *pb.InstanceGroupManagerStatusStateful {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerStatusStateful{}
	out.HasStatefulConfig = in.HasStatefulConfig
	out.PerInstanceConfigs = InstanceGroupManagerStatusStatefulPerInstanceConfigs_v1beta1_ToProto(mapCtx, in.PerInstanceConfigs)
	return out
}

func InstanceGroupManagerStatusStatefulPerInstanceConfigs_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManagerStatusStatefulPerInstanceConfigs) *krm.InstanceGroupManagerPerInstanceConfigsStatus {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerPerInstanceConfigsStatus{}
	out.AllEffective = in.AllEffective
	return out
}

func InstanceGroupManagerStatusStatefulPerInstanceConfigs_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerPerInstanceConfigsStatus) *pb.InstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerStatusStatefulPerInstanceConfigs{}
	out.AllEffective = in.AllEffective
	return out
}

func InstanceGroupManagerStatusVersionTarget_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManagerStatusVersionTarget) *krm.InstanceGroupManagerVersionTargetStatus {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerVersionTargetStatus{}
	out.IsReached = in.IsReached
	return out
}

func InstanceGroupManagerStatusVersionTarget_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerVersionTargetStatus) *pb.InstanceGroupManagerStatusVersionTarget {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerStatusVersionTarget{}
	out.IsReached = in.IsReached
	return out
}

func InstanceGroupManagerUpdatePolicyStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupManagerUpdatePolicy) *krm.InstanceGroupManagerUpdatePolicyStatus {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerUpdatePolicyStatus{}
	out.MaxSurge = FixedOrPercentMaxSurgeStatus_v1beta1_FromProto(mapCtx, in.MaxSurge)
	out.MaxUnavailable = FixedOrPercentMaxUnavailableStatus_v1beta1_FromProto(mapCtx, in.MaxUnavailable)
	return out
}

func InstanceGroupManagerUpdatePolicyStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerUpdatePolicyStatus) *pb.InstanceGroupManagerUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupManagerUpdatePolicy{}
	out.MaxSurge = FixedOrPercentMaxSurgeStatus_v1beta1_ToProto(mapCtx, in.MaxSurge)
	out.MaxUnavailable = FixedOrPercentMaxUnavailableStatus_v1beta1_ToProto(mapCtx, in.MaxUnavailable)
	return out
}

func FixedOrPercentMaxSurgeStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.FixedOrPercent) *krm.InstanceGroupManagerMaxSurgeStatus {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerMaxSurgeStatus{}
	out.Calculated = direct.PtrInt32ToPtrInt64(in.Calculated)
	return out
}

func FixedOrPercentMaxSurgeStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerMaxSurgeStatus) *pb.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &pb.FixedOrPercent{}
	out.Calculated = direct.PtrInt64ToPtrInt32(in.Calculated)
	return out
}

func FixedOrPercentMaxUnavailableStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.FixedOrPercent) *krm.InstanceGroupManagerMaxUnavailableStatus {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupManagerMaxUnavailableStatus{}
	out.Calculated = direct.PtrInt32ToPtrInt64(in.Calculated)
	return out
}

func FixedOrPercentMaxUnavailableStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupManagerMaxUnavailableStatus) *pb.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &pb.FixedOrPercent{}
	out.Calculated = direct.PtrInt64ToPtrInt32(in.Calculated)
	return out
}

func ComputeInstanceGroupManagerSpec_TargetPools_FromProto(mapCtx *direct.MapContext, in []string) []krm.ComputeInstanceGroupManagerTargetPoolRef {
	if in == nil {
		return nil
	}
	var out []krm.ComputeInstanceGroupManagerTargetPoolRef
	for _, i := range in {
		out = append(out, krm.ComputeInstanceGroupManagerTargetPoolRef{
			External: i,
		})
	}
	return out
}

func ComputeInstanceGroupManagerSpec_TargetPools_ToProto(mapCtx *direct.MapContext, in []krm.ComputeInstanceGroupManagerTargetPoolRef) []string {
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
