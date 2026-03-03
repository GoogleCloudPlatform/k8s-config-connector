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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
)

// BackupIdentity defines the resource reference to MemorystoreInstanceBackup, which "External" field
// holds the GCP identifier for the KRM object.
type BackupIdentity struct {
	parent string
	id     string
}

func (i *BackupIdentity) String() string {
	return i.parent + "/backups/" + i.id
}

func (i *BackupIdentity) ID() string {
	return i.id
}

func (i *BackupIdentity) Parent() string {
	return i.parent
}

// New builds a BackupIdentity from the Config Connector Backup object.
func NewBackupIdentity(ctx context.Context, obj *MemorystoreInstanceBackup) (*BackupIdentity, error) {
	parent := common.ValueOf(obj.Spec.BackupCollection)
	if parent == "" {
		return nil, fmt.Errorf("cannot resolve parent")
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
		if resourceID == "" {
			return nil, fmt.Errorf("cannot resolve resource ID")
		}
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseBackupExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent != parent {
			return nil, fmt.Errorf("spec.backupCollection changed, expect %s, got %s", actualParent, parent)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &BackupIdentity{
		parent: parent,
		id:     resourceID,
	}, nil
}

func ParseBackupExternal(external string) (string, string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "backupCollections" || tokens[6] != "backups" {
		return "", "", fmt.Errorf("format of MemorystoreInstanceBackup external=%q was not known (use projects/{{projectID}}/locations/{{location}}/backupCollections/{{backupCollectionName}}/backups/{{backupID}})", external)
	}
	return fmt.Sprintf("projects/%s/locations/%s/backupCollections/%s", tokens[1], tokens[3], tokens[5]), tokens[7], nil
}
