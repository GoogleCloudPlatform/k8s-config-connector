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

// ExternalAddressIdentity defines the resource reference to VMwareEngineExternalAddress, which "External" field
// holds the GCP identifier for the KRM object.
type ExternalAddressIdentity struct {
	parent *ExternalAddressParent
	id     string
}

func (i *ExternalAddressIdentity) String() string {
	return i.parent.String() + "/externalAddresses/" + i.id
}

func (i *ExternalAddressIdentity) ID() string {
	return i.id
}

func (i *ExternalAddressIdentity) Parent() *ExternalAddressParent {
	return i.parent
}

type ExternalAddressParent struct {
	PrivateCloud string
}

func (p *ExternalAddressParent) String() string {
	return p.PrivateCloud
}

// New builds a ExternalAddressIdentity from the Config Connector ExternalAddress object.
func NewExternalAddressIdentity(ctx context.Context, reader client.Reader, obj *VMwareEngineExternalAddress) (*ExternalAddressIdentity, error) {
	// Get Parent
	privateCloud, err := obj.Spec.PrivateCloudRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
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
		actualParent, actualResourceID, err := ParseExternalAddressExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.PrivateCloud != privateCloud {
			return nil, fmt.Errorf("spec.privateCloudRef changed, expect %s, got %s", actualParent.PrivateCloud, privateCloud)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &ExternalAddressIdentity{
		parent: &ExternalAddressParent{
			PrivateCloud: privateCloud,
		},
		id: resourceID,
	}, nil
}

func ParseExternalAddressExternal(external string) (parent *ExternalAddressParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "privateClouds" || tokens[6] != "externalAddresses" {
		return nil, "", fmt.Errorf("format of VMwareEngineExternalAddress external=%q was not known (use projects/{{projectID}}/locations/{{location}}/privateClouds/{{privatecloudID}}/externalAddresses/{{externaladdressID}})", external)
	}
	privateCloud := strings.Join(tokens[:len(tokens)-2], "/")
	parent = &ExternalAddressParent{
		PrivateCloud: privateCloud,
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
