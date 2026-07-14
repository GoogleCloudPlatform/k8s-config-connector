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

package databasemigrationconversionworkspace

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/datamigration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzDatabaseMigrationConversionWorkspace())
}

func fuzzDatabaseMigrationConversionWorkspace() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ConversionWorkspace{},
		datamigration.DatabaseMigrationConversionWorkspaceSpec_FromProto, datamigration.DatabaseMigrationConversionWorkspaceSpec_ToProto,
		datamigration.DatabaseMigrationConversionWorkspaceObservedState_FromProto, datamigration.DatabaseMigrationConversionWorkspaceObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")

	f.SpecField(".source")
	f.SpecField(".source.engine")
	f.SpecField(".source.version")
	f.SpecField(".destination")
	f.SpecField(".destination.engine")
	f.SpecField(".destination.version")
	f.SpecField(".global_settings")
	f.SpecField(".display_name")

	f.StatusField(".has_uncommitted_changes")
	f.StatusField(".latest_commit_id")
	f.StatusField(".latest_commit_time")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	return f
}
