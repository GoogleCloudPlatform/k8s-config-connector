// Copyright 2024 Google LLC
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

// NOTE: If references a KMSKey, use `KMSKeyRef_OneOf` instead!
// todo: use unexported variable kmsCryptoKeyRef to avoid referencing to the KMS crypto key.

// KMSCryptoKeyRef is a reference to a KMSCryptoKey.
// Deprecated: Use kmsv1beta1.KMSCryptoKeyRef instead.

type KMSCryptoKeyRef struct {
	// A reference to an externally managed KMSCryptoKey.
	// Should be in the format `projects/[kms_project_id]/locations/[region]/keyRings/[key_ring_id]/cryptoKeys/[key]`.
	External string `json:"external,omitempty"`

	// The `name` of a `KMSCryptoKey` resource.
	Name string `json:"name,omitempty"`

	// The `namespace` of a `KMSCryptoKey` resource.
	Namespace string `json:"namespace,omitempty"`
}
