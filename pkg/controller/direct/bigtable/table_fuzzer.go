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
// proto.message: google.bigtable.admin.v2.Table
// api.group: bigtable.cnrm.cloud.google.com

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(bigtableTableFuzzer())
}

func bigtableTableFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Table{},
		BigtableTableSpec_v1beta1_FromProto, BigtableTableSpec_v1beta1_ToProto,
		BigtableTableObservedState_v1beta1_FromProto, BigtableTableObservedState_v1beta1_ToProto,
	)

	f.SpecFields.Insert("column_families")
	f.SpecFields.Insert("granularity")
	f.SpecFields.Insert("change_stream_config")
	f.SpecFields.Insert("deletion_protection")
	f.SpecFields.Insert("automated_backup_policy")
	f.SpecFields.Insert("row_key_schema")

	f.StatusFields.Insert("cluster_states")
	f.StatusFields.Insert("restore_info")

	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".change_stream_config")
	f.UnimplementedFields.Insert(".granularity")
	f.UnimplementedFields.Insert(".automated_backup_policy")
	f.UnimplementedFields.Insert(".column_families.gc_rule")
	f.UnimplementedFields.Insert(".column_families.value_type")

	return f
}
