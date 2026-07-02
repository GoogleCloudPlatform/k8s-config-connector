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
	_ identity.IdentityV2 = &MarketingPlatformAdminAnalyticsAccountLinkIdentity{}
	_ identity.Resource   = &MarketingPlatformAdminAnalyticsAccountLink{}
)

var MarketingPlatformAdminAnalyticsAccountLinkIdentityFormat = gcpurls.Template[MarketingPlatformAdminAnalyticsAccountLinkIdentity]("marketingplatformadmin.googleapis.com", "organizations/{organization}/analyticsAccountLinks/{analytics_account_link}")

// +k8s:deepcopy-gen=false
type MarketingPlatformAdminAnalyticsAccountLinkIdentity struct {
	Organization         string
	AnalyticsAccountLink string
}

func (i *MarketingPlatformAdminAnalyticsAccountLinkIdentity) String() string {
	return MarketingPlatformAdminAnalyticsAccountLinkIdentityFormat.ToString(*i)
}

func (i *MarketingPlatformAdminAnalyticsAccountLinkIdentity) FromExternal(ref string) error {
	parsed, match, err := MarketingPlatformAdminAnalyticsAccountLinkIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MarketingPlatformAdminAnalyticsAccountLink external=%q was not known (use %s): %w", ref, MarketingPlatformAdminAnalyticsAccountLinkIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MarketingPlatformAdminAnalyticsAccountLink external=%q was not known (use %s)", ref, MarketingPlatformAdminAnalyticsAccountLinkIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MarketingPlatformAdminAnalyticsAccountLinkIdentity) Host() string {
	return MarketingPlatformAdminAnalyticsAccountLinkIdentityFormat.Host()
}

func getIdentityFromMarketingPlatformAdminAnalyticsAccountLinkSpec(ctx context.Context, reader client.Reader, obj *MarketingPlatformAdminAnalyticsAccountLink) (*MarketingPlatformAdminAnalyticsAccountLinkIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	var organizationID string
	if obj.Spec.OrganizationRef != nil {
		org, err := refs.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve organization: %w", err)
		}
		if org != nil {
			organizationID = org.OrganizationID
		}
	}
	if organizationID == "" {
		return nil, fmt.Errorf("cannot resolve organization ID")
	}

	identity := &MarketingPlatformAdminAnalyticsAccountLinkIdentity{
		Organization:         organizationID,
		AnalyticsAccountLink: resourceID,
	}
	return identity, nil
}

func (obj *MarketingPlatformAdminAnalyticsAccountLink) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromMarketingPlatformAdminAnalyticsAccountLinkSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &MarketingPlatformAdminAnalyticsAccountLinkIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change MarketingPlatformAdminAnalyticsAccountLink identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
