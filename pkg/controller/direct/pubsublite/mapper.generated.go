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

package pubsublite

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/pubsublite/apiv1/pubsublitepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsublite/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func PubsubliteReservationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.PubsubliteReservationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PubsubliteReservationObservedState{}
	// MISSING: Name
	// MISSING: ThroughputCapacity
	return out
}
func PubsubliteReservationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PubsubliteReservationObservedState) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	// MISSING: Name
	// MISSING: ThroughputCapacity
	return out
}
func PubsubliteReservationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.PubsubliteReservationSpec {
	if in == nil {
		return nil
	}
	out := &krm.PubsubliteReservationSpec{}
	// MISSING: Name
	// MISSING: ThroughputCapacity
	return out
}
func PubsubliteReservationSpec_ToProto(mapCtx *direct.MapContext, in *krm.PubsubliteReservationSpec) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	// MISSING: Name
	// MISSING: ThroughputCapacity
	return out
}
func Reservation_FromProto(mapCtx *direct.MapContext, in *pb.Reservation) *krm.Reservation {
	if in == nil {
		return nil
	}
	out := &krm.Reservation{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ThroughputCapacity = direct.LazyPtr(in.GetThroughputCapacity())
	return out
}
func Reservation_ToProto(mapCtx *direct.MapContext, in *krm.Reservation) *pb.Reservation {
	if in == nil {
		return nil
	}
	out := &pb.Reservation{}
	out.Name = direct.ValueOf(in.Name)
	out.ThroughputCapacity = direct.ValueOf(in.ThroughputCapacity)
	return out
}
