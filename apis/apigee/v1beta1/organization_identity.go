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

const (
	OrganizationIDToken  = "organizations"
	OrganizationIDFormat = OrganizationIDToken + "/{{organizationID}}"
)

// OrganizationIdentity uniquely defines a ApigeeOrganization object.
type OrganizationIdentity struct {
	ResourceID string
}

func (i *OrganizationIdentity) String() string {
	return OrganizationIDToken + "/" + i.ResourceID
}

// NewOrganizationIdentity parses a string-format ApigeeOrganization reference into a OrganizationIdentity object.
func NewOrganizationIdentity(ref string) (*OrganizationIdentity, error) {
	requiredTokens := len(strings.Split(OrganizationIDFormat, "/"))

	tokens := strings.Split(ref, "/")
	if len(tokens) != requiredTokens || tokens[len(tokens)-2] != OrganizationIDToken {
		return nil, fmt.Errorf("format of ApigeeOrganization ref=%q was not known (use %q)", ref, OrganizationIDFormat)
	}

	resourceID := tokens[len(tokens)-1]

	id := &OrganizationIdentity{
		ResourceID: resourceID,
	}

	return id, nil
}

// GetIdentity reads the identity from the ApigeeOrganization resource.
func (obj *ApigeeOrganization) GetIdentity(ctx context.Context, reader client.Reader) (*OrganizationIdentity, error) {
	// Get resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	id := &OrganizationIdentity{
		ResourceID: resourceID,
	}

	// Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		previousID, err := NewOrganizationIdentity(externalRef)
		if err != nil {
			return nil, err
		}
		if id.String() != previousID.String() {
			return nil, fmt.Errorf("cannot update ApigeeOrganization identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
		}
	}

	return id, nil
}
