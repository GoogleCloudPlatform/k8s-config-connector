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
func DimensionsInfo_FromProto(mapCtx *direct.MapContext, in *pb.DimensionsInfo) *krm.DimensionsInfo {
	if in == nil {
		return nil
	}
	out := &krm.DimensionsInfo{}
	out.Dimensions = in.Dimensions
	out.Details = QuotaDetails_FromProto(mapCtx, in.GetDetails())
	out.ApplicableLocations = in.ApplicableLocations
	return out
}
func DimensionsInfo_ToProto(mapCtx *direct.MapContext, in *krm.DimensionsInfo) *pb.DimensionsInfo {
	if in == nil {
		return nil
	}
	out := &pb.DimensionsInfo{}
	out.Dimensions = in.Dimensions
	out.Details = QuotaDetails_ToProto(mapCtx, in.Details)
	out.ApplicableLocations = in.ApplicableLocations
	return out
}
func QuotaDetails_FromProto(mapCtx *direct.MapContext, in *pb.QuotaDetails) *krm.QuotaDetails {
	if in == nil {
		return nil
	}
	out := &krm.QuotaDetails{}
	out.Value = direct.LazyPtr(in.GetValue())
	out.RolloutInfo = RolloutInfo_FromProto(mapCtx, in.GetRolloutInfo())
	return out
}
func QuotaDetails_ToProto(mapCtx *direct.MapContext, in *krm.QuotaDetails) *pb.QuotaDetails {
	if in == nil {
		return nil
	}
	out := &pb.QuotaDetails{}
	out.Value = direct.ValueOf(in.Value)
	out.RolloutInfo = RolloutInfo_ToProto(mapCtx, in.RolloutInfo)
	return out
}
func QuotaIncreaseEligibility_FromProto(mapCtx *direct.MapContext, in *pb.QuotaIncreaseEligibility) *krm.QuotaIncreaseEligibility {
	if in == nil {
		return nil
	}
	out := &krm.QuotaIncreaseEligibility{}
	out.IsEligible = direct.LazyPtr(in.GetIsEligible())
	out.IneligibilityReason = direct.Enum_FromProto(mapCtx, in.GetIneligibilityReason())
	return out
}
func QuotaIncreaseEligibility_ToProto(mapCtx *direct.MapContext, in *krm.QuotaIncreaseEligibility) *pb.QuotaIncreaseEligibility {
	if in == nil {
		return nil
	}
	out := &pb.QuotaIncreaseEligibility{}
	out.IsEligible = direct.ValueOf(in.IsEligible)
	out.IneligibilityReason = direct.Enum_ToProto[pb.QuotaIncreaseEligibility_IneligibilityReason](mapCtx, in.IneligibilityReason)
	return out
}
func QuotaInfo_FromProto(mapCtx *direct.MapContext, in *pb.QuotaInfo) *krm.QuotaInfo {
	if in == nil {
		return nil
	}
	out := &krm.QuotaInfo{}
	out.Name = direct.LazyPtr(in.GetName())
	out.QuotaID = direct.LazyPtr(in.GetQuotaId())
	out.Metric = direct.LazyPtr(in.GetMetric())
	out.Service = direct.LazyPtr(in.GetService())
	out.IsPrecise = direct.LazyPtr(in.GetIsPrecise())
	out.RefreshInterval = direct.LazyPtr(in.GetRefreshInterval())
	out.ContainerType = direct.Enum_FromProto(mapCtx, in.GetContainerType())
	out.Dimensions = in.Dimensions
	out.MetricDisplayName = direct.LazyPtr(in.GetMetricDisplayName())
	out.QuotaDisplayName = direct.LazyPtr(in.GetQuotaDisplayName())
	out.MetricUnit = direct.LazyPtr(in.GetMetricUnit())
	out.QuotaIncreaseEligibility = QuotaIncreaseEligibility_FromProto(mapCtx, in.GetQuotaIncreaseEligibility())
	out.IsFixed = direct.LazyPtr(in.GetIsFixed())
	out.DimensionsInfos = direct.Slice_FromProto(mapCtx, in.DimensionsInfos, DimensionsInfo_FromProto)
	out.IsConcurrent = direct.LazyPtr(in.GetIsConcurrent())
	out.ServiceRequestQuotaURI = direct.LazyPtr(in.GetServiceRequestQuotaUri())
	return out
}
func QuotaInfo_ToProto(mapCtx *direct.MapContext, in *krm.QuotaInfo) *pb.QuotaInfo {
	if in == nil {
		return nil
	}
	out := &pb.QuotaInfo{}
	out.Name = direct.ValueOf(in.Name)
	out.QuotaId = direct.ValueOf(in.QuotaID)
	out.Metric = direct.ValueOf(in.Metric)
	out.Service = direct.ValueOf(in.Service)
	out.IsPrecise = direct.ValueOf(in.IsPrecise)
	out.RefreshInterval = direct.ValueOf(in.RefreshInterval)
	out.ContainerType = direct.Enum_ToProto[pb.QuotaInfo_ContainerType](mapCtx, in.ContainerType)
	out.Dimensions = in.Dimensions
	out.MetricDisplayName = direct.ValueOf(in.MetricDisplayName)
	out.QuotaDisplayName = direct.ValueOf(in.QuotaDisplayName)
	out.MetricUnit = direct.ValueOf(in.MetricUnit)
	out.QuotaIncreaseEligibility = QuotaIncreaseEligibility_ToProto(mapCtx, in.QuotaIncreaseEligibility)
	out.IsFixed = direct.ValueOf(in.IsFixed)
	out.DimensionsInfos = direct.Slice_ToProto(mapCtx, in.DimensionsInfos, DimensionsInfo_ToProto)
	out.IsConcurrent = direct.ValueOf(in.IsConcurrent)
	out.ServiceRequestQuotaUri = direct.ValueOf(in.ServiceRequestQuotaURI)
	return out
}
func RolloutInfo_FromProto(mapCtx *direct.MapContext, in *pb.RolloutInfo) *krm.RolloutInfo {
	if in == nil {
		return nil
	}
	out := &krm.RolloutInfo{}
	out.OngoingRollout = direct.LazyPtr(in.GetOngoingRollout())
	return out
}
func RolloutInfo_ToProto(mapCtx *direct.MapContext, in *krm.RolloutInfo) *pb.RolloutInfo {
	if in == nil {
		return nil
	}
	out := &pb.RolloutInfo{}
	out.OngoingRollout = direct.ValueOf(in.OngoingRollout)
	return out
}
