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
	_ identity.IdentityV2 = &KMSKeyRingImportJobIdentity{}
	_ identity.Resource   = &KMSKeyRingImportJob{}
)

var KMSKeyRingImportJobIdentityFormat = gcpurls.Template[KMSKeyRingImportJobIdentity]("cloudkms.googleapis.com", "projects/{project}/locations/{location}/keyRings/{keyring}/importJobs/{importjob}")

var keyRingFormat = gcpurls.Template[keyRingIdentity]("cloudkms.googleapis.com", "projects/{project}/locations/{location}/keyRings/{keyring}")

type keyRingIdentity struct {
	Project  string
	Location string
	KeyRing  string
}

// KMSKeyRingImportJobIdentity is the identity of a GCP KMSKeyRingImportJob resource.
// +k8s:deepcopy-gen=false
type KMSKeyRingImportJobIdentity struct {
	Project   string
	Location  string
	KeyRing   string
	ImportJob string
}

func (i *KMSKeyRingImportJobIdentity) String() string {
	return KMSKeyRingImportJobIdentityFormat.ToString(*i)
}

func (i *KMSKeyRingImportJobIdentity) FromExternal(ref string) error {
	parsed, match, err := KMSKeyRingImportJobIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of KMSKeyRingImportJob external=%q was not known (use %s): %w", ref, KMSKeyRingImportJobIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of KMSKeyRingImportJob external=%q was not known (use %s)", ref, KMSKeyRingImportJobIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *KMSKeyRingImportJobIdentity) Host() string {
	return KMSKeyRingImportJobIdentityFormat.Host()
}

func (i *KMSKeyRingImportJobIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", i.Project, i.Location, i.KeyRing)
}

func getIdentityFromKMSKeyRingImportJobSpec(ctx context.Context, reader client.Reader, obj *KMSKeyRingImportJob) (*KMSKeyRingImportJobIdentity, error) {
	if obj.Spec.KeyRing == "" {
		return nil, fmt.Errorf(".spec.keyRing must be set")
	}

	parsed, match, err := keyRingFormat.Parse(obj.Spec.KeyRing)
	if err != nil {
		return nil, fmt.Errorf("cannot parse .spec.keyRing: %w", err)
	}
	if !match {
		return nil, fmt.Errorf("format of .spec.keyRing was not known (use %s): %q", keyRingFormat.CanonicalForm(), obj.Spec.KeyRing)
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.Spec.ImportJobId
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	identity := &KMSKeyRingImportJobIdentity{
		Project:   parsed.Project,
		Location:  parsed.Location,
		KeyRing:   parsed.KeyRing,
		ImportJob: resourceID,
	}
	return identity, nil
}

func (obj *KMSKeyRingImportJob) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromKMSKeyRingImportJobSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	statusName := common.ValueOf(obj.Status.Name)
	if statusName != "" {
		statusIdentity := &KMSKeyRingImportJobIdentity{}
		if err := statusIdentity.FromExternal(statusName); err != nil {
			return nil, err
		}

		if specIdentity.ImportJob == "" {
			specIdentity.ImportJob = statusIdentity.ImportJob
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change KMSKeyRingImportJob identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
