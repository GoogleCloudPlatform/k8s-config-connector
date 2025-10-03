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

package privilegedaccessmanager

import (
	"fmt"
	"strings"
)

// The Identifier for Config Connector to track the PrivilegedAccessManagerEntitlement resource from the GCP service.
type PrivilegedAccessManagerEntitlementIdentity struct {
	Parent      *parent
	Entitlement string
}

type parent struct {
	Container string
	Location  string
}

func (p *parent) String() string {
	return fmt.Sprintf("%s/locations/%s", p.Container, p.Location)
}

// FullyQualifiedName returns both parent and resource ID in the full url format.
func (c *PrivilegedAccessManagerEntitlementIdentity) FullyQualifiedName() string {
	return fmt.Sprintf("%s/entitlements/%s", c.Parent, c.Entitlement)
}

// AsExternalRef builds a externalRef from a PrivilegedAccessManagerEntitlement.
func (c *PrivilegedAccessManagerEntitlementIdentity) AsExternalRef() *string {
	e := serviceDomain + "/" + c.FullyQualifiedName()
	return &e
}

// asID builds a PrivilegedAccessManagerEntitlementIdentity from a `status.externalRef`
func asID(externalRef string) (*PrivilegedAccessManagerEntitlementIdentity, error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	tokens := strings.Split(path, "/")

	if len(tokens) == 6 &&
		(tokens[0] == "projects" || tokens[0] == "folders" || tokens[0] == "organizations") &&
		tokens[2] == "locations" && tokens[4] == "entitlements" {

		return &PrivilegedAccessManagerEntitlementIdentity{
			Parent: &parent{
				Container: fmt.Sprintf("%s/%s", tokens[0], tokens[1]),
				Location:  tokens[3],
			},
			Entitlement: tokens[5],
		}, nil
	}

	return nil, fmt.Errorf("externalRef should be one of "+
		"%s/projects/<project>/locations/<location>/entitlements/<entitlement>, "+
		"%s/folders/<folder>/locations/<location>/entitlements/<entitlement> and "+
		"%s/organizations/<organization>/locations/<location>/entitlements/<entitlement>,"+
		"got %s", serviceDomain, serviceDomain, serviceDomain, externalRef)
}

// BuildID builds the ID for Config Connector to track the PrivilegedAccessManagerEntitlement resource from the GCP service.
func BuildID(container, location, resourceID string) *PrivilegedAccessManagerEntitlementIdentity {
	return &PrivilegedAccessManagerEntitlementIdentity{
		Parent:      &parent{Container: container, Location: location},
		Entitlement: resourceID,
	}
}
