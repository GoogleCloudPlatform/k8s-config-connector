package v1beta1

import (
	"testing"
)

func TestApigeeOrganizationIdentity_FromExternal(t *testing.T) {
	testCases := []struct {
		name       string
		external   string
		expectErr  bool
		expectID   string
	}{
		{
			name:      "valid external ref",
			external:  "organizations/my-org",
			expectErr: false,
			expectID:  "my-org",
		},
		{
			name:      "full url",
			external:  "//apigee.googleapis.com/organizations/my-org",
			expectErr: false,
			expectID:  "my-org",
		},
		{
			name:      "invalid format",
			external:  "projects/my-project",
			expectErr: true,
		},
		{
			name:      "invalid collection",
			external:  "instances/my-instance",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id := &ApigeeOrganizationIdentity{}
			err := id.FromExternal(tc.external)
			if tc.expectErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if id.ResourceID != tc.expectID {
					t.Errorf("expected ResourceID %q, got %q", tc.expectID, id.ResourceID)
				}
			}
		})
	}
}
