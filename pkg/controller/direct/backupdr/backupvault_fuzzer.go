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
		BackupDRBackupVaultSpec_FromProto, BackupDRBackupVaultSpec_ToProto,
		BackupDRBackupVaultObservedState_FromProto, BackupDRBackupVaultObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".backup_minimum_enforced_retention_duration")
	f.SpecFields.Insert(".etag")
	f.SpecFields.Insert(".effective_time")
	f.SpecFields.Insert(".annotations")
	f.SpecFields.Insert(".access_restriction")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".deletable")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".backup_count")
	f.StatusFields.Insert(".service_account")
	f.StatusFields.Insert(".total_stored_bytes")
	f.StatusFields.Insert(".uid")

	f.UnimplementedFields.Insert(".name")

	return f
}
