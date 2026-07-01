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
// proto.message: google.cloud.sql.v1beta4.BackupRun
// api.group: sqladmin.cnrm.cloud.google.com

package sqladmin

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/sql/v1beta4"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(sqlAdminBackupFuzzer())
}

func sqlAdminBackupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.BackupRun{},
		SQLAdminBackupSpec_FromProto, SQLAdminBackupSpec_ToProto,
		SQLAdminBackupObservedState_FromProto, SQLAdminBackupObservedState_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".instance")
	f.SpecField(".location")
	f.SpecField(".disk_encryption_configuration")
	f.SpecField(".backup_kind")

	f.StatusField(".status")
	f.StatusField(".enqueued_time")
	f.StatusField(".id")
	f.StatusField(".start_time")
	f.StatusField(".end_time")
	f.StatusField(".error")
	f.StatusField(".type")
	f.StatusField(".window_start_time")
	f.StatusField(".disk_encryption_status")
	f.StatusField(".time_zone")

	f.Unimplemented_Identity(".kind")
	f.Unimplemented_NotYetTriaged(".self_link")
	f.Unimplemented_NotYetTriaged(".max_chargeable_bytes")

	return f
}
