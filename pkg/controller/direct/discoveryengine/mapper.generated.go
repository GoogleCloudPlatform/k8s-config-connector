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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
)
func DiscoveryengineSampleQueryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SampleQuery) *krm.DiscoveryengineSampleQueryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineSampleQueryObservedState{}
	// MISSING: QueryEntry
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func DiscoveryengineSampleQueryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineSampleQueryObservedState) *pb.SampleQuery {
	if in == nil {
		return nil
	}
	out := &pb.SampleQuery{}
	// MISSING: QueryEntry
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func DiscoveryengineSampleQuerySpec_FromProto(mapCtx *direct.MapContext, in *pb.SampleQuery) *krm.DiscoveryengineSampleQuerySpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineSampleQuerySpec{}
	// MISSING: QueryEntry
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func DiscoveryengineSampleQuerySpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineSampleQuerySpec) *pb.SampleQuery {
	if in == nil {
		return nil
	}
	out := &pb.SampleQuery{}
	// MISSING: QueryEntry
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func SampleQuery_FromProto(mapCtx *direct.MapContext, in *pb.SampleQuery) *krm.SampleQuery {
	if in == nil {
		return nil
	}
	out := &krm.SampleQuery{}
	out.QueryEntry = SampleQuery_QueryEntry_FromProto(mapCtx, in.GetQueryEntry())
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	return out
}
func SampleQuery_ToProto(mapCtx *direct.MapContext, in *krm.SampleQuery) *pb.SampleQuery {
	if in == nil {
		return nil
	}
	out := &pb.SampleQuery{}
	if oneof := SampleQuery_QueryEntry_ToProto(mapCtx, in.QueryEntry); oneof != nil {
		out.Content = &pb.SampleQuery_QueryEntry_{QueryEntry: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	return out
}
func SampleQueryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SampleQuery) *krm.SampleQueryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SampleQueryObservedState{}
	// MISSING: QueryEntry
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func SampleQueryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SampleQueryObservedState) *pb.SampleQuery {
	if in == nil {
		return nil
	}
	out := &pb.SampleQuery{}
	// MISSING: QueryEntry
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
func SampleQuery_QueryEntry_FromProto(mapCtx *direct.MapContext, in *pb.SampleQuery_QueryEntry) *krm.SampleQuery_QueryEntry {
	if in == nil {
		return nil
	}
	out := &krm.SampleQuery_QueryEntry{}
	out.Query = direct.LazyPtr(in.GetQuery())
	out.Targets = direct.Slice_FromProto(mapCtx, in.Targets, SampleQuery_QueryEntry_Target_FromProto)
	return out
}
func SampleQuery_QueryEntry_ToProto(mapCtx *direct.MapContext, in *krm.SampleQuery_QueryEntry) *pb.SampleQuery_QueryEntry {
	if in == nil {
		return nil
	}
	out := &pb.SampleQuery_QueryEntry{}
	out.Query = direct.ValueOf(in.Query)
	out.Targets = direct.Slice_ToProto(mapCtx, in.Targets, SampleQuery_QueryEntry_Target_ToProto)
	return out
}
func SampleQuery_QueryEntry_Target_FromProto(mapCtx *direct.MapContext, in *pb.SampleQuery_QueryEntry_Target) *krm.SampleQuery_QueryEntry_Target {
	if in == nil {
		return nil
	}
	out := &krm.SampleQuery_QueryEntry_Target{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.PageNumbers = in.PageNumbers
	out.Score = in.Score
	return out
}
func SampleQuery_QueryEntry_Target_ToProto(mapCtx *direct.MapContext, in *krm.SampleQuery_QueryEntry_Target) *pb.SampleQuery_QueryEntry_Target {
	if in == nil {
		return nil
	}
	out := &pb.SampleQuery_QueryEntry_Target{}
	out.Uri = direct.ValueOf(in.URI)
	out.PageNumbers = in.PageNumbers
	out.Score = in.Score
	return out
}
