// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicedirectory

import (
	"testing"
)

func TestNormalizeServiceDirectoryEndpointNetwork(t *testing.T) {
	cases := []struct {
		name           string
		input          string
		defaultProject string
		expected       string
	}{
		{
			name:           "Canonical format",
			input:          "projects/123456789012/locations/global/networks/test-network",
			defaultProject: "sample-project",
			expected:       "projects/123456789012/locations/global/networks/test-network",
		},
		{
			name:           "Canonical format with project ID",
			input:          "projects/sample-project/locations/global/networks/test-network",
			defaultProject: "sample-project",
			expected:       "projects/sample-project/locations/global/networks/test-network",
		},
		{
			name:           "Standard Compute selfLink",
			input:          "https://www.googleapis.com/compute/v1/projects/sample-project/global/networks/test-network",
			defaultProject: "",
			expected:       "projects/sample-project/locations/global/networks/test-network",
		},
		{
			name:           "Sovereign custom universe Compute selfLink",
			input:          "https://compute.custom.universe.goog/compute/v1/projects/sample-project/global/networks/test-network",
			defaultProject: "",
			expected:       "projects/sample-project/locations/global/networks/test-network",
		},
		{
			name:           "Relative compute path",
			input:          "projects/sample-project/global/networks/test-network",
			defaultProject: "",
			expected:       "projects/sample-project/locations/global/networks/test-network",
		},
		{
			name:           "Bare network name with default project",
			input:          "test-network",
			defaultProject: "sample-project",
			expected:       "projects/sample-project/locations/global/networks/test-network",
		},
		{
			name:           "Empty string",
			input:          "",
			defaultProject: "sample-project",
			expected:       "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := normalizeServiceDirectoryEndpointNetwork(tc.input, tc.defaultProject)
			if actual != tc.expected {
				t.Errorf("normalizeServiceDirectoryEndpointNetwork(%q, %q) = %q, expected %q",
					tc.input, tc.defaultProject, actual, tc.expected)
			}
		})
	}
}

func TestServiceDirectoryEndpointNetworkDiffSuppress(t *testing.T) {
	cases := []struct {
		name     string
		old      string
		new      string
		expected bool
	}{
		{
			name:     "Exact match canonical",
			old:      "projects/123456789012/locations/global/networks/test-network",
			new:      "projects/123456789012/locations/global/networks/test-network",
			expected: true,
		},
		{
			name:     "Project number from API vs project ID from config",
			old:      "projects/123456789012/locations/global/networks/test-network",
			new:      "projects/sample-project/locations/global/networks/test-network",
			expected: true,
		},
		{
			name:     "Project number from API vs Compute selfLink from config",
			old:      "projects/123456789012/locations/global/networks/test-network",
			new:      "https://compute.custom.universe.goog/compute/v1/projects/sample-project/global/networks/test-network",
			expected: true,
		},
		{
			name:     "Project number from API vs standard Compute selfLink",
			old:      "projects/123456789012/locations/global/networks/test-network",
			new:      "https://www.googleapis.com/compute/v1/projects/sample-project/global/networks/test-network",
			expected: true,
		},
		{
			name:     "Project number from API vs relative Compute path",
			old:      "projects/123456789012/locations/global/networks/test-network",
			new:      "projects/sample-project/global/networks/test-network",
			expected: true,
		},
		{
			name:     "Different network names should not be suppressed",
			old:      "projects/123456789012/locations/global/networks/test-network-1",
			new:      "projects/123456789012/locations/global/networks/test-network-2",
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := serviceDirectoryEndpointNetworkDiffSuppress("network", tc.old, tc.new, nil)
			if actual != tc.expected {
				t.Errorf("serviceDirectoryEndpointNetworkDiffSuppress(%q, %q) = %v, expected %v",
					tc.old, tc.new, actual, tc.expected)
			}
		})
	}
}
