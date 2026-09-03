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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NetworkSecurityMirroringDeploymentGroupIdentity{}
	_ identity.Resource   = &NetworkSecurityMirroringDeploymentGroup{}
)

var (
	NetworkSecurityMirroringDeploymentGroupIdentityFormat = gcpurls.Template[NetworkSecurityMirroringDeploymentGroupIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/mirroringDeploymentGroups/{mirroringdeploymentgroup}")
)

// NetworkSecurityMirroringDeploymentGroupIdentity is the identity of a GCP NetworkSecurityMirroringDeploymentGroup resource.
// +k8s:deepcopy-gen=false
type NetworkSecurityMirroringDeploymentGroupIdentity struct {
	Project                  string
	Location                 string
	MirroringDeploymentGroup string
}

func (i *NetworkSecurityMirroringDeploymentGroupIdentity) String() string {
	return NetworkSecurityMirroringDeploymentGroupIdentityFormat.ToString(*i)
}

func (i *NetworkSecurityMirroringDeploymentGroupIdentity) Host() string {
	return NetworkSecurityMirroringDeploymentGroupIdentityFormat.Host()
}

func (i *NetworkSecurityMirroringDeploymentGroupIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkSecurityMirroringDeploymentGroupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkSecurityMirroringDeploymentGroupIdentity external=%q was not known (use %s): %w", ref, NetworkSecurityMirroringDeploymentGroupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkSecurityMirroringDeploymentGroupIdentity external=%q was not known (use %s)", ref, NetworkSecurityMirroringDeploymentGroupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkSecurityMirroringDeploymentGroupIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromNetworkSecurityMirroringDeploymentGroupSpec(ctx context.Context, reader client.Reader, obj *NetworkSecurityMirroringDeploymentGroup) (*NetworkSecurityMirroringDeploymentGroupIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &NetworkSecurityMirroringDeploymentGroupIdentity{
		Project:                  projectID,
		Location:                 location,
		MirroringDeploymentGroup: resourceID,
	}
	return identity, nil
}

func (obj *NetworkSecurityMirroringDeploymentGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecurityMirroringDeploymentGroupSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &NetworkSecurityMirroringDeploymentGroupIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkSecurityMirroringDeploymentGroup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier returns the external identifier for the resource.
func (obj *NetworkSecurityMirroringDeploymentGroup) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
