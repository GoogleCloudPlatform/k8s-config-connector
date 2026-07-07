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
	_ identity.IdentityV2 = &SecurityCenterManagementEventThreatDetectionCustomModuleIdentity{}
	_ identity.Resource   = &SecurityCenterManagementEventThreatDetectionCustomModule{}
)

var SecurityCenterManagementEventThreatDetectionCustomModuleIdentityFormat = gcpurls.Template[SecurityCenterManagementEventThreatDetectionCustomModuleIdentity]("securitycentermanagement.googleapis.com", "organizations/{organization}/locations/{location}/eventThreatDetectionCustomModules/{eventThreatDetectionCustomModule}")

// +k8s:deepcopy-gen=false
type SecurityCenterManagementEventThreatDetectionCustomModuleIdentity struct {
	Organization                     string
	Location                         string
	EventThreatDetectionCustomModule string
}

func (i *SecurityCenterManagementEventThreatDetectionCustomModuleIdentity) String() string {
	return SecurityCenterManagementEventThreatDetectionCustomModuleIdentityFormat.ToString(*i)
}

func (i *SecurityCenterManagementEventThreatDetectionCustomModuleIdentity) FromExternal(ref string) error {
	parsed, match, err := SecurityCenterManagementEventThreatDetectionCustomModuleIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of SecurityCenterManagementEventThreatDetectionCustomModule external=%q was not known (use %s): %w", ref, SecurityCenterManagementEventThreatDetectionCustomModuleIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of SecurityCenterManagementEventThreatDetectionCustomModule external=%q was not known (use %s)", ref, SecurityCenterManagementEventThreatDetectionCustomModuleIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *SecurityCenterManagementEventThreatDetectionCustomModuleIdentity) ParentString() string {
	return fmt.Sprintf("organizations/%s/locations/%s", i.Organization, i.Location)
}

func (i *SecurityCenterManagementEventThreatDetectionCustomModuleIdentity) Host() string {
	return SecurityCenterManagementEventThreatDetectionCustomModuleIdentityFormat.Host()
}

func getIdentityFromSecurityCenterManagementEventThreatDetectionCustomModuleSpec(ctx context.Context, reader client.Reader, obj *SecurityCenterManagementEventThreatDetectionCustomModule) (*SecurityCenterManagementEventThreatDetectionCustomModuleIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
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

	if obj.Spec.Location == nil || *obj.Spec.Location == "" {
		return nil, fmt.Errorf("spec.location is required")
	}

	identity := &SecurityCenterManagementEventThreatDetectionCustomModuleIdentity{
		Organization:                     organizationID,
		Location:                         *obj.Spec.Location,
		EventThreatDetectionCustomModule: resourceID,
	}
	return identity, nil
}

func (obj *SecurityCenterManagementEventThreatDetectionCustomModule) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromSecurityCenterManagementEventThreatDetectionCustomModuleSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &SecurityCenterManagementEventThreatDetectionCustomModuleIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change SecurityCenterManagementEventThreatDetectionCustomModule identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *SecurityCenterManagementEventThreatDetectionCustomModule) ExternalIdentifier() string {
	if obj.Status.ExternalRef != nil {
		return *obj.Status.ExternalRef
	}
	return ""
}
