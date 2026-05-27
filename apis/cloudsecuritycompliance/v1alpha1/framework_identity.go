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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &CloudSecurityFrameworkIdentity{}
	_ identity.Resource   = &CloudSecurityFramework{}
)

var CloudSecurityFrameworkIdentityFormat = gcpurls.Template[CloudSecurityFrameworkIdentity]("cloudsecuritycompliance.googleapis.com", "organizations/{organization}/locations/{location}/frameworks/{framework}")

// +k8s:deepcopy-gen=false
type CloudSecurityFrameworkIdentity struct {
	Organization string
	Location     string
	Framework    string
}

func (i *CloudSecurityFrameworkIdentity) String() string {
	return CloudSecurityFrameworkIdentityFormat.ToString(*i)
}

func (i *CloudSecurityFrameworkIdentity) FromExternal(ref string) error {
	parsed, match, err := CloudSecurityFrameworkIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CloudSecurityFramework external=%q was not known (use %s): %w", ref, CloudSecurityFrameworkIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CloudSecurityFramework external=%q was not known (use %s)", ref, CloudSecurityFrameworkIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CloudSecurityFrameworkIdentity) Host() string {
	return CloudSecurityFrameworkIdentityFormat.Host()
}

func getIdentityFromCloudSecurityFrameworkSpec(ctx context.Context, reader client.Reader, obj client.Object) (*CloudSecurityFrameworkIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	// For OrganizationRef
	var organization string
	frameworkObj := obj.(*CloudSecurityFramework)
	if frameworkObj.Spec.OrganizationRef != nil {
		orgIdentity, err := refs.ResolveOrganization(ctx, reader, frameworkObj, frameworkObj.Spec.OrganizationRef)
		if err != nil {
			return nil, err
		}
		organization = orgIdentity.OrganizationID
	} else {
		return nil, fmt.Errorf("cannot resolve organization: spec.organizationRef must be provided")
	}

	identity := &CloudSecurityFrameworkIdentity{
		Organization: organization,
		Location:     location,
		Framework:    resourceID,
	}
	return identity, nil
}

func (obj *CloudSecurityFramework) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCloudSecurityFrameworkSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &CloudSecurityFrameworkIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change CloudSecurityFramework identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
