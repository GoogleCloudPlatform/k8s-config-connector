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
	_ identity.IdentityV2 = &SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity{}
	_ identity.Resource   = &SecurityCenterManagementSecurityHealthAnalyticsCustomModule{}
)

var SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentityFormat = gcpurls.Template[SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity](
	"securitycentermanagement.googleapis.com",
	"projects/{project}/locations/{location}/securityHealthAnalyticsCustomModules/{securitycentermanagementsecurityhealthanalyticscustommodule}",
)

// SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity is the identity of a GCP SecurityCenterManagementSecurityHealthAnalyticsCustomModule resource.
// +k8s:deepcopy-gen=false
type SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity struct {
	Project                                                     string
	Location                                                    string
	SecurityCenterManagementSecurityHealthAnalyticsCustomModule string
}

func (i *SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity) String() string {
	return SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentityFormat.ToString(*i)
}

func (i *SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity) FromExternal(ref string) error {
	parsed, match, err := SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of SecurityCenterManagementSecurityHealthAnalyticsCustomModule external=%q was not known (use %s): %w", ref, SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of SecurityCenterManagementSecurityHealthAnalyticsCustomModule external=%q was not known (use %s)", ref, SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity) Host() string {
	return SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentityFormat.Host()
}

func (i *SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromSecurityCenterManagementSecurityHealthAnalyticsCustomModuleSpec(ctx context.Context, reader client.Reader, obj *SecurityCenterManagementSecurityHealthAnalyticsCustomModule) (*SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	if obj.Spec.Location == nil || *obj.Spec.Location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}
	location := *obj.Spec.Location

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity{
		Project:  projectID,
		Location: location,
		SecurityCenterManagementSecurityHealthAnalyticsCustomModule: resourceID,
	}
	return identity, nil
}

func (obj *SecurityCenterManagementSecurityHealthAnalyticsCustomModule) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromSecurityCenterManagementSecurityHealthAnalyticsCustomModuleSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &SecurityCenterManagementSecurityHealthAnalyticsCustomModuleIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change SecurityCenterManagementSecurityHealthAnalyticsCustomModule identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
