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
// proto.message: google.cloud.alloydb.v1beta.Backup

package alloydbbackup

import (
	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(backupFuzzer())
}

func backupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Backup{},
		AlloyDBBackupSpec_FromProto, AlloyDBBackupSpec_ToProto,
		AlloyDBBackupStatus_FromProto, AlloyDBBackupStatus_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".cluster_name")
	f.SpecField(".encryption_config")

	f.StatusField(".name")
	f.StatusField(".uid")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".state")
	f.StatusField(".reconciling")
	f.StatusField(".etag")
	f.StatusField(".encryption_info")

	f.Unimplemented_Internal(".satisfies_pzs")

	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_LabelsAnnotations(".annotations")

	f.Unimplemented_NotYetTriaged(".display_name")
	f.Unimplemented_NotYetTriaged(".delete_time")
	f.Unimplemented_NotYetTriaged(".create_completion_time")
	f.Unimplemented_NotYetTriaged(".cluster_uid")
	f.Unimplemented_NotYetTriaged(".type")
	f.Unimplemented_NotYetTriaged(".size_bytes")
	f.Unimplemented_NotYetTriaged(".expiry_time")
	f.Unimplemented_NotYetTriaged(".expiry_quantity")
	f.Unimplemented_NotYetTriaged(".database_version")
	f.Unimplemented_NotYetTriaged(".tags")

	return f
}
