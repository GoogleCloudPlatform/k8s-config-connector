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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CustomConstraintIdentity defines the resource reference to OrgpolicyCustomConstraint, which "External" field
// holds the GCP identifier for the KRM object.
type CustomConstraintIdentity struct {
	parent *CustomConstraintParent
	id     string
}

func (i *CustomConstraintIdentity) String() string {
	return i.parent.String() + "/customconstraints/" + i.id
}

func (i *CustomConstraintIdentity) ID() string {
	return i.id
}

func (i *CustomConstraintIdentity) Parent() *CustomConstraintParent {
	return i.parent
}

type CustomConstraintParent struct {
	OrganizationID string
}

func (p *CustomConstraintParent) String() string {
	return "organizations/" + p.OrganizationID
}

// New builds a CustomConstraintIdentity from the Config Connector CustomConstraint object.
func NewCustomConstraintIdentity(ctx context.Context, reader client.Reader, obj *OrgpolicyCustomConstraint) (*CustomConstraintIdentity, error) {

	// Get Parent
	organizationRef, err := refsv1beta1.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)
	if err != nil {
		return nil, err
	}
	organizationID := organizationRef.OrganizationID
	if organizationID == "" {
		return nil, fmt.Errorf("cannot resolve organization")
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
		actualParent, actualResourceID, err := ParseCustomConstraintExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.OrganizationID != organizationID {
			return nil, fmt.Errorf("spec.organizationRef changed, expect %s, got %s", actualParent.OrganizationID, organizationID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &CustomConstraintIdentity{
		parent: &CustomConstraintParent{
			OrganizationID: organizationID,
		},
		id: resourceID,
	}, nil
}

func ParseCustomConstraintExternal(external string) (parent *CustomConstraintParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "organizations" || tokens[2] != "customconstraints" {
		return nil, "", fmt.Errorf("format of OrgpolicyCustomConstraint external=%q was not known (use organizations/{{organizationID}}/customconstraints/{{customconstraintID}})", external)
	}
	parent = &CustomConstraintParent{
		OrganizationID: tokens[1],
	}
	resourceID = tokens[3]
	return parent, resourceID, nil
}
