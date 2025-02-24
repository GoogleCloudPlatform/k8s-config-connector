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

package storagetransfer

import (
	pb "cloud.google.com/go/storagetransfer/apiv1/storagetransferpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storagetransfer/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AgentPool_BandwidthLimit_FromProto(mapCtx *direct.MapContext, in *pb.AgentPool_BandwidthLimit) *krm.AgentPool_BandwidthLimit {
	if in == nil {
		return nil
	}
	out := &krm.AgentPool_BandwidthLimit{}
	out.LimitMbps = direct.LazyPtr(in.GetLimitMbps())
	return out
}
func AgentPool_BandwidthLimit_ToProto(mapCtx *direct.MapContext, in *krm.AgentPool_BandwidthLimit) *pb.AgentPool_BandwidthLimit {
	if in == nil {
		return nil
	}
	out := &pb.AgentPool_BandwidthLimit{}
	out.LimitMbps = direct.ValueOf(in.LimitMbps)
	return out
}
func StorageTransferAgentPoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AgentPool) *krm.StorageTransferAgentPoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StorageTransferAgentPoolObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: BandwidthLimit
	return out
}
func StorageTransferAgentPoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StorageTransferAgentPoolObservedState) *pb.AgentPool {
	if in == nil {
		return nil
	}
	out := &pb.AgentPool{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: BandwidthLimit
	return out
}
func StorageTransferAgentPoolSpec_FromProto(mapCtx *direct.MapContext, in *pb.AgentPool) *krm.StorageTransferAgentPoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.StorageTransferAgentPoolSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: BandwidthLimit
	return out
}
func StorageTransferAgentPoolSpec_ToProto(mapCtx *direct.MapContext, in *krm.StorageTransferAgentPoolSpec) *pb.AgentPool {
	if in == nil {
		return nil
	}
	out := &pb.AgentPool{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: BandwidthLimit
	return out
}
