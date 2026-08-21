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
	_ identity.IdentityV2 = &DiscoveryEngineLicenseConfigIdentity{}
	_ identity.Resource   = &DiscoveryEngineLicenseConfig{}
)

var DiscoveryEngineLicenseConfigIdentityFormat = gcpurls.Template[DiscoveryEngineLicenseConfigIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/licenseConfigs/{licenseConfig}")

// DiscoveryEngineLicenseConfigIdentity is the identity of a GCP DiscoveryEngineLicenseConfig resource.
// +k8s:deepcopy-gen=false
type DiscoveryEngineLicenseConfigIdentity struct {
	Project       string
	Location      string
	LicenseConfig string
}

func (i *DiscoveryEngineLicenseConfigIdentity) String() string {
	return DiscoveryEngineLicenseConfigIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineLicenseConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineLicenseConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineLicenseConfig external=%q was not known (use %s): %w", ref, DiscoveryEngineLicenseConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineLicenseConfig external=%q was not known (use %s)", ref, DiscoveryEngineLicenseConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineLicenseConfigIdentity) Host() string {
	return DiscoveryEngineLicenseConfigIdentityFormat.Host()
}

func (i *DiscoveryEngineLicenseConfigIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromDiscoveryEngineLicenseConfigSpec(ctx context.Context, reader client.Reader, obj *DiscoveryEngineLicenseConfig) (*DiscoveryEngineLicenseConfigIdentity, error) {
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

	identity := &DiscoveryEngineLicenseConfigIdentity{
		Project:       projectID,
		Location:      location,
		LicenseConfig: resourceID,
	}
	return identity, nil
}

func (obj *DiscoveryEngineLicenseConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineLicenseConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveryEngineLicenseConfigIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineLicenseConfig identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (obj *DiscoveryEngineLicenseConfig) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
