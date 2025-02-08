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

package parallelstore

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/parallelstore/apiv1/parallelstorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/parallelstore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.CapacityGib = direct.LazyPtr(in.GetCapacityGib())
	// MISSING: DaosVersion
	// MISSING: AccessPoints
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.ReservedIPRange = direct.LazyPtr(in.GetReservedIpRange())
	// MISSING: EffectiveReservedIPRange
	out.FileStripeLevel = direct.Enum_FromProto(mapCtx, in.GetFileStripeLevel())
	out.DirectoryStripeLevel = direct.Enum_FromProto(mapCtx, in.GetDirectoryStripeLevel())
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.CapacityGib = direct.ValueOf(in.CapacityGib)
	// MISSING: DaosVersion
	// MISSING: AccessPoints
	out.Network = direct.ValueOf(in.Network)
	out.ReservedIpRange = direct.ValueOf(in.ReservedIPRange)
	// MISSING: EffectiveReservedIPRange
	out.FileStripeLevel = direct.Enum_ToProto[pb.FileStripeLevel](mapCtx, in.FileStripeLevel)
	out.DirectoryStripeLevel = direct.Enum_ToProto[pb.DirectoryStripeLevel](mapCtx, in.DirectoryStripeLevel)
	return out
}
func InstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.InstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceObservedState{}
	// MISSING: Name
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: CapacityGib
	out.DaosVersion = direct.LazyPtr(in.GetDaosVersion())
	out.AccessPoints = in.AccessPoints
	// MISSING: Network
	// MISSING: ReservedIPRange
	out.EffectiveReservedIPRange = direct.LazyPtr(in.GetEffectiveReservedIpRange())
	// MISSING: FileStripeLevel
	// MISSING: DirectoryStripeLevel
	return out
}
func InstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: CapacityGib
	out.DaosVersion = direct.ValueOf(in.DaosVersion)
	out.AccessPoints = in.AccessPoints
	// MISSING: Network
	// MISSING: ReservedIPRange
	out.EffectiveReservedIpRange = direct.ValueOf(in.EffectiveReservedIPRange)
	// MISSING: FileStripeLevel
	// MISSING: DirectoryStripeLevel
	return out
}
func ParallelstoreInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.ParallelstoreInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ParallelstoreInstanceObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: CapacityGib
	// MISSING: DaosVersion
	// MISSING: AccessPoints
	// MISSING: Network
	// MISSING: ReservedIPRange
	// MISSING: EffectiveReservedIPRange
	// MISSING: FileStripeLevel
	// MISSING: DirectoryStripeLevel
	return out
}
func ParallelstoreInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ParallelstoreInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: CapacityGib
	// MISSING: DaosVersion
	// MISSING: AccessPoints
	// MISSING: Network
	// MISSING: ReservedIPRange
	// MISSING: EffectiveReservedIPRange
	// MISSING: FileStripeLevel
	// MISSING: DirectoryStripeLevel
	return out
}
func ParallelstoreInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.ParallelstoreInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ParallelstoreInstanceSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: CapacityGib
	// MISSING: DaosVersion
	// MISSING: AccessPoints
	// MISSING: Network
	// MISSING: ReservedIPRange
	// MISSING: EffectiveReservedIPRange
	// MISSING: FileStripeLevel
	// MISSING: DirectoryStripeLevel
	return out
}
func ParallelstoreInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ParallelstoreInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: CapacityGib
	// MISSING: DaosVersion
	// MISSING: AccessPoints
	// MISSING: Network
	// MISSING: ReservedIPRange
	// MISSING: EffectiveReservedIPRange
	// MISSING: FileStripeLevel
	// MISSING: DirectoryStripeLevel
	return out
}
