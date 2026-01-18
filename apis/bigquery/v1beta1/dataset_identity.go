// Copyright 2024 Google LLC
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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &DatasetIdentity{}

var (
	DatasetIdentityFormat = gcpurls.Template[DatasetIdentity]("bigquery.googleapis.com", "projects/{project}/datasets/{dataset}")
)

// DatasetIdentity defines the resource reference to BigQueryDataset, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type DatasetIdentity struct {
	Project string
	Dataset string
}

func (i *DatasetIdentity) String() string {
	return DatasetIdentityFormat.ToString(*i)
}

func (i *DatasetIdentity) FromExternal(ref string) error {
	parsed, match, err := DatasetIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigQueryDataset external=%q was not known (use %s): %w", ref, DatasetIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigQueryDataset external=%q was not known (use %s)", ref, DatasetIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DatasetIdentity) ID() string {
	return i.Dataset
}

func (i *DatasetIdentity) Parent() *DatasetParent {
	return &DatasetParent{ProjectID: i.Project}
}

type DatasetParent struct {
	ProjectID string
}

func (p *DatasetParent) String() string {
	return "projects/" + p.ProjectID
}

// New builds a DatasetIdentity from the Config Connector Dataset object.
func NewDatasetIdentity(ctx context.Context, reader client.Reader, obj *BigQueryDataset) (*DatasetIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	if projectRef == nil {
		return nil, fmt.Errorf("cannot resolve projectRef: obj.Spec.ProjectRef not defined")
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualIdentity := &DatasetIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actualIdentity.Project != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.Project, projectID)
		}
		if actualIdentity.Dataset != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.Dataset)
		}
	}
	return &DatasetIdentity{
		Project: projectID,
		Dataset: resourceID,
	}, nil
}

// Deprecated: prefer FromExternal
func ParseDatasetExternal(external string) (parent *DatasetParent, resourceID string, err error) {
	id := &DatasetIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, "", err
	}
	return id.Parent(), id.ID(), nil
}
