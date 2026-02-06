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
	_ identity.IdentityV2 = &InternalRangeIdentity{}
	_ identity.Resource   = &NetworkConnectivityInternalRange{}
)

var InternalRangeIdentityFormat = gcpurls.Template[InternalRangeIdentity]("networkconnectivity.googleapis.com", "projects/{project}/locations/{location}/internalRanges/{internalRange}")

// +k8s:deepcopy-gen=false
type InternalRangeIdentity struct {
	Project       string
	Location      string
	InternalRange string
}

func (i *InternalRangeIdentity) String() string {
	return InternalRangeIdentityFormat.ToString(*i)
}

func (i *InternalRangeIdentity) FromExternal(ref string) error {
	parsed, match, err := InternalRangeIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkConnectivityInternalRange external=%q was not known (use %s): %w", ref, InternalRangeIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkConnectivityInternalRange external=%q was not known (use %s)", ref, InternalRangeIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *InternalRangeIdentity) Host() string {
	return InternalRangeIdentityFormat.Host()
}

func getIdentityFromInternalRangeSpec(ctx context.Context, reader client.Reader, obj client.Object) (*InternalRangeIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &InternalRangeIdentity{
		Project:       projectID,
		Location:      location,
		InternalRange: resourceID,
	}
	return identity, nil
}

func (obj *NetworkConnectivityInternalRange) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromInternalRangeSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &InternalRangeIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkConnectivityInternalRange identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}