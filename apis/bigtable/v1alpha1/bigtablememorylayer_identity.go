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
	"strings"

	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MemoryLayerIdentity is the identity of a BigtableMemoryLayer.
//
// A memory layer is a singleton of its cluster, so the identity carries no ID
// of its own; the parent cluster provides uniqueness.
type MemoryLayerIdentity struct {
	parent *ClusterIdentity
}

func (i *MemoryLayerIdentity) String() string {
	return i.ParentString() + "/memoryLayer"
}

func (i *MemoryLayerIdentity) Parent() *ClusterIdentity {
	return i.parent
}

func (i *MemoryLayerIdentity) ParentString() string {
	return i.parent.String()
}

// NewMemoryLayerIdentity builds a MemoryLayerIdentity from the Config Connector
// BigtableMemoryLayer object.
func NewMemoryLayerIdentity(ctx context.Context, reader client.Reader, obj *BigtableMemoryLayer) (*MemoryLayerIdentity, error) {
	// Get Parent
	if obj.Spec.ClusterRef == nil {
		return nil, fmt.Errorf("spec.clusterRef is required")
	}
	clusterRef, err := obj.Spec.ClusterRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	instanceParent, clusterID, err := ParseClusterExternal(clusterRef)
	if err != nil {
		return nil, err
	}

	id := &MemoryLayerIdentity{
		parent: &ClusterIdentity{
			parent: instanceParent,
			id:     clusterID,
		},
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actual, err := ParseMemoryLayerExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actual.ParentString() != id.ParentString() {
			return nil, fmt.Errorf("cannot reset `spec.clusterRef` to %s, since it has already assigned to %s",
				id.ParentString(), actual.ParentString())
		}
	}
	return id, nil
}

func ParseMemoryLayerExternal(external string) (*MemoryLayerIdentity, error) {
	tokens := strings.Split(strings.TrimPrefix(external, "/"), "/")
	if len(tokens) != 7 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "clusters" || tokens[6] != "memoryLayer" {
		return nil, fmt.Errorf("format of BigtableMemoryLayer external=%q was not known (use projects/{{projectID}}/instances/{{instanceID}}/clusters/{{clusterID}}/memoryLayer)", external)
	}
	return &MemoryLayerIdentity{
		parent: &ClusterIdentity{
			parent: &bigtablev1beta1.InstanceIdentity{
				Parent: &parent.ProjectParent{ProjectID: tokens[1]},
				Id:     tokens[3],
			},
			id: tokens[5],
		},
	}, nil
}
