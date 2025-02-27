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

// BackupScheduleIdentity defines the resource reference to SpannerBackupSchedule, which "External" field
// holds the GCP identifier for the KRM object.
type BackupScheduleIdentity struct {
	parent *BackupScheduleParent
	id     string
}

func (i *BackupScheduleIdentity) String() string {
	return i.parent.String() + "/backupSchedules/" + i.id
}

func (i *BackupScheduleIdentity) ID() string {
	return i.id
}

func (i *BackupScheduleIdentity) Parent() *BackupScheduleParent {
	return i.parent
}

type BackupScheduleParent struct {
	ProjectID  string
	InstanceID string
	DatabaseID string
}

func (p *BackupScheduleParent) String() string {
	return "projects/" + p.ProjectID + "/instances/" + p.InstanceID + "/databases/" + p.DatabaseID
}

// New builds a BackupScheduleIdentity from the Config Connector BackupSchedule object.
func NewBackupScheduleIdentity(ctx context.Context, reader client.Reader, obj *SpannerBackupSchedule) (*BackupScheduleIdentity, error) {

	// Get Parent
	databaseRef, err := refsv1beta1.ResolveSpannerDatabaseRef(ctx, reader, obj, obj.Spec.DatabaseRef)
	if err != nil {
		return nil, err
	}

	databaseID := databaseRef.DatabaseID
	projectID := databaseRef.ProjectID
	instanceID := databaseRef.InstanceID

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
		actualParent, actualResourceID, err := ParseBackupScheduleExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.InstanceID != instanceID {
			return nil, fmt.Errorf("spec.instanceRef changed, expect %s, got %s", actualParent.InstanceID, instanceID)
		}
		if actualParent.DatabaseID != databaseID {
			return nil, fmt.Errorf("spec.databaseRef changed, expect %s, got %s", actualParent.DatabaseID, databaseID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &BackupScheduleIdentity{
		parent: &BackupScheduleParent{
			ProjectID:  projectID,
			InstanceID: instanceID,
			DatabaseID: databaseID,
		},
		id: resourceID,
	}, nil
}

func ParseBackupScheduleExternal(external string) (parent *BackupScheduleParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "databases" || tokens[6] != "backupSchedules" {
		return nil, "", fmt.Errorf("format of SpannerBackupSchedule external=%q was not known (use projects/{{projectID}}/instances/{{instance}}/databases/{{database}}/backupSchedules/{{backupscheduleID}})", external)
	}
	parent = &BackupScheduleParent{
		ProjectID:  tokens[1],
		InstanceID: tokens[3],
		DatabaseID: tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
