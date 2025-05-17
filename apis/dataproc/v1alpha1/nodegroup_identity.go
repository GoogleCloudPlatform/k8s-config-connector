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

// NodeGroupIdentity defines the resource reference to DataprocNodeGroup, which "External" field
// holds the GCP identifier for the KRM object.
type NodeGroupIdentity struct {
	parent *NodeGroupParent
	id     string
}

func (i *NodeGroupIdentity) String() string {
	return i.parent.String() + "/nodeGroups/" + i.id
}

func (i *NodeGroupIdentity) ID() string {
	return i.id
}

func (i *NodeGroupIdentity) Parent() *NodeGroupParent {
	return i.parent
}

type NodeGroupParent struct {
	ProjectID string
	Location  string
	Cluster   string
}

func (p *NodeGroupParent) String() string {
	return "projects/" + p.ProjectID + "/regions/" + p.Location + "/clusters/" + p.Cluster
}

// New builds a NodeGroupIdentity from the Config Connector NodeGroup object.
func NewNodeGroupIdentity(ctx context.Context, reader client.Reader, obj *DataprocNodeGroup) (*NodeGroupIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := obj.Spec.Location
	cluster := obj.Spec.ClusterRef.Name

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
		actualParent, actualResourceID, err := ParseNodeGroupExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualParent.Cluster != cluster {
			return nil, fmt.Errorf("spec.clusterRef changed, expect %s, got %s", actualParent.Cluster, cluster)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &NodeGroupIdentity{
		parent: &NodeGroupParent{
			ProjectID: projectID,
			Location:  location,
			Cluster:   cluster,
		},
		id: resourceID,
	}, nil
}

func ParseNodeGroupExternal(external string) (parent *NodeGroupParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "regions" || tokens[4] != "clusters" || tokens[6] != "nodeGroups" {
		return nil, "", fmt.Errorf("format of DataprocNodeGroup external=%q was not known (use projects/{{projectID}}/regions/{{region}}/clusters/{{cluster}}/nodeGroups/{{nodegroupID}})", external)
	}
	parent = &NodeGroupParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
		Cluster:   tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
