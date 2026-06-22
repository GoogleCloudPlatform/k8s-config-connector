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
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestKMSSecretCiphertextIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *KMSSecretCiphertextIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/keyRings/my-keyring/cryptoKeys/my-key/ChQAt3M",
			want: &KMSSecretCiphertextIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				KeyRing:    "my-keyring",
				CryptoKey:  "my-key",
				Ciphertext: "ChQAt3M",
			},
		},
		{
			name: "valid reference with slashes in base64 ciphertext",
			ref:  "projects/my-project/locations/us-central1/keyRings/my-keyring/cryptoKeys/my-key/ChQAt/3M+abc/xyz",
			want: &KMSSecretCiphertextIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				KeyRing:    "my-keyring",
				CryptoKey:  "my-key",
				Ciphertext: "ChQAt/3M+abc/xyz",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://cloudkms.googleapis.com/projects/my-project/locations/us-central1/keyRings/my-keyring/cryptoKeys/my-key/ChQAt3M",
			want: &KMSSecretCiphertextIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				KeyRing:    "my-keyring",
				CryptoKey:  "my-key",
				Ciphertext: "ChQAt3M",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &KMSSecretCiphertextIdentity{}
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
