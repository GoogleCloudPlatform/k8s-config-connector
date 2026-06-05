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
	_ identity.IdentityV2 = &VertexAITrainingPipelineIdentity{}
	_ identity.Resource   = &VertexAITrainingPipeline{}
)

// VertexAITrainingPipelineIdentityFormat is the template format for VertexAITrainingPipeline.
var VertexAITrainingPipelineIdentityFormat = gcpurls.Template[VertexAITrainingPipelineIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/trainingPipelines/{trainingpipeline}")

// VertexAITrainingPipelineIdentity is the identity of a GCP VertexAITrainingPipeline resource.
// +k8s:deepcopy-gen=false
type VertexAITrainingPipelineIdentity struct {
	Project          string
	Location         string
	TrainingPipeline string
}

func (i *VertexAITrainingPipelineIdentity) String() string {
	return VertexAITrainingPipelineIdentityFormat.ToString(*i)
}

func (i *VertexAITrainingPipelineIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAITrainingPipelineIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAITrainingPipeline external=%q was not known (use %s): %w", ref, VertexAITrainingPipelineIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAITrainingPipeline external=%q was not known (use %s)", ref, VertexAITrainingPipelineIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAITrainingPipelineIdentity) Host() string {
	return VertexAITrainingPipelineIdentityFormat.Host()
}

func getIdentityFromVertexAITrainingPipelineSpec(ctx context.Context, reader client.Reader, obj *VertexAITrainingPipeline) (*VertexAITrainingPipelineIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &VertexAITrainingPipelineIdentity{
		Project:          projectID,
		Location:         location,
		TrainingPipeline: resourceID,
	}
	return identity, nil
}

// GetIdentity returns the identity of the VertexAITrainingPipeline resource.
func (obj *VertexAITrainingPipeline) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAITrainingPipelineSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAITrainingPipelineIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAITrainingPipeline identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
