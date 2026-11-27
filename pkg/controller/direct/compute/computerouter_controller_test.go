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

package compute

import (
	"context"
	"testing"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
)

func TestCompareComputeRouter(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name     string
		actual   *computepb.Router
		desired  *computepb.Router
		wantDiff bool
	}{
		{
			name: "equal with full URL network and region on actual, relative on desired",
			actual: &computepb.Router{
				Name:        proto.String("r1"),
				Description: proto.String("desc"),
				Network:     proto.String("https://www.googleapis.com/compute/v1/projects/p1/global/networks/default"),
				Region:      proto.String("https://www.googleapis.com/compute/v1/projects/p1/regions/us-west1"),
			},
			desired: &computepb.Router{
				Name:        proto.String("r1"),
				Description: proto.String("desc"),
				Network:     proto.String("projects/p1/global/networks/default"),
				Region:      proto.String("us-west1"),
			},
			wantDiff: false,
		},
		{
			name: "equal with short name network and region",
			actual: &computepb.Router{
				Name:        proto.String("r1"),
				Description: proto.String("desc"),
				Network:     proto.String("https://www.googleapis.com/compute/v1/projects/p1/global/networks/default"),
				Region:      proto.String("https://www.googleapis.com/compute/v1/projects/p1/regions/us-west1"),
			},
			desired: &computepb.Router{
				Name:        proto.String("r1"),
				Description: proto.String("desc"),
				Network:     proto.String("default"),
				Region:      proto.String("us-west1"),
			},
			wantDiff: false,
		},
		{
			name: "diff in description",
			actual: &computepb.Router{
				Name:        proto.String("r1"),
				Description: proto.String("old desc"),
				Network:     proto.String("https://www.googleapis.com/compute/v1/projects/p1/global/networks/default"),
				Region:      proto.String("https://www.googleapis.com/compute/v1/projects/p1/regions/us-west1"),
			},
			desired: &computepb.Router{
				Name:        proto.String("r1"),
				Description: proto.String("new desc"),
				Network:     proto.String("projects/p1/global/networks/default"),
				Region:      proto.String("us-west1"),
			},
			wantDiff: true,
		},
		{
			name: "diff in network",
			actual: &computepb.Router{
				Name:        proto.String("r1"),
				Description: proto.String("desc"),
				Network:     proto.String("https://www.googleapis.com/compute/v1/projects/p1/global/networks/default"),
				Region:      proto.String("https://www.googleapis.com/compute/v1/projects/p1/regions/us-west1"),
			},
			desired: &computepb.Router{
				Name:        proto.String("r1"),
				Description: proto.String("desc"),
				Network:     proto.String("projects/p1/global/networks/other"),
				Region:      proto.String("us-west1"),
			},
			wantDiff: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diff, _, err := compareComputeRouter(ctx, tc.actual, tc.desired)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff.HasDiff() != tc.wantDiff {
				t.Errorf("HasDiff() = %v, want %v", diff.HasDiff(), tc.wantDiff)
			}
		})
	}
}
