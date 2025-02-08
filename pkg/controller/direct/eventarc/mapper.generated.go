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

package eventarc

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Enrollment_FromProto(mapCtx *direct.MapContext, in *pb.Enrollment) *krm.Enrollment {
	if in == nil {
		return nil
	}
	out := &krm.Enrollment{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.CelMatch = direct.LazyPtr(in.GetCelMatch())
	out.MessageBus = direct.LazyPtr(in.GetMessageBus())
	out.Destination = direct.LazyPtr(in.GetDestination())
	return out
}
func Enrollment_ToProto(mapCtx *direct.MapContext, in *krm.Enrollment) *pb.Enrollment {
	if in == nil {
		return nil
	}
	out := &pb.Enrollment{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.CelMatch = direct.ValueOf(in.CelMatch)
	out.MessageBus = direct.ValueOf(in.MessageBus)
	out.Destination = direct.ValueOf(in.Destination)
	return out
}
func EnrollmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Enrollment) *krm.EnrollmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EnrollmentObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: CelMatch
	// MISSING: MessageBus
	// MISSING: Destination
	return out
}
func EnrollmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EnrollmentObservedState) *pb.Enrollment {
	if in == nil {
		return nil
	}
	out := &pb.Enrollment{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	out.Etag = direct.ValueOf(in.Etag)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: CelMatch
	// MISSING: MessageBus
	// MISSING: Destination
	return out
}
func EventarcEnrollmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Enrollment) *krm.EventarcEnrollmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventarcEnrollmentObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: CelMatch
	// MISSING: MessageBus
	// MISSING: Destination
	return out
}
func EventarcEnrollmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventarcEnrollmentObservedState) *pb.Enrollment {
	if in == nil {
		return nil
	}
	out := &pb.Enrollment{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: CelMatch
	// MISSING: MessageBus
	// MISSING: Destination
	return out
}
func EventarcEnrollmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Enrollment) *krm.EventarcEnrollmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcEnrollmentSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: CelMatch
	// MISSING: MessageBus
	// MISSING: Destination
	return out
}
func EventarcEnrollmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcEnrollmentSpec) *pb.Enrollment {
	if in == nil {
		return nil
	}
	out := &pb.Enrollment{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: CelMatch
	// MISSING: MessageBus
	// MISSING: Destination
	return out
}
