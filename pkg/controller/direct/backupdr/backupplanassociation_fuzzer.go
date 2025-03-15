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
		BackupDRBackupPlanAssociationSpec_FromProto, BackupDRBackupPlanAssociationSpec_ToProto,
		BackupDRBackupPlanAssociationObservedState_FromProto, BackupDRBackupPlanAssociationObservedState_ToProto,
	)

	f.SpecFields.Insert(".resource_type")
	f.SpecFields.Insert(".resource")
	f.SpecFields.Insert(".backup_plan")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".rules_config_info")
	f.StatusFields.Insert(".data_source")

	f.UnimplementedFields.Insert(".name")

	return f
}
