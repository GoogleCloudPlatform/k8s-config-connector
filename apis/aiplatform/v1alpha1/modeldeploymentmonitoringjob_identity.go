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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &VertexAIModelDeploymentMonitoringJobIdentity{}
	_ identity.Resource   = &VertexAIModelDeploymentMonitoringJob{}
)

var VertexAIModelDeploymentMonitoringJobIdentityFormat = gcpurls.Template[VertexAIModelDeploymentMonitoringJobIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{modelDeploymentMonitoringJob}")

// +k8s:deepcopy-gen=false
type VertexAIModelDeploymentMonitoringJobIdentity struct {
	Project                      string
	Location                     string
	ModelDeploymentMonitoringJob string
}

func (i *VertexAIModelDeploymentMonitoringJobIdentity) String() string {
	return VertexAIModelDeploymentMonitoringJobIdentityFormat.ToString(*i)
}

func (i *VertexAIModelDeploymentMonitoringJobIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIModelDeploymentMonitoringJobIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIModelDeploymentMonitoringJob external=%q was not known (use %s): %w", ref, VertexAIModelDeploymentMonitoringJobIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIModelDeploymentMonitoringJob external=%q was not known (use %s)", ref, VertexAIModelDeploymentMonitoringJobIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAIModelDeploymentMonitoringJobIdentity) Host() string {
	return VertexAIModelDeploymentMonitoringJobIdentityFormat.Host()
}

func ParseVertexAIModelDeploymentMonitoringJobIdentity(external string) (*VertexAIModelDeploymentMonitoringJobIdentity, error) {
	id := &VertexAIModelDeploymentMonitoringJobIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func (i *VertexAIModelDeploymentMonitoringJobIdentity) ID() string {
	return i.ModelDeploymentMonitoringJob
}

func getIdentityFromVertexAIModelDeploymentMonitoringJobSpec(ctx context.Context, reader client.Reader, obj client.Object) (*VertexAIModelDeploymentMonitoringJobIdentity, error) {
	resourceID, err := refsv1beta1.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refsv1beta1.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &VertexAIModelDeploymentMonitoringJobIdentity{
		Project:                      projectID,
		Location:                     location,
		ModelDeploymentMonitoringJob: resourceID,
	}
	return identity, nil
}

func (obj *VertexAIModelDeploymentMonitoringJob) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAIModelDeploymentMonitoringJobSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAIModelDeploymentMonitoringJobIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAIModelDeploymentMonitoringJob identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *VertexAIModelDeploymentMonitoringJob) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
