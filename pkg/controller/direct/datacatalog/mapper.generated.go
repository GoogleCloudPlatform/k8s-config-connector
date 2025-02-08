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

package datacatalog

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datacatalog/lineage/apiv1/lineagepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DatacatalogLineageEventObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LineageEvent) *krm.DatacatalogLineageEventObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatacatalogLineageEventObservedState{}
	// MISSING: Name
	// MISSING: Links
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func DatacatalogLineageEventObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatacatalogLineageEventObservedState) *pb.LineageEvent {
	if in == nil {
		return nil
	}
	out := &pb.LineageEvent{}
	// MISSING: Name
	// MISSING: Links
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func DatacatalogLineageEventSpec_FromProto(mapCtx *direct.MapContext, in *pb.LineageEvent) *krm.DatacatalogLineageEventSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatacatalogLineageEventSpec{}
	// MISSING: Name
	// MISSING: Links
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func DatacatalogLineageEventSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatacatalogLineageEventSpec) *pb.LineageEvent {
	if in == nil {
		return nil
	}
	out := &pb.LineageEvent{}
	// MISSING: Name
	// MISSING: Links
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func EntityReference_FromProto(mapCtx *direct.MapContext, in *pb.EntityReference) *krm.EntityReference {
	if in == nil {
		return nil
	}
	out := &krm.EntityReference{}
	out.FullyQualifiedName = direct.LazyPtr(in.GetFullyQualifiedName())
	return out
}
func EntityReference_ToProto(mapCtx *direct.MapContext, in *krm.EntityReference) *pb.EntityReference {
	if in == nil {
		return nil
	}
	out := &pb.EntityReference{}
	out.FullyQualifiedName = direct.ValueOf(in.FullyQualifiedName)
	return out
}
func EventLink_FromProto(mapCtx *direct.MapContext, in *pb.EventLink) *krm.EventLink {
	if in == nil {
		return nil
	}
	out := &krm.EventLink{}
	out.Source = EntityReference_FromProto(mapCtx, in.GetSource())
	out.Target = EntityReference_FromProto(mapCtx, in.GetTarget())
	return out
}
func EventLink_ToProto(mapCtx *direct.MapContext, in *krm.EventLink) *pb.EventLink {
	if in == nil {
		return nil
	}
	out := &pb.EventLink{}
	out.Source = EntityReference_ToProto(mapCtx, in.Source)
	out.Target = EntityReference_ToProto(mapCtx, in.Target)
	return out
}
func LineageEvent_FromProto(mapCtx *direct.MapContext, in *pb.LineageEvent) *krm.LineageEvent {
	if in == nil {
		return nil
	}
	out := &krm.LineageEvent{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Links = direct.Slice_FromProto(mapCtx, in.Links, EventLink_FromProto)
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func LineageEvent_ToProto(mapCtx *direct.MapContext, in *krm.LineageEvent) *pb.LineageEvent {
	if in == nil {
		return nil
	}
	out := &pb.LineageEvent{}
	out.Name = direct.ValueOf(in.Name)
	out.Links = direct.Slice_ToProto(mapCtx, in.Links, EventLink_ToProto)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
