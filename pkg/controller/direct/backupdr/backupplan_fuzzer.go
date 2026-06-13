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
// proto.message: google.cloud.backupdr.v1.BackupPlan
// api.group: backupdr.cnrm.cloud.google.com

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(backupDRBackupPlanFuzzer())
}

func backupDRBackupPlanFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.BackupPlan{},
		BackupDRBackupPlanSpec_v1beta1_FromProto, BackupDRBackupPlanSpec_v1beta1_ToProto,
		BackupDRBackupPlanObservedState_v1beta1_FromProto, BackupDRBackupPlanObservedState_v1beta1_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".backup_rules")
	f.SpecField(".resource_type")
	f.SpecField(".backup_vault")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".state")
	f.StatusField(".backup_vault_service_account")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".log_retention_days")
	f.Unimplemented_NotYetTriaged(".revision_id")
	f.Unimplemented_NotYetTriaged(".revision_name")
	f.Unimplemented_NotYetTriaged(".supported_resource_types")
	f.Unimplemented_Etag()
	f.Unimplemented_LabelsAnnotations(".labels")

	return f
}
