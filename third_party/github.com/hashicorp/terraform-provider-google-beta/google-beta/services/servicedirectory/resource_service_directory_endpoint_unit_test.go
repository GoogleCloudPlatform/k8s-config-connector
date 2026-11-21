// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicedirectory

import "testing"

func TestExpandServiceDirectoryEndpointNetwork(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "compute v1 selfLink",
			input: "https://www.googleapis.com/compute/v1/projects/123456789/global/networks/my-net",
			want:  "projects/123456789/global/networks/my-net",
		},
		{
			name:  "custom universe compute selfLink",
			input: "https://compute.custom.universe.goog/compute/v1/projects/123456789/global/networks/my-net",
			want:  "projects/123456789/global/networks/my-net",
		},
		{
			name:  "relative compute network path",
			input: "projects/123456789/global/networks/my-net",
			want:  "projects/123456789/global/networks/my-net",
		},
		{
			name:  "service directory locations/global network path",
			input: "projects/123456789/locations/global/networks/my-net",
			want:  "projects/123456789/locations/global/networks/my-net",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := expandServiceDirectoryEndpointNetwork(tc.input, nil, nil)
			if err != nil {
				t.Fatalf("expandServiceDirectoryEndpointNetwork(%q) unexpected error: %v", tc.input, err)
			}
			if got != tc.want {
				t.Errorf("expandServiceDirectoryEndpointNetwork(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestCompareServiceDirectoryEndpointNetwork(t *testing.T) {
	oldVal := "projects/123456789/locations/global/networks/my-net"
	newSelfLink := "https://www.googleapis.com/compute/v1/projects/123456789/global/networks/my-net"
	newCustomSelfLink := "https://compute.custom.universe.goog/compute/v1/projects/123456789/global/networks/my-net"
	differentNet := "projects/123456789/locations/global/networks/other-net"

	if !compareServiceDirectoryEndpointNetwork("network", oldVal, newSelfLink, nil) {
		t.Errorf("expected diff suppression between %q and %q", oldVal, newSelfLink)
	}
	if !compareServiceDirectoryEndpointNetwork("network", oldVal, newCustomSelfLink, nil) {
		t.Errorf("expected diff suppression between %q and %q", oldVal, newCustomSelfLink)
	}
	if compareServiceDirectoryEndpointNetwork("network", oldVal, differentNet, nil) {
		t.Errorf("expected diff between %q and %q", oldVal, differentNet)
	}
}
