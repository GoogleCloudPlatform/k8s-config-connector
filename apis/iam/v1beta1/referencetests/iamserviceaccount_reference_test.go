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

package referencetests

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
)

func TestIAMServiceAccountRef_ParseAndAsHelpers(t *testing.T) {
	tests := []struct {
		name                     string
		external                 string
		wantProject              string
		wantAccountID            string
		wantEmail                string
		wantRelativeResourceName string
		wantErr                  bool
	}{
		{
			name:                     "Relative resource name with account ID",
			external:                 "projects/my-project/serviceAccounts/my-sa",
			wantProject:              "my-project",
			wantAccountID:            "my-sa",
			wantEmail:                "my-sa@my-project.iam.gserviceaccount.com",
			wantRelativeResourceName: "projects/my-project/serviceAccounts/my-sa",
			wantErr:                  false,
		},
		{
			name:                     "Relative resource name with host prefix",
			external:                 "https://iam.googleapis.com/projects/my-project/serviceAccounts/my-sa",
			wantProject:              "my-project",
			wantAccountID:            "my-sa",
			wantEmail:                "my-sa@my-project.iam.gserviceaccount.com",
			wantRelativeResourceName: "projects/my-project/serviceAccounts/my-sa",
			wantErr:                  false,
		},
		{
			name:     "Invalid format: relative resource name with email",
			external: "projects/my-project/serviceAccounts/my-sa@my-project.iam.gserviceaccount.com",
			wantErr:  true,
		},
		{
			name:     "Invalid format: email address format",
			external: "my-sa@my-project.iam.gserviceaccount.com",
			wantErr:  true,
		},
		{
			name:     "Invalid relative resource name (empty project)",
			external: "projects//serviceAccounts/my-sa",
			wantErr:  true,
		},
		{
			name:     "Invalid format entirely",
			external: "not-a-service-account",
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ref := &v1beta1.IAMServiceAccountRef{External: tc.external}
			err := ref.ValidateExternal(tc.external)
			if (err != nil) != tc.wantErr {
				t.Fatalf("ValidateExternal() error = %v, wantErr %v", err, tc.wantErr)
			}

			if tc.wantErr {
				return
			}

			email, err := ref.AsEmail()
			if err != nil {
				t.Fatalf("AsEmail() unexpected error: %v", err)
			}
			if email != tc.wantEmail {
				t.Errorf("AsEmail() = %q, want %q", email, tc.wantEmail)
			}

			rrn, err := ref.AsRelativeResourceName()
			if err != nil {
				t.Fatalf("AsRelativeResourceName() unexpected error: %v", err)
			}
			if rrn != tc.wantRelativeResourceName {
				t.Errorf("AsRelativeResourceName() = %q, want %q", rrn, tc.wantRelativeResourceName)
			}

			identity, err := ref.ParseExternalToIdentity()
			if err != nil {
				t.Fatalf("ParseExternalToIdentity() unexpected error: %v", err)
			}
			id, ok := identity.(*v1beta1.IAMServiceAccountIdentity)
			if !ok {
				t.Fatalf("expected *v1beta1.IAMServiceAccountIdentity, got %T", identity)
			}
			if id.Project != tc.wantProject {
				t.Errorf("Identity.Project = %q, want %q", id.Project, tc.wantProject)
			}
			if id.Account != tc.wantAccountID {
				t.Errorf("Identity.Account = %q, want %q", id.Account, tc.wantAccountID)
			}
		})
	}
}

func TestLegacyIAMServiceAccountRef_ParseAndAsHelpers(t *testing.T) {
	tests := []struct {
		name                     string
		external                 string
		wantProject              string
		wantAccountID            string
		wantEmail                string
		wantRelativeResourceName string
		wantErr                  bool
	}{
		{
			name:                     "Email address format",
			external:                 "my-sa@my-project.iam.gserviceaccount.com",
			wantProject:              "my-project",
			wantAccountID:            "my-sa",
			wantEmail:                "my-sa@my-project.iam.gserviceaccount.com",
			wantRelativeResourceName: "projects/my-project/serviceAccounts/my-sa",
			wantErr:                  false,
		},
		{
			name:     "Invalid format: relative resource name with account ID",
			external: "projects/my-project/serviceAccounts/my-sa",
			wantErr:  true,
		},
		{
			name:     "Invalid format: relative resource name with host prefix",
			external: "https://iam.googleapis.com/projects/my-project/serviceAccounts/my-sa",
			wantErr:  true,
		},
		{
			name:     "Invalid format: relative resource name with email",
			external: "projects/my-project/serviceAccounts/my-sa@my-project.iam.gserviceaccount.com",
			wantErr:  true,
		},
		{
			name:     "Invalid email (empty project/account)",
			external: "@.iam.gserviceaccount.com",
			wantErr:  true,
		},
		{
			name:     "Invalid format entirely",
			external: "not-a-service-account",
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ref := &v1beta1.LegacyIAMServiceAccountRef{External: tc.external}
			err := ref.ValidateExternal(tc.external)
			if (err != nil) != tc.wantErr {
				t.Fatalf("ValidateExternal() error = %v, wantErr %v", err, tc.wantErr)
			}

			if tc.wantErr {
				return
			}

			email, err := ref.AsEmail()
			if err != nil {
				t.Fatalf("AsEmail() unexpected error: %v", err)
			}
			if email != tc.wantEmail {
				t.Errorf("AsEmail() = %q, want %q", email, tc.wantEmail)
			}

			rrn, err := ref.AsRelativeResourceName()
			if err != nil {
				t.Fatalf("AsRelativeResourceName() unexpected error: %v", err)
			}
			if rrn != tc.wantRelativeResourceName {
				t.Errorf("AsRelativeResourceName() = %q, want %q", rrn, tc.wantRelativeResourceName)
			}

			identity, err := ref.ParseExternalToIdentity()
			if err != nil {
				t.Fatalf("ParseExternalToIdentity() unexpected error: %v", err)
			}
			id, ok := identity.(*v1beta1.IAMServiceAccountIdentity)
			if !ok {
				t.Fatalf("expected *v1beta1.IAMServiceAccountIdentity, got %T", identity)
			}
			if id.Project != tc.wantProject {
				t.Errorf("Identity.Project = %q, want %q", id.Project, tc.wantProject)
			}
			if id.Account != tc.wantAccountID {
				t.Errorf("Identity.Account = %q, want %q", id.Account, tc.wantAccountID)
			}
		})
	}
}
