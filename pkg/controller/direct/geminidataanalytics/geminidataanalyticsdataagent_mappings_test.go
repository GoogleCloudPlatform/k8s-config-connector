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

package geminidataanalytics

import (
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/geminidataanalytics/v1alpha1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestGeminiDataAnalyticsDataAgentSpec_ToProto_KMSKey(t *testing.T) {
	tests := []struct {
		name       string
		kmsKeyRef  *kmsv1beta1.KMSCryptoKeyRef
		wantKmsKey *string
	}{
		{
			name:       "nil KMSKeyRef leaves out.KmsKey nil",
			kmsKeyRef:  nil,
			wantKmsKey: nil,
		},
		{
			name:       "empty External leaves out.KmsKey nil",
			kmsKeyRef:  &kmsv1beta1.KMSCryptoKeyRef{External: ""},
			wantKmsKey: nil,
		},
		{
			name:       "non-empty External sets out.KmsKey",
			kmsKeyRef:  &kmsv1beta1.KMSCryptoKeyRef{External: "projects/test-project/locations/us-central1/keyRings/ring/cryptoKeys/key"},
			wantKmsKey: direct.LazyPtr("projects/test-project/locations/us-central1/keyRings/ring/cryptoKeys/key"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			in := &krm.GeminiDataAnalyticsDataAgentSpec{
				KMSKeyRef: tc.kmsKeyRef,
			}
			out := GeminiDataAnalyticsDataAgentSpec_ToProto(mapCtx, in)
			if tc.wantKmsKey == nil {
				if out.KmsKey != nil {
					t.Errorf("got out.KmsKey = %q, want nil", *out.KmsKey)
				}
			} else {
				if out.KmsKey == nil {
					t.Fatalf("got out.KmsKey = nil, want %q", *tc.wantKmsKey)
				}
				if *out.KmsKey != *tc.wantKmsKey {
					t.Errorf("got out.KmsKey = %q, want %q", *out.KmsKey, *tc.wantKmsKey)
				}
			}
		})
	}
}
