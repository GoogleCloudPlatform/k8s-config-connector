// Copyright 2025 Google LLC
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
// proto.message: google.firestore.admin.v1.Index
// api.group: firestore.cnrm.cloud.google.com

package firestore

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	pb "google.golang.org/genproto/googleapis/firestore/admin/v1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(firestoreDatabaseFuzzer())
}

func firestoreDatabaseFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Database{},
		FirestoreDatabaseSpec_FromProto, FirestoreDatabaseSpec_ToProto,
		FirestoreDatabaseObservedState_FromProto, FirestoreDatabaseObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")

	f.UnimplementedFields.Insert(".delete_time")
	f.UnimplementedFields.Insert(".key_prefix")
	f.UnimplementedFields.Insert(".cmek_config")
	f.UnimplementedFields.Insert(".previous_id")
	f.UnimplementedFields.Insert(".source_info")

	// Default value fields set by controller
	f.UnimplementedFields.Insert(".type")
	f.UnimplementedFields.Insert(".app_engine_integration_mode")

	f.SpecField(".location_id")
	f.SpecField(".concurrency_mode")
	f.SpecField(".point_in_time_recovery_enablement")
	f.SpecField(".delete_protection_state")

	f.StatusField(".uid")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".version_retention_period")
	f.StatusField(".earliest_version_time")
	f.StatusField(".etag")

	f.Unimplemented_NotYetTriaged(".free_tier")
	f.Unimplemented_NotYetTriaged(".tags")
	f.Unimplemented_NotYetTriaged(".database_edition")

	return f
}
