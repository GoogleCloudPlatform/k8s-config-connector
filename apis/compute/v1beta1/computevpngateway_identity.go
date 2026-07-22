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
	_ identity.IdentityV2 = &ComputeVPNGatewayIdentity{}
	_ identity.Resource   = &ComputeVPNGateway{}
)

var ComputeVPNGatewayIdentityFormat = gcpurls.Template[ComputeVPNGatewayIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/vpnGateways/{vpngateway}")

// ComputeVPNGatewayIdentity is the identity of a GCP ComputeVPNGateway resource.
// +k8s:deepcopy-gen=false
type ComputeVPNGatewayIdentity struct {
	Project    string
	Region     string
	VpnGateway string
}

func (i *ComputeVPNGatewayIdentity) String() string {
	return ComputeVPNGatewayIdentityFormat.ToString(*i)
}

func (i *ComputeVPNGatewayIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeVPNGatewayIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeVPNGateway external=%q was not known (use %s): %w", ref, ComputeVPNGatewayIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeVPNGateway external=%q was not known (use %s)", ref, ComputeVPNGatewayIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeVPNGatewayIdentity) Host() string {
	return ComputeVPNGatewayIdentityFormat.Host()
}

func (i *ComputeVPNGatewayIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Region)
}

func ParseComputeVPNGatewayExternal(external string) (*ComputeVPNGatewayIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeVPNGateway external value")
	}
	id := &ComputeVPNGatewayIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeVPNGatewaySpec(ctx context.Context, reader client.Reader, obj *ComputeVPNGateway) (*ComputeVPNGatewayIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeVPNGatewayIdentity{
		Project:    projectID,
		Region:     obj.Spec.Region,
		VpnGateway: resourceID,
	}
	return identity, nil
}

func (obj *ComputeVPNGateway) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeVPNGatewaySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &ComputeVPNGatewayIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeVPNGateway identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
