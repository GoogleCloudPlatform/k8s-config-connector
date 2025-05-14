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

import (
	pb "cloud.google.com/go/workstations/apiv1/workstationspb"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func WorkstationClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationClusterSpec) *pb.WorkstationCluster {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationCluster{
		DisplayName:          direct.ValueOf(in.DisplayName),
		Annotations:          WorkstationAnnotations_ToProto(mapCtx, in.Annotations),
		Labels:               WorkstationLabels_ToProto(mapCtx, in.Labels),
		Network:              in.NetworkRef.External,
		Subnetwork:           in.SubnetworkRef.External,
		PrivateClusterConfig: WorkstationCluster_PrivateClusterConfig_ToProto(mapCtx, in.PrivateClusterConfig),
	}
	return out
}

func WorkstationCluster_PrivateClusterConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationCluster_PrivateClusterConfig) *pb.WorkstationCluster_PrivateClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationCluster_PrivateClusterConfig{
		EnablePrivateEndpoint: direct.ValueOf(in.EnablePrivateEndpoint),
		AllowedProjects:       WorkstationClusterAllowedProjects_ToProto(mapCtx, in.AllowedProjects),
	}
	return out
}

func WorkstationClusterAllowedProjects_ToProto(mapCtx *direct.MapContext, in []refs.ProjectRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, p := range in {
		out = append(out, p.External)
	}
	return out
}

func WorkstationClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationCluster) *krm.WorkstationClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationClusterSpec{
		DisplayName: direct.LazyPtr(in.GetDisplayName()),
		Annotations: WorkstationAnnotations_FromProto(mapCtx, in.GetAnnotations()),
		Labels:      WorkstationLabels_FromProto(mapCtx, in.GetLabels()),
		NetworkRef: computev1beta1.ComputeNetworkRef{
			External: in.GetNetwork(),
		},
		SubnetworkRef: computev1beta1.ComputeSubnetworkRef{
			External: in.GetSubnetwork(),
		},
		PrivateClusterConfig: WorkstationCluster_PrivateClusterConfig_FromProto(mapCtx, in.GetPrivateClusterConfig()),
	}
	return out
}

func WorkstationCluster_PrivateClusterConfig_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationCluster_PrivateClusterConfig) *krm.WorkstationCluster_PrivateClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationCluster_PrivateClusterConfig{
		EnablePrivateEndpoint: direct.LazyPtr(in.GetEnablePrivateEndpoint()),
		AllowedProjects:       WorkstationClusterAllowedProjects_FromProto(mapCtx, in.AllowedProjects),
	}
	return out
}

func WorkstationClusterAllowedProjects_FromProto(mapCtx *direct.MapContext, in []string) []refs.ProjectRef {
	if in == nil {
		return nil
	}
	var out []refs.ProjectRef
	for _, p := range in {
		out = append(out, refs.ProjectRef{
			External: p,
		})
	}
	return out
}

func WorkstationClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationCluster) *krm.WorkstationClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationClusterObservedState{
		Uid:                  direct.LazyPtr(in.GetUid()),
		Reconciling:          direct.LazyPtr(in.GetReconciling()),
		CreateTime:           direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime()),
		UpdateTime:           direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime()),
		DeleteTime:           direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime()),
		Etag:                 direct.LazyPtr(in.GetEtag()),
		ControlPlaneIP:       direct.LazyPtr(in.GetControlPlaneIp()),
		ClusterHostname:      WorkstationClusterClusterHostname_FromProto(mapCtx, in.PrivateClusterConfig),
		ServiceAttachmentURI: WorkstationClusterServiceAttachmentUri_FromProto(mapCtx, in.PrivateClusterConfig),
		Degraded:             direct.LazyPtr(in.GetDegraded()),
		GCPConditions:        WorkstationGCPConditions_FromProto(mapCtx, in.GetConditions()),
	}
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
	out.ControlPlaneIp = direct.ValueOf(in.ControlPlaneIP)
	out.Degraded = direct.ValueOf(in.Degraded)
	out.Conditions = WorkstationGCPConditions_ToProto(mapCtx, in.GCPConditions)
	return out
}

func WorkstationClusterClusterHostname_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationCluster_PrivateClusterConfig) *string {
	if in == nil {
		return nil
	}
	return direct.LazyPtr(in.GetClusterHostname())
}

func WorkstationClusterServiceAttachmentUri_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationCluster_PrivateClusterConfig) *string {
	if in == nil {
		return nil
	}
	return direct.LazyPtr(in.GetServiceAttachmentUri())
}
