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

package v1alpha1

import (
	"context"
	"fmt"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &KMSCryptoKeyVersionRef{}
var KMSCryptoKeyVersionGVK = GroupVersion.WithKind("KMSCryptoKeyVersion")

// KMSCryptoKeyVersionRef defines the resource reference to KMSCryptoKeyVersion, which "External" field
// holds the GCP identifier for the KRM object.
type KMSCryptoKeyVersionRef struct {
	// A reference to an externally managed cryptoKeyVersion.
	// Should be in the format `projects/{{kms_project_id}}/locations/{{region}}/keyRings/{{key_ring_id}}/cryptoKeys/{{key}}/cryptoKeyVersions/{{version}}`.
	External string `json:"external,omitempty"`

	// The `name` of a `KMSCryptoKey` resource.
	//Name string `json:"name,omitempty"`
	// The `namespace` of a `KMSCryptoKey` resource.
	//Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on KMSCryptoKeyVersionRef.
// If the "External" is given in the other resource's spec.KMSCryptoKeyVersionRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual KMSCryptoKeyVersionRef object from the cluster.
func (r *KMSCryptoKeyVersionRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	// todo: Currently cannot reference to KMSCryptoKeyVersion name. Implemented once the KMSCryptoKeyVersion resource is supported by KCC
	//if r.External != "" && r.Name != "" {
	//	return "", fmt.Errorf("cannot specify both name and external on %s reference", KMSCryptoKeyVersionGVK.Kind)
	//}

	// From given External
	// External should be in the `projects/{{kms_project_id}}/locations/{{region}}/keyRings/{{key_ring_id}}/cryptoKeys/{{key}}/cryptoKeyVersions/{{version}}` format
	if r.External != "" {
		//if _, err := ParseKMSCryptoKeyVersionExternal(r.External); err != nil {
		//	return "", err
		//}
		return r.External, nil
	}

	// From the Config Connector object
	return "", fmt.Errorf("no External specified")
}
