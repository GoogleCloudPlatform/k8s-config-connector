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
// proto.message: google.cloud.bigquery.migration.v2alpha.MigrationWorkflow
// api.group: bigquerymigration.cnrm.cloud.google.com

package bigquerymigration

import (
	pb "cloud.google.com/go/bigquery/migration/apiv2alpha/migrationpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(bigQueryMigrationMigrationWorkflowFuzzer())
}

func bigQueryMigrationMigrationWorkflowFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.MigrationWorkflow{},
		BigQueryMigrationMigrationWorkflowSpec_FromProto, BigQueryMigrationMigrationWorkflowSpec_ToProto,
		BigQueryMigrationMigrationWorkflowObservedState_FromProto, BigQueryMigrationMigrationWorkflowObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".tasks")

	f.StatusField(".state")
	f.StatusField(".create_time")
	f.StatusField(".last_update_time")

	f.Unimplemented_Identity(".name")

	f.FilterSpec = func(in *pb.MigrationWorkflow) {
		for _, task := range in.Tasks {
			task.Id = ""
			task.State = pb.MigrationTask_STATE_UNSPECIFIED
			task.ProcessingError = nil
			task.OrchestrationResult = nil
			if details, ok := task.TaskDetails.(*pb.MigrationTask_TranslationTaskDetails); ok && details != nil && details.TranslationTaskDetails != nil {
				details.TranslationTaskDetails.SpecialTokenMap = nil
			}
		}
		cleanEmptyMessages(in.ProtoReflect())
	}

	f.FilterStatus = func(in *pb.MigrationWorkflow) {
		cleanEmptyMessages(in.ProtoReflect())
	}

	return f
}

func cleanEmptyMessages(m protoreflect.Message) {
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.Kind() == protoreflect.MessageKind {
			if fd.IsList() || fd.IsMap() {
				return true
			}
			sub := v.Message()
			cleanEmptyMessages(sub)
			// check if sub has any populated fields now
			hasFields := false
			sub.Range(func(fd2 protoreflect.FieldDescriptor, v2 protoreflect.Value) bool {
				hasFields = true
				return false
			})
			if !hasFields {
				m.Clear(fd)
			}
		}
		return true
	})
}
