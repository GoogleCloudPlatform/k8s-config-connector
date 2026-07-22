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
	_ identity.IdentityV2 = &ConfigDeploymentGroupIdentity{}
	_ identity.Resource   = &ConfigDeploymentGroup{}
)

var ConfigDeploymentGroupIdentityFormat = gcpurls.Template[ConfigDeploymentGroupIdentity]("config.googleapis.com", "projects/{project}/locations/{location}/deploymentGroups/{deploymentGroup}")

// ConfigDeploymentGroupIdentity is the identity of a GCP ConfigDeploymentGroup resource.
// +k8s:deepcopy-gen=false
type ConfigDeploymentGroupIdentity struct {
	Project         string
	Location        string
	DeploymentGroup string
}

func (i *ConfigDeploymentGroupIdentity) String() string {
	return ConfigDeploymentGroupIdentityFormat.ToString(*i)
}

func (i *ConfigDeploymentGroupIdentity) FromExternal(ref string) error {
	parsed, match, err := ConfigDeploymentGroupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ConfigDeploymentGroup external=%q was not known (use %s): %w", ref, ConfigDeploymentGroupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ConfigDeploymentGroup external=%q was not known (use %s)", ref, ConfigDeploymentGroupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ConfigDeploymentGroupIdentity) Host() string {
	return ConfigDeploymentGroupIdentityFormat.Host()
}

func (i *ConfigDeploymentGroupIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromConfigDeploymentGroupSpec(ctx context.Context, reader client.Reader, obj client.Object) (*ConfigDeploymentGroupIdentity, error) {
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

	identity := &ConfigDeploymentGroupIdentity{
		Project:         projectID,
		Location:        location,
		DeploymentGroup: resourceID,
	}
	return identity, nil
}

func (obj *ConfigDeploymentGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromConfigDeploymentGroupSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ConfigDeploymentGroupIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ConfigDeploymentGroup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
