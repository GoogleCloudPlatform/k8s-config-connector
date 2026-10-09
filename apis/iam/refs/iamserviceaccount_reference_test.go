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

package iamrefs

import (
	"context"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestIAMServiceAccountRef_Normalize(t *testing.T) {
	ctx := context.Background()

	readySA := &unstructured.Unstructured{}
	readySA.SetGroupVersionKind(IAMServiceAccountGVK)
	readySA.SetName("my-sa")
	readySA.SetNamespace("default")
	if err := unstructured.SetNestedField(readySA.Object, "my-sa@test-project.iam.gserviceaccount.com", "status", "email"); err != nil {
		t.Fatalf("failed to set status.email: %v", err)
	}

	notReadySA := &unstructured.Unstructured{}
	notReadySA.SetGroupVersionKind(IAMServiceAccountGVK)
	notReadySA.SetName("not-ready-sa")
	notReadySA.SetNamespace("default")

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readySA, notReadySA).Build()

	tests := []struct {
		name             string
		ref              *IAMServiceAccountRef
		defaultNamespace string
		wantErr          bool
		wantExternal     string
	}{
		{
			name: "external reference valid",
			ref: &IAMServiceAccountRef{
				External: "custom-sa@custom-project.iam.gserviceaccount.com",
			},
			defaultNamespace: "default",
			wantErr:          false,
			wantExternal:     "custom-sa@custom-project.iam.gserviceaccount.com",
		},
		{
			name: "external reference invalid format",
			ref: &IAMServiceAccountRef{
				External: "projects/test-project/serviceAccounts/my-sa",
			},
			defaultNamespace: "default",
			wantErr:          true,
		},
		{
			name: "resolve by name ready",
			ref: &IAMServiceAccountRef{
				Name: "my-sa",
			},
			defaultNamespace: "default",
			wantErr:          false,
			wantExternal:     "my-sa@test-project.iam.gserviceaccount.com",
		},
		{
			name: "resolve by name not found",
			ref: &IAMServiceAccountRef{
				Name: "missing-sa",
			},
			defaultNamespace: "default",
			wantErr:          true,
		},
		{
			name: "resolve by name not ready (no status.email)",
			ref: &IAMServiceAccountRef{
				Name: "not-ready-sa",
			},
			defaultNamespace: "default",
			wantErr:          true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.ref.Normalize(ctx, reader, tc.defaultNamespace)
			if (err != nil) != tc.wantErr {
				t.Fatalf("Normalize() error = %v, wantErr %v", err, tc.wantErr)
			}
			if !tc.wantErr && tc.ref.External != tc.wantExternal {
				t.Errorf("Normalize() External = %q, want %q", tc.ref.External, tc.wantExternal)
			}
		})
	}
}

func TestParseIAMServiceAccountEmail(t *testing.T) {
	tests := []struct {
		email       string
		wantAccount string
		wantProject string
		wantErr     bool
	}{
		{
			email:       "test-sa@my-project.iam.gserviceaccount.com",
			wantAccount: "test-sa",
			wantProject: "my-project",
			wantErr:     false,
		},
		{
			email:   "invalid-email",
			wantErr: true,
		},
		{
			email:   "test-sa@my-project.example.com",
			wantErr: true,
		},
		{
			email:   "@my-project.iam.gserviceaccount.com",
			wantErr: true,
		},
		{
			email:   "test-sa@.iam.gserviceaccount.com",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		account, project, err := ParseIAMServiceAccountEmail(tc.email)
		if (err != nil) != tc.wantErr {
			t.Errorf("ParseIAMServiceAccountEmail(%q) error = %v, wantErr %v", tc.email, err, tc.wantErr)
			continue
		}
		if !tc.wantErr {
			if account != tc.wantAccount || project != tc.wantProject {
				t.Errorf("ParseIAMServiceAccountEmail(%q) = (%q, %q), want (%q, %q)", tc.email, account, project, tc.wantAccount, tc.wantProject)
			}
		}
	}
}
