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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/analytics/admin/apiv1beta/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/analytics/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
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
func DataStream_FromProto(mapCtx *direct.MapContext, in *pb.DataStream) *krm.DataStream {
	if in == nil {
		return nil
	}
	out := &krm.DataStream{}
	out.WebStreamData = DataStream_WebStreamData_FromProto(mapCtx, in.GetWebStreamData())
	out.AndroidAppStreamData = DataStream_AndroidAppStreamData_FromProto(mapCtx, in.GetAndroidAppStreamData())
	out.IosAppStreamData = DataStream_IosAppStreamData_FromProto(mapCtx, in.GetIosAppStreamData())
	// MISSING: Name
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func DataStream_ToProto(mapCtx *direct.MapContext, in *krm.DataStream) *pb.DataStream {
	if in == nil {
		return nil
	}
	out := &pb.DataStream{}
	if oneof := DataStream_WebStreamData_ToProto(mapCtx, in.WebStreamData); oneof != nil {
		out.StreamData = &pb.DataStream_WebStreamData_{WebStreamData: oneof}
	}
	if oneof := DataStream_AndroidAppStreamData_ToProto(mapCtx, in.AndroidAppStreamData); oneof != nil {
		out.StreamData = &pb.DataStream_AndroidAppStreamData_{AndroidAppStreamData: oneof}
	}
	if oneof := DataStream_IosAppStreamData_ToProto(mapCtx, in.IosAppStreamData); oneof != nil {
		out.StreamData = &pb.DataStream_IosAppStreamData_{IosAppStreamData: oneof}
	}
	// MISSING: Name
	out.Type = direct.Enum_ToProto[pb.DataStream_DataStreamType](mapCtx, in.Type)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func DataStreamObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataStream) *krm.DataStreamObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataStreamObservedState{}
	out.WebStreamData = DataStream_WebStreamDataObservedState_FromProto(mapCtx, in.GetWebStreamData())
	out.AndroidAppStreamData = DataStream_AndroidAppStreamDataObservedState_FromProto(mapCtx, in.GetAndroidAppStreamData())
	out.IosAppStreamData = DataStream_IosAppStreamDataObservedState_FromProto(mapCtx, in.GetIosAppStreamData())
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Type
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func DataStreamObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataStreamObservedState) *pb.DataStream {
	if in == nil {
		return nil
	}
	out := &pb.DataStream{}
	if oneof := DataStream_WebStreamDataObservedState_ToProto(mapCtx, in.WebStreamData); oneof != nil {
		out.StreamData = &pb.DataStream_WebStreamData_{WebStreamData: oneof}
	}
	if oneof := DataStream_AndroidAppStreamDataObservedState_ToProto(mapCtx, in.AndroidAppStreamData); oneof != nil {
		out.StreamData = &pb.DataStream_AndroidAppStreamData_{AndroidAppStreamData: oneof}
	}
	if oneof := DataStream_IosAppStreamDataObservedState_ToProto(mapCtx, in.IosAppStreamData); oneof != nil {
		out.StreamData = &pb.DataStream_IosAppStreamData_{IosAppStreamData: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Type
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func DataStream_AndroidAppStreamData_FromProto(mapCtx *direct.MapContext, in *pb.DataStream_AndroidAppStreamData) *krm.DataStream_AndroidAppStreamData {
	if in == nil {
		return nil
	}
	out := &krm.DataStream_AndroidAppStreamData{}
	// MISSING: FirebaseAppID
	out.PackageName = direct.LazyPtr(in.GetPackageName())
	return out
}
func DataStream_AndroidAppStreamData_ToProto(mapCtx *direct.MapContext, in *krm.DataStream_AndroidAppStreamData) *pb.DataStream_AndroidAppStreamData {
	if in == nil {
		return nil
	}
	out := &pb.DataStream_AndroidAppStreamData{}
	// MISSING: FirebaseAppID
	out.PackageName = direct.ValueOf(in.PackageName)
	return out
}
func DataStream_AndroidAppStreamDataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataStream_AndroidAppStreamData) *krm.DataStream_AndroidAppStreamDataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataStream_AndroidAppStreamDataObservedState{}
	out.FirebaseAppID = direct.LazyPtr(in.GetFirebaseAppId())
	// MISSING: PackageName
	return out
}
func DataStream_AndroidAppStreamDataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataStream_AndroidAppStreamDataObservedState) *pb.DataStream_AndroidAppStreamData {
	if in == nil {
		return nil
	}
	out := &pb.DataStream_AndroidAppStreamData{}
	out.FirebaseAppId = direct.ValueOf(in.FirebaseAppID)
	// MISSING: PackageName
	return out
}
func DataStream_IosAppStreamData_FromProto(mapCtx *direct.MapContext, in *pb.DataStream_IosAppStreamData) *krm.DataStream_IosAppStreamData {
	if in == nil {
		return nil
	}
	out := &krm.DataStream_IosAppStreamData{}
	// MISSING: FirebaseAppID
	out.BundleID = direct.LazyPtr(in.GetBundleId())
	return out
}
func DataStream_IosAppStreamData_ToProto(mapCtx *direct.MapContext, in *krm.DataStream_IosAppStreamData) *pb.DataStream_IosAppStreamData {
	if in == nil {
		return nil
	}
	out := &pb.DataStream_IosAppStreamData{}
	// MISSING: FirebaseAppID
	out.BundleId = direct.ValueOf(in.BundleID)
	return out
}
func DataStream_IosAppStreamDataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataStream_IosAppStreamData) *krm.DataStream_IosAppStreamDataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataStream_IosAppStreamDataObservedState{}
	out.FirebaseAppID = direct.LazyPtr(in.GetFirebaseAppId())
	// MISSING: BundleID
	return out
}
func DataStream_IosAppStreamDataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataStream_IosAppStreamDataObservedState) *pb.DataStream_IosAppStreamData {
	if in == nil {
		return nil
	}
	out := &pb.DataStream_IosAppStreamData{}
	out.FirebaseAppId = direct.ValueOf(in.FirebaseAppID)
	// MISSING: BundleID
	return out
}
func DataStream_WebStreamData_FromProto(mapCtx *direct.MapContext, in *pb.DataStream_WebStreamData) *krm.DataStream_WebStreamData {
	if in == nil {
		return nil
	}
	out := &krm.DataStream_WebStreamData{}
	// MISSING: MeasurementID
	// MISSING: FirebaseAppID
	out.DefaultURI = direct.LazyPtr(in.GetDefaultUri())
	return out
}
func DataStream_WebStreamData_ToProto(mapCtx *direct.MapContext, in *krm.DataStream_WebStreamData) *pb.DataStream_WebStreamData {
	if in == nil {
		return nil
	}
	out := &pb.DataStream_WebStreamData{}
	// MISSING: MeasurementID
	// MISSING: FirebaseAppID
	out.DefaultUri = direct.ValueOf(in.DefaultURI)
	return out
}
func DataStream_WebStreamDataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataStream_WebStreamData) *krm.DataStream_WebStreamDataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataStream_WebStreamDataObservedState{}
	out.MeasurementID = direct.LazyPtr(in.GetMeasurementId())
	out.FirebaseAppID = direct.LazyPtr(in.GetFirebaseAppId())
	// MISSING: DefaultURI
	return out
}
func DataStream_WebStreamDataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataStream_WebStreamDataObservedState) *pb.DataStream_WebStreamData {
	if in == nil {
		return nil
	}
	out := &pb.DataStream_WebStreamData{}
	out.MeasurementId = direct.ValueOf(in.MeasurementID)
	out.FirebaseAppId = direct.ValueOf(in.FirebaseAppID)
	// MISSING: DefaultURI
	return out
}
