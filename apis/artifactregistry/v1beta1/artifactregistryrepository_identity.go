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

func (obj *ArtifactRegistryRepository) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Get Location
	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	// Get Project
	projectID := obj.GetAnnotations()["cnrm.cloud.google.com/project-id"]
	if projectID == "" {
		projectID = obj.GetNamespace()
	}
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	// Use approved External
	// Status.Name matches the google.devtools.artifactregistry.v1.Repository.name field, which is the full resource name.
	externalRef := common.ValueOf(obj.Status.Name)
	if externalRef != "" {
		// Validate desired with actual
		actualIdentity := &ArtifactRegistryRepositoryIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actualIdentity.Project != projectID {
			return nil, fmt.Errorf("project changed, expect %s, got %s", actualIdentity.Project, projectID)
		}
		if actualIdentity.Location != location {
			return nil, fmt.Errorf("location changed, expect %s, got %s", actualIdentity.Location, location)
		}
		if actualIdentity.Repository != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.Repository)
		}
	}

	return &ArtifactRegistryRepositoryIdentity{
		Project:    projectID,
		Location:   location,
		Repository: resourceID,
	}, nil
}
