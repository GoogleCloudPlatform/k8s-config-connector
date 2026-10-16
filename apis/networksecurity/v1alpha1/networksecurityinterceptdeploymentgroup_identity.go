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
	_ identity.IdentityV2 = &NetworkSecurityInterceptDeploymentGroupIdentity{}
	_ identity.Resource   = &NetworkSecurityInterceptDeploymentGroup{}
)

var (
	NetworkSecurityInterceptDeploymentGroupIdentityFormat = gcpurls.Template[NetworkSecurityInterceptDeploymentGroupIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/interceptDeploymentGroups/{interceptdeploymentgroup}")
)

// NetworkSecurityInterceptDeploymentGroupIdentity is the identity of a GCP NetworkSecurityInterceptDeploymentGroup resource.
// +k8s:deepcopy-gen=false
type NetworkSecurityInterceptDeploymentGroupIdentity struct {
	Project                  string
	Location                 string
	InterceptDeploymentGroup string
}

func (i *NetworkSecurityInterceptDeploymentGroupIdentity) String() string {
	return NetworkSecurityInterceptDeploymentGroupIdentityFormat.ToString(*i)
}

func (i *NetworkSecurityInterceptDeploymentGroupIdentity) Host() string {
	return NetworkSecurityInterceptDeploymentGroupIdentityFormat.Host()
}

func (i *NetworkSecurityInterceptDeploymentGroupIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkSecurityInterceptDeploymentGroupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkSecurityInterceptDeploymentGroupIdentity external=%q was not known (use %s): %w", ref, NetworkSecurityInterceptDeploymentGroupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkSecurityInterceptDeploymentGroupIdentity external=%q was not known (use %s)", ref, NetworkSecurityInterceptDeploymentGroupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkSecurityInterceptDeploymentGroupIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromNetworkSecurityInterceptDeploymentGroupSpec(ctx context.Context, reader client.Reader, obj *NetworkSecurityInterceptDeploymentGroup) (*NetworkSecurityInterceptDeploymentGroupIdentity, error) {
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

	identity := &NetworkSecurityInterceptDeploymentGroupIdentity{
		Project:                  projectID,
		Location:                 location,
		InterceptDeploymentGroup: resourceID,
	}
	return identity, nil
}

func (obj *NetworkSecurityInterceptDeploymentGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecurityInterceptDeploymentGroupSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &NetworkSecurityInterceptDeploymentGroupIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkSecurityInterceptDeploymentGroup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier returns the external identifier for the resource.
func (obj *NetworkSecurityInterceptDeploymentGroup) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
