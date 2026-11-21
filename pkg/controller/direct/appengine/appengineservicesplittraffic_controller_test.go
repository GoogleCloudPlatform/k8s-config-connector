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

package appengine

import (
	"context"
	"testing"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestSplitTrafficDiff(t *testing.T) {
	ip := direct.LazyPtr("IP")
	cookie := direct.LazyPtr("COOKIE")

	tests := []struct {
		name string
		// desired spec.split
		desiredShardBy     *string
		desiredAllocations map[string]string
		// split returned by App Engine
		liveShardBy     pb.TrafficSplit_ShardBy
		liveAllocations map[string]float64

		wantChange bool
		// wantSentShardBy is the shardBy sent with the update, when there is a change.
		wantSentShardBy pb.TrafficSplit_ShardBy
	}{
		{
			name:               "one version, App Engine dropped shardBy: no change",
			desiredShardBy:     ip,
			desiredAllocations: map[string]string{"v2": "1"},
			liveShardBy:        pb.TrafficSplit_UNSPECIFIED,
			liveAllocations:    map[string]float64{"v2": 1},
		},
		{
			name:               "one version, shardBy omitted: no change",
			desiredAllocations: map[string]string{"v2": "1"},
			liveShardBy:        pb.TrafficSplit_UNSPECIFIED,
			liveAllocations:    map[string]float64{"v2": 1},
		},
		{
			name:               "traffic moves to another single version: change, sent with the desired shardBy",
			desiredShardBy:     ip,
			desiredAllocations: map[string]string{"v2": "1"},
			liveShardBy:        pb.TrafficSplit_UNSPECIFIED,
			liveAllocations:    map[string]float64{"v1": 1},
			wantChange:         true,
			wantSentShardBy:    pb.TrafficSplit_IP,
		},
		{
			name:               "one version to two versions: change, sent with the desired shardBy",
			desiredShardBy:     ip,
			desiredAllocations: map[string]string{"v1": "0.5", "v2": "0.5"},
			liveShardBy:        pb.TrafficSplit_UNSPECIFIED,
			liveAllocations:    map[string]float64{"v2": 1},
			wantChange:         true,
			wantSentShardBy:    pb.TrafficSplit_IP,
		},
		{
			name:               "two versions to one version: change, sent with the desired shardBy",
			desiredShardBy:     ip,
			desiredAllocations: map[string]string{"v2": "1"},
			liveShardBy:        pb.TrafficSplit_IP,
			liveAllocations:    map[string]float64{"v1": 0.5, "v2": 0.5},
			wantChange:         true,
			wantSentShardBy:    pb.TrafficSplit_IP,
		},
		{
			name:               "two versions, shardBy changed: change",
			desiredShardBy:     cookie,
			desiredAllocations: map[string]string{"v1": "0.5", "v2": "0.5"},
			liveShardBy:        pb.TrafficSplit_IP,
			liveAllocations:    map[string]float64{"v1": 0.5, "v2": 0.5},
			wantChange:         true,
			wantSentShardBy:    pb.TrafficSplit_COOKIE,
		},
		{
			name:               "two versions without shardBy are left alone: change",
			desiredShardBy:     ip,
			desiredAllocations: map[string]string{"v1": "0.5", "v2": "0.5"},
			liveShardBy:        pb.TrafficSplit_UNSPECIFIED,
			liveAllocations:    map[string]float64{"v1": 0.5, "v2": 0.5},
			wantChange:         true,
			wantSentShardBy:    pb.TrafficSplit_IP,
		},
		{
			name:               "no versions without shardBy are left alone: change",
			desiredShardBy:     ip,
			desiredAllocations: map[string]string{"v2": "1"},
			liveShardBy:        pb.TrafficSplit_UNSPECIFIED,
			wantChange:         true,
			wantSentShardBy:    pb.TrafficSplit_IP,
		},
		{
			name:               "two versions, unchanged: no change",
			desiredShardBy:     ip,
			desiredAllocations: map[string]string{"v1": "0.5", "v2": "0.5"},
			liveShardBy:        pb.TrafficSplit_IP,
			liveAllocations:    map[string]float64{"v1": 0.5, "v2": 0.5},
		},
		{
			name:               "one version, App Engine kept a different shardBy: change",
			desiredShardBy:     cookie,
			desiredAllocations: map[string]string{"v2": "1"},
			liveShardBy:        pb.TrafficSplit_IP,
			liveAllocations:    map[string]float64{"v2": 1},
			wantChange:         true,
			wantSentShardBy:    pb.TrafficSplit_COOKIE,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			live := &pb.Service{
				Name:  "apps/my-project/services/default",
				Id:    "default",
				Split: &pb.TrafficSplit{ShardBy: tc.liveShardBy, Allocations: tc.liveAllocations},
			}
			a := &splitTrafficAdapter{
				desired: &krm.AppEngineServiceSplitTraffic{
					Spec: krm.AppEngineServiceSplitTrafficSpec{
						ResourceID: direct.LazyPtr("default"),
						Split: &krm.AppEngineServiceSplitTrafficSplit{
							ShardBy:     tc.desiredShardBy,
							Allocations: tc.desiredAllocations,
						},
					},
				},
				actual: live,
			}

			diffRes, err := a.diffSplit(context.Background())
			if err != nil {
				t.Fatalf("diffSplit failed: %v", err)
			}

			if gotChange := !diffRes.Empty(); gotChange != tc.wantChange {
				t.Fatalf("change detected = %v, want %v (diff: %+v)", gotChange, tc.wantChange, diffRes.Diff)
			}
			if tc.wantChange {
				if got := diffRes.MergedDesired.GetSplit().GetShardBy(); got != tc.wantSentShardBy {
					t.Errorf("shardBy sent with the update = %v, want %v", got, tc.wantSentShardBy)
				}
			}
			if got := a.actual.GetSplit().GetShardBy(); got != tc.liveShardBy {
				t.Errorf("diffSplit changed the live service: shardBy = %v, want %v", got, tc.liveShardBy)
			}
		})
	}
}
