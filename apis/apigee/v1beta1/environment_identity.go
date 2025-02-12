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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// EnvironmentIdentity defines the resource reference to ApigeeEnvironment, which "External" field
// holds the GCP identifier for the KRM object.
type EnvironmentIdentity struct {
	parent *OrganizationIdentity
	id     string
}

func (i *EnvironmentIdentity) String() string {
	return i.parent.String() + "/environments/" + i.id
}

func (i *EnvironmentIdentity) ID() string {
	return i.id
}

func (i *EnvironmentIdentity) Parent() *OrganizationIdentity {
	return i.parent
}

// New builds a EnvironmentIdentity from the Config Connector Environment object.
func NewEnvironmentIdentity(ctx context.Context, reader client.Reader, obj *ApigeeEnvironment) (*EnvironmentIdentity, error) {
	// Get Parent
	orgExternal, err := obj.Spec.ApigeeOrganizationRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	parentOrgID, err := NewOrganizationIdentityFromNormalizedExternal(orgExternal)
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
		if actualParent.String() != parentOrgID.String() {
			return nil, fmt.Errorf("parent organization (spec.apigeeOrganizationRef) changed, expected %s, got %s", actualParent.String(), parentOrgID.String())
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &EnvironmentIdentity{
		parent: parentOrgID,
		id:     resourceID,
	}, nil
}

func ParseEnvironmentExternal(external string) (parent *OrganizationIdentity, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "organizations" || tokens[2] != "environments" {
		return nil, "", fmt.Errorf("format of ApigeeEnvironment external=%q was not known (use organizations/{{organizationID}}/environments/{{environmentID}})", external)
	}
	parent, err = NewOrganizationIdentityFromNormalizedExternal(strings.Join(tokens[:2], "/"))
	if err != nil {
		return nil, "", err
	}
	resourceID = tokens[3]
	return parent, resourceID, nil
}
