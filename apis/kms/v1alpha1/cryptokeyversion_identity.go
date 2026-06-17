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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.ServerGeneratedIdentity = &KMSCryptoKeyVersionIdentity{}
	_ identity.Resource                = &KMSCryptoKeyVersion{}
)

var KMSCryptoKeyVersionIdentityFormat = gcpurls.Template[KMSCryptoKeyVersionIdentity]("cloudkms.googleapis.com", "projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{cryptokey}/cryptoKeyVersions/{cryptokeyversion}")

var cryptoKeyFormat = gcpurls.Template[cryptoKeyIdentity]("cloudkms.googleapis.com", "projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{cryptokey}")

type cryptoKeyIdentity struct {
	Project   string
	Location  string
	Keyring   string
	Cryptokey string
}

// KMSCryptoKeyVersionIdentity is the identity of a GCP KMSCryptoKeyVersion resource.
// +k8s:deepcopy-gen=false
type KMSCryptoKeyVersionIdentity struct {
	Project          string
	Location         string
	KeyRing          string
	CryptoKey        string
	CryptoKeyVersion string
}

func (i *KMSCryptoKeyVersionIdentity) HasIdentitySpecified() bool {
	return i.CryptoKeyVersion != ""
}

func (i *KMSCryptoKeyVersionIdentity) String() string {
	return KMSCryptoKeyVersionIdentityFormat.ToString(*i)
}

func (i *KMSCryptoKeyVersionIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", i.Project, i.Location, i.KeyRing, i.CryptoKey)
}

func (i *KMSCryptoKeyVersionIdentity) FromExternal(ref string) error {
	parsed, match, err := KMSCryptoKeyVersionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of KMSCryptoKeyVersion external=%q was not known (use %s): %w", ref, KMSCryptoKeyVersionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of KMSCryptoKeyVersion external=%q was not known (use %s)", ref, KMSCryptoKeyVersionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *KMSCryptoKeyVersionIdentity) Host() string {
	return KMSCryptoKeyVersionIdentityFormat.Host()
}

func getIdentityFromKMSCryptoKeyVersionSpec(ctx context.Context, reader client.Reader, obj *KMSCryptoKeyVersion) (*KMSCryptoKeyVersionIdentity, error) {
	if obj.Spec.CryptoKey == "" {
		return nil, fmt.Errorf(".spec.cryptoKey must be set")
	}

	// We use a self-contained format template (cryptoKeyFormat) to parse obj.Spec.CryptoKey rather than
	// importing the KMSCryptoKeyIdentity format from the v1beta1 package. This prevents a circular go import
	// dependency, because v1beta1 types sometimes need to reference v1alpha1 definitions (or vice-versa).
	parsed, match, err := cryptoKeyFormat.Parse(obj.Spec.CryptoKey)
	if err != nil {
		return nil, fmt.Errorf("cannot parse .spec.cryptoKey: %w", err)
	}
	if !match {
		return nil, fmt.Errorf("format of .spec.cryptoKey was not known (use %s): %q", cryptoKeyFormat.CanonicalForm(), obj.Spec.CryptoKey)
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)

	identity := &KMSCryptoKeyVersionIdentity{
		Project:          parsed.Project,
		Location:         parsed.Location,
		KeyRing:          parsed.Keyring,
		CryptoKey:        parsed.Cryptokey,
		CryptoKeyVersion: resourceID,
	}
	return identity, nil
}

func (obj *KMSCryptoKeyVersion) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromKMSCryptoKeyVersionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	statusName := common.ValueOf(obj.Status.Name)
	if statusName != "" {
		statusIdentity := &KMSCryptoKeyVersionIdentity{}
		if err := statusIdentity.FromExternal(statusName); err != nil {
			return nil, err
		}

		if specIdentity.CryptoKeyVersion == "" {
			specIdentity.CryptoKeyVersion = statusIdentity.CryptoKeyVersion
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change KMSCryptoKeyVersion identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
