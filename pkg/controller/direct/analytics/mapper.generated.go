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
	pb "cloud.google.com/go/analytics/admin/apiv1beta/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/analytics/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
func AnalyticsDataStreamObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataStream) *krm.AnalyticsDataStreamObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsDataStreamObservedState{}
	// MISSING: WebStreamData
	// MISSING: AndroidAppStreamData
	// MISSING: IosAppStreamData
	// MISSING: Name
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AnalyticsDataStreamObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsDataStreamObservedState) *pb.DataStream {
	if in == nil {
		return nil
	}
	out := &pb.DataStream{}
	// MISSING: WebStreamData
	// MISSING: AndroidAppStreamData
	// MISSING: IosAppStreamData
	// MISSING: Name
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AnalyticsDataStreamSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataStream) *krm.AnalyticsDataStreamSpec {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsDataStreamSpec{}
	// MISSING: WebStreamData
	// MISSING: AndroidAppStreamData
	// MISSING: IosAppStreamData
	// MISSING: Name
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AnalyticsDataStreamSpec_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsDataStreamSpec) *pb.DataStream {
	if in == nil {
		return nil
	}
	out := &pb.DataStream{}
	// MISSING: WebStreamData
	// MISSING: AndroidAppStreamData
	// MISSING: IosAppStreamData
	// MISSING: Name
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AnalyticsFirebaseLinkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FirebaseLink) *krm.AnalyticsFirebaseLinkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsFirebaseLinkObservedState{}
	// MISSING: Name
	// MISSING: Project
	// MISSING: CreateTime
	return out
}
func AnalyticsFirebaseLinkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsFirebaseLinkObservedState) *pb.FirebaseLink {
	if in == nil {
		return nil
	}
	out := &pb.FirebaseLink{}
	// MISSING: Name
	// MISSING: Project
	// MISSING: CreateTime
	return out
}
func AnalyticsFirebaseLinkSpec_FromProto(mapCtx *direct.MapContext, in *pb.FirebaseLink) *krm.AnalyticsFirebaseLinkSpec {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsFirebaseLinkSpec{}
	// MISSING: Name
	// MISSING: Project
	// MISSING: CreateTime
	return out
}
func AnalyticsFirebaseLinkSpec_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsFirebaseLinkSpec) *pb.FirebaseLink {
	if in == nil {
		return nil
	}
	out := &pb.FirebaseLink{}
	// MISSING: Name
	// MISSING: Project
	// MISSING: CreateTime
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
func FirebaseLink_FromProto(mapCtx *direct.MapContext, in *pb.FirebaseLink) *krm.FirebaseLink {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseLink{}
	// MISSING: Name
	out.Project = direct.LazyPtr(in.GetProject())
	// MISSING: CreateTime
	return out
}
func FirebaseLink_ToProto(mapCtx *direct.MapContext, in *krm.FirebaseLink) *pb.FirebaseLink {
	if in == nil {
		return nil
	}
	out := &pb.FirebaseLink{}
	// MISSING: Name
	out.Project = direct.ValueOf(in.Project)
	// MISSING: CreateTime
	return out
}
func FirebaseLinkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FirebaseLink) *krm.FirebaseLinkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseLinkObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Project
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func FirebaseLinkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FirebaseLinkObservedState) *pb.FirebaseLink {
	if in == nil {
		return nil
	}
	out := &pb.FirebaseLink{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Project
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
