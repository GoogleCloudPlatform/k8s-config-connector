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

package deploy

import (
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/deploy/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AnthosCluster_FromProto(mapCtx *direct.MapContext, in *pb.AnthosCluster) *krm.AnthosCluster {
	if in == nil {
		return nil
	}
	out := &krm.AnthosCluster{}
	out.Membership = direct.LazyPtr(in.GetMembership())
	return out
}
func AnthosCluster_ToProto(mapCtx *direct.MapContext, in *krm.AnthosCluster) *pb.AnthosCluster {
	if in == nil {
		return nil
	}
	out := &pb.AnthosCluster{}
	out.Membership = direct.ValueOf(in.Membership)
	return out
}
func AssociatedEntities_FromProto(mapCtx *direct.MapContext, in *pb.AssociatedEntities) *krm.AssociatedEntities {
	if in == nil {
		return nil
	}
	out := &krm.AssociatedEntities{}
	out.GkeClusters = direct.Slice_FromProto(mapCtx, in.GkeClusters, GkeCluster_FromProto)
	out.AnthosClusters = direct.Slice_FromProto(mapCtx, in.AnthosClusters, AnthosCluster_FromProto)
	return out
}
func AssociatedEntities_ToProto(mapCtx *direct.MapContext, in *krm.AssociatedEntities) *pb.AssociatedEntities {
	if in == nil {
		return nil
	}
	out := &pb.AssociatedEntities{}
	out.GkeClusters = direct.Slice_ToProto(mapCtx, in.GkeClusters, GkeCluster_ToProto)
	out.AnthosClusters = direct.Slice_ToProto(mapCtx, in.AnthosClusters, AnthosCluster_ToProto)
	return out
}
func CloudRunLocation_FromProto(mapCtx *direct.MapContext, in *pb.CloudRunLocation) *krm.CloudRunLocation {
	if in == nil {
		return nil
	}
	out := &krm.CloudRunLocation{}
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func CloudRunLocation_ToProto(mapCtx *direct.MapContext, in *krm.CloudRunLocation) *pb.CloudRunLocation {
	if in == nil {
		return nil
	}
	out := &pb.CloudRunLocation{}
	out.Location = direct.ValueOf(in.Location)
	return out
}
func CustomTarget_FromProto(mapCtx *direct.MapContext, in *pb.CustomTarget) *krm.CustomTarget {
	if in == nil {
		return nil
	}
	out := &krm.CustomTarget{}
	out.CustomTargetType = direct.LazyPtr(in.GetCustomTargetType())
	return out
}
func CustomTarget_ToProto(mapCtx *direct.MapContext, in *krm.CustomTarget) *pb.CustomTarget {
	if in == nil {
		return nil
	}
	out := &pb.CustomTarget{}
	out.CustomTargetType = direct.ValueOf(in.CustomTargetType)
	return out
}
func DefaultPool_FromProto(mapCtx *direct.MapContext, in *pb.DefaultPool) *krm.DefaultPool {
	if in == nil {
		return nil
	}
	out := &krm.DefaultPool{}
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.ArtifactStorage = direct.LazyPtr(in.GetArtifactStorage())
	return out
}
func DefaultPool_ToProto(mapCtx *direct.MapContext, in *krm.DefaultPool) *pb.DefaultPool {
	if in == nil {
		return nil
	}
	out := &pb.DefaultPool{}
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.ArtifactStorage = direct.ValueOf(in.ArtifactStorage)
	return out
}
func DeployTargetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Target) *krm.DeployTargetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployTargetObservedState{}
	// MISSING: Name
	// MISSING: TargetID
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: RequireApproval
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Gke
	// MISSING: AnthosCluster
	// MISSING: Run
	// MISSING: MultiTarget
	// MISSING: CustomTarget
	// MISSING: AssociatedEntities
	// MISSING: Etag
	// MISSING: ExecutionConfigs
	// MISSING: DeployParameters
	return out
}
func DeployTargetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployTargetObservedState) *pb.Target {
	if in == nil {
		return nil
	}
	out := &pb.Target{}
	// MISSING: Name
	// MISSING: TargetID
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: RequireApproval
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Gke
	// MISSING: AnthosCluster
	// MISSING: Run
	// MISSING: MultiTarget
	// MISSING: CustomTarget
	// MISSING: AssociatedEntities
	// MISSING: Etag
	// MISSING: ExecutionConfigs
	// MISSING: DeployParameters
	return out
}
func DeployTargetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Target) *krm.DeployTargetSpec {
	if in == nil {
		return nil
	}
	out := &krm.DeployTargetSpec{}
	// MISSING: Name
	// MISSING: TargetID
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: RequireApproval
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Gke
	// MISSING: AnthosCluster
	// MISSING: Run
	// MISSING: MultiTarget
	// MISSING: CustomTarget
	// MISSING: AssociatedEntities
	// MISSING: Etag
	// MISSING: ExecutionConfigs
	// MISSING: DeployParameters
	return out
}
func DeployTargetSpec_ToProto(mapCtx *direct.MapContext, in *krm.DeployTargetSpec) *pb.Target {
	if in == nil {
		return nil
	}
	out := &pb.Target{}
	// MISSING: Name
	// MISSING: TargetID
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: RequireApproval
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Gke
	// MISSING: AnthosCluster
	// MISSING: Run
	// MISSING: MultiTarget
	// MISSING: CustomTarget
	// MISSING: AssociatedEntities
	// MISSING: Etag
	// MISSING: ExecutionConfigs
	// MISSING: DeployParameters
	return out
}
func ExecutionConfig_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionConfig) *krm.ExecutionConfig {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionConfig{}
	out.Usages = direct.EnumSlice_FromProto(mapCtx, in.Usages)
	out.DefaultPool = DefaultPool_FromProto(mapCtx, in.GetDefaultPool())
	out.PrivatePool = PrivatePool_FromProto(mapCtx, in.GetPrivatePool())
	out.WorkerPool = direct.LazyPtr(in.GetWorkerPool())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.ArtifactStorage = direct.LazyPtr(in.GetArtifactStorage())
	out.ExecutionTimeout = direct.StringDuration_FromProto(mapCtx, in.GetExecutionTimeout())
	out.Verbose = direct.LazyPtr(in.GetVerbose())
	return out
}
func ExecutionConfig_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionConfig) *pb.ExecutionConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionConfig{}
	out.Usages = direct.EnumSlice_ToProto[pb.ExecutionConfig_ExecutionEnvironmentUsage](mapCtx, in.Usages)
	if oneof := DefaultPool_ToProto(mapCtx, in.DefaultPool); oneof != nil {
		out.ExecutionEnvironment = &pb.ExecutionConfig_DefaultPool{DefaultPool: oneof}
	}
	if oneof := PrivatePool_ToProto(mapCtx, in.PrivatePool); oneof != nil {
		out.ExecutionEnvironment = &pb.ExecutionConfig_PrivatePool{PrivatePool: oneof}
	}
	out.WorkerPool = direct.ValueOf(in.WorkerPool)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.ArtifactStorage = direct.ValueOf(in.ArtifactStorage)
	out.ExecutionTimeout = direct.StringDuration_ToProto(mapCtx, in.ExecutionTimeout)
	out.Verbose = direct.ValueOf(in.Verbose)
	return out
}
func GkeCluster_FromProto(mapCtx *direct.MapContext, in *pb.GkeCluster) *krm.GkeCluster {
	if in == nil {
		return nil
	}
	out := &krm.GkeCluster{}
	out.Cluster = direct.LazyPtr(in.GetCluster())
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.ProxyURL = direct.LazyPtr(in.GetProxyUrl())
	out.DNSEndpoint = direct.LazyPtr(in.GetDnsEndpoint())
	return out
}
func GkeCluster_ToProto(mapCtx *direct.MapContext, in *krm.GkeCluster) *pb.GkeCluster {
	if in == nil {
		return nil
	}
	out := &pb.GkeCluster{}
	out.Cluster = direct.ValueOf(in.Cluster)
	out.InternalIp = direct.ValueOf(in.InternalIP)
	out.ProxyUrl = direct.ValueOf(in.ProxyURL)
	out.DnsEndpoint = direct.ValueOf(in.DNSEndpoint)
	return out
}
func MultiTarget_FromProto(mapCtx *direct.MapContext, in *pb.MultiTarget) *krm.MultiTarget {
	if in == nil {
		return nil
	}
	out := &krm.MultiTarget{}
	out.TargetIds = in.TargetIds
	return out
}
func MultiTarget_ToProto(mapCtx *direct.MapContext, in *krm.MultiTarget) *pb.MultiTarget {
	if in == nil {
		return nil
	}
	out := &pb.MultiTarget{}
	out.TargetIds = in.TargetIds
	return out
}
func PrivatePool_FromProto(mapCtx *direct.MapContext, in *pb.PrivatePool) *krm.PrivatePool {
	if in == nil {
		return nil
	}
	out := &krm.PrivatePool{}
	out.WorkerPool = direct.LazyPtr(in.GetWorkerPool())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.ArtifactStorage = direct.LazyPtr(in.GetArtifactStorage())
	return out
}
func PrivatePool_ToProto(mapCtx *direct.MapContext, in *krm.PrivatePool) *pb.PrivatePool {
	if in == nil {
		return nil
	}
	out := &pb.PrivatePool{}
	out.WorkerPool = direct.ValueOf(in.WorkerPool)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.ArtifactStorage = direct.ValueOf(in.ArtifactStorage)
	return out
}
func Target_FromProto(mapCtx *direct.MapContext, in *pb.Target) *krm.Target {
	if in == nil {
		return nil
	}
	out := &krm.Target{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: TargetID
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	out.RequireApproval = direct.LazyPtr(in.GetRequireApproval())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Gke = GkeCluster_FromProto(mapCtx, in.GetGke())
	out.AnthosCluster = AnthosCluster_FromProto(mapCtx, in.GetAnthosCluster())
	out.Run = CloudRunLocation_FromProto(mapCtx, in.GetRun())
	out.MultiTarget = MultiTarget_FromProto(mapCtx, in.GetMultiTarget())
	out.CustomTarget = CustomTarget_FromProto(mapCtx, in.GetCustomTarget())
	// MISSING: AssociatedEntities
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.ExecutionConfigs = direct.Slice_FromProto(mapCtx, in.ExecutionConfigs, ExecutionConfig_FromProto)
	out.DeployParameters = in.DeployParameters
	return out
}
func Target_ToProto(mapCtx *direct.MapContext, in *krm.Target) *pb.Target {
	if in == nil {
		return nil
	}
	out := &pb.Target{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: TargetID
	// MISSING: Uid
	out.Description = direct.ValueOf(in.Description)
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	out.RequireApproval = direct.ValueOf(in.RequireApproval)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	if oneof := GkeCluster_ToProto(mapCtx, in.Gke); oneof != nil {
		out.DeploymentTarget = &pb.Target_Gke{Gke: oneof}
	}
	if oneof := AnthosCluster_ToProto(mapCtx, in.AnthosCluster); oneof != nil {
		out.DeploymentTarget = &pb.Target_AnthosCluster{AnthosCluster: oneof}
	}
	if oneof := CloudRunLocation_ToProto(mapCtx, in.Run); oneof != nil {
		out.DeploymentTarget = &pb.Target_Run{Run: oneof}
	}
	if oneof := MultiTarget_ToProto(mapCtx, in.MultiTarget); oneof != nil {
		out.DeploymentTarget = &pb.Target_MultiTarget{MultiTarget: oneof}
	}
	if oneof := CustomTarget_ToProto(mapCtx, in.CustomTarget); oneof != nil {
		out.DeploymentTarget = &pb.Target_CustomTarget{CustomTarget: oneof}
	}
	// MISSING: AssociatedEntities
	out.Etag = direct.ValueOf(in.Etag)
	out.ExecutionConfigs = direct.Slice_ToProto(mapCtx, in.ExecutionConfigs, ExecutionConfig_ToProto)
	out.DeployParameters = in.DeployParameters
	return out
}
func TargetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Target) *krm.TargetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TargetObservedState{}
	// MISSING: Name
	out.TargetID = direct.LazyPtr(in.GetTargetId())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: RequireApproval
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Gke
	// MISSING: AnthosCluster
	// MISSING: Run
	// MISSING: MultiTarget
	// MISSING: CustomTarget
	// MISSING: AssociatedEntities
	// MISSING: Etag
	// MISSING: ExecutionConfigs
	// MISSING: DeployParameters
	return out
}
func TargetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TargetObservedState) *pb.Target {
	if in == nil {
		return nil
	}
	out := &pb.Target{}
	// MISSING: Name
	out.TargetId = direct.ValueOf(in.TargetID)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: RequireApproval
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Gke
	// MISSING: AnthosCluster
	// MISSING: Run
	// MISSING: MultiTarget
	// MISSING: CustomTarget
	// MISSING: AssociatedEntities
	// MISSING: Etag
	// MISSING: ExecutionConfigs
	// MISSING: DeployParameters
	return out
}
