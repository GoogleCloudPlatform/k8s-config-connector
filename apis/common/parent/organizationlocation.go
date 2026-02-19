// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package parent

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
)

var _ identity.Identity = &OrganizationLocationParent{}

// OrganizationLocationParent specifies the resource reference to a GCP Organization and Location.
type OrganizationLocationParent struct {
	OrganizationID string
	Location       string
}

func (p *OrganizationLocationParent) String() string {
	return "organizations/" + p.OrganizationID + "/locations/" + p.Location
}

func (p *OrganizationLocationParent) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/")
	if len(tokens) != 4 || tokens[0] != "organizations" || tokens[2] != "locations" {
		return fmt.Errorf("format of OrganizationLocation external=%q was not known (use organizations/<organizationId>/locations/<location>)", ref)
	}
	p.OrganizationID = tokens[1]
	p.Location = tokens[3]
	return nil
}

func (p *OrganizationLocationParent) MatchActual(actualI Parent) error {
	actual, ok := actualI.(*OrganizationLocationParent)
	if !ok {
		return fmt.Errorf("parent format changed, desired %T", p)
	}
	if p.OrganizationID != actual.OrganizationID {
		return fmt.Errorf("spec.organizationRef changed, desired %s, actual %s", p.OrganizationID, actual.OrganizationID)
	}
	if p.Location != actual.Location {
		return fmt.Errorf("spec.location changed, desired %s, actual %s", p.Location, actual.Location)
	}
	return nil
}
