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
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.ServerGeneratedIdentity = &KMSSecretCiphertextIdentity{}
	_ identity.Resource                = &KMSSecretCiphertext{}
)

var KMSSecretCiphertextIdentityFormat = gcpurls.Template[KMSSecretCiphertextIdentity]("cloudkms.googleapis.com", "projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{cryptokey}/{ciphertext}")

// KMSSecretCiphertextIdentity is the identity of a GCP KMSSecretCiphertext resource.
// +k8s:deepcopy-gen=false
type KMSSecretCiphertextIdentity struct {
	Project    string
	Location   string
	KeyRing    string
	CryptoKey  string
	Ciphertext string
}

func (i *KMSSecretCiphertextIdentity) HasIdentitySpecified() bool {
	return i.Ciphertext != ""
}

func (i *KMSSecretCiphertextIdentity) String() string {
	return KMSSecretCiphertextIdentityFormat.ToString(*i)
}

func (i *KMSSecretCiphertextIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", i.Project, i.Location, i.KeyRing, i.CryptoKey)
}

func (i *KMSSecretCiphertextIdentity) FromExternal(ref string) error {
	s := strings.TrimPrefix(ref, "https:")
	s = strings.TrimPrefix(s, "http:")
	s = strings.TrimPrefix(s, "//")

	host := KMSSecretCiphertextIdentityFormat.Host()
	if strings.HasPrefix(s, host+"/") {
		s = strings.TrimPrefix(s, host)
	}

	s = strings.Trim(s, "/")
	parts := strings.Split(s, "/")

	if len(parts) >= 9 &&
		parts[0] == "projects" &&
		parts[2] == "locations" &&
		parts[4] == "keyRings" &&
		parts[6] == "cryptoKeys" {
		i.Project = parts[1]
		i.Location = parts[3]
		i.KeyRing = parts[5]
		i.CryptoKey = parts[7]
		i.Ciphertext = strings.Join(parts[8:], "/")
		return nil
	}

	parsed, match, err := KMSSecretCiphertextIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of KMSSecretCiphertext external=%q was not known (use %s): %w", ref, KMSSecretCiphertextIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of KMSSecretCiphertext external=%q was not known (use %s)", ref, KMSSecretCiphertextIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *KMSSecretCiphertextIdentity) Host() string {
	return KMSSecretCiphertextIdentityFormat.Host()
}

func getIdentityFromKMSSecretCiphertextSpec(ctx context.Context, reader client.Reader, obj *KMSSecretCiphertext) (*KMSSecretCiphertextIdentity, error) {
	if obj.Spec.CryptoKey == "" {
		return nil, fmt.Errorf(".spec.cryptoKey must be set")
	}

	parsed, match, err := cryptoKeyFormat.Parse(obj.Spec.CryptoKey)
	if err != nil {
		return nil, fmt.Errorf("cannot parse .spec.cryptoKey: %w", err)
	}
	if !match {
		return nil, fmt.Errorf("format of .spec.cryptoKey was not known (use %s): %q", cryptoKeyFormat.CanonicalForm(), obj.Spec.CryptoKey)
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)

	identity := &KMSSecretCiphertextIdentity{
		Project:    parsed.Project,
		Location:   parsed.Location,
		KeyRing:    parsed.Keyring,
		CryptoKey:  parsed.Cryptokey,
		Ciphertext: resourceID,
	}
	return identity, nil
}

func (obj *KMSSecretCiphertext) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromKMSSecretCiphertextSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	statusCiphertext := common.ValueOf(obj.Status.Ciphertext)
	if statusCiphertext != "" {
		statusIdentity := &KMSSecretCiphertextIdentity{
			Project:    specIdentity.Project,
			Location:   specIdentity.Location,
			KeyRing:    specIdentity.KeyRing,
			CryptoKey:  specIdentity.CryptoKey,
			Ciphertext: statusCiphertext,
		}

		if specIdentity.Ciphertext == "" {
			specIdentity.Ciphertext = statusIdentity.Ciphertext
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change KMSSecretCiphertext identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
