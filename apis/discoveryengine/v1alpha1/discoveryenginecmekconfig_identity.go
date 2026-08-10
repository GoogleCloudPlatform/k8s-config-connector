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
	_ identity.IdentityV2 = &DiscoveryEngineCMEKConfigIdentity{}
	_ identity.Resource   = &DiscoveryEngineCMEKConfig{}
)

var DiscoveryEngineCMEKConfigIdentityFormat = gcpurls.Template[DiscoveryEngineCMEKConfigIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/cmekConfig")

// DiscoveryEngineCMEKConfigIdentity is the identity of a GCP DiscoveryEngineCMEKConfig resource.
// +k8s:deepcopy-gen=false
type DiscoveryEngineCMEKConfigIdentity struct {
	Project  string
	Location string
}

func (i *DiscoveryEngineCMEKConfigIdentity) String() string {
	return DiscoveryEngineCMEKConfigIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineCMEKConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineCMEKConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineCMEKConfig external=%q was not known (use %s): %w", ref, DiscoveryEngineCMEKConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineCMEKConfig external=%q was not known (use %s)", ref, DiscoveryEngineCMEKConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineCMEKConfigIdentity) Host() string {
	return DiscoveryEngineCMEKConfigIdentityFormat.Host()
}

func (i *DiscoveryEngineCMEKConfigIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromDiscoveryEngineCMEKConfigSpec(ctx context.Context, reader client.Reader, obj *DiscoveryEngineCMEKConfig) (*DiscoveryEngineCMEKConfigIdentity, error) {
	if obj.Spec.Location == nil || *obj.Spec.Location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}
	location := *obj.Spec.Location

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &DiscoveryEngineCMEKConfigIdentity{
		Project:  projectID,
		Location: location,
	}
	return identity, nil
}

func (obj *DiscoveryEngineCMEKConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineCMEKConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status externalRef, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveryEngineCMEKConfigIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineCMEKConfig identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (obj *DiscoveryEngineCMEKConfig) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
