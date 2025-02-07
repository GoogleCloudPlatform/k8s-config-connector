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
func AnalyticsAccountSummaryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccountSummary) *krm.AnalyticsAccountSummaryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsAccountSummaryObservedState{}
	// MISSING: Name
	// MISSING: Account
	// MISSING: DisplayName
	// MISSING: PropertySummaries
	return out
}
func AnalyticsAccountSummaryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsAccountSummaryObservedState) *pb.AccountSummary {
	if in == nil {
		return nil
	}
	out := &pb.AccountSummary{}
	// MISSING: Name
	// MISSING: Account
	// MISSING: DisplayName
	// MISSING: PropertySummaries
	return out
}
func AnalyticsAccountSummarySpec_FromProto(mapCtx *direct.MapContext, in *pb.AccountSummary) *krm.AnalyticsAccountSummarySpec {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsAccountSummarySpec{}
	// MISSING: Name
	// MISSING: Account
	// MISSING: DisplayName
	// MISSING: PropertySummaries
	return out
}
func AnalyticsAccountSummarySpec_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsAccountSummarySpec) *pb.AccountSummary {
	if in == nil {
		return nil
	}
	out := &pb.AccountSummary{}
	// MISSING: Name
	// MISSING: Account
	// MISSING: DisplayName
	// MISSING: PropertySummaries
	return out
}
func AnalyticsConversionEventObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversionEvent) *krm.AnalyticsConversionEventObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsConversionEventObservedState{}
	// MISSING: Name
	// MISSING: EventName
	// MISSING: CreateTime
	// MISSING: Deletable
	// MISSING: Custom
	// MISSING: CountingMethod
	// MISSING: DefaultConversionValue
	return out
}
func AnalyticsConversionEventObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsConversionEventObservedState) *pb.ConversionEvent {
	if in == nil {
		return nil
	}
	out := &pb.ConversionEvent{}
	// MISSING: Name
	// MISSING: EventName
	// MISSING: CreateTime
	// MISSING: Deletable
	// MISSING: Custom
	// MISSING: CountingMethod
	// MISSING: DefaultConversionValue
	return out
}
func AnalyticsConversionEventSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConversionEvent) *krm.AnalyticsConversionEventSpec {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsConversionEventSpec{}
	// MISSING: Name
	// MISSING: EventName
	// MISSING: CreateTime
	// MISSING: Deletable
	// MISSING: Custom
	// MISSING: CountingMethod
	// MISSING: DefaultConversionValue
	return out
}
func AnalyticsConversionEventSpec_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsConversionEventSpec) *pb.ConversionEvent {
	if in == nil {
		return nil
	}
	out := &pb.ConversionEvent{}
	// MISSING: Name
	// MISSING: EventName
	// MISSING: CreateTime
	// MISSING: Deletable
	// MISSING: Custom
	// MISSING: CountingMethod
	// MISSING: DefaultConversionValue
	return out
}
func AnalyticsCustomDimensionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomDimension) *krm.AnalyticsCustomDimensionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsCustomDimensionObservedState{}
	// MISSING: Name
	// MISSING: ParameterName
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Scope
	// MISSING: DisallowAdsPersonalization
	return out
}
func AnalyticsCustomDimensionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsCustomDimensionObservedState) *pb.CustomDimension {
	if in == nil {
		return nil
	}
	out := &pb.CustomDimension{}
	// MISSING: Name
	// MISSING: ParameterName
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Scope
	// MISSING: DisallowAdsPersonalization
	return out
}
func AnalyticsCustomDimensionSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomDimension) *krm.AnalyticsCustomDimensionSpec {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsCustomDimensionSpec{}
	// MISSING: Name
	// MISSING: ParameterName
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Scope
	// MISSING: DisallowAdsPersonalization
	return out
}
func AnalyticsCustomDimensionSpec_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsCustomDimensionSpec) *pb.CustomDimension {
	if in == nil {
		return nil
	}
	out := &pb.CustomDimension{}
	// MISSING: Name
	// MISSING: ParameterName
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Scope
	// MISSING: DisallowAdsPersonalization
	return out
}
func AnalyticsCustomMetricObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomMetric) *krm.AnalyticsCustomMetricObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsCustomMetricObservedState{}
	// MISSING: Name
	// MISSING: ParameterName
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: MeasurementUnit
	// MISSING: Scope
	// MISSING: RestrictedMetricType
	return out
}
func AnalyticsCustomMetricObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsCustomMetricObservedState) *pb.CustomMetric {
	if in == nil {
		return nil
	}
	out := &pb.CustomMetric{}
	// MISSING: Name
	// MISSING: ParameterName
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: MeasurementUnit
	// MISSING: Scope
	// MISSING: RestrictedMetricType
	return out
}
func AnalyticsCustomMetricSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomMetric) *krm.AnalyticsCustomMetricSpec {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsCustomMetricSpec{}
	// MISSING: Name
	// MISSING: ParameterName
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: MeasurementUnit
	// MISSING: Scope
	// MISSING: RestrictedMetricType
	return out
}
func AnalyticsCustomMetricSpec_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsCustomMetricSpec) *pb.CustomMetric {
	if in == nil {
		return nil
	}
	out := &pb.CustomMetric{}
	// MISSING: Name
	// MISSING: ParameterName
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: MeasurementUnit
	// MISSING: Scope
	// MISSING: RestrictedMetricType
	return out
}
func AnalyticsDataRetentionSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataRetentionSettings) *krm.AnalyticsDataRetentionSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsDataRetentionSettingsObservedState{}
	// MISSING: Name
	// MISSING: EventDataRetention
	// MISSING: ResetUserDataOnNewActivity
	return out
}
func AnalyticsDataRetentionSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsDataRetentionSettingsObservedState) *pb.DataRetentionSettings {
	if in == nil {
		return nil
	}
	out := &pb.DataRetentionSettings{}
	// MISSING: Name
	// MISSING: EventDataRetention
	// MISSING: ResetUserDataOnNewActivity
	return out
}
func AnalyticsDataRetentionSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataRetentionSettings) *krm.AnalyticsDataRetentionSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsDataRetentionSettingsSpec{}
	// MISSING: Name
	// MISSING: EventDataRetention
	// MISSING: ResetUserDataOnNewActivity
	return out
}
func AnalyticsDataRetentionSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsDataRetentionSettingsSpec) *pb.DataRetentionSettings {
	if in == nil {
		return nil
	}
	out := &pb.DataRetentionSettings{}
	// MISSING: Name
	// MISSING: EventDataRetention
	// MISSING: ResetUserDataOnNewActivity
	return out
}
func AnalyticsDataSharingSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataSharingSettings) *krm.AnalyticsDataSharingSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsDataSharingSettingsObservedState{}
	// MISSING: Name
	// MISSING: SharingWithGoogleSupportEnabled
	// MISSING: SharingWithGoogleAssignedSalesEnabled
	// MISSING: SharingWithGoogleAnySalesEnabled
	// MISSING: SharingWithGoogleProductsEnabled
	// MISSING: SharingWithOthersEnabled
	return out
}
func AnalyticsDataSharingSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsDataSharingSettingsObservedState) *pb.DataSharingSettings {
	if in == nil {
		return nil
	}
	out := &pb.DataSharingSettings{}
	// MISSING: Name
	// MISSING: SharingWithGoogleSupportEnabled
	// MISSING: SharingWithGoogleAssignedSalesEnabled
	// MISSING: SharingWithGoogleAnySalesEnabled
	// MISSING: SharingWithGoogleProductsEnabled
	// MISSING: SharingWithOthersEnabled
	return out
}
func AnalyticsDataSharingSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataSharingSettings) *krm.AnalyticsDataSharingSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsDataSharingSettingsSpec{}
	// MISSING: Name
	// MISSING: SharingWithGoogleSupportEnabled
	// MISSING: SharingWithGoogleAssignedSalesEnabled
	// MISSING: SharingWithGoogleAnySalesEnabled
	// MISSING: SharingWithGoogleProductsEnabled
	// MISSING: SharingWithOthersEnabled
	return out
}
func AnalyticsDataSharingSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsDataSharingSettingsSpec) *pb.DataSharingSettings {
	if in == nil {
		return nil
	}
	out := &pb.DataSharingSettings{}
	// MISSING: Name
	// MISSING: SharingWithGoogleSupportEnabled
	// MISSING: SharingWithGoogleAssignedSalesEnabled
	// MISSING: SharingWithGoogleAnySalesEnabled
	// MISSING: SharingWithGoogleProductsEnabled
	// MISSING: SharingWithOthersEnabled
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
func AnalyticsGoogleAdsLinkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GoogleAdsLink) *krm.AnalyticsGoogleAdsLinkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsGoogleAdsLinkObservedState{}
	// MISSING: Name
	// MISSING: CustomerID
	// MISSING: CanManageClients
	// MISSING: AdsPersonalizationEnabled
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CreatorEmailAddress
	return out
}
func AnalyticsGoogleAdsLinkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsGoogleAdsLinkObservedState) *pb.GoogleAdsLink {
	if in == nil {
		return nil
	}
	out := &pb.GoogleAdsLink{}
	// MISSING: Name
	// MISSING: CustomerID
	// MISSING: CanManageClients
	// MISSING: AdsPersonalizationEnabled
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CreatorEmailAddress
	return out
}
func AnalyticsGoogleAdsLinkSpec_FromProto(mapCtx *direct.MapContext, in *pb.GoogleAdsLink) *krm.AnalyticsGoogleAdsLinkSpec {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsGoogleAdsLinkSpec{}
	// MISSING: Name
	// MISSING: CustomerID
	// MISSING: CanManageClients
	// MISSING: AdsPersonalizationEnabled
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CreatorEmailAddress
	return out
}
func AnalyticsGoogleAdsLinkSpec_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsGoogleAdsLinkSpec) *pb.GoogleAdsLink {
	if in == nil {
		return nil
	}
	out := &pb.GoogleAdsLink{}
	// MISSING: Name
	// MISSING: CustomerID
	// MISSING: CanManageClients
	// MISSING: AdsPersonalizationEnabled
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: CreatorEmailAddress
	return out
}
func AnalyticsKeyEventObservedState_FromProto(mapCtx *direct.MapContext, in *pb.KeyEvent) *krm.AnalyticsKeyEventObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsKeyEventObservedState{}
	// MISSING: Name
	// MISSING: EventName
	// MISSING: CreateTime
	// MISSING: Deletable
	// MISSING: Custom
	// MISSING: CountingMethod
	// MISSING: DefaultValue
	return out
}
func AnalyticsKeyEventObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsKeyEventObservedState) *pb.KeyEvent {
	if in == nil {
		return nil
	}
	out := &pb.KeyEvent{}
	// MISSING: Name
	// MISSING: EventName
	// MISSING: CreateTime
	// MISSING: Deletable
	// MISSING: Custom
	// MISSING: CountingMethod
	// MISSING: DefaultValue
	return out
}
func AnalyticsKeyEventSpec_FromProto(mapCtx *direct.MapContext, in *pb.KeyEvent) *krm.AnalyticsKeyEventSpec {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsKeyEventSpec{}
	// MISSING: Name
	// MISSING: EventName
	// MISSING: CreateTime
	// MISSING: Deletable
	// MISSING: Custom
	// MISSING: CountingMethod
	// MISSING: DefaultValue
	return out
}
func AnalyticsKeyEventSpec_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsKeyEventSpec) *pb.KeyEvent {
	if in == nil {
		return nil
	}
	out := &pb.KeyEvent{}
	// MISSING: Name
	// MISSING: EventName
	// MISSING: CreateTime
	// MISSING: Deletable
	// MISSING: Custom
	// MISSING: CountingMethod
	// MISSING: DefaultValue
	return out
}
func AnalyticsMeasurementProtocolSecretObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MeasurementProtocolSecret) *krm.AnalyticsMeasurementProtocolSecretObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsMeasurementProtocolSecretObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SecretValue
	return out
}
func AnalyticsMeasurementProtocolSecretObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsMeasurementProtocolSecretObservedState) *pb.MeasurementProtocolSecret {
	if in == nil {
		return nil
	}
	out := &pb.MeasurementProtocolSecret{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SecretValue
	return out
}
func AnalyticsMeasurementProtocolSecretSpec_FromProto(mapCtx *direct.MapContext, in *pb.MeasurementProtocolSecret) *krm.AnalyticsMeasurementProtocolSecretSpec {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsMeasurementProtocolSecretSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SecretValue
	return out
}
func AnalyticsMeasurementProtocolSecretSpec_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsMeasurementProtocolSecretSpec) *pb.MeasurementProtocolSecret {
	if in == nil {
		return nil
	}
	out := &pb.MeasurementProtocolSecret{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SecretValue
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
func DataRetentionSettings_FromProto(mapCtx *direct.MapContext, in *pb.DataRetentionSettings) *krm.DataRetentionSettings {
	if in == nil {
		return nil
	}
	out := &krm.DataRetentionSettings{}
	// MISSING: Name
	out.EventDataRetention = direct.Enum_FromProto(mapCtx, in.GetEventDataRetention())
	out.ResetUserDataOnNewActivity = direct.LazyPtr(in.GetResetUserDataOnNewActivity())
	return out
}
func DataRetentionSettings_ToProto(mapCtx *direct.MapContext, in *krm.DataRetentionSettings) *pb.DataRetentionSettings {
	if in == nil {
		return nil
	}
	out := &pb.DataRetentionSettings{}
	// MISSING: Name
	out.EventDataRetention = direct.Enum_ToProto[pb.DataRetentionSettings_RetentionDuration](mapCtx, in.EventDataRetention)
	out.ResetUserDataOnNewActivity = direct.ValueOf(in.ResetUserDataOnNewActivity)
	return out
}
func DataRetentionSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataRetentionSettings) *krm.DataRetentionSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataRetentionSettingsObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: EventDataRetention
	// MISSING: ResetUserDataOnNewActivity
	return out
}
func DataRetentionSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataRetentionSettingsObservedState) *pb.DataRetentionSettings {
	if in == nil {
		return nil
	}
	out := &pb.DataRetentionSettings{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: EventDataRetention
	// MISSING: ResetUserDataOnNewActivity
	return out
}
