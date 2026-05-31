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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NetworkSecurityClientTLSPolicyIdentity{}
	_ identity.Resource   = &NetworkSecurityClientTLSPolicy{}
)

var NetworkSecurityClientTLSPolicyIdentityFormat = gcpurls.Template[NetworkSecurityClientTLSPolicyIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/clientTlsPolicies/{clientTlsPolicy}")

// +k8s:deepcopy-gen=false
type NetworkSecurityClientTLSPolicyIdentity struct {
	Project         string
	Location        string
	ClientTlsPolicy string
}

func (i *NetworkSecurityClientTLSPolicyIdentity) String() string {
	return NetworkSecurityClientTLSPolicyIdentityFormat.ToString(*i)
}

func (i *NetworkSecurityClientTLSPolicyIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkSecurityClientTLSPolicyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkSecurityClientTLSPolicy external=%q was not known (use %s): %w", ref, NetworkSecurityClientTLSPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkSecurityClientTLSPolicy external=%q was not known (use %s)", ref, NetworkSecurityClientTLSPolicyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkSecurityClientTLSPolicyIdentity) Host() string {
	return NetworkSecurityClientTLSPolicyIdentityFormat.Host()
}

func getIdentityFromNetworkSecurityClientTLSPolicySpec(ctx context.Context, reader client.Reader, obj client.Object) (*NetworkSecurityClientTLSPolicyIdentity, error) {
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

	identity := &NetworkSecurityClientTLSPolicyIdentity{
		Project:         projectID,
		Location:        location,
		ClientTlsPolicy: resourceID,
	}
	return identity, nil
}

func (obj *NetworkSecurityClientTLSPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecurityClientTLSPolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
