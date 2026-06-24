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
// proto.message: google.cloud.backupdr.v1.BackupPlanAssociation
// api.group: backupdr.cnrm.cloud.google.com

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(backupDRBackupPlanAssociationFuzzer())
}

func backupDRBackupPlanAssociationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.BackupPlanAssociation{},
		BackupDRBackupPlanAssociationSpec_v1beta1_FromProto, BackupDRBackupPlanAssociationSpec_v1beta1_ToProto,
		BackupDRBackupPlanAssociationObservedState_v1beta1_FromProto, BackupDRBackupPlanAssociationObservedState_v1beta1_ToProto,
	)

	f.SpecField(".resource_type")
	f.SpecField(".backup_plan")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".state")
	f.StatusField(".rules_config_info")
	f.StatusField(".data_source")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".resource") // special field, the value has to be a URL representing ComputeInstance.

	f.Unimplemented_NotYetTriaged(".backup_plan_revision_name")
	f.Unimplemented_NotYetTriaged(".backup_plan_revision_id")
	f.Unimplemented_NotYetTriaged(".backup_plan_revision_name")
	f.Unimplemented_NotYetTriaged(".cloud_sql_instance_backup_plan_association_properties")
	f.Unimplemented_NotYetTriaged(".resource_properties")

	f.Unimplemented_NotYetTriaged(".rules_config_info[].last_backup_error.details")
	f.Unimplemented_NotYetTriaged(".rules_config_info[].last_backup_error.details[].value")
	f.Unimplemented_NotYetTriaged(".rules_config_info[].last_backup_error.details[].type_url")

	return f
}
