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

var _ identity.Identity = &DatabaseIdentity{}

// Database is the parent of BigLakeTable.
var _ parent.Parent = &DatabaseIdentity{}

// DatabaseIdentity defines the resource reference to BigLakeDatabase, which "External" field
// holds the GCP identifier for the KRM object.
type DatabaseIdentity struct {
	parent *CatalogIdentity
	id     string
}

func (i *DatabaseIdentity) String() string {
	return i.parent.String() + "/databases/" + i.id
}

func (i *DatabaseIdentity) ID() string {
	return i.id
}

func (i *DatabaseIdentity) Parent() *CatalogIdentity {
	return i.parent
}

func (i *DatabaseIdentity) URL() string {
	return i.parent.URL() + "/databases/{{databaseID}}"
}

func (i *DatabaseIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/databases/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of BigLakeDatabase external=%q was not known (use %s/databases/{{databaseID}})", ref, i.parent.String())
	}
	parentID := &DatabaseIdentity{}
	if err := parentID.FromExternal(tokens[0]); err != nil {
		return fmt.Errorf("get parent databse URL: %w", err)
	}
	i.id = tokens[1]
	return nil
}

func (i *DatabaseIdentity) MatchActual(actualI parent.Parent) error {
	actual, ok := actualI.(*DatabaseIdentity)
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

// New builds a DatabaseIdentity from the Config Connector Database object.
func NewDatabaseIdentity(ctx context.Context, reader client.Reader, obj *BigLakeDatabase) (*DatabaseIdentity, error) {
	database := &DatabaseIdentity{}

	// Get ID from user-configurable field
	database.id = common.ValueOf(obj.Spec.ResourceID)
	if database.id == "" {
		database.id = obj.GetName()
	}
	if database.id == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Get parent from user-configurable field
	parentExternal, err := obj.Spec.Parent.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, fmt.Errorf("cannot build parent: %w", err)
	}
	if err := database.parent.FromExternal(parentExternal); err != nil {
		return nil, fmt.Errorf("get parent catalog: %w", err)
	}

	// Verify the user-configured id matches the actual ID (reserved in .status.externalRef)
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		idFromExternal := &DatabaseIdentity{}
		if err := idFromExternal.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if err := database.MatchActual(idFromExternal); err != nil {
			return nil, err
		}
	}

	return database, nil
}
