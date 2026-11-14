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
	"testing"

	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func TestCompareRedisCluster_UnsetClusterEndpointsDoesNotClearLiveEndpoints(t *testing.T) {
	ctx := context.Background()
	actual := redisClusterWithUserCreatedPSC()

	// spec.clusterEndpoints omitted: ToProto leaves the repeated field empty.
	desired := &pb.Cluster{}

	diffs, err := compareRedisCluster(ctx, actual, desired)
	if err != nil {
		t.Fatalf("compareRedisCluster() unexpected error: %v", err)
	}
	if hasDiffField(diffs, "cluster_endpoints") {
		t.Fatalf("unset spec.clusterEndpoints must not diff cluster_endpoints (that update deletes user-created PSC connections); fields: %v", diffFieldIDs(diffs))
	}
}

func TestCompareRedisCluster_ExplicitClusterEndpointsStillDiff(t *testing.T) {
	ctx := context.Background()
	actual := redisClusterWithUserCreatedPSC()
	desired := &pb.Cluster{
		ClusterEndpoints: []*pb.ClusterEndpoint{
			{
				Connections: []*pb.ConnectionDetail{
					{
						Connection: &pb.ConnectionDetail_PscAutoConnection{
							PscAutoConnection: &pb.PscAutoConnection{
								ProjectId: "other-project",
								Network:   "projects/other-project/global/networks/other",
							},
						},
					},
				},
			},
		},
	}

	diffs, err := compareRedisCluster(ctx, actual, desired)
	if err != nil {
		t.Fatalf("compareRedisCluster() unexpected error: %v", err)
	}
	if !hasDiffField(diffs, "cluster_endpoints") {
		t.Fatalf("an explicit spec.clusterEndpoints that differs from the live cluster must still diff cluster_endpoints; fields: %v", diffFieldIDs(diffs))
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

func hasDiffField(diff *structuredreporting.Diff, id string) bool {
	if diff == nil {
		return false
	}
	for _, field := range diff.Fields {
		if field.ID == id {
			return true
		}
	}
	return false
}

func diffFieldIDs(diff *structuredreporting.Diff) []string {
	if diff == nil {
		return nil
	}
	ids := make([]string, 0, len(diff.Fields))
	for _, field := range diff.Fields {
		ids = append(ids, field.ID)
	}
	return ids
}
