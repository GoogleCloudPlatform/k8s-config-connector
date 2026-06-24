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
	_ identity.IdentityV2 = &NetworkSecurityAddressGroupIdentity{}
	_ identity.Resource   = &NetworkSecurityAddressGroup{}
)

var NetworkSecurityAddressGroupIdentityFormat = gcpurls.Template[NetworkSecurityAddressGroupIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/addressGroups/{addressgroup}")

// +k8s:deepcopy-gen=false
type NetworkSecurityAddressGroupIdentity struct {
	Project      string
	Location     string
	Addressgroup string
}

func (i *NetworkSecurityAddressGroupIdentity) String() string {
	return NetworkSecurityAddressGroupIdentityFormat.ToString(*i)
}

func (i *NetworkSecurityAddressGroupIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkSecurityAddressGroupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkSecurityAddressGroup external=%q was not known (use %s): %w", ref, NetworkSecurityAddressGroupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkSecurityAddressGroup external=%q was not known (use %s)", ref, NetworkSecurityAddressGroupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkSecurityAddressGroupIdentity) Host() string {
	return NetworkSecurityAddressGroupIdentityFormat.Host()
}

func getIdentityFromNetworkSecurityAddressGroupSpec(ctx context.Context, reader client.Reader, obj client.Object) (*NetworkSecurityAddressGroupIdentity, error) {
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

	identity := &NetworkSecurityAddressGroupIdentity{
		Project:      projectID,
		Location:     location,
		Addressgroup: resourceID,
	}
	return identity, nil
}

func (obj *NetworkSecurityAddressGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecurityAddressGroupSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &NetworkSecurityAddressGroupIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkSecurityAddressGroup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier returns the GCP external identifier (the GCP URL).
func (obj *NetworkSecurityAddressGroup) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
