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
	_ identity.IdentityV2 = &AuthzExtensionIdentity{}
	_ identity.Resource   = &NetworkServicesAuthzExtension{}
)

var AuthzExtensionIdentityFormat = gcpurls.Template[AuthzExtensionIdentity](
	"networkservices.googleapis.com",
	"projects/{project}/locations/{location}/authzExtensions/{authzExtension}",
)

// AuthzExtensionIdentity is the identity of a NetworkServicesAuthzExtension.
// +k8s:deepcopy-gen=false
type AuthzExtensionIdentity struct {
	Project        string
	Location       string
	AuthzExtension string
}

func (i *AuthzExtensionIdentity) String() string {
	return AuthzExtensionIdentityFormat.ToString(*i)
}

func (i *AuthzExtensionIdentity) FromExternal(ref string) error {
	parsed, match, err := AuthzExtensionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkServicesAuthzExtension external=%q was not known (use %s): %w", ref, AuthzExtensionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkServicesAuthzExtension external=%q was not known (use %s)", ref, AuthzExtensionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *AuthzExtensionIdentity) Host() string {
	return AuthzExtensionIdentityFormat.Host()
}

func (i *AuthzExtensionIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromAuthzExtensionSpec(ctx context.Context, reader client.Reader, obj client.Object) (*AuthzExtensionIdentity, error) {
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

	identity := &AuthzExtensionIdentity{
		Project:        projectID,
		Location:       location,
		AuthzExtension: resourceID,
	}
	return identity, nil
}

func (obj *NetworkServicesAuthzExtension) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromAuthzExtensionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &AuthzExtensionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkServicesAuthzExtension identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// NewAuthzExtensionIdentity is a helper used by the direct controller.
func NewAuthzExtensionIdentity(ctx context.Context, reader client.Reader, obj *NetworkServicesAuthzExtension) (*AuthzExtensionIdentity, error) {
	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	return identity.(*AuthzExtensionIdentity), nil
}

func (obj *NetworkServicesAuthzExtension) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
