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

var _ identity.Identity = &TableIdentity{}

// TableIdentity defines the resource reference to BigLakeTable, which "External" field
// holds the GCP identifier for the KRM object.
type TableIdentity struct {
	parent *DatabaseIdentity
	id     string
}

func (i *TableIdentity) String() string {
	return i.parent.String() + "/tables/" + i.id
}

func (i *TableIdentity) ID() string {
	return i.id
}

func (i *TableIdentity) Parent() *DatabaseIdentity {
	return i.parent
}

func (i *TableIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/tables/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of BigLakeTable external=%q was not known (use {{databaseURL}}/tables/{{tableID}})", i.parent.URL())
	}
	parentID := &DatabaseIdentity{}
	if err := parentID.FromExternal(tokens[0]); err != nil {
		return fmt.Errorf("get parent databse URL: %w", err)
	}
	i.id = tokens[1]
	return nil
}

func (i *TableIdentity) MatchActual(actualI parent.Parent) error {
	actual, ok := actualI.(*TableIdentity)
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

// New builds a TableIdentity from the Config Connector Table object.
func NewTableIdentity(ctx context.Context, reader client.Reader, obj *BigLakeTable) (*TableIdentity, error) {
	table := &TableIdentity{}

	// Get ID from user-configurable field
	table.id = common.ValueOf(obj.Spec.ResourceID)
	if table.id == "" {
		table.id = obj.GetName()
	}
	if table.id == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Get parent from user-configurable field
	parentExternal, err := obj.Spec.Parent.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, fmt.Errorf("cannot build parent: %w", err)
	}
	if err := table.parent.FromExternal(parentExternal); err != nil {
		return nil, fmt.Errorf("get parent catalog: %w", err)
	}

	// Verify the user-configured id matches the actual ID (reserved in .status.externalRef)
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		idFromExternal := &DatabaseIdentity{}
		if err := idFromExternal.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if err := table.MatchActual(idFromExternal); err != nil {
			return nil, err
		}
	}

	return table, nil
}
