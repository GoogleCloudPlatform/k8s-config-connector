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

// +tool:fuzz-gen
// proto.message: google.bigtable.admin.v2.Instance
// api.group: bigtable.cnrm.cloud.google.com

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(bigtableInstanceFuzzer())
}

func bigtableInstanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.Instance{},
		BigtableInstanceSpec_v1beta1_FromProto, BigtableInstanceSpec_v1beta1_ToProto,
	)

	f.SpecFields.Insert("display_name")
	f.SpecFields.Insert("instance_type")
	f.SpecFields.Insert("labels")
	f.SpecFields.Insert("cluster")

	f.StatusFields.Insert("state")
	f.StatusFields.Insert("type")
	f.StatusFields.Insert("create_time")
	f.StatusFields.Insert("satisfies_pzs")

	f.UnimplementedFields.Insert(".name")

	return f
}
