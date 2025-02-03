// Copyright 2024 Google LLC
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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// EnvironmentGroupIdentity defines the resource reference to ApigeeEnvgroup, which "External" field
// holds the GCP identifier for the KRM object.
type EnvironmentGroupIdentity struct {
	parent *EnvironmentGroupParent
	id     string
}

func (i *EnvironmentGroupIdentity) String() string {
	return fmt.Sprintf("%s/envgroups/%s", i.parent, i.id)
}

func (i *EnvironmentGroupIdentity) ID() string {
	return i.id
}

func (i *EnvironmentGroupIdentity) Parent() *EnvironmentGroupParent {
	return i.parent
}

type EnvironmentGroupParent struct {
	Organization string
}

func (p *EnvironmentGroupParent) String() string {
	return "organizations/" + p.Organization
}

// New builds a NewEnvironmentGroupIdentity from the Config Connector ApigeeEnvgroup object.
func NewEnvironmentGroupIdentity(ctx context.Context, reader client.Reader, obj *ApigeeEnvgroup) (*EnvironmentGroupIdentity, error) {
	// Get Parent
	orgRef := obj.Spec.Parent.OrganizationRef
	if orgRef == nil {
		return nil, fmt.Errorf("no parent organization")
	}

	orgExternal, err := orgRef.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve organization: %w", err)
	}

	org, err := refs.ParseApigeeOrganizationExternal(orgExternal)
	if err != nil {
		return nil, fmt.Errorf("cannot parse external organization: %w", err)
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseEnvironmentGroupExternal(externalRef)
		if err != nil {
			return nil, err
		}

		if actualParent.Organization != org {
			return nil, fmt.Errorf("ApigeeEnvgroup %s/%s has Spec.Parent.OrganizationRef changed, expect %s, got %s",
				obj.GetNamespace(), obj.GetName(), actualParent.Organization, org)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("ApigeeEnvgroup %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				obj.GetNamespace(), obj.GetName(), actualResourceID, resourceID)
		}
	}

	return &EnvironmentGroupIdentity{
		parent: &EnvironmentGroupParent{
			Organization: org,
		},
		id: resourceID,
	}, nil
}

func ParseEnvironmentGroupExternal(external string) (parent *EnvironmentGroupParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "organizations" || tokens[2] != "envgroups" {
		return nil, "", fmt.Errorf("format of ApigeeEnvgroup external=%q was not known (use organizations/{{organization}}/envgroups/{{envgroup}})",
			external)
	}
	parent = &EnvironmentGroupParent{
		Organization: tokens[1],
	}
	resourceID = tokens[3]
	return parent, resourceID, nil
}
