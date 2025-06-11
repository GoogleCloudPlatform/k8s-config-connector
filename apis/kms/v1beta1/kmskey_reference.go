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

import (
	"context"
	"fmt"
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.ExternalNormalizer = &KMSKeyRef_OneOf{}

// A reference to the KMSCryptoKey(manual management), or the KMSKeyHandle(automated management)
type KMSKeyRef_OneOf struct {
	// Default KMS crypto key. This is for API backward compatibility and cannot be changed.
	*kmsCryptoKeyRef `json:",inline"`

	// A reference to the Autokey `KMSKeyHandle`, which auto generates a crypto key.
	KMSKeyHandleRef *kmsKeyHandleRef `json:"keyHandleRef,omitempty"`

	// A reference to an externally managed KMSCryptoKey or KMSKeyHandle(AutoKey).
	// Should be in the format `projects/{{kms_project_id}}/locations/{{region}}/keyRings/{{key_ring_id}}/cryptoKeys/{{key}}`.
	// For AutoKey, replace {{key_ring_id}} to `autokey`, i.e. `projects/{{kms_project_id}}/locations/{{region}}/keyRings/autokey/cryptoKeys/{{key}}`.
	External string `json:"external,omitempty"`
}

func (r *KMSKeyRef_OneOf) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.validateOneOf(); err != nil {
		return "", err
	}

	if r.External != "" {
		// Resolve the external managed reference resource by external value
		tokens := strings.Split(r.External, "/")
		if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "keyRings" && tokens[6] == "cryptoKeys" {
			return r.External, nil
		}
		return "", fmt.Errorf("format of KMSKeyRef external=%q was not known (use projects/{{kms_project_id}}/locations/{{region}}/keyRings/{{key_ring_id}}/cryptoKeys/{{key}})", r.External)
	}

	// Resolve the KCC managed reference resource by its name
	if r.kmsCryptoKeyRef != nil {
		// Use KMSCryptoKey
		cryptoKey, err := r.kmsCryptoKeyRef.normalizedExternal(ctx, reader, otherNamespace)
		if err != nil {
			return "", err
		}
		r.External = cryptoKey
	} else {
		// Use KMSAutoKey
		autoKey, err := r.KMSKeyHandleRef.normalizedExternal(ctx, reader, otherNamespace)
		if err != nil {
			return "", err
		}
		r.External = autoKey
	}
	return r.External, nil
}

// validateOneOf checks that exactly one of the key reference fields is set
func (r *KMSKeyRef_OneOf) validateOneOf() error {
	numOfNonNil := 0
	if r.kmsCryptoKeyRef != nil {
		numOfNonNil++
	}
	if r.KMSKeyHandleRef != nil {
		numOfNonNil++
	}
	if r.External != "" {
		numOfNonNil++
	}
	if numOfNonNil == 0 {
		return fmt.Errorf("a key reference must be provided: specify one of '.name', '.keyHandleRef.name' or '.external'")
	}
	if numOfNonNil > 1 {
		return fmt.Errorf("exactly one of '.name', '.keyHandleRef.name' or '.external' must be specified, but %d were found", numOfNonNil)
	}
	return nil
}
