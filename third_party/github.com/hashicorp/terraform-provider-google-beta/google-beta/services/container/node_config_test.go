// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package container

import (
	"testing"
)

func TestContainerNodePoolResourceLabelsDiffSuppress(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		key      string
		oldVal   string
		newVal   string
		expected bool
	}{
		"suppress server-added goog-gke provisioning model label": {
			key:      "node_config.0.resource_labels.goog-gke-node-pool-provisioning-model",
			oldVal:   "on-demand",
			newVal:   "",
			expected: true,
		},
		"suppress other server-added goog-gke label": {
			key:      "node_pool.0.node_config.0.resource_labels.goog-gke-volume",
			oldVal:   "true",
			newVal:   "",
			expected: true,
		},
		"do not suppress user-specified goog-gke label change": {
			key:      "node_config.0.resource_labels.goog-gke-node-pool-provisioning-model",
			oldVal:   "on-demand",
			newVal:   "spot",
			expected: false,
		},
		"suppress map count diff": {
			key:      "node_config.0.resource_labels.%",
			oldVal:   "2",
			newVal:   "1",
			expected: true,
		},
		"do not suppress user label removal": {
			key:      "node_config.0.resource_labels.service_name",
			oldVal:   "my-service",
			newVal:   "",
			expected: false,
		},
		"do not suppress user label modification": {
			key:      "node_config.0.resource_labels.service_name",
			oldVal:   "old-service",
			newVal:   "new-service",
			expected: false,
		},
		"do not suppress user label addition": {
			key:      "node_config.0.resource_labels.env",
			oldVal:   "",
			newVal:   "prod",
			expected: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := containerNodePoolResourceLabelsDiffSuppress(tc.key, tc.oldVal, tc.newVal, nil)
			if got != tc.expected {
				t.Errorf("containerNodePoolResourceLabelsDiffSuppress(%q, %q, %q) = %v, want %v", tc.key, tc.oldVal, tc.newVal, got, tc.expected)
			}
		})
	}
}
