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
// proto.message: google.cloud.migrationcenter.v1.PreferenceSet
// api.group: migrationcenter.cnrm.cloud.google.com

package migrationcenter

import (
	pb "cloud.google.com/go/migrationcenter/apiv1/migrationcenterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(migrationCenterPreferenceSetFuzzer())
}

func migrationCenterPreferenceSetFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.PreferenceSet{},
		MigrationCenterPreferenceSetSpec_FromProto, MigrationCenterPreferenceSetSpec_ToProto,
		MigrationCenterPreferenceSetObservedState_FromProto, MigrationCenterPreferenceSetObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".virtual_machine_preferences.target_product")
	f.SpecField(".virtual_machine_preferences.region_preferences.preferred_regions")
	f.SpecField(".virtual_machine_preferences.commitment_plan")
	f.SpecField(".virtual_machine_preferences.sizing_optimization_strategy")
	f.SpecField(".virtual_machine_preferences.compute_engine_preferences.machine_preferences.allowed_machine_series")
	f.SpecField(".virtual_machine_preferences.compute_engine_preferences.machine_preferences.allowed_machine_series.code")
	f.SpecField(".virtual_machine_preferences.compute_engine_preferences.license_type")
	f.SpecField(".virtual_machine_preferences.vmware_engine_preferences.cpu_overcommit_ratio")
	f.SpecField(".virtual_machine_preferences.vmware_engine_preferences.memory_overcommit_ratio")
	f.SpecField(".virtual_machine_preferences.vmware_engine_preferences.storage_deduplication_compression_ratio")
	f.SpecField(".virtual_machine_preferences.vmware_engine_preferences.commitment_plan")
	f.SpecField(".virtual_machine_preferences.sole_tenancy_preferences.cpu_overcommit_ratio")
	f.SpecField(".virtual_machine_preferences.sole_tenancy_preferences.host_maintenance_policy")
	f.SpecField(".virtual_machine_preferences.sole_tenancy_preferences.commitment_plan")
	f.SpecField(".virtual_machine_preferences.sole_tenancy_preferences.node_types")
	f.SpecField(".virtual_machine_preferences.sole_tenancy_preferences.node_types.node_name")

	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.Unimplemented_Identity(".name")

	f.FilterSpec = func(in *pb.PreferenceSet) {
		cleanEmptyMessages(in.ProtoReflect())
	}

	f.FilterStatus = func(in *pb.PreferenceSet) {
		cleanEmptyMessages(in.ProtoReflect())
	}

	return f
}
