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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeRouterPeerIdentity{}
	_ identity.Resource   = &ComputeRouterPeer{}
)

var ComputeRouterPeerIdentityFormat = gcpurls.Template[ComputeRouterPeerIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/routers/{router}/bgpPeers/{computerouterpeer}")

// ComputeRouterPeerIdentity is the identity of a GCP ComputeRouterPeer resource.
// +k8s:deepcopy-gen=false
type ComputeRouterPeerIdentity struct {
	Project           string
	Region            string
	Router            string
	ComputeRouterPeer string
}

func (i *ComputeRouterPeerIdentity) String() string {
	return ComputeRouterPeerIdentityFormat.ToString(*i)
}

func (i *ComputeRouterPeerIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeRouterPeerIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeRouterPeer external=%q was not known (use %s): %w", ref, ComputeRouterPeerIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeRouterPeer external=%q was not known (use %s)", ref, ComputeRouterPeerIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeRouterPeerIdentity) Host() string {
	return ComputeRouterPeerIdentityFormat.Host()
}

func (i *ComputeRouterPeerIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/regions/%s/routers/%s", i.Project, i.Region, i.Router)
}

func ParseComputeRouterPeerExternal(external string) (*ComputeRouterPeerIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeRouterPeer external value")
	}
	id := &ComputeRouterPeerIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeRouterPeerSpec(ctx context.Context, reader client.Reader, obj *ComputeRouterPeer) (*ComputeRouterPeerIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	region := obj.Spec.Region
	if region == "" {
		return nil, fmt.Errorf("cannot resolve region: spec.region is empty")
	}

	routerRefExternal, err := obj.Spec.RouterRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, fmt.Errorf("cannot resolve routerRef: %w", err)
	}

	trimmedRouterRef := apirefs.TrimComputeURIPrefix(routerRefExternal)
	tokens := strings.Split(trimmedRouterRef, "/")
	if len(tokens) == 0 {
		return nil, fmt.Errorf("invalid routerRef resolved external: %q", routerRefExternal)
	}
	routerID := tokens[len(tokens)-1]

	identity := &ComputeRouterPeerIdentity{
		Project:           projectID,
		Region:            region,
		Router:            routerID,
		ComputeRouterPeer: resourceID,
	}
	return identity, nil
}

func (obj *ComputeRouterPeer) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeRouterPeerSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.externalRef, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &ComputeRouterPeerIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if *specIdentity != *statusIdentity {
			return nil, fmt.Errorf("spec identity %s doesn't match status identity %s", specIdentity.String(), statusIdentity.String())
		}
	}

	return specIdentity, nil
}
