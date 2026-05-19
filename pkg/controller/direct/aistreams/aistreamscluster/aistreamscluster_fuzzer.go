// Copyright 2026 Google LLC
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
// proto.message: google.partner.aistreams.v1alpha1.Cluster
// api.group: aistreams.cnrm.cloud.google.com

package aistreamscluster

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	pb "google.golang.org/genproto/googleapis/partner/aistreams/v1alpha1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(aistreamsClusterFuzzer())
}

func aistreamsClusterFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Cluster{},
		AIStreamsClusterSpec_FromProto, AIStreamsClusterSpec_ToProto,
		AIStreamsClusterObservedState_FromProto, AIStreamsClusterObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".labels")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".certificate")
	f.StatusFields.Insert(".service_endpoint")

	return f
}
