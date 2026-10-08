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
	_ identity.IdentityV2 = &LBEdgeExtensionIdentity{}
	_ identity.Resource   = &NetworkServicesLBEdgeExtension{}
)

var LBEdgeExtensionIdentityFormat = gcpurls.Template[LBEdgeExtensionIdentity](
	"networkservices.googleapis.com",
	"projects/{project}/locations/{location}/lbEdgeExtensions/{lbEdgeExtension}",
)

// LBEdgeExtensionIdentity is the identity of a NetworkServicesLBEdgeExtension.
// +k8s:deepcopy-gen=false
type LBEdgeExtensionIdentity struct {
	Project         string
	Location        string
	LbEdgeExtension string
}

func (i *LBEdgeExtensionIdentity) String() string {
	return LBEdgeExtensionIdentityFormat.ToString(*i)
}

func (i *LBEdgeExtensionIdentity) FromExternal(ref string) error {
	parsed, match, err := LBEdgeExtensionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkServicesLBEdgeExtension external=%q was not known (use %s): %w", ref, LBEdgeExtensionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkServicesLBEdgeExtension external=%q was not known (use %s)", ref, LBEdgeExtensionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *LBEdgeExtensionIdentity) Host() string {
	return LBEdgeExtensionIdentityFormat.Host()
}

func getIdentityFromLBEdgeExtensionSpec(ctx context.Context, reader client.Reader, obj client.Object) (*LBEdgeExtensionIdentity, error) {
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

	identity := &LBEdgeExtensionIdentity{
		Project:         projectID,
		Location:        location,
		LbEdgeExtension: resourceID,
	}
	return identity, nil
}

func (obj *NetworkServicesLBEdgeExtension) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromLBEdgeExtensionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &LBEdgeExtensionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkServicesLBEdgeExtension identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// NewLBEdgeExtensionIdentity is a helper used by the direct controller.
func NewLBEdgeExtensionIdentity(ctx context.Context, reader client.Reader, obj *NetworkServicesLBEdgeExtension) (*LBEdgeExtensionIdentity, error) {
	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	return identity.(*LBEdgeExtensionIdentity), nil
}
