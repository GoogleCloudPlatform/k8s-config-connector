// Copyright 2024 Google LLC
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
// proto.message: google.cloud.redis.cluster.v1.Cluster
// api.group: redis.cnrm.cloud.google.com

package cluster

import (
	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(redisClusterFuzzer())
}

func redisClusterFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Cluster{},
		RedisClusterSpec_FromProto, RedisClusterSpec_ToProto,
		RedisClusterObservedState_FromProto, RedisClusterObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // Identifier

	// New Fields in the updated version of Redis Cluster
	f.UnimplementedFields.Insert(".backup_collection")
	f.UnimplementedFields.Insert(".managed_backup_source")
	f.UnimplementedFields.Insert(".psc_service_attachments")
	f.UnimplementedFields.Insert(".psc_connections[].psc_connection_status")
	f.UnimplementedFields.Insert(".psc_connections[].service_attachment")
	f.UnimplementedFields.Insert(".psc_connections[].connection_type")
	f.UnimplementedFields.Insert(".cross_cluster_replication_config")
	f.UnimplementedFields.Insert(".kms_key")
	f.UnimplementedFields.Insert(".maintenance_policy")
	f.UnimplementedFields.Insert(".maintenance_schedule")
	f.UnimplementedFields.Insert(".automated_backup_config")
	f.UnimplementedFields.Insert(".encryption_info")
	f.UnimplementedFields.Insert(".gcs_source")
	f.UnimplementedFields.Insert(".cluster_endpoints")
	f.UnimplementedFields.Insert(".labels")

	f.SpecFields.Insert(".authorization_mode")
	f.SpecFields.Insert(".transit_encryption_mode")
	f.SpecFields.Insert(".shard_count")
	f.SpecFields.Insert(".psc_configs")
	f.SpecFields.Insert(".node_type")
	f.SpecFields.Insert(".persistence_config")
	f.SpecFields.Insert(".redis_configs")
	f.SpecFields.Insert(".replica_count")
	f.SpecFields.Insert(".zone_distribution_config")
	f.SpecFields.Insert(".deletion_protection_enabled")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".size_gb")
	f.StatusFields.Insert(".discovery_endpoints")
	f.StatusFields.Insert(".psc_connections")
	f.StatusFields.Insert(".state_info")
	f.StatusFields.Insert(".precise_size_gb")

	return f
}
