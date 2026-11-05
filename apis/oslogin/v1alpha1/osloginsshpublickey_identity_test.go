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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestOSLoginSSHPublicKeyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name                  string
		ref                   string
		wantErr               bool
		want                  *OSLoginSSHPublicKeyIdentity
		wantIdentitySpecified bool
	}{
		{
			name: "valid reference",
			ref:  "users/test-user@example.com/sshPublicKeys/1234567890abcdef",
			want: &OSLoginSSHPublicKeyIdentity{
				User:        "test-user@example.com",
				Fingerprint: "1234567890abcdef",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "full url with scheme",
			ref:  "https://oslogin.googleapis.com/users/test-user@example.com/sshPublicKeys/1234567890abcdef",
			want: &OSLoginSSHPublicKeyIdentity{
				User:        "test-user@example.com",
				Fingerprint: "1234567890abcdef",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "double slash url",
			ref:  "//oslogin.googleapis.com/users/test-user@example.com/sshPublicKeys/1234567890abcdef",
			want: &OSLoginSSHPublicKeyIdentity{
				User:        "test-user@example.com",
				Fingerprint: "1234567890abcdef",
			},
			wantIdentitySpecified: true,
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "missing fingerprint",
			ref:     "users/test-user@example.com/sshPublicKeys/",
			wantErr: true,
		},
		{
			name:    "missing user segment",
			ref:     "users//sshPublicKeys/1234567890abcdef",
			wantErr: true,
		},
		{
			name:    "unrelated gcp url",
			ref:     "projects/p1/locations/l1/connections/c1",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &OSLoginSSHPublicKeyIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.HasIdentitySpecified(); got != tt.wantIdentitySpecified {
					t.Errorf("HasIdentitySpecified() = %v, want %v", got, tt.wantIdentitySpecified)
				}
			}
		})
	}
}

func TestOSLoginSSHPublicKeyIdentity_Methods(t *testing.T) {
	identity := &OSLoginSSHPublicKeyIdentity{
		User:        "test-user@example.com",
		Fingerprint: "1234567890abcdef",
	}

	wantString := "users/test-user@example.com/sshPublicKeys/1234567890abcdef"
	if diff := cmp.Diff(wantString, identity.String()); diff != "" {
		t.Errorf("String() mismatch (-want +got):\n%s", diff)
	}

	wantParent := "users/test-user@example.com"
	if diff := cmp.Diff(wantParent, identity.ParentString()); diff != "" {
		t.Errorf("ParentString() mismatch (-want +got):\n%s", diff)
	}

	wantHost := "oslogin.googleapis.com"
	if diff := cmp.Diff(wantHost, identity.Host()); diff != "" {
		t.Errorf("Host() mismatch (-want +got):\n%s", diff)
	}
}

func TestOSLoginSSHPublicKey_GetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name                  string
		obj                   *OSLoginSSHPublicKey
		wantErr               bool
		want                  *OSLoginSSHPublicKeyIdentity
		wantIdentitySpecified bool
	}{
		{
			name: "GetIdentity with specified resourceID",
			obj: &OSLoginSSHPublicKey{
				Spec: OSLoginSSHPublicKeySpec{
					User:       "test-user@example.com",
					ResourceID: common.LazyPtr("1234567890abcdef"),
				},
			},
			want: &OSLoginSSHPublicKeyIdentity{
				User:        "test-user@example.com",
				Fingerprint: "1234567890abcdef",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with full external format in resourceID",
			obj: &OSLoginSSHPublicKey{
				Spec: OSLoginSSHPublicKeySpec{
					User:       "test-user@example.com",
					ResourceID: common.LazyPtr("users/test-user@example.com/sshPublicKeys/1234567890abcdef"),
				},
			},
			want: &OSLoginSSHPublicKeyIdentity{
				User:        "test-user@example.com",
				Fingerprint: "1234567890abcdef",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with full external resourceID but mismatched user",
			obj: &OSLoginSSHPublicKey{
				Spec: OSLoginSSHPublicKeySpec{
					User:       "user1@example.com",
					ResourceID: common.LazyPtr("users/user2@example.com/sshPublicKeys/1234567890abcdef"),
				},
			},
			wantErr: true,
		},
		{
			name: "GetIdentity with server-generated identity (empty spec resourceID, defaulted from status)",
			obj: &OSLoginSSHPublicKey{
				Spec: OSLoginSSHPublicKeySpec{
					User: "test-user@example.com",
				},
				Status: OSLoginSSHPublicKeyStatus{
					Fingerprint: common.LazyPtr("1234567890abcdef"),
				},
			},
			want: &OSLoginSSHPublicKeyIdentity{
				User:        "test-user@example.com",
				Fingerprint: "1234567890abcdef",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with no spec resourceID and empty status (not yet created)",
			obj: &OSLoginSSHPublicKey{
				Spec: OSLoginSSHPublicKeySpec{
					User: "test-user@example.com",
				},
			},
			want: &OSLoginSSHPublicKeyIdentity{
				User:        "test-user@example.com",
				Fingerprint: "",
			},
			wantIdentitySpecified: false,
		},
		{
			name: "GetIdentity with server-generated identity conflict with spec resourceID",
			obj: &OSLoginSSHPublicKey{
				Spec: OSLoginSSHPublicKeySpec{
					User:       "test-user@example.com",
					ResourceID: common.LazyPtr("original-fingerprint"),
				},
				Status: OSLoginSSHPublicKeyStatus{
					Fingerprint: common.LazyPtr("different-fingerprint"),
				},
			},
			wantErr: true,
		},
		{
			name: "GetIdentity with empty user in spec",
			obj: &OSLoginSSHPublicKey{
				Spec: OSLoginSSHPublicKeySpec{
					User: "",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIdentity, err := tt.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				got, ok := gotIdentity.(*OSLoginSSHPublicKeyIdentity)
				if !ok {
					t.Fatalf("returned identity is not *OSLoginSSHPublicKeyIdentity, got %T", gotIdentity)
				}
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
				if gotSpecified := got.HasIdentitySpecified(); gotSpecified != tt.wantIdentitySpecified {
					t.Errorf("got.HasIdentitySpecified() = %v, want %v", gotSpecified, tt.wantIdentitySpecified)
				}
			}
		})
	}
}

func TestOSLoginSSHPublicKeyRef(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)

	readyObj := &OSLoginSSHPublicKey{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-key",
			Namespace: "test-ns",
		},
		Spec: OSLoginSSHPublicKeySpec{
			User: "test-user@example.com",
		},
		Status: OSLoginSSHPublicKeyStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
			Fingerprint: common.LazyPtr("1234567890abcdef"),
		},
	}

	notReadyObj := &OSLoginSSHPublicKey{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "not-ready-key",
			Namespace: "test-ns",
		},
		Spec: OSLoginSSHPublicKeySpec{
			User: "test-user@example.com",
		},
		Status: OSLoginSSHPublicKeyStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "False",
				},
			},
		},
	}

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readyObj, notReadyObj).Build()

	t.Run("ValidateExternal and ParseExternalToIdentity", func(t *testing.T) {
		ref := &OSLoginSSHPublicKeyRef{
			External: "users/test-user@example.com/sshPublicKeys/1234567890abcdef",
		}
		if err := ref.ValidateExternal(ref.External); err != nil {
			t.Errorf("ValidateExternal() unexpected error: %v", err)
		}
		id, err := ref.ParseExternalToIdentity()
		if err != nil {
			t.Fatalf("ParseExternalToIdentity() unexpected error: %v", err)
		}
		want := &OSLoginSSHPublicKeyIdentity{
			User:        "test-user@example.com",
			Fingerprint: "1234567890abcdef",
		}
		if diff := cmp.Diff(want, id); diff != "" {
			t.Errorf("ParseExternalToIdentity() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("ValidateExternal with invalid ref", func(t *testing.T) {
		ref := &OSLoginSSHPublicKeyRef{
			External: "invalid/format",
		}
		if err := ref.ValidateExternal(ref.External); err == nil {
			t.Errorf("ValidateExternal() expected error, got nil")
		}
	})

	t.Run("Normalize with ready object", func(t *testing.T) {
		ref := &OSLoginSSHPublicKeyRef{
			Name:      "test-key",
			Namespace: "test-ns",
		}
		if err := ref.Normalize(ctx, fakeClient, "test-ns"); err != nil {
			t.Fatalf("Normalize() unexpected error: %v", err)
		}
		wantExternal := "users/test-user@example.com/sshPublicKeys/1234567890abcdef"
		if ref.External != wantExternal {
			t.Errorf("Normalize() external = %q, want %q", ref.External, wantExternal)
		}
	})

	t.Run("Normalize with not ready object", func(t *testing.T) {
		ref := &OSLoginSSHPublicKeyRef{
			Name:      "not-ready-key",
			Namespace: "test-ns",
		}
		if err := ref.Normalize(ctx, fakeClient, "test-ns"); err == nil {
			t.Errorf("Normalize() expected error for not-ready object, got nil")
		}
	})
}
