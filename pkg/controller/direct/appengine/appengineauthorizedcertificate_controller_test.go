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

package appengine

import (
	"context"
	"testing"

	appenginepb "cloud.google.com/go/appengine/apiv1/appenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestAuthorizedCertificateAdapter_Export(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		id          *krm.AppEngineAuthorizedCertificateIdentity
		actual      *appenginepb.AuthorizedCertificate
		wantName    string
		wantProject string
		wantDisplay string
	}{
		{
			name: "basic certificate export",
			id: &krm.AppEngineAuthorizedCertificateIdentity{
				App:                   "test-project",
				AuthorizedCertificate: "12345",
			},
			actual: &appenginepb.AuthorizedCertificate{
				Name:        "apps/test-project/authorizedCertificates/12345",
				DisplayName: "Test Certificate",
			},
			wantName:    "12345",
			wantProject: "test-project",
			wantDisplay: "Test Certificate",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			adapter := &AuthorizedCertificateAdapter{
				id:     tc.id,
				actual: tc.actual,
			}

			u, err := adapter.Export(ctx)
			if err != nil {
				t.Fatalf("Export() returned unexpected error: %v", err)
			}

			if got := u.GetName(); got != tc.wantName {
				t.Errorf("u.GetName() = %q, want %q", got, tc.wantName)
			}

			if got := u.GetKind(); got != "AppEngineAuthorizedCertificate" {
				t.Errorf("u.GetKind() = %q, want AppEngineAuthorizedCertificate", got)
			}

			projectRefExternal, found, err := unstructured.NestedString(u.Object, "spec", "projectRef", "external")
			if err != nil || !found {
				t.Fatalf("spec.projectRef.external not found: %v", err)
			}
			if projectRefExternal != tc.wantProject {
				t.Errorf("spec.projectRef.external = %q, want %q", projectRefExternal, tc.wantProject)
			}

			displayName, found, err := unstructured.NestedString(u.Object, "spec", "displayName")
			if err != nil || !found {
				t.Fatalf("spec.displayName not found: %v", err)
			}
			if displayName != tc.wantDisplay {
				t.Errorf("spec.displayName = %q, want %q", displayName, tc.wantDisplay)
			}
		})
	}
}
