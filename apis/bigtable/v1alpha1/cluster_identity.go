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

// ClusterIdentity defines the resource reference to BigtableCluster, which "External" field
// holds the GCP identifier for the KRM object.
type ClusterIdentity struct {
	parent *ClusterParent
	id     string
}

func (i *ClusterIdentity) String() string {
	return i.parent.String() + "/clusters/" + i.id
}

func (i *ClusterIdentity) ID() string {
	return i.id
}

func (i *ClusterIdentity) Parent() *ClusterParent {
	return i.parent
}

type ClusterParent struct {
	ProjectID   string
	InstanceRef string
}

func (p *ClusterParent) String() string {
	return "projects/" + p.ProjectID + "/instances/" + p.InstanceRef
}

// New builds a ClusterIdentity from the Config Connector Cluster object.
func NewClusterIdentity(ctx context.Context, reader client.Reader, obj *BigtableCluster) (*ClusterIdentity, error) {
	// Resolve the ClusterParent fields from the Parent fields.
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	instanceRef, err := obj.Spec.InstanceRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
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
		actualParent, actualResourceID, err := ParseClusterExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.InstanceRef != instanceRef {
			return nil, fmt.Errorf("spec.instanceRef changed, expect %s, got %s", actualParent.InstanceRef, instanceRef)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &ClusterIdentity{
		parent: &ClusterParent{
			ProjectID:   projectID,
			InstanceRef: instanceRef,
		},
		id: resourceID,
	}, nil
}

func ParseClusterExternal(external string) (parent *ClusterParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "clusters" {
		return nil, "", fmt.Errorf("format of BigtableCluster external=%q was not known (use projects/{{projectID}}/instances/{{instanceID}}/clusters/{{clusterID}})", external)
	}
	parent = &ClusterParent{
		ProjectID:   tokens[1],
		InstanceRef: tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
