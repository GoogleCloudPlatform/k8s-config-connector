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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeOrganizationSecurityPolicyAssociationIdentity{}
	_ identity.Resource   = &ComputeOrganizationSecurityPolicyAssociation{}
)

var ComputeOrganizationSecurityPolicyAssociationIdentityFormat = gcpurls.Template[ComputeOrganizationSecurityPolicyAssociationIdentity](
	"compute.googleapis.com",
	"locations/global/securityPolicies/{securityPolicy}/associations/{association}",
)

// ComputeOrganizationSecurityPolicyAssociationIdentity is the identity of a GCP ComputeOrganizationSecurityPolicyAssociation resource.
// +k8s:deepcopy-gen=false
type ComputeOrganizationSecurityPolicyAssociationIdentity struct {
	SecurityPolicy string
	Association    string
}

func (i *ComputeOrganizationSecurityPolicyAssociationIdentity) String() string {
	return ComputeOrganizationSecurityPolicyAssociationIdentityFormat.ToString(*i)
}

func (i *ComputeOrganizationSecurityPolicyAssociationIdentity) FromExternal(ref string) error {
	ref = apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeOrganizationSecurityPolicyAssociationIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ComputeOrganizationSecurityPolicyAssociation external=%q was not known (use %s): %w", ref, ComputeOrganizationSecurityPolicyAssociationIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeOrganizationSecurityPolicyAssociation external=%q was not known (use %s)", ref, ComputeOrganizationSecurityPolicyAssociationIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeOrganizationSecurityPolicyAssociationIdentity) Host() string {
	return ComputeOrganizationSecurityPolicyAssociationIdentityFormat.Host()
}

func (i *ComputeOrganizationSecurityPolicyAssociationIdentity) ParentString() string {
	return fmt.Sprintf("locations/global/securityPolicies/%s", i.SecurityPolicy)
}

func ParseComputeOrganizationSecurityPolicyAssociationExternal(external string) (*ComputeOrganizationSecurityPolicyAssociationIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeOrganizationSecurityPolicyAssociation external value")
	}
	id := &ComputeOrganizationSecurityPolicyAssociationIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func parseSecurityPolicyID(policyID string) (string, error) {
	if policyID == "" {
		return "", fmt.Errorf("spec.policyId is empty")
	}
	trimmed := apirefs.TrimComputeURIPrefix(policyID)
	if strings.Contains(trimmed, "/securityPolicies/") {
		parts := strings.Split(trimmed, "/securityPolicies/")
		if len(parts) == 2 && parts[1] != "" {
			subparts := strings.Split(parts[1], "/")
			return subparts[0], nil
		}
	}
	if strings.Contains(trimmed, "/") {
		return "", fmt.Errorf("invalid policyId format: %q", policyID)
	}
	return trimmed, nil
}

func getIdentityFromComputeOrganizationSecurityPolicyAssociationSpec(ctx context.Context, reader client.Reader, obj *ComputeOrganizationSecurityPolicyAssociation) (*ComputeOrganizationSecurityPolicyAssociationIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	var policyID string
	if obj.Spec.PolicyID != nil && *obj.Spec.PolicyID != "" {
		policyID, err = parseSecurityPolicyID(*obj.Spec.PolicyID)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve security policy ID: %w", err)
		}
	} else if obj.Spec.SecurityPolicyRef != nil {
		ref := obj.Spec.SecurityPolicyRef
		if err := ref.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, fmt.Errorf("cannot normalize securityPolicyRef: %w", err)
		}
		id, err := ref.ParseExternalToIdentity()
		if err != nil {
			return nil, fmt.Errorf("cannot parse securityPolicyRef: %w", err)
		}
		policyIdentity, ok := id.(*ComputeOrganizationSecurityPolicyIdentity)
		if !ok {
			return nil, fmt.Errorf("unexpected identity type for securityPolicyRef: %T", id)
		}
		policyID = policyIdentity.SecurityPolicy
	} else {
		return nil, fmt.Errorf("cannot resolve security policy ID: exactly one of spec.policyId or spec.securityPolicyRef must be set")
	}

	identity := &ComputeOrganizationSecurityPolicyAssociationIdentity{
		SecurityPolicy: policyID,
		Association:    resourceID,
	}
	return identity, nil
}

func (obj *ComputeOrganizationSecurityPolicyAssociation) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeOrganizationSecurityPolicyAssociationSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	return specIdentity, nil
}
