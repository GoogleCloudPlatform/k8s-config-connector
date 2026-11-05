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
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.ServerGeneratedIdentity = &ComputeOrganizationSecurityPolicyIdentity{}
	_ identity.Resource                = &ComputeOrganizationSecurityPolicy{}
)

var ComputeOrganizationSecurityPolicyIdentityFormat = gcpurls.Template[ComputeOrganizationSecurityPolicyIdentity](
	"compute.googleapis.com",
	"locations/global/securityPolicies/{securityPolicy}",
)

// ComputeOrganizationSecurityPolicyIdentity is the identity of a GCP ComputeOrganizationSecurityPolicy resource.
// +k8s:deepcopy-gen=false
type ComputeOrganizationSecurityPolicyIdentity struct {
	SecurityPolicy string
}

func (i *ComputeOrganizationSecurityPolicyIdentity) HasIdentitySpecified() bool {
	return i.SecurityPolicy != ""
}

func (i *ComputeOrganizationSecurityPolicyIdentity) String() string {
	return ComputeOrganizationSecurityPolicyIdentityFormat.ToString(*i)
}

func (i *ComputeOrganizationSecurityPolicyIdentity) FromExternal(ref string) error {
	ref = apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeOrganizationSecurityPolicyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ComputeOrganizationSecurityPolicy external=%q was not known (use %s): %w", ref, ComputeOrganizationSecurityPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeOrganizationSecurityPolicy external=%q was not known (use %s)", ref, ComputeOrganizationSecurityPolicyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeOrganizationSecurityPolicyIdentity) Host() string {
	return ComputeOrganizationSecurityPolicyIdentityFormat.Host()
}

func (i *ComputeOrganizationSecurityPolicyIdentity) ParentString() string {
	return "locations/global"
}

func ParseComputeOrganizationSecurityPolicyExternal(external string) (*ComputeOrganizationSecurityPolicyIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeOrganizationSecurityPolicy external value")
	}
	id := &ComputeOrganizationSecurityPolicyIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeOrganizationSecurityPolicySpec(ctx context.Context, reader client.Reader, obj *ComputeOrganizationSecurityPolicy) (*ComputeOrganizationSecurityPolicyIdentity, error) {
	// For ComputeOrganizationSecurityPolicy, resourceID is optional and can be empty.
	// We retrieve it directly from Spec.ResourceID to avoid falling back to GetName().
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	identity := &ComputeOrganizationSecurityPolicyIdentity{
		SecurityPolicy: resourceID,
	}
	return identity, nil
}

func (obj *ComputeOrganizationSecurityPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeOrganizationSecurityPolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.policyId, if present.
	policyID := common.ValueOf(obj.Status.PolicyId)
	if policyID != "" {
		if specIdentity.SecurityPolicy == "" {
			specIdentity.SecurityPolicy = policyID
		}

		if policyID != specIdentity.SecurityPolicy {
			return nil, fmt.Errorf("cannot change ComputeOrganizationSecurityPolicy identity (old=%q, new=%q)", policyID, specIdentity.SecurityPolicy)
		}
	}

	return specIdentity, nil
}
