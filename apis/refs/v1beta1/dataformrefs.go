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
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DataformRepositoryRef struct {
	/* The `projectID` field of a project, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `Project` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `Project` resource. */
	Namespace string `json:"namespace,omitempty"`
}

func (r *DataformRepositoryRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "dataform.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "DataformRepository",
	}
}

func (r *DataformRepositoryRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *DataformRepositoryRef) GetExternal() string {
	return r.External
}

func (r *DataformRepositoryRef) SetExternal(external string) {
	r.External = external
}

func (r *DataformRepositoryRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return Normalize(ctx, reader, r, defaultNamespace)
}

func (r *DataformRepositoryRef) ValidateExternal(external string) error {
	if !strings.HasPrefix(external, "projects/") {
		return fmt.Errorf("external DataformRepository reference must be in format projects/{project}/locations/{location}/repositories/{repository}, got %s", external)
	}
	return nil
}

func (r *DataformRepositoryRef) FullyQualifiedName(ctx context.Context, reader client.Reader, defaultNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, defaultNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}

type DataformRepositoryReleaseConfigRef struct {
	/* The `projectID` field of a project, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `Project` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `Project` resource. */
	Namespace string `json:"namespace,omitempty"`
}

func (r *DataformRepositoryReleaseConfigRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "dataform.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "DataformRepositoryReleaseConfig",
	}
}

func (r *DataformRepositoryReleaseConfigRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *DataformRepositoryReleaseConfigRef) GetExternal() string {
	return r.External
}

func (r *DataformRepositoryReleaseConfigRef) SetExternal(external string) {
	r.External = external
}

func (r *DataformRepositoryReleaseConfigRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return Normalize(ctx, reader, r, defaultNamespace)
}

func (r *DataformRepositoryReleaseConfigRef) ValidateExternal(external string) error {
	if !strings.HasPrefix(external, "projects/") {
		return fmt.Errorf("external DataformRepositoryReleaseConfig reference must be in format projects/{project}/locations/{location}/repositories/{repository}/releaseConfigs/{release_config}, got %s", external)
	}
	return nil
}

func (r *DataformRepositoryReleaseConfigRef) FullyQualifiedName(ctx context.Context, reader client.Reader, defaultNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, defaultNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}
