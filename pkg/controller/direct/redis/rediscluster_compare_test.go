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

package redis

import (
	"context"
	"slices"
	"testing"

	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1beta1"
)

// An unset spec.clusterEndpoints must not be compared: otherwise the update mask
// would include cluster_endpoints and delete user-created PSC connections.
func TestCompareRedisCluster_UnsetClusterEndpointsIsNotCompared(t *testing.T) {
	res, err := compareRedisCluster(context.Background(), &krm.RedisClusterSpec{}, redisClusterWithUserCreatedPSC())
	if err != nil {
		t.Fatalf("compareRedisCluster() unexpected error: %v", err)
	}
	if slices.Contains(res.FieldMask.GetPaths(), "cluster_endpoints") {
		t.Fatalf("unset spec.clusterEndpoints must not diff cluster_endpoints; paths: %v", res.FieldMask.GetPaths())
	}
}

func redisClusterWithUserCreatedPSC() *pb.Cluster {
	return &pb.Cluster{
		ClusterEndpoints: []*pb.ClusterEndpoint{
			{
				Connections: []*pb.ConnectionDetail{
					{
						Connection: &pb.ConnectionDetail_PscAutoConnection{
							PscAutoConnection: &pb.PscAutoConnection{
								ProjectId:           "PROJECT",
								Network:             "projects/PROJECT/global/networks/NETWORK",
								PscConnectionStatus: pb.PscConnectionStatus_PSC_CONNECTION_STATUS_ACTIVE,
							},
						},
					},
				},
			},
			{
				Connections: []*pb.ConnectionDetail{
					{
						Connection: &pb.ConnectionDetail_PscConnection{
							PscConnection: &pb.PscConnection{
								ForwardingRule:      "projects/OTHER_PROJECT/regions/europe-west4/forwardingRules/fr-1",
								Network:             "projects/OTHER_PROJECT/global/networks/OTHER_NETWORK",
								ProjectId:           "OTHER_PROJECT",
								PscConnectionId:     "111",
								PscConnectionStatus: pb.PscConnectionStatus_PSC_CONNECTION_STATUS_ACTIVE,
							},
						},
					},
				},
			},
		},
	}
}
