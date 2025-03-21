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

package cloudquota

import (
	pb "cloud.google.com/go/cloudquotas/apiv1beta/cloudquotaspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudquota/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func APIQuotaPreferenceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QuotaPreference) *krm.APIQuotaPreferenceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.APIQuotaPreferenceObservedState{}
	// MISSING: Name
	out.QuotaConfig = QuotaConfigObservedState_FromProto(mapCtx, in.GetQuotaConfig())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	return out
}
func APIQuotaPreferenceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.APIQuotaPreferenceObservedState) *pb.QuotaPreference {
	if in == nil {
		return nil
	}
	out := &pb.QuotaPreference{}
	// MISSING: Name
	out.QuotaConfig = QuotaConfigObservedState_ToProto(mapCtx, in.QuotaConfig)
	out.Etag = direct.ValueOf(in.Etag)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	return out
}
func APIQuotaPreferenceSpec_FromProto(mapCtx *direct.MapContext, in *pb.QuotaPreference) *krm.APIQuotaPreferenceSpec {
	if in == nil {
		return nil
	}
	out := &krm.APIQuotaPreferenceSpec{}
	// MISSING: Name
	out.Dimensions = in.Dimensions
	out.QuotaConfig = QuotaConfig_FromProto(mapCtx, in.GetQuotaConfig())
	out.Service = direct.LazyPtr(in.GetService())
	out.QuotaID = direct.LazyPtr(in.GetQuotaId())
	out.Justification = direct.LazyPtr(in.GetJustification())
	out.ContactEmail = direct.LazyPtr(in.GetContactEmail())
	return out
}
func APIQuotaPreferenceSpec_ToProto(mapCtx *direct.MapContext, in *krm.APIQuotaPreferenceSpec) *pb.QuotaPreference {
	if in == nil {
		return nil
	}
	out := &pb.QuotaPreference{}
	// MISSING: Name
	out.Dimensions = in.Dimensions
	out.QuotaConfig = QuotaConfig_ToProto(mapCtx, in.QuotaConfig)
	out.Service = direct.ValueOf(in.Service)
	out.QuotaId = direct.ValueOf(in.QuotaID)
	out.Justification = direct.ValueOf(in.Justification)
	out.ContactEmail = direct.ValueOf(in.ContactEmail)
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
	out.GrantedValue = direct.LazyPtr(direct.Int64Value_FromProto(mapCtx, in.GetGrantedValue()))
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
	out.GrantedValue = direct.Int64Value_ToProto(mapCtx, *in.GrantedValue)
	out.TraceId = direct.ValueOf(in.TraceID)
	// MISSING: Annotations
	out.RequestOrigin = direct.Enum_ToProto[pb.QuotaConfig_Origin](mapCtx, in.RequestOrigin)
	return out
}
