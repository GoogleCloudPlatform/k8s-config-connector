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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/artifactregistry/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ArtifactRegistryRepositoryRef{}

// ArtifactRegistryRepositoryRef defines the resource reference to ArtifactRegistryRepository, which "External" field
// holds the GCP identifier for the KRM object.
// +kcc:proto=google.devtools.artifactregistry.v1.Repository
type ArtifactRegistryRepositoryRef struct {
	// A reference to an externally managed ArtifactRegistryRepository resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/repositories/{{repositoryID}}".
	External string `json:"external,omitempty"`

	// The name of a ArtifactRegistryRepository resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ArtifactRegistryRepository resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ArtifactRegistryRepositoryRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "artifactregistry.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ArtifactRegistryRepository",
	}
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
	if !strings.HasPrefix(ref, "projects/") || !strings.Contains(ref, "/locations/") || !strings.Contains(ref, "/repositories/") {
		return fmt.Errorf("ArtifactRegistryRepositoryRef.External must be in the format \"projects/{{projectID}}/locations/{{location}}/repositories/{{repositoryID}}\", got %s", ref)
	}
	return nil
}

func (r *ArtifactRegistryRepositoryRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" && r.Name != "" {
		return fmt.Errorf("cannot specify both name and external on ArtifactRegistryRepositoryRef")
	}
	if r.External != "" {
		return r.ValidateExternal(r.External)
	}

	if r.Namespace == "" {
		r.Namespace = defaultNamespace
	}

	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &krm.ArtifactRegistryRepository{}
	if err := reader.Get(ctx, key, u); err != nil {
		return err
	}

	projectID, err := resolveProject(ctx, reader, u)
	if err != nil {
		return err
	}

	location := u.Spec.Location
	repositoryID := u.Spec.ResourceID
	if repositoryID == nil || *repositoryID == "" {
		repositoryID = &u.Name
	}

	r.External = fmt.Sprintf("projects/%s/locations/%s/repositories/%s", projectID, location, *repositoryID)
	return nil
}

func resolveProject(ctx context.Context, reader client.Reader, obj *krm.ArtifactRegistryRepository) (string, error) {
	// Check annotation on the object
	if projectID := obj.GetAnnotations()["cnrm.cloud.google.com/project-id"]; projectID != "" {
		return projectID, nil
	}

	// Check annotation on the namespace
	ns := &corev1.Namespace{}
	if err := reader.Get(ctx, types.NamespacedName{Name: obj.GetNamespace()}, ns); err != nil {
		return "", fmt.Errorf("error getting namespace %q: %w", obj.GetNamespace(), err)
	}
	if projectID := ns.GetAnnotations()["cnrm.cloud.google.com/project-id"]; projectID != "" {
		return projectID, nil
	}

	// Fallback to namespace name
	return obj.GetNamespace(), nil
}
