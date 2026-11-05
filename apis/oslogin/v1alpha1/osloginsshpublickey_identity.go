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
	_ identity.ServerGeneratedIdentity = &OSLoginSSHPublicKeyIdentity{}
	_ identity.Resource                = &OSLoginSSHPublicKey{}
)

var OSLoginSSHPublicKeyIdentityFormat = gcpurls.Template[OSLoginSSHPublicKeyIdentity]("oslogin.googleapis.com", "users/{user}/sshPublicKeys/{fingerprint}")

// OSLoginSSHPublicKeyIdentity is the identity of a GCP OSLoginSSHPublicKey resource.
// +k8s:deepcopy-gen=false
type OSLoginSSHPublicKeyIdentity struct {
	User        string
	Fingerprint string
}

func (i *OSLoginSSHPublicKeyIdentity) HasIdentitySpecified() bool {
	return i.Fingerprint != ""
}

func (i *OSLoginSSHPublicKeyIdentity) String() string {
	return OSLoginSSHPublicKeyIdentityFormat.ToString(*i)
}

func (i *OSLoginSSHPublicKeyIdentity) ParentString() string {
	return "users/" + i.User
}

func (i *OSLoginSSHPublicKeyIdentity) FromExternal(ref string) error {
	parsed, match, err := OSLoginSSHPublicKeyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of OSLoginSSHPublicKey external=%q was not known (use %s): %w", ref, OSLoginSSHPublicKeyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of OSLoginSSHPublicKey external=%q was not known (use %s)", ref, OSLoginSSHPublicKeyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *OSLoginSSHPublicKeyIdentity) Host() string {
	return OSLoginSSHPublicKeyIdentityFormat.Host()
}

func getIdentityFromOSLoginSSHPublicKeySpec(ctx context.Context, reader client.Reader, obj *OSLoginSSHPublicKey) (*OSLoginSSHPublicKeyIdentity, error) {
	if obj.Spec.User == "" {
		return nil, fmt.Errorf(".spec.user must be set")
	}

	// For OSLoginSSHPublicKey, resourceID is optional and can be empty.
	// We retrieve it directly from Spec.ResourceID to avoid falling back to GetName().
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if strings.Contains(resourceID, "/") {
		id := &OSLoginSSHPublicKeyIdentity{}
		if err := id.FromExternal(resourceID); err != nil {
			return nil, err
		}
		if id.User != obj.Spec.User {
			return nil, fmt.Errorf("cannot change user in OSLoginSSHPublicKey identity (spec=%q, resourceID=%q)", obj.Spec.User, id.User)
		}
		return id, nil
	}

	identity := &OSLoginSSHPublicKeyIdentity{
		User:        obj.Spec.User,
		Fingerprint: resourceID,
	}
	return identity, nil
}

func (obj *OSLoginSSHPublicKey) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromOSLoginSSHPublicKeySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	statusFingerprint := common.ValueOf(obj.Status.Fingerprint)
	if statusFingerprint != "" {
		if strings.Contains(statusFingerprint, "/") {
			statusIdentity := &OSLoginSSHPublicKeyIdentity{}
			if err := statusIdentity.FromExternal(statusFingerprint); err != nil {
				return nil, err
			}
			statusFingerprint = statusIdentity.Fingerprint
		}

		if specIdentity.Fingerprint == "" {
			specIdentity.Fingerprint = statusFingerprint
		}

		if statusFingerprint != specIdentity.Fingerprint {
			return nil, fmt.Errorf("cannot change OSLoginSSHPublicKey identity (old=%q, new=%q)", statusFingerprint, specIdentity.Fingerprint)
		}
	}

	return specIdentity, nil
}
