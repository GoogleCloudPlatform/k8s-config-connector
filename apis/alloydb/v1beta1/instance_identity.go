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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// InstanceIdentity defines the resource reference to AlloyDBInstance.
type InstanceIdentity struct {
	parent *InstanceParent
	id     string
}

func (i *InstanceIdentity) String() string {
	return i.parent.String() + "/instances/" + i.id
}

func (i *InstanceIdentity) ID() string {
	return i.id
}

func (i *InstanceIdentity) Parent() *InstanceParent {
	return i.parent
}

type InstanceParent struct {
	projectID string
	location  string
	clusterID string
}

func (p *InstanceParent) String() string {
	return "projects/" + p.projectID + "/locations/" + p.location + "/clusters/" + p.clusterID
}

// New builds a InstanceIdentity from the Config Connector Instance object.
func NewInstanceIdentity(ctx context.Context, reader client.Reader, obj *AlloyDBInstance) (*InstanceIdentity, error) {

	// Get Parent
	clusterRef, err := refsv1beta1.ResolveAlloyDBCluster(ctx, reader, obj, obj.Spec.ClusterRef)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve AlloyDBCluster ref: %w", err)
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved ExternalRef
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseInstanceExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.String() != clusterRef.String() {
			return nil, fmt.Errorf("spec.clusterRef changed, expect %s, got %s", actualParent.String(), clusterRef)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &InstanceIdentity{
		parent: &InstanceParent{
			projectID: clusterRef.ProjectID,
			location:  clusterRef.Location,
			clusterID: clusterRef.ClusterID,
		},
		id: resourceID,
	}, nil
}
