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

// GoogleCloudApigeeV1EnvironmentGroupIdentity defines the resource reference to ApigeeEnvgroup, which "External" field
// holds the GCP identifier for the KRM object.
type GoogleCloudApigeeV1EnvironmentGroupIdentity struct {
	parent *GoogleCloudApigeeV1EnvironmentGroupParent
	id     string
}

func (i *GoogleCloudApigeeV1EnvironmentGroupIdentity) String() string {
	return fmt.Sprintf("%s/envgroups/%s", i.parent.String(), i.id)
}

func (i *GoogleCloudApigeeV1EnvironmentGroupIdentity) ID() string {
	return i.id
}

func (i *GoogleCloudApigeeV1EnvironmentGroupIdentity) Parent() *GoogleCloudApigeeV1EnvironmentGroupParent {
	return i.parent
}

type GoogleCloudApigeeV1EnvironmentGroupParent struct {
	Organization string
}

func (p *GoogleCloudApigeeV1EnvironmentGroupParent) String() string {
	return "organizations/" + p.Organization
}

// New builds a GoogleCloudApigeeV1EnvironmentGroupIdentity from the Config Connector GoogleCloudApigeeV1EnvironmentGroup object.
func NewGoogleCloudApigeeV1EnvironmentGroupIdentity(ctx context.Context, reader client.Reader, obj *ApigeeEnvgroup) (*GoogleCloudApigeeV1EnvironmentGroupIdentity, error) {
	// Get Parent
	orgRef, err := refs.ResolveOrganization(ctx, reader, obj, obj.Spec.Parent.OrganizationRef)
	if err != nil {
		return nil, err
	}
	orgName := orgRef.OrganizationName
	if orgName == "" {
		return nil, fmt.Errorf("cannot resolve organization")
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
		actualParent, actualResourceID, err := ParseGoogleCloudApigeeV1EnvironmentGroupExternal(externalRef)
		if err != nil {
			return nil, err
		}

		if actualParent.Organization != orgName {
			return nil, fmt.Errorf("ApigeeEnvgroup %s/%s has Spec.Parent.OrganizationRef changed, expect %s, got %s",
				obj.GetNamespace(), obj.GetName(), actualParent.Organization, orgRef.OrganizationName)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("ApigeeEnvgroup %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				obj.GetNamespace(), obj.GetName(), actualResourceID, resourceID)
		}
	}

	return &GoogleCloudApigeeV1EnvironmentGroupIdentity{
		parent: &GoogleCloudApigeeV1EnvironmentGroupParent{
			Organization: orgName,
		},
		id: resourceID,
	}, nil
}

func ParseGoogleCloudApigeeV1EnvironmentGroupExternal(external string) (parent *GoogleCloudApigeeV1EnvironmentGroupParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "organizations" || tokens[2] != "envgroups" {
		return nil, "", fmt.Errorf("format of ApigeeEnvgroup external=%q was not known (use organizations/{{organization}}/envgroups/{{envgroup}})",
			external)
	}
	parent = &GoogleCloudApigeeV1EnvironmentGroupParent{
		Organization: tokens[1],
	}
	resourceID = tokens[3]
	return parent, resourceID, nil
}
