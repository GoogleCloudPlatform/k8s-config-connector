// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package resourcemanager

import "testing"

func TestNormalizeIAMCustomRoleProjectAndName(t *testing.T) {
	projectTests := []struct {
		in   string
		want string
	}{
		{in: "sample-project", want: "sample-project"},
		{in: "partition:sample-project", want: "sample-project"},
		{in: "fr0:project-paris", want: "project-paris"},
	}
	for _, tc := range projectTests {
		if got := normalizeIAMCustomRoleProject(tc.in); got != tc.want {
			t.Errorf("normalizeIAMCustomRoleProject(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}

	nameTests := []struct {
		in   string
		want string
	}{
		{
			in:   "projects/sample-project/roles/kcc_audit_custom_role",
			want: "projects/sample-project/roles/kcc_audit_custom_role",
		},
		{
			in:   "projects/partition:sample-project/roles/kcc_audit_custom_role",
			want: "projects/sample-project/roles/kcc_audit_custom_role",
		},
		{
			in:   "organizations/123456789/roles/kcc_audit_custom_role",
			want: "organizations/123456789/roles/kcc_audit_custom_role",
		},
	}
	for _, tc := range nameTests {
		if got := normalizeIAMCustomRoleName(tc.in); got != tc.want {
			t.Errorf("normalizeIAMCustomRoleName(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
