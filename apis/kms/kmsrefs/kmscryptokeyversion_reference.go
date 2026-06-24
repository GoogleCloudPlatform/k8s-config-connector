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

package kmsrefs

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/kccscheme"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var KMSCryptoKeyVersionGVK = schema.GroupVersionKind{
	Group:   "kms.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "KMSCryptoKeyVersion",
}

var (
	_ refs.Ref         = &KMSCryptoKeyVersionRef{}
	_ refs.ExternalRef = &KMSCryptoKeyVersionRef{}
)

type KMSCryptoKeyVersionRef struct {
	// A reference to an externally managed cryptoKeyVersion.
	// Should be in the format `projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}/cryptoKeyVersions/{{version}}`.
	External string `json:"external,omitempty"`

	// The `name` of a `KMSCryptoKeyVersion` resource.
	Name string `json:"name,omitempty"`

	// The `namespace` of a `KMSCryptoKeyVersion` resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	kccscheme.RegisterRef(&KMSCryptoKeyVersionRef{}, KMSCryptoKeyVersionGVK)
}

func (r *KMSCryptoKeyVersionRef) GetGVK() schema.GroupVersionKind {
	return KMSCryptoKeyVersionGVK
}

func (r *KMSCryptoKeyVersionRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *KMSCryptoKeyVersionRef) GetExternal() string {
	return r.External
}

func (r *KMSCryptoKeyVersionRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *KMSCryptoKeyVersionRef) ValidateExternal(ref string) error {
	tokens := strings.Split(ref, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "keyRings" && tokens[6] == "cryptoKeys" && tokens[8] == "cryptoKeyVersions" {
		return nil
	}
	return fmt.Errorf("format of KMSCryptoKeyVersionRef external=%q was not known (use projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}/cryptoKeyVersions/{{version}})", ref)
}

func (r *KMSCryptoKeyVersionRef) ParseExternalToIdentity() (identity.Identity, error) {
	return nil, fmt.Errorf("parsing KMSCryptoKeyVersion to identity is not yet supported")
}

func (r *KMSCryptoKeyVersionRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
