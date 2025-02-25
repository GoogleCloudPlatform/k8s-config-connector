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
// proto.message: google.spanner.admin.database.v1.BackupSchedule
// api.group: spannerbackupschedule.cnrm.cloud.google.com

package spannerbackupschedules

import (
	pb "cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(spannerBackupScheduleFuzzer())
}

func spannerBackupScheduleFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.BackupSchedule{},
		SpannerBackupScheduleSpec_FromProto, SpannerBackupScheduleSpec_ToProto,
		SpannerBackupScheduleObservedState_FromProto, SpannerBackupScheduleObservedState_ToProto,
	)

	f.SpecFields.Insert(".spec")
	f.SpecFields.Insert(".retention_duration")
	f.SpecFields.Insert(".encryption_config")
	f.SpecFields.Insert(".full_backup_spec")
	f.SpecFields.Insert(".incremental_backup_spec")

	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".cron_spec")
	f.StatusFields.Insert(".spec.cron_spec.time_zone")
	f.StatusFields.Insert(".spec.cron_spec.creation_window")

	f.UnimplementedFields.Insert(".name") // Identifier, output only

	return f
}
