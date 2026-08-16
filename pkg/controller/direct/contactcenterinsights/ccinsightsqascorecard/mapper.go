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
	"time"

	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"google.golang.org/protobuf/types/known/timestamppb"

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
	if in.CreateTime != nil {
		out.CreateTime = direct.LazyPtr(in.CreateTime.AsTime().Format(time.RFC3339Nano))
	}
	if in.UpdateTime != nil {
		out.UpdateTime = direct.LazyPtr(in.UpdateTime.AsTime().Format(time.RFC3339Nano))
	}
	return out
}

func CCInsightsQAScorecardObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CCInsightsQAScorecardObservedState) *pb.QaScorecard {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecard{}
	if in.CreateTime != nil {
		t, err := time.Parse(time.RFC3339Nano, direct.ValueOf(in.CreateTime))
		if err != nil {
			mapCtx.Errorf("parsing createTime: %v", err)
		} else {
			out.CreateTime = timestamppb.New(t)
		}
	}
	if in.UpdateTime != nil {
		t, err := time.Parse(time.RFC3339Nano, direct.ValueOf(in.UpdateTime))
		if err != nil {
			mapCtx.Errorf("parsing updateTime: %v", err)
		} else {
			out.UpdateTime = timestamppb.New(t)
		}
	}
	return out
}
