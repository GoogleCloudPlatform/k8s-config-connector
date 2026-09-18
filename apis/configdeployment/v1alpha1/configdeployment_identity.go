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
	_ identity.IdentityV2 = &ConfigDeploymentIdentity{}
	_ identity.Resource   = &ConfigDeployment{}
)

var ConfigDeploymentIdentityFormat = gcpurls.Template[ConfigDeploymentIdentity]("config.googleapis.com", "projects/{project}/locations/{location}/deployments/{deployment}")

// ConfigDeploymentIdentity is the identity of a GCP ConfigDeployment resource.
// +k8s:deepcopy-gen=false
type ConfigDeploymentIdentity struct {
	Project    string
	Location   string
	Deployment string
}

func (i *ConfigDeploymentIdentity) String() string {
	return ConfigDeploymentIdentityFormat.ToString(*i)
}

func (i *ConfigDeploymentIdentity) FromExternal(ref string) error {
	parsed, match, err := ConfigDeploymentIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ConfigDeployment external=%q was not known (use %s): %w", ref, ConfigDeploymentIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ConfigDeployment external=%q was not known (use %s)", ref, ConfigDeploymentIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ConfigDeploymentIdentity) Host() string {
	return ConfigDeploymentIdentityFormat.Host()
}

func (i *ConfigDeploymentIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromConfigDeploymentSpec(ctx context.Context, reader client.Reader, obj *ConfigDeployment) (*ConfigDeploymentIdentity, error) {
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

	identity := &ConfigDeploymentIdentity{
		Project:    projectID,
		Location:   location,
		Deployment: resourceID,
	}
	return identity, nil
}

func (obj *ConfigDeployment) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromConfigDeploymentSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ConfigDeploymentIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ConfigDeployment identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
