// Copyright 2026 Google LLC
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

package bigquerymigration

import (
	pb "cloud.google.com/go/bigquery/migration/apiv2alpha/migrationpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzBigQueryMigrationMigrationWorkflow())
}

func fuzzBigQueryMigrationMigrationWorkflow() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.MigrationWorkflow{},
		BigQueryMigrationMigrationWorkflowSpec_FromProto, BigQueryMigrationMigrationWorkflowSpec_ToProto,
		BigQueryMigrationMigrationWorkflowObservedState_FromProto, BigQueryMigrationMigrationWorkflowObservedState_ToProto,
	)
	f.UnimplementedFields.Insert(".name")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".tasks")

	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".last_update_time")

	// Unimplemented nested tasks/translation fields not in KRM
	f.Unimplemented_NotYetTriaged(".tasks[].id")
	f.Unimplemented_NotYetTriaged(".tasks[].state")
	f.Unimplemented_NotYetTriaged(".tasks[].processing_error")
	f.Unimplemented_NotYetTriaged(".tasks[].orchestration_result")
	f.Unimplemented_NotYetTriaged(".tasks[].translation_task_details.special_token_map")

	f.FilterSpec = func(in *pb.MigrationWorkflow) {
		for _, task := range in.Tasks {
			if task != nil {
				task.Id = ""
				task.State = pb.MigrationTask_STATE_UNSPECIFIED
				task.ProcessingError = nil
				task.OrchestrationResult = nil
				if details := task.GetTranslationTaskDetails(); details != nil {
					details.SpecialTokenMap = nil
				}
			}
		}
	}

	return f
}
