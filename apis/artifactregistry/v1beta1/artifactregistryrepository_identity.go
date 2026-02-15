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

package v1beta1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ArtifactRegistryRepositoryIdentity{}
	_ identity.Resource   = &ArtifactRegistryRepository{}
)

var ArtifactRegistryRepositoryIdentityFormat = gcpurls.Template[ArtifactRegistryRepositoryIdentity]("artifactregistry.googleapis.com", "projects/{project}/locations/{location}/repositories/{repository}")

// +k8s:deepcopy-gen=false
type ArtifactRegistryRepositoryIdentity struct {
	Project    string
	Location   string
	Repository string
}

func (i *ArtifactRegistryRepositoryIdentity) String() string {
	return ArtifactRegistryRepositoryIdentityFormat.ToString(*i)
}

func (i *ArtifactRegistryRepositoryIdentity) FromExternal(ref string) error {
	parsed, match, err := ArtifactRegistryRepositoryIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ArtifactRegistryRepository external=%q was not known (use %s): %w", ref, ArtifactRegistryRepositoryIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ArtifactRegistryRepository external=%q was not known (use %s)", ref, ArtifactRegistryRepositoryIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ArtifactRegistryRepositoryIdentity) Host() string {
	return ArtifactRegistryRepositoryIdentityFormat.Host()
}

func getIdentityFromArtifactRegistryRepositorySpec(ctx context.Context, reader client.Reader, obj client.Object) (*ArtifactRegistryRepositoryIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ArtifactRegistryRepositoryIdentity{
		Project:    projectID,
		Location:   location,
		Repository: resourceID,
	}
	return identity, nil
}

func (obj *ArtifactRegistryRepository) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromArtifactRegistryRepositorySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	// NOTE: ArtifactRegistryRepository does not yet support status.externalRef, but we parse `name` instead
	externalRef := common.ValueOf(obj.Status.Name)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ArtifactRegistryRepositoryIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ArtifactRegistryRepository identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
