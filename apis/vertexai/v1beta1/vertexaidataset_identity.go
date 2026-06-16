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

package v1beta1

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
	_ identity.IdentityV2 = &VertexAIDatasetIdentity{}
	_ identity.Resource   = &VertexAIDataset{}
)

var VertexAIDatasetIdentityFormat = gcpurls.Template[VertexAIDatasetIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/datasets/{dataset}")

// VertexAIDatasetIdentity is the identity of a Google Cloud VertexAIDataset resource.
// +k8s:deepcopy-gen=false
type VertexAIDatasetIdentity struct {
	Project  string
	Location string
	Dataset  string
}

func (i *VertexAIDatasetIdentity) String() string {
	return VertexAIDatasetIdentityFormat.ToString(*i)
}

func (i *VertexAIDatasetIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAIDatasetIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAIDataset external=%q was not known (use %s): %w", ref, VertexAIDatasetIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAIDataset external=%q was not known (use %s)", ref, VertexAIDatasetIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAIDatasetIdentity) Host() string {
	return VertexAIDatasetIdentityFormat.Host()
}

func (i *VertexAIDatasetIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromVertexAIDatasetSpec(ctx context.Context, reader client.Reader, obj *VertexAIDataset) (*VertexAIDatasetIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := ""
	if obj.Spec.Region != nil {
		location = *obj.Spec.Region
	}
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &VertexAIDatasetIdentity{
		Project:  projectID,
		Location: location,
		Dataset:  resourceID,
	}
	return identity, nil
}

func (obj *VertexAIDataset) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAIDatasetSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	var externalRef string
	if obj.Status.ObservedState != nil {
		externalRef = common.ValueOf(obj.Status.ObservedState.Name)
	}
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAIDatasetIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAIDataset identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
