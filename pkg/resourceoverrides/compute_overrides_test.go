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

package resourceoverrides

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestComputeForwardingRulePSCOverrides(t *testing.T) {
	gvk := schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeForwardingRule",
	}

	tests := []struct {
		name                 string
		target               string
		initialConfig        map[string]interface{}
		initialRaw           map[string]interface{}
		expectedSchemeConfig interface{}
		expectedSchemeRaw    interface{}
	}{
		{
			name:   "PSC target without load_balancing_scheme should set it to empty string",
			target: "projects/my-project/regions/us-central1/serviceAttachments/my-sa",
			initialConfig: map[string]interface{}{
				"target": "projects/my-project/regions/us-central1/serviceAttachments/my-sa",
			},
			initialRaw: map[string]interface{}{
				"target": "projects/my-project/regions/us-central1/serviceAttachments/my-sa",
			},
			expectedSchemeConfig: "",
			expectedSchemeRaw:    "",
		},
		{
			name:   "PSC target with explicit load_balancing_scheme should preserve it",
			target: "projects/my-project/regions/us-central1/serviceAttachments/my-sa",
			initialConfig: map[string]interface{}{
				"target":                "projects/my-project/regions/us-central1/serviceAttachments/my-sa",
				"load_balancing_scheme": "INTERNAL",
			},
			initialRaw: map[string]interface{}{
				"target":                "projects/my-project/regions/us-central1/serviceAttachments/my-sa",
				"load_balancing_scheme": "INTERNAL",
			},
			expectedSchemeConfig: "INTERNAL",
			expectedSchemeRaw:    "INTERNAL",
		},
		{
			name:   "Non-PSC target should not set load_balancing_scheme to empty string",
			target: "projects/my-project/regions/us-central1/targetVpnGateways/my-vpn",
			initialConfig: map[string]interface{}{
				"target": "projects/my-project/regions/us-central1/targetVpnGateways/my-vpn",
			},
			initialRaw: map[string]interface{}{
				"target": "projects/my-project/regions/us-central1/targetVpnGateways/my-vpn",
			},
			expectedSchemeConfig: nil,
			expectedSchemeRaw:    nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			op := &operations.PreTerraformApply{
				TerraformConfig: &terraform.ResourceConfig{
					Config: tc.initialConfig,
					Raw:    tc.initialRaw,
				},
				LiveState: &terraform.InstanceState{},
			}
			if err := Handler.PreTerraformApply(context.Background(), gvk, op); err != nil {
				t.Fatalf("PreTerraformApply failed: %v", err)
			}
			if val := op.TerraformConfig.Config["load_balancing_scheme"]; val != tc.expectedSchemeConfig {
				t.Errorf("Config['load_balancing_scheme'] = %v, want %v", val, tc.expectedSchemeConfig)
			}
			if val := op.TerraformConfig.Raw["load_balancing_scheme"]; val != tc.expectedSchemeRaw {
				t.Errorf("Raw['load_balancing_scheme'] = %v, want %v", val, tc.expectedSchemeRaw)
			}
		})
	}
}
