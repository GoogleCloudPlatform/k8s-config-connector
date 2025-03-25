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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// SnapshotIdentity defines the resource reference to PubSubSnapshot, which "External" field
// holds the GCP identifier for the KRM object.
type SnapshotIdentity struct {
	parent *SnapshotParent
	id     string
}

func (i *SnapshotIdentity) String() string {
	return i.parent.String() + "/snapshots/" + i.id
}

func (i *SnapshotIdentity) ID() string {
	return i.id
}

func (i *SnapshotIdentity) Parent() *SnapshotParent {
	return i.parent
}

type SnapshotParent struct {
	ProjectID string
}

func (p *SnapshotParent) String() string {
	return "projects/" + p.ProjectID
}

// New builds a SnapshotIdentity from the Config Connector Snapshot object.
func NewSnapshotIdentity(ctx context.Context, reader client.Reader, obj *PubSubSnapshot) (*SnapshotIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), *obj.Spec.ProjectRef)
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
		actualParent, actualResourceID, err := ParseSnapshotExternal(externalRef)
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
	return &SnapshotIdentity{
		parent: &SnapshotParent{
			ProjectID: projectID,
		},
		id: resourceID,
	}, nil
}

func ParseSnapshotExternal(external string) (parent *SnapshotParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "snapshots" {
		return nil, "", fmt.Errorf("format of PubSubSnapshot external=%q was not known (use projects/{{projectID}}/snapshots/{{snapshotID}})", external)
	}
	parent = &SnapshotParent{
		ProjectID: tokens[1],
	}
	resourceID = tokens[3]
	return parent, resourceID, nil
}
