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
// proto.message: google.cloud.memorystore.v1.Instance
// api.group: memorystore.cnrm.cloud.google.com

package memorystore

import (
	pb "cloud.google.com/go/memorystore/apiv1/memorystorepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(memorystoreInstanceEndpointFuzzer())
}

func memorystoreInstanceEndpointFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Instance{},
		MemorystoreInstanceEndpointSpec_FromProto, MemorystoreInstanceEndpointSpec_ToProto,
		MemorystoreInstanceEndpointObservedState_FromProto, MemorystoreInstanceEndpointObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".allow_fewer_zones_deployment")
	f.UnimplementedFields.Insert(".async_instance_endpoints_deletion_enabled")
	f.UnimplementedFields.Insert(".authorization_mode")
	f.UnimplementedFields.Insert(".automated_backup_config")
	f.UnimplementedFields.Insert(".available_maintenance_versions")
	f.UnimplementedFields.Insert(".backup_collection")
	f.UnimplementedFields.Insert(".create_time")
	f.UnimplementedFields.Insert(".cross_instance_replication_config")
	f.UnimplementedFields.Insert(".deletion_protection_enabled")
	f.UnimplementedFields.Insert(".discovery_endpoints")
	f.UnimplementedFields.Insert(".effective_maintenance_version")
	f.UnimplementedFields.Insert(".encryption_info")
	f.UnimplementedFields.Insert(".endpoints[].connections[].psc_auto_connection")
	f.UnimplementedFields.Insert(".engine_configs")
	f.UnimplementedFields.Insert(".engine_version")
	f.UnimplementedFields.Insert(".gcs_source")
	f.UnimplementedFields.Insert(".kms_key")
	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".maintenance_policy")
	f.UnimplementedFields.Insert(".maintenance_schedule")
	f.UnimplementedFields.Insert(".maintenance_version")
	f.UnimplementedFields.Insert(".managed_backup_source")
	f.UnimplementedFields.Insert(".mode")
	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".node_config")
	f.UnimplementedFields.Insert(".node_type")
	f.UnimplementedFields.Insert(".ondemand_maintenance")
	f.UnimplementedFields.Insert(".persistence_config")
	f.UnimplementedFields.Insert(".psc_attachment_details")
	f.UnimplementedFields.Insert(".psc_auto_connections")
	f.UnimplementedFields.Insert(".replica_count")
	f.UnimplementedFields.Insert(".satisfies_pzi")
	f.UnimplementedFields.Insert(".satisfies_pzs")
	f.UnimplementedFields.Insert(".shard_count")
	f.UnimplementedFields.Insert(".simulate_maintenance_event")
	f.UnimplementedFields.Insert(".state")
	f.UnimplementedFields.Insert(".state_info")
	f.UnimplementedFields.Insert(".state_info.update_info")
	f.UnimplementedFields.Insert(".transit_encryption_mode")
	f.UnimplementedFields.Insert(".uid")
	f.UnimplementedFields.Insert(".update_time")
	f.UnimplementedFields.Insert(".zone_distribution_config")

	f.SpecFields.Insert(".endpoints")

	f.StatusFields.Insert(".endpoints")

	return f
}
