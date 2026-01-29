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
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	// refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// APIGatewayAPIConfigIdentity defines the resource identity for APIGatewayAPIConfig.
type APIGatewayAPIConfigIdentity struct {
	parent *ApiIdentity
	id     string
}

func (i *APIGatewayAPIConfigIdentity) String() string {
	return i.parent.String() + "/configs/" + i.id
}

func (i *APIGatewayAPIConfigIdentity) Parent() *ApiIdentity {
	return i.parent
}

func (i *APIGatewayAPIConfigIdentity) ID() string {
	return i.id
}

// NewAPIGatewayAPIConfigIdentity parses the Identity from the resource URL.
func NewAPIGatewayAPIConfigIdentity(ctx context.Context, reader client.Reader, obj *APIGatewayAPIConfig) (*APIGatewayAPIConfigIdentity, error) {
	// Get Parent
	parentID, err := ResolveAPIGatewayAPI(ctx, reader, obj)
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
		actualParent, actualResourceID, err := ParseAPIGatewayAPIConfigExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.String() != parentID.String() {
			return nil, fmt.Errorf("spec.apiRef changed, expect %s, got %s", actualParent.String(), parentID.String())
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}

	return &APIGatewayAPIConfigIdentity{
		parent: parentID,
		id:     resourceID,
	}, nil
}

func ResolveAPIGatewayAPI(ctx context.Context, reader client.Reader, obj *APIGatewayAPIConfig) (*ApiIdentity, error) {
	if obj.Spec.APIRef != nil {
		// Resolve APIRef
		apiRef := obj.Spec.APIRef
		// Object to target
		target := &APIGatewayAPI{}

		// If external is provided, use it directly
		if apiRef.External != "" {
			parent, _, err := ParseApiExternal(apiRef.External)
			if err != nil {
				return nil, err
			}
			// ParseApiExternal returns *ApiParent and resourceID
			// We need *ApiIdentity
			// ParseApiExternal splits: projects/{p}/locations/{l}/apis/{id}
			// ApiIdentity needs parent *ApiParent and id string
			_, id, _ := ParseApiExternal(apiRef.External)
			return &ApiIdentity{
				parent: parent,
				id:     id,
			}, nil
		}

		// Resolve by Name/Namespace
		nn := client.ObjectKey{
			Name:      apiRef.Name,
			Namespace: apiRef.Namespace,
		}
		if nn.Namespace == "" {
			nn.Namespace = obj.GetNamespace()
		}

		if err := reader.Get(ctx, nn, target); err != nil {
			return nil, err
		}

		// Use the Identity of the target object
		return NewApiIdentity(ctx, reader, target)
	}

	// Fallback? If APIRef is nil, fail.
	return nil, fmt.Errorf("spec.apiRef is required")
}

func ParseAPIGatewayAPIConfigExternal(external string) (parent *ApiIdentity, resourceID string, err error) {
	// projects/{project}/locations/{location}/apis/{api}/configs/{config}
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "apis" || tokens[6] != "configs" {
		return nil, "", fmt.Errorf("format of APIGatewayAPIConfig external=%q was not known (use projects/{{projectID}}/locations/global/apis/{{apiID}}/configs/{{configID}})", external)
	}

	apiParent := &ApiParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	apiID := tokens[5]

	parent = &ApiIdentity{
		parent: apiParent,
		id:     apiID,
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
