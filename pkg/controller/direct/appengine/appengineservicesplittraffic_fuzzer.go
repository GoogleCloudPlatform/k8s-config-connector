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
// proto.message: google.appengine.v1.Service
// api.group: appengine.cnrm.cloud.google.com

package appengine

import (
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(appEngineServiceSplitTrafficFuzzer())
}

func appEngineServiceSplitTrafficFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Service{},
		AppEngineServiceSplitTrafficSpec_v1alpha1_FromProto, AppEngineServiceSplitTrafficSpec_v1alpha1_ToProto,
		AppEngineServiceSplitTrafficObservedState_v1alpha1_FromProto, AppEngineServiceSplitTrafficObservedState_v1alpha1_ToProto,
	)

	// +-------------------------+---------------------------+----------------------------------------------+
	// | KRM Spec Field          | Proto Field               | Fuzzer/Mapping Status                        |
	// +-------------------------+---------------------------+----------------------------------------------+
	// | migrateTraffic          | N/A                       | KRM spec-only field                          |
	// | project                 | N/A                       | KRM spec-only GCP project reference          |
	// | resourceID              | id                        | SpecField (".id")                            |
	// | split                   | split                     | SpecField (".split")                         |
	// | split.allocations       | split.allocations         | SpecField (".split.allocations")             |
	// | split.shardBy           | split.shard_by            | SpecField (".split.shard_by")                |
	// +-------------------------+---------------------------+----------------------------------------------+

	// Spec fields
	f.SpecField(".id")
	f.SpecField(".split")

	// Identity fields
	f.Unimplemented_Identity(".name")

	// Unimplemented fields
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_NotYetTriaged(".network_settings")

	return f
}
