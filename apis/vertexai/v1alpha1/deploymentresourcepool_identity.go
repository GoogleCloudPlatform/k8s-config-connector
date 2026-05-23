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
	_ identity.IdentityV2 = &VertexAIDeploymentResourcePoolIdentity{}
	_ identity.Resource   = &VertexAIDeploymentResourcePool{}
)

var VertexAIDeploymentResourcePoolIdentityFormat = gcpurls.Template[VertexAIDeploymentResourcePoolIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/deploymentResourcePools/{deploymentresourcepool}")

// +k8s:deepcopy-gen=false
type VertexAIDeploymentResourcePoolIdentity struct {
	Project                string
	Location               string
	DeploymentResourcePool string
}

func (i *VertexAIDeploymentResourcePoolIdentity) String() string {
	return VertexAIDeploymentResourcePoolIdentityFormat.ToString(*i)
}

func (i *VertexAIDeploymentResourcePoolIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIDeploymentResourcePoolIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIDeploymentResourcePool external=%q was not known (use %s): %w", ref, VertexAIDeploymentResourcePoolIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIDeploymentResourcePool external=%q was not known (use %s)", ref, VertexAIDeploymentResourcePoolIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAIDeploymentResourcePoolIdentity) Host() string {
	return VertexAIDeploymentResourcePoolIdentityFormat.Host()
}

func NewVertexAIDeploymentResourcePoolIdentity(ctx context.Context, reader client.Reader, obj *VertexAIDeploymentResourcePool) (*VertexAIDeploymentResourcePoolIdentity, error) {
	return getIdentityFromVertexAIDeploymentResourcePoolSpec(ctx, reader, obj)
}

func getIdentityFromVertexAIDeploymentResourcePoolSpec(ctx context.Context, reader client.Reader, obj client.Object) (*VertexAIDeploymentResourcePoolIdentity, error) {
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

	identity := &VertexAIDeploymentResourcePoolIdentity{
		Project:                projectID,
		Location:               location,
		DeploymentResourcePool: resourceID,
	}
	return identity, nil
}

func (obj *VertexAIDeploymentResourcePool) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAIDeploymentResourcePoolSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAIDeploymentResourcePoolIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAIDeploymentResourcePool identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
