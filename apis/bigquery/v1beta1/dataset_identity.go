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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DatasetIdentity defines the resource reference to BigQueryDataset, which "External" field
// holds the GCP identifier for the KRM object.
type DatasetIdentity struct {
	parent *DatasetParent
	id     string
}

func (i *DatasetIdentity) String() string {
	return i.parent.String() + "/datasets/" + i.id
}

func (i *DatasetIdentity) ID() string {
	return i.id
}

func (i *DatasetIdentity) Parent() *DatasetParent {
	return i.parent
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
		actualParent, actualResourceID, err := ParseDatasetExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &DatasetIdentity{
		parent: &DatasetParent{
			ProjectID: projectID,
		},
		id: resourceID,
	}, nil
}

func ParseDatasetExternal(external string) (parent *DatasetParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")

	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "datasets" {
		return nil, "", fmt.Errorf("format of BigQueryDataset external=%q was not known (use projects/<projectId>/datasets/<datasetID>)", external)
	}
	parent = &DatasetParent{
		ProjectID: tokens[1],
	}

	resourceID = tokens[3]

	return parent, resourceID, nil
}
