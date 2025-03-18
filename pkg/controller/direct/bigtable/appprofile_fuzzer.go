// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.bigtable.admin.v2.AppProfile
// api.group: bigtable.cnrm.cloud.google.com

package bigtable

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(bigtableAppProfileFuzzer())
}

func bigtableAppProfileFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.AppProfile{},
		BigtableAppProfileSpec_FromProto, BigtableAppProfileSpec_ToProto,
		BigtableAppProfileObservedState_FromProto, BigtableAppProfileObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".multi_cluster_routing_use_any")
	f.SpecFields.Insert(".single_cluster_routing")
	f.SpecFields.Insert(".priority")
	f.SpecFields.Insert(".standard_isolation")
	f.SpecFields.Insert(".data_boost_isolation_read_only")

	f.StatusFields.Insert(".etag")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
