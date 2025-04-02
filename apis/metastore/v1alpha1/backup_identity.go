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

	// TODO: Add import for the parent service reference if needed, e.g.,
	// metastorev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/metastore/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// BackupIdentity defines the resource reference to MetastoreBackup, which "External" field
// holds the GCP identifier for the KRM object.
type BackupIdentity struct {
	parent *BackupParent
	id     string
}

func (i *BackupIdentity) String() string {
	// URL uses "backups", which matches the proto pattern.
	return i.parent.String() + "/backups/" + i.id
}

func (i *BackupIdentity) ID() string {
	return i.id
}

func (i *BackupIdentity) Parent() *BackupParent {
	return i.parent
}

// BackupParent defines the parent resource hierarchy for a MetastoreBackup.
// Based on the pattern: projects/{project}/locations/{location}/services/{service}
type BackupParent struct {
	ProjectID string
	Location  string
	ServiceID string
}

func (p *BackupParent) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/services/%s", p.ProjectID, p.Location, p.ServiceID)
}

// New builds a BackupIdentity from the Config Connector Backup object.
func NewBackupIdentity(ctx context.Context, reader client.Reader, obj *MetastoreBackup) (*BackupIdentity, error) {

	// Get Parent components
	serviceRef := obj.Spec.MetastoreBackupParent.ServiceRef
	serviceExternalRef, err := serviceRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	serviceParent, serviceID, err := ParseServiceExternal(serviceExternalRef)
	if err != nil {
		return nil, err
	}

	projectID := serviceParent.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve projectID from serviceRef")
	}

	// Assuming Location is directly in Spec as per the original code and backup_types.go.
	// If Location needs to be resolved from the parent Service, this logic would change.
	location := serviceParent.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location from serviceRef")
	}

	if serviceID == "" {
		return nil, fmt.Errorf("cannot determine serviceID from serviceRef")
	}

	// Get desired Backup ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	// Use approved External reference if available
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired fields against the actual external reference
		actualParent, actualResourceID, err := ParseBackupExternal(externalRef)
		if err != nil {
			// If parsing fails, it might indicate an old format or corruption.
			return nil, fmt.Errorf("failed to parse existing externalRef %q: %w", externalRef, err)
		}

		// Check immutability constraints
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s (immutable field)", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s (immutable field)", actualParent.Location, location)
		}
		if actualParent.ServiceID != serviceID {
			// Assuming ServiceID is also immutable or derived from immutable fields.
			return nil, fmt.Errorf("parent service changed, expect %s, got %s (immutable field)", actualParent.ServiceID, serviceID)
		}
		if actualResourceID != resourceID {
			// Resource ID (backup name) might be mutable or immutable depending on GCP API. Assume immutable based on original code.
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %q, since it has already been assigned to %q in externalRef",
				resourceID, actualResourceID)
		}
	}

	// Construct the BackupIdentity
	return &BackupIdentity{
		parent: &BackupParent{
			ProjectID: projectID,
			Location:  location,
			ServiceID: serviceID, // Include the resolved ServiceID
		},
		id: resourceID,
	}, nil
}

// ParseBackupExternal parses the external resource identifier string.
// Expected format: projects/{projectID}/locations/{location}/services/{serviceID}/backups/{backupID}
func ParseBackupExternal(external string) (parent *BackupParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	// Expected format: projects/PROJECT/locations/LOCATION/services/SERVICE/backups/BACKUP
	// Indices:          0       1       2         3          4        5       6        7
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "services" || tokens[6] != "backups" {
		return nil, "", fmt.Errorf("format of MetastoreBackup external=%q was not known (use projects/{{projectID}}/locations/{{location}}/services/{{serviceID}}/backups/{{backupID}})", external)
	}
	parent = &BackupParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
		ServiceID: tokens[5], // Extract ServiceID
	}
	resourceID = tokens[7] // Extract BackupID
	return parent, resourceID, nil
}
