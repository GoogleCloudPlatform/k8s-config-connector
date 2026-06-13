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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &FirewallIdentity{}
	_ identity.Resource   = &ComputeFirewall{}
)

var ComputeFirewallIdentityFormat = gcpurls.Template[FirewallIdentity]("compute.googleapis.com", "projects/{project}/global/firewalls/{firewall}")

// FirewallIdentity is the identity of a GCP ComputeFirewall resource.
// +k8s:deepcopy-gen=false
type FirewallIdentity struct {
	Project  string
	Firewall string
}

func (i *FirewallIdentity) String() string {
	return ComputeFirewallIdentityFormat.ToString(*i)
}

func (i *FirewallIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeFirewallIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeFirewall external=%q was not known (use %s): %w", ref, ComputeFirewallIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeFirewall external=%q was not known (use %s)", ref, ComputeFirewallIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *FirewallIdentity) Host() string {
	return ComputeFirewallIdentityFormat.Host()
}

// Deprecated: ParseComputeFirewallExternal is deprecated and should not be used.
func ParseComputeFirewallExternal(external string) (*FirewallIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeFirewall external value")
	}
	id := &FirewallIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeFirewallSpec(ctx context.Context, reader client.Reader, obj *ComputeFirewall) (*FirewallIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &FirewallIdentity{
		Project:  projectID,
		Firewall: resourceID,
	}
	return identity, nil
}

func (obj *ComputeFirewall) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeFirewallSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &FirewallIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeFirewall identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
