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
	_ identity.IdentityV2 = &MapManagementMapConfigIdentity{}
	_ identity.Resource   = &MapManagementMapConfig{}
)

var MapManagementMapConfigIdentityFormat = gcpurls.Template[MapManagementMapConfigIdentity]("mapmanagement.googleapis.com", "projects/{project}/mapConfigs/{mapConfig}")

// +k8s:deepcopy-gen=false
type MapManagementMapConfigIdentity struct {
	Project   string
	MapConfig string
}

func (i *MapManagementMapConfigIdentity) String() string {
	return MapManagementMapConfigIdentityFormat.ToString(*i)
}

func (i *MapManagementMapConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := MapManagementMapConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MapManagementMapConfig external=%q was not known (use %s): %w", ref, MapManagementMapConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MapManagementMapConfig external=%q was not known (use %s)", ref, MapManagementMapConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MapManagementMapConfigIdentity) Host() string {
	return MapManagementMapConfigIdentityFormat.Host()
}

func getIdentityFromMapManagementMapConfigSpec(ctx context.Context, reader client.Reader, obj *MapManagementMapConfig) (*MapManagementMapConfigIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &MapManagementMapConfigIdentity{
		Project:   projectID,
		MapConfig: resourceID,
	}
	return identity, nil
}

func (obj *MapManagementMapConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromMapManagementMapConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &MapManagementMapConfigIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change MapManagementMapConfig identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier returns the GCP external identifier (the GCP URL).
func (obj *MapManagementMapConfig) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
