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
// krm.version: v1beta1
// proto.service: google.pubsub.v1

package pubsub

import (
	pb "cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func PubSubSnapshotObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.PubSubSnapshotObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PubSubSnapshotObservedState{}
	// MISSING: Name
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	return out
}
func PubSubSnapshotObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PubSubSnapshotObservedState) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	// MISSING: Name
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	return out
}
