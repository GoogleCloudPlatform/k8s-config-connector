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

	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AuthorizedViewIdentity defines the resource reference to BigtableAuthorizedView, which "External" field
// holds the GCP identifier for the KRM object.
type AuthorizedViewIdentity struct {
	parent *bigtablev1beta1.BigtableTableIdentity
	id     string
}

func (i *AuthorizedViewIdentity) String() string {
	return i.parent.String() + "/authorizedViews/" + i.id
}

func (i *AuthorizedViewIdentity) ID() string {
	return i.id
}

func (i *AuthorizedViewIdentity) FromExternal(external string) error {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "tables" || tokens[6] != "authorizedViews" {
		return fmt.Errorf("format of BigtableAuthorizedView external=%q was not known (use projects/{{projectID}}/instances/{{instanceID}}/tables/{{tableID}}/authorizedViews/{{authorizedViewID}})", external)
	}
	i.parent = &bigtablev1beta1.BigtableTableIdentity{
		Project:  tokens[1],
		Instance: tokens[3],
		Table:    tokens[5],
	}
	i.id = tokens[7]
	return nil
}

// New builds a AuthorizedViewIdentity from the Config Connector AuthorizedView object.
func NewAuthorizedViewIdentity(ctx context.Context, reader client.Reader, obj *BigtableAuthorizedView) (*AuthorizedViewIdentity, error) {
	// Get Parent
	if err := obj.Spec.TableRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	tableExternal := obj.Spec.TableRef.External
	tableIdentity := &bigtablev1beta1.BigtableTableIdentity{}
	if err := tableIdentity.FromExternal(tableExternal); err != nil {
		return nil, err
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &AuthorizedViewIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if statusIdentity.parent.Project != tableIdentity.Project {
			return nil, fmt.Errorf("spec.tableRef ProjectID changed, expect %s, got %s", statusIdentity.parent.Project, tableIdentity.Project)
		}
		if statusIdentity.parent.Instance != tableIdentity.Instance {
			return nil, fmt.Errorf("spec.tableRef InstanceID changed, expect %s, got %s", statusIdentity.parent.Instance, tableIdentity.Instance)
		}
		if statusIdentity.parent.Table != tableIdentity.Table {
			return nil, fmt.Errorf("spec.tableRef tableID changed, expect %s, got %s", statusIdentity.parent.Table, tableIdentity.Table)
		}
		if statusIdentity.id != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, statusIdentity.id)
		}
	}
	return &AuthorizedViewIdentity{
		parent: tableIdentity,
		id:     resourceID,
	}, nil
}

func ParseAuthorizedViewExternal(external string) (*bigtablev1beta1.BigtableTableIdentity, string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "tables" || tokens[6] != "authorizedViews" {
		return nil, "", fmt.Errorf("format of BigtableAuthorizedView external=%q was not known (use projects/{{projectID}}/instances/{{instanceID}}/tables/{{tableID}}/authorizedViews/{{authorizedViewID}})", external)
	}
	p := &bigtablev1beta1.BigtableTableIdentity{
		Project:  tokens[1],
		Instance: tokens[3],
		Table:    tokens[5],
	}
	resourceID := tokens[7]
	return p, resourceID, nil
}
