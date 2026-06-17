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
// proto.message: google.cloud.backupdr.v1.BackupVault
// api.group: backupdr.cnrm.cloud.google.com

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(backupDRBackupVaultFuzzer())
}

func backupDRBackupVaultFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.BackupVault{},
		BackupDRBackupVaultSpec_v1beta1_FromProto, BackupDRBackupVaultSpec_v1beta1_ToProto,
		BackupDRBackupVaultObservedState_v1beta1_FromProto, BackupDRBackupVaultObservedState_v1beta1_ToProto,
	)

	f.Unimplemented_Identity(".name")

	f.SpecField(".description")
	f.SpecField(".backup_minimum_enforced_retention_duration")
	f.SpecField(".effective_time")
	f.SpecField(".annotations")
	f.SpecField(".access_restriction")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".deletable")
	f.StatusField(".state")
	f.StatusField(".backup_count")
	f.StatusField(".service_account")
	f.StatusField(".total_stored_bytes")
	f.StatusField(".uid")
	f.StatusField(".etag")

	f.Unimplemented_LabelsAnnotations(".labels")

	f.Unimplemented_NotYetTriaged(".backup_retention_inheritance")
	f.Unimplemented_NotYetTriaged(".encryption_config")

	return f
}
