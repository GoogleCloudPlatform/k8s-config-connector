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
	Project      string
	Folder       string
	Organization string
	Location     string
}

func (p *parent) String() string {
	if p.Project != "" {
		return fmt.Sprintf("projects/%s/locations/%s", p.Project, p.Location)
	} else if p.Folder != "" {
		return fmt.Sprintf("folders/%s/locations/%s", p.Folder, p.Location)
	} else {
		return fmt.Sprintf("organizations/%s/locations/%s", p.Organization, p.Location)
	}
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

	if len(tokens) != 6 ||
		!(tokens[0] == "projects" || tokens[0] == "folders" || tokens[0] == "organizations") ||
		tokens[2] != "locations" || tokens[4] != "entitlements" {
		return nil, fmt.Errorf("externalRef should be one of "+
			"%s/projects/<project>/locations/<location>/entitlements/<entitlement>, "+
			"%s/folders/<folder>/locations/<location>/entitlements/<entitlement> and "+
			"%s/organizations/<organization>/locations/<location>/entitlements/<entitlement>,"+
			"got %s", serviceDomain, serviceDomain, serviceDomain, externalRef)
	}
	p := parent{Location: tokens[3]}
	if tokens[0] == "projects" {
		p.Project = tokens[1]
	} else if tokens[0] == "folders" {
		p.Folder = tokens[1]
	} else if tokens[0] == "organizations" {
		p.Organization = tokens[1]
	}
	return &PrivilegedAccessManagerEntitlementIdentity{
		Parent:      &p,
		Entitlement: tokens[5],
	}, nil
}

// BuildID builds the ID for Config Connector to track the PrivilegedAccessManagerEntitlement resource from the GCP service.
func BuildID(project, folder, organization, location, resourceID string) *PrivilegedAccessManagerEntitlementIdentity {
	return &PrivilegedAccessManagerEntitlementIdentity{
		Parent:      &parent{Project: project, Folder: folder, Organization: organization, Location: location},
		Entitlement: resourceID,
	}
}
