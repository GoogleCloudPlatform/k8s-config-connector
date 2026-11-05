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

package dataproc

import (
	"context"
	"testing"

	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestCompareDataprocAutoscalingPolicy(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name     string
		actual   *pb.AutoscalingPolicy
		desired  *pb.AutoscalingPolicy
		wantDiff bool
	}{
		{
			name: "identical policy",
			actual: &pb.AutoscalingPolicy{
				Id:   "test-policy",
				Name: "projects/test-project/regions/us-central1/autoscalingPolicies/test-policy",
				Algorithm: &pb.AutoscalingPolicy_BasicAlgorithm{
					BasicAlgorithm: &pb.BasicAutoscalingAlgorithm{
						CooldownPeriod: &durationpb.Duration{Seconds: 120},
						Config: &pb.BasicAutoscalingAlgorithm_YarnConfig{
							YarnConfig: &pb.BasicYarnAutoscalingConfig{
								GracefulDecommissionTimeout: &durationpb.Duration{Seconds: 3600},
								ScaleUpFactor:               0.5,
								ScaleDownFactor:             0.5,
							},
						},
					},
				},
				WorkerConfig: &pb.InstanceGroupAutoscalingPolicyConfig{
					MinInstances: 2,
					MaxInstances: 10,
					Weight:       1,
				},
				SecondaryWorkerConfig: &pb.InstanceGroupAutoscalingPolicyConfig{
					MinInstances: 0,
					MaxInstances: 5,
					Weight:       1,
				},
			},
			desired: &pb.AutoscalingPolicy{
				Id:   "test-policy",
				Name: "projects/test-project/regions/us-central1/autoscalingPolicies/test-policy",
				Algorithm: &pb.AutoscalingPolicy_BasicAlgorithm{
					BasicAlgorithm: &pb.BasicAutoscalingAlgorithm{
						CooldownPeriod: &durationpb.Duration{Seconds: 120},
						Config: &pb.BasicAutoscalingAlgorithm_YarnConfig{
							YarnConfig: &pb.BasicYarnAutoscalingConfig{
								GracefulDecommissionTimeout: &durationpb.Duration{Seconds: 3600},
								ScaleUpFactor:               0.5,
								ScaleDownFactor:             0.5,
							},
						},
					},
				},
				WorkerConfig: &pb.InstanceGroupAutoscalingPolicyConfig{
					MinInstances: 2,
					MaxInstances: 10,
					Weight:       1,
				},
				SecondaryWorkerConfig: &pb.InstanceGroupAutoscalingPolicyConfig{
					MinInstances: 0,
					MaxInstances: 5,
					Weight:       1,
				},
			},
			wantDiff: false,
		},
		{
			name: "actual has non-spec server fields (labels)",
			actual: &pb.AutoscalingPolicy{
				Id:     "test-policy",
				Name:   "projects/test-project/regions/us-central1/autoscalingPolicies/test-policy",
				Labels: map[string]string{"server-label": "value"},
				WorkerConfig: &pb.InstanceGroupAutoscalingPolicyConfig{
					MinInstances: 2,
					MaxInstances: 10,
					Weight:       1,
				},
			},
			desired: &pb.AutoscalingPolicy{
				Id:   "test-policy",
				Name: "projects/test-project/regions/us-central1/autoscalingPolicies/test-policy",
				WorkerConfig: &pb.InstanceGroupAutoscalingPolicyConfig{
					MinInstances: 2,
					MaxInstances: 10,
					Weight:       1,
				},
			},
			wantDiff: false,
		},
		{
			name: "worker config changed",
			actual: &pb.AutoscalingPolicy{
				Id:   "test-policy",
				Name: "projects/test-project/regions/us-central1/autoscalingPolicies/test-policy",
				WorkerConfig: &pb.InstanceGroupAutoscalingPolicyConfig{
					MinInstances: 2,
					MaxInstances: 10,
					Weight:       1,
				},
			},
			desired: &pb.AutoscalingPolicy{
				Id:   "test-policy",
				Name: "projects/test-project/regions/us-central1/autoscalingPolicies/test-policy",
				WorkerConfig: &pb.InstanceGroupAutoscalingPolicyConfig{
					MinInstances: 4,
					MaxInstances: 20,
					Weight:       1,
				},
			},
			wantDiff: true,
		},
		{
			name: "secondary worker config changed",
			actual: &pb.AutoscalingPolicy{
				Id:   "test-policy",
				Name: "projects/test-project/regions/us-central1/autoscalingPolicies/test-policy",
				SecondaryWorkerConfig: &pb.InstanceGroupAutoscalingPolicyConfig{
					MinInstances: 0,
					MaxInstances: 5,
					Weight:       1,
				},
			},
			desired: &pb.AutoscalingPolicy{
				Id:   "test-policy",
				Name: "projects/test-project/regions/us-central1/autoscalingPolicies/test-policy",
				SecondaryWorkerConfig: &pb.InstanceGroupAutoscalingPolicyConfig{
					MinInstances: 1,
					MaxInstances: 10,
					Weight:       2,
				},
			},
			wantDiff: true,
		},
		{
			name: "algorithm changed",
			actual: &pb.AutoscalingPolicy{
				Id:   "test-policy",
				Name: "projects/test-project/regions/us-central1/autoscalingPolicies/test-policy",
				Algorithm: &pb.AutoscalingPolicy_BasicAlgorithm{
					BasicAlgorithm: &pb.BasicAutoscalingAlgorithm{
						CooldownPeriod: &durationpb.Duration{Seconds: 120},
					},
				},
			},
			desired: &pb.AutoscalingPolicy{
				Id:   "test-policy",
				Name: "projects/test-project/regions/us-central1/autoscalingPolicies/test-policy",
				Algorithm: &pb.AutoscalingPolicy_BasicAlgorithm{
					BasicAlgorithm: &pb.BasicAutoscalingAlgorithm{
						CooldownPeriod: &durationpb.Duration{Seconds: 300},
					},
				},
			},
			wantDiff: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diff, err := compareDataprocAutoscalingPolicy(ctx, tc.actual, tc.desired)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if gotDiff := diff.HasDiff(); gotDiff != tc.wantDiff {
				t.Errorf("compareDataprocAutoscalingPolicy() diff.HasDiff() = %v, want %v (diff entries: %v)", gotDiff, tc.wantDiff, diff.Fields)
			}
		})
	}
}
