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

package v1beta1

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestIAMServiceAccountIdentityParse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      *IAMServiceAccountIdentity
		wantError bool
	}{
		{
			name:      "Parse with email (invalid)",
			input:     "projects/myProject/serviceAccounts/my-sa@myProject.iam.gserviceaccount.com",
			want:      nil,
			wantError: true,
		},
		{
			name:  "Normal parse with account ID",
			input: "projects/myProject/serviceAccounts/my-sa",
			want: &IAMServiceAccountIdentity{
				Project: "myProject",
				Account: "my-sa",
			},
			wantError: false,
		},
		{
			name:      "Parse with leading slash and email (invalid)",
			input:     "/projects/p1/serviceAccounts/sa@p1.iam.gserviceaccount.com",
			want:      nil,
			wantError: true,
		},
		{
			name:  "Parse with leading slash and account ID",
			input: "/projects/p1/serviceAccounts/sa",
			want: &IAMServiceAccountIdentity{
				Project: "p1",
				Account: "sa",
			},
			wantError: false,
		},
		{
			name:      "Parse with domain and email (invalid)",
			input:     "//iam.googleapis.com/projects/first/serviceAccounts/second@first.iam.gserviceaccount.com",
			want:      nil,
			wantError: true,
		},
		{
			name:  "Parse with domain and account ID",
			input: "//iam.googleapis.com/projects/first/serviceAccounts/second",
			want: &IAMServiceAccountIdentity{
				Project: "first",
				Account: "second",
			},
			wantError: false,
		},
		{
			name:      "Empty string",
			input:     "",
			want:      nil,
			wantError: true,
		},
		{
			name:      "Wrong format - missing serviceAccounts prefix",
			input:     "projects/myProject/my-sa@myProject.iam.gserviceaccount.com",
			want:      nil,
			wantError: true,
		},
		{
			name:      "Wrong format - wrong project key",
			input:     "orgs/myProject/serviceAccounts/my-sa",
			want:      nil,
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := &IAMServiceAccountIdentity{}
			err := got.FromExternal(tc.input)
			if tc.wantError {
				if err == nil {
					t.Errorf("FromExternal(%q) expected error but got none", tc.input)
				}
				return
			}
			if err != nil {
				t.Errorf("FromExternal(%q) unexpected error: %v", tc.input, err)
				return
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("FromExternal(%q) mismatch (-want +got):\n%s", tc.input, diff)
			}
		})
	}
}

func TestIAMServiceAccount_GetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	if err := SchemeBuilder.AddToScheme(scheme); err != nil {
		t.Fatalf("failed to add to scheme: %v", err)
	}
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name    string
		obj     *IAMServiceAccount
		want    *IAMServiceAccountIdentity
		wantErr bool
	}{
		{
			name: "Normal GetIdentity when Status.ExternalRef is empty",
			obj: &IAMServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-sa",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: IAMServiceAccountSpec{
					ResourceID: common.LazyPtr("my-sa"),
				},
			},
			want: &IAMServiceAccountIdentity{
				Project: "my-project",
				Account: "my-sa",
			},
			wantErr: false,
		},
		{
			name: "Normal GetIdentity when Status.ExternalRef matches the spec identity (non-email format)",
			obj: &IAMServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-sa",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: IAMServiceAccountSpec{
					ResourceID: common.LazyPtr("my-sa"),
				},
				Status: IAMServiceAccountStatus{
					ExternalRef: common.LazyPtr("projects/my-project/serviceAccounts/my-sa"),
				},
			},
			want: &IAMServiceAccountIdentity{
				Project: "my-project",
				Account: "my-sa",
			},
			wantErr: false,
		},
		{
			name: "Error when Status.ExternalRef has email suffix",
			obj: &IAMServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-sa",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: IAMServiceAccountSpec{
					ResourceID: common.LazyPtr("my-sa"),
				},
				Status: IAMServiceAccountStatus{
					ExternalRef: common.LazyPtr("projects/my-project/serviceAccounts/my-sa@my-project.iam.gserviceaccount.com"),
				},
			},
			wantErr: true,
		},
		{
			name: "Error when Status.ExternalRef is a mismatched account name",
			obj: &IAMServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-sa",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: IAMServiceAccountSpec{
					ResourceID: common.LazyPtr("my-sa"),
				},
				Status: IAMServiceAccountStatus{
					ExternalRef: common.LazyPtr("projects/my-project/serviceAccounts/mismatched-sa"),
				},
			},
			wantErr: true,
		},
		{
			name: "Error when Status.ExternalRef is a full path but mismatches the project name",
			obj: &IAMServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-sa",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: IAMServiceAccountSpec{
					ResourceID: common.LazyPtr("my-sa"),
				},
				Status: IAMServiceAccountStatus{
					ExternalRef: common.LazyPtr("projects/another-project/serviceAccounts/my-sa"),
				},
			},
			wantErr: true,
		},
		{
			name: "Error when Status.ExternalRef is in an unexpected format that fails to parse",
			obj: &IAMServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-sa",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: IAMServiceAccountSpec{
					ResourceID: common.LazyPtr("my-sa"),
				},
				Status: IAMServiceAccountStatus{
					ExternalRef: common.LazyPtr("invalid-format"),
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotIdentity, err := tc.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetIdentity() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !tc.wantErr {
				got, ok := gotIdentity.(*IAMServiceAccountIdentity)
				if !ok {
					t.Fatalf("returned identity is not *IAMServiceAccountIdentity, got %T", gotIdentity)
				}
				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestIAMServiceAccountRef(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	if err := SchemeBuilder.AddToScheme(scheme); err != nil {
		t.Fatalf("failed to add to scheme: %v", err)
	}

	t.Run("ValidateExternal and ParseExternalToIdentity", func(t *testing.T) {
		tests := []struct {
			name     string
			external string
			wantProj string
			wantAcc  string
			wantErr  bool
		}{
			{
				name:     "Valid email address",
				external: "my-sa@my-project.iam.gserviceaccount.com",
				wantProj: "my-project",
				wantAcc:  "my-sa",
				wantErr:  false,
			},
			{
				name:     "Invalid format: missing suffix",
				external: "my-sa@my-project",
				wantErr:  true,
			},
			{
				name:     "Invalid format: missing account",
				external: "@my-project.iam.gserviceaccount.com",
				wantErr:  true,
			},
			{
				name:     "Invalid format: projects/ prefix",
				external: "projects/my-project/serviceAccounts/my-sa",
				wantErr:  true,
			},
			{
				name:     "Empty",
				external: "",
				wantErr:  true,
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				ref := &IAMServiceAccountRef{External: tc.external}
				err := ref.ValidateExternal(tc.external)
				if (err != nil) != tc.wantErr {
					t.Fatalf("ValidateExternal() error = %v, wantErr %v", err, tc.wantErr)
				}

				id, err := ref.ParseExternalToIdentity()
				if (err != nil) != tc.wantErr {
					t.Fatalf("ParseExternalToIdentity() error = %v, wantErr %v", err, tc.wantErr)
				}

				if !tc.wantErr {
					saId, ok := id.(*IAMServiceAccountIdentity)
					if !ok {
						t.Fatalf("expected *IAMServiceAccountIdentity, got %T", id)
					}
					if saId.Project != tc.wantProj || saId.Account != tc.wantAcc {
						t.Errorf("expected Proj=%q, Acc=%q; got Proj=%q, Acc=%q", tc.wantProj, tc.wantAcc, saId.Project, saId.Account)
					}
				}
			})
		}
	})

	t.Run("Normalize from local reference", func(t *testing.T) {
		sa := &IAMServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "local-sa",
				Namespace: "my-ns",
			},
			Status: IAMServiceAccountStatus{
				Email: common.LazyPtr("local-sa@my-ns.iam.gserviceaccount.com"),
			},
		}
		fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(sa).Build()

		ref := &IAMServiceAccountRef{
			Name:      "local-sa",
			Namespace: "my-ns",
		}

		if err := ref.Normalize(ctx, fakeClient, "my-ns"); err != nil {
			t.Fatalf("Normalize() unexpected error: %v", err)
		}

		wantEmail := "local-sa@my-ns.iam.gserviceaccount.com"
		if ref.External != wantEmail {
			t.Errorf("expected External=%q, got %q", wantEmail, ref.External)
		}
	})

	t.Run("Normalize local reference - not ready", func(t *testing.T) {
		sa := &IAMServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "local-sa",
				Namespace: "my-ns",
			},
			// Email is empty, so not ready
		}
		fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(sa).Build()

		ref := &IAMServiceAccountRef{
			Name:      "local-sa",
			Namespace: "my-ns",
		}

		err := ref.Normalize(ctx, fakeClient, "my-ns")
		if err == nil {
			t.Fatalf("Normalize() expected error for unready resource, got none")
		}
	})
}
