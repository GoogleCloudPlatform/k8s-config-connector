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
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestIsPSCForwardingRule(t *testing.T) {
	tests := []struct {
		name string
		spec *krm.ComputeForwardingRuleSpec
		want bool
	}{
		{
			name: "nil spec",
			spec: nil,
			want: false,
		},
		{
			name: "no target",
			spec: &krm.ComputeForwardingRuleSpec{
				Location: direct.LazyPtr("us-central1"),
			},
			want: false,
		},
		{
			name: "serviceAttachmentRef target",
			spec: &krm.ComputeForwardingRuleSpec{
				Location: direct.LazyPtr("us-central1"),
				Target: &krm.Target{
					ServiceAttachmentRef: &refs.ComputeServiceAttachmentRef{
						Name: "my-sa",
					},
				},
			},
			want: true,
		},
		{
			name: "memorystoreInstanceServiceAttachment target",
			spec: &krm.ComputeForwardingRuleSpec{
				Location: direct.LazyPtr("us-central1"),
				Target: &krm.Target{
					MemorystoreInstanceServiceAttachment: &krm.MemorystoreInstanceServiceAttachment{},
				},
			},
			want: true,
		},
		{
			name: "redisClusterServiceAttachment target",
			spec: &krm.ComputeForwardingRuleSpec{
				Location: direct.LazyPtr("us-central1"),
				Target: &krm.Target{
					RedisClusterServiceAttachment: &krm.RedisClusterServiceAttachment{},
				},
			},
			want: true,
		},
		{
			name: "googleAPIsBundle all-apis target",
			spec: &krm.ComputeForwardingRuleSpec{
				Location: direct.LazyPtr("global"),
				Target: &krm.Target{
					GoogleAPIsBundle: direct.LazyPtr("all-apis"),
				},
			},
			want: true,
		},
		{
			name: "googleAPIsBundle vpc-sc target",
			spec: &krm.ComputeForwardingRuleSpec{
				Location: direct.LazyPtr("global"),
				Target: &krm.Target{
					GoogleAPIsBundle: direct.LazyPtr("vpc-sc"),
				},
			},
			want: true,
		},
		{
			name: "targetHTTPProxyRef target is not PSC",
			spec: &krm.ComputeForwardingRuleSpec{
				Location: direct.LazyPtr("global"),
				Target: &krm.Target{
					TargetHTTPProxyRef: &refs.ComputeTargetHTTPProxyRef{
						Name: "my-http-proxy",
					},
				},
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isPSCForwardingRule(tc.spec); got != tc.want {
				t.Errorf("isPSCForwardingRule() = %v, want %v", got, tc.want)
			}
		})
	}
}
