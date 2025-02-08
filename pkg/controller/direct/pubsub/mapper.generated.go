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

package pubsub

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1alpha1"
)
func PubsubSnapshotObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.PubsubSnapshotObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PubsubSnapshotObservedState{}
	// MISSING: Name
	// MISSING: Topic
	// MISSING: ExpireTime
	// MISSING: Labels
	return out
}
func PubsubSnapshotObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PubsubSnapshotObservedState) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	// MISSING: Name
	// MISSING: Topic
	// MISSING: ExpireTime
	// MISSING: Labels
	return out
}
func PubsubSnapshotSpec_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.PubsubSnapshotSpec {
	if in == nil {
		return nil
	}
	out := &krm.PubsubSnapshotSpec{}
	// MISSING: Name
	// MISSING: Topic
	// MISSING: ExpireTime
	// MISSING: Labels
	return out
}
func PubsubSnapshotSpec_ToProto(mapCtx *direct.MapContext, in *krm.PubsubSnapshotSpec) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	// MISSING: Name
	// MISSING: Topic
	// MISSING: ExpireTime
	// MISSING: Labels
	return out
}
func Snapshot_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.Snapshot {
	if in == nil {
		return nil
	}
	out := &krm.Snapshot{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Topic = direct.LazyPtr(in.GetTopic())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Labels = in.Labels
	return out
}
func Snapshot_ToProto(mapCtx *direct.MapContext, in *krm.Snapshot) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	out.Name = direct.ValueOf(in.Name)
	out.Topic = direct.ValueOf(in.Topic)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.Labels = in.Labels
	return out
}
