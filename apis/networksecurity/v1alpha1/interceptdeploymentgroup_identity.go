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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.IdentityV2 = &NetworkSecurityInterceptDeploymentGroupIdentity{}
var _ identity.Resource = &NetworkSecurityInterceptDeploymentGroup{}

var (
	NetworkSecurityInterceptDeploymentGroupFormat             = gcpurls.Template[NetworkSecurityInterceptDeploymentGroupIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/interceptDeploymentGroups/{intercept_deployment_group}")
	NetworkSecurityInterceptDeploymentGroupOrganizationFormat = gcpurls.Template[NetworkSecurityInterceptDeploymentGroupIdentity]("networksecurity.googleapis.com", "organizations/{organization}/locations/{location}/interceptDeploymentGroups/{intercept_deployment_group}")
)

// +k8s:deepcopy-gen=false
type NetworkSecurityInterceptDeploymentGroupIdentity struct {
	Project                    string
	Organization               string
	Location                   string
	Intercept_deployment_group string
}

func (i *NetworkSecurityInterceptDeploymentGroupIdentity) String() string {
	if i.Organization != "" {
		return NetworkSecurityInterceptDeploymentGroupOrganizationFormat.ToString(*i)
	}
	return NetworkSecurityInterceptDeploymentGroupFormat.ToString(*i)
}

func (i *NetworkSecurityInterceptDeploymentGroupIdentity) FromExternal(ref string) error {
	if parsed, match, err := NetworkSecurityInterceptDeploymentGroupFormat.Parse(ref); match {
		*i = *parsed
		return nil
	} else if err != nil {
		return err
	}
	if parsed, match, err := NetworkSecurityInterceptDeploymentGroupOrganizationFormat.Parse(ref); match {
		*i = *parsed
		return nil
	} else if err != nil {
		return err
	}
	return fmt.Errorf("format of NetworkSecurityInterceptDeploymentGroup external=%q was not known (use %s or %s)", ref, NetworkSecurityInterceptDeploymentGroupFormat.CanonicalForm(), NetworkSecurityInterceptDeploymentGroupOrganizationFormat.CanonicalForm())
}

func (i *NetworkSecurityInterceptDeploymentGroupIdentity) Host() string {
	return "networksecurity.googleapis.com"
}

func (r *NetworkSecurityInterceptDeploymentGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromNetworkSecurityInterceptDeploymentGroupSpec(ctx, reader, r)
}

func getIdentityFromNetworkSecurityInterceptDeploymentGroupSpec(ctx context.Context, reader client.Reader, obj client.Object) (*NetworkSecurityInterceptDeploymentGroupIdentity, error) {
	project, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	actualLocation, err := refs.GetLocation(obj)
	if err != nil {
		return nil, err
	}

	actualResourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	id := &NetworkSecurityInterceptDeploymentGroupIdentity{
		Project:                    project,
		Location:                   actualLocation,
		Intercept_deployment_group: actualResourceID,
	}

	// If there is an external ref, and it parses successfully, it must match.
	// We want to verify that the externalRef matched what we expected,
	// but there might be multiple representations of the same resource.
	var externalRef string
	switch obj := obj.(type) {
	case *NetworkSecurityInterceptDeploymentGroup:
		if obj.Status.ExternalRef != nil {
			externalRef = *obj.Status.ExternalRef
		}
	case *unstructured.Unstructured:
		externalRef, _, _ = unstructured.NestedString(obj.Object, "status", "externalRef")
	}

	if externalRef != "" {
		externalId := &NetworkSecurityInterceptDeploymentGroupIdentity{}
		if err := externalId.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("parsing externalRef %q: %w", externalRef, err)
		}

		if externalId.Project != id.Project || externalId.Organization != id.Organization || externalId.Location != id.Location || externalId.Intercept_deployment_group != id.Intercept_deployment_group {
			return nil, fmt.Errorf("externalRef %q does not match identity %q", externalRef, id.String())
		}
	}

	return id, nil
}

func (obj *NetworkSecurityInterceptDeploymentGroup) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil && *obj.Status.ExternalRef != "" {
		return obj.Status.ExternalRef
	}
	return nil
}
