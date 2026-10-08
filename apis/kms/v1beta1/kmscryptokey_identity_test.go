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
	"errors"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestKMSCryptoKeyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *KMSCryptoKeyIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/keyRings/my-keyring/cryptoKeys/my-key",
			want: &KMSCryptoKeyIdentity{
				Project:   "my-project",
				Location:  "us-central1",
				KeyRing:   "my-keyring",
				CryptoKey: "my-key",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://cloudkms.googleapis.com/projects/my-project/locations/us-central1/keyRings/my-keyring/cryptoKeys/my-key",
			want: &KMSCryptoKeyIdentity{
				Project:   "my-project",
				Location:  "us-central1",
				KeyRing:   "my-keyring",
				CryptoKey: "my-key",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &KMSCryptoKeyIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestKMSCryptoKeyRef_Normalize(t *testing.T) {
	ctx := context.Background()

	readyKey := &unstructured.Unstructured{}
	readyKey.SetGroupVersionKind(KMSCryptoKeyGVK)
	readyKey.SetName("ready-key")
	readyKey.SetNamespace("test-ns")
	if err := unstructured.SetNestedField(readyKey.Object, "projects/my-project/locations/us-central1/keyRings/my-keyring/cryptoKeys/ready-key", "status", "selfLink"); err != nil {
		t.Fatalf("failed to set status.selfLink: %v", err)
	}

	unreadyKey := &unstructured.Unstructured{}
	unreadyKey.SetGroupVersionKind(KMSCryptoKeyGVK)
	unreadyKey.SetName("unready-key")
	unreadyKey.SetNamespace("test-ns")
	if err := unstructured.SetNestedField(unreadyKey.Object, "unready-key", "spec", "resourceID"); err != nil {
		t.Fatalf("failed to set spec.resourceID: %v", err)
	}

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readyKey, unreadyKey).Build()

	t.Run("resolves status.selfLink when ready", func(t *testing.T) {
		ref := &KMSCryptoKeyRef{Name: "ready-key"}
		if err := ref.Normalize(ctx, reader, "test-ns"); err != nil {
			t.Fatalf("Normalize() unexpected error: %v", err)
		}
		want := "projects/my-project/locations/us-central1/keyRings/my-keyring/cryptoKeys/ready-key"
		if ref.External != want {
			t.Errorf("Normalize() External = %q, want %q", ref.External, want)
		}
	})

	t.Run("returns ReferenceNotReadyError when status.selfLink is empty", func(t *testing.T) {
		ref := &KMSCryptoKeyRef{Name: "unready-key"}
		err := ref.Normalize(ctx, reader, "test-ns")
		if err == nil {
			t.Fatalf("Normalize() expected ReferenceNotReadyError, got nil (External=%q)", ref.External)
		}
		var notReadyErr *k8s.ReferenceNotReadyError
		if !errors.As(err, &notReadyErr) {
			t.Errorf("Normalize() error = %T (%v), want *k8s.ReferenceNotReadyError", err, err)
		}
	})
}
