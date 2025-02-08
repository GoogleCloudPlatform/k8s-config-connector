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

package spanner

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/spanner/apiv1/spannerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Session_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.Session {
	if in == nil {
		return nil
	}
	out := &krm.Session{}
	// MISSING: Name
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: ApproximateLastUseTime
	out.CreatorRole = direct.LazyPtr(in.GetCreatorRole())
	out.Multiplexed = direct.LazyPtr(in.GetMultiplexed())
	return out
}
func Session_ToProto(mapCtx *direct.MapContext, in *krm.Session) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	// MISSING: Name
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: ApproximateLastUseTime
	out.CreatorRole = direct.ValueOf(in.CreatorRole)
	out.Multiplexed = direct.ValueOf(in.Multiplexed)
	return out
}
func SessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.SessionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SessionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.ApproximateLastUseTime = direct.StringTimestamp_FromProto(mapCtx, in.GetApproximateLastUseTime())
	// MISSING: CreatorRole
	// MISSING: Multiplexed
	return out
}
func SessionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SessionObservedState) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.ApproximateLastUseTime = direct.StringTimestamp_ToProto(mapCtx, in.ApproximateLastUseTime)
	// MISSING: CreatorRole
	// MISSING: Multiplexed
	return out
}
func SpannerSessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.SpannerSessionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpannerSessionObservedState{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: ApproximateLastUseTime
	// MISSING: CreatorRole
	// MISSING: Multiplexed
	return out
}
func SpannerSessionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpannerSessionObservedState) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: ApproximateLastUseTime
	// MISSING: CreatorRole
	// MISSING: Multiplexed
	return out
}
func SpannerSessionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.SpannerSessionSpec {
	if in == nil {
		return nil
	}
	out := &krm.SpannerSessionSpec{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: ApproximateLastUseTime
	// MISSING: CreatorRole
	// MISSING: Multiplexed
	return out
}
func SpannerSessionSpec_ToProto(mapCtx *direct.MapContext, in *krm.SpannerSessionSpec) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: ApproximateLastUseTime
	// MISSING: CreatorRole
	// MISSING: Multiplexed
	return out
}
