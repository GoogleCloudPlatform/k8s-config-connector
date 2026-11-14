// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package resourcemanager

import "testing"

func TestNormalizeResourceManagerLienParent(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "projects/sample-project",
			want: "projects/sample-project",
		},
		{
			in:   "sample-project",
			want: "projects/sample-project",
		},
		{
			in:   "123456789012",
			want: "projects/123456789012",
		},
		{
			in:   "projects/123456789012",
			want: "projects/123456789012",
		},
		{
			in:   "projects/partition:sample-project",
			want: "projects/sample-project",
		},
		{
			in:   "partition:sample-project",
			want: "projects/sample-project",
		},
	}

	for _, tc := range tests {
		got, err := normalizeResourceManagerLienParent(tc.in, nil, nil)
		if err != nil {
			t.Fatalf("normalizeResourceManagerLienParent(%q) unexpected error: %v", tc.in, err)
		}
		if got != tc.want {
			t.Errorf("normalizeResourceManagerLienParent(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
