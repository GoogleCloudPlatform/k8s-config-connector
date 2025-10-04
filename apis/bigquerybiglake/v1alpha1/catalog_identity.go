// Copyright 2025 Google LLC
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
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &CatalogIdentity{}

// Catalog is the parent of BigLakeDatabase.
var _ parent.Parent = &CatalogIdentity{}

// CatalogIdentity defines the resource reference to BigLakeCatalog, which "External" field
// holds the GCP identifier for the KRM object.
type CatalogIdentity struct {
	parent *parent.ProjectAndLocation
	id     string
}

func (i *CatalogIdentity) String() string {
	return i.parent.String() + "/catalogs/" + i.id
}

func (i *CatalogIdentity) URL() string {
	return i.parent.URL() + "/catalogs/{{catalogID}}"
}

func (i *CatalogIdentity) ID() string {
	return i.id
}

func (i *CatalogIdentity) Parent() *parent.ProjectAndLocation {
	return i.parent
}

func (i *CatalogIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/catalogs/")

	if len(tokens) != 2 {
		return fmt.Errorf("format of BigLakeCatalog external=%q was not known (use %s/catalogs/{{catalogID}})", i.parent.URL())
	}
	i.parent = &parent.ProjectAndLocation{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return fmt.Errorf("get parent project and location: %w", err)
	}
	i.id = tokens[1]
	return nil
}

func (i *CatalogIdentity) MatchActual(actualI parent.Parent) error {
	actual, ok := actualI.(*CatalogIdentity)
	if !ok {
		return fmt.Errorf("identity format changed, desired %T", i)
	}
	if i.id != actual.id {
		return fmt.Errorf("spec.resourceID changed, desired %s, actual %s", i.id, actual.id)
	}
	if err := i.parent.MatchActual(actual.parent); err != nil {
		return err
	}
	return nil
}

// New builds a CatalogIdentity from the Config Connector Catalog object.
func NewCatalogIdentity(ctx context.Context, reader client.Reader, obj *BigLakeCatalog) (*CatalogIdentity, error) {
	catalog := &CatalogIdentity{}

	// Get ID from user-configurable field
	catalog.id = common.ValueOf(obj.Spec.ResourceID)
	if catalog.id == "" {
		catalog.id = obj.GetName()
	}
	if catalog.id == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Get parent from user-configurable field
	parentExternalRef, err := obj.Spec.Parent.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, fmt.Errorf("cannot build parent: %w", err)
	}
	if err := catalog.parent.FromExternal(parentExternalRef); err != nil {
		return nil, fmt.Errorf("get parent project and location: %w", err)
	}

	// Verify the user-configured id matches the actual ID (reserved in .status.externalRef)
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		idFromExternal := &CatalogIdentity{}
		if err := idFromExternal.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if err := catalog.MatchActual(idFromExternal); err != nil {
			return nil, err
		}
	}

	return catalog, nil
}
