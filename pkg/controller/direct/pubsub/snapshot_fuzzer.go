// Copyright 2024 Google LLC
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

// +tool:fuzz-gen
// proto.message: google.pubsub.v1.Snapshot
// api.group: pubsub.cnrm.cloud.google.com

package pubsub

import (
	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(pubSubSnapshotFuzzer())
}

func pubSubSnapshotFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Snapshot{},
		PubSubSnapshotSpec_FromProto, PubSubSnapshotSpec_ToProto,
		PubSubSnapshotObservedState_FromProto, PubSubSnapshotObservedState_ToProto,
	)

	f.SpecFields.Insert(".topic")
	f.SpecFields.Insert(".expire_time")
	f.SpecFields.Insert(".labels")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
