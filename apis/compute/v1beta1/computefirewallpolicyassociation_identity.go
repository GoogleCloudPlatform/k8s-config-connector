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

package v1beta1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeFirewallPolicyAssociationIdentity{}
	_ identity.Resource   = &ComputeFirewallPolicyAssociation{}
)

var ComputeFirewallPolicyAssociationIdentityFormat = gcpurls.Template[ComputeFirewallPolicyAssociationIdentity]("compute.googleapis.com", "locations/global/firewallPolicies/{firewallPolicy}/associations/{association}")

// ComputeFirewallPolicyAssociationIdentity is the identity of a GCP ComputeFirewallPolicyAssociation resource.
// +k8s:deepcopy-gen=false
type ComputeFirewallPolicyAssociationIdentity struct {
	FirewallPolicy string
	Association    string
}

func (i *ComputeFirewallPolicyAssociationIdentity) String() string {
	return ComputeFirewallPolicyAssociationIdentityFormat.ToString(*i)
}

func (i *ComputeFirewallPolicyAssociationIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeFirewallPolicyAssociationIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeFirewallPolicyAssociation external=%q was not known (use %s): %w", ref, ComputeFirewallPolicyAssociationIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeFirewallPolicyAssociation external=%q was not known (use %s)", ref, ComputeFirewallPolicyAssociationIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeFirewallPolicyAssociationIdentity) Host() string {
	return ComputeFirewallPolicyAssociationIdentityFormat.Host()
}

func (i *ComputeFirewallPolicyAssociationIdentity) ParentString() string {
	return "locations/global/firewallPolicies/" + i.FirewallPolicy
}

func ParseComputeFirewallPolicyAssociationExternal(external string) (*ComputeFirewallPolicyAssociationIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeFirewallPolicyAssociation external value")
	}
	id := &ComputeFirewallPolicyAssociationIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeFirewallPolicyAssociationSpec(ctx context.Context, reader client.Reader, obj *ComputeFirewallPolicyAssociation) (*ComputeFirewallPolicyAssociationIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	// Resolve the firewallPolicyRef
	firewallPolicyRef := &obj.Spec.FirewallPolicyRef
	if err := firewallPolicyRef.Normalize(ctx, reader, obj.Namespace); err != nil {
		return nil, fmt.Errorf("cannot normalize firewallPolicyRef: %w", err)
	}
	firewallPolicyExternal := firewallPolicyRef.External
	if firewallPolicyExternal == "" {
		return nil, fmt.Errorf("cannot resolve firewallPolicyRef")
	}

	firewallPolicyID, err := ParseComputeFirewallPolicyExternal(firewallPolicyExternal)
	if err != nil {
		return nil, fmt.Errorf("cannot parse resolved firewallPolicyRef external=%q: %w", firewallPolicyExternal, err)
	}

	identity := &ComputeFirewallPolicyAssociationIdentity{
		FirewallPolicy: firewallPolicyID.FirewallPolicy,
		Association:    resourceID,
	}
	return identity, nil
}

func (obj *ComputeFirewallPolicyAssociation) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeFirewallPolicyAssociationSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
