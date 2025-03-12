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
// proto.message: google.cloud.gkebackup.v1.BackupPlan
// api.group: gkebackup.cnrm.cloud.google.com

package gkebackup

import (
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(GKEBackupBackupPlanFuzzer())
}

func GKEBackupBackupPlanFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.BackupPlan{},
		GKEBackupBackupPlanSpec_FromProto, GKEBackupBackupPlanSpec_ToProto,
		GKEBackupBackupPlanObservedState_FromProto, GKEBackupBackupPlanObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".cluster")
	f.SpecFields.Insert(".retention_policy")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".backup_schedule")
	f.SpecFields.Insert(".deactivated")
	f.SpecFields.Insert(".backup_config")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".backup_schedule")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".protected_pod_count")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".state_reason")
	f.StatusFields.Insert(".rpo_risk_level")
	f.StatusFields.Insert(".rpo_risk_reason")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
