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
	_ identity.IdentityV2 = &NetworkSecurityBackendAuthenticationConfigIdentity{}
	_ identity.Resource   = &NetworkSecurityBackendAuthenticationConfig{}
)

var NetworkSecurityBackendAuthenticationConfigIdentityFormat = gcpurls.Template[NetworkSecurityBackendAuthenticationConfigIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/backendAuthenticationConfigs/{backend_authentication_config}")

// +k8s:deepcopy-gen=false
type NetworkSecurityBackendAuthenticationConfigIdentity struct {
	Project                     string
	Location                    string
	BackendAuthenticationConfig string `gcpurls:"backend_authentication_config"`
}

func (i *NetworkSecurityBackendAuthenticationConfigIdentity) String() string {
	return NetworkSecurityBackendAuthenticationConfigIdentityFormat.ToString(*i)
}

func (i *NetworkSecurityBackendAuthenticationConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkSecurityBackendAuthenticationConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkSecurityBackendAuthenticationConfig external=%q was not known (use %s): %w", ref, NetworkSecurityBackendAuthenticationConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkSecurityBackendAuthenticationConfig external=%q was not known (use %s)", ref, NetworkSecurityBackendAuthenticationConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkSecurityBackendAuthenticationConfigIdentity) Host() string {
	return NetworkSecurityBackendAuthenticationConfigIdentityFormat.Host()
}

func getIdentityFromNetworkSecurityBackendAuthenticationConfigSpec(ctx context.Context, reader client.Reader, obj client.Object) (*NetworkSecurityBackendAuthenticationConfigIdentity, error) {
	resource, ok := obj.(*NetworkSecurityBackendAuthenticationConfig)
	if !ok {
		return nil, fmt.Errorf("object is not a NetworkSecurityBackendAuthenticationConfig")
	}
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := resource.Spec.Location
	if location == nil || *location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &NetworkSecurityBackendAuthenticationConfigIdentity{
		Project:                     projectID,
		Location:                    *location,
		BackendAuthenticationConfig: resourceID,
	}
	return identity, nil
}

func (obj *NetworkSecurityBackendAuthenticationConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecurityBackendAuthenticationConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &NetworkSecurityBackendAuthenticationConfigIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkSecurityBackendAuthenticationConfig identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
