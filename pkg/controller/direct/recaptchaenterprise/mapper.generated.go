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

package recaptchaenterprise

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/recaptchaenterprise/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ChallengeMetrics_FromProto(mapCtx *direct.MapContext, in *pb.ChallengeMetrics) *krm.ChallengeMetrics {
	if in == nil {
		return nil
	}
	out := &krm.ChallengeMetrics{}
	out.PageloadCount = direct.LazyPtr(in.GetPageloadCount())
	out.NocaptchaCount = direct.LazyPtr(in.GetNocaptchaCount())
	out.FailedCount = direct.LazyPtr(in.GetFailedCount())
	out.PassedCount = direct.LazyPtr(in.GetPassedCount())
	return out
}
func ChallengeMetrics_ToProto(mapCtx *direct.MapContext, in *krm.ChallengeMetrics) *pb.ChallengeMetrics {
	if in == nil {
		return nil
	}
	out := &pb.ChallengeMetrics{}
	out.PageloadCount = direct.ValueOf(in.PageloadCount)
	out.NocaptchaCount = direct.ValueOf(in.NocaptchaCount)
	out.FailedCount = direct.ValueOf(in.FailedCount)
	out.PassedCount = direct.ValueOf(in.PassedCount)
	return out
}
func Metrics_FromProto(mapCtx *direct.MapContext, in *pb.Metrics) *krm.Metrics {
	if in == nil {
		return nil
	}
	out := &krm.Metrics{}
	// MISSING: Name
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.ScoreMetrics = direct.Slice_FromProto(mapCtx, in.ScoreMetrics, ScoreMetrics_FromProto)
	out.ChallengeMetrics = direct.Slice_FromProto(mapCtx, in.ChallengeMetrics, ChallengeMetrics_FromProto)
	return out
}
func Metrics_ToProto(mapCtx *direct.MapContext, in *krm.Metrics) *pb.Metrics {
	if in == nil {
		return nil
	}
	out := &pb.Metrics{}
	// MISSING: Name
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.ScoreMetrics = direct.Slice_ToProto(mapCtx, in.ScoreMetrics, ScoreMetrics_ToProto)
	out.ChallengeMetrics = direct.Slice_ToProto(mapCtx, in.ChallengeMetrics, ChallengeMetrics_ToProto)
	return out
}
func MetricsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Metrics) *krm.MetricsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetricsObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: StartTime
	// MISSING: ScoreMetrics
	// MISSING: ChallengeMetrics
	return out
}
func MetricsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetricsObservedState) *pb.Metrics {
	if in == nil {
		return nil
	}
	out := &pb.Metrics{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: StartTime
	// MISSING: ScoreMetrics
	// MISSING: ChallengeMetrics
	return out
}
func RecaptchaenterpriseMetricsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Metrics) *krm.RecaptchaenterpriseMetricsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RecaptchaenterpriseMetricsObservedState{}
	// MISSING: Name
	// MISSING: StartTime
	// MISSING: ScoreMetrics
	// MISSING: ChallengeMetrics
	return out
}
func RecaptchaenterpriseMetricsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RecaptchaenterpriseMetricsObservedState) *pb.Metrics {
	if in == nil {
		return nil
	}
	out := &pb.Metrics{}
	// MISSING: Name
	// MISSING: StartTime
	// MISSING: ScoreMetrics
	// MISSING: ChallengeMetrics
	return out
}
func RecaptchaenterpriseMetricsSpec_FromProto(mapCtx *direct.MapContext, in *pb.Metrics) *krm.RecaptchaenterpriseMetricsSpec {
	if in == nil {
		return nil
	}
	out := &krm.RecaptchaenterpriseMetricsSpec{}
	// MISSING: Name
	// MISSING: StartTime
	// MISSING: ScoreMetrics
	// MISSING: ChallengeMetrics
	return out
}
func RecaptchaenterpriseMetricsSpec_ToProto(mapCtx *direct.MapContext, in *krm.RecaptchaenterpriseMetricsSpec) *pb.Metrics {
	if in == nil {
		return nil
	}
	out := &pb.Metrics{}
	// MISSING: Name
	// MISSING: StartTime
	// MISSING: ScoreMetrics
	// MISSING: ChallengeMetrics
	return out
}
func ScoreDistribution_FromProto(mapCtx *direct.MapContext, in *pb.ScoreDistribution) *krm.ScoreDistribution {
	if in == nil {
		return nil
	}
	out := &krm.ScoreDistribution{}
	// MISSING: ScoreBuckets
	return out
}
func ScoreDistribution_ToProto(mapCtx *direct.MapContext, in *krm.ScoreDistribution) *pb.ScoreDistribution {
	if in == nil {
		return nil
	}
	out := &pb.ScoreDistribution{}
	// MISSING: ScoreBuckets
	return out
}
func ScoreMetrics_FromProto(mapCtx *direct.MapContext, in *pb.ScoreMetrics) *krm.ScoreMetrics {
	if in == nil {
		return nil
	}
	out := &krm.ScoreMetrics{}
	out.OverallMetrics = ScoreDistribution_FromProto(mapCtx, in.GetOverallMetrics())
	// MISSING: ActionMetrics
	return out
}
func ScoreMetrics_ToProto(mapCtx *direct.MapContext, in *krm.ScoreMetrics) *pb.ScoreMetrics {
	if in == nil {
		return nil
	}
	out := &pb.ScoreMetrics{}
	out.OverallMetrics = ScoreDistribution_ToProto(mapCtx, in.OverallMetrics)
	// MISSING: ActionMetrics
	return out
}
