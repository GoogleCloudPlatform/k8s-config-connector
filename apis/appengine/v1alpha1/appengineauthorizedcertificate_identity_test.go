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

package v1alpha1

import (
	"context"
	"testing"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
)

func TestAppEngineAuthorizedCertificateIdentity_SpecProjectRef(t *testing.T) {
	ctx := context.Background()
	obj := &AppEngineAuthorizedCertificate{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-cert",
		},
		Spec: AppEngineAuthorizedCertificateSpec{
			ProjectRef: &refsv1beta1.ProjectRef{
				External: "project-from-spec",
			},
			ResourceID: ptr.To("12345"),
		},
	}

	id, err := getIdentityFromAppEngineAuthorizedCertificateSpec(ctx, nil, obj)
	if err != nil {
		t.Fatalf("unexpected error resolving identity: %v", err)
	}
	if id.App != "project-from-spec" {
		t.Errorf("got app %q, want %q", id.App, "project-from-spec")
	}
	if id.AuthorizedCertificate != "12345" {
		t.Errorf("got authorizedCertificate %q, want %q", id.AuthorizedCertificate, "12345")
	}
}

func TestAppEngineAuthorizedCertificateIdentity_ResolveProjectIDFallback(t *testing.T) {
	ctx := context.Background()
	obj := &AppEngineAuthorizedCertificate{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-cert",
			Annotations: map[string]string{
				"cnrm.cloud.google.com/project-id": "project-from-annotation",
			},
		},
	}

	id, err := getIdentityFromAppEngineAuthorizedCertificateSpec(ctx, nil, obj)
	if err != nil {
		t.Fatalf("unexpected error resolving identity: %v", err)
	}
	if id.App != "project-from-annotation" {
		t.Errorf("got app %q, want %q", id.App, "project-from-annotation")
	}
	// For server-generated ID without spec.resourceID, AuthorizedCertificate should be empty string
	if id.AuthorizedCertificate != "" {
		t.Errorf("got authorizedCertificate %q, want empty string", id.AuthorizedCertificate)
	}
}

func TestAppEngineAuthorizedCertificateIdentity_ErrorWhenUnset(t *testing.T) {
	ctx := context.Background()
	obj := &AppEngineAuthorizedCertificate{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-cert",
		},
	}

	_, err := getIdentityFromAppEngineAuthorizedCertificateSpec(ctx, nil, obj)
	if err == nil {
		t.Fatalf("expected error resolving identity when project is unset, got nil")
	}
}
