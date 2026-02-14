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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
)

// BackupIdentity defines the resource reference to MemorystoreInstanceBackup, which "External" field
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
	ProjectID        string
	Location         string
	BackupCollection string
}

func (p *BackupParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/backupCollections/" + p.BackupCollection
}

// New builds a BackupIdentity from the Config Connector Backup object.
func NewBackupIdentity(ctx context.Context, obj *MemorystoreInstanceBackup) (*BackupIdentity, error) {
	backupCollectionExternal := ""
	if obj.Spec.BackupCollectionExternal != nil {
		backupCollectionExternal = *(obj.Spec.BackupCollectionExternal)
	}
	if backupCollectionExternal == "" {
		return nil, fmt.Errorf("cannot resolve backup collection external")
	}
	parent, err := ParseBackupParentExternal(backupCollectionExternal)
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
		if actualParent.String() != parent.String() {
			return nil, fmt.Errorf("spec.BackupCollectionExternal changed, expect %s, got %s", actualParent.String(), parent.String())
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return nil, nil
}

func ParseBackupParentExternal(external string) (*BackupParent, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "backupCollections" {
		return nil, fmt.Errorf("format of BackupCollection external=%q was not known (use projects/{{projectID}}/locations/{{location}}/backupCollections/{{backupCollectionName}})", external)
	}
	return &BackupParent{
		ProjectID:        tokens[1],
		Location:         tokens[3],
		BackupCollection: tokens[5],
	}, nil
}

func ParseBackupExternal(external string) (*BackupParent, string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "backupCollections" || tokens[6] != "backups" {
		return nil, "", fmt.Errorf("format of MemorystoreInstanceBackup external=%q was not known (use projects/{{projectID}}/locations/{{location}}/backupCollections/{{backupCollectionName}}/backups/{{backupID}})", external)
	}
	return &BackupParent{
		ProjectID:        tokens[1],
		Location:         tokens[3],
		BackupCollection: tokens[5],
	}, tokens[7], nil
}
