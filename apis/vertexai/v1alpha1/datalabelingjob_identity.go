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
	_ identity.IdentityV2 = &VertexAIDataLabelingJobIdentity{}
	_ identity.Resource   = &VertexAIDataLabelingJob{}
)

var VertexAIDataLabelingJobIdentityFormat = gcpurls.Template[VertexAIDataLabelingJobIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/dataLabelingJobs/{datalabelingjob}")

// +k8s:deepcopy-gen=false
type VertexAIDataLabelingJobIdentity struct {
	Project         string
	Location        string
	DataLabelingJob string
}

func (i *VertexAIDataLabelingJobIdentity) String() string {
	return VertexAIDataLabelingJobIdentityFormat.ToString(*i)
}

func (i *VertexAIDataLabelingJobIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIDataLabelingJobIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIDataLabelingJob external=%q was not known (use %s): %w", ref, VertexAIDataLabelingJobIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIDataLabelingJob external=%q was not known (use %s)", ref, VertexAIDataLabelingJobIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAIDataLabelingJobIdentity) Host() string {
	return VertexAIDataLabelingJobIdentityFormat.Host()
}

func getIdentityFromVertexAIDataLabelingJobSpec(ctx context.Context, reader client.Reader, obj client.Object) (*VertexAIDataLabelingJobIdentity, error) {
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

	identity := &VertexAIDataLabelingJobIdentity{
		Project:         projectID,
		Location:        location,
		DataLabelingJob: resourceID,
	}
	return identity, nil
}

func NewVertexAIDataLabelingJobIdentity(ctx context.Context, reader client.Reader, obj *VertexAIDataLabelingJob) (*VertexAIDataLabelingJobIdentity, error) {
	specIdentity, err := getIdentityFromVertexAIDataLabelingJobSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &VertexAIDataLabelingJobIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.Location != specIdentity.Location {
			return nil, fmt.Errorf("cannot change VertexAIDataLabelingJob location (old=%q, new=%q)", statusIdentity.Location, specIdentity.Location)
		}

		return statusIdentity, nil
	}

	return specIdentity, nil
}

func (obj *VertexAIDataLabelingJob) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return NewVertexAIDataLabelingJobIdentity(ctx, reader, obj)
}
