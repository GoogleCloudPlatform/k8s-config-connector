// Copyright 2024 Google LLC
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

package sql

import (
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/sqladmin/v1beta4"
)

func TestParseInstanceExternal(t *testing.T) {
	tests := []struct {
		name             string
		external         string
		wantInstanceID   string
		wantProjectID    string
	}{
		{
			name:           "plain instance name",
			external:       "my-instance",
			wantInstanceID: "my-instance",
			wantProjectID:  "",
		},
		{
			name:           "relative path",
			external:       "projects/my-project/instances/my-instance",
			wantInstanceID: "my-instance",
			wantProjectID:  "my-project",
		},
		{
			name:           "full selfLink URL",
			external:       "https://sqladmin.googleapis.com/sql/v1beta4/projects/my-project/instances/my-instance",
			wantInstanceID: "my-instance",
			wantProjectID:  "my-project",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			instanceID, projectID := parseInstanceExternal(tc.external)
			if instanceID != tc.wantInstanceID {
				t.Errorf("instanceID = %q, want %q", instanceID, tc.wantInstanceID)
			}
			if projectID != tc.wantProjectID {
				t.Errorf("projectID = %q, want %q", projectID, tc.wantProjectID)
			}
		})
	}
}

func TestSQLUserKRMToGCP_ForceSendFields(t *testing.T) {
	tests := []struct {
		name                string
		password            string
		policyAllowed       *int64
		policyEnableCheck   *bool
		wantForceSendUser   []string
		wantForceSendPolicy []string
	}{
		{
			name:              "password set forces send",
			password:          "secret",
			wantForceSendUser: []string{"Password"},
		},
		{
			name:              "empty password does not force send",
			password:          "",
			wantForceSendUser: nil,
		},
		{
			name:                "zero AllowedFailedAttempts forces send",
			policyAllowed:       direct.PtrTo(int64(0)),
			policyEnableCheck:   direct.PtrTo(false),
			wantForceSendPolicy: []string{"AllowedFailedAttempts", "EnableFailedAttemptsCheck"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			obj := &krm.SQLUser{}
			obj.Spec.ResourceID = direct.PtrTo("testuser")
			if tc.policyAllowed != nil || tc.policyEnableCheck != nil {
				obj.Spec.PasswordPolicy = &krm.SQLUserPasswordPolicy{
					AllowedFailedAttempts:    tc.policyAllowed,
					EnableFailedAttemptsCheck: tc.policyEnableCheck,
				}
			}

			got, err := SQLUserKRMToGCP(obj, tc.password, nil)
			if err != nil {
				t.Fatal(err)
			}

			if !stringSliceEqual(got.ForceSendFields, tc.wantForceSendUser) {
				t.Errorf("User.ForceSendFields = %v, want %v", got.ForceSendFields, tc.wantForceSendUser)
			}
			if got.PasswordPolicy != nil {
				if !stringSliceEqual(got.PasswordPolicy.ForceSendFields, tc.wantForceSendPolicy) {
					t.Errorf("PasswordPolicy.ForceSendFields = %v, want %v", got.PasswordPolicy.ForceSendFields, tc.wantForceSendPolicy)
				}
			}
		})
	}
}

func TestSQLUserKRMToGCP_NullFields(t *testing.T) {
	t.Run("clearing PasswordPolicy adds NullFields", func(t *testing.T) {
		obj := &krm.SQLUser{}
		obj.Spec.ResourceID = direct.PtrTo("testuser")
		// No PasswordPolicy in desired, but actual has one.
		actual := &api.User{
			PasswordPolicy: &api.UserPasswordValidationPolicy{
				AllowedFailedAttempts: 3,
			},
		}

		got, err := SQLUserKRMToGCP(obj, "", actual)
		if err != nil {
			t.Fatal(err)
		}

		if !containsString(got.NullFields, "PasswordPolicy") {
			t.Errorf("NullFields = %v, want to contain 'PasswordPolicy'", got.NullFields)
		}
	})

	t.Run("clearing PasswordExpirationDuration adds NullFields", func(t *testing.T) {
		obj := &krm.SQLUser{}
		obj.Spec.ResourceID = direct.PtrTo("testuser")
		obj.Spec.PasswordPolicy = &krm.SQLUserPasswordPolicy{
			EnableFailedAttemptsCheck: direct.PtrTo(true),
			// PasswordExpirationDuration deliberately nil.
		}
		actual := &api.User{
			PasswordPolicy: &api.UserPasswordValidationPolicy{
				PasswordExpirationDuration: "3600s",
			},
		}

		got, err := SQLUserKRMToGCP(obj, "", actual)
		if err != nil {
			t.Fatal(err)
		}

		if got.PasswordPolicy == nil {
			t.Fatal("PasswordPolicy should not be nil")
		}
		if !containsString(got.PasswordPolicy.NullFields, "PasswordExpirationDuration") {
			t.Errorf("PasswordPolicy.NullFields = %v, want to contain 'PasswordExpirationDuration'", got.PasswordPolicy.NullFields)
		}
	})
}

func TestUserHasDiff(t *testing.T) {
	tests := []struct {
		name     string
		desired  *api.User
		actual   *api.User
		wantDiff bool
	}{
		{
			name:     "identical users no diff",
			desired:  &api.User{Type: "BUILT_IN"},
			actual:   &api.User{Type: "BUILT_IN"},
			wantDiff: false,
		},
		{
			name:     "password set always diffs",
			desired:  &api.User{Type: "BUILT_IN", Password: "secret"},
			actual:   &api.User{Type: "BUILT_IN"},
			wantDiff: true,
		},
		{
			name:     "type changed",
			desired:  &api.User{Type: "CLOUD_IAM_USER"},
			actual:   &api.User{Type: "BUILT_IN"},
			wantDiff: true,
		},
		{
			name: "policy added",
			desired: &api.User{
				Type: "BUILT_IN",
				PasswordPolicy: &api.UserPasswordValidationPolicy{
					EnableFailedAttemptsCheck: true,
				},
			},
			actual:   &api.User{Type: "BUILT_IN"},
			wantDiff: true,
		},
		{
			name:    "policy removed via NullFields",
			desired: &api.User{Type: "BUILT_IN", NullFields: []string{"PasswordPolicy"}},
			actual: &api.User{
				Type: "BUILT_IN",
				PasswordPolicy: &api.UserPasswordValidationPolicy{
					EnableFailedAttemptsCheck: true,
				},
			},
			wantDiff: true,
		},
		{
			name: "same policy no diff",
			desired: &api.User{
				Type: "BUILT_IN",
				PasswordPolicy: &api.UserPasswordValidationPolicy{
					AllowedFailedAttempts:     3,
					EnableFailedAttemptsCheck: true,
				},
			},
			actual: &api.User{
				Type: "BUILT_IN",
				PasswordPolicy: &api.UserPasswordValidationPolicy{
					AllowedFailedAttempts:     3,
					EnableFailedAttemptsCheck: true,
					// Status is output-only, should be ignored.
					Status: &api.PasswordStatus{
						Locked: true,
					},
				},
			},
			wantDiff: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			adapter := &sqlUserAdapter{actual: tc.actual}
			got := adapter.userHasDiff(tc.desired)
			if got != tc.wantDiff {
				t.Errorf("userHasDiff() = %v, want %v", got, tc.wantDiff)
			}
		})
	}
}

func stringSliceEqual(a, b []string) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func containsString(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
