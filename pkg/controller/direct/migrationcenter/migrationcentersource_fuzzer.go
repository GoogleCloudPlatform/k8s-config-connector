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
// proto.message: google.cloud.migrationcenter.v1.Source
// api.group: migrationcenter.cnrm.cloud.google.com

package migrationcenter

import (
	pb "cloud.google.com/go/migrationcenter/apiv1/migrationcenterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(migrationCenterSourceFuzzer())
}

func migrationCenterSourceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Source{},
		MigrationCenterSourceSpec_FromProto, MigrationCenterSourceSpec_ToProto,
		MigrationCenterSourceObservedState_FromProto, MigrationCenterSourceObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".type")
	f.SpecField(".priority")
	f.SpecField(".managed")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".pending_frame_count")
	f.StatusField(".error_frame_count")
	f.StatusField(".state")

	f.IdentityField(".name")

	f.FilterSpec = func(in *pb.Source) {
		cleanEmptyMessages(in.ProtoReflect())
	}

	f.FilterStatus = func(in *pb.Source) {
		cleanEmptyMessages(in.ProtoReflect())
	}

	return f
}
