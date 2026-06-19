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
// proto.message: google.cloud.redis.v1.Instance
// api.group: redis.cnrm.cloud.google.com

package redis

import (
	redispb "cloud.google.com/go/redis/apiv1/redispb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(redisInstanceFuzzer())
}

func redisInstanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&redispb.Instance{},
		RedisInstanceSpec_FromProto, RedisInstanceSpec_ToProto,
		InstanceObservedStateStatus_FromProto, InstanceObservedStateStatus_ToProto,
	)

	// Field comparisons:
	// - alternativeLocationId maps to .alternative_location_id
	// - authEnabled maps to .auth_enabled
	// - authString is NOT in redispb.Instance proto, but in separate RPC; handled separately or not mapped
	// - authorizedNetworkRef maps to .authorized_network
	// - connectMode maps to .connect_mode
	// - customerManagedKeyRef maps to .customer_managed_key
	// - displayName maps to .display_name
	// - locationId maps to .location_id
	// - maintenancePolicy maps to .maintenance_policy
	// - maintenanceSchedule maps to .maintenance_schedule
	// - memorySizeGb maps to .memory_size_gb
	// - persistenceConfig maps to .persistence_config
	// - readReplicasMode maps to .read_replicas_mode
	// - redisConfigs maps to .redis_configs
	// - redisVersion maps to .redis_version
	// - region maps to parent/URL resource name (not direct spec field)
	// - replicaCount maps to .replica_count
	// - reservedIpRange maps to .reserved_ip_range
	// - resourceID maps to GCP resource Name (handled by Unimplemented_Identity)
	// - secondaryIpRange maps to .secondary_ip_range
	// - tier maps to .tier
	// - transitEncryptionMode maps to .transit_encryption_mode

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".maintenance_schedule.can_reschedule")

	f.SpecField(".alternative_location_id")
	f.SpecField(".auth_enabled")
	f.SpecField(".authorized_network")
	f.SpecField(".connect_mode")
	f.SpecField(".customer_managed_key")
	f.SpecField(".display_name")
	f.SpecField(".location_id")
	f.SpecField(".maintenance_policy")
	f.SpecField(".maintenance_schedule")
	f.SpecField(".memory_size_gb")
	f.SpecField(".persistence_config")
	f.SpecField(".read_replicas_mode")
	f.SpecField(".redis_configs")
	f.SpecField(".redis_version")
	f.SpecField(".replica_count")
	f.SpecField(".reserved_ip_range")
	f.SpecField(".secondary_ip_range")
	f.SpecField(".tier")
	f.SpecField(".transit_encryption_mode")

	// Unimplemented top-level status fields and untriaged fields:
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_NotYetTriaged(".host")
	f.Unimplemented_NotYetTriaged(".port")
	f.Unimplemented_NotYetTriaged(".current_location_id")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".state")
	f.Unimplemented_NotYetTriaged(".status_message")
	f.Unimplemented_NotYetTriaged(".persistence_iam_identity")
	f.Unimplemented_NotYetTriaged(".server_ca_certs")
	f.Unimplemented_NotYetTriaged(".nodes")
	f.Unimplemented_NotYetTriaged(".read_endpoint")
	f.Unimplemented_NotYetTriaged(".read_endpoint_port")
	f.Unimplemented_NotYetTriaged(".suspension_reasons")
	f.Unimplemented_NotYetTriaged(".maintenance_version")
	f.Unimplemented_NotYetTriaged(".available_maintenance_versions")

	f.FilterSpec = func(in *redispb.Instance) {
		if in.MaintenancePolicy != nil {
			for _, w := range in.MaintenancePolicy.WeeklyMaintenanceWindow {
				if w != nil && w.StartTime != nil {
					if w.StartTime.Hours == 0 && w.StartTime.Minutes == 0 && w.StartTime.Seconds == 0 && w.StartTime.Nanos == 0 {
						w.StartTime = nil
					}
				}
			}
		}
	}

	return f
}
