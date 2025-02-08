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

package video

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/video/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Pool_FromProto(mapCtx *direct.MapContext, in *pb.Pool) *krm.Pool {
	if in == nil {
		return nil
	}
	out := &krm.Pool{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.NetworkConfig = Pool_NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	return out
}
func Pool_ToProto(mapCtx *direct.MapContext, in *krm.Pool) *pb.Pool {
	if in == nil {
		return nil
	}
	out := &pb.Pool{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.NetworkConfig = Pool_NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	return out
}
func PoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Pool) *krm.PoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PoolObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: NetworkConfig
	return out
}
func PoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PoolObservedState) *pb.Pool {
	if in == nil {
		return nil
	}
	out := &pb.Pool{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: NetworkConfig
	return out
}
func Pool_NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.Pool_NetworkConfig) *krm.Pool_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.Pool_NetworkConfig{}
	out.PeeredNetwork = direct.LazyPtr(in.GetPeeredNetwork())
	return out
}
func Pool_NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.Pool_NetworkConfig) *pb.Pool_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.Pool_NetworkConfig{}
	out.PeeredNetwork = direct.ValueOf(in.PeeredNetwork)
	return out
}
func VideoPoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Pool) *krm.VideoPoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VideoPoolObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: NetworkConfig
	return out
}
func VideoPoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VideoPoolObservedState) *pb.Pool {
	if in == nil {
		return nil
	}
	out := &pb.Pool{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: NetworkConfig
	return out
}
func VideoPoolSpec_FromProto(mapCtx *direct.MapContext, in *pb.Pool) *krm.VideoPoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.VideoPoolSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: NetworkConfig
	return out
}
func VideoPoolSpec_ToProto(mapCtx *direct.MapContext, in *krm.VideoPoolSpec) *pb.Pool {
	if in == nil {
		return nil
	}
	out := &pb.Pool{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: NetworkConfig
	return out
}
