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
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NetworkSecurityMirroringEndpointGroupIdentity{}
	_ identity.Resource   = &NetworkSecurityMirroringEndpointGroup{}
)

var NetworkSecurityMirroringEndpointGroupIdentityFormat = gcpurls.Template[NetworkSecurityMirroringEndpointGroupIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/mirroringEndpointGroups/{mirroringEndpointGroup}")

// +k8s:deepcopy-gen=false
type NetworkSecurityMirroringEndpointGroupIdentity struct {
	Project                string
	Location               string
	MirroringEndpointGroup string
}

func (i *NetworkSecurityMirroringEndpointGroupIdentity) String() string {
	return NetworkSecurityMirroringEndpointGroupIdentityFormat.ToString(*i)
}

func (i *NetworkSecurityMirroringEndpointGroupIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkSecurityMirroringEndpointGroupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkSecurityMirroringEndpointGroup external=%q was not known (use %s): %w", ref, NetworkSecurityMirroringEndpointGroupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkSecurityMirroringEndpointGroup external=%q was not known (use %s)", ref, NetworkSecurityMirroringEndpointGroupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkSecurityMirroringEndpointGroupIdentity) Host() string {
	return NetworkSecurityMirroringEndpointGroupIdentityFormat.Host()
}

func getIdentityFromNetworkSecurityMirroringEndpointGroupSpec(ctx context.Context, reader client.Reader, obj client.Object) (*NetworkSecurityMirroringEndpointGroupIdentity, error) {
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

	identity := &NetworkSecurityMirroringEndpointGroupIdentity{
		Project:                projectID,
		Location:               location,
		MirroringEndpointGroup: resourceID,
	}
	return identity, nil
}

func (obj *NetworkSecurityMirroringEndpointGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecurityMirroringEndpointGroupSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &NetworkSecurityMirroringEndpointGroupIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkSecurityMirroringEndpointGroup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier returns the external identifier for the resource.
func (obj *NetworkSecurityMirroringEndpointGroup) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	// Fallback to generating it if possible
	id := obj.GetName()
	if obj.Spec.ResourceID != nil {
		id = *obj.Spec.ResourceID
	}
	if obj.Spec.ProjectRef != nil && obj.Spec.ProjectRef.External != "" && obj.Spec.Location != nil {
		return ptr.To(fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroups/%s", obj.Spec.ProjectRef.External, *obj.Spec.Location, id))
	}
	return nil
}
