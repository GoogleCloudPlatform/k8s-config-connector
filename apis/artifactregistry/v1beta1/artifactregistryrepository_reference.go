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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var ArtifactRegistryRepositoryGVK = schema.GroupVersionKind{
	Group:   "artifactregistry.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ArtifactRegistryRepository",
}

var _ refs.Ref = &ArtifactRegistryRepositoryRef{}

// ArtifactRegistryRepositoryRef defines the resource reference to ArtifactRegistryRepository, which "External" field
// holds the GCP identifier for the KRM object.
type ArtifactRegistryRepositoryRef struct {
	// A reference to an externally managed ArtifactRegistryRepository resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/repositories/{{repositoryID}}".
	External string `json:"external,omitempty"`

	// The name of a ArtifactRegistryRepository resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ArtifactRegistryRepository resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ArtifactRegistryRepositoryRef{})
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
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ArtifactRegistryRepositoryRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ArtifactRegistryRepositoryIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ArtifactRegistryRepositoryRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		resourceID, err := refs.GetResourceID(u)
		if err != nil {
			return ""
		}

		location, err := refs.GetLocation(u)
		if err != nil {
			return ""
		}

		projectID, err := refs.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}

		return fmt.Sprintf("projects/%s/locations/%s/repositories/%s", projectID, location, resourceID)
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
