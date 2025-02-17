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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/cloudquotas/apiv1/cloudquotaspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/api/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
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
