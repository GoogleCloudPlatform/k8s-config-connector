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

package edgenetwork

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/edgenetwork/apiv1/edgenetworkpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/edgenetwork/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func EdgenetworkZoneObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Zone) *krm.EdgenetworkZoneObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EdgenetworkZoneObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: LayoutName
	return out
}
func EdgenetworkZoneObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EdgenetworkZoneObservedState) *pb.Zone {
	if in == nil {
		return nil
	}
	out := &pb.Zone{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: LayoutName
	return out
}
func EdgenetworkZoneSpec_FromProto(mapCtx *direct.MapContext, in *pb.Zone) *krm.EdgenetworkZoneSpec {
	if in == nil {
		return nil
	}
	out := &krm.EdgenetworkZoneSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: LayoutName
	return out
}
func EdgenetworkZoneSpec_ToProto(mapCtx *direct.MapContext, in *krm.EdgenetworkZoneSpec) *pb.Zone {
	if in == nil {
		return nil
	}
	out := &pb.Zone{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: LayoutName
	return out
}
func Zone_FromProto(mapCtx *direct.MapContext, in *pb.Zone) *krm.Zone {
	if in == nil {
		return nil
	}
	out := &krm.Zone{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.LayoutName = direct.LazyPtr(in.GetLayoutName())
	return out
}
func Zone_ToProto(mapCtx *direct.MapContext, in *krm.Zone) *pb.Zone {
	if in == nil {
		return nil
	}
	out := &pb.Zone{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.LayoutName = direct.ValueOf(in.LayoutName)
	return out
}
func ZoneObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Zone) *krm.ZoneObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ZoneObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: LayoutName
	return out
}
func ZoneObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ZoneObservedState) *pb.Zone {
	if in == nil {
		return nil
	}
	out := &pb.Zone{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: LayoutName
	return out
}
