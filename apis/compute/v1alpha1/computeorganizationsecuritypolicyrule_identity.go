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
	_ identity.IdentityV2 = &ComputeOrganizationSecurityPolicyRuleIdentity{}
	_ identity.Resource   = &ComputeOrganizationSecurityPolicyRule{}
)

var ComputeOrganizationSecurityPolicyRuleIdentityFormat = gcpurls.Template[ComputeOrganizationSecurityPolicyRuleIdentity](
	"compute.googleapis.com",
	"locations/global/securityPolicies/{securityPolicy}/rules/{priority}",
)

// ComputeOrganizationSecurityPolicyRuleIdentity is the identity of a GCP ComputeOrganizationSecurityPolicyRule resource.
// +k8s:deepcopy-gen=false
type ComputeOrganizationSecurityPolicyRuleIdentity struct {
	SecurityPolicy string
	Priority       string
}

func (i *ComputeOrganizationSecurityPolicyRuleIdentity) String() string {
	return ComputeOrganizationSecurityPolicyRuleIdentityFormat.ToString(*i)
}

func (i *ComputeOrganizationSecurityPolicyRuleIdentity) FromExternal(ref string) error {
	ref = apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeOrganizationSecurityPolicyRuleIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ComputeOrganizationSecurityPolicyRule external=%q was not known (use %s): %w", ref, ComputeOrganizationSecurityPolicyRuleIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeOrganizationSecurityPolicyRule external=%q was not known (use %s)", ref, ComputeOrganizationSecurityPolicyRuleIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeOrganizationSecurityPolicyRuleIdentity) Host() string {
	return ComputeOrganizationSecurityPolicyRuleIdentityFormat.Host()
}

func (i *ComputeOrganizationSecurityPolicyRuleIdentity) ParentString() string {
	return fmt.Sprintf("locations/global/securityPolicies/%s", i.SecurityPolicy)
}

func ParseComputeOrganizationSecurityPolicyRuleExternal(external string) (*ComputeOrganizationSecurityPolicyRuleIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeOrganizationSecurityPolicyRule external value")
	}
	id := &ComputeOrganizationSecurityPolicyRuleIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func parseSecurityPolicyRuleParentID(policyID string) (string, error) {
	if policyID == "" {
		return "", fmt.Errorf("spec.policyId is empty")
	}
	trimmed := apirefs.TrimComputeURIPrefix(policyID)
	trimmed = strings.TrimPrefix(trimmed, "/")
	if strings.HasPrefix(trimmed, "securityPolicies/") {
		trimmed = strings.TrimPrefix(trimmed, "securityPolicies/")
	} else if strings.Contains(trimmed, "/securityPolicies/") {
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

func getIdentityFromComputeOrganizationSecurityPolicyRuleSpec(ctx context.Context, reader client.Reader, obj *ComputeOrganizationSecurityPolicyRule) (*ComputeOrganizationSecurityPolicyRuleIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	policyID, err := parseSecurityPolicyRuleParentID(obj.Spec.PolicyID)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve security policy ID: %w", err)
	}

	identity := &ComputeOrganizationSecurityPolicyRuleIdentity{
		SecurityPolicy: policyID,
		Priority:       resourceID,
	}
	return identity, nil
}

func (obj *ComputeOrganizationSecurityPolicyRule) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeOrganizationSecurityPolicyRuleSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	return specIdentity, nil
}
