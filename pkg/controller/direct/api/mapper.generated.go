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

package api

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/api/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/cloudquotas/apiv1/cloudquotaspb"
)
func ApiQuotaInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QuotaInfo) *krm.ApiQuotaInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApiQuotaInfoObservedState{}
	// MISSING: Name
	// MISSING: QuotaID
	// MISSING: Metric
	// MISSING: Service
	// MISSING: IsPrecise
	// MISSING: RefreshInterval
	// MISSING: ContainerType
	// MISSING: Dimensions
	// MISSING: MetricDisplayName
	// MISSING: QuotaDisplayName
	// MISSING: MetricUnit
	// MISSING: QuotaIncreaseEligibility
	// MISSING: IsFixed
	// MISSING: DimensionsInfos
	// MISSING: IsConcurrent
	// MISSING: ServiceRequestQuotaURI
	return out
}
func ApiQuotaInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApiQuotaInfoObservedState) *pb.QuotaInfo {
	if in == nil {
		return nil
	}
	out := &pb.QuotaInfo{}
	// MISSING: Name
	// MISSING: QuotaID
	// MISSING: Metric
	// MISSING: Service
	// MISSING: IsPrecise
	// MISSING: RefreshInterval
	// MISSING: ContainerType
	// MISSING: Dimensions
	// MISSING: MetricDisplayName
	// MISSING: QuotaDisplayName
	// MISSING: MetricUnit
	// MISSING: QuotaIncreaseEligibility
	// MISSING: IsFixed
	// MISSING: DimensionsInfos
	// MISSING: IsConcurrent
	// MISSING: ServiceRequestQuotaURI
	return out
}
func ApiQuotaInfoSpec_FromProto(mapCtx *direct.MapContext, in *pb.QuotaInfo) *krm.ApiQuotaInfoSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApiQuotaInfoSpec{}
	// MISSING: Name
	// MISSING: QuotaID
	// MISSING: Metric
	// MISSING: Service
	// MISSING: IsPrecise
	// MISSING: RefreshInterval
	// MISSING: ContainerType
	// MISSING: Dimensions
	// MISSING: MetricDisplayName
	// MISSING: QuotaDisplayName
	// MISSING: MetricUnit
	// MISSING: QuotaIncreaseEligibility
	// MISSING: IsFixed
	// MISSING: DimensionsInfos
	// MISSING: IsConcurrent
	// MISSING: ServiceRequestQuotaURI
	return out
}
func ApiQuotaInfoSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApiQuotaInfoSpec) *pb.QuotaInfo {
	if in == nil {
		return nil
	}
	out := &pb.QuotaInfo{}
	// MISSING: Name
	// MISSING: QuotaID
	// MISSING: Metric
	// MISSING: Service
	// MISSING: IsPrecise
	// MISSING: RefreshInterval
	// MISSING: ContainerType
	// MISSING: Dimensions
	// MISSING: MetricDisplayName
	// MISSING: QuotaDisplayName
	// MISSING: MetricUnit
	// MISSING: QuotaIncreaseEligibility
	// MISSING: IsFixed
	// MISSING: DimensionsInfos
	// MISSING: IsConcurrent
	// MISSING: ServiceRequestQuotaURI
	return out
}
func ApiQuotaPreferenceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QuotaPreference) *krm.ApiQuotaPreferenceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApiQuotaPreferenceObservedState{}
	// MISSING: Name
	// MISSING: Dimensions
	// MISSING: QuotaConfig
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Service
	// MISSING: QuotaID
	// MISSING: Reconciling
	// MISSING: Justification
	// MISSING: ContactEmail
	return out
}
func ApiQuotaPreferenceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApiQuotaPreferenceObservedState) *pb.QuotaPreference {
	if in == nil {
		return nil
	}
	out := &pb.QuotaPreference{}
	// MISSING: Name
	// MISSING: Dimensions
	// MISSING: QuotaConfig
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Service
	// MISSING: QuotaID
	// MISSING: Reconciling
	// MISSING: Justification
	// MISSING: ContactEmail
	return out
}
func ApiQuotaPreferenceSpec_FromProto(mapCtx *direct.MapContext, in *pb.QuotaPreference) *krm.ApiQuotaPreferenceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApiQuotaPreferenceSpec{}
	// MISSING: Name
	// MISSING: Dimensions
	// MISSING: QuotaConfig
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Service
	// MISSING: QuotaID
	// MISSING: Reconciling
	// MISSING: Justification
	// MISSING: ContactEmail
	return out
}
func ApiQuotaPreferenceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApiQuotaPreferenceSpec) *pb.QuotaPreference {
	if in == nil {
		return nil
	}
	out := &pb.QuotaPreference{}
	// MISSING: Name
	// MISSING: Dimensions
	// MISSING: QuotaConfig
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Service
	// MISSING: QuotaID
	// MISSING: Reconciling
	// MISSING: Justification
	// MISSING: ContactEmail
	return out
}
func QuotaConfig_FromProto(mapCtx *direct.MapContext, in *pb.QuotaConfig) *krm.QuotaConfig {
	if in == nil {
		return nil
	}
	out := &krm.QuotaConfig{}
	out.PreferredValue = direct.LazyPtr(in.GetPreferredValue())
	// MISSING: StateDetail
	// MISSING: GrantedValue
	// MISSING: TraceID
	out.Annotations = in.Annotations
	// MISSING: RequestOrigin
	return out
}
func QuotaConfig_ToProto(mapCtx *direct.MapContext, in *krm.QuotaConfig) *pb.QuotaConfig {
	if in == nil {
		return nil
	}
	out := &pb.QuotaConfig{}
	out.PreferredValue = direct.ValueOf(in.PreferredValue)
	// MISSING: StateDetail
	// MISSING: GrantedValue
	// MISSING: TraceID
	out.Annotations = in.Annotations
	// MISSING: RequestOrigin
	return out
}
func QuotaConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QuotaConfig) *krm.QuotaConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QuotaConfigObservedState{}
	// MISSING: PreferredValue
	out.StateDetail = direct.LazyPtr(in.GetStateDetail())
	out.GrantedValue = direct.Int64Value_FromProto(mapCtx, in.GetGrantedValue())
	out.TraceID = direct.LazyPtr(in.GetTraceId())
	// MISSING: Annotations
	out.RequestOrigin = direct.Enum_FromProto(mapCtx, in.GetRequestOrigin())
	return out
}
func QuotaConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QuotaConfigObservedState) *pb.QuotaConfig {
	if in == nil {
		return nil
	}
	out := &pb.QuotaConfig{}
	// MISSING: PreferredValue
	out.StateDetail = direct.ValueOf(in.StateDetail)
	out.GrantedValue = direct.Int64Value_ToProto(mapCtx, in.GrantedValue)
	out.TraceId = direct.ValueOf(in.TraceID)
	// MISSING: Annotations
	out.RequestOrigin = direct.Enum_ToProto[pb.QuotaConfig_Origin](mapCtx, in.RequestOrigin)
	return out
}
func QuotaPreference_FromProto(mapCtx *direct.MapContext, in *pb.QuotaPreference) *krm.QuotaPreference {
	if in == nil {
		return nil
	}
	out := &krm.QuotaPreference{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Dimensions = in.Dimensions
	out.QuotaConfig = QuotaConfig_FromProto(mapCtx, in.GetQuotaConfig())
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Service = direct.LazyPtr(in.GetService())
	out.QuotaID = direct.LazyPtr(in.GetQuotaId())
	// MISSING: Reconciling
	out.Justification = direct.LazyPtr(in.GetJustification())
	out.ContactEmail = direct.LazyPtr(in.GetContactEmail())
	return out
}
func QuotaPreference_ToProto(mapCtx *direct.MapContext, in *krm.QuotaPreference) *pb.QuotaPreference {
	if in == nil {
		return nil
	}
	out := &pb.QuotaPreference{}
	out.Name = direct.ValueOf(in.Name)
	out.Dimensions = in.Dimensions
	out.QuotaConfig = QuotaConfig_ToProto(mapCtx, in.QuotaConfig)
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Service = direct.ValueOf(in.Service)
	out.QuotaId = direct.ValueOf(in.QuotaID)
	// MISSING: Reconciling
	out.Justification = direct.ValueOf(in.Justification)
	out.ContactEmail = direct.ValueOf(in.ContactEmail)
	return out
}
func QuotaPreferenceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QuotaPreference) *krm.QuotaPreferenceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QuotaPreferenceObservedState{}
	// MISSING: Name
	// MISSING: Dimensions
	out.QuotaConfig = QuotaConfigObservedState_FromProto(mapCtx, in.GetQuotaConfig())
	// MISSING: Etag
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Service
	// MISSING: QuotaID
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	// MISSING: Justification
	// MISSING: ContactEmail
	return out
}
func QuotaPreferenceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QuotaPreferenceObservedState) *pb.QuotaPreference {
	if in == nil {
		return nil
	}
	out := &pb.QuotaPreference{}
	// MISSING: Name
	// MISSING: Dimensions
	out.QuotaConfig = QuotaConfigObservedState_ToProto(mapCtx, in.QuotaConfig)
	// MISSING: Etag
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Service
	// MISSING: QuotaID
	out.Reconciling = direct.ValueOf(in.Reconciling)
	// MISSING: Justification
	// MISSING: ContactEmail
	return out
}
