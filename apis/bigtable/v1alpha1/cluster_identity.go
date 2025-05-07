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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"

	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ClusterIdentity defines the resource reference to BigtableCluster, which "External" field
// holds the GCP identifier for the KRM object.
type ClusterIdentity struct {
	parent *bigtablev1beta1.InstanceIdentity
	id     string
}

func (i *ClusterIdentity) String() string {
	return i.ParentString() + "/clusters/" + i.id
}

func (i *ClusterIdentity) ID() string {
	return i.id
}

func (i *ClusterIdentity) ParentString() string {
	return i.parent.String()
}

// New builds a ClusterIdentity from the Config Connector Cluster object.
func NewClusterIdentity(ctx context.Context, reader client.Reader, obj *BigtableCluster) (*ClusterIdentity, error) {
	// Resolve parent
	instanceRef, err := obj.Spec.InstanceRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	instanceParent, instanceID, err := bigtablev1beta1.ParseInstanceExternal(instanceRef)
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
		if actualParent.Parent.ProjectID != instanceParent.ProjectID {
			return nil, fmt.Errorf("spec.instanceRef ProjectID changed, expect %s, got %s", actualParent.Parent.ProjectID, instanceParent.ProjectID)
		}
		if actualParent.Id != instanceID {
			return nil, fmt.Errorf("spec.instanceRef ID changed, expect %s, got %s", actualParent.Id, instanceID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &ClusterIdentity{
		parent: &bigtablev1beta1.InstanceIdentity{
			Parent: &parent.ProjectParent{
				ProjectID: instanceParent.ProjectID},
			Id: instanceID,
		},
		id: resourceID,
	}, nil
}

func ParseClusterExternal(external string) (*bigtablev1beta1.InstanceIdentity, string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "clusters" {
		return nil, "", fmt.Errorf("format of BigtableCluster external=%q was not known (use projects/{{projectID}}/instances/{{instanceID}}/clusters/{{clusterID}})", external)
	}
	p := &bigtablev1beta1.InstanceIdentity{
		Parent: &parent.ProjectParent{ProjectID: tokens[1]},
		Id:     tokens[3],
	}
	resourceID := tokens[5]
	return p, resourceID, nil
}
