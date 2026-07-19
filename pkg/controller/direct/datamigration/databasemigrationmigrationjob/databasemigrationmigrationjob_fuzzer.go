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

package databasemigrationmigrationjob

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/datamigration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzDatabaseMigrationMigrationJob())
}

func fuzzDatabaseMigrationMigrationJob() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.MigrationJob{},
		datamigration.DatabaseMigrationMigrationJobSpec_FromProto, datamigration.DatabaseMigrationMigrationJobSpec_ToProto,
		datamigration.DatabaseMigrationMigrationJobObservedState_FromProto, datamigration.DatabaseMigrationMigrationJobObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")

	f.SpecField(".labels")
	f.SpecField(".display_name")
	f.SpecField(".type")
	f.SpecField(".dump_path")
	f.SpecField(".dump_flags")
	f.SpecField(".source")
	f.SpecField(".destination")
	f.SpecField(".reverse_ssh_connectivity")
	f.SpecField(".vpc_peering_connectivity")
	f.SpecField(".static_ip_connectivity")
	f.SpecField(".source_database")
	f.SpecField(".destination_database")
	f.SpecField(".conversion_workspace")
	f.SpecField(".filter")
	f.SpecField(".cmek_key_name")
	f.SpecField(".performance_config")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".phase")
	f.StatusField(".duration")
	f.StatusField(".error")
	f.StatusField(".end_time")
	f.StatusField(".state")

	f.Unimplemented_NotYetTriaged(".error.details")
	f.Unimplemented_NotYetTriaged(".error.details[].value")
	f.Unimplemented_NotYetTriaged(".error.details[].type_url")

	return f
}
