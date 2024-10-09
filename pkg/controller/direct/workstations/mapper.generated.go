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

package workstations

/*
import (
	pb "cloud.google.com/go/workstations/apiv1/workstationspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func WorkstationClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationCluster) *krm.WorkstationClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationClusterObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: Network
	// MISSING: Subnetwork
	out.ControlPlaneIp = direct.LazyPtr(in.GetControlPlaneIp())
	out.Degraded = direct.LazyPtr(in.GetDegraded())
	// MISSING: Conditions
	return out
}
func WorkstationClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationClusterObservedState) *pb.WorkstationCluster {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationCluster{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: Network
	// MISSING: Subnetwork
	out.ControlPlaneIp = direct.ValueOf(in.ControlPlaneIp)
	out.Degraded = direct.ValueOf(in.Degraded)
	// MISSING: Conditions
	return out
}
func WorkstationClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationCluster) *krm.WorkstationClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationClusterSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	// MISSING: Network
	// MISSING: Subnetwork
	out.PrivateClusterConfig = WorkstationCluster_PrivateClusterConfig_FromProto(mapCtx, in.GetPrivateClusterConfig())
	// MISSING: Conditions
	return out
}
func WorkstationClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationClusterSpec) *pb.WorkstationCluster {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationCluster{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	// MISSING: Network
	// MISSING: Subnetwork
	out.PrivateClusterConfig = WorkstationCluster_PrivateClusterConfig_ToProto(mapCtx, in.PrivateClusterConfig)
	// MISSING: Conditions
	return out
}
func WorkstationCluster_PrivateClusterConfig_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationCluster_PrivateClusterConfig) *krm.WorkstationCluster_PrivateClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationCluster_PrivateClusterConfig{}
	out.EnablePrivateEndpoint = direct.LazyPtr(in.GetEnablePrivateEndpoint())
	// MISSING: ClusterHostname
	// MISSING: ServiceAttachmentUri
	out.AllowedProjects = in.AllowedProjects
	return out
}
func WorkstationCluster_PrivateClusterConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationCluster_PrivateClusterConfig) *pb.WorkstationCluster_PrivateClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationCluster_PrivateClusterConfig{}
	out.EnablePrivateEndpoint = direct.ValueOf(in.EnablePrivateEndpoint)
	// MISSING: ClusterHostname
	// MISSING: ServiceAttachmentUri
	out.AllowedProjects = in.AllowedProjects
	return out
}
*/
