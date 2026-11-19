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

package v1beta1_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestKMSCryptoKeyRef_Normalize(t *testing.T) {
	tests := []struct {
		name         string
		ref          v1beta1.KMSCryptoKeyRef
		objs         []client.Object
		wantExternal string
		wantNotReady bool
		wantNotFound bool
	}{
		{
			name: "canonicalizes_host_prefixed_external",
			ref: v1beta1.KMSCryptoKeyRef{
				External: "//cloudkms.googleapis.com/projects/my-proj/locations/us-west1/keyRings/kr/cryptoKeys/my-key",
			},
			wantExternal: "projects/my-proj/locations/us-west1/keyRings/kr/cryptoKeys/my-key",
		},
		{
			name: "canonicalizes_https_external",
			ref: v1beta1.KMSCryptoKeyRef{
				External: "https://cloudkms.googleapis.com/projects/my-proj/locations/us-west1/keyRings/kr/cryptoKeys/my-key",
			},
			wantExternal: "projects/my-proj/locations/us-west1/keyRings/kr/cryptoKeys/my-key",
		},
		{
			name: "resolves_and_canonicalizes_status_self_link",
			ref: v1beta1.KMSCryptoKeyRef{
				Name: "my-key",
			},
			objs: []client.Object{
				newUnstructuredKMSKey("my-key", "test-ns", "//cloudkms.googleapis.com/projects/my-proj/locations/us/keyRings/kr/cryptoKeys/my-key"),
			},
			wantExternal: "projects/my-proj/locations/us/keyRings/kr/cryptoKeys/my-key",
		},
		{
			name:         "missing_key_returns_not_found",
			ref:          v1beta1.KMSCryptoKeyRef{Name: "missing-key"},
			wantNotFound: true,
		},
		{
			name: "key_without_self_link_returns_not_ready",
			ref:  v1beta1.KMSCryptoKeyRef{Name: "unready-key"},
			objs: []client.Object{
				newUnstructuredKMSKey("unready-key", "test-ns", ""),
			},
			wantNotReady: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reader := fake.NewClientBuilder().WithScheme(runtime.NewScheme()).WithObjects(tc.objs...).Build()
			r := tc.ref

			err := r.Normalize(t.Context(), reader, "test-ns")

			if tc.wantNotReady || tc.wantNotFound {
				if err == nil {
					t.Fatalf("Normalize(%+v) = nil, want error", tc.ref)
				}
				if tc.wantNotReady && !k8s.IsReferenceNotReadyError(err) {
					t.Errorf("Normalize(%+v) error = %v, want ReferenceNotReadyError", tc.ref, err)
				}
				if tc.wantNotFound && !k8s.IsReferenceNotFoundError(err) {
					t.Errorf("Normalize(%+v) error = %v, want ReferenceNotFoundError", tc.ref, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("Normalize(%+v) unexpected error: %v", tc.ref, err)
			}
			if r.External != tc.wantExternal {
				t.Errorf("Normalize(%+v) External = %q, want %q", tc.ref, r.External, tc.wantExternal)
			}
		})
	}
}

func newUnstructuredKMSKey(name, namespace, selfLink string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(v1beta1.KMSCryptoKeyGVK)
	u.SetName(name)
	u.SetNamespace(namespace)
	if selfLink != "" {
		u.Object["status"] = map[string]any{"selfLink": selfLink}
	}
	return u
}
