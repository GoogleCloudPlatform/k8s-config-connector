// Copyright 2025 Google LLC
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

import "testing"

func TestValidateOneOf(t *testing.T) {
	testCases := []struct {
		name          string
		input         *KMSKeyRef_OneOf
		expectedError string
	}{
		{
			name:          "no fields set",
			input:         &KMSKeyRef_OneOf{},
			expectedError: "a key reference must be provided: specify one of '.name', '.keyHandleRef.name' or '.external'",
		},
		{
			name:          "empty external",
			input:         &KMSKeyRef_OneOf{External: ""},
			expectedError: "a key reference must be provided: specify one of '.name', '.keyHandleRef.name' or '.external'",
		},
		{
			name: "only kmsCryptoKeyRef is set",
			input: &KMSKeyRef_OneOf{
				kmsCryptoKeyRef: &kmsCryptoKeyRef{Name: "test-key"},
			},
			expectedError: "",
		},
		{
			name: "only keyHandleRef is set",
			input: &KMSKeyRef_OneOf{
				KMSKeyHandleRef: &kmsKeyHandleRef{Name: "auto-key"},
			},
			expectedError: "",
		},
		{
			name: "only External is set",
			input: &KMSKeyRef_OneOf{
				External: "external-key",
			},
			expectedError: "",
		},
		{
			name: "kmsCryptoKeyRef and kmsKeyHandleRef are set",
			input: &KMSKeyRef_OneOf{
				kmsCryptoKeyRef: &kmsCryptoKeyRef{Name: "test-key"},
				KMSKeyHandleRef: &kmsKeyHandleRef{Name: "auto-key"},
			},
			expectedError: "exactly one of '.name', '.keyHandleRef.name' or '.external' must be specified, but 2 were found",
		},
		{
			name: "kmsCryptoKeyRef and External are set",
			input: &KMSKeyRef_OneOf{
				kmsCryptoKeyRef: &kmsCryptoKeyRef{Name: "test-key"},
				External:        "external-key",
			},
			expectedError: "exactly one of '.name', '.keyHandleRef.name' or '.external' must be specified, but 2 were found",
		},
		{
			name: "kmsKeyHandleRef and External are set",
			input: &KMSKeyRef_OneOf{
				KMSKeyHandleRef: &kmsKeyHandleRef{Name: "auto-key"},
				External:        "external-key",
			},
			expectedError: "exactly one of '.name', '.keyHandleRef.name' or '.external' must be specified, but 2 were found",
		},
		{
			name: "all three fields are set",
			input: &KMSKeyRef_OneOf{
				kmsCryptoKeyRef: &kmsCryptoKeyRef{Name: "test-key"},
				KMSKeyHandleRef: &kmsKeyHandleRef{Name: "auto-key"},
				External:        "external-key",
			},
			expectedError: "exactly one of '.name', '.keyHandleRef.name' or '.external' must be specified, but 3 were found",
		},
		{
			name: "all three fields are set2",
			input: &KMSKeyRef_OneOf{
				kmsCryptoKeyRef: &kmsCryptoKeyRef{Name: "test-key"},
				KMSKeyHandleRef: &kmsKeyHandleRef{Namespace: "auto-key-namespace"},
				External:        "",
			},
			expectedError: "exactly one of '.name', '.keyHandleRef.name' or '.external' must be specified, but 2 were found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.validateOneOf()

			if tc.expectedError == "" {
				if err != nil {
					t.Errorf("expected no error, but got: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected error '%s', but got none", tc.expectedError)
				} else if err.Error() != tc.expectedError {
					t.Errorf("expected error '%s', but got '%s'", tc.expectedError, err.Error())
				}
			}
		})
	}
}
