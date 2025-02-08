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
func DiscoveryengineSampleQuerySetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SampleQuerySet) *krm.DiscoveryengineSampleQuerySetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineSampleQuerySetObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: Description
	return out
}
func DiscoveryengineSampleQuerySetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineSampleQuerySetObservedState) *pb.SampleQuerySet {
	if in == nil {
		return nil
	}
	out := &pb.SampleQuerySet{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: Description
	return out
}
func DiscoveryengineSampleQuerySetSpec_FromProto(mapCtx *direct.MapContext, in *pb.SampleQuerySet) *krm.DiscoveryengineSampleQuerySetSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineSampleQuerySetSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: Description
	return out
}
func DiscoveryengineSampleQuerySetSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineSampleQuerySetSpec) *pb.SampleQuerySet {
	if in == nil {
		return nil
	}
	out := &pb.SampleQuerySet{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: Description
	return out
}
func SampleQuerySet_FromProto(mapCtx *direct.MapContext, in *pb.SampleQuerySet) *krm.SampleQuerySet {
	if in == nil {
		return nil
	}
	out := &krm.SampleQuerySet{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func SampleQuerySet_ToProto(mapCtx *direct.MapContext, in *krm.SampleQuerySet) *pb.SampleQuerySet {
	if in == nil {
		return nil
	}
	out := &pb.SampleQuerySet{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	out.Description = direct.ValueOf(in.Description)
	return out
}
func SampleQuerySetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SampleQuerySet) *krm.SampleQuerySetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SampleQuerySetObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Description
	return out
}
func SampleQuerySetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SampleQuerySetObservedState) *pb.SampleQuerySet {
	if in == nil {
		return nil
	}
	out := &pb.SampleQuerySet{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Description
	return out
}
