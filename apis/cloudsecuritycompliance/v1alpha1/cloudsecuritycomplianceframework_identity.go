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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &CloudSecurityComplianceFrameworkIdentity{}
	_ identity.Resource   = &CloudSecurityComplianceFramework{}
)

var CloudSecurityComplianceFrameworkIdentityFormat = gcpurls.Template[CloudSecurityComplianceFrameworkIdentity]("cloudsecuritycompliance.googleapis.com", "organizations/{organization}/locations/{location}/frameworks/{framework}")

// +k8s:deepcopy-gen=false
type CloudSecurityComplianceFrameworkIdentity struct {
	Organization string
	Location     string
	Framework    string
}

func (i *CloudSecurityComplianceFrameworkIdentity) String() string {
	return CloudSecurityComplianceFrameworkIdentityFormat.ToString(*i)
}

func (i *CloudSecurityComplianceFrameworkIdentity) FromExternal(ref string) error {
	parsed, match, err := CloudSecurityComplianceFrameworkIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CloudSecurityComplianceFramework external=%q was not known (use %s): %w", ref, CloudSecurityComplianceFrameworkIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CloudSecurityComplianceFramework external=%q was not known (use %s)", ref, CloudSecurityComplianceFrameworkIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CloudSecurityComplianceFrameworkIdentity) Host() string {
	return CloudSecurityComplianceFrameworkIdentityFormat.Host()
}

func getIdentityFromCloudSecurityComplianceFrameworkSpec(ctx context.Context, reader client.Reader, obj client.Object) (*CloudSecurityComplianceFrameworkIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	organizationID, err := resolveFrameworkOrganizationID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve organization: %w", err)
	}

	identity := &CloudSecurityComplianceFrameworkIdentity{
		Organization: organizationID,
		Location:     location,
		Framework:    resourceID,
	}
	return identity, nil
}

func resolveFrameworkOrganizationID(ctx context.Context, reader client.Reader, obj client.Object) (string, error) {
	if cc, ok := obj.(*CloudSecurityComplianceFramework); ok {
		if cc.Spec.OrganizationRef != nil {
			org, err := refs.ResolveOrganization(ctx, reader, obj, cc.Spec.OrganizationRef)
			if err != nil {
				return "", err
			}
			return org.OrganizationID, nil
		}
	} else if u, ok := obj.(*unstructured.Unstructured); ok {
		return refs.ResolveOrganizationID(ctx, reader, u)
	}
	return "", fmt.Errorf("organizationRef is required")
}

func (obj *CloudSecurityComplianceFramework) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromCloudSecurityComplianceFrameworkSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &CloudSecurityComplianceFrameworkIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change CloudSecurityComplianceFramework identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
