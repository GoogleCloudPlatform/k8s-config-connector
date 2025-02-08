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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ContactcenterinsightsIssueObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Issue) *krm.ContactcenterinsightsIssueObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsIssueObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SampleUtterances
	// MISSING: DisplayDescription
	return out
}
func ContactcenterinsightsIssueObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsIssueObservedState) *pb.Issue {
	if in == nil {
		return nil
	}
	out := &pb.Issue{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SampleUtterances
	// MISSING: DisplayDescription
	return out
}
func ContactcenterinsightsIssueSpec_FromProto(mapCtx *direct.MapContext, in *pb.Issue) *krm.ContactcenterinsightsIssueSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsIssueSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SampleUtterances
	// MISSING: DisplayDescription
	return out
}
func ContactcenterinsightsIssueSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsIssueSpec) *pb.Issue {
	if in == nil {
		return nil
	}
	out := &pb.Issue{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SampleUtterances
	// MISSING: DisplayDescription
	return out
}
func Issue_FromProto(mapCtx *direct.MapContext, in *pb.Issue) *krm.Issue {
	if in == nil {
		return nil
	}
	out := &krm.Issue{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SampleUtterances
	out.DisplayDescription = direct.LazyPtr(in.GetDisplayDescription())
	return out
}
func Issue_ToProto(mapCtx *direct.MapContext, in *krm.Issue) *pb.Issue {
	if in == nil {
		return nil
	}
	out := &pb.Issue{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SampleUtterances
	out.DisplayDescription = direct.ValueOf(in.DisplayDescription)
	return out
}
func IssueObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Issue) *krm.IssueObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IssueObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.SampleUtterances = in.SampleUtterances
	// MISSING: DisplayDescription
	return out
}
func IssueObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IssueObservedState) *pb.Issue {
	if in == nil {
		return nil
	}
	out := &pb.Issue{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.SampleUtterances = in.SampleUtterances
	// MISSING: DisplayDescription
	return out
}
