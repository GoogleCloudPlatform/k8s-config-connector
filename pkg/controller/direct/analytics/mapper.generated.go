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

package analytics

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/analytics/admin/apiv1beta/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/analytics/v1alpha1"
)
func Account_FromProto(mapCtx *direct.MapContext, in *pb.Account) *krm.Account {
	if in == nil {
		return nil
	}
	out := &krm.Account{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.RegionCode = direct.LazyPtr(in.GetRegionCode())
	// MISSING: Deleted
	// MISSING: GmpOrganization
	return out
}
func Account_ToProto(mapCtx *direct.MapContext, in *krm.Account) *pb.Account {
	if in == nil {
		return nil
	}
	out := &pb.Account{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.RegionCode = direct.ValueOf(in.RegionCode)
	// MISSING: Deleted
	// MISSING: GmpOrganization
	return out
}
func AccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Account) *krm.AccountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccountObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: DisplayName
	// MISSING: RegionCode
	out.Deleted = direct.LazyPtr(in.GetDeleted())
	out.GmpOrganization = direct.LazyPtr(in.GetGmpOrganization())
	return out
}
func AccountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccountObservedState) *pb.Account {
	if in == nil {
		return nil
	}
	out := &pb.Account{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: DisplayName
	// MISSING: RegionCode
	out.Deleted = direct.ValueOf(in.Deleted)
	out.GmpOrganization = direct.ValueOf(in.GmpOrganization)
	return out
}
func AnalyticsAccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Account) *krm.AnalyticsAccountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsAccountObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: RegionCode
	// MISSING: Deleted
	// MISSING: GmpOrganization
	return out
}
func AnalyticsAccountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsAccountObservedState) *pb.Account {
	if in == nil {
		return nil
	}
	out := &pb.Account{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: RegionCode
	// MISSING: Deleted
	// MISSING: GmpOrganization
	return out
}
func AnalyticsAccountSpec_FromProto(mapCtx *direct.MapContext, in *pb.Account) *krm.AnalyticsAccountSpec {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsAccountSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: RegionCode
	// MISSING: Deleted
	// MISSING: GmpOrganization
	return out
}
func AnalyticsAccountSpec_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsAccountSpec) *pb.Account {
	if in == nil {
		return nil
	}
	out := &pb.Account{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: RegionCode
	// MISSING: Deleted
	// MISSING: GmpOrganization
	return out
}
