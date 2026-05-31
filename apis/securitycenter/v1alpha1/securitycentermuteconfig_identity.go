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
	_ identity.IdentityV2 = &SecurityCenterMuteConfigIdentity{}
	_ identity.Resource   = &SecurityCenterMuteConfig{}
)

var SecurityCenterMuteConfigIdentityFormat = gcpurls.Template[SecurityCenterMuteConfigIdentity]("securitycenter.googleapis.com", "organizations/{organization}/muteConfigs/{muteConfig}")

// +k8s:deepcopy-gen=false
type SecurityCenterMuteConfigIdentity struct {
	Organization string
	MuteConfig   string
}

func (i *SecurityCenterMuteConfigIdentity) String() string {
	return SecurityCenterMuteConfigIdentityFormat.ToString(*i)
}

func (i *SecurityCenterMuteConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := SecurityCenterMuteConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of SecurityCenterMuteConfig external=%q was not known (use %s): %w", ref, SecurityCenterMuteConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of SecurityCenterMuteConfig external=%q was not known (use %s)", ref, SecurityCenterMuteConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *SecurityCenterMuteConfigIdentity) Host() string {
	return SecurityCenterMuteConfigIdentityFormat.Host()
}

func getIdentityFromSecurityCenterMuteConfigSpec(ctx context.Context, reader client.Reader, obj client.Object) (*SecurityCenterMuteConfigIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	var organizationID string
	if u, ok := obj.(*unstructured.Unstructured); ok {
		orgID, err := refs.ResolveOrganizationID(ctx, reader, u)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve organization: %w", err)
		}
		organizationID = orgID
	} else if typed, ok := obj.(*SecurityCenterMuteConfig); ok {
		if typed.Spec.OrganizationRef != nil {
			org, err := refs.ResolveOrganization(ctx, reader, typed, typed.Spec.OrganizationRef)
			if err != nil {
				return nil, fmt.Errorf("cannot resolve organization: %w", err)
			}
			if org != nil {
				organizationID = org.OrganizationID
			}
		}
	}
	if organizationID == "" {
		return nil, fmt.Errorf("cannot resolve organization ID")
	}

	identity := &SecurityCenterMuteConfigIdentity{
		Organization: organizationID,
		MuteConfig:   resourceID,
	}
	return identity, nil
}

func (obj *SecurityCenterMuteConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromSecurityCenterMuteConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &SecurityCenterMuteConfigIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change SecurityCenterMuteConfig identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
