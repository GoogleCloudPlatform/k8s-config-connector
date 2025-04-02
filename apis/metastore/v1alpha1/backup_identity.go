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
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	// Assuming Location is directly in Spec as per the original code and backup_types.go.
	// If Location needs to be resolved from the parent Service, this logic would change.
	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("spec.location must be set")
	}

	// TODO: Resolve ServiceID. The current MetastoreBackupSpec in backup_types.go doesn't explicitly contain a ServiceRef.
	// This resolution logic depends on how the parent service is specified in the MetastoreBackup CRD.
	// Placeholder: This needs to be adapted based on the actual spec design.
	var serviceID string
	// Example: If Service ID comes from a field like obj.Spec.ServiceRef
	// serviceRef, err := resolveMetastoreServiceRef(ctx, reader, obj, obj.Spec.ServiceRef)
	// if err != nil {
	// 	 return nil, fmt.Errorf("cannot resolve service: %w", err)
	// }
	// serviceID = serviceRef.ServiceID // Extract the ID

	// Example: If Service ID comes from the spec.name field (less ideal)
	// if obj.Spec.Name != nil {
	// 	 nameParts := strings.Split(*obj.Spec.Name, "/")
	// 	 if len(nameParts) == 8 && nameParts[0] == "projects" && nameParts[2] == "locations" && nameParts[4] == "services" && nameParts[6] == "backups" {
	// 		 serviceID = nameParts[5]
	// 	 }
	// }

	if serviceID == "" {
		// If ServiceID cannot be determined from the spec, this will fail.
		// This indicates a potential mismatch between the identity logic and the CRD definition.
		// Or, the ServiceID might need to be retrieved differently (e.g., from parent Service status).
		return nil, fmt.Errorf("cannot determine service ID from MetastoreBackup spec - resolution logic needs implementation based on CRD design")
	}

	// Get desired Backup ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID (metadata.name or spec.resourceID)")
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
