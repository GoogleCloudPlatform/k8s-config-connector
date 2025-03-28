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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// EnvironmentIdentity defines the resource reference to DataplexEnvironment, which "External" field
// holds the GCP identifier for the KRM object.
type EnvironmentIdentity struct {
	parent *LakeIdentity
	id     string
}

func (i *EnvironmentIdentity) String() string {
	return i.parent.String() + "/environments/" + i.id
}

func (i *EnvironmentIdentity) ID() string {
	return i.id
}

// New builds a EnvironmentIdentity from the Config Connector Environment object.
func NewEnvironmentIdentity(ctx context.Context, reader client.Reader, obj *DataplexEnvironment) (*EnvironmentIdentity, error) {

	// Get Parent
	lakeExternal, err := obj.Spec.LakeRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}

	parent, lakeId, err := ParseLakeExternal(lakeExternal)
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
		actualParent, actualResourceID, err := ParseEnvironmentExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.parent != parent {
			return nil, fmt.Errorf("parent changed, expect %s, got %s", actualParent.parent, parent)
		}
		if actualParent.id != lakeId {
			return nil, fmt.Errorf("spec.lakeRef changed, expect %s, got %s", actualParent.id, lakeId)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &EnvironmentIdentity{
		parent: &LakeIdentity{
			parent: parent,
			id:     lakeId,
		},
		id: resourceID,
	}, nil
}

func ParseEnvironmentExternal(external string) (parent *LakeIdentity, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "lakes" || tokens[6] != "environments" {
		return nil, "", fmt.Errorf("format of DataplexContent external=%q was not known (use projects/{{projectID}}/locations/{{location}}/lakes/{{lakeID}}/environments/{{environmentsID}})", external)
	}
	parent = &LakeIdentity{
		parent: &LakeParent{
			ProjectID: tokens[1],
			Location:  tokens[3],
		},
		id: tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
