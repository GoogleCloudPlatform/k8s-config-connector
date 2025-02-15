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

	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// InstanceIdentity defines the resource reference to ApigeeInstance, which "External" field
// holds the GCP identifier for the KRM object.
type InstanceIdentity struct {
	parent *apigeev1beta1.OrganizationIdentity
	id     string
}

func (i *InstanceIdentity) String() string {
	return i.parent.String() + "/instances/" + i.id
}

func (i *InstanceIdentity) ID() string {
	return i.id
}

func (i *InstanceIdentity) Parent() *apigeev1beta1.OrganizationIdentity {
	return i.parent
}

// New builds a InstanceIdentity from the Config Connector Instance object.
func NewApigeeInstanceIdentity(ctx context.Context, reader client.Reader, obj *ApigeeInstance) (*InstanceIdentity, error) {
	// Get Parent
	orgExternal, err := obj.Spec.OrganizationRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	parentOrgID, err := apigeev1beta1.NewOrganizationIdentityFromNormalizedExternal(orgExternal)
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
		actualParent, actualResourceID, err := ParseInstanceExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.String() != parentOrgID.String() {
			return nil, fmt.Errorf("parent organization (spec.organizationRef) changed, expected %s, got %s", actualParent.String(), parentOrgID.String())
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &InstanceIdentity{
		parent: parentOrgID,
		id:     resourceID,
	}, nil
}

// NewInstanceIdentityFromNormalizedExternal builds an InstanceIdentity from
// the normalized string format of an Instance external reference.
func NewInstanceIdentityFromNormalizedExternal(externalRef string) (*InstanceIdentity, error) {
	parent, id, err := ParseInstanceExternal(externalRef)
	if err != nil {
		return nil, err
	}
	instanceID := &InstanceIdentity{
		parent: parent,
		id:     id,
	}
	return instanceID, nil
}

func ParseInstanceExternal(external string) (parent *apigeev1beta1.OrganizationIdentity, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "organizations" || tokens[2] != "instances" {
		return nil, "", fmt.Errorf("format of ApigeeInstance external=%q was not known (use organizations/{{organizationID}}/instances/{{instanceID}})", external)
	}
	parent, err = apigeev1beta1.NewOrganizationIdentityFromNormalizedExternal(strings.Join(tokens[:2], "/"))
	if err != nil {
		return nil, "", err
	}
	resourceID = tokens[3]
	return parent, resourceID, nil
}
