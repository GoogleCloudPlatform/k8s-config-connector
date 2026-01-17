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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

var _ refsv1beta1.ExternalNormalizer = &ArtifactRegistryRepositoryRef{}

// NormalizedExternal provision the "External" value for other resource that depends on ArtifactRegistryRepository.
// If the "External" is given in the other resource's spec.ArtifactRegistryRepositoryRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ArtifactRegistryRepository object from the cluster.
func (r *ArtifactRegistryRepositoryRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ArtifactRegistryRepositoryGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if err := r.ValidateExternal(r.External); err != nil {
			return "", err
		}
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ArtifactRegistryRepositoryGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", ArtifactRegistryRepositoryGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef != "" {
		r.External = actualExternalRef
		return r.External, nil
	}

	// Backward compatible to Terraform/DCL based resource, which does not have status.externalRef.
	// But ArtifactRegistryRepository is direct, so it should have externalRef.
	// We can try to construct it if missing, or just fail/return empty.
	// Let's try to resolve it.
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return "", err
	}

	location, _, err := unstructured.NestedString(u.Object, "spec", "location")
	if err != nil {
		return "", fmt.Errorf("reading spec.location: %w", err)
	}
	if location == "" {
		return "", fmt.Errorf("spec.location is empty")
	}

	resourceID, err := refsv1beta1.GetResourceID(u)
	if err != nil {
		return "", err
	}

	// projects/{project}/locations/{location}/repositories/{repository}
	identity := &ArtifactRegistryRepositoryIdentity{
		Project:    projectID,
		Location:   location,
		Repository: resourceID,
	}
	r.External = identity.String()
	return r.External, nil
}

func (r *ArtifactRegistryRepositoryRef) GetGVK() schema.GroupVersionKind {
	return ArtifactRegistryRepositoryGVK
}

func (r *ArtifactRegistryRepositoryRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ArtifactRegistryRepositoryRef) GetExternal() string {
	return r.External
}

func (r *ArtifactRegistryRepositoryRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ArtifactRegistryRepositoryRef) ValidateExternal(ref string) error {
	id := &ArtifactRegistryRepositoryIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *ArtifactRegistryRepositoryRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}
		
		location, _, err := unstructured.NestedString(u.Object, "spec", "location")
		if err != nil {
			return ""
		}
		if location == "" {
			return ""
		}

		resourceID, err := refsv1beta1.GetResourceID(u)
		if err != nil {
			return ""
		}

		identity := &ArtifactRegistryRepositoryIdentity{
			Project:    projectID,
			Location:   location,
			Repository: resourceID,
		}
		return identity.String()
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}