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
	_ identity.IdentityV2 = &SecurityCenterBigQueryExportIdentity{}
	_ identity.Resource   = &SecurityCenterBigQueryExport{}
)

var SecurityCenterBigQueryExportIdentityFormat = gcpurls.Template[SecurityCenterBigQueryExportIdentity]("securitycenter.googleapis.com", "organizations/{organization}/locations/{location}/bigQueryExports/{export}")

// +k8s:deepcopy-gen=false
type SecurityCenterBigQueryExportIdentity struct {
	Organization string
	Location     string
	Export       string
}

func (i *SecurityCenterBigQueryExportIdentity) String() string {
	return SecurityCenterBigQueryExportIdentityFormat.ToString(*i)
}

func (i *SecurityCenterBigQueryExportIdentity) FromExternal(ref string) error {
	parsed, match, err := SecurityCenterBigQueryExportIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of SecurityCenterBigQueryExport external=%q was not known (use %s): %w", ref, SecurityCenterBigQueryExportIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of SecurityCenterBigQueryExport external=%q was not known (use %s)", ref, SecurityCenterBigQueryExportIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *SecurityCenterBigQueryExportIdentity) Host() string {
	return SecurityCenterBigQueryExportIdentityFormat.Host()
}

func getIdentityFromSecurityCenterBigQueryExportSpec(ctx context.Context, reader client.Reader, obj client.Object) (*SecurityCenterBigQueryExportIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	// Resolve Organization
	var organizationID string
	if u, ok := obj.(*unstructured.Unstructured); ok {
		orgID, err := refs.ResolveOrganizationID(ctx, reader, u)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve organization: %w", err)
		}
		organizationID = orgID
	} else if typed, ok := obj.(*SecurityCenterBigQueryExport); ok {
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
		return nil, fmt.Errorf("cannot resolve organization")
	}

	identity := &SecurityCenterBigQueryExportIdentity{
		Organization: organizationID,
		Location:     location,
		Export:       resourceID,
	}
	return identity, nil
}

func (obj *SecurityCenterBigQueryExport) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromSecurityCenterBigQueryExportSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &SecurityCenterBigQueryExportIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change SecurityCenterBigQueryExport identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
