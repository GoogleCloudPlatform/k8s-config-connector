// Copyright 2026 Google LLC
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

package ccinsightsqascorecard

import (
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CCInsightsQAScorecardSpec_ToProto(mapCtx *direct.MapContext, in *krm.CCInsightsQAScorecardSpec) *pb.QaScorecard {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecard{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	return out
}

func CCInsightsQAScorecardSpec_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecard) *krm.CCInsightsQAScorecardSpec {
	if in == nil {
		return nil
	}
	out := &krm.CCInsightsQAScorecardSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}

func CCInsightsQAScorecardObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecard) *krm.CCInsightsQAScorecardObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CCInsightsQAScorecardObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func CCInsightsQAScorecardObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CCInsightsQAScorecardObservedState) *pb.QaScorecard {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecard{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
