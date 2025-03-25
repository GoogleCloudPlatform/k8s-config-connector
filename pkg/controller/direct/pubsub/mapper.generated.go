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

// +generated:mapper
// krm.group: pubsub.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.pubsub.v1

package pubsub

import (
	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func PubSubSnapshotObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.PubSubSnapshotObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PubSubSnapshotObservedState{}
	// MISSING: Name
	return out
}
func PubSubSnapshotObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PubSubSnapshotObservedState) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	// MISSING: Name
	return out
}
func PubSubSnapshotSpec_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.PubSubSnapshotSpec {
	if in == nil {
		return nil
	}
	out := &krm.PubSubSnapshotSpec{}
	// MISSING: Name
	out.Topic = direct.LazyPtr(in.GetTopic())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Labels = in.Labels
	return out
}
func PubSubSnapshotSpec_ToProto(mapCtx *direct.MapContext, in *krm.PubSubSnapshotSpec) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	// MISSING: Name
	out.Topic = direct.ValueOf(in.Topic)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.Labels = in.Labels
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
