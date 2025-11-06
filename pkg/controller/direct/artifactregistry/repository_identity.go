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

package artifactregistry

import (
	"context"
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ArtifactRegistryRepositoryIdentity struct {
	Parent     ArtifactRegistryRepositoryParent
	ResourceID string
}

type ArtifactRegistryRepositoryParent struct {
	ProjectID string
	Location  string
}

func (p *ArtifactRegistryRepositoryParent) String() string {
	return fmt.Sprintf("projects/%s/locations/%s", p.ProjectID, p.Location)
}

func (i *ArtifactRegistryRepositoryIdentity) FullyQualifiedName() string {
	return fmt.Sprintf("%s/repositories/%s", i.Parent.String(), i.ResourceID)
}

func NewArtifactRegistryRepositoryIdentity(ctx context.Context, reader client.Reader, obj *krm.ArtifactRegistryRepository) (*ArtifactRegistryRepositoryIdentity, error) {
	// Get project ID from annotation (standard Config Connector approach)
	var projectID string
	annotations := obj.GetAnnotations()
	if annotations != nil {
		if pid, ok := annotations["cnrm.cloud.google.com/project-id"]; ok {
			projectID = pid
		}
	}
	if projectID == "" {
		return nil, fmt.Errorf("missing required annotation cnrm.cloud.google.com/project-id")
	}

	// Get location
	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("missing required field spec.location")
	}

	// Get resource ID
	resourceID := obj.Spec.ResourceID
	if resourceID == nil {
		resourceID = &obj.ObjectMeta.Name
	}

	return &ArtifactRegistryRepositoryIdentity{
		Parent: ArtifactRegistryRepositoryParent{
			ProjectID: projectID,
			Location:  location,
		},
		ResourceID: *resourceID,
	}, nil
}