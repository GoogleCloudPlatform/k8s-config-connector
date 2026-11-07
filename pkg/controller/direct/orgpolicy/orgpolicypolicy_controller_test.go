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

package orgpolicy

import (
	"testing"
	"time"

	orgpolicypb "cloud.google.com/go/orgpolicy/apiv2/orgpolicypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestPolicySpecEqual(t *testing.T) {
	now := timestamppb.New(time.Now())

	tests := []struct {
		name    string
		actual  *orgpolicypb.PolicySpec
		desired *orgpolicypb.PolicySpec
		want    bool
	}{
		{
			name:    "both nil",
			actual:  nil,
			desired: nil,
			want:    true,
		},
		{
			name:   "actual nil, desired non-nil",
			actual: nil,
			desired: &orgpolicypb.PolicySpec{
				Reset_: true,
			},
			want: false,
		},
		{
			name: "actual non-nil, desired nil",
			actual: &orgpolicypb.PolicySpec{
				Reset_: true,
			},
			desired: nil,
			want:    false,
		},
		{
			name: "equal when actual has output-only etag and update_time",
			actual: &orgpolicypb.PolicySpec{
				Etag:       "actual-etag",
				UpdateTime: now,
				Rules: []*orgpolicypb.PolicySpec_PolicyRule{
					{
						Kind: &orgpolicypb.PolicySpec_PolicyRule_Enforce{
							Enforce: true,
						},
					},
				},
			},
			desired: &orgpolicypb.PolicySpec{
				Rules: []*orgpolicypb.PolicySpec_PolicyRule{
					{
						Kind: &orgpolicypb.PolicySpec_PolicyRule_Enforce{
							Enforce: true,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "not equal when rules differ despite output-only fields",
			actual: &orgpolicypb.PolicySpec{
				Etag:       "actual-etag",
				UpdateTime: now,
				Rules: []*orgpolicypb.PolicySpec_PolicyRule{
					{
						Kind: &orgpolicypb.PolicySpec_PolicyRule_Enforce{
							Enforce: true,
						},
					},
				},
			},
			desired: &orgpolicypb.PolicySpec{
				Rules: []*orgpolicypb.PolicySpec_PolicyRule{
					{
						Kind: &orgpolicypb.PolicySpec_PolicyRule_Enforce{
							Enforce: false,
						},
					},
				},
			},
			want: false,
		},
		{
			name: "not equal when inheritFromParent differs",
			actual: &orgpolicypb.PolicySpec{
				Etag:              "actual-etag",
				UpdateTime:        now,
				InheritFromParent: true,
			},
			desired: &orgpolicypb.PolicySpec{
				InheritFromParent: false,
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := policySpecEqual(tc.actual, tc.desired); got != tc.want {
				t.Errorf("policySpecEqual() = %v, want %v", got, tc.want)
			}
		})
	}
}
