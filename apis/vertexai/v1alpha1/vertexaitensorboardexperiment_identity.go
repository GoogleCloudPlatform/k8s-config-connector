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
	_ identity.IdentityV2 = &VertexAITensorboardExperimentIdentity{}
	_ identity.Resource   = &VertexAITensorboardExperiment{}
)

var VertexAITensorboardExperimentIdentityFormat = gcpurls.Template[VertexAITensorboardExperimentIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/tensorboards/{tensorboard}/experiments/{experiment}")

// VertexAITensorboardExperimentIdentity is the identity of a GCP VertexAITensorboardExperiment resource.
// +k8s:deepcopy-gen=false
type VertexAITensorboardExperimentIdentity struct {
	Project     string
	Location    string
	Tensorboard string
	Experiment  string
}

func (i *VertexAITensorboardExperimentIdentity) String() string {
	return VertexAITensorboardExperimentIdentityFormat.ToString(*i)
}

func (i *VertexAITensorboardExperimentIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAITensorboardExperimentIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAITensorboardExperiment external=%q was not known (use %s): %w", ref, VertexAITensorboardExperimentIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAITensorboardExperiment external=%q was not known (use %s)", ref, VertexAITensorboardExperimentIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAITensorboardExperimentIdentity) Host() string {
	return VertexAITensorboardExperimentIdentityFormat.Host()
}

func getIdentityFromVertexAITensorboardExperimentSpec(ctx context.Context, reader client.Reader, obj client.Object) (*VertexAITensorboardExperimentIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	vertexaiObj := obj.(*VertexAITensorboardExperiment)
	location := common.ValueOf(vertexaiObj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	tensorboardRef := vertexaiObj.Spec.TensorboardRef
	if tensorboardRef == nil {
		return nil, fmt.Errorf("spec.tensorboardRef must be specified")
	}
	if err := tensorboardRef.Normalize(ctx, reader, vertexaiObj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.tensorboardRef: %w", err)
	}
	tensorboardIdentityRaw, err := tensorboardRef.ParseExternalToIdentity()
	if err != nil {
		return nil, fmt.Errorf("parsing tensorboardRef: %w", err)
	}
	tensorboardIdentity, ok := tensorboardIdentityRaw.(*VertexAITensorboardIdentity)
	if !ok {
		return nil, fmt.Errorf("expected *VertexAITensorboardIdentity from tensorboardRef")
	}
	parentTensorboard := tensorboardIdentity.Tensorboard
	identity := &VertexAITensorboardExperimentIdentity{
		Project:     projectID,
		Location:    location,
		Tensorboard: parentTensorboard,
		Experiment:  resourceID,
	}
	return identity, nil
}

func (obj *VertexAITensorboardExperiment) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAITensorboardExperimentSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAITensorboardExperimentIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAITensorboardExperiment identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
