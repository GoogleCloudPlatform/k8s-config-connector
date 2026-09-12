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
	_ identity.IdentityV2 = &ComputeNetworkPeeringIdentity{}
	_ identity.Resource   = &ComputeNetworkPeering{}
)

var ComputeNetworkPeeringIdentityFormat = gcpurls.Template[ComputeNetworkPeeringIdentity]("compute.googleapis.com", "projects/{project}/global/networks/{network}/networkPeerings/{computenetworkpeering}")

// ComputeNetworkPeeringIdentity is the identity of a GCP ComputeNetworkPeering resource.
// +k8s:deepcopy-gen=false
type ComputeNetworkPeeringIdentity struct {
	Project               string
	Network               string
	ComputeNetworkPeering string
}

func (i *ComputeNetworkPeeringIdentity) String() string {
	return ComputeNetworkPeeringIdentityFormat.ToString(*i)
}

func (i *ComputeNetworkPeeringIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeNetworkPeeringIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeNetworkPeering external=%q was not known (use %s): %w", ref, ComputeNetworkPeeringIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeNetworkPeering external=%q was not known (use %s)", ref, ComputeNetworkPeeringIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeNetworkPeeringIdentity) Host() string {
	return ComputeNetworkPeeringIdentityFormat.Host()
}

func (i *ComputeNetworkPeeringIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/global/networks/%s", i.Project, i.Network)
}

func ParseComputeNetworkPeeringExternal(external string) (*ComputeNetworkPeeringIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeNetworkPeering external value")
	}
	id := &ComputeNetworkPeeringIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeNetworkPeeringSpec(ctx context.Context, reader client.Reader, obj *ComputeNetworkPeering) (*ComputeNetworkPeeringIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	networkRef := &obj.Spec.NetworkRef
	if err := networkRef.Normalize(ctx, reader, obj.Namespace); err != nil {
		return nil, fmt.Errorf("cannot normalize networkRef: %w", err)
	}
	networkExternal := networkRef.External
	if networkExternal == "" {
		return nil, fmt.Errorf("cannot resolve networkRef")
	}
	networkIdentity, err := ParseComputeNetworkExternal(networkExternal)
	if err != nil {
		return nil, fmt.Errorf("cannot parse resolved networkRef external=%q: %w", networkExternal, err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeNetworkPeeringIdentity{
		Project:               projectID,
		Network:               networkIdentity.Network,
		ComputeNetworkPeering: resourceID,
	}
	return identity, nil
}

func (obj *ComputeNetworkPeering) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeNetworkPeeringSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
