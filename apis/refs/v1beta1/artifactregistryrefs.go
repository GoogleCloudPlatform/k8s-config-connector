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

import (
	"context"
	"fmt"
	"strings"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ArtifactRegistryRepositoryRef struct {
	// A reference to an externally managed ArtifactRegistryRepository.
	// Should be in the format `projects/[project_id]/locations/[location]/repositories/[repository_id]`.
	External string `json:"external,omitempty"`

	// The `name` of an `ArtifactRegistryRepository` resource.
	Name string `json:"name,omitempty"`

	// The `namespace` of an `ArtifactRegistryRepository` resource.
	Namespace string `json:"namespace,omitempty"`
}

// ResolveArtifactRegistryRepositoryRef will resolve an ArtifactRegistryRepositoryRef to an ArtifactRegistryRepository.
func ResolveArtifactRegistryRepositoryRef(ctx context.Context, reader client.Reader, src client.Object, ref *ArtifactRegistryRepositoryRef) (*ArtifactRegistryRepositoryRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on ArtifactRegistryRepositoryRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on ArtifactRegistryRepositoryRef")
	}

	// External should be in the `projects/[project_id]/locations/[location]/repositories/[repository_id]` format
	if ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "repositories" {
			ref = &ArtifactRegistryRepositoryRef{
				External: fmt.Sprintf("projects/%s/locations/%s/repositories/%s", tokens[1], tokens[3], tokens[5]),
			}
			return ref, nil
		}
		return nil, fmt.Errorf("format of ArtifactRegistryRepositoryRef external=%q was not known (use projects/[project_id]/locations/[location]/repositories/[repository_id])", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	// Fetch object from k8s cluster to construct the external form
	repo := &unstructured.Unstructured{}
	repo.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "artifactregistry.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ArtifactRegistryRepository",
	})
	if err := reader.Get(ctx, key, repo); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced ArtifactRegistryRepository %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced ArtifactRegistryRepository %v: %w", key, err)
	}

	repoResourceID, err := GetResourceID(repo)
	if err != nil {
		return nil, err
	}

	projectID, err := ResolveProjectID(ctx, reader, repo)
	if err != nil {
		return nil, err
	}

	location, err := GetLocation(repo)
	if err != nil {
		return nil, err
	}

	ref = &ArtifactRegistryRepositoryRef{
		External: fmt.Sprintf("projects/%s/locations/%s/repositories/%s", projectID, location, repoResourceID),
	}

	return ref, nil
}