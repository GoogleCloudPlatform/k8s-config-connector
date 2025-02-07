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
	pb "cloud.google.com/go/cloudquotas/apiv1beta/cloudquotaspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/api/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ApiQuotaAdjusterSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QuotaAdjusterSettings) *krm.ApiQuotaAdjusterSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApiQuotaAdjusterSettingsObservedState{}
	// MISSING: Name
	// MISSING: Enablement
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func ApiQuotaAdjusterSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApiQuotaAdjusterSettingsObservedState) *pb.QuotaAdjusterSettings {
	if in == nil {
		return nil
	}
	out := &pb.QuotaAdjusterSettings{}
	// MISSING: Name
	// MISSING: Enablement
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func ApiQuotaAdjusterSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.QuotaAdjusterSettings) *krm.ApiQuotaAdjusterSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApiQuotaAdjusterSettingsSpec{}
	// MISSING: Name
	// MISSING: Enablement
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func ApiQuotaAdjusterSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApiQuotaAdjusterSettingsSpec) *pb.QuotaAdjusterSettings {
	if in == nil {
		return nil
	}
	out := &pb.QuotaAdjusterSettings{}
	// MISSING: Name
	// MISSING: Enablement
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func QuotaAdjusterSettings_FromProto(mapCtx *direct.MapContext, in *pb.QuotaAdjusterSettings) *krm.QuotaAdjusterSettings {
	if in == nil {
		return nil
	}
	out := &krm.QuotaAdjusterSettings{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Enablement = direct.Enum_FromProto(mapCtx, in.GetEnablement())
	// MISSING: UpdateTime
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func QuotaAdjusterSettings_ToProto(mapCtx *direct.MapContext, in *krm.QuotaAdjusterSettings) *pb.QuotaAdjusterSettings {
	if in == nil {
		return nil
	}
	out := &pb.QuotaAdjusterSettings{}
	out.Name = direct.ValueOf(in.Name)
	out.Enablement = direct.Enum_ToProto[pb.QuotaAdjusterSettings_Enablement](mapCtx, in.Enablement)
	// MISSING: UpdateTime
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func QuotaAdjusterSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QuotaAdjusterSettings) *krm.QuotaAdjusterSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QuotaAdjusterSettingsObservedState{}
	// MISSING: Name
	// MISSING: Enablement
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Etag
	return out
}
func QuotaAdjusterSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QuotaAdjusterSettingsObservedState) *pb.QuotaAdjusterSettings {
	if in == nil {
		return nil
	}
	out := &pb.QuotaAdjusterSettings{}
	// MISSING: Name
	// MISSING: Enablement
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Etag
	return out
}
