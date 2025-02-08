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

package discoveryengine

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func SiteVerificationInfo_FromProto(mapCtx *direct.MapContext, in *pb.SiteVerificationInfo) *krm.SiteVerificationInfo {
	if in == nil {
		return nil
	}
	out := &krm.SiteVerificationInfo{}
	out.SiteVerificationState = direct.Enum_FromProto(mapCtx, in.GetSiteVerificationState())
	out.VerifyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetVerifyTime())
	return out
}
func SiteVerificationInfo_ToProto(mapCtx *direct.MapContext, in *krm.SiteVerificationInfo) *pb.SiteVerificationInfo {
	if in == nil {
		return nil
	}
	out := &pb.SiteVerificationInfo{}
	out.SiteVerificationState = direct.Enum_ToProto[pb.SiteVerificationInfo_SiteVerificationState](mapCtx, in.SiteVerificationState)
	out.VerifyTime = direct.StringTimestamp_ToProto(mapCtx, in.VerifyTime)
	return out
}
func TargetSite_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite) *krm.TargetSite {
	if in == nil {
		return nil
	}
	out := &krm.TargetSite{}
	// MISSING: Name
	out.ProvidedURIPattern = direct.LazyPtr(in.GetProvidedUriPattern())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	// MISSING: GeneratedURIPattern
	// MISSING: RootDomainURI
	// MISSING: SiteVerificationInfo
	// MISSING: IndexingStatus
	// MISSING: UpdateTime
	// MISSING: FailureReason
	return out
}
func TargetSite_ToProto(mapCtx *direct.MapContext, in *krm.TargetSite) *pb.TargetSite {
	if in == nil {
		return nil
	}
	out := &pb.TargetSite{}
	// MISSING: Name
	out.ProvidedUriPattern = direct.ValueOf(in.ProvidedURIPattern)
	out.Type = direct.Enum_ToProto[pb.TargetSite_Type](mapCtx, in.Type)
	out.ExactMatch = direct.ValueOf(in.ExactMatch)
	// MISSING: GeneratedURIPattern
	// MISSING: RootDomainURI
	// MISSING: SiteVerificationInfo
	// MISSING: IndexingStatus
	// MISSING: UpdateTime
	// MISSING: FailureReason
	return out
}
func TargetSiteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite) *krm.TargetSiteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TargetSiteObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ProvidedURIPattern
	// MISSING: Type
	// MISSING: ExactMatch
	out.GeneratedURIPattern = direct.LazyPtr(in.GetGeneratedUriPattern())
	out.RootDomainURI = direct.LazyPtr(in.GetRootDomainUri())
	out.SiteVerificationInfo = SiteVerificationInfo_FromProto(mapCtx, in.GetSiteVerificationInfo())
	out.IndexingStatus = direct.Enum_FromProto(mapCtx, in.GetIndexingStatus())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.FailureReason = TargetSite_FailureReason_FromProto(mapCtx, in.GetFailureReason())
	return out
}
func TargetSiteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TargetSiteObservedState) *pb.TargetSite {
	if in == nil {
		return nil
	}
	out := &pb.TargetSite{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ProvidedURIPattern
	// MISSING: Type
	// MISSING: ExactMatch
	out.GeneratedUriPattern = direct.ValueOf(in.GeneratedURIPattern)
	out.RootDomainUri = direct.ValueOf(in.RootDomainURI)
	out.SiteVerificationInfo = SiteVerificationInfo_ToProto(mapCtx, in.SiteVerificationInfo)
	out.IndexingStatus = direct.Enum_ToProto[pb.TargetSite_IndexingStatus](mapCtx, in.IndexingStatus)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.FailureReason = TargetSite_FailureReason_ToProto(mapCtx, in.FailureReason)
	return out
}
func TargetSite_FailureReason_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite_FailureReason) *krm.TargetSite_FailureReason {
	if in == nil {
		return nil
	}
	out := &krm.TargetSite_FailureReason{}
	out.QuotaFailure = TargetSite_FailureReason_QuotaFailure_FromProto(mapCtx, in.GetQuotaFailure())
	return out
}
func TargetSite_FailureReason_ToProto(mapCtx *direct.MapContext, in *krm.TargetSite_FailureReason) *pb.TargetSite_FailureReason {
	if in == nil {
		return nil
	}
	out := &pb.TargetSite_FailureReason{}
	if oneof := TargetSite_FailureReason_QuotaFailure_ToProto(mapCtx, in.QuotaFailure); oneof != nil {
		out.Failure = &pb.TargetSite_FailureReason_QuotaFailure_{QuotaFailure: oneof}
	}
	return out
}
func TargetSite_FailureReason_QuotaFailure_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite_FailureReason_QuotaFailure) *krm.TargetSite_FailureReason_QuotaFailure {
	if in == nil {
		return nil
	}
	out := &krm.TargetSite_FailureReason_QuotaFailure{}
	out.TotalRequiredQuota = direct.LazyPtr(in.GetTotalRequiredQuota())
	return out
}
func TargetSite_FailureReason_QuotaFailure_ToProto(mapCtx *direct.MapContext, in *krm.TargetSite_FailureReason_QuotaFailure) *pb.TargetSite_FailureReason_QuotaFailure {
	if in == nil {
		return nil
	}
	out := &pb.TargetSite_FailureReason_QuotaFailure{}
	out.TotalRequiredQuota = direct.ValueOf(in.TotalRequiredQuota)
	return out
}
