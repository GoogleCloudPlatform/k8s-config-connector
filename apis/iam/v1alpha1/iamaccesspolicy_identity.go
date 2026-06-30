// Copyright 2026 Google LLC
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &IamAccessPolicyIdentity{}
	_ identity.Resource   = &IamAccessPolicy{}
)

// IamAccessPolicyIdentityFormat matches organizations/{organization}/locations/{location}/accessPolicies/{accessPolicy}
var IamAccessPolicyIdentityFormat = gcpurls.Template[IamAccessPolicyIdentity]("accesscontextmanager.googleapis.com", "organizations/{organization}/locations/{location}/accessPolicies/{accessPolicy}")

// +k8s:deepcopy-gen=false
type IamAccessPolicyIdentity struct {
	Organization string
	Location     string
	AccessPolicy string
}

func (i *IamAccessPolicyIdentity) String() string {
	return IamAccessPolicyIdentityFormat.ToString(*i)
}

func (i *IamAccessPolicyIdentity) FromExternal(ref string) error {
	parsed, match, err := IamAccessPolicyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of IamAccessPolicy external=%q was not known (use %s): %w", ref, IamAccessPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of IamAccessPolicy external=%q was not known (use %s)", ref, IamAccessPolicyIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *IamAccessPolicyIdentity) Host() string {
	return IamAccessPolicyIdentityFormat.Host()
}

func getIdentityFromIamAccessPolicySpec(ctx context.Context, reader client.Reader, obj *IamAccessPolicy) (*IamAccessPolicyIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	var organizationID string
	if obj.Spec.OrganizationRef == nil {
		return nil, fmt.Errorf("organizationRef must be set")
	}

	org, err := refsv1beta1.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)
	if err != nil {
		return nil, err
	}
	organizationID = org.OrganizationID

	if obj.Spec.Location == "" {
		return nil, fmt.Errorf("location must be set")
	}

	identity := &IamAccessPolicyIdentity{
		Organization: organizationID,
		Location:     obj.Spec.Location,
		AccessPolicy: resourceID,
	}
	return identity, nil
}

func (obj *IamAccessPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromIamAccessPolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &IamAccessPolicyIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change IamAccessPolicy identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
