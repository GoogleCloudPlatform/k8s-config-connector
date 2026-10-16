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

	backupdrv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BackupIdentity{}
	_ identity.Resource   = &BackupDRBackup{}
)

var BackupIdentityFormat = gcpurls.Template[BackupIdentity]("backupdr.googleapis.com", "projects/{project}/locations/{location}/backupVaults/{backupvault}/dataSources/{datasource}/backups/{backup}")

// BackupIdentity is the identity of a GCP BackupDRBackup resource.
// +k8s:deepcopy-gen=false
type BackupIdentity struct {
	Project     string
	Location    string
	BackupVault string
	DataSource  string
	Backup      string
}

func (i *BackupIdentity) String() string {
	return BackupIdentityFormat.ToString(*i)
}

func (i *BackupIdentity) FromExternal(ref string) error {
	parsed, match, err := BackupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BackupDRBackup external=%q was not known (use %s): %w", ref, BackupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BackupDRBackup external=%q was not known (use %s)", ref, BackupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BackupIdentity) Host() string {
	return BackupIdentityFormat.Host()
}

func (i *BackupIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/dataSources/%s", i.Project, i.Location, i.BackupVault, i.DataSource)
}

func getIdentityFromBackupSpec(ctx context.Context, reader client.Reader, obj *BackupDRBackup) (*BackupIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	location := common.ValueOf(obj.Spec.Location)

	var backupVault string
	if obj.Spec.BackupVaultRef == nil {
		return nil, fmt.Errorf("spec.backupVaultRef must be set")
	}
	if obj.Spec.BackupVaultRef.External != "" {
		parsed, match, err := backupdrv1beta1.BackupVaultIdentityFormat.Parse(obj.Spec.BackupVaultRef.External)
		if err == nil && match {
			backupVault = parsed.BackupVault
			if projectID == "" {
				projectID = parsed.Project
			} else if parsed.Project != "" && projectID != parsed.Project {
				return nil, fmt.Errorf("project %q does not match backupVaultRef project %q", projectID, parsed.Project)
			}
			if location == "" {
				location = parsed.Location
			} else if parsed.Location != "" && location != parsed.Location {
				return nil, fmt.Errorf("location %q does not match backupVaultRef location %q", location, parsed.Location)
			}
		} else if refs.IsShortName(obj.Spec.BackupVaultRef.External) {
			backupVault = obj.Spec.BackupVaultRef.External
		} else {
			return nil, fmt.Errorf("cannot parse backupVaultRef external=%q: %w", obj.Spec.BackupVaultRef.External, err)
		}
	} else if obj.Spec.BackupVaultRef.Name != "" {
		if reader != nil {
			if err := obj.Spec.BackupVaultRef.Normalize(ctx, reader, obj.GetNamespace()); err == nil && obj.Spec.BackupVaultRef.External != "" {
				parsed, match, err := backupdrv1beta1.BackupVaultIdentityFormat.Parse(obj.Spec.BackupVaultRef.External)
				if err == nil && match {
					backupVault = parsed.BackupVault
					if projectID == "" {
						projectID = parsed.Project
					}
					if location == "" {
						location = parsed.Location
					}
				}
			}
		}
		if backupVault == "" {
			backupVault = obj.Spec.BackupVaultRef.Name
		}
	}

	var dataSource string
	if obj.Spec.DataSourceRef == nil {
		return nil, fmt.Errorf("spec.dataSourceRef must be set")
	}
	if obj.Spec.DataSourceRef.External == "" {
		return nil, fmt.Errorf("spec.dataSourceRef.external must be set")
	}
	parsedDS, matchDS, errDS := DataSourceIdentityFormat.Parse(obj.Spec.DataSourceRef.External)
	if errDS == nil && matchDS {
		dataSource = parsedDS.DataSource
		if projectID == "" {
			projectID = parsedDS.Project
		} else if parsedDS.Project != "" && projectID != parsedDS.Project {
			return nil, fmt.Errorf("project %q does not match dataSourceRef project %q", projectID, parsedDS.Project)
		}
		if location == "" {
			location = parsedDS.Location
		} else if parsedDS.Location != "" && location != parsedDS.Location {
			return nil, fmt.Errorf("location %q does not match dataSourceRef location %q", location, parsedDS.Location)
		}
		if backupVault == "" {
			backupVault = parsedDS.BackupVault
		} else if parsedDS.BackupVault != "" && backupVault != parsedDS.BackupVault {
			return nil, fmt.Errorf("backupVault %q does not match dataSourceRef backupVault %q", backupVault, parsedDS.BackupVault)
		}
	} else if refs.IsShortName(obj.Spec.DataSourceRef.External) {
		dataSource = obj.Spec.DataSourceRef.External
	} else {
		return nil, fmt.Errorf("cannot parse dataSourceRef external=%q: %w", obj.Spec.DataSourceRef.External, errDS)
	}

	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project: please set spec.projectRef or ensure parent references specify project")
	}
	if location == "" {
		return nil, fmt.Errorf("spec.location must be set")
	}
	if backupVault == "" {
		return nil, fmt.Errorf("spec.backupVaultRef must be set")
	}
	if dataSource == "" {
		return nil, fmt.Errorf("spec.dataSourceRef must be set")
	}

	identity := &BackupIdentity{
		Project:     projectID,
		Location:    location,
		BackupVault: backupVault,
		DataSource:  dataSource,
		Backup:      resourceID,
	}
	return identity, nil
}

func (obj *BackupDRBackup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBackupSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &BackupIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BackupDRBackup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (obj *BackupDRBackup) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
