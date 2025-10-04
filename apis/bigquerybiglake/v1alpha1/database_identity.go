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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &DatabaseIdentity{}

const (
	DatabaseIDTokens = "databases"
	DatabaseIDURL    = CatalogIDURL + "/databases/{{databaseID}}"
)

// DatabaseIdentity defines the resource reference to BigLakeDatabase, which "External" field
// holds the GCP identifier for the KRM object.
type DatabaseIdentity struct {
	parent *CatalogIdentity
	id     string
}

func (i *DatabaseIdentity) String() string {
	return i.parent.String() + "/" + DatabaseIDTokens + "/" + i.id
}

func (i *DatabaseIdentity) ID() string {
	return i.id
}

func (i *DatabaseIdentity) Parent() *CatalogIdentity {
	return i.parent
}

func (i *DatabaseIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/"+DatabaseIDTokens+"/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of BigLakeDatabase external=%q was not known (use %s)", ref, DatabaseIDURL)
	}
	i.parent = &CatalogIdentity{}
	if err := i.parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	i.id = tokens[1]
	if i.id == "" {
		return fmt.Errorf("databaseID was empty in external=%q", ref)
	}
	return nil
}

// New builds a DatabaseIdentity from the Config Connector Database object.
func NewDatabaseIdentity(ctx context.Context, reader client.Reader, obj *BigLakeDatabase) (*DatabaseIdentity, error) {
	newIdentity := &DatabaseIdentity{}

	// Resolve Parent
	if err := obj.Spec.ParentRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.parentRef: %w", err)
	}
	newIdentity.parent = &CatalogIdentity{}
	if err := newIdentity.parent.FromExternal(obj.Spec.ParentRef.GetExternal()); err != nil {
		return nil, fmt.Errorf("parsing parentRef.external=%q: %w", obj.Spec.ParentRef.GetExternal(), err)
	}

	// Get desired ID
	newIdentity.id = common.ValueOf(obj.Spec.ResourceID)
	if newIdentity.id == "" {
		newIdentity.id = obj.GetName()
	}
	if newIdentity.id == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Validate against preserved ID, if any
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		preserved := &DatabaseIdentity{}
		if err := preserved.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if preserved.String() != newIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, newIdentity.String())
		}
	}
	return newIdentity, nil
}
