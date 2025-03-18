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
// proto.message: google.bigtable.admin.v2.Backup
// api.group: bigtable.cnrm.cloud.google.com

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/bigtablepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(bigtableBackupFuzzer())
}

func bigtableBackupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Backup{},
		BigtableBackupSpec_FromProto, BigtableBackupSpec_ToProto,
		BigtableBackupObservedState_FromProto, BigtableBackupObservedState_ToProto,
	)

	f.SpecFields.Insert(".source_table")
	f.SpecFields.Insert(".expire_time")
	f.SpecFields.Insert(".backup_type")
	f.SpecFields.Insert(".hot_to_standard_time")

	f.StatusFields.Insert(".source_backup")
	f.StatusFields.Insert(".start_time")
	f.StatusFields.Insert(".end_time")
	f.StatusFields.Insert(".size_bytes")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".encryption_info")

	f.UnimplementedFields.Insert(".name")
	return f
}
