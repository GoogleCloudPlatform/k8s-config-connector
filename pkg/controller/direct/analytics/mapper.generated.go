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
func AnalyticsPropertyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Property) *krm.AnalyticsPropertyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsPropertyObservedState{}
	// MISSING: Name
	// MISSING: PropertyType
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Parent
	// MISSING: DisplayName
	// MISSING: IndustryCategory
	// MISSING: TimeZone
	// MISSING: CurrencyCode
	// MISSING: ServiceLevel
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Account
	return out
}
func AnalyticsPropertyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsPropertyObservedState) *pb.Property {
	if in == nil {
		return nil
	}
	out := &pb.Property{}
	// MISSING: Name
	// MISSING: PropertyType
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Parent
	// MISSING: DisplayName
	// MISSING: IndustryCategory
	// MISSING: TimeZone
	// MISSING: CurrencyCode
	// MISSING: ServiceLevel
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Account
	return out
}
func AnalyticsPropertySpec_FromProto(mapCtx *direct.MapContext, in *pb.Property) *krm.AnalyticsPropertySpec {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsPropertySpec{}
	// MISSING: Name
	// MISSING: PropertyType
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Parent
	// MISSING: DisplayName
	// MISSING: IndustryCategory
	// MISSING: TimeZone
	// MISSING: CurrencyCode
	// MISSING: ServiceLevel
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Account
	return out
}
func AnalyticsPropertySpec_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsPropertySpec) *pb.Property {
	if in == nil {
		return nil
	}
	out := &pb.Property{}
	// MISSING: Name
	// MISSING: PropertyType
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Parent
	// MISSING: DisplayName
	// MISSING: IndustryCategory
	// MISSING: TimeZone
	// MISSING: CurrencyCode
	// MISSING: ServiceLevel
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Account
	return out
}
func Property_FromProto(mapCtx *direct.MapContext, in *pb.Property) *krm.Property {
	if in == nil {
		return nil
	}
	out := &krm.Property{}
	// MISSING: Name
	out.PropertyType = direct.Enum_FromProto(mapCtx, in.GetPropertyType())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Parent = direct.LazyPtr(in.GetParent())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.IndustryCategory = direct.Enum_FromProto(mapCtx, in.GetIndustryCategory())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	out.CurrencyCode = direct.LazyPtr(in.GetCurrencyCode())
	// MISSING: ServiceLevel
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	out.Account = direct.LazyPtr(in.GetAccount())
	return out
}
func Property_ToProto(mapCtx *direct.MapContext, in *krm.Property) *pb.Property {
	if in == nil {
		return nil
	}
	out := &pb.Property{}
	// MISSING: Name
	out.PropertyType = direct.Enum_ToProto[pb.PropertyType](mapCtx, in.PropertyType)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Parent = direct.ValueOf(in.Parent)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.IndustryCategory = direct.Enum_ToProto[pb.IndustryCategory](mapCtx, in.IndustryCategory)
	out.TimeZone = direct.ValueOf(in.TimeZone)
	out.CurrencyCode = direct.ValueOf(in.CurrencyCode)
	// MISSING: ServiceLevel
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	out.Account = direct.ValueOf(in.Account)
	return out
}
func PropertyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Property) *krm.PropertyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PropertyObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: PropertyType
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Parent
	// MISSING: DisplayName
	// MISSING: IndustryCategory
	// MISSING: TimeZone
	// MISSING: CurrencyCode
	out.ServiceLevel = direct.Enum_FromProto(mapCtx, in.GetServiceLevel())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	// MISSING: Account
	return out
}
func PropertyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PropertyObservedState) *pb.Property {
	if in == nil {
		return nil
	}
	out := &pb.Property{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: PropertyType
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Parent
	// MISSING: DisplayName
	// MISSING: IndustryCategory
	// MISSING: TimeZone
	// MISSING: CurrencyCode
	out.ServiceLevel = direct.Enum_ToProto[pb.ServiceLevel](mapCtx, in.ServiceLevel)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	// MISSING: Account
	return out
}
