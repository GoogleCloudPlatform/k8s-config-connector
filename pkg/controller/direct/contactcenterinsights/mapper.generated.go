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

package contactcenterinsights

import (
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func ContactcenterinsightsQaScorecardObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecard) *krm.ContactcenterinsightsQaScorecardObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsQaScorecardObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ContactcenterinsightsQaScorecardObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsQaScorecardObservedState) *pb.QaScorecard {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecard{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ContactcenterinsightsQaScorecardSpec_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecard) *krm.ContactcenterinsightsQaScorecardSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsQaScorecardSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ContactcenterinsightsQaScorecardSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsQaScorecardSpec) *pb.QaScorecard {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecard{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func QaScorecard_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecard) *krm.QaScorecard {
	if in == nil {
		return nil
	}
	out := &krm.QaScorecard{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func QaScorecard_ToProto(mapCtx *direct.MapContext, in *krm.QaScorecard) *pb.QaScorecard {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecard{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func QaScorecardObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QaScorecard) *krm.QaScorecardObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QaScorecardObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func QaScorecardObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QaScorecardObservedState) *pb.QaScorecard {
	if in == nil {
		return nil
	}
	out := &pb.QaScorecard{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
