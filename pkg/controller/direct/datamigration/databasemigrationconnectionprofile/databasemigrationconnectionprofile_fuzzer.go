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

package databasemigrationconnectionprofile

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/datamigration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzDatabaseMigrationConnectionProfile())
}

func fuzzDatabaseMigrationConnectionProfile() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ConnectionProfile{},
		datamigration.DatabaseMigrationConnectionProfileSpec_FromProto, datamigration.DatabaseMigrationConnectionProfileSpec_ToProto,
		datamigration.DatabaseMigrationConnectionProfileObservedState_FromProto, datamigration.DatabaseMigrationConnectionProfileObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")

	f.SpecField(".labels")
	f.SpecField(".display_name")
	f.SpecField(".mysql")
	f.SpecField(".postgresql")
	f.SpecField(".oracle")
	f.SpecField(".cloudsql")
	f.SpecField(".alloydb")
	f.SpecField(".provider")

	f.StatusField(".state")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".error")

	f.Unimplemented_NotYetTriaged(".error.details")
	f.Unimplemented_NotYetTriaged(".error.details[].value")
	f.Unimplemented_NotYetTriaged(".error.details[].type_url")

	return f
}
