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

	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// BackupIdentity defines the resource reference to BigtableBackup, which "External" field
// holds the GCP identifier for the KRM object.
type BackupIdentity struct {
	parent *ClusterIdentity
	id     string
}

func (i *BackupIdentity) String() string {
	return i.ParentString() + "/backups/" + i.id
}

func (i *BackupIdentity) ID() string {
	return i.id
}

func (i *BackupIdentity) ParentString() string {
	return i.parent.String()
}

// New builds a BackupIdentity from the Config Connector Backup object.
func NewBackupIdentity(ctx context.Context, reader client.Reader, obj *BigtableBackup) (*BackupIdentity, error) {

	// Get Parent
	clusterRef, err := obj.Spec.ClusterRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	clusterParent, clusterID, err := ParseClusterExternal(clusterRef)
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
		actualParent, actualResourceID, err := ParseBackupExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.parent.Parent.ProjectID != clusterParent.Parent.ProjectID {
			return nil, fmt.Errorf("spec.clusterRef ProjectID changed, expect %s, got %s", actualParent.parent.Parent.ProjectID, clusterParent.Parent.ProjectID)
		}
		if actualParent.parent.Id != clusterParent.Id {
			return nil, fmt.Errorf("spec.clusterRef instanceID changed, expect %s, got %s", actualParent.parent.Id, clusterParent.Id)
		}
		if actualParent.id != clusterID {
			return nil, fmt.Errorf("spec.clusterRef clusterID changed, expect %s, got %s", actualParent.id, clusterID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &BackupIdentity{
		parent: &ClusterIdentity{
			parent: &bigtablev1beta1.InstanceIdentity{
				Parent: &parent.ProjectParent{ProjectID: clusterParent.Parent.ProjectID},
				Id:     clusterParent.Id,
			},
			id: clusterID,
		},
		id: resourceID,
	}, nil
}

func ParseBackupExternal(external string) (*ClusterIdentity, string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "clusters" || tokens[6] != "backups" {
		return nil, "", fmt.Errorf("format of BigtableBackup external=%q was not known (use projects/{{projectID}}/instances/{{instanceID}}/clusters/{{clusterID}}/backups/{{backupID}})", external)
	}
	p := &ClusterIdentity{
		parent: &bigtablev1beta1.InstanceIdentity{
			Parent: &parent.ProjectParent{ProjectID: tokens[1]},
			Id:     tokens[3],
		},
		id: tokens[5],
	}
	resourceID := tokens[7]
	return p, resourceID, nil
}
