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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// BackupIdentity defines the resource reference to GKEBackupBackup, which "External" field
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
	BackupPlan string
}

func (p *BackupParent) String() string {
	return p.BackupPlan
}

// New builds a BackupIdentity from the Config Connector Backup object.
func NewBackupIdentity(ctx context.Context, reader client.Reader, obj *GKEBackupBackup) (*BackupIdentity, error) {
	// Get Parent
	backupPlanRef := obj.Spec.BackupPlanRef
	backupPlan, err := backupPlanRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
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
		if actualParent.BackupPlan != backupPlan {
			return nil, fmt.Errorf("spec.backupPlanRef changed, expect %s, got %s", actualParent.BackupPlan, backupPlan)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &BackupIdentity{
		parent: &BackupParent{
			BackupPlan: backupPlan,
		},
		id: resourceID,
	}, nil
}

func ParseBackupExternal(external string) (parent *BackupParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "backupPlans" || tokens[6] != "backups" {
		return nil, "", fmt.Errorf("format of GKEBackupBackup external=%q was not known (use projects/{{projectID}}/locations/{{location}}/backupPlans/{{backupPlanID}}/backups/{{backupID}})", external)
	}
	backupPlan := strings.Join(tokens[:len(tokens)-2], "/")
	parent = &BackupParent{
		BackupPlan: backupPlan,
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
