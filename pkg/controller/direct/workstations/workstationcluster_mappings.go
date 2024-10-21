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
	status "google.golang.org/genproto/googleapis/rpc/status"

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
		Annotations:          WorkstationClusterAnnotations_ToProto(mapCtx, in.Annotations),
		Labels:               WorkstationClusterLabels_ToProto(mapCtx, in.Labels),
		Network:              in.NetworkRef.External,
		Subnetwork:           in.SubnetworkRef.External,
		PrivateClusterConfig: WorkstationCluster_PrivateClusterConfig_ToProto(mapCtx, in.PrivateClusterConfig),
	}
	return out
}

func WorkstationClusterAnnotations_ToProto(mapCtx *direct.MapContext, in []krm.WorkstationClusterAnnotation) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string)
	for _, a := range in {
		out[a.Key] = a.Value
	}
	return out
}

func WorkstationClusterLabels_ToProto(mapCtx *direct.MapContext, in []krm.WorkstationClusterLabel) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string)
	for _, l := range in {
		out[l.Key] = l.Value
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
		Annotations: WorkstationClusterAnnotations_FromProto(mapCtx, in.GetAnnotations()),
		Labels:      WorkstationClusterLabels_FromProto(mapCtx, in.GetLabels()),
		NetworkRef: refs.ComputeNetworkRef{
			External: in.GetNetwork(),
		},
		SubnetworkRef: refs.ComputeSubnetworkRef{
			External: in.GetSubnetwork(),
		},
		PrivateClusterConfig: WorkstationCluster_PrivateClusterConfig_FromProto(mapCtx, in.GetPrivateClusterConfig()),
	}
	return out
}

func WorkstationClusterAnnotations_FromProto(mapCtx *direct.MapContext, in map[string]string) []krm.WorkstationClusterAnnotation {
	if in == nil {
		return nil
	}
	var out []krm.WorkstationClusterAnnotation
	for k, v := range in {
		out = append(out, krm.WorkstationClusterAnnotation{
			Key:   k,
			Value: v,
		})
	}
	return out
}

func WorkstationClusterLabels_FromProto(mapCtx *direct.MapContext, in map[string]string) []krm.WorkstationClusterLabel {
	if in == nil {
		return nil
	}
	var out []krm.WorkstationClusterLabel
	for k, v := range in {
		out = append(out, krm.WorkstationClusterLabel{
			Key:   k,
			Value: v,
		})
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
		GCPConditions:        WorkstationClusterGCPConditions_FromProto(mapCtx, in.GetConditions()),
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
	out.Conditions = WorkstationClusterGCPConditions_ToProto(mapCtx, in.GCPConditions)
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

func WorkstationClusterGCPConditions_FromProto(mapCtx *direct.MapContext, in []*status.Status) []krm.WorkstationClusterGCPCondition {
	if in == nil {
		return nil
	}
	var out []krm.WorkstationClusterGCPCondition
	for _, c := range in {
		out = append(out, krm.WorkstationClusterGCPCondition{
			Code:    direct.LazyPtr(c.Code),
			Message: direct.LazyPtr(c.Message),
		})
	}
	return out
}
func WorkstationClusterGCPConditions_ToProto(mapCtx *direct.MapContext, in []krm.WorkstationClusterGCPCondition) []*status.Status {
	if in == nil {
		return nil
	}
	var out []*status.Status
	for _, c := range in {
		out = append(out, &status.Status{
			Code:    direct.ValueOf(c.Code),
			Message: direct.ValueOf(c.Message),
		})
	}
	return out
}
