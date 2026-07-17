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
// proto.message: google.cloud.redis.cluster.v1.Cluster
// api.group: redis.cnrm.cloud.google.com

package redis

import (
	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(redisClusterEndpointFuzzer())
}

func redisClusterEndpointFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Cluster{},
		RedisClusterEndpointSpec_FromProto, RedisClusterEndpointSpec_ToProto,
		RedisClusterEndpointObservedState_FromProto, RedisClusterEndpointObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // Identifier
	f.UnimplementedFields.Insert(".authorization_mode")
	f.UnimplementedFields.Insert(".transit_encryption_mode")
	f.UnimplementedFields.Insert(".shard_count")
	f.UnimplementedFields.Insert(".psc_configs")
	f.UnimplementedFields.Insert(".node_type")
	f.UnimplementedFields.Insert(".persistence_config")
	f.UnimplementedFields.Insert(".redis_configs")
	f.UnimplementedFields.Insert(".replica_count")
	f.UnimplementedFields.Insert(".zone_distribution_config")
	f.UnimplementedFields.Insert(".deletion_protection_enabled")
	f.UnimplementedFields.Insert(".automated_backup_config")
	f.UnimplementedFields.Insert(".maintenance_policy")
	f.UnimplementedFields.Insert(".kms_key")
	f.UnimplementedFields.Insert(".cross_cluster_replication_config")
	f.UnimplementedFields.Insert(".create_time")
	f.UnimplementedFields.Insert(".state")
	f.UnimplementedFields.Insert(".uid")
	f.UnimplementedFields.Insert(".size_gb")
	f.UnimplementedFields.Insert(".discovery_endpoints")
	f.UnimplementedFields.Insert(".psc_connections")
	f.UnimplementedFields.Insert(".psc_service_attachments")
	f.UnimplementedFields.Insert(".state_info")
	f.UnimplementedFields.Insert(".precise_size_gb")
	f.UnimplementedFields.Insert(".maintenance_schedule")
	f.UnimplementedFields.Insert(".encryption_info")
	f.UnimplementedFields.Insert(".backup_collection")
	f.UnimplementedFields.Insert(".managed_backup_source")
	f.UnimplementedFields.Insert(".gcs_source")
	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".rotate_server_certificate")
	f.UnimplementedFields.Insert(".server_ca_mode")
	f.UnimplementedFields.Insert(".server_ca_pool")

	f.SpecFields.Insert(".cluster_endpoints")
	f.StatusFields.Insert(".cluster_endpoints")

	return f
}
