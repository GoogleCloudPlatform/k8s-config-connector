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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AuthorizedViewIdentity defines the resource reference to BigtableAuthorizedView, which "External" field
// holds the GCP identifier for the KRM object.
type AuthorizedViewIdentity struct {
	parent *bigtablev1beta1.TableIdentity
	id     string
}

func (i *AuthorizedViewIdentity) String() string {
	return i.parent.String() + "/authorizedViews/" + i.id
}

func (i *AuthorizedViewIdentity) ID() string {
	return i.id
}

// New builds a AuthorizedViewIdentity from the Config Connector AuthorizedView object.
func NewAuthorizedViewIdentity(ctx context.Context, reader client.Reader, obj *BigtableAuthorizedView) (*AuthorizedViewIdentity, error) {

	// Get Parent
	tableExternal, err := obj.Spec.TableRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	instanceIdentity, tableID, err := bigtablev1beta1.ParseTableExternal(tableExternal)
	if err != nil {
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
		actualParent, actualResourceID, err := ParseAuthorizedViewExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.Id != tableID {
			return nil, fmt.Errorf("spec.groupRef changed, expect %s, got %s", actualParent.Id, tableID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &AuthorizedViewIdentity{
		parent: &bigtablev1beta1.TableIdentity{
			Parent: instanceIdentity,
			Id:     tableID,
		},
		id: resourceID,
	}, nil
}

func ParseAuthorizedViewExternal(external string) (*bigtablev1beta1.TableIdentity, string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "tables" || tokens[6] != "authorizedViews" {
		return nil, "", fmt.Errorf("format of BigtableAuthorizedView external=%q was not known (use projects/{{projectID}}/instances/{{instanceID}}/tables/{{tableID}}/authorizedViews/{{authorizedViewID}})", external)
	}
	p := &bigtablev1beta1.TableIdentity{
		Parent: &bigtablev1beta1.InstanceIdentity{
			Parent: &parent.ProjectParent{
				ProjectID: tokens[1],
			},
			Id: tokens[3],
		},
		Id: tokens[5],
	}
	resourceID := tokens[7]
	return p, resourceID, nil
}
