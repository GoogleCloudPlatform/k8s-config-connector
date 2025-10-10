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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MaterializedViewIdentity defines the resource reference to BigtableMaterializedView, which "External" field
// holds the GCP identifier for the KRM object.
type MaterializedViewIdentity struct {
	parent *v1beta1.InstanceIdentity
	id     string
}

func (i *MaterializedViewIdentity) String() string {
	return i.ParentString() + "/materializedViews/" + i.id
}

func (i *MaterializedViewIdentity) ID() string {
	return i.id
}

func (i *MaterializedViewIdentity) Parent() *v1beta1.InstanceIdentity {
	return i.parent
}

func (i *MaterializedViewIdentity) ParentString() string {
	return i.parent.String()
}

func (i *MaterializedViewIdentity) ParentInstanceIdString() string {
	return i.parent.Id
}

// New builds a MaterializedViewIdentity from the Config Connector MaterializedView object.
func NewMaterializedViewIdentity(ctx context.Context, reader client.Reader, obj *BigtableMaterializedView) (*MaterializedViewIdentity, error) {

	// Get Parent
	instanceRef, err := obj.Spec.InstanceRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	instanceParent, instanceID, err := v1beta1.ParseInstanceExternal(instanceRef)
	if err != nil {
		return nil, err
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve materialized view name")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseMaterializedViewExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.Parent.ProjectID != instanceParent.ProjectID {
			return nil, fmt.Errorf("ProjectID changed, expect %s, got %s", actualParent.Parent.ProjectID, instanceParent.ProjectID)
		}
		if actualParent.Id != instanceID {
			return nil, fmt.Errorf("InstanceID changed, expect %s, got %s", actualParent.Id, instanceID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &MaterializedViewIdentity{
		parent: &v1beta1.InstanceIdentity{
			Parent: instanceParent,
			Id:     instanceID,
		},
		id: resourceID,
	}, nil
}

func ParseMaterializedViewExternal(external string) (*v1beta1.InstanceIdentity, string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "materializedViews" {
		return nil, "", fmt.Errorf("format of BigtableMaterializedView external=%q was not known (use projects/{{projectID}}/instances/{{instance}}/materializedViews/{{materializedViewID}})", external)
	}
	return &v1beta1.InstanceIdentity{Parent: &parent.ProjectParent{ProjectID: tokens[1]}, Id: tokens[3]}, tokens[5], nil
}
