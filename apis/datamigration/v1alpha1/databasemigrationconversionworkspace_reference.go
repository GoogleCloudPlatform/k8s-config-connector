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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &DatabaseMigrationConversionWorkspaceRef{}

// DatabaseMigrationConversionWorkspaceRef is a reference to a DatabaseMigrationConversionWorkspace.
type DatabaseMigrationConversionWorkspaceRef struct {
	// A reference to an externally managed DatabaseMigrationConversionWorkspace resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/conversionWorkspaces/{{conversionWorkspaceID}}".
	External string `json:"external,omitempty"`

	// The name of a DatabaseMigrationConversionWorkspace resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DatabaseMigrationConversionWorkspace resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&DatabaseMigrationConversionWorkspaceRef{}, &DatabaseMigrationConversionWorkspace{})
}

func (r *DatabaseMigrationConversionWorkspaceRef) GetGVK() schema.GroupVersionKind {
	return DatabaseMigrationConversionWorkspaceGVK
}

func (r *DatabaseMigrationConversionWorkspaceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DatabaseMigrationConversionWorkspaceRef) GetExternal() string {
	return r.External
}

func (r *DatabaseMigrationConversionWorkspaceRef) SetExternal(ref string) {
	r.External = ref
}

func (r *DatabaseMigrationConversionWorkspaceRef) ValidateExternal(ref string) error {
	id := &DatabaseMigrationConversionWorkspaceIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *DatabaseMigrationConversionWorkspaceRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DatabaseMigrationConversionWorkspaceIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DatabaseMigrationConversionWorkspaceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
