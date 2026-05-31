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
	_ identity.IdentityV2 = &NetworkSecurityInterceptDeploymentIdentity{}
	_ identity.Resource   = &NetworkSecurityInterceptDeployment{}
)

var NetworkSecurityInterceptDeploymentIdentityFormat = gcpurls.Template[NetworkSecurityInterceptDeploymentIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/interceptDeployments/{interceptdeployment}")

// +k8s:deepcopy-gen=false
type NetworkSecurityInterceptDeploymentIdentity struct {
	Project             string
	Location            string
	InterceptDeployment string
}

func (i *NetworkSecurityInterceptDeploymentIdentity) String() string {
	return NetworkSecurityInterceptDeploymentIdentityFormat.ToString(*i)
}

func (i *NetworkSecurityInterceptDeploymentIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkSecurityInterceptDeploymentIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkSecurityInterceptDeployment external=%q was not known (use %s): %w", ref, NetworkSecurityInterceptDeploymentIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkSecurityInterceptDeployment external=%q was not known (use %s)", ref, NetworkSecurityInterceptDeploymentIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkSecurityInterceptDeploymentIdentity) Host() string {
	return NetworkSecurityInterceptDeploymentIdentityFormat.Host()
}

func getIdentityFromNetworkSecurityInterceptDeploymentSpec(ctx context.Context, reader client.Reader, obj client.Object) (*NetworkSecurityInterceptDeploymentIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &NetworkSecurityInterceptDeploymentIdentity{
		Project:             projectID,
		Location:            location,
		InterceptDeployment: resourceID,
	}
	return identity, nil
}

func (obj *NetworkSecurityInterceptDeployment) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecurityInterceptDeploymentSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &NetworkSecurityInterceptDeploymentIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkSecurityInterceptDeployment identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *NetworkSecurityInterceptDeployment) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
