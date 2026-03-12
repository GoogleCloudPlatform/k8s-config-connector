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
// proto.message: google.firestore.admin.v1.BackupSchedule
// api.group: firestore.cnrm.cloud.google.com
package firestore

import (
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(backupScheduleFuzzer())
}

func backupScheduleFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.BackupSchedule{},
		FirestoreBackupScheduleSpec_v1alpha1_FromProto, FirestoreBackupScheduleSpec_v1alpha1_ToProto,
		FirestoreBackupScheduleObservedState_v1alpha1_FromProto, FirestoreBackupScheduleObservedState_v1alpha1_ToProto,
	)

	f.SpecField(".retention")
	f.SpecField(".daily_recurrence")
	f.SpecField(".weekly_recurrence")

	f.StatusField(".name") // Server generated ID, so surface in observedState
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	return f
}
