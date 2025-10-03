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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ApigeeOrganizationIDToken  = "organizations"
	ApigeeOrganizationIDFormat = ApigeeOrganizationIDToken + "/{{organizationID}}"
)

var _ identity.Identity = &ApigeeInstanceIdentity{}

type ApigeeOrganizationIdentity struct {
	ResourceID string
}

func (i *ApigeeOrganizationIdentity) String() string {
	return ApigeeOrganizationIDToken + "/" + i.ResourceID
}

func (i *ApigeeOrganizationIdentity) FromExternal(ref string) error {
	requiredTokens := len(strings.Split(ApigeeOrganizationIDFormat, "/"))

	tokens := strings.Split(ref, "/")
	if len(tokens) != requiredTokens || tokens[len(tokens)-2] != ApigeeOrganizationIDToken {
		return fmt.Errorf("format of ApigeeOrganization ref=%q was not known (use %q)", ref, ApigeeOrganizationIDFormat)
	}

	resourceID := tokens[len(tokens)-1]

	i.ResourceID = resourceID

	return nil
}

var _ identity.Resource = &ApigeeOrganization{}

func (obj *ApigeeOrganization) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	id := &ApigeeOrganizationIdentity{
		ResourceID: resourceID,
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		previousID := &ApigeeOrganizationIdentity{}
		if err := previousID.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if id.String() != previousID.String() {
			return nil, fmt.Errorf("cannot update ApigeeOrganization identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
		}
	}

	return id, nil
}

func (obj *ApigeeOrganization) GetParentIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return nil, nil
}
