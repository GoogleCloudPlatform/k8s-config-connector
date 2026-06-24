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
	"strings"

	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ApigeeAPIProductIdentity{}
	_ identity.Resource   = &ApigeeAPIProduct{}
)

var ApigeeAPIProductIdentityFormat = gcpurls.Template[ApigeeAPIProductIdentity]("apigee.googleapis.com", "organizations/{organization}/apiproducts/{apiproduct}")

// +k8s:deepcopy-gen=false
type ApigeeAPIProductIdentity struct {
	Organization string
	Apiproduct   string
}

func (i *ApigeeAPIProductIdentity) String() string {
	return ApigeeAPIProductIdentityFormat.ToString(*i)
}

func (i *ApigeeAPIProductIdentity) FromExternal(ref string) error {
	parsed, match, err := ApigeeAPIProductIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ApigeeAPIProduct external=%q was not known (use %s): %w", ref, ApigeeAPIProductIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ApigeeAPIProduct external=%q was not known (use %s)", ref, ApigeeAPIProductIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ApigeeAPIProductIdentity) Host() string {
	return ApigeeAPIProductIdentityFormat.Host()
}

func (i *ApigeeAPIProductIdentity) ID() string {
	return i.Apiproduct
}

func (i *ApigeeAPIProductIdentity) ParentString() string {
	return fmt.Sprintf("organizations/%s", i.Organization)
}

func getIdentityFromApigeeAPIProductSpec(ctx context.Context, reader client.Reader, obj *ApigeeAPIProduct) (*ApigeeAPIProductIdentity, error) {
	if obj.Spec.OrganizationRef == nil {
		return nil, fmt.Errorf("spec.organizationRef is required")
	}
	organizationRef := obj.Spec.OrganizationRef.DeepCopy()

	// Special case: ApigeeOrganizationRef doesn't automatically convert "projects/foo" to "organizations/foo"
	// if it is done within Normalize. However, Normalize() may fail if the resource isn't actually ready or available in a mock reader,
	// so for external refs starting with "projects/", we parse them manually here for identity evaluation.
	orgExternal := organizationRef.External
	if strings.HasPrefix(orgExternal, "projects/") {
		projectRef := &refs.ProjectRef{External: orgExternal}
		parsedProject := &refs.ProjectIdentity{}
		if err := parsedProject.FromExternal(projectRef.External); err == nil {
			orgExternal = fmt.Sprintf("organizations/%s", parsedProject.ProjectID)
			organizationRef.External = orgExternal
		}
	}

	if err := organizationRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}

	parentID := &apigeev1beta1.ApigeeOrganizationIdentity{}
	if err := parentID.FromExternal(organizationRef.External); err != nil {
		return nil, err
	}

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	return &ApigeeAPIProductIdentity{
		Organization: parentID.ResourceID,
		Apiproduct:   resourceID,
	}, nil
}

func (obj *ApigeeAPIProduct) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromApigeeAPIProductSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ApigeeAPIProductIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ApigeeAPIProduct identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
