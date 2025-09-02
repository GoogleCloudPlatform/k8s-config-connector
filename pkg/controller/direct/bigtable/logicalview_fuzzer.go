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
// proto.message: google.bigtable.admin.v2.LogicalView
// api.group: bigtable.cnrm.cloud.google.com

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(bigtableLogicalViewFuzzer())
}

func bigtableLogicalViewFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.LogicalView{},
		BigtableLogicalViewSpec_FromProto, BigtableLogicalViewSpec_ToProto,
		BigtableLogicalViewObservedState_FromProto, BigtableLogicalViewObservedState_ToProto,
	)

	f.SpecFields.Insert(".query")
	f.SpecFields.Insert(".deletion_protection")

	f.UnimplementedFields.Insert(".name") // special field
	f.UnimplementedFields.Insert(".etag")

	f.Unimplemented_NotYetTriaged(".deletion_protection")

	return f
}
