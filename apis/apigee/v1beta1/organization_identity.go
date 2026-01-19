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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	ApigeeOrganizationIdentityFormat = gcpurls.Template[ApigeeOrganizationIdentity]("apigee.googleapis.com", "organizations/{resourceID}")
)

const (
	ApigeeOrganizationIDToken   = "organizations"
	ApigeeOrganizationURLFormat = ApigeeOrganizationIDToken + "/{{organizationID}}"
)

var _ identity.Identity = &ApigeeOrganizationIdentity{}

type ApigeeOrganizationIdentity struct {
	ResourceID string
}

func (i *ApigeeOrganizationIdentity) String() string {
	return ApigeeOrganizationIdentityFormat.ToString(*i)
}

func (i *ApigeeOrganizationIdentity) FromExternal(ref string) error {
	parsed, match, err := ApigeeOrganizationIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ApigeeOrganization external=%q was not known (use %s): %w", ref, ApigeeOrganizationURLFormat, err)
	}
	if !match {
		return fmt.Errorf("format of ApigeeOrganization external=%q was not known (use %s)", ref, ApigeeOrganizationURLFormat)
	}

	*i = *parsed
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
