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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveArtifactRegistryRepositoryRefs(ctx context.Context, kube client.Reader, obj *krm.ArtifactRegistryRepository) error {
	// Resolve repository references in virtual repository config
	if err := resolveVirtualRepositoryRefs(ctx, kube, obj); err != nil {
		return err
	}
	
	return nil
}

func resolveVirtualRepositoryRefs(ctx context.Context, kube client.Reader, obj *krm.ArtifactRegistryRepository) error {
	if obj.Spec.VirtualRepositoryConfig == nil {
		return nil
	}

	for i := range obj.Spec.VirtualRepositoryConfig.UpstreamPolicies {
		upstreamPolicy := &obj.Spec.VirtualRepositoryConfig.UpstreamPolicies[i]
		
		if upstreamPolicy.RepositoryRef == nil {
			continue
		}

		repositoryRef := upstreamPolicy.RepositoryRef

		if repositoryRef.External != nil && repositoryRef.Name != nil {
			return fmt.Errorf("cannot specify both external and name in repositoryRef")
		}

		if repositoryRef.External != nil {
			// Already resolved
			continue
		} else if repositoryRef.Name != nil {
			if repositoryRef.Namespace == nil {
				repositoryRef.Namespace = &obj.Namespace
			}

			key := types.NamespacedName{
				Namespace: *repositoryRef.Namespace,
				Name:      *repositoryRef.Name,
			}

			repository := &unstructured.Unstructured{}
			repository.SetGroupVersionKind(krm.ArtifactRegistryRepositoryGVK)
			if err := kube.Get(ctx, key, repository); err != nil {
				if apierrors.IsNotFound(err) {
					return k8s.NewReferenceNotFoundError(krm.ArtifactRegistryRepositoryGVK, key)
				}
				return fmt.Errorf("error reading referenced ArtifactRegistryRepository %v: %w", key, err)
			}

			// Get the resourceID (defaults to name if not specified)
			resourceID, _, err := unstructured.NestedString(repository.Object, "spec", "resourceID")
			if err != nil {
				return fmt.Errorf("error reading resourceID from %v: %w", key, err)
			}
			if resourceID == "" {
				resourceID = *repositoryRef.Name
			}

			// Get project ID
			projectID, err := refs.ResolveProjectID(ctx, kube, repository)
			if err != nil {
				return fmt.Errorf("cannot resolve project for repository %v: %w", key, err)
			}

			// Get location
			location, found, err := unstructured.NestedString(repository.Object, "spec", "location")
			if err != nil || !found {
				return fmt.Errorf("cannot read location from repository %v: %w", key, err)
			}

			// Set the external reference
			external := fmt.Sprintf("projects/%s/locations/%s/repositories/%s", projectID, location, resourceID)
			repositoryRef.External = &external

		} else {
			return fmt.Errorf("must specify either external or name in repositoryRef")
		}
	}

	return nil
}