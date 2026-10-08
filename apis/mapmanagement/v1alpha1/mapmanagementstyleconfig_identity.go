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
	_ identity.IdentityV2 = &MapManagementStyleConfigIdentity{}
	_ identity.Resource   = &MapManagementStyleConfig{}
)

var MapManagementStyleConfigIdentityFormat = gcpurls.Template[MapManagementStyleConfigIdentity]("mapmanagement.googleapis.com", "projects/{project}/styleConfigs/{styleConfig}")

// MapManagementStyleConfigIdentity is the identity of a GCP MapManagementStyleConfig resource.
// +k8s:deepcopy-gen=false
type MapManagementStyleConfigIdentity struct {
	Project     string
	StyleConfig string
}

func (i *MapManagementStyleConfigIdentity) String() string {
	return MapManagementStyleConfigIdentityFormat.ToString(*i)
}

func (i *MapManagementStyleConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := MapManagementStyleConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MapManagementStyleConfig external=%q was not known (use %s): %w", ref, MapManagementStyleConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MapManagementStyleConfig external=%q was not known (use %s)", ref, MapManagementStyleConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MapManagementStyleConfigIdentity) Host() string {
	return MapManagementStyleConfigIdentityFormat.Host()
}

func (i *MapManagementStyleConfigIdentity) ParentString() string {
	return "projects/" + i.Project
}

func getIdentityFromMapManagementStyleConfigSpec(ctx context.Context, reader client.Reader, obj *MapManagementStyleConfig) (*MapManagementStyleConfigIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &MapManagementStyleConfigIdentity{
		Project:     projectID,
		StyleConfig: resourceID,
	}
	return identity, nil
}

func (obj *MapManagementStyleConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromMapManagementStyleConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &MapManagementStyleConfigIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if specIdentity.StyleConfig == "" {
			return statusIdentity, nil
		}

		if specIdentity.StyleConfig != statusIdentity.StyleConfig {
			return nil, fmt.Errorf("cannot change MapManagementStyleConfig identity (old=%q, new=%q)", statusIdentity.StyleConfig, specIdentity.StyleConfig)
		}
		return statusIdentity, nil
	}

	return specIdentity, nil
}

// ExternalIdentifier returns the GCP external identifier (the GCP URL).
func (obj *MapManagementStyleConfig) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
