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
// proto.message: google.cloud.memorystore.v1.Instance
// api.group: memorystore.cnrm.cloud.google.com

package memorystore

import (
	pb "cloud.google.com/go/memorystore/apiv1/memorystorepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(memorystoreInstanceFuzzer())
}

func memorystoreInstanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Instance{},
		MemorystoreInstanceSpec_FromProto, MemorystoreInstanceSpec_ToProto,
		MemorystoreInstanceObservedState_FromProto, MemorystoreInstanceObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // Special field: resource name
	f.UnimplementedFields.Insert(".satisfies_pzi")
	f.UnimplementedFields.Insert(".satisfies_pzs")
	f.UnimplementedFields.Insert(".psc_auto_connections")
	f.UnimplementedFields.Insert(".discovery_endpoints")
	f.UnimplementedFields.Insert(".async_instance_endpoints_deletion_enabled")
	f.UnimplementedFields.Insert(".ondemand_maintenance")

	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".replica_count")
	f.SpecFields.Insert(".authorization_mode")
	f.SpecFields.Insert(".transit_encryption_mode")
	f.SpecFields.Insert(".shard_count")
	f.SpecFields.Insert(".node_type")
	f.SpecFields.Insert(".persistence_config")
	f.SpecFields.Insert(".engine_version")
	f.SpecFields.Insert(".engine_configs")
	f.SpecFields.Insert(".zone_distribution_config")
	f.SpecFields.Insert(".deletion_protection_enabled")
	f.SpecFields.Insert(".gcs_source")
	f.SpecFields.Insert(".managed_backup_source")
	f.SpecFields.Insert(".endpoints")
	f.SpecFields.Insert(".mode")
	f.SpecFields.Insert(".maintenance_policy")
	f.SpecFields.Insert(".cross_instance_replication_config")
	f.SpecFields.Insert(".automated_backup_config")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".state_info")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".node_config")
	f.StatusFields.Insert(".psc_attachment_details")
	f.StatusFields.Insert(".endpoints")
	f.StatusFields.Insert(".maintenance_policy")
	f.StatusFields.Insert(".maintenance_schedule")
	f.StatusFields.Insert(".cross_instance_replication_config")
	f.StatusFields.Insert(".backup_collection")

	f.Unimplemented_NotYetTriaged(".available_maintenance_versions")
	f.Unimplemented_NotYetTriaged(".effective_maintenance_version")
	f.Unimplemented_NotYetTriaged(".maintenance_version")
	f.Unimplemented_NotYetTriaged(".allow_fewer_zones_deployment")
	f.Unimplemented_NotYetTriaged(".encryption_info")
	f.Unimplemented_NotYetTriaged(".simulate_maintenance_event")
	f.Unimplemented_NotYetTriaged(".kms_key")

	return f
}
