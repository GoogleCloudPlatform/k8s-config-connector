// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package tpgresource

import "testing"

func TestIsZone(t *testing.T) {
	tests := []struct {
		location string
		want     bool
	}{
		{location: "global", want: false},
		{location: "us-central1", want: false},
		{location: "u-region-1", want: false},
		{location: "u-france-south1", want: false},
		{location: "eu-sovereign-central1", want: false},
		{location: "us-central1-a", want: true},
		{location: "u-region-1-a", want: true},
		{location: "u-france-south1-b", want: true},
		{location: "eu-sovereign-central1-c", want: true},
	}

	for _, tc := range tests {
		if got := IsZone(tc.location); got != tc.want {
			t.Errorf("IsZone(%q) = %v, want %v", tc.location, got, tc.want)
		}
	}
}
