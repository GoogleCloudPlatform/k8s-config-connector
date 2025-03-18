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

// BackupIdentity defines the resource reference to BigtableBackup, which "External" field
// holds the GCP identifier for the KRM object.
type BackupIdentity struct {
	parent *BackupParent
	id     string
}

func (i *BackupIdentity) String() string {
	return i.parent.String() + "/backups/" + i.id
}

func (i *BackupIdentity) ID() string {
	return i.id
}

func (i *BackupIdentity) Parent() *BackupParent {
	return i.parent
}

type BackupParent struct {
	ProjectID string
	Instance  string
	Cluster   string
}

func (p *BackupParent) String() string {
	return "projects/" + p.ProjectID + "/instances/" + p.Instance + "/clusters/" + p.Cluster
}

// New builds a BackupIdentity from the Config Connector Backup object.
func NewBackupIdentity(ctx context.Context, reader client.Reader, obj *BigtableBackup) (*BackupIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	instance := obj.Spec.Instance
	cluster := obj.Spec.Cluster

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
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Instance != instance {
			return nil, fmt.Errorf("spec.instance changed, expect %s, got %s", actualParent.Instance, instance)
		}
		if actualParent.Cluster != cluster {
			return nil, fmt.Errorf("spec.cluster changed, expect %s, got %s", actualParent.Cluster, cluster)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &BackupIdentity{
		parent: &BackupParent{
			ProjectID: projectID,
			Instance:  instance,
			Cluster:   cluster,
		},
		id: resourceID,
	}, nil
}

func ParseBackupExternal(external string) (parent *BackupParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "clusters" || tokens[6] != "backups" {
		return nil, "", fmt.Errorf("format of BigtableBackup external=%q was not known (use projects/{{projectID}}/instances/{{instanceID}}/clusters/{{clusterID}}/backups/{{backupID}})", external)
	}
	parent = &BackupParent{
		ProjectID: tokens[1],
		Instance:  tokens[3],
		Cluster:   tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
