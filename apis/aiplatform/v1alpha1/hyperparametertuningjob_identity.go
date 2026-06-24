// Copyright 2025 Google LLC
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
	_ identity.IdentityV2 = &HyperparameterTuningJobIdentity{}
	_ identity.Resource   = &VertexAIHyperparameterTuningJob{}
)

var HyperparameterTuningJobIdentityFormat = gcpurls.Template[HyperparameterTuningJobIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/hyperparameterTuningJobs/{hyperparametertuningjob}")

// +k8s:deepcopy-gen=false
type HyperparameterTuningJobIdentity struct {
	Project                 string
	Location                string
	HyperparameterTuningJob string
}

func (i *HyperparameterTuningJobIdentity) String() string {
	return HyperparameterTuningJobIdentityFormat.ToString(*i)
}

func (i *HyperparameterTuningJobIdentity) FromExternal(ref string) error {
	parsed, match, err := HyperparameterTuningJobIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIHyperparameterTuningJob external=%q was not known (use %s): %w", ref, HyperparameterTuningJobIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIHyperparameterTuningJob external=%q was not known (use %s)", ref, HyperparameterTuningJobIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *HyperparameterTuningJobIdentity) Host() string {
	return HyperparameterTuningJobIdentityFormat.Host()
}

func getIdentityFromVertexAIHyperparameterTuningJobSpec(ctx context.Context, reader client.Reader, obj client.Object) (*HyperparameterTuningJobIdentity, error) {
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

	identity := &HyperparameterTuningJobIdentity{
		Project:                 projectID,
		Location:                location,
		HyperparameterTuningJob: resourceID,
	}
	return identity, nil
}

func (obj *VertexAIHyperparameterTuningJob) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAIHyperparameterTuningJobSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &HyperparameterTuningJobIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAIHyperparameterTuningJob identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// NewHyperparameterTuningJobIdentity builds a HyperparameterTuningJobIdentity from the object.
func NewHyperparameterTuningJobIdentity(ctx context.Context, reader client.Reader, obj *VertexAIHyperparameterTuningJob) (*HyperparameterTuningJobIdentity, error) {
	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	return id.(*HyperparameterTuningJobIdentity), nil
}

// Parent returns the parent identity.
func (i *HyperparameterTuningJobIdentity) Parent() *HyperparameterTuningJobParent {
	return &HyperparameterTuningJobParent{
		ProjectID: i.Project,
		Location:  i.Location,
	}
}

// ID returns the resource ID.
func (i *HyperparameterTuningJobIdentity) ID() string {
	return i.HyperparameterTuningJob
}

type HyperparameterTuningJobParent struct {
	ProjectID string
	Location  string
}

func (p *HyperparameterTuningJobParent) String() string {
	return fmt.Sprintf("projects/%s/locations/%s", p.ProjectID, p.Location)
}
