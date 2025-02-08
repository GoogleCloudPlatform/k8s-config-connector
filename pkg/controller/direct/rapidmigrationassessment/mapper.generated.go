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

package rapidmigrationassessment

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/rapidmigrationassessment/apiv1/rapidmigrationassessmentpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/rapidmigrationassessment/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Collector_FromProto(mapCtx *direct.MapContext, in *pb.Collector) *krm.Collector {
	if in == nil {
		return nil
	}
	out := &krm.Collector{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	// MISSING: Bucket
	out.ExpectedAssetCount = direct.LazyPtr(in.GetExpectedAssetCount())
	// MISSING: State
	// MISSING: ClientVersion
	// MISSING: GuestOsScan
	// MISSING: VsphereScan
	out.CollectionDays = direct.LazyPtr(in.GetCollectionDays())
	out.EulaURI = direct.LazyPtr(in.GetEulaUri())
	return out
}
func Collector_ToProto(mapCtx *direct.MapContext, in *krm.Collector) *pb.Collector {
	if in == nil {
		return nil
	}
	out := &pb.Collector{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	// MISSING: Bucket
	out.ExpectedAssetCount = direct.ValueOf(in.ExpectedAssetCount)
	// MISSING: State
	// MISSING: ClientVersion
	// MISSING: GuestOsScan
	// MISSING: VsphereScan
	out.CollectionDays = direct.ValueOf(in.CollectionDays)
	out.EulaUri = direct.ValueOf(in.EulaURI)
	return out
}
func CollectorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Collector) *krm.CollectorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CollectorObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ServiceAccount
	out.Bucket = direct.LazyPtr(in.GetBucket())
	// MISSING: ExpectedAssetCount
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ClientVersion = direct.LazyPtr(in.GetClientVersion())
	out.GuestOsScan = GuestOsScan_FromProto(mapCtx, in.GetGuestOsScan())
	out.VsphereScan = VSphereScan_FromProto(mapCtx, in.GetVsphereScan())
	// MISSING: CollectionDays
	// MISSING: EulaURI
	return out
}
func CollectorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CollectorObservedState) *pb.Collector {
	if in == nil {
		return nil
	}
	out := &pb.Collector{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ServiceAccount
	out.Bucket = direct.ValueOf(in.Bucket)
	// MISSING: ExpectedAssetCount
	out.State = direct.Enum_ToProto[pb.Collector_State](mapCtx, in.State)
	out.ClientVersion = direct.ValueOf(in.ClientVersion)
	out.GuestOsScan = GuestOsScan_ToProto(mapCtx, in.GuestOsScan)
	out.VsphereScan = VSphereScan_ToProto(mapCtx, in.VsphereScan)
	// MISSING: CollectionDays
	// MISSING: EulaURI
	return out
}
func GuestOsScan_FromProto(mapCtx *direct.MapContext, in *pb.GuestOsScan) *krm.GuestOsScan {
	if in == nil {
		return nil
	}
	out := &krm.GuestOsScan{}
	out.CoreSource = direct.LazyPtr(in.GetCoreSource())
	return out
}
func GuestOsScan_ToProto(mapCtx *direct.MapContext, in *krm.GuestOsScan) *pb.GuestOsScan {
	if in == nil {
		return nil
	}
	out := &pb.GuestOsScan{}
	out.CoreSource = direct.ValueOf(in.CoreSource)
	return out
}
func RapidmigrationassessmentCollectorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Collector) *krm.RapidmigrationassessmentCollectorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RapidmigrationassessmentCollectorObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ServiceAccount
	// MISSING: Bucket
	// MISSING: ExpectedAssetCount
	// MISSING: State
	// MISSING: ClientVersion
	// MISSING: GuestOsScan
	// MISSING: VsphereScan
	// MISSING: CollectionDays
	// MISSING: EulaURI
	return out
}
func RapidmigrationassessmentCollectorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RapidmigrationassessmentCollectorObservedState) *pb.Collector {
	if in == nil {
		return nil
	}
	out := &pb.Collector{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ServiceAccount
	// MISSING: Bucket
	// MISSING: ExpectedAssetCount
	// MISSING: State
	// MISSING: ClientVersion
	// MISSING: GuestOsScan
	// MISSING: VsphereScan
	// MISSING: CollectionDays
	// MISSING: EulaURI
	return out
}
func RapidmigrationassessmentCollectorSpec_FromProto(mapCtx *direct.MapContext, in *pb.Collector) *krm.RapidmigrationassessmentCollectorSpec {
	if in == nil {
		return nil
	}
	out := &krm.RapidmigrationassessmentCollectorSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ServiceAccount
	// MISSING: Bucket
	// MISSING: ExpectedAssetCount
	// MISSING: State
	// MISSING: ClientVersion
	// MISSING: GuestOsScan
	// MISSING: VsphereScan
	// MISSING: CollectionDays
	// MISSING: EulaURI
	return out
}
func RapidmigrationassessmentCollectorSpec_ToProto(mapCtx *direct.MapContext, in *krm.RapidmigrationassessmentCollectorSpec) *pb.Collector {
	if in == nil {
		return nil
	}
	out := &pb.Collector{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ServiceAccount
	// MISSING: Bucket
	// MISSING: ExpectedAssetCount
	// MISSING: State
	// MISSING: ClientVersion
	// MISSING: GuestOsScan
	// MISSING: VsphereScan
	// MISSING: CollectionDays
	// MISSING: EulaURI
	return out
}
func VSphereScan_FromProto(mapCtx *direct.MapContext, in *pb.VSphereScan) *krm.VSphereScan {
	if in == nil {
		return nil
	}
	out := &krm.VSphereScan{}
	out.CoreSource = direct.LazyPtr(in.GetCoreSource())
	return out
}
func VSphereScan_ToProto(mapCtx *direct.MapContext, in *krm.VSphereScan) *pb.VSphereScan {
	if in == nil {
		return nil
	}
	out := &pb.VSphereScan{}
	out.CoreSource = direct.ValueOf(in.CoreSource)
	return out
}
