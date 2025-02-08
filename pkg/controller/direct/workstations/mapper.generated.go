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

package workstations

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/workstations/apiv1/workstationspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1alpha1"
)
func WorkstationsWorkstationClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationCluster) *krm.WorkstationsWorkstationClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationsWorkstationClusterObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	// MISSING: Network
	// MISSING: Subnetwork
	// MISSING: ControlPlaneIP
	// MISSING: PrivateClusterConfig
	// MISSING: Degraded
	// MISSING: Conditions
	return out
}
func WorkstationsWorkstationClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationsWorkstationClusterObservedState) *pb.WorkstationCluster {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationCluster{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	// MISSING: Network
	// MISSING: Subnetwork
	// MISSING: ControlPlaneIP
	// MISSING: PrivateClusterConfig
	// MISSING: Degraded
	// MISSING: Conditions
	return out
}
func WorkstationsWorkstationClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.WorkstationCluster) *krm.WorkstationsWorkstationClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationsWorkstationClusterSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	// MISSING: Network
	// MISSING: Subnetwork
	// MISSING: ControlPlaneIP
	// MISSING: PrivateClusterConfig
	// MISSING: Degraded
	// MISSING: Conditions
	return out
}
func WorkstationsWorkstationClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationsWorkstationClusterSpec) *pb.WorkstationCluster {
	if in == nil {
		return nil
	}
	out := &pb.WorkstationCluster{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	// MISSING: Network
	// MISSING: Subnetwork
	// MISSING: ControlPlaneIP
	// MISSING: PrivateClusterConfig
	// MISSING: Degraded
	// MISSING: Conditions
	return out
}
