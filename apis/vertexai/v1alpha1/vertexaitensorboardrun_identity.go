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
	_ identity.IdentityV2 = &VertexAITensorboardRunIdentity{}
	_ identity.Resource   = &VertexAITensorboardRun{}
)

// VertexAITensorboardRunIdentityFormat is the URL template format for VertexAITensorboardRun.
var VertexAITensorboardRunIdentityFormat = gcpurls.Template[VertexAITensorboardRunIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/tensorboards/{tensorboard}/experiments/{experiment}/runs/{run}")

// VertexAITensorboardRunIdentity is the identity of a GCP VertexAITensorboardRun resource.
// +k8s:deepcopy-gen=false
type VertexAITensorboardRunIdentity struct {
	Project     string
	Location    string
	Tensorboard string
	Experiment  string
	Run         string
}

func (i *VertexAITensorboardRunIdentity) String() string {
	return VertexAITensorboardRunIdentityFormat.ToString(*i)
}

func (i *VertexAITensorboardRunIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAITensorboardRunIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAITensorboardRun external=%q was not known (use %s): %w", ref, VertexAITensorboardRunIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAITensorboardRun external=%q was not known (use %s)", ref, VertexAITensorboardRunIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAITensorboardRunIdentity) Host() string {
	return VertexAITensorboardRunIdentityFormat.Host()
}

func (i *VertexAITensorboardRunIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/tensorboards/%s/experiments/%s", i.Project, i.Location, i.Tensorboard, i.Experiment)
}

func getIdentityFromVertexAITensorboardRunSpec(ctx context.Context, reader client.Reader, obj *VertexAITensorboardRun) (*VertexAITensorboardRunIdentity, error) {
	if obj.Spec.TensorboardRef == nil {
		return nil, fmt.Errorf("spec.tensorboardRef must be set")
	}

	if err := obj.Spec.TensorboardRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}

	tensorboardIdentity, err := obj.Spec.TensorboardRef.ParseExternalToIdentity()
	if err != nil {
		return nil, err
	}
	tbID, ok := tensorboardIdentity.(*VertexAITensorboardIdentity)
	if !ok {
		return nil, fmt.Errorf("resolved tensorboardIdentity is not of type *VertexAITensorboardIdentity")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("spec.location must be set")
	}

	if tbID.Project != projectID {
		return nil, fmt.Errorf("spec.projectRef (%s) must match tensorboardRef project (%s)", projectID, tbID.Project)
	}

	if tbID.Location != location {
		return nil, fmt.Errorf("spec.location (%s) must match tensorboardRef location (%s)", location, tbID.Location)
	}

	experimentID := obj.Spec.TensorboardExperimentID
	if experimentID == "" {
		return nil, fmt.Errorf("spec.tensorboardExperimentID must be set")
	}

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	identity := &VertexAITensorboardRunIdentity{
		Project:     projectID,
		Location:    location,
		Tensorboard: tbID.Tensorboard,
		Experiment:  experimentID,
		Run:         resourceID,
	}
	return identity, nil
}

func (obj *VertexAITensorboardRun) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAITensorboardRunSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAITensorboardRunIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAITensorboardRun identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
